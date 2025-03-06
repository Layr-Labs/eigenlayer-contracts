import "./AllocationManagerValidState.spec";

use invariant maxMagnitudeHistoryKeysMonotonicInc;
use invariant maxMagnitudeHistoryKeysLessThanCurrentBlock;
use invariant maxMagnitudeMonotonicallyDecreasing;
use invariant SetInRegisteredIFFStatusIsTrue;
use invariant encumberedMagnitudeEqSumOfCurrentMagnitudesAndPositivePending;
use invariant negativePendingDiffAtMostCurrentMagnitude;
use invariant deallocationQueueDataUniqueness;
use invariant noZeroKeyInDealocationQ;
use invariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne;
use invariant deallocationQueueEffectBlockAscesndingOrder;
use invariant noPositivePendingDiffInDeallocationQ;
use invariant effectBlockZeroHasNoPendingDiff;

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
