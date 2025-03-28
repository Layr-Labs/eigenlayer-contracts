// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Delegate_Deposit_Queue_Complete is IntegrationCheckUtils {
    function testFuzz_delegate_deposit_queue_completeAsShares(uint24 _random) public rand(_random) {
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        // Check that the deposit increased operator shares the staker is delegated to
        check_Deposit_State(staker, strategies, shares);
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");

        // 3. Queue Withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_delegate_deposit_queue_completeAsTokens(uint24 _random) public rand(_random) {
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);

        // Check that the deposit increased operator shares the staker is delegated to
        check_Deposit_State(staker, strategies, shares);
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");

        // 3. Queue Withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }
}
