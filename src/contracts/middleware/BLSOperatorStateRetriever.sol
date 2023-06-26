// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./BLSRegistryCoordinatorWithIndices.sol";

import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";


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

    IBLSRegistryCoordinatorWithIndices public registryCoordinator;
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IIndexRegistry public indexRegistry;

    constructor(IBLSRegistryCoordinatorWithIndices _registryCoordinator) {
        registryCoordinator = _registryCoordinator;

        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
        indexRegistry = _registryCoordinator.indexRegistry();
    }

    /**
     * @notice returns the ordered list of operators (id and stake) for each quorum
     * @param operatorId the id of the operator calling the function
     * @return 2d array of operators. For each quorum the provided operaor is a part of, a ordered list of operators
     */
    function getOperatorState(bytes32 operatorId) external view returns (Operator[][] memory) {
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(registryCoordinator.getCurrentQuorumBitmapByOperatorId(operatorId));

        return getOperatorState(quorumNumbers);
    }

    /**
     * @notice returns the ordered list of operators (id and stake) for each quorum
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @return 2d array of operators. For each quorum, a ordered list of operators
     */
    function getOperatorState(bytes memory quorumNumbers) public view returns(Operator[][] memory) {
        Operator[][] memory operators = new Operator[][](quorumNumbers.length);
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbers[i]);
            bytes32[] memory operatorIds = indexRegistry.getOperatorListForQuorum(quorumNumber);
            operators[i] = new Operator[](operatorIds.length);
            for (uint256 j = 0; j < operatorIds.length; j++) {
                bytes32 operatorId = bytes32(operatorIds[j]);
                operators[i][j] = Operator({
                    operatorId: operatorId,
                    stake: stakeRegistry.getMostRecentStakeUpdateByOperatorId(operatorId, quorumNumber).stake
                });
            }
        }
            
        return operators;
    }

    /**
     * @notice returns 
     *          1) the indices of the quorumBitmaps for each of given operators at the given blocknumber
     *          2) the indices of the total stakes entries for the given quorums at the given blocknumber
     *          3) the indices of the stakes of each of the nonsigners in each of the quorums they are a 
     *             part of (for each nonsigner, an array of length the number of quorums they a part of 
     *             that are also part of the provided quorumNumbers) at the given blocknumber
     *          4) the indices of the quourm apks for each of the provided quorums at the given blocknumber
     * @param referenceBlockNumber is the block number to get the indices for
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @param nonSignerOperatorIds are the ids of the nonsigning operators
     */
    function getCheckSignaturesData(
        uint32 referenceBlockNumber, 
        bytes calldata quorumNumbers, 
        bytes32[] calldata nonSignerOperatorIds
    ) external view returns (CheckSignaturesIndices memory) {
        uint256 quorumBitmap = BitmapUtils.bytesArrayToBitmap(quorumNumbers);

        CheckSignaturesIndices memory checkSignaturesIndices;

        checkSignaturesIndices.nonSignerQuorumBitmapIndices = registryCoordinator.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(referenceBlockNumber, nonSignerOperatorIds);
        checkSignaturesIndices.totalStakeIndices = stakeRegistry.getTotalStakeIndicesByQuorumNumbersAtBlockNumber(referenceBlockNumber, quorumNumbers);
        
      checkSignaturesIndices.nonSignerStakeIndices = new uint32[][](nonSignerOperatorIds.length);
        for (uint i = 0; i < nonSignerOperatorIds.length; i++) {
            uint192 nonSignerQuorumBitmap = 
                registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                    nonSignerOperatorIds[i], 
                    referenceBlockNumber, 
                    checkSignaturesIndices.nonSignerQuorumBitmapIndices[i]
                );
            // the number of quorums the operator is a part of that are also part of the provided quorumNumbers
            checkSignaturesIndices.nonSignerStakeIndices[i] = new uint32[](BitmapUtils.countNumOnes(nonSignerQuorumBitmap & quorumBitmap));

            for (uint8 j = 0; j < 192; j++) {
                // if the operator is a part of the quorum and the quorum is a part of the provided quorums
                if (nonSignerQuorumBitmap >> j & (quorumBitmap >> j & 1) == 1) {
                    checkSignaturesIndices.nonSignerStakeIndices[i][j] = stakeRegistry.getStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber(
                        nonSignerOperatorIds[i],
                        uint8(j),
                        referenceBlockNumber
                    );
                }
            }
        }

        checkSignaturesIndices.quorumApkIndices = blsPubkeyRegistry.getApkIndicesForQuorumsAtBlockNumber(quorumNumbers, referenceBlockNumber);

        return checkSignaturesIndices;
    }
}
