import "./EigenPodMethodsAndSimplifications.spec";

// Tracks changes to _validatorPubkeyHashToInfo[key].status
ghost mathint validatorsActivated {
    init_state axiom validatorsActivated == 0;
}

// Tracks changes to _validatorPubkeyHashToInfo[key].status in validatorsActivated
hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash].status IEigenPod.VALIDATOR_STATUS newValue (IEigenPod.VALIDATOR_STATUS oldValue) 
{
    if (oldValue != IEigenPod.VALIDATOR_STATUS.ACTIVE && newValue == IEigenPod.VALIDATOR_STATUS.ACTIVE) validatorsActivated = validatorsActivated + 1;
    if (oldValue == IEigenPod.VALIDATOR_STATUS.ACTIVE && newValue != IEigenPod.VALIDATOR_STATUS.ACTIVE) validatorsActivated = validatorsActivated - 1;
}

/// @title activeValidatorsCount equals to count(validator v where v.isActive)
rule activeValidatorsCount_correctness(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    require validatorsActivated == 0;
    mathint activeValsBefore = activeValidatorCount();
    bytes32 validatorHash;
    calldataarg args;
    f(e, args);
    mathint activeValsAfter = activeValidatorCount();

    assert activeValsAfter - activeValsBefore == validatorsActivated;
}


///////////////////   IN DEVELOPMENT / OBSOLETE    ////////


ghost mathint sumOfValidatorRestakedbalancesWei 
{
    // NOTE: this commented out line is broken, as calling functions in axioms is currently disallowed, but this is what we'd run ideally. 
    // init_state axiom sumOfValidatorRestakedbalancesWei == to_mathint(get_podOwnerShares()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei() * 1000000000);
    // since both of these variables are zero at construction, just set the ghost to zero in the axiom

    // Certora: since we only track changes, the zero initial value is fine
    //          If we want the specified initial value, we can make it undetermined here 
    //          and set it at the beginning of each rule that uses the ghost
    init_state axiom sumOfValidatorRestakedbalancesWei == 0;
}

hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash].restakedBalanceGwei uint64 newValue (uint64 oldValue) {
    sumOfValidatorRestakedbalancesWei = (
        sumOfValidatorRestakedbalancesWei + 
        to_mathint(newValue) * 1000000000 -
        to_mathint(oldValue) * 1000000000
    );
}

// Rule to capture the following behavior (or at least most of it):
 //  The core invariant that ought to be maintained across the EPM and the EPs is that
 //  podOwnerShares[podOwner] + sum(sharesInQueuedWithdrawals) =
 //  sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) + withdrawableRestakedExecutionLayerGwei
 //  idea: if we ignore shares in queued withdrawals and rearrange, then we have:
 //  sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) = 
 //  EigenPodManager.podOwnerShares(podOwner) - withdrawableRestakedExecutionLayerGwei
 //  we can track changes to the '_validatorPubkeyHashToInfo' mapping and check this with ghost variables
/*
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
    
    //  handling for weird, unrealistic edge case where calling `initialize` causes the pod owner to change, so the 
    //  call to `get_podOwnerShares` queries the shares for a different address.
    //  calling `initialize` should *not* change user shares, so it is unrealistic to simulate it doing so.
    
    if (f.selector == sig:initialize(address).selector) {
        podOwnerSharesAfter = podOwnerSharesBefore;
    }
    // check post-state
    // TODO: this check is still broken for `withdrawRestakedBeaconChainETH` since it does a low-level call to transfer the ETH, which triggers optimistic fallback dispatching
    // special handling for one function
    if (f.selector == sig:withdrawRestakedBeaconChainETH(address,uint256).selector) {
        // TODO: un-comment this once the dispatching is handled correctly
        // assert(sumOfValidatorRestakedbalancesWei ==
        //     to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter)
        //     // adjustment term for the ETH balance of the contract changing
        //     + to_mathint(eigenPodBalanceBefore) - to_mathint(eigenPodBalanceAfter),
        //     "invalid post-state");
        
        // TODO: delete this once the above is salvaged (was added since CVL forbids empty blocks)
        assert(true);
    // outside of special case, we don't need the adjustment term
    } else {
        assert(sumOfValidatorRestakedbalancesWei ==
            to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter),
            "invalid post-state");
    }
}
*/

//get_withdrawableRestakedExecutionLayerGwei == podOwnerShares() - withdrawableRestakedExecutionLayerGwei
rule baseInvariant(method f, env e) 
{
    require sumOfValidatorRestakedbalancesWei == 0;
    int256 podOwnerSharesBefore = get_podOwnerShares();
    uint256 restakedBefore = get_withdrawableRestakedExecutionLayerGwei();
    mathint deltaBefore = podOwnerSharesBefore - restakedBefore;
    calldataarg args;
    f(e, args);
    int256 podOwnerSharesAfter = get_podOwnerShares();
    uint256 restakedAfter = get_withdrawableRestakedExecutionLayerGwei();
    mathint deltaAfter = podOwnerSharesAfter - restakedAfter;
    assert sumOfValidatorRestakedbalancesWei == deltaAfter - deltaBefore;
}
