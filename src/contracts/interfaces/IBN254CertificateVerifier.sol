// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {BN254} from "../libraries/BN254.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";
import "./IBaseCertificateVerifier.sol";

interface IBN254CertificateVerifierTypes is IOperatorTableCalculatorTypes {
    /**
     * @notice A witness for an operator
     * @param operatorIndex the index of the nonsigner in the `BN254OperatorInfo` tree
     * @param operatorInfoProofs merkle proofs of the nonsigner at the index. Empty if operator is in cache.
     * @param operatorInfo the `BN254OperatorInfo` for the operator. Empty if operator is in cache
     */
    struct BN254OperatorInfoWitness {
        uint32 operatorIndex;
        bytes operatorInfoProof;
        BN254OperatorInfo operatorInfo;
    }

    /**
     * @notice A BN254 Certificate
     * @param referenceTimestamp the timestamp at which the certificate was created
     * @param messageHash the hash of the message that was signed by operators and used to verify the aggregated signature
     * @param signature the G1 signature of the message
     * @param apk the G2 aggregate public key
     * @param nonSignerWitnesses an array of witnesses of non-signing operators
     */
    struct BN254Certificate {
        uint32 referenceTimestamp;
        bytes32 messageHash;
        BN254.G1Point signature;
        BN254.G2Point apk;
        BN254OperatorInfoWitness[] nonSignerWitnesses;
    }
}

interface IBN254CertificateVerifierEvents is IBN254CertificateVerifierTypes {
    /// @notice Emitted when an BN254 table is updated
    event TableUpdated(OperatorSet operatorSet, uint32 referenceTimestamp, BN254OperatorSetInfo operatorSetInfo);
}

interface IBN254CertificateVerifierErrors {
    ///@notice thrown when operator index provided in certificate is invalid
    error InvalidOperatorIndex();
}

/// @notice An interface for verifying BN254 certificates
/// @notice This implements the base `IBaseCertificateVerifier` interface
interface IBN254CertificateVerifier is
    IBN254CertificateVerifierEvents,
    IBaseCertificateVerifier,
    IBN254CertificateVerifierErrors
{
    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos were sourced
     * @param operatorSetInfo the operatorInfos to update the operator table with
     * @param operatorSetConfig the configuration of the operatorSet
     * @dev only callable by the operatorTableUpdater for the given operatorSet
     * @dev The `referenceTimestamp` must be greater than the latest reference timestamp for the given operatorSet
     */
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig calldata operatorSetConfig
    ) external;

    /**
     * @notice verifies a certificate
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @return signedStakes amount of stake that signed the certificate for each stake
     * type. Each index corresponds to a stake type in the `totalWeights` array in the `BN254OperatorSetInfo`.
     */
    function verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) external returns (uint256[] memory signedStakes);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake on the AVS
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @param totalStakeProportionThresholds the proportion, in BPS,of total stake that
     * the signed stake of the certificate should meet. Each index corresponds to
     * a stake type in the `totalWeights` array in the `BN254OperatorSetInfo`.
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyCertificateProportion(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided nominal stake thresholds
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @param totalStakeNominalThresholds the nominal amount of stake that
     * the signed stake of the certificate should meet. Each index corresponds to
     * a stake type in the `totalWeights` array in the `BN254OperatorSetInfo`.
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyCertificateNominal(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool);

    /**
     * @notice Attempts signature verification with gas limit for safety
     * @param msgHash The message hash that was signed
     * @param aggPubkey The aggregate public key of signers
     * @param apkG2 The G2 point representation of the aggregate public key
     * @param signature The BLS signature to verify
     * @return pairingSuccessful Whether the pairing operation completed successfully
     * @return signatureValid Whether the signature is valid
     */
    function trySignatureVerification(
        bytes32 msgHash,
        BN254.G1Point memory aggPubkey,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory signature
    ) external view returns (bool pairingSuccessful, bool signatureValid);

    /**
     * @notice Get cached nonsigner operator info
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The operator index
     * @return The cached operator info
     * @dev If the operator is not in the cache, the operator info will be empty
     */
    function getNonsignerOperatorInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (BN254OperatorInfo memory);

    /**
     * @notice Check if a nonsigner is cached
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The operator index
     * @return Whether the operator is cached
     */
    function isNonsignerCached(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (bool);

    /**
     * @notice Get operator set info for a timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The operator set info
     */
    function getOperatorSetInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (BN254OperatorSetInfo memory);
}
