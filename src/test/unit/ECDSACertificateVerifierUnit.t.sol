// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "forge-std/Test.sol";

import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/multichain/ECDSACertificateVerifier.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/test/utils/Random.sol";

/**
 * @title ECDSACertificateVerifierUnitTests
 * @notice Base contract for all ECDSACertificateVerifier unit tests
 */
contract ECDSACertificateVerifierUnitTests is
    EigenLayerMultichainUnitTestSetup,
    IECDSACertificateVerifierTypes,
    IECDSACertificateVerifierEvents,
    IBaseCertificateVerifierErrors
{
    using OperatorSetLib for OperatorSet;

    // Contracts
    ECDSACertificateVerifier ecdsaCertificateVerifierImplementation;
    ECDSACertificateVerifier verifier;

    // Test accounts
    address operatorSetOwner = address(0x4);

    // Defaults
    uint32 numOperators = 4;
    uint32 defaultMaxStaleness = 3600; // 1 hour max staleness
    OperatorSet defaultOperatorSet;
    ICrossChainRegistryTypes.OperatorSetConfig defaultOperatorSetConfig;

    // ECDSA signature specific fields
    bytes32 defaultMsgHash = keccak256(abi.encodePacked("test message"));

    function setUp() public virtual override {
        super.setUp();

        defaultOperatorSet = OperatorSet({avs: address(0x5), id: 0});
        defaultOperatorSetConfig =
            ICrossChainRegistryTypes.OperatorSetConfig({owner: operatorSetOwner, maxStalenessPeriod: defaultMaxStaleness});

        // Deploy implementation
        ecdsaCertificateVerifierImplementation =
            new ECDSACertificateVerifier(IOperatorTableUpdater(address(operatorTableUpdaterMock)), "1.0.0");

        // Deploy proxy and initialize
        verifier = ECDSACertificateVerifier(
            address(new TransparentUpgradeableProxy(address(ecdsaCertificateVerifierImplementation), address(eigenLayerProxyAdmin), ""))
        );
    }

    // Helper functions

    /**
     * @notice Generate signer and non-signer private keys
     * @param pseudoRandomNumber Pseudo random number for generating keys
     * @param numSigners Number of signers to generate
     * @param numNonSigners Number of non-signers to generate
     * @return signerPrivKeys Array of signer private keys
     * @return nonSignerPrivKeys Array of non-signer private keys
     */
    function _generateSignerAndNonSignerPrivateKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        pure
        returns (uint[] memory signerPrivKeys, uint[] memory nonSignerPrivKeys)
    {
        signerPrivKeys = new uint[](numSigners);
        nonSignerPrivKeys = new uint[](numNonSigners);

        // Generate signer keys
        for (uint i = 0; i < numSigners; i++) {
            signerPrivKeys[i] = uint(keccak256(abi.encodePacked("signerPrivateKey", pseudoRandomNumber, i))) % MAX_PRIVATE_KEY;
            if (signerPrivKeys[i] == 0) signerPrivKeys[i] = 1; // Ensure non-zero
        }

        // Generate non-signer keys
        for (uint i = 0; i < numNonSigners; i++) {
            nonSignerPrivKeys[i] = uint(keccak256(abi.encodePacked("nonSignerPrivateKey", pseudoRandomNumber, i))) % MAX_PRIVATE_KEY;
            if (nonSignerPrivKeys[i] == 0) nonSignerPrivKeys[i] = 1; // Ensure non-zero
        }
    }

    /**
     * @notice Create operators with split keys
     * @param pseudoRandomNumber Pseudo random number for generating operator data
     * @param numSigners Number of signers
     * @param numNonSigners Number of non-signers
     * @return operators Array of operator infos
     * @return nonSignerIndices Array of non-signer indices
     * @return signerPrivKeys Array of signer private keys
     */
    function _createOperatorsWithSplitKeys(uint pseudoRandomNumber, uint numSigners, uint numNonSigners)
        internal
        view
        returns (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        )
    {
        require(numSigners + numNonSigners == numOperators, "Total operators mismatch");

        // Generate private keys
        uint[] memory nonSignerPrivKeys;
        (signerPrivKeys, nonSignerPrivKeys) = _generateSignerAndNonSignerPrivateKeys(pseudoRandomNumber, numSigners, numNonSigners);

        // Create all operators
        operators = new IECDSACertificateVerifierTypes.ECDSAOperatorInfo[](numOperators);

        // Track indices of non-signers
        nonSignerIndices = new uint32[](numNonSigners);

        // Create signers first
        for (uint32 i = 0; i < numSigners; i++) {
            operators[i].pubkey = vm.addr(signerPrivKeys[i]);
            operators[i].weights = new uint[](2);
            operators[i].weights[0] = uint(100 + i * 10);
            operators[i].weights[1] = uint(200 + i * 20);
        }

        // Create non-signers
        for (uint32 i = 0; i < numNonSigners; i++) {
            uint32 idx = uint32(numSigners + i);
            operators[idx].pubkey = vm.addr(nonSignerPrivKeys[i]);
            operators[idx].weights = new uint[](2);
            operators[idx].weights[0] = uint(100 + idx * 10);
            operators[idx].weights[1] = uint(200 + idx * 20);
            nonSignerIndices[i] = idx;
        }
    }

    /**
     * @notice Create a certificate with ECDSA signatures
     * @param referenceTimestamp Reference timestamp for the certificate
     * @param messageHash Message hash to sign
     * @param nonSignerIndices Array of non-signer indices
     * @param operators Array of operator infos
     * @param signerPrivKeys Array of signer private keys
     * @return cert The created certificate
     */
    function _createCertificate(
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint32[] memory nonSignerIndices,
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
        uint[] memory signerPrivKeys
    ) internal view returns (IECDSACertificateVerifierTypes.ECDSACertificate memory cert) {
        // Use the contract's digest calculation
        bytes32 signableDigest = verifier.calculateCertificateDigest(referenceTimestamp, messageHash);

        // Collect signers and their private keys
        uint numSigners = signerPrivKeys.length;
        address[] memory signerAddresses = new address[](numSigners);
        bytes[] memory signaturesArr = new bytes[](numSigners);

        // Get signer addresses and create signatures
        for (uint i = 0; i < numSigners; i++) {
            signerAddresses[i] = vm.addr(signerPrivKeys[i]);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPrivKeys[i], signableDigest);
            signaturesArr[i] = abi.encodePacked(r, s, v);
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

        cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            sig: signatures
        });
    }

    /**
     * @notice Update operator table with randomness
     * @param r Randomness for generating operator data
     * @param numSigners Number of signers
     * @param numNonSigners Number of non-signers
     * @return operators Array of operator infos
     * @return operatorSetConfig Operator set configuration
     * @return referenceTimestamp Reference timestamp
     * @return nonSignerIndices Array of non-signer indices
     * @return signerPrivKeys Array of signer private keys
     */
    function _updateOperatorTable(Randomness r, uint numSigners, uint numNonSigners)
        internal
        returns (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        )
    {
        // Generate seed and reference timestamp
        uint seed = r.Uint256();
        referenceTimestamp = r.Uint32(uint32(block.timestamp + 1), uint32(block.timestamp + 1000));

        // Create operators
        (operators, nonSignerIndices, signerPrivKeys) = _createOperatorsWithSplitKeys(seed, numSigners, numNonSigners);
        operatorSetConfig = defaultOperatorSetConfig;

        // Update operator table
        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, operatorSetConfig);
    }

    /**
     * @notice Initialize operator table for basic tests
     * @return referenceTimestamp The reference timestamp used
     */
    function _initializeOperatorTableBase() internal returns (uint32 referenceTimestamp) {
        referenceTimestamp = uint32(block.timestamp);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, numOperators, 0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);
    }
}

