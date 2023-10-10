## EigenPodManager

<!-- Technical details on the EigenPod subsystem as it functions during M2. Includes:
* EigenPodManager
* EigenPods
* Native restaking
* Beacon chain proofs
* Stake / proof / withdrawal flows -->

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`EigenPodManager.sol`](../../src/contracts/pods/EigenPodManager.sol) | Singleton | Transparent proxy |
| [`EigenPod.sol`](../../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy |
| [`DelayedWithdrawalRouter.sol`](../../src/contracts/pods/DelayedWithdrawalRouter.sol) | Singleton | Transparent proxy |
| [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/telepathy-contracts/blob/main/external/integrations/eigenlayer/EigenLayerBeaconOracle.sol) | Singleton | UUPS Proxy |

The `EigenPodManager` and `EigenPod` contracts allow Stakers to restake beacon chain ETH which can then be delegated to Operators via the `DelegationManager`.

The `EigenPodManager` is the entry point for this process, allowing Stakers to deploy an `EigenPod` and begin restaking. While actively restaking, a Staker uses their `EigenPod` to validate various beacon chain state proofs of validator balance and withdrawal status. When exiting, a Staker uses the `DelegationManager` to undelegate or queue a withdrawal from EigenLayer.

`EigenPods` serve as the withdrawal credentials for one or more beacon chain validators controlled by a Staker. Their primary role is to validate beacon chain proofs for each of the Staker's validators. Beacon chain proofs are used to verify a validator's:
* `EigenPod.verifyWithdrawalCredentials`: withdrawal credentials and effective balance
* `EigenPod.verifyBalanceUpdate`: current balance
* `EigenPod.verifyAndProcessWithdrawals`: withdrawable epoch, and processed withdrawals within historical block summary

See [`./proofs`](./proofs/) for detailed documentation on each of the state proofs used in these methods. Additionally, proofs are checked against a beacon chain block root supplied by Succinct's Telepathy protocol ([docs link](https://docs.telepathy.xyz/)).

#### High-level Concepts

The functions of the `EigenPodManager` and `EigenPod` contracts are tightly linked. Rather than writing two separate docs pages, documentation for both contracts is included in this file. This document organizes methods according to the following themes (click each to be taken to the relevant section):
* [Depositing Into EigenLayer](#depositing-into-eigenlayer)
* [Restaking Beacon Chain ETH](#restaking-beacon-chain-eth)
* [Withdrawal Processing](#withdrawal-processing)
* [System Configuration](#system-configuration)
* [Other Methods](#other-methods)

#### Important State Variables

* `EigenPodManager`:
    * `mapping(address => IEigenPod) public ownerToPod`: Tracks the deployed `EigenPod` for each Staker
    * `mapping(address => int256) public podOwnerShares`: Keeps track of the actively restaked beacon chain ETH for each Staker. 
        * In some cases, a beacon chain balance update may cause a Staker's balance to drop below zero. This is because when queueing for a withdrawal in the `DelegationManager`, the Staker's current shares are fully removed. If the Staker's beacon chain balance drops after this occurs, their `podOwnerShares` may go negative. This is a temporary change to account for the drop in balance, and is ultimately corrected when the withdrawal is finally processed.
* `EigenPod`:
    * `_validatorPubkeyHashToInfo[bytes32] -> (ValidatorInfo)`: individual validators are identified within an `EigenPod` according the their public key hash. This mapping keeps track of the following for each validator:
        * `validatorStatus`: (`INACTIVE`, `ACTIVE`, `WITHDRAWN`)
        * `validatorIndex`: A `uint40` that is unique for each validator making a successful deposit via the deposit contract
        * `mostRecentBalanceUpdateTimestamp`: A timestamp that represents the most recent successful proof of the validator's effective balance
        * `restakedBalanceGwei`: Calculated against the validator's proven effective balance using `_calculateRestakedBalanceGwei` (see definitions below)

#### Important Definitions

* "Pod Owner": A Staker who has deployed an `EigenPod` is a Pod Owner. The terms are used interchangably in this document.
    * Pod Owners can only deploy a single `EigenPod`, but can restake any number of beacon chain validators from the same `EigenPod`.
    * Pod Owners can delegate their `EigenPodManager` shares to Operators (via `DelegationManager`).
    * These shares correspond to the amount of provably-restaked beacon chain ETH held by the Pod Owner via their `EigenPod`.
* `EigenPod`:
    * `_calculateRestakedBalanceGwei(uint64 effectiveBalance) -> (uint64)`:
        * This method is used by an `EigenPod` to calculate a "pessimistic" view of a validator's effective balance to avoid the need for repeated balance updates when small balance fluctuations occur.
        * The calculation subtracts an offset (`RESTAKED_BALANCE_OFFSET_GWEI`) from the validator's proven balance, and round down to the nearest ETH
        * Related: `uint64 RESTAKED_BALANCE_OFFSET_GWEI`
            * As of M2, this is 0.75 ETH (in Gwei)
        * Related: `uint64 MAX_VALIDATOR_BALANCE_GWEI`
            * As of M2, this is 31 ETH (in Gwei)
            * This is the maximum amount of restaked ETH a single validator can be credited with in EigenLayer
    * `_podWithdrawalCredentials() -> (bytes memory)`:
        * Gives `abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(EigenPod))`
        * These are the `0x01` withdrawal credentials of the `EigenPod`, used as a validator's withdrawal credentials on the beacon chain.
    

### Depositing Into EigenLayer

Before a Staker begins restaking beacon chain ETH, they need to deploy an `EigenPod`, stake, and start a beacon chain validator:
* [`EigenPodManager.createPod`](#eigenpodmanagercreatepod)
* [`EigenPodManager.stake`](#eigenpodmanagerstake)
    * [`EigenPod.stake`](#eigenpodstake)

To complete the deposit process, the Staker needs to prove that the validator's withdrawal credentials are pointed at the `EigenPod`:
* [`EigenPod.verifyWithdrawalCredentials`](#eigenpodverifywithdrawalcredentials)

#### `EigenPodManager.createPod`

```solidity
function createPod() external
```

Allows a Staker to deploy an `EigenPod` instance, if they have not done so already.

Each Staker can only deploy a single `EigenPod` instance, but a single `EigenPod` can serve as the withdrawal credentials for any number of beacon chain validators. Each `EigenPod` is created using Create2 and the beacon proxy pattern, using the Staker's address as the Create2 salt.

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

*As of M2*:
* `EigenPods` are initialized with restaking enabled by default (`hasRestaked = true`). Pods deployed before M2 may not have this enabled, and will need to call `EigenPod.activateRestaking()`.

#### `EigenPodManager.stake`

```solidity
function stake(
    bytes calldata pubkey, 
    bytes calldata signature, 
    bytes32 depositDataRoot
) 
    external 
    payable
```

Allows a Staker to deposit 32 ETH into the beacon chain deposit contract, providing the credentials for the Staker's beacon chain validator. The `EigenPod.stake` method is called, which automatically calculates the correct withdrawal credentials for the pod and passes these to the deposit contract along with the 32 ETH.

*Effects*:
* Deploys and initializes an `EigenPod` on behalf of Staker, if this has not been done already
* See [`EigenPod.stake`](#eigenpodstake)

*Requirements*:
* If deploying an `EigenPod`, pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`
* See [`EigenPod.stake`](#eigenpodstake)

##### `EigenPod.stake`

```solidity
function stake(
    bytes calldata pubkey,
    bytes calldata signature,
    bytes32 depositDataRoot
) 
    external 
    payable 
    onlyEigenPodManager
```

Handles the call to the beacon chain deposit contract. Only called via `EigenPodManager.stake`.

*Effects*:
* Deposits 32 ETH into the beacon chain deposit contract, and provides the pod's address as the deposit's withdrawal credentials

*Requirements*:
* Caller MUST be the `EigenPodManager`
* Call value MUST be 32 ETH
* Deposit contract `deposit` method MUST succeed given the provided `pubkey`, `signature`, and `depositDataRoot`

#### `EigenPod.verifyWithdrawalCredentials`

```solidity
function verifyWithdrawalCredentials(
    uint64 oracleTimestamp,
    BeaconChainProofs.StateRootProof calldata stateRootProof,
    uint40[] calldata validatorIndices,
    bytes[] calldata validatorFieldsProofs,
    bytes32[][] calldata validatorFields
)
    external
    onlyEigenPodOwner
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
    proofIsForValidTimestamp(oracleTimestamp)
    hasEnabledRestaking
```

Once a Pod Owner has deposited ETH into the beacon chain deposit contract, they can call this method to "fully restake" one or more validators by proving the validator's withdrawal credentials are pointed at the `EigenPod`. This activation will mean that the ETH in those validators:
* is awarded to the Staker/Pod Owner in `EigenPodManager.podOwnerShares`
* is delegatable to an Operator (via the `DelegationManager`)

For each successfully proven validator, that validator's status becomes `VALIDATOR_STATUS.ACTIVE`, and the sum of restakable ether across all newly-proven validators is provided to [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate), where it is added to the Pod Owner's shares. If the Pod Owner is delegated to an Operator via the `DelegationManager`, this sum is also added to the Operator's delegated shares for the beacon chain ETH strategy.

For each validator the Pod Owner wants to verify, the Pod Owner must supply:
* `validatorIndices`: their validator's `ValidatorIndex` (see [consensus specs](https://eth2book.info/capella/part3/config/types/#validatorindex))
* `validatorFields`: the fields of the `Validator` container associated with the validator (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator))
* `stateRootProof`: a proof that will verify `stateRootProof.beaconStateRoot` against an oracle-provided beacon block root
* `validatorFieldsProofs`: a proof that the `Validator` container belongs to the associated validator at the given `ValidatorIndex` within `stateRootProof.beaconStateRoot`
* `oracleTimestamp`: a timestamp used to fetch a beacon block root from `EigenPodManager.beaconChainOracle`

*Effects*:
* For each validator (`_validatorPubkeyHashToInfo[pubkeyHash]`) the validator's info is set for the first time:
    * `VALIDATOR_STATUS` moves from `INACTIVE` to `ACTIVE`
    * `validatorIndex` is recorded
    * `mostRecentBalanceUpdateTimestamp` is set to the `oracleTimestamp` used to fetch the beacon block root
    * `restakedBalanceGwei` is set to `_calculateRestakedBalanceGwei(effectiveBalance)`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

*Requirements*:
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CREDENTIALS`
* Pod MUST have enabled restaking
* All input array lengths MUST be equal
* `oracleTimestamp`:
    * MUST be greater than the `mostRecentWithdrawalTimestamp`
    * MUST be no more than `VERIFY_BALANCE_UPDATE_WINDOW_SECONDS` (~4.5 hrs) old
    * MUST be queryable via `EigenPodManager.getBlockRootAtTimestamp` (fails if `stateRoot == 0`)
* `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot` MUST verify the provided `beaconStateRoot` against the oracle-provided `latestBlockRoot`
* For each validator:
    * The validator's status MUST initially be `VALIDATOR_STATUS.INACTIVE`
    * `BeaconChainProofs.verifyValidatorFields` MUST verify the provided `validatorFields` against the `beaconStateRoot`
    * The aforementioned proofs MUST show that the validator's withdrawal credentials are set to the `EigenPod`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

*As of M2*:
* Restaking is enabled by default for pods deployed after M2. See `activateRestaking` for more info.

---

### Restaking Beacon Chain ETH

At this point, a Staker/Pod Owner has deployed their `EigenPod`, started their beacon chain validator, and proven that its withdrawal credentials are pointed to their `EigenPod`. They are now free to delegate to an Operator (if they have not already), or start up + verify additional beacon chain validators that also withdraw to the same `EigenPod`.

The following methods are the primary operations that concern actively-restaked validators:
* [`EigenPod.verifyBalanceUpdate`](#eigenpodverifybalanceupdate)
* [`EigenPod.verifyAndProcessWithdrawals`](#eigenpodverifyandprocesswithdrawals)
* In conjunction with [`DelegationManager.undelegate`](./DelegationManager.md#undelegate):
    * [`EigenPodManager.forceIntoUndelegationLimbo`](#eigenpodmanagerforceintoundelegationlimbo)
    * [`EigenPodManager.exitUndelegationLimbo`](#eigenpodmanagerexitundelegationlimbo)

#### `EigenPod.verifyBalanceUpdate`

```solidity
function verifyBalanceUpdate(
    uint64 oracleTimestamp,
    uint40 validatorIndex,
    BeaconChainProofs.StateRootProof calldata stateRootProof,
    BeaconChainProofs.BalanceUpdateProof calldata balanceUpdateProof,
    bytes32[] calldata validatorFields
) 
    external 
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE)
```

Anyone (not just the Pod Owner) may call this method with a valid balance update proof to record an balance update in one of the `EigenPod's` validators.

A successful balance update proof updates the `EigenPod's` view of a validator's balance. If the validator's balance has changed, the difference is sent to `EigenPodManager.recordBeaconChainETHBalanceUpdate`, which updates the Pod Owner's shares. If the Pod Owner is delegated to an Operator, this delta is also sent to the `DelegationManager` to update the Operator's delegated beacon chain ETH shares.

For the validator whose balance should be updated, the caller must supply:
* `validatorIndex`: the validator's `ValidatorIndex` (see [consensus specs](https://eth2book.info/capella/part3/config/types/#validatorindex))
* `stateRootProof`: a proof that will verify `stateRootProof.beaconStateRoot` against an oracle-provided beacon block root
* `balanceUpdateProof`: contains the `validatorFieldsProof` mentioned in `verifyWithdrawalCredentials`, as well as a proof that `balanceUpdateProof.balanceRoot` contains the validator's current balance
* `validatorFields`: the fields of the `Validator` container associated with the validator (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator))
* `oracleTimestamp`: a timestamp used to fetch a beacon block root from `EigenPodManager.beaconChainOracle`

*Effects*:
* Updates the validator's stored info:
    * `restakedBalanceGwei` is updated to the newly-proven balance
    * `mostRecentBalanceUpdateTimestamp` is set to the `oracleTimestamp` used to fetch the beacon block root
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE`
* `oracleTimestamp`:
    * MUST be no more than `VERIFY_BALANCE_UPDATE_WINDOW_SECONDS` (~4.5 hrs) old
    * MUST be newer than the validator's `mostRecentBalanceUpdateTimestamp`
    * MUST be queryable via `EigenPodManager.getBlockRootAtTimestamp` (fails if `stateRoot == 0`)
* `validatorFields[0]` MUST be a pubkey hash corresponding to a validator whose withdrawal credentials have been proven, and is not yet withdrawn (`VALIDATOR_STATUS.ACTIVE`)
* `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot` MUST verify the provided `beaconStateRoot` against the oracle-provided `latestBlockRoot`
* `BeaconChainProofs.verifyValidatorFields` MUST verify the provided `validatorFields` against the `beaconStateRoot`
* `BeaconChainProofs.verifyValidatorBalance` MUST verify the provided `balanceRoot` against the `beaconStateRoot`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

#### `EigenPod.verifyAndProcessWithdrawals`

```solidity
function verifyAndProcessWithdrawals(
    uint64 oracleTimestamp,
    BeaconChainProofs.StateRootProof calldata stateRootProof,
    BeaconChainProofs.WithdrawalProof[] calldata withdrawalProofs,
    bytes[] calldata validatorFieldsProofs,
    bytes32[][] calldata validatorFields,
    bytes32[][] calldata withdrawalFields
) 
    external 
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_WITHDRAWAL) 
    onlyNotFrozen
```

Anyone (not just the Pod Owner) can call this method to prove that one or more validators associated with an `EigenPod` have performed a full or partial withdrawal from the beacon chain. 

Whether each withdrawal is a full or partial withdrawal is determined by the validator's "withdrawable epoch" in the `Validator` contained given by `validatorFields` (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator)). If the withdrawal proof timestamp is after this epoch, the withdrawal is a full withdrawal.
* Partial withdrawals are performed automatically by the beacon chain when a validator has an effective balance over 32 ETH. This method can be used to prove that these withdrawals occured, allowing the Pod Owner to withdraw the excess ETH (via [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)).
* Full withdrawals are performed when a Pod Owner decides to fully exit a validator from the beacon chain. To do this, the Pod Owner needs to: (i) exit their validator from the beacon chain, (ii) provide a withdrawal proof to this method, and (iii) enter the withdrawal queue via [`EigenPodManager.queueWithdrawal`](#eigenpodmanagerqueuewithdrawal).

*Effects*:
* For each proven withdrawal:
    * The validator in question is recorded as having a proven withdrawal at the timestamp given by `withdrawalProof.timestampRoot`
        * This is to prevent the same withdrawal from being proven twice
    * If this is a full withdrawal:
        * Any withdrawal amount in excess of `_calculateRestakedBalanceGwei(MAX_VALIDATOR_BALANCE_GWEI)` is immediately withdrawn (see [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal))
            * The remainder must be withdrawn through `EigenPodManager.queueWithdrawal`, but in the meantime is added to `EigenPod.withdrawableRestakedExecutionLayerGwei`
        * If the amount being withdrawn is not equal to the current accounted-for validator balance, a `shareDelta` is calculated to be sent to ([`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)).
        * The validator's info is updated to reflect its `WITHDRAWN` status:
            * `restakedBalanceGwei` is set to 0
            * `mostRecentBalanceUpdateTimestamp` is updated to the timestamp given by `withdrawalProof.timestampRoot`
    * If this is a partial withdrawal:
        * The withdrawal amount is immediately withdrawn (see [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal))

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_WITHDRAWAL`
* All input array lengths MUST be equal
* `oracleTimestamp` MUST be queryable via `EigenPodManager.getBlockRootAtTimestamp` (fails if `stateRoot == 0`)
* `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot` MUST verify the provided `beaconStateRoot` against the oracle-provided `latestBlockRoot`
* For each withdrawal being proven:
    * The time of the withdrawal (`withdrawalProof.timestampRoot`) must be AFTER the `EigenPod's` `mostRecentWithdrawalTimestamp`
    * The validator MUST be in either status: `ACTIVE` or `WITHDRAWN`
        * `WITHDRAWN` is permitted because technically, it's possible to deposit additional ETH into an exited validator and have the ETH be auto-withdrawn.
        * If the withdrawal is a full withdrawal, only `ACTIVE` is permitted
    * The validator MUST NOT have already proven a withdrawal at the `withdrawalProof.timestampRoot`
    * `BeaconChainProofs.verifyWithdrawal` MUST verify the provided `withdrawalFields` against the provided `beaconStateRoot`
    * `BeaconChainProofs.verifyValidatorFields` MUST verify the provided `validatorFields` against the `beaconStateRoot`

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

##### `EigenPodManager.recordBeaconChainETHBalanceUpdate`

```solidity
function recordBeaconChainETHBalanceUpdate(
    address podOwner,
    int256 sharesDelta
) 
    external 
    onlyEigenPod(podOwner) 
    nonReentrant
```

This method is called by an `EigenPod` during a balance update or withdrawal. It accepts a positive or negative `sharesDelta`, which is added/subtracted against the Pod Owner's shares.

If the Pod Owner is not in undelegation limbo and is delegated to an Operator, the `sharesDelta` is also sent to the `DelegationManager` to either increase or decrease the Operator's delegated shares.

*Effects*:
* Adds or removes `sharesDelta` from the Pod Owner's shares
* If the Pod Owner is NOT in undelegation limbo:
    * If `sharesDelta` is negative: see [`DelegationManager.decreaseDelegatedShares`](./DelegationManager.md#decreasedelegatedshares)
    * If `sharesDelta` is positive: see [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Caller MUST be the `EigenPod` associated with the passed-in `podOwner`
* `sharesDelta`: 
    * MUST NOT be 0
    * If negative, `sharesDelta` MUST NOT remove more shares than the Pod Owner has

#### `EigenPodManager.forceIntoUndelegationLimbo`

```solidity
function forceIntoUndelegationLimbo(
    address podOwner,
    address delegatedTo
)
    external
    onlyDelegationManager
    onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
    onlyNotFrozen(podOwner)
    nonReentrant
    returns (uint256 sharesRemovedFromDelegation)
```

This method is called by the `DelegationManager` when a Staker is undelegated from an Operator. If the Staker has `EigenPodManager` shares and isn't in undelegation limbo, the Staker is placed into undelegation limbo and their previously-active shares are returned to the `DelegationManager` to be removed from the Operator.

See *Helpful Definitions* at the top for more info on undelegation limbo.

*Entry Points*:
* `DelegationManager.undelegate`

*Effects*:
* If the Staker has shares and isn't in undelegation limbo, this places them into undelegation limbo and records:
    * `startBlock`: the current block, used to impose a delay on exiting undelegation limbo
    * `delegatedAddress`: the Operator the Staker is being undelegated from

*Requirements*:
* Caller MUST be the `DelegationManager`
* Pause status MUST NOT be set: `PAUSED_WITHDRAW_RESTAKED_ETH`

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

#### `EigenPodManager.exitUndelegationLimbo`

```solidity
function exitUndelegationLimbo(
    uint256 middlewareTimesIndex,
    bool withdrawFundsFromEigenLayer
) 
    external 
    onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH) 
    onlyNotFrozen(msg.sender) 
    nonReentrant
```

Called by a Staker who is currently in undelegation limbo to exit undelegation limbo and either (i) withdraw beacon chain ETH from EigenLayer, or (ii) stay in EigenLayer and continue restaking.

See *Helpful Definitions* at the top for more info on undelegation limbo.

*Effects*:
* Staker's undelegation limbo status is removed
* If withdrawing funds from EigenLayer (funds must already be in `EigenPod`):
    * Staker's `EigenPod.withdrawableRestakedExecutionLayerGwei` is decreased by their current shares
    * `EigenPod.withdrawRestakedBeaconChainETH` directly sends `amountWei` to the Staker
    * Staker's shares are set to 0
* If not withdrawing funds, see [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_WITHDRAW_RESTAKED_ETH`
* Caller MUST be a Staker currently in undelegation limbo, and the current block MUST be at least `StrategyManager.withdrawalDelayBlocks` after the limbo's `startBlock`
* If withdrawing funds from EigenLayer, the funds MUST already be in the Caller's `EigenPod` (and any necessary withdrawals proven via `EigenPod.verifyAndProcessWithdrawals`)

*As of M2*:
* `slasher.canWithdraw` is a no-op
* The `onlyNotFrozen` modifier is a no-op
* The `middlewareTimesIndex` parameter is unused

---

### Withdrawal Processing

* [`EigenPodManager.removeShares`](#eigenpodmanagerremoveshares)
* [`EigenPodManager.addShares`](#eigenpodmanageraddshares)
* [`EigenPodManager.withdrawSharesAsTokens`](#eigenpodmanagerwithdrawsharesastokens)
* [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)

#### `EigenPodManager.removeShares`

TODO

#### `EigenPodManager.addShares`

TODO

#### `EigenPodManager.withdrawSharesAsTokens`

TODO

#### `DelayedWithdrawalRouter.createDelayedWithdrawal`

TODO

#### `EigenPodManager.queueWithdrawal`

```solidity
function queueWithdrawal(
    uint256 amountWei,
    address withdrawer
)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (bytes32)
```

Called by a Pod Owner to initiate a withdrawal of some amount of their beacon chain ETH. The Pod Owner should first call `verifyAndProcessWithdrawals` to prove that a withdrawal can be made in the first place.

Withdrawals can be completed after a delay via `EigenPodManager.completeQueuedWithdrawal`.

*Effects*:
* The Pod Owner's share balance is decreased by `amountWei`
    * If the Pod Owner is not in undelegation limbo, and is delegated to an Operator: (see [`DelegationManager.decreaseDelegatedShares`](./DelegationManager.md#decreasedelegatedshares))
* The Pod Owner's `EigenPod.withdrawableRestakedExecutionLayerGwei` is decreased by `amountWei`
* A `BeaconChainQueuedWithdrawal` is created, recording:
    * The shares being withdrawn (`amountWei`)
    * The Pod Owner (`msg.sender`)
    * The Pod Owner's current withdrawal nonce
    * The "start block" of the withdrawal (`block.number`)
    * The address to which the Pod Owner is currently delegated (can be zero)
    * The specified `withdrawer` address

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_WITHDRAW_RESTAKED_ETH`
* Caller MUST be a Pod Owner who is NOT in undelegation limbo
* The withdrawal amount (`amountWei`):
    * MUST NOT be zero
    * MUST NOT be greater than the Pod Owner's shares
    * MUST be equally divisible by 1e9 (only whole amounts of Gwei may be withdrawn)
    * MUST already exist in the Pod Owner's `EigenPod.withdrawableRestakedExecutionLayerGwei`, meaning the amount being queued has been proven to have been withdrawn by a validator within the `EigenPod`. (see [`EigenPod.verifyAndProcessWithdrawals`](#eigenpodverifyandprocesswithdrawals))

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

### System Configuration

* [`EigenPodManager.setMaxPods`](#eigenpodmanagersetmaxpods)
* [`EigenPodManager.updateBeaconChainOracle`](#eigenpodmanagerupdatebeaconchainoracle)

#### `EigenPodManager.setMaxPods`

```solidity
function setMaxPods(uint256 newMaxPods) external onlyUnpauser
```

*Effects*:
* TODO

*Requirements*:
* TODO

#### `EigenPodManager.updateBeaconChainOracle`

```solidity
function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyOwner
```

*Effects*:
* TODO

*Requirements*:
* TODO

### Other Methods

This section details various methods that don't fit well into other sections.

*For pods deployed prior to M2*, the following methods are callable:
* [`EigenPod.activateRestaking`](#eigenpodactivaterestaking)
* [`EigenPod.withdrawBeforeRestaking`](#eigenpodwithdrawbeforerestaking)

`EigenPod` also includes two token recovery mechanisms:
* [`EigenPod.withdrawNonBeaconChainETHBalanceWei`](#eigenpodwithdrawnonbeaconchainethbalancewei)
* [`EigenPod.recoverTokens`](#eigenpodrecovertokens)

#### `EigenPod.activateRestaking`

```solidity
function activateRestaking()
    external
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
    onlyEigenPodOwner
    hasNeverRestaked
```

Note: This method is only callable on pods deployed before M2. After M2, restaking is enabled by default. 

`activateRestaking` allows a Pod Owner to designate their pod (and any future ETH sent to it) as being restaked. Calling this method first withdraws any ETH in the `EigenPod` via the `DelayedWithdrawalRouter`, and then prevents further calls to `withdrawBeforeRestaking`.

Withdrawing any future ETH sent via beacon chain withdrawal to the `EigenPod` requires providing beacon chain state proofs. However, ETH sent to the pod's `receive` function should be withdrawable without proofs (see [`withdrawNonBeaconChainETHBalanceWei`](#eigenpodwithdrawnonbeaconchainethbalancewei)).

*Effects*:
* Sets `hasRestaked = true`
* Updates the pod's most recent withdrawal timestamp to the current time
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*Requirements*:
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`
* Pod MUST NOT have already activated restaking
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*As of M2*: restaking is automatically activated for newly-deployed `EigenPods` (`hasRestaked = true`). However, for `EigenPods` deployed *before* M2, restaking may not be active (unless the Pod Owner has called this method).

#### `EigenPod.withdrawBeforeRestaking`

```solidity
function withdrawBeforeRestaking() 
    external 
    onlyEigenPodOwner 
    hasNeverRestaked
```

Note: This method is only callable on pods deployed before M2. After M2, restaking is enabled by default. 

Allows the Pod Owner to withdraw any ETH in the `EigenPod` via the `DelayedWithdrawalRouter`, assuming restaking has not yet been activated. See [`EigenPod.activateRestaking`](#eigenpodactivaterestaking) for more details.

*Effects*:
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*Requirements*:
* Caller MUST be the Pod Owner
* Pod MUST NOT have already activated restaking
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*As of M2*: restaking is automatically activated for newly-deployed `EigenPods`, making this method uncallable for pods deployed after M2. However, for `EigenPods` deployed *before* M2, restaking may not be active, and this method may be callable.

#### `EigenPod.withdrawNonBeaconChainETHBalanceWei`

TODO

#### `EigenPod.recoverTokens`

TODO