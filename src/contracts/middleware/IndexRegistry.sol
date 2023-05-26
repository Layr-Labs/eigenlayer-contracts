// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../libraries/BN254.sol";


contract IndexRegistry is IIndexRegistry {

    IRegistryCoordinator public registryCoordinator;
    bytes32[] public globalOperatorList;

    mapping(uint8 => bytes32[]) public quorumToOperatorList;
    mapping(bytes32 => mapping(uint8 => OperatorIndex[])) public operatorIdToIndexHistory;
    mapping(uint8 => OperatorIndex[]) public totalOperatorsHistory;

    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "IndexRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ){
        registryCoordinator = _registryCoordinator;
    }

    /**
     * @notice Registers the operator with the specified `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId is the id of the operator that is being registered
     * @param quorumNumbers is the quorum numbers the operator is registered for
     * @dev access restricted to the RegistryCoordinator
     */
    function registerOperator(bytes32 operatorId, uint8[] memory quorumNumbers) external onlyRegistryCoordinator {
        //add operator to operatorList
        globalOperatorList.push(operatorId);

        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumToOperatorList[quorumNumber].push(operatorId);
            _updateOperatorIdToIndexHistory(operatorId, quorumNumbers[i], quorumToOperatorList[quorumNumber].length - 1);
            _updateTotalOperatorHistory(quorumNumbers[i]);
        }
    }

    /**
     * @notice Deregisters the operator with the specified `operatorId` for the quorums specified by `quorumBitmap`.
     * @param operatorId is the id of the operator that is being deregistered
     * @param quorumNumbers is the quorum numbers the operator is deregistered for
     * @param quorumToOperatorListIndexes is an array of indexes for each quorum as witnesses for the last operators to swap for each quorum
     * @param globalOperatorListIndex is the index of the operator that is to be removed from the list
     * @dev access restricted to the RegistryCoordinator
     */
    function deregisterOperator(bytes32 operatorId, uint8[] memory quorumNumbers, uint32[] memory quorumToOperatorListIndexes, uint32 globalOperatorListIndex) external onlyRegistryCoordinator {
        require(quorumNumbers.length == indexes.length, "IndexRegistry.deregisterOperator: quorumNumbers and indexes must be the same length");
        _removeOperatorFromGlobalOperatorList(globalOperatorListIndex);  

        for (uint i = 0; i < quorumNumbers.length; i++) {
            _removeOperatorFromQuorumToOperatorList(quorumNumbers[i], quorumToOperatorListIndexes[i]);
            _updateTotalOperatorHistory(quorumNumbers[i]);
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

        //blocknumber must be before the "index'th" entry's toBlockNumber
        require(blockNumber <= operatorIndexToCheck.toBlockNumber, "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number");
       
        //if there is an index update before the "index'th" update, the blocknumber must be after than the previous entry's toBlockNumber
        if(index != 0){
            OperatorIndex memory previousOperatorIndex = operatorIdToIndexHistory[operatorId][quorumNumber][index - 1];
            require(blockNumber > previousOperatorIndex.toBlockNumber, "IndexRegistry.getOperatorIndexForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number");
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

        //blocknumber must be before the "index'th" entry's toBlockNumber
        require(blockNumber <= operatorIndexToCheck.toBlockNumber, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the future for provided block number");
        
        //if there is an index update before the "index'th" update, the blocknumber must be after than the previous entry's toBlockNumber
        if (index != 0){
            OperatorIndex memory previousOperatorIndex = totalOperatorsHistory[quorumNumber][index - 1];
            require(blockNumber > previousOperatorIndex.toBlockNumber, "IndexRegistry.getTotalOperatorsForQuorumAtBlockNumberByIndex: provided index is too far in the past for provided block number");
        }
        return operatorIndexToCheck.index;
    }

    function totalOperatorsForQuorum(uint8 quorumNumber) external view returns (uint32){
        return uint32(quorumToOperatorList[quorumNumber].length);
    }


    /// @notice Returns the current number of operators of this service.
    function totalOperators() external view returns (uint32){
        return uint32(globalOperatorList.length);
    }


    function _updateTotalOperatorHistory(uint8 quorumNumber) internal {

        //if there is a prior entry, update its "toBlockNumber"
        if (totalOperatorsHistory[quorumNumber].length > 0) {
            totalOperatorsHistory[quorumNumber][totalOperatorsHistory[quorumNumber].length - 1].toBlockNumber = block.number;
        }

        OperatorIndex memory totalOperatorUpdate;
        totalOperatorUpdate.index = quorumToOperatorList[quorumNumber].length - 1;
        totalOperatorsHistory[quorumNumber].push(totalOperatorUpdate);
    }

    /// 
    /// @param operatorId operatorId of the operator to update
    /// @param quorumNumber quorumNumber of the operator to update
    /// @param index the latest index of that operator in quorumToOperatorList
    function _updateOperatorIdToIndexHistory(bytes32 operatorId, uint8 quorumNumber, uint32 index) internal {
        uint256 operatorIdToIndexHistoryLength = operatorIdToIndexHistory[operatorId][quorumNumber].length;

        //if there is a prior entry for the operator, set the previous entry's toBlocknumber
        if (operatorIdToIndexHistoryLength > 0) {
            operatorIdToIndexHistory[operatorIdToSwap][quorumNumber][operatorIdToIndexHistoryLength - 1].toBlockNumber = block.number;
        }
        OperatorIndex memory latestOperatorIndex;
        latestOperatorIndex.index = index; 
        operatorIdToIndexHistory[operatorId][quorumNumber].push(latestOperatorIndex);
    }

    /// @notice when we remove an operator from quorumToOperatorList, we swap the last operator in 
    ///         quorumToOperatorList[quorumNumber] to the position of the operator we are removing
    /// @param quorumNumber quorum number of the operator to remove
    /// @param indexToRemove index of the operator to remove
    function _removeOperatorFromQuorumToOperatorList(uint8 quorumNumber, uint32 indexToRemove) internal {   
        
        uint32 quorumToOperatorListLastIndex  = uint32(quorumToOperatorList[quorumNumber].length - 1);

        // if the operator is not the last in the list, we must swap the last operator into their positon
        if(indexToRemove != quorumToOperatorListLastIndex){
            bytes32 operatorIdToSwap = quorumToOperatorList[quorumNumber][quorumToOperatorListLastIndex];
            //update the swapped operator's operatorIdToIndexHistory list with a new entry, as their index has now changed
            _updateOperatorIdToIndexHistory(operatorIdToSwap, quorumNumber, indexToRemove);

            //set last operator in the list to removed operator's position in the array
            quorumToOperatorList[quorumNumber][indexToRemove] = operatorIdToSwap;
        }
        //marking the final entry in the deregistering operator's operatorIdToIndexHistory entry with the deregistration block number
        operatorIdToIndexHistory[operatorId][quorumNumber][operatorIdToIndexHistory[operatorId][quorumNumber].length - 1].toBlockNumber = block.number;
        //removing the swapped operator from the list
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