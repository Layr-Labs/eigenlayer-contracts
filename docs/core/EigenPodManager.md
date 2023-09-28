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
| [TODO - BeaconChainOracle](#TODO) | TODO | TODO |

The `EigenPodManager` and `EigenPod` contracts allow Stakers to restake beacon chain ETH which can then be delegated to Operators via the `DelegationManager`.

`EigenPods` enable a Staker to restake beacon chain ETH by serving as the withdrawal credentials for a beacon chain validator. This allows EigenLayer to impose conditions on validator withdrawals - namely, that withdrawals (i) are supported with a beacon chain state proof, and (ii) are processed through a withdrawal queue.

*Important state variables*:
* `mapping(address => IEigenPod) public ownerToPod`
* `mapping(address => uint256) public podOwnerShares`
* TODO

*Helpful definitions*:
* `EigenPodManager`:
    * TODO
* `EigenPod`:
    * `_podWithdrawalCredentials() -> (bytes memory)`:
        * TODO
    * `_validatorPubkeyHashToInfo`
        * TODO (validator status is here)

### Stakers (Before Restaking)

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

*Effects*:
* TODO

*Requirements*:
* TODO

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

TODO only callable by above method

*Effects*:
* TODO

*Requirements*:
* TODO

### Stakers (Active Restaking)

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