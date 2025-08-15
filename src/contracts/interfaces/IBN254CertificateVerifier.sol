// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {BN254} from "../libraries/BN254.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";
import "./IBaseCertificateVerifier.sol";

interface IBN254CertificateVerifierTypes is IOperatorTableCalculatorTypes {
    /**
     * @notice A witness for an operator, used to identify the non-signers for a given certificate
     * @param operatorIndex the index of the nonsigner in the `BN254OperatorInfo` tree
     * @param operatorInfoProofs merkle proof of the nonsigner at the index.
     *        Leave empty if the non-signing operator is already stored from a previous verification at the same `referenceTimestamp`
     * @param operatorInfo the `BN254OperatorInfo` for the operator.
     *        Leave empty if the non-signing operator is already stored from a previous verification at the same `referenceTimestamp`
     * @dev Non-signing operators are stored in the `BN254CertificateVerifier` upon the first successful certificate verification that includes a merkle proof for the non-signing operator.
     *      This is done to avoid the need to provide proofs for non-signing operators for each certificate with the same `referenceTimestamp`
     */
    struct BN254OperatorInfoWitness {
        uint32 operatorIndex;
        bytes operatorInfoProof;
        BN254OperatorInfo operatorInfo;
    }

    /**
     * @notice A BN254 Certificate
     * @param referenceTimestamp a reference timestamp that corresponds to a timestamp at which an operator table was updated for the operatorSet.
     * @param messageHash the hash of a task that was completed by operators. The messageHash is defined by the AVS, see `TaskMailbox.sol` for an example implementation.
     *                    NOTE: This value is NOT the message that is signed by operators - see `calculateCertificateDigest` for the signable digest.
     * @param signature the G1 signature of the message. The signature is over the signable digest, which is calculated by `calculateCertificateDigest`
     * @param apk the G2 aggregate public key
     * @param nonSignerWitnesses an array of witnesses of non-signing operators
     * @dev The `referenceTimestamp` is used to key into the operatorSet's stake weights. It is NOT the timestamp at which the certificate was generated off-chain
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
    /// @notice thrown when operator index provided in certificate is invalid
    /// @dev Error code: 0x03f4a78e
    /// @dev We enforce that operator indices are within valid bounds to prevent out-of-bounds access in the merkle tree verification
    error InvalidOperatorIndex();
}

/// @notice An interface for verifying BN254 certificates
/// @notice This implements the base `IBaseCertificateVerifier` interface
interface IBN254CertificateVerifier is
    IBN254CertificateVerifierEvents,
    IBaseCertificateVerifier,
    IBN254CertificateVerifierErrors
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
     * @param referenceTimestamp the timestamp at which the operatorSetInfo (i.e. operator table) was sourced
     * @param operatorSetInfo the operator table for this operatorSet. This includes the `totalWeights`, `operatorInfoTreeRoot`, `aggregatePubkey`, and `numOperators`.
     *        See `IOperatorTableCalculator.BN254OperatorSetInfo` for more details
     * @param operatorSetConfig the configuration of the operatorSet, which includes the owner and max staleness period
     * @dev This function can only be called by the `OperatorTableUpdater` contract, which is itself permissionless to call
     * @dev The `referenceTimestamp` must correspond to a reference timestamp for a globalTableRoot stored in the `OperatorTableUpdater`
     *      In addition, it must be greater than the latest reference timestamp for the given operatorSet
     * @dev Reverts for:
     *      - OnlyTableUpdater: Caller is not the authorized OperatorTableUpdater
     *      - TableUpdateStale: referenceTimestamp is not greater than the latest reference timestamp
     * @dev Emits the following events:
     *      - TableUpdated: When operator table is successfully updated
     */
    function updateOperatorTable(
        OperatorSet calldata operatorSet,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo memory operatorSetInfo,
        OperatorSetConfig calldata operatorSetConfig
    ) external;

    /**
     * @notice verifies a certificate against the operator table for a given reference timestamp
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert a certificate
     * @return totalSignedStakeWeights total stake weight that signed the certificate for each stake type. Each
     *         index corresponds to a stake type in the `weights` array in the `BN254OperatorSetInfo` struct
     * @dev The `referenceTimestamp` in the `BN254Certificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts if the certificate's `referenceTimestamp` is too stale with respect to the `maxStalenessPeriod` of the operatorSet
     * @dev This function is *non-view* because it caches non-signing operator info upon a successful certificate verification. See `getNonsignerOperatorInfo` for more details
     * @dev Reverts for:
     *      - CertificateStale: Certificate referenceTimestamp is too stale per maxStalenessPeriod
     *      - ReferenceTimestampDoesNotExist: No operator table exists for the referenceTimestamp
     *      - RootDisabled: The global table root for this timestamp has been disabled
     *      - InvalidOperatorIndex: Operator index provided in nonSigner witness is invalid
     *      - VerificationFailed: Merkle proof verification failed or BLS signature verification failed
     */
    function verifyCertificate(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert
    ) external returns (uint256[] memory totalSignedStakeWeights);

