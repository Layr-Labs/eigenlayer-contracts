// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../utils/MockAVSDeployer.sol";
import "./util/G2Operations.sol";
import "../../contracts/middleware/BLSSignatureChecker.sol";

contract FFIBase is MockAVSDeployer, G2Operations {
    using BN254 for BN254.G1Point;

    uint256[] public nonSignerPrivateKeys;

    function _getRandomNonSignerStakeAndSignatures(
        uint64 pseudoRandomNumber, 
        uint64 numOperators, 
        uint64 numNonSigners, 
        uint64 numQuorums,
        uint256 setQuorumBitmap
    ) internal returns(
        bytes32 msgHash,
        bytes memory allQuorums,
        uint32 referenceBlockNumber, 
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) {

        nonSignerStakesAndSignature.quorumApks = new BN254.G1Point[](numQuorums);
        nonSignerStakesAndSignature.nonSignerPubkeys = new BN254.G1Point[](numNonSigners);
        bytes32[] memory nonSignerOperatorIds = new bytes32[](numNonSigners);
        
        msgHash = keccak256(abi.encodePacked("msgHash", pseudoRandomNumber));
        BN254.G1Point memory sigma;
        BN254.G2Point memory aggSignerApkG2;

        for(uint256 i = 0; i < numOperators; i++) {

            uint256 privKey;
            BN254.G1Point memory pubkey;
            if(i < numNonSigners) {
                privKey = nonSignerPrivateKeys[i];
                pubkey = BN254.generatorG1().scalar_mul(privKey);
                nonSignerStakesAndSignature.nonSignerPubkeys[i] = pubkey;
                nonSignerOperatorIds[i] = pubkey.hashG1Point();
            } else {
                privKey = uint256(keccak256(abi.encodePacked("operatorPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
                pubkey = BN254.generatorG1().scalar_mul(privKey);
            }

            uint256 quorumBitmap;
            if(setQuorumBitmap > 0){
                quorumBitmap = setQuorumBitmap;
            } else {
                quorumBitmap = privKey & (1 << numQuorums - 1) | 1;
                if(i < numQuorums) {
                    quorumBitmap = quorumBitmap | (1 << i);
                }
            }

            {
                address operator = _incrementAddress(defaultOperator, i);
                cheats.roll(registrationBlockNumber + blocksBetweenRegistrations * i);
                _registerOperatorWithCoordinator(operator, quorumBitmap, pubkey);
            }

            bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

            for(uint256 j = 0; j < quorumNumbers.length; j++) {
                uint8 quorumNumber = uint8(quorumNumbers[j]);
                nonSignerStakesAndSignature.quorumApks[quorumNumber] = nonSignerStakesAndSignature.quorumApks[quorumNumber].plus(pubkey);
            }
            
            if(i >= numNonSigners) {
                BN254.G1Point memory sig = BN254.hashToG1(msgHash).scalar_mul(privKey);
                sigma = sigma.plus(sig.scalar_mul(quorumNumbers.length));

                BN254.G2Point memory pkG2 = G2Operations.mulGen(privKey);
                aggSignerApkG2 = G2Operations.add(aggSignerApkG2, G2Operations.mul(pkG2, quorumNumbers.length));
            } 
        }
        
        allQuorums = BitmapUtils.bitmapToBytesArray((1 << numQuorums) - 1);
        referenceBlockNumber = registrationBlockNumber + blocksBetweenRegistrations * uint32(numOperators) + 1;
        cheats.roll(referenceBlockNumber + 100);
        
        BLSOperatorStateRetriever.CheckSignaturesIndices memory checkSignaturesIndices = operatorStateRetriever.getCheckSignaturesIndices(
            registryCoordinator,
            referenceBlockNumber, 
            allQuorums, 
            nonSignerOperatorIds
        );

        nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices = checkSignaturesIndices.nonSignerQuorumBitmapIndices;
        nonSignerStakesAndSignature.apkG2 = aggSignerApkG2;
        nonSignerStakesAndSignature.sigma = sigma;
        nonSignerStakesAndSignature.quorumApkIndices = checkSignaturesIndices.quorumApkIndices;
        nonSignerStakesAndSignature.totalStakeIndices = checkSignaturesIndices.totalStakeIndices;
        nonSignerStakesAndSignature.nonSignerStakeIndices = checkSignaturesIndices.nonSignerStakeIndices;

        return (msgHash, allQuorums, referenceBlockNumber, nonSignerStakesAndSignature);
    }

    function _setNonSignerPrivKeys(uint256 numNonSigners, uint256 pseudoRandomNumber) internal {
        nonSignerPrivateKeys = new uint256[](numNonSigners);
        for (uint256 i = 0; i < numNonSigners; i++) {
            nonSignerPrivateKeys[i] = uint256(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
            uint256 j = 0;
            if(i != 0) {
                while(BN254.generatorG1().scalar_mul(nonSignerPrivateKeys[i]).hashG1Point() <= BN254.generatorG1().scalar_mul(nonSignerPrivateKeys[i-1]).hashG1Point()){
                    nonSignerPrivateKeys[i] = uint256(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, j))) % BN254.FR_MODULUS;
                    j++;
                }
            }
        }
    }

}