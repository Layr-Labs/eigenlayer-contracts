// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Register_QueueWithdrawal_Complete is IntegrationChecks {
    function testFuzz_deposit_registerOperator_queueWithdrawal_completeAsShares(uint24) public {
        // Create a staker with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();

        // 1. Staker deposits into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Staker registers as an operator
        staker.registerAsOperator();
        assertTrue(delegationManager.isOperator(address(staker)), "Staker should be registered as an operator");

        // 3. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, staker, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal as Shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, staker, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_deposit_registerOperator_queueWithdrawal_completeAsTokens(uint24) public {
        // Create a staker with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        // 1. Staker deposits into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Staker registers as an operator
        staker.registerAsOperator();
        assertTrue(delegationManager.isOperator(address(staker)), "Staker should be registered as an operator");

        // 3. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, staker, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal as Tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, staker, withdrawals[i], shares, expectedTokens);
        }
    }
}
