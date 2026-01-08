import "./AllocationManagerValidState.spec";
using DelegationManager as DelegationManager;

definition DEFAULT_BURN_ADDRESS() returns address = 0x00000000000000000000000000000000000E16E4;

/*
This rule applies to all methods except modifyAllocation because this function makes valid transitions, 
but only at the memory level. first, it loads allocation from memory, makes several valid transitions, 
and then stores it in the storage again, which can lead to an overall violation.
*/
/// @title pendingDiff can transition from 0 to negative or positive, from positive zero only, from negative to negative or to zero.
/// @property pendingDiffs state transitions
rule pendingDiffStateTransitions(method f, address operator, bytes32 setKey, address strategy)
filtered {f -> f.selector != sig:modifyAllocations(address,IAllocationManagerTypes.AllocateParams[]).selector} {
    env e;
    calldataarg args;
    requireValidState();

    require didMakeInvalidPendingDiffTransition == false;

        f(e, args);

    assert didMakeInvalidPendingDiffTransition == false;
}

// this rule is for modifyAllocations where when all the effectblock are in the future, this function cannot make more than one transition at a time.
rule pendingDiffStateTransitionModifyAllocations(method f)
filtered {f -> f.selector == sig:modifyAllocations(address,IAllocationManagerTypes.AllocateParams[]).selector} {
    env e;
    calldataarg args;
    requireValidState();
    require forall address operator . forall bytes32 setKey . forall address strategy . allocationsEffectBlock[operator][setKey][strategy] > e.block.number;

    require didMakeInvalidPendingDiffTransition == false;

        f(e, args);

    assert didMakeInvalidPendingDiffTransition == false;
}

