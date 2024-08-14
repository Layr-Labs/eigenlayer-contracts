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

    satisfy canIncreasePodOwnerShares(f) => sharesAfter > sharesBefore;
    satisfy canDecreasePodOwnerShares(f) => sharesAfter < sharesBefore;
}

invariant noPodNoShares(address owner)
    get_podByOwner(owner) == 0 => get_podOwnerShares(owner) == 0;

rule addShares_reverts(env e)
{
    uint256 shares;
    address owner;
    addShares@withrevert(e, owner, shares); 
    bool reverted = lastReverted;
    satisfy !reverted;
    assert shares == 0 => reverted;
}

rule removeShares_reverts(env e)
{
    uint256 shares;
    address owner;
    removeShares@withrevert(e, owner, shares); 
    bool reverted = lastReverted;
    assert shares == 0 => reverted;
}

rule addShares_additivity(env e)
{
    uint256 shares1; uint256 shares2; uint256 sharesSum;
    require shares1 + shares2 == sharesSum * 1;
    address owner;
    storage init = lastStorage;
    addShares(e, owner, shares1); 
    addShares(e, owner, shares2);
    storage after12 = lastStorage;
    addShares(e, owner, sharesSum) at init;
    storage afterSum = lastStorage;
    assert after12 == afterSum;
}

rule removeShares_additivity(env e)
{
    uint256 shares1; uint256 shares2; uint256 sharesSum;
    require shares1 + shares2 == sharesSum * 1;
    address owner;
    storage init = lastStorage;
    removeShares(e, owner, shares1); 
    removeShares(e, owner, shares2);
    storage after12 = lastStorage;
    removeShares(e, owner, sharesSum) at init;
    storage afterSum = lastStorage;
    assert after12 == afterSum;
}

rule withdrawShares_additivity(env e)
{
    uint256 shares1; uint256 shares2; uint256 sharesSum;
    require shares1 + shares2 == sharesSum * 1;
    address owner; address receiver;
    storage init = lastStorage;
    withdrawSharesAsTokens(e, owner, receiver, shares1); 
    withdrawSharesAsTokens(e, owner, receiver, shares2);
    storage after12 = lastStorage;
    withdrawSharesAsTokens(e, owner, receiver, sharesSum) at init;
    storage afterSum = lastStorage;
    assert after12 == afterSum;
}

rule add_remove_inverse(env e)
{
    uint256 shares;
    address owner;
    mathint sharesBefore = get_podOwnerShares(owner);
    addShares(e, owner, shares); 
    removeShares(e, owner, shares);
    mathint sharesAfter = get_podOwnerShares(owner);
    assert sharesBefore == sharesAfter;
}

rule addShares_integrity(env e)
{
    uint256 shares;
    address owner;
    mathint sharesBefore = get_podOwnerShares(owner);
    addShares(e, owner, shares); 
    mathint sharesAfter = get_podOwnerShares(owner);
    assert sharesBefore + shares == sharesAfter;
}

rule removeShares_integrity(env e)
{
    uint256 shares;
    address owner;
    mathint sharesBefore = get_podOwnerShares(owner);
    removeShares(e, owner, shares); 
    mathint sharesAfter = get_podOwnerShares(owner);
    assert sharesBefore - shares == sharesAfter;
}

rule addShares_independence(env e)
{
    uint256 shares1; uint256 shares2;
    address owner;
    storage init = lastStorage;
    addShares(e, owner, shares1);
    addShares(e, owner, shares2); 
    storage storageAfter12 = lastStorage;

    addShares(e, owner, shares2) at init;
    addShares(e, owner, shares1); 
    storage storageAfter21 = lastStorage;
    assert storageAfter12 == storageAfter21;
}

rule add_remove_independence(env e)
{
    uint256 shares1; uint256 shares2;
    address owner;
    storage init = lastStorage;
    addShares(e, owner, shares1);
    removeShares(e, owner, shares2); 
    storage storageAfter12 = lastStorage;

    removeShares(e, owner, shares2) at init;
    addShares(e, owner, shares1); 
    storage storageAfter21 = lastStorage;
    assert storageAfter12 == storageAfter21;
}

rule addShares_revertsWhenNoPod(env e)
{
    uint256 shares; 
    address owner;
    bool hasPod = get_podByOwner(owner) != 0;
    satisfy true;
    addShares@withrevert(e, owner, shares);
    bool reverted = lastReverted;

    satisfy !hasPod && reverted;
    assert !hasPod => reverted;
}

rule removeShares_revertsWhenNoPod(env e)
{
    uint256 shares; 
    address owner;
    bool hasPod = get_podByOwner(owner) != 0;
    removeShares@withrevert(e, owner, shares);
    bool reverted = lastReverted;
    assert !hasPod => reverted;
}

////////////////////  IN DEVELOPMENT / OBSOLETE   //////////////


// to check the conditions under which the method works correctly
// loop_iter, hashing_length_bound, optimistic_loop, optimistic_hashing
rule addShares_alwaysReverts(env e)
{
    uint256 shares;
    address owner;
    addShares@withrevert(e, owner, shares); 
    bool reverted = lastReverted;
    assert reverted;
}

// to check the conditions under which the method works correctly
// loop_iter, hashing_length_bound, optimistic_loop, optimistic_hashing
rule methodsAlwaysRevert(env e, method f) filtered 
    {
        f -> 
            f.selector == sig:EigenPodManagerHarness.initialize(address, address, uint256).selector ||
            f.selector == sig:EigenPodManagerHarness.addShares(address,uint256).selector ||
            f.selector == sig:EigenPodManagerHarness.recordBeaconChainETHBalanceUpdate(address,int256).selector ||
            f.selector == sig:DummyEigenPodA.verifyWithdrawalCredentials(uint64, BeaconChainProofs.StateRootProof, uint40[], bytes[], bytes32[][]).selector
    }
{
    calldataarg args;
    f@withrevert(e, args);
    bool reverted = lastReverted;
    //satisfy !reverted;
    assert reverted;
}

// to check the conditions under which the method works correctly
// loop_iter, hashing_length_bound, optimistic_loop, optimistic_hashing
rule methodsDontAlwaysRevert(env e, method f) 
{
    calldataarg args;
    f@withrevert(e, args);
    bool reverted = lastReverted;
    satisfy !reverted;
}