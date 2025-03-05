## DelegationManager

| File | Notes |
| -------- | -------- |
| [`DelegationManager.sol`](../../src/contracts/core/DelegationManager.sol) | |
| [`DelegationManagerStorage.sol`](../../src/contracts/core/DelegationManagerStorage.sol) | state variables |
| [`IDelegationManager.sol`](../../src/contracts/interfaces/IDelegationManager.sol) | interface |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`PermissionControllerMixin.sol`](../../src/contracts/mixins/PermissionControllerMixin.sol) | account delegation |
| [`SignatureUtils.sol`](../../src/contracts/mixins/SignatureUtils.sol) | signature validation |
| [`Pausable.sol`](../../src/contracts/permissions/Pausable.sol) | |
| [`SlashingLib.sol`](../../src/contracts/libraries/SlashingLib.sol) | slashing math |
| [`Snapshots.sol`](../../src/contracts/libraries/Snapshots.sol) | historical state |

## Prior Reading

* [ELIP-002: Slashing via Unique Stake and Operator Sets](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md)
* [Shares Accounting](./accounting/SharesAccounting.md)

## Overview

The `DelegationManager` is the intersection between the two sides of the protocol. It (i) allows stakers to delegate/undelegate to/from operators, (ii) handles withdrawals and withdrawal processing for assets in both the `StrategyManager` and `EigenPodManager`, and (iii) manages accounting around slashing for stakers and operators.

When operators are slashed by AVSs, it receives share burning directives from the `AllocationManager`. When stakers deposit assets using the `StrategyManager/EigenPodManager`, it tracks share/delegation accounting changes. The `DelegationManager` combines inputs from both sides of the protocol into a staker's "deposit scaling factor," which serves as the primary conversion vehicle between a staker's _raw deposited assets_ and the _amount they can withdraw_.

