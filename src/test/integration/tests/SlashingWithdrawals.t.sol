// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract SlashingWithdrawals is IntegrationCheckUtils {
    
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;

    /// Shared setup:
    /// 
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator registers for operator set (default allocation delay)
    /// 5. Operator allocates to operator set
    function _init() internal override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // TODO invariant checks here
        operator.registerForOperatorSet(operatorSet);
        // TODO invariant checks here
        
        // 4. Allocate to operator set
        allocateParams = operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );
    }

    function testFuzz_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 4. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should once again have original token initTokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_slash_undelegate_completeAsShares(
        uint24 _random
    ) public rand(_random) {
        // 4. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_queue_slash_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 4. Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should once again have original token initTokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_queue_slash_completeAsShares(
        uint24 _random
    ) public rand(_random) {
        // 4. Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deallocate_slash_queue_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 4. Deallocate all.
        IAllocationManagerTypes.AllocateParams memory deallocateParams = operator.deallocateAll(operatorSet);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, deallocateParams, slashingParams, "staker deposit shares should be slashed");
        }
        
        // 6. Queue withdrawals
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, initTokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 7. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }
        
        // Check Final State
        assert_HasUnderlyingTokenBalances(
            staker,
            allocateParams.strategies,
            initTokenBalances,
            "staker should have withdrawn all shares"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deregister_slash(
        uint24 _random
    ) public rand(_random) {
        // 4. Deregister.
        operator.deregisterFromOperatorSet(operatorSet);

        // 5. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _randStrategiesAndWadsToSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }
    }
}