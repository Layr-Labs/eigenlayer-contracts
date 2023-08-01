// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../libraries/BN254.sol";

/**
 * @title Minimal interface for the `BLSPublicKeyCompendium` contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IBLSPublicKeyCompendium {

    // DATA STRUCTURES
    struct BN254Pubkeys {
        // G1 public key over BN254 curve for the operator
        BN254.G1Point pubkeyG1;
        // G2 public key over BN254 curve for the operator
        BN254.G2Point pubkeyG2;
    }

    /**
     * @notice mapping from operator address to pubkey hash.
     * Returns *zero* if the `operator` has never registered, and otherwise returns the hash of the public key of the operator.
     */
    function operatorToPubkeyHash(address operator) external view returns (bytes32);

    /**
     * @notice mapping from pubkey hash to operator address.
     * Returns *zero* if no operator has ever registered the public key corresponding to `pubkeyHash`,
     * and otherwise returns the (unique) registered operator who owns the BLS public key that is the preimage of `pubkeyHash`.
     */
    function pubkeyHashToOperator(bytes32 pubkeyHash) external view returns (address);

    /**
     * @notice mapping from pubkey hash to BN254 pubkeys.
     **/
    function getBN254PubkeysFromOperatorID(bytes32 operatorID) external view returns (BN254Pubkeys memory);

    /**
     * @notice Called by an operator to register themselves as the owner of a BLS public key and reveal their G1 and G2 public key.
     * @param signedMessageHash is the registration message hash signed by the private key of the operator
     * @param pubkeyG1 is the corresponding G1 public key of the operator 
     * @param pubkeyG2 is the corresponding G2 public key of the operator
     */
    function registerBLSPublicKey(BN254.G1Point memory signedMessageHash, BN254.G1Point memory pubkeyG1, BN254.G2Point memory pubkeyG2) external;
}
