// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../MultichainIntegrationChecks.t.sol";

/**
 * @title Multichain_Generation_Reservation_Removal
 * @notice Integration tests for generation reservation removal functionality
 * @dev Tests the behavior when generation reservations are removed and tables can no longer be transported
 */
contract Multichain_Generation_Reservation_Removal is MultichainIntegrationCheckUtils {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    /**
     * @notice Test that tables cannot be transported after removing generation reservation
     * @dev Test case 2: Verify that after removing generation reservation, tables are not transported
     */
    function test_RemoveGenerationReservation_TablesNotTransported() external {
        console.log("Testing generation reservation removal - tables not transported:");
        vm.warp(50_000);

        // Setup test environment with staker having nonzero shares
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);

        // Create staker with shares
        User staker;
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;
        (staker, strategies, tokenBalances) = _newRandomStaker();

        // Deposit tokens to get nonzero shares
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory expectedShares = _calculateExpectedShares(strategies, tokenBalances);
        assert_HasExpectedShares(staker, strategies, expectedShares, "Staker should have nonzero shares");

        // Create operator and register
        User operator = _newRandomOperator_NoAssets();

        // Delegate staker to operator
        staker.delegateTo(operator);
        assertEq(delegationManager.delegatedTo(address(staker)), address(operator), "Staker should be delegated to operator");

        // Create and configure operator set
        _createTestOperatorSet(1);
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: 1});

        // Configure operator set for BN254 curve (randomly choose BN254 for this test)
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        // Setup AVS and chains
        _setupAVSAndChains();

        // Register operator key
        User[] memory operators = new User[](1);
        operators[0] = operator;
        _registerSingleBN254Key(operator, operatorSet, _getBN254PrivateKeys()[0]);

        // Operator allocation (simulate allocation to the operator set)
        // This would typically involve AllocationManager operations

        // Create generation reservation
        _createGenerationReservation(operatorSet);

        // Verify generation reservation exists
        check_GenerationReservation_Exists_State(operatorSet);

        // Test that we can generate operator table bytes before removal
        check_OperatorTableBytes_Generation_State(operatorSet, true);
        console.log("Successfully generated operator table bytes before removing generation reservation");

        // Generate operator table and transport it successfully
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = _generateBN254OperatorTable(operatorSet, operators);
        uint32 referenceTimestamp = uint32(block.timestamp - 1000);
        _confirmGlobalTableRootAndUpdateBN254(operatorSet, operatorSetInfo, referenceTimestamp);
        console.log("Successfully transported tables before removing generation reservation");

        // Remove generation reservation
        vm.prank(address(this)); // AVS owner removes the reservation
        crossChainRegistry.removeGenerationReservation(operatorSet);
        console.log("Removed generation reservation");

        // Verify generation reservation no longer exists
        check_GenerationReservation_NotExists_State(operatorSet);

        // Try to generate operator table bytes after removal - this should fail
        // because the operator table calculator and config have been deleted
        check_OperatorTableBytes_Generation_State(operatorSet, false);
        console.log("Successfully verified that operator table bytes cannot be generated after removing generation reservation");

        // Verify that operator set config and calculator are cleared
        check_OperatorSetConfig_Cleared_State(operatorSet);

        console.log("State validation passed: All operator set data cleared after generation reservation removal");
        console.log("Tables cannot be transported because source chain data is no longer available");
    }
}
