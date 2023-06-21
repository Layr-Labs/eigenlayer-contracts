# DelegationManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/core/DelegationManager.sol)

**Inherits:**
Initializable, OwnableUpgradeable, [Pausable](/docs/docgen/src/src/contracts/permissions/Pausable.sol/contract.Pausable.md), [DelegationManagerStorage](/docs/docgen/src/src/contracts/core/DelegationManagerStorage.sol/abstract.DelegationManagerStorage.md)

**Author:**
Layr Labs, Inc.

This is the contract for delegation in EigenLayer. The main functionalities of this contract are
- enabling anyone to register as an operator in EigenLayer
- allowing new operators to provide a DelegationTerms-type contract, which may mediate their interactions with stakers who delegate to them
- enabling any staker to delegate its stake to the operator of its choice
- enabling a staker to undelegate its assets from an operator (performed as part of the withdrawal process, initiated through the StrategyManager)


## State Variables
### PAUSED_NEW_DELEGATION

```solidity
uint8 internal constant PAUSED_NEW_DELEGATION = 0;
```


### ERC1271_MAGICVALUE

```solidity
bytes4 internal constant ERC1271_MAGICVALUE = 0x1626ba7e;
```


### ORIGINAL_CHAIN_ID

```solidity
uint256 immutable ORIGINAL_CHAIN_ID;
```


## Functions
### onlyStrategyManager

Simple permission for functions that are only callable by the StrategyManager contract.


```solidity
modifier onlyStrategyManager();
```

### constructor


```solidity
constructor(IStrategyManager _strategyManager, ISlasher _slasher)
    DelegationManagerStorage(_strategyManager, _slasher);
```

### initialize


```solidity
function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus)
    external
    initializer;
```

### registerAsOperator

This will be called by an operator to register itself as an operator that stakers can choose to delegate to.

*An operator can set `dt` equal to their own address (or another EOA address), in the event that they want to split payments
in a more 'trustful' manner.*

*In the present design, once set, there is no way for an operator to ever modify the address of their DelegationTerms contract.*


```solidity
function registerAsOperator(IDelegationTerms dt) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`dt`|`IDelegationTerms`|is the `DelegationTerms` contract that the operator has for those who delegate to them.|


### delegateTo

This will be called by a staker to delegate its assets to some operator.


```solidity
function delegateTo(address operator) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator to whom staker (msg.sender) is delegating its assets|


### delegateToBySignature

Delegates from `staker` to `operator`.

*requires that:
1) if `staker` is an EOA, then `signature` is valid ECDSA signature from `staker`, indicating their intention for this action
2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271*


```solidity
function delegateToBySignature(address staker, address operator, uint256 expiry, bytes memory signature) external;
```

### undelegate

check validity of signature:
1) if `staker` is an EOA, then `signature` must be a valid ECDSA signature from `staker`,
indicating their intention for this action
2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271

Undelegates `staker` from the operator who they are delegated to.

Callable only by the StrategyManager

*Should only ever be called in the event that the `staker` has no active deposits in EigenLayer.*


```solidity
function undelegate(address staker) external onlyStrategyManager;
```

### increaseDelegatedShares

Increases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker has further deposits into EigenLayer

*Callable only by the StrategyManager*


```solidity
function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares) external onlyStrategyManager;
```

### decreaseDelegatedShares

Decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`, typically called when the staker withdraws from EigenLayer

*Callable only by the StrategyManager*


```solidity
function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares)
    external
    onlyStrategyManager;
```

### _delegationReceivedHook

Makes a low-level call to `dt.onDelegationReceived(staker, strategies, shares)`, ignoring reverts and with a gas budget
equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).

**If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
`returnData` is *only the first 32 bytes* returned by the call to `dt`.*


```solidity
function _delegationReceivedHook(
    IDelegationTerms dt,
    address staker,
    IStrategy[] memory strategies,
    uint256[] memory shares
) internal;
```

### _delegationWithdrawnHook

We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.

Makes a low-level call to `dt.onDelegationWithdrawn(staker, strategies, shares)`, ignoring reverts and with a gas budget
equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).

**If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
`returnData` is *only the first 32 bytes* returned by the call to `dt`.*


```solidity
function _delegationWithdrawnHook(
    IDelegationTerms dt,
    address staker,
    IStrategy[] memory strategies,
    uint256[] memory shares
) internal;
```

### _delegate

We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.

Internal function implementing the delegation *from* `staker` *to* `operator`.

*Ensures that the operator has registered as a delegate (`address(dt) != address(0)`), verifies that `staker` is not already
delegated, and records the new delegation.*


```solidity
function _delegate(address staker, address operator) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|The address to delegate *from* -- this address is delegating control of its own assets.|
|`operator`|`address`|The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services|


### isDelegated

Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.


```solidity
function isDelegated(address staker) public view returns (bool);
```

### isNotDelegated

Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.


```solidity
function isNotDelegated(address staker) public view returns (bool);
```

### isOperator

Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.


```solidity
function isOperator(address operator) public view returns (bool);
```

## Events
### OnDelegationReceivedCallFailure
*Emitted when a low-level call to `delegationTerms.onDelegationReceived` fails, returning `returnData`*


```solidity
event OnDelegationReceivedCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
```

### OnDelegationWithdrawnCallFailure
*Emitted when a low-level call to `delegationTerms.onDelegationWithdrawn` fails, returning `returnData`*


```solidity
event OnDelegationWithdrawnCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
```

### RegisterAsOperator
*Emitted when an entity registers itself as an operator in the DelegationManager*


```solidity
event RegisterAsOperator(address indexed operator, IDelegationTerms indexed delegationTerms);
```

