# Solidity API

## IWhitelister

### whitelist

```solidity
function whitelist(address operator) external
```

### getStaker

```solidity
function getStaker(address operator) external returns (address)
```

### depositIntoStrategy

```solidity
function depositIntoStrategy(address staker, contract IStrategy strategy, contract IERC20 token, uint256 amount) external returns (bytes)
```

### queueWithdrawal

```solidity
function queueWithdrawal(address staker, uint256[] strategyIndexes, contract IStrategy[] strategies, uint256[] shares, address withdrawer, bool undelegateIfPossible) external returns (bytes)
```

### completeQueuedWithdrawal

```solidity
function completeQueuedWithdrawal(address staker, struct IStrategyManager.QueuedWithdrawal queuedWithdrawal, contract IERC20[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) external returns (bytes)
```

### transfer

```solidity
function transfer(address staker, address token, address to, uint256 amount) external returns (bytes)
```

### callAddress

```solidity
function callAddress(address to, bytes data) external payable returns (bytes)
```

