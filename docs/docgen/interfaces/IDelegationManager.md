# Solidity API

## IDelegationManager

This is the contract for delegation in EigenLayer. The main functionalities of this contract are
- enabling anyone to register as an operator in EigenLayer
- allowing new operators to provide a DelegationTerms-type contract, which may mediate their interactions with stakers who delegate to them
- enabling any staker to delegate its stake to the operator of its choice
- enabling a staker to undelegate its assets from an operator (performed as part of the withdrawal process, initiated through the StrategyManager)

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

### delegatedTo

```solidity
function delegatedTo(address staker) external view returns (address)
```

returns the address of the operator that `staker` is delegated to.

### delegationTerms

```solidity
function delegationTerms(address operator) external view returns (contract IDelegationTerms)
```

returns the DelegationTerms of the `operator`, which may mediate their interactions with stakers who delegate to them.

### operatorShares

```solidity
function operatorShares(address operator, contract IStrategy strategy) external view returns (uint256)
```

returns the total number of shares in `strategy` that are delegated to `operator`.

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

### isDelegated

```solidity
function isDelegated(address staker) external view returns (bool)
```

Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.

### isNotDelegated

```solidity
function isNotDelegated(address staker) external view returns (bool)
```

Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.

### isOperator

```solidity
function isOperator(address operator) external view returns (bool)
```

Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.

