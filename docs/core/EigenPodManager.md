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
        * Since balances on the consensus layer are stored only in Gwei amounts, the EigenPodManager enforces the invariant that `podOwnerShares` is always a whole Gwei amount for every staker, i.e. `podOwnerShares[staker] % 1e9 == 0` always.
* `EigenPod`:
    * `_validatorPubkeyHashToInfo[bytes32] -> (ValidatorInfo)`: individual validators are identified within an `EigenPod` according the their public key hash. This mapping keeps track of the following for each validator:
        * `validatorStatus`: (`INACTIVE`, `ACTIVE`, `WITHDRAWN`)
        * `validatorIndex`: A `uint40` that is unique for each validator making a successful deposit via the deposit contract
        * `mostRecentBalanceUpdateTimestamp`: A timestamp that represents the most recent successful proof of the validator's effective balance
        * `restakedBalanceGwei`: set to the validator's balance.
    * `withdrawableRestakedExecutionLayerGwei`: When a Staker proves that a validator has exited from the beacon chain, the withdrawal amount is added to this variable. When completing a withdrawal of beacon chain ETH, the withdrawal amount is subtracted from this variable. See also:
        * [`DelegationManager`: "Undelegating and Withdrawing"](./DelegationManager.md#undelegating-and-withdrawing)
        * [`EigenPodManager`: "Withdrawal Processing"](#withdrawal-processing)

#### Important Definitions

* "Pod Owner": A Staker who has deployed an `EigenPod` is a Pod Owner. The terms are used interchangeably in this document.
    * Pod Owners can only deploy a single `EigenPod`, but can restake any number of beacon chain validators from the same `EigenPod`.
    * Pod Owners can delegate their `EigenPodManager` shares to Operators (via `DelegationManager`).
    * These shares correspond to the amount of provably-restaked beacon chain ETH held by the Pod Owner via their `EigenPod`.
* `EigenPod`:
    * `_podWithdrawalCredentials() -> (bytes memory)`:
        * Gives `abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(EigenPod))`
        * These are the `0x01` withdrawal credentials of the `EigenPod`, used as a validator's withdrawal credentials on the beacon chain.

---    

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

*Beacon chain proofs used*:
* [`verifyStateRootAgainstLatestBlockRoot`](./proofs/BeaconChainProofs.md#beaconchainproofsverifystaterootagainstlatestblockroot)
* [`verifyValidatorFields`](./proofs/BeaconChainProofs.md#beaconchainproofsverifyvalidatorfields)

*Effects*:
* For each validator (`_validatorPubkeyHashToInfo[pubkeyHash]`) the validator's info is set for the first time:
    * `VALIDATOR_STATUS` moves from `INACTIVE` to `ACTIVE`
    * `validatorIndex` is recorded
    * `mostRecentBalanceUpdateTimestamp` is set to the `oracleTimestamp` used to fetch the beacon block root
    * `restakedBalanceGwei` is set to the validator's effective balance
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

The primary method concerning actively restaked validators is:
* [`EigenPod.verifyBalanceUpdate`](#eigenpodverifybalanceupdate)

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

Note that if a validator's balance has decreased, this method will result in shares being removed from the Pod Owner in `EigenPodManager.recordBeaconChainETHBalanceUpdate`. This may cause the Pod Owner's balance to go negative in some cases, representing a "deficit" that must be repaid before any withdrawals can be processed. One example flow where this might occur is:
* Pod Owner calls `DelegationManager.undelegate`, which queues a withdrawal in the `DelegationManager`. The Pod Owner's shares are set to 0 while the withdrawal is in the queue.
* Pod Owner's beacon chain ETH balance decreases (maybe due to slashing), and someone provides a proof of this to `EigenPod.verifyBalanceUpdate`. In this case, the Pod Owner will have negative shares in the `EigenPodManager`.
* After a delay, the Pod Owner calls `DelegationManager.completeQueuedWithdrawal`. The negative shares are then repaid out of the withdrawn assets.

For the validator whose balance should be updated, the caller must supply:
* `validatorIndex`: the validator's `ValidatorIndex` (see [consensus specs](https://eth2book.info/capella/part3/config/types/#validatorindex))
* `stateRootProof`: a proof that will verify `stateRootProof.beaconStateRoot` against an oracle-provided beacon block root
* `balanceUpdateProof`: contains the `validatorFieldsProof` mentioned in `verifyWithdrawalCredentials`, as well as a proof that `balanceUpdateProof.balanceRoot` contains the validator's current balance
* `validatorFields`: the fields of the `Validator` container associated with the validator (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator))
* `oracleTimestamp`: a timestamp used to fetch a beacon block root from `EigenPodManager.beaconChainOracle`

*Beacon chain proofs used*:
* [`verifyStateRootAgainstLatestBlockRoot`](./proofs/BeaconChainProofs.md#beaconchainproofsverifystaterootagainstlatestblockroot)
* [`verifyValidatorFields`](./proofs/BeaconChainProofs.md#beaconchainproofsverifyvalidatorfields)
* [`verifyValidatorBalance`](./proofs/BeaconChainProofs.md#beaconchainproofsverifyvalidatorbalance)

*Effects*:
* Updates the validator's stored info:
    * `restakedBalanceGwei` is updated to the newly-proven balance
    * `mostRecentBalanceUpdateTimestamp` is set to the `oracleTimestamp` used to fetch the beacon block root
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE`
* Balance updates should only be made before a validator has fully exited. If the validator has exited, any further proofs should follow the `verifyAndProcessWithdrawals` path.
    * This is to prevent someone providing a "balance update" on an exited validator that "proves" a balance of 0, when we want to process that update as a withdrawal instead.
* `oracleTimestamp`:
    * MUST be no more than `VERIFY_BALANCE_UPDATE_WINDOW_SECONDS` (~4.5 hrs) old
    * MUST be newer than the validator's `mostRecentBalanceUpdateTimestamp`
    * MUST be queryable via `EigenPodManager.getBlockRootAtTimestamp` (fails if `stateRoot == 0`)
* `validatorFields[0]` MUST be a pubkey hash corresponding to a validator whose withdrawal credentials have been proven, and is not yet withdrawn (`VALIDATOR_STATUS.ACTIVE`)
* `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot` MUST verify the provided `beaconStateRoot` against the oracle-provided `latestBlockRoot`
* `BeaconChainProofs.verifyValidatorFields` MUST verify the provided `validatorFields` against the `beaconStateRoot`
* `BeaconChainProofs.verifyValidatorBalance` MUST verify the provided `balanceRoot` against the `beaconStateRoot`
* See [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

---

### Withdrawal Processing

The `DelegationManager` is the entry point for all undelegation and withdrawals, which must be queued for a time before being completed. When a withdrawal is initiated, the following method is used:
* [`EigenPodManager.removeShares`](#eigenpodmanagerremoveshares)

When completing a queued undelegation or withdrawal, the `DelegationManager` calls one of these two methods:
* [`EigenPodManager.addShares`](#eigenpodmanageraddshares)
* [`EigenPodManager.withdrawSharesAsTokens`](#eigenpodmanagerwithdrawsharesastokens)
    * [`EigenPod.withdrawRestakedBeaconChainETH`](#eigenpodwithdrawrestakedbeaconchaineth)

If a Staker wishes to fully withdraw their beacon chain ETH (via `withdrawSharesAsTokens`), they need to exit their validator and prove the withdrawal *prior to* completing the queued withdrawal. They do so using this method:
* [`EigenPod.verifyAndProcessWithdrawals`](#eigenpodverifyandprocesswithdrawals)

Some withdrawals are sent to their destination via the `DelayedWithdrawalRouter`:
* [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)
* [`DelayedWithdrawalRouter.claimDelayedWithdrawals`](#delayedwithdrawalrouterclaimdelayedwithdrawals)

#### `EigenPodManager.removeShares`

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

#### `EigenPodManager.addShares`

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

#### `EigenPodManager.withdrawSharesAsTokens`

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

Note that because this method entails withdrawing and sending beacon chain ETH, two conditions must be met for this method to succeed: (i) the ETH being withdrawn should already be in the `EigenPod`, and (ii) the beacon chain withdrawals responsible for the ETH should already be proven.

This means that before completing their queued withdrawal, a Pod Owner needs to prove their beacon chain withdrawals via `EigenPod.verifyAndProcessWithdrawals`.

Also note that, like `addShares`, if the original Pod Owner has a share deficit (negative shares), the deficit is repaid out of the withdrawn `shares` before any native ETH is withdrawn.

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* If `podOwner's` share balance is negative, `shares` are added until the balance hits 0
    * Any remaining shares are withdrawn and sent to `destination` (see [`EigenPod.withdrawRestakedBeaconChainETH`](#eigenpodwithdrawrestakedbeaconchaineth))

*Requirements*:
* `podOwner` MUST NOT be zero
* `destination` MUST NOT be zero
* `shares` MUST NOT be negative when converted to an `int256`
* `shares` MUST be a whole Gwei amount
* See [`EigenPod.withdrawRestakedBeaconChainETH`](#eigenpodwithdrawrestakedbeaconchaineth)

##### `EigenPod.withdrawRestakedBeaconChainETH`

```solidity
function withdrawRestakedBeaconChainETH(
    address recipient, 
    uint256 amountWei
) 
    external 
    onlyEigenPodManager
```

The `EigenPodManager` calls this method when withdrawing a Pod Owner's shares as tokens (native ETH). The input `amountWei` is converted to Gwei and subtracted from `withdrawableRestakedExecutionLayerGwei`, which tracks Gwei that has been provably withdrawn (via `EigenPod.verifyAndProcessWithdrawals`).

As such:
* If a withdrawal has not been proven that sufficiently raises `withdrawableRestakedExecutionLayerGwei`, this method will revert.
* If the `EigenPod` does not have `amountWei` available to transfer, this method will revert

*Effects*:
* Decreases the pod's `withdrawableRestakedExecutionLayerGwei` by `amountWei / GWEI_TO_WEI`
* Sends `amountWei` ETH to `recipient`

*Requirements*:
* `amountWei / GWEI_TO_WEI` MUST NOT be greater than the proven `withdrawableRestakedExecutionLayerGwei`
* Pod MUST have at least `amountWei` ETH balance
* `recipient` MUST NOT revert when transferred `amountWei`
* `amountWei` MUST be a whole Gwei amount

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

Whether each withdrawal is a full or partial withdrawal is determined by the validator's "withdrawable epoch" in the `Validator` container given by `validatorFields` (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator)). If the withdrawal proof timestamp is after this epoch, the withdrawal is a full withdrawal.
* Partial withdrawals are performed automatically by the beacon chain when a validator has an effective balance over 32 ETH. This method can be used to prove that these withdrawals occurred, allowing the Pod Owner to withdraw the excess ETH (via [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)).
* Full withdrawals are performed when a Pod Owner decides to fully exit a validator from the beacon chain. To do this, the Pod Owner should follow these steps: 
    1. Undelegate or queue a withdrawal (via the `DelegationManager`: ["Undelegating and Withdrawing"](./DelegationManager.md#undelegating-and-withdrawing))
    2. Exit their validator from the beacon chain and provide a proof to this method
    3. Complete their withdrawal (via [`DelegationManager.completeQueuedWithdrawal`](./DelegationManager.md#completequeuedwithdrawal)).

If the Pod Owner only exits their validator, the ETH of the pod owner is still staked through EigenLayer and can be used to service AVSs, even though their ETH has been withdrawn from the beacon chain. The protocol allows for this edge case. 

*Beacon chain proofs used*:
* [`verifyStateRootAgainstLatestBlockRoot`](./proofs/BeaconChainProofs.md#beaconchainproofsverifystaterootagainstlatestblockroot)
* [`verifyWithdrawal`](./proofs/BeaconChainProofs.md#beaconchainproofsverifywithdrawal)
* [`verifyValidatorFields`](./proofs/BeaconChainProofs.md#beaconchainproofsverifyvalidatorfields)

*Effects*:
* For each proven withdrawal:
    * The validator in question is recorded as having a proven withdrawal at the timestamp given by `withdrawalProof.timestampRoot`
        * This is to prevent the same withdrawal from being proven twice
    * If this is a full withdrawal:
        * Any withdrawal amount in excess of `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR` is immediately withdrawn (see [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal))
            * The remainder (`MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR`) must be withdrawn through the `DelegationManager's` withdrawal flow, but in the meantime is added to `EigenPod.withdrawableRestakedExecutionLayerGwei`
        * If the amount being withdrawn is not equal to the current accounted-for validator balance, a `shareDelta` is calculated to be sent to ([`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)).
        * The validator's info is updated to reflect its `WITHDRAWN` status, and `restakedBalanceGwei` is set to 0
    * If this is a partial withdrawal:
        * The withdrawal amount is added to `sumOfPartialWithdrawalsClaimedGwei`
        * The withdrawal amount is withdrawn (via [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal))

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

#### `DelayedWithdrawalRouter.createDelayedWithdrawal`

```solidity
function createDelayedWithdrawal(
    address podOwner,
    address recipient
) 
    external 
    payable 
    onlyEigenPod(podOwner) 
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS)
```

Used by `EigenPods` to queue a withdrawal of beacon chain ETH that can be claimed by a `recipient` after `withdrawalDelayBlocks` have passed.

*Effects*:
* Creates a `DelayedWithdrawal` for the `recipient` in the amount of `msg.value`, starting at the current block

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_DELAYED_WITHDRAWAL_CLAIMS`
* Caller MUST be the `EigenPod` associated with the `podOwner`
* `recipient` MUST NOT be zero

#### `DelayedWithdrawalRouter.claimDelayedWithdrawals`

```solidity
function claimDelayedWithdrawals(
    address recipient,
    uint256 maxNumberOfDelayedWithdrawalsToClaim
) 
    external 
    nonReentrant 
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS)

// (Uses `msg.sender` as `recipient`)
function claimDelayedWithdrawals(
    uint256 maxNumberOfDelayedWithdrawalsToClaim
) 
    external 
    nonReentrant 
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS)
```

After `withdrawalDelayBlocks`, withdrawals can be claimed using these methods. Claims may be processed on behalf of someone else by passing their address in as the `recipient`. Otherwise, claims are processed on behalf of `msg.sender`.

This method loops over up to `maxNumberOfDelayedWithdrawalsToClaim` withdrawals, tallys each withdrawal amount, and sends the total to the `recipient`.

*Effects*:
* Updates the `recipient's` `delayedWithdrawalsCompleted`
* Sends ETH from completed withdrawals to the `recipient`

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_DELAYED_WITHDRAWAL_CLAIMS`

---

### System Configuration

* [`EigenPodManager.setMaxPods`](#eigenpodmanagersetmaxpods)
* [`EigenPodManager.updateBeaconChainOracle`](#eigenpodmanagerupdatebeaconchainoracle)
* [`DelayedWithdrawalRouter.setWithdrawalDelayBlocks`](#delayedwithdrawalroutersetwithdrawaldelayblocks)

#### `EigenPodManager.setMaxPods`

```solidity
function setMaxPods(uint256 newMaxPods) external onlyUnpauser
```

Allows the unpauser to update the maximum number of `EigenPods` that the `EigenPodManager` can create.

*Effects*:
* Updates `EigenPodManager.maxPods`

*Requirements*:
* Caller MUST be the unpauser

#### `EigenPodManager.updateBeaconChainOracle`

```solidity
function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyOwner
```

Allows the owner to update the address of the oracle used by `EigenPods` to retrieve beacon chain state roots (used when verifying beacon chain state proofs).

*Effects*:
* Updates `EigenPodManager.beaconChainOracle`

*Requirements*:
* Caller MUST be the owner

#### `DelayedWithdrawalRouter.setWithdrawalDelayBlocks`

```solidity
function setWithdrawalDelayBlocks(uint256 newValue) external onlyOwner
```

Allows the `DelayedWithdrawalRouter` to update the delay between withdrawal creation and claimability.

The new delay can't exceed `MAX_WITHDRAWAL_DELAY_BLOCKS`.

*Effects*:
* Updates `DelayedWithdrawalRouter.withdrawalDelayBlocks`

*Requirements*:
* Caller MUST be the owner
* `newValue` MUST NOT be greater than `MAX_WITHDRAWAL_DELAY_BLOCKS`

---

### Other Methods

This section details various methods that don't fit well into other sections.

Stakers' balance updates are accounted for when the Staker's `EigenPod` calls this method:
* [`EigenPodManager.recordBeaconChainETHBalanceUpdate`](#eigenpodmanagerrecordbeaconchainethbalanceupdate)

*For pods deployed prior to M2*, the following methods are callable:
* [`EigenPod.activateRestaking`](#eigenpodactivaterestaking)
* [`EigenPod.withdrawBeforeRestaking`](#eigenpodwithdrawbeforerestaking)

The `EigenPod` also includes two token recovery mechanisms:
* [`EigenPod.withdrawNonBeaconChainETHBalanceWei`](#eigenpodwithdrawnonbeaconchainethbalancewei)
* [`EigenPod.recoverTokens`](#eigenpodrecovertokens)

#### `EigenPodManager.recordBeaconChainETHBalanceUpdate`

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

*Entry Points*:
* `EigenPod.verifyWithdrawalCredentials`
* `EigenPod.verifyBalanceUpdate`
* `EigenPod.verifyAndProcessWithdrawals`

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
    * MUST be a whole Gwei amount

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
* Sets the pod's `nonBeaconChainETHBalanceWei` to 0 (only incremented in the fallback function)
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
* Sets the pod's `nonBeaconChainETHBalanceWei` to 0 (only incremented in the fallback function)
* Updates the pod's most recent withdrawal timestamp to the current time
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*Requirements*:
* Caller MUST be the Pod Owner
* Pod MUST NOT have already activated restaking
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*As of M2*: restaking is automatically activated for newly-deployed `EigenPods`, making this method uncallable for pods deployed after M2. However, for `EigenPods` deployed *before* M2, restaking may not be active, and this method may be callable.

#### `EigenPod.withdrawNonBeaconChainETHBalanceWei`

```solidity
function withdrawNonBeaconChainETHBalanceWei(
    address recipient,
    uint256 amountToWithdraw
) 
    external 
    onlyEigenPodOwner
```

Allows the Pod Owner to withdraw ETH accidentally sent to the contract's `receive` function.

The `receive` function updates `nonBeaconChainETHBalanceWei`, which this function uses to calculate how much can be withdrawn.

Withdrawals from this function are sent via the `DelayedWithdrawalRouter`, and can be claimed by the passed-in `recipient`.

*Effects:*
* Decrements `nonBeaconChainETHBalanceWei`
* Sends `amountToWithdraw` wei to [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)

*Requirements:*
* Caller MUST be the Pod Owner
* `amountToWithdraw` MUST NOT be greater than the amount sent to the contract's `receive` function
* See [`DelayedWithdrawalRouter.createDelayedWithdrawal`](#delayedwithdrawalroutercreatedelayedwithdrawal)

#### `EigenPod.recoverTokens`

```solidity
function recoverTokens(
    IERC20[] memory tokenList,
    uint256[] memory amountsToWithdraw,
    address recipient
) 
    external 
    onlyEigenPodOwner
```

Allows the Pod Owner to rescue ERC20 tokens accidentally sent to the `EigenPod`.

*Effects:*
* Calls `transfer` on each of the ERC20's in `tokenList`, sending the corresponding `amountsToWithdraw` to the `recipient`

*Requirements:*
* `tokenList` and `amountsToWithdraw` MUST have equal lengths
* Caller MUST be the Pod Owner