// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Queue is IntegrationCheckUtils {

    // function _newEmptyStaker() internal returns (User) {
    //     string memory stakerName;

    //     User staker
    // }

    function testFuzz_deposit_queue(uint24 _r) public rand(_r) {
        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // // Perform 50 random operations:
        // // 1. Deposit some amount
        // // 2. Queue/complete some amount
        // for (uint i = 0; i < 50; i++) {

        // }


        // IStrategy strategy = _selectSingle(strategies);

        
        

        // // 3. Queue Withdrawal
        // IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        // bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        // check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawals, withdrawalRoots);

        // // 4. Complete Queued Withdrawal
        // _rollBlocksForCompleteWithdrawals(withdrawals);
        // for (uint i = 0; i < withdrawals.length; i++) {
        //     staker.completeWithdrawalAsShares(withdrawals[i]);
        //     check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        // }        
    }
}