// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";

contract OperatorTableUpdaterUnitTests is
    EigenLayerMultichainUnitTestSetup,
    IOperatorTableUpdaterErrors,
    IOperatorTableUpdaterEvents,
    IBN254CertificateVerifierTypes,
    IECDSACertificateVerifierTypes,
    ICrossChainRegistryTypes,
    IKeyRegistrarTypes
{
    using StdStyle for *;
    using ArrayLib for *;

    /// @notice Pointers to the operatorTableUpdater and its implementation
    OperatorTableUpdater operatorTableUpdater;
    OperatorTableUpdater operatorTableUpdaterImplementation;

    /// @notice The default operatorSet
    OperatorSet defaultOperatorSet = OperatorSet({avs: address(this), id: 0});

    /// @notice The global confirmer operatorSet params
    OperatorSet globalRootConfirmerSet = OperatorSet({avs: address(0xDEADBEEF), id: 0});
    uint16 internal constant GLOBAL_ROOT_CONFIRMATION_THRESHOLD = 10_000; // 100% signoff threshold

    function setUp() public virtual override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Setup a mock Bn254OperatorSetInfo for the initial table update on initialization
        BN254OperatorSetInfo memory initialOperatorSetInfo = BN254OperatorSetInfo({
            operatorInfoTreeRoot: bytes32(0),
            numOperators: 1,
            aggregatePubkey: BN254.G1Point({X: 1, Y: 2}),
            totalWeights: new uint[](1)
        });
        OperatorSetConfig memory initialOperatorSetConfig = OperatorSetConfig({owner: address(0xDEADBEEF), maxStalenessPeriod: 10_000});

        operatorTableUpdater =
            OperatorTableUpdater(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));

        // Deploy operatorTableUpdater
        operatorTableUpdaterImplementation = new OperatorTableUpdater(
            IBN254CertificateVerifier(address(bn254CertificateVerifierMock)),
            IECDSACertificateVerifier(address(ecdsaCertificateVerifierMock)),
            "1.0.0"
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(operatorTableUpdater))),
            address(operatorTableUpdaterImplementation),
            abi.encodeWithSelector(
                OperatorTableUpdater.initialize.selector,
                address(this), // owner
                globalRootConfirmerSet, // globalRootConfirmerSet
                GLOBAL_ROOT_CONFIRMATION_THRESHOLD, // globalRootConfirmationThreshold
                block.timestamp - 1, // referenceTimestamp
                initialOperatorSetInfo, // globalRootConfirmerSetInfo
                initialOperatorSetConfig // globalRootConfirmerSetConfig
            )
        );
    }

    /// @dev Sets a valid certificate for the BN254 certificate verifier
    function _setIsValidCertificate(BN254Certificate memory certificate, bool isValid) internal {
        bn254CertificateVerifierMock.setIsValidCertificate(certificate, isValid);
    }

    /// @dev Sets a valid certificate for the ECDSA certificate verifier
    function _setIsValidCertificate(ECDSACertificate memory certificate, bool isValid) internal {
        ecdsaCertificateVerifierMock.setIsValidCertificate(certificate, isValid);
    }

    /// @dev Generates a random BN254 operator set info
    function _generateRandomBN254OperatorSetInfo(Randomness r) internal returns (BN254OperatorSetInfo memory) {
        BN254OperatorSetInfo memory operatorSetInfo;
        uint maxWeightLength = r.Uint256(1, 16);
        operatorSetInfo.operatorInfoTreeRoot = bytes32(r.Uint256());
        operatorSetInfo.numOperators = r.Uint256();
        operatorSetInfo.aggregatePubkey = BN254.G1Point({X: r.Uint256(), Y: r.Uint256()});
        operatorSetInfo.totalWeights = r.Uint256Array({len: maxWeightLength, min: 2, max: 10_000 ether});
        return operatorSetInfo;
    }

    /// @dev Generates a random ECDSA operator set info
    function _generateRandomECDSAOperatorInfos(Randomness r) internal returns (ECDSAOperatorInfo[] memory) {
        uint numOperators = r.Uint256(1, 50);
        uint maxWeightLength = r.Uint256(1, 16);
        ECDSAOperatorInfo[] memory operatorInfos = new ECDSAOperatorInfo[](numOperators);

        for (uint i = 0; i < numOperators; i++) {
            operatorInfos[i].pubkey = r.Address();
            operatorInfos[i].weights = r.Uint256Array({len: maxWeightLength, min: 2, max: 10_000 ether});
        }

        return operatorInfos;
    }

    /// @dev Generates a random operator set config
    function _generateRandomOperatorSetConfig(Randomness r) internal returns (OperatorSetConfig memory) {
        OperatorSetConfig memory operatorSetConfig;
        operatorSetConfig.owner = r.Address();
        operatorSetConfig.maxStalenessPeriod = r.Uint32(0, type(uint32).max);
        return operatorSetConfig;
    }

    /// @dev Creates a global table root with a single operatorSetLeafHash
    function _createGlobalTableRoot(Randomness r, bytes32 operatorSetLeafHash) internal returns (bytes32, uint32, bytes32[] memory) {
        // Generate a random power of 2 between 2 and 2^16
        uint exponent = r.Uint256(1, 16);
        uint numLeaves = 2 ** exponent;

        // Create leaves array with the specified size
        bytes32[] memory leaves = new bytes32[](numLeaves);

        // Set a random index to be the operatorSetLeafHash
        uint32 randomIndex = uint32(r.Uint256(0, numLeaves - 1));

        // Fill remaining leaves with random data
        for (uint i = 0; i < numLeaves; i++) {
            if (i == randomIndex) leaves[i] = operatorSetLeafHash;
            else leaves[i] = bytes32(r.Uint256());
        }

        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        return (globalTableRoot, randomIndex, leaves);
    }

    /// @dev Creates a global table root with multiple operatorSetLeafHashes
    function _createGlobalTableRoot(Randomness r, bytes32[] memory operatorSetLeafHashes)
        internal
        returns (bytes32, uint32[] memory, bytes32[] memory)
    {
        // Generate a random power of 2 between the number of operatorSetLeafHashes and 2^16
        uint exponent = r.Uint256(operatorSetLeafHashes.length, 16);
        uint numLeaves = 2 ** exponent;

        // Create leaves array with the specified size
        bytes32[] memory leaves = new bytes32[](numLeaves);

        uint32[] memory operatorSetIndices = new uint32[](operatorSetLeafHashes.length);

        // Set the first n indices to be the operatorSetLeafHashes
        uint32 randomIndex = uint32(r.Uint256(0, numLeaves - 1));

        // Fill the leaves
        for (uint i = 0; i < operatorSetLeafHashes.length; i++) {
            leaves[i] = operatorSetLeafHashes[i];
            operatorSetIndices[i] = uint32(i);
        }

        // Fill the remaining leaves with random data
        for (uint i = operatorSetLeafHashes.length; i < numLeaves; i++) {
            leaves[i] = bytes32(r.Uint256());
        }

        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        return (globalTableRoot, operatorSetIndices, leaves);
    }

    /// @dev Updates the global table root in the BN254 certificate verifier
    function _updateGlobalTableRoot(bytes32 globalTableRoot) internal {
        BN254Certificate memory mockCertificate;
        uint32 referenceBlockNumber = uint32(block.number);
        mockCertificate.messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, uint32(block.timestamp), referenceBlockNumber);
        _setIsValidCertificate(mockCertificate, true);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, globalTableRoot, uint32(block.timestamp), referenceBlockNumber);
    }
}

