# MiddlewareRegistryMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/MiddlewareRegistryMock.sol)

**Inherits:**
[IRegistry](/src/contracts/interfaces/IRegistry.sol/interface.IRegistry.md), DSTest


## State Variables
### serviceManager

```solidity
IServiceManager public serviceManager;
```


### strategyManager

```solidity
IStrategyManager public strategyManager;
```


### slasher

```solidity
ISlasher public slasher;
```


## Functions
### constructor


```solidity
constructor(IServiceManager _serviceManager, IStrategyManager _strategyManager);
```

### registerOperator


```solidity
function registerOperator(address operator, uint32 serveUntil) public;
```

### deregisterOperator


```solidity
function deregisterOperator(address operator) public;
```

### propagateStakeUpdate


```solidity
function propagateStakeUpdate(address operator, uint32 blockNumber, uint256 prevElement) external;
```

### isActiveOperator


```solidity
function isActiveOperator(address operator) external pure returns (bool);
```

