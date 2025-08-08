// SPDX-License-Identifier: MIT
// Adapted from OpenZeppelin Contracts (last updated v4.8.0) (utils/cryptography/MerkleProof.sol)

pragma solidity ^0.8.0;

/**
 * @dev These functions deal with verification of Merkle Tree proofs.
 *
 * WARNING: You should avoid using leaf values that are 64 bytes long prior to
 * hashing, salt the leaves, or hash the leaves with a hash function other than
 * what is used for the Merkle tree's internal nodes. This is because the
 * concatenation of a sorted pair of internal nodes in the Merkle tree could
 * be reinterpreted as a leaf value.
 */
library Merkle {
    /// @notice Thrown when the provided proof was not a multiple of 32, or was empty for SHA256.
    /// @dev Error code: 0x4dc5f6a4
    error InvalidProofLength();

    /// @notice Thrown when the provided index was outside the max index for the tree.
    /// @dev Error code: 0x63df8171
    error InvalidIndex();

    /// @notice Thrown when the provided leaves' length was not a power of two.
    /// @dev Error code: 0xf6558f51
    error LeavesNotPowerOfTwo();

    /// @notice Thrown when the provided leaves' length was 0.
    /// @dev Error code: 0xbaec3d9a
    error NoLeaves();

    /// @notice Thrown when the provided leaves' length was insufficient.
    /// @dev Error code: 0xf8ef0367
    /// @dev This is used for the SHA256 Merkle tree, where the tree must have more than 1 leaf.
    error NotEnoughLeaves();

    /// @notice Thrown when the root is empty.
    /// @dev Error code: 0x53ce4ece
    /// @dev Empty roots should never be valid. We prevent them to avoid issues like the Nomad bridge attack: <https://medium.com/nomad-xyz-blog/nomad-bridge-hack-root-cause-analysis-875ad2e5aacd>
    error EmptyRoot();

    /**
     * @notice Verifies that a given leaf is included in a Merkle tree
     * @param proof The proof of inclusion for the leaf
     * @param root The root of the Merkle tree
     * @param leaf The leaf to verify
     * @param index The index of the leaf in the Merkle tree
     * @return True if the leaf is included in the Merkle tree, false otherwise
     * @dev A `proof` is valid if and only if the rebuilt hash matches the root of the tree.
     * @dev Reverts for:
     *      - InvalidProofLength: proof.length is not a multiple of 32.
     *      - InvalidIndex: index is not 0 at conclusion of computation (implying outside the max index for the tree).
     */
    function verifyInclusionKeccak(
        bytes memory proof,
        bytes32 root,
        bytes32 leaf,
        uint256 index
    ) internal pure returns (bool) {
        require(root != bytes32(0), EmptyRoot());
        return processInclusionProofKeccak(proof, leaf, index) == root;
    }

    /**
     * @notice Returns the rebuilt hash obtained by traversing a Merkle tree up
     * from `leaf` using `proof`.
     * @param proof The proof of inclusion for the leaf
     * @param leaf The leaf to verify
     * @param index The index of the leaf in the Merkle tree
     * @return The rebuilt hash
     * @dev Reverts for:
     *      - InvalidProofLength: proof.length is not a multiple of 32.
     *      - InvalidIndex: index is not 0 at conclusion of computation (implying outside the max index for the tree).
     * @dev The tree is built assuming `leaf` is the 0 indexed `index`'th leaf from the bottom left of the tree.
     */
    function processInclusionProofKeccak(
        bytes memory proof,
        bytes32 leaf,
        uint256 index
    ) internal pure returns (bytes32) {
        if (proof.length == 0) {
            return leaf;
        }

        require(proof.length % 32 == 0, InvalidProofLength());

        bytes32 computedHash = leaf;
        for (uint256 i = 32; i <= proof.length; i += 32) {
            if (index % 2 == 0) {
                // if index is even, then computedHash is a left sibling
                assembly {
                    mstore(0x00, computedHash)
                    mstore(0x20, mload(add(proof, i)))
                    computedHash := keccak256(0x00, 0x40)
                    index := div(index, 2)
                }
            } else {
                // if index is odd, then computedHash is a right sibling
                assembly {
                    mstore(0x00, mload(add(proof, i)))
                    mstore(0x20, computedHash)
                    computedHash := keccak256(0x00, 0x40)
                    index := div(index, 2)
                }
            }
        }

        // Confirm proof was fully consumed by end of computation
        require(index == 0, InvalidIndex());

        return computedHash;
    }

    /**
     * @notice Verifies that a given leaf is included in a Merkle tree
     * @param proof The proof of inclusion for the leaf
     * @param root The root of the Merkle tree
     * @param leaf The leaf to verify
     * @param index The index of the leaf in the Merkle tree
     * @return True if the leaf is included in the Merkle tree, false otherwise
     * @dev A `proof` is valid if and only if the rebuilt hash matches the root of the tree.
     * @dev Reverts for:
     *      - InvalidProofLength: proof.length is 0 or not a multiple of 32.
     *      - InvalidIndex: index is not 0 at conclusion of computation (implying outside the max index for the tree).
     */
    function verifyInclusionSha256(
        bytes memory proof,
        bytes32 root,
        bytes32 leaf,
        uint256 index
    ) internal view returns (bool) {
        require(root != bytes32(0), EmptyRoot());
        return processInclusionProofSha256(proof, leaf, index) == root;
    }

    /**
     * @notice Returns the rebuilt hash obtained by traversing a Merkle tree up
     * from `leaf` using `proof`.
     * @param proof The proof of inclusion for the leaf
     * @param leaf The leaf to verify
     * @param index The index of the leaf in the Merkle tree
     * @return The rebuilt hash
     * @dev Reverts for:
     *      - InvalidProofLength: proof.length is 0 or not a multiple of 32.
     *      - InvalidIndex: index is not 0 at conclusion of computation (implying outside the max index for the tree).
     * @dev The tree is built assuming `leaf` is the 0 indexed `index`'th leaf from the bottom left of the tree.
     */
    function processInclusionProofSha256(
        bytes memory proof,
        bytes32 leaf,
        uint256 index
    ) internal view returns (bytes32) {
        require(proof.length != 0 && proof.length % 32 == 0, InvalidProofLength());
        bytes32[1] memory computedHash = [leaf];
        for (uint256 i = 32; i <= proof.length; i += 32) {
            if (index % 2 == 0) {
                // if index is even, then computedHash is a left sibling
                assembly {
                    mstore(0x00, mload(computedHash))
                    mstore(0x20, mload(add(proof, i)))
                    if iszero(staticcall(sub(gas(), 2000), 2, 0x00, 0x40, computedHash, 0x20)) { revert(0, 0) }
                    index := div(index, 2)
                }
            } else {
                // if index is odd, then computedHash is a right sibling
                assembly {
                    mstore(0x00, mload(add(proof, i)))
                    mstore(0x20, mload(computedHash))
                    if iszero(staticcall(sub(gas(), 2000), 2, 0x00, 0x40, computedHash, 0x20)) { revert(0, 0) }
                    index := div(index, 2)
                }
            }
        }

        // Confirm proof was fully consumed by end of computation
        require(index == 0, InvalidIndex());

        return computedHash[0];
    }

    /**
     * @notice Returns the Merkle root of a tree created from a set of leaves using SHA-256 as its hash function
     * @param leaves the leaves of the Merkle tree
     * @return The computed Merkle root of the tree.
     * @dev Reverts for:
     *      - NotEnoughLeaves: leaves.length is less than 2.
     *      - LeavesNotPowerOfTwo: leaves.length is not a power of two.
     * @dev Unlike the Keccak version, this function does not allow a single-leaf tree.
     */
    function merkleizeSha256(
        bytes32[] memory leaves
    ) internal pure returns (bytes32) {
        require(leaves.length > 1, NotEnoughLeaves());
        require(isPowerOfTwo(leaves.length), LeavesNotPowerOfTwo());

        // There are half as many nodes in the layer above the leaves
        uint256 numNodesInLayer = leaves.length / 2;
        // Create a layer to store the internal nodes
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        // Fill the layer with the pairwise hashes of the leaves
        for (uint256 i = 0; i < numNodesInLayer; i++) {
            layer[i] = sha256(abi.encodePacked(leaves[2 * i], leaves[2 * i + 1]));
        }

        // While we haven't computed the root
        while (numNodesInLayer != 1) {
            // The next layer above has half as many nodes
            numNodesInLayer /= 2;
            // Overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint256 i = 0; i < numNodesInLayer; i++) {
                layer[i] = sha256(abi.encodePacked(layer[2 * i], layer[2 * i + 1]));
            }
        }
        // The first node in the layer is the root
        return layer[0];
    }

    /**
     * @notice Returns the Merkle root of a tree created from a set of leaves using Keccak as its hash function
     * @param leaves the leaves of the Merkle tree
     * @return The computed Merkle root of the tree.
     * @dev Reverts for:
     *      - NoLeaves: leaves.length is 0.
     */
    function merkleizeKeccak(
        bytes32[] memory leaves
    ) internal pure returns (bytes32) {
        require(leaves.length > 0, NoLeaves());

        uint256 numNodesInLayer;
        if (!isPowerOfTwo(leaves.length)) {
            // Pad to the next power of 2
            numNodesInLayer = 1;
            while (numNodesInLayer < leaves.length) {
                numNodesInLayer *= 2;
            }
        } else {
            numNodesInLayer = leaves.length;
        }

        // Create a layer to store the internal nodes
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        for (uint256 i = 0; i < leaves.length; i++) {
            layer[i] = leaves[i];
        }

        // While we haven't computed the root
        while (numNodesInLayer != 1) {
            // The next layer above has half as many nodes
            numNodesInLayer /= 2;
            // Overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint256 i = 0; i < numNodesInLayer; i++) {
                layer[i] = keccak256(abi.encodePacked(layer[2 * i], layer[2 * i + 1]));
            }
        }
        // The first node in the layer is the root
        return layer[0];
    }

    /**
     * @notice Returns the Merkle proof for a given index in a tree created from a set of leaves using Keccak as its hash function
     * @param leaves the leaves of the Merkle tree
     * @param index the index of the leaf to get the proof for
     * @return proof The computed Merkle proof for the leaf at index.
     * @dev Reverts for:
     *      - InvalidIndex: index is outside the max index for the tree.
     */
    function getProofKeccak(bytes32[] memory leaves, uint256 index) internal pure returns (bytes memory proof) {
        require(leaves.length > 0, NoLeaves());
        // TODO: very inefficient, use ZERO_HASHES
        // pad to the next power of 2
        uint256 numNodesInLayer = 1;
        while (numNodesInLayer < leaves.length) {
            numNodesInLayer *= 2;
        }
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        for (uint256 i = 0; i < leaves.length; i++) {
            layer[i] = leaves[i];
        }

        if (index >= layer.length) revert InvalidIndex();

        // While we haven't computed the root
        while (numNodesInLayer != 1) {
            // Flip the least significant bit of index to get the sibling index
            uint256 siblingIndex = index ^ 1;
            // Add the sibling to the proof
            proof = abi.encodePacked(proof, layer[siblingIndex]);
            index /= 2;

            // The next layer above has half as many nodes
            numNodesInLayer /= 2;
            // Overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint256 i = 0; i < numNodesInLayer; i++) {
                layer[i] = keccak256(abi.encodePacked(layer[2 * i], layer[2 * i + 1]));
            }
        }
    }

    /**
     * @notice Returns the Merkle proof for a given index in a tree created from a set of leaves using SHA-256 as its hash function
     * @param leaves the leaves of the Merkle tree
     * @param index the index of the leaf to get the proof for
     * @return proof The computed Merkle proof for the leaf at index.
     * @dev Reverts for:
     *      - NotEnoughLeaves: leaves.length is less than 2.
     * @dev Unlike the Keccak version, this function does not allow a single-leaf proof.
     */
    function getProofSha256(bytes32[] memory leaves, uint256 index) internal pure returns (bytes memory proof) {
        require(leaves.length > 1, NotEnoughLeaves());
        // TODO: very inefficient, use ZERO_HASHES
        // pad to the next power of 2
        uint256 numNodesInLayer = 1;
        while (numNodesInLayer < leaves.length) {
            numNodesInLayer *= 2;
        }
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        for (uint256 i = 0; i < leaves.length; i++) {
            layer[i] = leaves[i];
        }

        if (index >= layer.length) revert InvalidIndex();

        // While we haven't computed the root
        while (numNodesInLayer != 1) {
            // Flip the least significant bit of index to get the sibling index
            uint256 siblingIndex = index ^ 1;
            // Add the sibling to the proof
            proof = abi.encodePacked(proof, layer[siblingIndex]);
            index /= 2;

            // The next layer above has half as many nodes
            numNodesInLayer /= 2;
            // Overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint256 i = 0; i < numNodesInLayer; i++) {
                layer[i] = sha256(abi.encodePacked(layer[2 * i], layer[2 * i + 1]));
            }
        }
    }

    /**
     * @notice Returns whether the input is a power of two
     * @param value the value to check
     * @return True if the input is a power of two, false otherwise
     */
    function isPowerOfTwo(
        uint256 value
    ) internal pure returns (bool) {
        return value != 0 && (value & (value - 1)) == 0;
    }
}
