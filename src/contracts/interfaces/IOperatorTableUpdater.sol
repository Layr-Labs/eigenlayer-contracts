// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../libraries/OperatorSetLib.sol";

import "./IECDSATableCalculator.sol";
import "./IBN254TableCalculator.sol";
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
}

interface IOperatorTableUpdaterEvents {
    /**
     * @notice Emitted when a new global table root is set
     * @param referenceTimestamp the timestamp of the global table root
     * @param globalTableRoot the root of the global table
     */
    event NewGlobalTableRoot(uint32 indexed referenceTimestamp, bytes32 indexed globalTableRoot);

    /**
     * @notice Emitted when the global root confirmer set is updated
     * @param operatorSet The operatorSet which certifies against global roots
     */
    event GlobalRootConfirmerSetUpdated(OperatorSet operatorSet);

    /**
     * @notice Emitted when the global root confirmation threshold is updated
     * @param bps The threshold, in bps, for a global root to be signed off on and updated
     */
    event GlobalRootConfirmationThresholdUpdated(uint16 bps);
}

interface IOperatorTableUpdater is
    IOperatorTableUpdaterErrors,
    IOperatorTableUpdaterEvents,
    IECDSACertificateVerifierTypes,
    IBN254CertificateVerifierTypes,
    IKeyRegistrarTypes,
    ICrossChainRegistryTypes
{
    /**
     * @notice Confirms Global operator table root
     * @param globalTableRootCert certificate of the root
     * @param globalTableRoot merkle root of all operatorSet tables
     * @param referenceTimestamp timestamp of the root
     * @param referenceBlockNumber block number of the root
     * @dev Any entity can submit with a valid certificate signed off by the `globalRootConfirmerSet`
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
    function setGlobalRootConfirmerSet(
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
    function getGlobalRootConfirmerSet() external view returns (OperatorSet memory);

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
     * @notice Get the reference timestamp of the global confirmer set
     * @return The reference timestamp of the global confirmer set
     * @dev In V1, we only update the table of the global root confirmer set on initial deployment, and never update it again.
     */
    function getGlobalConfirmerSetReferenceTimestamp() external view returns (uint32);
}
