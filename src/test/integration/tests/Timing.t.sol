// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/tests/SlashingWithdrawals.t.sol";

/**
 * @notice These tests check for specific withdrawal correctness around timing
 * @dev These tests assume the following:
 * - The staker has a positive balance in all given strategies
 * - The staker has no pending withdrawals
 * - The staker is delegated to the operator
 */
contract Integration_WithdrawalTiming is Integration_ALMSlashBase {
    ///////////////////////////////
    /// WITHDRAWAL TIMING TESTS ///
    ///////////////////////////////

    /**
     * @notice Test that a slash works correctly just before a _partial_ withdrawal is completed
     */
    function testFuzz_queuePartialWithdrawal_slashBeforeWithdrawalDelay_completeAsTokens(uint24 _r) public rand(_r) {
        uint[] memory depositSharesToWithdraw = new uint[](initDepositShares.length);
        /// 0. Calculate partial withdrawal amounts
        for (uint i = 0; i < initDepositShares.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial but 10 as
            // the maximum divisor is more arbitrary
            depositSharesToWithdraw[i] = initDepositShares[i] / _randUint(2, 10);
        }

        /// 1. Queue withdrawal
        uint[] memory withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositSharesToWithdraw);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Validate correctly queued withdrawals
        check_QueuedWithdrawal_State(
            staker, operator, strategies, depositSharesToWithdraw, withdrawableShares, withdrawals, withdrawalRoots
        );

        /// 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        vm.roll(block.number - 1);

        // Verify that the withdrawals are _still_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) <= slashableUntil);
        }

        /// 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);

        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        /// 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // Verify that the withdrawals are _no longer_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) > slashableUntil);
        }

        /// 5. Complete withdrawals
        // Note: expectedTokens must be recalculated because the withdrawable shares have changed due to the slash
        withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);

        // Verify that the withdrawals were completed correctly
        for (uint i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        /// 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");
    }

    /**
     * @notice Test that a slash works correctly just before a _total_ withdrawal is completed
     */
    function testFuzz_queueTotalWithdrawal_slashBeforeWithdrawalDelay_completeAsTokens(uint24 _r) public rand(_r) {
        /// 1. Queue withdrawal
        uint[] memory withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Validate correctly queued withdrawals
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        /// 2. Move time forward to _just before_ withdrawal block
        // Expected behavior: Withdrawals are still pending and cannot be completed, but slashes can still be performed
        _rollBlocksForCompleteWithdrawals(withdrawals);
        vm.roll(block.number - 1);

        // Verify that the withdrawals are _still_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) <= slashableUntil);
        }

        /// 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);

        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        /// 4. Move time forward to withdrawal block
        vm.roll(block.number + 1);

        // Verify that the withdrawals are _no longer_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) > slashableUntil);
        }

        /// 5. Complete withdrawals
        // Note: expectedTokens must be recalculated because the withdrawable shares have changed due to the slash
        withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);

        // Verify that the withdrawals were completed correctly
        for (uint i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        /// 6. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, 0, "all strategies should have no shares remaining");
    }

    /**
     * @notice Test that a staker can still complete a partial withdrawal even after a slash has been performed
     */
    function testFuzz_queuePartialWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(uint24 _r) public rand(_r) {
        uint[] memory depositSharesToWithdraw = new uint[](initDepositShares.length);
        /// 0. Calculate partial withdrawal amounts
        for (uint i = 0; i < initDepositShares.length; ++i) {
            // Note: 2 is specifically chosen as the minimum divisor to ensure that the withdrawal is partial
            // but 10 as the maximum divisor is more arbitrary
            depositSharesToWithdraw[i] = initDepositShares[i] / _randUint(2, 10);
        }

        /// 1. Queue withdrawal
        uint[] memory withdrawableShares = _calcWithdrawable(staker, strategies, depositSharesToWithdraw);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositSharesToWithdraw);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Validate correctly queued withdrawals
        check_QueuedWithdrawal_State(
            staker, operator, strategies, depositSharesToWithdraw, withdrawableShares, withdrawals, withdrawalRoots
        );

        /// 2. Move time forward to _just after_ withdrawal block
        // Expected behavior: Withdrawals are no longer slashable, so slashes no longer affect the staker
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Verify that the withdrawals are _no longer_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) > slashableUntil);
        }

        /// 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);

        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        /// 4. Complete withdrawals
        // Note: expectedTokens must be recalculated because the withdrawable shares have changed due to the slash
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, withdrawableShares);
        staker.completeWithdrawalsAsTokens(withdrawals);

        // Verify that the withdrawals were completed correctly
        for (uint i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, expectedTokens);
        }

        /// 5. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, strategies.length, "all strategies should have some shares remaining");
    }

    /**
     * @notice Test that a staker is unaffected by a slash after the withdrawal delay has passed
     */
    function testFuzz_queueTotalWithdrawal_slashAfterWithdrawalDelay_completeAsTokens(uint24 _r) public rand(_r) {
        /// 1. Queue withdrawal
        uint[] memory withdrawableShares = _calcWithdrawable(staker, strategies, initDepositShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Validate correctly queued withdrawals
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        /// 2. Move time forward to withdrawal block
        // Expected behavior: Withdrawals are no longer slashable, so slashes no longer affect the staker
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Verify that the withdrawals are _no longer_ slashable
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint32 slashableUntil = withdrawals[i].startBlock + delegationManager.minWithdrawalDelayBlocks();
            assert(uint32(block.number) > slashableUntil);
        }

        /// 3. Slash operator
        SlashingParams memory slashingParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(operator, operatorSet.id, slashingParams.strategies, slashingParams.wadsToSlash);

        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashingParams);

        /// 4. Complete withdrawals
        staker.completeWithdrawalsAsTokens(withdrawals);

        // Verify that the withdrawals were completed correctly
        for (uint i = 0; i < withdrawals.length; ++i) {
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, withdrawableShares, tokens, initTokenBalances);
        }

        /// 5. Check final state

        // Assert that all strategies have some shares remaining
        (IStrategy[] memory strats,) = delegationManager.getDepositedShares(address(staker));
        assertEq(strats.length, 0, "all strategies should have no shares remaining");
    }
}

