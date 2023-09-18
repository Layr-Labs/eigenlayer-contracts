## DelegationManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](#TODO) | Singleton | Transparent proxy |

The `DelegationManager` sits between the EigenPod and Strategy subsystems to keep track of shares allocated to Operators as Stakers delegate/undelegate to them.

### Operators

Operators interact with the following functions:

#### `registerAsOperator`

Registers the caller as an Operator in EigenLayer. The new Operator provides the `OperatorDetails`, a struct containing:
* `address earningsReceiver`: the address that will receive earnings as the Operator provides services to AVSs *(currently unused)*
* `address delegationApprover`: if set, this address must sign and approve new delegation from Stakers to this Operator *(optional)*
* `uint32 stakerOptOutWindowBlocks`: the minimum delay (in blocks) between beginning and completing registration for an AVS. *(currently unused)*

**Effects**: `registerAsOperator` cements the Operator's `OperatorDetails`, and self-delegates the Operator to themselves - permanently marking the caller as an Operator.

They cannot "deregister" as an Operator - however, they can exit the system by withdrawing their funds via the `EigenPodManager` or `StrategyManager`.

**Requirements**:
* Caller MUST NOT already be an Operator
* Caller MUST NOT already be delegated to an Operator
* `earningsReceiver != address(0)`
* `stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`: (~15 days)
* Unpaused if not in pause status: `PAUSED_NEW_DELEGATION`

#### `modifyOperatorDetails`

Allows an Operator to update their stored `OperatorDetails`.

**Requirements**:
* Caller MUST already be an Operator
* `new earningsReceiver != address(0)`
* `new stakerOptOutWindowBlocks >= old stakerOptOutWindowBlocks`
* `new stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`

#### `updateOperatorMetadataURI`

Allows an Operator to emit an `OperatorMetadataURIUpdated` event. No other state changes occur.

**Requirements**:
* Caller MUST already be an Operator

#### `forceUndelegation`

### Stakers

Stakers interact with the following functions:

#### `delegateTo`

Allows the caller to delegate their stake to an Operator.

**Effects**: `delegateTo` delegates a Staker's assets to an Operator, increasing the Operator's shares in each strategy the caller has assets in.

**Requirements**:
* The `operator` MUST already be an Operator
* If the `operator` has a `delegationApprover`, the caller MUST provide a valid `approverSignatureAndExpiry` and `approverSalt`
* The caller MUST NOT already be delegated to an Operator
* Unpaused if not in pause status: `PAUSED_NEW_DELEGATION`

#### `delegateToBySignature`

Allows a Staker to delegate to an Operator by way of signature. This function can be called by three different parties:
* If the Staker calls this method, they need to submit both the `stakerSignatureAndExpiry` AND `approverSignatureAndExpiry`
* If the Operator calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the Operator's `delegationApprover` calls this method, they need to submit only the `stakerSignatureAndExpiry`

**Effects**: See `delegateTo` above.

**Requirements**:
* See `delegateTo` above
    * If caller is either the Operator's `delegationApprover` or the Operator, the `approverSignatureAndExpiry` and `approverSalt` can be empty
* `stakerSignatureAndExpiry` MUST be a valid, unexpired signature over the correct hash and nonce

### Other

The Strategy and EigenPod subsystems may call the following functions:

#### `undelegate`

#### `increaseDelegatedShares`

#### `decreaseDelegatedShares`
