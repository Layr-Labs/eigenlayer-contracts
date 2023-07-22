# MerkleDelegationTerms
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/operators/MerkleDelegationTerms.sol)

**Inherits:**
Ownable, [IDelegationTerms](/src/contracts/interfaces/IDelegationTerms.sol/interface.IDelegationTerms.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This contract specifies the delegation terms of a given operator. When a staker delegates its stake to an operator,
it has to agrees to the terms set in the operator's 'Delegation Terms' contract. Payments to an operator are routed through
their specified 'Delegation Terms' contract for subsequent distribution of earnings to individual stakers.
There are also hooks that call into an operator's DelegationTerms contract when a staker delegates to or undelegates from
the operator.

*This contract uses a system in which the operator posts roots of a *sparse Merkle tree*. Each leaf of the tree is expected
to contain the **cumulative** earnings of a staker. This will reduce the total number of actions that stakers who claim only rarely
have to take, while allowing stakers to claim their earnings as often as new Merkle roots are posted.*


## State Variables
### MAX_HEIGHT

```solidity
uint256 internal constant MAX_HEIGHT = 256;
```


### cumulativeClaimedByStakerOfToken
staker => token => cumulative amount *claimed*


```solidity
mapping(address => mapping(IERC20 => uint256)) public cumulativeClaimedByStakerOfToken;
```


### merkleRoots
Array of Merkle roots with heights, each posted by the operator (contract owner)


```solidity
MerkleRootAndTreeHeight[] public merkleRoots;
```


## Functions
### operatorWithdrawal

Used by the operator to withdraw tokens directly from this contract.


```solidity
function operatorWithdrawal(TokenAndAmount[] calldata tokensAndAmounts) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`tokensAndAmounts`|`TokenAndAmount[]`|ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw.|


### postMerkleRoot

Used by the operator to post an updated root of the stakers' all-time earnings


```solidity
function postMerkleRoot(bytes32 newRoot, uint256 height) external onlyOwner;
```

### proveEarningsAndWithdraw

Called by a staker to prove the inclusion of their earnings in a Merkle root (posted by the operator) and claim them.


```solidity
function proveEarningsAndWithdraw(
    TokenAndAmount[] calldata tokensAndAmounts,
    bytes memory proof,
    uint256 nodeIndex,
    uint256 rootIndex
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`tokensAndAmounts`|`TokenAndAmount[]`|ERC20 tokens to withdraw and the amount of each respective ERC20 token to withdraw.|
|`proof`|`bytes`|Merkle proof showing that a leaf containing `(msg.sender, tokensAndAmounts)` was included in the `rootIndex`-th Merkle root posted by the operator.|
|`nodeIndex`|`uint256`|Specifies the node inside the Merkle tree corresponding to the specified root, `merkleRoots[rootIndex].root`.|
|`rootIndex`|`uint256`|Specifies the Merkle root to look up, using `merkleRoots[rootIndex]`|


### calculateLeafHash

Helper function for calculating a leaf in a Merkle tree formatted as `(address staker, TokenAndAmount[] calldata tokensAndAmounts)`


```solidity
function calculateLeafHash(address staker, TokenAndAmount[] calldata tokensAndAmounts) public pure returns (bytes32);
```

### payForService


```solidity
function payForService(IERC20, uint256) external payable;
```

### onDelegationReceived

Hook for receiving new delegation


```solidity
function onDelegationReceived(address, IStrategy[] memory, uint256[] memory) external pure returns (bytes memory);
```

### onDelegationWithdrawn

Hook for withdrawing delegation


```solidity
function onDelegationWithdrawn(address, IStrategy[] memory, uint256[] memory) external pure returns (bytes memory);
```

## Events
### NewMerkleRootPosted

```solidity
event NewMerkleRootPosted(bytes32 newRoot, uint256 height);
```

## Structs
### TokenAndAmount

```solidity
struct TokenAndAmount {
    IERC20 token;
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

