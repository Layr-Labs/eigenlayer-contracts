// SPDX-License-Identifier: GPL-2.0-or-later

using MerkleTree as MerkleTree;
using Util as Util;

methods {
    function root() external returns bytes32 envfree;
    function ipfsHash() external returns bytes32 envfree;
    function claimed(address, address) external returns(uint256) envfree;
    function claim(address, address, uint256, bytes32[]) external returns(uint256) envfree;
    function pendingRoot() external returns(bytes32, bytes32, uint256) envfree;

    function MerkleTree.getValue(address, address) external returns(uint256) envfree;
    function MerkleTree.getHash(bytes32) external returns(bytes32) envfree;
    function MerkleTree.wellFormedPath(bytes32, bytes32[]) external envfree;

    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

    function Util.balanceOf(address, address) external returns(uint256) envfree;
}

// Don't need
// paymentCoordinator.submitRoot() adds a merkle root to the contract with a 7 day activationDelay

// Check how accept root changes the storage.
rule acceptRootStorageChange(env e) {
    bytes32 pendingRoot; bytes32 pendingIpfsHash;
    pendingRoot, pendingIpfsHash, _ = pendingRoot();

    acceptRoot(e);

    assert root() == pendingRoot;
    assert ipfsHash() == pendingIpfsHash;
}

// Check an account claimed amount is correctly updated.
rule updatedClaimedAmount(address account, address reward, uint256 claimable, bytes32[] proof) {
    claim(account, reward, claimable, proof);

    assert claimable == claimed(account, reward);
}

// Check an account can only claim greater amounts each time.
rule increasingClaimedAmounts(address account, address reward, uint256 claimable, bytes32[] proof) {
    uint256 claimed = claimed(account, reward);

    claim(account, reward, claimable, proof);

    assert claimable > claimed;
}

// 1. For any claim that checkClaim(claim) == true && contains duplicate tokenLeafs, processClaim() will revert


// 2. Given a claim, processClaim passes, second processClaim reverts


// 3. Keep rule
// Check that transfer amount in paymentCoordinator.prcocessClaim() is equal to leaf.cumulativeEarnings - cumulativeClaimed

// create invariant with hooks keeping a sum of earner claims everytime processClaim() is called, this should also equal cumulativeClaimed storage mapping
// sum of earner claims = sum of earner
// Check that the transferred amount is equal to the claimed amount minus the previous claimed amount.
rule transferredTokens(address account, address reward, uint256 claimable, bytes32[] proof) {
    // Assume that the rewards distributor itself is not receiving the tokens, to simplify this rule.
    require account != currentContract;

    uint256 balanceBefore = Util.balanceOf(reward, account);
    uint256 claimedBefore = claimed(account, reward);

    // Safe require because the sum is capped by the total supply.
    require balanceBefore + Util.balanceOf(reward, currentContract) < 2^256;

    claim(account, reward, claimable, proof);

    uint256 balanceAfter = Util.balanceOf(reward, account);

    assert balanceAfter - balanceBefore == claimable - claimedBefore;
}


// 4. Keep rule
// paymentCoordinator.checkClaim == True => processClaim success 

// The main correctness result of the verification.
// It ensures that if the root is setup according to a well-formed Merkle tree, then claiming will result in receiving the rewards stored in the tree for that particular pair of account and reward.
rule claimCorrectness(address account, address reward, uint256 claimable, bytes32[] proof) {
    bytes32 node;

    // Assume that root is the hash of node in the tree.
    require MerkleTree.getHash(node) == root();

    // No need to make sure that node is equal to currRoot: one can pass an internal node instead.

    // Assume that the tree is well-formed.
    MerkleTree.wellFormedPath(node, proof);

    claim(account, reward, claimable, proof);

    assert claimable == MerkleTree.getValue(account, reward);
}