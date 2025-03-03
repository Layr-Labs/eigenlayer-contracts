## EigenPodManager

| File | Notes |
| -------- | -------- |
| [`EigenPodManager.sol`](../../src/contracts/pods/EigenPodManager.sol) | |
| [`EigenPodManagerStorage.sol`](../../src/contracts/pods/EigenPodManagerStorage.sol) | state variables |
| [`IEigenPodManager.sol`](../../src/contracts/interfaces/IEigenPodManager.sol) | interface |

Libraries and Mixins:
| File | Notes |
| -------- | -------- |
| [`EigenPodPausingConstants.sol`](../../src/contracts/pods/EigenPodPausingConstants.sol) | pause values for `EigenPod/EigenPodManager` methods |

## Prior Reading

* [Shares Accounting](./accounting/SharesAccounting.md), especially [_Handling Beacon Chain Balance Decreases in EigenPods_](./accounting/SharesAccounting.md#handling-beacon-chain-balance-decreases-in-eigenpods)

## Overview

The `EigenPodManager` manages the "beacon chain ETH strategy", a virtual strategy that stakers can hold delegatable shares in, similar to the strategies managed by the `StrategyManager`. Whereas the `StrategyManager's` strategy shares are backed by deposited ERC20 tokens, beacon chain ETH strategy shares are backed either by beacon chain validators or native ETH attributed to `EigenPods`.

The `EigenPodManager` allows any staker to deploy an `EigenPod`. `EigenPods` contains beacon chain state proof logic that enable a staker to point either/both a validator's _withdrawal credentials_ and/or _fee recipient_ addresses to their pod. After submitting beacon chain state proofs to their pod, the staker is awarded deposit shares in the beacon chain ETH strategy, which are then delegated to their operator in the `DelegationManager` (if applicable). For more details, see [`EigenPod.md`](./EigenPod.md).

As an `EigenPod` receives balance updates, they are forwarded to the `EigenPodManager`. Balance increases resulting from validator activity on the beacon chain or ETH received in the `EigenPod` will result in an increase in the staker's _deposit shares_ for the beacon chain ETH strategy.

Balance decreases resulting from validator inactivity or beacon chain slashing do NOT decrease the staker's deposit shares. Instead, they result in a _decrease_ to the staker's _beacon chain slashing factor_. Among other factors, the `DelegationManager` uses the beacon chain slashing factor when determining:
* How many of a staker's _deposit shares_ can actually be withdrawn
* How many of a staker's _deposit shares_ can be delegated to an operator

Note that the number of _withdrawable shares_ a staker's _deposit shares_ represent can be queried using `DelegationManager.getWithdrawableShares(staker, strategies)`.

The `EigenPodManager's` responsibilities can be broken down into the following concepts:
* [Depositing Into EigenLayer](#depositing-into-eigenlayer)
* [Withdrawal Processing](#withdrawal-processing)
* [Other Methods](#other-methods)

## Parameterization

* `beaconChainETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0`
    * A virtual strategy used to represent beacon chain ETH internally. The `DelegationManager` uses this address to denote the beacon chain ETH strategy managed by the `EigenPodManager`. However, it does not correspond to an actual contract!
* `ethPOS = 0x00000000219ab540356cBB839Cbe05303d7705Fa`
    * The address of the beacon chain deposit contract
* `beaconProxyBytecode` (defined in `EigenPodManagerStorage.sol`)
    * `EigenPods` are deployed using a beacon proxy. This bytecode is a constant, containing the _creation bytecode_ calculated by compiling [OpenZeppelin's `BeaconProxy` contract at version 4.7.1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/beacon/BeaconProxy.sol). Compilation used solc version `0.8.12`, optimization enabled, 200 runs. Example verified contract [here](https://etherscan.io/address/0xA6f93249580EC3F08016cD3d4154AADD70aC3C96#code).

---

## Depositing Into EigenLayer

Before a staker begins restaking beacon chain ETH, they need to deploy an `EigenPod`, stake, and start a beacon chain validator:
* [`createPod`](#createpod)
* [`stake`](#stake)

#### `createPod`

```solidity
/**
 * @notice Creates an EigenPod for the sender.
 * @dev Function will revert if the `msg.sender` already has an EigenPod.
 * @dev Returns EigenPod address
 */
function createPod()
    external
    onlyWhenNotPaused(PAUSED_NEW_EIGENPODS)
    returns (address)
```

Allows a staker to deploy an `EigenPod` instance, if they have not done so already.

Each staker can only deploy a single `EigenPod` instance, but a single `EigenPod` can serve as the fee recipient / withdrawal credentials for any number of beacon chain validators. Each `EigenPod` is created using Create2 and the beacon proxy pattern, using the staker's address as the Create2 salt.

As part of the `EigenPod` deployment process, the staker is made the Pod Owner, a permissioned role within the `EigenPod`.

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
/**
 * @notice Stakes for a new beacon chain validator on the sender's EigenPod.
 * Also creates an EigenPod for the sender if they don't have one already.
 * @param pubkey The 48 bytes public key of the beacon chain validator.
 * @param signature The validator's signature of the deposit data.
 * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
 */
function stake(
    bytes calldata pubkey,
    bytes calldata signature,
    bytes32 depositDataRoot
)
    external
    payable
    onlyWhenNotPaused(PAUSED_NEW_EIGENPODS)
```

Allows a staker to deposit 32 ETH into the beacon chain deposit contract, providing the credentials for the staker's beacon chain validator. The `EigenPod.stake` method is called, which automatically calculates the correct withdrawal credentials for the pod and passes these to the deposit contract along with the 32 ETH.

*Effects*:
* Deploys and initializes an `EigenPod` on behalf of staker, if this has not been done already
* See [`EigenPod.stake`](./EigenPod.md#stake)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`
* See [`EigenPod.stake`](./EigenPod.md#stake)

---

## Withdrawal Processing

These methods are callable ONLY by the DelegationManager, and are used when processing undelegations and withdrawals.

**Concepts**:
* [Shares Accounting - Handling Beacon Chain Balance Decreases](./accounting/SharesAccounting.md#handling-beacon-chain-balance-decreases-in-eigenpods)
* [Deposit Shares and Beacon Chain Slashing](#deposit-shares-and-beacon-chain-slashing)

**Methods**:
* [`removeDepositShares`](#removeDepositShares)
* [`addShares`](#addshares)
* [`withdrawSharesAsTokens`](#withdrawsharesastokens)

#### Deposit Shares and Beacon Chain Slashing

The `EigenPodManager` tracks a staker's _deposit shares_ and _beacon chain slashing factor_ using the following state variables:

```solidity
/**
 * @notice mapping from pod owner to the deposit shares they have in the virtual beacon chain ETH strategy
 * 
 * @dev When an EigenPod registers a balance increase, deposit shares are increased. When registering a balance
 * decrease, however, deposit shares are NOT decreased. Instead, the pod owner's beacon chain slashing factor
 * is decreased proportional to the balance decrease. This impacts the number of shares that will be withdrawn
 * when the deposit shares are queued for withdrawal in the DelegationManager.
 * 
 * Note that prior to the slashing release, deposit shares were decreased when balance decreases occurred.
 * In certain cases, a combination of queueing a withdrawal plus registering a balance decrease could result
 * in a staker having negative deposit shares in this mapping. This negative value would be corrected when the
 * staker completes a withdrawal (as tokens or as shares).
 *
 * With the slashing release, negative shares are no longer possible. However, a staker can still have negative
 * shares if they met the conditions for them before the slashing release. If this is the case, that staker
 * should complete any outstanding queued withdrawal in the DelegationManager ("as shares"). This will correct
 * the negative share count and allow the staker to continue using their pod as normal.
 */
mapping(address podOwner => int256 shares) public podOwnerDepositShares;

/**
 * @notice The amount of beacon chain slashing experienced by a pod owner as a proportion of WAD
 * @param isSet whether the slashingFactor has ever been updated. Used to distinguish between
 * a value of "0" and an uninitialized value.
 * @param slashingFactor the proportion of the pod owner's balance that has been decreased due to
 * slashing or other beacon chain balance decreases.
 * @dev NOTE: if !isSet, `slashingFactor` should be treated as WAD. `slashingFactor` is monotonically
 * decreasing and can hit 0 if fully slashed.
 */
struct BeaconChainSlashingFactor {
    bool isSet;
    uint64 slashingFactor;
}

/// @notice Returns the slashing factor applied to the `staker` for the `beaconChainETHStrategy`
/// Note: this value starts at 1 WAD (1e18) for all stakers, and is updated when a staker's pod registers
/// a balance decrease.
mapping(address staker => BeaconChainSlashingFactor) internal _beaconChainSlashingFactor;
```

#### `removeDepositShares`

```solidity
/**
 * @notice Used by the DelegationManager to remove a pod owner's deposit shares when they enter the withdrawal queue.
 * Simply decreases the `podOwner`'s shares by `shares`, down to a minimum of zero.
 * @dev This function reverts if it would result in `podOwnerDepositShares[podOwner]` being less than zero, i.e. it is forbidden for this function to
 * result in the `podOwner` incurring a "share deficit". This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
 * shares from the operator to whom the staker is delegated.
 * @dev The delegation manager validates that the podOwner is not address(0)
 */
function removeDepositShares(
    address podOwner,
    IStrategy strategy,
    uint256 depositSharesToRemove
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a staker queues a withdrawal (or undelegates, which also queues a withdrawal). The shares are removed while the withdrawal is in the queue, and when the queue completes, the shares will either be re-awarded or withdrawn as tokens (`addShares` and `withdrawSharesAsTokens`, respectively).

The staker's share balance is decreased by `depositSharesToRemove`.

This method is not allowed to cause the `staker's` balance to go negative. This prevents a staker from queueing a withdrawal for more shares than they have (or more shares than they delegated to an operator).

Note that the amount of deposit shares removed while in the withdrawal queue may not equal the amount credited when the withdrawal is completed. The staker may receive fewer if slashing occurred; see [`DelegationManager.md`](./DelegationManager.md) for details.

*Effects*:
* Removes `depositSharesToRemove` from `podOwner` share balance in `podOwnerDepositShares`
* Emits a `NewTotalShares` event

*Requirements*:
* Caller MUST be the `DelegationManager`
* `strategy` MUST be `beaconChainETHStrategy`
* `staker` MUST NOT be zero
* `depositSharesToRemove` MUST be less than `staker` share balance in `podOwnerDepositShares`

#### `addShares`

```solidity
/**
 * @notice Increases the `podOwner`'s shares by `shares`, paying off negative shares if needed.
 * Used by the DelegationManager to award a pod owner shares on exiting the withdrawal queue
 * @return existingDepositShares the pod owner's shares prior to any additions. Returns 0 if negative
 * @return addedShares the number of shares added to the staker's balance above 0. This means that if,
 * after shares are added, the staker's balance is non-positive, this will return 0.
 */
function addShares(
    address staker,
    IStrategy strategy,
    uint256 shares
)
    external
    onlyDelegationManager
    returns (uint256, uint256)
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal "as shares" (rather than as the underlying tokens). A staker might want to do this in order to change their delegated operator without needing to fully exit their validators.

This method credits the input deposit shares to the staker. In most cases, the input `shares` equal the same shares originally removed when the withdrawal was queued. However, if the staker's operator was slashed (or the staker experienced beacon chain slashing), they may receive less. See [`DelegationManager.md`](./DelegationManager.md) for details.

If the staker has a share deficit (negative shares), the deficit is repaid out of the added `shares`. If the Pod Owner's positive share count increases, this change is returned to the `DelegationManager` to be delegated to the staker's operator (if they have one).

*Effects*:
* Increases `staker`'s deposit share balance in `podOwnerDepositShares` by `shares`

*Requirements*:
* Caller MUST be the `DelegationManager`
* `strategy` MUST be `beaconChainETHStrategy`
* `staker` MUST NOT be `address(0)`
* `shares` MUST NOT be negative when converted to an `int256`
* Emits `PodSharesUpdated` and `NewTotalShares` events

#### `withdrawSharesAsTokens`

```solidity
/**
 * @notice Used by the DelegationManager to complete a withdrawal, sending tokens to the pod owner
 * @dev Prioritizes decreasing the podOwner's share deficit, if they have one
 * @dev This function assumes that `removeShares` has already been called by the delegationManager, hence why
 *      we do not need to update the podOwnerDepositShares if `currentpodOwnerDepositShares` is positive
 */
function withdrawSharesAsTokens(
    address podOwner,
    IStrategy strategy,
    IERC20,
    uint256 shares
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as the tokens underlying the shares. This can be used to "fully exit" some amount of native ETH and send it to the pod owner (via `EigenPod.withdrawRestakedBeaconChainETH`).

Note that because this method entails withdrawing and sending native ETH, two conditions must be met for this method to succeed: (i) the ETH being withdrawn should already be in the `EigenPod`, and (ii) the amount being withdrawn should be accounted for in `EigenPod.withdrawableExecutionLayerGwei`. This latter condition can be achieved by completing an `EigenPod` checkpoint just prior to completing a queued `DelegationManager` withdrawal (see [EigenPod: Checkpointing Validators](./EigenPod.md#checkpointing-validators) for details).

Also note that, like `addShares`, if the original Pod Owner has a share deficit (negative shares), the deficit is repaid out of the withdrawn `shares` before any native ETH is withdrawn.

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

## Other Methods

**Methods**:
* [`recordBeaconChainETHBalanceUpdate`](#recordbeaconchainethbalanceupdate)
* [`increaseBurnableShares`](#increaseburnableshares)

#### `recordBeaconChainETHBalanceUpdate`

```solidity
/**
 * @notice Adds any positive share delta to the pod owner's deposit shares, and delegates them to the pod
 * owner's operator (if applicable). A negative share delta does NOT impact the pod owner's deposit shares,
 * but will reduce their beacon chain slashing factor and delegated shares accordingly.
 * @param podOwner is the pod owner whose balance is being updated.
 * @param prevRestakedBalanceWei is the total amount restaked through the pod before the balance update, including
 * any amount currently in the withdrawal queue.
 * @param balanceDeltaWei is the amount the balance changed
 * @dev Callable only by the podOwner's EigenPod contract.
 * @dev Reverts if `sharesDelta` is not a whole Gwei amount
 */
function recordBeaconChainETHBalanceUpdate(
    address podOwner,
    uint256 prevRestakedBalanceWei,
    int256 balanceDeltaWei
)
    external
    onlyEigenPod(podOwner)
    nonReentrant
```

This method is called by an `EigenPod` to report a change in its pod owner's shares. It accepts a positive or negative `balanceDeltaWei`. A positive delta is added to the pod owner's _deposit shares,_ and delegated to their operator if applicable. A negative delta is NOT removed from the pod owner's deposit shares. Instead, the proportion of the balance decrease is used to update the pod owner's beacon chain slashing factor and decrease the number of shares delegated to their operator (if applicable). A zero delta results in no change. 

**Note** that prior to the slashing release, negative balance deltas subtracted from the pod owner's shares, and could, in certain cases, result in a negative share balance. As of the slashing release, negative balance deltas no longer subtract from share balances, updating the beacon chain slashing factor instead. 

If a staker has negative shares as of the slashing release, this method will REVERT, preventing any further balance updates from their pod while the negative share balance persists. In order to fix this and restore the use of their pod, the staker should complete any outstanding withdrawals in the `DelegationManager` "as shares," which will correct the share deficit.

*Effects*:
* If `balanceDeltaWei` is zero, do nothing
* If `balanceDeltaWei` is positive:
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

#### `increaseBurnableShares`

```solidity
/**
 * @notice Increase the amount of burnable shares for a given Strategy. This is called by the DelegationManager
 * when an operator is slashed in EigenLayer.
 * @param strategy The strategy to burn shares in.
 * @param addedSharesToBurn The amount of added shares to burn.
 * @dev This function is only called by the DelegationManager when an operator is slashed.
 */
function increaseBurnableShares(
    IStrategy strategy, 
    uint256 addedSharesToBurn
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when an operator is slashed, calculating the number of slashable shares and marking them for burning here.

Unlike in the `StrategyManager`, there is no current mechanism to burn these shares, as burning requires the Pectra hard fork to be able to eject validators. This will be added in a future update.

*Effects*:
* Increases `burnableShares` for the beacon chain ETH strategy by `addedSharesToBurn`

*Requirements*:
* Can only be called by the `DelegationManager`