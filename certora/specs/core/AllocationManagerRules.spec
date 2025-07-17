import "./AllocationManagerValidState.spec";
import "../ptaHelpers.spec";
using DelegationManager as DelegationManager;

methods {
    function getOperatorKey(address, uint32) external returns (bytes32) envfree;
    function getOperatorSetFromKey(bytes32) external returns (AllocationManagerHarness.OperatorSet) envfree;
}

definition DEFAULT_BURN_ADDRESS() returns address = 0x00000000000000000000000000000000000E16E4;


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

invariant maxMagnitudeGECurrentMagnitude(bytes32 operatorKey, address operator, address strategy)
    currentContract.allocations[operator][operatorKey][strategy].currentMagnitude <= getMaxMagnitude(operator, strategy)
    {
        preserved with (env e) {
            requireValidState();
            SumTrackingSetup();
            requireNoOverflow(e);
            // magnitudes cannot go beyond 1e18
            requireInvariant maxMagnitudeLeqWAD(operator, strategy);
            requireInvariant currentMagnitudeLeqWAD(operatorKey, operator, strategy);
            requireInvariant encumberedMagnitudeLeqWAD(e, operator, strategy);
            requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, operator, strategy);

            requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            require (e.block.number < currentContract.allocations[operator][operatorKey][strategy].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");
        }
        preserved slashOperator(
            address avs2, 
            IAllocationManagerTypes.SlashingParams params
        ) with (env e) {
            bytes32 opKey = getOperatorKey(avs2, params.operatorSetId);
            require (opKey == operatorKey, "need to ensure that the parameter opKey is the same operatorKey as used in the invariant"); 

            requireValidState();
            SumTrackingSetup();
            requireNoOverflow(e);
            // magnitudes cannot go beyond 1e18
            requireInvariant maxMagnitudeLeqWAD(operator, strategy);
            requireInvariant currentMagnitudeLeqWAD(operatorKey, operator, strategy);
            requireInvariant encumberedMagnitudeLeqWAD(e, operator, strategy);
            requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, operator, strategy);

            requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            require (e.block.number < currentContract.allocations[operator][operatorKey][strategy].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");
        } 
    }

invariant maxMagnitudeLeqWAD(address operator, address strategy)
    getMaxMagnitude(operator, strategy) <= WAD();

invariant currentMagnitudeLeqWAD(bytes32 operatorKey, address operator, address strategy)
    currentContract.allocations[operator][operatorKey][strategy].currentMagnitude <= WAD() {
        preserved with (env e) {
            SumTrackingSetup();
            requireValidState();
            requireNoOverflow(e);
            requireInvariant maxMagnitudeLeqWAD(operator, strategy);
            requireInvariant encumberedMagnitudeLeqWAD(e, operator, strategy);
            requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            requireInvariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(operatorKey, operator, strategy);
            require (e.block.number < currentContract.allocations[operator][operatorKey][strategy].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");
        }
    }

invariant encumberedMagnitudeLeqWAD(env e, address operator, address strategy)
    getEncumberedMagnitude(e, operator, strategy) <= WAD() {
        preserved {
            requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            SumTrackingSetup();
            requireValidState();
            requireNoOverflow(e);
            requireInvariant maxMagnitudeLeqWAD(operator, strategy);
        }
    }

invariant sumOfPendingDiffCurrentMagnitudeRespectsWAD(bytes32 operatorKey, address operator, address strategy)
     currentContract.allocations[operator][operatorKey][strategy].currentMagnitude + currentContract.allocations[operator][operatorKey][strategy].pendingDiff <= WAD() {
            preserved with (env e) {
                SumTrackingSetup();
                requireValidState();
                requireNoOverflow(e);
                requireInvariant maxMagnitudeLeqWAD(operator, strategy);
                requireInvariant encumberedMagnitudeLeqWAD(e, operator, strategy);
                requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            }
            preserved modifyAllocations(
                address op,
                IAllocationManagerTypes.AllocateParams[] params
            ) with (env e) {
                // require (forall uint256 i. forall uint256 j. params[i].newMagnitudes[j] <= WAD(), "assume new magnitudes respect WAD()");

                // AllocationManager.OperatorSet opSet = getOperatorSetFromKey(operatorKey);
                // require (forall uint256 i. forall uint256 j. params[i].operatorSet.avs == opSet.avs => 
                // params[i].newMagnitudes[j] + currentContract.allocations[operator][operatorKey][strategy].pendingDiff <= WAD(), "assume new magnitude + current pending diff respect WAD()");
                
                require (e.block.number < currentContract.allocations[operator][operatorKey][strategy].effectBlock, "require to not trigger a change of values throughout computation by _getUpdatedAllocation()");

                SumTrackingSetup();
                requireValidState();
                requireNoOverflow(e);
                requireInvariant maxMagnitudeLeqWAD(operator, strategy);
                requireInvariant encumberedMagnitudeLeqWAD(e, operator, strategy);
                requireInvariant maxMagnitudeGEencumberedMagnitude(operator, strategy);
            }
        }

/// @title redistributionRecipient should not be dead address after a successful run
rule redistributionRecipientCannotBeDeadAddress(AllocationManager.OperatorSet operatorSet) {
    env e;
  
    IAllocationManagerTypes.CreateSetParams[] params;
    require (params.operatorSetId == operatorSet.id, "assume we're setting the recipient for the fixed operatorSet");

    address[] recipients;
    createRedistributingOperatorSets(e, operatorSet.avs, params, recipients);
    
    address finalRecipient = getRedistributionRecipient(e, operatorSet);
       
    assert params.length != 0 => finalRecipient != DEFAULT_BURN_ADDRESS();
}

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
    require forall address operator. forall bytes32 setKey. forall address strategy. currentContract.allocations[operator][setKey][strategy].effectBlock > e.block.number;

    require didMakeInvalidPendingDiffTransition == false;

        f(e, args);

    assert didMakeInvalidPendingDiffTransition == false;
}
