// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./util/G2Operations.sol";
import "../utils/MockAVSDeployer.sol";
import "../../contracts/middleware/BLSSignatureChecker.sol";


contract BLSSignatureCheckerFFITests is MockAVSDeployer, G2Operations {

    using BN254 for BN254.G1Point;

    bytes32 msgHash = keccak256(abi.encodePacked("hello world"));
    uint256 aggSignerPrivKey;
    BN254.G2Point aggSignerApkG2;
    BN254.G2Point oneHundredQuorumApkG2;
    BN254.G1Point sigma;

    BLSSignatureChecker blsSignatureChecker;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();

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

    function _setAggregatePublicKeysAndSignature(uint256 pseudoRandomNumber) internal {
        if(pseudoRandomNumber > type(uint256).max / 100) {
            pseudoRandomNumber = type(uint256).max / 100;
        }
        aggSignerPrivKey = pseudoRandomNumber;
        aggSignerApkG2 = G2Operations.mul(aggSignerPrivKey);
        oneHundredQuorumApkG2 = G2Operations.mul(100 * aggSignerPrivKey);
        sigma = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);
    }

    function _generateSignerAndNonSignerPrivateKeys(uint256 pseudoRandomNumber, uint256 numSigners, uint256 numNonSigners) internal returns (uint256[] memory, uint256[] memory) {
        _setAggregatePublicKeysAndSignature(pseudoRandomNumber);

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
        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        // randomly combine signer and non-signer private keys
        uint256[] memory privateKeys = new uint256[](maxOperatorsToRegister);
        // generate addresses and public keys
        address[] memory operators = new address[](maxOperatorsToRegister);
        BN254.G1Point[] memory pubkeys = new BN254.G1Point[](maxOperatorsToRegister);
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature;
        nonSignerStakesAndSignature.quorumApks = new BN254.G1Point[](quorumNumbers.length);
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

                // add the public key to each quorum
                for (uint j = 0; j < nonSignerStakesAndSignature.quorumApks.length; j++) {
                    nonSignerStakesAndSignature.quorumApks[j] = nonSignerStakesAndSignature.quorumApks[j].plus(pubkeys[i]);
                }
            }
        }

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

        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices = checkSignaturesIndices.nonSignerQuorumBitmapIndices;
        nonSignerStakesAndSignature.apkG2 = aggSignerApkG2;
        nonSignerStakesAndSignature.sigma = sigma;
        nonSignerStakesAndSignature.quorumApkIndices = checkSignaturesIndices.quorumApkIndices;
        nonSignerStakesAndSignature.totalStakeIndices = checkSignaturesIndices.totalStakeIndices;
        nonSignerStakesAndSignature.nonSignerStakeIndices = checkSignaturesIndices.nonSignerStakeIndices;

        return (referenceBlockNumber, nonSignerStakesAndSignature);
    }

}