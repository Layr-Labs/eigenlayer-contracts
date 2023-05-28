// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../../contracts/interfaces/IRegistryCoordinator.sol";


contract RegistryCoordinatorMock is IRegistryCoordinator {
    /// @notice Returns the bitmap of the quroums the operator is registered for.
    function operatorIdToQuorumBitmap(bytes32 pubkeyHash) external view returns (uint256){}

    /// @notice Returns the stored id for the specified `operator`.
    function getOperatorId(address operator) external view returns (bytes32){}

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32){}

    /// @notice registers the sender as an operator for the `quorumNumbers` with additional bytes for registry interaction data
    function registerOperator(uint8[] memory quorumNumbers, bytes calldata) external returns (bytes32){}

    /// @notice deregisters the sender with additional bytes for registry interaction data
    function deregisterOperator(bytes calldata) external returns (bytes32){}

     /// @notice Returns the registry at the desired index
    function registries(uint256) external view returns (address){}

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256){}
}