The `DelegationManager's` responsibilities can be broken down into the following concepts:
* [Becoming an Operator](#becoming-an-operator)
* [Delegation and Withdrawals](#delegation-and-withdrawals)
* [Slashing and Accounting](#slashing-and-accounting)

## Parameterization

* `MIN_WITHDRAWAL_DELAY_BLOCKS`: The delay in blocks before withdrawals can be completed.
    * Mainnet: `100800 blocks` (14 days).
    * Testnet: `50 blocks` (10 minutes).
* `beaconChainETHStrategy`: a pseudo strategy used to represent beacon chain ETH internally. This is not a real contract!
    * Value: `0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0`

---

## Becoming an Operator

The `DelegationManager` tracks operator-related state in the following mappings:

```solidity
/// @notice Returns the `operator` a `staker` is delegated to, or address(0) if not delegated.
/// Note: operators are delegated to themselves
mapping(address staker => address operator) public delegatedTo;

/// @notice Returns the operator details for a given `operator`.
/// Note: two of the `OperatorDetails` fields are deprecated. The only relevant field
/// is `OperatorDetails.delegationApprover`.
mapping(address operator => OperatorDetails) internal _operatorDetails;

/**
 * @notice Tracks the current balance of shares an `operator` is delegated according to each `strategy`. 
 * Updated by both the `StrategyManager` and `EigenPodManager` when a staker's delegatable balance changes,
 * and by the `AllocationManager` when the `operator` is slashed.
 *
 * @dev The following invariant should hold for each `strategy`:
 *
 * operatorShares[operator] = sum(withdrawable shares of all stakers delegated to operator)
 */
mapping(address operator => mapping(IStrategy strategy => uint256 shares)) public operatorShares;
```

**Methods**:
* [`DelegationManager.registerAsOperator`](#registerasoperator)
* [`DelegationManager.modifyOperatorDetails`](#modifyoperatordetails)
* [`DelegationManager.updateOperatorMetadataURI`](#updateoperatormetadatauri)

#### `registerAsOperator`

```solidity
/**
 * @notice Registers the caller as an operator in EigenLayer.
 * @param initDelegationApprover is an address that, if set, must provide a signature when stakers delegate
 * to an operator.
 * @param allocationDelay The delay before allocations take effect.
 * @param metadataURI is a URI for the operator's metadata, i.e. a link providing more details on the operator.
 *
 * @dev Once an operator is registered, they cannot 'deregister' as an operator, and they will forever be considered "delegated to themself".
 * @dev This function will revert if the caller is already delegated to an operator.
 * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
 */
function registerAsOperator(
    address initDelegationApprover,
    uint32 allocationDelay,
    string calldata metadataURI
) external nonReentrant;
```

Registers the caller as an operator in EigenLayer. The new operator provides the following input parameters:
* `address initDelegationApprover`: *(OPTIONAL)* if set to a non-zero address, this address must sign and approve new delegation from stakers to this operator (See [`delegateTo`](#delegateto))
* `uint32 allocationDelay`: the delay (in blocks) before slashable stake allocations will take effect. This is passed to the `AllocationManager` (See [`AllocationManager.md#setAllocationDelay`](./AllocationManager.md#setallocationdelay))
* `string calldata metadataURI`: emits this input in the event `OperatorMetadataURIUpdated`. Does not store the value anywhere.

`registerAsOperator` cements the operator's delegation approver and allocation delay in storage, and self-delegates the operator to themselves - permanently marking the caller as an operator. They cannot "deregister" as an operator - however, if they have deposited funds, they can still withdraw them (See [Delegation and Withdrawals](#delegation-and-withdrawals)).

*Effects*:
* Sets `_operatorDetails[operator].delegationApprover`. Note that the other `OperatorDetails` fields are deprecated; only the `delegationApprover` is used.
* Delegates the operator to themselves
    * Tabulates any deposited shares across the `EigenPodManager` and `StrategyManager`, and delegates these shares to themselves
    * For each strategy in which the operator holds assets, updates the operator's `depositScalingFactor` for that strategy

*Requirements*:
* Caller MUST NOT already be delegated
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`
* For each strategy in which the operator holds assets, their `slashingFactor` for that strategy MUST be non-zero.

#### `modifyOperatorDetails`

```solidity
/**
 * @notice Updates an operator's stored `delegationApprover`.
 * @param operator is the operator to update the delegationApprover for
 * @param newDelegationApprover is the new delegationApprover for the operator
 *
 * @dev The caller must have previously registered as an operator in EigenLayer.
 */
function modifyOperatorDetails(
    address operator, 
    address newDelegationApprover
) 
    external 
    checkCanCall(operator)
    nonReentrant
```

_Note: this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

Allows an operator to update their stored `delegationApprover`.

*Requirements*:
* `address operator` MUST already be an operator.
* Caller MUST be authorized: either the operator themselves, or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))

#### `updateOperatorMetadataURI`

```solidity
/**
 * @notice Called by an operator to emit an `OperatorMetadataURIUpdated` event indicating the information has updated.
 * @param operator The operator to update metadata for
 * @param metadataURI The URI for metadata associated with an operator
 * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
 */
function updateOperatorMetadataURI(
    address operator, 
    string calldata metadataURI
) 
    external 
    checkCanCall(operator)
```

_Note: this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

Allows an operator to emit an `OperatorMetadataURIUpdated` event. No other state changes occur.

*Requirements*:
* `address operator` MUST already be an operator.
* Caller MUST be authorized: either the operator themselves, or an admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))

---

## Delegation and Withdrawals

**Concepts**:
* [Shares Accounting](./accounting/SharesAccounting.md)
* [Legacy and Post-Slashing Withdrawals](#legacy-and-post-slashing-withdrawals)
* [Slashing Factors and Scaling Shares](#slashing-factors-and-scaling-shares)

**Methods**:
* [`DelegationManager.delegateTo`](#delegateto)
* [`DelegationManager.undelegate`](#undelegate)
* [`DelegationManager.redelegate`](#redelegate)
* [`DelegationManager.queueWithdrawals`](#queuewithdrawals)
* [`DelegationManager.completeQueuedWithdrawal`](#completequeuedwithdrawal)
* [`DelegationManager.completeQueuedWithdrawals`](#completequeuedwithdrawals)

#### Legacy and Post-Slashing Withdrawals

The `DelegationManager` tracks withdrawal-related state in the following mappings:

```solidity
/**
 * @dev A struct representing an existing queued withdrawal. After the withdrawal delay has elapsed, this withdrawal can be completed via `completeQueuedWithdrawal`.
 * A `Withdrawal` is created by the `DelegationManager` when `queueWithdrawals` is called. The `withdrawalRoots` hashes returned by `queueWithdrawals` can be used
 * to fetch the corresponding `Withdrawal` from storage (via `queuedWithdrawals`).
 *
 * @param staker The address that queued the withdrawal
 * @param delegatedTo The address that the staker was delegated to at the time the withdrawal was queued. Used to determine if additional slashing occurred before
 * this withdrawal became completable.
 * @param withdrawer The address that will call the contract to complete the withdrawal. Note that this will always equal `staker`; alternate withdrawers are not
 * supported at this time.
 * @param nonce The staker's `cumulativeWithdrawalsQueued` at time of queuing. Used to ensure withdrawals have unique hashes.
 * @param startBlock The block number when the withdrawal was queued.
 * @param strategies The strategies requested for withdrawal when the withdrawal was queued
 * @param scaledShares The staker's deposit shares requested for withdrawal, scaled by the staker's `depositScalingFactor`. Upon completion, these will be
 * scaled by the appropriate slashing factor as of the withdrawal's completable block. The result is what is actually withdrawable.
 */
struct Withdrawal {
    address staker;
    address delegatedTo;
    address withdrawer;
    uint256 nonce;
    uint32 startBlock;
    IStrategy[] strategies;
    uint256[] scaledShares;
}

/// @dev Returns whether a withdrawal is pending for a given `withdrawalRoot`.
/// @dev This variable will be deprecated in the future, values should only be read or deleted.
mapping(bytes32 withdrawalRoot => bool pending) public pendingWithdrawals;

/// @notice Returns the total number of withdrawals that have been queued for a given `staker`.
/// @dev This only increments (doesn't decrement), and is used to help ensure that otherwise identical withdrawals have unique hashes.
mapping(address staker => uint256 totalQueued) public cumulativeWithdrawalsQueued;

/// @notice Returns a list of queued withdrawals for a given `staker`.
/// @dev Entries are removed when the withdrawal is completed.
/// @dev This variable only reflects withdrawals that were made after the slashing release.
mapping(address staker => EnumerableSet.Bytes32Set withdrawalRoots) internal _stakerQueuedWithdrawalRoots;

/// @notice Returns the details of a queued withdrawal given by `withdrawalRoot`.
/// @dev This variable only reflects withdrawals that were made after the slashing release.
mapping(bytes32 withdrawalRoot => Withdrawal withdrawal) public queuedWithdrawals;

/// @notice Contains history of the total cumulative staker withdrawals for an operator and a given strategy.
/// Used to calculate burned StrategyManager shares when an operator is slashed.
/// @dev Stores scaledShares instead of total withdrawn shares to track current slashable shares, dependent on the maxMagnitude
mapping(address operator => mapping(IStrategy strategy => Snapshots.DefaultZeroHistory)) internal
    _cumulativeScaledSharesHistory;
```

Prior to the slashing release, withdrawals were only stored as hashes in the `pendingWithdrawals` mapping. 

With the slashing release, withdrawals are now stored entirely in state, and two new mappings have been added to support this:
* `_stakedQueuedWithdrawalRoots`: a list of all the currently-queued withdrawal hashes belonging to a staker
* `queuedWithdrawals`: maps queued withdrawal hash to `Withdrawal` struct

Legacy withdrawals remain completable using the same methods as new withdrawals. The primary difference between the two is that it is not possible to query the corresponding `Withdrawal` struct for a legacy withdrawal hash. When determining what `Withdrawal` struct to supply to the contract to complete a legacy withdrawal, the caller will need to derive the original `Withdrawal` struct generated when the withdrawal was queued.

#### Slashing Factors and Scaling Shares

_See the [Shares Accounting](./accounting/SharesAccounting.md) doc for a more thorough explanation with examples._

Throughout the `DelegationManager`, a staker's _deposit shares_ can be converted into their current _withdrawable shares_ by applying two factors: the _slashing factor_ and the _deposit scaling factor_. These two values are scaling factors that act as numerators when scaling shares. By default, these values start at `1 WAD` (`1e18`). `1 WAD` also acts as the denominator when scaling.

```solidity
/// @dev All scaling factors have `1e18` as an initial/default value. This value is represented
/// by the constant `WAD`, which is used to preserve precision with uint256 math.
///
/// When applying scaling factors, they are typically multiplied/divided by `WAD`, allowing this
/// constant to act as a "1" in mathematical formulae.
uint64 constant WAD = 1e18;
```

The _deposit scaling factor_ is represented in `DelegationManager` storage, and can be thought of as a way to normalize newly-deposited shares using the _current_ slashing factor, so that _future_ withdrawals can be scaled appropriately if the slashing factor has changed:

```solidity
/*
 * There are 2 types of shares:
 *      1. deposit shares
 *          - These can be converted to an amount of tokens given a strategy
 *              - by calling `sharesToUnderlying` on the strategy address (they're already tokens 
 *              in the case of EigenPods)
 *          - These live in the storage of the EigenPodManager and individual StrategyManager strategies 
 *      2. withdrawable shares
 *          - For a staker, this is the amount of shares that they can withdraw
 *          - For an operator, the shares delegated to them are equal to the sum of their stakers'
 *            withdrawable shares
 *
 * Along with a slashing factor, the DepositScalingFactor is used to convert between the two share types.
 */
struct DepositScalingFactor {
    uint256 _scalingFactor;
}

/// @notice Returns the scaling factor applied to a `staker` for a given `strategy`
mapping(address staker => mapping(IStrategy strategy => DepositScalingFactor)) internal _depositScalingFactor;
```

Calculating the _slashing factor_ varies depending on the strategy in question. _For all strategies_, the slashing factor is the max magnitude of the staker's delegated `operator` in the `AllocationManager` (See [Max vs Encumbered Magnitude](./AllocationManager.md#max-vs-encumbered-magnitude)). If the staker is NOT delegated, this is `WAD` (aka "1").

_For the `beaconChainETHStrategy`_, the slashing factor _also_ includes the staker's `beaconChainSlashingFactor`, which acts like the `operator's` max magnitude, but for a staker's beacon chain assets. This means that, for the `beaconChainETHStrategy` specifically, _slashing factors_ can be applied because of EITHER/BOTH:
* `operator` got slashed for this strategy by an AVS
* `staker` got slashed on the beacon chain

From `DelegationManager.sol`:

```solidity
/// @dev Calculate the amount of slashing to apply to the staker's shares
function _getSlashingFactor(
    address staker,
    IStrategy strategy,
    uint64 operatorMaxMagnitude
) internal view returns (uint256) {
    if (strategy == beaconChainETHStrategy) {
        uint64 beaconChainSlashingFactor = eigenPodManager.beaconChainSlashingFactor(staker);
        return operatorMaxMagnitude.mulWad(beaconChainSlashingFactor);
    }

    return operatorMaxMagnitude;
}
```

#### `delegateTo`

```solidity
// @notice Struct that bundles together a signature and an expiration time for the signature. Used primarily for stack management.
struct SignatureWithExpiry {
    // the signature itself, formatted as a single bytes object
    bytes signature;
    // the expiration timestamp (UTC) of the signature
    uint256 expiry;
}

/**
 * @notice Caller delegates their stake to an operator.
 * @param operator The account (`msg.sender`) is delegating its assets to for use in serving applications built on EigenLayer.
 * @param approverSignatureAndExpiry (optional) Verifies the operator approves of this delegation
 * @param approverSalt (optional) A unique single use value tied to an individual signature.
 * @dev The signature/salt are used ONLY if the operator has configured a delegationApprover.
 * If they have not, these params can be left empty.
 */
function delegateTo(
    address operator, 
    SignatureWithExpiry memory approverSignatureAndExpiry, 
    bytes32 approverSalt
) 
    external
    nonReentrant
```

Allows a staker to delegate their assets to an operator. Delegation is all-or-nothing: when a staker delegates to an operator, they delegate ALL their assets. Stakers can only be delegated to one operator at a time.

For each strategy the staker has deposit shares in, the `DelegationManager` will:
* Query the staker's deposit shares from the `StrategyManager/EigenPodManager`
* Get the slashing factor for this `(staker, operator, strategy)` and use it to update the staker's deposit scaling factor (See [Slashing Factors and Scaling Shares](#slashing-factors-and-scaling-shares))
* Add the deposit shares to the operator's `operatorShares` directly. _Note_ that the initial delegation to an operator is a special case where deposit shares == withdrawable shares.

*Effects*:
* Delegates the caller to the `operator`
    * Tabulates any deposited shares across the `EigenPodManager` and `StrategyManager`, and delegates these shares to the `operator`
    * For each strategy in which the caller holds assets, updates the caller's `depositScalingFactor` for that strategy

*Requirements*:
* The caller MUST NOT already be delegated to an operator
* The `operator` MUST already be an operator
* If the `operator` has a `delegationApprover`, the caller MUST provide a valid `approverSignatureAndExpiry` and `approverSalt`
* Pause status MUST NOT be set: `PAUSED_NEW_DELEGATION`
* For each strategy in which the staker holds assets, the `slashingFactor` for that strategy MUST be non-zero.

#### `undelegate`

```solidity
/**
 * @notice Undelegates the staker from their operator and queues a withdrawal for all of their shares
 * @param staker The account to be undelegated
 * @return withdrawalRoots The roots of the newly queued withdrawals, if a withdrawal was queued. Returns 
 * an empty array if none was queued.
 *
 * @dev Reverts if the `staker` is also an operator, since operators are not allowed to undelegate from themselves.
 * @dev Reverts if the caller is not the staker, nor the operator who the staker is delegated to, nor the operator's specified "delegationApprover"
 * @dev Reverts if the `staker` is not delegated to an operator
 */
function undelegate(
    address staker
) 
    external
    nonReentrant
    returns (bytes32[] memory withdrawalRoots);
```

_Note: this method can be called directly by an operator, or by a caller authorized by the operator. See [`PermissionController.md`](../permissions/PermissionController.md) for details._

`undelegate` can be called EITHER by a staker to undelegate themselves, OR by an operator to force-undelegate a staker from them. Force-undelegation is primarily useful if an operator has a `delegationApprover`, as this role is the only way to prevent a staker from delegating back to the operator once force-undelegated.

Undelegation immediately sets the staker's delegated operator to 0, decreases the prior operator's delegated shares, and queues withdrawals for all of the staker's deposited assets. For UX reasons, one withdrawal is queued for each strategy in which the staker has deposited assets. Queued withdrawals mimic the behavior of the [`queueWithdrawals`](#queuewithdrawals) method; see that method's documentation for details.

Just as with a normal queued withdrawal, these withdrawals can be completed by the staker after `MIN_WITHDRAWAL_DELAY_BLOCKS`. These withdrawals do not require the staker to "fully exit" from the system -- the staker may choose to keep their assets in the system once withdrawals are completed (See [`completeQueuedWithdrawal`](#completequeuedwithdrawal) for details).

*Effects*: 
* The `staker` is undelegated from their operator
* If the `staker` has no deposit shares, there is no withdrawal queued or further effects
* For each strategy held by the `staker`, a `Withdrawal` is queued:
    * _Deposit shares_ are removed from the staker's deposit share balances
        * See [`EigenPodManager.removeDepositShares`](./EigenPodManager.md#removedepositshares)
        * See [`StrategyManager.removeDepositShares`](./StrategyManager.md#removedepositshares)
    * _Deposit shares_ are converted to _withdrawable shares_ (See [Slashing Factors and Scaling Shares](#slashing-factors-and-scaling-shares)). These are decremented from the operator's delegated shares.
    * _Deposit shares_ are converted to _scaled shares_  (See [Shares Accounting - Queue Withdrawals](./accounting/SharesAccounting.md#queue-withdrawal)), which are stored in the `Withdrawal` struct
    * _Scaled shares_ are pushed to `_cumulativeScaledSharesHistory`, which is used for burning slashed shares
    * The `Withdrawal` is saved to storage
        * The hash of the `Withdrawal` is marked as "pending"
        * The hash of the `Withdrawal` is set in a mapping to the `Withdrawal` struct itself
        * The hash of the `Withdrawal` is pushed to `_stakerQueuedWithdrawalRoots`
    * The staker's withdrawal nonce is increased by 1 for each `Withdrawal`

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_ENTER_WITHDRAWAL_QUEUE`
* `staker` MUST be delegated to an operator
* `staker` MUST NOT be an operator
* `staker` parameter MUST NOT be zero
* Caller MUST be authorized: either the `staker` themselves, their operator, their operator's `delegationApprover`, or their operator's admin/appointee (see [`PermissionController.md`](../permissions/PermissionController.md))
* See [`EigenPodManager.removeDepositShares`](./EigenPodManager.md#removedepositshares)
* See [`StrategyManager.removeDepositShares`](./StrategyManager.md#removedepositshares)

#### `redelegate`

```solidity
/**
 * @notice Undelegates the staker from their current operator, and redelegates to `newOperator`
 * Queues a withdrawal for all of the staker's withdrawable shares. These shares will only be
 * delegated to `newOperator` AFTER the withdrawal is completed.
 * @dev This method acts like a call to `undelegate`, then `delegateTo`
 * @param newOperator the new operator that will be delegated all assets
 * @dev NOTE: the following 2 params are ONLY checked if `newOperator` has a `delegationApprover`.
 * If not, they can be left empty.
 * @param newOperatorApproverSig A signature from the operator's `delegationApprover`
 * @param approverSalt A unique single use value tied to the approver's signature
 */
 function redelegate(
    address newOperator,
    SignatureWithExpiry memory newOperatorApproverSig,
    bytes32 approverSalt
) 
    external 
    returns (bytes32[] memory withdrawalRoots);
```

`redelegate` is a convenience method that combines the effects of `undelegate` and `delegateTo`. `redelegate` allows a staker to switch their delegated operator to `newOperator` with a single call. **Note**, though, that the staker's assets will not be fully delegated to `newOperator` until the withdrawals queued during the undelegation portion of this method are completed.

*Effects*: 
* See [`delegateTo`](#delegateto) and [`undelegate`](#undelegate)

*Requirements*:
* See [`delegateTo`](#delegateto) and [`undelegate`](#undelegate)

#### `queueWithdrawals`

```solidity
/**
 * @param strategies The strategies to withdraw from
 * @param depositShares For each strategy, the number of deposit shares to withdraw. Deposit shares can
 * be queried via `getDepositedShares`.
 * NOTE: The number of shares ultimately received when a withdrawal is completed may be lower depositShares
 * if the staker or their delegated operator has experienced slashing.
 * @param __deprecated_withdrawer This field is ignored. The only party that may complete a withdrawal
 * is the staker that originally queued it. Alternate withdrawers are not supported.
 */
struct QueuedWithdrawalParams {
    IStrategy[] strategies;
    uint256[] depositShares;
    address __deprecated_withdrawer;
}

/**
 * @notice Allows a staker to queue a withdrawal of their deposit shares. The withdrawal can be 
 * completed after the MIN_WITHDRAWAL_DELAY_BLOCKS via either of the completeQueuedWithdrawal methods.
 * 
 * While in the queue, these shares are removed from the staker's balance, as well as from their operator's
 * delegated share balance (if applicable). Note that while in the queue, deposit shares are still subject
 * to slashing. If any slashing has occurred, the shares received may be less than the queued deposit shares.
 *
 * @dev To view all the staker's strategies/deposit shares that can be queued for withdrawal, see `getDepositedShares`
 * @dev To view the current conversion between a staker's deposit shares and withdrawable shares, see `getWithdrawableShares`
 */
function queueWithdrawals(
    QueuedWithdrawalParams[] calldata queuedWithdrawalParams
) 
    external 
    onlyWhenNotPaused(PAUSED_ENTER_WITHDRAWAL_QUEUE)
    nonReentrant
    returns (bytes32[] memory)
```

Allows the caller to queue their deposit shares for withdrawal across any strategy. Withdrawals can be completed after `MIN_WITHDRAWAL_DELAY_BLOCKS`, by calling [`completeQueuedWithdrawal`](#completequeuedwithdrawal). This method accepts _deposit shares_ as input - however, the amounts received upon completion may be lower if the staker has experienced slashing (See [Shares Accounting](./accounting/SharesAccounting.md) and [Slashing Factors and Scaling Shares](#slashing-factors-and-scaling-shares)).

For each `QueuedWithdrawalParams` passed as input, a `Withdrawal` is created in storage (See [Legacy and Post-Slashing Withdrawals](#legacy-and-post-slashing-withdrawals) for details on structure and querying). Queueing a withdrawal involves multiple transformations to a staker's _deposit shares_, serving a few different purposes:
* The raw _deposit shares_ are removed from the staker's deposit share balance in the corresponding share manager (`EigenPodManager` or `StrategyManager`).
* _Scaled shares_ are calculated by applying the staker's _deposit scaling factor_ to their _deposit shares_. Scaled shares:
    * are stored in the `Withdrawal` itself and used during withdrawal completion
    * are added to the operator's `cumulativeScaledSharesHistory`, where they can be burned if slashing occurs while the withdrawal is in the queue
* _Withdrawable shares_ are calculated by applying both the staker's _deposit scaling factor_ AND any appropriate _slashing factor_ to the staker's _deposit shares_. These "currently withdrawable shares" are removed from the operator's delegated shares (if applicable).

Note that the `QueuedWithdrawalParams.__deprecated_withdrawer` field is ignored. Originally, this was used to create withdrawals that could be completed by a third party. This functionality was removed during the M2 release due to growing concerns over the phish risk this presented. Until the slashing release, this field was explicitly checked for equivalence with `msg.sender`; however, at present it is ignored. All `Withdrawals` are created with `withdrawer == staker` regardless of this field's value.

*Effects*:
* For each `QueuedWithdrawalParams` element:
    * _Deposit shares_ are removed from the staker's deposit share balances
        * See [`EigenPodManager.removeDepositShares`](./EigenPodManager.md#removedepositshares)
        * See [`StrategyManager.removeDepositShares`](./StrategyManager.md#removedepositshares)
    * _Deposit shares_ are converted to _withdrawable shares_ (See [Slashing Factors and Scaling Deposits](#slashing-factors-and-scaling-shares)). These are decremented from their operator's delegated shares (if applicable)
    * _Deposit shares_ are converted to _scaled shares_  (See [Shares Accounting - Queue Withdrawals](./accounting/SharesAccounting.md#queue-withdrawal)), which are stored in the `Withdrawal` struct
    * If the caller is delegated to an operator, _scaled shares_ are pushed to that operator's `_cumulativeScaledSharesHistory`, which may be burned if slashing occurs.
    * The `Withdrawal` is saved to storage
        * The hash of the `Withdrawal` is marked as "pending"
        * The hash of the `Withdrawal` is set in a mapping to the `Withdrawal` struct itself
        * The hash of the `Withdrawal` is pushed to `_stakerQueuedWithdrawalRoots`
    * The staker's withdrawal nonce is increased by 1

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_ENTER_WITHDRAWAL_QUEUE`
* For each `QueuedWithdrawalParams` element:
    * `strategies.length` MUST equal `depositShares.length`
    * The `withdrawer` MUST equal `msg.sender`
    * `strategies.length` MUST NOT be equal to 0
    * See [`EigenPodManager.removeDepositShares`](./EigenPodManager.md#removedepositshares)
    * See [`StrategyManager.removeDepositShares`](./StrategyManager.md#removedepositshares)

#### `completeQueuedWithdrawal`

```solidity
/**
 * @dev A struct representing an existing queued withdrawal. After the withdrawal delay has elapsed, this withdrawal can be completed via `completeQueuedWithdrawal`.
 * A `Withdrawal` is created by the `DelegationManager` when `queueWithdrawals` is called. The `withdrawalRoots` hashes returned by `queueWithdrawals` can be used
 * to fetch the corresponding `Withdrawal` from storage (via `getQueuedWithdrawal`).
 *
 * @param staker The address that queued the withdrawal
 * @param delegatedTo The address that the staker was delegated to at the time the withdrawal was queued. Used to determine if additional slashing occurred before
 * this withdrawal became completable.
 * @param withdrawer The address that will call the contract to complete the withdrawal. Note that this will always equal `staker`; alternate withdrawers are not
 * supported at this time.
 * @param nonce The staker's `cumulativeWithdrawalsQueued` at time of queuing. Used to ensure withdrawals have unique hashes.
 * @param startBlock The block number when the withdrawal was queued.
 * @param strategies The strategies requested for withdrawal when the withdrawal was queued
 * @param scaledShares The staker's deposit shares requested for withdrawal, scaled by the staker's `depositScalingFactor`. Upon completion, these will be
 * scaled by the appropriate slashing factor as of the withdrawal's completable block. The result is what is actually withdrawable.
 */
struct Withdrawal {
    address staker;
    address delegatedTo;
    address withdrawer;
    uint256 nonce;
    uint32 startBlock;
    IStrategy[] strategies;
    uint256[] scaledShares;
}

/**
 * @notice Used to complete a queued withdrawal
 * @param withdrawal The withdrawal to complete
 * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `withdrawal.strategies` array.
 * @param tokens For each `withdrawal.strategies`, the underlying token of the strategy
 * NOTE: if `receiveAsTokens` is false, the `tokens` array is unused and can be filled with default values. However, `tokens.length` MUST still be equal to `withdrawal.strategies.length`.
 * NOTE: For the `beaconChainETHStrategy`, the corresponding `tokens` value is ignored (can be 0).
 * @param receiveAsTokens If true, withdrawn shares will be converted to tokens and sent to the caller. If false, the caller receives shares that can be delegated to an operator.
 * NOTE: if the caller receives shares and is currently delegated to an operator, the received shares are
 * automatically delegated to the caller's current operator.
 */
function completeQueuedWithdrawal(
    Withdrawal calldata withdrawal,
    IERC20[] calldata tokens,
    bool receiveAsTokens
) 
    external 
    onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE)
    nonReentrant
```

`MIN_WITHDRAWAL_DELAY_BLOCKS` after queueing, a staker can complete a `Withdrawal` by calling this method. The staker can elect to receive _either_ tokens OR shares depending on the value of the `receiveAsTokens` parameter. 

Before processing a withdrawal, this method will calculate the slashing factor at the withdrawal's completion block (`withdrawal.startBlock + MIN_WITHDRAWAL_DELAY_BLOCKS`), according to the operator that was delegated to when the withdrawal was queued (`withdrawal.delegatedTo`). This slashing factor is used to determine if any additional slashing occurred while the withdrawal was in the queue. If so, this slashing is applied now.

For each `Withdrawal`, `withdrawal.scaledShares` are converted into _withdrawable shares_, accounting for any slashing that occurred during the withdrawal period (See [Shares Accounting - Complete Withdrawal](./accounting/SharesAccounting.md#complete-withdrawal)).

If the staker chooses to receive the withdrawal _as tokens_, the withdrawable shares are converted to tokens via the corresponding share manager (`EigenPodManager`/`StrategyManager`), and sent to the caller.

If the staker chooses to receive the withdrawal _as shares_, the withdrawable shares are credited to the staker via the corresponding share manager (`EigenPodManager`/`StrategyManager`). Additionally, if the caller is delegated to an operator, the new slashing factor for the given `(staker, operator, strategy)` determines how many shares are awarded to the operator (and how the staker's deposit scaling factor is updated) (See [Slashing Factors and Scaling Shares](#slashing-factors-and-scaling-shares)). In receiving the withdrawal as shares, this amount is credited as deposit shares for the staker. Due to known rounding error, the amount of withdrawable shares after completing the withdrawal may be slightly less than what was originally withdrawable.

**Note:** if the staker (i) receives the withdrawal as shares, (ii) has `MAX_STAKER_STRATEGY_LIST_LENGTH` unique deposit strategies in the `StrategyManager`, and (iii) is withdrawing to a `StrategyManager` strategy in which they do not currently have shares, this will revert. The staker cannot withdraw such that their `stakerStrategyList` length exceeds the maximum; this withdrawal will have to be completed as tokens instead.

**Note:** if the staker receives a `beaconChainETHStrategy` withdrawal as tokens, the staker's `EigenPod` MUST have sufficient `withdrawableExecutionLayerGwei` to honor the withdrawal.

*Effects*:
* The hash of the `Withdrawal` is removed from the pending withdrawals
* The hash of the `Withdrawal` is removed from the enumerable set of staker queued withdrawals
* The `Withdrawal` struct is removed from the queued withdrawals 
* If `receiveAsTokens`:
    * See [`StrategyManager.withdrawSharesAsTokens`](./StrategyManager.md#withdrawsharesastokens)
    * See [`EigenPodManager.withdrawSharesAsTokens`](./EigenPodManager.md#withdrawsharesastokens)
* If `!receiveAsTokens`:
    * Withdrawable shares are awarded to the caller and delegated to the caller's current operator if applicable
    * See [`StrategyManager.addShares`](./StrategyManager.md#addshares)
    * See [`EigenPodManager.addShares`](./EigenPodManager.md#addshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EXIT_WITHDRAWAL_QUEUE`
* `tokens.length` must equal `withdrawal.strategies.length`
* Caller MUST be the `staker/withdrawer` specified in the `Withdrawal`
* At least `MIN_WITHDRAWAL_DELAY_BLOCKS` MUST have passed before `completeQueuedWithdrawal` is called
* The hash of the passed-in `Withdrawal` MUST correspond to a pending withdrawal
* If `receiveAsTokens`:
    * The caller MUST pass in the underlying `IERC20[] tokens` being withdrawn in the appropriate order according to the strategies in the `Withdrawal`.
    * See [`StrategyManager.withdrawSharesAsTokens`](./StrategyManager.md#withdrawsharesastokens)
    * See [`EigenPodManager.withdrawSharesAsTokens`](./EigenPodManager.md#withdrawsharesastokens)
* If `!receiveAsTokens`:
    * See [`StrategyManager.addShares`](./StrategyManager.md#addshares)
    * See [`EigenPodManager.addShares`](./EigenPodManager.md#addshares)

#### `completeQueuedWithdrawals`

```solidity
/**
 * @notice Used to complete multiple queued withdrawals
 * @param withdrawals Array of Withdrawals to complete. See `completeQueuedWithdrawal` for the usage of a single Withdrawal.
 * @param tokens Array of tokens for each Withdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
 * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `completeQueuedWithdrawal` for the usage of a single boolean.
 * @dev See `completeQueuedWithdrawal` for relevant dev tags
 */
function completeQueuedWithdrawals(
    Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    bool[] calldata receiveAsTokens
) 
    external 
    onlyWhenNotPaused(PAUSED_EXIT_WITHDRAWAL_QUEUE) 
    nonReentrant
```

This method is the plural version of [`completeQueuedWithdrawal`](#completequeuedwithdrawal).

---

## Slashing and Accounting

These methods are all called by other system contracts: the `AllocationManager` calls `slashOperatorShares` during slashing, and the `EigenPodManager/StrategyManager` call `increaseDelegatedShares/decreaseDelegatedShares` when stakers' deposit shares (or when beacon chain balance decreases occur).

**Methods**:
* [`DelegationManager.slashOperatorShares`](#slashoperatorshares)
* [`DelegationManager.increaseDelegatedShares`](#increasedelegatedshares)
* [`DelegationManager.decreaseDelegatedShares`](#decreasedelegatedshares)

#### `slashOperatorShares`

```solidity
/**
 * @notice Decreases the operators shares in storage after a slash and increases the burnable shares by calling
 * into either the StrategyManager or EigenPodManager (if the strategy is beaconChainETH).
 * @param operator The operator to decrease shares for
 * @param strategy The strategy to decrease shares for
 * @param prevMaxMagnitude the previous maxMagnitude of the operator
 * @param newMaxMagnitude the new maxMagnitude of the operator
 * @dev Callable only by the AllocationManager
 * @dev Note: Assumes `prevMaxMagnitude <= newMaxMagnitude`. This invariant is maintained in
 * the AllocationManager.
 */
function slashOperatorShares(
    address operator,
    IStrategy strategy,
    uint64 prevMaxMagnitude,
    uint64 newMaxMagnitude
) 
    external
    onlyAllocationManager
    nonReentrant
```

_See [Shares Accounting - Slashing](https://github.com/Layr-Labs/eigenlayer-contracts/blob/slashing-magnitudes/docs/core/accounting/SharesAccounting.md#slashing) for a description of the accounting in this method._

This method is called by the `AllocationManager` when processing an AVS's slash of an operator. Slashing occurs instantly, with this method directly reducing the operator's delegated shares proportional to the slash.

Additionally, any _slashable shares_ in the withdrawal queue are marked for burning according to the same slashing proportion (shares in the withdrawal queue remain slashable for `MIN_WITHDRAWAL_DELAY_BLOCKS`). For the slashed strategy, the corresponding share manager (`EigenPodManager/StrateyManager`) is called, increasing the burnable shares for that strategy.

**Note**: native ETH does not currently possess a burning mechanism, as this requires Pectra to be able to force exit validators. Currently, slashing for the `beaconChainETHStrategy` is realized by modifying the amount stakers are able to withdraw.

*Effects*:
* The `operator's` `operatorShares` are reduced for the given `strategy`, according to the proportion given by `prevMaxMagnitude` and `newMaxMagnitude`
* Any slashable shares in the withdrawal queue are marked for burning according to the same proportion
* See [`StrategyManager.increaseBurnableShares`](./StrategyManager.md#increaseBurnableShares)
* See [`EigenPodManager.increaseBurnableShares`](./EigenPodManager.md#increaseBurnableShares)


*Requirements*:
* The amount slashed from the operator must not result in underflow of their `operatorShares` for the given `strategy`

#### `increaseDelegatedShares`

```solidity
/**
 * @notice Called by a share manager when a staker's deposit share balance in a strategy increases.
 * This method delegates any new shares to an operator (if applicable), and updates the staker's 
 * deposit scaling factor regardless.
 * @param staker The address whose deposit shares have increased
 * @param strategy The strategy in which shares have been deposited
 * @param prevDepositShares The number of deposit shares the staker had in the strategy prior to the increase
 * @param addedShares The number of deposit shares added by the staker
 *
 * @dev Note that if either the staker's current operator has been slashed 100% for `strategy`, OR the
 * staker has been slashed 100% on the beacon chain such that the calculated slashing factor is 0, this
 * method WILL REVERT.
 */
function increaseDelegatedShares(
    address staker,
    IStrategy strategy,
    uint256 prevDepositShares,
    uint256 addedShares
) 
    external
    onlyStrategyManagerOrEigenPodManager
    nonReentrant
```

Called by either the `StrategyManager` or `EigenPodManager` when a staker's deposit shares for one or more strategies increase.

If the staker is delegated to an operator, the new deposit shares are directly added to that operator's `operatorShares`. Regardless of delegation status, the staker's deposit scaling factor is updated.

**Note** that if either the staker's current operator has been slashed 100% for `strategy`, OR the staker has been slashed 100% on the beacon chain such that the calculated slashing factor is 0, this method WILL REVERT. See [Shares Accounting - Fully Slashed](./accounting/SharesAccountingEdgeCases.md#fully-slashed-for-a-strategy) for details. This doesn't block delegation to an operator if the staker has 0 deposit shares for a strategy which has a slashing factor of 0, but any subsequent deposits that call `increaseDelegatedShares` will revert from the **Fully Slashed** edge case.

*Effects*:
* If the staker is delegated to an operator, `addedShares` are added to the operator's shares
* The staker's deposit scaling factor is updated

*Requirements*:
* Caller MUST be either the `StrategyManager` or `EigenPodManager`

#### `decreaseDelegatedShares`

```solidity
/**
 * @notice If the staker is delegated, decreases its operator's shares in response to
 * a decrease in balance in the beaconChainETHStrategy
 * @param staker the staker whose operator's balance will be decreased
 * @param curDepositShares the current deposit shares held by the staker
 * @param beaconChainSlashingFactorDecrease the amount that the staker's beaconChainSlashingFactor has decreased by
 * @dev Note: `beaconChainSlashingFactorDecrease` are assumed to ALWAYS be < 1 WAD.
 * These invariants are maintained in the EigenPodManager.
 */
function decreaseDelegatedShares(
    address staker,
    uint256 curDepositShares,
    uint64 beaconChainSlashingFactorDecrease
) 
    external
    onlyEigenPodManager
    nonReentrant
```

Called by the `EigenPodManager` when a staker's shares decrease due to a checkpointed balance decrease on the beacon chain. If the staker is delegated to an operator, the operator's shares are decreased in return. Otherwise, this method does nothing.

*Effects*: If the staker in question is delegated to an operator, the operator's shares for the `beaconChainETHStrategy` are decreased by the amount the staker's withdrawable shares have decreased by
* This method is a no-op if the staker is not delegated to an operator.

*Requirements*:
* Caller MUST be the `EigenPodManager`