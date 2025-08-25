// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/test/utils/Murky.sol";

abstract contract MerkleBaseTest is Test, MurkyBase {
    uint internal constant MIN_LEAVES = 2; // Minimum number of leaves.
    uint internal constant MAX_LEAVES = 65; // Maximum number of leaves.

    bytes32[] leaves; // The contents of the merkle tree (unsorted).
    bytes32 root; // The root of the merkle tree.
    bytes[] proofs; // The proofs for each leaf in the tree.

    /// @dev Takes in `seed` which should be provided as a fuzz input.
    /// Ensures that RNG is deterministic, and tests are easily reproducible.
    modifier rng(uint seed) {
        _rng(seed);
        _;
    }

    function _rng(uint seed) internal {
        vm.setSeed(seed);
        leaves = _genLeaves(vm.randomUint(MIN_LEAVES, MAX_LEAVES));
        proofs = _genProofs(leaves);
        root = _genRoot(leaves);
    }

    /// -----------------------------------------------------------------------
    /// Keccak + Sha256 Tests
    /// -----------------------------------------------------------------------

    /// @notice Verifies that (Murky's) proofs are compatible with our implementation.
    function testFuzz_verifyInclusion_ValidProof(uint seed) public rng(seed) {
        _checkAllProofs(true);
    }

    /// @notice Verifies that an empty proof(s) is invalid.
    function testFuzz_verifyInclusion_EmptyProofs(uint seed) public rng(seed) {
        if (!usingSha()) vm.skip(true); // TODO: Breaking change, add in future.
        proofs = new bytes[](proofs.length);
        _checkAllProofs(false);
    }

    /// @notice Verifies that using wrong root fails verification
    function testFuzz_verifyInclusion_WrongRoot(uint seed) public rng(seed) {
        root = bytes32(vm.randomUint());
        _checkAllProofs(false);
    }

    /// @notice Verifies valid proofs cannot be used to prove invalid leaves.
    function testFuzz_verifyInclusion_WrongProofs(uint seed) public rng(seed) {
        bytes memory proof0 = proofs[0];
        bytes memory proof1 = proofs[1];
        (proofs[0], proofs[1]) = (proof1, proof0);
        _checkSingleProof(false, 0);
        _checkSingleProof(false, 1);
    }

    /// @notice Verifies that a valid proof with excess data appended is invalid.
    function testFuzz_verifyInclusion_ExcessProofLength(uint seed) public rng(seed) {
        unchecked {
            proofs[0] = abi.encodePacked(proofs[0], vm.randomBytes(vm.randomUint(1, 10) * vm.randomUint(31, 32)));
        }
        _checkSingleProof(false, 0);
    }

    /// @notice Verifies that a valid proof with a truncated length is invalid.
    function testFuzz_verifyInclusion_TruncatedProofLength(uint seed) public rng(seed) {
        if (!usingSha()) vm.skip(true); // TODO: Breaking change, add in future.
        bytes memory proof = proofs[0];
        console.log("proof length %d", proof.length); // 32
        /// @solidity memory-safe-assembly
        assembly {
            mstore(proof, sub(mload(proof), 32))
        }
        console.log("proof length %d", proof.length); // 0
        proofs[0] = proof;
        _checkSingleProof(false, 0); // Should revert, but doesn't...
    }

    /// @notice Verifies that a valid proof with a manipulated word is invalid.
    function testFuzz_verifyInclusion_ManipulatedProof(uint seed) public rng(seed) {
        bytes memory proof = proofs[0];
        /// @solidity memory-safe-assembly
        assembly {
            let m := add(proof, 0x20)
            let manipulated := shr(8, mload(m)) // Shift the first word to the right by 8 bits.
            mstore(m, manipulated)
        }
        proofs[0] = proof;
        _checkSingleProof(false, 0);
    }

    /// @notice Verifies that an out-of-bounds index reverts.
    function testFuzz_verifyInclusion_IndexOutOfBounds(uint seed) public rng(seed) {
        uint index = vm.randomUint(leaves.length, type(uint).max);
        vm.expectRevert(stdError.indexOOBError);
        _checkSingleProof(false, index);
    }

    /// @notice Verifies that an internal node cannot be used as a proof.
    function testFuzz_verifyInclusion_InternalNodeAsProof(uint seed) public rng(seed) {
        // Generate a tree with at least 4 leaves to ensure internal nodes exist
        leaves = _genLeaves(vm.randomUint(4, 8));
        proofs = _genProofs(leaves);
        root = _genRoot(leaves);
        function (bytes memory proof, bytes32 root, bytes32 leaf, uint256 index) view returns (bool) verifyInclusion =
            usingSha() ? Merkle.verifyInclusionSha256 : Merkle.verifyInclusionKeccak;
        assertFalse(verifyInclusion(proofs[2], root, hashLeafPairs(leaves[0], leaves[1]), 2));
    }

    /// @notice Verifies behavior with duplicate leaves in the tree
    function testFuzz_verifyInclusion_DuplicateLeaves(uint seed) public rng(seed) {
        leaves = _genLeaves(vm.randomUint(4, 8));
        leaves[0] = leaves[vm.randomUint(1, leaves.length - 1)];
        proofs = _genProofs(leaves);
        root = _genRoot(leaves);
        _checkAllProofs(true);
    }

    /// -----------------------------------------------------------------------
    /// Assertions
    /// -----------------------------------------------------------------------

    /// @dev Checks that all proofs are valid for their respective leaves.
    function _checkAllProofs(bool status) internal virtual {
        function (bytes memory proof, bytes32 root, bytes32 leaf, uint256 index) returns (bool) verifyInclusion =
            usingSha() ? Merkle.verifyInclusionSha256 : Merkle.verifyInclusionKeccak;
        for (uint i = 0; i < leaves.length; ++i) {
            if (proofs[i].length == 0 || proofs[i].length % 32 != 0) vm.expectRevert(Merkle.InvalidProofLength.selector);
            assertEq(verifyInclusion(proofs[i], root, leaves[i], i), status);
        }
    }

    /// @dev Checks that a single proof is valid for its respective leaf.
    function _checkSingleProof(bool status, uint index) internal virtual {
        function (bytes memory proof, bytes32 root, bytes32 leaf, uint256 index) view returns (bool) verifyInclusion =
            usingSha() ? Merkle.verifyInclusionSha256 : Merkle.verifyInclusionKeccak;
        if (proofs[index].length == 0 || proofs[index].length % 32 != 0) vm.expectRevert(Merkle.InvalidProofLength.selector);
        assertEq(verifyInclusion(proofs[index], root, leaves[index], index), status);
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /// @dev Efficiently pads the length of leaves to the next power of 2 by appending zeros.
    function _padLeaves(bytes32[] memory leaves) internal view virtual returns (bytes32[] memory paddedLeaves) {
        uint numLeaves = _roundUpPow2(leaves.length);
        paddedLeaves = new bytes32[](numLeaves);
        for (uint i = 0; i < leaves.length; ++i) {
            paddedLeaves[i] = leaves[i];
        }
    }

    /// @dev Generates a random list of leaves without iterative hashing.
    function _genLeaves(uint numLeaves) internal view virtual returns (bytes32[] memory leaves) {
        bytes memory _leavesAsBytes = vm.randomBytes(numLeaves * 32);
        /// @solidity memory-safe-assembly
        assembly {
            leaves := _leavesAsBytes // Typecast bytes -> bytes32[].
            mstore(leaves, numLeaves) // Update length n*32 -> n.
        }
    }

    /// @dev Generates proofs for each leaf in the tree.
    function _genProofs(bytes32[] memory leaves) internal view virtual returns (bytes[] memory proofs) {
        uint numLeaves = _roundUpPow2(leaves.length);
        bytes32[] memory paddedLeaves = _padLeaves(leaves);
        proofs = new bytes[](leaves.length);
        for (uint i = 0; i < leaves.length; ++i) {
            proofs[i] = abi.encodePacked(getProof(paddedLeaves, i));
        }
    }

    /// @dev Computes the merkle root using the appropriate hash function
    function _genRoot(bytes32[] memory leaves) internal view virtual returns (bytes32) {
        function (bytes32[] memory leaves) view returns (bytes32) merkleize = usingSha() ? Merkle.merkleizeSha256 : Merkle.merkleizeKeccak;
        if (usingSha()) leaves = _padLeaves(leaves);
        return merkleize(leaves);
    }

    /// @dev Rounds up to the next power of 2.
    ///      https://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
    function _roundUpPow2(uint v) internal pure returns (uint) {
        unchecked {
            v -= 1;
            v |= v >> 1;
            v |= v >> 2;
            v |= v >> 4;
            v |= v >> 8;
            v |= v >> 16;
            v |= v >> 32;
            v |= v >> 64;
            v |= v >> 128;
            return v + 1;
        }
    }

    function usingSha() internal view virtual returns (bool);
}

contract MerkleKeccakTest is MerkleBaseTest, MerkleKeccak {
    function usingSha() internal view virtual override returns (bool) {
        return false;
    }
}

contract MerkleShaTest is MerkleBaseTest, MerkleSha {
    function usingSha() internal view virtual override returns (bool) {
        return true;
    }
}
