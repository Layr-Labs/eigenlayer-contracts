import "./EigenPodMethodsAndSimplifications.spec";

// Defines the allowed validator status transitions
definition validatorStatusTransitionAllowed(IEigenPod.VALIDATOR_STATUS statusBefore, IEigenPod.VALIDATOR_STATUS statusAfter) returns bool =
    (statusBefore == IEigenPod.VALIDATOR_STATUS.INACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.ACTIVE)
    || (statusBefore == IEigenPod.VALIDATOR_STATUS.ACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.WITHDRAWN);

/// @title Only the 2 allowed transitions of validator status occur
rule validatorStatusTransitionsCorrect(bytes32 pubkeyHash) {
    IEigenPod.VALIDATOR_STATUS statusBefore = validatorStatus(pubkeyHash);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    IEigenPod.VALIDATOR_STATUS statusAfter = validatorStatus(pubkeyHash);
    assert(
        (statusBefore == statusAfter)
        || validatorStatusTransitionAllowed(statusBefore, statusAfter),
        "disallowed validator status transition occurred"
    );
}

/// @title if a validator is marked as 'INACTIVE', then it has no other entries set in its ValidatorInfo
invariant inactiveValidatorsHaveEmptyInfo(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) => (
        get_validatorIndex(pubkeyHash) == 0
        && get_restakedBalanceGwei(pubkeyHash) == 0
        && get_mostRecentBalanceUpdateTimestamp(pubkeyHash) == 0);

/// @title _validatorPubkeyHashToInfo[validatorPubkeyHash].validatorIndex can be set initially but otherwise can't change
// this can be understood as the only allowed transitions of index being of the form: 0 => anything (otherwise the index must stay the same)
rule validatorIndexSetOnlyOnce(bytes32 pubkeyHash) {
    requireInvariant inactiveValidatorsHaveEmptyInfo(pubkeyHash);
    uint64 validatorIndexBefore = get_validatorIndex(pubkeyHash);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    uint64 validatorIndexAfter = get_validatorIndex(pubkeyHash);
    assert(validatorIndexBefore == 0 || validatorIndexAfter == validatorIndexBefore,
        "validator index modified from nonzero value");
}

/// @title once a validator has its status set to WITHDRAWN, its ‘restakedBalanceGwei’ is *and always remains* zero
invariant withdrawnValidatorsHaveZeroRestakedGwei(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) =>
        (get_restakedBalanceGwei(pubkeyHash) == 0)
;

////******************** Added by Certora *************//////////

/// @title Active validators have nonzero balance
invariant activeValidatorsHaveNonZeroRestakedGwei(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.ACTIVE) =>
        (get_restakedBalanceGwei(pubkeyHash) != 0)
;

/// @title _validatorPubkeyHashToInfo[validatorPubkeyHash].lastCheckpointed cannot be greater than LastCheckpointTimestamp
invariant lastCheckpointedNoGreaterThanLastTimestamp(bytes32 validatorPubkeyHash)
    get_validatorLastCheckpointed(validatorPubkeyHash) <= 
        max(get_lastCheckpointTimestamp(), get_currentCheckpointTimestamp())
    { preserved with (env e) 
        { require timestampsNotFromFuture(e) && validatorDataNotFromFuture(e, validatorPubkeyHash); } 
}

/// @title _validatorPubkeyHashToInfo[validatorPubkeyHash].mostRecentBalanceUpdateTimestamp can ONLY increase (or remain the same)
rule mostRecentBalanceUpdateTimestampOnlyIncreases2(env e, bytes32 validatorPubkeyHash) {
    requireInvariant lastCheckpointedNoGreaterThanLastTimestamp(validatorPubkeyHash);
    requireInvariant checkpointsTimestampRemainsCorrect();
    require timestampsNotFromFuture(e) && validatorDataNotFromFuture(e, validatorPubkeyHash);
    uint64 validatorCheckpointedBefore = get_validatorLastCheckpointed(validatorPubkeyHash);
    method f;
    calldataarg args;
    f(e, args);
    uint64 validatorCheckpointedAfter = get_validatorLastCheckpointed(validatorPubkeyHash);
    assert(validatorCheckpointedAfter >= validatorCheckpointedBefore);
}

/// @title Only specified methods can increase/decrease validator.lastTimestamped
rule whoCanChangeBalanceUpdateTimestamp(bytes32 validatorPubkeyHash, env e, method f) 
{
    requireInvariant checkpointsTimestampRemainsCorrect();
    requireInvariant inactiveValidatorsHaveEmptyInfo(validatorPubkeyHash);
    require timestampsNotFromFuture(e) && validatorDataNotFromFuture(e, validatorPubkeyHash);

    uint64 timestampBefore = get_mostRecentBalanceUpdateTimestamp(validatorPubkeyHash);
    calldataarg args;
    f(e,args);
    uint64 timestampAfter = get_mostRecentBalanceUpdateTimestamp(validatorPubkeyHash);
    
    assert timestampAfter > timestampBefore => canIncreaseBalanceUpdateTimestamp(f);
    assert timestampAfter < timestampBefore => canDecreaseBalanceUpdateTimestamp(f);
}

