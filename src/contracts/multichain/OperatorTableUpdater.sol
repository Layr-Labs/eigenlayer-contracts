// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../libraries/Merkle.sol";
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
    ) OperatorTableUpdaterStorage(_bn254CertificateVerifier, _ecdsaCertificateVerifier) SemVerMixin(_version) {
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
            globalRootConfirmerSet, globalTableRootCert, stakeProportionThresholds
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
        // Check that the `referenceTimestamp` is greater than the latest reference timestamp
        require(
            referenceTimestamp > bn254CertificateVerifier.latestReferenceTimestamp(operatorSet),
            TableUpdateForPastTimestamp()
        );

        bytes32 operatorSetLeafHash = keccak256(abi.encode(operatorSetInfo, config));

        // Verify the operator table update
        _verifyOperatorTableUpdate({
            referenceTimestamp: referenceTimestamp,
            globalTableRoot: globalTableRoot,
            operatorSetIndex: operatorSetIndex,
            proof: proof,
            operatorSetLeafHash: operatorSetLeafHash
        });

        // Update the operator table
        bn254CertificateVerifier.updateOperatorTable(operatorSet, referenceTimestamp, operatorSetInfo, config);
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
        // Check that the `referenceTimestamp` is greater than the latest reference timestamp
        require(
            referenceTimestamp > ecdsaCertificateVerifier.latestReferenceTimestamp(operatorSet),
            TableUpdateForPastTimestamp()
        );

        bytes32 operatorSetLeafHash = keccak256(abi.encode(operatorInfos, config));

        // Verify the operator table update
        _verifyOperatorTableUpdate({
            referenceTimestamp: referenceTimestamp,
            globalTableRoot: globalTableRoot,
            operatorSetIndex: operatorSetIndex,
            proof: proof,
            operatorSetLeafHash: operatorSetLeafHash
        });

        // Update the operator table
        ecdsaCertificateVerifier.updateOperatorTable(operatorSet, referenceTimestamp, operatorInfos, config);
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
     * @notice Verifies that the operator table update is valid by checking the `proof` against a `globalTableRoot`
     * @param referenceTimestamp The reference timestamp of the operator table update
     * @param globalTableRoot The global table root of the operator table update
     * @param operatorSetIndex The index of the operator set in the operator table
     * @param proof The proof of the operator table update
     * @param operatorSetLeafHash The leaf hash of the operator set
     * @dev Reverts if there does not exist a `globalTableRoot` for the given `referenceTimestamp`
     */
    function _verifyOperatorTableUpdate(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes32 operatorSetLeafHash
    ) internal view {
        // Check that the `globalTableRoot` matches the `referenceTimestamp`
        require(_globalTableRoots[referenceTimestamp] == globalTableRoot, InvalidGlobalTableRoot());

        // Verify inclusion of the operatorSet and operatorSetLeaf in the merkle tree
        require(
            Merkle.verifyInclusionKeccak({
                proof: proof,
                root: globalTableRoot,
                leaf: operatorSetLeafHash,
                index: operatorSetIndex
            }),
            InvalidOperatorSetProof()
        );
    }

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
