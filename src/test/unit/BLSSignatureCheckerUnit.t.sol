// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/BLSSignatureChecker.sol";
import "../utils/BLSMockAVSDeployer.sol";

contract BLSSignatureCheckerUnitTests is BLSMockAVSDeployer {
    using BN254 for BN254.G1Point;

    BLSSignatureChecker blsSignatureChecker;

    function setUp() virtual public {
        _setUpBLSMockAVSDeployer();

        blsSignatureChecker = new BLSSignatureChecker(registryCoordinator);
    }

    // this test checks that a valid signature from maxOperatorsToRegister with a random number of nonsigners is checked
    // correctly on the BLSSignatureChecker contract when all operators are only regsitered for a single quorum and
    // the signature is only checked for stakes on that quorum
    function testBLSSignatureChecker_SingleQuorum_Valid(uint256 pseudoRandomNumber) public { 
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 1);

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        uint256 gasBefore = gasleft();
        (
            BLSSignatureChecker.QuorumStakeTotals memory quorumStakeTotals,
            /* bytes32 signatoryRecordHash */
        ) = blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);
        assertTrue(quorumStakeTotals.signedStakeForQuorum[0] > 0);

        // 0 nonSigners: 159908
        // 1 nonSigner: 178683
        // 2 nonSigners: 197410
    }

    // this test checks that a valid signature from maxOperatorsToRegister with a random number of nonsigners is checked
    // correctly on the BLSSignatureChecker contract when all operators are registered for the first 100 quorums
    // and the signature is only checked for stakes on those quorums
    function testBLSSignatureChecker_100Quorums_Valid(uint256 pseudoRandomNumber) public { 
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 1);

        // 100 set bits
        uint256 quorumBitmap = (1 << 100) - 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        nonSignerStakesAndSignature.sigma = sigma.scalar_mul(quorumNumbers.length);
        nonSignerStakesAndSignature.apkG2 = oneHundredQuorumApkG2;

        uint256 gasBefore = gasleft();
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);
    }

    function testBLSSignatureChecker_IncorrectQuorumBitmapIndex_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
        
        // record a quorumBitmap update
        registryCoordinator.recordOperatorQuorumBitmapUpdate(nonSignerStakesAndSignature.nonSignerPubkeys[0].hashG1Point(), uint192(quorumBitmap | 2));

        // set the nonSignerQuorumBitmapIndices to a different value
        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices[0] = 1;

        cheats.expectRevert("BLSRegistryCoordinatorWithIndices.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }

    function testBLSSignatureChecker_IncorrectTotalStakeIndex_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
        
        // set the totalStakeIndices to a different value
        nonSignerStakesAndSignature.totalStakeIndices[0] = 0;

        cheats.expectRevert("StakeRegistry._validateOperatorStakeAtBlockNumber: there is a newer operatorStakeUpdate available before blockNumber");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }

    function testBLSSignatureChecker_IncorrectNonSignerStakeIndex_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        bytes32 nonSignerOperatorId = nonSignerStakesAndSignature.nonSignerPubkeys[0].hashG1Point();
        
        // record a stake update
        stakeRegistry.recordOperatorStakeUpdate(
            nonSignerOperatorId, 
            uint8(quorumNumbers[0]), 
            IStakeRegistry.OperatorStakeUpdate({
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0,
                stake: 1234
            })
        );
        
        // set the nonSignerStakeIndices to a different value
        nonSignerStakesAndSignature.nonSignerStakeIndices[0][0] = 1;

        cheats.expectRevert("StakeRegistry._validateOperatorStakeAtBlockNumber: operatorStakeUpdate is from after blockNumber");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );

    }

    function testBLSSignatureChecker_IncorrectQuorumAPKIndex_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        // set the quorumApkIndices to a different value
        nonSignerStakesAndSignature.quorumApkIndices[0] = 0;

        cheats.expectRevert("BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: not latest apk update");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }

    function testBLSSignatureChecker_IncorrectQuorumAPK_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
        
        // set the quorumApk to a different value
        nonSignerStakesAndSignature.quorumApks[0] = nonSignerStakesAndSignature.quorumApks[0].negate();

        cheats.expectRevert("BLSSignatureChecker.checkSignatures: quorumApk hash in storage does not match provided quorum apk");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }

    function testBLSSignatureChecker_IncorrectSignature_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
        
        // set the sigma to a different value
        nonSignerStakesAndSignature.sigma = nonSignerStakesAndSignature.sigma.negate();

        cheats.expectRevert("BLSSignatureChecker.checkSignatures: signature is invalid");
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }

    function testBLSSignatureChecker_InvalidSignature_Reverts(uint256 pseudoRandomNumber) public {
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 2) + 1;

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
        
        // set the sigma to a different value
        nonSignerStakesAndSignature.sigma.X++;

        // expect a non-specific low-level revert, since this call will ultimately fail as part of the precompile call
        cheats.expectRevert();
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }
}