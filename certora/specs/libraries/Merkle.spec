methods
{
    function merkleizeSha256(bytes32[]) external returns bytes32 envfree;
    function processInclusionProofSha256(bytes, bytes32, uint256) external returns (bytes32) envfree;
    function processInclusionProofKeccak(bytes, bytes32, uint256) external returns (bytes32) envfree;
    function equals(bytes32[], bytes32[]) external returns (bool) envfree;
}


/// @title The hashing doesn't colide.
rule merkleizeSha256IsInjective()
{
	bytes32[] leaves1; bytes32[] leaves2;
    bytes32 res1 = merkleizeSha256(leaves1);
    bytes32 res2 = merkleizeSha256(leaves2);

    //satisfy true;
    require isPowerOfTwo(leaves1.length) && isPowerOfTwo(leaves2.length);
        // && leaves1.length >= 32 && leaves2.length >= 32;
    
    satisfy res1 == res2;
    assert res1 == res2 => equals(leaves1, leaves2);
}

/// @title The hashing doesn't colide on inputs with the same lengths 
rule merkleizeSha256IsInjective_onSameLengths()
{
	bytes32[] leaves1; bytes32[] leaves2;
    bytes32 res1 = merkleizeSha256(leaves1);
    bytes32 res2 = merkleizeSha256(leaves2);

    require leaves1.length == leaves2.length;
    require isPowerOfTwo(leaves1.length) && isPowerOfTwo(leaves2.length);
        // && leaves1.length >= 32 && leaves2.length >= 32;
    
    satisfy res1 == res2;
    assert res1 == res2 => equals(leaves1, leaves2);
}

/// @title If only the leaf changes, the result must also change for processInclusionProofKeccak
//Proof1.length == Proof2.length && index1 == index2 && leaf1 != leaf2) => 
//  processInclusionProofKeccak(proof1, leaf1, index1) != processInclusionProofKeccak(proof2, leaf2, index2)           
rule processInclusionProofKeccak_correctness()
{
	bytes proof1; bytes32 leaf1; uint256 index1;
    bytes proof2; bytes32 leaf2; uint256 index2;
    bytes32 res1 = processInclusionProofKeccak(proof1, leaf1, index1);
    bytes32 res2 = processInclusionProofKeccak(proof2, leaf2, index2);

    satisfy res1 != res2;
    assert proof1.length == proof2.length && index1 == index2 && leaf1 != leaf2 => res1 != res2;
}

function isPowerOfTwo(uint256 x) returns bool
{
    return x == 0 || x == 1 || x == 2 ||
        x == 4 || x == 8 || x == 16 || x == 32 ||
        x == 64 || x == 128 || x == 256 || x == 512 ||
        x == 1024 || x == 2048 || x == 4096 || x == 8192;
}


///////////////////   IN DEVELOPMENT / OBSOLETE    ////////

/// @title If only the leaf changes, the result must also change for processInclusionProofSha256
// Proof1.length == Proof2.length && index1 == index2 && leaf1 != leaf2) => 
// processInclusionProofSha256(proof1, leaf1, index1) != processInclusionProofSha256(proof2, leaf2, index2)           
// TODO
rule processInclusionProofSha256_correctness()
{
	bytes proof1; bytes32 leaf1; uint256 index1;
    bytes proof2; bytes32 leaf2; uint256 index2;
    bytes32 res1 = processInclusionProofSha256(proof1, leaf1, index1);
    bytes32 res2 = processInclusionProofSha256(proof2, leaf2, index2);

    satisfy true;
    satisfy res1 != res2;
    assert proof1.length == proof2.length && index1 == index2 && leaf1 != leaf2 => res1 != res2;
}

// to check the conditions under which the method works correctly
// loop_iter, hashing_length_bound, optimistic_loop, optimistic_hashing
rule processInclusionProofSha256_SingleValue()
{
	bytes proof1; bytes32 leaf1; uint256 index1;
    bytes proof2; bytes32 leaf2; uint256 index2;
    bytes32 res1 = processInclusionProofSha256(proof1, leaf1, index1);
    bytes32 res2 = processInclusionProofSha256(proof2, leaf2, index2);

    satisfy true;
    assert res1 == res2;
}
