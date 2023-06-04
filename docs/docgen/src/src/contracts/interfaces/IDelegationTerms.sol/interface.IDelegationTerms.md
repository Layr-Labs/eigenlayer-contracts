# IDelegationTerms
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IDelegationTerms.sol)

**Author:**
Layr Labs, Inc.

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

