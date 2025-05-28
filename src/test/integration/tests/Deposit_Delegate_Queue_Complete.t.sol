// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Deposit_Delegate_Queue_Complete is IntegrationChecks {
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
    function testFuzz_deposit_delegate_queue_completeAsTokens(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], shares, expectedTokens);
        }
    }

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a ALL shares
    /// 4. completes the queued withdrawal as shares
    function testFuzz_deposit_delegate_queue_completeAsShares(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
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
    function testFuzz_deposit_delegate_revert_alreadyDelegated(uint24) public {
        _configAssetTypes(NO_ASSETS | HOLDS_LST | HOLDS_ETH | HOLDS_ALL);

        /// 0. Create a staker and operator
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Register staker as an operator
        staker.registerAsOperator();

        // 3. Attempt to delegate to an operator
        //    This should fail as the staker is already delegated to themselves.
        cheats.expectRevert();
        staker.delegateTo(operator);
    }
}
