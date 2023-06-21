# VoteWeigherBaseStorage
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/VoteWeigherBaseStorage.sol)

**Inherits:**
Initializable, [IVoteWeigher](/docs/docgen/src/src/contracts/interfaces/IVoteWeigher.sol/interface.IVoteWeigher.md)

**Author:**
Layr Labs, Inc.

This storage contract is separate from the logic to simplify the upgrade process.


## State Variables
### WEIGHTING_DIVISOR
Constant used as a divisor in calculating weights.


```solidity
uint256 internal constant WEIGHTING_DIVISOR = 1e18;
```


### MAX_WEIGHING_FUNCTION_LENGTH
Maximum length of dynamic arrays in the `strategiesConsideredAndMultipliers` mapping.


```solidity
uint8 internal constant MAX_WEIGHING_FUNCTION_LENGTH = 32;
```


### MAX_BIPS
Constant used as a divisor in dealing with BIPS amounts.


```solidity
uint256 internal constant MAX_BIPS = 10000;
```


### delegation
The address of the Delegation contract for EigenLayer.


```solidity
IDelegationManager public immutable delegation;
```


### strategyManager
The address of the StrategyManager contract for EigenLayer.


```solidity
IStrategyManager public immutable strategyManager;
```


### slasher
The address of the Slasher contract for EigenLayer.


```solidity
ISlasher public immutable slasher;
```


### serviceManager
The ServiceManager contract for this middleware, where tasks are created / initiated.


```solidity
IServiceManager public immutable serviceManager;
```


### NUMBER_OF_QUORUMS
Number of quorums that are being used by the middleware.


```solidity
uint256 public immutable NUMBER_OF_QUORUMS;
```


### strategiesConsideredAndMultipliers
mapping from quorum number to the list of strategies considered and their
corresponding multipliers for that specific quorum


```solidity
mapping(uint256 => StrategyAndWeightingMultiplier[]) public strategiesConsideredAndMultipliers;
```


### quorumBips
This defines the earnings split between different quorums. Mapping is quorumNumber => BIPS which the quorum earns, out of the total earnings.

*The sum of all entries, i.e. sum(quorumBips[0] through quorumBips[NUMBER_OF_QUORUMS - 1]) should *always* be 10,000!*


```solidity
mapping(uint256 => uint256) public quorumBips;
```


## Functions
### constructor


```solidity
constructor(IStrategyManager _strategyManager, IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS);
```

## Structs
### StrategyAndWeightingMultiplier
In weighing a particular strategy, the amount of underlying asset for that strategy is
multiplied by its multiplier, then divided by WEIGHTING_DIVISOR


```solidity
struct StrategyAndWeightingMultiplier {
    IStrategy strategy;
    uint96 multiplier;
}
```

