// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_EigenPod_SlashedEigenpod_BC is IntegrationChecks {
    using ArrayLib for *;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();
        avs = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies); // Should only return ETH

        // Ensure the staker has at least 64 ETH to deposit.
        if (initTokenBalances[0] < 64 ether) {
            initTokenBalances[0] = 64 ether;
            cheats.deal(address(staker), 64 ether);
        }

        // Deposit staker
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);
        validators = staker.getActiveValidators();

        //Slash on Beacon chain
        slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();

        // Checkpoint post slash
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, slashedGwei);
    }

    function testFuzz_delegateSlashedStaker_dsfWad(uint24) public {
        uint[] memory initDelegatableShares = _getWithdrawableShares(staker, strategies);
        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDelegatableShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, initDelegatableShares);
        }

        (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
            delegationManager().getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDelegatableShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(
            withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal"
        );
    }

    function testFuzz_delegateSlashedStaker_dsfNonWad(uint24) public {
        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);

        uint[] memory initDelegatableShares = _getWithdrawableShares(staker, strategies);
        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDelegatableShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, initDelegatableShares);
        }

        (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
            delegationManager().getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDelegatableShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(
            withdrawableSharesAfter[0], depositSharesAfter[0], 1000, "Withdrawable shares should equal deposit shares after withdrawal"
        );
    }

    function testFuzz_delegateSlashedStaker_slashedOperator(uint24) public {
        (User staker2,,) = _newRandomStaker();
        (uint40[] memory validators2,,) = staker2.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker2.verifyWithdrawalCredentials(validators2);
        staker2.startCheckpoint();
        staker2.completeCheckpoint();
        staker2.delegateTo(operator);

        //randomize additional deposit to eigenpod
        if (_randBool()) {
            uint amount = 32 ether * _randUint({min: 1, max: 5});
            cheats.deal(address(staker), amount);
            (uint40[] memory newValidators,,) = staker.startValidators();
            beaconChain.advanceEpoch_NoWithdrawNoRewards();
            staker.verifyWithdrawalCredentials(newValidators);

            staker.startCheckpoint();
            staker.completeCheckpoint();
        }

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        //Slash operator before delegation
        SlashingParams memory slashingParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashingParams);
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        (uint[] memory delegatedShares,) = delegationManager().getWithdrawableShares(address(staker), strategies);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatedShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, delegatedShares);
        }

        (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
            delegationManager().getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(
            withdrawableSharesAfter[0], depositSharesAfter[0], 1000, "Withdrawable shares should equal deposit shares after withdrawal"
        );
    }

    function testFuzz_delegateSlashedStaker_redelegate_complete(uint24) public {
        User operator2 = _newRandomOperator();

        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);

        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
        (uint[] memory delegatedShares,) = delegationManager().getWithdrawableShares(address(staker), strategies);

        // Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operator2, withdrawals, withdrawalRoots, strategies, delegatedShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(
                staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, delegatedShares
            );
        }

        (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
            delegationManager().getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(
            withdrawableSharesAfter[0], depositSharesAfter[0], 1000, "Withdrawable shares should equal deposit shares after withdrawal"
        );
    }

    function testFuzz_delegateSlashedStaker_slashedOperator_withdrawAllShares_complete(uint24) public {
        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);

        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        // Slash operator before delegation
        SlashingParams memory slashingParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashingParams);
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
        (uint[] memory delegatedShares,) = delegationManager().getWithdrawableShares(address(staker), strategies);

        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // Withdraw all shares
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawableShares);
        }

        (uint[] memory withdrawableSharesAfter, uint[] memory depositSharesAfter) =
            delegationManager().getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(
            withdrawableSharesAfter[0], depositSharesAfter[0], 1000, "Withdrawable shares should equal deposit shares after withdrawal"
        );
    }

    function testFuzz_redeposit_queue_completeAsTokens(uint24) public {
        // Prove an additional validator
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);

        // Queue withdrawal for all tokens
        uint[] memory depositShares = _getStakerDepositShares(staker, strategies);
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, withdrawalRoots
        );

        // Complete withdrawal as tokens
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawableShares, expectedTokens);
        }
    }

    function testFuzz_redeposit_queue_completeAsShares(uint24) public {
        // Prove an additional validator
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);

        // Queue withdrawal for all
        uint[] memory depositShares = _getStakerDepositShares(staker, strategies);
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositShares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(
            staker, User(payable(address(0))), strategies, depositShares, withdrawableShares, withdrawals, withdrawalRoots
        );

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawableShares);
        }
    }

    function testFuzz_verifyAdditionalValidator_delegateSlashedStaker(uint24) public {
        // Create operatorSet
        operatorSet = avs.createOperatorSet(strategies);

        // Allocate to operatorSet
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_Base_IncrAlloc_State(operator, allocateParams); // No increase in allocated stake since staker not delegated
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // Register to opSet
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_PendingAllocation(operator, allocateParams);

        // Slash operator - should have non-WAD magnitude
        SlashingParams memory slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // Verify additional validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);

        // Delegate to operator
        uint[] memory initDepositShares = _getStakerDepositShares(staker, strategies); // Deposit shares increased after verifying validator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
    }
}

