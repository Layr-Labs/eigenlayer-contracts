// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistry.sol";

/**
 * @title Interface for a `Registry` that keeps track of stakes of operators for up to 256 quroums.
 * @author Layr Labs, Inc.
 */
interface IStakeRegistry is IRegistry {
    // DATA STRUCTURES

    /// @notice struct used to store the stakes of an individual operator or the sum of all operators' stakes, for storage
    struct OperatorStakeUpdate {
        // the block number at which the stake amounts were updated and stored
        uint32 updateBlockNumber;
        // the block number at which the *next update* occurred.
        /// @notice This entry has the value **0** until another update takes place.
        uint32 nextUpdateBlockNumber;
        // stake weight for the quorum
        uint96 stake;
    }

    /**
     * @notice Registers the `operator` with `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operator The address of the operator to register.
     * @param operatorId The id of the operator to register.
     * @param quorumBitmap The bitmap of the quorums the operator is registering for.
     * @dev Permissioned by RegistryCoordinator
     */
    function registerOperator(address operator, bytes32 operatorId, uint8 quorumBitmap) external;

    /**
     * @notice Deregisters the operator with `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId The id of the operator to deregister.
     * @param quorumBitmap The bitmap of the quorums the operator is deregistering from.
     * @dev Permissioned by RegistryCoordinator
     */
    function deregisterOperator(bytes32 operatorId, uint8 quorumBitmap) external;

    function getLengthOfTotalStakeHistoryForQuorum(uint8 quorumNumber) external view returns (uint256);

    /**
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory` for quorum `quorumNumber`.
     * @dev Function will revert in the event that `index` is out-of-bounds.
     */
    function getTotalStakeUpdateForQuorumFromIndex(uint8 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory);

    /**
     * @notice Returns the `index`-th entry in the `operatorIdToStakeHistory[operatorId][quorumNumber]` array.
     * @param quorumNumber The quorum number to get the stake for.
     * @param operatorId The id of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `operatorIdToStakeHistory[operatorId][quorumNumber]`.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeUpdateForQuorumFromOperatorIdAndIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index)
        external
        view
        returns (OperatorStakeUpdate memory);

    /**
     * @notice Returns the stake weight corresponding to `operatorId` for quorum `quorumNumber`, at the
     * `index`-th entry in the `operatorIdToStakeHistory[operatorId][quorumNumber]` array if it was the operator's
     * stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param operatorId The id of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `operatorIdToStakeHistory[operatorId][quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index)
        external
        view
        returns (uint96);

    /**
     * @notice Returns the total stake weight for quorum `quorumNumber`, at the `index`-th entry in the 
     * `totalStakeHistory[quorumNumber]` array if it was the stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     */
    function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (uint96);

    /**
     * @notice Checks that the `operator` was active at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.
     * @param operator is the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had stake in
     * @param stakeHistoryIndex specifies an index in `operatorIdToStakeHistory[operatorId]`, where `operatorId` is looked up
     * in `registryCoordinator.getOperatorId(operator)`
     * @return 'true' if it is succesfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `operatorIdToStakeHistory[operatorId][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `operatorIdToStakeHistory[operatorId][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `operatorIdToStakeHistory[operatorId][quorumNumber][index].stake > 0` i.e. the operator had nonzero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was inactive at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     */
    function checkOperatorActiveAtBlockNumber(
        address operator,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool);

    /**
     * @notice Checks that the `operator` was inactive at the `blockNumber`, using the specified `stakeHistoryIndex` for `quorumNumber` as proof.
     * @param operator is the operator of interest
     * @param blockNumber is the block number of interest
     * @param quorumNumber is the quorum number which the operator had no stake in
     * @param stakeHistoryIndex specifies an index in `operatorIdToStakeHistory[operatorId]`, where `operatorId` is looked up
     * in `registryCoordinator.getOperatorId(operator)`
     * @return 'true' if it is succesfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise
     * @dev In order for this function to return 'true', the inputs must satisfy all of the following list:
     * 1) `operatorIdToStakeHistory[operatorId][quorumNumber][index].updateBlockNumber <= blockNumber`
     * 2) `operatorIdToStakeHistory[operatorId][quorumNumber][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
     * is must be strictly greater than `blockNumber`
     * 3) `operatorIdToStakeHistory[operatorId][quorumNumber][index].stake == 0` i.e. the operator had zero stake
     * @dev Note that a return value of 'false' does not guarantee that the `operator` was active at `blockNumber`, since a
     * bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.
     */
    function checkOperatorInactiveAtBlockNumber(
        address operator,
        uint256 blockNumber,
        uint8 quorumNumber,
        uint256 stakeHistoryIndex
        ) external view returns (bool);

    /**
     * @notice Returns the most recent stake weight for the `operator` for quorum `quorumNumber`
     * @dev Function returns weight of **0** in the event that the operator has no stake history
     */
    function getCurrentOperatorStakeForQuorum(address operator, uint8 quorumNumber) external view returns (uint96);

    /// @notice Returns the stake weight from the latest entry in `totalStakeHistory` for quorum `quorumNumber`.
    function getCurrentTotalStakeForQuorum(uint8 quorumNumber) external view returns (uint96);

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operatorIds are the ids of the nodes whose deposit information is getting updated
     * @param prevElements are the elements before this middleware in the operator's linked list within the slasher
     * @dev updates the stakes of the operators in storage and communicates the updates to the service manager that sends them to the slasher
     */
    function updateStakes(bytes32[] memory operatorIds, uint256[] memory prevElements) external;
}