# Merkle
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/libraries/Merkle.sol)

*These functions deal with verification of Merkle Tree proofs.
The tree and the proofs can be generated using our
https://github.com/OpenZeppelin/merkle-tree[JavaScript library].
You will find a quickstart guide in the readme.
WARNING: You should avoid using leaf values that are 64 bytes long prior to
hashing, or use a hash function other than keccak256 for hashing leaves.
This is because the concatenation of a sorted pair of internal nodes in
the merkle tree could be reinterpreted as a leaf value.
OpenZeppelin's JavaScript library generates merkle trees that are safe
against this attack out of the box.*


## Functions
### verifyInclusionKeccak

*Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is
the 0 indexed `index`'th leaf from the bottom left of the tree.
Note this is for a Merkle tree using the keccak/sha3 hash function*


```solidity
function verifyInclusionKeccak(bytes memory proof, bytes32 root, bytes32 leaf, uint256 index)
    internal
    pure
    returns (bool);
```

### processInclusionProofKeccak

*Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is
the 0 indexed `index`'th leaf from the bottom left of the tree.
_Available since v4.4._
Note this is for a Merkle tree using the keccak/sha3 hash function*


```solidity
function processInclusionProofKeccak(bytes memory proof, bytes32 leaf, uint256 index) internal pure returns (bytes32);
```

### verifyInclusionSha256

*Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is
the 0 indexed `index`'th leaf from the bottom left of the tree.
Note this is for a Merkle tree using the sha256 hash function*


```solidity
function verifyInclusionSha256(bytes memory proof, bytes32 root, bytes32 leaf, uint256 index)
    internal
    view
    returns (bool);
```

### processInclusionProofSha256

*Returns the rebuilt hash obtained by traversing a Merkle tree up
from `leaf` using `proof`. A `proof` is valid if and only if the rebuilt
hash matches the root of the tree. The tree is built assuming `leaf` is
the 0 indexed `index`'th leaf from the bottom left of the tree.
_Available since v4.4._
Note this is for a Merkle tree using the sha256 hash function*


```solidity
function processInclusionProofSha256(bytes memory proof, bytes32 leaf, uint256 index) internal view returns (bytes32);
```

### merkleizeSha256

this function returns the merkle root of a tree created from a set of leaves using sha256 as its hash function

*A pre-condition to this function is that leaves.length is a power of two.  If not, the function will merkleize the inputs incorrectly.*


```solidity
function merkleizeSha256(bytes32[] memory leaves) internal pure returns (bytes32);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`leaves`|`bytes32[]`|the leaves of the merkle tree|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes32`|The computed Merkle root of the tree.|


