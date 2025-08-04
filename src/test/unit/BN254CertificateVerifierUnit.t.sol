// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/interfaces/IBaseCertificateVerifier.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";

/**
 * @title BN254CertificateVerifierUnitTests
 * @notice Base contract for all BN254CertificateVerifier unit tests
 */
contract BN254CertificateVerifierUnitTests is
    EigenLayerMultichainUnitTestSetup,
    IBaseCertificateVerifierErrors,
    IBN254CertificateVerifierTypes,
    IBN254CertificateVerifierErrors,
    ICrossChainRegistryTypes
{
    using BN254 for BN254.G1Point;
    using Merkle for bytes;
    using OperatorSetLib for OperatorSet;

    // Constants
    uint16 constant BPS_DENOMINATOR = 10_000;

    // Contracts
    BN254CertificateVerifier bn254CertificateVerifierImplementation;
    BN254CertificateVerifier verifier;

    // Test accounts
    address owner = address(0x1);
    address operatorSetOwner = address(0x4);

    // Defaults
    uint32 numOperators = 4;
    uint32 defaultMaxStaleness = 3600; // 1 hour max staleness
    OperatorSet defaultOperatorSet;
    OperatorSetConfig defaultOperatorSetConfig;

    // BLS signature specific fields
    bytes32 defaultMsgHash = keccak256(abi.encodePacked("test message"));
    uint aggSignerPrivKey = 69;
    BN254.G2Point aggSignerApkG2; // G2 public key corresponding to aggSignerPrivKey

    function setUp() public virtual override {
        // Setup Mocks
        super.setUp();

        defaultOperatorSet = OperatorSet({avs: address(0x5), id: 0});

        defaultOperatorSetConfig = OperatorSetConfig({owner: operatorSetOwner, maxStalenessPeriod: defaultMaxStaleness});

        // Deploy Contracts
        bn254CertificateVerifierImplementation =
            new BN254CertificateVerifier(IOperatorTableUpdater(address(operatorTableUpdaterMock)), "1.0.0");
        verifier = BN254CertificateVerifier(
            address(new TransparentUpgradeableProxy(address(bn254CertificateVerifierImplementation), address(eigenLayerProxyAdmin), ""))
        );

        // Set up the aggregate public key in G2
        aggSignerApkG2.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
        aggSignerApkG2.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
        aggSignerApkG2.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
        aggSignerApkG2.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
    }

    // Helper functions
    function _generateSignerAndNonSignerPrivateKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        view
        returns (uint[] memory signerPrivKeys, uint[] memory nonSignerPrivKeys)
    {
        signerPrivKeys = new uint[](numSigners);
        uint sum = 0;

        // Generate numSigners-1 random keys
        for (uint i = 0; i < numSigners - 1; i++) {
            signerPrivKeys[i] = uint(keccak256(abi.encodePacked("signerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
            sum = addmod(sum, signerPrivKeys[i], BN254.FR_MODULUS);
        }

        // Last key makes the total sum equal to aggSignerPrivKey
        signerPrivKeys[numSigners - 1] = addmod(aggSignerPrivKey, BN254.FR_MODULUS - sum % BN254.FR_MODULUS, BN254.FR_MODULUS);

        // Generate non-signer keys
        nonSignerPrivKeys = new uint[](numNonSigners);
        for (uint i = 0; i < numNonSigners; i++) {
            nonSignerPrivKeys[i] = uint(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
        }

        // Sort nonSignerPrivateKeys in order of ascending pubkeyHash
        for (uint i = 1; i < nonSignerPrivKeys.length; i++) {
            uint privateKey = nonSignerPrivKeys[i];
            bytes32 pubkeyHash = _toPubkeyHash(privateKey);
            uint j = i;

            while (j > 0 && _toPubkeyHash(nonSignerPrivKeys[j - 1]) > pubkeyHash) {
                nonSignerPrivKeys[j] = nonSignerPrivKeys[j - 1];
                j--;
            }
            nonSignerPrivKeys[j] = privateKey;
        }
    }

    function _toPubkeyHash(uint privKey) internal view returns (bytes32) {
        return BN254.generatorG1().scalar_mul(privKey).hashG1Point();
    }

    function _createOperatorsWithSplitKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        view
        returns (BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
    {
        require(numSigners + numNonSigners == numOperators, "Total operators mismatch");

        // Generate private keys
        (uint[] memory signerPrivKeys, uint[] memory nonSignerPrivKeys) =
            _generateSignerAndNonSignerPrivateKeys(pseudoRandomNumber, numSigners, numNonSigners);

        // Create all operators
        operators = new BN254OperatorInfo[](numOperators);

        // Track indices of non-signers
        nonSignerIndices = new uint32[](numNonSigners);

        // Create signers first
        for (uint32 i = 0; i < numSigners; i++) {
            operators[i].pubkey = BN254.generatorG1().scalar_mul(signerPrivKeys[i]);
            operators[i].weights = new uint[](2);
            operators[i].weights[0] = uint(100 + i * 10);
            operators[i].weights[1] = uint(200 + i * 20);
        }

        // Create non-signers
        for (uint32 i = 0; i < numNonSigners; i++) {
            uint32 idx = uint32(numSigners + i);
            operators[idx].pubkey = BN254.generatorG1().scalar_mul(nonSignerPrivKeys[i]);
            operators[idx].weights = new uint[](2);
            operators[idx].weights[0] = uint(100 + idx * 10);
            operators[idx].weights[1] = uint(200 + idx * 20);
            nonSignerIndices[i] = idx;
        }

        // Calculate aggregate signature for the signers
        signature = BN254.hashToG1(defaultMsgHash).scalar_mul(aggSignerPrivKey);
    }

    function _getMerkleRoot(BN254OperatorInfo[] memory ops) internal view returns (bytes32 root) {
        bytes32[] memory leaves = new bytes32[](ops.length);
        for (uint i = 0; i < ops.length; i++) {
            leaves[i] = verifier.calculateOperatorInfoLeaf(ops[i]);
        }
        root = Merkle.merkleizeKeccak(leaves);
    }

    function _getMerkleProof(BN254OperatorInfo[] memory ops, uint32 operatorIndex) internal view returns (bytes memory proof) {
        bytes32[] memory leaves = new bytes32[](ops.length);
        for (uint i = 0; i < ops.length; i++) {
            leaves[i] = verifier.calculateOperatorInfoLeaf(ops[i]);
        }
        proof = Merkle.getProofKeccak(leaves, operatorIndex);
    }

    function _createOperatorSetInfo(BN254OperatorInfo[] memory ops) internal view returns (BN254OperatorSetInfo memory) {
        uint32 _numOperators = uint32(ops.length);
        bytes32 operatorInfoTreeRoot = _getMerkleRoot(ops);

        // Create aggregate public key
        BN254.G1Point memory aggregatePubkey = BN254.G1Point(0, 0);

        // Create total weights
        uint[] memory _totalWeights = new uint[](2);

        for (uint32 i = 0; i < _numOperators; i++) {
            aggregatePubkey = aggregatePubkey.plus(ops[i].pubkey);
            for (uint j = 0; j < 2; j++) {
                _totalWeights[j] += ops[i].weights[j];
            }
        }

        return BN254OperatorSetInfo({
            numOperators: _numOperators,
            aggregatePubkey: aggregatePubkey,
            totalWeights: _totalWeights,
            operatorInfoTreeRoot: operatorInfoTreeRoot
        });
    }

    function _createCertificate(
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint32[] memory nonSignerIndices,
        BN254OperatorInfo[] memory ops,
        BN254.G1Point memory signature
    ) internal view returns (BN254Certificate memory) {
        // Create witnesses for non-signers
        BN254OperatorInfoWitness[] memory witnesses = new BN254OperatorInfoWitness[](nonSignerIndices.length);

        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            bytes memory proof = _getMerkleProof(ops, nonSignerIndex);

            witnesses[i] =
                BN254OperatorInfoWitness({operatorIndex: nonSignerIndex, operatorInfoProof: proof, operatorInfo: ops[nonSignerIndex]});
        }

        return BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: witnesses
        });
    }

    function _updateOperatorTable(Randomness r, uint numSigners, uint numNonsigners)
        internal
        returns (
            BN254OperatorInfo[] memory operatorInfos,
            BN254OperatorSetInfo memory operatorSetInfo,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            BN254.G1Point memory signature
        )
    {
        // Generate seed and reference timestamp
        uint seed = r.Uint256();
        referenceTimestamp = r.Uint32(uint32(block.timestamp + 1), uint32(block.timestamp + 1000));

        // Create operators
        (operatorInfos, nonSignerIndices, signature) = _createOperatorsWithSplitKeys(seed, numSigners, numNonsigners);
        operatorSetInfo = _createOperatorSetInfo(operatorInfos);

        // Update operator table
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);
    }

    function _initializeOperatorTableBase() internal returns (uint32 referenceTimestamp) {
        (BN254OperatorInfo[] memory operatorInfos, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operatorInfos);
        vm.prank(address(operatorTableUpdaterMock));
        referenceTimestamp = uint32(block.timestamp);
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);
        return referenceTimestamp;
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_updateOperatorTable
 * @notice Unit tests for BN254CertificateVerifier.updateOperatorTable
 */
contract BN254CertificateVerifierUnitTests_updateOperatorTable is BN254CertificateVerifierUnitTests, IBN254CertificateVerifierEvents {
    function test_revert_notTableUpdater() public {
        uint32 referenceTimestamp = uint32(block.timestamp);
        (BN254OperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 3, 1);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        vm.prank(address(0x100));
        vm.expectRevert(OnlyTableUpdater.selector);
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);
    }

    function test_revert_staleTimestamp() public {
        vm.warp(1000);
        uint32 firstTimestamp = uint32(block.timestamp);
        uint32 staleTimestamp = firstTimestamp - 100;

        (BN254OperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 3, 1);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        vm.startPrank(address(operatorTableUpdaterMock));

        // First update should succeed
        verifier.updateOperatorTable(defaultOperatorSet, firstTimestamp, operatorSetInfo, defaultOperatorSetConfig);

        // Second update with earlier timestamp should fail
        vm.expectRevert(TableUpdateStale.selector);
        verifier.updateOperatorTable(defaultOperatorSet, staleTimestamp, operatorSetInfo, defaultOperatorSetConfig);

        vm.stopPrank();
    }

    function testFuzz_updateOperatorTable_correctness(Randomness r) public rand(r) {
        uint seed = r.Uint256();
        uint numSigners = r.Uint256(1, numOperators);
        uint nonSigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(seed, numSigners, nonSigners);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        // Expect event
        uint32 referenceTimestamp = uint32(block.timestamp);
        vm.expectEmit(true, true, true, true, address(verifier));
        emit TableUpdated(defaultOperatorSet, referenceTimestamp, operatorSetInfo);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);

        // Verify storage updates
        assertEq(verifier.latestReferenceTimestamp(defaultOperatorSet), referenceTimestamp, "Reference timestamp mismatch");
        assertEq(verifier.getOperatorSetOwner(defaultOperatorSet), operatorSetOwner, "Operator set owner mismatch");
        assertEq(verifier.maxOperatorTableStaleness(defaultOperatorSet), defaultMaxStaleness, "Max staleness mismatch");

        // Verify operator set info was stored correctly
        BN254OperatorSetInfo memory storedOperatorSetInfo = verifier.getOperatorSetInfo(defaultOperatorSet, referenceTimestamp);

        assertEq(storedOperatorSetInfo.numOperators, operatorSetInfo.numOperators, "Num operators mismatch");
        assertEq(storedOperatorSetInfo.aggregatePubkey.X, operatorSetInfo.aggregatePubkey.X, "Aggregate pubkey X mismatch");
        assertEq(storedOperatorSetInfo.aggregatePubkey.Y, operatorSetInfo.aggregatePubkey.Y, "Aggregate pubkey Y mismatch");
        assertEq(storedOperatorSetInfo.operatorInfoTreeRoot, operatorSetInfo.operatorInfoTreeRoot, "Tree root mismatch");
    }

    function test_multiple() public {
        // Create two different operator sets
        OperatorSet memory operatorSet1 = OperatorSet({avs: address(0x10), id: 1});
        OperatorSet memory operatorSet2 = OperatorSet({avs: address(0x20), id: 2});

        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators for each set
        (BN254OperatorInfo[] memory operators1,,) = _createOperatorsWithSplitKeys(111, 3, 1);
        (BN254OperatorInfo[] memory operators2,,) = _createOperatorsWithSplitKeys(222, 2, 2);

        // Create operator set infos
        BN254OperatorSetInfo memory operatorSetInfo1 = _createOperatorSetInfo(operators1);
        BN254OperatorSetInfo memory operatorSetInfo2 = _createOperatorSetInfo(operators2);

        // Create operator set configs with different owners
        OperatorSetConfig memory config1 = OperatorSetConfig({
            owner: address(0x100),
            maxStalenessPeriod: 1800 // 30 minutes
        });
        OperatorSetConfig memory config2 = OperatorSetConfig({
            owner: address(0x200),
            maxStalenessPeriod: 7200 // 2 hours
        });

        vm.startPrank(address(operatorTableUpdaterMock));

        // Update both operator tables
        verifier.updateOperatorTable(operatorSet1, referenceTimestamp, operatorSetInfo1, config1);
        verifier.updateOperatorTable(operatorSet2, referenceTimestamp, operatorSetInfo2, config2);

        vm.stopPrank();

        // Verify that both operator sets are stored correctly and independently
        assertEq(verifier.getOperatorSetOwner(operatorSet1), address(0x100), "OperatorSet1 owner incorrect");
        assertEq(verifier.getOperatorSetOwner(operatorSet2), address(0x200), "OperatorSet2 owner incorrect");

        assertEq(verifier.maxOperatorTableStaleness(operatorSet1), 1800, "OperatorSet1 staleness incorrect");
        assertEq(verifier.maxOperatorTableStaleness(operatorSet2), 7200, "OperatorSet2 staleness incorrect");

        assertEq(verifier.latestReferenceTimestamp(operatorSet1), referenceTimestamp, "OperatorSet1 timestamp incorrect");
        assertEq(verifier.latestReferenceTimestamp(operatorSet2), referenceTimestamp, "OperatorSet2 timestamp incorrect");

        // Verify operator set infos are stored independently
        BN254OperatorSetInfo memory storedInfo1 = verifier.getOperatorSetInfo(operatorSet1, referenceTimestamp);
        BN254OperatorSetInfo memory storedInfo2 = verifier.getOperatorSetInfo(operatorSet2, referenceTimestamp);

        assertEq(storedInfo1.numOperators, 4, "OperatorSet1 should have 4 operators");
        assertEq(storedInfo2.numOperators, 4, "OperatorSet2 should have 4 operators");

        // Verify they have different tree roots (since operators are different)
        assertTrue(storedInfo1.operatorInfoTreeRoot != storedInfo2.operatorInfoTreeRoot, "Tree roots should be different");
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_verifyCertificate
 * @notice Unit tests for BN254CertificateVerifier.verifyCertificate
 */
contract BN254CertificateVerifierUnitTests_verifyCertificate is BN254CertificateVerifierUnitTests {
    function test_revert_certificateStale() public {
        uint32 referenceTimestamp = _initializeOperatorTableBase();
        BN254Certificate memory cert;
        cert.referenceTimestamp = referenceTimestamp;

        // Jump forward in time beyond the max staleness
        vm.warp(block.timestamp + defaultMaxStaleness + 1);

        vm.expectRevert(CertificateStale.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_zeroStalenessPeriod_neverStale() public {
        // Create operator set config with 0 staleness period
        OperatorSetConfig memory zeroStalenessConfig = OperatorSetConfig({
            owner: operatorSetOwner,
            maxStalenessPeriod: 0 // 0 means certificate never becomes stale
        });

        uint32 referenceTimestamp = uint32(block.timestamp);
        (BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, 4, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, zeroStalenessConfig);

        // Create certificate
        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Jump forward in time far beyond any reasonable staleness period
        vm.warp(block.timestamp + 365 days);

        // Certificate should still be valid with 0 staleness period
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Verify all stakes are signed
        assertEq(signedStakes[0], operatorSetInfo.totalWeights[0], "All stake should be signed for type 0");
        assertEq(signedStakes[1], operatorSetInfo.totalWeights[1], "All stake should be signed for type 1");
    }

    function test_revert_referenceTimestampDoesNotExist() public {
        uint32 referenceTimestamp = _initializeOperatorTableBase();
        uint32 nonExistentTimestamp = referenceTimestamp + 1000;

        BN254Certificate memory cert;
        cert.referenceTimestamp = nonExistentTimestamp;

        vm.expectRevert(ReferenceTimestampDoesNotExist.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_isReferenceTimestampSet() public {
        uint32 referenceTimestamp = uint32(block.timestamp);
        (BN254OperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 4, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        // Before updating operator table, timestamp should not be set
        assertFalse(verifier.isReferenceTimestampSet(defaultOperatorSet, referenceTimestamp), "Timestamp should not be set before update");

        // Update operator table
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);

        // After updating, timestamp should be set
        assertTrue(verifier.isReferenceTimestampSet(defaultOperatorSet, referenceTimestamp), "Timestamp should be set after update");

        // A different timestamp should not be set
        assertFalse(verifier.isReferenceTimestampSet(defaultOperatorSet, referenceTimestamp + 1), "Different timestamp should not be set");
    }

    function test_revert_rootDisabled() public {
        // Initialize operator table with a valid root
        uint32 referenceTimestamp = _initializeOperatorTableBase();

        // Mock the operatorTableUpdater to return false for isRootValidByTimestamp
        operatorTableUpdaterMock.invalidateRoot(referenceTimestamp);

        BN254Certificate memory cert;
        cert.referenceTimestamp = referenceTimestamp;

        vm.expectRevert(RootDisabled.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_invalidOperatorIndex() public {
        // Update operator table
        (BN254OperatorInfo[] memory operatorInfos, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operatorInfos);
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, uint32(block.timestamp), operatorSetInfo, defaultOperatorSetConfig);

        // Create certificate with invalid operator index
        BN254OperatorInfoWitness[] memory witnesses = new BN254OperatorInfoWitness[](1);
        witnesses[0] = BN254OperatorInfoWitness({
            operatorIndex: uint32(operatorSetInfo.numOperators), // Out of bounds, since we only have 1 nonsigner
            operatorInfoProof: new bytes(32),
            operatorInfo: operatorInfos[0]
        });

        BN254Certificate memory cert = BN254Certificate({
            referenceTimestamp: uint32(block.timestamp),
            messageHash: defaultMsgHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: witnesses
        });

        vm.expectRevert(InvalidOperatorIndex.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_invalidMerkleProof() public {
        // Update operator table
        (BN254OperatorInfo[] memory operatorInfos, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, 3, 1);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operatorInfos);
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, uint32(block.timestamp), operatorSetInfo, defaultOperatorSetConfig);

        // Create certificate with invalid merkle proof
        BN254OperatorInfoWitness[] memory witnesses = new BN254OperatorInfoWitness[](1);
        witnesses[0] = BN254OperatorInfoWitness({
            operatorIndex: nonSignerIndices[0],
            operatorInfoProof: new bytes(32), // Invalid proof
            operatorInfo: operatorInfos[nonSignerIndices[0]]
        });

        BN254Certificate memory cert = BN254Certificate({
            referenceTimestamp: uint32(block.timestamp),
            messageHash: defaultMsgHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: witnesses
        });

        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_invalidSignature() public {
        // Update operator table
        (BN254OperatorInfo[] memory operatorInfos, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operatorInfos);
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, uint32(block.timestamp), operatorSetInfo, defaultOperatorSetConfig);

        // Create certificate with wrong message hash
        bytes32 wrongHash = keccak256("wrong message");

        BN254Certificate memory cert = _createCertificate(
            uint32(block.timestamp),
            wrongHash, // Wrong hash
            nonSignerIndices,
            operatorInfos,
            signature // Signature is for original defaultMsgHash
        );

        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function testFuzz_verifyCertificate_allSigners(Randomness r) public rand(r) {
        (, BN254OperatorSetInfo memory allSignerOpSetInfo, uint32 referenceTimestamp,, BN254.G1Point memory signature) =
            _updateOperatorTable(r, numOperators, 0);

        // Create certificate with no non-signers
        BN254Certificate memory cert = BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: new BN254OperatorInfoWitness[](0)
        });

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // All stakes should be signed
        assertEq(signedStakes[0], allSignerOpSetInfo.totalWeights[0], "All stake should be signed for type 0");
        assertEq(signedStakes[1], allSignerOpSetInfo.totalWeights[1], "All stake should be signed for type 1");
    }

    function testFuzz_verifyCertificate_someNonSigners(Randomness r) public rand(r) {
        // Create operators and update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint nonSigners = numOperators - numSigners;
        (
            BN254OperatorInfo[] memory operators,
            BN254OperatorSetInfo memory newOperatorSetInfo,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            BN254.G1Point memory signature
        ) = _updateOperatorTable(r, numSigners, nonSigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Check that the signed stakes are correct
        assertEq(signedStakes.length, 2, "Wrong number of stake types");

        // Calculate expected signed stakes
        uint[] memory expectedSignedStakes = new uint[](2);
        expectedSignedStakes[0] = newOperatorSetInfo.totalWeights[0];
        expectedSignedStakes[1] = newOperatorSetInfo.totalWeights[1];

        // Subtract non-signer stakes
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            expectedSignedStakes[0] -= operators[nonSignerIndex].weights[0];
            expectedSignedStakes[1] -= operators[nonSignerIndex].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");
    }

    function test_verifyCertificate_cachesOperatorInfo(Randomness r) public rand(r) {
        // Create operators
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint nonSigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, nonSigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // First verification should cache the operator infos
        verifier.verifyCertificate(defaultOperatorSet, cert);

        // Check that operator infos are now cached
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            BN254OperatorInfo memory cachedInfo = verifier.getNonsignerOperatorInfo(defaultOperatorSet, referenceTimestamp, nonSignerIndex);

            // Cached info should match original
            assertEq(cachedInfo.pubkey.X, operators[nonSignerIndex].pubkey.X, "Cached pubkey X mismatch");
            assertEq(cachedInfo.pubkey.Y, operators[nonSignerIndex].pubkey.Y, "Cached pubkey Y mismatch");
            assertEq(cachedInfo.weights[0], operators[nonSignerIndex].weights[0], "Cached weight 0 mismatch");
            assertEq(cachedInfo.weights[1], operators[nonSignerIndex].weights[1], "Cached weight 1 mismatch");
            assertTrue(verifier.isNonsignerCached(defaultOperatorSet, referenceTimestamp, nonSignerIndex), "Operator should be cached");
        }

        // Second verification should use cached data. We don't have to provide an operatorInfoProof for the non-signers
        // since they are in the cache.
        BN254OperatorInfo memory emptyOperatorInfo;
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            cert.nonSignerWitnesses[i].operatorInfoProof = new bytes(0);
            cert.nonSignerWitnesses[i].operatorInfo = emptyOperatorInfo;
        }
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_verifyCertificate_olderTimestamp() public {
        // First update with initial operators
        uint32 firstTimestamp = uint32(block.timestamp);
        (BN254OperatorInfo[] memory firstOperators, uint32[] memory firstNonSignerIndices, BN254.G1Point memory firstSignature) =
            _createOperatorsWithSplitKeys(123, 3, 1);
        BN254OperatorSetInfo memory firstOperatorSetInfo = _createOperatorSetInfo(firstOperators);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, firstTimestamp, firstOperatorSetInfo, defaultOperatorSetConfig);

        // Advance time and update with new operators (making this the latest timestamp)
        vm.warp(block.timestamp + 100);
        uint32 secondTimestamp = uint32(block.timestamp);
        (BN254OperatorInfo[] memory secondOperators,,) = _createOperatorsWithSplitKeys(456, 2, 2);
        BN254OperatorSetInfo memory secondOperatorSetInfo = _createOperatorSetInfo(secondOperators);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, secondTimestamp, secondOperatorSetInfo, defaultOperatorSetConfig);

        // Verify that the second timestamp is now the latest
        assertEq(verifier.latestReferenceTimestamp(defaultOperatorSet), secondTimestamp, "Second timestamp should be latest");

        // Create certificate for the FIRST (older) timestamp
        BN254Certificate memory cert =
            _createCertificate(firstTimestamp, defaultMsgHash, firstNonSignerIndices, firstOperators, firstSignature);

        // Verify certificate for older timestamp should succeed
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Calculate expected signed stakes from first operators (total minus non-signers)
        uint[] memory expectedSignedStakes = new uint[](2);
        expectedSignedStakes[0] = firstOperatorSetInfo.totalWeights[0];
        expectedSignedStakes[1] = firstOperatorSetInfo.totalWeights[1];

        // Subtract non-signer stakes
        for (uint i = 0; i < firstNonSignerIndices.length; i++) {
            uint32 nonSignerIndex = firstNonSignerIndices[i];
            expectedSignedStakes[0] -= firstOperators[nonSignerIndex].weights[0];
            expectedSignedStakes[1] -= firstOperators[nonSignerIndex].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");

        // Verify the signed stakes match expected calculation
        assertEq(signedStakes.length, 2, "Wrong number of stake types returned");
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_verifyCertificateProportion
 * @notice Unit tests for BN254CertificateVerifier.verifyCertificateProportion
 */
contract BN254CertificateVerifierUnitTests_verifyCertificateProportion is BN254CertificateVerifierUnitTests {
    function testFuzz_revert_arrayLengthMismatch(Randomness r) public rand(r) {
        // Update operator table
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numOperators, 0); // 0 nonsigners

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        uint16[] memory wrongLengthThresholds = new uint16[](1); // Should be 2
        wrongLengthThresholds[0] = 6000;

        vm.expectRevert(ArrayLengthMismatch.selector);
        verifier.verifyCertificateProportion(defaultOperatorSet, cert, wrongLengthThresholds);
    }

    function testFuzz_verifyCertificateProportion_meetsThresholds(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint numNonsigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Set thresholds at 60% of total stake for each type
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 6000; // 60%
        thresholds[1] = 6000; // 60%

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate expected result based on the number of signers
        // With numSigners out of numOperators, check if it meets 60% threshold
        uint signedPercentage = (numSigners * 10_000) / numOperators;
        bool expectedResult = signedPercentage >= 6000;

        assertEq(meetsThresholds, expectedResult, "Certificate threshold check incorrect");
    }

    function testFuzz_verifyCertificateProportion_doesNotMeetThresholds(Randomness r) public rand(r) {
        // Update operator table with a specific split to ensure some thresholds won't be met
        uint numSigners = r.Uint256(1, numOperators / 2); // At most half signers
        uint numNonsigners = numOperators - numSigners;
        (
            BN254OperatorInfo[] memory operators,
            BN254OperatorSetInfo memory operatorSetInfo,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            BN254.G1Point memory signature
        ) = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Try with higher threshold that shouldn't be met
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 9000; // 90%
        thresholds[1] = 9000; // 90%

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate percentage of signed stakes to determine if it should meet threshold
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);
        uint signedPercentage0 = (signedStakes[0] * 10_000) / operatorSetInfo.totalWeights[0];
        uint signedPercentage1 = (signedStakes[1] * 10_000) / operatorSetInfo.totalWeights[1];

        bool shouldMeetThreshold = (signedPercentage0 >= 9000) && (signedPercentage1 >= 9000);
        assertEq(meetsThresholds, shouldMeetThreshold, "Certificate threshold check incorrect");
    }

    /// @notice Fuzz against random thresholds
    function testFuzz_verifyCertificateProportion_thresholds(Randomness r, uint16 threshold0, uint16 threshold1) public rand(r) {
        // Update operator table with random split
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonsigners = numOperators - numSigners;
        (
            BN254OperatorInfo[] memory operators,
            BN254OperatorSetInfo memory operatorSetInfo,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            BN254.G1Point memory signature
        ) = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        threshold0 = uint16(bound(threshold0, 0, 10_000)); // 0% to 100%
        threshold1 = uint16(bound(threshold1, 0, 10_000)); // 0% to 100%

        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = threshold0;
        thresholds[1] = threshold1;

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate expected result
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);
        bool expectedResult = true;
        for (uint i = 0; i < 2; i++) {
            uint threshold = (operatorSetInfo.totalWeights[i] * thresholds[i]) / 10_000;
            if (signedStakes[i] < threshold) {
                expectedResult = false;
                break;
            }
        }

        assertEq(meetsThresholds, expectedResult, "Threshold calculation mismatch");
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_verifyCertificateNominal
 * @notice Unit tests for BN254CertificateVerifier.verifyCertificateNominal
 */
contract BN254CertificateVerifierUnitTests_verifyCertificateNominal is BN254CertificateVerifierUnitTests {
    function testFuzz_revert_arrayLengthMismatch(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint numNonsigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        uint[] memory wrongLengthThresholds = new uint[](1); // Should be 2
        wrongLengthThresholds[0] = 100;

        vm.expectRevert(ArrayLengthMismatch.selector);
        verifier.verifyCertificateNominal(defaultOperatorSet, cert, wrongLengthThresholds);
    }

    function testFuzz_verifyCertificateNominal_meetsThresholds(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonsigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Get the signed stakes for reference
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Test with thresholds lower than signed stakes (should pass)
        uint[] memory passThresholds = new uint[](2);
        passThresholds[0] = signedStakes[0] > 10 ? signedStakes[0] - 10 : 0;
        passThresholds[1] = signedStakes[1] > 10 ? signedStakes[1] - 10 : 0;

        bool meetsThresholds = verifier.verifyCertificateNominal(defaultOperatorSet, cert, passThresholds);
        assertTrue(meetsThresholds, "Certificate should meet nominal thresholds");
    }

    function testFuzz_verifyCertificateNominal_doesNotMeetThresholds(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators - 1); // Ensure at least one non-signer
        uint numNonsigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Get the signed stakes for reference
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Test with thresholds higher than signed stakes (should fail)
        uint[] memory failThresholds = new uint[](2);
        failThresholds[0] = signedStakes[0] + 10;
        failThresholds[1] = signedStakes[1] + 10;

        bool meetsThresholds = verifier.verifyCertificateNominal(defaultOperatorSet, cert, failThresholds);
        assertFalse(meetsThresholds, "Certificate should not meet impossible nominal thresholds");
    }

    /// @notice Fuzz against random thresholds
    function testFuzz_verifyCertificateNominal_thresholds(Randomness r, uint threshold0, uint threshold1) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonsigners = numOperators - numSigners;
        (BN254OperatorInfo[] memory operators,, uint32 referenceTimestamp, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = _updateOperatorTable(r, numSigners, numNonsigners);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Get the signed stakes for reference
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Bound thresholds to reasonable values
        threshold0 = bound(threshold0, 0, signedStakes[0] * 2);
        threshold1 = bound(threshold1, 0, signedStakes[1] * 2);

        uint[] memory thresholds = new uint[](2);
        thresholds[0] = threshold0;
        thresholds[1] = threshold1;

        bool meetsThresholds = verifier.verifyCertificateNominal(defaultOperatorSet, cert, thresholds);

        // Expected result
        bool expectedResult = (signedStakes[0] >= threshold0) && (signedStakes[1] >= threshold1);

        assertEq(meetsThresholds, expectedResult, "Nominal threshold check mismatch");
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_ViewFunctions
 * @notice Unit tests for BN254CertificateVerifier view functions
 */
contract BN254CertificateVerifierUnitTests_ViewFunctions is BN254CertificateVerifierUnitTests {
    uint32 referenceTimestamp;
    BN254OperatorSetInfo operatorSetInfo;
    OperatorSetConfig operatorSetConfig;

    function setUp() public override {
        super.setUp();
        referenceTimestamp = uint32(block.timestamp);

        (BN254OperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 3, 1);
        operatorSetInfo = _createOperatorSetInfo(operators);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operatorSetInfo, defaultOperatorSetConfig);
    }

    function test_getOperatorSetOwner() public view {
        address owner = verifier.getOperatorSetOwner(defaultOperatorSet);
        assertEq(owner, operatorSetOwner, "Operator set owner mismatch");
    }

    function test_maxOperatorTableStaleness() public view {
        uint32 staleness = verifier.maxOperatorTableStaleness(defaultOperatorSet);
        assertEq(staleness, defaultMaxStaleness, "Max staleness mismatch");
    }

    function test_latestReferenceTimestamp() public view {
        uint32 timestamp = verifier.latestReferenceTimestamp(defaultOperatorSet);
        assertEq(timestamp, referenceTimestamp, "Latest reference timestamp mismatch");
    }

    function test_getOperatorSetInfo() public view {
        BN254OperatorSetInfo memory retrievedInfo = verifier.getOperatorSetInfo(defaultOperatorSet, referenceTimestamp);

        assertEq(retrievedInfo.numOperators, operatorSetInfo.numOperators, "Num operators mismatch");
        assertEq(retrievedInfo.aggregatePubkey.X, operatorSetInfo.aggregatePubkey.X, "Aggregate pubkey X mismatch");
        assertEq(retrievedInfo.aggregatePubkey.Y, operatorSetInfo.aggregatePubkey.Y, "Aggregate pubkey Y mismatch");
        assertEq(retrievedInfo.operatorInfoTreeRoot, operatorSetInfo.operatorInfoTreeRoot, "Tree root mismatch");
        assertEq(retrievedInfo.totalWeights.length, operatorSetInfo.totalWeights.length, "Total weights length mismatch");
        for (uint i = 0; i < retrievedInfo.totalWeights.length; i++) {
            assertEq(retrievedInfo.totalWeights[i], operatorSetInfo.totalWeights[i], "Total weight mismatch");
        }
    }

    function test_getNonsignerOperatorInfo() public {
        // First cache some operator info
        (BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, 3, 1);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Verify certificate to cache operator info
        verifier.verifyCertificate(defaultOperatorSet, cert);

        // Get cached operator info
        uint32 operatorIndex = nonSignerIndices[0];
        BN254OperatorInfo memory cachedInfo = verifier.getNonsignerOperatorInfo(defaultOperatorSet, referenceTimestamp, operatorIndex);

        assertEq(cachedInfo.pubkey.X, operators[operatorIndex].pubkey.X, "Cached pubkey X mismatch");
        assertEq(cachedInfo.pubkey.Y, operators[operatorIndex].pubkey.Y, "Cached pubkey Y mismatch");
        assertEq(cachedInfo.weights[0], operators[operatorIndex].weights[0], "Cached weight 0 mismatch");
        assertEq(cachedInfo.weights[1], operators[operatorIndex].weights[1], "Cached weight 1 mismatch");
    }

    function test_isNonsignerCached() public {
        // First cache some operator info
        (BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature) =
            _createOperatorsWithSplitKeys(123, 3, 1);

        BN254Certificate memory cert = _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signature);

        // Non-signer index should be not cached prior to verification
        assertFalse(
            verifier.isNonsignerCached(defaultOperatorSet, referenceTimestamp, nonSignerIndices[0]),
            "Nonsigner operator info should not be cached"
        );

        // Verify certificate to cache operator info
        verifier.verifyCertificate(defaultOperatorSet, cert);

        // Check that the non-signer operator info is cached
        assertTrue(
            verifier.isNonsignerCached(defaultOperatorSet, referenceTimestamp, nonSignerIndices[0]),
            "Nonsigner operator info should be cached"
        );

        // Assert that the signers are NOT cached
        for (uint i = 0; i < operators.length; i++) {
            if (i != nonSignerIndices[0]) {
                assertFalse(
                    verifier.isNonsignerCached(defaultOperatorSet, referenceTimestamp, i), "Signer operator info should not be cached"
                );
            }
        }
    }

    function test_isReferenceTimestampSet_view() public view {
        // Check that the reference timestamp is set (from setUp)
        assertTrue(verifier.isReferenceTimestampSet(defaultOperatorSet, referenceTimestamp), "Reference timestamp should be set");

        // Check that a different timestamp is not set
        assertFalse(verifier.isReferenceTimestampSet(defaultOperatorSet, referenceTimestamp + 1), "Different timestamp should not be set");

        // Check for a different operator set
        OperatorSet memory differentOperatorSet = OperatorSet({avs: address(0x99), id: 10});
        assertFalse(
            verifier.isReferenceTimestampSet(differentOperatorSet, referenceTimestamp), "Different operator set timestamp should not be set"
        );
    }

    function test_getTotalStakeWeights() public view {
        // Test with valid timestamp - should return the totalWeights from operatorSetInfo
        uint[] memory totalStakes = verifier.getTotalStakeWeights(defaultOperatorSet, referenceTimestamp);

        assertEq(totalStakes.length, operatorSetInfo.totalWeights.length, "Total stakes length mismatch");
        for (uint i = 0; i < totalStakes.length; i++) {
            assertEq(totalStakes[i], operatorSetInfo.totalWeights[i], "Total stake mismatch at index");
        }

        // Test with non-existent timestamp - should return empty array
        uint32 nonExistentTimestamp = referenceTimestamp + 1000;
        uint[] memory emptyStakes = verifier.getTotalStakeWeights(defaultOperatorSet, nonExistentTimestamp);
        assertEq(emptyStakes.length, 0, "Should return empty array for non-existent timestamp");

        // Test with different operator set - should return empty array
        OperatorSet memory differentOperatorSet = OperatorSet({avs: address(0x99), id: 10});
        uint[] memory differentSetStakes = verifier.getTotalStakeWeights(differentOperatorSet, referenceTimestamp);
        assertEq(differentSetStakes.length, 0, "Should return empty array for unconfigured operator set");
    }

    function test_getOperatorCount() public view {
        // Test with valid timestamp - should return the operator count from operatorSetInfo
        uint operatorCount = verifier.getOperatorCount(defaultOperatorSet, referenceTimestamp);
        assertEq(operatorCount, operatorSetInfo.numOperators, "Operator count mismatch");

        // Test with non-existent timestamp - should return 0
        uint32 nonExistentTimestamp = referenceTimestamp + 1000;
        uint zeroCount = verifier.getOperatorCount(defaultOperatorSet, nonExistentTimestamp);
        assertEq(zeroCount, 0, "Should return 0 for non-existent timestamp");

        // Test with different operator set - should return 0
        OperatorSet memory differentOperatorSet = OperatorSet({avs: address(0x99), id: 10});
        uint differentSetCount = verifier.getOperatorCount(differentOperatorSet, referenceTimestamp);
        assertEq(differentSetCount, 0, "Should return 0 for unconfigured operator set");
    }
}

/**
 * @title BN254CertificateVerifierUnitTests_trySignatureVerification
 * @notice Unit tests for BN254CertificateVerifier.trySignatureVerification
 */
contract BN254CertificateVerifierUnitTests_trySignatureVerification is BN254CertificateVerifierUnitTests {
    function testFuzz_trySignatureVerification_validSignature(Randomness r) public rand(r) {
        // Create all operators as signers
        (BN254OperatorInfo[] memory operators,, BN254.G1Point memory validSignature) =
            _createOperatorsWithSplitKeys(r.Uint256(), numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        // Verify signature
        (bool pairingSuccessful, bool signatureValid) =
            verifier.trySignatureVerification(defaultMsgHash, operatorSetInfo.aggregatePubkey, aggSignerApkG2, validSignature);

        assertTrue(pairingSuccessful, "Pairing should be successful");
        assertTrue(signatureValid, "Signature should be valid");
    }

    function testFuzz_trySignatureVerification_invalidPairing(Randomness r) public rand(r) {
        // Create valid signature with one set of operators
        (BN254OperatorInfo[] memory operators,, BN254.G1Point memory validSignature) =
            _createOperatorsWithSplitKeys(r.Uint256(), numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        // Try to verify signature with invalid message hash
        bytes32 invalidMsgHash = bytes32(uint(1));
        (bool pairingSuccessful, bool signatureValid) =
            verifier.trySignatureVerification(invalidMsgHash, operatorSetInfo.aggregatePubkey, aggSignerApkG2, validSignature);

        assertFalse(pairingSuccessful, "Pairing should be unsuccessful");
        assertTrue(signatureValid, "Signature should be valid");
    }

    function testFuzz_trySignatureVerification_invalidSigAndPairing(Randomness r) public rand(r) {
        // Create valid signature with one set of operators
        (BN254OperatorInfo[] memory operators,, BN254.G1Point memory validSignature) =
            _createOperatorsWithSplitKeys(r.Uint256(), numOperators, 0);
        BN254OperatorSetInfo memory operatorSetInfo = _createOperatorSetInfo(operators);

        // Try to verify signature with invalid apk
        aggSignerApkG2.X[0] += 1; // Invalid apk
        (bool pairingSuccessful, bool signatureValid) =
            verifier.trySignatureVerification(defaultMsgHash, operatorSetInfo.aggregatePubkey, aggSignerApkG2, validSignature);

        assertFalse(pairingSuccessful, "Pairing should be unsuccessful");
        assertFalse(signatureValid, "Signature should be invalid");
    }
}
