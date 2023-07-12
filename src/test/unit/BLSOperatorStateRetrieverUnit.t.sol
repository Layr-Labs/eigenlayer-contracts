// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";

contract BLSOperatorStateRetrieverUnitTests is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    uint8 maxQuorumsToRegisterFor = 4;
    uint256 maxOperatorsToRegister = 4;
    uint32 registrationBlockNumber = 100;
    uint32 blocksBetweenRegistrations = 10;

    struct OperatorMetadata {
        uint256 quorumBitmap;
        address operator;
        bytes32 operatorId;
        BN254.G1Point pubkey;
        uint96[] stakes; // in every quorum for simplicity
    }

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();
    }

    function testGetOperatorState_Valid(uint256 pseudoRandomNumber) public {
        (
            OperatorMetadata[] memory operatorMetadatas,
            uint256[][] memory expectedOperatorOverallIndices
        ) = _registerRandomOperators(pseudoRandomNumber);

        for (uint i = 0; i < operatorMetadatas.length; i++) {
            uint32 blockNumber = uint32(registrationBlockNumber + blocksBetweenRegistrations * i);

            (uint256 quorumBitmap, BLSOperatorStateRetriever.Operator[][] memory operators) = 
                operatorStateRetriever.getOperatorState(registryCoordinator, operatorMetadatas[i].operatorId, blockNumber);

            assertEq(operatorMetadatas[i].quorumBitmap, quorumBitmap);
            bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
            
            _assertExpectedOperators(
                quorumNumbers,
                operators,
                expectedOperatorOverallIndices,
                operatorMetadatas
            );
        }

        uint256 operatorIndexToDeregister = pseudoRandomNumber % maxOperatorsToRegister;
        bytes memory quorumNumbersToDeregister = BitmapUtils.bitmapToBytesArray(operatorMetadatas[operatorIndexToDeregister].quorumBitmap);
        bytes32[] memory operatorIdsToSwap = new bytes32[](quorumNumbersToDeregister.length);
        for (uint i = 0; i < quorumNumbersToDeregister.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbersToDeregister[i]);
            operatorIdsToSwap[i] = operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][expectedOperatorOverallIndices[quorumNumber].length - 1]].operatorId;
        }

        uint32 deregistrationBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * (uint32(operatorMetadatas.length) + 1);
        cheats.roll(deregistrationBlockNumber);

        cheats.prank(_incrementAddress(defaultOperator, operatorIndexToDeregister));
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbersToDeregister, operatorMetadatas[operatorIndexToDeregister].pubkey, operatorIdsToSwap);

        // modify expectedOperatorOverallIndices accordingly
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

        assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices.length, 0);
        assertEq(checkSignaturesIndices.quorumApkIndices.length, allInclusiveQuorumNumbers.length);
        assertEq(checkSignaturesIndices.totalStakeIndices.length, allInclusiveQuorumNumbers.length);
        assertEq(checkSignaturesIndices.nonSignerStakeIndices.length, 0);

        // assert the indices are the number of registered operators for the quorum minus 1
        for (uint8 i = 0; i < allInclusiveQuorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(allInclusiveQuorumNumbers[i]);
            assertEq(checkSignaturesIndices.quorumApkIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1);
            assertEq(checkSignaturesIndices.totalStakeIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1);
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

        assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices.length, nonSignerOperatorIds.length);
        assertEq(checkSignaturesIndices.quorumApkIndices.length, allInclusiveQuorumNumbers.length);
        assertEq(checkSignaturesIndices.totalStakeIndices.length, allInclusiveQuorumNumbers.length);
        assertEq(checkSignaturesIndices.nonSignerStakeIndices.length, nonSignerOperatorIds.length);

        // assert the indices are the number of registered operators for the quorum minus 1
        for (uint8 i = 0; i < allInclusiveQuorumNumbers.length; i++) {
            uint8 quorumNumber = uint8(allInclusiveQuorumNumbers[i]);
            assertEq(checkSignaturesIndices.quorumApkIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1);
            assertEq(checkSignaturesIndices.totalStakeIndices[i], expectedOperatorOverallIndices[quorumNumber].length - 1);
        }

        // assert the indices are zero because there have been no kicks or stake updates
        for (uint i = 0; i < nonSignerOperatorIds.length; i++) {
            assertEq(checkSignaturesIndices.nonSignerQuorumBitmapIndices[i], 0);
            for (uint j = 0; j < checkSignaturesIndices.nonSignerStakeIndices[i].length; j++) {
                assertEq(checkSignaturesIndices.nonSignerStakeIndices[i][j], 0);
            }
        }
    }

    function _registerRandomOperators(uint256 pseudoRandomNumber) internal returns(OperatorMetadata[] memory, uint256[][] memory) {
        OperatorMetadata[] memory operatorMetadatas = new OperatorMetadata[](maxOperatorsToRegister);
        for (uint i = 0; i < operatorMetadatas.length; i++) {
            // limit to 16 quorums so we don't run out of gas, make them all register for quorum 0 as well
            operatorMetadatas[i].quorumBitmap = uint256(keccak256(abi.encodePacked("quorumBitmap", pseudoRandomNumber, i))) & (maxQuorumsToRegisterFor << 1 - 1) | 1;
            operatorMetadatas[i].operator = _incrementAddress(defaultOperator, i);
            operatorMetadatas[i].pubkey = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", pseudoRandomNumber, i)));
            operatorMetadatas[i].operatorId = operatorMetadatas[i].pubkey.hashG1Point();
            operatorMetadatas[i].stakes = new uint96[](maxQuorumsToRegisterFor);
            for (uint j = 0; j < maxQuorumsToRegisterFor; j++) {
                operatorMetadatas[i].stakes[j] = uint96(uint64(uint256(keccak256(abi.encodePacked("stakes", pseudoRandomNumber, i, j)))));
            }
        }

        // get the index in quorumBitmaps of each operator in each quorum in the order they will register
        uint256[][] memory expectedOperatorOverallIndices = new uint256[][](numQuorums);
        for (uint i = 0; i < numQuorums; i++) {
            uint32 numOperatorsInQuorum;
            // for each quorumBitmap, check if the i'th bit is set
            for (uint j = 0; j < operatorMetadatas.length; j++) {
                if (operatorMetadatas[j].quorumBitmap >> i & 1 == 1) {
                    numOperatorsInQuorum++;
                }
            }
            expectedOperatorOverallIndices[i] = new uint256[](numOperatorsInQuorum);
            uint256 numOperatorCounter;
            for (uint j = 0; j < operatorMetadatas.length; j++) {
                if (operatorMetadatas[j].quorumBitmap >> i & 1 == 1) {
                    expectedOperatorOverallIndices[i][numOperatorCounter] = j;
                    numOperatorCounter++;
                }
            }
        }

        // register operators
        for (uint i = 0; i < operatorMetadatas.length; i++) {
            cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);
            
            _registerOperatorWithCoordinator(operatorMetadatas[i].operator, operatorMetadatas[i].quorumBitmap, operatorMetadatas[i].pubkey, operatorMetadatas[i].stakes);
        }

        return (operatorMetadatas, expectedOperatorOverallIndices);
    }

    function _assertExpectedOperators(
        bytes memory quorumNumbers,
        BLSOperatorStateRetriever.Operator[][] memory operators,
        uint256[][] memory expectedOperatorOverallIndices,
        OperatorMetadata[] memory operatorMetadatas
    ) internal {
        // for each quorum
        for (uint j = 0; j < quorumNumbers.length; j++) {
            // make sure the each operator id is correct
            for (uint k = 0; k < operators[j].length; k++) {
                uint8 quorumNumber = uint8(quorumNumbers[j]);
                assertEq(operators[j][k].operatorId, operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][k]].operatorId);
                assertEq(operators[j][k].stake, operatorMetadatas[expectedOperatorOverallIndices[quorumNumber][k]].stakes[quorumNumber]);
            }
        }
    }
}