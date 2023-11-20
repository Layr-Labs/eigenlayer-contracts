// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";

contract Integration_Deposit_Delegate_Undelegate_Complete is IntegrationBase {

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete their queued withdrawal as tokens
    function testFuzz_deposit_undelegate_completeAsTokens(uint24 _random) public {   
        // When new Users are created, they will choose a random configuration from these params: 
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST,
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
            assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal memory expectedWithdrawal = _getExpectedWithdrawalStruct(staker);
        {
            /// 3. Undelegate from an operator
            //
            // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqeuued,
            //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
            //     and operator have reduced shares

            bytes32 withdrawalRoot = staker.undelegate();

            assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");
            assert_ValidWithdrawalHash(expectedWithdrawal, withdrawalRoot, "calculated withdrawl should match returned root");
            assert_WithdrawalPending(withdrawalRoot, "staker's withdrawal should now be pending");
            assert_Snap_IncrementQueuedWithdrawals(staker, "staker should have increased nonce by 1");
            assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete the withdrawal as tokens
            // 
            // ... check that the withdrawal is not pending, that the staker received the expected tokens, and that the total shares of each 
            //     strategy withdrawn from decreased
            uint[] memory expectedTokens = _calculateExpectedTokens(expectedWithdrawal.strategies, expectedWithdrawal.shares);
            IERC20[] memory tokens = staker.completeQueuedWithdrawal(expectedWithdrawal, true);

            assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(expectedWithdrawal), "staker's withdrawal should no longer be pending");
            assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
            assert_Snap_Removed_StrategyShares(strategies, shares, "strategies should have total shares decremented");
        }
    }

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete their queued withdrawal as shares
    function testFuzz_deposit_undelegate_completeAsShares(uint24 _random) public {  
        // When new Users are created, they will choose a random configuration from these params: 
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST,
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
            assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
            assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
        }

        IDelegationManager.Withdrawal memory expectedWithdrawal = _getExpectedWithdrawalStruct(staker);
        {
            /// 3. Undelegate from an operator
            //
            // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqeuued,
            //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
            //     and operator have reduced shares

            bytes32 withdrawalRoot = staker.undelegate();

            assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");
            assert_ValidWithdrawalHash(expectedWithdrawal, withdrawalRoot, "calculated withdrawl should match returned root");
            assert_WithdrawalPending(withdrawalRoot, "staker's withdrawal should now be pending");
            assert_Snap_IncrementQueuedWithdrawals(staker, "staker should have increased nonce by 1");
            assert_Snap_Removed_OperatorShares(operator, strategies, shares, "failed to remove operator shares");
            assert_Snap_Removed_StakerShares(staker, strategies, shares, "failed to remove staker shares");
        }

        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        {
            /// 4. Complete withdrawal(s):
            // The staker will complete the withdrawal as tokens
            // 
            // ... check that the withdrawal is not pending, that the withdrawer received the expected shares, and that the total shares of each 
            //     strategy withdrawn remains unchanged 
            staker.completeQueuedWithdrawal(expectedWithdrawal, false);

            assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(expectedWithdrawal), "staker's withdrawal should no longer be pending");
            assert_Snap_Added_StakerShares(staker, expectedWithdrawal.strategies, expectedWithdrawal.shares, "staker should have received expected tokens");
            assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");
        } 
    }

    /// @notice Assumes staker and withdrawer are the same
    function _getExpectedWithdrawalStruct(User staker) internal view returns (IDelegationManager.Withdrawal memory) {
        (IStrategy[] memory strategies, uint[] memory shares)
            = delegationManager.getDelegatableShares(address(staker));

        return IDelegationManager.Withdrawal({
            staker: address(staker),
            delegatedTo: delegationManager.delegatedTo(address(staker)),
            withdrawer: address(staker),
            nonce: delegationManager.cumulativeWithdrawalsQueued(address(staker)),
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: shares
        });
    }

    function _getExpectedWithdrawalStruct_diffWithdrawer(User staker, address withdrawer) internal view returns (IDelegationManager.Withdrawal memory) {
        IDelegationManager.Withdrawal memory withdrawal = _getExpectedWithdrawalStruct(staker);
        withdrawal.withdrawer = withdrawer;
        return withdrawal;
    }
}