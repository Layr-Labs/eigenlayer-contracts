// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IOperatorTableUpdater.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IECDSACertificateVerifier.sol";

abstract contract OperatorTableUpdaterStorage is IOperatorTableUpdater {
    // Constants

    /// @notice Index for flag that pauses calling `updateGlobalTableRoot`
    uint8 internal constant PAUSED_GLOBAL_ROOT_UPDATE = 0;

    /// @notice Index for flag that pauses calling `updateOperatorTable`
    uint8 internal constant PAUSED_OPERATOR_TABLE_UPDATE = 1;

    // OPERATOR_TABLE_LEAF_SALT is now inherited from LeafCalculatorMixin

    bytes32 public constant GLOBAL_TABLE_ROOT_CERT_TYPEHASH =
        keccak256("GlobalTableRootCert(bytes32 globalTableRoot,uint32 referenceTimestamp,uint32 referenceBlockNumber)");

    /// @notice The maximum BPS value
    uint16 public constant MAX_BPS = 10_000;

    /// @notice Dummy initial global table root to break circular dependency for certificate verification
    bytes32 public constant GENERATOR_GLOBAL_TABLE_ROOT = keccak256("GENERATOR_GLOBAL_TABLE_ROOT");

    /// @notice The reference timestamp for the generator
    uint32 public constant GENERATOR_REFERENCE_TIMESTAMP = 1;

    /// @notice The `maxStalenessPeriods` for the generator
    /// @dev This is set to 0 to allow certificates to always be valid, regardless of the `referenceTimestamp`
    uint32 public constant GENERATOR_MAX_STALENESS_PERIOD = 0;

    // Immutable Storage

    /// @notice The BN254 certificate verifier
    IBN254CertificateVerifier public immutable bn254CertificateVerifier;

    /// @notice The ECDSA certificate verifier
    IECDSACertificateVerifier public immutable ecdsaCertificateVerifier;

    // Mutable Storage

    /// @notice The threshold, in bps, for a global root to be signed off on and updated
    uint16 public globalRootConfirmationThreshold;

    /// @notice The latest reference timestamp
    uint32 internal _latestReferenceTimestamp;

    /// @notice The operatorSet which certifies against global roots
    OperatorSet internal _generator;

    /// @notice The global table roots by timestamp
    mapping(uint32 timestamp => bytes32 globalTableRoot) internal _globalTableRoots;

    /// @notice Mapping from latest reference timestamp to reference block number
    mapping(uint32 referenceTimestamp => uint32 referenceBlockNumber) internal _referenceBlockNumbers;

    /// @notice Mapping from reference block number to reference timestamp
    mapping(uint32 referenceBlockNumber => uint32 referenceTimestamp) internal _referenceTimestamps;

    /// @notice Mapping from global table root to validity status
    mapping(bytes32 globalTableRoot => bool valid) internal _isRootValid;

    /// @notice The operatorSetConfig for the generator
    /// @dev The `maxStalenessPeriod` is set to `GENERATOR_MAX_STALENESS_PERIOD` to allow certificates to always be valid, regardless of the `referenceTimestamp`
    /// @dev The `owner` is set to the address of the `OperatorTableUpdater`
    OperatorSetConfig internal _generatorConfig;

    // Constructor
    constructor(
        IBN254CertificateVerifier _bn254CertificateVerifier,
        IECDSACertificateVerifier _ecdsaCertificateVerifier
    ) {
        bn254CertificateVerifier = _bn254CertificateVerifier;
        ecdsaCertificateVerifier = _ecdsaCertificateVerifier;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
