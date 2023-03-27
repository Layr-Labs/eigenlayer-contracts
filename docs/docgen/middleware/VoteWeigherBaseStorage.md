# Solidity API

## VoteWeigherBaseStorage

This storage contract is separate from the logic to simplify the upgrade process.

### StrategyAndWeightingMultiplier

```solidity
struct StrategyAndWeightingMultiplier {
  contract IStrategy strategy;
  uint96 multiplier;
}
```

### WEIGHTING_DIVISOR

```solidity
uint256 WEIGHTING_DIVISOR
```

Constant used as a divisor in calculating weights.

### MAX_WEIGHING_FUNCTION_LENGTH

```solidity
uint8 MAX_WEIGHING_FUNCTION_LENGTH
```

Maximum length of dynamic arrays in the `strategiesConsideredAndMultipliers` mapping.

### MAX_BIPS

```solidity
uint256 MAX_BIPS
```

Constant used as a divisor in dealing with BIPS amounts.

### delegation

```solidity
contract IDelegationManager delegation
```

The address of the Delegation contract for EigenLayer.

### strategyManager

```solidity
contract IStrategyManager strategyManager
```

The address of the StrategyManager contract for EigenLayer.

### slasher

```solidity
contract ISlasher slasher
```

The address of the Slasher contract for EigenLayer.

### serviceManager

```solidity
contract IServiceManager serviceManager
```

The ServiceManager contract for this middleware, where tasks are created / initiated.

### NUMBER_OF_QUORUMS

```solidity
uint256 NUMBER_OF_QUORUMS
```

Number of quorums that are being used by the middleware.

### strategiesConsideredAndMultipliers

```solidity
mapping(uint256 => struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[]) strategiesConsideredAndMultipliers
```

mapping from quorum number to the list of strategies considered and their
corresponding multipliers for that specific quorum

### quorumBips

```solidity
mapping(uint256 => uint256) quorumBips
```

This defines the earnings split between different quorums. Mapping is quorumNumber => BIPS which the quorum earns, out of the total earnings.

_The sum of all entries, i.e. sum(quorumBips[0] through quorumBips[NUMBER_OF_QUORUMS - 1]) should *always* be 10,000!_

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS) internal
```

