// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../mixins/PermissionControllerMixin.sol";
import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
import "../libraries/OperatorSetLib.sol";
import "./AllocationManagerStorage.sol";

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable,
    PermissionControllerMixin
{
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using EnumerableSet for *;

    using Snapshots for Snapshots.DefaultWadHistory;
    using OperatorSetLib for OperatorSet;
    using SlashingLib for uint256;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the DelegationManager address, the deallocation delay, and the allocation configuration delay.
     */
    constructor(
        IDelegationManager _delegation,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    )
        AllocationManagerStorage(_delegation, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
        Pausable(_pauserRegistry)
        PermissionControllerMixin(_permissionController)
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
        address avs,
        SlashingParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SLASHING) checkCanCall(avs) {
        // Check that the operator set exists and the operator is registered to it
        OperatorSet memory operatorSet = OperatorSet(avs, params.operatorSetId);
        bool isOperatorSlashable = _isOperatorSlashable(params.operator, operatorSet);
        require(params.strategies.length == params.wadsToSlash.length, InputArrayLengthMismatch());
        require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());
        require(isOperatorSlashable, OperatorNotSlashable());

        uint256[] memory wadSlashed = new uint256[](params.strategies.length);

        // For each strategy in the operator set, slash any existing allocation
        for (uint256 i = 0; i < params.strategies.length; i++) {
            // Check that `strategies` is in ascending order.
            require(
                i == 0 || uint160(address(params.strategies[i])) > uint160(address(params.strategies[i - 1])),
                StrategiesMustBeInAscendingOrder()
            );
            // Check that `wadToSlash` is within acceptable bounds.
            require(0 < params.wadsToSlash[i] && params.wadsToSlash[i] <= WAD, InvalidWadToSlash());
            // Check that the operator set contains the strategy.
            require(
                _operatorSetStrategies[operatorSet.key()].contains(address(params.strategies[i])),
                StrategyNotInOperatorSet()
            );

            // 1. Get the operator's allocation info for the strategy and operator set
            (StrategyInfo memory info, Allocation memory allocation) =
                _getUpdatedAllocation(params.operator, operatorSet.key(), params.strategies[i]);

            // 2. Skip if the operator does not have a slashable allocation
            // NOTE: this "if" is equivalent to: `if (!_isAllocationSlashable)`, because the other
            // conditions in this method are already true (isOperatorSlashable + operatorSetStrategies.contains)
            if (allocation.currentMagnitude == 0) {
                continue;
            }

            // 3. Calculate the amount of magnitude being slashed, and subtract from
            // the operator's currently-allocated magnitude, as well as the strategy's
            // max and encumbered magnitudes
            uint64 slashedMagnitude = uint64(uint256(allocation.currentMagnitude).mulWadRoundUp(params.wadsToSlash[i]));
            uint64 prevMaxMagnitude = info.maxMagnitude;
            wadSlashed[i] = uint256(slashedMagnitude).divWad(info.maxMagnitude);

            allocation.currentMagnitude -= slashedMagnitude;
            info.maxMagnitude -= slashedMagnitude;
            info.encumberedMagnitude -= slashedMagnitude;

            // 4. If there is a pending deallocation, reduce the pending deallocation proportionally.
            // This ensures that when the deallocation is completed, less magnitude is freed.
            if (allocation.pendingDiff < 0) {
                uint64 slashedPending =
                    uint64(uint256(uint128(-allocation.pendingDiff)).mulWadRoundUp(params.wadsToSlash[i]));
                allocation.pendingDiff += int128(uint128(slashedPending));

                emit AllocationUpdated(
                    params.operator,
                    operatorSet,
                    params.strategies[i],
                    _addInt128(allocation.currentMagnitude, allocation.pendingDiff),
                    allocation.effectBlock
                );
            }

            // 5. Update state
            _updateAllocationInfo(params.operator, operatorSet.key(), params.strategies[i], info, allocation);

            // Emit an event for the updated allocation
            emit AllocationUpdated(
                params.operator, operatorSet, params.strategies[i], allocation.currentMagnitude, uint32(block.number)
            );

            _updateMaxMagnitude(params.operator, params.strategies[i], info.maxMagnitude);

            // 6. Decrease and burn operators shares in the DelegationManager
            delegation.burnOperatorShares({
                operator: params.operator,
                strategy: params.strategies[i],
                prevMaxMagnitude: prevMaxMagnitude,
                newMaxMagnitude: info.maxMagnitude
            });
        }

        emit OperatorSlashed(params.operator, operatorSet, params.strategies, wadSlashed, params.description);
    }

    /// @inheritdoc IAllocationManager
    function modifyAllocations(
        address operator,
        AllocateParams[] memory params
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        // Check that the caller is allowed to modify allocations on behalf of the operator
        // We do not use a modifier to avoid `stack too deep` errors
        require(_checkCanCall(operator), InvalidCaller());

        // Check that the operator exists and has configured an allocation delay
        uint32 operatorAllocationDelay;
        {
            (bool isSet, uint32 delay) = getAllocationDelay(operator);
            require(isSet, UninitializedAllocationDelay());
            operatorAllocationDelay = delay;
        }

        for (uint256 i = 0; i < params.length; i++) {
            require(params[i].strategies.length == params[i].newMagnitudes.length, InputArrayLengthMismatch());

            // Check that the operator set exists and get the operator's registration status
            // Operators do not need to be registered for an operator set in order to allocate
            // slashable magnitude to the set. In fact, it is expected that operators will
            // allocate magnitude before registering, as AVS's will likely only accept
            // registrations from operators that are already slashable.
            OperatorSet memory operatorSet = params[i].operatorSet;
            require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());

            bool isOperatorSlashable = _isOperatorSlashable(operator, operatorSet);

            for (uint256 j = 0; j < params[i].strategies.length; j++) {
                IStrategy strategy = params[i].strategies[j];

                // 1. If the operator has any pending deallocations for this strategy, clear them
                // to free up magnitude for allocation. Fetch the operator's up to date allocation
                // info and ensure there is no remaining pending modification.
                _clearDeallocationQueue(operator, strategy, type(uint16).max);

                (StrategyInfo memory info, Allocation memory allocation) =
                    _getUpdatedAllocation(operator, operatorSet.key(), strategy);
                require(allocation.pendingDiff == 0, ModificationAlreadyPending());

                // 2. Check whether the operator's allocation is slashable. If not, we allow instant
                // deallocation.
                bool isSlashable = _isAllocationSlashable(operatorSet, strategy, allocation, isOperatorSlashable);

                // 3. Calculate the change in magnitude
                allocation.pendingDiff = _calcDelta(allocation.currentMagnitude, params[i].newMagnitudes[j]);
                require(allocation.pendingDiff != 0, SameMagnitude());

                // 4. Handle deallocation/allocation
                if (allocation.pendingDiff < 0) {
                    if (isSlashable) {
                        // If the operator is slashable, deallocated magnitude will be freed after
                        // the deallocation delay. This magnitude remains slashable until then.
                        deallocationQueue[operator][strategy].pushBack(operatorSet.key());

                        // deallocations are slashable in the window [block.number, block.number + deallocationDelay]
                        // therefore, the effectBlock is set to the block right after the slashable window
                        allocation.effectBlock = uint32(block.number) + DEALLOCATION_DELAY + 1;
                    } else {
                        // Deallocation immediately updates/frees magnitude if the operator is not slashable
                        info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, allocation.pendingDiff);

                        allocation.currentMagnitude = params[i].newMagnitudes[j];
                        allocation.pendingDiff = 0;
                        allocation.effectBlock = uint32(block.number);
                    }
                } else if (allocation.pendingDiff > 0) {
                    // Allocation immediately consumes available magnitude, but the additional
                    // magnitude does not become slashable until after the allocation delay
                    info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, allocation.pendingDiff);
                    require(info.encumberedMagnitude <= info.maxMagnitude, InsufficientMagnitude());

                    allocation.effectBlock = uint32(block.number) + operatorAllocationDelay;
                }

                // 5. Update state
                _updateAllocationInfo(operator, operatorSet.key(), strategy, info, allocation);

                // 6. Emit an event for the updated allocation
                emit AllocationUpdated(
                    operator,
                    OperatorSetLib.decode(operatorSet.key()),
                    strategy,
                    _addInt128(allocation.currentMagnitude, allocation.pendingDiff),
                    allocation.effectBlock
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

        for (uint256 i = 0; i < strategies.length; ++i) {
            _clearDeallocationQueue({operator: operator, strategy: strategies[i], numToClear: numToClear[i]});
        }
    }

    /// @inheritdoc IAllocationManager
    function registerForOperatorSets(
        address operator,
        RegisterParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) checkCanCall(operator) {
        // Check that the operator exists
        require(delegation.isOperator(operator), InvalidOperator());

        for (uint256 i = 0; i < params.operatorSetIds.length; i++) {
            // Check the operator set exists and the operator is not currently registered to it
            OperatorSet memory operatorSet = OperatorSet(params.avs, params.operatorSetIds[i]);
            require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());
            require(!_isOperatorSlashable(operator, operatorSet), AlreadyMemberOfSet());

            // Add operator to operator set
            registeredSets[operator].add(operatorSet.key());
            _operatorSetMembers[operatorSet.key()].add(operator);
            emit OperatorAddedToOperatorSet(operator, operatorSet);

            // Mark the operator registered
            registrationStatus[operator][operatorSet.key()].registered = true;
        }

        // Call the AVS to complete registration. If the AVS reverts, registration will fail.
        getAVSRegistrar(params.avs).registerOperator(operator, params.operatorSetIds, params.data);
    }

    /// @inheritdoc IAllocationManager
    function deregisterFromOperatorSets(
        DeregisterParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        // Check that the caller is either authorized on behalf of the operator or AVS
        require(_checkCanCall(params.operator) || _checkCanCall(params.avs), InvalidCaller());

        for (uint256 i = 0; i < params.operatorSetIds.length; i++) {
            // Check the operator set exists and the operator is registered to it
            OperatorSet memory operatorSet = OperatorSet(params.avs, params.operatorSetIds[i]);
            require(_operatorSets[params.avs].contains(operatorSet.id), InvalidOperatorSet());
            require(registrationStatus[params.operator][operatorSet.key()].registered, NotMemberOfSet());

            // Remove operator from operator set
            registeredSets[params.operator].remove(operatorSet.key());
            _operatorSetMembers[operatorSet.key()].remove(params.operator);
            emit OperatorRemovedFromOperatorSet(params.operator, operatorSet);

            // Mark operator deregistered until the DEALLOCATION_DELAY passes
            // forgefmt: disable-next-item
            registrationStatus[params.operator][operatorSet.key()] = RegistrationStatus({
                registered: false,
                slashableUntil: uint32(block.number) + DEALLOCATION_DELAY
            });
        }

        // Call the AVS to complete deregistration. Even if the AVS reverts, the operator is
        // considered deregistered
        try getAVSRegistrar(params.avs).deregisterOperator(params.operator, params.operatorSetIds) {} catch {}
    }

    /// @inheritdoc IAllocationManager
    function setAllocationDelay(address operator, uint32 delay) external {
        if (msg.sender != address(delegation)) {
            require(_checkCanCall(operator), InvalidCaller());
            require(delegation.isOperator(operator), InvalidOperator());
        }
        _setAllocationDelay(operator, delay);
    }

    /// @inheritdoc IAllocationManager
    function setAVSRegistrar(address avs, IAVSRegistrar registrar) external checkCanCall(avs) {
        _avsRegistrar[avs] = registrar;
        emit AVSRegistrarSet(avs, getAVSRegistrar(avs));
    }

    /// @inheritdoc IAllocationManager
    function updateAVSMetadataURI(address avs, string calldata metadataURI) external checkCanCall(avs) {
        emit AVSMetadataURIUpdated(avs, metadataURI);
    }

    /// @inheritdoc IAllocationManager
    function createOperatorSets(address avs, CreateSetParams[] calldata params) external checkCanCall(avs) {
        for (uint256 i = 0; i < params.length; i++) {
            OperatorSet memory operatorSet = OperatorSet(avs, params[i].operatorSetId);

            // Create the operator set, ensuring it does not already exist
            require(_operatorSets[avs].add(operatorSet.id), InvalidOperatorSet());
            emit OperatorSetCreated(OperatorSet(avs, operatorSet.id));

            // Add strategies to the operator set
            bytes32 operatorSetKey = operatorSet.key();
            for (uint256 j = 0; j < params[i].strategies.length; j++) {
                _operatorSetStrategies[operatorSetKey].add(address(params[i].strategies[j]));
                emit StrategyAddedToOperatorSet(operatorSet, params[i].strategies[j]);
            }
        }
    }

    /// @inheritdoc IAllocationManager
    function addStrategiesToOperatorSet(
        address avs,
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external checkCanCall(avs) {
        OperatorSet memory operatorSet = OperatorSet(avs, operatorSetId);
        bytes32 operatorSetKey = operatorSet.key();

        require(_operatorSets[avs].contains(operatorSet.id), InvalidOperatorSet());

        for (uint256 i = 0; i < strategies.length; i++) {
            require(_operatorSetStrategies[operatorSetKey].add(address(strategies[i])), StrategyAlreadyInOperatorSet());
            emit StrategyAddedToOperatorSet(operatorSet, strategies[i]);
        }
    }

    /// @inheritdoc IAllocationManager
    function removeStrategiesFromOperatorSet(
        address avs,
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external checkCanCall(avs) {
        OperatorSet memory operatorSet = OperatorSet(avs, operatorSetId);
        require(_operatorSets[avs].contains(operatorSet.id), InvalidOperatorSet());

        bytes32 operatorSetKey = operatorSet.key();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(_operatorSetStrategies[operatorSetKey].remove(address(strategies[i])), StrategyNotInOperatorSet());
            emit StrategyRemovedFromOperatorSet(operatorSet, strategies[i]);
        }
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
            (StrategyInfo memory info, Allocation memory allocation) =
                _getUpdatedAllocation(operator, operatorSetKey, strategy);

            // If we've reached a pending deallocation that isn't completable yet,
            // we can stop. Any subsequent deallocation will also be uncompletable.
            if (block.number < allocation.effectBlock) {
                break;
            }

            // Update state. This completes the deallocation, because `_getUpdatedAllocation`
            // gave us strategy/allocation info as if the deallocation was already completed.
            _updateAllocationInfo(operator, operatorSetKey, strategy, info, allocation);

            // Remove the deallocation from the queue
            deallocationQueue[operator][strategy].popFront();
            ++numCleared;
            --length;
        }
    }

    /**
     * @dev Sets the operator's allocation delay. This is the number of blocks between an operator
     * allocating magnitude to an operator set, and the magnitude becoming slashable.
     * @param operator The operator to set the delay on behalf of.
     * @param delay The allocation delay in blocks.
     */
    function _setAllocationDelay(address operator, uint32 delay) internal {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        // If there is a pending delay that can be applied now, set it
        if (info.effectBlock != 0 && block.number >= info.effectBlock) {
            info.delay = info.pendingDelay;
            info.isSet = true;
        }

        info.pendingDelay = delay;
        info.effectBlock = uint32(block.number) + ALLOCATION_CONFIGURATION_DELAY;

        _allocationDelayInfo[operator] = info;
        emit AllocationDelaySet(operator, delay, info.effectBlock);
    }

    /// @notice returns whether the operator is slashable in the given operator set
    function _isOperatorSlashable(address operator, OperatorSet memory operatorSet) internal view returns (bool) {
        RegistrationStatus memory status = registrationStatus[operator][operatorSet.key()];

        // slashableUntil returns the last block the operator is slashable in so we check for
        // less than or equal to
        return status.registered || block.number <= status.slashableUntil;
    }

    /// @notice returns whether the operator's allocation is slashable in the given operator set
    function _isAllocationSlashable(
        OperatorSet memory operatorSet,
        IStrategy strategy,
        Allocation memory allocation,
        bool isOperatorSlashable
    ) internal view returns (bool) {
        /// forgefmt: disable-next-item
        return 
            // If the operator set does not use this strategy, any allocation from it is not slashable
            _operatorSetStrategies[operatorSet.key()].contains(address(strategy)) &&
            // If the operator is not slashable by the operatorSet, any allocation is not slashable
            isOperatorSlashable &&
            // If there is nothing allocated, the allocation is not slashable
            allocation.currentMagnitude != 0;
    }

    /**
     * @dev For an operator set, get the operator's effective allocated magnitude.
     * If the operator set has a pending deallocation that can be completed at the
     * current block number, this method returns a view of the allocation as if the deallocation
     * was completed.
     * @return info the effective allocated and pending magnitude for the operator set, and
     * the effective encumbered magnitude for all operator sets belonging to this strategy
     */
    function _getUpdatedAllocation(
        address operator,
        bytes32 operatorSetKey,
        IStrategy strategy
    ) internal view returns (StrategyInfo memory, Allocation memory) {
        StrategyInfo memory info = StrategyInfo({
            maxMagnitude: _maxMagnitudeHistory[operator][strategy].latest(),
            encumberedMagnitude: encumberedMagnitude[operator][strategy]
        });

        Allocation memory allocation = allocations[operator][operatorSetKey][strategy];

        // If the pending change can't be completed yet, return as-is
        if (block.number < allocation.effectBlock) {
            return (info, allocation);
        }

        // Otherwise, complete the pending change and return updated info
        allocation.currentMagnitude = _addInt128(allocation.currentMagnitude, allocation.pendingDiff);

        // If the completed change was a deallocation, update used magnitude
        if (allocation.pendingDiff < 0) {
            info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, allocation.pendingDiff);
        }

        allocation.effectBlock = 0;
        allocation.pendingDiff = 0;

        return (info, allocation);
    }

    function _updateAllocationInfo(
        address operator,
        bytes32 operatorSetKey,
        IStrategy strategy,
        StrategyInfo memory info,
        Allocation memory allocation
    ) internal {
        // Update encumbered magnitude if it has changed
        // The mapping should NOT be updated when there is a deallocation on a delay
        if (encumberedMagnitude[operator][strategy] != info.encumberedMagnitude) {
            encumberedMagnitude[operator][strategy] = info.encumberedMagnitude;
            emit EncumberedMagnitudeUpdated(operator, strategy, info.encumberedMagnitude);
        }

        // Update allocation for this operator set from the strategy
        // We emit an `AllocationUpdated` from the `modifyAllocations` and `slashOperator` functions.
        // `clearDeallocationQueue` does not emit an `AllocationUpdated` event since it was
        // emitted when the deallocation was queued
        allocations[operator][operatorSetKey][strategy] = allocation;

        // Note: these no-op if the sets already contain the added values (or do not contain removed ones)
        if (allocation.pendingDiff != 0) {
            // If we have a pending modification, ensure the allocation is in the operator's
            // list of enumerable strategies/sets.
            allocatedStrategies[operator][operatorSetKey].add(address(strategy));
            allocatedSets[operator].add(operatorSetKey);
        } else if (allocation.currentMagnitude == 0) {
            // If we do NOT have a pending modification, and no existing magnitude, remove the
            // allocation from the operator's lists.
            allocatedStrategies[operator][operatorSetKey].remove(address(strategy));

            if (allocatedStrategies[operator][operatorSetKey].length() == 0) {
                allocatedSets[operator].remove(operatorSetKey);
            }
        }
    }

    function _updateMaxMagnitude(address operator, IStrategy strategy, uint64 newMaxMagnitude) internal {
        _maxMagnitudeHistory[operator][strategy].push({key: uint32(block.number), value: newMaxMagnitude});
        emit MaxMagnitudeUpdated(operator, strategy, newMaxMagnitude);
    }

    function _calcDelta(uint64 currentMagnitude, uint64 newMagnitude) internal pure returns (int128) {
        return int128(uint128(newMagnitude)) - int128(uint128(currentMagnitude));
    }

    function _addInt128(uint64 a, int128 b) internal pure returns (uint64) {
        return uint64(uint128(int128(uint128(a)) + b));
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IAllocationManager
    function getOperatorSetCount(
        address avs
    ) external view returns (uint256) {
        return _operatorSets[avs].length();
    }

    /// @inheritdoc IAllocationManager
    function getAllocatedSets(
        address operator
    ) external view returns (OperatorSet[] memory) {
        uint256 length = allocatedSets[operator].length();

        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        for (uint256 i = 0; i < length; i++) {
            operatorSets[i] = OperatorSetLib.decode(allocatedSets[operator].at(i));
        }

        return operatorSets;
    }

    /// @inheritdoc IAllocationManager
    function getAllocatedStrategies(
        address operator,
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory) {
        address[] memory values = allocatedStrategies[operator][operatorSet.key()].values();
        IStrategy[] memory strategies;

        assembly {
            strategies := values
        }

        return strategies;
    }

    /// @inheritdoc IAllocationManager
    function getAllocation(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) public view returns (Allocation memory) {
        (, Allocation memory allocation) = _getUpdatedAllocation(operator, operatorSet.key(), strategy);

        return allocation;
    }

    /// @inheritdoc IAllocationManager
    function getAllocations(
        address[] memory operators,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) external view returns (Allocation[] memory) {
        Allocation[] memory _allocations = new Allocation[](operators.length);

        for (uint256 i = 0; i < operators.length; i++) {
            _allocations[i] = getAllocation(operators[i], operatorSet, strategy);
        }

        return _allocations;
    }

    /// @inheritdoc IAllocationManager
    function getStrategyAllocations(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, Allocation[] memory) {
        uint256 length = allocatedSets[operator].length();

        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        Allocation[] memory _allocations = new Allocation[](length);

        for (uint256 i = 0; i < length; i++) {
            OperatorSet memory operatorSet = OperatorSetLib.decode(allocatedSets[operator].at(i));

            operatorSets[i] = operatorSet;
            _allocations[i] = getAllocation(operator, operatorSet, strategy);
        }

        return (operatorSets, _allocations);
    }

    /// @inheritdoc IAllocationManager
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        // This method needs to simulate clearing any pending deallocations.
        // This roughly mimics the calculations done in `_clearDeallocationQueue` and
        // `_getUpdatedAllocation`, while operating on a `curEncumberedMagnitude`
        // rather than continually reading/updating state.
        uint64 curEncumberedMagnitude = encumberedMagnitude[operator][strategy];

        uint256 length = deallocationQueue[operator][strategy].length();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 operatorSetKey = deallocationQueue[operator][strategy].at(i);
            Allocation memory allocation = allocations[operator][operatorSetKey][strategy];

            // If we've reached a pending deallocation that isn't completable yet,
            // we can stop. Any subsequent modifications will also be uncompletable.
            if (block.number < allocation.effectBlock) {
                break;
            }

            // The diff is a deallocation. Add to encumbered magnitude. Note that this is a deallocation
            // queue and allocations aren't considered because encumbered magnitude
            // is updated as soon as the allocation is created.
            curEncumberedMagnitude = _addInt128(curEncumberedMagnitude, allocation.pendingDiff);
        }

        // The difference between the operator's max magnitude and its encumbered magnitude
        // is the magnitude that can be allocated.
        return _maxMagnitudeHistory[operator][strategy].latest() - curEncumberedMagnitude;
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitude(address operator, IStrategy strategy) public view returns (uint64) {
        return _maxMagnitudeHistory[operator][strategy].latest();
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitudes(
        address operator,
        IStrategy[] memory strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = getMaxMagnitude(operator, strategies[i]);
        }

        return maxMagnitudes;
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitudes(address[] memory operators, IStrategy strategy) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](operators.length);

        for (uint256 i = 0; i < operators.length; ++i) {
            maxMagnitudes[i] = getMaxMagnitude(operators[i], strategy);
        }

        return maxMagnitudes;
    }

    /// @inheritdoc IAllocationManager
    function getMaxMagnitudesAtBlock(
        address operator,
        IStrategy[] memory strategies,
        uint32 blockNumber
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup({key: blockNumber});
        }

        return maxMagnitudes;
    }

    /// @inheritdoc IAllocationManager
    function getAllocationDelay(
        address operator
    ) public view returns (bool, uint32) {
        AllocationDelayInfo memory info = _allocationDelayInfo[operator];

        uint32 delay = info.delay;
        bool isSet = info.isSet;

        // If there is a pending delay that can be applied, apply it
        if (info.effectBlock != 0 && block.number >= info.effectBlock) {
            delay = info.pendingDelay;
            isSet = true;
        }

        return (isSet, delay);
    }

    /// @inheritdoc IAllocationManager
    function getRegisteredSets(
        address operator
    ) public view returns (OperatorSet[] memory) {
        uint256 length = registeredSets[operator].length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);

        for (uint256 i = 0; i < length; ++i) {
            operatorSets[i] = OperatorSetLib.decode(registeredSets[operator].at(i));
        }

        return operatorSets;
    }

    /// @inheritdoc IAllocationManager
    function isMemberOfOperatorSet(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetMembers[operatorSet.key()].contains(operator);
    }

    /// @inheritdoc IAllocationManager
    function isOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (bool) {
        return _operatorSets[operatorSet.avs].contains(operatorSet.id);
    }

    /// @inheritdoc IAllocationManager
    function getMembers(
        OperatorSet memory operatorSet
    ) external view returns (address[] memory) {
        return _operatorSetMembers[operatorSet.key()].values();
    }

    /// @inheritdoc IAllocationManager
    function getMemberCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[operatorSet.key()].length();
    }

    /// @inheritdoc IAllocationManager
    function getAVSRegistrar(
        address avs
    ) public view returns (IAVSRegistrar) {
        IAVSRegistrar registrar = _avsRegistrar[avs];

        return address(registrar) == address(0) ? IAVSRegistrar(avs) : registrar;
    }

    /// @inheritdoc IAllocationManager
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory) {
        address[] memory values = _operatorSetStrategies[operatorSet.key()].values();
        IStrategy[] memory strategies;

        assembly {
            strategies := values
        }

        return strategies;
    }

    /// @inheritdoc IAllocationManager
    function getMinimumSlashableStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint32 futureBlock
    ) external view returns (uint256[][] memory slashableStake) {
        slashableStake = new uint256[][](operators.length);
        uint256[][] memory delegatedStake = delegation.getOperatorsShares(operators, strategies);

        for (uint256 i = 0; i < operators.length; i++) {
            address operator = operators[i];
            slashableStake[i] = new uint256[](strategies.length);

            for (uint256 j = 0; j < strategies.length; j++) {
                IStrategy strategy = strategies[j];

                // Fetch the max magnitude and allocation for the operator/strategy.
                // Prevent division by 0 if needed. This mirrors the "FullySlashed" checks
                // in the DelegationManager
                uint64 maxMagnitude = _maxMagnitudeHistory[operator][strategy].latest();
                if (maxMagnitude == 0) {
                    continue;
                }

                Allocation memory alloc = getAllocation(operator, operatorSet, strategy);

                // If the pending change takes effect before `futureBlock`, include it in `currentMagnitude`
                // However, ONLY include the pending change if it is a deallocation, since this method
                // is supposed to return the minimum slashable stake between now and `futureBlock`
                if (alloc.effectBlock <= futureBlock && alloc.pendingDiff < 0) {
                    alloc.currentMagnitude = _addInt128(alloc.currentMagnitude, alloc.pendingDiff);
                }

                uint256 slashableProportion = uint256(alloc.currentMagnitude).divWad(maxMagnitude);
                slashableStake[i][j] = delegatedStake[i][j].mulWad(slashableProportion);
            }
        }
    }
}
