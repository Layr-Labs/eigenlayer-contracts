# Solidity API

## DelegationManager

This is the contract for delegation in EigenLayer. The main functionalities of this contract are
- enabling anyone to register as an operator in EigenLayer
- allowing new operators to provide a DelegationTerms-type contract, which may mediate their interactions with stakers who delegate to them
- enabling any staker to delegate its stake to the operator of its choice
- enabling a staker to undelegate its assets from an operator (performed as part of the withdrawal process, initiated through the StrategyManager)

### PAUSED_NEW_DELEGATION

```solidity
uint8 PAUSED_NEW_DELEGATION
```

### ERC1271_MAGICVALUE

```solidity
bytes4 ERC1271_MAGICVALUE
```

### onlyStrategyManager

```solidity
modifier onlyStrategyManager()
```

Simple permission for functions that are only callable by the StrategyManager contract.

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract ISlasher _slasher) public
```

### OnDelegationReceivedCallFailure

```solidity
event OnDelegationReceivedCallFailure(contract IDelegationTerms delegationTerms, bytes32 returnData)
```

_Emitted when a low-level call to `delegationTerms.onDelegationReceived` fails, returning `returnData`_

### OnDelegationWithdrawnCallFailure

```solidity
event OnDelegationWithdrawnCallFailure(contract IDelegationTerms delegationTerms, bytes32 returnData)
```

_Emitted when a low-level call to `delegationTerms.onDelegationWithdrawn` fails, returning `returnData`_

### initialize

```solidity
function initialize(address initialOwner, contract IPauserRegistry _pauserRegistry, uint256 initialPausedStatus) external
```

### registerAsOperator

```solidity
function registerAsOperator(contract IDelegationTerms dt) external
```

This will be called by an operator to register itself as an operator that stakers can choose to delegate to.

_An operator can set `dt` equal to their own address (or another EOA address), in the event that they want to split payments
in a more 'trustful' manner.
In the present design, once set, there is no way for an operator to ever modify the address of their DelegationTerms contract._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| dt | contract IDelegationTerms | is the `DelegationTerms` contract that the operator has for those who delegate to them. |

### delegateTo

```solidity
function delegateTo(address operator) external
```

@notice This will be called by a staker to delegate its assets to some operator.
 @param operator is the operator to whom staker (msg.sender) is delegating its assets

### delegateToBySignature

```solidity
function delegateToBySignature(address staker, address operator, uint256 expiry, bytes signature) external
```

Delegates from `staker` to `operator`.

_requires that:
1) if `staker` is an EOA, then `signature` is valid ECSDA signature from `staker`, indicating their intention for this action
2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271_

### undelegate

```solidity
function undelegate(address staker) external
```

Undelegates `staker` from the operator who they are delegated to.
Callable only by the StrategyManager

_Should only ever be called in the event that the `staker` has no active deposits in EigenLayer._

### increaseDelegatedShares

```solidity
function increaseDelegatedShares(address staker, contract IStrategy strategy, uint256 shares) external
```

Increases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker has further deposits into EigenLayer

_Callable only by the StrategyManager_

### decreaseDelegatedShares

```solidity
function decreaseDelegatedShares(address staker, contract IStrategy[] strategies, uint256[] shares) external
```

Decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`, typically called when the staker withdraws from EigenLayer

_Callable only by the StrategyManager_

### _delegationReceivedHook

```solidity
function _delegationReceivedHook(contract IDelegationTerms dt, address staker, contract IStrategy[] strategies, uint256[] shares) internal
```

Makes a low-level call to `dt.onDelegationReceived(staker, strategies, shares)`, ignoring reverts and with a gas budget 
equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).

_*If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
`returnData` is *only the first 32 bytes* returned by the call to `dt`._

### _delegationWithdrawnHook

```solidity
function _delegationWithdrawnHook(contract IDelegationTerms dt, address staker, contract IStrategy[] strategies, uint256[] shares) internal
```

Makes a low-level call to `dt.onDelegationWithdrawn(staker, strategies, shares)`, ignoring reverts and with a gas budget 
equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).

_*If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
`returnData` is *only the first 32 bytes* returned by the call to `dt`._

### _delegate

```solidity
function _delegate(address staker, address operator) internal
```

Internal function implementing the delegation *from* `staker` *to* `operator`.

_Ensures that the operator has registered as a delegate (`address(dt) != address(0)`), verifies that `staker` is not already
delegated, and records the new delegation._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| staker | address | The address to delegate *from* -- this address is delegating control of its own assets. |
| operator | address | The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services |

### isDelegated

```solidity
function isDelegated(address staker) public view returns (bool)
```

Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.

### isNotDelegated

```solidity
function isNotDelegated(address staker) public view returns (bool)
```

Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.

### isOperator

```solidity
function isOperator(address operator) public view returns (bool)
```

Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.

