// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/interfaces/IStakeRegistry.sol";

/**
 * @title Interface for a `Registry` that keeps track of stakes of operators for up to 256 quorums.
 * @author Layr Labs, Inc.
 */
contract StakeRegistryMock is IStakeRegistry {

    function registryCoordinator() external view returns (IRegistryCoordinator) {}

    /**
     * @notice Registers the `operator` with `operatorId` for the specified `quorumNumbers`.
     * @param operator The address of the operator to register.
     * @param operatorId The id of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already registered
     */
    function registerOperator(address operator, bytes32 operatorId, bytes memory quorumNumbers) external {}

    /**
     * @notice Deregisters the operator with `operatorId` for the specified `quorumNumbers`.
     * @param operatorId The id of the operator to deregister.
     * @param quorumNumbers The quorum numbers the operator is deregistering from, where each byte is an 8 bit integer quorumNumber.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions (these are assumed, not validated in this contract):
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already deregistered
     *         5) `quorumNumbers` is a subset of the quorumNumbers that the operator is registered for
     */
    function deregisterOperator(bytes32 operatorId, bytes memory quorumNumbers) external {}

    /// @notice In order to register for a quorum i, an operator must have at least `minimumStakeForQuorum[i]`
    function minimumStakeForQuorum(uint256 quorumNumber) external view returns (uint96) {}

    function getLengthOfTotalStakeHistoryForQuorum(uint8 quorumNumber) external view returns (uint256) {}

    /**
     * @notice Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory` for quorum `quorumNumber`.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     */
    function getTotalStakeUpdateForQuorumFromIndex(uint8 quorumNumber, uint256 index) external view returns (OperatorStakeUpdate memory) {}

    /// @notice Returns the indices of the operator stakes for the provided `quorumNumber` at the given `blockNumber`
    function getStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber)
        external
        view
        returns (uint32) {}

    /// @notice Returns the indices of the total stakes for the provided `quorumNumbers` at the given `blockNumber`
    function getTotalStakeIndicesByQuorumNumbersAtBlockNumber(uint32 blockNumber, bytes calldata quorumNumbers) external view returns(uint32[] memory) {}

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
        returns (OperatorStakeUpdate memory) {}

    /**
     * @notice Returns the most recent stake weight for the `operatorId` for a certain quorum
     * @dev Function returns an OperatorStakeUpdate struct with **every entry equal to 0** in the event that the operator has no stake history
     */
    function getMostRecentStakeUpdateByOperatorId(bytes32 operatorId, uint8 quorumNumber) external view returns (OperatorStakeUpdate memory) {}

    /**
     * @notice Returns the stake weight corresponding to `operatorId` for quorum `quorumNumber`, at the
     * `index`-th entry in the `operatorIdToStakeHistory[operatorId][quorumNumber]` array if the entry 
     * corresponds to the operator's stake at `blockNumber`. Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param operatorId The id of the operator of interest.
     * @param index Array index for lookup, within the dynamic array `operatorIdToStakeHistory[operatorId][quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     * @dev used the BLSSignatureChecker to get past stakes of signing operators
     */
    function getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index)
        external
        view
        returns (uint96) {}

    /**
     * @notice Returns the total stake weight for quorum `quorumNumber`, at the `index`-th entry in the 
     * `totalStakeHistory[quorumNumber]` array if the entry corresponds to the total stake at `blockNumber`. 
     * Reverts otherwise.
     * @param quorumNumber The quorum number to get the stake for.
     * @param index Array index for lookup, within the dynamic array `totalStakeHistory[quorumNumber]`.
     * @param blockNumber Block number to make sure the stake is from.
     * @dev Function will revert if `index` is out-of-bounds.
     * @dev used the BLSSignatureChecker to get past stakes of signing operators
     */
    function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (uint96) {}

    /**
     * @notice Returns the most recent stake weight for the `operatorId` for quorum `quorumNumber`
     * @dev Function returns weight of **0** in the event that the operator has no stake history
     */
    function getCurrentOperatorStakeForQuorum(bytes32 operatorId, uint8 quorumNumber) external view returns (uint96) {}

    /// @notice Returns the stake of the operator for the provided `quorumNumber` at the given `blockNumber`
    function getStakeForOperatorIdForQuorumAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber)
        external
        view
        returns (uint96){}

    /**
     * @notice Returns the stake weight from the latest entry in `_totalStakeHistory` for quorum `quorumNumber`.
     * @dev Will revert if `_totalStakeHistory[quorumNumber]` is empty.
     */
    function getCurrentTotalStakeForQuorum(uint8 quorumNumber) external view returns (uint96) {}

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the addresses of the operators whose stake information is getting updated
     */
    function updateStakes(address[] memory operators) external {
        for (uint256 i = 0; i < operators.length; i++) {
            emit StakeUpdate(
                bytes32(uint256(keccak256(abi.encodePacked(operators[i], "operatorId")))),
                uint8(uint256(keccak256(abi.encodePacked(operators[i], i, "quorumNumber")))),
                uint96(uint256(keccak256(abi.encodePacked(operators[i], i, "stake"))))
            );
        }
    }

    function getMockOperatorId(address operator) external returns(bytes32) {
        return bytes32(uint256(keccak256(abi.encodePacked(operator, "operatorId"))));
    }
}