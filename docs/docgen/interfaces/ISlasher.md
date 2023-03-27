# Solidity API

## ISlasher

See the `Slasher` contract itself for implementation details.

### MiddlewareTimes

```solidity
struct MiddlewareTimes {
  uint32 stalestUpdateBlock;
  uint32 latestServeUntil;
}
```

### MiddlewareDetails

```solidity
struct MiddlewareDetails {
  uint32 contractCanSlashOperatorUntil;
  uint32 latestUpdateBlock;
}
```

### optIntoSlashing

```solidity
function optIntoSlashing(address contractAddress) external
```

Gives the `contractAddress` permission to slash the funds of the caller.

_Typically, this function must be called prior to registering for a middleware._

### freezeOperator

```solidity
function freezeOperator(address toBeFrozen) external
```

Used for 'slashing' a certain operator.

_Technically the operator is 'frozen' (hence the name of this function), and then subject to slashing pending a decision by a human-in-the-loop.
The operator must have previously given the caller (which should be a contract) the ability to slash them, through a call to `optIntoSlashing`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| toBeFrozen | address | The operator to be frozen. |

### resetFrozenStatus

```solidity
function resetFrozenStatus(address[] frozenAddresses) external
```

Removes the 'frozen' status from each of the `frozenAddresses`

_Callable only by the contract owner (i.e. governance)._

### recordFirstStakeUpdate

```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntil) external
```

this function is a called by middlewares during an operator's registration to make sure the operator's stake at registration 
        is slashable until serveUntil

_adds the middleware's slashing contract to the operator's linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| serveUntil | uint32 | the timestamp until which the operator's stake at the current block is slashable |

### recordStakeUpdate

```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntil, uint256 insertAfter) external
```

this function is a called by middlewares during a stake update for an operator (perhaps to free pending withdrawals)
        to make sure the operator's stake at updateBlock is slashable until serveUntil

_insertAfter should be calculated offchain before making the transaction that calls this. this is subject to race conditions, 
     but it is anticipated to be rare and not detrimental._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| updateBlock | uint32 | the block for which the stake update is being recorded |
| serveUntil | uint32 | the timestamp until which the operator's stake at updateBlock is slashable |
| insertAfter | uint256 | the element of the operators linked list that the currently updating middleware should be inserted after |

### recordLastStakeUpdateAndRevokeSlashingAbility

```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntil) external
```

this function is a called by middlewares during an operator's deregistration to make sure the operator's stake at deregistration 
        is slashable until serveUntil

_removes the middleware's slashing contract to the operator's linked list and revokes the middleware's (i.e. caller's) ability to
slash `operator` once `serveUntil` is reached_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| serveUntil | uint32 | the timestamp until which the operator's stake at the current block is slashable |

### isFrozen

```solidity
function isFrozen(address staker) external view returns (bool)
```

Used to determine whether `staker` is actively 'frozen'. If a staker is frozen, then they are potentially subject to
slashing of their funds, and cannot cannot deposit or withdraw from the strategyManager until the slashing process is completed
and the staker's status is reset (to 'unfrozen').

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | Returns 'true' if `staker` themselves has their status set to frozen, OR if the staker is delegated to an operator who has their status set to frozen. Otherwise returns 'false'. |

### canSlash

```solidity
function canSlash(address toBeSlashed, address slashingContract) external view returns (bool)
```

Returns true if `slashingContract` is currently allowed to slash `toBeSlashed`.

### contractCanSlashOperatorUntil

```solidity
function contractCanSlashOperatorUntil(address operator, address serviceContract) external view returns (uint32)
```

Returns the UTC timestamp until which `serviceContract` is allowed to slash the `operator`.

### latestUpdateBlock

```solidity
function latestUpdateBlock(address operator, address serviceContract) external view returns (uint32)
```

Returns the block at which the `serviceContract` last updated its view of the `operator`'s stake

### getCorrectValueForInsertAfter

```solidity
function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) external view returns (uint256)
```

A search routine for finding the correct input value of `insertAfter` to `recordStakeUpdate` / `_updateMiddlewareList`.

### canWithdraw

```solidity
function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) external returns (bool)
```

Returns 'true' if `operator` can currently complete a withdrawal started at the `withdrawalStartBlock`, with `middlewareTimesIndex` used
to specify the index of a `MiddlewareTimes` struct in the operator's list (i.e. an index in `operatorToMiddlewareTimes[operator]`). The specified
struct is consulted as proof of the `operator`'s ability (or lack thereof) to complete the withdrawal.
This function will return 'false' if the operator cannot currently complete a withdrawal started at the `withdrawalStartBlock`, *or* in the event
that an incorrect `middlewareTimesIndex` is supplied, even if one or more correct inputs exist.

_The correct `middlewareTimesIndex` input should be computable off-chain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | Either the operator who queued the withdrawal themselves, or if the withdrawing party is a staker who delegated to an operator, this address is the operator *who the staker was delegated to* at the time of the `withdrawalStartBlock`. |
| withdrawalStartBlock | uint32 | The block number at which the withdrawal was initiated. |
| middlewareTimesIndex | uint256 | Indicates an index in `operatorToMiddlewareTimes[operator]` to consult as proof of the `operator`'s ability to withdraw |

### operatorToMiddlewareTimes

```solidity
function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) external view returns (struct ISlasher.MiddlewareTimes)
```

operator => 
 [
     (
         the least recent update block of all of the middlewares it's serving/served, 
         latest time that the stake bonded at that update needed to serve until
     )
 ]

### middlewareTimesLength

```solidity
function middlewareTimesLength(address operator) external view returns (uint256)
```

Getter function for fetching `operatorToMiddlewareTimes[operator].length`

### getMiddlewareTimesIndexBlock

```solidity
function getMiddlewareTimesIndexBlock(address operator, uint32 index) external view returns (uint32)
```

Getter function for fetching `operatorToMiddlewareTimes[operator][index].stalestUpdateBlock`.

### getMiddlewareTimesIndexServeUntil

```solidity
function getMiddlewareTimesIndexServeUntil(address operator, uint32 index) external view returns (uint32)
```

Getter function for fetching `operatorToMiddlewareTimes[operator][index].latestServeUntil`.

### operatorWhitelistedContractsLinkedListSize

```solidity
function operatorWhitelistedContractsLinkedListSize(address operator) external view returns (uint256)
```

Getter function for fetching `_operatorToWhitelistedContractsByUpdate[operator].size`.

### operatorWhitelistedContractsLinkedListEntry

```solidity
function operatorWhitelistedContractsLinkedListEntry(address operator, address node) external view returns (bool, uint256, uint256)
```

Getter function for fetching a single node in the operator's linked list (`_operatorToWhitelistedContractsByUpdate[operator]`).

