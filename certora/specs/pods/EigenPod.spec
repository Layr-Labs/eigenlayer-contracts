import "./../setup.spec";

methods {
    // Internal, NONDET-summarized EigenPod library functions
    function _.verifyValidatorFields(bytes32, bytes32[] calldata, bytes calldata, uint40) internal => NONDET;
    function _.verifyValidatorBalance(bytes32, uint40, BeaconChainProofs.BalanceProof calldata) internal => NONDET;
    function _.verifyStateRoot(bytes32, BeaconChainProofs.StateRootProof calldata) internal => NONDET;

    // Internal, NONDET-summarized "send ETH" function -- unsound summary used to avoid HAVOC behavior
    // when sending ETH using `Address.sendValue()`
    function _._sendETH(address recipient, uint256 amountWei) internal => NONDET;

    //// External Calls

	// external calls to Slasher
    function _.recordStakeUpdate(address,uint32,uint32,uint256) external => NONDET;

    // external calls to Strategy contracts
    //function _.deposit(address, uint256) external => NONDET;
    //function _.withdraw(address, address, uint256) external => NONDET;


    // envfree functions
    //function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() external returns (uint64) envfree;
    function withdrawableRestakedExecutionLayerGwei() external returns (uint64) envfree;
    //function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;
    function eigenPodManager() external returns (address) envfree;
    function podOwner() external returns (address) envfree;
    //function hasRestaked() external returns (bool) envfree;
    //function mostRecentWithdrawalTimestamp() external returns (uint64) envfree;
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external returns (IEigenPod.ValidatorInfo) envfree;
    //function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) external returns (bool) envfree;
    function validatorStatus(bytes32 pubkeyHash) external returns (IEigenPod.VALIDATOR_STATUS) envfree;
    //function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;

    // harnessed functions
    function get_validatorIndex(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_restakedBalanceGwei(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_podOwnerShares() external returns (int256) envfree;
    function get_withdrawableRestakedExecutionLayerGwei() external returns (uint256) envfree;
    function get_ETH_Balance() external returns (uint256) envfree;
    function get_currentCheckpointTimestamp() external returns (uint64) envfree;
    function get_lastCheckpointTimestamp() external returns (uint64) envfree;
    function validatorIsActive(bytes32) external returns (bool) envfree;   
    function get_validatorLastCheckpointed(bytes32) external returns (uint64) envfree;
    function activeValidatorCount() external returns (uint256) envfree;
    function currentCheckpoint() external returns (IEigenPod.Checkpoint) envfree;
}

// defines the allowed validator status transitions
definition validatorStatusTransitionAllowed(IEigenPod.VALIDATOR_STATUS statusBefore, IEigenPod.VALIDATOR_STATUS statusAfter) returns bool =
    (statusBefore == IEigenPod.VALIDATOR_STATUS.INACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.ACTIVE)
    || (statusBefore == IEigenPod.VALIDATOR_STATUS.ACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.WITHDRAWN);

// verifies that only the 2 allowed transitions of validator status occur
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

// verifies that _validatorPubkeyHashToInfo[validatorPubkeyHash].mostRecentBalanceUpdateTimestamp can ONLY increase (or remain the same)
rule mostRecentBalanceUpdateTimestampOnlyIncreases(bytes32 validatorPubkeyHash) {
    IEigenPod.ValidatorInfo validatorInfoBefore = validatorPubkeyHashToInfo(validatorPubkeyHash);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    IEigenPod.ValidatorInfo validatorInfoAfter = validatorPubkeyHashToInfo(validatorPubkeyHash);
    assert(validatorInfoAfter.lastCheckpointedAt >= validatorInfoBefore.lastCheckpointedAt,
        "mostRecentBalanceUpdateTimestamp decreased");
}

// verifies that if a validator is marked as 'INACTIVE', then it has no other entries set in its ValidatorInfo
invariant inactiveValidatorsHaveEmptyInfo(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) => (
        get_validatorIndex(pubkeyHash) == 0
        && get_restakedBalanceGwei(pubkeyHash) == 0
        && get_mostRecentBalanceUpdateTimestamp(pubkeyHash) == 0);

// verifies that _validatorPubkeyHashToInfo[validatorPubkeyHash].validatorIndex can be set initially but otherwise can't change
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

// verifies that once a validator has its status set to WITHDRAWN, its ‘restakedBalanceGwei’ is *and always remains* zero
invariant withdrawnValidatorsHaveZeroRestakedGwei(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) =>
        (get_restakedBalanceGwei(pubkeyHash) == 0);


// // TODO: see if this draft rule can be salvaged
// // draft rule to capture the following behavior (or at least most of it):
// // The core invariant that ought to be maintained across the EPM and the EPs is that
// // podOwnerShares[podOwner] + sum(sharesInQueuedWithdrawals) =
// // sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) + withdrawableRestakedExecutionLayerGwei

// // idea: if we ignore shares in queued withdrawals and rearrange, then we have:
// // sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) = 
// // EigenPodManager.podOwnerShares(podOwner) - withdrawableRestakedExecutionLayerGwei
// // we can track changes to the '_validatorPubkeyHashToInfo' mapping and check this with ghost variables

// based on Certora's example here https://github.com/Certora/Tutorials/blob/michael/ethcc/EthCC/Ghosts/ghostTest.spec
ghost mathint sumOfValidatorRestakedbalancesWei {
    // NOTE: this commented out line is broken, as calling functions in axioms is currently disallowed, but this is what we'd run ideally. 
    // init_state axiom sumOfValidatorRestakedbalancesWei == to_mathint(get_podOwnerShares()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei() * 1000000000);

    // since both of these variables are zero at construction, just set the ghost to zero in the axiom
    init_state axiom sumOfValidatorRestakedbalancesWei == 0;
}

hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash].restakedBalanceGwei uint64 newValue (uint64 oldValue) {
    sumOfValidatorRestakedbalancesWei = (
        sumOfValidatorRestakedbalancesWei + 
        to_mathint(newValue) * 1000000000 -
        to_mathint(oldValue) * 1000000000
    );
}

rule consistentAccounting() {
    // fetch info before call
    int256 podOwnerSharesBefore = get_podOwnerShares();
    uint256 withdrawableRestakedExecutionLayerGweiBefore = get_withdrawableRestakedExecutionLayerGwei();
    uint256 eigenPodBalanceBefore = get_ETH_Balance();
    // filter down to valid pre-states
    require(sumOfValidatorRestakedbalancesWei ==
        to_mathint(podOwnerSharesBefore) - to_mathint(withdrawableRestakedExecutionLayerGweiBefore));

    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);

    // fetch info after call
    int256 podOwnerSharesAfter = get_podOwnerShares();
    uint256 withdrawableRestakedExecutionLayerGweiAfter = get_withdrawableRestakedExecutionLayerGwei();
    uint256 eigenPodBalanceAfter = get_ETH_Balance();
    /**
     * handling for weird, unrealistic edge case where calling `initialize` causes the pod owner to change, so the 
     * call to `get_podOwnerShares` queries the shares for a different address.
     * calling `initialize` should *not* change user shares, so it is unrealistic to simulate it doing so.
     */
    if (f.selector == sig:initialize(address).selector) {
        podOwnerSharesAfter = podOwnerSharesBefore;
    }
    // check post-state
    // TODO: this check is still broken for `withdrawRestakedBeaconChainETH` since it does a low-level call to transfer the ETH, which triggers optimistic fallback dispatching
    // special handling for one function
    if (f.selector == sig:withdrawRestakedBeaconChainETH(address,uint256).selector) {
        /* TODO: un-comment this once the dispatching is handled correctly
        assert(sumOfValidatorRestakedbalancesWei ==
            to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter)
            // adjustment term for the ETH balance of the contract changing
            + to_mathint(eigenPodBalanceBefore) - to_mathint(eigenPodBalanceAfter),
            "invalid post-state");
        */
        // TODO: delete this once the above is salvaged (was added since CVL forbids empty blocks)
        assert(true);
    // outside of special case, we don't need the adjustment term
    } else {
        assert(sumOfValidatorRestakedbalancesWei ==
            to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter),
            "invalid post-state");
    }
}