/// @notice This is not considered dual slashing since the operator is pre-slashed
contract Integration_EigenPod_SlashedOperator_SlashedEigenpod_Base is IntegrationChecks {
    using ArrayLib for *;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();
        avs = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // Steps 3-4-5
        // 3. Allocate to operatorSet
        // 4. Register to opSet
        // 5. Slash operator
        allocateParams = _giveOperatorNonWadMagnitude(operator);
    }

    /// @dev Assumes operatorSet has been initialized with relevant strategies for staker
    function _giveOperatorNonWadMagnitude(User _operator) internal returns (AllocateParams memory _allocateParams) {
        // Allocate to operatorSet
        _allocateParams = _genAllocation_AllAvailable(_operator, operatorSet);
        _operator.modifyAllocations(_allocateParams);
        check_Base_IncrAlloc_State(_operator, _allocateParams); // No increase in allocated stake since staker not delegated
        _rollBlocksForCompleteAllocation(_operator, operatorSet, strategies);

        // Register to opSet
        _operator.registerForOperatorSet(operatorSet);
        check_Registration_State_PendingAllocation(_operator, _allocateParams);

        // Slash operator
        SlashingParams memory slashParams = _genSlashing_Rand(_operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(_operator, _allocateParams, slashParams);
    }
}

contract Integration_EigenPod_SlashedOperator_SlashedEigenpod is Integration_EigenPod_SlashedOperator_SlashedEigenpod_Base {
    function _init() internal virtual override {
        // 1-5
        super._init();

        // 6. Slash on BC
        uint40[] memory validators = staker.getActiveValidators();
        uint40[] memory slashedValidators = _choose(validators);
        uint64 slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoRewards();
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, slashedValidators, slashedGwei);

        // 7. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 8. Verify Additional Validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);
        initDepositShares = _getStakerDepositShares(staker, strategies); // Deposit shares increased after verifying validator
    }

    function testFuzz_slashOnBC_delegate_verifyValidator_undelegate_completeAsTokens(uint24) public {
        // 9. Undelegate
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 10. Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawableShares, expectedTokens);
        }
    }

    function testFuzz_slashOnBC_delegate_verifyValidator_undelegate_completeAsShares(uint24) public {
        // 9. Undelegate
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 10. Complete withdrawal as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawableShares);
        }
    }

    function testFuzz_slashOnBC_delegate_verifyValidator_redelegate_completeAsTokens(uint24) public {
        // Create new operator
        User operator2 = _newRandomOperator();

        // Randomly give the operator a non-WAD magnitude
        if (_randBool()) _giveOperatorNonWadMagnitude(operator2);

        // 9. Redelegate
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operator2, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 10. Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], withdrawableShares, expectedTokens);
        }
    }

    function testFuzz_slashOnBC_delegate_verifyValidator_redelegate_completeAsShares(uint24) public {
        // Create new operator
        User operator2 = _newRandomOperator();

        // Randomly give the operator a non-WAD magnitude
        if (_randBool()) _giveOperatorNonWadMagnitude(operator2);

        // 9. Redelegate
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operator2, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 10. Complete withdrawal as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(
                staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, withdrawableShares
            );
        }
    }
}

