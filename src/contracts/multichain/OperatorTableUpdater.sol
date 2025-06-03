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
        OperatorSet calldata _globalRootConfirmerSet,
        uint16 _globalRootConfirmationThreshold
    ) external initializer {
        _setGlobalRootConfirmerSet(_globalRootConfirmerSet);
        _setGlobalRootConfirmationThreshold(_globalRootConfirmationThreshold);
    }

    /**
     *
     *                         ACTIONS
     *
     */

    /// @inheritdoc IOperatorTableUpdater
    function confirmGlobalTableRoot(
        BN254Certificate calldata globalTableRootCert,
        bytes32 globalTableRoot,
        uint32 referenceTimestamp
    ) external {
        // Verify certificate by using the stake proportion thresholds
        uint16[] memory stakeProportionThresholds = new uint16[](1);
        stakeProportionThresholds[0] = globalRootConfirmationThreshold;
        bool isValid = bn254CertificateVerifier.verifyCertificateProportion(
            globalRootConfirmerSet,
            globalTableRootCert,
            stakeProportionThresholds
        );

        require(isValid, CertificateInvalid());

        // Update the global table root
        _currentGlobalTableRoot = globalTableRoot;
        _globalTableRoots[referenceTimestamp] = globalTableRoot;

        emit NewGlobalTableRoot(referenceTimestamp, globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateBN254OperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        OperatorSet calldata operatorSet,
        BN254OperatorSetInfo calldata operatorSetInfo,
        OperatorSetConfig calldata config
    ) external {
        // TODO: Implement
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateECDSAOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        OperatorSet calldata operatorSet,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata config
    ) external {
        // TODO: Implement
    }

    /**
     *
     *                         SETTERS
     *
     */

    /// @inheritdoc IOperatorTableUpdater
    function setGlobalRootConfirmerSet(
        OperatorSet calldata operatorSet
    ) external onlyOwner {
        _setGlobalRootConfirmerSet(operatorSet);
    }

    /// @inheritdoc IOperatorTableUpdater
    function setGlobalRootConfirmationThreshold(
        uint16 bps
    ) external onlyOwner {
        _setGlobalRootConfirmationThreshold(bps);
    }

    /**
     *
     *                         GETTERS
     *
     */

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalTableRootByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (bytes32) {
        return _globalTableRoots[referenceTimestamp];
    }

    /// @inheritdoc IOperatorTableUpdater
    function getCurrentGlobalTableRoot() external view returns (bytes32) {
        return _currentGlobalTableRoot;
    }

    /**
     *
     *                         INTERNAL HELPERS
     *
     */

    /**
     * @notice Sets the global root confirmer set
     * @param operatorSet The operatorSet which certifies against global roots
     */
    function _setGlobalRootConfirmerSet(
        OperatorSet calldata operatorSet
    ) internal {
        globalRootConfirmerSet = operatorSet;
        emit GlobalRootConfirmerSetUpdated(operatorSet);
    }

    /**
     * @notice Sets the global root confirmation threshold
     * @param bps The threshold, in bps, for a global root to be signed off on and updated
     */
    function _setGlobalRootConfirmationThreshold(
        uint16 bps
    ) internal {
        globalRootConfirmationThreshold = bps;
        emit GlobalRootConfirmationThresholdUpdated(bps);
    }
}