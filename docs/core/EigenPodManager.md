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

`EigenPods` enable a Staker to restake beacon chain ETH by serving as the withdrawal credentials for a beacon chain validator. This allows EigenLayer to impose conditions on validator withdrawals - namely, that withdrawals (i) are supported with a beacon chain state proof, and (ii) are processed through a withdrawal queue.

*Important state variables*:
* `mapping(address => IEigenPod) public ownerToPod`
* `mapping(address => uint256) public podOwnerShares`
* `EigenPod`:
    * `_validatorPubkeyHashToInfo[bytes32] -> (ValidatorInfo)`: individual validators are identified within an `EigenPod` according the their public key hash. This mapping keeps track of the following for each validator:
        * `validatorStatus`: (`INACTIVE`, `ACTIVE`, `WITHDRAWN`)
        * `validatorIndex`: A `uint40` that is unique for each validator making a successful deposit via the deposit contract
        * `mostRecentBalanceUpdateTimestamp`: A timestamp that represents the most recent successful proof of the validator's effective balance
        * `restakedBalanceGwei`: Calculated against the validator's proven effective balance using `_calculateRestakedBalanceGwei`

*Helpful definitions*:
* `EigenPodManager`:
    * TODO
* `EigenPod`:
    * `_podWithdrawalCredentials() -> (bytes memory)`:
        * TODO
    * `_calculateRestakedBalanceGwei`

### Stakers (Before Activating Validator)

Before a Staker "enters" the EigenLayer system, they need to:
1. Deploy an `EigenPod`
2. Start a beacon chain validator and point its withdrawal credentials at the `EigenPod`
    * In case of a validator with BLS withdrawal credentials, its withdrawal credentials need to be updated to point at the `EigenPod`
3. Provide a beacon chain state proof that shows their validator has sufficient balance and has withdrawal credentials pointed at their `EigenPod`

