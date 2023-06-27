// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistry.sol";

/**
 * @title Interface for a `Registry`-type contract that keeps track of an ordered list of operators for up to 256 quorums.
 * @author Layr Labs, Inc.
 */
interface IIndexRegistry is IRegistry {
    // EVENTS
    // emitted when an operator's index in the orderd operator list for the quorum with number `quorumNumber` is updated
    event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newIndex);
    // emitted when an operator's index in the global operator list is updated
    event GlobalIndexUpdate(bytes32 indexed operatorId, uint32 newIndex);

    // DATA STRUCTURES

    // struct used to give definitive ordering to operators at each blockNumber. 
    // NOTE: this struct is slightly abused for also storing the total number of operators for each quorum over time
    struct OperatorIndexUpdate {
        // blockNumber number at which operator index changed
        // note that the operator's index is different *for this block number*, i.e. the *new* index is *inclusive* of this value
        uint32 toBlockNumber;
        // index of the operator in array of operators, or the total number of operators if in the 'totalOperatorsHistory'
        uint32 index;
    }

    /**
     * @notice Registers the operator with the specified `operatorId` for the quorums specified by `quorumNumbers`.
     * @param operatorId is the id of the operator that is being registered
     * @param quorumNumbers is the quorum numbers the operator is registered for
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already registered
     */
    function registerOperator(bytes32 operatorId, bytes calldata quorumNumbers) external;

    /**
     * @notice Deregisters the operator with the specified `operatorId` for the quorums specified by `quorumNumbers`.
     * @param operatorId is the id of the operator that is being deregistered
     * @param completeDeregistration Whether the operator is deregistering from all quorums or just some.
     * @param quorumNumbers is the quorum numbers the operator is deregistered for
     * @param operatorIdsToSwap is the list of operatorIds that have the largest indexes in each of the `quorumNumbers`
     * they will be swapped the operators current index
     * @param globalOperatorListIndex is the index of the operator that is to be removed from the list
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already deregistered
     *         5) `quorumNumbers` is the same as the parameter use when registering
     */
    function deregisterOperator(bytes32 operatorId, bool completeDeregistration, bytes calldata quorumNumbers, bytes32[] memory operatorIdsToSwap, uint32 globalOperatorListIndex) external;

    /// @notice Returns the _operatorIdToIndexHistory entry for the specified `operatorId` and `quorumNumber` at the specified `index`
    function getOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(bytes32 operatorId, uint8 quorumNumber, uint32 index) external view returns (OperatorIndexUpdate memory);

    /// @notice Returns the _totalOperatorsHistory entry for the specified `quorumNumber` at the specified `index`
    function getTotalOperatorsUpdateForQuorumAtIndex(uint8 quorumNumber, uint32 index) external view returns (OperatorIndexUpdate memory);

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