contract OperatorTableUpdaterUnitTests_initialize is OperatorTableUpdaterUnitTests {
    function test_initialize_success() public view {
        OperatorSet memory confirmerSet = operatorTableUpdater.getGlobalRootConfirmerSet();
        assertEq(confirmerSet.avs, address(0xDEADBEEF));
        assertEq(confirmerSet.id, 0);
        assertEq(operatorTableUpdater.getLatestReferenceTimestamp(), uint32(block.timestamp - 1));
        assertEq(operatorTableUpdater.getGlobalConfirmerSetReferenceTimestamp(), uint32(block.timestamp - 1));

        uint16 threshold = operatorTableUpdater.globalRootConfirmationThreshold();
        assertEq(threshold, GLOBAL_ROOT_CONFIRMATION_THRESHOLD);
    }

    function test_initialize_revert_reinitialization() public {
        BN254OperatorSetInfo memory initialOperatorSetInfo;
        OperatorSetConfig memory initialOperatorSetConfig;
        cheats.expectRevert("Initializable: contract is already initialized");
        operatorTableUpdater.initialize(
            address(this),
            globalRootConfirmerSet,
            GLOBAL_ROOT_CONFIRMATION_THRESHOLD,
            uint32(block.timestamp - 1),
            initialOperatorSetInfo,
            initialOperatorSetConfig
        );
    }
}

