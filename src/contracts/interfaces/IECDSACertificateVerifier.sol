// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IBaseCertificateVerifier.sol";
import "./IOperatorTableCalculator.sol";

interface IECDSACertificateVerifierErrors {
    /// @notice Thrown when the signature length is invalid
    /// @dev Error code: 0x4be6321b
    /// @dev We require valid signature lengths (65 bytes per signature) for proper ECDSA signature verification and recovery
    error InvalidSignatureLength();

    /// @notice Thrown when the signatures are not ordered by signer address to validate unique signers
    /// @dev Error code: 0xb550c570
    /// @dev We order signers by address as a gas optimization for verification and to ensure unique signers without additional storage
    error SignersNotOrdered();

    /// @notice Thrown when the operator count is zero
    /// @dev Error code: 0x40a42054
    /// @dev We require a non-zero operator count to ensure there are operators available for certificate verification
    error OperatorCountZero();

    /// @notice Thrown when the operator index is out of bounds
    /// @dev Error code: 0x40a42054
    /// @dev We require a valid operator index to ensure the operator exists in the operator table
    error IndexOutOfBounds();
}

interface IECDSACertificateVerifierTypes is IOperatorTableCalculatorTypes {
    /**
     * @notice A Certificate used to verify a set of ECDSA signatures for an off-chain task
     * @param referenceTimestamp a reference timestamp that corresponds to a timestamp at which an operator table was updated for the operatorSet
     * @param messageHash the hash of a task that was completed by operators. The messageHash is defined by the AVS, see `TaskMailbox.sol` for an example implementation.
     *                    NOTE: This value is NOT the message that is signed by operators - see `calculateCertificateDigest` for the signable digest.
     * @param sig the concatenated signature of each signing operator, in ascending order of signer address. The signature should be over the signable digest,
     *            which is calculated by `calculateCertificateDigest`
     * @dev The signers can be sorted via OZ sort library
     * @dev ECDSA certificates DO NOT support smart contract signatures
     * @dev The `referenceTimestamp` is used to key into the operatorSet's stake weights. It is NOT the timestamp at which the certificate was generated off-chain
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
    /// @notice The following steps describe the certificate verification process, in order:
    /// 1. The AVS configures the following parameters in EigenLayer core:
    ///    a. AllocationManager.createOperatorSet: Creates an operatorSet
    ///    b. KeyRegistrar.configureOperatorSet: Configures the curve type of the operatorSet
    ///    c. CrossChainRegistry.makeGenerationReservation: Registers the operatorSet to be transported by the multichain protocol. This includes
    ///      the `owner`, `maxStalenessPeriod`, and `operatorTableCalculator` for the operatorSet. The output of the `OperatorTableCalculator`
    ///      are the operatorSet's stake weights (i.e. operator table) and is transported by the multichain protocol, along with the `maxStalenessPeriod` and `owner`
    /// 2. The multichain protocol calculates the operatorTable of an operatorSet. The time at which the table is calculated is the reference timestamp. The protocol
    ///    will then call `updateOperatorTable` to update the operatorSet's operator table for a given referenceTimestamp
    /// 3. A task is created and certificate is generated, off-chain, by the AVS to validate the completion of a task.
    ///    The reference timestamp in the certificate is used to key into the operator table that was updated in step 2.
    /// 4. The certificate is verified, either normally, proportionally, or nominally.
    /// @dev The `referenceTimestamp` is used to key into the operatorSet's stake weights. It is NOT when the certificate was generated off-chain
    /// @dev The `maxStalenessPeriod` configured in step 1c denotes if a certificate is too stale with respect to the `referenceTimestamp`
    /// @dev Operator tables for ALL operatorSets with an active generation reservation are updated at a set cadence. See `crossChainRegistry.tableUpdateCadence` for the frequency of table updates
    /// @dev To ensure that tables do not become stale between table updates (i.e. a large operator has joined or been ejected), the multichain protocol updates tables for operatorSets when the following events are emitted:
    ///      - AllocationManager: `OperatorSlashed`
    ///      - AllocationManager: `OperatorAddedToOperatorSet`
    ///      - AllocationManager: `OperatorRemovedFromOperatorSet`
    /// @dev Certificates can be replayed across all destination chains
    /// @dev Race conditions should be handled by the AVS. The protocol makes no guarantees about how certificates should be verified (eg. preventing certificates against tables that are NOT the latest)
    ///      Some examples of race conditions include:
    ///      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
    ///         wants to only confirm tasks against the *latest* certificate
    ///      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet

    /**
     * @notice updates the operatorSet with the operator table (i.e. stake weights) and its configuration
     * @param operatorSet the operatorSet to update the operator table for
     * @param referenceTimestamp the timestamp at which the operatorInfos (i.e. operator table) was sourced
     * @param operatorInfos the operatorInfos to update the operator table with.
     *        See `IOperatorTableCalculator.ECDSAOperatorInfo` for more details
     * @param operatorSetConfig the configuration of the operatorSet, which includes the owner and max staleness period
     * @dev This function can only be called by the `OperatorTableUpdater` contract, which is itself permissionless to call
     * @dev The `referenceTimestamp` must correspond to a reference timestamp for a globalTableRoot stored in the `OperatorTableUpdater`
     *      In addition, it must be greater than the latest reference timestamp for the given operatorSet
     * @dev Reverts for:
     *      - OnlyTableUpdater: Caller is not the operatorTableUpdater
     *      - TableUpdateStale: The referenceTimestamp is not greater than the latest reference timestamp
     * @dev Emits the following events:
     *      - TableUpdated: When the operator table is successfully updated
     */
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata operatorSetConfig
    ) external;

    /**
     * @notice verifies a certificate against the operator table for a given reference timestamp
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @return totalSignedStakeWeights total stake weight that signed the certificate for each stake type. Each
     * index corresponds to a stake type in the `weights` array in the `ECDSAOperatorInfo` struct
     * @return signers array of addresses that signed the certificate
     * @dev This function DOES NOT support smart contact signatures
     * @dev The `referenceTimestamp` in the `ECDSACertificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts for:
     *      - CertificateStale: The certificate's referenceTimestamp is too stale with respect to the maxStalenessPeriod of the operatorSet
     *      - ReferenceTimestampDoesNotExist: The root at referenceTimestamp does not exist
     *      - RootDisabled: The root at referenceTimestamp is not valid
     *      - InvalidSignatureLength: Signatures are not proper length
     *      - InvalidSignature: Each signature is not valid
     *      - SignersNotOrdered: Signatures are not ordered by signer address ascending
     *      - ReferenceTimestampDoesNotExist: The operatorSet has not been updated for the referenceTimestamp
     *      - OperatorCountZero: There are zero operators for the referenceTimestamp
     *      - VerificationFailed: Any signer is not a registered operator
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
     * @dev This function DOES NOT support smart contact signatures
     * @dev The `referenceTimestamp` in the `ECDSACertificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts for:
     *      - All requirements from verifyCertificate
     *      - ArrayLengthMismatch: signedStakes.length does not equal totalStakeProportionThresholds.length
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
     * @dev This function DOES NOT support smart contact signatures
     * @dev The `referenceTimestamp` in the `ECDSACertificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts for:
     *      - All requirements from verifyCertificate
     *      - ArrayLengthMismatch: signedStakes.length does not equal totalStakeNominalThresholds.length
     */
    function verifyCertificateNominal(
        OperatorSet calldata operatorSet,
        ECDSACertificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external view returns (bool, address[] memory signers);

    /**
     * @notice Get operator infos for a timestamp, which for each operator is the operator's signing key and stake weights
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
     * @dev The index is at most the number of operators in the operatorSet at the given reference timestamp,
     *      which is given by `getOperatorCount`
     * @dev Reverts for:
     *      - IndexOutOfBounds: operatorIndex is greater than or equal to the number of operators
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
     * @notice Calculate the EIP-712 digest bytes for a certificate, returning the raw bytes of the digest
     * @param referenceTimestamp The reference timestamp
     * @param messageHash The message hash of the task
     * @return The EIP-712 digest
     * @dev EIP-712 is a standard ECDSA signature verification framework. See https://eips.ethereum.org/EIPS/eip-712 for more details
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
     * @notice Calculate the EIP-712 digest for a certificate, returning the hash of the digest
     * @param referenceTimestamp The reference timestamp
     * @param messageHash The message hash of the task
     * @return The EIP-712 digest
     * @dev EIP-712 is a standard ECDSA signature verification framework. See https://eips.ethereum.org/EIPS/eip-712 for more details
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
