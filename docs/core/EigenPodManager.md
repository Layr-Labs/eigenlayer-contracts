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

* `mapping(address => IEigenPod) public ownerToPod`: Tracks the deployed `EigenPod` for each Staker
* `mapping(address => int256) public podOwnerShares`: Keeps track of the actively restaked beacon chain ETH for each Staker. 
    * In some cases, a beacon chain balance update may cause a Staker's balance to drop below zero. This is because when queueing for a withdrawal in the `DelegationManager`, the Staker's current shares are fully removed. If the Staker's beacon chain balance drops after this occurs, their `podOwnerShares` may go negative. This is a temporary change to account for the drop in balance, and is ultimately corrected when the withdrawal is finally processed.
    * Since balances on the consensus layer are stored only in Gwei amounts, the EigenPodManager enforces the invariant that `podOwnerShares` is always a whole Gwei amount for every staker, i.e. `podOwnerShares[staker] % 1e9 == 0` always.

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
* [`removeShares`](#removeshares)

When completing a queued undelegation or withdrawal, the `DelegationManager` calls one of these two methods:
* [`addShares`](#addshares)
* [`withdrawSharesAsTokens`](#withdrawsharesastokens)

#### `removeShares`

```solidity
function removeShares(
    address podOwner, 
    uint256 shares
) 
    external 
    onlyDelegationManager
```

The `DelegationManager` calls this method when a Staker queues a withdrawal (or undelegates, which also queues a withdrawal). The shares are removed while the withdrawal is in the queue, and when the queue completes, the shares will either be re-awarded or withdrawn as tokens (`addShares` and `withdrawSharesAsTokens`, respectively).

The Staker's share balance is decreased by the removed `shares`. 

This method is not allowed to cause the `Staker's` balance to go negative. This prevents a Staker from queueing a withdrawal for more shares than they have (or more shares than they delegated to an Operator).

*Entry Points*:
* `DelegationManager.undelegate`
* `DelegationManager.queueWithdrawals`

*Effects*:
* Removes `shares` from `podOwner's` share balance

*Requirements*:
* `podOwner` MUST NOT be zero
* `shares` MUST NOT be negative when converted to `int256`
* `shares` MUST NOT be greater than `podOwner's` share balance
* `shares` MUST be a whole Gwei amount

#### `addShares`

```solidity
function addShares(
    address podOwner,
    uint256 shares
) 
    external 
    onlyDelegationManager 
    returns (uint256)
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as "shares" (rather than as the underlying tokens). A Pod Owner might want to do this in order to change their delegated Operator without needing to fully exit their validators.

Note that typically, shares from completed withdrawals are awarded to a `withdrawer` specified when the withdrawal is initiated in `DelegationManager.queueWithdrawals`. However, because beacon chain ETH shares are linked to proofs provided to a Pod Owner's `EigenPod`, this method is used to award shares to the original Pod Owner.

If the Pod Owner has a share deficit (negative shares), the deficit is repaid out of the added `shares`. If the Pod Owner's positive share count increases, this change is returned to the `DelegationManager` to be delegated to the Pod Owner's Operator (if they have one).

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* The `podOwner's` share balance is increased by `shares`

*Requirements*:
* `podOwner` MUST NOT be zero
* `shares` MUST NOT be negative when converted to an `int256`
* `shares` MUST be a whole Gwei amount

#### `withdrawSharesAsTokens`

```solidity
function withdrawSharesAsTokens(
    address podOwner, 
    address destination, 
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
* If `podOwner's` share balance is negative, `shares` are added until the balance hits 0
    * Any remaining shares are converted 1:1 to ETH and sent to `destination` (see [`EigenPod.withdrawRestakedBeaconChainETH`](./EigenPod.md#withdrawrestakedbeaconchaineth))

*Requirements*:
* `podOwner` MUST NOT be zero
* `destination` MUST NOT be zero
* `shares` MUST NOT be negative when converted to an `int256`
* `shares` MUST be a whole Gwei amount
* See [`EigenPod.withdrawRestakedBeaconChainETH`](./EigenPod.md#withdrawrestakedbeaconchaineth)

---

### Other Methods

#### `recordBeaconChainETHBalanceUpdate`

```solidity
function recordBeaconChainETHBalanceUpdate(
    address podOwner,
    int256 sharesDelta
) 
    external 
    onlyEigenPod(podOwner) 
    nonReentrant
```

This method is called by an `EigenPod` to report a change in its Pod Owner's shares. It accepts a positive or negative `sharesDelta`, which is added/subtracted against the Pod Owner's shares. The delta is also communicated to the `DelegationManager`, which updates the number of shares the Pod Owner has delegated to an Operator.

Note that this method _may_ result in a Pod Owner's shares going negative. This can occur when:
* The Pod Owner has queued a withdrawal for all their Beacon Chain ETH shares via `DelegationManager.queueWithdrawals`
    * This will set the `EigenPodManager.podOwnerShares[podOwner]` to 0
* The Pod Owner's pod reports a negative delta, perhaps due to the Pod Owner getting slashed on the beacon chain.

In this case, the Pod Owner's `podOwnerShares` will go negative.

*Entry Points*:
* `EigenPod.verifyWithdrawalCredentials`
* `EigenPod.startCheckpoint`
* `EigenPod.verifyCheckpointProofs`

*Effects*:
* Adds or removes `sharesDelta` from the Pod Owner's shares
* If `sharesDelta` is negative: see [`DelegationManager.decreaseDelegatedShares`](./DelegationManager.md#decreasedelegatedshares)
* If `sharesDelta` is positive: see [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Caller MUST be the `EigenPod` associated with the passed-in `podOwner`
* `sharesDelta` MUST be a whole Gwei amount
