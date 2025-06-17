// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "forge-std/Test.sol";

import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/multichain/ECDSACertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";

/**
 * @title ECDSACertificateVerifierTest
 * @notice Unit tests for ECDSACertificateVerifier contract
 */
contract ECDSACertificateVerifierTest is Test {
    // Contract being tested
    ECDSACertificateVerifier verifier;

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

    // ECDSA signature specific fields
    bytes32 msgHash;
    uint signerPrivateKey = 0x1234;
    address signerAddress;

    // Events
    event TableUpdated(
        OperatorSet indexed operatorSet, uint32 referenceTimestamp, IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] operatorInfos
    );

    function setUp() public virtual {
        vm.warp(1_000_000); // Set block timestamp

        testOperatorSet.avs = address(0x5);
        testOperatorSet.id = 1;

        // Deploy implementation
        ECDSACertificateVerifier implementation = new ECDSACertificateVerifier(IOperatorTableUpdater(tableUpdater), "1.0.0");

        // Deploy proxy and initialize
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), "");

        verifier = ECDSACertificateVerifier(address(proxy));

        // Set standard test message hash
        msgHash = keccak256(abi.encodePacked("test message"));

        // Generate signer address from private key
        signerAddress = vm.addr(signerPrivateKey);
    }

    // Helper to create operator infos with specified weights
    function createOperatorInfos(uint numSigners, uint numNonSigners)
        internal
        view
        returns (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory,
            address[] memory,
            uint[] memory // signer private keys
        )
    {
        require(numSigners + numNonSigners == numOperators, "Total operators mismatch");

        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory ops = new IECDSACertificateVerifierTypes.ECDSAOperatorInfo[](numOperators);
        address[] memory nonSigners = new address[](numNonSigners);
        uint[] memory signerPrivKeys = new uint[](numSigners);

        // Create signers
        for (uint i = 0; i < numSigners; i++) {
            uint privateKey = uint(keccak256(abi.encodePacked("signer", i)));
            address pubkey = vm.addr(privateKey);

            ops[i].pubkey = pubkey;
            ops[i].weights = new uint[](2);
            ops[i].weights[0] = uint(100 + i * 10);
            ops[i].weights[1] = uint(200 + i * 20);
            signerPrivKeys[i] = privateKey;
        }

        // Create non-signers
        for (uint i = 0; i < numNonSigners; i++) {
            uint idx = numSigners + i;
            uint privateKey = uint(keccak256(abi.encodePacked("nonsigner", i)));
            address pubkey = vm.addr(privateKey);

            ops[idx].pubkey = pubkey;
            ops[idx].weights = new uint[](2);
            ops[idx].weights[0] = uint(100 + idx * 10);
            ops[idx].weights[1] = uint(200 + idx * 20);
            nonSigners[i] = pubkey;
        }

        return (ops, nonSigners, signerPrivKeys);
    }

    // Helper to create operator set config
    function createOperatorSetConfig() internal view returns (ICrossChainRegistryTypes.OperatorSetConfig memory) {
        return ICrossChainRegistryTypes.OperatorSetConfig({owner: operatorSetOwner, maxStalenessPeriod: maxStaleness});
    }

    // Helper function to mirror _calculateSignableDigest
    function calculateSignableDigest(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", verifier.domainSeparator(), structHash));
    }

    // Helper to create a certificate with real ECDSA signatures
    function createCertificate(
        uint32 referenceTimestamp,
        bytes32 messageHash,
        address[] memory nonSigners,
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory ops,
        uint[] memory signerPrivKeys
    ) internal view returns (IECDSACertificateVerifierTypes.ECDSACertificate memory) {
        // Use the contract's digest calculation
        bytes32 signableDigest = verifier.calculateCertificateDigest(referenceTimestamp, messageHash);

        // Collect signers and their private keys
        uint numSigners = signerPrivKeys.length;
        address[] memory signerAddresses = new address[](numSigners);
        bytes[] memory signaturesArr = new bytes[](numSigners);
        uint signerIdx = 0;
        for (uint i = 0; i < ops.length; i++) {
            bool isNonSigner = false;
            for (uint j = 0; j < nonSigners.length; j++) {
                if (ops[i].pubkey == nonSigners[j]) {
                    isNonSigner = true;
                    break;
                }
            }
            if (!isNonSigner) {
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPrivKeys[signerIdx], signableDigest);
                bytes memory signature = abi.encodePacked(r, s, v);
                signerAddresses[signerIdx] = ops[i].pubkey;
                signaturesArr[signerIdx] = signature;
                signerIdx++;
            }
        }
        // Sort signers and signatures by address (ascending)
        for (uint i = 0; i < numSigners; i++) {
            for (uint j = i + 1; j < numSigners; j++) {
                if (signerAddresses[j] < signerAddresses[i]) {
                    // Swap addresses
                    address tmpAddr = signerAddresses[i];
                    signerAddresses[i] = signerAddresses[j];
                    signerAddresses[j] = tmpAddr;
                    // Swap signatures
                    bytes memory tmpSig = signaturesArr[i];
                    signaturesArr[i] = signaturesArr[j];
                    signaturesArr[j] = tmpSig;
                }
            }
        }
        // Concatenate signatures in sorted order
        bytes memory signatures;
        for (uint i = 0; i < numSigners; i++) {
            signatures = bytes.concat(signatures, signaturesArr[i]);
        }
        return IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            sig: signatures
        });
    }

    // Test updating the operator table
    function testUpdateOperatorTable() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.startPrank(tableUpdater);

        // Update the operator table
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        vm.stopPrank();

        // Verify storage updates
        assertEq(verifier.latestReferenceTimestamp(testOperatorSet), referenceTimestamp, "Reference timestamp not updated correctly");
        assertEq(verifier.getOperatorSetOwner(testOperatorSet), operatorSetOwner, "Operator set owner not stored correctly");
        assertEq(verifier.maxOperatorTableStaleness(testOperatorSet), maxStaleness, "Max staleness not stored correctly");

        // Verify operator infos were stored correctly
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOperators =
            verifier.getOperatorInfos(testOperatorSet, referenceTimestamp);

        assertEq(storedOperators.length, operators.length, "Number of operators not stored correctly");
        for (uint i = 0; i < operators.length; i++) {
            assertEq(storedOperators[i].pubkey, operators[i].pubkey, "Operator pubkey not stored correctly");
            assertEq(storedOperators[i].weights[0], operators[i].weights[0], "Operator weight 0 not stored correctly");
            assertEq(storedOperators[i].weights[1], operators[i].weights[1], "Operator weight 1 not stored correctly");
        }
    }

    // Test verifyCertificate with actual ECDSA signature validation
    function testVerifyCertificate() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Create certificate with real ECDSA signatures
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSigners, operators, signerPrivKeys);

        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);

        // Check that the signed stakes are correct
        assertEq(signedStakes.length, 2, "Wrong number of stake types");

        // Calculate expected signed stakes
        uint[] memory expectedSignedStakes = new uint[](2);

        // Start with total stakes
        for (uint i = 0; i < operators.length; i++) {
            expectedSignedStakes[0] += operators[i].weights[0];
            expectedSignedStakes[1] += operators[i].weights[1];
        }

        // Subtract non-signer stakes
        for (uint i = 0; i < nonSigners.length; i++) {
            for (uint j = 0; j < operators.length; j++) {
                if (operators[j].pubkey == nonSigners[i]) {
                    expectedSignedStakes[0] -= operators[j].weights[0];
                    expectedSignedStakes[1] -= operators[j].weights[1];
                    break;
                }
            }
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");
    }

    // Test verifyCertificateProportion
    function testVerifyCertificateProportion() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Create certificate with real ECDSA signatures
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSigners, operators, signerPrivKeys);

        // Set thresholds at 60% of total stake for each type
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 6000; // 60%
        thresholds[1] = 6000; // 60%

        // Verify certificate meets thresholds
        bool meetsThresholds = verifier.verifyCertificateProportion(testOperatorSet, cert, thresholds);

        // With 3 signers out of 4, should meet 60% threshold
        assertTrue(meetsThresholds, "Certificate should meet thresholds");

        // Try with higher threshold that shouldn't be met
        thresholds[0] = 9000; // 90%
        thresholds[1] = 9000; // 90%

        meetsThresholds = verifier.verifyCertificateProportion(testOperatorSet, cert, thresholds);

        // Calculate percentage of signed stakes to determine if it should meet threshold
        uint[] memory signedStakes = verifier.verifyCertificate(testOperatorSet, cert);
        uint totalStake0 = 0;
        uint totalStake1 = 0;
        for (uint i = 0; i < operators.length; i++) {
            totalStake0 += operators[i].weights[0];
            totalStake1 += operators[i].weights[1];
        }
        uint signedPercentage0 = (signedStakes[0] * 10_000) / totalStake0;
        uint signedPercentage1 = (signedStakes[1] * 10_000) / totalStake1;

        bool shouldMeetThreshold = (signedPercentage0 >= 9000) && (signedPercentage1 >= 9000);
        assertEq(meetsThresholds, shouldMeetThreshold, "Certificate threshold check incorrect");
    }

    // Test verifyCertificateNominal
    function testVerifyCertificateNominal() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Create certificate with real ECDSA signatures
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSigners, operators, signerPrivKeys);

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

    // Test with invalid signature (wrong message hash)
    function testVerifyCertificateInvalidSignature() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Create a certificate with signatures for the original message hash
        bytes32 originalHash = msgHash;
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(referenceTimestamp, originalHash, nonSigners, operators, signerPrivKeys);

        // Now modify the certificate to use a different message hash
        // This should make the signatures invalid since they were created for originalHash
        bytes32 differentHash = keccak256("different message");
        cert.messageHash = differentHash;

        // Verification should fail because the signatures were created for originalHash
        // but we're trying to verify them against differentHash
        vm.expectRevert(abi.encodeWithSignature("VerificationFailed()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test certificate with stale timestamp (should fail)
    function testVerifyCertificateStaleTimestamp() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Create certificate with real ECDSA signatures
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(referenceTimestamp, msgHash, nonSigners, operators, signerPrivKeys);

        // Jump forward in time beyond the max staleness
        vm.warp(block.timestamp + maxStaleness + 1);

        // Verification should fail due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        verifier.verifyCertificate(testOperatorSet, cert);
    }

    // Test with invalid reference timestamp
    function testInvalidReferenceTimestamp() public {
        uint32 existingReferenceTimestamp = uint32(block.timestamp);
        uint32 nonExistentTimestamp = existingReferenceTimestamp + 1000;

        // Create operators - 3 signers, 1 non-signer
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        // Create operator set config
        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, existingReferenceTimestamp, operators, operatorSetConfig);

        // Create certificate using a non-existent timestamp
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            createCertificate(nonExistentTimestamp, msgHash, nonSigners, operators, signerPrivKeys);

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

        // Create different operators for each set
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators1, address[] memory nonSigners1, uint[] memory signerPrivKeys1)
        = createOperatorInfos(3, 1);

        // Create second set of operators with different private keys
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators2 = new IECDSACertificateVerifierTypes.ECDSAOperatorInfo[](4);
        address[] memory nonSigners2 = new address[](2);
        uint[] memory signerPrivKeys2 = new uint[](2);

        for (uint i = 0; i < 4; i++) {
            uint privateKey = uint(keccak256(abi.encodePacked("set2_signer", i)));
            address pubkey = vm.addr(privateKey);

            operators2[i].pubkey = pubkey;
            operators2[i].weights = new uint[](2);
            operators2[i].weights[0] = uint(100 + i * 10);
            operators2[i].weights[1] = uint(200 + i * 20);

            if (i >= 2) signerPrivKeys2[i - 2] = privateKey;
            else nonSigners2[i] = pubkey;
        }

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
        verifier.updateOperatorTable(operatorSet1, referenceTimestamp, operators1, config1);
        verifier.updateOperatorTable(operatorSet2, referenceTimestamp, operators2, config2);

        vm.stopPrank();

        // Verify that both operator sets are stored correctly and independently
        assertEq(verifier.getOperatorSetOwner(operatorSet1), address(0x100), "OperatorSet1 owner incorrect");
        assertEq(verifier.getOperatorSetOwner(operatorSet2), address(0x200), "OperatorSet2 owner incorrect");

        assertEq(verifier.maxOperatorTableStaleness(operatorSet1), 1800, "OperatorSet1 staleness incorrect");
        assertEq(verifier.maxOperatorTableStaleness(operatorSet2), 7200, "OperatorSet2 staleness incorrect");

        assertEq(verifier.latestReferenceTimestamp(operatorSet1), referenceTimestamp, "OperatorSet1 timestamp incorrect");
        assertEq(verifier.latestReferenceTimestamp(operatorSet2), referenceTimestamp, "OperatorSet2 timestamp incorrect");

        // Verify operator infos are stored independently
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOps1 = verifier.getOperatorInfos(operatorSet1, referenceTimestamp);
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOps2 = verifier.getOperatorInfos(operatorSet2, referenceTimestamp);

        assertEq(storedOps1.length, 4, "OperatorSet1 should have 4 operators");
        assertEq(storedOps2.length, 4, "OperatorSet2 should have 4 operators");

        // Verify they have different operators
        assertTrue(storedOps1[0].pubkey != storedOps2[0].pubkey, "Operators should be different");
    }

    // Test access control for operator table updates
    function testOperatorTableUpdateAccessControl() public {
        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        // Try to update with non-authorized account
        vm.prank(nonOwner);
        vm.expectRevert(abi.encodeWithSignature("OnlyTableUpdater()"));
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);

        // Should succeed with authorized account
        vm.prank(tableUpdater);
        verifier.updateOperatorTable(testOperatorSet, referenceTimestamp, operators, operatorSetConfig);
    }

    // Test stale table update (timestamp not increasing)
    function testStaleTableUpdate() public {
        uint32 firstTimestamp = uint32(block.timestamp);
        uint32 staleTimestamp = firstTimestamp - 100; // Earlier timestamp

        // Create operators
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators, address[] memory nonSigners, uint[] memory signerPrivKeys) =
            createOperatorInfos(3, 1);

        ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig = createOperatorSetConfig();

        vm.startPrank(tableUpdater);

        // First update should succeed
        verifier.updateOperatorTable(testOperatorSet, firstTimestamp, operators, operatorSetConfig);

        // Second update with earlier timestamp should fail
        vm.expectRevert(abi.encodeWithSignature("TableUpdateStale()"));
        verifier.updateOperatorTable(testOperatorSet, staleTimestamp, operators, operatorSetConfig);

        vm.stopPrank();
    }
}
