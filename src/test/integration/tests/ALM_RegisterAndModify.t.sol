// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_ALMBase is IntegrationCheckUtils {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// Shared setup:
    ///
    /// 1. Generate staker with deposited assets, operator, and AVS
    /// 2. Deposit asssets and delegate to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    function _init() internal virtual override {
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
}

contract Integration_InitRegistered is Integration_ALMBase {
    /// @dev Integration test variants that start with the operator being registered
    /// for the operator set
    function _init() internal virtual override {
        super._init();

        // Register for operator set before allocating to any strategies
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);
    }

    function testFuzz_allocate_deallocate_deregister(uint24 _r) public rand(_r) {
        // 1. Allocate to the operator set
        allocateParams = _genAllocation_Rand(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        // 2. Roll to the allocation's effect block
        _rollForward_AllocationDelay(operator);

        // 3. Deallocate fully from the operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 4. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_PendingAllocation(operator, operatorSet);

        // 5. Check the operator is fully deallocated after the deallocation delay
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_allocate_deallocate_waitDeallocate_deregister(uint24 _r) public rand(_r) {
        // 1. Allocate to the operator set
        allocateParams = _genAllocation_Rand(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        // 2. Roll to the allocation's effect block
        _rollForward_AllocationDelay(operator);

        // 3. Deallocate fully from the operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 4. Check the operator is fully deallocated after the deallocation delay
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);

        // 5. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_NoAllocation(operator, operatorSet);
    }

    function testFuzz_deregister_waitDeregister_allocate_deallocate(uint24 _r) public rand(_r) {
        // 1. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_NoAllocation(operator, operatorSet);

        // 2. Wait for deallocation delay. Operator is no longer slashable
        _rollForward_DeallocationDelay();

        // 3. Allocate to operator set
        allocateParams = _genAllocation_Rand(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_NotSlashable(operator, allocateParams);

        // 3. Wait for allocation delay
        _rollForward_AllocationDelay(operator);

        // 4. Deallocate operator from operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_NotSlashable(operator, deallocateParams);
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_deregister_allocate_waitAllocate_deallocate(uint24 _r) public rand(_r) {
        // 1. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_NoAllocation(operator, operatorSet);

        // 2. Before deregistration is complete, allocate to operator set
        // The operator should be slashable after the allocation delay
        allocateParams = _genAllocation_Rand(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        // 3. Wait for allocation delay. Operator remains slashable
        _rollForward_AllocationDelay(operator);

        // 4. Deallocate operator from operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 5. Check the operator is fully deallocated after the deallocation delay
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_deregister_allocate_waitDeregister_deallocate(uint24 _r) public rand(_r) {
        // 1. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_NoAllocation(operator, operatorSet);

        // 2. Before deregistration is complete, allocate to operator set
        // The operator should be slashable after the allocation delay
        allocateParams = _genAllocation_Rand(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        // 3. Wait for deallocation delay so operator is no longer slashable
        _rollForward_DeallocationDelay();

        // 4. Instant-deallocate operator from operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_NotSlashable(operator, deallocateParams);
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }
}

contract Integration_InitAllocated is Integration_ALMBase {
    /// @dev Integration test variants that start with the operator being allocated
    /// for the operator set
    function _init() internal virtual override {
        super._init();

        // Allocate fully to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_NotSlashable(operator, allocateParams);
    }

    function testFuzz_register_deallocate_deregister(uint24 _r) public rand(_r) {
        // 1. Register for the operator set
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_PendingAllocation(operator, allocateParams);

        // 2. Roll to the allocation's effect block
        _rollForward_AllocationDelay(operator);

        // 3. Deallocate fully from the operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 4. Deregister from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_PendingAllocation(operator, operatorSet);

        // 5. Roll forward to the deallocation's effect block and check the operator is fully deallocated
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }

    function testFuzz_waitAllocation_register_deallocate(uint24 _r) public rand(_r) {
        _rollForward_AllocationDelay(operator);

        // 1. Register for the operator set. The allocation immediately becomes slashable
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_ActiveAllocation(operator, allocateParams);

        // 2. Deallocate fully from the operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // 3. Roll forward to the deallocation's effect block and check the operator is fully deallocated
        _rollForward_DeallocationDelay();
        check_FullyDeallocated_State(operator, allocateParams, deallocateParams);
    }
}
