# Slasher
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/core/Slasher.sol)

**Inherits:**
Initializable, OwnableUpgradeable, [ISlasher](/src/contracts/interfaces/ISlasher.sol/interface.ISlasher.md), [Pausable](/src/contracts/permissions/Pausable.sol/contract.Pausable.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This contract specifies details on slashing. The functionalities are:
- adding contracts who have permission to perform slashing,
- revoking permission for slashing from specified contracts,
- tracking historic stake updates to ensure that withdrawals can only be completed once no middlewares have slashing rights
over the funds being withdrawn


## State Variables
### HEAD

```solidity
uint256 private constant HEAD = 0;
```


### PAUSED_OPT_INTO_SLASHING

```solidity
uint8 internal constant PAUSED_OPT_INTO_SLASHING = 0;
```


### PAUSED_FIRST_STAKE_UPDATE

```solidity
uint8 internal constant PAUSED_FIRST_STAKE_UPDATE = 1;
```


### PAUSED_NEW_FREEZING

```solidity
uint8 internal constant PAUSED_NEW_FREEZING = 2;
```


### strategyManager
The central StrategyManager contract of EigenLayer


```solidity
IStrategyManager public immutable strategyManager;
```


### delegation
The DelegationManager contract of EigenLayer


```solidity
IDelegationManager public immutable delegation;
```


### _whitelistedContractDetails

```solidity
mapping(address => mapping(address => MiddlewareDetails)) internal _whitelistedContractDetails;
```


### frozenStatus

```solidity
mapping(address => bool) internal frozenStatus;
```


### MAX_CAN_SLASH_UNTIL

```solidity
uint32 internal constant MAX_CAN_SLASH_UNTIL = type(uint32).max;
```


### _operatorToWhitelistedContractsByUpdate
operator => a linked list of the addresses of the whitelisted middleware with permission to slash the operator, i.e. which
the operator is serving. Sorted by the block at which they were last updated (content of updates below) in ascending order.
This means the 'HEAD' (i.e. start) of the linked list will have the stalest 'updateBlock' value.


```solidity
mapping(address => StructuredLinkedList.List) internal _operatorToWhitelistedContractsByUpdate;
```


### _operatorToMiddlewareTimes
operator =>
[
(
the least recent update block of all of the middlewares it's serving/served,
latest time that the stake bonded at that update needed to serve until
)
]


```solidity
mapping(address => MiddlewareTimes[]) internal _operatorToMiddlewareTimes;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[46] private __gap;
```


## Functions
### constructor


```solidity
constructor(IStrategyManager _strategyManager, IDelegationManager _delegation);
```

### onlyRegisteredForService

Ensures that the operator has opted into slashing by the caller, and that the caller has never revoked its slashing ability.


```solidity
modifier onlyRegisteredForService(address operator);
```

### initialize


```solidity
function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus)
    external
    initializer;
```

### optIntoSlashing

Gives the `contractAddress` permission to slash the funds of the caller.

*Typically, this function must be called prior to registering for a middleware.*


```solidity
function optIntoSlashing(address contractAddress) external onlyWhenNotPaused(PAUSED_OPT_INTO_SLASHING);
```

### freezeOperator

Used for 'slashing' a certain operator.

*Technically the operator is 'frozen' (hence the name of this function), and then subject to slashing pending a decision by a human-in-the-loop.*

*The operator must have previously given the caller (which should be a contract) the ability to slash them, through a call to `optIntoSlashing`.*


```solidity
function freezeOperator(address toBeFrozen) external onlyWhenNotPaused(PAUSED_NEW_FREEZING);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`toBeFrozen`|`address`|The operator to be frozen.|


### resetFrozenStatus

Removes the 'frozen' status from each of the `frozenAddresses`

*Callable only by the contract owner (i.e. governance).*


```solidity
function resetFrozenStatus(address[] calldata frozenAddresses) external onlyOwner;
```

### recordFirstStakeUpdate

this function is a called by middlewares during an operator's registration to make sure the operator's stake at registration
is slashable until serveUntilBlock

*adds the middleware's slashing contract to the operator's linked list*


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock)
    external
    onlyWhenNotPaused(PAUSED_FIRST_STAKE_UPDATE)
    onlyRegisteredForService(operator);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at the current block is slashable|


### recordStakeUpdate

this function is a called by middlewares during a stake update for an operator (perhaps to free pending withdrawals)
to make sure the operator's stake at updateBlock is slashable until serveUntilBlock

*insertAfter should be calculated offchain before making the transaction that calls this. this is subject to race conditions,
but it is anticipated to be rare and not detrimental.*


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter)
    external
    onlyRegisteredForService(operator);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`updateBlock`|`uint32`|the block for which the stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at updateBlock is slashable|
|`insertAfter`|`uint256`|the element of the operators linked list that the currently updating middleware should be inserted after|


### recordLastStakeUpdateAndRevokeSlashingAbility

Move the middleware to its correct update position, determined by `updateBlock` and indicated via `insertAfter`.
If the the middleware is the only one in the list, then no need to mutate the list

this function is a called by middlewares during an operator's deregistration to make sure the operator's stake at deregistration
is slashable until serveUntilBlock

*removes the middleware's slashing contract to the operator's linked list and revokes the middleware's (i.e. caller's) ability to
slash `operator` once `serveUntilBlock` is reached*


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock)
    external
    onlyRegisteredForService(operator);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the operator whose stake update is being recorded|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at the current block is slashable|


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

### whitelistedContractDetails


```solidity
function whitelistedContractDetails(address operator, address serviceContract)
    external
    view
    returns (MiddlewareDetails memory);
```

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
function canSlash(address toBeSlashed, address slashingContract) public view returns (bool);
```

### canWithdraw

Returns 'true' if `operator` can currently complete a withdrawal started at the `withdrawalStartBlock`, with `middlewareTimesIndex` used
to specify the index of a `MiddlewareTimes` struct in the operator's list (i.e. an index in `_operatorToMiddlewareTimes[operator]`). The specified
struct is consulted as proof of the `operator`'s ability (or lack thereof) to complete the withdrawal.
This function will return 'false' if the operator cannot currently complete a withdrawal started at the `withdrawalStartBlock`, *or* in the event
that an incorrect `middlewareTimesIndex` is supplied, even if one or more correct inputs exist.

*The correct `middlewareTimesIndex` input should be computable off-chain.*


```solidity
function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex)
    external
    view
    returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|Either the operator who queued the withdrawal themselves, or if the withdrawing party is a staker who delegated to an operator, this address is the operator *who the staker was delegated to* at the time of the `withdrawalStartBlock`.|
|`withdrawalStartBlock`|`uint32`|The block number at which the withdrawal was initiated.|
|`middlewareTimesIndex`|`uint256`|Indicates an index in `_operatorToMiddlewareTimes[operator]` to consult as proof of the `operator`'s ability to withdraw|


### operatorToMiddlewareTimes

Case-handling for if the operator is not registered for any middlewares (i.e. they previously registered but are no longer registered for any),
AND the withdrawal was initiated after the 'stalestUpdateBlock' of the MiddlewareTimes struct specified by the provided `middlewareTimesIndex`.
NOTE: we check the 2nd of these 2 conditions first for gas efficiency, to help avoid an extra SLOAD in all other cases.
In this case, we just check against the 'latestServeUntilBlock' of the last MiddlewareTimes struct. This is because the operator not being registered
for any middlewares (i.e. `_operatorToWhitelistedContractsByUpdate.size == 0`) means no new MiddlewareTimes structs will be being pushed, *and* the operator
will not be undertaking any new obligations (so just checking against the last entry is OK, unlike when the operator is actively registered for >=1 middleware).

Getter function for fetching `_operatorToMiddlewareTimes[operator][arrayIndex]`.


```solidity
function operatorToMiddlewareTimes(address operator, uint256 arrayIndex)
    external
    view
    returns (MiddlewareTimes memory);
```

### middlewareTimesLength

Getter function for fetching `_operatorToMiddlewareTimes[operator].length`.


```solidity
function middlewareTimesLength(address operator) external view returns (uint256);
```

### getMiddlewareTimesIndexStalestUpdateBlock

Getter function for fetching `_operatorToMiddlewareTimes[operator][index].stalestUpdateBlock`.


```solidity
function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) external view returns (uint32);
```

### getMiddlewareTimesIndexServeUntilBlock

Getter function for fetching `_operatorToMiddlewareTimes[operator][index].latestServeUntilBlock`.


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

### getCorrectValueForInsertAfter

A search routine for finding the correct input value of `insertAfter` to `recordStakeUpdate` / `_updateMiddlewareList`.

*Used within this contract only as a fallback in the case when an incorrect value of `insertAfter` is supplied as an input to `_updateMiddlewareList`.*

*The return value should *either* be 'HEAD' (i.e. zero) in the event that the node being inserted in the linked list has an `updateBlock`
that is less than the HEAD of the list, *or* the return value should specify the last `node` in the linked list for which
`_whitelistedContractDetails[operator][node].latestUpdateBlock <= updateBlock`,
i.e. the node such that the *next* node either doesn't exist,
OR
`_whitelistedContractDetails[operator][nextNode].latestUpdateBlock > updateBlock`.*


```solidity
function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) public view returns (uint256);
```

### getPreviousWhitelistedContractByUpdate

Special case:
If the node being inserted in the linked list has an `updateBlock` that is less than the HEAD of the list, then we set `insertAfter = HEAD`.
In _updateMiddlewareList(), the new node will be pushed to the front (HEAD) of the list.
`node` being zero (i.e. equal to 'HEAD') indicates an empty/non-existent node, i.e. reaching the end of the linked list.
Since the linked list is ordered in ascending order of update blocks, we simply start from the head of the list and step through until
we find a the *last* `node` for which `_whitelistedContractDetails[operator][node].latestUpdateBlock <= updateBlock`, or
otherwise reach the end of the list.

gets the node previous to the given node in the operators middleware update linked list

*used in offchain libs for updating stakes*


```solidity
function getPreviousWhitelistedContractByUpdate(address operator, uint256 node) external view returns (bool, uint256);
```

### _optIntoSlashing


```solidity
function _optIntoSlashing(address operator, address contractAddress) internal;
```

### _revokeSlashingAbility


```solidity
function _revokeSlashingAbility(address operator, address contractAddress, uint32 serveUntilBlock) internal;
```

### _freezeOperator


```solidity
function _freezeOperator(address toBeFrozen, address slashingContract) internal;
```

### _resetFrozenStatus


```solidity
function _resetFrozenStatus(address previouslySlashedAddress) internal;
```

### _recordUpdateAndAddToMiddlewareTimes

records the most recent updateBlock for the currently updating middleware and appends an entry to the operator's list of
MiddlewareTimes if relavent information has updated

*this function is only called during externally called stake updates by middleware contracts that can slash operator*


```solidity
function _recordUpdateAndAddToMiddlewareTimes(address operator, uint32 updateBlock, uint32 serveUntilBlock) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|the entity whose stake update is being recorded|
|`updateBlock`|`uint32`|the block number for which the currently updating middleware is updating the serveUntilBlock for|
|`serveUntilBlock`|`uint32`|the block until which the operator's stake at updateBlock is slashable|


