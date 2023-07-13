// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/interfaces/IBLSPublicKeyCompendium.sol";
import "../../contracts/libraries/BN254.sol";
import "forge-std/Test.sol";

/**
 * @title A shared contract for EigenLayer operators to register their BLS public keys.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
contract BLSPublicKeyCompendiumMock is IBLSPublicKeyCompendium, DSTest {

    /// @notice mapping from operator address to pubkey hash
    mapping(address => bytes32) public operatorToPubkeyHash;
    /// @notice mapping from pubkey hash to operator address
    mapping(bytes32 => address) public pubkeyHashToOperator;

    /**
     * @notice Called by an operator to register themselves as the owner of a BLS public key and reveal their G1 and G2 public key.
     * @param s is the field element of the operator's Schnorr signature
     * @param rPoint is the group element of the operator's Schnorr signature
     * @param pubkeyG1 is the the G1 pubkey of the operator
     * @param pubkeyG2 is the G2 with the same private key as the pubkeyG1
     */
    function registerBLSPublicKey(uint256 s, BN254.G1Point memory rPoint, BN254.G1Point memory pubkeyG1, BN254.G2Point memory pubkeyG2) external {
    }


    function registerPublicKey(BN254.G1Point memory pk) external {

        bytes32 pubkeyHash = BN254.hashG1Point(pk);
        // store updates
        operatorToPubkeyHash[msg.sender] = pubkeyHash;
        pubkeyHashToOperator[pubkeyHash] = msg.sender;
    }

    function setBLSPublicKey(address account, BN254.G1Point memory pk) external {

        bytes32 pubkeyHash = BN254.hashG1Point(pk);
        // store updates
        operatorToPubkeyHash[account] = pubkeyHash;
        pubkeyHashToOperator[pubkeyHash] = account;
    }
}