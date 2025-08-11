// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../MultichainIntegrationChecks.t.sol";

/**
 * @title Multichain_Timing_Tests
 * @notice Integration tests for multichain timing constraints and edge cases
 * @dev Tests the temporal behavior of global table root confirmations and operator table updates
 */
contract Integration_Multichain_Timing_Tests_Base is MultichainIntegrationCheckUtils {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    OperatorSet operatorSet;
    User[] operators;

    function setUp() public virtual override {
        super.setUp();

        // Warp time to avoid stale reference timestamp
        vm.warp(block.timestamp + 50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(1);
        operatorSet = OperatorSet({avs: address(this), id: 1});
    }

    /**
     * @notice Helper function to confirm global table root and update BN254 table with custom staleness period
     */
    function _confirmGlobalTableRootAndUpdateBN254WithStaleness(
        OperatorSet memory _operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory _operatorSetInfo,
        uint32 referenceTimestamp,
        uint32 stalenessPeriod
    ) internal {
        // Create the operator table data for transport with custom staleness period
        bytes memory operatorTable = abi.encode(
            _operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: stalenessPeriod}),
            abi.encode(_operatorSetInfo)
        );

        // Create global table root containing the operator table
        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update the global table root with confirmation
        _updateGlobalTableRoot(globalTableRoot, referenceTimestamp);

        // Update the operator table using the confirmed global root
        uint32 operatorSetIndex = 0; // Single leaf in merkle tree
        bytes memory proof = new bytes(0); // Empty proof for single leaf

        operatorTableUpdater.updateOperatorTable(referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    /**
     * @notice Helper function to confirm global table root and update ECDSA table with custom staleness period
     */
    function _confirmGlobalTableRootAndUpdateECDSAWithStaleness(
        OperatorSet memory _operatorSet,
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos,
        uint32 referenceTimestamp,
        uint32 stalenessPeriod
    ) internal {
        // Create the operator table data for transport with custom staleness period
        bytes memory operatorTable = abi.encode(
            _operatorSet,
            IKeyRegistrarTypes.CurveType.ECDSA,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: stalenessPeriod}),
            abi.encode(operatorInfos)
        );

        // Create global table root containing the operator table
        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update the global table root with confirmation
        _updateGlobalTableRoot(globalTableRoot, referenceTimestamp);

        // Update the operator table using the confirmed global root
        uint32 operatorSetIndex = 0; // Single leaf in merkle tree
        bytes memory proof = new bytes(0); // Empty proof for single leaf

        operatorTableUpdater.updateOperatorTable(referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    /**
     * @notice Helper function to post a new global root
     */
    function _postNewGlobalRoot(
        OperatorSet memory _operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory _operatorSetInfo,
        uint32 timestamp
    ) internal returns (bytes memory operatorTable, bytes32 globalTableRoot) {
        operatorTable = abi.encode(
            _operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3700}), // Change this value so the root is different
            abi.encode(_operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update global root with new timestamp
        _updateGlobalTableRoot(globalTableRoot, timestamp);
    }

    /**
     * @notice Helper function to test stale table transport
     */
    function _testStaleTableTransport(
        OperatorSet memory _operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory _operatorSetInfo,
        uint32 staleTimestamp
    ) internal {
        // Create operator table data with stale timestamp
        bytes memory staleOperatorTable = abi.encode(
            _operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(_operatorSetInfo)
        );

        // Use the currently valid global table root (not a fresh one)
        bytes32 validGlobalTableRoot = operatorTableUpdater.getCurrentGlobalTableRoot();

        // This should revert with TableUpdateForPastTimestamp error
        // because the timestamp is stale, even though the global root is valid
        check_TableTransport_StaleTimestamp_Reverts(staleTimestamp, validGlobalTableRoot, staleOperatorTable);
    }
}

contract Integration_Multichain_Timing_Tests_GlobalTableRoot is Integration_Multichain_Timing_Tests_Base {
    // BN254 operator info - using BN254 for posting
    IOperatorTableCalculatorTypes.BN254OperatorSetInfo operatorSetInfo;

    function setUp() public override {
        super.setUp();

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        operators = _registerBN254Keys(operatorSet);
        operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);
    }

    /**
     * @notice Test that posting a root with the same reference timestamp as the latest reverts
     *
     */
    function test_PostRoot_SameLatestReferenceTimestamp_Reverts() external {
        console.log("Testing post root with same latestReferenceTimestamp - should revert:");

        // Post initial root successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully posted initial root with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp from the OperatorTableUpdater
        uint32 latestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        console.log("Latest reference timestamp:", latestReferenceTimestamp);

        // Verify that the latest timestamp matches what we posted
        check_LatestReferenceTimestamp_Updated_State(firstReferenceTimestamp);

        // Try to post another root with the same timestamp - this should noop
        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Generate certificate for global root confirmation
        bytes32 messageHash = operatorTableUpdater.getGlobalTableUpdateMessageHash(
            globalTableRoot,
            firstReferenceTimestamp, // Same timestamp as before
            uint32(block.number)
        );

        IBN254CertificateVerifierTypes.BN254Certificate memory confirmationCertificate = _generateGlobalRootConfirmationCertificate(
            operatorTableUpdater.getGenerator(), operatorTableUpdater.getGeneratorReferenceTimestamp(), messageHash
        );

        // This should revert with GlobalTableRootStale error
        check_GlobalTableRoot_PostSameTimestamp_Reverts(globalTableRoot, firstReferenceTimestamp, confirmationCertificate);

        console.log("Successfully verified that posting root with same timestamp reverts");
    }

    /**
     * @notice Test that posting a root right after the latest reference timestamp succeeds
     * @dev This tests the boundary condition where timestamp = latestReferenceTimestamp + 1
     */
    function test_PostRoot_RightAfterLatestReferenceTimestamp_Succeeds() external {
        console.log("Testing post root right after latestReferenceTimestamp - should succeed:");

        // Post initial root successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully posted initial root with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp from the OperatorTableUpdater
        check_LatestReferenceTimestamp_Updated_State(firstReferenceTimestamp);
        uint32 latestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        console.log("Latest reference timestamp:", latestReferenceTimestamp);

        // Post another root with timestamp = latestReferenceTimestamp + 1
        uint32 nextReferenceTimestamp = latestReferenceTimestamp + 1;
        console.log("Attempting to post root with timestamp:", nextReferenceTimestamp);

        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3700}), // Change this value so the root is different
            abi.encode(operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Generate certificate for global root confirmation
        bytes32 messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, nextReferenceTimestamp, uint32(block.number));

        IBN254CertificateVerifierTypes.BN254Certificate memory confirmationCertificate = _generateGlobalRootConfirmationCertificate(
            operatorTableUpdater.getGenerator(), operatorTableUpdater.getGeneratorReferenceTimestamp(), messageHash
        );

        // This should succeed
        check_GlobalTableRoot_PostSuccess_State(globalTableRoot, nextReferenceTimestamp, confirmationCertificate);

        console.log("Successfully posted root with timestamp right after latest reference timestamp");

        // Verify that the latest reference timestamp has been updated
        uint32 newLatestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        console.log("New latest reference timestamp:", newLatestReferenceTimestamp);
    }

    /**
     * @notice Test that transporting tables with a stale reference timestamp reverts
     * @dev This tests that operator table updates must have timestamps > latest for that operator set
     */
    function test_TransportTables_StaleReferenceTimestamp_Reverts() external {
        console.log("Testing transport tables with stale reference timestamp - should revert:");

        // Post initial root and transport tables successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully transported initial table with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp for this operator set from the certificate verifier
        uint32 operatorSetLatestTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        console.log("Operator set latest reference timestamp:", operatorSetLatestTimestamp);

        // Verify that the operator set timestamp matches what we posted
        check_OperatorSetLatestTimestamp_State(operatorSet, firstReferenceTimestamp);

        // Post a new global root with a later timestamp
        uint32 newGlobalTimestamp = firstReferenceTimestamp + 500;
        _postNewGlobalRoot(operatorSet, operatorSetInfo, newGlobalTimestamp);
        console.log("Successfully posted new global root with timestamp:", newGlobalTimestamp);

        // Try to transport tables with a stale reference timestamp (older than operatorSetLatestTimestamp)
        uint32 staleTimestamp = operatorSetLatestTimestamp - 100; // Older than current operator set timestamp
        console.log("Attempting to transport table with stale timestamp:", staleTimestamp);

        // This should revert with TableUpdateForPastTimestamp error
        _testStaleTableTransport(operatorSet, operatorSetInfo, staleTimestamp);

        console.log("Successfully verified that transporting tables with stale timestamp reverts");

        // Verify that the operator set timestamp hasn't changed
        check_OperatorSetLatestTimestamp_State(operatorSet, operatorSetLatestTimestamp);
        uint32 finalTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        console.log("Operator set timestamp remains:", finalTimestamp);
    }

    /**
     * @notice Test that transporting tables right after updating the latest reference timestamp succeeds
     * @dev This tests that table transport works immediately after the global table root is updated
     */
    function test_TransportTables_RightAfterLatestReferenceTimestamp_Succeeds() external {
        console.log("Testing transport tables right after latest reference timestamp - should succeed:");

        // Post initial root and transport tables successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully transported initial table with timestamp:", firstReferenceTimestamp);

        // Verify that global and operator set reference timestamps are correct
        check_LatestReferenceTimestamp_Updated_State(firstReferenceTimestamp);
        check_OperatorSetLatestTimestamp_State(operatorSet, firstReferenceTimestamp);

        // Post a new global root with timestamp = latestReferenceTimestamp + 1
        uint32 newGlobalTimestamp = firstReferenceTimestamp + 1;
        (bytes memory operatorTable, bytes32 globalTableRoot) = _postNewGlobalRoot(operatorSet, operatorSetInfo, newGlobalTimestamp);
        console.log("Successfully posted new global root with timestamp:", newGlobalTimestamp);

        // Immediately transport the table with the new timestamp (should succeed)
        uint32 transportTimestamp = newGlobalTimestamp; // Same timestamp as global root
        console.log("Attempting to transport table with timestamp:", transportTimestamp);

        // Transport should succeed since timestamp matches the global root timestamp
        uint32 operatorSetIndex = 0; // Single leaf in merkle tree
        bytes memory proof = new bytes(0); // Empty proof for single leaf

        operatorTableUpdater.updateOperatorTable(transportTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTable);

        console.log("Successfully transported table right after latest reference timestamp");

        // Verify that the operator set timestamp was updated correctly
        check_OperatorSetLatestTimestamp_State(operatorSet, transportTimestamp);
        uint32 finalOperatorSetTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        console.log("Final operator set timestamp:", finalOperatorSetTimestamp);
    }
}