/**
 * @notice These tests check for specific deallocation correctness around timing
 * @dev These tests assume the following:
 * - The operator is registered and allocated to the operator set
 */
contract Integration_OperatorDeallocationTiming is Integration_ALMSlashBase {
    //////////////////////////////////////////
    /// OPERATOR DEALLOCATION TIMING TESTS ///
    //////////////////////////////////////////

    function testFuzz_deallocateFully_slashBeforeDelay(uint24 _r) public rand(_r) {
        /// 1. Deallocate
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        // Validate the deallocation
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        /// 2. Move time forward to _just before_ deallocation delay
        // Expected behavior: Deallocation delay is not yet passed, so slashes can still be performed
        _rollForward_DeallocationDelay();
        rollBackward(1);
        // Verify that the operator is _still_ slashable
        check_IsSlashable_State(operator, operatorSet, strategies);

        /// 3. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }

    function testFuzz_deallocateFully_slashAfterDelay(uint24 _r) public rand(_r) {
        /// 1. Deallocate
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        // Validate the deallocation
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        /// 2. Move time forward to deallocation delay
        _rollForward_DeallocationDelay();
        // Verify that the operator is fully deallocated.
        // Note: even though the operator is technically still slashable, a slash is effectively useless as the
        // operator is no longer allocated.
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);

        /// 3. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        // Verify that the slash has no impact on the operator
        // Note: emptySlashParams remains uninitialized, as the slash _should not_ impact the operator.
        SlashingParams memory emptySlashParams;
        check_Base_Slashing_State(operator, allocateParams, emptySlashParams);
    }
}

