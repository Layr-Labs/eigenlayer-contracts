// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../../contracts/interfaces/IRegistryCoordinator.sol";


contract RegistryCoordinatorMock is IRegistryCoordinator {
    /// @notice Returns the bitmap of the quroums the operator is registered for.
    function operatorIdToQuorumBitmap(bytes32 pubkeyHash) external view returns (uint256){}

    function getOperator(address operator) external view returns (Operator memory){}

    /// @notice Returns the stored id for the specified `operator`.
    function getOperatorId(address operator) external view returns (bytes32){}

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32){}

    function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] memory operatorIds) external view returns (uint32[] memory){}

    /// @notice Returns the quorum bitmap for the given `operatorId` at the given `blockNumber` via the `index`
    function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external view returns (uint192) {}

    /// @notice Returns the current quorum bitmap for the given `operatorId`
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external view returns (uint192) {}

    function numRegistries() external view returns (uint256){}

    function registries(uint256) external view returns (address){}

    function registerOperatorWithCoordinator(bytes memory quorumNumbers, bytes calldata) external {}

    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata) external {}

}
