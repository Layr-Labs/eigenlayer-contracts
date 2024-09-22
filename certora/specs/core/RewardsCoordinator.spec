
using ERC20 as token1;
using ERC20PresetFixedSupply as token2;

methods {
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

    function _.createAVSRewardsSubmission(IRewardsCoordinator.RewardsSubmission[]) external => DISPATCHER(true);
    function _.createRewardsForAllSubmission(IRewardsCoordinator.RewardsSubmission[]) external => DISPATCHER(true);
    function _.submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) external => DISPATCHER(true);
    function _.processClaim(IRewardsCoordinator.RewardsMerkleClaim claim, address recipient) external => DISPATCHER(true);

    function _.verifyInclusionKeccak(bytes memory, bytes32, bytes32, uint256) internal => CONSTANT;
    function cumulativeClaimed(address, address) external returns(uint256) envfree;
    function submissionNonce(address) external returns(uint256) envfree;

    function _.strategyIsWhitelistedForDeposit(address) external => NONDET;
    function _.isPauser(address) external => NONDET;
    function _.unpauser() external => NONDET;
    function _._ external => DISPATCH [
        _.transferFrom(address, address, uint256),
        _.transfer(address, uint256)
    ] default NONDET;
}


// Given a processClaim passes, second processClaim reverts
// Assumptions
// - checkClaim(claim) == true and does not revert
// - claim has non-empty tokenLeaves
/// status: pass
rule claimWithduplicateTokenLeafs(env e, IRewardsCoordinator.RewardsMerkleClaim claim, address recipient) {
    require claim.tokenLeaves.length == 1;
    
    checkClaim(e, claim);

    if ((claimerFor(e, claim.earnerLeaf.earner)) != 0) {
        require e.msg.sender == claimerFor(e, claim.earnerLeaf.earner);
    } else {
        require e.msg.sender == claim.earnerLeaf.earner;
    }

    processClaim(e, claim, recipient);
    processClaim@withrevert(e, claim, recipient);

	assert (
        lastReverted,
        "First claim passing with token claims means second claim should revert"
    );
}

/// status: pass
rule cumulativeClaimedStrictlyIncreasing(env e, address claimToken, address earner) {
    uint256 cumulativeClaimedBefore = cumulativeClaimed(earner, claimToken);
    calldataarg args;
    method f;
    f(e, args);
    uint256 cumulativeClaimedAfter = cumulativeClaimed(earner, claimToken);

    assert(cumulativeClaimedAfter >= cumulativeClaimedBefore);
}


/// Check that transfer amount in RewardsCoordinator.prcocessClaim() is equal to leaf.cumulativeEarnings - cumulativeClaimed
/// cumulativeEarnings - cumulativeClaimed == balanceAfter - balanceBefore
/// Initially failed from havoc from token safeTransfer in the calldata. Fixed by having unresolved catch-all calls defined
/// status: pass
rule transferredTokensFromClaim(env e, IRewardsCoordinator.RewardsMerkleClaim claim, address recipient) {
    require claim.tokenLeaves.length == 2;
    require token1 == claim.tokenLeaves[0].token;
    require token2 == claim.tokenLeaves[1].token;
    // if recipient is the RewardsCoordinator, rule breaks since balanceAfter == balanceBefore 
    require recipient != currentContract;

    address earner = claim.earnerLeaf.earner;
    uint256 token1BalanceBefore = token1.balanceOf(e, recipient);
    uint256 token2BalanceBefore = token2.balanceOf(e, recipient);
    uint256 token1CumulativeClaimed = cumulativeClaimed(earner, claim.tokenLeaves[0].token);
    uint256 token2CumulativeClaimed = cumulativeClaimed(earner, claim.tokenLeaves[1].token);

    processClaim(e, claim, recipient);

    uint256 token1BalanceAfter = token1.balanceOf(e, recipient);
    uint256 token2BalanceAfter = token2.balanceOf(e, recipient);

    if (token1 == token2) {
        // In practice this shouldn't occur as we will not construct leaves with multiple tokenLeaves with the same token address
        assert claim.tokenLeaves[1].cumulativeEarnings - token1CumulativeClaimed == token2BalanceAfter - token2BalanceBefore;
    } else {
        assert claim.tokenLeaves[0].cumulativeEarnings - token1CumulativeClaimed == token1BalanceAfter - token1BalanceBefore;
        assert claim.tokenLeaves[1].cumulativeEarnings - token2CumulativeClaimed == token2BalanceAfter - token2BalanceBefore;
    }
}

/// RewardsCoordinator.checkClaim == True => processClaim success 
/// status: pass
rule claimCorrectness(env e, IRewardsCoordinator.RewardsMerkleClaim claim, address recipient) {
    // Call checkClaim to ensure it does not revert in rule since it never returns false
    checkClaim(e, claim);

    bool canWork = claimerFor(e, claim.earnerLeaf.earner) != 0 ? 
        e.msg.sender == claimerFor(e, claim.earnerLeaf.earner) :
        e.msg.sender == claim.earnerLeaf.earner;

    processClaim@withrevert(e, claim, recipient);

    assert !lastReverted => canWork;
}

/// checkClaim either returns True or reverts otherwise
invariant checkClaimNeverFalse(env e, IRewardsCoordinator.RewardsMerkleClaim claim)
    !checkClaim(e, claim);

