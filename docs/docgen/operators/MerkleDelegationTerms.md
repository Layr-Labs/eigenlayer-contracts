# Solidity API

## MerkleDelegationTerms

This contract specifies the delegation terms of a given operator. When a staker delegates its stake to an operator,
it has to agrees to the terms set in the operator's 'Delegation Terms' contract. Payments to an operator are routed through
their specified 'Delegation Terms' contract for subsequent distribution of earnings to individual stakers.
There are also hooks that call into an operator's DelegationTerms contract when a staker delegates to or undelegates from
the operator.

_This contract uses a system in which the operator posts roots of a *sparse Merkle tree*. Each leaf of the tree is expected
to contain the **cumulative** earnings of a staker. This will reduce the total number of actions that stakers who claim only rarely
have to take, while allowing stakers to claim their earnings as often as new Merkle roots are posted._

### TokenAndAmount

```solidity
struct TokenAndAmount {
  contract IERC20 token;
  uint256 amount;
}
```

### MerkleRootAndTreeHeight

```solidity
struct MerkleRootAndTreeHeight {
  bytes32 root;
  uint256 height;
}
```

### MAX_HEIGHT

```solidity
uint256 MAX_HEIGHT
```

### cumulativeClaimedByStakerOfToken

```solidity
mapping(address => mapping(contract IERC20 => uint256)) cumulativeClaimedByStakerOfToken
```

staker => token => cumulative amount *claimed*

### merkleRoots

```solidity
struct MerkleDelegationTerms.MerkleRootAndTreeHeight[] merkleRoots
```

Array of Merkle roots with heights, each posted by the operator (contract owner)

### NewMerkleRootPosted

```solidity
event NewMerkleRootPosted(bytes32 newRoot, uint256 height)
```

### operatorWithdrawal

```solidity
function operatorWithdrawal(struct MerkleDelegationTerms.TokenAndAmount[] tokensAndAmounts) external
```

Used by the operator to withdraw tokens directly from this contract.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokensAndAmounts | struct MerkleDelegationTerms.TokenAndAmount[] | ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw. |

### postMerkleRoot

```solidity
function postMerkleRoot(bytes32 newRoot, uint256 height) external
```

Used by the operator to post an updated root of the stakers' all-time earnings

### proveEarningsAndWithdraw

```solidity
function proveEarningsAndWithdraw(struct MerkleDelegationTerms.TokenAndAmount[] tokensAndAmounts, bytes proof, uint256 nodeIndex, uint256 rootIndex) external
```

Called by a staker to prove the inclusion of their earnings in a Merkle root (posted by the operator) and claim them.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokensAndAmounts | struct MerkleDelegationTerms.TokenAndAmount[] | ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw. |
| proof | bytes | Merkle proof showing that a leaf containing `(msg.sender, tokensAndAmounts)` was included in the `rootIndex`-th Merkle root posted by the operator. |
| nodeIndex | uint256 | Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`. |
| rootIndex | uint256 | Specifies the Merkle root to look up, using `merkleRoots[rootIndex]` |

### calculateLeafHash

```solidity
function calculateLeafHash(address staker, struct MerkleDelegationTerms.TokenAndAmount[] tokensAndAmounts) public pure returns (bytes32)
```

Helper function for calculating a leaf in a Merkle tree formatted as `(address staker, TokenAndAmount[] calldata tokensAndAmounts)`

### payForService

```solidity
function payForService(contract IERC20, uint256) external payable
```

### onDelegationReceived

```solidity
function onDelegationReceived(address, contract IStrategy[], uint256[]) external pure returns (bytes)
```

Hook for receiving new delegation

### onDelegationWithdrawn

```solidity
function onDelegationWithdrawn(address, contract IStrategy[], uint256[]) external pure returns (bytes)
```

Hook for withdrawing delegation

