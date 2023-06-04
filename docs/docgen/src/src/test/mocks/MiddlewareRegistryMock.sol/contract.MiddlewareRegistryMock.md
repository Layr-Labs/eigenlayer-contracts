# MiddlewareRegistryMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/MiddlewareRegistryMock.sol)

**Inherits:**
[IRegistry](/docs/docgen/src/src/contracts/interfaces/IRegistry.sol/interface.IRegistry.md), DSTest


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

