// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../libraries/MiddlewareUtils.sol";
import "../libraries/BN254.sol";
import "../libraries/BitmapUtils.sol";

/**
 * @title Used for checking BLS aggregate signatures from the operators of a `BLSRegistry`.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This is the contract for checking the validity of aggregate operator signatures.
 */
contract BLSSignatureChecker {
    using BN254 for BN254.G1Point;

    // DATA STRUCTURES

    /*
    NonSignerStakesAndSignature is passed as input to checkSignatures() to verify an aggregate signature.

    The Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
    updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
    Each of these events push a new entry to its respective datatype array.

    The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
    Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
    One can use getCheckSignaturesIndices() in BLSOperatorStateRetriever.sol to fetch the indices given a block number.

    The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
    apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
    quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
    nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
    Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
    block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.

    fields:
    @nonSignersPubkeysG1: G1 pubkey of all non-signers
    @quorumApksG1: G1 aggregated pubkey of each quorum (aggregated over all operators opted in to the quorum at the given blocknumber)
    @signersApkG2: aggregated G2 pubkey of the signers
    @signersAggSigG1: aggregated signature of the signers
    @nonSignersQuorumBitmapIndices: the indices of the quorumBitmaps for each of the operators in the @param nonSignerOperatorIds array at the given blocknumber
    @quorumApkIndices: the indices of the quorum apks for each of the provided quorums at the given blocknumber
    @quorumTotalStakeIndices: the indices of the total stakes entries for the given quorums at the given blocknumber
    @quorumNonSignersStakeIndices: the indices of the stakes of each of the nonsigners in each of the quorums they were a
                                   part of (for each nonsigner, an array of length the number of quorums they were a part of
                                   that are also part of the provided quorumNumbers) at the given blocknumber
    
    Each field array can either be over signers, nonsigners, or quorums.
    We use the first word of the variable to represent what the array is over
      eg: []nonSignersPubkeysG1 -> array of nonSigners
      eg: []quorumApkG1 -> array of quorums
    */
    struct NonSignerStakesAndSignature {
        BN254.G1Point[] nonSignersPubkeysG1; // nonSignersPubkeysG1[nonSignerIndex]
        BN254.G1Point[] quorumApksG1; // quorumApksG1[quorumIndex]
        BN254.G2Point signersApkG2;
        BN254.G1Point signersAggSigG1;
        uint32[] nonSignersQuorumBitmapIndices; // nonSignersQuorumBitmapIndices[nonSignerIndex]
        uint32[] quorumApkIndices; // quorumApkIndices[quorumIndex]
        uint32[] quorumTotalStakeIndices; // quorumTotalStakeIndices[quorumIndex]
        uint32[][] quorumNonSignersStakeIndices; // quorumNonSignersStakeIndices[quorumIndex][nonSignerIndex]
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
    ) public view returns (QuorumStakeTotals memory, bytes32) {
        // verify the provided apk was the apk at referenceBlockNumber
        // loop through every quorumNumber and keep track of the apk
        BN254.G1Point memory apk = BN254.G1Point(0, 0);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            require(
                bytes24(
                    nonSignerStakesAndSignature.quorumApksG1[i].hashG1Point()
                ) ==
                    blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(
                        uint8(quorumNumbers[i]),
                        referenceBlockNumber,
                        nonSignerStakesAndSignature.quorumApkIndices[i]
                    ),
                "BLSSignatureChecker.checkSignatures: quorumApk hash in storage does not match provided quorum apk"
            );
            apk = apk.plus(nonSignerStakesAndSignature.quorumApksG1[i]);
        }

        // initialize memory for the quorumStakeTotals
        QuorumStakeTotals memory quorumStakeTotals;
        quorumStakeTotals.totalStakeForQuorum = new uint96[](
            quorumNumbers.length
        );
        quorumStakeTotals.signedStakeForQuorum = new uint96[](
            quorumNumbers.length
        );
        // the pubkeyHashes of the nonSigners
        bytes32[] memory nonSignerPubkeyHashes = new bytes32[](
            nonSignerStakesAndSignature.nonSignersPubkeysG1.length
        );
        {
            // the quorumBitmaps of the nonSigners
            uint256[] memory nonSignerQuorumBitmaps = new uint256[](
                nonSignerStakesAndSignature.nonSignersPubkeysG1.length
            );
            {
                // the bitmap of the quorumNumbers
                uint256 signingQuorumBitmap = BitmapUtils.bytesArrayToBitmap(
                    quorumNumbers
                );

                for (
                    uint i = 0;
                    i < nonSignerStakesAndSignature.nonSignersPubkeysG1.length;
                    i++
                ) {
                    nonSignerPubkeyHashes[i] = nonSignerStakesAndSignature
                        .nonSignersPubkeysG1[i]
                        .hashG1Point();

                    // check that the nonSignerPubkeys are sorted and free of duplicates
                    if (i != 0) {
                        require(
                            uint256(nonSignerPubkeyHashes[i]) >
                                uint256(nonSignerPubkeyHashes[i - 1]),
                            "BLSSignatureChecker.checkSignatures: nonSignerPubkeys not sorted"
                        );
                    }

                    nonSignerQuorumBitmaps[i] = registryCoordinator
                        .getQuorumBitmapByOperatorIdAtBlockNumberByIndex(
                            nonSignerPubkeyHashes[i],
                            referenceBlockNumber,
                            nonSignerStakesAndSignature
                                .nonSignersQuorumBitmapIndices[i]
                        );

                    // subtract the nonSignerPubkey from the running apk to get the apk of all signers
                    apk = apk.plus(
                        nonSignerStakesAndSignature
                            .nonSignersPubkeysG1[i]
                            .negate()
                            .scalar_mul_tiny(
                                BitmapUtils.countNumOnes(
                                    nonSignerQuorumBitmaps[i] &
                                        signingQuorumBitmap
                                ) // we subtract the nonSignerPubkey from each quorum that they are a part of, TODO:
                            )
                    );
                }
            }
            // loop through each quorum number
            for (
                uint8 quorumNumberIndex = 0;
                quorumNumberIndex < quorumNumbers.length;

            ) {
                // get the quorum number
                uint8 quorumNumber = uint8(quorumNumbers[quorumNumberIndex]);
                // get the totalStake for the quorum at the referenceBlockNumber
                quorumStakeTotals.totalStakeForQuorum[
                    quorumNumberIndex
                ] = stakeRegistry.getTotalStakeAtBlockNumberFromIndex(
                    quorumNumber,
                    referenceBlockNumber,
                    nonSignerStakesAndSignature.quorumTotalStakeIndices[
                        quorumNumberIndex
                    ]
                );
                // copy total stake to signed stake
                quorumStakeTotals.signedStakeForQuorum[
                    quorumNumberIndex
                ] = quorumStakeTotals.totalStakeForQuorum[quorumNumberIndex];
                // loop through all nonSigners, checking that they are a part of the quorum via their quorumBitmap
                // if so, load their stake at referenceBlockNumber and subtract it from running stake signed
                for (
                    uint32 i = 0;
                    i < nonSignerStakesAndSignature.nonSignersPubkeysG1.length;
                    i++
                ) {
                    // keep track of the nonSigners index in the quorum
                    uint32 nonSignerForQuorumIndex = 0;
                    // if the nonSigner is a part of the quorum, subtract their stake from the running total
                    if ((nonSignerQuorumBitmaps[i] >> quorumNumber) & 1 == 1) {
                        quorumStakeTotals.signedStakeForQuorum[
                            quorumNumberIndex
                        ] -= stakeRegistry
                            .getStakeForQuorumAtBlockNumberFromOperatorIdAndIndex(
                                quorumNumber,
                                referenceBlockNumber,
                                nonSignerPubkeyHashes[i],
                                nonSignerStakesAndSignature
                                    .quorumNonSignersStakeIndices[quorumNumberIndex][
                                        nonSignerForQuorumIndex
                                    ]
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
            (
                bool pairingSuccessful,
                bool sigantureIsValid
            ) = trySignatureAndApkVerification(
                    msgHash,
                    apk,
                    nonSignerStakesAndSignature.signersApkG2,
                    nonSignerStakesAndSignature.signersAggSigG1
                );
            require(
                pairingSuccessful,
                "BLSSignatureChecker.checkSignatures: pairing precompile call failed"
            );
            require(
                sigantureIsValid,
                "BLSSignatureChecker.checkSignatures: signature is invalid"
            );
        }
        // set signatoryRecordHash variable used for fraudproofs
        bytes32 signatoryRecordHash = MiddlewareUtils
            .computeSignatoryRecordHash(
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
    ) public view returns (bool pairingSuccessful, bool siganatureIsValid) {
        // gamma = keccak256(abi.encodePacked(msgHash, apk, apkG2, sigma))
        uint256 gamma = uint256(
            keccak256(
                abi.encodePacked(
                    msgHash,
                    apk.X,
                    apk.Y,
                    apkG2.X[0],
                    apkG2.X[1],
                    apkG2.Y[0],
                    apkG2.Y[1],
                    sigma.X,
                    sigma.Y
                )
            )
        ) % BN254.FR_MODULUS;
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
