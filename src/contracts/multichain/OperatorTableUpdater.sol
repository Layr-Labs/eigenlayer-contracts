// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../libraries/Merkle.sol";
import "../mixins/SemVerMixin.sol";
import "./OperatorTableUpdaterStorage.sol";

contract OperatorTableUpdater is Initializable, OwnableUpgradeable, OperatorTableUpdaterStorage, SemVerMixin {
    enum KeyType {
        NONE,
        BN254,
        ECDSA
    }

    error InvalidKeyType();


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
     * @param owner The owner of the OperatorTableUpdater
     * @param _globalRootConfirmerSet The operatorSet which certifies against global roots
     * @param _globalRootConfirmationThreshold The threshold, in bps, for a global root to be signed off on and updated
     * @param referenceTimestamp The reference timestamp for the global root confirmer set
     * @param globalRootConfirmerSetInfo The operatorSetInfo for the global root confirmer set
     * @param globalRootConfirmerSetConfig The operatorSetConfig for the global root confirmer set
     * @dev We also update the operator table for the global root confirmer set, to begin signing off on global roots
     */
    function initialize(
        address owner,
        OperatorSet calldata _globalRootConfirmerSet,
        uint16 _globalRootConfirmationThreshold,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo calldata globalRootConfirmerSetInfo,
        OperatorSetConfig calldata globalRootConfirmerSetConfig
    ) external initializer {
        _transferOwnership(owner);
        _setGlobalRootConfirmerSet(_globalRootConfirmerSet);
        _setGlobalRootConfirmationThreshold(_globalRootConfirmationThreshold);

        // Update the operator table for the global root confirmer set
        bn254CertificateVerifier.updateOperatorTable(
            _globalRootConfirmerSet, referenceTimestamp, globalRootConfirmerSetInfo, globalRootConfirmerSetConfig
        );
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
        // Table roots can only be updated for current or past timestamps and after the latest reference timestamp
        require(referenceTimestamp <= block.timestamp, GlobalTableRootInFuture());
        require(referenceTimestamp > latestReferenceTimestamp, GlobalTableRootStale());
        require(globalTableRoot == globalTableRootCert.messageHash, TableRootNotInCertificate());

        // Verify certificate by using the stake proportion thresholds
        uint16[] memory stakeProportionThresholds = new uint16[](1);
        stakeProportionThresholds[0] = globalRootConfirmationThreshold;
        bool isValid = bn254CertificateVerifier.verifyCertificateProportion(
            _globalRootConfirmerSet, globalTableRootCert, stakeProportionThresholds
        );

        require(isValid, CertificateInvalid());

        // Update the global table root
        latestReferenceTimestamp = referenceTimestamp;
        _currentGlobalTableRoot = globalTableRoot;
        _globalTableRoots[referenceTimestamp] = globalTableRoot;

        emit NewGlobalTableRoot(referenceTimestamp, globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes calldata tableInfo
    ) external {
        (OperatorSet memory operatorSet, KeyType keyType, OperatorSetConfig memory operatorSetConfig) = _getOperatorTableInfo(tableInfo);

        // Check that the `referenceTimestamp` is greater than the latest reference timestamp
        require(
            referenceTimestamp > IBaseCertificateVerifier(_getCertificateVerifier(keyType)).latestReferenceTimestamp(operatorSet),
            TableUpdateForPastTimestamp()
        );

        bytes32 operatorSetLeafHash = keccak256(tableInfo);

        _verifyOperatorTableUpdate({
            referenceTimestamp: referenceTimestamp,
            globalTableRoot: globalTableRoot,
            operatorSetIndex: operatorSetIndex,
            proof: proof,
            operatorSetLeafHash: operatorSetLeafHash
        });

        if (keyType == KeyType.BN254) {
            bn254CertificateVerifier.updateOperatorTable(operatorSet, referenceTimestamp, _getBN254OperatorSetInfo(tableInfo), operatorSetConfig);
        } else if (keyType == KeyType.ECDSA) {
            ecdsaCertificateVerifier.updateOperatorTable(operatorSet, referenceTimestamp, _getECDSAOperatorSetInfo(tableInfo), operatorSetConfig);
        } else {
            revert InvalidKeyType();
        }
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

        bytes32 operatorSetLeafHash = keccak256(abi.encode(operatorSet.key(), operatorSetInfo, config));

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

        bytes32 operatorSetLeafHash = keccak256(abi.encode(operatorSet.key(), operatorInfos, config));

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

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalRootConfirmerSet() external view returns (OperatorSet memory) {
        return _globalRootConfirmerSet;
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
        _globalRootConfirmerSet = operatorSet;
        emit GlobalRootConfirmerSetUpdated(operatorSet);
    }

    /**
     * @notice Sets the global root confirmation threshold
     * @param bps The threshold, in bps, for a global root to be signed off on and updated
     */
    function _setGlobalRootConfirmationThreshold(
        uint16 bps
    ) internal {
        require(bps <= MAX_BPS, InvalidConfirmationThreshold());
        globalRootConfirmationThreshold = bps;
        emit GlobalRootConfirmationThresholdUpdated(bps);
    }

    function _getOperatorTableInfo(
        bytes calldata tableInfo
    ) internal pure returns (OperatorSet memory operatorSet, KeyType keyType, OperatorSetConfig memory operatorSetInfo) {
        (operatorSet, keyType, operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig));
    }

    function _getBN254OperatorSetInfo(
        bytes calldata tableInfo
    ) internal pure returns (BN254OperatorSetInfo memory operatorSetInfo) {
        (,,,operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig, BN254OperatorSetInfo));
    }

    function _getECDSAOperatorSetInfo(
        bytes calldata tableInfo
    ) internal pure returns (ECDSAOperatorInfo[] memory operatorSetInfo) {
        (,,,operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig, ECDSAOperatorInfo[]));
    }

    function _getCertificateVerifier(
        KeyType keyType
    ) internal view returns (address) {
        if (keyType == KeyType.BN254) {
            return address(bn254CertificateVerifier);
        } else if (keyType == KeyType.ECDSA) {
            return address(ecdsaCertificateVerifier);
        } else {
            revert InvalidKeyType();
        }
    }
}
