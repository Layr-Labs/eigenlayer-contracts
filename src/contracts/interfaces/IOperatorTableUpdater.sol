// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../libraries/OperatorSetLib.sol";

import "./IECDSACertificateVerifier.sol";
import "./IBN254CertificateVerifier.sol";
import "./IKeyRegistrar.sol";
import "./ICrossChainRegistry.sol";

interface IOperatorTableUpdaterErrors {
    /// @notice Thrown when the global table root is in the future
    error GlobalTableRootInFuture();
    /// @notice Thrown when the global table root is stale
    error GlobalTableRootStale();
    /// @notice Thrown when the table root does not match what is in the certificate
    error InvalidMessageHash();
    /// @notice Thrown when the GlobalTableRoot update fails
    error CertificateInvalid();
    /// @notice Thrown when the table has been updated for the timestamp
    error TableUpdateForPastTimestamp();
    /// @notice Thrown when the global table root does not match what is in storage
    error InvalidGlobalTableRoot();
    /// @notice Thrown when the operator set proof is invalid
    error InvalidOperatorSetProof();
    /// @notice Thrown when the confirmation threshold is invalid
    error InvalidConfirmationThreshold();
    /// @notice Thrown when the curve type is invalid
    error InvalidCurveType();
    /// @notice Thrown when a root is invalid
    error InvalidRoot();
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
     * @notice Confirms Global operator table root
     * @param globalTableRootCert certificate of the root
     * @param globalTableRoot merkle root of all operatorSet tables
     * @param referenceTimestamp timestamp of the root
     * @param referenceBlockNumber block number of the root
     * @dev Any entity can submit with a valid certificate signed off by the `Generator`
     * @dev The `msgHash` in the `globalOperatorTableRootCert` is the hash of the `globalOperatorTableRoot`
     */
    function confirmGlobalTableRoot(
        BN254Certificate calldata globalTableRootCert,
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) external;

    /**
     * @notice Set the operatorSet which certifies against global roots
     * @param operatorSet the operatorSet which certifies against global roots
     * @dev The `operatorSet` is used to verify the certificate of the global table root
     * @dev Only callable by the owner of the contract
     */
    function setGenerator(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice The threshold, in bps, for a global root to be signed off on and updated
     * @dev Only callable by the owner of the contract
     */
    function setGlobalRootConfirmationThreshold(
        uint16 bps
    ) external;

    /**
     * @notice Updates the operator table for the generator
     * @param referenceTimestamp The reference timestamp of the operator table update
     * @param GeneratorInfo The operatorSetInfo for the generator
     * @param GeneratorConfig The operatorSetConfig for the generator
     * @dev We have a separate function for updating this operatorSet since it's not transported and updated
     *      in the same way as the other operatorSets
     * @dev Only callable by the owner of the contract
     */
    function updateGenerator(
        uint32 referenceTimestamp,
        BN254OperatorSetInfo calldata GeneratorInfo,
        OperatorSetConfig calldata GeneratorConfig
    ) external;

    /**
     * @notice Updates an operator table
     * @param referenceTimestamp the reference block number of the globalTableRoot
     * @param globalTableRoot the new globalTableRoot
     * @param operatorSetIndex the index of the given operatorSet being updated
     * @param proof the proof of the leaf at index against the globalTableRoot
     * @param operatorTableBytes the bytes of the operator table
     * @dev Depending on the decoded KeyType, the tableInfo will be decoded
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
     * @dev In V1, we only update the table of the generator on initial deployment, and never update it again.
     */
    function getGeneratorReferenceTimestamp() external view returns (uint32);

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
