# Utils
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/Utils.sol)


## State Variables
### dummyAdmin

```solidity
address constant dummyAdmin = address(uint160(uint256(keccak256("DummyAdmin"))));
```


## Functions
### deployNewStrategy


```solidity
function deployNewStrategy(
    IERC20 token,
    IStrategyManager strategyManager,
    IPauserRegistry pauserRegistry,
    address admin
) public returns (StrategyBase);
```

