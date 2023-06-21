# IDelegationManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IDelegationManager.sol)

**Author:**
Layr Labs, Inc.

This is the contract for delegation in EigenLayer. The main functionalities of this contract are
- enabling anyone to register as an operator in EigenLayer
- allowing new operators to provide a DelegationTerms-type contract, which may mediate their interactions with stakers who delegate to them
- enabling any staker to delegate its stake to the operator of its choice
- enabling a staker to undelegate its assets from an operator (performed as part of the withdrawal process, initiated through the StrategyManager)


## Functions
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

Undelegates `staker` from the operator who they are delegated to.

Callable only by the StrategyManager

*Should only ever be called in the event that the `staker` has no active deposits in EigenLayer.*


```solidity
function undelegate(address staker) external;
```

### delegatedTo

returns the address of the operator that `staker` is delegated to.


```solidity
function delegatedTo(address staker) external view returns (address);
```

### delegationTerms

returns the DelegationTerms of the `operator`, which may mediate their interactions with stakers who delegate to them.


```solidity
function delegationTerms(address operator) external view returns (IDelegationTerms);
```

### operatorShares

returns the total number of shares in `strategy` that are delegated to `operator`.


```solidity
function operatorShares(address operator, IStrategy strategy) external view returns (uint256);
```

### increaseDelegatedShares

Increases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker has further deposits into EigenLayer

*Callable only by the StrategyManager*


```solidity
function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares) external;
```

### decreaseDelegatedShares

Decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`, typically called when the staker withdraws from EigenLayer

*Callable only by the StrategyManager*


```solidity
function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares) external;
```

### isDelegated

Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.


```solidity
function isDelegated(address staker) external view returns (bool);
```

### isNotDelegated

Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.


```solidity
function isNotDelegated(address staker) external view returns (bool);
```

### isOperator

Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.


```solidity
function isOperator(address operator) external view returns (bool);
```