### _updateMiddlewareList

A routine for updating the `operator`'s linked list of middlewares, inside `recordStakeUpdate`.


```solidity
function _updateMiddlewareList(address operator, uint32 updateBlock, uint256 insertAfter) internal;
```

### _addressToUint

boolean used to track if the `insertAfter input to this function is incorrect. If it is, then `runFallbackRoutine` will
be flipped to 'true', and we will use `getCorrectValueForInsertAfter` to find the correct input. This routine helps solve
a race condition where the proper value of `insertAfter` changes while a transaction is pending.
Make sure `insertAfter` specifies a node for which the most recent updateBlock was *at or before* updateBlock.
Again, if not,  we will use the fallback routine to find the correct value for `insertAfter`.


```solidity
function _addressToUint(address addr) internal pure returns (uint256);
```

### _uintToAddress


```solidity
function _uintToAddress(uint256 x) internal pure returns (address);
```

## Events
### MiddlewareTimesAdded
Emitted when a middleware times is added to `operator`'s array.


```solidity
event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock);
```

### OptedIntoSlashing
Emitted when `operator` begins to allow `contractAddress` to slash them.


```solidity
event OptedIntoSlashing(address indexed operator, address indexed contractAddress);
```

### SlashingAbilityRevoked
Emitted when `contractAddress` signals that it will no longer be able to slash `operator` after the `contractCanSlashOperatorUntilBlock`.


```solidity
event SlashingAbilityRevoked(
    address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock
);
```

### OperatorFrozen
Emitted when `slashingContract` 'freezes' the `slashedOperator`.

*The `slashingContract` must have permission to slash the `slashedOperator`, i.e. `canSlash(slasherOperator, slashingContract)` must return 'true'.*


```solidity
event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract);
```

### FrozenStatusReset
Emitted when `previouslySlashedAddress` is 'unfrozen', allowing them to again move deposited funds within EigenLayer.


```solidity
event FrozenStatusReset(address indexed previouslySlashedAddress);
```

