# Solidity API

## Merkle

_These functions deal with verification of Merkle Tree proofs.

The tree and the proofs can be generated using our
https://github.com/OpenZeppelin/merkle-tree[JavaScript library].
You will find a quickstart guide in the readme.

WARNING: You should avoid using leaf values that are 64 bytes long prior to
hashing, or use a hash function other than keccak256 for hashing leaves.
This is because the concatenation of a sorted pair of internal nodes in
the merkle tree could be reinterpreted as a leaf value.
OpenZeppelin's JavaScript library generates merkle trees that are safe
against this attack out of the box._

### verifyInclusionKeccak

```solidity
function verifyInclusionKeccak(bytes proof, bytes32 root, bytes32 leaf, uint256 index) internal pure returns (bool)
```

_Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is 
the 0 indexed `index`'th leaf from the bottom left of the tree.

Note this is for a Merkle tree using the keccak/sha3 hash function_

### processInclusionProofKeccak

```solidity
function processInclusionProofKeccak(bytes proof, bytes32 leaf, uint256 index) internal pure returns (bytes32)
```

_Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is 
the 0 indexed `index`'th leaf from the bottom left of the tree.

_Available since v4.4._

Note this is for a Merkle tree using the keccak/sha3 hash function_

### verifyInclusionSha256

```solidity
function verifyInclusionSha256(bytes proof, bytes32 root, bytes32 leaf, uint256 index) internal view returns (bool)
```

_Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is 
the 0 indexed `index`'th leaf from the bottom left of the tree.

Note this is for a Merkle tree using the sha256 hash function_

### processInclusionProofSha256

```solidity
function processInclusionProofSha256(bytes proof, bytes32 leaf, uint256 index) internal view returns (bytes32)
```

_Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is 
the 0 indexed `index`'th leaf from the bottom left of the tree.

_Available since v4.4._

Note this is for a Merkle tree using the sha256 hash function_

### merkleizeSha256

```solidity
function merkleizeSha256(bytes32[] leaves) internal pure returns (bytes32)
```

this function returns the merkle root of a tree created from a set of leaves using sha256 as its hash function
     @param leaves the leaves of the merkle tree

     @notice requires the leaves.length is a power of 2

