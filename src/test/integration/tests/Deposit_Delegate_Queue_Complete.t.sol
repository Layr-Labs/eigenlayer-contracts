// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";

contract Deposit_Delegate_Queue_Complete is IntegrationBase {

    /*******************************************************************************
                                FULL WITHDRAWALS
    *******************************************************************************/

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a ALL shares
    /// 4. completes the queued withdrawal as tokens
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
            assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_Snap_Unchanged_StakerShares(staker, "staker shares should be unchanged after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for all strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned withdrawals
            //     match now-pending withdrawal roots, and that the staker and operator have
            //     reduced shares.
            withdrawals = staker.queueWithdrawals(strategies, shares);
            withdrawalRoots = _getWithdrawalHashes(withdrawals);

            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker received their tokens
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                uint[] memory expectedTokens = _calculateExpectedTokens(withdrawal.strategies, withdrawal.shares);
                IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawal, true);

                assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
                assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
                assert_Snap_Unchanged_StakerShares(staker, "staker shares should not have changed");
                assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
            }
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "staker should once again have original token balances");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a ALL shares
    /// 4. completes the queued withdrawal as shares
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
            assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_Snap_Unchanged_StakerShares(staker, "staker shares should be unchanged after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for all strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned withdrawals
            //     match now-pending withdrawal roots, and that the staker and operator have
            //     reduced shares.
            withdrawals = staker.queueWithdrawals(strategies, shares);
            withdrawalRoots = _getWithdrawalHashes(withdrawals);

            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker and operator received their shares and that neither
            // have any change in token balances
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                staker.completeQueuedWithdrawal(withdrawal, false);

                assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
                assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
                assert_Snap_Added_StakerShares(staker, withdrawal.strategies, withdrawal.shares, "staker should have received shares");
                assert_Snap_Added_OperatorShares(operator, withdrawal.strategies, withdrawal.shares, "operator should have received shares");
            }
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /*******************************************************************************
                              RANDOM WITHDRAWALS
    *******************************************************************************/

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a random subset of shares
    /// 4. completes the queued withdrawal as tokens
    function testFuzz_deposit_delegate_queueRand_completeAsTokens(uint24 _random) public {
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
            assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_Snap_Unchanged_StakerShares(staker, "staker shares should be unchanged after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        // Randomly select one or more assets to withdraw
        (
            IStrategy[] memory withdrawStrats,
            uint[] memory withdrawShares
        ) = _randWithdrawal(strategies, shares);

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for the selected strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned roots
            //     match the hashes of each withdrawal, and that the staker and operator have
            //     reduced shares.
            withdrawals = staker.queueWithdrawals(withdrawStrats, withdrawShares);
            withdrawalRoots = _getWithdrawalHashes(withdrawals);

            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_Removed_OperatorShares(operator, withdrawStrats, withdrawShares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, withdrawStrats, withdrawShares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker received their tokens and that the staker/operator
            // have unchanged share amounts
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                uint[] memory expectedTokens = _calculateExpectedTokens(withdrawal.strategies, withdrawal.shares);
                IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawal, true);

                assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
                assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
                assert_Snap_Unchanged_StakerShares(staker, "staker shares should not have changed");
                assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
            }
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// 2. delegates to an operator
    /// 3. queues a withdrawal for a random subset of shares
    /// 4. completes the queued withdrawal as shares
    function testFuzz_deposit_delegate_queueRand_completeAsShares(uint24 _random) public {
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
            assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Delegate to an operator:
            //
            // ... check that the staker is now delegated to the operator, and that the operator
            //     was awarded the staker's shares
            staker.delegateTo(operator);

            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
            assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
            assert_Snap_Unchanged_StakerShares(staker, "staker shares should be unchanged after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        // Randomly select one or more assets to withdraw
        (
            IStrategy[] memory withdrawStrats,
            uint[] memory withdrawShares
        ) = _randWithdrawal(strategies, shares);

        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;

        {
            /// 3. Queue withdrawal(s):
            // The staker will queue one or more withdrawals for the selected strategies and shares
            //
            // ... check that each withdrawal was successfully enqueued, that the returned roots
            //     match the hashes of each withdrawal, and that the staker and operator have
            //     reduced shares.
            withdrawals = staker.queueWithdrawals(withdrawStrats, withdrawShares);
            withdrawalRoots = _getWithdrawalHashes(withdrawals);

            assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots, "calculated withdrawals should match returned roots");
            assert_AllWithdrawalsPending(withdrawalRoots, "staker's withdrawals should now be pending");
            assert_Snap_Added_QueuedWithdrawals(staker, withdrawals, "staker should have increased nonce by withdrawals.length");
            assert_Snap_Removed_OperatorShares(operator, withdrawStrats, withdrawShares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, withdrawStrats, withdrawShares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete each withdrawal as tokens
            // 
            // ... check that the staker received their tokens and that the staker/operator
            // have unchanged share amounts
            for (uint i = 0; i < withdrawals.length; i++) {
                IDelegationManager.Withdrawal memory withdrawal = withdrawals[i];

                staker.completeQueuedWithdrawal(withdrawal, false);

                assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
                assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
                assert_Snap_Added_StakerShares(staker, withdrawal.strategies, withdrawal.shares, "staker should have received shares");
                assert_Snap_Added_OperatorShares(operator, withdrawal.strategies, withdrawal.shares, "operator should have received shares");
            }
        }

        // Check final state:
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /*******************************************************************************
                               UNHAPPY PATH TESTS
    *******************************************************************************/

    /// Generates a random staker and operator. The staker:
    /// 1. deposits all assets into strategies
    /// --- registers as an operator
    /// 2. delegates to an operator
    ///
    /// ... we check that the final step fails
    function testFuzz_deposit_delegate_revert_alreadyDelegated(uint24 _random) public {
        _configRand({
            _randomSeed: _random,
            _assetTypes: NO_ASSETS | HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        /// 0. Create a staker and operator
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
            assert_Snap_Added_StakerShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        }

        {
            /// 2. Register the staker as an operator, then attempt to delegate to an operator. 
            ///    This should fail as the staker is already delegated to themselves.
            staker.registerAsOperator();
            assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");

            cheats.expectRevert("DelegationManager._delegate: staker is already actively delegated");
            staker.delegateTo(operator);
        }
    }
}