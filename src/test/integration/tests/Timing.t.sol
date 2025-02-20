// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Integration_ALMSlashBase, IStrategy, IERC20} from "src/test/integration/tests/SlashingWithdrawals.t.sol";

/**
 * @notice These tests check for specific timing edge case behavior correctness
 * @dev These tests assume the following:
 * - The staker has a positive balance in all given trategies
 * - The staker has no pending withdrawals
 * - The staker has no pending slash requests
 * - The staker is delegated to the operator
 */
contract Integration_Timing is Integration_ALMSlashBase {

    /**
     * @notice Test that a slash works correctly just before a partial withdrawal is completed
     */
    function testFuzz_queuePartialWithdrawal_slashBeforeWithdrawal_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 0. Calculate partial withdrawal amounts
        for (uint256 i = 0; i < initTokenBalances.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial
            // but 10 as the maximum divisor is more arbitrary
            initTokenBalances[i] /= _randUint(2, 10);
        }

        // 1. Queue withdrawal
        uint256[] memory sharesToWithdraw = _calculateExpectedShares(strategies, initTokenBalances);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, sharesToWithdraw);

        // 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        // Note: This is done to ensure that the withdrawal is still pending
        vm.roll(block.number - 1);

        // 3. Slash operator
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);
        SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);

        // Assert that the operator allocations and staker withdrawable shares are slashed, but the staker deposit shares are not
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
        assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");

        // 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // 5. Complete withdrawals
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens = _calculateExpectedTokens(
                withdrawals[i].strategies,
                withdrawals[i].scaledShares
            );
            staker.completeWithdrawalAsTokens(withdrawals[i]);

            // Assert that the withdrawal is completed as expected
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }

        // 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");

        // Assert that the staker has withdrawn the expected tokens
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should have withdrawn tokens minus slashed"
        );

        // Assert that all withdrawals are removed from pending
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * @notice Test that a slash works correctly just before a total withdrawal is completed
     */
    function testFuzz_queueTotalWithdrawal_slashBeforeWithdrawal_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 1. Undelegate (i.e. queue a total withdrawal from every strategy)
        Withdrawal[] memory withdrawals = staker.undelegate();

        // 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        // Note: This is done to ensure that the withdrawal is still pending
        vm.roll(block.number - 1);

        // 3. Slash operator
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);
        SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);

        // Assert that the operator allocations and staker withdrawable shares are slashed, but the staker deposit shares are not
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
        assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");

        // 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // 5. Complete withdrawals
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens = _calculateExpectedTokens(
                withdrawals[i].strategies,
                withdrawals[i].scaledShares
            );
            staker.completeWithdrawalAsTokens(withdrawals[i]);

            // Assert that the withdrawal is completed as expected
            check_Withdrawal_AsTokens_State_AfterSlash(
                staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens
            );
        }

        // 6. Check final state

        // Assert that no strategies have any shares remaining
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");

        // Assert that the staker has withdrawn the expected tokens
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            initTokenBalances,
            "staker should have withdrawn tokens minus slashed"
        );

        // Assert that all withdrawals are removed from pending
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * @notice Test that a staker can still complete a partial withdrawal even after a slash has been performed
     */
    function testFuzz_queuePartialWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 0. Calculate partial withdrawal amounts
        for (uint256 i = 0; i < initTokenBalances.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial
            // but 10 as the maximum divisor is more arbitrary
            initTokenBalances[i] /= _randUint(2, 10);
        }

        // 1. Queue withdrawal
        uint256[] memory sharesToWithdraw = _calculateExpectedShares(strategies, initTokenBalances);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, sharesToWithdraw);

        // 2. Move time forward to _exactly_ the withdrawal block
        // Expected behavior: Withdrawals are still pending, but can be completed, and slashes can no longer be 
        // performed _on the partial withdrawals_. The remaining shares may still be affected.
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 3. Slash operator
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);
        // Note: This will _partially_ impact the staker's withdrawable shares as they have not withdrawn all shares
        SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);

        // Assert:
        // * Operator allocations are slashed
        // * Staker withdrawable shares are _partially_ affected due to the partial withdrawal
        // * Staker deposit shares are unaffected
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
        assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should be slashed");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");

        // 4. Complete withdrawal
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens = _calculateExpectedTokens(
                withdrawals[i].strategies,
                withdrawals[i].scaledShares
            );
            staker.completeWithdrawalAsTokens(withdrawals[i]);

            IERC20[] memory tokens = new IERC20[](withdrawals[i].strategies.length);
            for (uint256 j = 0; j < withdrawals[i].strategies.length; ++j) {
                tokens[j] = withdrawals[i].strategies[j].underlyingToken();
            }

            // Assert that the withdrawal is completed as expected, with NO IMPACT from the slash 
            // as these withdrawals are past the delay
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], strategies, sharesToWithdraw, tokens, expectedTokens
            );
        }

        // 5. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");

        // Assert that the staker has withdrawn the expected tokens
        assert_HasUnderlyingTokenBalances(
            staker,
            strategies,
            initTokenBalances,
            "staker should have withdrawn tokens"
        );

        // Assert that all withdrawals are removed from pending
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
    
    
    /**
     * @notice Test that a staker is unaffected by a slash after the withdrawal delay has passed
     */
    function testFuzz_queueTotalWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 1. Undelegate (i.e. queue a total withdrawal from every strategy)
        uint256[] memory sharesToWithdraw = _getWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();

        // 2. Move time forward to _exactly_ the withdrawal block
        // Expected behavior: Withdrawals are still pending, but can be completed, and slashes can no longer affect 
        // this staker
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 3. Slash operator
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);
        // Note: This should have no effect on the STAKER as the withdrawal is already past the withdrawal block
        SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
        
        // Assert:
        // * Operator allocations are slashed
        // * Staker withdrawable shares are unaffected by the slash
        // * Staker deposit shares are unaffected
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
        assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker withdrawable shares should NOT be slashed"); // TODO: verify if this should be passing
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");

        // 4. Complete withdrawal
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens = _calculateExpectedTokens(
                withdrawals[i].strategies,
                withdrawals[i].scaledShares
            );
            staker.completeWithdrawalAsTokens(withdrawals[i]);

            IERC20[] memory tokens = new IERC20[](withdrawals[i].strategies.length);
            for (uint256 j = 0; j < withdrawals[i].strategies.length; ++j) {
                tokens[j] = withdrawals[i].strategies[j].underlyingToken();
            }

            // Assert that the withdrawal is completed as expected, with NO IMPACT from the slash
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], strategies, sharesToWithdraw, tokens, expectedTokens
            );
        }

        // 5. Check final state

        // Assert that all strategies have some shares remaining
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");

        // Assert that the staker has withdrawn the expected tokens
        assert_HasUnderlyingTokenBalances(
            staker,
            strategies,
            initTokenBalances,
            "staker should have withdrawn tokens"
        );

        // Assert that all withdrawals are removed from pending
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}