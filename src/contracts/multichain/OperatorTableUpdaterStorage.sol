// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IOperatorTableUpdater.sol";
import "../interfaces/IBN254CertificateVerifier.sol";
import "../interfaces/IECDSACertificateVerifier.sol";

abstract contract OperatorTableUpdaterStorage is IOperatorTableUpdater {
    // Immutable Storage

    /// @notice The BN254 certificate verifier
    IBN254CertificateVerifier public immutable bn254CertificateVerifier;

    /// @notice The ECDSA certificate verifier
    IECDSACertificateVerifier public immutable ecdsaCertificateVerifier;

    // Mutable Storage

    /// @notice The threshold, in bps, for a global root to be signed off on and updated
    uint16 public globalRootConfirmationThreshold;

    /// @notice The operatorSet which certifies against global roots
    OperatorSet public globalRootConfirmerSet;

    /// @notice The current global table root
    bytes32 internal _currentGlobalTableRoot;

    /// @notice The global table roots by timestamp
    mapping(uint32 timestamp => bytes32 globalTableRoot) internal _globalTableRoots;

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
    uint256[46] private __gap;
}
