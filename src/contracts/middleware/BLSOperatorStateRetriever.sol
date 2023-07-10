// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSRegistryCoordinatorWithIndices.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";

/**
 * @title BLSOperatorStateRetriever with view functions that allow to retrieve the state of an AVSs registry system.
 * @author Layr Labs Inc.
 */
contract BLSOperatorStateRetriever {
    struct Operator {
        bytes32 operatorId;
        uint96 stake;
    }

    struct CheckSignaturesIndices {
        uint32[] nonSignerQuorumBitmapIndices;
        uint32[] quorumApkIndices;
        uint32[] totalStakeIndices;  
        uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
    }

    /**
     * @notice This function is intended to to be called by AVS operators every time a new task is created (i.e.)
     * the AVS coordinator makes a request to AVS operators. Since all of the crucial information is kept onchain, 
     * operators don't need to run indexers to fetch the data.
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param operatorId the id of the operator to fetch the quorums lists 
     * @param blockNumber is the block number to get the operator state for
     * @return 1) the quorumBitmap of the operator at the given blockNumber
     *         2) 2d array of Operator structs. For each quorum the provided operator 
     *            was a part of at `blockNumber`, an ordered list of operators.
     */
    function getOperatorState(IBLSRegistryCoordinatorWithIndices registryCoordinator, bytes32 operatorId, uint32 blockNumber) external returns (uint256, Operator[][] memory) {
        bytes32[] memory operatorIds = new bytes32[](1);
        operatorIds[0] = operatorId;
        uint256 index = registryCoordinator.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber, operatorIds)[0];
    
        uint256 quorumBitmap = registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId, blockNumber, index);

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        return (quorumBitmap, getOperatorState(registryCoordinator, quorumNumbers, blockNumber));
    }

    /**
     * @notice returns the ordered list of operators (id and stake) for each quorum. The AVS coordinator 
     * may call this function directly to get the operator state for a given block number
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @param blockNumber is the block number to get the operator state for
     * @return 2d array of operators. For each quorum, an ordered list of operators
     */
    function getOperatorState(IBLSRegistryCoordinatorWithIndices registryCoordinator, bytes memory quorumNumbers, uint32 blockNumber) public returns(Operator[][] memory) {
        IStakeRegistry stakeRegistry = registryCoordinator.stakeRegistry();
        IIndexRegistry indexRegistry = registryCoordinator.indexRegistry();

        Operator[][] memory operators = new Operator[][](quorumNumbers.length);
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbers[i]);
            bytes32[] memory operatorIds = indexRegistry.getOperatorListForQuorumAtBlockNumber(quorumNumber, blockNumber);
            operators[i] = new Operator[](operatorIds.length);
            for (uint256 j = 0; j < operatorIds.length; j++) {
                bytes32 operatorId = bytes32(operatorIds[j]);
                operators[i][j] = Operator({
                    operatorId: operatorId,
                    stake: stakeRegistry.getStakeForOperatorIdForQuorumAtBlockNumber(operatorId, quorumNumber, blockNumber)
                });
            }
        }
            
        return operators;
    }

    /**
     * @notice this is called by the AVS operator to get the relevant indices for the checkSignatures function
     * if they are not running an indexer    
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param referenceBlockNumber is the block number to get the indices for
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @param nonSignerOperatorIds are the ids of the nonsigning operators
     * @return 1) the indices of the quorumBitmaps for each of the operators in the @param nonSignerOperatorIds array at the given blocknumber
     *         2) the indices of the total stakes entries for the given quorums at the given blocknumber
     *         3) the indices of the stakes of each of the nonsigners in each of the quorums they were a 
     *            part of (for each nonsigner, an array of length the number of quorums they were a part of
     *            that are also part of the provided quorumNumbers) at the given blocknumber
     *         4) the indices of the quorum apks for each of the provided quorums at the given blocknumber
     */
    function getCheckSignaturesIndices(
        IBLSRegistryCoordinatorWithIndices registryCoordinator,
        uint32 referenceBlockNumber, 
        bytes calldata quorumNumbers, 
        bytes32[] calldata nonSignerOperatorIds
    ) external view returns (CheckSignaturesIndices memory) {
        IStakeRegistry stakeRegistry = registryCoordinator.stakeRegistry();
        CheckSignaturesIndices memory checkSignaturesIndices;

        checkSignaturesIndices.nonSignerQuorumBitmapIndices = registryCoordinator.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(referenceBlockNumber, nonSignerOperatorIds);
        checkSignaturesIndices.totalStakeIndices = stakeRegistry.getTotalStakeIndicesByQuorumNumbersAtBlockNumber(referenceBlockNumber, quorumNumbers);
        
        checkSignaturesIndices.nonSignerStakeIndices = new uint32[][](quorumNumbers.length);
        for (uint8 quorumNumberIndex = 0; quorumNumberIndex < quorumNumbers.length; quorumNumberIndex++) {
            uint256 numNonSignersForQuorum = 0;
            // this array's length will be at most the number of nonSignerOperatorIds
            checkSignaturesIndices.nonSignerStakeIndices[quorumNumberIndex] = new uint32[](nonSignerOperatorIds.length);

            for (uint i = 0; i < nonSignerOperatorIds.length; i++) {
                uint192 nonSignerQuorumBitmap = 
                    registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                        nonSignerOperatorIds[i], 
                        referenceBlockNumber, 
                        checkSignaturesIndices.nonSignerQuorumBitmapIndices[i]
                    );
                
                // if the operator was a part of the quorum and the quorum is a part of the provided quorumNumbers
                if (nonSignerQuorumBitmap >> uint8(quorumNumbers[quorumNumberIndex]) & 1 == 1) {
                    checkSignaturesIndices.nonSignerStakeIndices[quorumNumberIndex][numNonSignersForQuorum] = stakeRegistry.getStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber(
                        nonSignerOperatorIds[i],
                        uint8(quorumNumbers[quorumNumberIndex]),
                        referenceBlockNumber
                    );
                    numNonSignersForQuorum++;
                }
            }

            // resize the array to the number of nonSigners for this quorum
            uint32[] memory nonSignerStakeIndicesForQuorum = new uint32[](numNonSignersForQuorum);
            for (uint i = 0; i < numNonSignersForQuorum; i++) {
                nonSignerStakeIndicesForQuorum[i] = checkSignaturesIndices.nonSignerStakeIndices[quorumNumberIndex][i];
            }
            checkSignaturesIndices.nonSignerStakeIndices[quorumNumberIndex] = nonSignerStakeIndicesForQuorum;
        }

        IBLSPubkeyRegistry blsPubkeyRegistry = registryCoordinator.blsPubkeyRegistry();
        checkSignaturesIndices.quorumApkIndices = blsPubkeyRegistry.getApkIndicesForQuorumsAtBlockNumber(quorumNumbers, referenceBlockNumber);

        return checkSignaturesIndices;
    }
}
