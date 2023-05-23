// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistry.sol";

/**
 * @title Interface for a `Registry`-type contract that uses either 1 or 2 quorums.
 * @author Layr Labs, Inc.
 * @notice This contract does not currently support n-quorums where n >= 3.
 * Note in particular the presence of only `firstQuorumStake` and `secondQuorumStake` in the `OperatorStake` struct.
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
     * @param quorumBitmap is the bitmap of the quorums the operator is registered for
     * @dev Permissioned by RegistryCoordinator
     */
    function registerOperator(bytes32 operatorId, uint8 quorumBitmap) external;

    /**
     * @notice Deregisters the operator with the specified `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId is the id of the operator that is being deregistered
     * @param quorumBitmap is the bitmap of the quorums the operator is deregistered for
     * @param indexes is an array of indexes for each quorum as witnesses for where to remove the operator from the quorum
     * @dev Permissioned by RegistryCoordinator
     */
    function deregisterOperator(bytes32 operatorId, uint8 quorumBitmap, uint32[] memory indexes) external;

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
}