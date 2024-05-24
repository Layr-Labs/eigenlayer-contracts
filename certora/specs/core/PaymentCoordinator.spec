
using ERC20 as token;

methods {


    // function MerkleTree.getValue(address, address) external returns(uint256) envfree;
    // function MerkleTree.getHash(bytes32) external returns(bytes32) envfree;
    // function MerkleTree.wellFormedPath(bytes32, bytes32[]) external envfree;

    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

    function _.payForRange(IPaymentCoordinator.RangePayment[]) external => DISPATCHER(true);
    function _.payAllForRange(IPaymentCoordinator.RangePayment[]) external => DISPATCHER(true);
    function _.submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) external => DISPATCHER(true);
    function _.processClaim(IPaymentCoordinator.PaymentMerkleClaim claim, address recipient) external => DISPATCHER(true);

    // function PaymentCoordinator._checkClaim(
    //     IPaymentCoordinator.PaymentMerkleC   laim calldata claim,
    //     IPaymentCoordinator.DistributionRoot memory root
    // ) internal => NONDET;
}

// Do not need
// paymentCoordinator.submitRoot() adds a merkle root to the contract with a 7 day activationDelay

// // Check how accept root changes the storage.
// rule acceptRootStorageChange(env e) {
//     bytes32 pendingRoot; bytes32 pendingIpfsHash;
//     pendingRoot, pendingIpfsHash, _ = pendingRoot();

//     acceptRoot(e);

//     assert root() == pendingRoot;
//     assert ipfsHash() == pendingIpfsHash;
// }

// // Check an account claimed amount is correctly updated.
// rule updatedClaimedAmount(address account, address reward, uint256 claimable, bytes32[] proof) {
//     claim(account, reward, claimable, proof);

//     assert claimable == claimed(account, reward);
// }

// // Check an account can only claim greater amounts each time.
// rule increasingClaimedAmounts(address account, address reward, uint256 claimable, bytes32[] proof) {
//     uint256 claimed = claimed(account, reward);

//     claim(account, reward, claimable, proof);

//     assert claimable > claimed;
// }


// 2. Given a processClaim passes, second processClaim reverts
// Rule status: FAILS
// Assumptions
// - checkClaim(claim) == true and does not revert
// - claim has non-empty tokenLeaves
rule claimWithduplicateTokenLeafs(env e, IPaymentCoordinator.PaymentMerkleClaim claim, address recipient) {
    require claim.tokenLeaves.length == 1;
    
    checkClaim(e, claim);
    // bool canWork = claimerFor(e, claim.earnerLeaf.earner) != 0 ? 
    //     e.msg.sender == claimerFor(e, claim.earnerLeaf.earner) :
    //     e.msg.sender == claim.earnerLeaf.earner;

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


// Check that transfer amount in paymentCoordinator.prcocessClaim() is equal to leaf.cumulativeEarnings - cumulativeClaimed
/// cumulativeEarnings - cumulativeClaimed == balanceAfter - balanceBefore
/// status: Fail, havoc from token safeTransfer in the calldata
rule transferredTokensFromClaim(env e, IPaymentCoordinator.PaymentMerkleClaim claim, address recipient) {
    require claim.tokenLeaves.length >= 1;
    require token == claim.tokenLeaves[0].token;

    address earner = claim.earnerLeaf.earner;
    uint256 balanceBefore = token.balanceOf(e, recipient);
    uint256 cumulativeClaimed = cumulativeClaimed(e, earner, claim.tokenLeaves[0].token);

    processClaim(e, claim, recipient);

    uint256 balanceAfter = token.balanceOf(e, recipient);

    assert claim.tokenLeaves[0].cumulativeEarnings - cumulativeClaimed == balanceAfter - balanceBefore;
}

/// paymentCoordinator.checkClaim == True => processClaim success 
/// status: pass
rule claimCorrectness(env e, IPaymentCoordinator.PaymentMerkleClaim claim, address recipient) {
    // Call checkClaim to ensure it does not revert in rule since it never returns false
    checkClaim(e, claim);

    bool canWork = claimerFor(e, claim.earnerLeaf.earner) != 0 ? 
        e.msg.sender == claimerFor(e, claim.earnerLeaf.earner) :
        e.msg.sender == claim.earnerLeaf.earner;

    processClaim@withrevert(e, claim, recipient);

    assert !lastReverted => canWork;
}

/// checkClaim either returns True or reverts otherwise
invariant checkClaimNeverFalse(env e, IPaymentCoordinator.PaymentMerkleClaim claim)
    !checkClaim(e, claim);
