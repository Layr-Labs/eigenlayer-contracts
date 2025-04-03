// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_QueueWithdrawal_Complete is IntegrationChecks {
    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. queueWithdrawal
    /// 3. completeQueuedWithdrawal"
    function testFuzz_deposit_queueWithdrawal_completeAsTokens(uint24) public {
        // Create a staker with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // Ensure staker is not delegated to anyone post deposit
        assertFalse(delegationManager().isDelegated(address(staker)), "Staker should not be delegated after deposit");

        // 2. Queue Withdrawal
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        // 3. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, User(payable(0)), withdrawals[i], shares, expectedTokens);
        }

        // Ensure staker is still not delegated to anyone post withdrawal completion
        assertFalse(delegationManager().isDelegated(address(staker)), "Staker should still not be delegated after withdrawal");
    }

    function testFuzz_deposit_queueWithdrawal_completeAsShares(uint24) public {
        // Create a staker with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // Ensure staker is not delegated to anyone post deposit
        assertFalse(delegationManager().isDelegated(address(staker)), "Staker should not be delegated after deposit");

        // 2. Queue Withdrawal
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        // 3. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, User(payable(0)), withdrawals[i], strategies, shares);
        }

        // Ensure staker is still not delegated to anyone post withdrawal completion
        assertFalse(delegationManager().isDelegated(address(staker)), "Staker should still not be delegated after withdrawal");
    }
}
