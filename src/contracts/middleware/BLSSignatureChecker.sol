// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../libraries/MiddlewareUtils.sol";
import "../libraries/BN254.sol";
import "../libraries/BitmapUtils.sol";

import "forge-std/Test.sol";

/**
 * @title Used for checking BLS aggregate signatures from the operators of a `BLSRegistry`.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This is the contract for checking the validity of aggregate operator signatures.
 */
contract BLSSignatureChecker is Test {
    using BN254 for BN254.G1Point;    

    // DATA STRUCTURES

    struct NonSignerStakesAndSignature {
        uint32[] nonSignerQuorumBitmapIndices;
        BN254.G1Point[] nonSignerPubkeys;
        BN254.G1Point[] quorumApks;
        BN254.G2Point apkG2;
        BN254.G1Point sigma;
        uint32[] quorumApkIndices;
        uint32[] totalStakeIndices;  
        uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
    }

    /**
     * @notice this data structure is used for recording the details on the total stake of the registered
     * operators and those operators who are part of the quorum for a particular taskNumber
     */

    struct QuorumStakeTotals {
        // total stake of the operators in each quorum
        uint96[] signedStakeForQuorum;
        // total amount staked by all operators in each quorum
        uint96[] totalStakeForQuorum;
    }
    
    // CONSTANTS & IMMUTABLES

    // gas cost of multiplying 2 pairings
    // TODO: verify this
    uint256 constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    IRegistryCoordinator public immutable registryCoordinator;
    IStakeRegistry public immutable stakeRegistry;
    IBLSPubkeyRegistry public immutable blsPubkeyRegistry;

    constructor(IBLSRegistryCoordinatorWithIndices _registryCoordinator) {
        registryCoordinator = IRegistryCoordinator(_registryCoordinator);
        stakeRegistry = _registryCoordinator.stakeRegistry();
        blsPubkeyRegistry = _registryCoordinator.blsPubkeyRegistry();
    }

    /**
     * @notice This function is called by disperser when it has aggregated all the signatures of the operators
     * that are part of the quorum for a particular taskNumber and is asserting them into onchain. The function
     * checks that the claim for aggregated signatures are valid.
     *
     * The thesis of this procedure entails:
     * - getting the aggregated pubkey of all registered nodes at the time of pre-commit by the
     * disperser (represented by apk in the parameters),
     * - subtracting the pubkeys of all the signers not in the quorum (nonSignerPubkeys) and storing 
     * the output in apk to get aggregated pubkey of all operators that are part of quorum.
     * - use this aggregated pubkey to verify the aggregated signature under BLS scheme.
     * 
     * @dev Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `referenceBlockNumber`
     * is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update
     * for the total stake (or the operator) or latest before the referenceBlockNumber.
     */
    function checkSignatures(
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
        // verify the provided apk was the apk at referenceBlockNumber
        // loop through every quorumNumber and keep track of the apk
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
        
        // initialize memory for the quorumStakeTotals
        QuorumStakeTotals memory quorumStakeTotals;
        quorumStakeTotals.totalStakeForQuorum = new uint96[](quorumNumbers.length);
        quorumStakeTotals.signedStakeForQuorum = new uint96[](quorumNumbers.length);
        // the pubkeyHashes of the nonSigners
        bytes32[] memory nonSignerPubkeyHashes = new bytes32[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
        {
            // the quorumBitmaps of the nonSigners
            uint256[] memory nonSignerQuorumBitmaps = new uint256[](nonSignerStakesAndSignature.nonSignerPubkeys.length);
            {
                // the bitmap of the quorumNumbers
                uint256 signingQuorumBitmap = BitmapUtils.bytesArrayToBitmap(quorumNumbers);

                for (uint i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    nonSignerPubkeyHashes[i] = nonSignerStakesAndSignature.nonSignerPubkeys[i].hashG1Point();
                    // if (i != 0) {
                    //     require(uint256(nonSignerPubkeyHashes[i]) > uint256(nonSignerPubkeyHashes[i - 1]), "BLSSignatureChecker.checkSignatures: nonSignerPubkeys not sorted");
                    // }
                    nonSignerQuorumBitmaps[i] = 
                        registryCoordinator.getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                            nonSignerPubkeyHashes[i], 
                            referenceBlockNumber, 
                            nonSignerStakesAndSignature.nonSignerQuorumBitmapIndices[i]
                        );
                    
                    // subtract the nonSignerPubkey from the running apk to get the apk of all signers
                    apk = apk.plus(
                        nonSignerStakesAndSignature.nonSignerPubkeys[i]
                            .negate()
                            .scalar_mul(
                                BitmapUtils.countNumOnes(nonSignerQuorumBitmaps[i] & signingQuorumBitmap) // we subtract the nonSignerPubkey from each quorum that they are a part of, TODO: 
                            )
                    );
                }
            }
            // loop through each quorum number
            for (uint8 quorumNumberIndex = 0; quorumNumberIndex < quorumNumbers.length;) {
                // get the quorum number
                uint8 quorumNumber = uint8(quorumNumbers[quorumNumberIndex]);
                // get the totalStake for the quorum at the referenceBlockNumber
                quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex] = 
                    stakeRegistry.getTotalStakeAtBlockNumberFromIndex(quorumNumber, referenceBlockNumber, nonSignerStakesAndSignature.totalStakeIndices[quorumNumberIndex]);
                // copy total stake to signed stake
                quorumStakeTotals.signedStakeForQuorum[quorumNumberIndex] = quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex];
                // loop through all nonSigners, checking that they are a part of the quorum via their quorumBitmap
                // if so, load their stake at referenceBlockNumber and subtract it from running stake signed
                for (uint32 i = 0; i < nonSignerStakesAndSignature.nonSignerPubkeys.length; i++) {
                    // keep track of the nonSigners index in the quorum
                    uint32 nonSignerForQuorumIndex = 0;
                    // if the nonSigner is a part of the quorum, subtract their stake from the running total
                    if (nonSignerQuorumBitmaps[i] >> quorumNumber & 1 == 1) {
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
            // verify the signature
            (bool pairingSuccessful, bool sigantureIsValid) = trySignatureAndApkVerification(msgHash, apk, nonSignerStakesAndSignature.apkG2, nonSignerStakesAndSignature.sigma);
            require(pairingSuccessful, "BLSSignatureChecker.checkSignatures: pairing precompile call failed");
            require(sigantureIsValid, "BLSSignatureChecker.checkSignatures: signature is invalid");
        }
        // set signatoryRecordHash variable used for fraudproofs
        bytes32 signatoryRecordHash = MiddlewareUtils.computeSignatoryRecordHash(
            referenceBlockNumber,
            nonSignerPubkeyHashes
        );

        // return the total stakes that signed for each quorum, and a hash of the information required to prove the exact signers and stake
        return (quorumStakeTotals, signatoryRecordHash);
    }

    /**
     * trySignatureAndApkVerification verifies a BLS aggregate signature and the veracity of a calculated G1 Public key
     * @param msgHash is the hash being signed
     * @param apk is the claimed G1 public key
     * @param apkG2 is provided G2 public key
     * @param sigma is the G1 point signature
     * @return pairingSuccessful is true if the pairing precompile call was successful
     * @return siganatureIsValid is true if the signature is valid
     */
    function trySignatureAndApkVerification(
        bytes32 msgHash,
        BN254.G1Point memory apk,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory sigma
    ) internal view returns(bool pairingSuccessful, bool siganatureIsValid) {
        // gamma = keccak256(abi.encodePacked(msgHash, apk, apkG2, sigma))
        uint256 gamma = uint256(keccak256(abi.encodePacked(msgHash, apk.X, apk.Y, apkG2.X[0], apkG2.X[1], apkG2.Y[0], apkG2.Y[1], sigma.X, sigma.Y))) % BN254.FR_MODULUS;
        // verify the signature
        (pairingSuccessful, siganatureIsValid) = BN254.safePairing(
                sigma.plus(apk.scalar_mul(gamma)),
                BN254.negGeneratorG2(),
                BN254.hashToG1(msgHash).plus(BN254.generatorG1().scalar_mul(gamma)),
                apkG2,
                PAIRING_EQUALITY_CHECK_GAS
            );
    }
}