The following top-level callable methods concern these steps:
* [`EigenPodManager.createPod`](#eigenpodmanagercreatepod)
* [`EigenPodManager.stake`](#eigenpodmanagerstake)
* [`EigenPod.activateRestaking`](#eigenpodactivaterestaking)
* [`EigenPod.withdrawBeforeRestaking`](#eigenpodwithdrawbeforerestaking)
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
* `EigenPod.initialize(msg.sender)`: initializes the pod, setting the caller as the Pod Owner
* Maps the new pod to the Pod Owner (each address can only deploy a single `EigenPod`)

*Requirements*:
* Caller MUST NOT have deployed an `EigenPod` already
* Pause status MUST NOT be set: `PAUSED_NEW_EIGENPODS`

*As of M2*:
* `EigenPods` are initialized with restaking enabled by default (`hasRestaked = true`). Pods deployed before M2 may not have this enabled, and will need to call `EigenPod.activateRestaking()`.

#### `EigenPodManager.stake`

```solidity
function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable
```

Allows a Staker to deposit 32 ETH into the beacon chain deposit contract, provided the credentials for the Staker's beacon chain validator. The `EigenPod.stake` method is called, which automatically calculates the correct withdrawal credentials for the pod and passes these to the deposit contract along with the 32 ETH.

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
) external payable onlyEigenPodManager
```

Handles the call to the beacon chain deposit contract. Only called via `EigenPodManager.stake`.

*Effects*:
* Deposits 32 ETH into the beacon chain deposit contract, and provides the pod's address as the deposit's withdrawal credentials

*Requirements*:
* Caller MUST be the `EigenPodManager`
* Call value MUST be 32 ETH
* Deposit contract `deposit` method MUST succeed given the provided `pubkey`, `signature`, and `depositDataRoot`

#### `EigenPod.activateRestaking`

```solidity
function activateRestaking()
    external
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
    onlyEigenPodOwner
    hasNeverRestaked
```

This method allows a Pod Owner to designate their pod (and any future ETH sent to it) as being restaked. Calling this method first withdraws any ETH in the `EigenPod` via the `DelayedWithdrawalRouter`, and then prevents further calls to `withdrawBeforeRestaking`.

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

Allows the Pod Owner to withdraw any ETH in the `EigenPod` via the `DelayedWithdrawalRouter`, assuming restaking has not yet been activated. See [`EigenPod.activateRestaking`](#eigenpodactivaterestaking) for more details.

*Effects*:
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*Requirements*:
* Caller MUST be the Pod Owner
* Pod MUST NOT have already activated restaking
* See [DelayedWithdrawalRouter.createDelayedWithdrawal](#delayedwithdrawalroutercreatedelayedwithdrawal)

*As of M2*: restaking is automatically activated for newly-deployed `EigenPods`, making this method uncallable for pods deployed after M2. However, for `EigenPods` deployed *before* M2, restaking may not be active, and this method may be callable.

#### `EigenPod.verifyWithdrawalCredentials`

```solidity
function verifyWithdrawalCredentials(
    uint64 oracleTimestamp,
    uint40[] calldata validatorIndices,
    BeaconChainProofs.WithdrawalCredentialProof[] calldata withdrawalCredentialProofs,
    bytes32[][] calldata validatorFields
)
    external
    onlyEigenPodOwner
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
    proofIsForValidTimestamp(oracleTimestamp)
    hasEnabledRestaking
```

Once a Pod Owner has deposited ETH into the beacon chain deposit contract, they can call this method to "restake" one or more validators - meaning the ETH in those validators:
* is awarded to the Staker/Pod Owner in `EigenPodManager.podOwnerShares`
* is delegatable to an Operator (via the `DelegationManager`)

For each validator the Pod Owner wants to restake, they must supply:
* `validatorIndices`: their validator's `ValidatorIndex` (see [consensus specs](https://eth2book.info/capella/part3/config/types/#validatorindex))
* `validatorFields`: the fields of the `Validator` container associated with the validator (see [consensus specs](https://eth2book.info/capella/part3/containers/dependencies/#validator))
* `withdrawalCredentialProofs`: a proof that the `Validator` container belongs to the associated `ValidatorIndex`, according to the beacon chain state
* `oracleTimestamp`: a timestamp used to fetch a beacon block root from `EigenPodManager.beaconChainOracle`

For each successfully proven validator, that validator's status becomes `VALIDATOR_STATUS.ACTIVE`, and the sum of restakable ether across all proven validators is provided to `EigenPodManager.restakeBeaconChainETH`, where it is added to the Pod Owner's shares. If the Pod Owner is delegated to an Operator via the `DelegationManager`, this sum is also added to the Operator's delegated shares for the beacon chain ETH strategy.

*Effects*:
* For each validator (`_validatorPubkeyHashToInfo[pubkeyHash]`) the validator's info is set for the first time:
    * `VALIDATOR_STATUS` moves from `INACTIVE` to `ACTIVE`
    * `validatorIndex` is recorded
    * `mostRecentBalanceUpdateTimestamp` is set to the `oracleTimestamp` used to fetch the beacon block root
    * `restakedBalanceGwei` is set to `_calculateRestakedBalanceGwei(effectiveBalance)`
* See [`EigenPodManager.restakeBeaconChainETH`](#eigenpodmanagerrestakebeaconchaineth)

*Requirements*:
* Caller MUST be the Pod Owner
* Pause status MUST NOT be set: `PAUSED_EIGENPODS_VERIFY_CREDENTIALS`
* Pod MUST have enabled restaking
* All input array lengths MUST be equal
* `oracleTimestamp`:
    * MUST be greater than the `mostRecentWithdrawalTimestamp`
    * MUST be no more than `VERIFY_BALANCE_UPDATE_WINDOW_SECONDS` (~4.5 hrs) old
    * MUST be queryable via `EigenPodManager.getBlockRootAtTimestamp` (fails if `stateRoot == 0`)
* For each validator:
    * The validator's status MUST initially be `VALIDATOR_STATUS.INACTIVE`
    * `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot` MUST verify the provided `beaconStateRoot` against the oracle-provided `latestBlockRoot`
    * `BeaconChainProofs.verifyValidatorFields` MUST verify the provided `validatorFields` against the `beaconStateRoot`
    * The aforementioned proofs MUST show that the validator's withdrawal credentials are set to the `EigenPod`
* See [`EigenPodManager.restakeBeaconChainETH`](#eigenpodmanagerrestakebeaconchaineth)

*As of M2*:
* Restaking is enabled by default for pods deployed after M2. See `activateRestaking` for more info.

##### `EigenPodManager.restakeBeaconChainETH`

```solidity
function restakeBeaconChainETH(
    address podOwner,
    uint256 amountWei
) 
    external 
    onlyEigenPod(podOwner) 
    onlyNotFrozen(podOwner) 
    nonReentrant
```

This method is only called by `EigenPod.verifyWithdrawalCredentials`, during which the `EigenPod` verifies a validator's effective balance and withdrawal credentials using a beacon chain state proof. 

After verifying the balance of one or more validators, the `EigenPod` will sum the "restakable" balance of each validator and pass it to this method, which adds this balance to the Pod Owner's shares.

If the Pod Owner is not in undelegation limbo, the added shares are also sent to `DelegationManager.increaseDelegatedShares`, where they will be awarded to the Staker/Pod Owner's delegated Operator.

*Effects*:
* Adds `amountWei` shares to the Pod Owner's `EigenPodManager` shares
* If the Pod Owner is NOT in undelegation limbo: see [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Caller MUST be the `EigenPod` associated with the passed-in Pod Owner
* `amountWei` MUST NOT be zero
* If the Pod Owner is NOT in undelegation limbo: see [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

---

### Stakers (After Activating Validator)

TODO intro

#### `EigenPod.verifyBalanceUpdate`

```solidity

```

*Effects*:
* TODO

*Requirements*:
* TODO

#### `EigenPod.verifyAndProcessWithdrawals`

```solidity

```

*Effects*:
* TODO

*Requirements*:
* TODO

##### `EigenPodManager.recordBeaconChainETHBalanceUpdate`

```solidity

```

TODO only callable by above two methods

*Effects*:
* TODO

*Requirements*:
* TODO

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

*Effects*:
* TODO

*Requirements*:
* TODO

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

#### `EigenPodManager.completeQueuedWithdrawal`

```solidity
function completeQueuedWithdrawal(
    BeaconChainQueuedWithdrawal memory queuedWithdrawal,
    uint256 middlewareTimesIndex
)
    external
    onlyNotFrozen(queuedWithdrawal.delegatedAddress)
    nonReentrant
    onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
```

*Effects*:
* TODO

*Requirements*:
* TODO

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

### Other

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

*Effects*:
* TODO

*Requirements*:
* TODO

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

*Effects*:
* TODO

*Requirements*:
* TODO

*As of M2*:
* The `onlyNotFrozen` modifier is a no-op

#### `DelayedWithdrawalRouter.createDelayedWithdrawal`

#### `EigenPod.withdrawNonBeaconChainETHBalanceWei`

#### `EigenPod.recoverTokens`

### Unpauser

#### `EigenPodManager.setMaxPods`

```solidity
function setMaxPods(uint256 newMaxPods) external onlyUnpauser
```

*Effects*:
* TODO

*Requirements*:
* TODO

### Owner

#### `EigenPodManager.updateBeaconChainOracle`

```solidity
function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyOwner
```

*Effects*:
* TODO

*Requirements*:
* TODO