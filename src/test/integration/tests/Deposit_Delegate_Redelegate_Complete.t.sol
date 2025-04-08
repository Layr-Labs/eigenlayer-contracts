// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Delegate_Redelegate_Complete is IntegrationChecks {
    using ArrayLib for *;

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. delegate to an operator
    /// 3. undelegates from the operator
    /// 4. complete queued withdrawal as shares
    /// 5. delegate to a new operator
    /// 5. queueWithdrawal
    /// 7. complete their queued withdrawal as tokens
    function testFuzz_deposit_delegate_reDelegate_completeAsTokens(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        //delegatable shares equals deposit shares here because no bc slashing
        uint[] memory delegatableShares = shares;

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        check_Delegation_State(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(
                staker, operator1, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
            );
        }

        // 5. Delegate to a new operator
        staker.delegateTo(operator2);
        check_Delegation_State(staker, operator2, strategies, shares);

        // 6. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator2, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 7. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawals
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
        }
    }

    function testFuzz_deposit_delegate_reDelegate_completeAsShares(uint24) public {
        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        //delegatable shares equals deposit shares here because no bc slashing
        uint[] memory delegatableShares = shares;

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        check_Delegation_State(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(
                staker, operator1, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
            );
        }

        // 5. Delegate to a new operator
        staker.delegateTo(operator2);
        check_Delegation_State(staker, operator2, strategies, shares);
        assertNotEq(address(operator1), delegationManager().delegatedTo(address(staker)), "staker should not be delegated to operator1");

        // 6. Queue Withdrawal
        withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator2, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);

        // 7. Complete withdrawal
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete all but last withdrawal as tokens
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, staker, withdrawals[i], shares, expectedTokens);
        }
    }

    function testFuzz_deposit_delegate_reDelegate_depositAfterRedelegate(uint24) public {
        _configAssetTypes(HOLDS_LST); // not holding ETH since we can only deposit 32 ETH multiples

        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        {
            // Divide shares by 2 in new array to do deposits after redelegate
            uint[] memory numTokensToDeposit = new uint[](initTokenBalances.length);
            uint[] memory numTokensRemaining = new uint[](initTokenBalances.length);
            for (uint i = 0; i < shares.length; i++) {
                numTokensToDeposit[i] = initTokenBalances[i] / 2;
                numTokensRemaining[i] = initTokenBalances[i] - numTokensToDeposit[i];
            }
            uint[] memory halfShares = _calculateExpectedShares(strategies, numTokensToDeposit);
            //delegatable shares equals deposit shares here because no bc slashing
            uint[] memory delegatableShares = halfShares;

            /// 1. Deposit Into Strategies
            staker.depositIntoEigenlayer(strategies, numTokensToDeposit);
            check_Deposit_State_PartialDeposit(staker, strategies, halfShares, numTokensRemaining);

            // 2. Delegate to an operator
            staker.delegateTo(operator1);
            check_Delegation_State(staker, operator1, strategies, halfShares);

            // 3. Undelegate from an operator
            Withdrawal[] memory withdrawals = staker.undelegate();
            withdrawalRoots = _getWithdrawalHashes(withdrawals);
            check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

            // 4. Complete withdrawal as shares
            // Fast forward to when we can complete the withdrawal
            _rollBlocksForCompleteWithdrawals(withdrawals);
            for (uint i = 0; i < withdrawals.length; ++i) {
                staker.completeWithdrawalAsShares(withdrawals[i]);
                check_Withdrawal_AsShares_Undelegated_State(
                    staker, operator1, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
                );
            }

            // 5. Delegate to a new operator
            staker.delegateTo(operator2);
            check_Delegation_State(staker, operator2, strategies, halfShares);
            assertNotEq(address(operator1), delegationManager().delegatedTo(address(staker)), "staker should not be delegated to operator1");

            // 6. Deposit into Strategies
            uint[] memory sharesAdded = _calculateExpectedShares(strategies, numTokensRemaining);
            staker.depositIntoEigenlayer(strategies, numTokensRemaining);
            initTokenBalances = _calculateExpectedTokens(strategies, shares);
            check_Deposit_State(staker, strategies, sharesAdded);
        }

        {
            // 7. Queue Withdrawal
            shares = _calculateExpectedShares(strategies, initTokenBalances);
            Withdrawal[] memory newWithdrawals = staker.queueWithdrawals(strategies, shares);
            bytes32[] memory newWithdrawalRoots = _getWithdrawalHashes(newWithdrawals);
            check_QueuedWithdrawal_State(staker, operator2, strategies, shares, shares, newWithdrawals, newWithdrawalRoots);

            // 8. Complete withdrawal
            // Fast forward to when we can complete the withdrawal
            _rollBlocksForCompleteWithdrawals(newWithdrawals);

            // Complete withdrawals
            for (uint i = 0; i < newWithdrawals.length; i++) {
                uint[] memory expectedTokens = _calculateExpectedTokens(newWithdrawals[i].strategies, newWithdrawals[i].scaledShares);
                staker.completeWithdrawalAsTokens(newWithdrawals[i]);
                check_Withdrawal_AsTokens_State(staker, operator2, newWithdrawals[i], shares, expectedTokens);
            }
        }
    }

    function testFuzz_deposit_delegate_reDelegate_depositBeforeRedelegate(uint24) public {
        _configAssetTypes(HOLDS_LST); // not holding ETH since we can only deposit 32 ETH multiples

        /// 0. Create an operator and a staker with:
        // - some nonzero underlying token balances
        // - corresponding to a random number of strategies
        //
        // ... check that the staker has no deleagatable shares and isn't delegated

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory totalShares = new uint[](strategies.length);
        {
            // Divide shares by 2 in new array to do deposits after redelegate
            uint[] memory numTokensToDeposit = new uint[](initTokenBalances.length);
            uint[] memory numTokensRemaining = new uint[](initTokenBalances.length);
            for (uint i = 0; i < strategies.length; i++) {
                numTokensToDeposit[i] = initTokenBalances[i] / 2;
                numTokensRemaining[i] = initTokenBalances[i] - numTokensToDeposit[i];
            }
            {
                uint[] memory sharesFromFirstDeposit = _calculateExpectedShares(strategies, numTokensToDeposit);
                //delegatable shares equals deposit shares here because no bc slashing
                uint[] memory delegatableShares = sharesFromFirstDeposit;

                /// 1. Deposit Into Strategies
                staker.depositIntoEigenlayer(strategies, numTokensToDeposit);
                check_Deposit_State_PartialDeposit(staker, strategies, sharesFromFirstDeposit, numTokensRemaining);

                // 2. Delegate to an operator
                staker.delegateTo(operator1);
                check_Delegation_State(staker, operator1, strategies, sharesFromFirstDeposit);

                // 3. Undelegate from an operator
                Withdrawal[] memory withdrawals = staker.undelegate();
                withdrawalRoots = _getWithdrawalHashes(withdrawals);
                check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

                // 4. Complete withdrawal as shares
                // Fast forward to when we can complete the withdrawal
                _rollBlocksForCompleteWithdrawals(withdrawals);
                for (uint i = 0; i < withdrawals.length; ++i) {
                    staker.completeWithdrawalAsShares(withdrawals[i]);
                    check_Withdrawal_AsShares_Undelegated_State(
                        staker, operator1, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares
                    );
                }

                // 5. Deposit into Strategies
                uint[] memory sharesFromSecondDeposit = _calculateExpectedShares(strategies, numTokensRemaining);
                for (uint i = 0; i < strategies.length; i++) {
                    totalShares[i] = sharesFromFirstDeposit[i] + sharesFromSecondDeposit[i];
                }

                staker.depositIntoEigenlayer(strategies, numTokensRemaining);
                initTokenBalances = _calculateExpectedTokens(strategies, totalShares);
                check_Deposit_State(staker, strategies, sharesFromSecondDeposit);
            }

            // 6. Delegate to a new operator
            staker.delegateTo(operator2);
            check_Delegation_State(staker, operator2, strategies, totalShares);
            assertNotEq(address(operator1), delegationManager().delegatedTo(address(staker)), "staker should not be delegated to operator1");
        }

        {
            // 7. Queue Withdrawal
            totalShares = _calculateExpectedShares(strategies, initTokenBalances);
            Withdrawal[] memory newWithdrawals = staker.queueWithdrawals(strategies, totalShares);
            bytes32[] memory newWithdrawalRoots = _getWithdrawalHashes(newWithdrawals);
            check_QueuedWithdrawal_State(staker, operator2, strategies, totalShares, totalShares, newWithdrawals, newWithdrawalRoots);

            // 8. Complete withdrawal
            // Fast forward to when we can complete the withdrawal
            _rollBlocksForCompleteWithdrawals(newWithdrawals);

            // Complete withdrawals
            for (uint i = 0; i < newWithdrawals.length; i++) {
                uint[] memory expectedTokens = _calculateExpectedTokens(newWithdrawals[i].strategies, newWithdrawals[i].scaledShares);
                staker.completeWithdrawalAsTokens(newWithdrawals[i]);
                check_Withdrawal_AsTokens_State(staker, operator2, newWithdrawals[i], totalShares, expectedTokens);
            }
        }
    }

    function testFuzz_deposit_delegate_undelegate_withdrawAsTokens_reDelegate_completeAsTokens(uint24) public {
        /// 0. Create operators and a staker
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        //delegatable shares equals deposit shares here because no bc slashing
        uint[] memory delegatableShares = shares;

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory withdrawnTokenBalances = _calculateExpectedTokens(strategies, shares);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        check_Delegation_State(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal as tokens
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator1, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
        }

        //5. Deposit into Strategies
        staker.depositIntoEigenlayer(strategies, withdrawnTokenBalances);
        shares = _calculateExpectedShares(strategies, withdrawnTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 6. Delegate to a new operator
        staker.delegateTo(operator2);
        check_Delegation_State(staker, operator2, strategies, shares);
        assertNotEq(address(operator1), delegationManager().delegatedTo(address(staker)), "staker should not be delegated to operator1");

        {
            // 7. Queue Withdrawal
            withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
            withdrawals = staker.queueWithdrawals(strategies, shares);
            withdrawalRoots = _getWithdrawalHashes(withdrawals);
            check_QueuedWithdrawal_State(staker, operator2, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);
        }

        // 8. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawals as tokens
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator2, withdrawals[i], shares, expectedTokens);
        }
    }

    function testFuzz_deposit_delegate_undelegate_withdrawAsTokens_reDelegate_completeAsShares(uint24) public {
        /// 0. Create operators and a staker
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        User operator1 = _newRandomOperator();
        User operator2 = _newRandomOperator();

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        //delegatable shares equals deposit shares here because no bc slashing
        uint[] memory delegatableShares = shares;

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        uint[] memory withdrawnTokenBalances = _calculateExpectedTokens(strategies, shares);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator1);
        check_Delegation_State(staker, operator1, strategies, shares);

        // 3. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator1, withdrawals, withdrawalRoots, strategies, delegatableShares);

        // 4. Complete withdrawal as Tokens
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator1, withdrawals[i], withdrawals[i].scaledShares, expectedTokens);
        }

        // 5. Deposit into Strategies
        shares = _calculateExpectedShares(strategies, withdrawnTokenBalances);
        staker.depositIntoEigenlayer(strategies, withdrawnTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 6. Delegate to a new operator
        staker.delegateTo(operator2);
        check_Delegation_State(staker, operator2, strategies, shares);
        assertNotEq(address(operator1), delegationManager().delegatedTo(address(staker)), "staker should not be delegated to operator1");

        // 7. Queue Withdrawal
        shares = _calculateExpectedShares(strategies, withdrawnTokenBalances);
        withdrawals = staker.queueWithdrawals(strategies, shares);
        withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator2, strategies, shares, shares, withdrawals, withdrawalRoots);

        // 8. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Complete withdrawals as shares
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator2, withdrawals[i], strategies, shares);
        }
    }
}
