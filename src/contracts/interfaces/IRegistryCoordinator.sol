// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

/**
 * @title Interface for a contract that coordinates between various registries for an AVS.
 * @author Layr Labs, Inc.
 */
interface IRegistryCoordinator {
    // DATA STRUCTURES
    enum OperatorStatus
    {
        // default is DEREGISTERED
        DEREGISTERED,
        REGISTERED
    }

    /**
     * @notice Data structure for storing info on operators
     */
    struct Operator {
        // the id of the operator, which is likely the keccak256 hash of the operator's public key if using BLSRegsitry
        bytes32 operatorId;
        // start taskNumber from which the  operator has been registered
        uint32 fromTaskNumber;
        // indicates whether the operator is actively registered for serving the middleware or not
        OperatorStatus status;
    }

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32);

    /// @notice Returns task number from when `operator` has been registered.
    function getFromTaskNumberForOperator(address operator) external view returns (uint32);

    /// @notice Returns the registry at the desired index
    function registries(uint256) external view returns (address);

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256);

    /// @notice registers the sender as an operator for the `quorumNumbers` with additional bytes for registry interaction data
    function registerOperator(bytes memory quorumNumbers, bytes calldata) external;

    /// @notice deregisters the sender with additional bytes for registry interaction data
    function deregisterOperator(bytes calldata) external;
}
