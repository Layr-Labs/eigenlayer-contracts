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
     * @notice returns the ordered list of operators (id and stake) for each quorum
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param operatorId the id of the operator to fetch the quorums lists 
     * @param blockNumber is the block number to get the operator state for
     * @return 1) the quorumBitmap of the operator at the given blockNumber
     *         2) 2d array of Operator tructs. For each quorum the provided operator 
     *            was a part of at `blockNumber`, an ordered list of operators.
     */
    function getOperatorState(IBLSRegistryCoordinatorWithIndices registryCoordinator, bytes32 operatorId, uint32 blockNumber) external view returns (uint256, Operator[][] memory) {
        bytes32[] memory operatorIds = new bytes32[](1);
        operatorIds[0] = operatorId;
        uint256 index = registryCoordinator.getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber, operatorIds)[0];
        uint256 quorumBitmap = registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId, blockNumber, index);

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        return (quorumBitmap, getOperatorState(registryCoordinator, quorumNumbers, blockNumber));
    }

    /**
     * @notice returns the ordered list of operators (id and stake) for each quorum
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @param blockNumber is the block number to get the operator state for
     * @return 2d array of operators. For each quorum, an ordered list of operators
     */
    function getOperatorState(IBLSRegistryCoordinatorWithIndices registryCoordinator, bytes memory quorumNumbers, uint32 blockNumber) public view returns(Operator[][] memory) {
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
     * @notice returns 
     *          1) the indices of the quorumBitmaps for each of given operators at the given blocknumber
     *          2) the indices of the total stakes entries for the given quorums at the given blocknumber
     *          3) the indices of the stakes of each of the nonsigners in each of the quorums they are a 
     *             part of (for each nonsigner, an array of length the number of quorums they a part of 
     *             that are also part of the provided quorumNumbers) at the given blocknumber
     *          4) the indices of the quorum apks for each of the provided quorums at the given blocknumber
     * @param registryCoordinator is the registry coordinator to fetch the AVS registry information from
     * @param referenceBlockNumber is the block number to get the indices for
     * @param quorumNumbers are the ids of the quorums to get the operator state for
     * @param nonSignerOperatorIds are the ids of the nonsigning operators
     */
    function getCheckSignaturesIndices(
        IBLSRegistryCoordinatorWithIndices registryCoordinator,
        uint32 referenceBlockNumber, 
        bytes calldata quorumNumbers, 
        bytes32[] calldata nonSignerOperatorIds
    ) external view returns (CheckSignaturesIndices memory) {
        uint256 quorumBitmap = BitmapUtils.bytesArrayToBitmap(quorumNumbers);
        IStakeRegistry stakeRegistry = registryCoordinator.stakeRegistry();
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

        IBLSPubkeyRegistry blsPubkeyRegistry = registryCoordinator.blsPubkeyRegistry();
        checkSignaturesIndices.quorumApkIndices = blsPubkeyRegistry.getApkIndicesForQuorumsAtBlockNumber(quorumNumbers, referenceBlockNumber);

        return checkSignaturesIndices;
    }
}
