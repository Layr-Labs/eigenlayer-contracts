// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IBaseCertificateVerifier.sol";
import "./IECDSATableCalculator.sol";

interface IECDSACertificateVerifierTypes is IECDSATableCalculatorTypes {
    // Errors
    error InvalidSignatureLength();

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

/// @notice An interface for verifying ECDSA certificates
/// @notice This implements the base `IBaseCertificateVerifier` interface
interface IECDSACertificateVerifier is IECDSACertificateVerifierEvents, IBaseCertificateVerifier {
    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos were sourced
     * @param operatorInfos the operatorInfos to update the operator table with
     * @param operatorSetConfig the configuration of the operatorSet
     * @dev only callable by the operatorTableUpdater for the given operatorSet
     * @dev The `referenceTimestamp` must be greater than the latest reference timestamp for the given operatorSet
     */
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata operatorSetConfig
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
    ) external returns (uint256[] memory signedStakes);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake on the AVS
     * @param cert a certificate
     * @param totalStakeProportionThresholds the proportion of total stake that
     * the signed stake of the certificate should meet
     * @return Whether or not the certificate is valid and meets thresholds
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
     * @return Whether or not the certificate is valid and meets thresholds
     */
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool);

    /**
     * @notice Get operator infos for a timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The operator infos
     */
    function getOperatorInfos(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) external view returns (ECDSAOperatorInfo[] memory);

    /**
     * @notice Get a single operator info by index
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The index of the operator
     * @return The operator info
     */
    function getOperatorInfo(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        uint32 operatorIndex
    ) external view returns (ECDSAOperatorInfo memory);

    /**
     * @notice Get the total number of operators for a given reference timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The number of operators
     */
    function getOperatorCount(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint32);

    /**
     * @notice Get the total stakes for all operators at a given reference timestamp
     * @param operatorSet The operator set to calculate stakes for
     * @param referenceTimestamp The reference timestamp
     * @return totalStakes The total stakes for all operators
     */
    function getTotalStakes(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint256[] memory);
}
