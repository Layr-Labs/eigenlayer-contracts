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
        IPauserRegistry _pauserRegistry,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    )
        AllocationManagerStorage(_delegation, _avsDirectory, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
        Pausable(_pauserRegistry)
    {
        _disableInitializers();
    }

    /// @inheritdoc IAllocationManager
    function initialize(address initialOwner, uint256 initialPausedStatus) external initializer {
        _setPausedStatus(initialPausedStatus);
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

        // Record the proportion of 1e18 that the operator's total shares that are being slashed
        uint256[] memory wadSlashed = new uint256[](params.strategies.length);

        for (uint256 i = 0; i < params.strategies.length; ++i) {
            PendingMagnitudeInfo memory info =
                _getPendingMagnitudeInfo(params.operator, params.strategies[i], operatorSetKey);

            require(info.currentMagnitude > 0, OperatorNotAllocated());

            // 1. Calculate slashing amount and update current/encumbered magnitude
            uint64 slashedMagnitude = uint64(uint256(info.currentMagnitude).mulWadRoundUp(params.wadToSlash));
            info.currentMagnitude -= slashedMagnitude;
            info.encumberedMagnitude -= slashedMagnitude;

            // 2. If there is a pending deallocation, reduce pending deallocation proportionally.
            // This ensures that when the deallocation is cleared, less magnitude is freed.
            if (info.pendingDiff < 0) {
                uint64 slashedPending = uint64(uint256(uint128(-info.pendingDiff)).mulWadRoundUp(params.wadToSlash));
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
            emit MaxMagnitudeUpdated(params.operator, params.strategies[i], maxMagnitudeAfterSlash);

            // 5. Decrease operators shares in the DelegationManager
            delegation.decreaseOperatorShares({
                operator: params.operator,
                strategy: params.strategies[i],
                previousMaxMagnitude: maxMagnitudeBeforeSlash,
                newMaxMagnitude: maxMagnitudeAfterSlash
            });

            // 6. Record the proportion of shares slashed
            wadSlashed[i] = uint256(slashedMagnitude).divWad(maxMagnitudeBeforeSlash);
        }

        emit OperatorSlashed(params.operator, operatorSet, params.strategies, wadSlashed, params.description);
    }

    /// @inheritdoc IAllocationManager
    function modifyAllocations(
        MagnitudeAllocation[] calldata allocations
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        (bool isSet, uint32 operatorAllocationDelay) = getAllocationDelay(msg.sender);
        require(isSet, UninitializedAllocationDelay());

        for (uint256 i = 0; i < allocations.length; ++i) {
            MagnitudeAllocation calldata allocation = allocations[i];
            require(allocation.operatorSets.length == allocation.magnitudes.length, InputArrayLengthMismatch());
            require(avsDirectory.isOperatorSetBatch(allocation.operatorSets), InvalidOperatorSet());

            // 1. Check current maxMagnitude matches expected value. This is to check for slashing race conditions
            // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
            // proprtional magnitudes relative to each other.
            uint64 maxMagnitude = _maxMagnitudeHistory[msg.sender][allocation.strategy].latest();
            require(maxMagnitude == allocation.expectedMaxMagnitude, InvalidExpectedMaxMagnitude());

            // 2. For the given (operator,strategy) clear any clearable pending deallocations to free up encumberedMagnitude
            _clearDeallocationQueue({operator: msg.sender, strategy: allocation.strategy, numToClear: type(uint16).max});

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

                    // Add the operatorSet to the deallocation queue
                    deallocationQueue[msg.sender][allocation.strategy].pushBack(operatorSetKey);
                } else if (info.pendingDiff > 0) {
                    info.effectTimestamp = uint32(block.timestamp) + operatorAllocationDelay;

                    // For allocations, immediately add to encumberedMagnitude to ensure the operator
                    // can't allocate more than their maximum
                    info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, info.pendingDiff);
                    require(info.encumberedMagnitude <= maxMagnitude, InsufficientAllocatableMagnitude());
                }

                // Update the modification in storage
                _updateMagnitudeInfo({
                    operator: msg.sender,
                    strategy: allocation.strategy,
                    operatorSetKey: operatorSetKey,
                    info: info
                });

                emit OperatorSetMagnitudeUpdated(
                    msg.sender,
                    allocation.operatorSets[j],
                    allocation.strategy,
                    _addInt128(info.currentMagnitude, info.pendingDiff),
                    info.effectTimestamp
                );
            }
        }
    }

    /// @inheritdoc IAllocationManager
    function clearDeallocationQueue(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata numToClear
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        require(strategies.length == numToClear.length, InputArrayLengthMismatch());
        require(delegation.isOperator(operator), OperatorNotRegistered());

        for (uint256 i = 0; i < strategies.length; ++i) {
            _clearDeallocationQueue({operator: operator, strategy: strategies[i], numToClear: numToClear[i]});
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
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @dev Clear one or more pending deallocations to a strategy's allocated magnitude
     * @param operator the operator whose pending deallocations will be cleared
     * @param strategy the strategy to update
     * @param numToClear the number of pending deallocations to clear
     */
    function _clearDeallocationQueue(address operator, IStrategy strategy, uint16 numToClear) internal {
        uint256 numCleared;
        uint256 length = deallocationQueue[operator][strategy].length();

        while (length > 0 && numCleared < numToClear) {
            bytes32 operatorSetKey = deallocationQueue[operator][strategy].front();
            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(operator, strategy, operatorSetKey);

            // If we've reached a pending deallocation that isn't clearable yet,
            // we can stop. Any subsequent deallocation will also be unclearable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // Update the operator's allocation in storage
            _updateMagnitudeInfo(operator, strategy, operatorSetKey, info);

            // Remove the deallocation from the queue
            deallocationQueue[operator][strategy].popFront();
            ++numCleared;
            --length;
        }
    }

    /**
     * @dev Sets the operator's allocation delay. This is the time between an operator
     * allocating magnitude to an operator set, and the magnitude becoming slashable.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in seconds.
     */
    function _setAllocationDelay(address operator, uint32 delay) internal {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        // If there is a pending delay that can be applied now, set it
        if (info.effectTimestamp != 0 && block.timestamp >= info.effectTimestamp) {
            info.delay = info.pendingDelay;
            info.isSet = true;
        }

        info.pendingDelay = delay;
        info.effectTimestamp = uint32(block.timestamp + ALLOCATION_CONFIGURATION_DELAY);

        _allocationDelayInfo[operator] = info;
        emit AllocationDelaySet(operator, delay, info.effectTimestamp);
    }

    /**
     * @dev For an operator set, get the operator's effective allocated magnitude.
     * If the operator set has a pending deallocation that can be cleared at the
     * current timestamp, this method returns a view of the allocation as if the deallocation
     * was cleared.
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

        // If the pending change can't be cleared yet
        if (block.timestamp < mInfo.effectTimestamp) {
            return PendingMagnitudeInfo({
                encumberedMagnitude: _encumberedMagnitude,
                currentMagnitude: mInfo.currentMagnitude,
                pendingDiff: mInfo.pendingDiff,
                effectTimestamp: mInfo.effectTimestamp
            });
        }

        // Pending change can be cleared - add delta to current magnitude
        info.currentMagnitude = _addInt128(mInfo.currentMagnitude, mInfo.pendingDiff);
        info.encumberedMagnitude = _encumberedMagnitude;
        info.effectTimestamp = 0;
        info.pendingDiff = 0;

        // If the cleared change was a deallocation, update encumbered magnitude
        if (mInfo.pendingDiff < 0) {
            info.encumberedMagnitude = _addInt128(_encumberedMagnitude, mInfo.pendingDiff);
        }

        return info;
    }

    /// @notice Update the operator's magnitude info in storage and their encumbered magnitude.
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

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IAllocationManager
    function getAllocationInfo(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, MagnitudeInfo[] memory) {
        OperatorSet[] memory operatorSets =
            avsDirectory.getOperatorSetsOfOperator({operator: operator, start: 0, length: type(uint256).max});
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
    function getAllocationInfo(
        OperatorSet calldata operatorSet,
        IStrategy[] calldata strategies,
        address[] calldata operators
    ) public view returns (MagnitudeInfo[][] memory) {
        MagnitudeInfo[][] memory infos = new MagnitudeInfo[][](operators.length);
        for (uint256 i = 0; i < operators.length; ++i) {
            for (uint256 j = 0; j < strategies.length; ++j) {
                PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo({
                    operator: operators[i],
                    strategy: strategies[j],
                    operatorSetKey: _encodeOperatorSet(operatorSet)
                });

                infos[i][j] = MagnitudeInfo({
                    currentMagnitude: info.currentMagnitude,
                    pendingDiff: info.pendingDiff,
                    effectTimestamp: info.effectTimestamp
                });
            }
        }

        return infos;
    }

    /// @inheritdoc IAllocationManager
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        // This method needs to simulate clearing any pending deallocations.
        // This roughly mimics the calculations done in `_clearDeallocationQueue` and
        // `_getPendingMagnitudeInfo`, while operating on a `curEncumberedMagnitude`
        // rather than continually reading/updating state.
        uint64 curEncumberedMagnitude = encumberedMagnitude[operator][strategy];

        uint256 length = deallocationQueue[operator][strategy].length();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 operatorSetKey = deallocationQueue[operator][strategy].at(i);
            MagnitudeInfo memory info = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];

            // If we've reached a pending deallocation that isn't completable yet,
            // we can stop. Any subsequent modificaitons will also be uncompletable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // The diff is a deallocation. Add to encumbered magnitude. Note that this is a deallocation
            // queue and allocations aren't considered because encumbered magnitude
            // is updated as soon as the allocation is created.
            curEncumberedMagnitude = _addInt128(curEncumberedMagnitude, info.pendingDiff);
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

    /// @inheritdoc IAllocationManager
    function getAllocationDelay(
        address operator
    ) public view returns (bool isSet, uint32 delay) {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        if (info.effectTimestamp != 0 && block.timestamp >= info.effectTimestamp) {
            delay = info.pendingDelay;
        } else {
            delay = info.delay;
        }

        // Check that the operator has a configured delay that has taken effect.
        // This is true if isSet is true OR block.timestamp >= effectTimestamp
        // meaning either a delay has been applied or there is a delay set and
        // the effectTimestamp has been reached
        isSet = info.isSet == true || (info.effectTimestamp != 0 && block.timestamp >= info.effectTimestamp);
        return (isSet, delay);
    }

    /// @inheritdoc IAllocationManager
    function getCurrentDelegatedAndSlashableOperatorShares(
        OperatorSet calldata operatorSet,
        address[] calldata operators,
        IStrategy[] calldata strategies
    ) external view returns (uint256[][] memory, uint256[][] memory) {
        return
            getMinDelegatedAndSlashableOperatorSharesBefore(operatorSet, operators, strategies, uint32(block.timestamp));
    }

    /// @inheritdoc IAllocationManager
    function getMinDelegatedAndSlashableOperatorSharesBefore(
        OperatorSet calldata operatorSet,
        address[] calldata operators,
        IStrategy[] calldata strategies,
        uint32 beforeTimestamp
    ) public view returns (uint256[][] memory, uint256[][] memory) {
        require(beforeTimestamp >= block.timestamp, InvalidTimestamp());
        bytes32 operatorSetKey = _encodeOperatorSet(operatorSet);
        uint256[][] memory delegatedShares = delegation.getOperatorsShares(operators, strategies);
        uint256[][] memory slashableShares = new uint256[][](operators.length);

        for (uint256 i = 0; i < operators.length; ++i) {
            address operator = operators[i];
            slashableShares[i] = new uint256[](strategies.length);
            for (uint256 j = 0; j < strategies.length; ++j) {
                IStrategy strategy = strategies[j];
                MagnitudeInfo memory mInfo = _operatorMagnitudeInfo[operator][strategy][operatorSetKey];
                uint64 slashableMagnitude = mInfo.currentMagnitude;
                if (mInfo.effectTimestamp <= beforeTimestamp) {
                    slashableMagnitude = _addInt128(slashableMagnitude, mInfo.pendingDiff);
                }
                slashableShares[i][j] = delegatedShares[i][j].mulWad(slashableMagnitude).divWad(
                    _maxMagnitudeHistory[operator][strategy].latest()
                );
            }
        }

        return (delegatedShares, slashableShares);
    }
}
