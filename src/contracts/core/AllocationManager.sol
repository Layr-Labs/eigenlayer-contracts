// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
import "../libraries/OperatorSetLib.sol";
import "./AllocationManagerStorage.sol";

interface IAVS {
    function registerOperator(address operator, uint32[] calldata operatorSetIds, bytes calldata data) external;
    function deregisterOperator(address operator, uint32[] calldata operatorSetIds) external;
}

contract AllocationManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AllocationManagerStorage,
    ReentrancyGuardUpgradeable
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
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManage,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) AllocationManagerStorage(_delegation, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY) {
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

        // Check that the operator set exists and the operator is registered to it
        OperatorSet memory operatorSet = OperatorSet(msg.sender, params.operatorSetId);
        bool isRegistered = _isRegistered(params.operator, operatorSet);
        require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());
        require(isRegistered, InvalidOperator());

        uint256 length = _operatorSetStrategies[operatorSet.key()].length();
        IStrategy[] memory strategiesSlashed = new IStrategy[](length);
        uint256[] memory wadSlashed = new uint256[](length);

        // For each strategy in the operator set, slash any existing allocation
        for (uint256 i = 0; i < length; i++) {
            // 1. Get the operator's allocation info for the strategy and operator set
            IStrategy strategy = IStrategy(_operatorSetStrategies[operatorSet.key()].at(i));
            (StrategyInfo memory info, Allocation memory allocation) =
                _getUpdatedAllocation(params.operator, operatorSet.key(), strategy);
            strategiesSlashed[i] = strategy;

            // 2. Skip if the operator does not have a slashable allocation
            // NOTE: this "if" is equivalent to: `if (!_isAllocationSlashable)`, because the other
            // conditions in this method are already true (isRegistered + operatorSetStrategies.contains)
            if (allocation.currentMagnitude == 0) {
                continue;
            }

            // 3. Calculate the amount of magnitude being slashed, and subtract from
            // the operator's currently-allocated magnitude, as well as the strategy's
            // max and encumbered magnitudes
            uint64 slashedMagnitude = uint64(uint256(allocation.currentMagnitude).mulWadRoundUp(params.wadToSlash));
            wadSlashed[i] = uint256(slashedMagnitude).divWad(info.maxMagnitude);

            allocation.currentMagnitude -= slashedMagnitude;
            info.maxMagnitude -= slashedMagnitude;
            info.encumberedMagnitude -= slashedMagnitude;

            // 4. If there is a pending deallocation, reduce the pending deallocation proportionally.
            // This ensures that when the deallocation is completed, less magnitude is freed.
            if (allocation.pendingDiff < 0) {
                uint64 slashedPending =
                    uint64(uint256(uint128(-allocation.pendingDiff)).mulWadRoundUp(params.wadToSlash));
                allocation.pendingDiff += int128(uint128(slashedPending));

                emit AllocationUpdated(
                    params.operator,
                    operatorSet,
                    strategy,
                    _addInt128(allocation.currentMagnitude, allocation.pendingDiff),
                    allocation.effectBlock
                );
            }

            // 5. Update state
            _updateAllocationInfo(params.operator, operatorSet.key(), strategy, info, allocation);
            _updateMaxMagnitude(params.operator, strategy, info.maxMagnitude);

            // 6. Decrease operators shares in the DelegationManager
            delegation.decreaseOperatorShares({
                operator: params.operator,
                strategy: strategy,
                previousMaxMagnitude: info.maxMagnitude + slashedMagnitude,
                newMaxMagnitude: info.maxMagnitude
            });
        }

        emit OperatorSlashed(params.operator, operatorSet, strategiesSlashed, wadSlashed, params.description);
    }

    /// @inheritdoc IAllocationManager
    function modifyAllocations(
        AllocateParams[] calldata params
    ) external onlyWhenNotPaused(PAUSED_MODIFY_ALLOCATIONS) {
        // Check that the operator exists and has configured an allocation delay
        (bool isSet, uint32 operatorAllocationDelay) = getAllocationDelay(msg.sender);
        require(isSet, UninitializedAllocationDelay());

        for (uint256 i = 0; i < params.length; i++) {
            require(params[i].strategies.length == params[i].newMagnitudes.length, InputArrayLengthMismatch());

            // Check that the operator set exists and get the operator's registration status
            // Operators do not need to be registered for an operator set in order to allocate
            // slashable magnitude to the set. In fact, it is expected that operators will
            // allocate magnitude before registering, as AVS's will likely only accept
            // registrations from operators that are already slashable.
            OperatorSet calldata operatorSet = params[i].operatorSet;
            bool isRegistered = _isRegistered(msg.sender, operatorSet);
            require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());

            for (uint256 j = 0; j < params[i].strategies.length; j++) {
                IStrategy strategy = params[i].strategies[j];

                // 1. If the operator has any pending deallocations for this strategy, clear them
                // to free up magnitude for allocation. Fetch the operator's up to date allocation
                // info and ensure there is no remaining pending modification.
                _clearDeallocationQueue(msg.sender, strategy, type(uint16).max);

                (StrategyInfo memory info, Allocation memory allocation) =
                    _getUpdatedAllocation(msg.sender, operatorSet.key(), strategy);
                require(allocation.pendingDiff == 0, ModificationAlreadyPending());

                // 2. Check whether the operator's allocation is slashable. If not, we allow instant
                // deallocation.
                bool isSlashable = _isAllocationSlashable(operatorSet, strategy, allocation, isRegistered);

                // 3. Calculate the change in magnitude
                allocation.pendingDiff = _calcDelta(allocation.currentMagnitude, params[i].newMagnitudes[j]);
                require(allocation.pendingDiff != 0, SameMagnitude());

                // 4. Handle deallocation/allocation
                if (allocation.pendingDiff < 0) {
                    if (isSlashable) {
                        // If the operator is slashable, deallocated magnitude will be freed after
                        // the deallocation delay. This magnitude remains slashable until then.
                        deallocationQueue[msg.sender][strategy].pushBack(operatorSet.key());

                        allocation.effectBlock = uint32(block.number) + DEALLOCATION_DELAY;
                    } else {
                        // Deallocation immediately updates/frees magnitude if the operator is not slashable
                        info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, allocation.pendingDiff);

                        allocation.currentMagnitude = params[i].newMagnitudes[j];
                        allocation.pendingDiff = 0;
                    }
                } else if (allocation.pendingDiff > 0) {
                    // Allocation immediately consumes available magnitude, but the additional
                    // magnitude does not become slashable until after the allocation delay
                    info.encumberedMagnitude = _addInt128(info.encumberedMagnitude, allocation.pendingDiff);
                    require(info.encumberedMagnitude <= info.maxMagnitude, InsufficientMagnitude());

                    allocation.effectBlock = uint32(block.number) + operatorAllocationDelay;
                }

                // 5. Update state
                _updateAllocationInfo(msg.sender, operatorSet.key(), strategy, info, allocation);
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
        RegisterParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        // Check that the operator exists
        require(delegation.isOperator(msg.sender), InvalidOperator());

        for (uint256 i = 0; i < params.operatorSetIds.length; i++) {
            // Check the operator set exists and the operator is not currently registered to it
            OperatorSet memory operatorSet = OperatorSet(params.avs, params.operatorSetIds[i]);
            require(_operatorSets[operatorSet.avs].contains(operatorSet.id), InvalidOperatorSet());
            require(!_isRegistered(msg.sender, operatorSet), AlreadyMemberOfSet());

            // Add operator to operator set
            registeredSets[msg.sender].add(operatorSet.key());
            _operatorSetMembers[operatorSet.key()].add(msg.sender);
            emit OperatorAddedToOperatorSet(msg.sender, operatorSet);

            // Mark the operator registered
            registrationStatus[msg.sender][operatorSet.key()].registered = true;
        }

        // Call the AVS to complete registration. If the AVS reverts, registration will fail.
        IAVS(params.avs).registerOperator(msg.sender, params.operatorSetIds, params.data);
    }

    /// @inheritdoc IAllocationManager
    function deregisterFromOperatorSets(
        DeregisterParams calldata params
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION) {
        require(msg.sender == params.operator || msg.sender == params.avs, InvalidCaller());

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
                registeredUntil: uint32(block.number) + DEALLOCATION_DELAY
            });
        }

        // Call the AVS to complete deregistration. Even if the AVS reverts, the operator is
        // considered deregistered
        try IAVS(params.avs).deregisterOperator(params.operator, params.operatorSetIds) {} catch {}
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
        CreateSetParams[] calldata params
    ) external {
        for (uint256 i = 0; i < params.length; i++) {
            OperatorSet memory operatorSet = OperatorSet(msg.sender, params[i].operatorSetId);

            // Create the operator set, ensuring it does not already exist
            require(_operatorSets[msg.sender].add(operatorSet.id) == false, InvalidOperatorSet());
            emit OperatorSetCreated(OperatorSet(msg.sender, operatorSet.id));

            // Add strategies to the operator set
            bytes32 operatorSetKey = operatorSet.key();
            for (uint256 j = 0; j < params[i].strategies.length; j++) {
                _operatorSetStrategies[operatorSetKey].add(address(params[i].strategies[j]));
                emit StrategyAddedToOperatorSet(operatorSet, params[i].strategies[j]);
            }
        }
    }

    /// @inheritdoc IAllocationManager
    function addStrategiesToOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(_operatorSets[msg.sender].contains(operatorSet.id), InvalidOperatorSet());

        bytes32 operatorSetKey = operatorSet.key();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(_operatorSetStrategies[operatorSetKey].add(address(strategies[i])), StrategyAlreadyInOperatorSet());
            emit StrategyAddedToOperatorSet(operatorSet, strategies[i]);
        }
    }

    /// @inheritdoc IAllocationManager
    function removeStrategiesFromOperatorSet(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet(msg.sender, operatorSetId);
        require(_operatorSets[msg.sender].contains(operatorSet.id), InvalidOperatorSet());

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
                _getUpdatedAllocation(msg.sender, operatorSetKey, strategy);

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

    function _isRegistered(address operator, OperatorSet memory operatorSet) internal view returns (bool) {
        RegistrationStatus memory status = registrationStatus[operator][operatorSet.key()];

        return status.registered || block.number < status.registeredUntil;
    }

    function _isAllocationSlashable(
        OperatorSet memory operatorSet,
        IStrategy strategy,
        Allocation memory allocation,
        bool isRegistered
    ) internal view returns (bool) {
        // If the operator set does not use this strategy, any allocation from it is not slashable
        if (!_operatorSetStrategies[operatorSet.key()].contains(address(strategy))) {
            return false;
        }

        // If the operator is not registered to the operator set, any allocation is not slashable
        if (!isRegistered) {
            return false;
        }

        // The allocation is not slashable if there is nothing allocated
        if (allocation.currentMagnitude == 0) {
            return false;
        }

        return true;
    }

    /**
     * @dev For an operator set, get the operator's effective allocated magnitude.
     * If the operator set has a pending deallocation that can be completed at the
     * current timestamp, this method returns a view of the allocation as if the deallocation
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
        // Update encumbered magnitude
        encumberedMagnitude[operator][strategy] = info.encumberedMagnitude;
        emit EncumberedMagnitudeUpdated(operator, strategy, info.encumberedMagnitude);

        // Update allocation for this operator set from the strategy
        allocations[operator][operatorSetKey][strategy] = allocation;
        emit AllocationUpdated(
            operator, OperatorSetLib.decode(operatorSetKey), strategy, allocation.currentMagnitude, uint32(block.number)
        );

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
        uint256 length = allocatedStrategies[operator][operatorSet.key()].length();

        IStrategy[] memory strategies = new IStrategy[](length);
        for (uint256 i = 0; i < length; i++) {
            strategies[i] = IStrategy(allocatedStrategies[operator][operatorSet.key()].at(i));
        }

        return strategies;
    }

    /// @inheritdoc IAllocationManager
    function getAllocation(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) external view returns (Allocation memory) {
        (, Allocation memory allocation) = _getUpdatedAllocation(operator, operatorSet.key(), strategy);

        return allocation;
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
            (, Allocation memory allocation) = _getUpdatedAllocation(operator, operatorSet.key(), strategy);

            operatorSets[i] = OperatorSetLib.decode(allocatedSets[operator].at(i));
            _allocations[i] = allocation;
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
            // we can stop. Any subsequent modificaitons will also be uncompletable.
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
    function getMaxMagnitudes(
        address operator,
        IStrategy[] memory strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].latest();
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
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup(blockNumber);
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
        address operator,
        uint256 start,
        uint256 count
    ) public view returns (OperatorSet[] memory) {
        uint256 maxCount = registeredSets[operator].length() - start;
        if (count > maxCount) count = maxCount;

        OperatorSet[] memory operatorSets = new OperatorSet[](count);
        for (uint256 i = 0; i < count; ++i) {
            // forgefmt: disable-next-item
            operatorSets[i] = OperatorSetLib.decode(
                registeredSets[operator].at(start + i)
            );
        }

        return operatorSets;
    }

    /// @inheritdoc IAllocationManager
    function getRegisteredSetCount(
        address operator
    ) external view returns (uint256) {
        return registeredSets[operator].length();
    }

    /// @inheritdoc IAllocationManager
    function getRegisteredSetAtIndex(address operator, uint256 index) external view returns (OperatorSet memory) {
        // forgefmt: disable-next-item
        return OperatorSetLib.decode(
            registeredSets[operator].at(index)
        );
    }

    /// @inheritdoc IAllocationManager
    function getMembers(
        OperatorSet memory operatorSet,
        uint256 start,
        uint256 count
    ) external view returns (address[] memory) {
        uint256 maxCount = _operatorSetMembers[operatorSet.key()].length() - start;
        if (count > maxCount) count = maxCount;

        address[] memory operators = new address[](count);
        for (uint256 i = 0; i < count; ++i) {
            operators[i] = _operatorSetMembers[operatorSet.key()].at(start + i);
        }

        return operators;
    }

    /// @inheritdoc IAllocationManager
    function getMemberCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[operatorSet.key()].length();
    }

    /// @inheritdoc IAllocationManager
    function getMemberAtIndex(OperatorSet memory operatorSet, uint256 index) external view returns (address) {
        return _operatorSetMembers[operatorSet.key()].at(index);
    }

    /// @inheritdoc IAllocationManager
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory) {
        uint256 length = _operatorSetStrategies[operatorSet.key()].length();

        IStrategy[] memory strategies = new IStrategy[](length);
        for (uint256 i = 0; i < length; ++i) {
            strategies[i] = IStrategy(_operatorSetStrategies[operatorSet.key()].at(i));
        }

        return strategies;
    }

    /// @inheritdoc IAllocationManager
    function getCurrentDelegatedAndSlashableOperatorShares(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies
    ) external view returns (uint256[][] memory, uint256[][] memory) {
        return getMinDelegatedAndSlashableOperatorSharesBefore(operatorSet, operators, strategies, uint32(block.number));
    }

    /// @inheritdoc IAllocationManager
    function getMinDelegatedAndSlashableOperatorSharesBefore(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint32 beforeBlock
    ) public view returns (uint256[][] memory, uint256[][] memory) {
        require(beforeBlock >= block.number, InvalidBlockNumber());
        bytes32 operatorSetKey = operatorSet.key();

        uint256[][] memory delegatedShares = delegation.getOperatorsShares(operators, strategies);
        uint256[][] memory slashableShares = new uint256[][](operators.length);

        for (uint256 i = 0; i < operators.length; ++i) {
            address operator = operators[i];
            slashableShares[i] = new uint256[](strategies.length);

            for (uint256 j = 0; j < strategies.length; ++j) {
                IStrategy strategy = strategies[j];
                Allocation memory allocation = allocations[operator][operatorSetKey][strategy];
                uint64 maxMagnitude = _maxMagnitudeHistory[operator][strategy].latest();

                if (beforeBlock >= allocation.effectBlock) {
                    allocation.currentMagnitude = _addInt128(allocation.currentMagnitude, allocation.pendingDiff);
                }

                // forgefmt: disable-next-item
                slashableShares[i][j] = 
                    delegatedShares[i][j]
                        .mulWad(allocation.currentMagnitude)
                        .divWad(maxMagnitude);
            }
        }

        return (delegatedShares, slashableShares);
    }
}