contract Integration_OperatorDeregistrationTiming is Integration_ALMSlashBase {
    ////////////////////////////////////////////
    /// OPERATOR DEREGISTRATION TIMING TESTS ///
    ////////////////////////////////////////////

    function testFuzz_deregister_slashBeforeDelay(uint24 _r) public rand(_r) {
        /// 1. Deregister
        operator.deregisterFromOperatorSet(operatorSet);
        // Validate the deregistration
        check_Deregistration_State_PendingAllocation(operator, operatorSet);

        /// 2. Move time forward to _just before_ deregistration delay
        // Expected behavior: Deregistration delay is not yet passed, so slashes can still be performed
        _rollForward_DeallocationDelay();
        rollBackward(1);
        // Verify that the operator is _still_ slashable
        check_IsSlashable_State(operator, operatorSet, strategies);

        /// 3. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }

    function testFuzz_deregister_slashAfterDelay(uint24 _r) public rand(_r) {
        /// 1. Deregister
        operator.deregisterFromOperatorSet(operatorSet);
        // Validate the deregistration
        check_Deregistration_State_PendingAllocation(operator, operatorSet);

        /// 2. Move time forward to deregistration delay
        _rollForward_DeallocationDelay();
        // Verify that the operator is _no longer_ slashable
        check_NotSlashable_State(operator, operatorSet);

        /// 3. Slash operator
        // Note: Unlike the deallocation case, the operator is no longer registered, so a slash will revert entirely.
        slashParams = _genSlashing_Rand(operator, operatorSet);
        vm.expectRevert(IAllocationManagerErrors.OperatorNotSlashable.selector);
        avs.slashOperator(slashParams);
    }
}

/**
 * @notice These tests check for specific allocation correctness around timing
 * @dev These tests inherit from IntegrationCheckUtils instead of Integration_ALMSlashBase because they require a
 * different initialization -- specifically, the allocation must be performed within the tests. As such, there are no
 * assumptions and many state variables are declared below.
 */
contract Integration_OperatorAllocationTiming is IntegrationCheckUtils {
    AVS avs;
    User operator;
    OperatorSet operatorSet;

    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    ////////////////////////////////////////
    /// OPERATOR ALLOCATION TIMING TESTS ///
    ////////////////////////////////////////

    function _init() internal virtual override {
        /// 0. Instantiate relevant objects
        _configAssetTypes(HOLDS_LST);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        /// 1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        // Validate the deposits
        check_Deposit_State(staker, strategies, initDepositShares);

        /// 2. Delegate to an operator
        staker.delegateTo(operator);
        // Validate the delegation
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        /// 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // Validate that the operator set was correctly created
        assertTrue(allocationManager.isOperatorSet(operatorSet));
    }

    function testFuzz_register_allocate_slashBeforeDelay(uint24 _r) public rand(_r) {
        /// 1. Create and register operator
        operator.registerForOperatorSet(operatorSet);
        // Validate registration
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        /// 2. Allocate
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        // Validate the allocation
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        /// 3. Move time forward to _just before_ allocation delay
        _rollForward_AllocationDelay(operator);
        rollBackward(1); // make sure that allocation delay is not yet passed

        /// 4. Slash operator
        // Note: This slash does not revert as the operator, even though it is not allocated, is
        // still registered. However, since there is no allocation, the slash has no material effect.
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);

        // Verify that the slash has no impact on the operator
        // Note: emptySlashParams remains uninitialized, as the slash _should not_ impact the operator.
        SlashingParams memory emptySlashParams;
        check_Base_Slashing_State(operator, allocateParams, emptySlashParams);
    }

    function testFuzz_allocate_register_slashBeforeDelay(uint24 _r) public rand(_r) {
        /// 1. Allocate
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        // Validate the allocation
        check_IncrAlloc_State_NotSlashable(operator, allocateParams);

        /// 2. Create and register operator
        operator.registerForOperatorSet(operatorSet);
        // Validate the registration
        check_Registration_State_PendingAllocation(operator, allocateParams);

        /// 3. Move time forward to _just before_ allocation delay
        _rollForward_AllocationDelay(operator);
        rollBackward(1); // make sure that allocation delay is not yet passed

        /// 4. Slash operator
        // Note: This slash does not revert as the operator, even though it is not allocated, is
        // still registered. However, since there is no allocation, the slash has no material effect.
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);

        // Verify that the slash has no impact on the operator
        // Note: emptySlashParams remains uninitialized, as the slash _should not_ impact the operator.
        SlashingParams memory emptySlashParams;
        check_Base_Slashing_State(operator, allocateParams, emptySlashParams);
    }

    function testFuzz_register_allocate_slashAfterDelay(uint24 _r) public rand(_r) {
        /// 1. Create and register operator
        operator.registerForOperatorSet(operatorSet);
        // Validate the registration
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        /// 2. Allocate
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        // Validate the allocation
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        /// 3. Move time forward to after the allocation delay completes
        _rollForward_AllocationDelay(operator);

        /// 4. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }

    function testFuzz_allocate_register_slashAfterDelay(uint24 _r) public rand(_r) {
        /// 1. Allocate
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        // Validate the allocation
        check_IncrAlloc_State_NotSlashable(operator, allocateParams);

        /// 2. Create and register operator
        operator.registerForOperatorSet(operatorSet);
        // Validate the registration
        check_Registration_State_PendingAllocation(operator, allocateParams);

        /// 3. Move time forward to after the allocation delay completes
        _rollForward_AllocationDelay(operator);

        /// 4. Slash operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        avs.slashOperator(slashParams);
        // Verify that the slash was performed correctly
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }
}