contract OperatorTableUpdaterUnitTests_confirmGlobalTableRoot is OperatorTableUpdaterUnitTests {
    BN254Certificate mockCertificate;

    function testFuzz_revert_futureTableRoot(Randomness r) public rand(r) {
        uint32 referenceTimestamp = r.Uint32(uint32(block.timestamp + 1), type(uint32).max);
        uint32 referenceBlockNumber = r.Uint32();

        cheats.expectRevert(GlobalTableRootInFuture.selector);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, bytes32(0), referenceTimestamp + 1, referenceBlockNumber);
    }

    function testFuzz_revert_staleCertificate(Randomness r) public rand(r) {
        uint32 referenceBlockNumber = uint32(block.number);
        mockCertificate.messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(bytes32(0), uint32(block.timestamp), referenceBlockNumber);
        _setIsValidCertificate(mockCertificate, true);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, bytes32(0), uint32(block.timestamp), referenceBlockNumber);

        uint32 referenceTimestamp = r.Uint32(0, operatorTableUpdater.getLatestReferenceTimestamp() - 1);
        cheats.expectRevert(GlobalTableRootStale.selector);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, bytes32(0), referenceTimestamp, referenceBlockNumber);
    }

    function testFuzz_revert_InvalidMessageHash(Randomness r) public rand(r) {
        bytes32 invalidTableRoot = bytes32(r.Uint256(1, type(uint).max));
        mockCertificate.messageHash = invalidTableRoot;
        uint32 referenceBlockNumber = r.Uint32();

        cheats.expectRevert(InvalidMessageHash.selector);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, bytes32(0), uint32(block.timestamp), referenceBlockNumber);
    }

    function test_revert_invalidCertificate() public {
        uint32 referenceBlockNumber = uint32(block.number);
        mockCertificate.messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(bytes32(0), uint32(block.timestamp), referenceBlockNumber);
        _setIsValidCertificate(mockCertificate, false);
        cheats.expectRevert(CertificateInvalid.selector);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, bytes32(0), uint32(block.timestamp), referenceBlockNumber);
    }

    function testFuzz_correctness(Randomness r) public rand(r) {
        uint32 referenceTimestamp = r.Uint32(operatorTableUpdater.getLatestReferenceTimestamp() + 1, type(uint32).max);
        uint32 referenceBlockNumber = r.Uint32();
        cheats.warp(uint(referenceTimestamp));
        bytes32 globalTableRoot = bytes32(r.Uint256(1, type(uint).max));
        mockCertificate.messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber);
        _setIsValidCertificate(mockCertificate, true);

        // Expect the global table root to be updated
        cheats.expectEmit(true, true, true, true);
        emit NewGlobalTableRoot(referenceTimestamp, globalTableRoot);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, globalTableRoot, referenceTimestamp, referenceBlockNumber);

        // Expect the global table root to be updated
        assertEq(operatorTableUpdater.getGlobalTableRootByTimestamp(referenceTimestamp), globalTableRoot);
        assertEq(operatorTableUpdater.getCurrentGlobalTableRoot(), globalTableRoot);
        assertEq(operatorTableUpdater.getLatestReferenceTimestamp(), referenceTimestamp);
        assertEq(operatorTableUpdater.getLatestReferenceBlockNumber(), referenceBlockNumber);
        assertEq(operatorTableUpdater.getReferenceBlockNumberByTimestamp(referenceTimestamp), referenceBlockNumber);
        assertEq(operatorTableUpdater.getReferenceTimestampByBlockNumber(referenceBlockNumber), referenceTimestamp);
    }
}

