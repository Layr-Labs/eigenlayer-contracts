// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../libraries/OperatorSetLib.sol";

import "./IECDSACertificateVerifier.sol";
import "./IBN254CertificateVerifier.sol";
import "./IKeyRegistrar.sol";
import "./ICrossChainRegistry.sol";

interface IOperatorTableUpdaterErrors {
    /// @notice Thrown when the global table root is in the future
    /// @dev 4byte error code: 0xb4233b6a
    /// @dev We enforce that reference timestamps are not in the future to prevent manipulation and ensure temporal consistency of operator table data
    error GlobalTableRootInFuture();
    /// @notice Thrown when the global table root is stale
    /// @dev 4byte error code: 0x1bfd4358
    /// @dev We enforce that new reference timestamps are greater than the latest to maintain chronological ordering and prevent replay attacks
    error GlobalTableRootStale();
    /// @notice Thrown when the table root does not match what is in the certificate
    /// @dev 4byte error code: 0x8b56642d
    /// @dev We enforce that the certificate's message hash matches the expected EIP-712 hash to ensure integrity of the global table root data
    error InvalidMessageHash();
    /// @notice Thrown when the GlobalTableRoot update fails
    /// @dev 4byte error code: 0xc108107c
    /// @dev We enforce valid certificate signatures to maintain cryptographic security guarantees for global table root confirmations
    error CertificateInvalid();
    /// @notice Thrown when the table has been updated for the timestamp
    /// @dev 4byte error code: 0x207617df
    /// @dev We enforce that operator table updates use newer timestamps than existing ones to prevent retroactive modifications
    error TableUpdateForPastTimestamp();
    /// @notice Thrown when the global table root does not match what is in storage
    /// @dev 4byte error code: 0xc73a136a
    /// @dev We enforce that the provided global table root matches the stored root at the reference timestamp to ensure data consistency
    error InvalidGlobalTableRoot();
    /// @notice Thrown when the operator set proof is invalid
    /// @dev 4byte error code: 0xafa42ca7
    /// @dev We enforce valid merkle proofs to ensure operator table data is properly included in the global table root
    error InvalidOperatorSetProof();
    /// @notice Thrown when the confirmation threshold is invalid
    /// @dev 4byte error code: 0x0e66de06
    /// @dev We enforce that confirmation thresholds do not exceed 100% (10000 BPS) to maintain valid percentage-based validation
    error InvalidConfirmationThreshold();
    /// @notice Thrown when the curve type is invalid
    /// @dev 4byte error code: 0xfdea7c09
    /// @dev We enforce that only supported curve types (BN254, ECDSA) are used to ensure proper certificate verification routing
    error InvalidCurveType();
    /// @notice Thrown when a root is invalid
    /// @dev 4byte error code: 0x504570e3
    /// @dev We enforce that roots are valid and not disabled to prevent operations against compromised or invalid operator table data
    error InvalidRoot();
    /// @notice Thrown when the generator is invalid (via a non-zero reference timestamp)
    /// @dev 4byte error code: 0x6446f917
    /// @dev We enforce that only new operator sets with zero reference timestamps can become generators to prevent conflicts with existing data
    error InvalidGenerator();
    /// @notice Thrown when the operator set to update is the generator
    /// @dev 4byte error code: 0x7ec5c154
    /// @dev We enforce that the generator operator set is updated separately to maintain the special handling required for global root confirmation
    error InvalidOperatorSet();
    /// @notice Thrown when the generator's global table root is being disabled
    /// @dev 4byte error code: 0x332415fa
    /// @dev We enforce that the generator's root cannot be disabled to maintain the system's ability to confirm new global table roots
    error CannotDisableGeneratorRoot();
}

interface IOperatorTableUpdaterEvents {
    /**
     * @notice Emitted when a new global table root is set
     * @param referenceTimestamp the timestamp of the global table root
     * @param globalTableRoot the root of the global table
     */
    event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot);

    /**
     * @notice Emitted when the generator is updated
     * @param operatorSet The operatorSet which certifies against global roots
     */
    event GeneratorUpdated(OperatorSet operatorSet);

    /**
     * @notice Emitted when the global root confirmation threshold is updated
     * @param bps The threshold, in bps, for a global root to be signed off on and updated
     */
    event GlobalRootConfirmationThresholdUpdated(uint16 bps);

    /**
     * @notice Emitted when a global table root is disabled
     * @param globalTableRoot the global table root that was disabled
     */
    event GlobalRootDisabled(bytes32 indexed globalTableRoot);
}