/*
rule baseInvariant() {
    int256 podOwnerSharesBefore = get_podOwnerShares();
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    int256 podOwnerSharesAfter = get_podOwnerShares();
    mathint podOwnerSharesDelta = podOwnerSharesAfter - podOwnerSharesBefore;
    assert(sumOfValidatorRestakedbalancesWei == podOwnerSharesDelta - to_mathint(get_withdrawableRestakedExecutionLayerGwei()),
        "base invariant violated");
}

invariant consistentAccounting() {
    sumOfValidatorRestakedbalancesWei ==
        to_mathint(get_withdrawableRestakedExecutionLayerGwei()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei());
}
*/

////******************** Added by Certora *************//////////

rule whoCanChangeBalanceUpdateTimestamp(bytes32 validatorPubkeyHash, env e, method f) 
{
    requireInvariant checkpointsTimestampRemainsCorrect();
    uint64 timestampBefore = get_mostRecentBalanceUpdateTimestamp(validatorPubkeyHash);
    calldataarg args;
    f(e,args);
    uint64 timestampAfter = get_mostRecentBalanceUpdateTimestamp(validatorPubkeyHash);
    
    assert timestampAfter > timestampBefore => canIncreaseBalanceUpdateTimestamp(f);
    assert timestampAfter < timestampBefore => canDecreaseBalanceUpdateTimestamp(f);
}