contract Integration_EigenPod_Redelegate_SlashOperator_SlashEigenpod is Integration_EigenPod_SlashedOperator_SlashedEigenpod_Base {
    using ArrayLib for *;

    User operator2;
    Withdrawal withdrawal;

    function _init() internal virtual override {
        // 1-5.
        super._init();

        // 6. Delegate staker
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 7. Create new operator & randomly give it a non-WAD magnitude
        operator2 = _newRandomOperator();
        if (_randBool()) _giveOperatorNonWadMagnitude(operator2);

        // 8. Redelegate
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        require(withdrawals.length == 1, "Expected 1 withdrawal");
        withdrawal = withdrawals[0];
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operator2, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 9. Slash original operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // 10. Slash on BC
        uint40[] memory validators = staker.getActiveValidators();
        uint40[] memory slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.startCheckpoint();
        staker.completeCheckpoint();
    }

    function testFuzz_delegate_redelegate_slashAVS_slashBC_verifyValidator_completeAsShares(uint24) public {
        // Populate withdrawal array
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;

        // 11. Verify additional validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);

        // 12. Complete withdrawal as shares
        uint[] memory withdrawableShares = _getWithdrawableSharesAfterCompletion(staker);
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(
                staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, withdrawableShares
            );
        }
    }

    function testFuzz_delegate_redelegate_slashAVS_slashBC_verifyValidator_completeAsTokens(uint24) public {
        // Populate withdrawal array
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;

        // 11. Verify additional validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);

        // 12. Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        uint[] memory withdrawableShares = _getWithdrawableSharesAfterCompletion(staker);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], withdrawableShares, expectedTokens);
        }
    }

    function testFuzz_delegate_redelegate_slashAVS_slashBC_completeAsShares_verifyValidator(uint24) public {
        // Populate withdrawal array
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;

        // 11. Complete withdrawal as shares
        uint[] memory withdrawableShares = _getWithdrawableSharesAfterCompletion(staker);
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(
                staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, withdrawableShares
            );
        }

        // 12. Verify additional validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);
    }

    function testFuzz_delegate_redelegate_slashAVS_slashBC_completeAsTokens_verifyValidator(uint24) public {
        // Populate withdrawal array
        Withdrawal[] memory withdrawals = new Withdrawal[](1);
        withdrawals[0] = withdrawal;

        // 11. Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        uint[] memory withdrawableShares = _getWithdrawableSharesAfterCompletion(staker);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawableShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], withdrawableShares, expectedTokens);
        }

        // 12. Verify additional validator
        cheats.deal(address(staker), 32 ether);
        uint64 beaconBalanceAddedGwei = uint64(32 ether / GWEI_TO_WEI);
        (uint40[] memory newValidators,,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, beaconBalanceAddedGwei);
    }
}

contract Integration_EigenPod_SlashedEigenpod_BC_HalfSlash is IntegrationChecks {
    using ArrayLib for *;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();
        avs = _newRandomAVS();

        // 1. Deposit staker
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        initDepositShares = shares;
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);
        uint40[] memory validators = staker.getActiveValidators();

        // 2. Slash on Beacon chain
        slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Half);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();

        // 3. Checkpoint post slash
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, slashedGwei);
    }

    /**
     * @notice Test sets up an EigenPod which has a non-WAD BCSF. After queue withdrawing all depositShares
     * which sets it to 0, they can then complete checkpoints repeatedly with 0 shares increase to increase the staker DSF each time
     */
    function test_completeCP_withNoAddedShares() public {
        // 4. queue withdraw all depositShares having it set to 0
        uint withdrawableSharesBefore = _getStakerWithdrawableShares(staker, strategies)[0];
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);

        // 5. advance epoch no rewards
        // start/complete cp repeatedly with 0 shares increase to increase the staker DSF each time
        for (uint i = 0; i < 10; i++) {
            beaconChain.advanceEpoch_NoWithdrawNoRewards();
            staker.startCheckpoint();
            staker.completeCheckpoint();
        }

        // Staker deposit shares should be 0 from queue withdrawing all depositShares
        // therefore the depositScalingFactor should also be reset WAD
        assertEq(eigenPodManager().podOwnerDepositShares(address(staker)), 0);
        assertEq(delegationManager().depositScalingFactor(address(staker), BEACONCHAIN_ETH_STRAT), WAD);

        // 6. deposit: can either verify wc or start/complete cp or complete the withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.completeWithdrawalsAsShares(withdrawals);

        // 7. delegateTo an operator
        staker.delegateTo(operator);
        // End state: staker and operator have much higher inflated withdrawable and delegated shares respectively
        // The staker's withdrawable shares should be <= from withdrawable shares before (should be equal but could be less due to rounding)
        uint withdrawableSharesAfter = _getStakerWithdrawableShares(staker, strategies)[0];
        uint operatorShares = delegationManager().operatorShares(address(operator), strategies[0]);
        assertLe(
            withdrawableSharesAfter, withdrawableSharesBefore, "staker withdrawable shares should be <= from withdrawable shares before"
        );
        assertLe(operatorShares, withdrawableSharesBefore, "operatorShares should be <= from withdrawable shares before");
    }
}
