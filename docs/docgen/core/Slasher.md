# Solidity API

## Slasher

This contract specifies details on slashing. The functionalities are:
- adding contracts who have permission to perform slashing,
- revoking permission for slashing from specified contracts,
- tracking historic stake updates to ensure that withdrawals can only be completed once no middlewares have slashing rights
over the funds being withdrawn

### HEAD

```solidity
uint256 HEAD
```

### PAUSED_OPT_INTO_SLASHING

```solidity
uint8 PAUSED_OPT_INTO_SLASHING
```

### PAUSED_FIRST_STAKE_UPDATE

```solidity
uint8 PAUSED_FIRST_STAKE_UPDATE
```

### PAUSED_NEW_FREEZING

```solidity
uint8 PAUSED_NEW_FREEZING
```

### strategyManager

```solidity
contract IStrategyManager strategyManager
```

The central StrategyManager contract of EigenLayer

### delegation

```solidity
contract IDelegationManager delegation
```

The DelegationManager contract of EigenLayer

### _whitelistedContractDetails

```solidity
mapping(address => mapping(address => struct ISlasher.MiddlewareDetails)) _whitelistedContractDetails
```

### frozenStatus

```solidity
mapping(address => bool) frozenStatus
```

### MAX_CAN_SLASH_UNTIL

```solidity
uint32 MAX_CAN_SLASH_UNTIL
```

### _operatorToWhitelistedContractsByUpdate

```solidity
mapping(address => struct StructuredLinkedList.List) _operatorToWhitelistedContractsByUpdate
```

operator => a linked list of the addresses of the whitelisted middleware with permission to slash the operator, i.e. which  
the operator is serving. Sorted by the block at which they were last updated (content of updates below) in ascending order.
This means the 'HEAD' (i.e. start) of the linked list will have the stalest 'updateBlock' value.

### _operatorToMiddlewareTimes

```solidity
mapping(address => struct ISlasher.MiddlewareTimes[]) _operatorToMiddlewareTimes
```

operator => 
 [
     (
         the least recent update block of all of the middlewares it's serving/served, 
         latest time that the stake bonded at that update needed to serve until
     )
 ]

### MiddlewareTimesAdded

```solidity
event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
```

Emitted when a middleware times is added to `operator`'s array.

### OptedIntoSlashing

```solidity
event OptedIntoSlashing(address operator, address contractAddress)
```

Emitted when `operator` begins to allow `contractAddress` to slash them.

### SlashingAbilityRevoked

```solidity
event SlashingAbilityRevoked(address operator, address contractAddress, uint32 contractCanSlashOperatorUntilBlock)
```

Emitted when `contractAddress` signals that it will no longer be able to slash `operator` after the `contractCanSlashOperatorUntilBlock`.

### OperatorFrozen

```solidity
event OperatorFrozen(address slashedOperator, address slashingContract)
```

Emitted when `slashingContract` 'freezes' the `slashedOperator`.

_The `slashingContract` must have permission to slash the `slashedOperator`, i.e. `canSlash(slasherOperator, slashingContract)` must return 'true'._

### FrozenStatusReset

```solidity
event FrozenStatusReset(address previouslySlashedAddress)
```

