// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";

contract BLSOperatorStateRetrieverUnitTests is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    uint8 maxQuorumsToRegisterFor = 4;
    uint256 maxOperatorsToRegister = 100;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();
    }

    function testGetOperatorState_Valid(uint256 pseudoRandomNumber) public {
        uint32 registrationBlockNumber = 100;
        uint32 blocksBetweenRegistrations = 10;

        uint256[] memory quorumBitmaps = new uint256[](maxOperatorsToRegister);
        for (uint i = 0; i < quorumBitmaps.length; i++) {
            // limit to 16 quorums so we don't run out of gas, make them all register for quorum 0 as well
            quorumBitmaps[i] = uint256(keccak256(abi.encodePacked("quorumBitmap", pseudoRandomNumber, i))) & (maxQuorumsToRegisterFor << 1 - 1) | 1;
        }

        uint96[] memory stakes = new uint96[](maxOperatorsToRegister);
        for (uint i = 0; i < stakes.length; i++) {
            stakes[i] = uint96(uint64(uint256(keccak256(abi.encodePacked("stakes", pseudoRandomNumber, i)))));
        }

        bytes32[] memory operatorIds = new bytes32[](maxOperatorsToRegister);

        // get the index in quorumBitmaps of each operator in each quorum in the order they will register
        uint256[][] memory expectedOperatorOverallIndices = new uint256[][](numQuorums);
        for (uint i = 0; i < numQuorums; i++) {
            uint32 numOperatorsInQuorum;
            // for each quorumBitmap, check if the i'th bit is set
            for (uint j = 0; j < quorumBitmaps.length; j++) {
                if (quorumBitmaps[j] >> i & 1 == 1) {
                    numOperatorsInQuorum++;
                }
            }
            expectedOperatorOverallIndices[i] = new uint256[](numOperatorsInQuorum);
            uint256 numOperatorCounter;
            for (uint j = 0; j < quorumBitmaps.length; j++) {
                if (quorumBitmaps[j] >> i & 1 == 1) {
                    expectedOperatorOverallIndices[i][numOperatorCounter] = j;
                    numOperatorCounter++;
                }
            }
        }

        // register operators
        for (uint i = 0; i < quorumBitmaps.length; i++) {
            cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);

            BN254.G1Point memory pubkey = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", pseudoRandomNumber, i)));
            operatorIds[i] = pubkey.hashG1Point();
            address operator = _incrementAddress(defaultOperator, i);
            
            _registerOperatorWithCoordinator(operator, quorumBitmaps[i], pubkey, stakes[i]);
        }

        for (uint i = 0; i < quorumBitmaps.length; i++) {
            uint32 blockNumber = uint32(registrationBlockNumber + blocksBetweenRegistrations * i);

            (uint256 quorumBitmap, BLSOperatorStateRetriever.Operator[][] memory operators) = 
                operatorStateRetriever.getOperatorState(registryCoordinator, operatorIds[i], blockNumber);

            assertEq(quorumBitmaps[i], quorumBitmap);
            bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);
            
            // for each quorum
            for (uint j = 0; j < quorumNumbers.length; j++) {
                uint8 quorumNumber = uint8(quorumNumbers[j]);
                // make sure the each operator id is correct
                for (uint k = 0; k < operators[j].length; k++) {
                    assertEq(operators[j][k].operatorId, operatorIds[expectedOperatorOverallIndices[quorumNumber][k]]);
                    assertEq(operators[j][k].stake, stakes[expectedOperatorOverallIndices[quorumNumber][k]]);
                }
            }
        }

        uint256 operatorIndexToDeregister = pseudoRandomNumber % maxOperatorsToRegister;
        bytes memory quorumNumbersToDeregister = BitmapUtils.bitmapToBytesArray(quorumBitmaps[operatorIndexToDeregister]);
        BN254.G1Point memory pubkeyToDeregister = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", pseudoRandomNumber, operatorIndexToDeregister)));
        bytes32[] memory operatorIdsToSwap = new bytes32[](quorumNumbersToDeregister.length);
        for (uint i = 0; i < quorumNumbersToDeregister.length; i++) {
            uint8 quorumNumber = uint8(quorumNumbersToDeregister[i]);
            operatorIdsToSwap[i] = operatorIds[expectedOperatorOverallIndices[quorumNumber][expectedOperatorOverallIndices[quorumNumber].length - 1]];
        }

        uint32 deregistrationBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * (uint32(quorumBitmaps.length) + 1);
        cheats.roll(deregistrationBlockNumber);

        cheats.prank(_incrementAddress(defaultOperator, operatorIndexToDeregister));
        registryCoordinator.deregisterOperatorWithCoordinator(quorumNumbersToDeregister, pubkeyToDeregister, operatorIdsToSwap);

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

        BLSOperatorStateRetriever.Operator[][] memory operatorsAfterDeregistration = 
            operatorStateRetriever.getOperatorState(registryCoordinator, allQuorumNumbers, deregistrationBlockNumber);
        
        // for each quorum
        for (uint j = 0; j < allQuorumNumbers.length; j++) {
            uint8 quorumNumber = uint8(allQuorumNumbers[j]);
            // make sure the each operator id is correct
            for (uint k = 0; k < operatorsAfterDeregistration[j].length; k++) {
                assertEq(operatorsAfterDeregistration[j][k].operatorId, operatorIds[expectedOperatorOverallIndices[quorumNumber][k]]);
                assertEq(operatorsAfterDeregistration[j][k].stake, stakes[expectedOperatorOverallIndices[quorumNumber][k]]);
            }
        }
    }
}