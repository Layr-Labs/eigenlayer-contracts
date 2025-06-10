// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {BN254} from "./BN254.sol";

/**
 * @title BN254SignatureVerifier
 * @notice Library for BN254 signature verification
 * @dev Provides unified signature verification with consistent gamma calculation and hash-to-G1 conversion
 */
library BN254SignatureVerifier {
    using BN254 for BN254.G1Point;

    /**
     * @notice Core BN254 signature verification function with optional gas limiting
     * @param msgHash The message hash that was signed
     * @param signature The BLS signature to verify (G1 point)
     * @param pubkeyG1 The G1 component of the public key
     * @param pubkeyG2 The G2 component of the public key
     * @param useGasLimit Whether to use gas-limited safe pairing
     * @param pairingGas Gas limit for pairing (ignored if useGasLimit is false)
     * @return success True if verification succeeded (always true if useGasLimit=false due to revert)
     * @return pairingSuccessful True if pairing operation completed (only relevant when useGasLimit=true)
     */
    function verifySignature(
        bytes32 msgHash,
        BN254.G1Point memory signature,
        BN254.G1Point memory pubkeyG1,
        BN254.G2Point memory pubkeyG2,
        bool useGasLimit,
        uint256 pairingGas
    ) internal view returns (bool success, bool pairingSuccessful) {
        BN254.G1Point memory messagePoint = BN254.hashToG1(msgHash);
        uint256 gamma = _calculateGamma(msgHash, pubkeyG1, pubkeyG2, signature);

        // Calculate pairing inputs
        BN254.G1Point memory leftG1 = signature.plus(pubkeyG1.scalar_mul(gamma));
        BN254.G1Point memory rightG1 = messagePoint.plus(BN254.generatorG1().scalar_mul(gamma));

        if (useGasLimit) {
            // Use safe pairing with gas limit
            (pairingSuccessful, success) =
                BN254.safePairing(leftG1, BN254.negGeneratorG2(), rightG1, pubkeyG2, pairingGas);
        } else {
            success = BN254.pairing(leftG1, BN254.negGeneratorG2(), rightG1, pubkeyG2);
            if (success) {
                pairingSuccessful = true;
            }
        }
    }

    /**
     * @notice Internal function to calculate gamma value for signature verification
     * @param msgHash The message hash
     * @param pubkeyG1 The G1 component of the public key
     * @param pubkeyG2 The G2 component of the public key
     * @param signature The signature point
     * @return gamma The calculated gamma value
     */
    function _calculateGamma(
        bytes32 msgHash,
        BN254.G1Point memory pubkeyG1,
        BN254.G2Point memory pubkeyG2,
        BN254.G1Point memory signature
    ) internal pure returns (uint256 gamma) {
        gamma = uint256(
            keccak256(
                abi.encodePacked(
                    msgHash,
                    pubkeyG1.X,
                    pubkeyG1.Y,
                    pubkeyG2.X[0],
                    pubkeyG2.X[1],
                    pubkeyG2.Y[0],
                    pubkeyG2.Y[1],
                    signature.X,
                    signature.Y
                )
            )
        ) % BN254.FR_MODULUS;
    }
}
