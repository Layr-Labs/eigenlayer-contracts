// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_ALMSlashBase is IntegrationCheckUtils {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens; // underlying token for each strategy
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// Shared setup:
    ///
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator allocates to operator set
    /// 5. Operator registers for operator set
    /// NOTE: Steps 4 and 5 are done in random order, as these should not have an outcome on the test
    function _init() internal virtual override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // randomly choose between:
        // register -> allocate / allocate -> register
        if (_randBool()) {
            // register -> allocate
            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_Slashable(operator, allocateParams);
        } else {
            // allocate -> register
            allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
            operator.modifyAllocations(allocateParams);
            check_IncrAlloc_State_NotSlashable(operator, allocateParams);

            operator.registerForOperatorSet(operatorSet);
            check_Registration_State_PendingAllocation(operator, allocateParams);
        }

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }
}

contract Integration_SlashThenWithdraw is Integration_ALMSlashBase {
    User stakerB;
    uint[] initTokenBalancesB;
    uint[] initTokenSharesB;

    User operatorB;
    OperatorSet operatorSetB;
    AllocateParams allocateParamsB;
    SlashingParams slashParamsB;

    /// Init: Registered operator gets slashed one or more times for a random magnitude
    /// - Second operator (for redelegation) may or may not be slashed
    ///
    /// Tests: Operator is slashed one or more times, then staker withdraws in different ways
    function _init() internal override {
        super._init();

        /// Create second operator set with the same strategies as the first
        /// Create a second operator. Register and allocate to operatorSetB
        /// Also create a second staker with the same assets as the first,
        /// delegated to operatorB. This is to give operatorB initial assets that can
        /// be checked via invariants.
        {
            // Create operatorB
            operatorB = _newRandomOperator_NoAssets();

            // Create stakerB, deposit, and delegate to operatorB
            (stakerB, initTokenBalancesB) = _newStaker(strategies);

            stakerB.depositIntoEigenlayer(strategies, initTokenBalancesB);
            initTokenSharesB = _calculateExpectedShares(strategies, initTokenBalancesB);
            check_Deposit_State(stakerB, strategies, initTokenSharesB);

            stakerB.delegateTo(operatorB);
            check_Delegation_State(stakerB, operatorB, strategies, initTokenSharesB);

            // Create operatorSetB
            operatorSetB = avs.createOperatorSet(strategies);

            // Register and allocate fully to operatorSetB
            operatorB.registerForOperatorSet(operatorSetB);
            check_Registration_State_NoAllocation(operatorB, operatorSetB, allStrats);

            allocateParamsB = _genAllocation_AllAvailable(operatorB, operatorSetB);
            operatorB.modifyAllocations(allocateParamsB);
            check_IncrAlloc_State_Slashable(operatorB, allocateParamsB);

            _rollBlocksForCompleteAllocation(operatorB, operatorSetB, strategies);
        }

        /// Slash first operator one or more times
        /// Each slash is for 1 to 99%
        {
            uint numSlashes = _randUint(1, 10);
            for (uint i = 0; i < numSlashes; i++) {
                slashParams = _genSlashing_Rand(operator, operatorSet);
                avs.slashOperator(slashParams);
                check_Base_Slashing_State(operator, allocateParams, slashParams);

                // TODO - staker variant?
                // assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
                // assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashParams, "staker deposit shares should be slashed");
            }
        }

        /// Optionally do a single slash for the second operator
        /// This is to test redelegation where the new operator has already been slashed
        {
            bool slashSecondOperator = _randBool();
            if (slashSecondOperator) {
                slashParamsB = _genSlashing_Half(operatorB, operatorSetB);
                avs.slashOperator(slashParamsB);
                check_Base_Slashing_State(operatorB, allocateParamsB, slashParamsB);
            }
        }
    }

    function testFuzz_undelegate_completeAsTokens(uint24 _r) public rand(_r) {
        /// Undelegate from operatorA
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        /// Complete withdrawal as tokens
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    function testFuzz_redelegate_completeAsTokens(uint24 _r) public rand(_r) {
        /// Redelegate to operatorB
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operatorB);
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operatorB, withdrawals, roots, strategies, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        /// Complete withdrawal as tokens
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    function testFuzz_queueFull_completeAsTokens(uint24 _r) public rand(_r) {
        // Queue a withdrawal for all shares
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawal as tokens
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
    }

    function testFuzz_undelegate_completeAsShares(uint24 _r) public rand(_r) {
        // Undelegate from operatorA
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawal as shares
        staker.completeWithdrawalsAsShares(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_redelegate_completeAsShares(uint24 _r) public rand(_r) {
        // Redelegate to operatorB
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operatorB);
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operatorB, withdrawals, roots, strategies, shares);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawal as shares
        staker.completeWithdrawalsAsShares(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsShares_Redelegated_State(staker, operator, operatorB, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_queueFull_completeAsShares(uint24 _r) public rand(_r) {
        // Queue a withdrawal for all shares
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawal as shares
        staker.completeWithdrawalsAsShares(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, withdrawableShares);
        }
    }
}

contract Integration_QueueWithdrawalThenSlash is Integration_ALMSlashBase {
    function testFuzz_queue_slash_completeAsTokens(uint24 _r) public rand(_r) {
        // 4. Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // 5. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        // TODO - staker variants?

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, expectedShares, tokens, expectedTokens);
        }

        // Check Final State
        // check_FullyWithdrawn_State(staker, ..., ); TODO
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_queue_slash_completeAsShares(uint24 _r) public rand(_r) {
        // 4. Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // 5. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        // TODO - staker variants?

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, expectedShares);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}

contract Integration_DeallocateThenSlash is Integration_ALMSlashBase {
    function testFuzz_deallocate_slash_queue_completeAsTokens(uint24 _r) public rand(_r) {
        // 4. Deallocate all.
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        _rollForward_DeallocationDelay();

        // 5. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, deallocateParams, slashParams);
        // TODO - staker variants?

        // 6. Queue withdrawals
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // 7. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, initDepositShares, tokens, expectedTokens);
        }

        // Check Final State
        assert_HasUnderlyingTokenBalances(staker, allocateParams.strategies, initTokenBalances, "staker should have withdrawn all shares");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deregister_slash(uint24 _r) public rand(_r) {
        // 4. Deregister.
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_PendingAllocation(operator, operatorSet);

        // 5. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        // TODO - staker variants?
    }
}
