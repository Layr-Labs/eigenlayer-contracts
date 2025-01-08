## EigenPodManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`EigenPodManager.sol`](../../src/contracts/pods/EigenPodManager.sol) | Singleton | Transparent proxy |
| [`EigenPod.sol`](../../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy |

The `EigenPodManager` manages the relationship between a Staker's `EigenPod`, the delegatable shares each Staker holds in the beacon chain ETH strategy, and the withdrawal of those shares via the `DelegationManager`. These functions together support Stakers' ability to restake beacon chain ETH and delegate restaked ETH to Operators in order to earn additional yield.

The `EigenPodManager` is the entry point for this process, allowing Stakers to deploy an `EigenPod` and begin restaking. After a Staker deploys their `EigenPod`, the `EigenPodManager` receives various updates from the pod that add or remove shares from the Staker.

#### High-level Concepts

This document organizes methods according to the following themes (click each to be taken to the relevant section):
* [Depositing Into EigenLayer](#depositing-into-eigenlayer)
* [Withdrawal Processing](#withdrawal-processing)
* [Other Methods](#other-methods)

#### Important State Variables

* `mapping(address podOwner => IEigenPod) public ownerToPod`: Tracks the deployed `EigenPod` for each Staker
* `mapping(address podOwner => int256 shares) public podOwnerDepositShares`: Keeps track of the actively restaked beacon chain ETH for each Staker
    * In some cases prior to the slashing release, a beacon chain balance update may cause a Staker's balance to drop below zero. This is because when queueing for a withdrawal in the `DelegationManager`, the Staker's current shares are fully removed. If the Staker's beacon chain balance drops after this occurs, their `podOwnerDepositShares` may go negative. This is a temporary change to account for the drop in balance, and is ultimately corrected when the withdrawal is finally processed.
    * Since balances on the consensus layer are stored only in Gwei amounts, the EigenPodManager enforces the invariant that `podOwnerDepositShares` is always a whole Gwei amount for every staker, i.e. `podOwnerDepositShares[staker] % 1e9 == 0` always.
* `mapping(address staker => BeaconChainSlashingFactor) internal _beaconChainSlashingFactor`: Tracks the "slashing factor" for a given staker
    * This is a substitute for the previous model of decreasing the shares below 0. Instead, the slashing factor tracks the amount by which a staker has been slashed, which is used by the `DelegationManager` to determine how many shares to withhold due to beacon chain balance decreases and EigenLayer slashing events.
* `uint256 public burnableETHShares`: Tracks the total amount of ETH across pods which has been slashed on EigenLayer but not yet burned.

---

### Depositing Into EigenLayer

Before a Staker begins restaking beacon chain ETH, they need to deploy an `EigenPod`, stake, and start a beacon chain validator:
* [`createPod`](#createpod)
* [`stake`](#stake)

#### `createPod`

```solidity
function createPod()
    external
    onlyWhenNotPaused(PAUSED_NEW_EIGENPODS)
    returns (address)
```

Allows a Staker to deploy an `EigenPod` instance, if they have not done so already.

Each Staker can only deploy a single `EigenPod` instance, but a single `EigenPod` can serve as the fee recipient / withdrawal credentials for any number of beacon chain validators. Each `EigenPod` is created using Create2 and the beacon proxy pattern, using the Staker's address as the Create2 salt.

As part of the `EigenPod` deployment process, the Staker is made the Pod Owner, a permissioned role within the `EigenPod`.

*Effects*:
* Create2 deploys `EigenPodManager.beaconProxyBytecode`, appending the `eigenPodBeacon` address as a constructor argument. `bytes32(msg.sender)` is used as the salt.
    * `address eigenPodBeacon` is an [OpenZeppelin v4.7.1 `UpgradableBeacon`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/UpgradeableBeacon.sol), whose implementation address points to the current `EigenPod` implementation
    * `beaconProxyBytecode` is the constructor code for an [OpenZeppelin v4.7.1 `BeaconProxy`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol)
* `EigenPod.initialize(msg.sender)`: initializes the pod, setting the caller as the Pod Owner and activating restaking for any validators pointed at the pod.
* Maps the new pod to the Pod Owner (each address can only deploy a single `EigenPod`)

*Requirements*:
* Caller MUST NOT have deployed an `EigenPod` already
* Pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`

#### `stake`

```solidity
function stake(
    bytes calldata pubkey,
    bytes calldata signature,
    bytes32 depositDataRoot
)
    external
    payable
    onlyWhenNotPaused(PAUSED_NEW_EIGENPODS)
```

Allows a Staker to deposit 32 ETH into the beacon chain deposit contract, providing the credentials for the Staker's beacon chain validator. The `EigenPod.stake` method is called, which automatically calculates the correct withdrawal credentials for the pod and passes these to the deposit contract along with the 32 ETH.

*Effects*:
* Deploys and initializes an `EigenPod` on behalf of Staker, if this has not been done already
* See [`EigenPod.stake`](./EigenPod.md#stake)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`
* See [`EigenPod.stake`](./EigenPod.md#stake)

---

### Withdrawal Processing

The `DelegationManager` is the entry point for all undelegation and withdrawals, which must be queued for a time before being completed. When a withdrawal is initiated, the `DelegationManager` calls the following method:
* [`removeDepositShares`](#removeDepositShares)

When completing a queued undelegation or withdrawal, the `DelegationManager` calls one of these two methods:
* [`addShares`](#addshares)
* [`withdrawSharesAsTokens`](#withdrawsharesastokens)

#### `removeDepositShares`

```solidity
function removeDepositShares(
    address podOwner,
    IStrategy strategy,
    uint256 depositSharesToRemove
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a Staker queues a withdrawal (or undelegates, which also queues a withdrawal). The shares are removed while the withdrawal is in the queue, and when the queue completes, the shares will either be re-awarded or withdrawn as tokens (`addShares` and `withdrawSharesAsTokens`, respectively).

The Staker's share balance is decreased by `depositSharesToRemove`.

This method is not allowed to cause the `Staker's` balance to go negative. This prevents a Staker from queueing a withdrawal for more shares than they have (or more shares than they delegated to an Operator).

*Entry Points*:
* `DelegationManager.undelegate`
* `DelegationManager.queueWithdrawals`

*Effects*:
* Removes `depositSharesToRemove` from `podOwner` share balance in `podOwnerDepositShares`
* Emits a `NewTotalShares` event

*Requirements*:
* Caller MUST be the `DelegationManager`
* `strategy` MUST be `beaconChainETHStrategy`
* `staker` MUST NOT be zero
* `depositSharesToRemove` MUST be less than `staker` share balance in `podOwnerDepositShares`
* `depositSharesToRemove` MUST be a whole Gwei amount

#### `addShares`

```solidity
function addShares(
    address staker,
    IStrategy strategy,
    IERC20,
    uint256 shares
)
    external
    onlyDelegationManager
    returns (uint256, uint256)
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as "shares" (rather than as the underlying tokens). A staker might want to do this in order to change their delegated Operator without needing to fully exit their validators.

Note that typically, shares from completed withdrawals are awarded to a `withdrawer` specified when the withdrawal is initiated in `DelegationManager.queueWithdrawals`. However, because beacon chain ETH shares are linked to proofs provided to a staker's `EigenPod`, this method is used to award shares to the original staker.

If the staker has a share deficit (negative shares), the deficit is repaid out of the added `shares`. If the Pod Owner's positive share count increases, this change is returned to the `DelegationManager` to be delegated to the staker's Operator (if they have one).

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* Increases `staker`'s share balance in `podOwnerDepositShares` by `shares`

*Requirements*:
* Caller MUST be the `DelegationManager`
* `strategy` MUST be `beaconChainETHStrategy`
* `staker` MUST NOT be `address(0)`
* `shares` MUST NOT be negative when converted to an `int256`
* Emits `PodSharesUpdated` and `NewTotalShares` events

#### `withdrawSharesAsTokens`

```solidity
function withdrawSharesAsTokens(
    address podOwner,
    IStrategy strategy,
    IERC20,
    uint256 shares
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as tokens (rather than shares). This can be used to "fully exit" some amount of beacon chain ETH and send it to a recipient (via `EigenPod.withdrawRestakedBeaconChainETH`).

Note that because this method entails withdrawing and sending beacon chain ETH, two conditions must be met for this method to succeed: (i) the ETH being withdrawn should already be in the `EigenPod`, and (ii) the amount being withdrawn should be accounted for in `EigenPod.withdrawableRestakedExecutionLayerGwei`. This latter condition can be achieved by completing an `EigenPod` checkpoint just prior to completing a queued `DelegationManager` withdrawal (see [EigenPod: Checkpointing Validators](./EigenPod.md#checkpointing-validators) for details).

Also note that, like `addShares`, if the original Pod Owner has a share deficit (negative shares), the deficit is repaid out of the withdrawn `shares` before any native ETH is withdrawn.

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* If `staker`'s share balance in `podOwnerDepositShares` is negative (i.e. the staker has a *deficit*):
    * If `shares` is greater than the current deficit:
        * Sets `staker` balance in `podOwnerDepositShares` to 0
        * Subtracts `|deficit|` from `shares` and converts remaining shares 1:1 to ETH (see [`EigenPod.withdrawRestakedBeaconChainETH`](./EigenPod.md#withdrawrestakedbeaconchaineth))
    * If `shares` is less than or equal to the current deficit:
        * Increases `staker` negative balance in `podOwnerDepositShares` by `shares`, bringing it closer to 0
        * Does *not* withdraw any shares
* Emits `PodSharesUpdated` and `NewTotalShares` events

*Requirements*:
* Caller MUST be the `DelegationManager`
* `strategy` MUST be `beaconChainETHStrategy`
* `staker` MUST NOT be `address(0)`
* `shares` MUST NOT be negative when converted to an `int256`
* See [`EigenPod.withdrawRestakedBeaconChainETH`](./EigenPod.md#withdrawrestakedbeaconchaineth)

---

### Other Methods

#### `recordBeaconChainETHBalanceUpdate`

```solidity
function recordBeaconChainETHBalanceUpdate(
    address podOwner,
    uint256 prevRestakedBalanceWei,
    int256 balanceDeltaWei
)
    external
    onlyEigenPod(podOwner)
    nonReentrant
```

This method is called by an `EigenPod` to report a change in its Pod Owner's shares. It accepts a positive or negative `balanceDeltaWei`, which is added/subtracted against the Pod Owner's shares. The delta is also communicated to the `DelegationManager`, which updates the number of shares the Pod Owner has delegated to an Operator.

Note that this method _may_ result in a Pod Owner's shares going negative. This can occur when:
* The Pod Owner has queued a withdrawal for all their Beacon Chain ETH shares via `DelegationManager.queueWithdrawals`
    * This will set the `EigenPodManager.podOwnerDepositShares[podOwner]` to 0
* The Pod Owner's pod reports a negative delta, perhaps due to the Pod Owner getting slashed on the beacon chain.

In this case, the Pod Owner's `podOwnerDepositShares` will go negative.

*Entry Points*:
* `EigenPod.verifyWithdrawalCredentials`
* `EigenPod.startCheckpoint`
* `EigenPod.verifyCheckpointProofs`

*Effects*:
* Adds or removes `balanceDeltaWei` from the Pod Owner's shares
* If `balanceDeltaWei` is positive or 0:
  * Adds `shares` to `podOwnerDepositShares` for `podOwner`
  * Emits `PodSharesUpdated` and `NewTotalShares` events
  * Calls [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)
* If `balanceDeltaWei` is negative:
  * Reduces slashing factor proportional to relative balance decrease
  * Emits `BeaconChainSlashingFactorDecreased` event
  * Calls [`DelegationManager.decreaseDelegatedShares`](./DelegationManager.md#decreasedelegatedshares)

*Requirements*:
* Caller MUST be the `EigenPod` associated with the passed-in `podOwner`
* `podOwner` MUST NOT be `address(0)`
* `balanceDeltaWei` MUST be a whole Gwei amount
* Legacy withdrawals MUST be complete (i.e. `currentDepositShares >= 0`)
