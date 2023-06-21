# ISlasher
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/ISlasher.sol)

**Author:**
Layr Labs, Inc.

See the `Slasher` contract itself for implementation details.


## Functions
### optIntoSlashing

Gives the `contractAddress` permission to slash the funds of the caller.

*Typically, this function must be called prior to registering for a middleware.*


```solidity
function optIntoSlashing(address contractAddress) external;
```

### freezeOperator

Used for 'slashing' a certain operator.

*Technically the operator is 'frozen' (hence the name of this function), and then subject to slashing pending a decision by a human-in-the-loop.*

*The operator must have previously given the caller (which should be a contract) the ability to slash them, through a call to `optIntoSlashing`.*


```solidity
function freezeOperator(address toBeFrozen) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`toBeFrozen`|`address`|The operator to be frozen.|


### resetFrozenStatus

Removes the 'frozen' status from each of the `frozenAddresses`

*Callable only by the contract owner (i.e. governance).*


```solidity
function resetFrozenStatus(address[] calldata frozenAddresses) external;
```

### recordFirstStakeUpdate

this function is a called by middlewares during an operator's registration to make sure the operator's stake at registration
is slashable until serveUntil

*adds the middleware's slashing contract to the operator's linked list*


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at the current block is slashable|


### recordStakeUpdate

this function is a called by middlewares during a stake update for an operator (perhaps to free pending withdrawals)
to make sure the operator's stake at updateBlock is slashable until serveUntil

*insertAfter should be calculated offchain before making the transaction that calls this. this is subject to race conditions,
but it is anticipated to be rare and not detrimental.*


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`updateBlock`|`uint32`|the block for which the stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at updateBlock is slashable|
|`insertAfter`|`uint256`|the element of the operators linked list that the currently updating middleware should be inserted after|


### recordLastStakeUpdateAndRevokeSlashingAbility

this function is a called by middlewares during an operator's deregistration to make sure the operator's stake at deregistration
is slashable until serveUntil

*removes the middleware's slashing contract to the operator's linked list and revokes the middleware's (i.e. caller's) ability to
slash `operator` once `serveUntil` is reached*


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at the current block is slashable|


### isFrozen

Used to determine whether `staker` is actively 'frozen'. If a staker is frozen, then they are potentially subject to
slashing of their funds, and cannot cannot deposit or withdraw from the strategyManager until the slashing process is completed
and the staker's status is reset (to 'unfrozen').


```solidity
function isFrozen(address staker) external view returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|The staker of interest.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|Returns 'true' if `staker` themselves has their status set to frozen, OR if the staker is delegated to an operator who has their status set to frozen. Otherwise returns 'false'.|


### canSlash

Returns true if `slashingContract` is currently allowed to slash `toBeSlashed`.


```solidity
function canSlash(address toBeSlashed, address slashingContract) external view returns (bool);
```

### contractCanSlashOperatorUntilBlock

Returns the block until which `serviceContract` is allowed to slash the `operator`.


```solidity
function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) external view returns (uint32);
```

### latestUpdateBlock

Returns the block at which the `serviceContract` last updated its view of the `operator`'s stake


```solidity
function latestUpdateBlock(address operator, address serviceContract) external view returns (uint32);
```

### getCorrectValueForInsertAfter

A search routine for finding the correct input value of `insertAfter` to `recordStakeUpdate` / `_updateMiddlewareList`.


```solidity
function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) external view returns (uint256);
```

### canWithdraw

Returns 'true' if `operator` can currently complete a withdrawal started at the `withdrawalStartBlock`, with `middlewareTimesIndex` used
to specify the index of a `MiddlewareTimes` struct in the operator's list (i.e. an index in `operatorToMiddlewareTimes[operator]`). The specified
struct is consulted as proof of the `operator`'s ability (or lack thereof) to complete the withdrawal.
This function will return 'false' if the operator cannot currently complete a withdrawal started at the `withdrawalStartBlock`, *or* in the event
that an incorrect `middlewareTimesIndex` is supplied, even if one or more correct inputs exist.

*The correct `middlewareTimesIndex` input should be computable off-chain.*


```solidity
function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex)
    external
    returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|Either the operator who queued the withdrawal themselves, or if the withdrawing party is a staker who delegated to an operator, this address is the operator *who the staker was delegated to* at the time of the `withdrawalStartBlock`.|
|`withdrawalStartBlock`|`uint32`|The block number at which the withdrawal was initiated.|
|`middlewareTimesIndex`|`uint256`|Indicates an index in `operatorToMiddlewareTimes[operator]` to consult as proof of the `operator`'s ability to withdraw|


### operatorToMiddlewareTimes

operator =>
[
(
the least recent update block of all of the middlewares it's serving/served,
latest time that the stake bonded at that update needed to serve until
)
]


```solidity
function operatorToMiddlewareTimes(address operator, uint256 arrayIndex)
    external
    view
    returns (MiddlewareTimes memory);
```

### middlewareTimesLength

Getter function for fetching `operatorToMiddlewareTimes[operator].length`


```solidity
function middlewareTimesLength(address operator) external view returns (uint256);
```

### getMiddlewareTimesIndexBlock

Getter function for fetching `operatorToMiddlewareTimes[operator][index].stalestUpdateBlock`.


```solidity
function getMiddlewareTimesIndexBlock(address operator, uint32 index) external view returns (uint32);
```

### getMiddlewareTimesIndexServeUntilBlock

Getter function for fetching `operatorToMiddlewareTimes[operator][index].latestServeUntil`.


```solidity
function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) external view returns (uint32);
```

### operatorWhitelistedContractsLinkedListSize

Getter function for fetching `_operatorToWhitelistedContractsByUpdate[operator].size`.


```solidity
function operatorWhitelistedContractsLinkedListSize(address operator) external view returns (uint256);
```

### operatorWhitelistedContractsLinkedListEntry

Getter function for fetching a single node in the operator's linked list (`_operatorToWhitelistedContractsByUpdate[operator]`).


```solidity
function operatorWhitelistedContractsLinkedListEntry(address operator, address node)
    external
    view
    returns (bool, uint256, uint256);
```

## Structs
### MiddlewareTimes

```solidity
struct MiddlewareTimes {
    uint32 stalestUpdateBlock;
    uint32 latestServeUntilBlock;
}
```

### MiddlewareDetails

```solidity
struct MiddlewareDetails {
    uint32 contractCanSlashOperatorUntilBlock;
    uint32 latestUpdateBlock;
}
```

