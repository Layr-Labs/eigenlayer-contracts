// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

import "src/test/integration/utils/MultiStakersLib.sol";

contract Integration_MultiStaker_Deposit_Delegate_Queue_Complete is IntegrationCheckUtils {
    using MultiStakersLib for User[];

    /*******************************************************************************
                                FULL WITHDRAWALS
    *******************************************************************************/

    // TODO: fix test
    /// Generates multiple random stakers and an operator. The stakers:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for ALL shares
    /// 4. completes the queued withdrawal as tokens
    function testFuzz_deposit_delegate_queue_completeAsTokens(uint24 _random) public rand(_random) {   
        /// 0. Create multiple random stakers with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated

        ( 
            User[] memory stakers,
            IStrategy[][] memory strategiesPerStaker,
            uint[][] memory tokenBalancesPerStaker
        ) = _newRandomStakers({numNewStakers: _randUint(1, 10)});
        (User operator, ,) = _newRandomOperator();

        uint[][] memory sharesPerStaker = _calculateExpectedShares(strategiesPerStaker, tokenBalancesPerStaker);

        // 1. Deposit Into Strategies
        stakers.depositIntoEigenlayer(strategiesPerStaker, tokenBalancesPerStaker);
        // check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        stakers.delegateTo(operator);
        // check_Delegation_State(staker, operator, strategies, shares);

        // 3. Queue Withdrawals
        Withdrawal[][] memory withdrawals = stakers.queueWithdrawals(strategiesPerStaker, sharesPerStaker);

        // 4. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals[0]);
        stakers.completeWithdrawalsAsTokens(withdrawals);
    }
}
