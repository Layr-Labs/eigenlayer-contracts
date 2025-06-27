// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IOperatorTableUpdater.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IECDSACertificateVerifier.sol";

abstract contract OperatorTableUpdaterStorage is IOperatorTableUpdater {
    // Constants

    bytes32 public constant GLOBAL_TABLE_ROOT_CERT_TYPEHASH =
        keccak256("GlobalTableRootCert(bytes32 globalTableRoot,uint32 referenceTimestamp,uint32 referenceBlockNumber)");

    /// @notice The maximum BPS value
    uint16 public constant MAX_BPS = 10_000;

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
    OperatorSet internal _globalRootConfirmerSet;

    /// @notice The global table roots by timestamp
    mapping(uint32 timestamp => bytes32 globalTableRoot) internal _globalTableRoots;

    /// @notice Mapping from latest reference timestamp to reference block number
    mapping(uint32 referenceTimestamp => uint32 referenceBlockNumber) internal _referenceBlockNumbers;

    /// @notice Mapping from reference block number to reference timestamp
    mapping(uint32 referenceBlockNumber => uint32 referenceTimestamp) internal _referenceTimestamps;

    /// @notice Mapping from global table root to validity status
    mapping(bytes32 globalTableRoot => bool valid) internal _isRootValid;

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
    uint256[44] private __gap;
}
