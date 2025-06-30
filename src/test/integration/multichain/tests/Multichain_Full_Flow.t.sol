// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../MultichainIntegrationBase.t.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";

/**
 * @title Multichain_Full_Flow
 * @notice Integration tests for multichain functionality
 * @dev Tests the complete flow: key registration -> table calculation ->
 *      transport simulation -> table update -> certificate verification
 */
contract Multichain_Full_Flow is MultichainIntegrationBase {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    /**
     * @notice Test complete BN254 multichain happy path flow
     * @dev Covers key registration, table generation, cross-chain transport simulation,
     *      table updates, and certificate verification
     */
    function test_BN254_MultichainHappyPath() external {
        console.log("Testing BN254 multichain flow:");
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

        // Simulate cross-chain transport
        console.log("Simulating cross-chain transport from chain 1 to chain 137");

        // Update operator table on destination chain using recent timestamp to avoid staleness
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, referenceTimestamp);

        // Generate and verify certificate
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        uint[] memory signedStakes = bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);

        // Validate results
        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, 2, "Should have two stake types");
        console.log("BN254 certificate verified successfully with", signedStakes.length, "stake types");
    }

    /**
     * @notice Test complete ECDSA multichain happy path flow
     * @dev Covers key registration, table generation, cross-chain transport simulation,
     *      table updates, and certificate verification
     */
    function test_ECDSA_MultichainHappyPath() external {
        console.log("Testing ECDSA multichain flow:");
        vm.warp(50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(2);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 2});

        // Configure operator set for ECDSA curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        User[] memory operators = _registerECDSAKeys(operatorSet);
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos = _generateECDSAOperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);

        // Simulate cross-chain transport
        console.log("Simulating cross-chain transport from chain 1 to chain 137");

        // Update operator table on destination chain using recent timestamp to avoid staleness
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateECDSA(operatorSet, operatorInfos, referenceTimestamp);

        // Generate and verify certificate
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate =
            _generateECDSACertificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getECDSAPrivateKeys());

        uint[] memory signedStakes = ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);

        // Validate results
        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, 2, "Should have two stake types");
        console.log("ECDSA certificate verified successfully with", signedStakes.length, "stake types");
    }
}
