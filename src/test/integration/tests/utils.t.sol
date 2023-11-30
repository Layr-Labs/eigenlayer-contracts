// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";

/// @notice Contract that provides utility functions to reuse common test blocks & checks
contract IntegrationTestUtils is IntegrationBase {
    
    function assertDepositState(User staker, IStrategy[] memory strategies, uint256[] memory shares) internal {
        /// Deposit into strategies:
        // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
        // the staker calls the relevant deposit function, depositing all held assets.
        //
        // ... check that all underlying tokens were transferred to the correct destination
        //     and that the staker now has the expected amount of delegated shares in each strategy
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker should have transferred all underlying tokens");
        assert_Snap_AddedStakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
    }

    function assertDelegationState(User staker, User operator, IStrategy[] memory strategies, uint256[] memory shares) internal {
        /// Delegate to an operator:
        //
        // ... check that the staker is now delegated to the operator, and that the operator
        //     was awarded the staker's shares
        assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
        assert_Snap_AddedOperatorShares(operator, strategies, shares, "operator should have received shares");
    }

    function assertQueuedWithdrawalState(User staker, User operator, IStrategy[] memory strategies, uint256[] memory shares, IDelegationManager.Withdrawal[] memory withdrawals, bytes32[] memory withdrawalRoots) internal {
        // The staker will queue one or more withdrawals for the selected strategies and shares
        //
        // ... check that each withdrawal was successfully enqueued, that the returned roots
        //     match the hashes of each withdrawal, and that the staker and operator have
        //     reduced shares.
        assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
        assert_Snap_IncreasedQueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
        assert_Snap_RemovedOperatorShares(operator, strategies, shares, "failed to remove operator shares");
        assert_Snap_RemovedStakerShares(staker, strategies, shares, "failed to remove staker shares");
    }

    function assertUndelegateState(
        User staker, 
        User operator, 
        IDelegationManager.Withdrawal memory withdrawal,
        bytes32 withdrawalRoot,
        IStrategy[] memory strategies,
        uint256[] memory shares 
    ) internal {
        /// Undelegate from an operator
        //
        // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqeuued,
        //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
        //     and operator have reduced shares
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");
        assert_ValidWithdrawalHash(withdrawal, withdrawalRoot, "calculated withdrawl should match returned root");
        assert_withdrawalPending(withdrawalRoot, "staker's withdrawal should now be pending");
        assert_Snap_IncrementQueuedWithdrawals(staker, "staker should have increased nonce by 1");
        assert_Snap_RemovedOperatorShares(operator, strategies, shares, "failed to remove operator shares");
        assert_Snap_RemovedStakerShares(staker, strategies, shares, "failed to remove staker shares");
    }

    function assertWithdrawalAsTokensState(
        User staker,
        IDelegationManager.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint256[] memory shares,
        IERC20[] memory tokens,
        uint256[] memory expectedTokens
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as tokens
        // 
        // ... check that the withdrawal is not pending, that the withdrawer received the expected tokens, and that the total shares of each 
        //     strategy withdrawn decreases
        assert_withdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker's withdrawal should no longer be pending");
        assert_Snap_IncreasedTokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_DecreasedStrategyShares(strategies, shares, "strategies should have total shares decremented");
    }

    function assertWithdrawalAsSharesState(
        User staker,
        IDelegationManager.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint256[] memory shares
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as shares
        // 
        // ... check that the withdrawal is not pending, that the withdrawer received the expected shares, and that the total shares of each 
        //     strategy withdrawn remains unchanged 
        assert_withdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker's withdrawal should no longer be pending");
        assert_Snap_AddedStakerShares(staker, strategies, shares, "staker should have received expected shares");
        assert_Snap_UnchangedStrategyShares(strategies, "strategies should have total shares unchanged");
    }
}