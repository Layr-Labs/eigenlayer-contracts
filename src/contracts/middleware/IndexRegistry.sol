// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../libraries/BN254.sol";


contract IndexRegistry is IIndexRegistry {

    IRegistryCoordinator registryCoordinator;
    bytes32[] public globalOperatorList;

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
        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];
            quorumToOperatorList[quorumNumber].push(operatorId);
            globalOperatorList.push(operatorId);

            _updateOperatorIdToIndexHistory(operatorId, quorumNumber);
            _updateTotalOperatorHistory(quorumNumber);
        }
    }

    function deregisterOperator(bytes32 operatorId, uint8[] memory quorumNumbers, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) external onlyRegistryCoordinator {
        require(quorumNumbers.length == indexes.length, "IndexRegistry.deregisterOperator: quorumNumbers and indexes must be the same length");
        _removeOperatorFromGlobalOperatorList(globalOperatorListIndex);  

        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];

            _removeOperatorFromQuorumToOperatorList(quorumNumber, quorumToOperatorListIndexes[i]);
            _updateTotalOperatorHistory(quorumNumber);
        }
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
        OperatorIndex memory operatorIndexToCheck = operatorIdToIndexHistory[operatorId][quorumNumber][index];
        
        require(blockNumber >= operatorIndexToCheck.toBlockNumber, "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number");

        if(index < operatorIdToIndexHistory[operatorId][quorumNumber].length - 1){
            OperatorIndex memory nextOperatorIndex = operatorIdToIndexHistory[operatorId][quorumNumber][index + 1];
            require(blockNumber < nextOperatorIndex.toBlockNumber, "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number")
        }
        return operatorIndexToCheck.index;
    }

    /**
     * @notice Looks up the number of total operators for `quorumNumber` at the specified `blockNumber`.
     * @param quorumNumber is the quorum number for which the total number of operators is desired
     * @param blockNumber is the block number at which the total number of operators is desired
     * @param index is the index of the entry in the dynamic array `totalOperatorsHistory[quorumNumber]` to read data from
     */
    function getTotalOperatorsForQuorumAtBlockNumberByIndex(uint8 quorumNumber, uint32 blockNumber, uint32 index) external view returns (uint32){
        OperatorIndex memory operatorIndexToCheck = totalOperatorsHistory[quorumNumber][index];

        require(blockNumber >= operatorIndexToCheck.toBlockNumber, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number");
        if(index < totalOperatorsHistory[quorumNumber].length - 1){
            OperatorIndex memory nextOperatorIndex = totalOperatorsHistory[quorumNumber][index + 1];
            require(blockNumber < nextOperatorIndex.toBlockNumber, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number")
        }
        return operatorIndexToCheck.index;
    }

    function totalOperatorsForQuorum(uint8 quorumNumber) external view returns (uint32){
        return uint32(quorumToOperatorList[quorumNumber].length);
    }


    /// @notice Returns the current number of operators of this service.
    function totalOperators() external view returns (uint32){
        return numOperators;
    }


    function _updateTotalOperatorHistory(uint8 quorumNumber, uint256 numOperators) internal {
        uint256 numOperators = quorumToOperatorList[quorumNumber].length;

        OperatorIndex memory totalOperatorUpdate = OperatorIndex({
            blockNumber: block.number,
            index: uint32(numOperators)
        });
        totalOperatorsHistory[quorumNumber].push(totalOperatorUpdate);
    }

    function _updateOperatorIdToIndexHistory(bytes32 operatorId, uint8 quorumNumber) internal {
        if (operatorIdToIndexHistoryLength > 0) {
            uint256 operatorIdToIndexHistoryLength = operatorIdToIndexHistory[operatorIdToSwap][quorumNumber].length;
            operatorIdToIndexHistory[operatorIdToSwap][quorumNumber][operatorIdToIndexHistoryLength - 1].toBlockNumber = block.number;
        }
        OperatorIndex memory latestOperatorIndex;
        latestOperatorIndex.index = quorumToOperatorList[quorumNumber].length - 1;
        operatorIdToIndexHistory[operatorId][quorumNumber].push(latestOperatorIndex);
    }


    function _removeOperatorFromQuorumToOperatorList(uint8 quorumNumber, uint32 indexToRemove) internal {   

        quorumToOperatorListLastIndex  = quorumToOperatorList[quorumNumber].length - 1;

        if(indexToRemove != quorumToOperatorListLastIndex){
            bytes32 operatorIdToSwap = quorumToOperatorList[quorumNumber][quorumToOperatorListLastIndex];
            //update the swapped operator's
            _updateOperatorIdToIndexHistory(operatorIdToSwap, quorumNumber);

            quorumToOperatorList[quorumNumber][indexToRemove] = operatorIdToSwap;
        }
        quorumToOperatorList[quorumNumber].pop();
    }

    function _removeOperatorFromGlobalOperatorList(uint32 indexToRemove) internal {
        uint32 globalOperatorListLastIndex = globalOperatorList.length - 1;

        if(indexToRemove != globalOperatorListLastIndex){
            bytes32 operatorIdToSwap = globalOperatorList[globalOperatorListLastIndex];
            globalOperatorList[indexToRemove] = operatorIdToSwap;
        }
        globalOperatorList.pop();
    }


}