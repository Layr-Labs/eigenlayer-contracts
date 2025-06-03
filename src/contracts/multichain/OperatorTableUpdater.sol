// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../mixins/SemVerMixin.sol";
import "./OperatorTableUpdaterStorage.sol";


contract OperatorTableUpdater is Initializable, OwnableUpgradeable, OperatorTableUpdaterStorage, SemVerMixin {

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    constructor(
        IBN254CertificateVerifier _bn254CertificateVerifier,
        IECDSACertificateVerifier _ecdsaCertificateVerifier,
        string memory _version
    )
        OperatorTableUpdaterStorage(_bn254CertificateVerifier, _ecdsaCertificateVerifier)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /**
     * @notice Initializes the OperatorTableUpdater
     * @param _globalRootConfirmerSet The operatorSet which certifies against global roots
     * @param _globalRootConfirmationThreshold The threshold, in bps, for a global root to be signed off on and updated
     */
    function initialize(
        OperatorSet memory _globalRootConfirmerSet,
        uint16 _globalRootConfirmationThreshold
    ) external initializer {
        globalRootConfirmerSet = _globalRootConfirmerSet;
        globalRootConfirmationThreshold = _globalRootConfirmationThreshold;
    }
}