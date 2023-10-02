// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";

contract BLSOperatorStateRetrieverUnitTests is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();
    }

    function testGetOperatorState_Valid(uint256 pseudoRandomNumber) public {
        // register random operators and get the expected indices within the quorums and the metadata for the operators
        (
            OperatorMetadata[] memory operatorMetadatas,
            uint256[][] memory expectedOperatorOverallIndices
        ) = _registerRandomOperators(pseudoRandomNumber);

        for (uint i = 0; i < operatorMetadatas.length; i++) {
            uint32 blockNumber = uint32(registrationBlockNumber + blocksBetweenRegistrations * i);

            uint256 gasBefore = gasleft();
            // retrieve the ordered list of operators for each quorum along with their id and stake
            (uint256 quorumBitmap, BLSOperatorStateRetriever.Operator[][] memory operators) = 
                operatorStateRetriever.getOperatorState(registryCoordinator, operatorMetadatas[i].operatorId, blockNumber);
            uint256 gasAfter = gasleft();
            emit log_named_uint("gasUsed", gasBefore - gasAfter);

            assertEq(operatorMetadatas[i].quorumBitmap, quorumBitmap);
            bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
            
            // assert that the operators returned are the expected ones
            _assertExpectedOperators(
                quorumNumbers,
                operators,
                expectedOperatorOverallIndices,
                operatorMetadatas
            );
        }

        // choose a random operator to deregister
        uint256 operatorIndexToDeregister = pseudoRandomNumber % maxOperatorsToRegister;
        bytes memory quorumNumbersToDeregister = BitmapUtils.bitmapToBytesArray(operatorMetadatas[operatorIndexToDeregister].quorumBitmap);
        // get the operatorIds of the last operators in each quorum to swap with the operator to deregister
        bytes32[] memory operatorIdsToSwap = new bytes32[](quorumNumbersToDeregister.length);
        for (uint i = 0; i < quorumNumbersToDeregister.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbersToDeregister[i]);
            operatorIdsToSwap[i] = operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][expectedOperatorOverallIndices[quorumNumber].length - 1]].operatorId;
        }

        uint32 deregistrationBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * (uint32(operatorMetadatas.length) + 1);
        cheats.roll(deregistrationBlockNumber);

        cheats.prank(_incrementAddress(defaultOperator, operatorIndexToDeregister));
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbersToDeregister, operatorMetadatas[operatorIndexToDeregister].pubkey, operatorIdsToSwap);

        // modify expectedOperatorOverallIndices by moving th operatorIdsToSwap to the index where the operatorIndexToDeregister was
        for (uint i = 0; i < quorumNumbersToDeregister.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbersToDeregister[i]);
            // loop through indices till we find operatorIndexToDeregister, then move that last operator into that index
            for (uint j = 0; j < expectedOperatorOverallIndices[quorumNumber].length; j++) {
                if (expectedOperatorOverallIndices[quorumNumber][j] == operatorIndexToDeregister) {
                    expectedOperatorOverallIndices[quorumNumber][j] = expectedOperatorOverallIndices[quorumNumber][expectedOperatorOverallIndices[quorumNumber].length - 1];
                    break;
                }
            }
        }

        // make sure the state retriever returns the expected state after deregistration
        bytes memory allQuorumNumbers = new bytes(maxQuorumsToRegisterFor);
        for (uint8 i = 0; i < allQuorumNumbers.length; i++) {
            allQuorumNumbers[i] = bytes1(i);
        }
        
        _assertExpectedOperators(
            allQuorumNumbers,
            operatorStateRetriever.getOperatorState(registryCoordinator, allQuorumNumbers, deregistrationBlockNumber),
            expectedOperatorOverallIndices,
            operatorMetadatas
        );
    }

    function testCheckSignaturesIndices_NoNonSigners_Valid(uint256 pseudoRandomNumber) public {
        (
            OperatorMetadata[] memory operatorMetadatas,
            uint256[][] memory expectedOperatorOverallIndices
        ) = _registerRandomOperators(pseudoRandomNumber);

        uint32 cumulativeBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * uint32(operatorMetadatas.length);

        // get the quorum bitmap for which there is at least 1 operator
        uint256 allInclusiveQuorumBitmap = 0;
        for (uint8 i = 0; i < operatorMetadatas.length; i++) {
            allInclusiveQuorumBitmap |= operatorMetadatas[i].quorumBitmap;
        }

        bytes memory allInclusiveQuorumNumbers = BitmapUtils.bitmapToBytesArray(allInclusiveQuorumBitmap);

        bytes32[] memory nonSignerOperatorIds = new bytes32[](0);

        BLSOperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices =
            operatorStateRetriever.getCheckSignaturesIndices(
                registryCoordinator,
                cumulativeBlockNumber, 
                allInclusiveQuorumNumbers, 
                nonSignerOperatorIds
            );

        assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices.length, 0, "nonSignerQuorumBitmapIndices should be empty if no nonsigners");
        assertEq(checkSignaturesIndices.quorumApkIndices.length, allInclusiveQuorumNumbers.length, "quorumApkIndices should be the number of quorums queried for");
        assertEq(checkSignaturesIndices.totalStakeIndices.length, allInclusiveQuorumNumbers.length, "totalStakeIndices should be the number of quorums queried for");
        assertEq(checkSignaturesIndices.nonSignerStakeIndices.length, allInclusiveQuorumNumbers.length, "nonSignerStakeIndices should be the number of quorums queried for");

        // assert the indices are the number of registered operators for the quorum minus 1
        for (uint8 i = 0; i < allInclusiveQuorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(allInclusiveQuorumNumbers[i]);
            assertEq(checkSignaturesIndices.quorumApkIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1, "quorumApkIndex should be the number of registered operators for the quorum minus 1");
            assertEq(checkSignaturesIndices.totalStakeIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1, "totalStakeIndex should be the number of registered operators for the quorum minus 1");
        }
    }

    function testCheckSignaturesIndices_FewNonSigners_Valid(uint256 pseudoRandomNumber) public {
        (
            OperatorMetadata[] memory operatorMetadatas,
            uint256[][] memory expectedOperatorOverallIndices
        ) = _registerRandomOperators(pseudoRandomNumber);

        uint32 cumulativeBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * uint32(operatorMetadatas.length);

        // get the quorum bitmap for which there is at least 1 operator
        uint256 allInclusiveQuorumBitmap = 0;
        for (uint8 i = 0; i < operatorMetadatas.length; i++) {
            allInclusiveQuorumBitmap |= operatorMetadatas[i].quorumBitmap;
        }

        bytes memory allInclusiveQuorumNumbers = BitmapUtils.bitmapToBytesArray(allInclusiveQuorumBitmap);

        bytes32[] memory nonSignerOperatorIds = new bytes32[](pseudoRandomNumber % (operatorMetadatas.length - 1) + 1);
        uint256 randomIndex = uint256(keccak256(abi.encodePacked("nonSignerOperatorIds", pseudoRandomNumber))) % operatorMetadatas.length;
        for (uint i = 0; i < nonSignerOperatorIds.length; i++) {
            nonSignerOperatorIds[i] = operatorMetadatas[(randomIndex + i) % operatorMetadatas.length].operatorId;
        }

        BLSOperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices =
            operatorStateRetriever.getCheckSignaturesIndices(
                registryCoordinator,
                cumulativeBlockNumber, 
                allInclusiveQuorumNumbers, 
                nonSignerOperatorIds
            );

        assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices.length, nonSignerOperatorIds.length, "nonSignerQuorumBitmapIndices should be the number of nonsigners");
        assertEq(checkSignaturesIndices.quorumApkIndices.length, allInclusiveQuorumNumbers.length, "quorumApkIndices should be the number of quorums queried for");
        assertEq(checkSignaturesIndices.totalStakeIndices.length, allInclusiveQuorumNumbers.length, "totalStakeIndices should be the number of quorums queried for");
        assertEq(checkSignaturesIndices.nonSignerStakeIndices.length, allInclusiveQuorumNumbers.length, "nonSignerStakeIndices should be the number of quorums queried for");

        // assert the indices are the number of registered operators for the quorum minus 1
        for (uint8 i = 0; i < allInclusiveQuorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(allInclusiveQuorumNumbers[i]);
            assertEq(checkSignaturesIndices.quorumApkIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1, "quorumApkIndex should be the number of registered operators for the quorum minus 1");
            assertEq(checkSignaturesIndices.totalStakeIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1, "totalStakeIndex should be the number of registered operators for the quorum minus 1");
        }

        // assert the quorum bitmap and stake indices are zero because there have been no kicks or stake updates
        for (uint i = 0; i < nonSignerOperatorIds.length; i++) {
            assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices[i], 0, "nonSignerQuorumBitmapIndices should be zero because there have been no kicks");
        }
        for (uint i = 0; i < checkSignaturesIndices.nonSignerStakeIndices.length; i++) {
            for (uint j = 0; j < checkSignaturesIndices.nonSignerStakeIndices[i].length; j++) {
                assertEq(checkSignaturesIndices.nonSignerStakeIndices[i][j], 0, "nonSignerStakeIndices should be zero because there have been no stake updates past the first one");
            }
        }
    }

    function _assertExpectedOperators(
        bytes memory quorumNumbers,
        BLSOperatorStateRetriever.Operator[][] memory operators,
        uint256[][] memory expectedOperatorOverallIndices,
        OperatorMetadata[] memory operatorMetadatas
    ) internal {
        // for each quorum
        for (uint j = 0; j < quorumNumbers.length; j++) {
            // make sure the each operator id and stake is correct
            for (uint k = 0; k < operators[j].length; k++) {
                uint8 quorumNumber = uint8(quorumNumbers[j]);
                assertEq(operators[j][k].operatorId, operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][k]].operatorId);
                assertEq(operators[j][k].stake, operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][k]].stakes[quorumNumber]);
            }
        }
    }
}