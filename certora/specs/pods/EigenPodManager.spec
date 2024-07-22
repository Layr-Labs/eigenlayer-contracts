import "../setup.spec";

// verifies that podOwnerShares[podOwner] is never a non-whole Gwei amount
invariant podOwnerSharesAlwaysWholeGweiAmount(address podOwner)
    get_podOwnerShares(podOwner) % 1000000000 == 0
    { preserved with (env e) {   
        require !isPrivilegedSender(e); }
    }

// verifies that ownerToPod[podOwner] is set once (when podOwner deploys a pod), and can otherwise never be updated
rule podAddressNeverChanges(address podOwner) {
    address podAddressBefore = get_podByOwner(podOwner);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    address podAddressAfter = get_podByOwner(podOwner);
    assert(podAddressBefore == 0 || podAddressBefore == podAddressAfter,
        "pod address changed after being set!");
}

// verifies that podOwnerShares[podOwner] can become negative (i.e. go from zero/positive to negative)
// ONLY as a result of a call to `recordBeaconChainETHBalanceUpdate`
rule limitationOnNegativeShares(address podOwner) {
    int256 podOwnerSharesBefore = get_podOwnerShares(podOwner);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    int256 podOwnerSharesAfter = get_podOwnerShares(podOwner);
    if (podOwnerSharesAfter < 0) {
        if (podOwnerSharesBefore >= 0) {
            assert(f.selector == sig:recordBeaconChainETHBalanceUpdate(address, int256).selector,
                "pod owner shares became negative from calling an unqualified function!");
        } else {
            assert(
                (podOwnerSharesAfter >= podOwnerSharesBefore) ||
                (f.selector == sig:recordBeaconChainETHBalanceUpdate(address, int256).selector),
                "pod owner had negative shares decrease inappropriately"
            );
        }
    }
    // need this line to keep the prover happy :upside_down_face:
    assert(true);
}

////******************** Added by Certora *************//////////

rule whoCanChangePodOwnerShares(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    address owner;
    int256 sharesBefore = get_podOwnerShares(owner);
    
    calldataarg args;
    f(e, args);
    int256 sharesAfter = get_podOwnerShares(owner);

    assert sharesAfter > sharesBefore => canIncreasePodOwnerShares(f);
    assert sharesAfter < sharesBefore => canDecreasePodOwnerShares(f);
}

