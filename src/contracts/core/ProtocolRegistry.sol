// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "./storage/ProtocolRegistryStorage.sol";

contract ProtocolRegistry is Initializable, ProtocolRegistryStorage {
    constructor(
        IAllocationManager _allocationManager
    ) ProtocolRegistryStorage(_allocationManager) {
        _disableInitializers();
    }

    /// @inheritdoc IAllocationManagerView
    function getOperatorSetCount(
        address avs
    ) external view returns (uint256) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatedSets(
        address operator
    ) external view returns (OperatorSet[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatedStrategies(
        address operator,
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocation(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) external view returns (Allocation memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocations(
        address[] memory operators,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) external view returns (Allocation[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getStrategyAllocations(
        address operator,
        IStrategy strategy
    ) external view returns (OperatorSet[] memory, Allocation[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getEncumberedMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatableMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitudes(
        address[] calldata operators,
        IStrategy strategy
    ) external view returns (uint64[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMaxMagnitudesAtBlock(
        address operator,
        IStrategy[] calldata strategies,
        uint32 blockNumber
    ) external view returns (uint64[] memory) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocationDelay(
        address operator
    ) external view returns (bool isSet, uint32 delay) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getRegisteredSets(
        address operator
    ) external view returns (OperatorSet[] memory operatorSets) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function isMemberOfOperatorSet(address operator, OperatorSet memory operatorSet) external view returns (bool) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (bool) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMembers(
        OperatorSet memory operatorSet
    ) external view returns (address[] memory operators) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMemberCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAVSRegistrar(
        address avs
    ) external view returns (IAVSRegistrar) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getStrategiesInOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (IStrategy[] memory strategies) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getMinimumSlashableStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint32 futureBlock
    ) external view returns (uint256[][] memory slashableStake) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getAllocatedStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies
    ) external view returns (uint256[][] memory slashableStake) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getRedistributionRecipient(
        OperatorSet memory operatorSet
    ) external view returns (address) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function isRedistributingOperatorSet(
        OperatorSet memory operatorSet
    ) external view returns (bool) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function getSlashCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        _staticcall(address(allocationManager));
    }

    /// @inheritdoc IAllocationManagerView
    function isOperatorRedistributable(
        address operator
    ) external view returns (bool) {
        _staticcall(address(allocationManager));
    }

    /// @dev Makes a static call to `target`, with `data`.
    /// Bubbles up the revert if the call fails.
    /// Otherwise, copies the returndata and returns.
    function _staticcall(
        address target
    ) internal view {
        bytes calldata data = msg.data;
        /// @solidity memory-safe-assembly
        assembly {
            // staticcall(gas,address,argsOffset,argsSize,retOffset,retSize)
            if iszero(staticcall(gas(), target, data.offset, data.length, 0x00, 0x00)) {
                // Bubble up the revert if the call fails.
                returndatacopy(0x00, 0x00, returndatasize())
                revert(0x00, returndatasize())
            }
            // Otherwise, copy the returndata and return.
            returndatacopy(0x00, 0x00, returndatasize())
            return(0, returndatasize())
        }
    }
}
