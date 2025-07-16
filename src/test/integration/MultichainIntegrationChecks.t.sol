// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/MultichainIntegrationBase.t.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";

/// @notice Contract that provides utility functions to reuse common test blocks & checks for multichain functionality
contract MultichainIntegrationCheckUtils is MultichainIntegrationBase {
    using ArrayLib for *;
    using StdStyle for *;

    /**
     *
     *                              CERTIFICATE VERIFICATION CHECKS
     *
     */
    function check_BN254Certificate_Basic_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate
    ) internal {
        check_BN254Certificate_Basic_State(operatorSet, certificate, 2);
    }

    function check_BN254Certificate_Basic_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint expectedStakeTypes
    ) internal {
        uint[] memory signedStakes = bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);

        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, expectedStakeTypes, "Should have expected number of stake types");
    }

    function check_BN254Certificate_Basic_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint expectedStakeTypes,
        uint[] memory expectedMinStakes
    ) internal returns (uint[] memory signedStakes) {
        signedStakes = bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);

        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, expectedStakeTypes, "Should have expected number of stake types");

        if (expectedMinStakes.length > 0) {
            assertEq(signedStakes.length, expectedMinStakes.length, "Signed stakes length should match expected min stakes length");
            for (uint i = 0; i < signedStakes.length; i++) {
                assertGe(
                    signedStakes[i], expectedMinStakes[i], string.concat("Signed stake at index ", vm.toString(i), " should meet minimum")
                );
            }
        }
    }

    function check_BN254Certificate_ProportionalVerification_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint16[] memory proportionalThresholds,
        bool expectedResult,
        string memory message
    ) internal {
        bool meetsThresholds = bn254CertificateVerifier.verifyCertificateProportion(operatorSet, certificate, proportionalThresholds);

        if (expectedResult) assertTrue(meetsThresholds, message);
        else assertFalse(meetsThresholds, message);
    }

    function check_BN254Certificate_NominalVerification_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint[] memory nominalThresholds,
        bool expectedResult,
        string memory message
    ) internal {
        bool meetsThresholds = bn254CertificateVerifier.verifyCertificateNominal(operatorSet, certificate, nominalThresholds);

        if (expectedResult) assertTrue(meetsThresholds, message);
        else assertFalse(meetsThresholds, message);
    }

    function check_ECDSACertificate_Basic_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate
    ) internal {
        check_ECDSACertificate_Basic_State(operatorSet, certificate, 2, 1);
    }

    function check_ECDSACertificate_Basic_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate,
        uint expectedStakeTypes,
        uint expectedMinSigners
    ) internal {
        (uint[] memory signedStakes, address[] memory signers) = ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);

        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, expectedStakeTypes, "Should have expected number of stake types");
        assertGe(signers.length, expectedMinSigners, "Should have at least expected minimum number of signers");
    }

    function check_ECDSACertificate_Basic_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate,
        uint expectedStakeTypes,
        uint expectedMinSigners,
        uint[] memory expectedMinStakes
    ) internal returns (uint[] memory signedStakes, address[] memory signers) {
        (signedStakes, signers) = ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);

        assertGt(signedStakes.length, 0, "Should return signed stakes");
        assertEq(signedStakes.length, expectedStakeTypes, "Should have expected number of stake types");
        assertGe(signers.length, expectedMinSigners, "Should have at least expected minimum number of signers");

        if (expectedMinStakes.length > 0) {
            assertEq(signedStakes.length, expectedMinStakes.length, "Signed stakes length should match expected min stakes length");
            for (uint i = 0; i < signedStakes.length; i++) {
                assertGe(
                    signedStakes[i], expectedMinStakes[i], string.concat("Signed stake at index ", vm.toString(i), " should meet minimum")
                );
            }
        }
    }

    /**
     *
     *                              GENERATION RESERVATION CHECKS
     *
     */
    function check_GenerationReservation_Exists_State(OperatorSet memory operatorSet) internal {
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        bool reservationFound = false;
        for (uint i = 0; i < activeReservations.length; i++) {
            if (activeReservations[i].avs == operatorSet.avs && activeReservations[i].id == operatorSet.id) {
                reservationFound = true;
                break;
            }
        }
        assertTrue(reservationFound, "Generation reservation should exist");
    }

    function check_GenerationReservation_NotExists_State(OperatorSet memory operatorSet) internal {
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        bool reservationFound = false;
        for (uint i = 0; i < activeReservations.length; i++) {
            if (activeReservations[i].avs == operatorSet.avs && activeReservations[i].id == operatorSet.id) {
                reservationFound = true;
                break;
            }
        }
        assertFalse(reservationFound, "Generation reservation should not exist");
    }

    function check_OperatorTableBytes_Generation_State(OperatorSet memory operatorSet, bool shouldSucceed) internal {
        if (shouldSucceed) {
            bytes memory operatorTableBytes = crossChainRegistry.calculateOperatorTableBytes(operatorSet);
            assertTrue(operatorTableBytes.length > 0, "Should be able to generate operator table bytes");
        } else {
            vm.expectRevert();
            crossChainRegistry.calculateOperatorTableBytes(operatorSet);
        }
    }

    function check_OperatorSetConfig_Cleared_State(OperatorSet memory operatorSet) internal {
        IOperatorTableCalculator calculator = crossChainRegistry.getOperatorTableCalculator(operatorSet);
        assertEq(address(calculator), address(0), "Operator table calculator should be cleared");

        ICrossChainRegistryTypes.OperatorSetConfig memory config = crossChainRegistry.getOperatorSetConfig(operatorSet);
        assertEq(config.owner, address(0), "Operator set config owner should be cleared");
        assertEq(config.maxStalenessPeriod, 0, "Max staleness period should be cleared");
    }

    /**
     *
     *                              TIMESTAMP AND TIMING CHECKS
     *
     */
    function check_LatestReferenceTimestamp_Updated_State(uint32 expectedTimestamp) internal {
        uint32 latestReferenceTimestamp = operatorTableUpdater.getLatestReferenceTimestamp();
        assertEq(latestReferenceTimestamp, expectedTimestamp, "Latest reference timestamp should match expected value");
    }

    function check_OperatorSetLatestTimestamp_State(OperatorSet memory operatorSet, uint32 expectedTimestamp) internal {
        uint32 operatorSetLatestTimestamp = bn254CertificateVerifier.latestReferenceTimestamp(operatorSet);
        assertEq(operatorSetLatestTimestamp, expectedTimestamp, "Operator set latest timestamp should match expected value");
    }

    function check_ECDSAOperatorSetLatestTimestamp_State(OperatorSet memory operatorSet, uint32 expectedTimestamp) internal {
        uint32 operatorSetLatestTimestamp = ecdsaCertificateVerifier.latestReferenceTimestamp(operatorSet);
        assertEq(operatorSetLatestTimestamp, expectedTimestamp, "ECDSA operator set latest timestamp should match expected value");
    }

    /**
     *
     *                              GLOBAL TABLE ROOT CHECKS
     *
     */
    function check_GlobalTableRoot_PostSameTimestamp_Reverts(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        IBN254CertificateVerifierTypes.BN254Certificate memory confirmationCertificate
    ) internal {
        vm.expectRevert(abi.encodeWithSignature("GlobalTableRootStale()"));
        operatorTableUpdater.confirmGlobalTableRoot(confirmationCertificate, globalTableRoot, referenceTimestamp, uint32(block.number));
    }

    function check_GlobalTableRoot_PostSuccess_State(
        bytes32 globalTableRoot,
        uint32 referenceTimestamp,
        IBN254CertificateVerifierTypes.BN254Certificate memory confirmationCertificate
    ) internal {
        operatorTableUpdater.confirmGlobalTableRoot(confirmationCertificate, globalTableRoot, referenceTimestamp, uint32(block.number));
        check_LatestReferenceTimestamp_Updated_State(referenceTimestamp);
    }

    function check_TableTransport_StaleTimestamp_Reverts(uint32 staleTimestamp, bytes32 validGlobalTableRoot, bytes memory operatorTable)
        internal
    {
        check_TableTransport_StaleTimestamp_Reverts(staleTimestamp, validGlobalTableRoot, operatorTable, 0, new bytes(0));
    }

    function check_TableTransport_StaleTimestamp_Reverts(
        uint32 staleTimestamp,
        bytes32 validGlobalTableRoot,
        bytes memory operatorTable,
        uint32 operatorSetIndex,
        bytes memory proof
    ) internal {
        vm.expectRevert(abi.encodeWithSignature("TableUpdateForPastTimestamp()"));
        operatorTableUpdater.updateOperatorTable(staleTimestamp, validGlobalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    /**
     *
     *                              COMPLETE FLOW CHECKS
     *
     */
    function check_BN254MultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint32 referenceTimestamp
    ) internal {
        check_BN254MultichainFlow_Complete_State(operatorSet, certificate, referenceTimestamp, 2);
    }

    function check_BN254MultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint32 referenceTimestamp,
        uint expectedStakeTypes
    ) internal {
        // Verify certificate works
        check_BN254Certificate_Basic_State(operatorSet, certificate, expectedStakeTypes);

        // Verify timestamp is correctly set
        check_OperatorSetLatestTimestamp_State(operatorSet, referenceTimestamp);
    }

    function check_BN254MultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate,
        uint32 referenceTimestamp,
        uint expectedStakeTypes,
        uint[] memory expectedMinStakes
    ) internal returns (uint[] memory signedStakes) {
        // Verify certificate works and return signed stakes
        signedStakes = check_BN254Certificate_Basic_State(operatorSet, certificate, expectedStakeTypes, expectedMinStakes);

        // Verify timestamp is correctly set
        check_OperatorSetLatestTimestamp_State(operatorSet, referenceTimestamp);
    }

    function check_ECDSAMultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate,
        uint32 referenceTimestamp
    ) internal {
        check_ECDSAMultichainFlow_Complete_State(operatorSet, certificate, referenceTimestamp, 2, 1);
    }

    function check_ECDSAMultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate,
        uint32 referenceTimestamp,
        uint expectedStakeTypes,
        uint expectedMinSigners
    ) internal {
        // Verify certificate works
        check_ECDSACertificate_Basic_State(operatorSet, certificate, expectedStakeTypes, expectedMinSigners);

        // Verify timestamp is correctly set
        check_ECDSAOperatorSetLatestTimestamp_State(operatorSet, referenceTimestamp);
    }

    function check_ECDSAMultichainFlow_Complete_State(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate,
        uint32 referenceTimestamp,
        uint expectedStakeTypes,
        uint expectedMinSigners,
        uint[] memory expectedMinStakes
    ) internal returns (uint[] memory signedStakes, address[] memory signers) {
        // Verify certificate works and return signed stakes and signers
        (signedStakes, signers) =
            check_ECDSACertificate_Basic_State(operatorSet, certificate, expectedStakeTypes, expectedMinSigners, expectedMinStakes);

        // Verify timestamp is correctly set
        check_ECDSAOperatorSetLatestTimestamp_State(operatorSet, referenceTimestamp);
    }

    /**
     *
     *                              CERTIFICATE GENERATION HELPERS
     *
     */
    function check_CertificateGeneration_BN254_State(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        User[] memory operators,
        uint[] memory privateKeys
    ) internal returns (IBN254CertificateVerifierTypes.BN254Certificate memory certificate) {
        certificate = _generateBN254Certificate(operatorSet, referenceTimestamp, messageHash, operators, privateKeys);

        // Basic validation that certificate was generated correctly
        assertEq(certificate.referenceTimestamp, referenceTimestamp, "Certificate should have correct reference timestamp");
        assertEq(certificate.messageHash, messageHash, "Certificate should have correct message hash");
        // BN254 signature is a G1Point, so check that it's not the zero point
        assertTrue(certificate.signature.X != 0 || certificate.signature.Y != 0, "Certificate should have a valid signature");
    }

    function check_CertificateGeneration_ECDSA_State(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        User[] memory operators,
        uint[] memory privateKeys
    ) internal returns (IECDSACertificateVerifierTypes.ECDSACertificate memory certificate) {
        certificate = _generateECDSACertificate(operatorSet, referenceTimestamp, messageHash, operators, privateKeys);

        // Basic validation that certificate was generated correctly
        assertEq(certificate.referenceTimestamp, referenceTimestamp, "Certificate should have correct reference timestamp");
        assertEq(certificate.messageHash, messageHash, "Certificate should have correct message hash");
        assertTrue(certificate.sig.length > 0, "Certificate should have signatures");
    }

    function check_CertificateGeneration_ECDSA_State(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        User[] memory operators,
        uint[] memory privateKeys,
        uint expectedSigLength
    ) internal returns (IECDSACertificateVerifierTypes.ECDSACertificate memory certificate) {
        certificate = _generateECDSACertificate(operatorSet, referenceTimestamp, messageHash, operators, privateKeys);

        // Basic validation that certificate was generated correctly
        assertEq(certificate.referenceTimestamp, referenceTimestamp, "Certificate should have correct reference timestamp");
        assertEq(certificate.messageHash, messageHash, "Certificate should have correct message hash");
        assertEq(certificate.sig.length, expectedSigLength, "Certificate should have expected signature length");
    }

    /**
     *
     *                              GRANULAR VALIDATION HELPERS
     *
     */
    function check_SignedStakes_MinimumThresholds(uint[] memory signedStakes, uint[] memory minimumThresholds, string memory context)
        internal
        pure
    {
        require(signedStakes.length == minimumThresholds.length, "SignedStakes and thresholds length mismatch");

        for (uint i = 0; i < signedStakes.length; i++) {
            assertGe(
                signedStakes[i],
                minimumThresholds[i],
                string.concat(context, ": Stake type ", vm.toString(i), " should meet minimum threshold")
            );
        }
    }

    function check_SignedStakes_ExactValues(uint[] memory signedStakes, uint[] memory expectedValues, string memory context)
        internal
        pure
    {
        require(signedStakes.length == expectedValues.length, "SignedStakes and expected values length mismatch");

        for (uint i = 0; i < signedStakes.length; i++) {
            assertEq(
                signedStakes[i], expectedValues[i], string.concat(context, ": Stake type ", vm.toString(i), " should match expected value")
            );
        }
    }

    function check_SignedStakes_ProportionalToTotal(
        uint[] memory signedStakes,
        uint[] memory totalStakes,
        uint16[] memory expectedProportionsBPS,
        string memory context
    ) internal pure {
        require(signedStakes.length == totalStakes.length, "SignedStakes and totalStakes length mismatch");
        require(signedStakes.length == expectedProportionsBPS.length, "SignedStakes and proportions length mismatch");

        for (uint i = 0; i < signedStakes.length; i++) {
            uint expectedMinStake = (totalStakes[i] * expectedProportionsBPS[i]) / 10_000;
            assertGe(
                signedStakes[i],
                expectedMinStake,
                string.concat(context, ": Stake type ", vm.toString(i), " should meet proportional threshold")
            );
        }
    }

    function check_Signers_ValidAddresses(address[] memory signers, string memory context) internal pure {
        for (uint i = 0; i < signers.length; i++) {
            assertTrue(
                signers[i] != address(0), string.concat(context, ": Signer at index ", vm.toString(i), " should not be zero address")
            );
        }
    }

    function check_Signers_ExpectedAddresses(address[] memory signers, address[] memory expectedSigners, string memory context)
        internal
        pure
    {
        assertEq(signers.length, expectedSigners.length, string.concat(context, ": Signers length should match expected"));

        for (uint i = 0; i < signers.length; i++) {
            bool found = false;
            for (uint j = 0; j < expectedSigners.length; j++) {
                if (signers[i] == expectedSigners[j]) {
                    found = true;
                    break;
                }
            }
            assertTrue(found, string.concat(context, ": Signer ", vm.toString(signers[i]), " should be in expected signers list"));
        }
    }

    function check_Signers_MinimumCount(address[] memory signers, uint minimumCount, string memory context) internal pure {
        assertGe(signers.length, minimumCount, string.concat(context, ": Should have at least ", vm.toString(minimumCount), " signers"));
    }

    /**
     *
     *                              TRANSPORT AND UPDATE CHECKS
     *
     */
    function check_TableTransport_Success_State(OperatorSet memory operatorSet, uint32 referenceTimestamp, bytes32 expectedGlobalTableRoot)
        internal
    {
        // Verify that the global table root was updated correctly
        bytes32 currentGlobalTableRoot = operatorTableUpdater.getCurrentGlobalTableRoot();
        assertEq(currentGlobalTableRoot, expectedGlobalTableRoot, "Global table root should match expected value");

        // Verify that the latest reference timestamp was updated
        check_LatestReferenceTimestamp_Updated_State(referenceTimestamp);
    }

    function check_OperatorSetRegistration_State(OperatorSet memory operatorSet, IKeyRegistrarTypes.CurveType expectedCurveType)
        internal
        view
    {
        IKeyRegistrarTypes.CurveType actualCurveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertTrue(actualCurveType == expectedCurveType, "Operator set should be configured for expected curve type");
    }

    /**
     *
     *                              KEY REGISTRATION CHECKS
     *
     */
    function check_BN254KeyRegistration_State(User operator, OperatorSet memory operatorSet, BN254.G1Point memory expectedPubkey)
        internal
        view
    {
        (BN254.G1Point memory registeredG1Pubkey,) = keyRegistrar.getBN254Key(operatorSet, address(operator));
        assertTrue(
            registeredG1Pubkey.X == expectedPubkey.X && registeredG1Pubkey.Y == expectedPubkey.Y,
            "Registered BN254 pubkey should match expected"
        );
    }

    function check_ECDSAKeyRegistration_State(User operator, OperatorSet memory operatorSet, address expectedPubkey) internal view {
        address registeredPubkey = keyRegistrar.getECDSAAddress(operatorSet, address(operator));
        assertEq(registeredPubkey, expectedPubkey, "Registered ECDSA pubkey should match expected");
    }

    /**
     *
     *                              STALENESS PERIOD CHECKS
     *
     */
    function check_BN254Certificate_PostStaleness_Reverts(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate
    ) internal {
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        bn254CertificateVerifier.verifyCertificate(operatorSet, certificate);
    }

    function check_ECDSACertificate_PostStaleness_Reverts(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate
    ) internal {
        vm.expectRevert(abi.encodeWithSignature("CertificateStale()"));
        ecdsaCertificateVerifier.verifyCertificate(operatorSet, certificate);
    }

    function check_BN254Certificate_AtStalenessBoundary_Succeeds(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate
    ) internal {
        // Should succeed at the exact boundary
        check_BN254Certificate_Basic_State(operatorSet, certificate);
    }

    function check_ECDSACertificate_AtStalenessBoundary_Succeeds(
        OperatorSet memory operatorSet,
        IECDSACertificateVerifierTypes.ECDSACertificate memory certificate
    ) internal {
        // Should succeed at the exact boundary
        check_ECDSACertificate_Basic_State(operatorSet, certificate);
    }

    function check_BN254StalenessPeriod_Updated_State(OperatorSet memory operatorSet, uint32 expectedStalenessPeriod) internal view {
        uint32 actualStalenessPeriod = bn254CertificateVerifier.maxOperatorTableStaleness(operatorSet);
        assertEq(actualStalenessPeriod, expectedStalenessPeriod, "BN254 staleness period should match expected value");
    }

    function check_ECDSAStalenessPeriod_Updated_State(OperatorSet memory operatorSet, uint32 expectedStalenessPeriod) internal view {
        uint32 actualStalenessPeriod = ecdsaCertificateVerifier.maxOperatorTableStaleness(operatorSet);
        assertEq(actualStalenessPeriod, expectedStalenessPeriod, "ECDSA staleness period should match expected value");
    }

    function check_Certificate_WithinStalenessPeriod_Succeeds(
        OperatorSet memory operatorSet,
        IBN254CertificateVerifierTypes.BN254Certificate memory bn254Certificate,
        IECDSACertificateVerifierTypes.ECDSACertificate memory ecdsaCertificate,
        bool isBN254
    ) internal {
        if (isBN254) check_BN254Certificate_Basic_State(operatorSet, bn254Certificate);
        else check_ECDSACertificate_Basic_State(operatorSet, ecdsaCertificate);
    }

    function check_StalenessUpdate_CompleteFlow_State(
        OperatorSet memory operatorSet,
        uint32 initialStalenessPeriod,
        uint32 newStalenessPeriod,
        bool isBN254
    ) internal view {
        if (isBN254) check_BN254StalenessPeriod_Updated_State(operatorSet, newStalenessPeriod);
        else check_ECDSAStalenessPeriod_Updated_State(operatorSet, newStalenessPeriod);

        // Verify that the new staleness period is different from the initial one
        assertTrue(newStalenessPeriod != initialStalenessPeriod, "New staleness period should be different from initial");
    }

    /**
     *
     *                              TIME MANIPULATION HELPERS
     *
     */
    function check_TimeAdvance_PostStaleness_State(uint32 referenceTimestamp, uint32 stalenessPeriod, uint expectedBlockTimestamp)
        internal
        view
    {
        uint expectedStaleTimestamp = referenceTimestamp + stalenessPeriod;
        assertTrue(block.timestamp > expectedStaleTimestamp, "Should be past staleness period");
        assertEq(block.timestamp, expectedBlockTimestamp, "Block timestamp should match expected value");
    }

    function check_TimeAdvance_AtBoundary_State(uint32 referenceTimestamp, uint32 stalenessPeriod) internal view {
        uint expectedBoundaryTimestamp = referenceTimestamp + stalenessPeriod;
        assertEq(block.timestamp, expectedBoundaryTimestamp, "Should be exactly at staleness boundary");
    }

    /**
     *
     *                              COMBINED VALIDATION HELPERS
     *
     */
    function check_StalenessScenario_Complete_State(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        uint32 stalenessPeriod,
        bool certificateShouldSucceed,
        bool isBN254,
        string memory scenarioDescription
    ) internal view {
        // Check timing
        if (certificateShouldSucceed) {
            assertTrue(
                block.timestamp <= referenceTimestamp + stalenessPeriod,
                string.concat(scenarioDescription, ": Should be within staleness period")
            );
        } else {
            assertTrue(
                block.timestamp > referenceTimestamp + stalenessPeriod,
                string.concat(scenarioDescription, ": Should be past staleness period")
            );
        }

        // Check staleness period configuration
        if (isBN254) check_BN254StalenessPeriod_Updated_State(operatorSet, stalenessPeriod);
        else check_ECDSAStalenessPeriod_Updated_State(operatorSet, stalenessPeriod);
    }
}
