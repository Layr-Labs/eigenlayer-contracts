// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IBaseCertificateVerifier.sol";
import "./IOperatorTableCalculator.sol";

interface IECDSACertificateVerifierErrors {
    /// @notice Thrown when the signature length is invalid
    error InvalidSignatureLength();
    /// @notice Thrown when the signatures are not ordered by signer address
    error SignersNotOrdered();
    /// @notice Thrown when the operator count is zero
    error OperatorCountZero();
}

interface IECDSACertificateVerifierTypes is IOperatorTableCalculatorTypes {
    /**
     * @notice A Certificate used to verify a set of ECDSA signatures
     * @param referenceTimestamp a reference timestamp that corresponds to an operator table update
     * @param messageHash the hash of the message that was signed by the operators. The messageHash
     *        MUST be calculated using `calculateCertificateDigest`
     * @param sig the concatenated signature of each signing operator, in ascending order of signer address
     * @dev ECDSA certificates DO NOT support smart contract signatures
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
interface IECDSACertificateVerifier is
    IECDSACertificateVerifierEvents,
    IECDSACertificateVerifierErrors,
    IBaseCertificateVerifier
{
    /**
     * @notice updates the operator table
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos were sourced
     * @param operatorInfos the operatorInfos to update the operator table with
     * @param operatorSetConfig the configuration of the operatorSet
     * @dev Only callable by the `OperatorTableUpdater`
     * @dev The `referenceTimestamp` must correspond to a reference timestamp for a globalTableRoot stored in the `OperatorTableUpdater`
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
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @return totalSignedStakeWeights total stake weight that signed the certificate for each stake type. Each
     * index corresponds to a stake type in the `weights` array in the `ECDSAOperatorInfo`
     * @return signers array of addresses that signed the certificate
     * @dev This function DOES NOT support smart contact signatures
     */
    function verifyCertificate(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert
    ) external view returns (uint256[] memory totalSignedStakeWeights, address[] memory signers);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake weight on the AVS
     * @param operatorSet the operatorSet to verify the certificate for
     * @param cert a certificate
     * @param totalStakeProportionThresholds the proportion, in BPS, of total stake weight that
     * the signed stake of the certificate should meet. Each index corresponds to
     * a stake type in the `weights` array in the `ECDSAOperatorInfo`
     * @return Whether or not the certificate is valid and meets thresholds
     * @return signers array of addresses that signed the certificate
     */
    function verifyCertificateProportion(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert,
        uint16[] memory totalStakeProportionThresholds
    ) external view returns (bool, address[] memory signers);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided nominal stake thresholds
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @param totalStakeNominalThresholds the nominal amount of total stake weight that
     * the signed stake of the certificate should meet. Each index corresponds to
     * a stake type in the `weights` array in the `ECDSAOperatorInfo`
     * @return Whether or not the certificate is valid and meets thresholds
     * @return signers array of addresses that signed the certificate
     */
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external view returns (bool, address[] memory signers);

    /**
     * @notice Get operator infos for a timestamp
     * @param operatorSet The operator set
     * @param referenceTimestamp The reference timestamp
     * @return The operator infos, empty if the operatorSet has not been updated for the given reference timestamp
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
     * @return The operator info, empty if the operatorSet has not been updated for the given reference timestamp
     */
    function getOperatorInfo(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        uint256 operatorIndex
    ) external view returns (ECDSAOperatorInfo memory);

    /**
     * @notice Override domainSeparator to not include chainId
     * @return The domain separator hash without chainId
     * @dev This function overrides the base domainSeparator to not include chainId to replay
     *      certificates across multiple destination chains
     */
    function domainSeparator() external view returns (bytes32);

    /**
     * @notice Calculate the EIP-712 digest bytes for a certificate
     * @param referenceTimestamp The reference timestamp
     * @param messageHash The message hash
     * @return The EIP-712 digest
     * @dev This function is public to allow offchain tools to calculate the same digest
     * @dev Note: This does not support smart contract based signatures for multichain
     * @dev This is a chain-agnostic digest, so it can be used to verify certificates across
     *      multiple destination chains
     * @dev This function returns the raw bytes of the digest, which still need to be hashed
     *      before signing with ECDSA
     */
    function calculateCertificateDigestBytes(
        uint32 referenceTimestamp,
        bytes32 messageHash
    ) external view returns (bytes memory);

    /**
     * @notice Calculate the EIP-712 digest for a certificate
     * @param referenceTimestamp The reference timestamp
     * @param messageHash The message hash
     * @return The EIP-712 digest
     * @dev This function is public to allow offchain tools to calculate the same digest
     * @dev Note: This does not support smart contract based signatures for multichain
     * @dev This is a chain-agnostic digest, so it can be used to verify certificates across
     *      multiple destination chains
     */
    function calculateCertificateDigest(
        uint32 referenceTimestamp,
        bytes32 messageHash
    ) external view returns (bytes32);
}
