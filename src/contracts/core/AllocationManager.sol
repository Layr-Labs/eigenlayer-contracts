// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
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

    /// @inheritdoc IAllocationManager
    function slashOperator(
        SlashingParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING) {
        require(0 < params.wadToSlash && params.wadToSlash <= WAD, InvalidWadToSlash());

        // Check that the operator is registered and slashable
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: params.operatorSetId});
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        require(avsDirectory.isOperatorSlashable(params.operator, operatorSet), InvalidOperator());

        for (uint256 i = 0; i < params.strategies.length; ++i) {
            PendingMagnitudeInfo memory info =
                _getPendingMagnitudeInfo(params.operator, params.strategies[i], operatorSetKey);

            // 1. Calculate slashing amount and update current/ encumbered magnitude
            uint64 slashedMagnitude = uint64(uint256(info.currentMagnitude).mulWad(params.wadToSlash));
            info.currentMagnitude -= slashedMagnitude;
            info.encumberedMagnitude -= slashedMagnitude;

            // 2. If there is a pending deallocation, reduce pending deallocation proportionally.
            // This ensures that when the deallocation is completed, less magnitude is freed.
            if (info.pendingDiff < 0) {
                uint64 slashedPending = uint64(uint256(uint128(-info.pendingDiff)).mulWad(params.wadToSlash));
                info.pendingDiff += int128(uint128(slashedPending));

                emit OperatorSetMagnitudeUpdated(
                    params.operator,
                    operatorSet,
                    params.strategies[i],
                    _addInt128(info.currentMagnitude, info.pendingDiff),
                    info.effectTimestamp
                );
            }

            // 3. Update the operator's allocation in storage
            _updateMagnitudeInfo({
                operator: params.operator,
                strategy: params.strategies[i],
                operatorSetKey: operatorSetKey,
                info: info
            });

            emit OperatorSetMagnitudeUpdated(
                params.operator, operatorSet, params.strategies[i], info.currentMagnitude, uint32(block.timestamp)
            );

            // 4. Reduce the operator's max magnitude
            uint64 maxMagnitudeBeforeSlash = _maxMagnitudeHistory[params.operator][params.strategies[i]].latest();
            uint64 maxMagnitudeAfterSlash = maxMagnitudeBeforeSlash - slashedMagnitude;
            _maxMagnitudeHistory[params.operator][params.strategies[i]].push({
                key: uint32(block.timestamp),
                value: maxMagnitudeAfterSlash
            });
            emit TotalMagnitudeUpdated(params.operator, params.strategies[i], maxMagnitudeAfterSlash);

            // 5. Decrease operators shares in the DelegationManager
            delegation.decreaseOperatorShares({
                operator: params.operator,
                strategy: params.strategies[i],
                previousTotalMagnitude: maxMagnitudeBeforeSlash,
                newTotalMagnitude: maxMagnitudeAfterSlash
            });
        }

        // TODO: find a solution to connect operatorSlashed to magnitude updates
        emit OperatorSlashed(params.operator, operatorSet, params.strategies, params.wadToSlash, params.description);
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
        (bool isSet, uint32 operatorAllocationDelay) = getAllocationDelay(msg.sender);
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
            uint64 maxMagnitude = _maxMagnitudeHistory[msg.sender][allocation.strategy].latest();
            require(maxMagnitude == allocation.expectedMaxMagnitude, InvalidExpectedTotalMagnitude());

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                bytes32 operatorSetKey = _encodeOperatorSet(allocation.operatorSets[j]);

                // Ensure there is not already a pending modification
                PendingMagnitudeInfo memory info =
                    _getPendingMagnitudeInfo(msg.sender, allocation.strategy, operatorSetKey);
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

    /// @inheritdoc IAllocationManager
    function setAllocationDelay(address operator, uint32 delay) external {
        require(msg.sender == address(delegation), OnlyDelegationManager());
        _setAllocationDelay(operator, delay);
    }

    /// @inheritdoc IAllocationManager
    function setAllocationDelay(
        uint32 delay
    ) external {
        require(delegation.isOperator(msg.sender), OperatorNotRegistered());
        _setAllocationDelay(msg.sender, delay);
    }

    /**
     * @dev Clear one or more pending modifications to a strategy's allocated magnitude
     * @param operator the operator whose pending modifications will be cleared
     * @param strategy the strategy to update
     * @param numToClear the number of pending modifications to complete
     */
    function _clearModificationQueue(address operator, IStrategy strategy, uint16 numToClear) internal {
        uint256 numCompleted;
        uint256 length = modificationQueue[operator][strategy].length();

        while (length > 0 && numCompleted < numToClear) {
            bytes32 operatorSetKey = modificationQueue[operator][strategy].front();
            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(operator, strategy, operatorSetKey);

            // If we've reached a pending modification that isn't completable yet,
            // we can stop. Any subsequent modificaitons will also be uncompletable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // Update the operator's allocation in storage
            _updateMagnitudeInfo(operator, strategy, operatorSetKey, info);

            // Remove the modification from the queue
            modificationQueue[operator][strategy].popFront();
            ++numCompleted;
            ++length;
        }
    }

    /**
     * @dev Sets the operator's allocation delay. This is the time between an operator
     * allocating magnitude to an operator set, and the magnitude becoming slashable.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     */
    function _setAllocationDelay(address operator, uint32 delay) internal {
        require(delay != 0, InvalidDelay());

        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        if (info.pendingDelay != 0 && block.timestamp >= info.effectTimestamp) {
            info.delay = info.pendingDelay;
        }

        info.pendingDelay = delay;
        info.effectTimestamp = uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        _allocationDelayInfo[operator] = info;
        emit AllocationDelaySet(operator, delay, info.effectTimestamp);
    }

    /**
     * @param encumberedMagnitude the effective magnitude allocated to all operator sets
     * for the strategy
     * @param currentMagnitude the effective current magnitude allocated to a single operator set
     * for the strategy
     * @param pendingDiff the pending change in magnitude, if one exists
     * @param effectTimestamp the time after which `pendingDiff` will take effect
     */
    struct PendingMagnitudeInfo {
        uint64 encumberedMagnitude;
        uint64 currentMagnitude;
        int128 pendingDiff;
        uint32 effectTimestamp;
    }

    /**
     * @dev For an operator set, get the operator's effective allocated magnitude.
     * If the operator set has a pending modification that can be completed at the
     * current timestamp, this method returns a view of the allocation as if the modification
     * was completed.
     * @return info the effective allocated and pending magnitude for the operator set, and
     * the effective encumbered magnitude for all operator sets belonging to this strategy
     */
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
                encumberedMagnitude: _encumberedMagnitude,
                currentMagnitude: mInfo.currentMagnitude,
                pendingDiff: mInfo.pendingDiff,
                effectTimestamp: mInfo.effectTimestamp
            });
        }

        // Pending change can be completed - add delta to current magnitude
        info.currentMagnitude = _addInt128(mInfo.currentMagnitude, mInfo.pendingDiff);
        info.effectTimestamp = 0;
        info.pendingDiff = 0;

        // If the completed change was a deallocation, update encumbered magnitude
        if (mInfo.pendingDiff < 0) {
            info.encumberedMagnitude = _addInt128(_encumberedMagnitude, mInfo.pendingDiff);
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
            pendingDiff: info.pendingDiff,
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
                    slashableMagnitudes[i][j] = _addInt128(slashableMagnitudes[i][j], mInfo.pendingDiff);
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
            if (block.timestamp >= mInfo.effectTimestamp && mInfo.effectTimestamp != 0 && mInfo.pendingDiff < 0) {
                freeableMagnitude += uint64(uint128(-mInfo.pendingDiff));
            } else {
                break;
            }
        }
        return uint64(_maxMagnitudeHistory[operator][strategy].latest()) - encumberedMagnitude[operator][strategy]
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
            (bool exists,, uint64 value) = _maxMagnitudeHistory[operator][strategies[i]].latestSnapshot();
            if (!exists) {
                totalMagnitudes[i] = WAD;
            } else {
                totalMagnitudes[i] = value;
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
            totalMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup(timestamp);
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
                pendingMagnitudeDeltas[i] = opsetMagnitudeInfo.pendingDiff;
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

    /// @inheritdoc IAllocationManager
    function getAllocationInfo(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, MagnitudeInfo[] memory) {
        OperatorSet[] memory operatorSets = avsDirectory.getOperatorSetsOfOperator(operator, 0, type(uint256).max);
        MagnitudeInfo[] memory infos = getAllocationInfo(operator, strategy, operatorSets);

        return (operatorSets, infos);
    }

    /// @inheritdoc IAllocationManager
    function getAllocationInfo(
        address operator,
        IStrategy strategy,
        OperatorSet[] memory operatorSets
    ) public view returns (MagnitudeInfo[] memory) {
        MagnitudeInfo[] memory infos = new MagnitudeInfo[](operatorSets.length);

        for (uint256 i = 0; i < operatorSets.length; ++i) {
            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo({
                operator: operator,
                strategy: strategy,
                operatorSetKey: _encodeOperatorSet(operatorSets[i])
            });

            infos[i] = MagnitudeInfo({
                currentMagnitude: info.currentMagnitude,
                pendingDiff: info.pendingDiff,
                effectTimestamp: info.effectTimestamp
            });
        }

        return infos;
    }

    /// @inheritdoc IAllocationManager
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        // This method needs to simulate clearing any pending allocation modifications.
        // This roughly mimics the calculations done in `_clearModificationQueue` and
        // `_getPendingMagnitudeInfo`, while operating on a `curEncumberedMagnitude`
        // rather than continually reading/updating state.
        uint64 curEncumberedMagnitude = encumberedMagnitude[operator][strategy];

        uint256 length = modificationQueue[operator][strategy].length();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 operatorSetKey = modificationQueue[operator][strategy].at(i);
            MagnitudeInfo memory info = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];

            // If we've reached a pending modification that isn't completable yet,
            // we can stop. Any subsequent modificaitons will also be uncompletable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // If the diff is a deallocation, add to encumbered magnitude. Allocations
            // do not need to be considered, because encumbered magnitude is updated as
            // soon as the allocation is created.
            if (info.pendingDiff < 0) {
                curEncumberedMagnitude = _addInt128(curEncumberedMagnitude, info.pendingDiff);
            }
        }

        // The difference between the operator's max magnitude and its encumbered magnitude
        // is the magnitude that can be allocated.
        return _maxMagnitudeHistory[operator][strategy].latest() - curEncumberedMagnitude;
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].latest();
        }

        return maxMagnitudes;
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitudesAtTimestamp(
        address operator,
        IStrategy[] calldata strategies,
        uint32 timestamp
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup(timestamp);
        }

        return maxMagnitudes;
    }

    /**
     * @notice Returns the time in seconds between an operator allocating slashable magnitude
     * and the magnitude becoming slashable. If the delay has not been set, `isSet` will be false.
     * @dev The operator must have a configured delay before allocating magnitude
     * @param operator The operator to query
     * @return isSet Whether the operator has configured a delay
     * @return delay The time in seconds between allocating magnitude and magnitude becoming slashable
     */
    /// @inheritdoc IAllocationManager
    function getAllocationDelay(
        address operator
    ) public view returns (bool isSet, uint32 delay) {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        if (info.pendingDelay != 0 && block.timestamp >= info.effectTimestamp) {
            delay = info.pendingDelay;
        } else {
            delay = info.delay;
        }

        // Operators cannot configure their allocation delay to be zero, so the delay has been
        // set as long as it is nonzero.
        isSet = delay != 0;
        return (isSet, delay);
    }
}