contract Integration_Multichain_Timing_Tests_BN254 is Integration_Multichain_Timing_Tests_Base {
    // BN254 operator info
    IOperatorTableCalculatorTypes.BN254OperatorSetInfo operatorSetInfo;

    function setUp() public override {
        super.setUp();

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        operators = _registerBN254Keys(operatorSet);
        operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);
    }

    /**
     * @notice Test that verifying BN254 certificate after staleness period expires reverts
     * @dev This tests that certificates become invalid after the maxStalenessPeriod
     */
    function test_VerifyBN254Certificate_AfterStalenessPeriod_Reverts() external {
        console.log("Testing BN254 certificate verification after staleness period - should revert:");

        // Set a short staleness period for testing (1 hour)
        uint32 stalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254WithStaleness(operatorSet, operatorSetInfo, referenceTimestamp, stalenessPeriod);

        // Generate certificate with the reference timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // Verify certificate works initially
        check_BN254Certificate_Basic_State(operatorSet, certificate);
        console.log("Certificate verification successful before staleness period");

        // Jump forward past the staleness period
        _rollForwardTime(stalenessPeriod + 1);
        console.log("Jumped forward past staleness period to timestamp:", block.timestamp);

        // Verify certificate should now revert due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);

        console.log("Successfully verified that BN254 certificate verification reverts after staleness period");
    }

    /**
     * @notice Test that verifying BN254 certificate right at staleness period boundary succeeds
     * @dev This tests the exact boundary condition where block.timestamp = referenceTimestamp + maxStalenessPeriod
     */
    function test_VerifyBN254Certificate_RightOnStalenessPeriod_Succeeds() external {
        console.log("Testing BN254 certificate verification right on staleness period boundary - should succeed:");

        // Set a staleness period for testing (1 hour)
        uint32 stalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254WithStaleness(operatorSet, operatorSetInfo, referenceTimestamp, stalenessPeriod);

        // Generate certificate with the reference timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // Jump forward to exactly the staleness boundary (referenceTimestamp + stalenessPeriod)
        uint32 boundaryTimestamp = referenceTimestamp + stalenessPeriod;
        vm.warp(boundaryTimestamp);
        console.log("Set timestamp to staleness boundary:", boundaryTimestamp);

        // Verify certificate should still succeed at the boundary
        check_BN254Certificate_Basic_State(operatorSet, certificate);

        console.log("Successfully verified that BN254 certificate verification succeeds right at staleness period boundary");
    }

    /**
     * @notice Test complex scenario: verify after staleness -> revert -> update staleness period -> verify again -> pass (BN254)
     * @dev This tests the complete staleness period update flow for BN254
     */
    function test_BN254Certificate_StalenessUpdate_ComplexFlow() external {
        console.log("Testing BN254 complex staleness update flow:");

        // Set initial short staleness period (1 hour)
        uint32 initialStalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254WithStaleness(operatorSet, operatorSetInfo, referenceTimestamp, initialStalenessPeriod);

        // Generate certificate with the reference timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // Jump forward past the initial staleness period
        _rollForwardTime(initialStalenessPeriod + 1);
        console.log("Jumped past initial staleness period");

        // Verify certificate should revert due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);
        console.log("Certificate verification correctly reverted after staleness period");

        // Update the staleness period to be longer (2 hours from original reference timestamp)
        uint32 newStalenessPeriod = 7200;
        uint32 newReferenceTimestamp = uint32(block.timestamp);
        _confirmGlobalTableRootAndUpdateBN254WithStaleness(operatorSet, operatorSetInfo, newReferenceTimestamp, newStalenessPeriod);
        console.log("Updated staleness period to:", newStalenessPeriod);

        // Generate new certificate with the new reference timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory newCertificate =
            _generateBN254Certificate(operatorSet, newReferenceTimestamp, keccak256("test message 2"), operators, _getBN254PrivateKeys());

        // Verify new certificate should work with updated staleness period
        check_BN254Certificate_Basic_State(operatorSet, newCertificate);
        console.log("New certificate verification successful with updated staleness period");

        // Verify that the staleness period was actually updated
        uint32 retrievedStalenessPeriod = bn254CertificateVerifier.maxOperatorTableStaleness(operatorSet);
        assertEq(retrievedStalenessPeriod, newStalenessPeriod, "Staleness period should be updated");

        console.log("Successfully completed BN254 complex staleness update flow");
    }
}

