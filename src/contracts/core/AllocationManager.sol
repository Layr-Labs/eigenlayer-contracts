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
            uint64 maxMagnitude = uint64(_totalMagnitudeUpdate[msg.sender][allocation.strategy].latest());
            require(maxMagnitude == allocation.expectedTotalMagnitude, InvalidExpectedTotalMagnitude());

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[j]);

                // Ensure there is not already a pending modification
                PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(msg.sender, allocation.strategy, operatorSetKey);
                require(info.pendingDiff == 0, ModificationAlreadyPending());

                info.pendingDiff = _calcDelta(info.currentMagnitude, allocation.magnitudes[j]);
                require(info.pendingDiff != 0, SameMagnitude());

                // Calculate the effectTimestamp for the modification
                if (info.pendingDiff < 0) {
                    info.effectTimestamp = uint32(block.timestamp) + DEALLOCATION_DELAY;
                } else if (info.pendingDiff > 0) {
                    info.effectTimestamp = uint32(block.timestamp) + operatorAllocationDelay;

                    // For allocations, immediately add to encumberedMagnitude to ensure the operator
                    // can't allocate more than their maximum
                    info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, info.pendingDiff);
                    require(info.encumberedMagnitude <= maxMagnitude, InsufficientAllocatableMagnitude());
                }

                // Add the operatorSet to the modification queue and update the allocation
                // in storage
                modificationQueue[msg.sender][allocation.strategy].pushBack(operatorSetKey);
                _updateMagnitudeInfo({
                    operator: msg.sender,
                    strategy: allocation.strategy,
                    operatorSetKey: operatorSetKey,
                    info: info
                });

                emit OperatorSetMagnitudeUpdated(
                    msg.sender,
                    _decodeOperatorSet(operatorSetKey),
                    allocation.strategy, 
                    _addInt128(info.currentMagnitude, info.pendingDiff), 
                    info.effectTimestamp
                );
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
        require(wadToSlash > 0 && wadToSlash <= WAD, InvalidWadToSlash(wadToSlash));
        bytes32 operatorSetKey;
        {
            OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
            operatorSetKey = _encodeOperatorSet(operatorSet);
            require(avsDirectory.isOperatorSlashable(operator, operatorSet), InvalidOperator());
        }
        
        for (uint256 i = 0; i < strategies.length; ++i) {
            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(operator, strategies[i], operatorSetKey);

            // 1. Calculate slashing amount and update current/ encumbered magnitude
            uint64 slashedMagnitude = uint64(uint256(info.currentMagnitude).mulWad(wadToSlash));
            info.currentMagnitude -= slashedMagnitude;
            info.encumberedMagnitude -= slashedMagnitude;

            // 2. If there is a pending deallocation, reduce pending deallocation proportionally.
            // This ensures that when the deallocation is completed, less magnitude is freed.
            if (info.pendingDiff < 0) {
                uint64 slashedPending = uint64(uint256(uint128(-info.pendingDiff)).mulWad(wadToSlash));
                info.pendingDiff += int128(uint128(slashedPending));

                emit OperatorSetMagnitudeUpdated(
                    operator,
                    _decodeOperatorSet(operatorSetKey),
                    strategies[i],
                    _addInt128(info.currentMagnitude, info.pendingDiff),
                    info.effectTimestamp
                );
            }

            // 3. Update the operator's allocation in storage
            _updateMagnitudeInfo({
                operator: operator,
                strategy: strategies[i],
                operatorSetKey: operatorSetKey,
                info: info
            });

            emit OperatorSetMagnitudeUpdated(
                operator,
                _decodeOperatorSet(operatorSetKey),
                strategies[i],
                info.currentMagnitude,
                uint32(block.timestamp)
            );

            // 4. Reduce the operator's max magnitude
            Snapshots.DefaultWadHistory storage totalMagnitudes = _totalMagnitudeUpdate[operator][strategies[i]];
            uint64 totalMagnitudeAfterSlash = uint64(totalMagnitudes.latest()) - slashedMagnitude;
            totalMagnitudes.push({key: uint32(block.timestamp), value: totalMagnitudeAfterSlash});
            emit TotalMagnitudeUpdated(operator, strategies[i], totalMagnitudeAfterSlash);

            // 5. Decrease operators shares in the DelegationManager
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
            bytes32 opSetKey = modificationQueue[operator][strategy].front();

            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(operator, strategy, opSetKey);

            // If we've reached a pending modification that isn't completable yet,
            // we can stop. Any subsequent modificaitons will also be uncompletable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // Update the operator's allocation in storage
            _updateMagnitudeInfo(operator, strategy, opSetKey, info);

            // Remove the modification from the queue
            modificationQueue[operator][strategy].popFront();
            ++numCompleted;
        }
    }

    struct PendingMagnitudeInfo {
        // Current magnitude allocated to this operatorSet
        uint64 currentMagnitude;

        // Pending magnitude delta
        int128 pendingDiff;
        // The timestamp after which `pendingDiff` is completable
        uint32 effectTimestamp;
        
        // Current encumbered magnitude for all opSets in this strategy
        uint64 encumberedMagnitude;
    }

    function _getPendingMagnitudeInfo(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey
    ) internal view returns (PendingMagnitudeInfo memory info) {
        MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];
        uint64 _encumberedMagnitude = encumberedMagnitude[operator][strategy];

        // If the pending change can't be completed yet
        if (block.timestamp < mInfo.effectTimestamp) {
            return PendingMagnitudeInfo({
                currentMagnitude: mInfo.currentMagnitude,
                pendingDiff: mInfo.pendingMagnitudeDelta,
                effectTimestamp: mInfo.effectTimestamp,
                encumberedMagnitude: _encumberedMagnitude
            });
        }

        // Pending change can be completed - add delta to current magnitude
        info.currentMagnitude = _addInt128(mInfo.currentMagnitude, mInfo.pendingMagnitudeDelta);
        info.effectTimestamp = 0;
        info.pendingDiff = 0;

        // If the completed change was a deallocation, update encumbered magnitude
        if (mInfo.pendingMagnitudeDelta < 0) {
            info.encumberedMagnitude = _addInt128(_encumberedMagnitude, mInfo.pendingMagnitudeDelta);
        }

        return info;
    }

    function _updateMagnitudeInfo(
        address operator,
        IStrategy strategy,
        bytes32 operatorSetKey,
        PendingMagnitudeInfo memory info
    ) internal {
        _operatorMagnitudeInfo[operator][strategy][operatorSetKey] = MagnitudeInfo({
            currentMagnitude: info.currentMagnitude,
            pendingMagnitudeDelta: info.pendingDiff,
            effectTimestamp: info.effectTimestamp
        });

        encumberedMagnitude[operator][strategy] = info.encumberedMagnitude;
        emit EncumberedMagnitudeUpdated(operator, strategy, info.encumberedMagnitude);
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
                bytes32 operatorSetKey = _encodeOperatorSet(operatorSets[j]);
                MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategies[i]][operatorSetKey];
                slashableMagnitudes[i][j] = mInfo.currentMagnitude;
                if (block.timestamp >= mInfo.effectTimestamp && mInfo.effectTimestamp != 0) {
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
            if (
                block.timestamp >= mInfo.effectTimestamp && mInfo.effectTimestamp != 0
                    && mInfo.pendingMagnitudeDelta < 0
            ) {
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
            (bool exists,, uint224 value) = _totalMagnitudeUpdate[operator][strategies[i]].latestSnapshot();
            if (!exists) {
                totalMagnitudes[i] = WAD;
            } else {
                totalMagnitudes[i] = uint64(value);
            }
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
            MagnitudeInfo memory opsetMagnitudeInfo =
                _operatorMagnitudeInfo[operator][strategy][_encodeOperatorSet(operatorSets[i])];

            if (opsetMagnitudeInfo.effectTimestamp < block.timestamp && opsetMagnitudeInfo.effectTimestamp != 0) {
                pendingMagnitudeDeltas[i] = opsetMagnitudeInfo.pendingMagnitudeDelta;
                timestamps[i] = opsetMagnitudeInfo.effectTimestamp;
            }
        }
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
