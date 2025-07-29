// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Delegate_Deposit_Queue_Complete is IntegrationChecks {
    function _init() internal override {
        _configAssetTypes(HOLDS_LST);
    }

    function testFuzz_delegate_deposit_queue_completeAsShares(uint24) public {
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 3. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_delegate_deposit_queue_completeAsTokens(uint24) public {
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
        check_Deposit_State(staker, strategies, shares);

        // 3. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], shares, expectedTokens);
        }
    }
}
