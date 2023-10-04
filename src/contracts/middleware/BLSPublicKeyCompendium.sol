// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IBLSPublicKeyCompendium.sol";
import "../libraries/BN254.sol";

/**
 * @title A shared contract for EigenLayer operators to register their BLS public keys.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
contract BLSPublicKeyCompendium is IBLSPublicKeyCompendium {
    using BN254 for BN254.G1Point;

    /// @notice mapping from operator address to pubkey hash
    mapping(address => bytes32) public operatorToPubkeyHash;
    /// @notice mapping from pubkey hash to operator address
    mapping(bytes32 => address) public pubkeyHashToOperator;

    /**
     * @notice Called by an operator to register themselves as the owner of a BLS public key and reveal their G1 and G2 public key.
     * @param signedMessageHash is the registration message hash signed by the private key of the operator
     * @param pubkeyG1 is the corresponding G1 public key of the operator 
     * @param pubkeyG2 is the corresponding G2 public key of the operator
     */
    function registerBLSPublicKey(BN254.G1Point memory signedMessageHash, BN254.G1Point memory pubkeyG1, BN254.G2Point memory pubkeyG2) external {
        // H(m) 
        BN254.G1Point memory messageHash = BN254.hashToG1(keccak256(abi.encodePacked(
            msg.sender, 
            block.chainid, 
            "EigenLayer_BN254_Pubkey_Registration"
        )));

        // gamma = h(sigma, P, P', H(m))
        uint256 gamma = uint256(keccak256(abi.encodePacked(
            signedMessageHash.X, 
            signedMessageHash.Y, 
            pubkeyG1.X, 
            pubkeyG1.Y, 
            pubkeyG2.X, 
            pubkeyG2.Y, 
            messageHash.X, 
            messageHash.Y
        ))) % BN254.FR_MODULUS;
        
        // e(sigma + P * gamma, [-1]_2) = e(H(m) + [1]_1 * gamma, P') 
        require(BN254.pairing(
            signedMessageHash.plus(pubkeyG1.scalar_mul(gamma)),
            BN254.negGeneratorG2(),
            messageHash.plus(BN254.generatorG1().scalar_mul(gamma)),
            pubkeyG2
        ), "BLSPublicKeyCompendium.registerBLSPublicKey: G1 and G2 private key do not match");

        bytes32 pubkeyHash = BN254.hashG1Point(pubkeyG1);

        require(
            operatorToPubkeyHash[msg.sender] == bytes32(0),
            "BLSPublicKeyCompendium.registerBLSPublicKey: operator already registered pubkey"
        );
        require(
            pubkeyHashToOperator[pubkeyHash] == address(0),
            "BLSPublicKeyCompendium.registerBLSPublicKey: public key already registered"
        );

        operatorToPubkeyHash[msg.sender] = pubkeyHash;
        pubkeyHashToOperator[pubkeyHash] = msg.sender;

        emit NewPubkeyRegistration(msg.sender, pubkeyG1, pubkeyG2);
    }
}
