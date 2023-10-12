// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IBLSSignatureChecker.sol";

contract BLSSigCheckerExperimental is IBLSSignatureChecker {
    using BN254 for BN254.G1Point;
    
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    IRegistryCoordinator public immutable registryCoordinator;
    IStakeRegistry public immutable stakeRegistry;
    IBLSPubkeyRegistry public immutable blsPubkeyRegistry;

    constructor(IBLSRegistryCoordinatorWithIndices _registryCoordinator) {
        registryCoordinator = IRegistryCoordinator(_registryCoordinator);
        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
    }

    function checkSignatures_reg(
        bytes32 msgHash, 
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) 
        public 
        view
        returns (
            QuorumStakeTotals memory,
            bytes32
        )
    {   
        BN254.G1Point memory apk = BN254.G1Point(0, 0);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            require(
                bytes24(nonSignerStakesAndSignature.quorumApks[i].hashG1Point()) == 
                    blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(
                        uint8(quorumNumbers[i]), 
                        referenceBlockNumber, 
                        nonSignerStakesAndSignature.quorumApkIndices[i]
                    ),
                "BLSSignatureChecker.checkSignatures: quorumApk hash in storage does not match provided quorum apk"
            );
            apk = apk.plus(nonSignerStakesAndSignature.quorumApks[i]);
        }
        
        QuorumStakeTotals memory quorumStakeTotals;
        quorumStakeTotals.totalStakeForQuorum = new uint96[](quorumNumbers.length);
        quorumStakeTotals.signedStakeForQuorum = new uint96[](quorumNumbers.length);
        bytes32[] memory nonSignerPubkeyHashes = new bytes32[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
        {
            uint256[] memory nonSignerQuorumBitmaps = new uint256[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
            {
                uint256 signingQuorumBitmap = BitmapUtils.bytesArrayToBitmap(quorumNumbers);

                for (uint i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    nonSignerPubkeyHashes[i] = nonSignerStakesAndSignature.nonSignerPubkeys[i].hashG1Point();
                    
                    if (i != 0) {
                        require(uint256(nonSignerPubkeyHashes[i]) > uint256(nonSignerPubkeyHashes[i - 1]), "BLSSignatureChecker.checkSignatures: nonSignerPubkeys not sorted");
                    }

                    nonSignerQuorumBitmaps[i] = 
                        registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                            nonSignerPubkeyHashes[i], 
                            referenceBlockNumber, 
                            nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices[i]
                        );
                    
                    apk = apk.plus(
                        nonSignerStakesAndSignature.nonSignerPubkeys[i]
                            .negate()
                            .scalar_mul(
                                BitmapUtils.countNumOnes(nonSignerQuorumBitmaps[i] & signingQuorumBitmap) 
                            )
                    );
                }
            }
            for (uint8 quorumNumberIndex = 0; quorumNumberIndex < quorumNumbers.length;) {
                uint8 quorumNumber = uint8(quorumNumbers[quorumNumberIndex]);
                quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex] = 
                    stakeRegistry.getTotalStakeAtBlockNumberFromIndex(quorumNumber, referenceBlockNumber, nonSignerStakesAndSignature.totalStakeIndices[quorumNumberIndex]);
                quorumStakeTotals.signedStakeForQuorum[quorumNumberIndex] = quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex];
                for (uint32 i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    uint32 nonSignerForQuorumIndex = 0;
                    if (BitmapUtils.numberIsInBitmap(nonSignerQuorumBitmaps[i], quorumNumber)) {
                        quorumStakeTotals.signedStakeForQuorum[quorumNumberIndex] -=
                            stakeRegistry.getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(
                                quorumNumber,
                                referenceBlockNumber,
                                nonSignerPubkeyHashes[i],
                                nonSignerStakesAndSignature.nonSignerStakeIndices[quorumNumberIndex][nonSignerForQuorumIndex]
                            );
                        unchecked {
                            ++nonSignerForQuorumIndex;
                        }
                    }
                }

                unchecked {
                    ++quorumNumberIndex;
                }
            }
        }
        {
            (bool pairingSuccessful, bool signatureIsValid) = trySignatureAndApkVerification(
                msgHash, 
                apk, 
                nonSignerStakesAndSignature.apkG2, 
                nonSignerStakesAndSignature.sigma
            );
            require(pairingSuccessful, "BLSSignatureChecker.checkSignatures: pairing precompile call failed");
            require(signatureIsValid, "BLSSignatureChecker.checkSignatures: signature is invalid");
        }
        bytes32 signatoryRecordHash = keccak256(abi.encodePacked(referenceBlockNumber, nonSignerPubkeyHashes));

        return (quorumStakeTotals, signatoryRecordHash);
    }

    function checkSignatures_tiny(
        bytes32 msgHash, 
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) 
        public 
        view
        returns (
            QuorumStakeTotals memory,
            bytes32
        )
    {   
        BN254.G1Point memory apk = BN254.G1Point(0, 0);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            require(
                bytes24(nonSignerStakesAndSignature.quorumApks[i].hashG1Point()) == 
                    blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(
                        uint8(quorumNumbers[i]), 
                        referenceBlockNumber, 
                        nonSignerStakesAndSignature.quorumApkIndices[i]
                    ),
                "BLSSignatureChecker.checkSignatures: quorumApk hash in storage does not match provided quorum apk"
            );
            apk = apk.plus(nonSignerStakesAndSignature.quorumApks[i]);
        }
        
        QuorumStakeTotals memory quorumStakeTotals;
        quorumStakeTotals.totalStakeForQuorum = new uint96[](quorumNumbers.length);
        quorumStakeTotals.signedStakeForQuorum = new uint96[](quorumNumbers.length);
        bytes32[] memory nonSignerPubkeyHashes = new bytes32[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
        {
            uint256[] memory nonSignerQuorumBitmaps = new uint256[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
            {
                uint256 signingQuorumBitmap = BitmapUtils.bytesArrayToBitmap(quorumNumbers);

                for (uint i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    nonSignerPubkeyHashes[i] = nonSignerStakesAndSignature.nonSignerPubkeys[i].hashG1Point();
                    
                    if (i != 0) {
                        require(uint256(nonSignerPubkeyHashes[i]) > uint256(nonSignerPubkeyHashes[i - 1]), "BLSSignatureChecker.checkSignatures: nonSignerPubkeys not sorted");
                    }

                    nonSignerQuorumBitmaps[i] = 
                        registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                            nonSignerPubkeyHashes[i], 
                            referenceBlockNumber, 
                            nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices[i]
                        );
                    
                    apk = apk.plus(
                        nonSignerStakesAndSignature.nonSignerPubkeys[i]
                            .negate()
                            .scalar_mul_tiny(
                                BitmapUtils.countNumOnes(nonSignerQuorumBitmaps[i] & signingQuorumBitmap) 
                            )
                    );
                }
            }
            for (uint8 quorumNumberIndex = 0; quorumNumberIndex < quorumNumbers.length;) {
                uint8 quorumNumber = uint8(quorumNumbers[quorumNumberIndex]);
                quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex] = 
                    stakeRegistry.getTotalStakeAtBlockNumberFromIndex(quorumNumber, referenceBlockNumber, nonSignerStakesAndSignature.totalStakeIndices[quorumNumberIndex]);
                quorumStakeTotals.signedStakeForQuorum[quorumNumberIndex] = quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex];
                for (uint32 i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    uint32 nonSignerForQuorumIndex = 0;
                    if (BitmapUtils.numberIsInBitmap(nonSignerQuorumBitmaps[i], quorumNumber)) {
                        quorumStakeTotals.signedStakeForQuorum[quorumNumberIndex] -=
                            stakeRegistry.getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(
                                quorumNumber,
                                referenceBlockNumber,
                                nonSignerPubkeyHashes[i],
                                nonSignerStakesAndSignature.nonSignerStakeIndices[quorumNumberIndex][nonSignerForQuorumIndex]
                            );
                        unchecked {
                            ++nonSignerForQuorumIndex;
                        }
                    }
                }

                unchecked {
                    ++quorumNumberIndex;
                }
            }
        }
        {
            (bool pairingSuccessful, bool signatureIsValid) = trySignatureAndApkVerification(
                msgHash, 
                apk, 
                nonSignerStakesAndSignature.apkG2, 
                nonSignerStakesAndSignature.sigma
            );
            require(pairingSuccessful, "BLSSignatureChecker.checkSignatures: pairing precompile call failed");
            require(signatureIsValid, "BLSSignatureChecker.checkSignatures: signature is invalid");
        }
        bytes32 signatoryRecordHash = keccak256(abi.encodePacked(referenceBlockNumber, nonSignerPubkeyHashes));

        return (quorumStakeTotals, signatoryRecordHash);
    }

    function trySignatureAndApkVerification(
        bytes32 msgHash,
        BN254.G1Point memory apk,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory sigma
    ) public view returns(bool pairingSuccessful, bool siganatureIsValid) {
        uint256 gamma = uint256(keccak256(abi.encodePacked(msgHash, apk.X, apk.Y, apkG2.X[0], apkG2.X[1], apkG2.Y[0], apkG2.Y[1], sigma.X, sigma.Y))) % BN254.FR_MODULUS;
        (pairingSuccessful, siganatureIsValid) = BN254.safePairing(
                sigma.plus(apk.scalar_mul(gamma)),
                BN254.negGeneratorG2(),
                BN254.hashToG1(msgHash).plus(BN254.generatorG1().scalar_mul(gamma)),
                apkG2,
                PAIRING_EQUALITY_CHECK_GAS
            );
    }

    function checkSignatures(
        bytes32 msgHash, 
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external view returns (
        QuorumStakeTotals memory,
        bytes32
    ) {}
    
}