contract OperatorTableUpdaterUnitTests_updateOperatorTable_BN254 is OperatorTableUpdaterUnitTests {
    function _setLatestReferenceTimestampBN254(OperatorSet memory operatorSet, uint32 referenceTimestamp) internal {
        bn254CertificateVerifierMock.setLatestReferenceTimestamp(operatorSet, referenceTimestamp);
    }

    function testFuzz_BN254_revert_staleTableUpdate(Randomness r) public rand(r) {
        uint32 referenceTimestamp = r.Uint32(uint32(block.timestamp), type(uint32).max);
        _setLatestReferenceTimestampBN254(defaultOperatorSet, referenceTimestamp);
        BN254OperatorSetInfo memory emptyOperatorSetInfo;

        bytes memory operatorTable =
            abi.encode(defaultOperatorSet, CurveType.BN254, _generateRandomOperatorSetConfig(r), emptyOperatorSetInfo);

        cheats.expectRevert(TableUpdateForPastTimestamp.selector);
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), bytes32(0), 0, new bytes(0), operatorTable);
    }

    function testFuzz_BN254_revert_invalidGlobalTableRoot(Randomness r) public rand(r) {
        // Generate random operatorSetInfo and operatorSetConfig
        BN254OperatorSetInfo memory operatorSetInfo = _generateRandomBN254OperatorSetInfo(r);
        bytes memory operatorSetInfoBytes = abi.encode(operatorSetInfo);
        OperatorSetConfig memory operatorSetConfig = _generateRandomOperatorSetConfig(r);
        bytes memory operatorTable = abi.encode(defaultOperatorSet, CurveType.BN254, operatorSetConfig, operatorSetInfoBytes);
        bytes32 operatorSetLeafHash = keccak256(operatorTable);

        // Include the operatorSetInfo and operatorSetConfig in the global table root & set it
        (bytes32 globalTableRoot, uint32 operatorSetIndex, bytes32[] memory leaves) = _createGlobalTableRoot(r, operatorSetLeafHash);
        _updateGlobalTableRoot(globalTableRoot);

        // Generate proof for the operatorSetInfo and operatorSetConfig
        bytes memory proof = Merkle.getProofKeccak(leaves, operatorSetIndex);

        // Try to update with a different global table root
        bytes32 wrongGlobalTableRoot = bytes32(r.Uint256(1, type(uint).max));
        cheats.expectRevert(InvalidGlobalTableRoot.selector);
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), wrongGlobalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    function testFuzz_BN254_revert_invalidOperatorSetProof(Randomness r) public rand(r) {
        // Generate random operatorSetInfo and operatorSetConfig
        BN254OperatorSetInfo memory operatorSetInfo = _generateRandomBN254OperatorSetInfo(r);
        bytes memory operatorSetInfoBytes = abi.encode(operatorSetInfo);
        OperatorSetConfig memory operatorSetConfig = _generateRandomOperatorSetConfig(r);
        bytes memory operatorTable = abi.encode(defaultOperatorSet, CurveType.BN254, operatorSetConfig, operatorSetInfoBytes);
        bytes32 operatorSetLeafHash = keccak256(operatorTable);

        // Include the operatorSetInfo and operatorSetConfig in the global table root & set it
        (bytes32 globalTableRoot, uint32 operatorSetIndex, bytes32[] memory leaves) = _createGlobalTableRoot(r, operatorSetLeafHash);
        _updateGlobalTableRoot(globalTableRoot);

        // Generate an invalid proof
        bytes memory invalidProof = new bytes(32 * r.Uint256(1, 10));
        for (uint i = 0; i < invalidProof.length; i++) {
            invalidProof[i] = bytes1(uint8(r.Uint256(0, 255)));
        }

        // Try to update with invalid proof
        cheats.expectRevert(InvalidOperatorSetProof.selector);
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), globalTableRoot, operatorSetIndex, invalidProof, operatorTable);
    }

    function testFuzz_BN254_correctness(Randomness r) public rand(r) {
        // Generate random operatorSetInfo and operatorSetConfig
        BN254OperatorSetInfo memory operatorSetInfo = _generateRandomBN254OperatorSetInfo(r);
        // Encode the operatorSetInfo as bytes, as the CrossChainRegistry expects the operatorTable to be encoded as a bytes array;
        bytes memory operatorSetInfoBytes = abi.encode(operatorSetInfo);
        OperatorSetConfig memory operatorSetConfig = _generateRandomOperatorSetConfig(r);
        bytes memory operatorTable = abi.encode(defaultOperatorSet, CurveType.BN254, operatorSetConfig, operatorSetInfoBytes);
        bytes32 operatorSetLeafHash = keccak256(operatorTable);

        // Include the operatorSetInfo and operatorSetConfig in the global table root & set it
        (bytes32 globalTableRoot, uint32 operatorSetIndex, bytes32[] memory leaves) = _createGlobalTableRoot(r, operatorSetLeafHash);
        _updateGlobalTableRoot(globalTableRoot);

        // Generate proof for the operatorSetInfo and operatorSetConfig
        bytes memory proof = Merkle.getProofKeccak(leaves, operatorSetIndex);

        // Update the operator table
        cheats.expectCall(
            address(bn254CertificateVerifierMock),
            abi.encodeWithSelector(
                IBN254CertificateVerifier.updateOperatorTable.selector,
                defaultOperatorSet,
                uint32(block.timestamp),
                operatorSetInfo,
                operatorSetConfig
            )
        );
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), globalTableRoot, operatorSetIndex, proof, operatorTable);
    }
}

