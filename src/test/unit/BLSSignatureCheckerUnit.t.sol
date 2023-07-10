// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/BLSSignatureChecker.sol";
import "../utils/MockAVSDeployer.sol";

contract BLSSignatureCheckerUnitTests is MockAVSDeployer {
    using BN254 for BN254.G1Point;

    BLSSignatureChecker blsSignatureChecker;
    bytes32 msgHash = keccak256(abi.encodePacked("hello world"));
    uint256 aggSignerPrivKey = 69;
    BN254.G2Point aggSignerApkG2;
    BN254.G1Point sigma;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();

        blsSignatureChecker = new BLSSignatureChecker(registryCoordinator);

        // aggSignerPrivKey*g2
        aggSignerApkG2.X[1] = 19101821850089705274637533855249918363070101489527618151493230256975900223847;
        aggSignerApkG2.X[0] = 5334410886741819556325359147377682006012228123419628681352847439302316235957;
        aggSignerApkG2.Y[1] = 354176189041917478648604979334478067325821134838555150300539079146482658331;
        aggSignerApkG2.Y[0] = 4185483097059047421902184823581361466320657066600218863748375739772335928910;

        sigma = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);
    }

    function testBLSSignatureChecker_SingleQuorum_Valid(uint256 pseudoRandomNumber) public { 
        uint256 numNonSigners = pseudoRandomNumber % (maxOperatorsToRegister - 1);

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) = 
            _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        uint256 gasBefore = gasleft();
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfter = gasleft();
        emit log_named_uint("gasUsed", gasBefore - gasAfter);

        // 0 nonSigners: 159908
        // 1 nonSigner: 178683
        // 2 nonSigners: 197410
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

        cheats.expectRevert("BLSRegistryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex: quorumBitmapUpdate is from after blockNumber");
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

        cheats.expectRevert();
        blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
    }
    
    function _generateSignerAndNonSignerPrivateKeys(uint256 pseudoRandomNumber, uint256 numSigners, uint256 numNonSigners) internal view returns (uint256[] memory, uint256[] memory) {
        uint256[] memory signerPrivateKeys = new uint256[](numSigners);
        // generate numSigners numbers that add up to aggSignerPrivKey mod BN254.FR_MODULUS
        uint256 sum = 0;
        for (uint i = 0; i < numSigners - 1; i++) {
            signerPrivateKeys[i] = uint256(keccak256(abi.encodePacked("signerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
            sum = addmod(sum, signerPrivateKeys[i], BN254.FR_MODULUS);
        }
        // signer private keys need to add to aggSignerPrivKey
        signerPrivateKeys[numSigners - 1] = addmod(aggSignerPrivKey, BN254.FR_MODULUS - sum % BN254.FR_MODULUS, BN254.FR_MODULUS);

        uint256[] memory nonSignerPrivateKeys = new uint256[](numNonSigners);
        for (uint i = 0; i < numNonSigners; i++) {
            nonSignerPrivateKeys[i] = uint256(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
        }

        return (signerPrivateKeys, nonSignerPrivateKeys);
    }

    function _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(uint256 pseudoRandomNumber, uint256 numNonSigners, uint256 quorumBitmap) internal returns(uint32, BLSSignatureChecker.NonSignerStakesAndSignature memory) {
        (uint256[] memory signerPrivateKeys, uint256[] memory nonSignerPrivateKeys) = _generateSignerAndNonSignerPrivateKeys(pseudoRandomNumber, maxOperatorsToRegister - numNonSigners, numNonSigners);

        // randomly combine signer and non-signer private keys
        uint256[] memory privateKeys = new uint256[](maxOperatorsToRegister);
        // generate addresses and public keys
        address[] memory operators = new address[](maxOperatorsToRegister);
        BN254.G1Point[] memory pubkeys = new BN254.G1Point[](maxOperatorsToRegister);
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature;
        nonSignerStakesAndSignature.quorumApks = new BN254.G1Point[](1);
        nonSignerStakesAndSignature.nonSignerPubkeys = new BN254.G1Point[](numNonSigners);
        bytes32[] memory nonSignerOperatorIds = new bytes32[](numNonSigners);
        {
            uint256 signerIndex = 0;
            uint256 nonSignerIndex = 0;
            for (uint i = 0; i < maxOperatorsToRegister; i++) {
                uint256 randomSeed = uint256(keccak256(abi.encodePacked("privKeyCombination", i)));
                if (randomSeed % 2 == 0 && signerIndex < signerPrivateKeys.length) {
                    privateKeys[i] = signerPrivateKeys[signerIndex];
                    signerIndex++;
                } else if (nonSignerIndex < nonSignerPrivateKeys.length) {
                    privateKeys[i] = nonSignerPrivateKeys[nonSignerIndex];
                    nonSignerStakesAndSignature.nonSignerPubkeys[nonSignerIndex] = BN254.generatorG1().scalar_mul(privateKeys[i]);
                    nonSignerOperatorIds[nonSignerIndex] = nonSignerStakesAndSignature.nonSignerPubkeys[nonSignerIndex].hashG1Point();
                    nonSignerIndex++;
                } else {
                    privateKeys[i] = signerPrivateKeys[signerIndex];
                    signerIndex++;
                }

                operators[i] = _incrementAddress(defaultOperator, i);
                pubkeys[i] = BN254.generatorG1().scalar_mul(privateKeys[i]);
                nonSignerStakesAndSignature.quorumApks[0] = nonSignerStakesAndSignature.quorumApks[0].plus(pubkeys[i]);
            }
        }

        // register all operators for the first quorum
        for (uint i = 0; i < maxOperatorsToRegister; i++) {
            cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);
            _registerOperatorWithCoordinator(operators[i], quorumBitmap, pubkeys[i], defaultStake);
        }

        uint32 referenceBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * uint32(maxOperatorsToRegister) + 1;
        cheats.roll(referenceBlockNumber + 100);
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);       

        BLSOperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices = operatorStateRetriever.getCheckSignaturesIndices(
            registryCoordinator,
            referenceBlockNumber, 
            quorumNumbers, 
            nonSignerOperatorIds
        );

        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices = checkSignaturesIndices.nonSignerQuorumBitmapIndices;
        nonSignerStakesAndSignature.apkG2 = aggSignerApkG2;
        nonSignerStakesAndSignature.sigma = sigma;
        nonSignerStakesAndSignature.quorumApkIndices = checkSignaturesIndices.quorumApkIndices;
        nonSignerStakesAndSignature.totalStakeIndices = checkSignaturesIndices.totalStakeIndices;
        nonSignerStakesAndSignature.nonSignerStakeIndices = checkSignaturesIndices.nonSignerStakeIndices;

        return (referenceBlockNumber, nonSignerStakesAndSignature);
    }

}