// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {BN254} from "../libraries/BN254.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IBN254TableCalculator.sol";
import "./IBaseCertificateVerifier.sol";

interface IBN254CertificateVerifierTypes is IBN254TableCalculatorTypes {
    /**
     * @notice A witness for an operator
     * @param operatorIndex the index of the nonsigner in the `BN254OperatorInfo` tree
     * @param operatorInfoProofs merkle proofs of the nonsigner at the index. Empty if operator is in cache.
     * @param operatorInfo the `BN254OperatorInfo` for the operator
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
     * type
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
     * @param totalStakeProportionThresholds the proportion of total stake that
     * the signed stake of the certificate should meet
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
     * the signed stake of the certificate should meet
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyCertificateNominal(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool);
}