contract OperatorTableUpdaterUnitTests_updateOperatorTable_ECDSA is OperatorTableUpdaterUnitTests {
    function _setLatestReferenceTimestampECDSA(OperatorSet memory operatorSet, uint32 referenceTimestamp) internal {
        ecdsaCertificateVerifierMock.setLatestReferenceTimestamp(operatorSet, referenceTimestamp);
    }

    function testFuzz_revert_ECDSA_staleTableUpdate(Randomness r) public rand(r) {
        uint32 referenceTimestamp = r.Uint32(uint32(block.timestamp), type(uint32).max);
        _setLatestReferenceTimestampECDSA(defaultOperatorSet, referenceTimestamp);
        ECDSAOperatorInfo[] memory emptyOperatorSetInfo;

        bytes memory operatorTable =
            abi.encode(defaultOperatorSet, CurveType.ECDSA, _generateRandomOperatorSetConfig(r), emptyOperatorSetInfo);

        cheats.expectRevert(TableUpdateForPastTimestamp.selector);
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), bytes32(0), 0, new bytes(0), operatorTable);
    }

    function testFuzz_ECDSA_correctness(Randomness r) public rand(r) {
        // Generate random operatorSetInfo and operatorSetConfig
        ECDSAOperatorInfo[] memory operatorInfos = _generateRandomECDSAOperatorInfos(r);
        // Encode the operatorInfos as bytes, as the CrossChainRegistry expects the operatorTable to be encoded as a bytes array;
        bytes memory operatorInfosBytes = abi.encode(operatorInfos);
        OperatorSetConfig memory operatorSetConfig = _generateRandomOperatorSetConfig(r);
        bytes memory operatorTable = abi.encode(defaultOperatorSet, CurveType.ECDSA, operatorSetConfig, operatorInfosBytes);
        bytes32 operatorSetLeafHash = keccak256(operatorTable);

        // Include the operatorInfos and operatorSetConfig in the global table root & set it
        (bytes32 globalTableRoot, uint32 operatorSetIndex, bytes32[] memory leaves) = _createGlobalTableRoot(r, operatorSetLeafHash);
        _updateGlobalTableRoot(globalTableRoot);

        // Generate proof for the operatorInfos and operatorSetConfig
        bytes memory proof = Merkle.getProofKeccak(leaves, operatorSetIndex);

        // Update the operator table
        cheats.expectCall(
            address(ecdsaCertificateVerifierMock),
            abi.encodeWithSelector(
                IECDSACertificateVerifier.updateOperatorTable.selector,
                defaultOperatorSet,
                uint32(block.timestamp),
                operatorInfos,
                operatorSetConfig
            )
        );
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), globalTableRoot, operatorSetIndex, proof, operatorTable);
    }
}

