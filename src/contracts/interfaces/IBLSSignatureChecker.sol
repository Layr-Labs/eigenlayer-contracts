// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../libraries/BN254.sol";
import "../libraries/BitmapUtils.sol";

/**
 * @title Used for checking BLS aggregate signatures from the operators of a EigenLayer AVS with the RegistryCoordinator/BLSPubkeyRegistry/StakeRegistry architechture.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This is the contract for checking the validity of aggregate operator signatures.
 */
interface IBLSSignatureChecker {
    // DATA STRUCTURES

    struct NonSignerStakesAndSignature {
        uint32[] nonSignerQuorumBitmapIndices; // is the indices of all nonsigner quorum bitmaps
        BN254.G1Point[] nonSignerPubkeys; // is the G1 pubkeys of all nonsigners
        BN254.G1Point[] quorumApks; // is the aggregate G1 pubkey of each quorum
        BN254.G2Point apkG2; // is the aggregate G2 pubkey of all signers and non signers
        BN254.G1Point sigma; // is the aggregate G1 signature of all signers
        uint32[] quorumApkIndices; // is the indices of each quorum aggregate pubkey
        uint32[] totalStakeIndices; // is the indices of each quorums total stake
        uint32[][] nonSignerStakeIndices; // is the indices of each non signers stake within a quorum
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

    function registryCoordinator() external view returns (IRegistryCoordinator);
    function stakeRegistry() external view returns (IStakeRegistry);
    function blsPubkeyRegistry() external view returns (IBLSPubkeyRegistry);

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
        external 
        view
        returns (
            QuorumStakeTotals memory,
            bytes32
        );
}