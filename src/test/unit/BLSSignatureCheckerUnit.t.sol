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

        uint256 quorumBitmap = 1;
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        // register all operators for the first quorum
        for (uint i = 0; i < maxOperatorsToRegister; i++) {
            cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);
            _registerOperatorWithCoordinator(operators[i], quorumBitmap, pubkeys[i], defaultStake);
        }

        uint32 referenceBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * uint32(maxOperatorsToRegister) + 1;
        cheats.roll(referenceBlockNumber + 100);

        BLSOperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices = operatorStateRetriever.getCheckSignaturesIndices(
            registryCoordinator,
            referenceBlockNumber, 
            quorumNumbers, 
            nonSignerOperatorIds
        );

        // struct NonSignerStakesAndSignature {
        //     uint32[] nonSignerQuorumBitmapIndices;
        //     BN254.G1Point[] nonSignerPubkeys;
        //     BN254.G1Point[] quorumApks;
        //     BN254.G2Point apkG2;
        //     BN254.G1Point sigma;
        //     uint32[] quorumApkIndices;
        //     uint32[] totalStakeIndices;  
        //     uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
        // }

        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices = checkSignaturesIndices.nonSignerQuorumBitmapIndices;
        nonSignerStakesAndSignature.apkG2 = aggSignerApkG2;
        nonSignerStakesAndSignature.sigma = sigma;
        nonSignerStakesAndSignature.quorumApkIndices = checkSignaturesIndices.quorumApkIndices;
        nonSignerStakesAndSignature.totalStakeIndices = checkSignaturesIndices.totalStakeIndices;
        nonSignerStakesAndSignature.nonSignerStakeIndices = checkSignaturesIndices.nonSignerStakeIndices;

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

}