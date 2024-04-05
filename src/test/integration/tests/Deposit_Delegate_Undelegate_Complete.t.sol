// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Delegate_Undelegate_Complete is IntegrationCheckUtils {

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete their queued withdrawal as tokens
    function testFuzz_deposit_undelegate_completeAsTokens(uint24 _random) public {   
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
        (User operator, ,) = _newRandomOperator();
        // Upgrade contracts if forkType is not local
        _upgradeEigenLayerContracts();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(strategies);

        // Complete withdrawal
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].shares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares, tokens, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "staker should once again have original token balances");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
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
        (User operator, ,) = _newRandomOperator();
        // Upgrade contracts if forkType is not local
        _upgradeEigenLayerContracts();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Undelegate from an operator
        IDelegationManager.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(strategies);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);

            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares);
        }

        // Check final state:
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deposit_delegate_forceUndelegate_completeAsTokens(uint24 _random) public {
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
        (User operator, ,) = _newRandomOperator();
        // Upgrade contracts if forkType is not local
        _upgradeEigenLayerContracts();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Force undelegate
        IDelegationManager.Withdrawal[] memory withdrawals = operator.forceUndelegate(staker);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(strategies);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].shares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares, tokens, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "staker should once again have original token balances");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deposit_delegate_forceUndelegate_completeAsShares(uint24 _random) public {
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
        (User operator, ,) = _newRandomOperator();
        // Upgrade contracts if forkType is not local
        _upgradeEigenLayerContracts();

        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // 3. Force undelegate
        IDelegationManager.Withdrawal[] memory withdrawals = operator.forceUndelegate(staker);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, shares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(strategies);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].shares);
        }

        // Check final state:
        assert_HasExpectedShares(staker, strategies, shares, "staker should have all original shares");
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}
