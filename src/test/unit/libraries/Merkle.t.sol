// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/test/utils/Murky.sol";

abstract contract MerkleBaseTest is Test, MurkyBase {
    bool usingSha; // Whether to use Keccak or Sha256 for tree + proof generation.
    bytes32[] leaves; // The contents of the merkle tree (unsorted).
    bytes32 root; // The root of the merkle tree.
    bytes[] proofs; // The proofs for each leaf in the tree.

    /// -----------------------------------------------------------------------
    /// Keccak + Sha256 Tests
    /// -----------------------------------------------------------------------

    /// @notice Verifies that Murky's proofs are compatible with our tree and proof verification.
    function test_verifyInclusion_ValidProof() public {
        assertValidProofs();
    }

    /// @notice Verifies that an empty proof(s) is invalid.
    function test_verifyInclusion_EmptyProofs() public {
        proofs = new bytes[](proofs.length);
        assertInvalidProofs();
    }

    /// -----------------------------------------------------------------------
    /// Assertions
    /// -----------------------------------------------------------------------

    /// @dev Checks that all proofs are valid for their respective leaves.
    function assertValidProofs() internal virtual {
        function (bytes memory proof, bytes32 root, bytes32 leaf, uint256 index) returns (bool) verifyInclusion =
            usingSha ? Merkle.verifyInclusionSha256 : Merkle.verifyInclusionKeccak;
        for (uint i = 0; i < leaves.length; ++i) {
            assertTrue(verifyInclusion(proofs[i], root, leaves[i], i), "invalid proof");
        }
    }

    /// @dev Checks that all proofs are invalid for their respective leaves.
    function assertInvalidProofs() internal virtual {
        function (bytes memory proof, bytes32 root, bytes32 leaf, uint256 index) returns (bool) verifyInclusion =
            usingSha ? Merkle.verifyInclusionSha256 : Merkle.verifyInclusionKeccak;
        for (uint i = 0; i < leaves.length; ++i) {
            assertFalse(verifyInclusion(proofs[i], root, leaves[i], i), "valid proof");
        }
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /// @dev Effeciently generates a random list of leaves without iterative hashing.
    function getLeaves(uint numLeaves) internal view virtual returns (bytes32[] memory leaves) {
        bytes memory _leavesAsBytes = vm.randomBytes(numLeaves * 32);
        /// @solidity memory-safe-assembly
        assembly {
            leaves := _leavesAsBytes // Typecast bytes -> bytes32[].
            mstore(leaves, numLeaves) // Update length n*32 -> n.
        }
    }

    /// @dev Generates proofs for each leaf in the tree.
    ///         Intended to be overridden by the below child contracts.
    function getProofs(bytes32[] memory leaves) public view virtual returns (bytes[] memory proofs);
}

contract MerkleKeccakTest is MerkleBaseTest, MerkleKeccak {
    function setUp() public {
        usingSha = false;
        leaves = getLeaves(vm.randomBool() ? 9 : 10);
        root = Merkle.merkleizeKeccak(leaves);
        proofs = getProofs(leaves);
    }

    function nextPowerOf2(uint v) internal pure returns (uint) {
        unchecked {
            // Round up to the next power of 2 using the method described here:
            // https://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
            if (v == 0) return 0;
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

    function getProofs(bytes32[] memory leaves) public view virtual override returns (bytes[] memory proofs) {
        // Merkle.merkleizeKeccak pads to next power of 2, so we need to match that.
        uint numLeaves = nextPowerOf2(leaves.length);
        bytes32[] memory paddedLeaves = new bytes32[](numLeaves);
        for (uint i = 0; i < leaves.length; ++i) {
            // TODO: Point leaves to paddedLeaves using assembly to avoid loop.
            paddedLeaves[i] = leaves[i];
        }

        proofs = new bytes[](leaves.length);
        for (uint i = 0; i < leaves.length; ++i) {
            proofs[i] = abi.encodePacked(getProof(paddedLeaves, i));
        }
    }
}
