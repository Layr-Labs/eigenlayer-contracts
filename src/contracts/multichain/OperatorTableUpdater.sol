// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../libraries/Merkle.sol";
import "../permissions/Pausable.sol";
import "../mixins/SemVerMixin.sol";
import "../mixins/LeafCalculatorMixin.sol";
import "./OperatorTableUpdaterStorage.sol";

contract OperatorTableUpdater is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    OperatorTableUpdaterStorage,
    SemVerMixin,
    LeafCalculatorMixin,
    ReentrancyGuardUpgradeable
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
     * @param _owner The owner of the OperatorTableUpdater
     * @param initialPausedStatus The initial paused status of the OperatorTableUpdater
     * @param _initialGenerator The operatorSet which certifies against global roots
     * @param _globalRootConfirmationThreshold The threshold, in bps, for a global root to be signed off on and updated
     * @param generatorInfo The operatorSetInfo for the Generator
     * @dev We also update the operator table for the Generator, to begin signing off on global roots
     * @dev We set the `_latestReferenceTimestamp` to the current timestamp, so that only *new* roots can be confirmed
     */
    function initialize(
        address _owner,
        uint256 initialPausedStatus,
        OperatorSet calldata _initialGenerator,
        uint16 _globalRootConfirmationThreshold,
        BN254OperatorSetInfo calldata generatorInfo
    ) external initializer {
        _transferOwnership(_owner);
        _setPausedStatus(initialPausedStatus);

        // Set the `operatorSetConfig` for the `Generator`
        _generatorConfig.maxStalenessPeriod = GENERATOR_MAX_STALENESS_PERIOD;
        _generatorConfig.owner = address(this);

        _updateGenerator(_initialGenerator, generatorInfo);
        _setGlobalRootConfirmationThreshold(_globalRootConfirmationThreshold);

        // The generator's global table root is the `GENERATOR_GLOBAL_TABLE_ROOT`.
        // The constant is used to enable the call to `confirmGlobalTableRoot` to pass since the `BN254CertificateVerifier` expects the `Generator` to have a valid root associated with it.
        _globalTableRoots[GENERATOR_REFERENCE_TIMESTAMP] = GENERATOR_GLOBAL_TABLE_ROOT;
        _isRootValid[GENERATOR_GLOBAL_TABLE_ROOT] = true;

        // Set the `latestReferenceTimestamp` so that only *new* roots can be confirmed
        _latestReferenceTimestamp = uint32(block.timestamp);
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
    ) external onlyWhenNotPaused(PAUSED_GLOBAL_ROOT_UPDATE) nonReentrant {
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
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_UPDATE) nonReentrant {
        (
            OperatorSet memory operatorSet,
            CurveType curveType,
            OperatorSetConfig memory operatorSetConfig,
            bytes memory operatorTableInfo
        ) = _decodeOperatorTableBytes(operatorTableBytes);

        // Check that the `globalTableRoot` is not disabled
        require(_isRootValid[globalTableRoot], InvalidRoot());

        // Check that the `operatorSet` is not the `Generator`
        require(operatorSet.key() != _generator.key(), InvalidOperatorSet());

        // Silently return if the `referenceTimestamp` has already been updated for the `operatorSet`
        // We do this to avoid race conditions with the offchain transport of the operator table
        if (
            IBaseCertificateVerifier(getCertificateVerifier(curveType)).isReferenceTimestampSet(
                operatorSet, referenceTimestamp
            )
        ) {
            return;
        }

        // Check that the `referenceTimestamp` is greater than the latest reference timestamp
        require(
            referenceTimestamp
                > IBaseCertificateVerifier(getCertificateVerifier(curveType)).latestReferenceTimestamp(operatorSet),
            TableUpdateForPastTimestamp()
        );

        // Check that the `globalTableRoot` matches the `referenceTimestamp`
        require(_globalTableRoots[referenceTimestamp] == globalTableRoot, InvalidGlobalTableRoot());

        // Verify the operator table update
        _verifyMerkleInclusion({
            globalTableRoot: globalTableRoot,
            operatorSetIndex: operatorSetIndex,
            proof: proof,
            operatorSetLeafHash: calculateOperatorTableLeaf(operatorTableBytes)
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

        // Check that the root is not the generator's global table root
        require(globalTableRoot != GENERATOR_GLOBAL_TABLE_ROOT, CannotDisableGeneratorRoot());

        _isRootValid[globalTableRoot] = false;
        emit GlobalRootDisabled(globalTableRoot);
    }

    /// @inheritdoc IOperatorTableUpdater
    function updateGenerator(
        OperatorSet calldata generator,
        BN254OperatorSetInfo calldata generatorInfo
    ) external onlyOwner {
        _updateGenerator(generator, generatorInfo);
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
    function getGeneratorConfig() external view returns (OperatorSetConfig memory) {
        return _generatorConfig;
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

    /// @inheritdoc IOperatorTableUpdater
    function getGlobalTableUpdateSignableDigest(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        uint32 referenceBlockNumber
    ) public view returns (bytes32) {
        bytes32 messageHash = getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber);
        return bn254CertificateVerifier.calculateCertificateDigest(GENERATOR_REFERENCE_TIMESTAMP, messageHash);
    }

    /**
     *
     *                         INTERNAL HELPERS
     *
     */

    /**
     * @notice Verifies that the operator table update is valid by checking the `proof` against a `globalTableRoot`
     * @param globalTableRoot The global table root of the operator table update
     * @param operatorSetIndex The index of the operator set in the operator table
     * @param proof The proof of the operator table update
     * @param operatorSetLeafHash The leaf hash of the operator set
     * @dev Reverts if there does not exist a `globalTableRoot` for the given `referenceTimestamp`
     */
    function _verifyMerkleInclusion(
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        bytes32 operatorSetLeafHash
    ) internal pure {
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
     * @notice Updates the `Generator` to a new operatorSet
     * @param generator The operatorSet which certifies against global roots
     * @param generatorInfo The operatorSetInfo for the generator
     * @dev We have a separate function for updating this operatorSet since it's not transported and updated
     *      in the same way as the other operatorSets
     * @dev Only callable by the owner of the contract
     * @dev Uses GENERATOR_GLOBAL_TABLE_ROOT constant to break circular dependency for certificate verification
     * @dev We ensure that there are no collisions with other reference timestamps because we expect the generator to have an initial reference timestamp of 0
     * @dev The `_latestReferenceTimestamp` is not updated since this root is ONLY used for the `Generator`
     * @dev The `_referenceBlockNumber` and `_referenceTimestamps` mappings are not updated since they are only used for introspection for official operatorSets
     */
    function _updateGenerator(OperatorSet calldata generator, BN254OperatorSetInfo calldata generatorInfo) internal {
        // Set the generator
        _generator = generator;

        // Get the latest reference timestamp for the Generator
        uint32 referenceTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(generator);
        require(referenceTimestamp == 0, InvalidGenerator());

        // Update the operator table for the Generator
        bn254CertificateVerifier.updateOperatorTable(
            generator, GENERATOR_REFERENCE_TIMESTAMP, generatorInfo, _generatorConfig
        );

        emit GeneratorUpdated(generator);
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
