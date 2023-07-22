# VoteWeigherBase
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/middleware/VoteWeigherBase.sol)

**Inherits:**
[VoteWeigherBaseStorage](/src/contracts/middleware/VoteWeigherBaseStorage.sol/abstract.VoteWeigherBaseStorage.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This contract is used for
- computing the total weight of an operator for any of the quorums that are considered
by the middleware
- addition and removal of strategies and the associated weighting criteria that are assigned
by the middleware for each of the quorum(s)

**


## Functions
### onlyServiceManagerOwner

when applied to a function, ensures that the function is only callable by the current `owner` of the `serviceManager`


```solidity
modifier onlyServiceManagerOwner();
```

### constructor

Sets the (immutable) `strategyManager` and `serviceManager` addresses, as well as the (immutable) `NUMBER_OF_QUORUMS` variable


```solidity
constructor(IStrategyManager _strategyManager, IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS)
    VoteWeigherBaseStorage(_strategyManager, _serviceManager, _NUMBER_OF_QUORUMS);
```

### _initialize

Set the split in earnings between the different quorums.


```solidity
function _initialize(uint256[] memory _quorumBips) internal virtual onlyInitializing;
```

### weightOfOperator

This function computes the total weight of the @param operator in the quorum @param quorumNumber.

*returns zero in the case that `quorumNumber` is greater than or equal to `NUMBER_OF_QUORUMS`*


```solidity
function weightOfOperator(address operator, uint256 quorumNumber) public virtual returns (uint96);
```

### addStrategiesConsideredAndMultipliers

Adds new strategies and the associated multipliers to the @param quorumNumber.


```solidity
function addStrategiesConsideredAndMultipliers(
    uint256 quorumNumber,
    StrategyAndWeightingMultiplier[] memory _newStrategiesConsideredAndMultipliers
) external virtual onlyServiceManagerOwner;
```

### removeStrategiesConsideredAndMultipliers

This function is used for removing strategies and their associated weights from the
mapping strategiesConsideredAndMultipliers for a specific @param quorumNumber.

*higher indices should be *first* in the list of @param indicesToRemove, since otherwise
the removal of lower index entries will cause a shift in the indices of the other strategiesToRemove*


```solidity
function removeStrategiesConsideredAndMultipliers(
    uint256 quorumNumber,
    IStrategy[] calldata _strategiesToRemove,
    uint256[] calldata indicesToRemove
) external virtual onlyServiceManagerOwner;
```

### modifyStrategyWeights

This function is used for modifying the weights of strategies that are already in the
mapping strategiesConsideredAndMultipliers for a specific @param quorumNumber.


```solidity
function modifyStrategyWeights(
    uint256 quorumNumber,
    uint256[] calldata strategyIndices,
    uint96[] calldata newMultipliers
) external virtual onlyServiceManagerOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`quorumNumber`|`uint256`||
|`strategyIndices`|`uint256[]`|is a correctness-check input -- the supplied values must match the indices of the strategiesToModifyWeightsOf in strategiesConsideredAndMultipliers[quorumNumber]|
|`newMultipliers`|`uint96[]`||


### strategiesConsideredAndMultipliersLength

Returns the length of the dynamic array stored in `strategiesConsideredAndMultipliers[quorumNumber]`.

*Reverts if `quorumNumber` < `NUMBER_OF_QUORUMS`, i.e. the input is out of bounds.*


```solidity
function strategiesConsideredAndMultipliersLength(uint256 quorumNumber) public view returns (uint256);
```

### _addStrategiesConsideredAndMultipliers

Adds `_newStrategiesConsideredAndMultipliers` to the `quorumNumber`-th quorum.

*Checks to make sure that the *same* strategy cannot be added multiple times (checks against both against existing and new strategies).*

*This function has no check to make sure that the strategies for a single quorum have the same underlying asset. This is a concious choice,
since a middleware may want, e.g., a stablecoin quorum that accepts USDC, USDT, DAI, etc. as underlying assets and trades them as "equivalent".*


```solidity
function _addStrategiesConsideredAndMultipliers(
    uint256 quorumNumber,
    StrategyAndWeightingMultiplier[] memory _newStrategiesConsideredAndMultipliers
) internal;
```

## Events
### StrategyAddedToQuorum
emitted when `strategy` has been added to the array at `strategiesConsideredAndMultipliers[quorumNumber]`


```solidity
event StrategyAddedToQuorum(uint256 indexed quorumNumber, IStrategy strategy);
```

### StrategyRemovedFromQuorum
emitted when `strategy` has removed from the array at `strategiesConsideredAndMultipliers[quorumNumber]`


```solidity
event StrategyRemovedFromQuorum(uint256 indexed quorumNumber, IStrategy strategy);
```