contract OperatorTableUpdaterUnitTests_updateOperatorTable_InvalidCurveType is OperatorTableUpdaterUnitTests {
    function testFuzz_revert_invalidCurveType_inGetCertificateVerifier(Randomness r) public rand(r) {
        // Generate random operatorSetInfo and operatorSetConfig with invalid curve type
        BN254OperatorSetInfo memory operatorSetInfo = _generateRandomBN254OperatorSetInfo(r);
        bytes memory operatorSetInfoBytes = abi.encode(operatorSetInfo);
        OperatorSetConfig memory operatorSetConfig = _generateRandomOperatorSetConfig(r);
        // Use CurveType.NONE (0) which is invalid
        bytes memory operatorTable = abi.encode(defaultOperatorSet, CurveType.NONE, operatorSetConfig, operatorSetInfoBytes);
        bytes32 operatorSetLeafHash = keccak256(operatorTable);

        // Include the operatorSetInfo and operatorSetConfig in the global table root & set it
        (bytes32 globalTableRoot, uint32 operatorSetIndex, bytes32[] memory leaves) = _createGlobalTableRoot(r, operatorSetLeafHash);
        _updateGlobalTableRoot(globalTableRoot);

        // Generate proof for the operatorSetInfo and operatorSetConfig
        bytes memory proof = Merkle.getProofKeccak(leaves, operatorSetIndex);

        // Should revert with InvalidCurveType when trying to update with CurveType.NONE
        // This will revert in getCertificateVerifier, not in the else branch
        cheats.expectRevert(InvalidCurveType.selector);
        operatorTableUpdater.updateOperatorTable(uint32(block.timestamp), globalTableRoot, operatorSetIndex, proof, operatorTable);
    }
}

contract OperatorTableUpdaterUnitTests_multipleCurveTypes is OperatorTableUpdaterUnitTests {
    function testFuzz_correctness(Randomness r) public rand(r) {
        // Generate random BN254 operatorSetInfo and operatorSetConfig
        bytes memory bn254OperatorTable;
        {
            BN254OperatorSetInfo memory bn254OperatorSetInfo = _generateRandomBN254OperatorSetInfo(r);
            bytes memory bn254OperatorSetInfoBytes = abi.encode(bn254OperatorSetInfo);
            OperatorSetConfig memory bn254OperatorSetConfig = _generateRandomOperatorSetConfig(r);
            bn254OperatorTable = abi.encode(defaultOperatorSet, CurveType.BN254, bn254OperatorSetConfig, bn254OperatorSetInfoBytes);
        }

        // Generate random ECDSA operatorInfos and operatorSetConfig
        bytes memory ecdsaOperatorTable;
        {
            ECDSAOperatorInfo[] memory ecdsaOperatorInfos = _generateRandomECDSAOperatorInfos(r);
            bytes memory ecdsaOperatorInfosBytes = abi.encode(ecdsaOperatorInfos);
            OperatorSetConfig memory ecdsaOperatorSetConfig = _generateRandomOperatorSetConfig(r);
            ecdsaOperatorTable = abi.encode(defaultOperatorSet, CurveType.ECDSA, ecdsaOperatorSetConfig, ecdsaOperatorInfosBytes);
        }

        // Include the operatorInfos and operatorSetConfig in the global table root & set it
        bytes32[] memory operatorSetLeafHashes = new bytes32[](2);
        operatorSetLeafHashes[0] = keccak256(bn254OperatorTable);
        operatorSetLeafHashes[1] = keccak256(ecdsaOperatorTable);
        (bytes32 globalTableRoot, uint32[] memory operatorSetIndices, bytes32[] memory leaves) =
            _createGlobalTableRoot(r, operatorSetLeafHashes);
        _updateGlobalTableRoot(globalTableRoot);

        // Update the operator table for bn254
        bytes memory bn254Proof = Merkle.getProofKeccak(leaves, operatorSetIndices[0]);
        operatorTableUpdater.updateOperatorTable(
            uint32(block.timestamp), globalTableRoot, operatorSetIndices[0], bn254Proof, bn254OperatorTable
        );

        // Update the operator table for ecdsa
        bytes memory ecdsaProof = Merkle.getProofKeccak(leaves, operatorSetIndices[1]);
        operatorTableUpdater.updateOperatorTable(
            uint32(block.timestamp), globalTableRoot, operatorSetIndices[1], ecdsaProof, ecdsaOperatorTable
        );
    }
}