/**
 * @title ECDSACertificateVerifierUnitTests_updateOperatorTable
 * @notice Unit tests for ECDSACertificateVerifier.updateOperatorTable
 */
contract ECDSACertificateVerifierUnitTests_updateOperatorTable is ECDSACertificateVerifierUnitTests {
    function test_revert_notTableUpdater() public {
        // Empty data
        uint32 referenceTimestamp;
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators;

        // Update Table
        vm.prank(address(0x100));
        vm.expectRevert(OnlyTableUpdater.selector);
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);
    }

    function test_revert_staleTimestamp() public {
        vm.warp(1000);
        uint32 firstTimestamp = uint32(block.timestamp);
        uint32 staleTimestamp = firstTimestamp - 100;

        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 3, 1);

        vm.startPrank(address(operatorTableUpdaterMock));

        // First update should succeed
        verifier.updateOperatorTable(defaultOperatorSet, firstTimestamp, operators, defaultOperatorSetConfig);

        // Second update with earlier timestamp should fail
        vm.expectRevert(TableUpdateStale.selector);
        verifier.updateOperatorTable(defaultOperatorSet, staleTimestamp, operators, defaultOperatorSetConfig);

        vm.stopPrank();
    }

    function testFuzz_updateOperatorTable_correctness(Randomness r) public rand(r) {
        uint seed = r.Uint256();
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonSigners = numOperators - numSigners;
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) =
            _createOperatorsWithSplitKeys(seed, numSigners, numNonSigners);

        // Expect event
        uint32 referenceTimestamp = uint32(block.timestamp);
        vm.expectEmit(true, true, true, true, address(verifier));
        emit TableUpdated(defaultOperatorSet, referenceTimestamp, operators);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        // Verify storage updates
        assertEq(verifier.latestReferenceTimestamp(defaultOperatorSet), referenceTimestamp, "Reference timestamp mismatch");
        assertEq(verifier.getOperatorSetOwner(defaultOperatorSet), operatorSetOwner, "Operator set owner mismatch");
        assertEq(verifier.maxOperatorTableStaleness(defaultOperatorSet), defaultMaxStaleness, "Max staleness mismatch");

        // Verify operator infos were stored correctly
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOperators =
            verifier.getOperatorInfos(defaultOperatorSet, referenceTimestamp);

        assertEq(storedOperators.length, operators.length, "Number of operators mismatch");
        for (uint i = 0; i < operators.length; i++) {
            assertEq(storedOperators[i].pubkey, operators[i].pubkey, "Operator pubkey mismatch");
            assertEq(storedOperators[i].weights[0], operators[i].weights[0], "Operator weight 0 mismatch");
            assertEq(storedOperators[i].weights[1], operators[i].weights[1], "Operator weight 1 mismatch");
        }
    }

    function test_multiple() public {
        // Create two different operator sets
        OperatorSet memory operatorSet1 = OperatorSet({avs: address(0x10), id: 1});
        OperatorSet memory operatorSet2 = OperatorSet({avs: address(0x20), id: 2});

        uint32 referenceTimestamp = uint32(block.timestamp);

        // Create operators for each set
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators1,,) = _createOperatorsWithSplitKeys(111, 3, 1);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators2,,) = _createOperatorsWithSplitKeys(222, 2, 2);

        // Create operator set configs with different owners
        ICrossChainRegistryTypes.OperatorSetConfig memory config1 = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(0x100),
            maxStalenessPeriod: 1800 // 30 minutes
        });
        ICrossChainRegistryTypes.OperatorSetConfig memory config2 = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(0x200),
            maxStalenessPeriod: 7200 // 2 hours
        });

        vm.startPrank(address(operatorTableUpdaterMock));

        // Update both operator tables
        verifier.updateOperatorTable(operatorSet1, referenceTimestamp, operators1, config1);
        verifier.updateOperatorTable(operatorSet2, referenceTimestamp, operators2, config2);

        vm.stopPrank();

        // Verify that both operator sets are stored correctly and independently
        assertEq(verifier.getOperatorSetOwner(operatorSet1), address(0x100), "OperatorSet1 owner incorrect");
        assertEq(verifier.getOperatorSetOwner(operatorSet2), address(0x200), "OperatorSet2 owner incorrect");

        assertEq(verifier.maxOperatorTableStaleness(operatorSet1), 1800, "OperatorSet1 staleness incorrect");
        assertEq(verifier.maxOperatorTableStaleness(operatorSet2), 7200, "OperatorSet2 staleness incorrect");
    }
}

