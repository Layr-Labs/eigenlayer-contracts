# MiddlewareVoteWeigherMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/MiddlewareVoteWeigherMock.sol)

**Inherits:**
[RegistryBase](/src/contracts/middleware/RegistryBase.sol/abstract.RegistryBase.md)


## State Variables
### _NUMBER_OF_QUORUMS

```solidity
uint8 _NUMBER_OF_QUORUMS = 2;
```


## Functions
### constructor


```solidity
constructor(IDelegationManager _delegation, IStrategyManager _strategyManager, IServiceManager _serviceManager)
    RegistryBase(_strategyManager, _serviceManager, _NUMBER_OF_QUORUMS);
```

### initialize


```solidity
function initialize(
    uint256[] memory _quorumBips,
    StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
    StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers
) external initializer;
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

