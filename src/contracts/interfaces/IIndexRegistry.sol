// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistry.sol";

/**
 * @title Interface for a `Registry`-type contract that keeps track of an ordered list of operators for up to 256 quroums.
 * @author Layr Labs, Inc.
 */
interface IIndexRegistry is IRegistry {
    // DATA STRUCTURES

    // struct used to give definitive ordering to operators at each blockNumber
    struct OperatorIndex {
        // blockNumber number at which operator index changed
        // note that the operator's index is different *for this block number*, i.e. the *new* index is *inclusive* of this value
        uint32 toBlockNumber;
        // index of the operator in array of operators, or the total number of operators if in the 'totalOperatorsHistory'
        uint32 index;
    }

    /**
     * @notice Registers the operator with the specified `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId is the id of the operator that is being registered
     * @param quorumNumbers is the quorum numbers the operator is registered for
     * @dev access restricted to the RegistryCoordinator
     */
    function registerOperator(bytes32 operatorId, bytes calldata quorumNumbers) external;

    /**
     * @notice Deregisters the operator with the specified `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId is the id of the operator that is being deregistered
     * @param quorumNumbers is the quorum numbers the operator is deregistered for
     * @param quorumToOperatorListIndexes is an array of indexes for each quorum as witnesses for the last operators to swap for each quorum
     * @param globalOperatorListIndex is the index of the operator in the global operator list
     * @dev Permissioned by RegistryCoordinator
     */
    function deregisterOperator(bytes32 operatorId, bytes calldata quorumNumbers, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) external;

    /**
     * @notice Looks up the `operator`'s index for `quorumNumber` at the specified `blockNumber` using the `index`.
     * @param operatorId is the id of the operator for which the index is desired
     * @param quorumNumber is the quorum number for which the operator index is desired
     * @param blockNumber is the block number at which the index of the operator is desired
     * @param index Used to specify the entry within the dynamic array `operatorIdToIndexHistory[operatorId]` to 
     * read data from
     * @dev Function will revert in the event that the specified `index` input does not identify the appropriate entry in the
     * array `operatorIdToIndexHistory[operatorId][quorumNumber]` to pull the info from.
     */
    function getOperatorIndexForQuorumAtBlockNumberByIndex(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber, uint32 index) external view returns (uint32);

    /**
     * @notice Looks up the number of total operators for `quorumNumber` at the specified `blockNumber`.
     * @param quorumNumber is the quorum number for which the total number of operators is desired
     * @param blockNumber is the block number at which the total number of operators is desired
     * @param index is the index of the entry in the dynamic array `totalOperatorsHistory[quorumNumber]` to read data from
     */
    function getTotalOperatorsForQuorumAtBlockNumberByIndex(uint8 quorumNumber, uint32 blockNumber, uint32 index) external view returns (uint32);

    /// @notice Returns the current number of operators of this service for `quorumNumber`.
    function totalOperatorsForQuorum(uint8 quorumNumber) external view returns (uint32);

    /// @notice Returns the current number of operators of this service.
    function totalOperators() external view returns (uint32);

    /// @notice Returns list of current operators of this service.
    function getOperatorIds() external view returns (bytes32[] memory);
}