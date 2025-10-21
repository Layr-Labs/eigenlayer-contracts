// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../MultichainIntegrationChecks.t.sol";

/**
 * @title Multichain_Full_Flow
 * @notice Integration tests for multichain functionality
 * @dev Tests the complete flow: key registration -> table calculation ->
 *      transport simulation -> table update -> certificate verification
 */
contract Integration_Multichain_Full_Flow_BN254 is MultichainIntegrationCheckUtils {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    OperatorSet operatorSet;
    User[] operators;
    IOperatorTableCalculatorTypes.BN254OperatorSetInfo operatorSetInfo;

    function setUp() public override {
        super.setUp();

        vm.warp(block.timestamp + 50_000);

        // Setup test environment
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);
        _createTestOperatorSet(1);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 1});

        // Configure operator set for BN254 curve
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        _setupAVSAndChains();

        // Register operator keys and generate operator table
        operators = _registerBN254Keys(operatorSet);
        operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        _createGenerationReservation(operatorSet);
    }

    /**
     * @notice Test complete BN254 multichain happy path flow
     * @dev Covers key registration, table generation, cross-chain transport simulation,
     *      table updates, and certificate verification
     */
    function test_BN254_MultichainStandard() external {
        console.log("Testing BN254 multichain flow:");

        // Simulate cross-chain transport
        console.log("Simulating cross-chain transport from chain 1 to chain 137");

        // Update operator table on destination chain using recent timestamp to avoid staleness
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, referenceTimestamp);

        // Generate and verify certificate
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // Validate certificate verification
        check_BN254Certificate_Basic_State(operatorSet, certificate);
        check_BN254MultichainFlow_Complete_State(operatorSet, certificate, referenceTimestamp);
        console.log("BN254 certificate verified successfully");
    }

    /**
     * @notice Test BN254 multichain flow with proportional verification
     * @dev Test case 3: Verify BN254 certificate with proportional thresholds
     */
    function test_BN254_MultichainStandard_ProportionalVerification() external {
        console.log("Testing BN254 multichain flow with proportional verification:");

        // Simulate cross-chain transport
        console.log("Simulating cross-chain transport from chain 1 to chain 137");

        // Update operator table on destination chain
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, referenceTimestamp);

        // Generate certificate
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // First verify with regular method
        check_BN254Certificate_Basic_State(operatorSet, certificate);
        console.log("Regular verification completed");

        // Test proportional verification with 60% threshold
        uint16[] memory proportionalThresholds = new uint16[](2);
        proportionalThresholds[0] = 6000; // 60% in basis points
        proportionalThresholds[1] = 6000; // 60% in basis points

        check_BN254Certificate_ProportionalVerification_State(
            operatorSet, certificate, proportionalThresholds, true, "Certificate should meet 60% proportional threshold"
        );
        console.log("Proportional verification (60% threshold): PASSED");

        // Test with higher threshold that should still pass (80%)
        proportionalThresholds[0] = 8000; // 80%
        proportionalThresholds[1] = 8000; // 80%

        check_BN254Certificate_ProportionalVerification_State(
            operatorSet, certificate, proportionalThresholds, true, "Certificate should meet 80% proportional threshold"
        );
        console.log("Proportional verification (80% threshold): PASSED");

        // Test with very high threshold that should still pass (99%)
        proportionalThresholds[0] = 9900; // 99%
        proportionalThresholds[1] = 9900; // 99%

        check_BN254Certificate_ProportionalVerification_State(
            operatorSet, certificate, proportionalThresholds, true, "Certificate should meet 99% proportional threshold with 100% signing"
        );
        console.log("Proportional verification (99% threshold): PASSED");

        console.log("BN254 proportional verification tests completed successfully");
    }

    /**
     * @notice Test BN254 multichain flow with nominal verification
     * @dev Test case 4: Verify BN254 certificate with nominal thresholds
     */
    function test_BN254_MultichainStandard_NominalVerification() external {
        console.log("Testing BN254 multichain flow with nominal verification:");

        // Simulate cross-chain transport
        console.log("Simulating cross-chain transport from chain 1 to chain 137");

        // Update operator table on destination chain
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, referenceTimestamp);

        // Generate certificate
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate =
            _generateBN254Certificate(operatorSet, referenceTimestamp, keccak256("test message"), operators, _getBN254PrivateKeys());

        // First verify with regular method to get signed stakes
        uint[] memory signedStakes = bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);
        check_BN254Certificate_Basic_State(operatorSet, certificate);
        console.log("Regular verification completed");
        console.log("Signed stakes - Type 0:", signedStakes[0]);
        console.log("Signed stakes - Type 1:", signedStakes[1]);

        // Test nominal verification with thresholds below signed stakes (should pass)
        uint[] memory nominalThresholds = new uint[](2);
        nominalThresholds[0] = signedStakes[0] > 100 ? signedStakes[0] - 100 : 0;
        nominalThresholds[1] = signedStakes[1] > 100 ? signedStakes[1] - 100 : 0;

        check_BN254Certificate_NominalVerification_State(
            operatorSet, certificate, nominalThresholds, true, "Certificate should meet nominal thresholds below signed stakes"
        );
        console.log("Nominal verification (below signed stakes): PASSED");

        // Test with thresholds equal to signed stakes (should pass)
        nominalThresholds[0] = signedStakes[0];
        nominalThresholds[1] = signedStakes[1];

        check_BN254Certificate_NominalVerification_State(
            operatorSet, certificate, nominalThresholds, true, "Certificate should meet exact nominal thresholds"
        );
        console.log("Nominal verification (exact signed stakes): PASSED");

        // Test with thresholds above signed stakes (should fail)
        nominalThresholds[0] = signedStakes[0] + 100;
        nominalThresholds[1] = signedStakes[1] + 100;

        check_BN254Certificate_NominalVerification_State(
            operatorSet, certificate, nominalThresholds, false, "Certificate should not meet nominal thresholds above signed stakes"
        );
        console.log("Nominal verification (above signed stakes): CORRECTLY FAILED");

        console.log("BN254 nominal verification tests completed successfully");
    }
}

contract Integration_Multichain_Full_Flow_ECDSA is MultichainIntegrationCheckUtils {
    using StdStyle for *;

    /**
     * @notice Test complete ECDSA multichain happy path flow
     * @dev Covers key registration, table generation, cross-chain transport simulation,
     *      table updates, and certificate verification
     */
    function test_ECDSA_MultichainStandard() external {
        console.log("Testing ECDSA multichain flow:");
        vm.warp(block.timestamp + 50_000);

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

        // Validate certificate verification
        check_ECDSACertificate_Basic_State(operatorSet, certificate);
        check_ECDSAMultichainFlow_Complete_State(operatorSet, certificate, referenceTimestamp);
        console.log("ECDSA certificate verified successfully");
    }
}
