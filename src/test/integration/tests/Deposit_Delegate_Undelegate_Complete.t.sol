// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Delegate_Undelegate_Complete is IntegrationChecks {
    uint[] shares;
    uint[] delegatableShares;

    function _init() internal virtual override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        shares = _calculateExpectedShares(strategies, initTokenBalances);
        //delegatable shares equals deposit shares here because no bc slashing
        delegatableShares = shares;

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);
    }

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete their queued withdrawal as tokens
    function testFuzz_deposit_undelegate_completeAsTokens(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawal
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
        }
    }

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete their queued withdrawal as shares
    function testFuzz_deposit_undelegate_completeAsShares(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);

            check_Withdrawal_AsShares_Undelegated_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
            );
        }
    }

    function testFuzz_deposit_delegate_forceUndelegate_completeAsTokens(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        // 3. Force undelegate
        Withdrawal[] memory withdrawals = operator.forceUndelegate(staker);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
        }
    }

    function testFuzz_deposit_delegate_forceUndelegate_completeAsShares(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        // 3. Force undelegate
        Withdrawal[] memory withdrawals = operator.forceUndelegate(staker);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
            );
        }
    }

    function test_deposit_delegate_undelegate_completeAsTokens_Max_Strategies() public {
        _configAssetTypes(HOLDS_MAX);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator();

        if (eq(FOUNDRY_PROFILE(), "default"))  {
            assertEq(strategies.length, 9, "sanity");

            uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
            //delegatable shares equals deposit shares here because no bc slashing
            uint[] memory delegatableShares = shares;

            /// 1. Deposit Into Strategies
            staker.depositIntoEigenlayer(strategies, initTokenBalances);
            check_Deposit_State(staker, strategies, shares);

            // 2. Delegate to an operator
            staker.delegateTo(operator);
            check_Delegation_State(staker, operator, strategies, shares);

            // 3. Undelegate from an operator
            Withdrawal[] memory withdrawals = staker.undelegate();
            withdrawalRoots = _getWithdrawalHashes(withdrawals);
            check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, delegatableShares);

            // 4. Complete withdrawal
            // Fast forward to when we can complete the withdrawal
            _rollBlocksForCompleteWithdrawals(withdrawals);

            // Complete withdrawal
            for (uint i = 0; i < withdrawals.length; ++i) {
                uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
                staker.completeWithdrawalAsTokens(withdrawals[i]);
                check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
            }
        }
    }
}