/**
 * @title ECDSACertificateVerifierUnitTests_verifyCertificate
 * @notice Unit tests for ECDSACertificateVerifier.verifyCertificate
 */
contract ECDSACertificateVerifierUnitTests_verifyCertificate is ECDSACertificateVerifierUnitTests {
    function test_revert_certificateStale() public {
        uint32 referenceTimestamp = _initializeOperatorTableBase();
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert;
        cert.referenceTimestamp = referenceTimestamp;

        // Jump forward in time beyond the max staleness
        vm.warp(block.timestamp + defaultMaxStaleness + 1);

        vm.expectRevert(CertificateStale.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_emptySignatures() public {
        uint32 referenceTimestamp = _initializeOperatorTableBase();

        // Create certificate with empty signatures
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            sig: "" // Empty signatures
        });

        vm.expectRevert(InvalidSignatureLength.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_referenceTimestampDoesNotExist() public {
        uint32 referenceTimestamp = _initializeOperatorTableBase();
        uint32 nonExistentTimestamp = referenceTimestamp + 1000;

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert;
        cert.referenceTimestamp = nonExistentTimestamp;

        vm.expectRevert(ReferenceTimestampDoesNotExist.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_rootDisabled() public {
        // Initialize operator table with a valid root
        uint32 referenceTimestamp = _initializeOperatorTableBase();

        operatorTableUpdaterMock.invalidateRoot(referenceTimestamp);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert;
        cert.referenceTimestamp = referenceTimestamp;

        vm.expectRevert(RootDisabled.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_invalidSignature() public {
        // Update operator table
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _createOperatorsWithSplitKeys(123, 3, 1);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, uint32(block.timestamp), operators, defaultOperatorSetConfig);

        // Create a certificate with signatures for the original message hash
        bytes32 originalHash = defaultMsgHash;
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(uint32(block.timestamp), originalHash, nonSignerIndices, operators, signerPrivKeys);

        // Modify the certificate to use a different message hash
        cert.messageHash = keccak256("different message");

        // Verification should fail
        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function testFuzz_verifyCertificate_allSigners(Randomness r) public rand(r) {
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numOperators, 0);

        // Create certificate with no non-signers
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Calculate total stakes
        uint[] memory totalStakes = new uint[](2);
        for (uint i = 0; i < operators.length; i++) {
            totalStakes[0] += operators[i].weights[0];
            totalStakes[1] += operators[i].weights[1];
        }

        // All stakes should be signed
        assertEq(signedStakes[0], totalStakes[0], "All stake should be signed for type 0");
        assertEq(signedStakes[1], totalStakes[1], "All stake should be signed for type 1");
    }

    function testFuzz_verifyCertificate_someNonSigners(Randomness r) public rand(r) {
        // Create operators and update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Check that the signed stakes are correct
        assertEq(signedStakes.length, 2, "Wrong number of stake types");

        // Calculate expected signed stakes
        uint[] memory expectedSignedStakes = new uint[](2);
        for (uint i = 0; i < operators.length; i++) {
            expectedSignedStakes[0] += operators[i].weights[0];
            expectedSignedStakes[1] += operators[i].weights[1];
        }

        // Subtract non-signer stakes
        for (uint i = 0; i < nonSignerIndices.length; i++) {
            uint32 nonSignerIndex = nonSignerIndices[i];
            expectedSignedStakes[0] -= operators[nonSignerIndex].weights[0];
            expectedSignedStakes[1] -= operators[nonSignerIndex].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");
    }

    function test_updateStalenessAndVerify() public {
        // Initial setup with default staleness (3600)
        uint32 firstTimestamp = uint32(block.timestamp);
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _createOperatorsWithSplitKeys(123, 4, 0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, firstTimestamp, operators, defaultOperatorSetConfig);

        // Warp past the staleness period (3601 seconds)
        vm.warp(block.timestamp + 3601);

        // Update table again with new staleness period (7200) and new owner
        uint32 secondTimestamp = uint32(block.timestamp);
        ICrossChainRegistryTypes.OperatorSetConfig memory newConfig = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(0x999),
            maxStalenessPeriod: 7200 // 2 hours
        });

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, secondTimestamp, operators, newConfig);

        // Warp past old staleness period (3601) but within new staleness period
        vm.warp(block.timestamp + 3601);

        // Create and verify certificate - should succeed with new staleness period
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(secondTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Verify all stakes are signed
        uint[] memory totalStakes = new uint[](2);
        for (uint i = 0; i < operators.length; i++) {
            totalStakes[0] += operators[i].weights[0];
            totalStakes[1] += operators[i].weights[1];
        }
        assertEq(signedStakes[0], totalStakes[0], "All stake should be signed for type 0");
        assertEq(signedStakes[1], totalStakes[1], "All stake should be signed for type 1");

        // Verify the new config is applied
        assertEq(verifier.maxOperatorTableStaleness(defaultOperatorSet), 7200, "New staleness period not applied");
        assertEq(verifier.getOperatorSetOwner(defaultOperatorSet), address(0x999), "New owner not applied");
    }

    function test_updateOperatorsAndVerify() public {
        // First update with initial operators
        uint32 firstTimestamp = uint32(block.timestamp);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory firstOperators,,) = _createOperatorsWithSplitKeys(123, 3, 1);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, firstTimestamp, firstOperators, defaultOperatorSetConfig);

        // Advance time and update with new operators
        vm.warp(block.timestamp + 100);
        uint32 secondTimestamp = uint32(block.timestamp);
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory secondOperators,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _createOperatorsWithSplitKeys(456, 2, 2); // Different seed for different operators

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, secondTimestamp, secondOperators, defaultOperatorSetConfig);

        // Create certificate with new operators
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(secondTimestamp, defaultMsgHash, nonSignerIndices, secondOperators, signerPrivKeys);

        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);

        // Verify only signers' stakes are counted
        uint[] memory expectedSignedStakes = new uint[](2);
        // First 2 operators are signers
        for (uint i = 0; i < 2; i++) {
            expectedSignedStakes[0] += secondOperators[i].weights[0];
            expectedSignedStakes[1] += secondOperators[i].weights[1];
        }

        assertEq(signedStakes[0], expectedSignedStakes[0], "Wrong signed stake for type 0");
        assertEq(signedStakes[1], expectedSignedStakes[1], "Wrong signed stake for type 1");

        // Verify the latest timestamp is updated
        assertEq(verifier.latestReferenceTimestamp(defaultOperatorSet), secondTimestamp, "Latest timestamp not updated");

        // Verify new operators are stored
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOperators =
            verifier.getOperatorInfos(defaultOperatorSet, secondTimestamp);
        assertEq(storedOperators.length, secondOperators.length, "Operator count mismatch");
        for (uint i = 0; i < secondOperators.length; i++) {
            assertEq(storedOperators[i].pubkey, secondOperators[i].pubkey, "Operator pubkey mismatch");
        }
    }

    function test_revert_signerNotOperator() public {
        // Create operators and update the table
        uint32 referenceTimestamp = uint32(block.timestamp);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 4, 0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        // Create a private key for a non-operator signer
        uint nonOperatorPrivKey = uint(keccak256(abi.encodePacked("nonOperator", block.timestamp)));
        if (nonOperatorPrivKey == 0) nonOperatorPrivKey = 1;

        // Create certificate with the non-operator as a signer
        bytes32 signableDigest = verifier.calculateCertificateDigest(referenceTimestamp, defaultMsgHash);

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(nonOperatorPrivKey, signableDigest);
        bytes memory signature = abi.encodePacked(r, s, v);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            sig: signature
        });

        // Should revert because the signer is not an operator
        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_signaturesNotOrdered() public {
        // Create operators with at least 2 signers
        uint32 referenceTimestamp = uint32(block.timestamp);
        uint[] memory signerPrivKeys = new uint[](2);

        // Create two private keys where the first address is greater than the second
        signerPrivKeys[0] = uint(keccak256(abi.encodePacked("signer1")));
        signerPrivKeys[1] = uint(keccak256(abi.encodePacked("signer2")));

        // Ensure valid private keys
        if (signerPrivKeys[0] == 0) signerPrivKeys[0] = 1;
        if (signerPrivKeys[1] == 0) signerPrivKeys[1] = 1;

        // Make sure addr0 > addr1 by adjusting keys if needed
        address addr0 = vm.addr(signerPrivKeys[0]);
        address addr1 = vm.addr(signerPrivKeys[1]);

        // If addr0 < addr1, swap the keys to ensure we get the wrong order
        if (addr0 < addr1) {
            uint temp = signerPrivKeys[0];
            signerPrivKeys[0] = signerPrivKeys[1];
            signerPrivKeys[1] = temp;
            addr0 = vm.addr(signerPrivKeys[0]);
            addr1 = vm.addr(signerPrivKeys[1]);
        }

        // Create operators with these addresses
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = new IECDSACertificateVerifierTypes.ECDSAOperatorInfo[](2);

        operators[0].pubkey = addr0;
        operators[0].weights = new uint[](2);
        operators[0].weights[0] = 100;
        operators[0].weights[1] = 200;

        operators[1].pubkey = addr1;
        operators[1].weights = new uint[](2);
        operators[1].weights[0] = 100;
        operators[1].weights[1] = 200;

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        // Create signatures in the WRONG order (higher address first)
        bytes32 signableDigest = verifier.calculateCertificateDigest(referenceTimestamp, defaultMsgHash);

        // Sign with both keys but put them in wrong order (addr0 > addr1)
        (uint8 v0, bytes32 r0, bytes32 s0) = vm.sign(signerPrivKeys[0], signableDigest);
        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(signerPrivKeys[1], signableDigest);

        // Concatenate signatures in wrong order (higher address first)
        bytes memory signatures = bytes.concat(abi.encodePacked(r0, s0, v0), abi.encodePacked(r1, s1, v1));

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            sig: signatures
        });

        // Should revert because signatures are not ordered by address
        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_invalidSignatureRecoveryError() public {
        // Create operators and update the table
        uint32 referenceTimestamp = uint32(block.timestamp);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 4, 0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        // Create an invalid signature with invalid s value (too large)
        // This will cause ECDSA.tryRecover to return RecoverError.InvalidSignatureS
        bytes memory invalidSignature = new bytes(65);
        // Set r to a valid value
        bytes32 r = bytes32(uint(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef));
        // Set s to a value that's too large (> N/2)
        bytes32 s = bytes32(uint(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141));
        uint8 v = 27;

        // Pack the signature
        for (uint i = 0; i < 32; i++) {
            invalidSignature[i] = r[i];
            invalidSignature[i + 32] = s[i];
        }
        invalidSignature[64] = bytes1(v);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            sig: invalidSignature
        });

        // Should revert because signature recovery returns an error
        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }

    function test_revert_recoveredAddressZero() public {
        // Create operators and update the table
        uint32 referenceTimestamp = uint32(block.timestamp);
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 4, 0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        // Create a signature that will recover to address(0)
        // This happens when v, r, s are valid but don't correspond to any valid signature
        bytes memory zeroRecoverySignature = new bytes(65);

        // These specific values will cause ecrecover to return address(0)
        // while still being within valid ranges
        bytes32 r = bytes32(uint(1));
        bytes32 s = bytes32(uint(1));
        uint8 v = 27;

        // Pack the signature
        for (uint i = 0; i < 32; i++) {
            zeroRecoverySignature[i] = r[i];
            zeroRecoverySignature[i + 32] = s[i];
        }
        zeroRecoverySignature[64] = bytes1(v);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: defaultMsgHash,
            sig: zeroRecoverySignature
        });

        // Should revert because recovered address is zero
        vm.expectRevert(VerificationFailed.selector);
        verifier.verifyCertificate(defaultOperatorSet, cert);
    }
}

