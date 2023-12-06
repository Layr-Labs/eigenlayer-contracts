## DelegationManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](../../src/contracts/core/DelegationManager.sol) | Singleton | Transparent proxy |

The primary functions of the `DelegationManager` are (i) to allow Stakers to delegate to Operators, (ii) allow Stakers to be undelegated from Operators, and (iii) handle withdrawals and withdrawal processing for shares in both the `StrategyManager` and `EigenPodManager`.

Whereas the `EigenPodManager` and `StrategyManager` perform accounting for individual Stakers according to their native ETH or LST holdings respectively, the `DelegationManager` sits between these two contracts and tracks these accounting changes according to the Operators each Staker has delegated to. 

This means that each time a Staker's balance changes in either the `EigenPodManager` or `StrategyManager`, the `DelegationManager` is called to record this update to the Staker's delegated Operator (if they have one). For example, if a Staker is delegated to an Operator and deposits into a strategy, the `StrategyManager` will call the `DelegationManager` to update the Operator's delegated shares for that strategy.

Additionally, whether a Staker is delegated to an Operator or not, the `DelegationManager` is how a Staker queues (and later completes) a withdrawal.

#### High-level Concepts

This document organizes methods according to the following themes (click each to be taken to the relevant section):
* [Becoming an Operator](#becoming-an-operator)
* [Delegating to an Operator](#delegating-to-an-operator)
* [Undelegating and Withdrawing](#undelegating-and-withdrawing)
* [Accounting](#accounting)

#### Important state variables

* `mapping(address => address) public delegatedTo`: Staker => Operator.
    * If a Staker is not delegated to anyone, `delegatedTo` is unset.
    * Operators are delegated to themselves - `delegatedTo[operator] == operator`
* `mapping(address => mapping(IStrategy => uint256)) public operatorShares`: Tracks the current balance of shares an Operator is delegated according to each strategy. Updated by both the `StrategyManager` and `EigenPodManager` when a Staker's delegatable balance changes.
    * Because Operators are delegated to themselves, an Operator's own restaked assets are reflected in these balances.
    * A similar mapping exists in the `StrategyManager`, but the `DelegationManager` additionally tracks beacon chain ETH delegated via the `EigenPodManager`. The "beacon chain ETH" strategy gets its own special address for this mapping: `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0`.
* `uint256 public withdrawalDelayBlocks`:
    * As of M2, this is 50400 (roughly 1 week)
    * Stakers must wait this amount of time before a withdrawal can be completed
* `mapping(bytes32 => bool) public pendingWithdrawals;`:
    * `Withdrawals` are hashed and set to `true` in this mapping when a withdrawal is initiated. The hash is set to false again when the withdrawal is completed. A per-staker nonce provides a way to distinguish multiple otherwise-identical withdrawals.

#### Helpful definitions

* `isDelegated(address staker) -> (bool)`
    * True if `delegatedTo[staker] != address(0)`
* `isOperator(address operator) -> (bool)` 
    * True if `_operatorDetails[operator].earningsReceiver != address(0)`

---

### Becoming an Operator

Operators interact with the following functions to become an Operator:

* [`DelegationManager.registerAsOperator`](#registerasoperator)
* [`DelegationManager.modifyOperatorDetails`](#modifyoperatordetails)
* [`DelegationManager.updateOperatorMetadataURI`](#updateoperatormetadatauri)

#### `registerAsOperator`

```solidity
function registerAsOperator(OperatorDetails calldata registeringOperatorDetails, string calldata metadataURI) external
```

Registers the caller as an Operator in EigenLayer. The new Operator provides the `OperatorDetails`, a struct containing:
* `address earningsReceiver`: the address that will receive earnings as the Operator provides services to AVSs *(currently unused)*
* `address delegationApprover`: if set, this address must sign and approve new delegation from Stakers to this Operator *(optional)*
* `uint32 stakerOptOutWindowBlocks`: the minimum delay (in blocks) between beginning and completing registration for an AVS. *(currently unused)*

`registerAsOperator` cements the Operator's `OperatorDetails`, and self-delegates the Operator to themselves - permanently marking the caller as an Operator. They cannot "deregister" as an Operator - however, they can exit the system by withdrawing their funds via `queueWithdrawals`.

*Effects*:
* Sets `OperatorDetails` for the Operator in question
* Delegates the Operator to itself
* If the Operator has shares in the `EigenPodManager`, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Operator holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

*Requirements*:
* Caller MUST NOT already be an Operator
* Caller MUST NOT already be delegated to an Operator
* `earningsReceiver != address(0)`
* `stakerOptOutWindowBlocks <= MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`: (~180 days)
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`

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

### Delegating to an Operator

Stakers interact with the following functions to delegate their shares to an Operator:

* [`DelegationManager.delegateTo`](#delegateto)
* [`DelegationManager.delegateToBySignature`](#delegatetobysignature)

#### `delegateTo`

```solidity
function delegateTo(
    address operator, 
    SignatureWithExpiry memory approverSignatureAndExpiry, 
    bytes32 approverSalt
) 
    external
```

Allows the caller (a Staker) to delegate their shares to an Operator. Delegation is all-or-nothing: when a Staker delegates to an Operator, they delegate ALL their shares. For each strategy the Staker has shares in, the `DelegationManager` will update the Operator's corresponding delegated share amounts.

*Effects*:
* Records the Staker as being delegated to the Operator
* If the Staker has shares in the `EigenPodManager`, the `DelegationManager` adds these shares to the Operator's shares for the beacon chain ETH strategy.
* For each of the three strategies in the `StrategyManager`, if the Staker holds shares in that strategy they are added to the Operator's shares under the corresponding strategy.

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`
* The caller MUST NOT already be delegated to an Operator
* The `operator` MUST already be an Operator
* If the `operator` has a `delegationApprover`, the caller MUST provide a valid `approverSignatureAndExpiry` and `approverSalt`

#### `delegateToBySignature`

```solidity
function delegateToBySignature(
    address staker,
    address operator,
    SignatureWithExpiry memory stakerSignatureAndExpiry,
    SignatureWithExpiry memory approverSignatureAndExpiry,
    bytes32 approverSalt
) 
    external
```

Allows a Staker to delegate to an Operator by way of signature. This function can be called by three different parties:
* If the Operator calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the Operator's `delegationApprover` calls this method, they need to submit only the `stakerSignatureAndExpiry`
* If the anyone else calls this method, they need to submit both the `stakerSignatureAndExpiry` AND `approverSignatureAndExpiry`

*Effects*: See `delegateTo` above.

*Requirements*: See `delegateTo` above. Additionally:
* If caller is either the Operator's `delegationApprover` or the Operator, the `approverSignatureAndExpiry` and `approverSalt` can be empty
* `stakerSignatureAndExpiry` MUST be a valid, unexpired signature over the correct hash and nonce

---

### Undelegating and Withdrawing

These methods can be called by both Stakers AND Operators, and are used to (i) undelegate a Staker from an Operator, (ii) queue a withdrawal of a Staker/Operator's shares, or (iii) complete a queued withdrawal:

* [`DelegationManager.undelegate`](#undelegate)
* [`DelegationManager.queueWithdrawals`](#queuewithdrawals)
* [`DelegationManager.completeQueuedWithdrawal`](#completequeuedwithdrawal)
* [`DelegationManager.completeQueuedWithdrawals`](#completequeuedwithdrawals)

#### `undelegate`

```solidity
function undelegate(
    address staker
) 
    external 
    onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
    returns (bytes32 withdrawalRoot)
```

`undelegate` can be called by a Staker to undelegate themselves, or by a Staker's delegated Operator (or that Operator's `delegationApprover`). Undelegation (i) queues a withdrawal on behalf of the Staker for all their delegated shares, and (ii) decreases the Operator's delegated shares according to the amounts and strategies being withdrawn.

If the Staker has active shares in either the `EigenPodManager` or `StrategyManager`, they are removed while the withdrawal is in the queue.

The withdrawal can be completed by the Staker after `withdrawalDelayBlocks`, and does not require the Staker to "fully exit" from the system -- the Staker may choose to receive their shares back in full once the withdrawal is completed (see [`completeQueuedWithdrawal`](#completequeuedwithdrawal) for details).

Note that becoming an Operator is irreversible! Although Operators can withdraw, they cannot use this method to undelegate from themselves.

*Effects*: 
* Any shares held by the Staker in the `EigenPodManager` and `StrategyManager` are removed from the Operator's delegated shares.
* The Staker is undelegated from the Operator
* If the Staker has no delegatable shares, there is no withdrawal queued or further effects
* A `Withdrawal` is queued for the Staker, tracking the strategies and shares being withdrawn
    * The Staker's withdrawal nonce is increased
    * The hash of the `Withdrawal` is marked as "pending"
* See [`EigenPodManager.removeShares`](./EigenPodManager.md#eigenpodmanagerremoveshares)
* See [`StrategyManager.removeShares`](./StrategyManager.md#removeshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_ENTER_WITHDRAWAL_QUEUE`
* Staker MUST exist and be delegated to someone
* Staker MUST NOT be an Operator
* `staker` parameter MUST NOT be zero
* Caller must be either the Staker, their Operator, or their Operator's `delegationApprover`
* See [`EigenPodManager.removeShares`](./EigenPodManager.md#eigenpodmanagerremoveshares)
* See [`StrategyManager.removeShares`](./StrategyManager.md#removeshares)

#### `queueWithdrawals`

```solidity
function queueWithdrawals(
    QueuedWithdrawalParams[] calldata queuedWithdrawalParams
) 
    external 
    onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE) 
    returns (bytes32[] memory)
```

Allows the caller to queue one or more withdrawals of their held shares across any strategy (in either/both the `EigenPodManager` or `StrategyManager`). If the caller is delegated to an Operator, the `shares` and `strategies` being withdrawn are immediately removed from that Operator's delegated share balances. Note that if the caller is an Operator, this still applies, as Operators are essentially delegated to themselves.

`queueWithdrawals` works very similarly to `undelegate`, except that the caller is not undelegated, and also may:
* Choose which strategies and how many shares to withdraw (as opposed to ALL shares/strategies)
* Specify a `withdrawer` to receive withdrawn funds once the withdrawal is completed

All shares being withdrawn (whether via the `EigenPodManager` or `StrategyManager`) are removed while the withdrawals are in the queue.

Withdrawals can be completed by the `withdrawer` after `withdrawalDelayBlocks`, and does not require the `withdrawer` to "fully exit" from the system -- they may choose to receive their shares back in full once the withdrawal is completed (see [`completeQueuedWithdrawal`](#completequeuedwithdrawal) for details).

*Effects*:
* For each withdrawal:
    * If the caller is delegated to an Operator, that Operator's delegated balances are decreased according to the `strategies` and `shares` being withdrawn.
    * A `Withdrawal` is queued for the `withdrawer`, tracking the strategies and shares being withdrawn
        * The caller's withdrawal nonce is increased
        * The hash of the `Withdrawal` is marked as "pending"
    * See [`EigenPodManager.removeShares`](./EigenPodManager.md#eigenpodmanagerremoveshares)
    * See [`StrategyManager.removeShares`](./StrategyManager.md#removeshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_ENTER_WITHDRAWAL_QUEUE`
* For each withdrawal:
    * `strategies.length` MUST equal `shares.length`
    * `strategies.length` MUST NOT be equal to 0
    * The `withdrawer` MUST NOT be 0
    * See [`EigenPodManager.removeShares`](./EigenPodManager.md#eigenpodmanagerremoveshares)
    * See [`StrategyManager.removeShares`](./StrategyManager.md#removeshares)

#### `completeQueuedWithdrawal`

```solidity
function completeQueuedWithdrawal(
    Withdrawal calldata withdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) 
    external 
    onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE)
    nonReentrant
```

After waiting `withdrawalDelayBlocks`, this allows the `withdrawer` of a `Withdrawal` to finalize a withdrawal and receive either (i) the underlying tokens of the strategies being withdrawn from, or (ii) the shares being withdrawn. This choice is dependent on the passed-in parameter `receiveAsTokens`.

For each strategy/share pair in the `Withdrawal`:
* If the `withdrawer` chooses to receive tokens:
    * The shares are converted to their underlying tokens via either the `EigenPodManager` or `StrategyManager` and sent to the `withdrawer`.
* If the `withdrawer` chooses to receive shares (and the strategy belongs to the `StrategyManager`): 
    * The shares are awarded to the `withdrawer` via the `StrategyManager`
    * If the `withdrawer` is delegated to an Operator, that Operator's delegated shares are increased by the added shares (according to the strategy being added to).

`Withdrawals` concerning `EigenPodManager` shares have some additional nuance depending on whether a withdrawal is specified to be received as tokens vs shares (read more about "why" in [`EigenPodManager.md`](./EigenPodManager.md)):
* `EigenPodManager` withdrawals received as shares: 
    * Shares ALWAYS go back to the originator of the withdrawal (rather than the `withdrawer` address). 
    * Shares are also delegated to the originator's Operator, rather than the `withdrawer's` Operator.
    * Shares received by the originator may be lower than the shares originally withdrawn if the originator has debt.
* `EigenPodManager` withdrawals received as tokens:
    * Before the withdrawal can be completed, the originator needs to prove that a withdrawal occurred on the beacon chain (see [`EigenPod.verifyAndProcessWithdrawals`](./EigenPodManager.md#eigenpodverifyandprocesswithdrawals)).

*Effects*:
* The hash of the `Withdrawal` is removed from the pending withdrawals
* If `receiveAsTokens`:
    * See [`StrategyManager.withdrawSharesAsTokens`](./StrategyManager.md#withdrawsharesastokens)
    * See [`EigenPodManager.withdrawSharesAsTokens`](./EigenPodManager.md#eigenpodmanagerwithdrawsharesastokens)
* If `!receiveAsTokens`:
    * For `StrategyManager` strategies:
        * Shares are awarded to the `withdrawer` and delegated to the `withdrawer's` Operator
        * See [`StrategyManager.addShares`](./StrategyManager.md#addshares)
    * For the native beacon chain ETH strategy (`EigenPodManager`):
        * Shares are awarded to `withdrawal.staker`, and delegated to the Staker's Operator
        * See [`EigenPodManager.addShares`](./EigenPodManager.md#eigenpodmanageraddshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EXIT_WITHDRAWAL_QUEUE`
* The hash of the passed-in `Withdrawal` MUST correspond to a pending withdrawal
    * At least `withdrawalDelayBlocks` MUST have passed before `completeQueuedWithdrawal` is called
    * Caller MUST be the `withdrawer` specified in the `Withdrawal`
* If `receiveAsTokens`:
    * The caller MUST pass in the underlying `IERC20[] tokens` being withdrawn in the appropriate order according to the strategies in the `Withdrawal`.
    * See [`StrategyManager.withdrawSharesAsTokens`](./StrategyManager.md#withdrawsharesastokens)
    * See [`EigenPodManager.withdrawSharesAsTokens`](./EigenPodManager.md#eigenpodmanagerwithdrawsharesastokens)
* If `!receiveAsTokens`:
    * See [`StrategyManager.addShares`](./StrategyManager.md#addshares)
    * See [`EigenPodManager.addShares`](./EigenPodManager.md#eigenpodmanageraddshares)

*As of M2*:
* The `middlewareTimesIndex` parameter has to do with the Slasher, which currently does nothing. As of M2, this parameter has no bearing on anything and can be ignored.

#### `completeQueuedWithdrawals`

```solidity
function completeQueuedWithdrawals(
    Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexes,
    bool[] calldata receiveAsTokens
) 
    external 
    onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) 
    nonReentrant
```

This method is the plural version of [`completeQueuedWithdrawal`](#completequeuedwithdrawal).

---

### Accounting

These methods are called by the `StrategyManager` and `EigenPodManager` to update delegated share accounting when a Staker's balance changes (e.g. due to a deposit):

* [`DelegationManager.increaseDelegatedShares`](#increasedelegatedshares)
* [`DelegationManager.decreaseDelegatedShares`](#decreasedelegatedshares)

#### `increaseDelegatedShares`

```solidity
function increaseDelegatedShares(
    address staker, 
    IStrategy strategy, 
    uint256 shares
)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by either the `StrategyManager` or `EigenPodManager` when a Staker's shares for one or more strategies increase. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the increase.

*Entry Points*:
* `StrategyManager.depositIntoStrategy`
* `StrategyManager.depositIntoStrategyWithSignature`
* `EigenPod.verifyWithdrawalCredentials`
* `EigenPod.verifyBalanceUpdate`
* `EigenPod.verifyAndProcessWithdrawals`

*Effects*: If the Staker in question is delegated to an Operator, the Operator's shares for the `strategy` are increased.
* This method is a no-op if the Staker is not delegated to an Operator.

*Requirements*:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`

#### `decreaseDelegatedShares`

```solidity
function decreaseDelegatedShares(
    address staker, 
    IStrategy strategy, 
    uint256 shares
)
    external
    onlyStrategyManagerOrEigenPodManager
```

Called by the `EigenPodManager` when a Staker's shares decrease. This method is called to ensure that if the Staker is delegated to an Operator, that Operator's share count reflects the decrease.

*Entry Points*: This method may be called as a result of the following top-level function calls:
* `EigenPod.verifyBalanceUpdate`
* `EigenPod.verifyAndProcessWithdrawals`

*Effects*: If the Staker in question is delegated to an Operator, the Operator's delegated balance for the `strategy` is decreased by `shares`
* This method is a no-op if the Staker is not delegated to an Operator.

*Requirements*:
* Caller MUST be either the `StrategyManager` or `EigenPodManager` (although the `StrategyManager` doesn't use this method)