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

    function testFuzz_allocate_deallocate(
        uint24 _random
    ) public rand(_random) {
        // 1. Allocate to the operator set. The allocation will be slashable at the effect block
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_NotSlashable_Allocation_State(operator, allocateParams);

        // 2. Roll to the allocation's effect block. The allocation remains unslashable (operator not registered)
        _rollForward_AllocationDelay(operator);

        // 3. Deallocate fully from the operator set
        IAllocationManagerTypes.AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_NotSlashable_Deallocation_State(operator, allocateParams, deallocateParams);

        // 4. Check the operator is fully deallocated. We shouldn't need to roll to the effect block
        // because the operator is not slashable
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_register_allocate_deallocate(
        uint24 _random
    ) public rand(_random) {
        // 1. Register for the operator set before allocating
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 2. Allocate to the operator set. The allocation will be slashable at the effect block
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_Slashable_Allocation_State(operator, allocateParams, initDepositShares);

        // 3. Roll to the allocation's effect block. The allocation becomes slashable
        _rollForward_AllocationDelay(operator);

        // 4. Deallocate fully from the operator set
        IAllocationManagerTypes.AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_Slashable_Deallocation_State(operator, deallocateParams, initDepositShares);

        // 5. Roll forward to the deallocation's effect block and check the operator is fully deallocated
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_allocate_registerWhilePending_deallocate(
        uint24 _random
    ) public rand(_random) {
        // 1. Allocate 100% to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_NotSlashable_Allocation_State(operator, allocateParams);

        // 2. Register for the operator set while the allocation is still pending
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_PendingAllocation(operator, allocateParams);

        // 3. Roll to the allocation's effect block. The allocation becomes slashable
        _rollForward_AllocationDelay(operator);

        // 4. Deallocate fully from the operator set
        IAllocationManagerTypes.AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_Slashable_Deallocation_State(operator, deallocateParams, initDepositShares);

        // 5. Roll forward to the deallocation's effect block and check the operator is fully deallocated
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_allocate_registerWhenAllocated_deallocate(
        uint24 _random
    ) public rand(_random) {
        // 1. Allocate 100% to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_NotSlashable_Allocation_State(operator, allocateParams);
        
        // 2. Roll to the allocation's effect block
        _rollForward_AllocationDelay(operator);

        // 3. Register for the operator set. The allocation immediately becomes slashable
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_ActiveAllocation(operator, allocateParams);

        // 4. Deallocate fully from the operator set
        IAllocationManagerTypes.AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_Slashable_Deallocation_State(operator, deallocateParams, initDepositShares);

        // 5. Roll forward to the deallocation's effect block and check the operator is fully deallocated
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    // function testFuzz_allocatePartial_registerWhenActive(
    //     uint24 _random
    // ) public rand(_random) {
    //     // 1. Allocate 100% to a random subset of strategies
    //     IStrategy[] memory unallocated;
    //     (allocateParams, unallocated) = _genAllocation_RandStrats(operator, operatorSet, 1e18);
    //     operator.modifyAllocations(allocateParams);
    //     check_NotSlashable_Allocation_State(operator, allocateParams);

    //     // 2. Roll to the allocation's effect block
    //     _rollForward_AllocationDelay(operator);

    //     // 3. Register for the operator set. The allocation immediately becomes slashable
    //     operator.registerForOperatorSet(operatorSet);
    //     check_Registration_State()
    // }
}