rule rewardsSubmissionIntegrity(env e) 
    {
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmission1;
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmission2;
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmissions;

    // requiring that rewardsSubmissions = [rewardsSubmission1, rewardsSubmission2]
    require rewardsSubmission1.length == 1;
    require rewardsSubmission2.length == 1;
    require rewardsSubmissions.length == 2;
    require rewardsSubmission1[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmission2[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmissions[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmissions[1].strategiesAndMultipliers.length == 1;
    require rewardsSubmission1[0].strategiesAndMultipliers[0].strategy 
        == rewardsSubmissions[0].strategiesAndMultipliers[0].strategy;
    require rewardsSubmission2[0].strategiesAndMultipliers[0].strategy 
        == rewardsSubmissions[1].strategiesAndMultipliers[0].strategy;
    require rewardsSubmission1[0].strategiesAndMultipliers[0].multiplier 
        == rewardsSubmissions[0].strategiesAndMultipliers[0].multiplier;
    require rewardsSubmission2[0].strategiesAndMultipliers[0].multiplier 
        == rewardsSubmissions[1].strategiesAndMultipliers[0].multiplier;
    require rewardsSubmission1[0].token == rewardsSubmissions[0].token;
    require rewardsSubmission2[0].token == rewardsSubmissions[1].token;
    require rewardsSubmission1[0].amount == rewardsSubmissions[0].amount;
    require rewardsSubmission2[0].amount == rewardsSubmissions[1].amount;
    require rewardsSubmission1[0].startTimestamp == rewardsSubmissions[0].startTimestamp;
    require rewardsSubmission2[0].startTimestamp == rewardsSubmissions[1].startTimestamp;
    require rewardsSubmission1[0].duration == rewardsSubmissions[0].duration;
    require rewardsSubmission2[0].duration == rewardsSubmissions[1].duration;

    storage init = lastStorage;

    createAVSRewardsSubmission(e, rewardsSubmission1);
    createAVSRewardsSubmission(e, rewardsSubmission2);

    storage splitFinal = lastStorage;

    createAVSRewardsSubmission(e, rewardsSubmissions) at init;

    assert(lastStorage == splitFinal);
}

rule allRewardsSubmissionIntegrity(env e) 
    {
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmission1;
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmission2;
    IRewardsCoordinator.RewardsSubmission[] rewardsSubmissions;

    // requiring that rewardsSubmissions = [rewardsSubmission1, rewardsSubmission2]
    require rewardsSubmission1.length == 1;
    require rewardsSubmission2.length == 1;
    require rewardsSubmissions.length == 2;
    require rewardsSubmission1[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmission2[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmissions[0].strategiesAndMultipliers.length == 1;
    require rewardsSubmissions[1].strategiesAndMultipliers.length == 1;
    require rewardsSubmission1[0].strategiesAndMultipliers[0].strategy 
        == rewardsSubmissions[0].strategiesAndMultipliers[0].strategy;
    require rewardsSubmission2[0].strategiesAndMultipliers[0].strategy 
        == rewardsSubmissions[1].strategiesAndMultipliers[0].strategy;
    require rewardsSubmission1[0].strategiesAndMultipliers[0].multiplier 
        == rewardsSubmissions[0].strategiesAndMultipliers[0].multiplier;
    require rewardsSubmission2[0].strategiesAndMultipliers[0].multiplier 
        == rewardsSubmissions[1].strategiesAndMultipliers[0].multiplier;
    require rewardsSubmission1[0].token == rewardsSubmissions[0].token;
    require rewardsSubmission2[0].token == rewardsSubmissions[1].token;
    require rewardsSubmission1[0].amount == rewardsSubmissions[0].amount;
    require rewardsSubmission2[0].amount == rewardsSubmissions[1].amount;
    require rewardsSubmission1[0].startTimestamp == rewardsSubmissions[0].startTimestamp;
    require rewardsSubmission2[0].startTimestamp == rewardsSubmissions[1].startTimestamp;
    require rewardsSubmission1[0].duration == rewardsSubmissions[0].duration;
    require rewardsSubmission2[0].duration == rewardsSubmissions[1].duration;

    storage init = lastStorage;

    createRewardsForAllSubmission(e, rewardsSubmission1);
    createRewardsForAllSubmission(e, rewardsSubmission2);

    storage splitFinal = lastStorage;

    createRewardsForAllSubmission(e, rewardsSubmissions) at init;

    assert(lastStorage == splitFinal);
}

rule balanceIncreasesOnRewardsSubmission(IRewardsCoordinator.RewardsSubmission[] rewardsSubmission, env e) {
    
    uint256 balance1Before = token1.balanceOf(e, currentContract);
    uint256 balance2Before = token2.balanceOf(e, currentContract);
    
    createAVSRewardsSubmission(e, rewardsSubmission);

    uint256 balance1After = token1.balanceOf(e, currentContract);
    uint256 balance2After = token2.balanceOf(e, currentContract);

    assert(balance1After >= balance1Before);
    assert(balance2After >= balance2Before);
}

rule submissionNonceMonontonic( address account) {
    uint256 nonceBefore = submissionNonce(account);
    
    calldataarg args;
    method f; env e;
    f(e, args);

    uint256 nonceAfter = submissionNonce(account);

    assert(nonceAfter >= nonceBefore);
}