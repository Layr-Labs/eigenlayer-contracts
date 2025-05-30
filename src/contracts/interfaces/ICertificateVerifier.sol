// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IECDSATableCalculator.sol";
import "src/contracts/interfaces/IBN254TableCalculator.sol";

interface ICertificateVerifierTypes is
    IBN254TableCalculatorTypes,
    IECDSATableCalculatorTypes,
    ICrossChainRegistryTypes
{
    // TODO: use the `KeyType` from `KeyRegistrar`
    /**
     * @notice The type of key used by the operatorSet. An OperatorSet can
     * only generate one Operator Table for an OperatorSet for a given OperatorKeyType.
     */
    enum OperatorKeyType {
        ECDSA,
        BN254
    }

    /**
     * @notice A ECDSA Certificate
     * @param referenceTimestamp the timestamp at which the certificate was created
     * @param messageHash the hash of the message that was signed by operators
     * @param signature the concatenated signature of each signing operator
     */
    struct ECDSACertificate {
        uint32 referenceTimestamp;
        bytes32 messageHash;
        bytes sig;
    }

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

interface ICertificateVerifierEvents is ICertificateVerifierTypes {
    /// @notice Emitted when an ECDSA table is updated
    event ECDSATableUpdated(uint32 referenceTimestamp, ECDSAOperatorInfo[] operatorInfos);

    /// @notice Emitted when a BN254 table is updated
    event BN254TableUpdated(uint32 referenceTimestamp, BN254OperatorSetInfo operatorSetInfo);
}

interface ICertificateVerifierErrors {
    /// @notice Thrown when the table updater is not caller
    error OnlyTableUpdater();
    /// @notice Thrown when the table is too stale
    error TableStale();
    /// @notice Thrown when certificate verification fails
    error VerificationFailed();
}

/// @notice A base interface that verifies certificates for a given operatorSet
/// @notice This is a base interface that all curve certificate verifiers (eg. BN254, ECDSA) must implement
/// @dev A single `CertificateVerifier` can be used for ONLY 1 operatorSet
interface ICertificateVerifier is ICertificateVerifierEvents, ICertificateVerifierErrors {
    /* ECDSA CERTIFICATE VERIFIER INTERFACE */

    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos were sourced
     * @param operatorInfos the operatorInfos to update the operator table with
     * @param OperatorSetConfig the configuration of the operatorSet
     * @dev only callable by the operatorTableUpdater for the given operatorSet
     * @dev We pass in an `operatorSet` for future-proofing a global `TableManager` contract
     */
    function updateECDSAOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata OperatorSetConfig
    ) external;

    /**
     * @notice verifies a certificate
     * @param cert a certificate
     * @return signedStakes amount of stake that signed the certificate for each stake
     * type
     */
    function verifyECDSACertificate(
        ECDSACertificate memory cert
    ) external returns (uint96[] memory signedStakes);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake on the AVS
     * @param cert a certificate
     * @param totalStakeProportionThresholds the proportion of total stake that
     * the signed stake of the certificate should meet
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyECDSACertificateProportion(
        ECDSACertificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake on the AVS
     * @param cert a certificate
     * @param totalStakeNominalThresholds the proportion of total stake that
     * the signed stake of the certificate should meet
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyECDSACertificateNominal(
        ECDSACertificate memory cert,
        uint96[] memory totalStakeNominalThresholds
    ) external returns (bool);

    /* BN254 CERTIFICATE VERIFIER INTERFACE */

    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorSetInfo and
     * operatorInfoTreeRoot were sourced
     * @param operatorSetInfo the aggregate information about the operatorSet
     * @param OperatorSetConfig the configuration of the operatorSet
     * @dev only callable by the operatorTableUpdater for the given operatorSet
     * @dev We pass in an `operatorSet` for future-proofing a global `TableManager` contract
     */
    function updateBN254OperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig calldata OperatorSetConfig
    ) external;

    /**
     * @notice verifies a certificate
     * @param cert a certificate
     * @return signedStakes amount of stake that signed the certificate for each stake
     * type
     */
    function verifyBN254Certificate(
        BN254Certificate memory cert
    ) external returns (uint96[] memory signedStakes);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake on the AVS
     * @param cert a certificate
     * @param totalStakeProportionThresholds the proportion of total stake that
     * the signed stake of the certificate should meet
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyBN254CertificateProportion(
        BN254Certificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external returns (bool);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided nominal stake thresholds
     * @param cert a certificate
     * @param totalStakeNominalThresholds the nominal amount of stake that
     * the signed stake of the certificate should meet
     * @return whether or not certificate is valid and meets thresholds
     */
    function verifyBN254CertificateNominal(
        BN254Certificate memory cert,
        uint96[] memory totalStakeNominalThresholds
    ) external returns (bool);

    /* OPERATOR SET CONFIG INTERFACE */

    /// @notice the address of the owner of the OperatorSet
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external returns (address);

    /// @return the maximum amount of seconds that a operator table can be in the past for a given operatorSet
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external returns (uint32);

    /// @notice The latest reference timestamp of the operator table for a given operatorSet
    function latestReferenceTimestamp(
        OperatorSet memory operatorSet
    ) external returns (uint32);
}
