// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

/**
 * @title Interface for a contract that coordinates between various registries for an AVS.
 * @author Layr Labs, Inc.
 */
interface IRegistryCoordinator {
    // EVENTS
    /// Emits when an operator is registered
    event OperatorRegistered(address indexed operator, bytes32 indexed operatorId);

    /// Emits when an operator is deregistered
    event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId);
    
    // DATA STRUCTURES
    enum OperatorStatus
    {
        // default is NEVER_REGISTERED
        NEVER_REGISTERED,
        REGISTERED,
        DEREGISTERED
    }

    // STRUCTS

    /**
     * @notice Data structure for storing info on operators
     */
    struct Operator {
        // the id of the operator, which is likely the keccak256 hash of the operator's public key if using BLSRegsitry
        bytes32 operatorId;
        // indicates whether the operator is actively registered for serving the middleware or not
        OperatorStatus status;
    }

    /**
     * @notice Data structure for storing info on quorum bitmap updates where the `quorumBitmap` is the bitmap of the 
     * quorums the operator is registered for starting at (inclusive)`updateBlockNumber` and ending at (exclusive) `nextUpdateBlockNumber`
     * @dev nextUpdateBlockNumber is initialized to 0 for the latest update
     */
    struct QuorumBitmapUpdate {
        uint32 updateBlockNumber;
        uint32 nextUpdateBlockNumber;
        uint192 quorumBitmap;
    }

    /// @notice Returns the operator struct for the given `operator`
    function getOperator(address operator) external view returns (Operator memory);

    /// @notice Returns the operatorId for the given `operator`
    function getOperatorId(address operator) external view returns (bytes32);

    /// @notice Returns the status for the given `operator`
    function getOperatorStatus(address operator) external view returns (IRegistryCoordinator.OperatorStatus);

    /// @notice Returns the indices of the quorumBitmaps for the provided `operatorIds` at the given `blockNumber`
    function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] memory operatorIds) external view returns (uint32[] memory);

    /**
     * @notice Returns the quorum bitmap for the given `operatorId` at the given `blockNumber` via the `index`
     * @dev reverts if `index` is incorrect 
     */ 
    function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external view returns (uint192);

    /// @notice Returns the `index`th entry in the operator with `operatorId`'s bitmap history
    function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) external view returns (QuorumBitmapUpdate memory);

    /// @notice Returns the current quorum bitmap for the given `operatorId`
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external view returns (uint192);

    /// @notice Returns the length of the quorum bitmap history for the given `operatorId`
    function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) external view returns (uint256);

    /// @notice Returns the registry at the desired index
    function registries(uint256) external view returns (address);

    /// @notice Returns the number of registries
    function numRegistries() external view returns (uint256);

    /**
     * @notice Registers msg.sender as an operator with the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registering for
     * @param registrationData is the data that is decoded to get the operator's registration information
     */
    function registerOperatorWithCoordinator(bytes memory quorumNumbers, bytes calldata registrationData) external;

    /**
     * @notice Deregisters the msg.sender as an operator from the middleware
     * @param quorumNumbers are the bytes representing the quorum numbers that the operator is registered for
     * @param deregistrationData is the the data that is decoded to get the operator's deregistration information
     */
    function deregisterOperatorWithCoordinator(bytes calldata quorumNumbers, bytes calldata deregistrationData) external;
}