/**
 * @title ECDSACertificateVerifierUnitTests_verifyCertificateProportion
 * @notice Unit tests for ECDSACertificateVerifier.verifyCertificateProportion
 */
contract ECDSACertificateVerifierUnitTests_verifyCertificateProportion is ECDSACertificateVerifierUnitTests {
    function testFuzz_revert_arrayLengthMismatch(Randomness r) public rand(r) {
        // Update operator table
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numOperators, 0);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        // Set thresholds with wrong length
        uint16[] memory wrongLengthThresholds = new uint16[](1); // Should be 2
        wrongLengthThresholds[0] = 6000;

        vm.expectRevert(ArrayLengthMismatch.selector);
        verifier.verifyCertificateProportion(defaultOperatorSet, cert, wrongLengthThresholds);
    }

    function testFuzz_verifyCertificateProportion_meetsThresholds(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        // Set thresholds at 60% of total stake for each type
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 6000; // 60%
        thresholds[1] = 6000; // 60%

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate expected result based on the number of signers
        uint signedPercentage = (numSigners * 10_000) / numOperators;
        bool expectedResult = signedPercentage >= 6000;

        assertEq(meetsThresholds, expectedResult, "Certificate threshold check incorrect");
    }

    function testFuzz_verifyCertificateProportion_doesNotMeetThresholds(Randomness r) public rand(r) {
        // Update operator table with a specific split to ensure some thresholds won't be met
        uint numSigners = r.Uint256(1, numOperators / 2); // At most half signers
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        // Try with higher threshold that shouldn't be met
        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = 9000; // 90%
        thresholds[1] = 9000; // 90%

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate percentage of signed stakes to determine if it should meet threshold
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);
        uint[] memory totalStakes = verifier.getTotalStakes(defaultOperatorSet, referenceTimestamp);
        uint signedPercentage0 = (signedStakes[0] * 10_000) / totalStakes[0];
        uint signedPercentage1 = (signedStakes[1] * 10_000) / totalStakes[1];

        bool shouldMeetThreshold = (signedPercentage0 >= 9000) && (signedPercentage1 >= 9000);
        assertEq(meetsThresholds, shouldMeetThreshold, "Certificate threshold check incorrect");
    }

    /// @notice Fuzz against random thresholds
    function testFuzz_verifyCertificateProportion_thresholds(Randomness r, uint16 threshold0, uint16 threshold1) public rand(r) {
        // Update operator table with random split
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        threshold0 = uint16(bound(threshold0, 0, 10_000)); // 0% to 100%
        threshold1 = uint16(bound(threshold1, 0, 10_000)); // 0% to 100%

        uint16[] memory thresholds = new uint16[](2);
        thresholds[0] = threshold0;
        thresholds[1] = threshold1;

        bool meetsThresholds = verifier.verifyCertificateProportion(defaultOperatorSet, cert, thresholds);

        // Calculate expected result
        uint[] memory signedStakes = verifier.verifyCertificate(defaultOperatorSet, cert);
        uint[] memory totalStakes = verifier.getTotalStakes(defaultOperatorSet, referenceTimestamp);

        bool expectedResult = true;
        for (uint i = 0; i < 2; i++) {
            uint threshold = (totalStakes[i] * thresholds[i]) / 10_000;
            if (signedStakes[i] < threshold) {
                expectedResult = false;
                break;
            }
        }

        assertEq(meetsThresholds, expectedResult, "Threshold calculation mismatch");
    }
}