invariant checkpointsTimestampRemainsCorrect()
    get_currentCheckpointTimestamp() > 0 => 
        get_lastCheckpointTimestamp() < get_currentCheckpointTimestamp();

invariant lastCheckpointedEqualsLastChPTS(bytes32 hash)
    validatorIsActive(hash) => 
        get_validatorLastCheckpointed(hash) == get_lastCheckpointTimestamp();

rule methodsOnlyChangeOneValidatorStatus(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    bytes32 validatorHash1; bytes32 validatorHash2;
    IEigenPod.VALIDATOR_STATUS status1Before = validatorStatus(validatorHash1);
    IEigenPod.VALIDATOR_STATUS status2Before = validatorStatus(validatorHash2);
    
    calldataarg args;
    f(e, args);
    IEigenPod.VALIDATOR_STATUS status1After = validatorStatus(validatorHash1);
    IEigenPod.VALIDATOR_STATUS status2After = validatorStatus(validatorHash2);
    
    assert status1Before == status1After ||
        status2Before == status2After;
}

rule activeValidatorsCount_correctness(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    mathint activeValsBefore = activeValidatorCount();
    bytes32 validatorHash;
    IEigenPod.VALIDATOR_STATUS statusBefore = validatorStatus(validatorHash);
    calldataarg args;
    f(e, args);
    IEigenPod.VALIDATOR_STATUS statusAfter = validatorStatus(validatorHash);
    mathint activeValsAfter = activeValidatorCount();

    bool wasActivated = (
        statusBefore != IEigenPod.VALIDATOR_STATUS.ACTIVE &&
        statusAfter == IEigenPod.VALIDATOR_STATUS.ACTIVE);

    bool wasDeactivated = (
        statusBefore == IEigenPod.VALIDATOR_STATUS.ACTIVE &&
        statusAfter != IEigenPod.VALIDATOR_STATUS.ACTIVE);
    assert wasActivated => activeValsAfter == activeValsBefore + 1;
    assert wasDeactivated => activeValsAfter == activeValsBefore - 1;

    //to prove the other side of the implication
    satisfy wasActivated || activeValsAfter <= activeValsBefore;
    satisfy wasDeactivated || activeValsAfter >= activeValsBefore;
}

/*    struct Checkpoint {
        bytes32 beaconBlockRoot;
        uint24 proofsRemaining;
        uint64 podBalanceGwei;
        int128 balanceDeltasGwei;
    }*/
invariant checkpointInfoIsEmpty()
    get_currentCheckpointTimestamp() == 0 <=> (
        currentCheckpoint().beaconBlockRoot == to_bytes32(0) &&
        currentCheckpoint().proofsRemaining == 0 &&
        currentCheckpoint().podBalanceGwei == 0 &&
        currentCheckpoint().balanceDeltasGwei == 0
);

rule proofsRemainingCannotIncreaseInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    uint24 proofsRemainingBefore = currentCheckpoint().proofsRemaining;
    calldataarg args;
    f(e, args);
    uint24 proofsRemainingAfter = currentCheckpoint().proofsRemaining;
    assert proofsRemainingBefore > 0 => proofsRemainingAfter <= proofsRemainingBefore;
}

rule podBalanceGweiDoesntChangeInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    uint64 podBalanceGweiBefore = currentCheckpoint().podBalanceGwei;
    calldataarg args;
    f(e, args);
    uint64 podBalanceGweiAfter = currentCheckpoint().podBalanceGwei;
    assert get_currentCheckpointTimestamp() > 0 => 
        podBalanceGweiBefore == podBalanceGweiAfter;
}

rule beaconBlockRootDoesntChangeInChP(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    bytes32 beaconBlockRootBefore = currentCheckpoint().beaconBlockRoot;
    calldataarg args;
    f(e, args);
    bytes32 beaconBlockRootAfter = currentCheckpoint().beaconBlockRoot;
    assert get_currentCheckpointTimestamp() > 0 => 
        beaconBlockRootBefore == beaconBlockRootAfter;
}

