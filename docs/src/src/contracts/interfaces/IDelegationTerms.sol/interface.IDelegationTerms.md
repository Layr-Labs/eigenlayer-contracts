# IDelegationTerms
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IDelegationTerms.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

The gas budget provided to this contract in calls from EigenLayer contracts is limited.


## Functions
### payForService


```solidity
function payForService(IERC20 token, uint256 amount) external payable;
```

### onDelegationWithdrawn


```solidity
function onDelegationWithdrawn(address delegator, IStrategy[] memory stakerStrategyList, uint256[] memory stakerShares)
    external
    returns (bytes memory);
```

### onDelegationReceived


```solidity
function onDelegationReceived(address delegator, IStrategy[] memory stakerStrategyList, uint256[] memory stakerShares)
    external
    returns (bytes memory);
```