/**
 * @title ECDSACertificateVerifierUnitTests_verifyCertificateNominal
 * @notice Unit tests for ECDSACertificateVerifier.verifyCertificateNominal
 */
contract ECDSACertificateVerifierUnitTests_verifyCertificateNominal is ECDSACertificateVerifierUnitTests {
    function testFuzz_revert_arrayLengthMismatch(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators - 1);
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

        // Set thresholds with wrong length
        uint[] memory wrongLengthThresholds = new uint[](1); // Should be 2
        wrongLengthThresholds[0] = 100;

        vm.expectRevert(ArrayLengthMismatch.selector);
        verifier.verifyCertificateNominal(defaultOperatorSet, cert, wrongLengthThresholds);
    }

    function testFuzz_verifyCertificateNominal_meetsThresholds(Randomness r) public rand(r) {
        // Update operator table
        uint numSigners = r.Uint256(1, numOperators);
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

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
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

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
        uint numNonSigners = numOperators - numSigners;
        (
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            uint32 referenceTimestamp,
            uint32[] memory nonSignerIndices,
            uint[] memory signerPrivKeys
        ) = _updateOperatorTable(r, numSigners, numNonSigners);

        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createCertificate(referenceTimestamp, defaultMsgHash, nonSignerIndices, operators, signerPrivKeys);

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
 * @title ECDSACertificateVerifierUnitTests_ViewFunctions
 * @notice Unit tests for ECDSACertificateVerifier view functions
 */
contract ECDSACertificateVerifierUnitTests_ViewFunctions is ECDSACertificateVerifierUnitTests {
    uint32 referenceTimestamp;

    function setUp() public override {
        super.setUp();
        referenceTimestamp = uint32(block.timestamp);

        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);
    }

    function _getOperators() internal view returns (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory) {
        (IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators,,) = _createOperatorsWithSplitKeys(123, 3, 1);
        return operators;
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

    function test_getOperatorInfos() public view {
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory storedOperators =
            verifier.getOperatorInfos(defaultOperatorSet, referenceTimestamp);

        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();

        assertEq(storedOperators.length, operators.length, "Number of operators mismatch");
        for (uint i = 0; i < operators.length; i++) {
            assertEq(storedOperators[i].pubkey, operators[i].pubkey, "Operator pubkey mismatch");
            assertEq(storedOperators[i].weights.length, operators[i].weights.length, "Weights length mismatch");
            for (uint j = 0; j < operators[i].weights.length; j++) {
                assertEq(storedOperators[i].weights[j], operators[i].weights[j], "Weight value mismatch");
            }
        }
    }

    function test_getOperatorInfo() public view {
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();

        for (uint32 i = 0; i < operators.length; i++) {
            IECDSACertificateVerifierTypes.ECDSAOperatorInfo memory operatorInfo =
                verifier.getOperatorInfo(defaultOperatorSet, referenceTimestamp, i);

            assertEq(operatorInfo.pubkey, operators[i].pubkey, "Operator pubkey mismatch");
            assertEq(operatorInfo.weights[0], operators[i].weights[0], "Weight 0 mismatch");
            assertEq(operatorInfo.weights[1], operators[i].weights[1], "Weight 1 mismatch");
        }
    }

    function test_getOperatorCount() public view {
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();

        uint32 operatorCount = verifier.getOperatorCount(defaultOperatorSet, referenceTimestamp);
        assertEq(operatorCount, operators.length, "Operator count mismatch");
    }

    function test_getTotalStakes() public view {
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();
        uint[] memory totalStakes = verifier.getTotalStakes(defaultOperatorSet, referenceTimestamp);

        // Calculate expected total stakes
        uint[] memory expectedTotalStakes = new uint[](2);
        for (uint i = 0; i < operators.length; i++) {
            expectedTotalStakes[0] += operators[i].weights[0];
            expectedTotalStakes[1] += operators[i].weights[1];
        }

        assertEq(totalStakes[0], expectedTotalStakes[0], "Total stake 0 mismatch");
        assertEq(totalStakes[1], expectedTotalStakes[1], "Total stake 1 mismatch");
    }

    function test_calculateCertificateDigest() public view {
        bytes32 messageHash = keccak256("test");
        uint32 timestamp = uint32(block.timestamp);

        bytes32 digest = verifier.calculateCertificateDigest(timestamp, messageHash);

        // Verify digest is deterministic
        bytes32 digest2 = verifier.calculateCertificateDigest(timestamp, messageHash);
        assertEq(digest, digest2, "Digest should be deterministic");

        // Verify different inputs produce different digests
        bytes32 differentDigest = verifier.calculateCertificateDigest(timestamp + 1, messageHash);
        assertTrue(digest != differentDigest, "Different inputs should produce different digests");
    }

    function test_domainSeparator() public view {
        bytes32 domainSep = verifier.domainSeparator();

        // Verify domain separator is deterministic
        assertTrue(domainSep != bytes32(0), "Domain separator should not be zero");
    }

    function test_revert_indexOutOfBounds() public {
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = _getOperators();

        // Try to get operator info with out of bounds index
        uint32 outOfBoundsIndex = uint32(operators.length);

        vm.expectRevert("Operator index out of bounds");
        verifier.getOperatorInfo(defaultOperatorSet, referenceTimestamp, outOfBoundsIndex);
    }

    function test_revert_referenceTimestampDoesNotExist() public {
        uint32 nonExistentTimestamp = uint32(block.timestamp + 1000);

        vm.expectRevert(ReferenceTimestampDoesNotExist.selector);
        verifier.getTotalStakes(defaultOperatorSet, nonExistentTimestamp);
    }

    function test_revert_noOperators() public {
        uint32 referenceTimestamp = uint32(block.timestamp) + 1;

        // Create empty operator array
        IECDSACertificateVerifierTypes.ECDSAOperatorInfo[] memory operators = new IECDSACertificateVerifierTypes.ECDSAOperatorInfo[](0);

        vm.prank(address(operatorTableUpdaterMock));
        verifier.updateOperatorTable(defaultOperatorSet, referenceTimestamp, operators, defaultOperatorSetConfig);

        vm.expectRevert(ReferenceTimestampDoesNotExist.selector);
        verifier.getTotalStakes(defaultOperatorSet, referenceTimestamp);
    }
}
