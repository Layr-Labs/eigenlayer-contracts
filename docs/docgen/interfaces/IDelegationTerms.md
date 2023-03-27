# Solidity API

## IDelegationTerms

The gas budget provided to this contract in calls from EigenLayer contracts is limited.

### payForService

```solidity
function payForService(contract IERC20 token, uint256 amount) external payable
```

### onDelegationWithdrawn

```solidity
function onDelegationWithdrawn(address delegator, contract IStrategy[] stakerStrategyList, uint256[] stakerShares) external returns (bytes)
```

### onDelegationReceived

```solidity
function onDelegationReceived(address delegator, contract IStrategy[] stakerStrategyList, uint256[] stakerShares) external returns (bytes)
```

