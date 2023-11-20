// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/tests/utils.t.sol";

contract Integration_Deposit_Delegate_Undelegate_Complete is IntegrationTestUtils {

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

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        assertDepositState(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        assertDelegationState(staker, operator, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal memory expectedWithdrawal = _getExpectedWithdrawalStruct(staker);
        bytes32 withdrawalRoot = staker.undelegate();
        assertUndelegateState(staker, operator, expectedWithdrawal, withdrawalRoot, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        // Complete withdrawal
        uint[] memory expectedTokens = _calculateExpectedTokens(expectedWithdrawal.strategies, expectedWithdrawal.shares);
        IERC20[] memory tokens = staker.completeQueuedWithdrawal(expectedWithdrawal, true);

        assertWithdrawalAsTokensState(staker, expectedWithdrawal, strategies, shares, tokens, expectedTokens);
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

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        assertDepositState(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        assertDelegationState(staker, operator, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal memory expectedWithdrawal = _getExpectedWithdrawalStruct(staker);
        bytes32 withdrawalRoot = staker.undelegate();
        assertUndelegateState(staker, operator, expectedWithdrawal, withdrawalRoot, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());

        uint[] memory expectedTokens = _calculateExpectedTokens(expectedWithdrawal.strategies, expectedWithdrawal.shares);
        IERC20[] memory tokens = staker.completeQueuedWithdrawal(expectedWithdrawal, false);

        assertWithdrawalAsSharesState(staker, expectedWithdrawal, strategies, shares);
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