contract Integration_Multichain_Timing_Tests_ECDSA is Integration_Multichain_Timing_Tests_Base {
    // ECDSA operator info

    function setUp() public override {
        super.setUp();

        // Configure operator set for ECDSA curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        operators = _registerECDSAKeys(operatorSet);
        _createGenerationReservation(operatorSet);
    }

    /**
     * @notice Test that verifying ECDSA certificate after staleness period expires reverts
     * @dev This tests that ECDSA certificates become invalid after the maxStalenessPeriod
     */
    function test_VerifyECDSACertificate_AfterStalenessPeriod_Reverts() external {
        console.log("Testing ECDSA certificate verification after staleness period - should revert:");
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos = _generateECDSAOperatorTable(operatorSet, operators);

        // Set a short staleness period for testing (1 hour)
        uint32 stalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateECDSAWithStaleness(operatorSet, operatorInfos, referenceTimestamp, stalenessPeriod);

        // Generate certificate with the reference timestamp
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate =
            _generateECDSACertificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getECDSAPrivateKeys());

        // Verify certificate works initially
        check_ECDSACertificate_Basic_State(operatorSet, certificate);
        console.log("Certificate verification successful before staleness period");

        // Jump forward past the staleness period
        _rollForwardTime(stalenessPeriod + 1);
        console.log("Jumped forward past staleness period to timestamp:", block.timestamp);

        // Verify certificate should now revert due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);

        console.log("Successfully verified that ECDSA certificate verification reverts after staleness period");
    }

    /**
     * @notice Test that verifying ECDSA certificate right at staleness period boundary succeeds
     * @dev This tests the exact boundary condition for ECDSA certificates
     */
    function test_VerifyECDSACertificate_RightOnStalenessPeriod_Succeeds() external {
        console.log("Testing ECDSA certificate verification right on staleness period boundary - should succeed:");
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos = _generateECDSAOperatorTable(operatorSet, operators);

        // Set a staleness period for testing (1 hour)
        uint32 stalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateECDSAWithStaleness(operatorSet, operatorInfos, referenceTimestamp, stalenessPeriod);

        // Generate certificate with the reference timestamp
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate =
            _generateECDSACertificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getECDSAPrivateKeys());

        // Jump forward to exactly the staleness boundary
        uint32 boundaryTimestamp = referenceTimestamp + stalenessPeriod;
        vm.warp(boundaryTimestamp);
        console.log("Set timestamp to staleness boundary:", boundaryTimestamp);

        // Verify certificate should still succeed at the boundary
        check_ECDSACertificate_Basic_State(operatorSet, certificate);

        console.log("Successfully verified that ECDSA certificate verification succeeds right at staleness period boundary");
    }

    /**
     * @notice Test complex scenario: verify after staleness -> revert -> update staleness period -> verify again -> pass (ECDSA)
     * @dev This tests the complete staleness period update flow for ECDSA
     */
    function test_ECDSACertificate_StalenessUpdate_ComplexFlow() external {
        console.log("Testing ECDSA complex staleness update flow:");
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos = _generateECDSAOperatorTable(operatorSet, operators);

        // Set initial short staleness period (1 hour)
        uint32 initialStalenessPeriod = 3600;
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateECDSAWithStaleness(operatorSet, operatorInfos, referenceTimestamp, initialStalenessPeriod);

        // Generate certificate with the reference timestamp
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate =
            _generateECDSACertificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getECDSAPrivateKeys());

        // Jump forward past the initial staleness period
        _rollForwardTime(initialStalenessPeriod + 1);
        console.log("Jumped past initial staleness period");

        // Verify certificate should revert due to staleness
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);
        console.log("Certificate verification correctly reverted after staleness period");

        // Update the staleness period to be longer (2 hours from original reference timestamp)
        uint32 newStalenessPeriod = 7200;
        uint32 newReferenceTimestamp = uint32(block.timestamp);
        _confirmGlobalTableRootAndUpdateECDSAWithStaleness(operatorSet, operatorInfos, newReferenceTimestamp, newStalenessPeriod);
        console.log("Updated staleness period to:", newStalenessPeriod);

        // Generate new certificate with the new reference timestamp
        IECDSACertificateVerifierTypes.ECDSACertificate memory newCertificate =
            _generateECDSACertificate(operatorSet, newReferenceTimestamp, keccak256("test message 2"), operators, _getECDSAPrivateKeys());

        // Verify new certificate should work with updated staleness period
        check_ECDSACertificate_Basic_State(operatorSet, newCertificate);
        console.log("New certificate verification successful with updated staleness period");

        // Verify that the staleness period was actually updated
        uint32 retrievedStalenessPeriod = ecdsaCertificateVerifier.maxOperatorTableStaleness(operatorSet);
        assertEq(retrievedStalenessPeriod, newStalenessPeriod, "Staleness period should be updated");

        console.log("Successfully completed ECDSA complex staleness update flow");
    }
}
