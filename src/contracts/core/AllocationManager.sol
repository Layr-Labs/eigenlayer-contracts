// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
import "../libraries/OperatorSetLib.sol";
import "./AllocationManagerStorage.sol";

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable
{
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Snapshots for Snapshots.DefaultWadHistory;
    using SlashingLib for uint256;
    using OperatorSetLib for OperatorSet;

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

    /// @inheritdoc IAllocationManager
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
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, id: params.operatorSetId});
        require(isOperatorSlashable(params.operator, operatorSet), InvalidOperator());

        // Record the proportion of 1e18 that the operator's total shares that are being slashed
        uint256[] memory wadSlashed = new uint256[](params.strategies.length);

        for (uint256 i = 0; i < params.strategies.length; ++i) {
            PendingMagnitudeInfo memory info =
                _getPendingMagnitudeInfo(params.operator, params.strategies[i], operatorSet.key());

            require(info.currentMagnitude > 0, OperatorNotAllocated());

            // 1. Calculate slashing amount and update current/encumbered magnitude
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
                operatorSetKey: operatorSet.key(),
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
                previousTotalMagnitude: maxMagnitudeBeforeSlash,
                newTotalMagnitude: maxMagnitudeAfterSlash
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

            // 1. For the given (operator,strategy) complete any pending deallocation to free up encumberedMagnitude
            _clearDeallocationQueue({operator: msg.sender, strategy: allocation.strategy, numToClear: type(uint16).max});

            // 2. Check current totalMagnitude matches expected value. This is to check for slashing race conditions
            // where an operator gets slashed from an operatorSet and as a result all the configured allocations have larger
            // proprtional magnitudes relative to each other.
            uint64 maxMagnitude = _maxMagnitudeHistory[msg.sender][allocation.strategy].latest();
            require(maxMagnitude == allocation.expectedMaxMagnitude, InvalidExpectedTotalMagnitude());

            for (uint256 j = 0; j < allocation.operatorSets.length; ++j) {
                OperatorSet calldata operatorSet = allocation.operatorSets[j];
                require(isOperatorSet[operatorSet.avs][operatorSet.id], InvalidOperatorSet());

                bytes32 operatorSetKey = allocation.operatorSets[j].key();

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

    /// @inheritdoc IAllocationManager
    function createOperatorSets(
        uint32[] calldata operatorSetIds
    ) external {
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            require(!isOperatorSet[msg.sender][operatorSetIds[i]], InvalidOperatorSet());
            isOperatorSet[msg.sender][operatorSetIds[i]] = true;
            emit OperatorSetCreated(OperatorSet({avs: msg.sender, id: operatorSetIds[i]}));
        }
    }

    // TODO: Refactor migrations
    // /// @inheritdoc IAllocationManager
    // function migrateOperatorsToOperatorSets(
    //     address[] calldata operators,
    //     uint32[][] calldata operatorSetIds
    // ) external override onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
    //     // Assert that the AVS is an operator set AVS.
    //     require(allocationManager.isOperatorSetAVS(msg.sender), InvalidAVS());

    //     for (uint256 i = 0; i < operators.length; ++i) {
    //         // Assert that the operator is registered & has not been migrated.
    //         require(
    //             avsOperatorStatus[msg.sender][operators[i]] == OperatorAVSRegistrationStatus.REGISTERED,
    //             InvalidOperator()
    //         );

    //         // Migrate operator to operator sets.
    //         _registerToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);

    //         // Deregister operator from AVS - this prevents the operator from being migrated again since
    //         // the AVS can no longer use the legacy M2 registration path
    //         _deregisterOperatorFromAVS(msg.sender, operator);

    //         emit OperatorMigratedToOperatorSets(operators[i], msg.sender, operatorSetIds[i]);
    //     }
    // }

    /// @inheritdoc IAllocationManager
    function addOperatorToSets(
        address operator,
        uint32[] calldata operatorSetIds
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        // Check operator exists in delegationManager
        require(delegation.isOperator(operator), OperatorNotRegisteredToEigenLayer());

        _registerToOperatorSets(operator, msg.sender, operatorSetIds);
    }

    /// @inheritdoc IAllocationManager
    function removeOperatorFromSets(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        require(msg.sender == operator || msg.sender == avs, InvalidCaller());

        _deregisterFromOperatorSets(avs, operator, operatorSetIds);
    }

    /// @inheritdoc IAllocationManager
    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(isOperatorSet[msg.sender][operatorSetId], InvalidOperatorSet());

        bytes32 operatorSetKey = operatorSet.key();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(_operatorSetStrategies[operatorSetKey].add(address(strategies[i])), StrategyAlreadyInOperatorSet());
            emit StrategyAddedToOperatorSet(operatorSet, strategies[i]);
        }
    }

    /// @inheritdoc IAllocationManager
    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(isOperatorSet[msg.sender][operatorSetId], InvalidOperatorSet());

        bytes32 operatorSetKey = operatorSet.key();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(_operatorSetStrategies[operatorSetKey].remove(address(strategies[i])), StrategyNotInOperatorSet());
            emit StrategyRemovedFromOperatorSet(operatorSet, strategies[i]);
        }
    }

    /**
     * @dev Clear one or more pending deallocations to a strategy's allocated magnitude
     * @param operator the operator whose pending deallocations will be cleared
     * @param strategy the strategy to update
     * @param numToClear the number of pending deallocations to complete
     */
    function _clearDeallocationQueue(address operator, IStrategy strategy, uint16 numToClear) internal {
        uint256 numCompleted;
        uint256 length = deallocationQueue[operator][strategy].length();

        while (length > 0 && numCompleted < numToClear) {
            bytes32 operatorSetKey = deallocationQueue[operator][strategy].front();
            PendingMagnitudeInfo memory info = _getPendingMagnitudeInfo(operator, strategy, operatorSetKey);

            // If we've reached a pending deallocation that isn't completable yet,
            // we can stop. Any subsequent deallocation will also be uncompletable.
            if (block.timestamp < info.effectTimestamp) {
                break;
            }

            // Update the operator's allocation in storage
            _updateMagnitudeInfo(operator, strategy, operatorSetKey, info);

            // Remove the deallocation from the queue
            deallocationQueue[operator][strategy].popFront();
            ++numCompleted;
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
        require(delay != 0, InvalidAllocationDelay());

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
     * @dev For an operator set, get the operator's effective allocated magnitude.
     * If the operator set has a pending deallocation that can be completed at the
     * current timestamp, this method returns a view of the allocation as if the deallocation
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
        info.encumberedMagnitude = _encumberedMagnitude;
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
     * @notice Helper function used by migration & registration functions to register an operator to operator sets.
     * @param avs The AVS that the operator is registering to.
     * @param operator The operator to register.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _registerToOperatorSets(address operator, address avs, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and register `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);
            require(isOperatorSet[avs][operatorSetIds[i]], InvalidOperatorSet());

            bytes32 operatorSetKey = operatorSet.key();

            _memberOfSets[operator].add(operatorSetKey);
            _operatorSetMembers[operatorSetKey].add(operator);

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];

            require(!registrationStatus.registered, InvalidOperator());
            registrationStatus.registered = true;

            emit OperatorAddedToOperatorSet(operator, operatorSet);
        }
    }

    /**
     * @notice Internal function to deregister an operator from an operator set.
     *
     * @param avs The AVS that the operator is deregistering from.
     * @param operator The operator to deregister.
     * @param operatorSetIds The IDs of the operator sets.
     */
    function _deregisterFromOperatorSets(address avs, address operator, uint32[] calldata operatorSetIds) internal {
        // Loop over `operatorSetIds` array and deregister `operator` for each item.
        for (uint256 i = 0; i < operatorSetIds.length; ++i) {
            OperatorSet memory operatorSet = OperatorSet(avs, operatorSetIds[i]);
            require(isOperatorSet[avs][operatorSetIds[i]], InvalidOperatorSet());

            bytes32 operatorSetKey = operatorSet.key();

            _memberOfSets[operator].remove(operatorSetKey);
            _operatorSetMembers[operatorSetKey].remove(operator);

            OperatorSetRegistrationStatus storage registrationStatus =
                operatorSetStatus[avs][operator][operatorSetIds[i]];

            require(registrationStatus.registered, InvalidOperator());
            registrationStatus.registered = false;
            registrationStatus.lastDeregisteredTimestamp = uint32(block.timestamp);

            emit OperatorRemovedFromOperatorSet(operator, operatorSet);
        }
    }

    function _calcDelta(uint64 currentMagnitude, uint64 newMagnitude) internal pure returns (int128) {
        return int128(uint128(newMagnitude)) - int128(uint128(currentMagnitude));
    }

    function _addInt128(uint64 a, int128 b) internal pure returns (uint64) {
        return uint64(uint128(int128(uint128(a)) + b));
    }

    /// @inheritdoc IAllocationManager
    function getAllocationInfo(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, MagnitudeInfo[] memory) {
        OperatorSet[] memory operatorSets = getOperatorSetsOfOperator(operator, 0, type(uint256).max);
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
                operatorSetKey: operatorSets[i].key()
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
                    operatorSetKey: operatorSet.key()
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

    /// @inheritdoc IAllocationManager
    function getMinDelegatedAndSlashableOperatorShares(
        OperatorSet calldata operatorSet,
        address[] calldata operators,
        IStrategy[] calldata strategies,
        uint32 beforeTimestamp
    ) external view returns (uint256[][] memory, uint256[][] memory) {
        require(beforeTimestamp > block.timestamp, InvalidTimestamp());
        bytes32 operatorSetKey = operatorSet.key();
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

    /// @inheritdoc IAllocationManager
    function operatorSetsMemberOfAtIndex(address operator, uint256 index) external view returns (OperatorSet memory) {
        return OperatorSetLib.decode(_memberOfSets[operator].at(index));
    }

    /// @inheritdoc IAllocationManager
    function operatorSetMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address) {
        return _operatorSetMembers[operatorSet.key()].at(index);
    }

    /// @inheritdoc IAllocationManager
    function inTotalOperatorSets(
        address operator
    ) external view returns (uint256) {
        return _memberOfSets[operator].length();
    }

    /// @inheritdoc IAllocationManager
    function getOperatorSetsOfOperator(
        address operator,
        uint256 start,
        uint256 length
    ) public view returns (OperatorSet[] memory operatorSets) {
        uint256 maxLength = _memberOfSets[operator].length() - start;
        if (length > maxLength) length = maxLength;
        operatorSets = new OperatorSet[](length);
        for (uint256 i; i < length; ++i) {
            operatorSets[i] = OperatorSetLib.decode(_memberOfSets[operator].at(start + i));
        }
    }

    /// @inheritdoc IAllocationManager
    function getOperatorsInOperatorSet(
        OperatorSet memory operatorSet,
        uint256 start,
        uint256 length
    ) external view returns (address[] memory operators) {
        bytes32 operatorSetKey = operatorSet.key();
        uint256 maxLength = _operatorSetMembers[operatorSetKey].length() - start;
        if (length > maxLength) length = maxLength;
        operators = new address[](length);
        for (uint256 i; i < length; ++i) {
            operators[i] = _operatorSetMembers[operatorSetKey].at(start + i);
        }
    }

    /// @inheritdoc IAllocationManager
    function getNumOperatorsInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[operatorSet.key()].length();
    }

    /// @inheritdoc IAllocationManager
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory strategies) {
        bytes32 operatorSetKey = operatorSet.key();
        uint256 length = _operatorSetStrategies[operatorSetKey].length();

        strategies = new IStrategy[](length);
        for (uint256 i; i < length; ++i) {
            strategies[i] = IStrategy(_operatorSetStrategies[operatorSetKey].at(i));
        }
    }

    /// @inheritdoc IAllocationManager
    function isMember(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _memberOfSets[operator].contains(operatorSet.key());
    }

    /// @inheritdoc IAllocationManager
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        if (isMember(operator, operatorSet)) return true;

        OperatorSetRegistrationStatus memory status = operatorSetStatus[operatorSet.avs][operator][operatorSet.id];

        return block.timestamp < status.lastDeregisteredTimestamp + DEALLOCATION_DELAY;
    }
}
