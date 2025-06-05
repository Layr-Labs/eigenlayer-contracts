// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";

contract BN254CertificateVerifierTest is Test {
    using BN254 for BN254.G1Point;

    // Contract being tested
    BN254CertificateVerifier verifier;

    // Test accounts
    address owner = address(0x1);
    address tableUpdater = address(0x2);
    address nonOwner = address(0x3);
    address operatorSetOwner = address(0x4);

    // Test data
    uint32 numOperators = 4;
    uint32 maxStaleness = 3600; // 1 hour max staleness

    // Create an OperatorSet for testing
    OperatorSet testOperatorSet;

    // BLS signature specific fields
    bytes32 msgHash;
    uint aggSignerPrivKey = 69;
    BN254.G2Point aggSignerApkG2; // G2 public key corresponding to aggSignerPrivKey

    // Events
    event TableUpdated(
        OperatorSet indexed operatorSet, uint32 referenceTimestamp, IBN254TableCalculatorTypes.BN254OperatorSetInfo operatorSetInfo
    );

    function setUp() public virtual {
        vm.warp(1_000_000); // Set block timestamp

        testOperatorSet.avs = address(0x5);
        testOperatorSet.id = 1;

        // Deploy implementation
        BN254CertificateVerifier implementation = new BN254CertificateVerifier(tableUpdater);

        // Deploy proxy and initialize
        ERC1967Proxy proxy =
            new ERC1967Proxy(address(implementation), abi.encodeWithSelector(BN254CertificateVerifier.initialize.selector, owner));

        verifier = BN254CertificateVerifier(address(proxy));

        // Set standard test message hash
        msgHash = keccak256(abi.encodePacked("test message"));

        // Set up the aggregate public key in G2
        aggSignerApkG2.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
        aggSignerApkG2.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
        aggSignerApkG2.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
        aggSignerApkG2.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
    }

    // Generate signer and non-signer private keys
    function generateSignerAndNonSignerPrivateKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        view
        returns (uint[] memory, uint[] memory)
    {
        uint[] memory signerPrivKeys = new uint[](numSigners);
        uint sum = 0;

        // Generate numSigners-1 random keys
        for (uint i = 0; i < numSigners - 1; i++) {
            signerPrivKeys[i] = uint(keccak256(abi.encodePacked("signerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
            sum = addmod(sum, signerPrivKeys[i], BN254.FR_MODULUS);
        }

        // Last key makes the total sum equal to aggSignerPrivKey
        signerPrivKeys[numSigners - 1] = addmod(aggSignerPrivKey, BN254.FR_MODULUS - sum % BN254.FR_MODULUS, BN254.FR_MODULUS);

        // Generate non-signer keys
        uint[] memory nonSignerPrivKeys = new uint[](numNonSigners);
        for (uint i = 0; i < numNonSigners; i++) {
            nonSignerPrivKeys[i] = uint(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, i))) % BN254.FR_MODULUS;
        }

        // Sort nonSignerPrivateKeys in order of ascending pubkeyHash
        for (uint i = 1; i < nonSignerPrivKeys.length; i++) {
            uint privateKey = nonSignerPrivKeys[i];
            bytes32 pubkeyHash = toPubkeyHash(privateKey);
            uint j = i;

            // Move elements that are greater than the current key ahead
            while (j > 0 && toPubkeyHash(nonSignerPrivKeys[j - 1]) > pubkeyHash) {
                nonSignerPrivKeys[j] = nonSignerPrivKeys[j - 1];
                j--;
            }
            nonSignerPrivKeys[j] = privateKey;
        }

        return (signerPrivKeys, nonSignerPrivKeys);
    }

    // Helper to hash a public key
    function toPubkeyHash(uint privKey) internal view returns (bytes32) {
        return BN254.generatorG1().scalar_mul(privKey).hashG1Point();
    }

    // Create operators with split keys
    function createOperatorsWithSplitKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        view
        returns (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory, uint32[] memory, BN254.G1Point memory)
    {
        require(numSigners + numNonSigners == numOperators, "Total operators mismatch");

        // Generate private keys
        (uint[] memory signerPrivKeys, uint[] memory nonSignerPrivKeys) =
            generateSignerAndNonSignerPrivateKeys(pseudoRandomNumber, numSigners, numNonSigners);

        // Create all operators
        IBN254TableCalculatorTypes.BN254OperatorInfo[] memory ops = new IBN254TableCalculatorTypes.BN254OperatorInfo[](numOperators);

        // Track indices of non-signers
        uint32[] memory nonSignerIndices = new uint32[](numNonSigners);

        // Create signers first
        for (uint32 i = 0; i < numSigners; i++) {
            ops[i].pubkey = BN254.generatorG1().scalar_mul(signerPrivKeys[i]);
            ops[i].weights = new uint[](2);
            ops[i].weights[0] = uint(100 + i * 10);
            ops[i].weights[1] = uint(200 + i * 20);
        }

        // Create non-signers
        for (uint32 i = 0; i < numNonSigners; i++) {
            uint32 idx = uint32(numSigners + i);
            ops[idx].pubkey = BN254.generatorG1().scalar_mul(nonSignerPrivKeys[i]);
            ops[idx].weights = new uint[](2);
            ops[idx].weights[0] = uint(100 + idx * 10);
            ops[idx].weights[1] = uint(200 + idx * 20);
            nonSignerIndices[i] = idx;
        }

        // Calculate aggregate signature for the signers
        BN254.G1Point memory signature = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);

        return (ops, nonSignerIndices, signature);
    }

    // Build a complete and correct merkle tree from operator infos
    function getMerkleRoot(IBN254TableCalculatorTypes.BN254OperatorInfo[] memory ops) internal returns (bytes32 root) {
        bytes32[] memory leaves = new bytes32[](ops.length);
        for (uint i = 0; i < ops.length; i++) {
            leaves[i] = keccak256(abi.encode(ops[i]));
        }
        root = Merkle.merkleizeKeccak(leaves);
    }

    function getMerkleProof(IBN254TableCalculatorTypes.BN254OperatorInfo[] memory ops, uint32 operatorIndex)
        internal
        returns (bytes memory proof)
    {
        bytes32[] memory leaves = new bytes32[](ops.length);
        for (uint i = 0; i < ops.length; i++) {
            leaves[i] = keccak256(abi.encode(ops[i]));
        }
        proof = Merkle.getProofKeccak(leaves, operatorIndex);
    }

    // Create operator set info
    function createOperatorSetInfo(IBN254TableCalculatorTypes.BN254OperatorInfo[] memory ops)
        internal
        returns (IBN254TableCalculatorTypes.BN254OperatorSetInfo memory)
    {
        uint32 _numOperators = uint32(ops.length);
        bytes32 operatorInfoTreeRoot = getMerkleRoot(ops);

        // Create aggregate public key (sum of all operator pubkeys)
        BN254.G1Point memory aggregatePubkey = BN254.G1Point(0, 0);

        // Create total weights (sum of all operator weights)
        uint[] memory _totalWeights = new uint[](2);

        for (uint32 i = 0; i < _numOperators; i++) {
            // Add pubkey to aggregate
            aggregatePubkey = aggregatePubkey.plus(ops[i].pubkey);

            // Add weights to total
            for (uint j = 0; j < 2; j++) {
                _totalWeights[j] += ops[i].weights[j];
            }
        }

        // Create the operator set info
        return IBN254TableCalculatorTypes.BN254OperatorSetInfo({
            numOperators: _numOperators,
            aggregatePubkey: aggregatePubkey,
            totalWeights: _totalWeights,
            operatorInfoTreeRoot: operatorInfoTreeRoot
        });
    }

    // Helper to create a certificate with real BLS signature
    function createCertificate(
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint32[] memory nonSignerIndices,
        IBN254TableCalculatorTypes.BN254OperatorInfo[] memory ops,
        BN254.G1Point memory signature
    ) internal returns (IBN254CertificateVerifierTypes.BN254Certificate memory) {
        // Create witnesses for non-signers
        IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[] memory witnesses =
            new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](nonSignerIndices.length);

        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            bytes memory proof = getMerkleProof(ops, nonSignerIndex);

            witnesses[i] = IBN254CertificateVerifierTypes.BN254OperatorInfoWitness({
                operatorIndex: nonSignerIndex,
                operatorInfoProof: proof,
                operatorInfo: ops[nonSignerIndex]
            });
        }

        return IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: witnesses
        });
    }

    // Helper to create operator set config
    function createOperatorSetConfig() internal view returns (ICrossChainRegistryTypes.OperatorSetConfig memory) {
        return ICrossChainRegistryTypes.OperatorSetConfig({owner: operatorSetOwner, maxStalenessPeriod: maxStaleness});
    }

    // Test updating the operator table
    function testUpdateOperatorTable() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators,,) = createOperatorsWithSplitKeys(123, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.startPrank(tableUpdater);

        // Update the operator table
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        vm.stopPrank();

        // Verify storage updates
        assertEq(verifier.latestReferenceTimestamp(testOperatorSet), referenceTimestamp, "Reference timestamp not updated correctly");
        assertEq(verifier.getOperatorSetOwner(testOperatorSet), operatorSetOwner, "Operator set owner not stored correctly");
        assertEq(verifier.maxOperatorTableStaleness(testOperatorSet), maxStaleness, "Max staleness not stored correctly");

        // Verify operator set info was stored correctly
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory storedOperatorSetInfo =
            verifier.getOperatorSetInfo(testOperatorSet, referenceTimestamp);

        assertEq(storedOperatorSetInfo.numOperators, operatorSetInfo.numOperators, "Num operators not stored correctly");
        assertEq(storedOperatorSetInfo.aggregatePubkey.X, operatorSetInfo.aggregatePubkey.X, "Aggregate pubkey X not stored correctly");
        assertEq(storedOperatorSetInfo.aggregatePubkey.Y, operatorSetInfo.aggregatePubkey.Y, "Aggregate pubkey Y not stored correctly");
        assertEq(storedOperatorSetInfo.operatorInfoTreeRoot, operatorSetInfo.operatorInfoTreeRoot, "Tree root not stored correctly");
    }

    // Test verifyCertificate with actual BLS signature validation and multiple signers
    function testVerifyCertificate() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 123;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);

        // Check that the signed stakes are correct
        assertEq(signedStakes.length, 2, "Wrong number of stake types");

        // Calculate expected signed stakes
        uint[] memory expectedSignedStakes = new uint[](2);

        // Start with total stakes
        expectedSignedStakes[0] = operatorSetInfo.totalWeights[0];
        expectedSignedStakes[1] = operatorSetInfo.totalWeights[1];

        // Subtract non-signer stakes
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            expectedSignedStakes[0] -= operators[nonSignerIndex].weights[0];
            expectedSignedStakes[1] -= operators[nonSignerIndex].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");
    }

    // Test verifyCertificate with a different distribution of signers/non-signers
    function testVerifyCertificateDifferentSplit() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 2 signers, 2 non-signers
        uint pseudoRandomNumber = 456;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 2, 2);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // No mocking needed - verify certificate with real verification
        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);

        // Calculate expected signed stakes
        uint[] memory expectedSignedStakes = new uint[](2);

        // Start with total stakes
        expectedSignedStakes[0] = operatorSetInfo.totalWeights[0];
        expectedSignedStakes[1] = operatorSetInfo.totalWeights[1];

        // Subtract non-signer stakes
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            expectedSignedStakes[0] -= operators[nonSignerIndex].weights[0];
            expectedSignedStakes[1] -= operators[nonSignerIndex].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");
    }

    // Test verifyCertificateProportion
    function testVerifyCertificateProportion() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 789;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // Set thresholds at 60% of total stake for each type
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 6000; // 60%
        thresholds[1] = 6000; // 60%

        // Verify certificate meets thresholds - no mocking needed
        bool meetsThresholds = verifier.verifyCertificateProportion(testOperatorSet, cert, thresholds);

        // With 3 signers out of 4, should meet 60% threshold
        assertTrue(meetsThresholds, "Certificate should meet thresholds");

        // Try with higher threshold that shouldn't be met
        thresholds[0] = 9000; // 90%
        thresholds[1] = 9000; // 90%

        meetsThresholds = verifier.verifyCertificateProportion(testOperatorSet, cert, thresholds);

        // Calculate percentage of signed stakes to determine if it should meet threshold
        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);
        uint signedPercentage0 = (signedStakes[0] * 10_000) / operatorSetInfo.totalWeights[0];
        uint signedPercentage1 = (signedStakes[1] * 10_000) / operatorSetInfo.totalWeights[1];

        bool shouldMeetThreshold = (signedPercentage0 >= 9000) && (signedPercentage1 >= 9000);
        assertEq(meetsThresholds, shouldMeetThreshold, "Certificate threshold check incorrect");
    }

    // Test with invalid signature (wrong message hash)
    function testVerifyCertificateInvalidSignature() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 123;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature but WRONG message hash
        bytes32 wrongHash = keccak256("wrong message");

        IBN254CertificateVerifierTypes.BN254Certificate memory cert = createCertificate(
            referenceTimestamp,
            wrongHash, // Use wrong hash here
            nonSignerIndices,
            operators,
            signature // Signature is for original msgHash, not wrongHash
        );

        // Verification should fail without mocking
        vm.expectRevert(abi.encodeWithSignature("VerificationFailed()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test certificate with stale timestamp (should fail)
    function testVerifyCertificateStaleTimestamp() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 123;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // Jump forward in time beyond the max staleness
        vm.warp(block.timestamp + maxStaleness + 1);

        // Verification should fail due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test verifyCertificateNominal functionality
    function testVerifyCertificateNominal() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 567;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // Get the signed stakes first
        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);

        // Test with thresholds lower than signed stakes (should pass)
        uint[] memory passThresholds = new uint[](2);
        passThresholds[0] = signedStakes[0] - 10;
        passThresholds[1] = signedStakes[1] - 10;

        bool meetsThresholds = verifier.verifyCertificateNominal(testOperatorSet, cert, passThresholds);
        assertTrue(meetsThresholds, "Certificate should meet nominal thresholds");

        // Test with thresholds higher than signed stakes (should fail)
        uint[] memory failThresholds = new uint[](2);
        failThresholds[0] = signedStakes[0] + 10;
        failThresholds[1] = signedStakes[1] + 10;

        meetsThresholds = verifier.verifyCertificateNominal(testOperatorSet, cert, failThresholds);
        assertFalse(meetsThresholds, "Certificate should not meet impossible nominal thresholds");
    }

    // Test with no non-signers (all operators sign)
    function testVerifyCertificateAllSigners() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with all signers
        uint pseudoRandomNumber = 999;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators,, BN254.G1Point memory signature) =
            createOperatorsWithSplitKeys(pseudoRandomNumber, numOperators, 0);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with no non-signers
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: msgHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Verify certificate
        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);

        // All stakes should be signed
        assertEq(signedStakes[0], operatorSetInfo.totalWeights[0], "All stake should be signed for type 0");
        assertEq(signedStakes[1], operatorSetInfo.totalWeights[1], "All stake should be signed for type 1");
    }

    // Test with invalid operator tree root
    function testVerifyCertificateInvalidOperatorTreeRoot() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 123;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info with CORRECT tree root
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        // Modify the tree root to be invalid AFTER creating the info
        operatorSetInfo.operatorInfoTreeRoot = keccak256("invalid root");

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature but proofs won't match the invalid root
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // Verification should fail due to invalid merkle proofs
        vm.expectRevert(abi.encodeWithSignature("VerificationFailed()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test with invalid reference timestamp
    function testInvalidReferenceTimestamp() public {
        uint32 existingReferenceTimestamp = uint32(block.timestamp);
        uint32 nonExistentTimestamp = existingReferenceTimestamp + 1000;

        // Create operators with split keys - 3 signers, 1 non-signer
        uint pseudoRandomNumber = 123;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 3, 1);

        // Create operator set info
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, existingReferenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate using a non-existent timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: nonExistentTimestamp,
            messageHash: msgHash,
            signature: signature,
            apk: aggSignerApkG2,
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Verification should fail due to timestamp not existing
        vm.expectRevert(abi.encodeWithSignature("ReferenceTimestampDoesNotExist()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test multiple operator sets
    function testMultipleOperatorSets() public {
        // Create two different operator sets
        OperatorSet memory operatorSet1 = OperatorSet({avs: address(0x10), id: 1});
        OperatorSet memory operatorSet2 = OperatorSet({avs: address(0x20), id: 2});

        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators for each set
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators1,,) = createOperatorsWithSplitKeys(111, 3, 1);
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators2,,) = createOperatorsWithSplitKeys(222, 2, 2);

        // Create operator set infos
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo1 = createOperatorSetInfo(operators1);
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo2 = createOperatorSetInfo(operators2);

        // Create operator set configs with different owners
        ICrossChainRegistryTypes.OperatorSetConfig memory config1 = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(0x100),
            maxStalenessPeriod: 1800 // 30 minutes
        });
        ICrossChainRegistryTypes.OperatorSetConfig memory config2 = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(0x200),
            maxStalenessPeriod: 7200 // 2 hours
        });

        vm.startPrank(tableUpdater);

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
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory storedInfo1 = verifier.getOperatorSetInfo(operatorSet1, referenceTimestamp);
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory storedInfo2 = verifier.getOperatorSetInfo(operatorSet2, referenceTimestamp);

        assertEq(storedInfo1.numOperators, 4, "OperatorSet1 should have 4 operators");
        assertEq(storedInfo2.numOperators, 4, "OperatorSet2 should have 4 operators");

        // Verify they have different tree roots (since operators are different)
        assertTrue(storedInfo1.operatorInfoTreeRoot != storedInfo2.operatorInfoTreeRoot, "Tree roots should be different");
    }

    // Test access control for operator table updates
    function testOperatorTableUpdateAccessControl() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators,,) = createOperatorsWithSplitKeys(123, 3, 1);

        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        // Try to update with non-authorized account
        vm.prank(nonOwner);
        vm.expectRevert(abi.encodeWithSignature("OnlyTableUpdater()"));
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Should succeed with authorized account
        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);
    }

    // Test stale table update (timestamp not increasing)
    function testStaleTableUpdate() public {
        uint32 firstTimestamp = uint32(block.timestamp);
        uint32 staleTimestamp = firstTimestamp - 100; // Earlier timestamp

        // Create operators
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators,,) = createOperatorsWithSplitKeys(123, 3, 1);

        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.startPrank(tableUpdater);

        // First update should succeed
        verifier.updateOperatorTable(testOperatorSet, firstTimestamp, operatorSetInfo, operatorSetConfig);

        // Second update with earlier timestamp should fail
        vm.expectRevert(abi.encodeWithSignature("TableUpdateStale()"));
        verifier.updateOperatorTable(testOperatorSet, staleTimestamp, operatorSetInfo, operatorSetConfig);

        vm.stopPrank();
    }

    // Test operator info caching
    function testOperatorInfoCaching() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators with split keys - 2 signers, 2 non-signers
        uint pseudoRandomNumber = 456;
        (IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operators, uint32[] memory nonSignerIndices, BN254.G1Point memory signature)
        = createOperatorsWithSplitKeys(pseudoRandomNumber, 2, 2);

        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = createOperatorSetInfo(operators);
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);

        // Create certificate with real BLS signature
        IBN254CertificateVerifierTypes.BN254Certificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSignerIndices, operators, signature);

        // First verification should cache the operator infos
        verifier.verifyCertificate(testOperatorSet, cert);

        // Check that operator infos are now cached
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            IBN254TableCalculatorTypes.BN254OperatorInfo memory cachedInfo =
                verifier.getOperatorInfo(testOperatorSet, referenceTimestamp, nonSignerIndex);

            // Cached info should match original
            assertEq(cachedInfo.pubkey.X, operators[nonSignerIndex].pubkey.X, "Cached pubkey X mismatch");
            assertEq(cachedInfo.pubkey.Y, operators[nonSignerIndex].pubkey.Y, "Cached pubkey Y mismatch");
            assertEq(cachedInfo.weights[0], operators[nonSignerIndex].weights[0], "Cached weight 0 mismatch");
            assertEq(cachedInfo.weights[1], operators[nonSignerIndex].weights[1], "Cached weight 1 mismatch");
        }

        // Second verification should use cached data (should still work)
        verifier.verifyCertificate(testOperatorSet, cert);
    }
}
