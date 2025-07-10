// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../MultichainIntegrationBase.t.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";

/**
 * @title Multichain_Timing_Tests
 * @notice Integration tests for multichain timing constraints and edge cases
 * @dev Tests the temporal behavior of global table root confirmations and operator table updates
 */
contract Multichain_Timing_Tests is MultichainIntegrationBase {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    /**
     * @notice Test that posting a root with the same reference timestamp as the latest reverts
     * @dev This tests the constraint that new roots must have timestamps > latestReferenceTimestamp
     */
    function test_PostRoot_SameLatestReferenceTimestamp_Reverts() external {
        console.log("Testing post root with same latestReferenceTimestamp - should revert:");
        vm.warp(50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(1);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 1});

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        User[] memory operators = _registerBN254Keys(operatorSet);
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);

        // Post initial root successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully posted initial root with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp from the OperatorTableUpdater
        uint32 latestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        console.log("Latest reference timestamp:", latestReferenceTimestamp);

        // Verify that the latest timestamp matches what we posted
        assertEq(latestReferenceTimestamp, firstReferenceTimestamp, "Latest reference timestamp should match first posted timestamp");

        // Try to post another root with the same timestamp - this should fail
        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = keccak256(operatorTable);
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
        vm.expectRevert(abi.encodeWithSignature("GlobalTableRootStale()"));
        operatorTableUpdater.confirmGlobalTableRoot(
            confirmationCertificate,
            globalTableRoot,
            firstReferenceTimestamp, // Same timestamp
            uint32(block.number)
        );

        console.log("Successfully verified that posting root with same timestamp reverts");
    }

    /**
     * @notice Test that posting a root right after the latest reference timestamp succeeds
     * @dev This tests the boundary condition where timestamp = latestReferenceTimestamp + 1
     */
    function test_PostRoot_RightAfterLatestReferenceTimestamp_Succeeds() external {
        console.log("Testing post root right after latestReferenceTimestamp - should succeed:");
        vm.warp(50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(1);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 1});

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        User[] memory operators = _registerBN254Keys(operatorSet);
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);

        // Post initial root successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully posted initial root with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp from the OperatorTableUpdater
        uint32 latestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        console.log("Latest reference timestamp:", latestReferenceTimestamp);

        // Post another root with timestamp = latestReferenceTimestamp + 1
        uint32 nextReferenceTimestamp = latestReferenceTimestamp + 1;
        console.log("Attempting to post root with timestamp:", nextReferenceTimestamp);

        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = keccak256(operatorTable);
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
        operatorTableUpdater.confirmGlobalTableRoot(confirmationCertificate, globalTableRoot, nextReferenceTimestamp, uint32(block.number));

        console.log("Successfully posted root with timestamp right after latest reference timestamp");

        // Verify that the latest reference timestamp has been updated
        uint32 newLatestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        assertEq(newLatestReferenceTimestamp, nextReferenceTimestamp, "Latest reference timestamp should be updated");
        console.log("New latest reference timestamp:", newLatestReferenceTimestamp);
    }

    /**
     * @notice Test that transporting tables with a stale reference timestamp reverts
     * @dev This tests that operator table updates must have timestamps > latest for that operator set
     */
    function test_TransportTables_StaleReferenceTimestamp_Reverts() external {
        console.log("Testing transport tables with stale reference timestamp - should revert:");
        vm.warp(50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(1);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 1});

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        User[] memory operators = _registerBN254Keys(operatorSet);
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);

        // Post initial root and transport tables successfully
        uint32 firstReferenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, firstReferenceTimestamp);
        console.log("Successfully transported initial table with timestamp:", firstReferenceTimestamp);

        // Get the latest reference timestamp for this operator set from the certificate verifier
        uint32 operatorSetLatestTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        console.log("Operator set latest reference timestamp:", operatorSetLatestTimestamp);

        // Verify that the operator set timestamp matches what we posted
        assertEq(operatorSetLatestTimestamp, firstReferenceTimestamp, "Operator set timestamp should match first posted timestamp");

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
        uint32 finalTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        assertEq(finalTimestamp, operatorSetLatestTimestamp, "Operator set timestamp should remain unchanged after failed update");
        console.log("Operator set timestamp remains:", finalTimestamp);
    }

    /**
     * @notice Helper function to post a new global root
     */
    function _postNewGlobalRoot(
        OperatorSet memory operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo,
        uint32 timestamp
    ) internal {
        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        bytes32 operatorSetLeafHash = keccak256(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update global root with new timestamp
        _updateGlobalTableRoot(globalTableRoot, timestamp);
    }

    /**
     * @notice Helper function to test stale table transport
     */
    function _testStaleTableTransport(
        OperatorSet memory operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo,
        uint32 staleTimestamp
    ) internal {
        // Create operator table data with stale timestamp
        bytes memory staleOperatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        // Use the currently valid global table root (not a fresh one)
        bytes32 validGlobalTableRoot = operatorTableUpdater.getCurrentGlobalTableRoot();

        // This should revert with TableUpdateForPastTimestamp error
        // because the timestamp is stale, even though the global root is valid
        vm.expectRevert(abi.encodeWithSignature("TableUpdateForPastTimestamp()"));
        operatorTableUpdater.updateOperatorTable(
            staleTimestamp, // Stale timestamp
            validGlobalTableRoot, // Valid global root from latest timestamp
            0, // operatorSetIndex
            new bytes(0), // proof (will fail later, but timestamp check comes first)
            staleOperatorTable
        );
    }
}
