// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IBaseCertificateVerifier.sol";
import "src/contracts/interfaces/IECDSATableCalculator.sol";

interface IECDSACertificateVerifierTypes is IECDSATableCalculatorTypes {
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
}

interface IECDSACertificateVerifierEvents is IECDSACertificateVerifierTypes {
    /// @notice Emitted when an ECDSA table is updated
    event TableUpdated(OperatorSet operatorSet, uint32 referenceTimestamp, ECDSAOperatorInfo[] operatorInfos);
}

/// @notice A base interface that verifies certificates for a given operatorSet
/// @notice This is a base interface that all curve certificate verifiers (eg. BN254, ECDSA) must implement
/// @dev A single `CertificateVerifier` can be used for ONLY 1 operatorSet
interface IECDSACertificateVerifier is IECDSACertificateVerifierEvents, IBaseCertificateVerifier {
    /* ECDSA CERTIFICATE VERIFIER INTERFACE */

    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos were sourced
     * @param operatorInfos the operatorInfos to update the operator table with
     * @param OperatorSetConfig the configuration of the operatorSet
     * @dev only callable by the operatorTableUpdater for the given operatorSet
     * @dev We pass in an `operatorSet` for future-proofing a global `TableManager` contract
     * @dev The `referenceTimestamp` must be greater than the latest reference timestamp for the given operatorSet
     */
    function updateOperatorTable(
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
    function verifyCertificate(
        OperatorSet calldata operatorSet,
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
    function verifyCertificateProportion(
        OperatorSet calldata operatorSet,
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
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert,
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
