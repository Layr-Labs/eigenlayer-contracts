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

        // Set the latest reference timestamp
        _latestReferenceTimestamp = referenceTimestamp;
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
        require(referenceTimestamp > _latestReferenceTimestamp, GlobalTableRootStale());
        require(
            globalTableRootCert.messageHash == getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp),
            InvalidMessageHash()
        );

        // Verify certificate by using the stake proportion thresholds
        uint16[] memory stakeProportionThresholds = new uint16[](1);
        stakeProportionThresholds[0] = globalRootConfirmationThreshold;
        bool isValid = bn254CertificateVerifier.verifyCertificateProportion(
            _globalRootConfirmerSet, globalTableRootCert, stakeProportionThresholds
        );

        require(isValid, CertificateInvalid());

        // Update the global table root
        _latestReferenceTimestamp = referenceTimestamp;
        _globalTableRoots[referenceTimestamp] = globalTableRoot;

        emit NewGlobalTableRoot(referenceTimestamp, globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes calldata operatorTableBytes
    ) external {
        (
            OperatorSet memory operatorSet,
            CurveType curveType,
            OperatorSetConfig memory operatorSetConfig,
            bytes memory operatorTableInfo
        ) = _decodeOperatorTableBytes(operatorTableBytes);

        // Check that the `referenceTimestamp` is greater than the latest reference timestamp
        require(
            referenceTimestamp
                > IBaseCertificateVerifier(getCertificateVerifier(curveType)).latestReferenceTimestamp(operatorSet),
            TableUpdateForPastTimestamp()
        );

        // Verify the operator table update
        _verifyOperatorTableUpdate({
            referenceTimestamp: referenceTimestamp,
            globalTableRoot: globalTableRoot,
            operatorSetIndex: operatorSetIndex,
            proof: proof,
            operatorSetLeafHash: keccak256(operatorTableBytes)
        });

        // Update the operator table
        if (curveType == CurveType.BN254) {
            bn254CertificateVerifier.updateOperatorTable(
                operatorSet, referenceTimestamp, _getBN254OperatorInfo(operatorTableInfo), operatorSetConfig
            );
        } else if (curveType == CurveType.ECDSA) {
            ecdsaCertificateVerifier.updateOperatorTable(
                operatorSet, referenceTimestamp, _getECDSAOperatorInfo(operatorTableInfo), operatorSetConfig
            );
        } else {
            revert InvalidCurveType();
        }
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
        return _globalTableRoots[_latestReferenceTimestamp];
    }

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalRootConfirmerSet() external view returns (OperatorSet memory) {
        return _globalRootConfirmerSet;
    }

    /// @inheritdoc IOperatorTableUpdater
    function getCertificateVerifier(
        CurveType curveType
    ) public view returns (address) {
        if (curveType == CurveType.BN254) {
            return address(bn254CertificateVerifier);
        } else if (curveType == CurveType.ECDSA) {
            return address(ecdsaCertificateVerifier);
        } else {
            revert InvalidCurveType();
        }
    }

    /// @inheritdoc IOperatorTableUpdater
    function getLatestReferenceTimestamp() external view returns (uint32) {
        return _latestReferenceTimestamp;
    }

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalTableUpdateMessageHash(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(GLOBAL_TABLE_ROOT_CERT_TYPEHASH, globalTableRoot, referenceTimestamp));
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

    /**
     * @notice Gets the operator table info from a bytes array
     * @param operatorTable The bytes containing the operator table
     * @return operatorSet The operator set
     * @return curveType The curve type
     * @return operatorSetInfo The operator set info
     * @return operatorTableInfo The operator table info. This is encoded as a bytes array, and its value is dependent on the curve type, see `_getBN254OperatorInfo` and `_getECDSAOperatorInfo`
     */
    function _decodeOperatorTableBytes(
        bytes calldata operatorTable
    )
        internal
        pure
        returns (
            OperatorSet memory operatorSet,
            CurveType curveType,
            OperatorSetConfig memory operatorSetInfo,
            bytes memory operatorTableInfo
        )
    {
        (operatorSet, curveType, operatorSetInfo, operatorTableInfo) =
            abi.decode(operatorTable, (OperatorSet, CurveType, OperatorSetConfig, bytes));
    }

    /**
     * @notice Gets the BN254 operator set info from a bytes array
     * @param BN254OperatorSetInfoBytes The bytes containing the operator set info
     * @return operatorSetInfo The BN254 operator set info
     */
    function _getBN254OperatorInfo(
        bytes memory BN254OperatorSetInfoBytes
    ) internal pure returns (BN254OperatorSetInfo memory) {
        return abi.decode(BN254OperatorSetInfoBytes, (BN254OperatorSetInfo));
    }

    /**
     * @notice Gets the ECDSA operator set info from a bytes array
     * @param ECDSAOperatorInfoBytes The bytes containing the operator table info
     * @return operatorSetInfo The ECDSA operator set info
     */
    function _getECDSAOperatorInfo(
        bytes memory ECDSAOperatorInfoBytes
    ) internal pure returns (ECDSAOperatorInfo[] memory) {
        return abi.decode(ECDSAOperatorInfoBytes, (ECDSAOperatorInfo[]));
    }
}
