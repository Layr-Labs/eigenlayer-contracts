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

#### Important Definitions

**_Pod Owner_**: A Staker who has deployed an `EigenPod` is a _Pod Owner_. The terms are used interchangeably in this document.
* _Pod Owners_ can only deploy a single `EigenPod`, but can restake any number of beacon chain validators from the same `EigenPod`.
* _Pod Owners_ can delegate their `EigenPodManager` shares to Operators (via `DelegationManager`).
* These shares correspond to the amount of restaked beacon chain ETH held by the _Pod Owner_ via their `EigenPod`.

**_Proof Submitter_**: An address designated by the Pod Owner with permissions to call certain `EigenPod` methods. This role is provided to allow Pod Owners to manage their day-to-day `EigenPod` tasks via hot wallets, rather than the Pod Owner address which controls all funds. The Proof Submitter can call `verifyWithdrawalCredentials` and `startCheckpoint`. See [`setProofSubmitter` docs](#setproofsubmitter) for more details.

**_Active validator set_**: This term is used frequently in this document to describe the set of validators whose withdrawal credentials have been verified to be pointed at an `EigenPod`. The _active validator set_ is used to determine the number of proofs required to complete a checkpoint (see [Checkpointing Validators](#checkpointing-validators)).
* A validator enters the _active validator set_ when their withdrawal credentials are verified (see [`verifyWithdrawalCredentials`](#verifywithdrawalcredentials))
* A validator leaves the _active validator set_ when a checkpoint proof shows they have 0 balance (see [`verifyCheckpointProofs`](#verifycheckpointproofs))

In the implementation, the _active validator set_ is comprised of two state variables:
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
    onlyOwnerOrProofSubmitter
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
* `activation_epoch`: Initially set to `type(uint64).max`, this value is updated when a validator reaches a balance of at least 32 ETH, designating the validator is ready to become active on the beacon chain. **This method requires that a validator is either already active, or in the process of activating on the beacon chain.**
* `exit_epoch`: Initially set to `type(uint64).max`, this value is updated when a validator initiates exit from the beacon chain. **This method requires that a validator has not initiated an exit from the beacon chain.**
  * If a validator has been exited prior to calling `verifyWithdrawalCredentials`, their ETH can be accounted for, awarded shares, and/or withdrawn via the checkpoint system (see [Checkpointing Validators](#checkpointing-validators)).

_Note that it is not required to verify your validator's withdrawal credentials_, unless you want to receive shares for ETH on the beacon chain. You may choose to use your `EigenPod` without verifying withdrawal credentials; you will still be able to withdraw yield (or receive shares for yield) via the [checkpoint system](#checkpointing-validators).

*Effects*:
* For each set of unique verified withdrawal credentials:
    * `activeValidatorCount` is increased by 1
    * The validator's info is recorded in state (`_validatorPubkeyHashToInfo[pubkeyHash]`):
        * `validatorIndex` is recorded from the passed-in `validatorIndices`
        * `restakedBalanceGwei` is set to the validator's effective balance
        * `lastCheckpointedAt` is set to either the `lastCheckpointTimestamp` or `currentCheckpointTimestamp`
        * `VALIDATOR_STATUS` moves from `INACTIVE` to `ACTIVE`
* The Pod Owner is awarded shares according to the sum of effective balances proven. See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](./EigenPodManager.md#recordbeaconchainethbalanceupdate)

*Requirements*:
* Caller MUST be EITHER the Pod Owner or Proof Submitter
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CREDENTIALS`
* Input array lengths MUST be equal
* `beaconTimestamp`:
    * MUST be greater than `currentCheckpointTimestamp`
    * MUST be queryable via the [EIP-4788 oracle][eip-4788]. Generally, this means `beaconTimestamp` corresponds to a valid beacon block created within the last 8192 blocks (~27 hours).
* `stateRootProof` MUST verify a `beaconStateRoot` against the `beaconBlockRoot` returned from the EIP-4788 oracle
* For each validator:
    * The validator MUST NOT have been previously-verified (`VALIDATOR_STATUS` should be `INACTIVE`)
    * The validator's `activation_epoch` MUST NOT equal `type(uint64).max` (aka `FAR_FUTURE_EPOCH`)
    * The validator's `exit_epoch` MUST equal `type(uint64).max` (aka `FAR_FUTURE_EPOCH`)
    * The validator's `withdrawal_credentials` MUST be pointed to the `EigenPod`
    * `validatorFieldsProof` MUST be a valid merkle proof of the corresponding `validatorFields` under the `beaconStateRoot` at the given `validatorIndex`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](./EigenPodManager.md#recordbeaconchainethbalanceupdate)

---

### Checkpointing Validators

Checkpoint proofs comprise the bulk of proofs submitted to an `EigenPod`. Completing a checkpoint means submitting _one checkpoint proof for each validator_ in the pod's _active validator set._

`EigenPods` use checkpoints to detect:
* when validators have exited from the beacon chain, leaving the pod's _active validator set_
* when the pod has accumulated fees / partial withdrawals from validators
* whether any validators on the beacon chain have increased/decreased in balance

When a checkpoint is completed, shares are updated accordingly for each of these events. OwnedShares can be withdrawn via the `DelegationManager` withdrawal queue (see [DelegationManager: Undelegating and Withdrawing](./DelegationManager.md#undelegating-and-withdrawing)), which means an `EigenPod's` checkpoint proofs also play an important role in allowing Pod Owners to exit funds from the system.

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
    onlyOwnerOrProofSubmitter() 
    onlyWhenNotPaused(PAUSED_START_CHECKPOINT) 
```

This method allows a Pod Owner (or Proof Submitter) to start a checkpoint, beginning the process of proving a pod's _active validator set_. `startCheckpoint` takes a snapshot of three things:
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
    * The Pod Owner's shares are updated (see [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](./EigenPodManager.md#recordbeaconchainethbalanceupdate))

*Requirements*:
* Caller MUST be EITHER the Pod Owner or Proof Submitter
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
    * The Pod Owner's total share delta is calculated as the sum of `checkpoint.podBalanceGwei` and `checkpoint.balanceDeltasGwei`, and forwarded to the `EigenPodManager` (see [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](./EigenPodManager.md#recordbeaconchainethbalanceupdate))

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS`
* A checkpoint MUST currently be active (`currentCheckpointTimestamp != 0`)
* `balanceContainerProof` MUST contain a valid merkle proof of the beacon chain's balances container against `_currentCheckpoint.beaconBlockRoot`
* Each `proof` in `proofs` MUST contain a valid merkle proof of the validator's `balanceRoot` against `balanceContainerProof.balanceContainerRoot`

---

### Staleness Proofs

Regular checkpointing of validators plays an important role in the health of the system, as a completed checkpoint ensures that the pod's shares and backing assets are up to date.

Typically, checkpoints can only be started by the Pod Owner (see [`startCheckpoint`](#startcheckpoint)). This is because completing a checkpoint with a lot of validators has the potential to be an expensive operation, so gating `startCheckpoint` to only be callable by the Pod Owner prevents a griefing vector where anyone can cheaply force the Pod Owner to perform a checkpoint.

In most cases, Pod Owners are incentivized to perform their own regular checkpoints, as completing checkpoints is the only way to access yield sent to the pod. However, if beacon chain validators are slashed, it's possible that a Pod Owner no longer has an incentive to start/complete a checkpoint. After all, they would be losing shares equal to the slashed amount. Unless they have enough unclaimed yield in the pod to make up for this, they only stand to lose by completing a checkpoint.

In this case, `verifyStaleBalance` can be used to allow a third party to start a checkpoint on the Pod Owner's behalf.

*Methods*:
* [`verifyStaleBalance`](#verifystalebalance)

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

Allows anyone to prove that a validator in the pod's _active validator set_ was slashed on the beacon chain. A successful proof allows the caller to start a checkpoint. Note that if the pod currently has an active checkpoint, the existing checkpoint needs to be completed before `verifyStaleBalance` can start a checkpoint.

A valid proof has the following requirements:
* The `beaconTimestamp` MUST be newer than the timestamp the validator was last checkpointed at
* The validator in question MUST be in the _active validator set_ (have the status `VALIDATOR_STATUS.ACTIVE`)
* The proof MUST show that the validator has been slashed

If these requirements are met and the proofs are valid against a beacon block root given by `beaconTimestamp`, a checkpoint is started.

*Effects*:
* Sets `currentCheckpointTimestamp` to `block.timestamp`
* Creates a new `Checkpoint`:
    * `beaconBlockRoot`: set to the current block's parent beacon block root, fetched by querying the [EIP-4788 oracle][eip-4788] using `block.timestamp` as input.
    * `proofsRemaining`: set to the current value of `activeValidatorCount`
    * `podBalanceGwei`: set to the pod's native ETH balance, minus any balance already accounted for in previous checkpoints
    * `balanceDeltasGwei`: set to 0 initially

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_START_CHECKPOINT`
* Pause status MUST NOT be set: `PAUSED_VERIFY_STALE_BALANCE`
* A checkpoint MUST NOT be active (`currentCheckpointTimestamp == 0`)
* The last checkpoint completed MUST NOT be the current block
* For the validator given by `proof.validatorFields`:
    * `beaconTimestamp` MUST be greater than `validatorInfo.lastCheckpointedAt`
    * `validatorInfo.status` MUST be `VALIDATOR_STATUS.ACTIVE`
    * `proof.validatorFields` MUST show that the validator is slashed
* `stateRootProof` MUST verify a `beaconStateRoot` against the `beaconBlockRoot` returned from the EIP-4788 oracle
* The `ValidatorProof` MUST contain a valid merkle proof of the corresponding `validatorFields` under the `beaconStateRoot` at `validatorInfo.validatorIndex`

---

### Other Methods

Minor methods that do not fit well into other sections:
* [`setProofSubmitter`](#setproofsubmitter)
* [`stake`](#stake)
* [`withdrawRestakedBeaconChainETH`](#withdrawrestakedbeaconchaineth)
* [`recoverTokens`](#recovertokens)

#### `setProofSubmitter`

```solidity
function setProofSubmitter(address newProofSubmitter) external onlyEigenPodOwner
```

Allows the Pod Owner to update the Proof Submitter address for the `EigenPod`. The Proof Submitter can call `verifyWithdrawalCredentials` and `startCheckpoint` just like the Pod Owner. This is intended to allow the Pod Owner to create a hot wallet to manage calls to these methods.

If set, EITHER the Pod Owner OR Proof Submitter may call `verifyWithdrawalCredentials`/`startCheckpoint`.

The Pod Owner can call this with `newProofSubmitter == 0` to remove the current Proof Submitter. If there is no designated Proof Submitter, ONLY the Pod Owner can call `verifyWithdrawalCredentials`/`startCheckpoint`.

*Effects*:
* Updates `proofSubmitter` to `newProofSubmitter`

*Requirements*:
* Caller MUST be the Pod Owner

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

Handles the call to the beacon chain deposit contract. Only called via `EigenPodManager.stake`.

*Effects*:
* Deposits 32 ETH into the beacon chain deposit contract, and provides the pod's address as the deposit's withdrawal credentials

*Requirements*:
* Caller MUST be the `EigenPodManager`
* Call value MUST be 32 ETH
* Deposit contract `deposit` method MUST succeed given the provided `pubkey`, `signature`, and `depositDataRoot`

#### `withdrawRestakedBeaconChainETH`

```solidity
function withdrawRestakedBeaconChainETH(
    address recipient, 
    uint256 amountWei
)
    external 
    onlyEigenPodManager
```

The `EigenPodManager` calls this method when withdrawing a Pod Owner's shares as tokens (native ETH). The input `amountWei` is converted to Gwei and subtracted from `withdrawableRestakedExecutionLayerGwei`, which tracks native ETH balance that has been accounted for in a checkpoint (see [Checkpointing Validators](#checkpointing-validators)).

If the `EigenPod` does not have `amountWei` available to transfer, this method will revert

*Effects*:
* Decreases the pod's `withdrawableRestakedExecutionLayerGwei` by `amountWei / GWEI_TO_WEI`
* Sends `amountWei` ETH to `recipient`

*Requirements*:
* `amountWei / GWEI_TO_WEI` MUST NOT be greater than the proven `withdrawableRestakedExecutionLayerGwei`
* Pod MUST have at least `amountWei` ETH balance
* `recipient` MUST NOT revert when transferred `amountWei`
* `amountWei` MUST be a whole Gwei amount

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

Allows the Pod Owner to rescue ERC20 tokens accidentally sent to the `EigenPod`.

*Effects:*
* Calls `transfer` on each of the ERC20's in `tokenList`, sending the corresponding `amountsToWithdraw` to the `recipient`

*Requirements:*
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_NON_PROOF_WITHDRAWALS`
* `tokenList` and `amountsToWithdraw` MUST have equal lengths

