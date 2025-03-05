// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Deposit_Delegate_Queue_Complete is IntegrationCheckUtils {
    /**
     *
     *                             FULL WITHDRAWALS
     *
     */

    // TODO: fix test
    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a ALL shares
    /// 4. completes the queued withdrawal as tokens
    function testFuzz_deposit_delegate_queue_completeAsTokens(uint24 _random) public rand(_random) {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "staker should once again have original token balances");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a ALL shares
    /// 4. completes the queued withdrawal as shares
    function testFuzz_deposit_delegate_queue_completeAsShares(uint24 _random) public rand(_random) {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     *
     *                           RANDOM WITHDRAWALS
     *
     */

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a random subset of shares
    /// 4. completes the queued withdrawal as tokens
    function testFuzz_deposit_delegate_queueRand_completeAsTokens(uint24 _random) public rand(_random) {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        // Randomly select one or more assets to withdraw
        (IStrategy[] memory withdrawStrats, uint[] memory withdrawShares) = _randWithdrawal(strategies, shares);

        Withdrawal[] memory withdrawals = staker.queueWithdrawals(withdrawStrats, withdrawShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, withdrawStrats, withdrawShares, withdrawShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawals
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawStrats, withdrawShares, tokens, expectedTokens);
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a random subset of shares
    /// 4. completes the queued withdrawal as shares
    function testFuzz_deposit_delegate_queueRand_completeAsShares(uint24 _random) public rand(_random) {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        // Randomly select one or more assets to withdraw
        (IStrategy[] memory withdrawStrats, uint[] memory withdrawShares) = _randWithdrawal(strategies, shares);

        Withdrawal[] memory withdrawals = staker.queueWithdrawals(withdrawStrats, withdrawShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, withdrawStrats, withdrawShares, withdrawShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], withdrawStrats, withdrawShares);
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     *
     *                            UNHAPPY PATH TESTS
     *
     */

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// --- registers as an operator
    /// 2. delegates to an operator
    ///
    /// ... we check that the final step fails
    function testFuzz_deposit_delegate_revert_alreadyDelegated(uint24 _random) public rand(_random) {
        _configAssetTypes(NO_ASSETS | HOLDS_LST | HOLDS_ETH | HOLDS_ALL);

        /// 0. Create a staker and operator
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Register staker as an operator
        staker.registerAsOperator();
        assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");

        // 3. Attempt to delegate to an operator
        //    This should fail as the staker is already delegated to themselves.
        cheats.expectRevert();
        staker.delegateTo(operator);
    }
}
