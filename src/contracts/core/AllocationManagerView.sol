// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../mixins/Deprecated_OwnableUpgradeable.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "../permissions/Pausable.sol";
import "../libraries/SlashingLib.sol";
import "../libraries/OperatorSetLib.sol";
import "./AllocationManagerStorage.sol";

/// @notice Non-state mutating view functions, (static) called by the `AllocationManager`.
contract AllocationManagerView is // `AllocationManagerStorage` starts at slot 51.
    AllocationManagerStorage, IAllocationManagerView layout at 51 {
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using Snapshots for Snapshots.DefaultWadHistory;
    using OperatorSetLib for OperatorSet;
    using SlashingLib for uint256;
    using EnumerableSet for *;
    using SafeCast for *;

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
        IStrategy _eigenStrategy,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) AllocationManagerStorage(_delegation, _eigenStrategy, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY) {
        // _disableInitializers();
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

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

    /**
     * @dev Returns the minimum allocated stake at the future block.
     * @param operatorSet The operator set to get the minimum allocated stake for.
     * @param operators The operators to get the minimum allocated stake for.
     * @param strategies The strategies to get the minimum allocated stake for.
     * @param futureBlock The future block to get the minimum allocated stake for.
     */
    function _getMinimumAllocatedStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint32 futureBlock
    ) internal view returns (uint256[][] memory allocatedStake) {
        allocatedStake = new uint256[][](operators.length);
        uint256[][] memory delegatedStake = delegation.getOperatorsShares(operators, strategies);

        for (uint256 i = 0; i < operators.length; i++) {
            address operator = operators[i];

            allocatedStake[i] = new uint256[](strategies.length);

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
                allocatedStake[i][j] = delegatedStake[i][j].mulWad(slashableProportion);
            }
        }
    }

    /// @dev Use safe casting when downcasting to uint64
    function _addInt128(uint64 a, int128 b) internal pure returns (uint64) {
        return uint256(int256(int128(uint128(a)) + b)).toUint64();
    }

    /**
     * @notice Helper function to check if an operator is redistributable from a list of operator sets
     * @param operator The operator to check
     * @param operatorSets The list of operator sets to check
     * @return True if the operator is redistributable from any of the operator sets, false otherwise
     */
    function _isOperatorRedistributable(
        address operator,
        OperatorSet[] memory operatorSets
    ) internal view returns (bool) {
        for (uint256 i = 0; i < operatorSets.length; ++i) {
            if (isOperatorSlashable(operator, operatorSets[i]) && isRedistributingOperatorSet(operatorSets[i])) {
                return true;
            }
        }
        return false;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IAllocationManagerView
    function getOperatorSetCount(
        address avs
    ) external view returns (uint256) {
        return _operatorSets[avs].length();
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatedSets(
        address operator
    ) public view returns (OperatorSet[] memory) {
        uint256 length = allocatedSets[operator].length();

        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        for (uint256 i = 0; i < length; i++) {
            operatorSets[i] = OperatorSetLib.decode(allocatedSets[operator].at(i));
        }

        return operatorSets;
    }

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
    function getAllocation(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) public view returns (Allocation memory) {
        (, Allocation memory allocation) = _getUpdatedAllocation(operator, operatorSet.key(), strategy);

        return allocation;
    }

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
    function getEncumberedMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        (uint64 curEncumberedMagnitude,) = _getFreeAndUsedMagnitude(operator, strategy);
        return curEncumberedMagnitude;
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        (, uint64 curAllocatableMagnitude) = _getFreeAndUsedMagnitude(operator, strategy);
        return curAllocatableMagnitude;
    }

    /// @dev For an operator, returns up-to-date amounts for current encumbered and available
    /// magnitude. Note that these two values will always add up to the operator's max magnitude
    /// for the strategy
    function _getFreeAndUsedMagnitude(
        address operator,
        IStrategy strategy
    ) internal view returns (uint64 curEncumberedMagnitude, uint64 curAllocatableMagnitude) {
        // This method needs to simulate clearing any pending deallocations.
        // This roughly mimics the calculations done in `_clearDeallocationQueue` and
        // `_getUpdatedAllocation`, while operating on a `curEncumberedMagnitude`
        // rather than continually reading/updating state.
        curEncumberedMagnitude = encumberedMagnitude[operator][strategy];

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
        curAllocatableMagnitude = _maxMagnitudeHistory[operator][strategy].latest() - curEncumberedMagnitude;
        return (curEncumberedMagnitude, curAllocatableMagnitude);
    }

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitude(address operator, IStrategy strategy) public view returns (uint64) {
        return _maxMagnitudeHistory[operator][strategy].latest();
    }

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitudes(address[] memory operators, IStrategy strategy) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](operators.length);

        for (uint256 i = 0; i < operators.length; ++i) {
            maxMagnitudes[i] = getMaxMagnitude(operators[i], strategy);
        }

        return maxMagnitudes;
    }

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
    function isMemberOfOperatorSet(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        return _operatorSetMembers[operatorSet.key()].contains(operator);
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (bool) {
        return _operatorSets[operatorSet.avs].contains(operatorSet.id);
    }

    /// @inheritdoc IAllocationManagerView
    function getMembers(
        OperatorSet memory operatorSet
    ) external view returns (address[] memory) {
        return _operatorSetMembers[operatorSet.key()].values();
    }

    /// @inheritdoc IAllocationManagerView
    function getMemberCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetMembers[operatorSet.key()].length();
    }

    /// @inheritdoc IAllocationManagerView
    function getAVSRegistrar(
        address avs
    ) public view returns (IAVSRegistrar) {
        IAVSRegistrar registrar = _avsRegistrar[avs];

        return address(registrar) == address(0) ? IAVSRegistrar(avs) : registrar;
    }

    /// @inheritdoc IAllocationManagerView
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

    /// @inheritdoc IAllocationManagerView
    function getMinimumSlashableStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint32 futureBlock
    ) external view returns (uint256[][] memory slashableStake) {
        slashableStake = _getMinimumAllocatedStake(operatorSet, operators, strategies, futureBlock);

        for (uint256 i = 0; i < operators.length; i++) {
            // If the operator is not slashable by the opSet, all strategies should have a slashable stake of 0
            if (!isOperatorSlashable(operators[i], operatorSet)) {
                for (uint256 j = 0; j < strategies.length; j++) {
                    slashableStake[i][j] = 0;
                }
            }
        }
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatedStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies
    ) external view returns (uint256[][] memory) {
        /// This helper function returns the minimum allocated stake by taking into account deallocations at some `futureBlock`.
        /// We use the block.number, as the `futureBlock`, meaning that no **future** deallocations are considered.
        return _getMinimumAllocatedStake(operatorSet, operators, strategies, uint32(block.number));
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) public view returns (bool) {
        RegistrationStatus memory status = registrationStatus[operator][operatorSet.key()];

        // slashableUntil returns the last block the operator is slashable in so we check for
        // less than or equal to
        return status.registered || block.number <= status.slashableUntil;
    }

    /// @inheritdoc IAllocationManagerView
    function getRedistributionRecipient(
        OperatorSet memory operatorSet
    ) public view returns (address) {
        // Load the redistribution recipient and return it if set, otherwise return the default burn address.
        address redistributionRecipient = _redistributionRecipients[operatorSet.key()];
        return redistributionRecipient == address(0) ? DEFAULT_BURN_ADDRESS : redistributionRecipient;
    }

    /// @inheritdoc IAllocationManagerView
    function isRedistributingOperatorSet(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        return getRedistributionRecipient(operatorSet) != DEFAULT_BURN_ADDRESS;
    }

    /// @inheritdoc IAllocationManagerView
    function getSlashCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _slashIds[operatorSet.key()];
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorRedistributable(
        address operator
    ) external view returns (bool) {
        // Get the registered and allocated sets for the operator.
        // We get both sets, since:
        //    - Upon registration the operator allocation will be pending to a redistributing operator set, and as such not yet in RegisteredSets.
        //    - Upon deregistration the operator is removed from RegisteredSets, but is still allocated.
        OperatorSet[] memory registeredSets = getRegisteredSets(operator);
        OperatorSet[] memory allocatedSets = getAllocatedSets(operator);

        // Check if the operator is redistributable from any of the registered or allocated sets
        return
            _isOperatorRedistributable(operator, registeredSets) || _isOperatorRedistributable(operator, allocatedSets);
    }
}