/// @title lastCheckpointTimestamp cannot decrease
rule lastCheckpointTSOnlyIncreases(env e) {
    requireInvariant checkpointsTimestampRemainsCorrect();
    require timestampsNotFromFuture(e);
    uint64 lastTSBefore = get_lastCheckpointTimestamp();
    method f;
    calldataarg args;
    f(e, args);
    uint64 lastTSAfter = get_lastCheckpointTimestamp();
    assert lastTSAfter >= lastTSBefore;
}

/// @title During a checkpoint the current checkpoint timestamp is always greater than last checkpoint timestamp
invariant checkpointsTimestampRemainsCorrect()
    isDuringCheckpoint() => get_lastCheckpointTimestamp() < get_currentCheckpointTimestamp() 
{ preserved with (env e) 
        { require timestampsNotFromFuture(e); } 
}

/// @title If not inside a checkpoint, validator.lastTimestamped must be lastCheckpointTimestamp (for active validator)
invariant lastCheckpointedEqualsLastChPTS(bytes32 hash)
    !isDuringCheckpoint() && validatorIsActive(hash) => 
        get_validatorLastCheckpointed(hash) == get_lastCheckpointTimestamp()
{ preserved with (env e) 
    { 
        require validatorIsActive(hash) => activeValidatorCount() > 0;
        require timestampsNotFromFuture(e);
        require activeValidatorCount() < 2^23;  //otherwise the cast overflows at EigenPod.sol:608
    } 
}

/// @title The checkpoint info is empty iff not inside a checkpoint
invariant checkpointInfoIsEmpty()
    !isDuringCheckpoint() <=> (
        currentCheckpoint().beaconBlockRoot == to_bytes32(0) &&
        currentCheckpoint().proofsRemaining == 0 &&
        currentCheckpoint().podBalanceGwei == 0 &&
        currentCheckpoint().balanceDeltasGwei == 0)
    { preserved with (env e) 
        { require timestampsNotFromFuture(e); } 
}

/// @title During a checkpoint the proofsRemaining cannot increase
rule proofsRemainingCannotIncreaseInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    uint24 proofsRemainingBefore = currentCheckpoint().proofsRemaining;
    require isDuringCheckpoint();
    calldataarg args;
    f(e, args);
    uint24 proofsRemainingAfter = currentCheckpoint().proofsRemaining;
    assert proofsRemainingBefore > 0 => proofsRemainingAfter <= proofsRemainingBefore;
}

/// @title During a checkpoint the podBalanceGwei doesnt change
rule podBalanceGweiDoesntChangeInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    uint64 podBalanceGweiBefore = currentCheckpoint().podBalanceGwei;
    calldataarg args;
    require isDuringCheckpoint();
    f(e, args);
    uint64 podBalanceGweiAfter = currentCheckpoint().podBalanceGwei;
    assert isDuringCheckpoint() => 
        podBalanceGweiBefore == podBalanceGweiAfter;
}

/// @title During a checkpoint the beaconBlockRoot doesnt change
rule beaconBlockRootDoesntChangeInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    bytes32 beaconBlockRootBefore = currentCheckpoint().beaconBlockRoot;
    require isDuringCheckpoint();
    calldataarg args;
    f(e, args);
    bytes32 beaconBlockRootAfter = currentCheckpoint().beaconBlockRoot;
    assert isDuringCheckpoint() => 
        beaconBlockRootBefore == beaconBlockRootAfter;
}

function max(uint64 a, uint64 b) returns uint64
{
    if (a > b) return a;
    return b;
}


///////////////////   IN DEVELOPMENT / OBSOLETE    ////////


// front run call to verifyStaleBalance can update currentCheckpointTimestamp to the current block.timestamp, 
// causing verifyWithdrawalCredentials to revert.
// TODO this is supposed to be violated. Currently holds probably due to a vacuity
rule frontrunning_verifyStaleBalance_verifyWithdrawalCredentials(env e1, env e2)
{
    // activeValidator count > 0 if there are active validators
    uint64 beaconTimestamp; BeaconChainProofs.StateRootProof stateRootProof;
    uint40[] validatorIndices; bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
    require nativeBalances[currentContract] < 2^40;   // to avoid issues with casting at EigenPod.sol:596
    require get_podOwnerShares() < 2^40; //to avoid overflow at EigenPodManager.sol:115

    storage initialStorage = lastStorage;
    verifyWithdrawalCredentials(e1, beaconTimestamp, 
        stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields);
    //didn't revert before
    satisfy true;

    uint64 beaconTimestamp2;
    BeaconChainProofs.StateRootProof stateRootProof2; BeaconChainProofs.ValidatorProof proof2;
    verifyStaleBalance(e2, beaconTimestamp2, stateRootProof2, proof2) at initialStorage;
    satisfy true;
    require isDuringCheckpoint();   //not sure about this require. It's violated without it

    verifyWithdrawalCredentials@withrevert(e1, beaconTimestamp, 
        stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields);
    bool reverted = lastReverted;
    satisfy true;
    
    //doesn't revert after
    assert !lastReverted;
}
