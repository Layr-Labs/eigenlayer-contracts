// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Delegate_Redelegate_Complete is IntegrationCheckUtils {
    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete queued withdrawal as shares 
    /// 5. delegate to a new operator
    /// 5. queueWithdrawal
    /// 7. complete their queued withdrawal as tokens
    function testFuzz_deposit_delegate_reDelegate_completeAsTokens(uint24 _random) public {   
        // When new Users are created, they will choose a random configuration from these params: 
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        (
            User staker,
            IStrategy[] memory strategies, 
            uint[] memory tokenBalances
        ) = _newRandomStaker();
        (User operator1, ,) = _newRandomOperator();
        (User operator2, ,) = _newRandomOperator();
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        check_Delegation_State(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, shares);

        // 4. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator1, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares);
        }

        // 5. Delegate to a new operator
        staker.delegateTo(operator2);
        check_Delegation_State(staker, operator2, strategies, shares);
        assertNotEq(address(operator1), delegationManager.delegatedTo(address(staker)), "staker should not be delegated to operator1");

        // 6. Queue Withdrawal
        withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator2, strategies, shares, withdrawals, withdrawalRoots);

        // 7. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        // Complete withdrawals
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].shares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares, tokens, expectedTokens);
        }
    }

    //TODO: add complete last withdrawal as shares
    //TODO: add complete middle withdrawal as tokens and then restake and redelegate
    //TODO: additional deposit before delegating to new operator
    //TODO: additional deposit after delegating to new operator
}