    /**
     * @notice verifies a certificate and makes sure that the signed stakes meet
     * provided portions of the total stake weight on the AVS
     * @param operatorSet the operatorSet that the certificate is for
     * @param cert the certificate
     * @param totalStakeProportionThresholds the proportion, in BPS, of total stake weight that
     *        the signed stake of the certificate should meet. Each index corresponds to
     *        a stake type in the `totalWeights` array in the `BN254OperatorSetInfo`
     * @return Whether or not certificate is valid and meets proportion thresholds
     * @dev The `referenceTimestamp` in the `BN254Certificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts if the certificate's `referenceTimestamp` is too stale with respect to the `maxStalenessPeriod` of the operatorSet
     * @dev This function is *non-view* because it caches non-signing operator info upon a successful certificate verification. See `getNonsignerOperatorInfo` for more details
     * @dev Reverts for:
     *      - CertificateStale: Certificate referenceTimestamp is too stale per maxStalenessPeriod
     *      - ReferenceTimestampDoesNotExist: No operator table exists for the referenceTimestamp
     *      - RootDisabled: The global table root for this timestamp has been disabled
     *      - InvalidOperatorIndex: Operator index provided in nonSigner witness is invalid
     *      - VerificationFailed: Merkle proof verification failed or BLS signature verification failed
     *      - ArrayLengthMismatch: signedStakes length does not equal totalStakeProportionThresholds length
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
     * @param cert the certificate
     * @param totalStakeNominalThresholds the nominal amount of stake that
     *        the signed stake of the certificate should meet. Each index corresponds to
     *        a stake type in the `totalWeights` array in the `BN254OperatorSetInfo`
     * @return Whether or not certificate is valid and meets nominal thresholds
     * @dev The `referenceTimestamp` in the `BN254Certificate` is used to determine the operator table to use for the verification
     * @dev AVS' are responsible for managing potential race conditions when certificates are signed close to operator table updates. Some examples include:
     *      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
     *         wants to only confirm tasks against the *latest* certificate
     *      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet
     * @dev Reverts if the certificate's `referenceTimestamp` is too stale with respect to the `maxStalenessPeriod` of the operatorSet
     * @dev This function is *non-view* because it caches non-signing operator info upon a successful certificate verification. See `getNonsignerOperatorInfo` for more details
     * @dev Reverts for:
     *      - CertificateStale: Certificate referenceTimestamp is too stale per maxStalenessPeriod
     *      - ReferenceTimestampDoesNotExist: No operator table exists for the referenceTimestamp
     *      - RootDisabled: The global table root for this timestamp has been disabled
     *      - InvalidOperatorIndex: Operator index provided in nonSigner witness is invalid
     *      - VerificationFailed: Merkle proof verification failed or BLS signature verification failed
     *      - ArrayLengthMismatch: signedStakes length does not equal totalStakeNominalThresholds length
     */
    function verifyCertificateNominal(
        OperatorSet memory operatorSet,
        BN254Certificate memory cert,
        uint256[] memory totalStakeNominalThresholds
    ) external returns (bool);

    /**
     * @notice Convenience function to attempt signature verification with gas limit for safety
     * @param msgHash The message hash that was signed
     * @param aggPubkey The aggregate public key of signers
     * @param apkG2 The G2 point representation of the aggregate public key
     * @param signature The BLS signature to verify
     * @return pairingSuccessful Whether the pairing operation completed successfully
     * @return signatureValid Whether the signature is valid
     * @dev This function should be used off-chain to validate a signature. Careful consideration should be taken
     *      when parsing `pairingSuccessful` and `signatureValid`. Refer to our internal usage of this function
     */
    function trySignatureVerification(
        bytes32 msgHash,
        BN254.G1Point memory aggPubkey,
        BN254.G2Point memory apkG2,
        BN254.G1Point memory signature
    ) external view returns (bool pairingSuccessful, bool signatureValid);

    /**
     * @notice Get cached nonsigner operator info
     * @param operatorSet The operatorSet
     * @param referenceTimestamp The reference timestamp
     * @param operatorIndex The operator index
     * @return The cached operator info, empty if the operator is not in the cache
     * @dev The non-signing operatorInfo is stored upon a successful certificate verification. Once cached,
     *      merkle proofs for non-signing operators do not need to be passed in as part of the `BN254Certificate` for a given reference timestamp
     * @dev Non-signing operators are stored on the `operatorInfoTreeRoot` that is transported in the `BN254OperatorSetInfo` struct on an operator table update
     * @dev The tree structure of the `operatorInfoTreeRoot` is as follows. (Below is a tree of height 2 -- in practice, each tree will have a height appropriate for the total number of leaves.)
     * ```
     *                    OperatorInfoTreeRoot
     *                           |
     *                    ┌──────┴──────┐
     *                    │             │
     *              Internal Node    Internal Node
     *               (0-1)           (2-3)
     *                    │             │
     *              ┌─────┴─────┐   ┌───┴───┐
     *              │           │   │       │
     *         Leaf 0        Leaf 1  Leaf 2  Leaf 3
     *    BN254OperatorInfo  #1      #2      #3
     *         #0
     * ```
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
     * @dev The non-signing operatorInfo is stored upon a successful certificate verification. Once cached,
     *      merkle proofs for non-signing operators do not need to be passed in as part of the `BN254Certificate` for a given reference timestamp
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
     * @return The operator set info, empty if the operatorSet has not been updated for the given reference timestamp
     */
    function getOperatorSetInfo(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (BN254OperatorSetInfo memory);

    /**
     * @notice Calculate the digest for a certificate, which must be signed over by operators who complete a task
     * @param referenceTimestamp The reference timestamp
     * @param messageHash The message hash of the task
     * @return The digest
     * @dev This is a chain-agnostic digest, so it can be used to verify certificates across
     *      multiple destination chains
     */
    function calculateCertificateDigest(
        uint32 referenceTimestamp,
        bytes32 messageHash
    ) external pure returns (bytes32);
}
