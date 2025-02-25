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
contract Integration_WithdrawalTiming is Integration_ALMSlashBase {

    ///////////////////////////////
    /// WITHDRAWAL TIMING TESTS ///
    ///////////////////////////////

    /**
     * @notice Test that a slash works correctly just before a partial withdrawal is completed
     */
    function testFuzz_queuePartialWithdrawal_slashBeforeWithdrawal_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        uint256[] memory depositSharesToWithdraw = new uint256[](initDepositShares.length);
        // 0. Calculate partial withdrawal amounts
        for (uint256 i = 0; i < initDepositShares.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial
            // but 10 as the maximum divisor is more arbitrary
            depositSharesToWithdraw[i] = initDepositShares[i] / _randUint(2, 10);
        }

        // 1. Queue withdrawal
        uint256[] memory withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositSharesToWithdraw);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, depositSharesToWithdraw, withdrawableShares, withdrawals, withdrawalRoots);

        // 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        // Note: This is done to ensure that the withdrawal is still pending
        vm.roll(block.number - 1);

        // 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        // 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // 5. Complete withdrawals
        withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        // 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");

        // Assert that all withdrawals are removed from pending
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * @notice Test that a slash works correctly just before a total withdrawal is completed
     */
    function testFuzz_queueTotalWithdrawal_slashBeforeWithdrawal_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 1. Queue withdrawal
        uint256[] memory withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        // Note: This is done to ensure that the withdrawal is still pending
        vm.roll(block.number - 1);

        // 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        // 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // 5. Complete withdrawals
        withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        // 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, 0, "all strategies should have no shares remaining");

        // Assert that all withdrawals are removed from pending
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * @notice Test that a staker can still complete a partial withdrawal even after a slash has been performed
     */
    function testFuzz_queuePartialWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        uint256[] memory depositSharesToWithdraw = new uint256[](initDepositShares.length);
        // 0. Calculate partial withdrawal amounts
        for (uint256 i = 0; i < initDepositShares.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial
            // but 10 as the maximum divisor is more arbitrary
            depositSharesToWithdraw[i] = initDepositShares[i] / _randUint(2, 10);
        }

        // 1. Queue withdrawal
        uint256[] memory withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositSharesToWithdraw);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Note: Because this is a partial withdrawal, we pass in the initial token balances as the withdrawable shares
        // rather than `initDepositShares`
        // Moreover, the shares to withdraw are the expected shares calculated from the token balances
        check_QueuedWithdrawal_State(staker, operator, strategies, depositSharesToWithdraw, withdrawableShares, withdrawals, withdrawalRoots);

        // 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        // 5. Complete withdrawals
        // Note: expectedTokens must be recalculated because the withdrawable shares
        // have changed due to the slash
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        // 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");

        // Assert that all withdrawals are removed from pending
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * @notice Test that a staker is unaffected by a slash after the withdrawal delay has passed
     */
    function testFuzz_queueTotalWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(
        uint24 _random
    ) public rand(_random) {
        // 1. Queue withdrawal
        uint256[] memory withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // 2. Move time forward to withdrawal block
        // Expected behavior: Withdrawals are no longer pending, and slashes
        // no longer affect the staker
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        // 5. Complete withdrawals
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        // 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, 0, "all strategies should have no shares remaining");

        // Assert that all withdrawals are removed from pending
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}

contract Integration_OperatorDeallocationTiming is Integration_ALMSlashBase {

    //////////////////////////////////////////
    /// OPERATOR DEALLOCATION TIMING TESTS ///
    //////////////////////////////////////////

    function testFuzz_deallocateFully_slashBeforeDelay(uint24 _r) public rand(_r) {
        // 1. Deallocate
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 2. Move time forward to _just before_ deallocation delay
        // Expected behavior: Deallocation delay is not yet passed, so slashes can still be performed
        _rollForward_DeallocationDelay();
        rollBackward(1); // make sure that deallocation delay is not yet passed

        // 3. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, deallocateParams, slashParams);
    }

    function testFuzz_deallocateFully_slashAfterDelay(uint24 _r) public rand(_r) {
        // 1. Deallocate
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 2. Move time forward to deallocation delay
        _rollForward_DeallocationDelay();

        // 3. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, deallocateParams, slashParams);
    }
}
