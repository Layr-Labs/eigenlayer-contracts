// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";

contract Deposit_Delegate_Queue_Complete is IntegrationBase {

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. queue a withdrawal for all shares (withdrawer set to staker)
    /// 4. complete their queued withdrawal as tokens
    function testFuzz_deposit_delegate_queue_completeAsTokens(uint24 _random) public {   
        // When new Users are created, they will choose a random configuration from these params: 
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (
            User staker,
            IStrategy[] memory strategies, 
            uint[] memory tokenBalances
        ) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        {
            /// 1. Deposit into strategies:
            // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
            // the staker calls the relevant deposit function, depositing all held assets.
            //
            // ... check that all underlying tokens were transferred to the correct destination
            //     and that the staker now has the expected amount of delegated shares in each strategy
            staker.depositIntoEigenlayer(strategies, tokenBalances);

            assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker should have transferred all underlying tokens");
            assert_Snap_AddedStakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
            assert_Snap_AddedOperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for the selected strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned roots
            //     match the hashes of each withdrawal, and that the staker and operator have
            //     reduced shares.
            (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);

            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_Snap_IncreasedQueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_RemovedOperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_RemovedStakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker received their tokens
            // TODO - there are more checks to be made here but i want to wrap this up and get eyes on it
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                uint[] memory expectedTokens = _calculateExpectedTokens(withdrawal.strategies, withdrawal.shares);
                IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawal, true);

                assert_Snap_IncreasedTokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
            }
        }
    }

    function testFuzz_deposit_delegate_queue_completeAsShares(uint24 _random) public {   
        // When new Users are created, they will choose a random configuration from these params: 
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random subset of valid strategies (StrategyManager and/or EigenPodManager)
        //
        // ... check that the staker has no delegatable shares and isn't currently delegated
        (
            User staker,
            IStrategy[] memory strategies, 
            uint[] memory tokenBalances
        ) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        {
            /// 1. Deposit into strategies:
            // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
            // the staker calls the relevant deposit function, depositing all held assets.
            //
            // ... check that all underlying tokens were transferred to the correct destination
            //     and that the staker now has the expected amount of delegated shares in each strategy
            staker.depositIntoEigenlayer(strategies, tokenBalances);

            assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker should have transferred all underlying tokens");
            assert_Snap_AddedStakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
            assert_Snap_AddedOperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for the selected strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned roots
            //     match the hashes of each withdrawal, and that the staker and operator have
            //     reduced shares.
            (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);

            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_Snap_IncreasedQueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_RemovedOperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_RemovedStakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker received their tokens
            // TODO - there are more checks to be made here but i want to wrap this up and get eyes on it
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                // uint[] memory expectedTokens = _calculateExpectedTokens(withdrawal.strategies, withdrawal.shares);
                staker.completeQueuedWithdrawal(withdrawal, false);

                assert_Snap_AddedStakerShares(staker, withdrawal.strategies, withdrawal.shares, "staker should have received expected tokens");
            }
        }
    }
}