// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashingAllocations is IntegrationCheckUtils {
    
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// Shared setup:
    /// 
    /// 1. Generate staker with deposited assets, operator, and AVS
    /// 2. Deposit asssets and delegate to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    function _init() internal override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set containing the strategies held by the staker
        operatorSet = avs.createOperatorSet(strategies);
    }

    function testFuzz_register_allocate(
        uint24 _random
    ) public rand(_random) {
        operator.registerForOperatorSet(operatorSet);
        check_Unallocated_Registration_State(operator, operatorSet);

        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams); // idea: operator.allocateHalfAvailable(operatorSet)?
        check_Slashable_Allocation_State(operator, allocateParams, initDepositShares);

        // // idea - do this before/after the allocation is completable
        // // before - instant deallocation
        // // after - queued deallocation
        // operator.deregisterFromOperatorSet(operatorSet);
        // check_Allocated_Deregistration_State(operator, operatorSet);

        // _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);


    }

    function testFuzz_allocate_registerWhilePending(
        uint24 _random
    ) public rand(_random) {
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams); // idea: operator.allocateHalfAvailable(operatorSet)?
        check_NotSlashable_Allocation_State(operator, allocateParams);
        
        // _rollForward_AllocationDelay(operator);

        operator.registerForOperatorSet(operatorSet);
        check_PendingAllocated_Registration_State(operator, operatorSet, allocateParams, initDepositShares);

        // // idea - do this before/after the allocation is completable
        // // before - instant deallocation
        // // after - queued deallocation
        // operator.deregisterFromOperatorSet(operatorSet);
        // check_Allocated_Deregistration_State(operator, operatorSet);

        


    }

    function testFuzz_allocate_registerWhenAllocated(
        uint24 _random
    ) public rand(_random) {
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams); // idea: operator.allocateHalfAvailable(operatorSet)?
        check_NotSlashable_Allocation_State(operator, allocateParams);
        
        _rollForward_AllocationDelay(operator);

        operator.registerForOperatorSet(operatorSet);
        check_Allocated_Registration_State(operator, operatorSet, allocateParams, initDepositShares);

        // // idea - do this before/after the allocation is completable
        // // before - instant deallocation
        // // after - queued deallocation
        // operator.deregisterFromOperatorSet(operatorSet);
        // check_Allocated_Deregistration_State(operator, operatorSet);

        


    }
}