Emitted when `previouslySlashedAddress` is 'unfrozen', allowing them to again move deposited funds within EigenLayer.

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract IDelegationManager _delegation) public
```

### onlyRegisteredForService

```solidity
modifier onlyRegisteredForService(address operator)
```

Ensures that the operator has opted into slashing by the caller, and that the caller has never revoked its slashing ability.

### initialize

```solidity
function initialize(address initialOwner, contract IPauserRegistry _pauserRegistry, uint256 initialPausedStatus) external
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
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external
```

this function is a called by middlewares during an operator's registration to make sure the operator's stake at registration 
        is slashable until serveUntilBlock

_adds the middleware's slashing contract to the operator's linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| serveUntilBlock | uint32 | the block until which the operator's stake at the current block is slashable |

### recordStakeUpdate

```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) external
```

this function is a called by middlewares during a stake update for an operator (perhaps to free pending withdrawals)
        to make sure the operator's stake at updateBlock is slashable until serveUntilBlock

_insertAfter should be calculated offchain before making the transaction that calls this. this is subject to race conditions, 
     but it is anticipated to be rare and not detrimental._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| updateBlock | uint32 | the block for which the stake update is being recorded |
| serveUntilBlock | uint32 | the block until which the operator's stake at updateBlock is slashable |
| insertAfter | uint256 | the element of the operators linked list that the currently updating middleware should be inserted after |

### recordLastStakeUpdateAndRevokeSlashingAbility

```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external
```

this function is a called by middlewares during an operator's deregistration to make sure the operator's stake at deregistration 
        is slashable until serveUntilBlock

_removes the middleware's slashing contract to the operator's linked list and revokes the middleware's (i.e. caller's) ability to
slash `operator` once `serveUntilBlock` is reached_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the operator whose stake update is being recorded |
| serveUntilBlock | uint32 | the block until which the operator's stake at the current block is slashable |

### contractCanSlashOperatorUntilBlock

```solidity
function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) external view returns (uint32)
```

Returns the block until which `serviceContract` is allowed to slash the `operator`.

### latestUpdateBlock

```solidity
function latestUpdateBlock(address operator, address serviceContract) external view returns (uint32)
```

Returns the block at which the `serviceContract` last updated its view of the `operator`'s stake

### whitelistedContractDetails

```solidity
function whitelistedContractDetails(address operator, address serviceContract) external view returns (struct ISlasher.MiddlewareDetails)
```

### isFrozen

```solidity
function isFrozen(address staker) external view returns (bool)
```

Used to determine whether `staker` is actively 'frozen'. If a staker is frozen, then they are potentially subject to
slashing of their funds, and cannot cannot deposit or withdraw from the strategyManager until the slashing process is completed
and the staker's status is reset (to 'unfrozen').

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| staker | address | The staker of interest. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | Returns 'true' if `staker` themselves has their status set to frozen, OR if the staker is delegated to an operator who has their status set to frozen. Otherwise returns 'false'. |

### canSlash

```solidity
function canSlash(address toBeSlashed, address slashingContract) public view returns (bool)
```

Returns true if `slashingContract` is currently allowed to slash `toBeSlashed`.

### canWithdraw

```solidity
function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) external view returns (bool)
```

Returns 'true' if `operator` can currently complete a withdrawal started at the `withdrawalStartBlock`, with `middlewareTimesIndex` used
to specify the index of a `MiddlewareTimes` struct in the operator's list (i.e. an index in `_operatorToMiddlewareTimes[operator]`). The specified
struct is consulted as proof of the `operator`'s ability (or lack thereof) to complete the withdrawal.
This function will return 'false' if the operator cannot currently complete a withdrawal started at the `withdrawalStartBlock`, *or* in the event
that an incorrect `middlewareTimesIndex` is supplied, even if one or more correct inputs exist.

_The correct `middlewareTimesIndex` input should be computable off-chain._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | Either the operator who queued the withdrawal themselves, or if the withdrawing party is a staker who delegated to an operator, this address is the operator *who the staker was delegated to* at the time of the `withdrawalStartBlock`. |
| withdrawalStartBlock | uint32 | The block number at which the withdrawal was initiated. |
| middlewareTimesIndex | uint256 | Indicates an index in `_operatorToMiddlewareTimes[operator]` to consult as proof of the `operator`'s ability to withdraw |

### operatorToMiddlewareTimes

```solidity
function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) external view returns (struct ISlasher.MiddlewareTimes)
```

Getter function for fetching `_operatorToMiddlewareTimes[operator][arrayIndex]`.

### middlewareTimesLength

```solidity
function middlewareTimesLength(address operator) external view returns (uint256)
```

Getter function for fetching `_operatorToMiddlewareTimes[operator].length`.

### getMiddlewareTimesIndexBlock

```solidity
function getMiddlewareTimesIndexBlock(address operator, uint32 index) external view returns (uint32)
```

Getter function for fetching `_operatorToMiddlewareTimes[operator][index].stalestUpdateBlock`.

### getMiddlewareTimesIndexServeUntilBlock

```solidity
function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) external view returns (uint32)
```

Getter function for fetching `_operatorToMiddlewareTimes[operator][index].latestServeUntilBlock`.

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

### getCorrectValueForInsertAfter

```solidity
function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) public view returns (uint256)
```

A search routine for finding the correct input value of `insertAfter` to `recordStakeUpdate` / `_updateMiddlewareList`.

_Used within this contract only as a fallback in the case when an incorrect value of `insertAfter` is supplied as an input to `_updateMiddlewareList`.
The return value should *either* be 'HEAD' (i.e. zero) in the event that the node being inserted in the linked list has an `updateBlock`
that is less than the HEAD of the list, *or* the return value should specify the last `node` in the linked list for which
`_whitelistedContractDetails[operator][node].latestUpdateBlock <= updateBlock`,
i.e. the node such that the *next* node either doesn't exist,
OR
`_whitelistedContractDetails[operator][nextNode].latestUpdateBlock > updateBlock`._

### getPreviousWhitelistedContractByUpdate

```solidity
function getPreviousWhitelistedContractByUpdate(address operator, uint256 node) external view returns (bool, uint256)
```

gets the node previous to the given node in the operators middleware update linked list

_used in offchain libs for updating stakes_

### _optIntoSlashing

```solidity
function _optIntoSlashing(address operator, address contractAddress) internal
```

### _revokeSlashingAbility

```solidity
function _revokeSlashingAbility(address operator, address contractAddress, uint32 serveUntilBlock) internal
```

### _freezeOperator

```solidity
function _freezeOperator(address toBeFrozen, address slashingContract) internal
```

### _resetFrozenStatus

```solidity
function _resetFrozenStatus(address previouslySlashedAddress) internal
```

### _recordUpdateAndAddToMiddlewareTimes

```solidity
function _recordUpdateAndAddToMiddlewareTimes(address operator, uint32 updateBlock, uint32 serveUntilBlock) internal
```

records the most recent updateBlock for the currently updating middleware and appends an entry to the operator's list of 
        MiddlewareTimes if relavent information has updated

_this function is only called during externally called stake updates by middleware contracts that can slash operator_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | the entity whose stake update is being recorded |
| updateBlock | uint32 | the block number for which the currently updating middleware is updating the serveUntilBlock for |
| serveUntilBlock | uint32 | the block until which the operator's stake at updateBlock is slashable |

### _updateMiddlewareList

```solidity
function _updateMiddlewareList(address operator, uint32 updateBlock, uint256 insertAfter) internal
```

A routine for updating the `operator`'s linked list of middlewares, inside `recordStakeUpdate`.

### _addressToUint

```solidity
function _addressToUint(address addr) internal pure returns (uint256)
```

### _uintToAddress

```solidity
function _uintToAddress(uint256 x) internal pure returns (address)
```

### __gap

```solidity
uint256[46] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