/*
slashedMagnitude / currentMagnitudeBefore >= totalDepositSharesToSlash / (operatorSharesBefore + sharesInQueueBefore)
If this is false, it implies an AVS can “overslash” an operator, i.e. they can remove a larger percentage of funds.
We prove the mathematical equivalent:
slashedMagnitude * (totalShares + sharesInQueue)  >= totalDepositSharesToSlash * currentMagnitude
*/
rule noOverslashingOfOperatorShares(env e, address avs, IAllocationManagerTypes.SlashingParams params){

    bytes32 operatorKey = getOperatorKey(avs, params.operatorSetId);

    // Assume some already proven invariants
    requireValidState();
    requireNoOverflow(e);
    requireInvariant maxMagnitudeLeqWAD(params.operator, params.strategies[0]);
    requireInvariant currentMagnitudeLeqWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant encumberedMagnitudeLeqWAD(e, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGEencumberedMagnitude(params.operator, params.strategies[0]);
    requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGECurrentMagnitude(operatorKey, params.operator, params.strategies[0]);
    
    require (e.block.number < currentContract.allocations[params.operator][operatorKey][params.strategies[0]].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");

    // cannot allocate type IAllocationManagerTypes.Allocation in CVL directly but can access fields of primitive types
    uint64 currentMagnitudeBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    int128 pendingDiffBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].pendingDiff;
    require (pendingDiffBefore == 0, "prove an underapproximation assuming no pendingDiff, as pendingDiff changes currentMagnitude throughout the computation of slashOperator in _getUpdatedAllocation()");
    uint64 maxMagnitudeBefore = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesBefore = DelegationManager.operatorShares[params.operator][params.strategies[0]];
    uint256 queuedSharesBefore = DelegationManager.getSlashableSharesInQueue(e, params.operator, params.strategies[0]);
    // Note: the totalDepositSharesToSlash usually also account for slashableSharesInQueue that have been removed
    // Due to known rounding issues in the computation of DelegationManager.getSlashableSharesInQueue() 
    // we underapproximate and assume no slashableSharesInQueue but check no overslashing for operatorShares only
    require queuedSharesBefore == 0; 
    require (operatorSharesBefore > 1, "1 share is an edge case that will always fail due to mulWadUp" );

    slashOperator(e, avs, params);

    uint64 currentMagnitudeAfter = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    uint64 maxMagnitudeAfter = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesAfter = DelegationManager.operatorShares[params.operator][params.strategies[0]];

    // currentMagnitude should not increase from a slash
    assert(currentMagnitudeBefore >= currentMagnitudeAfter);
    // operatorShares should not increase from a slash
    assert(operatorSharesBefore >= operatorSharesAfter);
    
    uint64 slashedMagnitude = require_uint64(currentMagnitudeBefore - currentMagnitudeAfter);
    uint256 totalDepositSharesToSlash = require_uint256(operatorSharesBefore - operatorSharesAfter);

    // assert slashedMagnitude / currentMagnitude >= totalDepositSharesToSlash / (operatorSharesBefore + queuedSharesBefore);
    assert slashedMagnitude != 0 => slashedMagnitude * operatorSharesBefore >= totalDepositSharesToSlash * currentMagnitudeBefore;
}


/*
slashedMagnitude / currentMagnitudeBefore >= totalDepositSharesToSlash / (operatorSharesBefore + sharesInQueueBefore)
If this is false, it implies an AVS can “overslash” an operator, i.e. they can remove a larger percentage of funds.
We prove the mathematical equivalent:
slashedMagnitude * (totalShares + sharesInQueue)  >= totalDepositSharesToSlash * currentMagnitude
*/
rule noOverslashingOfShares(env e, address avs, IAllocationManagerTypes.SlashingParams params){

    bytes32 operatorKey = getOperatorKey(avs, params.operatorSetId);

    // Assume some already proven invariants
    requireValidState();
    requireNoOverflow(e);
    requireInvariant maxMagnitudeLeqWAD(params.operator, params.strategies[0]);
    requireInvariant currentMagnitudeLeqWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant encumberedMagnitudeLeqWAD(e, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGEencumberedMagnitude(params.operator, params.strategies[0]);
    requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGECurrentMagnitude(operatorKey, params.operator, params.strategies[0]);
    
    require (e.block.number < currentContract.allocations[params.operator][operatorKey][params.strategies[0]].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");

    // cannot allocate type IAllocationManagerTypes.Allocation in CVL directly but can access fields of primitive types
    uint64 currentMagnitudeBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    int128 pendingDiffBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].pendingDiff;
    require (pendingDiffBefore == 0, "prove an underapproximation assuming no pendingDiff, as pendingDiff changes currentMagnitude throughout the computation of slashOperator in _getUpdatedAllocation()");
    uint64 maxMagnitudeBefore = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesBefore = DelegationManager.operatorShares[params.operator][params.strategies[0]];
    uint256 queuedSharesBefore = DelegationManager.getSlashableSharesInQueue(e, params.operator, params.strategies[0]);
    // Note: the totalDepositSharesToSlash usually also account for slashableSharesInQueue that have been removed
    // Due to known rounding issues in the computation of DelegationManager.getSlashableSharesInQueue() 
    // we underapproximate and assume no slashableSharesInQueue but check no overslashing for operatorShares only
    require (operatorSharesBefore + queuedSharesBefore >= 1, "+" );

    slashOperator(e, avs, params);

    uint64 currentMagnitudeAfter = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    uint64 maxMagnitudeAfter = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesAfter = DelegationManager.operatorShares[params.operator][params.strategies[0]];
    uint256 queuedSharesAfter = DelegationManager.getSlashableSharesInQueue(e, params.operator, params.strategies[0]);

    // currentMagnitude should not increase from a slash
    assert(currentMagnitudeBefore >= currentMagnitudeAfter);
    // operatorShares should not increase from a slash
    assert(operatorSharesBefore >= operatorSharesAfter);
    
    uint64 slashedMagnitude = require_uint64(currentMagnitudeBefore - currentMagnitudeAfter);
    uint256 totalDepositSharesToSlash = require_uint256(operatorSharesBefore - operatorSharesAfter + queuedSharesBefore - queuedSharesAfter);

    // assert slashedMagnitude / currentMagnitude >= totalDepositSharesToSlash / (operatorSharesBefore + queuedSharesBefore);
    assert slashedMagnitude != 0 => assert_uint256(slashedMagnitude / currentMagnitudeBefore) >= assert_uint256(totalDepositSharesToSlash / (operatorSharesBefore + queuedSharesBefore));
}

/*
slashedMagnitude / currentMagnitudeBefore >= totalDepositSharesToSlash / (operatorSharesBefore + sharesInQueueBefore)
If this is false, it implies an AVS can “overslash” an operator, i.e. they can remove a larger percentage of funds.
We prove the mathematical equivalent:
slashedMagnitude * (totalShares + sharesInQueue)  >= totalDepositSharesToSlash * currentMagnitude
*/
rule overslashingOfSharesAtMostOne(env e, address avs, IAllocationManagerTypes.SlashingParams params){

    bytes32 operatorKey = getOperatorKey(avs, params.operatorSetId);

    // Assume some already proven invariants
    requireValidState();
    requireNoOverflow(e);
    requireInvariant maxMagnitudeLeqWAD(params.operator, params.strategies[0]);
    requireInvariant currentMagnitudeLeqWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant encumberedMagnitudeLeqWAD(e, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGEencumberedMagnitude(params.operator, params.strategies[0]);
    requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, params.operator, params.strategies[0]);
    requireInvariant maxMagnitudeGECurrentMagnitude(operatorKey, params.operator, params.strategies[0]);
    
    require (e.block.number < currentContract.allocations[params.operator][operatorKey][params.strategies[0]].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");

    // cannot allocate type IAllocationManagerTypes.Allocation in CVL directly but can access fields of primitive types
    uint64 currentMagnitudeBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    int128 pendingDiffBefore = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].pendingDiff;
    require (pendingDiffBefore == 0, "prove an underapproximation assuming no pendingDiff, as pendingDiff changes currentMagnitude throughout the computation of slashOperator in _getUpdatedAllocation()");
    uint64 maxMagnitudeBefore = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesBefore = DelegationManager.operatorShares[params.operator][params.strategies[0]];
    uint256 queuedSharesBefore = DelegationManager.getSlashableSharesInQueue(e, params.operator, params.strategies[0]);
    // Note: the totalDepositSharesToSlash usually also account for slashableSharesInQueue that have been removed
    // Due to known rounding issues in the computation of DelegationManager.getSlashableSharesInQueue() 
    // we underapproximate and assume no slashableSharesInQueue but check no overslashing for operatorShares only
    require (operatorSharesBefore + queuedSharesBefore >= 1, "+" );

    slashOperator(e, avs, params);

    uint64 currentMagnitudeAfter = currentContract.allocations[params.operator][operatorKey][params.strategies[0]].currentMagnitude;
    uint64 maxMagnitudeAfter = getMaxMagnitude(params.operator, params.strategies[0]);

    uint256 operatorSharesAfter = DelegationManager.operatorShares[params.operator][params.strategies[0]];
    uint256 queuedSharesAfter = DelegationManager.getSlashableSharesInQueue(e, params.operator, params.strategies[0]);

    // currentMagnitude should not increase from a slash
    assert(currentMagnitudeBefore >= currentMagnitudeAfter);
    // operatorShares should not increase from a slash
    assert(operatorSharesBefore >= operatorSharesAfter);
    
    uint64 slashedMagnitude = require_uint64(currentMagnitudeBefore - currentMagnitudeAfter);
    uint256 totalDepositSharesToSlash = require_uint256(operatorSharesBefore - operatorSharesAfter + queuedSharesBefore - queuedSharesAfter);

    // assert slashedMagnitude / currentMagnitude >= totalDepositSharesToSlash / (operatorSharesBefore + queuedSharesBefore);
    assert slashedMagnitude != 0 => assert_uint256(slashedMagnitude / currentMagnitudeBefore) + 1 >= assert_uint256(totalDepositSharesToSlash / (operatorSharesBefore + queuedSharesBefore));
}


/// @title redistributionRecipient should not be dead address after a successful run
rule redistributionRecipientCannotBeDeadAddress(AllocationManager.OperatorSet operatorSet) {
    env e;
  
    IAllocationManagerTypes.CreateSetParams[] params;
    require (params[0].operatorSetId == operatorSet.id, "assume we're setting the recipient for the fixed operatorSet");

    address[] recipients;
    createRedistributingOperatorSets(e, operatorSet.avs, params, recipients);
    
    address finalRecipient = getRedistributionRecipient(e, operatorSet);
       
    assert params.length != 0 => finalRecipient != DEFAULT_BURN_ADDRESS();
}
