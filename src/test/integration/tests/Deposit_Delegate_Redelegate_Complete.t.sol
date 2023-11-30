// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/tests/utils.t.sol";

contract Integration_Deposit_Delegate_Redelegate_Complete is IntegrationTestUtils {
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
        assertDepositState(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        assertDelegationState(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal memory expectedWithdrawal = _getExpectedWithdrawalStruct(staker);
        bytes32 withdrawalRoot = staker.undelegate();
        assertUndelegateState(staker, operator1, expectedWithdrawal, withdrawalRoot, strategies, shares);

        // 4. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        staker.completeQueuedWithdrawal(expectedWithdrawal, false);

        assertWithdrawalAsSharesState(staker, expectedWithdrawal, strategies, shares);

        // 5. Delegate to a new operator
        staker.delegateTo(operator2);
        assertDelegationState(staker, operator2, strategies, shares);
        assertNotEq(address(operator1), delegationManager.delegatedTo(address(staker)), "staker should not be delegated to operator1");

        // 6. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);
        assertQueuedWithdrawalState(staker, operator2, strategies, shares, withdrawals, withdrawalRoots);

        // 7. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        // Complete withdrawals
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawals[i], true);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].shares);
            assertWithdrawalAsTokensState(staker, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    /// @notice Assumes staker and withdrawer are the same
    function _getExpectedWithdrawalStruct(User staker) internal view returns (IDelegationManager.Withdrawal memory) {
        (IStrategy[] memory strategies, uint256[] memory shares)
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

    //TODO: add complete last withdrawal as shares
    //TODO: add complete middle withdrawal as tokens and then restake and redelegate
    //TODO: additional deposit before delegating to new operator
    //TODO: additional deposit after delegating to new operator

    function _getExpectedWithdrawalStruct_diffWithdrawer(User staker, address withdrawer) internal view returns (IDelegationManager.Withdrawal memory) {
        IDelegationManager.Withdrawal memory withdrawal = _getExpectedWithdrawalStruct(staker);
        withdrawal.withdrawer = withdrawer;
        return withdrawal;
    }
}