contract OperatorTableUpdaterUnitTests_getCertificateVerifier is OperatorTableUpdaterUnitTests {
    function test_revert_invalidCurveType() public {
        cheats.expectRevert(InvalidCurveType.selector);
        operatorTableUpdater.getCertificateVerifier(CurveType.NONE);
    }

    function testFuzz_correctness(Randomness r) public rand(r) {
        CurveType curveType = CurveType(r.Uint256(1, 2));
        // Get the certificate verifier
        address certificateVerifier = operatorTableUpdater.getCertificateVerifier(curveType);

        // Expect the certificate verifier to be the correct address
        if (curveType == CurveType.BN254) assertEq(certificateVerifier, address(bn254CertificateVerifierMock));
        else if (curveType == CurveType.ECDSA) assertEq(certificateVerifier, address(ecdsaCertificateVerifierMock));
        else revert("Invalid State");
    }
}

contract OperatorTableUpdaterUnitTests_getters is OperatorTableUpdaterUnitTests {
    function testFuzz_getGlobalTableRootByTimestamp(Randomness r) public rand(r) {
        uint32 referenceTimestamp = r.Uint32(operatorTableUpdater.getLatestReferenceTimestamp() + 1, type(uint32).max);
        uint32 referenceBlockNumber = r.Uint32();
        cheats.warp(uint(referenceTimestamp));
        bytes32 globalTableRoot = bytes32(r.Uint256(1, type(uint).max));

        // Set a global table root
        BN254Certificate memory mockCertificate;
        mockCertificate.messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber);
        _setIsValidCertificate(mockCertificate, true);
        operatorTableUpdater.confirmGlobalTableRoot(mockCertificate, globalTableRoot, referenceTimestamp, referenceBlockNumber);

        // Test getters
        assertEq(operatorTableUpdater.getGlobalTableRootByTimestamp(referenceTimestamp), globalTableRoot);
        assertEq(operatorTableUpdater.getCurrentGlobalTableRoot(), globalTableRoot);
        assertEq(operatorTableUpdater.getLatestReferenceTimestamp(), referenceTimestamp);
        assertEq(operatorTableUpdater.getLatestReferenceBlockNumber(), referenceBlockNumber);
        assertEq(operatorTableUpdater.getReferenceBlockNumberByTimestamp(referenceTimestamp), referenceBlockNumber);
        assertEq(operatorTableUpdater.getReferenceTimestampByBlockNumber(referenceBlockNumber), referenceTimestamp);
    }

    function test_getGlobalRootConfirmerSet() public view {
        OperatorSet memory confirmerSet = operatorTableUpdater.getGlobalRootConfirmerSet();
        assertEq(confirmerSet.avs, globalRootConfirmerSet.avs);
        assertEq(confirmerSet.id, globalRootConfirmerSet.id);
    }

    function test_getGlobalConfirmerSetReferenceTimestamp() public view {
        uint32 timestamp = operatorTableUpdater.getGlobalConfirmerSetReferenceTimestamp();
        assertEq(timestamp, uint32(block.timestamp - 1));
    }

    function testFuzz_getGlobalTableUpdateMessageHash(Randomness r) public rand(r) {
        bytes32 globalTableRoot = bytes32(r.Uint256());
        uint32 referenceTimestamp = r.Uint32();
        uint32 referenceBlockNumber = r.Uint32();

        bytes32 messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber);

        // Verify the hash is deterministic
        bytes32 messageHash2 =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, referenceTimestamp, referenceBlockNumber);
        assertEq(messageHash, messageHash2);

        // Verify hash changes with different inputs
        bytes32 differentHash = operatorTableUpdater.getGlobalTableUpdateMessageHash(
            bytes32(uint(globalTableRoot) + 1), referenceTimestamp, referenceBlockNumber
        );
        assertTrue(messageHash != differentHash);
    }
}

