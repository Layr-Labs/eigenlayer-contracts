## Merkle

| File | Notes |
| -------- | -------- |
| [`Merkle.sol`](../../../src/contracts/libraries/Merkle.sol) | Core Merkle tree library |

## Overview

The `Merkle` library provides cryptographically secure Merkle tree functionality for the EigenLayer protocol. It supports both Keccak256 and SHA-256 hash functions for different use cases across the system. The library enables efficient verification of data inclusion in large datasets without requiring the full dataset, which is essential for scalable proof systems in EigenLayer.

Key capabilities include:
- **Proof verification**: Verify that a leaf exists in a Merkle tree given a root and proof
- **Tree construction**: Build Merkle trees from arrays of leaves
- **Proof generation**: Generate inclusion proofs for specific leaves
- **Dual hash function support**: Both Keccak256 (for EVM compatibility) and SHA-256 (for beacon chain compatibility)

## Prior Reading

* Understanding of [Merkle trees](https://en.wikipedia.org/wiki/Merkle_tree) and cryptographic hash functions
* [EIP-197](https://eips.ethereum.org/EIPS/eip-197) for understanding precompiled contracts

## Usage in EigenLayer

The Merkle library is used extensively throughout EigenLayer for various proof systems:

- **[`BN254CertificateVerifier`](../../../src/contracts/multichain/BN254CertificateVerifier.sol)**: Verifies operator information inclusion in certificate Merkle trees (see [CertificateVerifier.md](../../multichain/destination/CertificateVerifier.md))
- **[`OperatorTableUpdater`](../../../src/contracts/multichain/OperatorTableUpdater.sol)**: Manages operator set proofs for multichain operations (see [OperatorTableUpdater.md](../../multichain/destination/OperatorTableUpdater.md))
- **[`RewardsCoordinator`](../../../src/contracts/core/RewardsCoordinator.sol)**: Verifies reward distribution claims (see [RewardsCoordinator.md](../RewardsCoordinator.md))
- **[`BeaconChainProofs`](../../../src/contracts/libraries/BeaconChainProofs.sol)**: Processes beacon chain state proofs (see [EigenPod.md](../EigenPod.md) for beacon chain proof usage)

## Security Considerations

### **Critical Security Warning**

**You should avoid using leaf values that are 64 bytes long prior to hashing, salt the leaves, or hash the leaves with a hash function other than what is used for the Merkle tree's internal nodes.** This prevents potential collision attacks where the concatenation of a sorted pair of internal nodes could be reinterpreted as a leaf value.

### **Zero Hash Padding**

When trees are not perfect powers of 2, the library pads with `bytes32(0)` values. For security-critical applications, consider using unique filler values to prevent potential collision attacks with legitimate zero-valued leaves.

---

## Proof Verification

### **Keccak256 Proof Verification**

#### `verifyInclusionKeccak`

```solidity
function verifyInclusionKeccak(
    bytes memory proof,
    bytes32 root,
    bytes32 leaf,
    uint256 index
) internal pure returns (bool)
```

Verifies that a given leaf is included in a Merkle tree using Keccak256.

*Effects:*
* Computes the root hash by traversing the Merkle proof path
* Compares computed root with expected root for verification

*Used in:*
* [`BN254CertificateVerifier.verifyOperatorInfoProof`](../../../src/contracts/multichain/BN254CertificateVerifier.sol)
* [`OperatorTableUpdater.checkGlobalTableHash`](../../../src/contracts/multichain/OperatorTableUpdater.sol)
* [`RewardsCoordinator`](../../../src/contracts/core/RewardsCoordinator.sol) for reward claim verification

#### `processInclusionProofKeccak`

```solidity
function processInclusionProofKeccak(
    bytes memory proof,
    bytes32 leaf,
    uint256 index
) internal pure returns (bytes32)
```

Returns the computed root hash by traversing up the tree from the leaf.

*Effects:*
* Traverses the Merkle tree from leaf to root using provided proof
* Computes hash at each level by combining current hash with proof siblings
* Returns the final computed root hash

*Requirements:*
* Proof length MUST be a multiple of 32 bytes (reverts with `InvalidProofLength` otherwise)
* Index MUST reach 0 after processing all proof elements (reverts with `InvalidIndex` if proof length mismatch)

### **SHA-256 Proof Verification**

#### `verifyInclusionSha256`

```solidity
function verifyInclusionSha256(
    bytes memory proof,
    bytes32 root,
    bytes32 leaf,
    uint256 index
) internal pure returns (bool)
```

Verifies inclusion using SHA-256 hash function via precompiled contract.

*Effects:*
* Computes the root hash using SHA-256 via precompiled contract
* Compares computed root with expected root for verification

*Used in:*
* [`BeaconChainProofs`](../../../src/contracts/libraries/BeaconChainProofs.sol) for beacon chain state verification

#### `processInclusionProofSha256`

```solidity
function processInclusionProofSha256(
    bytes memory proof,
    bytes32 leaf,
    uint256 index
) internal view returns (bytes32)
```

Returns the computed root hash by traversing up the tree from the leaf using SHA-256.

*Effects:*
* Traverses the Merkle tree from leaf to root using provided proof
* Computes hash at each level combining current hash with proof siblings
* Returns the final computed root hash

*Requirements:*
* Proof length MUST be non-zero and a multiple of 32 bytes (reverts with `InvalidProofLength` otherwise)
* Index MUST reach 0 after processing all proof elements (reverts with `InvalidIndex` if proof length mismatch)

---

## Tree Construction

### **Keccak256 Tree Construction**

#### `merkleizeKeccak`

```solidity
function merkleizeKeccak(bytes32[] memory leaves) internal pure returns (bytes32)
```

Constructs a Merkle tree root from an array of leaves using Keccak256. Accepts any non-empty array of leaves, including single-leaf trees, and automatically pads to the next power of 2 using `bytes32(0)` values.

*Effects:*
* Pads input array to next power of 2 (if needed) using `bytes32(0)` values
* Constructs binary Merkle tree bottom-up by hashing pairs using in-place array modification
* For single-leaf arrays, returns the leaf itself as the root
* Returns the single root hash representing the entire tree

*Algorithm:*
1. Pad leaves array to next power of 2
2. Iteratively hash pairs level by level until single root remains
3. Uses in-place array modification for gas efficiency

*Used in:*
* [`BN254CertificateVerifier`](../../../src/contracts/multichain/BN254CertificateVerifier.sol) for operator info trees
* [`OperatorTableUpdater`](../../../src/contracts/multichain/OperatorTableUpdater.sol) for operator set hashing

### **SHA-256 Tree Construction**

#### `merkleizeSha256`

```solidity
function merkleizeSha256(bytes32[] memory leaves) internal pure returns (bytes32)
```

Constructs a Merkle tree root using SHA-256 hash function.

*Effects:*
* Validates input meets strict requirements (power of 2, minimum 2 leaves)
* Constructs binary Merkle tree bottom-up using SHA-256 hashing
* Returns the single root hash representing the entire tree

*Requirements*:
* Input array MUST contain at least 2 leaves (rejects single-leaf trees with `NotEnoughLeaves` error)
* Input array length MUST be an exact power of 2 (validates with `LeavesNotPowerOfTwo` error)
* No auto-padding available - stricter requirements for beacon chain compatibility

*Used in:*
* [`BeaconChainProofs`](../../../src/contracts/libraries/BeaconChainProofs.sol) for beacon chain compatibility

---

## Proof Generation

### **Keccak256 Proof Generation**

#### `getProofKeccak`

```solidity
function getProofKeccak(bytes32[] memory leaves, uint256 index) internal pure returns (bytes memory proof)
```

Generates an inclusion proof for a specific leaf in a Keccak256 tree. Supports single-leaf trees (returns empty proof) and automatically handles non-power-of-2 leaf arrays through padding.

*Effects:*
* Constructs a Merkle tree from the provided leaves with automatic padding
* Traverses from specified leaf to root, collecting sibling hashes at each level
* For single-leaf trees, returns empty proof since root equals leaf
* Returns concatenated proof bytes for verification

*Algorithm:*
1. Pad leaves to next power of 2
2. For each tree level, find sibling of current index
3. Append sibling to proof bytes
4. Move up tree by dividing index by 2
5. Continue until reaching root

*Used in:*
* Test frameworks for generating proofs
* Off-chain proof generation systems

### **SHA-256 Proof Generation**

#### `getProofSha256`

```solidity
function getProofSha256(bytes32[] memory leaves, uint256 index) internal pure returns (bytes memory proof)
```

Generates SHA-256 inclusion proof with same algorithm as Keccak version but using SHA-256 hashing.

*Effects:*
* Validates input meets SHA-256 requirements (minimum 2 leaves)
* Constructs a Merkle tree using SHA-256 hashing
* Traverses from specified leaf to root, collecting sibling hashes
* Returns concatenated proof bytes for verification

*Requirements*:
* Input array MUST contain at least 2 leaves (rejects single-leaf trees with `NotEnoughLeaves` error)
* Cannot generate proofs for single-element arrays
* Follows stricter validation for beacon chain compatibility

---

## Utility Functions

### **Power of Two Check**

#### `isPowerOfTwo`

```solidity
function isPowerOfTwo(uint256 value) internal pure returns (bool)
```

Efficiently determines if a value is a power of 2 using bit manipulation.

*Effects:*
* Performs bitwise operations to check power-of-2 property

*Used internally* for validation in `merkleizeSha256` and optimization paths.

---

## Error Reference

| Error | Code | Description |
|-------|------|-------------|
| `InvalidProofLength` | `0x4dc5f6a4` | Proof length not multiple of 32 bytes |
| `InvalidIndex` | `0x63df8171` | Index outside valid range for tree |
| `LeavesNotPowerOfTwo` | `0xf6558f51` | Leaves array not power of 2 (SHA-256 only) |
| `NoLeaves` | `0xbaec3d9a` | Empty leaves array provided |
| `NotEnoughLeaves` | `0xf8ef0367` | Less than 2 leaves for SHA-256 operations |

---

## Implementation Optimizations

### **Assembly Usage**
- Proof verification uses inline assembly for gas efficiency
- Manual memory management avoids Solidity's safety overhead
- Direct opcode usage (Keccak256) vs precompile calls (SHA-256)

### **Memory Efficiency**
- Tree construction reuses input array space
- In-place modifications reduce memory allocation costs
- Sibling calculation uses XOR instead of arithmetic operations

### **Precompile Handling**
- SHA-256 functions reserve gas before precompile calls
- Static call pattern optimized for success case
- Explicit revert handling for precompile failures

---

## Implementation Notes

### **Keccak256 vs SHA-256 Differences**

The library provides two distinct implementations with different design philosophies:

| Feature | Keccak256 Functions | SHA-256 Functions |
|---------|-------------------|-------------------|
| **Single-leaf trees** | ✅ Supported | ❌ Rejected (`NotEnoughLeaves`) |
| **Input validation** | Flexible (any non-empty array) | Strict (≥2 leaves, power of 2) |
| **Auto-padding** | ✅ Pads to next power of 2 | ❌ Requires exact power of 2 |
| **Use case** | General EVM Merkle trees | Beacon chain compatibility |
| **Error handling** | Permissive | Strict validation |

### **Tree Padding Strategy**
The library uses different padding strategies for different hash functions:

- **Keccak256**: Pads with `bytes32(0)` to next power of 2
- **SHA-256**: Requires exact power of 2, no padding

### **Index Validation**
Index validation occurs during proof processing rather than upfront, allowing the tree traversal algorithm to naturally detect out-of-bounds conditions.

### **Memory Layout**
Proof bytes are laid out as concatenated 32-byte chunks: `[sibling₀, sibling₁, ..., siblingₙ]` where siblings are sequenced by tree depth, with `sibling₀` being the sibling at the leaf level, `sibling₁` at the next level up, and so on until reaching the root.
