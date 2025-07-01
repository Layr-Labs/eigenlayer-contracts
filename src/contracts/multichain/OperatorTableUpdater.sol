// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../libraries/Merkle.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "./OperatorTableUpdaterStorage.sol";

contract OperatorTableUpdater is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    OperatorTableUpdaterStorage,
    SemVerMixin
{
    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor(
        IBN254CertificateVerifier _bn254CertificateVerifier,
        IECDSACertificateVerifier _ecdsaCertificateVerifier,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        OperatorTableUpdaterStorage(_bn254CertificateVerifier, _ecdsaCertificateVerifier)
        Pausable(_pauserRegistry)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /**
     * @notice Initializes the OperatorTableUpdater
     * @param owner The owner of the OperatorTableUpdater
     * @param initialPausedStatus The initial paused status of the OperatorTableUpdater
     * @param _generator The operatorSet which certifies against global roots
     * @param _globalRootConfirmationThreshold The threshold, in bps, for a global root to be signed off on and updated
     * @param referenceTimestamp The reference timestamp for the global root confirmer set
     * @param generatorInfo The operatorSetInfo for the global root confirmer set
     * @param generatorConfig The operatorSetConfig for the global root confirmer set
     * @dev We also update the operator table for the global root confirmer set, to begin signing off on global roots
     * @dev Uses INITIAL_GLOBAL_TABLE_ROOT constant to break circular dependency for certificate verification
     */
    function initialize(
        address owner,
        uint256 initialPausedStatus,
        OperatorSet calldata _generator,
        uint16 _globalRootConfirmationThreshold,
        uint32 referenceTimestamp,
        BN254OperatorSetInfo calldata generatorInfo,
        OperatorSetConfig calldata generatorConfig
    ) external initializer {
        _transferOwnership(owner);
        _setPausedStatus(initialPausedStatus);
        _setGenerator(_generator);
        _setGlobalRootConfirmationThreshold(_globalRootConfirmationThreshold);
        _updateGenerator(referenceTimestamp, generatorInfo, generatorConfig);

        /// @dev The first global table root is the `INITIAL_GLOBAL_TABLE_ROOT`
        /// @dev This is used to enable the first call to `confirmGlobalTableRoot` to pass since it expects
        /// @dev the `Generator` to have a valid initial global table root
        _globalTableRoots[referenceTimestamp] = INITIAL_GLOBAL_TABLE_ROOT;
        _isRootValid[INITIAL_GLOBAL_TABLE_ROOT] = true;
        _referenceBlockNumbers[referenceTimestamp] = uint32(block.number);
        _referenceTimestamps[uint32(block.number)] = referenceTimestamp;

        // Set the latest reference timestamp
        _latestReferenceTimestamp = referenceTimestamp;

        // Emit the initial global table root event
        emit NewGlobalTableRoot(referenceTimestamp, INITIAL_GLOBAL_TABLE_ROOT);
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
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) external onlyWhenNotPaused(PAUSED_GLOBAL_ROOT_UPDATE) {
        // Table roots can only be updated for current or past timestamps and after the latest reference timestamp
        require(referenceTimestamp <= block.timestamp, GlobalTableRootInFuture());
        require(referenceTimestamp > _latestReferenceTimestamp, GlobalTableRootStale());
        require(
            globalTableRootCert.messageHash
                == getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber),
            InvalidMessageHash()
        );

        // Verify certificate by using the stake proportion thresholds
        uint16[] memory stakeProportionThresholds = new uint16[](1);
        stakeProportionThresholds[0] = globalRootConfirmationThreshold;
        bool isValid = bn254CertificateVerifier.verifyCertificateProportion(
            _generator, globalTableRootCert, stakeProportionThresholds
        );

        require(isValid, CertificateInvalid());

        // Update the global table root & reference timestamps
        _latestReferenceTimestamp = referenceTimestamp;
        _referenceBlockNumbers[referenceTimestamp] = referenceBlockNumber;
        _referenceTimestamps[referenceBlockNumber] = referenceTimestamp;
        _globalTableRoots[referenceTimestamp] = globalTableRoot;
        _isRootValid[globalTableRoot] = true;

        emit NewGlobalTableRoot(referenceTimestamp, globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes calldata operatorTableBytes
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_UPDATE) {
        (
            OperatorSet memory operatorSet,
            CurveType curveType,
            OperatorSetConfig memory operatorSetConfig,
            bytes memory operatorTableInfo
        ) = _decodeOperatorTableBytes(operatorTableBytes);

        // Check that the `globalTableRoot` is not disabled
        require(_isRootValid[globalTableRoot], InvalidRoot());

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
    function setGenerator(
        OperatorSet calldata operatorSet
    ) external onlyOwner {
        _setGenerator(operatorSet);
    }

    /// @inheritdoc IOperatorTableUpdater
    function setGlobalRootConfirmationThreshold(
        uint16 bps
    ) external onlyOwner {
        _setGlobalRootConfirmationThreshold(bps);
    }

    /// @inheritdoc IOperatorTableUpdater
    function disableRoot(
        bytes32 globalTableRoot
    ) external onlyPauser {
        // Check that the root already exists and is not disabled
        require(_isRootValid[globalTableRoot], InvalidRoot());

        _isRootValid[globalTableRoot] = false;
        emit GlobalRootDisabled(globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateGenerator(
        uint32 referenceTimestamp,
        BN254OperatorSetInfo calldata generatorInfo,
        OperatorSetConfig calldata generatorConfig
    ) external onlyOwner {
        _updateGenerator(referenceTimestamp, generatorInfo, generatorConfig);
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
    function getGenerator() external view returns (OperatorSet memory) {
        return _generator;
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
    function getLatestReferenceBlockNumber() external view returns (uint32) {
        return _referenceBlockNumbers[_latestReferenceTimestamp];
    }

    /// @inheritdoc IOperatorTableUpdater
    function getReferenceBlockNumberByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (uint32) {
        return _referenceBlockNumbers[referenceTimestamp];
    }

    /// @inheritdoc IOperatorTableUpdater
    function getReferenceTimestampByBlockNumber(
        uint32 referenceBlockNumber
    ) external view returns (uint32) {
        return _referenceTimestamps[referenceBlockNumber];
    }

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalTableUpdateMessageHash(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) public pure returns (bytes32) {
        return keccak256(
            abi.encode(GLOBAL_TABLE_ROOT_CERT_TYPEHASH, globalTableRoot, referenceTimestamp, referenceBlockNumber)
        );
    }

    /// @inheritdoc IOperatorTableUpdater
    function getGeneratorReferenceTimestamp() external view returns (uint32) {
        return IBaseCertificateVerifier(address(bn254CertificateVerifier)).latestReferenceTimestamp(_generator);
    }

    /// @inheritdoc IOperatorTableUpdater
    function isRootValid(
        bytes32 globalTableRoot
    ) public view returns (bool) {
        return _isRootValid[globalTableRoot];
    }

    /// @inheritdoc IOperatorTableUpdater
    function isRootValidByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (bool) {
        return _isRootValid[_globalTableRoots[referenceTimestamp]];
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
    function _setGenerator(
        OperatorSet calldata operatorSet
    ) internal {
        _generator = operatorSet;
        emit GeneratorUpdated(operatorSet);
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
     * @notice Updates the operator table for the global root confirmer set
     * @param referenceTimestamp The reference timestamp of the operator table update
     * @param generatorInfo The operatorSetInfo for the global root confirmer set
     * @param generatorConfig The operatorSetConfig for the global root confirmer set
     */
    function _updateGenerator(
        uint32 referenceTimestamp,
        BN254OperatorSetInfo calldata generatorInfo,
        OperatorSetConfig calldata generatorConfig
    ) internal {
        bn254CertificateVerifier.updateOperatorTable(_generator, referenceTimestamp, generatorInfo, generatorConfig);
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
