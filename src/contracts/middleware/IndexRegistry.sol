// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../libraries/BN254.sol";


contract IndexRegistry is IIndexRegistry {

    IRegistryCoordinator registryCoordinator;
    
    bytes32[] public operatorList;
    mapping(uint8 => bytes32[]) public quorumToOperatorList;

    mapping(bytes32 => mapping(uint8 => OperatorIndex[])) operatorIdToIndexHistory;
    mapping(uint8 => OperatorIndex[]) totalOperatorsHistory;

    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
        _;
    }

     
    constructor(
        IRegistryCoordinator _registryCoordinator
    ){
        registryCoordinator = _registryCoordinator;

    }
    function registerOperator(bytes32 operatorId, uint8[] memory quorumNumbers) external onlyRegistryCoordinator {
        //add operator to operatorList
        operatorList.push(operatorId);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];
            uint256 numOperators = quorumToOperatorList[quorumNumber].length;

            quorumToOperatorList[quorumNumber].push(operatorId);

            _updateTotalOperatorHistory(quorumNumber, numOperators);
            _updateoperatorIdToIndexHistory(operatorId, quorumNumber, numOperators - 1);
        }
    }

    function deregisterOperator(bytes32 operatorId, uint8[] memory quorumNumbers, uint32[] memory indexes) external onlyRegistryCoordinator {
        require(quorumNumbers.length == indexes.length, "IndexRegistry.deregisterOperator: quorumNumbers and indexes must be the same length");



    }

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
    function getOperatorIndexForQuorumAtBlockNumberByIndex(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber, uint32 index) external view returns (uint32){

    }

    /**
     * @notice Looks up the number of total operators for `quorumNumber` at the specified `blockNumber`.
     * @param quorumNumber is the quorum number for which the total number of operators is desired
     * @param blockNumber is the block number at which the total number of operators is desired
     * @param index is the index of the entry in the dynamic array `totalOperatorsHistory[quorumNumber]` to read data from
     */
    function getTotalOperatorsForQuorumAtBlockNumberByIndex(uint8 quorumNumber, uint32 blockNumber, uint32 index) external view returns (uint32){

    }

    /// @notice Returns the current number of operators of this service.
    function totalOperators() external view returns (uint32){

    }


    function _updateTotalOperatorHistory(uint8 quorumNumber, uint256 numOperators) internal {

        OperatorIndex memory totalOperatorUpdate = OperatorIndex({
            blockNumber: block.number,
            index: uint32(numOperators)
        });
        totalOperatorsHistory[quorumNumber].push(totalOperatorUpdate);
    }

    function _updateoperatorIdToIndexHistory(bytes32 operatorId, uint8 quorumNumber, uint256 numOperators) internal {

        OperatorIndex memory operatorIndex = OperatorIndex({
            blockNumber: block.number,
            index: uint32(numOperators - 1)
        });
        operatorIdToIndexHistory[operatorId][quorumNumber].push(operatorIndex);
    }


}