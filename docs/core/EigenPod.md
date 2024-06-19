[eip-4788]: https://eips.ethereum.org/EIPS/eip-4788
[custom-types]: https://eth2book.info/capella/part3/config/types/#custom-types
[validator-container]: https://eth2book.info/capella/part3/containers/dependencies/#validator

## EigenPod

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`EigenPod.sol`](../../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy |

An `EigenPod` is deployed via the `EigenPodManager` by a Staker (referred to in this doc as the _Pod Owner_). `EigenPods` allow a Pod Owner to restake one or more beacon chain validators, earning shares which can be delegated to Operators to earn yield. When a Pod Owner begins running a validator on the beacon chain, they choose _withdrawal credentials_ for that validator. Withdrawal credentials are the ETH address to which:
* A validator's _principal_ is sent when the validator exits the beacon chain
* A validator's _consensus rewards_ are sent as the validator proposes/attests to blocks on the beacon chain

Additionally, when running validator node software, a validator is configured with a _fee recipient_. The fee recipient receives:
* _Execution layer rewards_ when the validator proposes a block
* _MEV rewards_ if the validator is running MEV-boost/other custom block proposer software

**An `EigenPod` may serve as EITHER/BOTH the withdrawal credentials OR the fee recipient for your validators.** In prior releases, it was only possible to use an `EigenPod` for withdrawal credentials. However, this is no longer the case!

---

The **primary goal** of the `EigenPod` system is to **ensure that shares are backed 1:1** with ETH that is _either already in the `EigenPod`, or will eventually flow through the `EigenPod`._ To support this goal, `EigenPods`: 
* serve as the withdrawal credentials for one or more beacon chain validators controlled by the Pod Owner
* validate beacon chain state proofs
* interpret these proofs to add or remove shares in the beacon chain ETH strategy

Because beacon chain proofs are processed asynchronously from the beacon chain itself, there is an inherent _lag_ between an event on the beacon chain and a corresponding share update in any affected `EigenPods`. Therefore, the **secondary goals** of the `EigenPod` system are to **minimize lag where possible** and to **ensure various timing windows cannot (i) create unbacked shares or (ii) prevent the withdrawal of existing shares.**

#### High-level Concepts

* [Restaking Beacon Chain ETH](#restaking-beacon-chain-eth)
* [Checkpointing Validators](#checkpointing-validators)
* [Staleness Proofs](#staleness-proofs)
* [Other Methods](#other-methods)
* [Compatibility with Previous Versions](#compatibility-with-previous-versions)

#### Important Definitions

**_Pod Owner_**: A Staker who has deployed an `EigenPod` is a _Pod Owner_. The terms are used interchangeably in this document.
* _Pod Owners_ can only deploy a single `EigenPod`, but can restake any number of beacon chain validators from the same `EigenPod`.
* _Pod Owners_ can delegate their `EigenPodManager` shares to Operators (via `DelegationManager`).
* These shares correspond to the amount of restaked beacon chain ETH held by the _Pod Owner_ via their `EigenPod`.

**_Active validator set_**: This term is used frequently in this document to describe the set of validators whose withdrawal credentials have been verified to be pointed at an `EigenPod`. The _active validator set_ is used to determine the number of proofs required to complete a checkpoint (see [Checkpointing Validators](#checkpointing-validators)).
* A validator enters the _active validator set_ when their withdrawal credentials are verified (see [`verifyWithdrawalCredentials`](#verifywithdrawalcredentials))
* A validator leaves the _active validator set_ when a checkpoint proof shows they have 0 balance (see [`verifyCheckpointProofs`](#verifycheckpointproofs))

In the implemtation, the _active validator set_ is comprised of two state variables:
* `uint256 activeValidatorCount` 
    * incremented by 1 when a validator enters the _active validator set_
    * decremented by 1 when a validator leaves the _active validator set_
* `mapping(bytes32 => ValidatorInfo) _validatorPubkeyHashToInfo` (specifically, the `status` field)
    * `VALIDATOR_STATUS.INACTIVE -> VALIDATOR_STATUS.ACTIVE` when entering the _active validator set_
    * `VALIDATOR_STATUS.ACTIVE -> VALIDATOR_STATUS.WITHDRAWN` when leaving the _active validator set_

**_Checkpoint_**: A snapshot of `EigenPod` and beacon chain state used to update the _Pod Owner's_ shares based on a combination of beacon chain balance and native ETH balance. Checkpoints allow an `EigenPod` to account for validator exits, partial withdrawals of consensus rewards, or execution layer fees earned by their validators. Completing a checkpoint will account for these amounts in the `EigenPod`, enabling the _Pod Owner_ to compound their restaked shares or withdraw accumulated yield.

Only one _checkpoint_ can be active at a time in a given `EigenPod`. The pod's _current checkpoint_ is represented by the following data structure:

```solidity
struct Checkpoint {
    bytes32 beaconBlockRoot;  // proofs are verified against a beacon block root
    uint24 proofsRemaining;   // number of proofs remaining before the checkpoint is completed
    uint64 podBalanceGwei;    // native ETH that will be awarded shares when the checkpoint is completed
    int128 balanceDeltasGwei; // total change in beacon chain balance tracked across submitted proofs
}
```

Checkpoints are completed by submitting one beacon chain proof per validator in the pod's _active validator set_. See [Checkpointing Validators](#checkpointing-validators) for details.

---    

### Restaking Beacon Chain ETH

If a Pod Owner has validators whose withdrawal credentials are an `EigenPod`, the Pod Owner can use `verifyWithdrawalCredentials` to begin restaking ETH while it is still on the beacon chain. Once a validator's withdrawal credentials are verified:
* the Pod Owner receives delegatable shares via `EigenPodManager.podOwnerShares`
* the validator enters the pod's _active validator set_, and must be included in future checkpoint proofs (see [Checkpointing Validators](#checkpointing-validators))

_Methods:_
* [`verifyWithdrawalCredentials`](#verifywithdrawalcredentials)

#### `verifyWithdrawalCredentials`

```solidity
function verifyWithdrawalCredentials(
    uint64 beaconTimestamp,
    BeaconChainProofs.StateRootProof calldata stateRootProof,
    uint40[] calldata validatorIndices,
    bytes[] calldata validatorFieldsProofs,
    bytes32[][] calldata validatorFields
)
    external
    onlyEigenPodOwner
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)

struct StateRootProof {
    bytes32 beaconStateRoot;
    bytes proof;
}
```

This method first verifies a beacon state root against a beacon block root returned by the [EIP-4788 oracle][eip-4788]. Then, it verifies one or more withdrawal credential proofs against the beacon state root. Finally, the Pod Owner is awarded shares according to the sum of the effective balance of each verified validator (via `EigenPodManager.recordBeaconChainETHBalanceUpdate`).

A withdrawal credential proof uses a validator's [`ValidatorIndex`][custom-types] and a merkle proof to prove the existence of a [`Validator` container][validator-container] at a given block. The beacon chain `Validator` container holds important information used in this method:
* `pubkey`: A BLS pubkey hash, used to uniquely identify the validator within the `EigenPod`
* `withdrawal_credentials`: Used to verify that the validator will withdraw its principal to this `EigenPod` if it exits the beacon chain
* `effective_balance`: The balance of the validator, updated once per epoch and capped at 32 ETH. Used to award shares to the Pod Owner
* `exit_epoch`: Initially set to `type(uint64).max`, this value is updated when a validator initiates exit from the beacon chain. This method requires that a validator _has not initiated an exit from the beacon chain_

_Note that it is not required to verify your validator's withdrawal credentials_, unless you want to receive shares for ETH on the beacon chain. You may choose to use your `EigenPod` without verifying withdrawal credentials; you will still be able to withdraw yield (or receive shares for yield) via the [checkpoint system](#checkpointing-validators).

*Effects*:
* For each set of unique verified withdrawal credentials:
    * `activeValidatorCount` is increased by 1
    * The validator's info is recorded in state (`_validatorPubkeyHashToInfo[pubkeyHash]`):
        * `validatorIndex` is recorded from the passed-in `validatorIndices`
        * `restakedBalanceGwei` is set to the validator's effective balance
        * `lastCheckpointedAt` is set to either the `lastCheckpointTimestamp` or `currentCheckpointTimestamp`
        * `VALIDATOR_STATUS` moves from `INACTIVE` to `ACTIVE`
* The Pod Owner is awarded shares according to the sum of effective balances proven. See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#TODO)

*Requirements*:
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CREDENTIALS`
* `hasRestaked` MUST be set to `true`
    * Note: for the majority of pods, this is `true` by default (see [Compatibility with Previous Versions](#compatibility-with-previous-versions))
* Input array lengths MUST be equal
* `beaconTimestamp`:
    * MUST be greater than BOTH `lastCheckpointTimestamp` AND `currentCheckpointTimestamp`
    * MUST be queryable via the [EIP-4788 oracle][eip-4788]. Generally, this means `beaconTimestamp` corresponds to a valid beacon block created within the last 24 hours.
* `stateRootProof` MUST verify a `beaconStateRoot` against the `beaconBlockRoot` returned from the EIP-4788 oracle
* For each validator:
    * The validator MUST NOT have been previously-verified (`VALIDATOR_STATUS` should be `INACTIVE`)
    * The validator's `exit_epoch` MUST equal `type(uint64).max` (aka `FAR_FUTURE_EPOCH`)
    * The validator's `withdrawal_credentials` MUST be pointed to the `EigenPod`
    * `validatorFieldsProof` MUST be a valid merkle proof of the corresponding `validatorFields` under the `beaconStateRoot` at the given `validatorIndex`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#TODO)

---

### Checkpointing Validators

Checkpoint proofs comprise the bulk of proofs submitted to an `EigenPod`. Completing a checkpoint means submitting _one checkpoint proof for each validator_ in the pod's _active validator set._

`EigenPods` use checkpoints to detect:
* when validators have exited from the beacon chain, leaving the pod's _active validator set_
* when the pod has accumulated fees / partial withdrawals from validators
* whether any validators on the beacon chain have increased/decreased in balance

When a checkpoint is completed, shares are updated accordingly for each of these events. Shares can be withdrawn via the `DelegationManager` withdrawal queue (see [DelegationManager: Undelegating and Withdrawing](./DelegationManager.md#undelegating-and-withdrawing)), which means an `EigenPod's` checkpoint proofs also play an important role in allowing Pod Owners to exit funds from the system.

_Important Notes:_
* `EigenPods` can only have **one active checkpoint** at a given time, and once started, checkpoints **cannot be cancelled** (only completed)
* Checkpoint proofs are based entirely off of _current balance_ proofs. Even though partial/full withdrawals are processed via checkpoint proofs, this system does NOT use withdrawal proofs.

_Methods:_
* [`startCheckpoint`](#startcheckpoint)
* [`verifyCheckpointProofs`](#verifycheckpointproofs)

#### `startCheckpoint`

```solidity
function startCheckpoint(bool revertIfNoBalance)
    external
    onlyEigenPodOwner() 
    onlyWhenNotPaused(PAUSED_START_CHECKPOINT) 
```

This method allows a Pod Owner to start a checkpoint, beginning the process of proving a pod's _active validator set_. `startCheckpoint` takes a snapshot of three things:
* `podBalanceGwei`: the `EigenPod's` native ETH balance, minus any balance already credited with shares through previous checkpoints
    * Note: if `revertIfNoBalance == true`, this method will revert if `podBalanceGwei == 0`. This is to allow a Pod Owner to avoid creating a checkpoint unintentionally.
* `activeValidatorCount`: the number of validators in the pod's _active validator set_, aka the number of validators with verified withdrawal credentials who have NOT been proven exited via a previous checkpoint
    * This becomes the checkpoint's `proofsRemaining`, or the number of proofs that need to be submitted to `verifyCheckpointProofs` to complete the checkpoint
* `beaconBlockRoot`: the beacon block root of the previous slot, fetched by querying the [EIP-4788 oracle][eip-4788] with the current `block.timestamp`
    * This is used as the single source of truth for all proofs submitted for this checkpoint

`startCheckpoint` plays a very important role in the security of the checkpoint process: it guarantees that _the pod's native ETH balance and any beacon balances proven in the checkpoint are 100% distinct_. That is: if a partial/full exit is processed in the block before `startCheckpoint` is called, then:
* The withdrawn ETH is already in the pod when `startCheckpoint` is called, and is factored into `podBalanceGwei`
* A proof of the validator's current balance against `beaconBlockRoot` will NOT include the withdrawn ETH

This guarantee means that, if we use the checkpoint to sum up the beacon chain balance of the pod's _active validator set_, **we can award guaranteed-backed shares** according to the sum of the pod's beacon chain balance and its native ETH balance.

*Effects*:
* Sets `currentCheckpointTimestamp` to `block.timestamp`
* Creates a new `Checkpoint`:
    * `beaconBlockRoot`: set to the current block's parent beacon block root, fetched by querying the [EIP-4788 oracle][eip-4788] using `block.timestamp` as input.
    * `proofsRemaining`: set to the current value of `activeValidatorCount` (note that this value MAY be 0)
    * `podBalanceGwei`: set to the pod's native ETH balance, minus any balance already accounted for in previous checkpoints
    * `balanceDeltasGwei`: set to 0 initially
* If `checkpoint.proofsRemaining == 0`, the new checkpoint is auto-completed:
    * `withdrawableRestakedExecutionLayerGwei` is increased by `checkpoint.podBalanceGwei`
    * `lastCheckpointTimestamp` is set to `currentCheckpointTimestamp`
    * `currentCheckpointTimestamp` and `_currentCheckpoint` are deleted
    * The Pod Owner's shares are updated (see [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#TODO))
* If `hasRestaked == false`, sets `hasRestaked` to `true` (see [Compatibility with Previous Versions](#compatibility-with-previous-versions))

*Requirements*:
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_START_CHECKPOINT`
* A checkpoint MUST NOT be active (`currentCheckpointTimestamp == 0`)
* The last checkpoint completed MUST NOT have been started in the current block (`lastCheckpointTimestamp != block.timestamp`)
* If `revertIfNoBalance == true`, the pod's native ETH balance MUST contain some nonzero value not already accounted for in the _Pod Owner's_ shares

#### `verifyCheckpointProofs`

```solidity
function verifyCheckpointProofs(
    BeaconChainProofs.BalanceContainerProof calldata balanceContainerProof,
    BeaconChainProofs.BalanceProof[] calldata proofs
)
    external 
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) 

struct BalanceContainerProof {
    bytes32 balanceContainerRoot;
    bytes proof;
}

struct BalanceProof {
    bytes32 pubkeyHash;
    bytes32 balanceRoot;
    bytes proof;
}
```

`verifyCheckpointProofs` is used to make progress on (or complete) the pod's current checkpoint. This method accepts one or more merkle proofs of validators' _current balances_ against a `balanceContainerRoot`. Additionally, a `balanceContainerProof` verifies this `balanceContainerRoot` against the _current checkpoint's_ `beaconBlockRoot`.

Proofs submitted to this method concern a validator's _current balance,_ NOT their _effective balance_. The _current balance_ is updated every slot, while _effective balances_ are updated roughly once per epoch. Current balances are stored in the [`BeaconState.balances` field](https://eth2book.info/capella/part3/containers/state/#beaconstate).

For each validator submitted via `proofs`:
* The validator's `status` should be `ACTIVE`. That is, its withdrawal credentials are verified (see [`verifyWithdrawalCredentials`](#verifywithdrawalcredentials)), and it has a nonzero balance as of the last time it was seen in a checkpoint proof.
* The validator's `lastCheckpointedAt` should be less than `currentCheckpointTimestamp`. This is to prevent a validator from counting towards a checkpoint's progression more than once.

If either of these two conditions is not met, _the proof will be skipped but execution will continue_. Execution continues without reverting to prevent a potential griefing vector where anyone could frontrun a batch of proofs, submit one proof from the batch, and cause the batch to revert.

Each valid proof submitted decreases the _current checkpoint's_ `proofsRemaining` by 1. If `proofsRemaining` hits 0 the checkpoint is automatically completed, updating the Pod Owner's shares accordingly.

*Effects*:
* For each validator successfully checkpointed:
    * The number of proofs remaining in the checkpoint is decreased (`checkpoint.proofsRemaining--`)
    * A balance delta is calculated using the validator's previous `restakedBalanceGwei`. This delta is added to `checkpoint.balanceDeltasGwei` to track the total beacon chain balance delta.
    * The validator's `restakedBalanceGwei` and `lastCheckpointedAt` fields are updated. Additionally, if the proof shows that the validator has a balance of 0, the validator's status is moved to `VALIDATOR_STATUS.WITHDRAWN` and the pod's `activeValidatorCount` is decreased.
* If the checkpoint's `proofsRemaining` drops to 0, the checkpoint is automatically completed:
    * `checkpoint.podBalanceGwei` is added to `withdrawableRestakedExecutionLayerGwei`, rendering it accounted for in future checkpoints
    * `lastCheckpointTimestamp` is set to `currentCheckpointTimestamp`, and both `_currentCheckpoint` and `currentCheckpointTimestamp` are deleted.
    * The Pod Owner's total share delta is calculated as the sum of `checkpoint.podBalanceGwei` and `checkpoint.balanceDeltasGwei`, and forwarded to the `EigenPodManager` (see [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](TODO))

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS`
* A checkpoint MUST currently be active (`currentCheckpointTimestamp != 0`)
* `balanceContainerProof` MUST contain a valid merkle proof of the beacon chain's balances container against `_currentCheckpoint.beaconBlockRoot`
* Each `proof` in `proofs` MUST contain a valid merkle proof of the validator's `balanceRoot` against `balanceContainerProof.balanceContainerRoot`

---

### Staleness Proofs

TODO

#### `verifyStaleBalance`

```solidity
function verifyStaleBalance(
    uint64 beaconTimestamp,
    BeaconChainProofs.StateRootProof calldata stateRootProof,
    BeaconChainProofs.ValidatorProof calldata proof
)
    external
    onlyWhenNotPaused(PAUSED_START_CHECKPOINT) 
    onlyWhenNotPaused(PAUSED_VERIFY_STALE_BALANCE)
```

TODO

---

### Other Methods

TODO

#### `stake`

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

TODO

#### `withdrawRestakedBeaconChainETH`

```solidity
function withdrawRestakedBeaconChainETH(
    address recipient, 
    uint256 amountWei
)
    external 
    onlyEigenPodManager
```

TODO

#### `recoverTokens`

```solidity
function recoverTokens(
    IERC20[] memory tokenList,
    uint256[] memory amountsToWithdraw,
    address recipient
) 
    external 
    onlyEigenPodOwner 
    onlyWhenNotPaused(PAUSED_NON_PROOF_WITHDRAWALS)
```

TODO

---

### Compatibility with Previous Versions

Although all `EigenPods` are updated simultaneously via the `BeaconProxy` pattern, previous versions of `EigenPods` have had slightly different state models that need to be accounted for in the latest release. There are two prior major versions of `EigenPod`:
* Pods deployed after M2 mainnet ("M2 Pods") were deployed with `EigenPod.hasRestaked` set to `true`. This is the default behavior for any new `EigenPods`.
* Pods deployed during M1 mainnet ("M1 Pods") were deployed with `EigenPod.hasRestaked` set to `false`. After M2 mainnet, owners of M1 Pods had two options:
    * Choose not to "fully upgrade" to an M2 Pod, and continue withdrawing ETH from the pod at will via [`EigenPod.withdrawBeforeRestaking`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v0.2.5-mainnet-m2-minor-eigenpod-upgrade/src/contracts/pods/EigenPod.sol#L403-L406).
    * Call [`EigenPod.activateRestaking`](https://github.com/Layr-Labs/eigenlayer-contracts/blob/v0.2.5-mainnet-m2-minor-eigenpod-upgrade/src/contracts/pods/EigenPod.sol#L386-L401), setting `EigenPod.hasRestaked` to `true` and performing one final withdrawal of any ETH in the pod.

Generally, M1 Pods that chose to activate restaking are identical to M2 Pods - so for the sake of clarity, "M1 Pods" will refer to pods that were deployed _before M2_ and _did not call `activateRestaking`_ at any point.

This latest release brings M1 and M2 pods together, forcing both pods to use the same state model while supporting the original behavior that M1 Pod owners prefer. To clarify, M1 Pod owners have never called `verifyWithdrawalCredentials` for any validators, and have therefore never earned restaking shares for any validators. However, M1 Pod owners are also able to withdraw yield without supplying withdrawal proofs required by M2 Pods.

There are a few user groups supported by this release. Note that M2 Pod Owners don't need to do anything special and can begin using the new `EigenPod` ABI as-is.

#### M1 Pod Owners that want to maintain M1 Pod behavior

#### M1 Pod Owners that want to start restaking now that proofs are cheaper

