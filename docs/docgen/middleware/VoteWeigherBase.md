# Solidity API

## VoteWeigherBase

This contract is used for
- computing the total weight of an operator for any of the quorums that are considered
by the middleware
- addition and removal of strategies and the associated weighting criteria that are assigned
by the middleware for each of the quorum(s)
@dev

### StrategyAddedToQuorum

```solidity
event StrategyAddedToQuorum(uint256 quorumNumber, contract IStrategy strategy)
```

emitted when `strategy` has been added to the array at `strategiesConsideredAndMultipliers[quorumNumber]`

### StrategyRemovedFromQuorum

```solidity
event StrategyRemovedFromQuorum(uint256 quorumNumber, contract IStrategy strategy)
```

emitted when `strategy` has removed from the array at `strategiesConsideredAndMultipliers[quorumNumber]`

### onlyServiceManagerOwner

```solidity
modifier onlyServiceManagerOwner()
```

when applied to a function, ensures that the function is only callable by the current `owner` of the `serviceManager`

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS) internal
```

Sets the (immutable) `strategyManager` and `serviceManager` addresses, as well as the (immutable) `NUMBER_OF_QUORUMS` variable

### _initialize

```solidity
function _initialize(uint256[] _quorumBips) internal virtual
```

Set the split in earnings between the different quorums.

### weightOfOperator

```solidity
function weightOfOperator(address operator, uint256 quorumNumber) public virtual returns (uint96)
```

This function computes the total weight of the @param operator in the quorum @param quorumNumber.

_returns zero in the case that `quorumNumber` is greater than or equal to `NUMBER_OF_QUORUMS`_

### addStrategiesConsideredAndMultipliers

```solidity
function addStrategiesConsideredAndMultipliers(uint256 quorumNumber, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _newStrategiesConsideredAndMultipliers) external virtual
```

Adds new strategies and the associated multipliers to the @param quorumNumber.

### removeStrategiesConsideredAndMultipliers

```solidity
function removeStrategiesConsideredAndMultipliers(uint256 quorumNumber, contract IStrategy[] _strategiesToRemove, uint256[] indicesToRemove) external virtual
```

This function is used for removing strategies and their associated weights from the
mapping strategiesConsideredAndMultipliers for a specific @param quorumNumber.

_higher indices should be *first* in the list of @param indicesToRemove, since otherwise
the removal of lower index entries will cause a shift in the indices of the other strategiesToRemove_

### modifyStrategyWeights

```solidity
function modifyStrategyWeights(uint256 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) external virtual
```

This function is used for modifying the weights of strategies that are already in the
mapping strategiesConsideredAndMultipliers for a specific @param quorumNumber.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| quorumNumber | uint256 |  |
| strategyIndices | uint256[] | is a correctness-check input -- the supplied values must match the indices of the strategiesToModifyWeightsOf in strategiesConsideredAndMultipliers[quorumNumber] |
| newMultipliers | uint96[] |  |

### strategiesConsideredAndMultipliersLength

```solidity
function strategiesConsideredAndMultipliersLength(uint256 quorumNumber) public view returns (uint256)
```

Returns the length of the dynamic array stored in `strategiesConsideredAndMultipliers[quorumNumber]`.

_Reverts if `quorumNumber` < `NUMBER_OF_QUORUMS`, i.e. the input is out of bounds._

### _addStrategiesConsideredAndMultipliers

```solidity
function _addStrategiesConsideredAndMultipliers(uint256 quorumNumber, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _newStrategiesConsideredAndMultipliers) internal
```

Adds `_newStrategiesConsideredAndMultipliers` to the `quorumNumber`-th quorum.

_Checks to make sure that the *same* strategy cannot be added multiple times (checks against both against existing and new strategies).
This function has no check to make sure that the strategies for a single quorum have the same underlying asset. This is a concious choice,
since a middleware may want, e.g., a stablecoin quorum that accepts USDC, USDT, DAI, etc. as underlying assets and trades them as "equivalent"._

