## DelegationManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](#TODO) | Singleton | Transparent proxy |

The primary functions of the `DelegationManager` are (i) to allow Stakers to delegate to Operators and (ii) to keep an up-to-date record of the number of shares each Operator has been delegated for each strategy.

Whereas the `EigenPodManager` and `StrategyManager` perform accounting for individual Stakers according to their native ETH or LST holdings respectively, the `DelegationManager` sits between these two contracts and tracks these accounting changes according to the Operators each Staker has delegated to. This means that each time a Staker deposits or withdraws from either the `EigenPodManager` or `StrategyManager`, the `DelegationManager` is called to record this update to the Staker's delegated Operator (if they have one).

Could add:
* Pausability ("Pausers" section), or even Upgradability
* High-level flows (what high-level flows this contract has entry points for)

### Operators

Operators interact with the following functions:

#### `registerAsOperator`

```solidity
function registerAsOperator(OperatorDetails calldata registeringOperatorDetails, string calldata metadataURI) external
```

Registers the caller as an Operator in EigenLayer. The new Operator provides the `OperatorDetails`, a struct containing:
* `address earningsReceiver`: the address that will receive earnings as the Operator provides services to AVSs *(currently unused)*
* `address delegationApprover`: if set, this address must sign and approve new delegation from Stakers to this Operator *(optional)*
* `uint32 stakerOptOutWindowBlocks`: the minimum delay (in blocks) between beginning and completing registration for an AVS. *(currently unused)*

`registerAsOperator` cements the Operator's `OperatorDetails`, and self-delegates the Operator to themselves - permanently marking the caller as an Operator. They cannot "deregister" as an Operator - however, they can exit the system by withdrawing their funds via the `EigenPodManager` or `StrategyManager`.

**Effects**:
* Sets `OperatorDetails` for the Operator in question
* If the Operator has deposited into the `EigenPodManager` and is not in undelegation limbo, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Operator holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

**Requirements**:
* Caller MUST NOT already be an Operator
* Caller MUST NOT already be delegated to an Operator
* `earningsReceiver != address(0)`
* `stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`: (~15 days)
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`

*Unimplemented as of M2*:
* `require(!slasher.isFrozen(operator))` is currently a no-op

#### `modifyOperatorDetails`

```solidity
function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external
```

Allows an Operator to update their stored `OperatorDetails`.

**Requirements**:
* Caller MUST already be an Operator
* `new earningsReceiver != address(0)`
* `new stakerOptOutWindowBlocks >= old stakerOptOutWindowBlocks`
* `new stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`

#### `updateOperatorMetadataURI`

```solidity
function updateOperatorMetadataURI(string calldata metadataURI) external
```

Allows an Operator to emit an `OperatorMetadataURIUpdated` event. No other state changes occur.

**Requirements**:
* Caller MUST already be an Operator

#### `forceUndelegation`

```solidity
function forceUndelegation(address staker) external returns (bytes32)
```

Allows an Operator or its `delegationApprover` to force a Staker to undelegate from them. This can be useful in case the Operator wants to use its `delegationApprover` to manually accept Stakers, rather than allowing all delegation by default.

**Effects**: Invokes methods on both the `EigenPodManager` and `StrategyManager`:
* `EigenPodManager.forceIntoUndelegationLimbo`
* `StrategyManager.forceTotalWithdrawal`

If the Staker has shares in these contracts, each contract will call back into `DelegationManager.decreaseDelegatedShares`, decreasing the shares allocated to the Operator for the strategy in question. Depending on what shares the Staker has, one of the two calls will also call back into `DelegationManager.undelegate`, which undelegates the Staker from the Operator.

**Requirements**:
* Caller MUST be either the Staker's Operator, or that Operator's `delegationApprover`
* Staker being undelegated MUST NOT be an Operator
* From `EigenPodManager.forceIntoUndelegationLimbo`:
    * Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`
* From `StrategyManager.forceTotalWithdrawal`:
    * Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`

### Stakers

Stakers interact with the following functions:

#### `delegateTo`

```solidity
function delegateTo(address operator, SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) external
```

Allows the caller (a Staker) to delegate ALL their shares to an Operator (delegation is all-or-nothing). For each strategy the Staker has shares in, the `DelegationManager` will update the Operator's corresponding delegated share amounts.

**Effects**:
* Records the Staker as being delegated to the Operator
* If the Staker has deposited into the `EigenPodManager` and is not in undelegation limbo, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Staker holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

**Requirements**:
* The `operator` MUST already be an Operator
* If the `operator` has a `delegationApprover`, the caller MUST provide a valid `approverSignatureAndExpiry` and `approverSalt`
* The caller MUST NOT already be delegated to an Operator
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`

*Unimplemented as of M2*:
* `require(!slasher.isFrozen(operator))` is currently a no-op

#### `delegateToBySignature`

```solidity
function delegateToBySignature(
    address staker,
    address operator,
    SignatureWithExpiry memory stakerSignatureAndExpiry,
    SignatureWithExpiry memory approverSignatureAndExpiry,
    bytes32 approverSalt
) external
```

Allows a Staker to delegate to an Operator by way of signature. This function can be called by three different parties:
* If the Staker calls this method, they need to submit both the `stakerSignatureAndExpiry` AND `approverSignatureAndExpiry`
* If the Operator calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the Operator's `delegationApprover` calls this method, they need to submit only the `stakerSignatureAndExpiry`

**Effects**: See `delegateTo` above.

**Requirements**: See `delegateTo` above. Additionally:
* If caller is either the Operator's `delegationApprover` or the Operator, the `approverSignatureAndExpiry` and `approverSalt` can be empty
* `stakerSignatureAndExpiry` MUST be a valid, unexpired signature over the correct hash and nonce

*Unimplemented as of M2*:
* `require(!slasher.isFrozen(operator))` is currently a no-op

### Other

These functions are not directly callable. Instead, the Strategy and EigenPod subsystems may call the following functions as part of certain operations:

#### `undelegate`

```solidity
function undelegate(address staker) external onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker undelegates from an Operator.

**Entry Points**: This method may be called as a result of the following top-level function calls:
* `DelegationManager.forceUndelegation`
* `StrategyManager.undelegate`

**Effects**: If the Staker was delegated to an Operator, this function undelegates them. This is simply an update to the `delegatedTo` mapping; there are no accounting changes in this method.

**Requirements**:
* Staker MUST NOT be an Operator
* Caller MUST be either the `StrategyManager` or `EigenPodManager`

#### `increaseDelegatedShares`

```solidity
function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker's shares for one or more strategies increase. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the increase.

**Entry Points**: This method may be called as a result of the following top-level function calls:
* `StrategyManager.depositIntoStrategy`
* `StrategyManager.depositIntoStrategyWithSignature`

**Effects**: If the Staker in question is delegated to an Operator, the Operator's shares for the `strategy` are increased.
* This method is a no-op if the Staker is not delegated to an Operator.

**Requirements**:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`

#### `decreaseDelegatedShares`

```solidity
function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker's shares for one or more strategies decrease. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the decrease.

**Entry Points**: This method may be called as a result of the following top-level function calls:
* `DelegationManager.forceUndelegation`
* TODO

**Effects**: If the Staker in question is delegated to an Operator, the Operator's shares for each of the `strategies` are decreased (by the corresponding amount in the `shares` array).
* This method is a no-op if the Staker is not delegated an an Operator.

**Requirements**:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`