// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";

/// @notice Contract that provides utility functions to reuse common test blocks & checks
contract IntegrationCheckUtils is IntegrationBase {
    
    function check_Deposit_State(
        User staker, 
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal {
        /// Deposit into strategies:
        // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
        // the staker calls the relevant deposit function, depositing all held assets.
        //
        // ... check that all underlying tokens were transferred to the correct destination
        //     and that the staker now has the expected amount of delegated shares in each strategy
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker should have transferred all underlying tokens");
        assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
    }

    function check_Delegation_State(
        User staker, 
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal {
        /// Delegate to an operator:
        //
        // ... check that the staker is now delegated to the operator, and that the operator
        //     was awarded the staker shares
        assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
        assert_Snap_Unchanged_StakerShares(staker, "staker shares should be unchanged after delegating");
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
    }

    function check_QueuedWithdrawal_State(
        User staker, 
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory shares, 
        IDelegationManager.Withdrawal[] memory withdrawals, 
        bytes32[] memory withdrawalRoots
    ) internal {
        // The staker will queue one or more withdrawals for the selected strategies and shares
        //
        // ... check that each withdrawal was successfully enqueued, that the returned roots
        //     match the hashes of each withdrawal, and that the staker and operator have
        //     reduced shares.
        assert_AllWithdrawalsPending(withdrawalRoots, "staker withdrawals should now be pending");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
        assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
        assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
    }

    function check_Undelegate_State(
        User staker, 
        User operator, 
        IDelegationManager.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        IStrategy[] memory strategies,
        uint[] memory shares 
    ) internal {
        /// Undelegate from an operator
        //
        // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqeuued,
        //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
        //     and operator have reduced shares
        assertEq(withdrawalRoots.length, 1, "should only be one withdrawal root");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawl should match returned root");
        assert_AllWithdrawalsPending(withdrawalRoots, "stakers withdrawal should now be pending");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by 1");
        assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
        assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
    }

    function check_Withdrawal_AsTokens_State(
        User staker,
        User operator,
        IDelegationManager.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares,
        IERC20[] memory tokens,
        uint[] memory expectedTokens
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as tokens
        // 
        // ... check that the withdrawal is not pending, that the withdrawer received the expected tokens, and that the total shares of each 
        //     strategy withdrawn decreases
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
        assert_Snap_Unchanged_StakerShares(staker, "staker shares should not have changed");
        assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
        assert_Snap_Removed_StrategyShares(strategies, shares, "strategies should have total shares decremented");
    }

    function check_Withdrawal_AsShares_State(
        User staker,
        User operator,
        IDelegationManager.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as shares
        // 
        // ... check that the withdrawal is not pending, that the withdrawer received the expected shares, and that the total shares of each 
        //     strategy withdrawn remains unchanged 
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
        assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should have received expected shares");
        assert_Snap_Added_OperatorShares(operator, withdrawal.strategies, withdrawal.shares, "operator should have received shares");
        assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");
    }

    /// @notice Difference from above is that operator shares do not increase since staker is not delegated
    function check_Withdrawal_AsShares_Undelegated_State(
        User staker,
        User operator,
        IDelegationManager.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as shares
        // 
        // ... check that the withdrawal is not pending, that the token balances of the staker and operator are unchanged,
        //     that the withdrawer received the expected shares, and that that the total shares of each o
        //     strategy withdrawn remains unchanged 
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
        assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should have received expected shares");
        assert_Snap_Unchanged_OperatorShares(operator, "operator should have shares unchanged");
        assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");
    }
}