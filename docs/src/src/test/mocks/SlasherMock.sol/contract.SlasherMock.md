# SlasherMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/SlasherMock.sol)

**Inherits:**
[ISlasher](/src/contracts/interfaces/ISlasher.sol/interface.ISlasher.md), Test


## State Variables
### isFrozen

```solidity
mapping(address => bool) public isFrozen;
```


### _canWithdraw

```solidity
bool public _canWithdraw = true;
```


## Functions
### setCanWithdrawResponse


```solidity
function setCanWithdrawResponse(bool response) external;
```

### setOperatorFrozenStatus


```solidity
function setOperatorFrozenStatus(address operator, bool status) external;
```

### freezeOperator


```solidity
function freezeOperator(address toBeFrozen) external;
```

### optIntoSlashing


```solidity
function optIntoSlashing(address contractAddress) external;
```

### resetFrozenStatus


```solidity
function resetFrozenStatus(address[] calldata frozenAddresses) external;
```

### recordFirstStakeUpdate


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external;
```

### recordStakeUpdate


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter)
    external;
```

### recordLastStakeUpdateAndRevokeSlashingAbility


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external;
```

### canSlash

Returns true if `slashingContract` is currently allowed to slash `toBeSlashed`.


```solidity
function canSlash(address toBeSlashed, address slashingContract) external view returns (bool);
```

### contractCanSlashOperatorUntilBlock

Returns the UTC timestamp until which `serviceContract` is allowed to slash the `operator`.


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


```solidity
function canWithdraw(address, uint32, uint256) external view returns (bool);
```

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

### getMiddlewareTimesIndexStalestUpdateBlock

Getter function for fetching `operatorToMiddlewareTimes[operator][index].stalestUpdateBlock`.


```solidity
function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) external view returns (uint32);
```

### getMiddlewareTimesIndexServeUntilBlock

Getter function for fetching `operatorToMiddlewareTimes[operator][index].latestServeUntilBlock`.


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

