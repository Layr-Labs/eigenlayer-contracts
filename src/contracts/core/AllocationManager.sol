// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "../libraries/SlashingLib.sol";
import "./AllocationManagerStorage.sol";

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable
{
    using Snapshots for Snapshots.DefaultWadHistory;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using SlashingLib for uint256;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManage,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) AllocationManagerStorage(_delegation, _avsDirectory, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY) {
        _disableInitializers();
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Called by the delagation manager to set delay when operators register.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     * @dev msg.sender is assumed to be the delegation manager.
     */
    function setAllocationDelay(address operator, uint32 delay) external {
        require(msg.sender == address(delegation), OnlyDelegationManager());
        _setAllocationDelay(operator, delay);
    }

    /**
     * @notice Called by operators to set their allocation delay.
     * @param delay the allocation delay in seconds
     * @dev msg.sender is assumed to be the operator
     */
    function setAllocationDelay(
        uint32 delay
    ) external {
        require(delegation.isOperator(msg.sender), OperatorNotRegistered());
        _setAllocationDelay(msg.sender, delay);
    }

    /**
     * @notice This function takes a list of strategies and adds all completable modifications for each strategy,
     * updating the encumberedMagnitude of the operator as needed.
     *
     * @param operator address to complete modifications for
     * @param strategies a list of strategies to complete modifications for
     * @param numToClear a list of number of pending modifications to complete for each strategy
     *
     * @dev can be called permissionlessly by anyone
     */
    function clearModificationQueue(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToClear
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        require(strategies.length == numToClear.length, InputArrayLengthMismatch());
        require(delegation.isOperator(operator), OperatorNotRegistered());
        for (uint256 i = 0; i < strategies.length; ++i) {
            _clearModificationQueue({operator: operator, strategy: strategies[i], numToClear: numToClear[i]});
        }
    }

    /**
     * @notice Modifies the propotions of slashable stake allocated to a list of operatorSets for a set of strategies
     * @param allocations array of magnitude adjustments for multiple strategies and corresponding operator sets
     * @dev Updates encumberedMagnitude for the updated strategies
     * @dev msg.sender is used as operator
     * @dev For each allocation, allocation.operatorSets MUST be ordered in ascending order according to the
     * encoding of the operatorSet. This is to prevent duplicate operatorSets being passed in. The easiest way to ensure
     * ordering is to sort allocated operatorSets by address first, and then sort for each avs by ascending operatorSetIds.
     */
    function modifyAllocations(
        MagnitudeAllocation[] calldata allocations
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        (bool isSet, uint32 operatorAllocationDelay) = allocationDelay(msg.sender);
        require(isSet, UninitializedAllocationDelay());

        for (uint256 i = 0; i < allocations.length; ++i) {
            MagnitudeAllocation calldata allocation = allocations[i];
            require(allocation.operatorSets.length == allocation.magnitudes.length, InputArrayLengthMismatch());
            require(avsDirectory.isOperatorSetBatch(allocation.operatorSets), InvalidOperatorSet());

            // 1. For the given (operator,strategy) complete any pending modifications to free up encumberedMagnitude
            _clearModificationQueue({operator: msg.sender, strategy: allocation.strategy, numToClear: type(uint16).max});

            // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
            // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
            // proprtional magnitudes relative to each other.
            uint64 totalMagnitude = uint64(_totalMagnitudeUpdate[msg.sender][allocation.strategy].latest());
            require(totalMagnitude == allocation.expectedTotalMagnitude, InvalidExpectedTotalMagnitude());

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[j]);
                // Check that there are no pending allocations & deallocations for the operator, operatorSet, strategy
                (bool noPendingModification, MagnitudeInfo memory mInfo) =
                    _completePendingModification(msg.sender, allocation.strategy, operatorSetKey);
                require(noPendingModification, ModificationAlreadyPending());

                // Calculate the new pending diff with this modification
                mInfo.pendingMagnitudeDelta = _calcDelta(mInfo.currentMagnitude, allocation.magnitudes[j]);
                require(mInfo.pendingMagnitudeDelta != 0, SameMagnitude());

                // Handle deallocation/allocation and modification effect timestamp
                if (mInfo.pendingMagnitudeDelta < 0) {
                    // This is a deallocation

                    // 1. Update the effect timestamp for the deallocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;
                } else if (mInfo.pendingMagnitudeDelta > 0) {
                    // This is an allocation

                    // 1. Update the effectTimestamp for the allocation
                    mInfo.effectTimestamp = uint32(block.timestamp) + operatorAllocationDelay;

                    // 2. add to the encumbered magnitude and make sure it doesn't exceed the total magnitude
                    _updateEncumberedMagnitude({
                        operator: msg.sender,
                        strategy: allocation.strategy,
                        encumberedMagnitudeDelta: mInfo.pendingMagnitudeDelta
                    });
                    // TODO: gross effect then check
                    require(
                        totalMagnitude >= encumberedMagnitude[msg.sender][allocation.strategy],
                        InsufficientAllocatableMagnitude()
                    );
                }

                // Add the operatorSet to the modification queue
                modificationQueue[msg.sender][allocation.strategy].pushBack(operatorSetKey);
                // Set the magnitude info
                _setMagnitudeInfo({
                    operator: msg.sender,
                    strategy: allocation.strategy,
                    operatorSetKey: operatorSetKey,
                    mInfo: mInfo,
                    isCompletion: false
                });
            }
        }
    }

    /**
     * @notice Called by an AVS to slash an operator for given operatorSetId, list of strategies, and bipsToSlash.
     * For each given (operator, operatorSetId, strategy) tuple, bipsToSlash
     * bips of the operatorSet's slashable stake allocation will be slashed
     *
     * @param operator the address to slash
     * @param operatorSetId the ID of the operatorSet the operator is being slashed on behalf of
     * @param strategies the set of strategies to slash
     * @param wadToSlash the parts in 1e18 to slash, this will be proportional to the
     * @param description the description of the slashing provided by the AVS for legibility
     * operator's slashable stake allocation for the operatorSet
     */
    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint256 wadToSlash,
        string calldata description
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING) {
        require(wadToSlash > 0, InvalidWadToSlash(wadToSlash));
        bytes32 operatorSetKey;
        {
            OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
            operatorSetKey = _encodeOperatorSet(operatorSet);
            require(avsDirectory.isOperatorSlashable(operator, operatorSet), InvalidOperator());
        }

        for (uint256 i = 0; i < strategies.length; ++i) {
            // 1. Complete pending modifications for the operator, strategy, and operatorSet
            (, MagnitudeInfo memory mInfo) = _completePendingModification(operator, strategies[i], operatorSetKey);

            // 2. Calculate the slashed magnitude and update operatorMagnitudeInfo
            uint64 slashedMagnitude = uint64(uint256(mInfo.currentMagnitude).mulWad(wadToSlash));
            mInfo.currentMagnitude -= slashedMagnitude;

            // if there is a pending deallocation, slash pending deallocation proportionally
            // since the delta is negative, we increase it to make sure they can free less of it when completed
            if (_isModificationPending(mInfo) && mInfo.pendingMagnitudeDelta < 0) {
                mInfo.pendingMagnitudeDelta += int128(uint128(uint256(uint128(-mInfo.pendingMagnitudeDelta)).mulWad(wadToSlash)));
            }

            // update operatorMagnitudeInfo
            _setMagnitudeInfo({
                operator: operator,
                strategy: strategies[i],
                operatorSetKey: operatorSetKey,
                mInfo: mInfo,
                isCompletion: false
            });

            // 3. update encumberedMagnitude, subtract the slashed magnitude
            _updateEncumberedMagnitude({
                operator: operator,
                strategy: strategies[i],
                encumberedMagnitudeDelta: -int128(uint128(slashedMagnitude))
            });

            // 3. update totalMagnitude, subtract the slashed magnitude and push the new totalMagnitude to the history
            Snapshots.DefaultWadHistory storage totalMagnitudes = _totalMagnitudeUpdate[operator][strategies[i]];
            uint64 totalMagnitudeAfterSlash = uint64(totalMagnitudes.latest()) - slashedMagnitude;
            totalMagnitudes.push({key: uint32(block.timestamp), value: totalMagnitudeAfterSlash});
            emit TotalMagnitudeUpdated(operator, strategies[i], totalMagnitudeAfterSlash);

            // 4. Decrease operators shares in the DM
            delegation.decreaseOperatorShares({
                operator: operator,
                strategy: strategies[i],
                previousTotalMagnitude: totalMagnitudeAfterSlash + slashedMagnitude,
                newTotalMagnitude: totalMagnitudeAfterSlash
            });
        }

        // TODO: find a solution to connect operatorSlashed to magnitude updates
        emit OperatorSlashed(operator, _decodeOperatorSet(operatorSetKey), description);
    }

    /**
     * @notice Helper for setting an operators allocation delay.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     */
    function _setAllocationDelay(address operator, uint32 delay) internal {
        require(delay != 0, InvalidDelay());

        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        if (info.pendingDelay != 0 && block.timestamp >= info.pendingDelayEffectTimestamp) {
            info.delay = info.pendingDelay;
        }

        info.pendingDelay = delay;
        info.pendingDelayEffectTimestamp = uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        _allocationDelayInfo[operator] = info;
        emit AllocationDelaySet(operator, delay, info.pendingDelayEffectTimestamp);
    }

    /**
     * @notice For a single strategy, clear the queue of pending modifications for an operator, strategy by completing them if necessary
     * @param operator address to update allocatableMagnitude for
     * @param strategy the strategy to update allocatableMagnitude for
     * @param numToClear the number of pending modifications to complete
     */
    function _clearModificationQueue(address operator, IStrategy strategy, uint16 numToClear) internal {
        uint256 numCompleted;
        while (modificationQueue[operator][strategy].length() > 0 && numCompleted < numToClear) {
            bytes32 opsetKey = modificationQueue[operator][strategy].front();
            (bool completed,) = _completePendingModification(operator, strategy, opsetKey);
            if (completed) {
                // if the modification was completed, remove it from the queue
                modificationQueue[operator][strategy].popFront();
                ++numCompleted;
            } else {
                // else, the modification is still pending, and so are all subsequent modifications, so break
                break;
            }
        }
    }

    // returns whether the modification was completed
    function _completePendingModification(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey
    ) internal returns (bool, MagnitudeInfo memory) {
        MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];
        // if the effect timestamp is not reached, return false
        if (block.timestamp < mInfo.effectTimestamp) {
            return (false, mInfo);
        } else if (mInfo.effectTimestamp == 0) {
            // if there is no effect timestamp, return true and continue completing
            return (true, mInfo);
        }

        // update stored magnitudes
        mInfo.currentMagnitude = _addInt128(mInfo.currentMagnitude, mInfo.pendingMagnitudeDelta);
        mInfo.effectTimestamp = 0;
        if (mInfo.pendingMagnitudeDelta < 0) {
            // if this was a deallocation, update the encumbered magnitude
            _updateEncumberedMagnitude({
                operator: operator,
                strategy: strategy,
                encumberedMagnitudeDelta: mInfo.pendingMagnitudeDelta
            });
        }
        mInfo.pendingMagnitudeDelta = 0;

        // update the magnitude info in storage
        _setMagnitudeInfo({
            operator: operator,
            strategy: strategy,
            operatorSetKey: operatorSetKey,
            mInfo: mInfo,
            isCompletion: true
        });

        return (true, mInfo);
    }

    function _setMagnitudeInfo(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey,
        MagnitudeInfo memory mInfo,
        bool isCompletion
    ) internal {
        MagnitudeInfo memory mInfoStored = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];
        OperatorSet memory operatorSet = _decodeOperatorSet(operatorSetKey);

        // if this is not a completion, emit events
        // completions have their events emmitted in advance
        if (!isCompletion) {
            // if the current magnitude has changed, emit an event
            if (mInfoStored.currentMagnitude != mInfo.currentMagnitude) {
                emit OperatorSetMagnitudeUpdated(
                    operator, operatorSet, strategy, mInfo.currentMagnitude, uint32(block.timestamp)
                );
            }

            // if the pending magnitude delta has changed, emit an event
            // forgefmt: disable-next-item
            if(mInfoStored.pendingMagnitudeDelta != mInfo.pendingMagnitudeDelta) {
                emit OperatorSetMagnitudeUpdated(
                    operator, operatorSet, strategy, _addInt128(mInfo.currentMagnitude, mInfo.pendingMagnitudeDelta), mInfo.effectTimestamp
                );
            }
        }

        // actually set the magnitude info
        _operatorMagnitudeInfo[operator][strategy][operatorSetKey] = mInfo;
    }

    function _updateEncumberedMagnitude(
        address operator,
        IStrategy strategy,
        int128 encumberedMagnitudeDelta
    ) internal {
        uint64 newEncumberedMagnitude = _addInt128(encumberedMagnitude[operator][strategy], encumberedMagnitudeDelta);
        encumberedMagnitude[operator][strategy] = newEncumberedMagnitude;
        emit EncumberedMagnitudeUpdated(operator, strategy, newEncumberedMagnitude);
    }

    /**
     * @notice Returns the allocation delay of an operator
     * @param operator The operator to get the allocation delay for
     */
    function allocationDelay(
        address operator
    ) public view returns (bool isSet, uint32 delay) {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];
        if (info.pendingDelay != 0 && block.timestamp >= info.pendingDelayEffectTimestamp) {
            return (true, info.pendingDelay);
        } else {
            return (info.delay != 0, info.delay);
        }
    }

    /**
     * @param operator the operator to get the slashable magnitude for
     * @param strategies the strategies to get the slashable magnitude for
     *
     * @return operatorSets the operator sets the operator is a member of and the current slashable magnitudes for each strategy
     */
    function getSlashableMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (OperatorSet[] memory, uint64[][] memory) {
        OperatorSet[] memory operatorSets = avsDirectory.getOperatorSetsOfOperator(operator, 0, type(uint256).max);
        uint64[][] memory slashableMagnitudes = new uint64[][](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            slashableMagnitudes[i] = new uint64[](operatorSets.length);
            for (uint256 j = 0; j < operatorSets.length; ++j) {
                MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategies[i]][_encodeOperatorSet(operatorSets[j])];
                slashableMagnitudes[i][j] = mInfo.currentMagnitude;
                if (_isModificationCompletable(mInfo)) {
                    slashableMagnitudes[i][j] = _addInt128(slashableMagnitudes[i][j], mInfo.pendingMagnitudeDelta);
                }
            }
        }
        return (operatorSets, slashableMagnitudes);
    }

    /**
     * @notice Get the allocatable magnitude for an operator and strategy
     * @param operator the operator to get the allocatable magnitude for
     * @param strategy the strategy to get the allocatable magnitude for
     */
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        uint64 freeableMagnitude = 0;
        for (uint256 i = 0; i < modificationQueue[operator][strategy].length(); ++i) {
            bytes32 opsetKey = modificationQueue[operator][strategy].at(i);
            MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][opsetKey];
            if (_isModificationCompletable(mInfo) && mInfo.pendingMagnitudeDelta < 0) {
                freeableMagnitude += uint64(uint128(-mInfo.pendingMagnitudeDelta));
            } else {
                break;
            }
        }
        return uint64(_totalMagnitudeUpdate[operator][strategy].latest()) - encumberedMagnitude[operator][strategy]
            + freeableMagnitude;
    }

    /**
     * @notice Returns the current total magnitudes of an operator for a given set of strategies
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            totalMagnitudes[i] = uint64(_totalMagnitudeUpdate[operator][strategies[i]].latest());
        }
        return totalMagnitudes;
    }

    /**
     * @notice Returns the total magnitudes of an operator for a given set of strategies at a given timestamp
     * @param operator the operator to get the total magnitude for
     * @param strategies the strategies to get the total magnitudes for
     * @param timestamp the timestamp to get the total magnitudes at
     * @return totalMagnitudes the total magnitudes for each strategy
     */
    function getTotalMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory) {
        uint64[] memory totalMagnitudes = new uint64[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
            totalMagnitudes[i] = uint64(_totalMagnitudeUpdate[operator][strategies[i]].upperLookup(timestamp));
        }
        return totalMagnitudes;
    }

    /**
     * @notice Returns the pending modifications of an operator for a given strategy and operatorSets.
     * @param operator the operator to get the pending modification for
     * @param strategy the strategy to get the pending modification for
     * @param operatorSets the operatorSets to get the pending modification for
     * @return timestamps the timestamps for each pending dealloction
     * @return pendingMagnitudeDeltas the pending modification diffs for each operatorSet
     */
    function getPendingModifications(
        address operator,
        IStrategy strategy,
        OperatorSet[] calldata operatorSets
    ) external view returns (uint32[] memory timestamps, int128[] memory pendingMagnitudeDeltas) {
        timestamps = new uint32[](operatorSets.length);
        pendingMagnitudeDeltas = new int128[](operatorSets.length);

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            MagnitudeInfo memory mInfo =
                _operatorMagnitudeInfo[operator][strategy][_encodeOperatorSet(operatorSets[i])];

            if (_isModificationPending(mInfo)) {
                pendingMagnitudeDeltas[i] = mInfo.pendingMagnitudeDelta;
                timestamps[i] = mInfo.effectTimestamp;
            }
        }
    }

    function _isModificationPending(MagnitudeInfo memory mInfo) internal view returns (bool) {
        return mInfo.effectTimestamp != 0 && mInfo.effectTimestamp > block.timestamp;
    }

    function _isModificationCompletable(MagnitudeInfo memory mInfo) internal view returns (bool) {
        return mInfo.effectTimestamp != 0 && mInfo.effectTimestamp <= block.timestamp;
    }

    function _calcDelta(uint64 currentMagnitude, uint64 newMagnitude) internal pure returns (int128) {
        return int128(uint128(newMagnitude)) - int128(uint128(currentMagnitude));
    }

    function _addInt128(uint64 a, int128 b) internal pure returns (uint64) {
        return uint64(uint128(int128(uint128(a)) + b));
    }

    /// @dev Returns an `OperatorSet` encoded into a 32-byte value.
    /// @param operatorSet The `OperatorSet` to encode.
    function _encodeOperatorSet(
        OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(operatorSet.avs, uint96(operatorSet.operatorSetId)));
    }

    /// @dev Returns an `OperatorSet` decoded from an encoded 32-byte value.
    /// @param encoded The encoded `OperatorSet` to decode.
    /// @dev Assumes `encoded` is encoded via `_encodeOperatorSet(operatorSet)`.
    function _decodeOperatorSet(
        bytes32 encoded
    ) internal pure returns (OperatorSet memory) {
        return OperatorSet({
            avs: address(uint160(uint256(encoded) >> 96)),
            operatorSetId: uint32(uint256(encoded) & type(uint96).max)
        });
    }
}