contract OperatorTableUpdaterUnitTests_setGlobalRootConfirmerSet is OperatorTableUpdaterUnitTests {
    function testFuzz_revert_onlyOwner(Randomness r) public rand(r) {
        address invalidCaller = r.Address();
        cheats.assume(invalidCaller != address(this));
        OperatorSet memory newSet;

        // Should revert when called by non-owner
        cheats.prank(invalidCaller);
        cheats.expectRevert("Ownable: caller is not the owner");
        operatorTableUpdater.setGlobalRootConfirmerSet(newSet);
    }

    function testFuzz_correctness(Randomness r) public rand(r) {
        // Generate random operator set
        OperatorSet memory newOperatorSet = OperatorSet({avs: r.Address(), id: r.Uint32()});

        // Set the confirmer set
        cheats.expectEmit(true, true, true, true);
        emit GlobalRootConfirmerSetUpdated(newOperatorSet);
        operatorTableUpdater.setGlobalRootConfirmerSet(newOperatorSet);

        // Verify the storage was updated
        OperatorSet memory updatedConfirmerSet = operatorTableUpdater.getGlobalRootConfirmerSet();
        assertEq(updatedConfirmerSet.avs, newOperatorSet.avs);
        assertEq(updatedConfirmerSet.id, newOperatorSet.id);
    }
}

contract OperatorTableUpdaterUnitTests_setGlobalRootConfirmationThreshold is OperatorTableUpdaterUnitTests {
    function testFuzz_revert_onlyOwner(Randomness r) public rand(r) {
        address invalidCaller = r.Address();
        cheats.assume(invalidCaller != address(this));
        uint16 newThreshold;

        // Should revert when called by non-owner
        cheats.prank(invalidCaller);
        cheats.expectRevert("Ownable: caller is not the owner");
        operatorTableUpdater.setGlobalRootConfirmationThreshold(newThreshold);
    }

    function testFuzz_revert_invalidThreshold(Randomness r) public rand(r) {
        uint16 newThreshold = uint16(r.Uint256(10_001, type(uint16).max));
        cheats.expectRevert(InvalidConfirmationThreshold.selector);
        operatorTableUpdater.setGlobalRootConfirmationThreshold(newThreshold);
    }

    function testFuzz_correctness(Randomness r) public rand(r) {
        // Generate random threshold (0 to 10000 bps = 0% to 100%)
        uint16 newThreshold = uint16(r.Uint256(0, 10_000));

        // Expect the event to be emitted
        cheats.expectEmit(true, true, true, true);
        emit GlobalRootConfirmationThresholdUpdated(newThreshold);
        operatorTableUpdater.setGlobalRootConfirmationThreshold(newThreshold);

        // Verify the storage was updated
        uint16 updatedThreshold = operatorTableUpdater.globalRootConfirmationThreshold();
        assertEq(updatedThreshold, newThreshold);
    }
}