interface IOperatorTableUpdater is
    IOperatorTableUpdaterErrors,
    IOperatorTableUpdaterEvents,
    IBN254CertificateVerifierTypes,
    IECDSACertificateVerifierTypes,
    IKeyRegistrarTypes,
    ICrossChainRegistryTypes
{
    /**
     * @notice Sets the global table root
     * @param globalTableRootCert certificate of the global table root, signed by the `Generator`
     * @param globalTableRoot merkle root of all operatorSet tables
     * @param referenceTimestamp block timestamp at which the global table root was calculated
     * @param referenceBlockNumber block number, corresponding to the `referenceTimestamp` of the global table root
     * @dev Any entity can submit with a valid certificate signed off by the `Generator`
     * @dev The `msgHash` in the `globalOperatorTableRootCert` is the hash of the `globalTableRoot`, `referenceTimestamp`, and `referenceBlockNumber`
     * @dev The `referenceTimestamp` nested in the `globalTableRootCert` should be `getGeneratorReferenceTimestamp`, whereas
     *      the `referenceTimestamp` passed directly in the calldata is the block timestamp at which the global table root was calculated
     * @dev Reverts for:
     * - GlobalTableRootInFuture: referenceTimestamp is in the future
     * - GlobalTableRootStale: referenceTimestamp is not greater than latest reference timestamp
     * - InvalidMessageHash: certificate messageHash doesn't match expected EIP-712 hash
     * - CertificateInvalid: certificate signature is invalid or doesn't meet confirmation threshold
     * @dev Emits the following events:
     * - NewGlobalTableRoot: when the global table root is successfully confirmed
     */
    function confirmGlobalTableRoot(
        BN254Certificate calldata globalTableRootCert,
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) external;

    /**
     * @notice The threshold, in bps, for a global root to be signed off on and updated
     * @dev Only callable by the owner of the contract
     * @dev Reverts for:
     * - InvalidConfirmationThreshold: bps exceeds MAX_BPS (10000)
     * @dev Emits the following events:
     * - GlobalRootConfirmationThresholdUpdated: when threshold is successfully updated
     */
    function setGlobalRootConfirmationThreshold(
        uint16 bps
    ) external;

    /**
     * @notice Updates the `Generator` to a new operatorSet
     * @param generator The operatorSet which certifies against global roots
     * @param generatorInfo The operatorSetInfo for the generator
     * @dev We have a separate function for updating this operatorSet since it's not transported and updated
     *      in the same way as the other operatorSets
     * @dev Only callable by the owner of the contract
     * @dev Uses GENERATOR_GLOBAL_TABLE_ROOT constant to break circular dependency for certificate verification
     * @dev We ensure that there are no collisions with other reference timestamps because we expect the generator to have an initial reference timestamp of 0
     * @dev The `_latestReferenceTimestamp` is not updated since this root is ONLY used for the `Generator`
     * @dev The `_referenceBlockNumber` and `_referenceTimestamps` mappings are not updated since they are only used for introspection for official operatorSets
     * @dev Reverts for:
     * - InvalidGenerator: generator's latest reference timestamp is not zero
     * - OnlyTableUpdater: (from BN254CertificateVerifier) caller is not the operator table updater
     * - TableUpdateStale: (from BN254CertificateVerifier) reference timestamp validation fails
     * @dev Emits the following events:
     * - GeneratorUpdated: when the generator is successfully updated
     * - TableUpdated: (from BN254CertificateVerifier) when the generator's operator table is updated
     */
    function updateGenerator(OperatorSet calldata generator, BN254OperatorSetInfo calldata generatorInfo) external;

    /**
     * @notice Updates an operator table
     * @param referenceTimestamp the reference timestamp of the globalTableRoot
     * @param globalTableRoot the new globalTableRoot
     * @param operatorSetIndex the index of the given operatorSet being updated
     * @param proof the proof of the leaf at index against the globalTableRoot
     * @param operatorTableBytes the bytes of the operator table
     * @dev This function calls `updateOperatorTable` on the `ECDSACertificateVerifier` or `BN254CertificateVerifier`
     *      depending on the `KeyType` of the operatorSet, which is encoded in the `operatorTableBytes`
     * @dev Reverts for:
     * - InvalidRoot: globalTableRoot is disabled or invalid
     * - InvalidOperatorSet: operatorSet is the generator (must use updateGenerator instead)
     * - TableUpdateForPastTimestamp: referenceTimestamp is not greater than latest for the operatorSet
     * - InvalidGlobalTableRoot: provided globalTableRoot doesn't match stored root at referenceTimestamp
     * - InvalidOperatorSetProof: merkle proof verification fails
     * - InvalidCurveType: operatorTableBytes contains unsupported curve type
     * - OnlyTableUpdater: (from CertificateVerifier) caller validation fails
     * - TableUpdateStale: (from CertificateVerifier) timestamp validation fails
     * @dev Emits the following events:
     * - TableUpdated: (from CertificateVerifier) when the operator table is successfully updated
     */
    function updateOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes calldata operatorTableBytes
    ) external;

    /**
     * @notice Disables a global table root
     * @param globalTableRoot the global table root to disable
     * @dev Only callable by the pauser
     * @dev Cannot disable the GENERATOR_GLOBAL_TABLE_ROOT
     * @dev Reverts for:
     * - InvalidRoot: globalTableRoot doesn't exist or is already disabled
     * - CannotDisableGeneratorRoot: attempting to disable the generator's global table root
     * @dev Emits the following events:
     * - GlobalRootDisabled: when the global table root is successfully disabled
     */
    function disableRoot(
        bytes32 globalTableRoot
    ) external;

    /**
     * @notice Get the current global table root
     * @return globalTableRoot the current global table root
     */
    function getCurrentGlobalTableRoot() external view returns (bytes32 globalTableRoot);

    /**
     * @notice Get the table root by timestamp
     * @param referenceTimestamp the timestamp of the table root
     * @return tableRoot the table root at the given timestamp
     */
    function getGlobalTableRootByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (bytes32 tableRoot);

    /**
     * @notice Get the operatorSet which certifies against global roots
     * @return The operatorSet which certifies against global roots
     */
    function getGenerator() external view returns (OperatorSet memory);

    /**
     * @notice Get the certificate verifier for a given key type
     * @param curveType The curve type
     * @return The certificate verifier for the given key type
     */
    function getCertificateVerifier(
        CurveType curveType
    ) external view returns (address);

    /**
     * @notice Get the latest reference timestamp
     * @return The latest reference timestamp
     */
    function getLatestReferenceTimestamp() external view returns (uint32);

    /**
     * @notice Get the latest reference block number
     * @return The latest reference block number
     */
    function getLatestReferenceBlockNumber() external view returns (uint32);

    /**
     * @notice Get the reference block number for a given reference timestamp
     * @param referenceTimestamp the reference timestamp
     * @return The reference block number for the given reference timestamp
     */
    function getReferenceBlockNumberByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (uint32);

    /**
     * @notice Get the reference timestamp for a given reference block number
     * @param referenceBlockNumber the reference block number
     * @return The reference timestamp for the given reference block number
     */
    function getReferenceTimestampByBlockNumber(
        uint32 referenceBlockNumber
    ) external view returns (uint32);

    /**
     * @notice Get the message hash for the certificate of a global table root update
     * @param globalTableRoot the global table root
     * @param referenceTimestamp the reference timestamp
     * @param referenceBlockNumber the reference block number
     * @return The message hash for a global table root
     */
    function getGlobalTableUpdateMessageHash(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) external view returns (bytes32);

    /**
     * @notice Get the reference timestamp of the generator
     * @return The reference timestamp of the generator
     * @dev The `Generator's` referenceTimestamp is hardcoded to 1. See `GENERATOR_REFERENCE_TIMESTAMP` in `OperatorTableUpdaterStorage.sol`
     */
    function getGeneratorReferenceTimestamp() external view returns (uint32);

    /**
     * @notice Get the operator set config for the Generator
     * @return The operator set config for the Generator
     * @dev The Generator's config has maxStalenessPeriod = 0 and owner = address(operatorTableUpdater)
     */
    function getGeneratorConfig() external view returns (OperatorSetConfig memory);

    /**
     * @notice Get the validity status of a global table root
     * @param globalTableRoot the global table root
     * @return The validity status of the global table root
     */
    function isRootValid(
        bytes32 globalTableRoot
    ) external view returns (bool);

    /**
     * @notice Get the validity status of a global table root by timestamp
     * @param referenceTimestamp the reference timestamp
     * @return The validity status of the global table root
     */
    function isRootValidByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (bool);
}
