## DelegationManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](../../src/contracts/core/DelegationManager.sol) | Singleton | Transparent proxy |

The primary functions of the `DelegationManager` are (i) to allow Stakers to delegate to Operators and (ii) to keep an up-to-date record of the number of shares each Operator has been delegated for each strategy.

Whereas the `EigenPodManager` and `StrategyManager` perform accounting for individual Stakers according to their native ETH or LST holdings respectively, the `DelegationManager` sits between these two contracts and tracks these accounting changes according to the Operators each Staker has delegated to. This means that each time a Staker's balance changes in either the `EigenPodManager` or `StrategyManager`, the `DelegationManager` is called to record this update to the Staker's delegated Operator (if they have one).

*Important state variables*:
* `mapping(address => address) public delegatedTo`: Staker => Operator.
    * If a Staker is not delegated to anyone, `delegatedTo` is unset.
    * Operators are delegated to themselves - `delegatedTo[operator] == operator`
* `mapping(address => mapping(IStrategy => uint256)) public operatorShares`: Tracks the current balance of shares an Operator is delegated according to each strategy. Updated by both the `StrategyManager` and `EigenPodManager` when a Staker's delegatable balance changes.
    * Because Operators are delegated to themselves, an Operator's own restaked assets are reflected in these balances.
    * A similar mapping exists in the `StrategyManager`, but the `DelegationManager` additionally tracks beacon chain ETH delegated via the `EigenPodManager`. The "beacon chain ETH" strategy gets its own special address for this mapping: `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0`.

*Helpful definitions*:
* `isDelegated(address staker) -> (bool)`
    * True if `delegatedTo[staker] != address(0)`
* `isOperator(address operator) -> (bool)` 
    * True if `_operatorDetails[operator].earningsReceiver != address(0)`

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

*Effects*:
* Sets `OperatorDetails` for the Operator in question
* Delegates the Operator to itself
* If the Operator has deposited into the `EigenPodManager` and is not in undelegation limbo, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Operator holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

*Requirements*:
* Caller MUST NOT already be an Operator
* Caller MUST NOT already be delegated to an Operator
* `earningsReceiver != address(0)`
* `stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`: (~180 days)
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`

*As of M2*:
* `require(!slasher.isFrozen(operator))` is currently a no-op

#### `modifyOperatorDetails`

```solidity
function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external
```

Allows an Operator to update their stored `OperatorDetails`.

*Requirements*:
* Caller MUST already be an Operator
* `new earningsReceiver != address(0)`
* `new stakerOptOutWindowBlocks >= old stakerOptOutWindowBlocks`
* `new stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`

#### `updateOperatorMetadataURI`

```solidity
function updateOperatorMetadataURI(string calldata metadataURI) external
```

Allows an Operator to emit an `OperatorMetadataURIUpdated` event. No other state changes occur.

*Requirements*:
* Caller MUST already be an Operator

### Stakers

Stakers interact with the following functions:

#### `delegateTo`

```solidity
function delegateTo(address operator, SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) external
```

Allows the caller (a Staker) to delegate ALL their shares to an Operator (delegation is all-or-nothing). For each strategy the Staker has shares in, the `DelegationManager` will update the Operator's corresponding delegated share amounts.

*Effects*:
* Records the Staker as being delegated to the Operator
* If the Staker has deposited into the `EigenPodManager` and is not in undelegation limbo, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Staker holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

*Requirements*:
* The caller MUST NOT already be delegated to an Operator
* The `operator` MUST already be an Operator
* If the `operator` has a `delegationApprover`, the caller MUST provide a valid `approverSignatureAndExpiry` and `approverSalt`
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`

*As of M2*:
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
* If the Operator calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the Operator's `delegationApprover` calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the anyone else calls this method, they need to submit both the `stakerSignatureAndExpiry` AND `approverSignatureAndExpiry`

*Effects*: See `delegateTo` above.

*Requirements*: See `delegateTo` above. Additionally:
* If caller is either the Operator's `delegationApprover` or the Operator, the `approverSignatureAndExpiry` and `approverSalt` can be empty
* `stakerSignatureAndExpiry` MUST be a valid, unexpired signature over the correct hash and nonce

*As of M2*:
* `require(!slasher.isFrozen(operator))` is currently a no-op

### Other

#### `undelegate`

```solidity
function undelegate(address staker) external returns (bytes32 withdrawalRoot)
```

This method undelegates a Staker from an Operator, decreasing the Operator's shares in the strategies held by the Staker. This can be called by a Staker to undelegate themselves, or by a Staker's delegated Operator (or that Operator's `delegationApprover`).

If the Staker has active shares in the `EigenPodManager`, the Staker is placed into "undelegation limbo."

If the Staker has active shares in any strategy in the `StrategyManager`, this initiates a withdrawal of the Staker's shares.

*Effects*: 
* See [`EigenPodManager.forceIntoUndelegationLimbo`](./EigenPodManager.md#eigenpodmanagerforceintoundelegationlimbo)
* See [`StrategyManager.forceTotalWithdrawal`](./StrategyManager.md#forcetotalwithdrawal)
* Any shares held by the Staker in the `EigenPodManager` and `StrategyManager` are removed from the Operator's delegated shares.
* If the Staker was delegated to an Operator, this function undelegates them.

*Requirements*:
* Staker MUST exist and be delegated to someone
* Staker MUST NOT be an Operator
* Caller must be either the Staker, their Operator, or their Operator's `delegationApprover`
* Pause status MUST NOT be set: `PAUSED_UNDELEGATION`
* See [`EigenPodManager.forceIntoUndelegationLimbo`](./EigenPodManager.md#eigenpodmanagerforceintoundelegationlimbo)
* See [`StrategyManager.forceTotalWithdrawal`](./StrategyManager.md#forcetotalwithdrawal)

#### `increaseDelegatedShares`

```solidity
function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker's shares for one or more strategies increase. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the increase.

*Entry Points*: This method may be called as a result of the following top-level function calls:
* `StrategyManager.depositIntoStrategy`
* `StrategyManager.depositIntoStrategyWithSignature`
* `StrategyManager.completeQueuedWithdrawal`
* `EigenPodManager.exitUndelegationLimbo`
* `EigenPod.verifyWithdrawalCredentials`
* `EigenPod.verifyBalanceUpdate`
* `EigenPod.verifyAndProcessWithdrawals`

*Effects*: If the Staker in question is delegated to an Operator, the Operator's shares for the `strategy` are increased.
* This method is a no-op if the Staker is not delegated to an Operator.

*Requirements*:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`

#### `decreaseDelegatedShares`

```solidity
function decreaseDelegatedShares(address staker, IStrategy[] calldata strategies, uint256[] calldata shares)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker's shares for one or more strategies decrease. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the decrease.

*Entry Points*: This method may be called as a result of the following top-level function calls:
* `EigenPodManager.queueWithdrawal`
* `EigenPod.verifyBalanceUpdate`
* `EigenPod.verifyAndProcessWithdrawals`
* `StrategyManager.queueWithdrawal`

*Effects*: If the Staker in question is delegated to an Operator, the Operator's shares for each of the `strategies` are decreased (by the corresponding amount in the `shares` array).
* This method is a no-op if the Staker is not delegated an an Operator.

*Requirements*:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`