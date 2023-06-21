# MiddlewareVoteWeigherMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/MiddlewareVoteWeigherMock.sol)

**Inherits:**
[RegistryBase](/docs/docgen/src/src/contracts/middleware/RegistryBase.sol/abstract.RegistryBase.md)


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

