# IWhitelister
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IWhitelister.sol)


## Functions
### whitelist


```solidity
function whitelist(address operator) external;
```

### getStaker


```solidity
function getStaker(address operator) external returns (address);
```

### depositIntoStrategy


```solidity
function depositIntoStrategy(address staker, IStrategy strategy, IERC20 token, uint256 amount)
    external
    returns (bytes memory);
```

### queueWithdrawal


```solidity
function queueWithdrawal(
    address staker,
    uint256[] calldata strategyIndexes,
    IStrategy[] calldata strategies,
    uint256[] calldata shares,
    address withdrawer,
    bool undelegateIfPossible
) external returns (bytes memory);
```

### completeQueuedWithdrawal


```solidity
function completeQueuedWithdrawal(
    address staker,
    IStrategyManager.QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) external returns (bytes memory);
```

### transfer


```solidity
function transfer(address staker, address token, address to, uint256 amount) external returns (bytes memory);
```

### callAddress


```solidity
function callAddress(address to, bytes memory data) external payable returns (bytes memory);
```

