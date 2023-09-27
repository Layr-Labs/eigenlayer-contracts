## EigenPodManager

Technical details on the EigenPod subsystem as it functions during M2. Includes:
* EigenPodManager
* EigenPods
* Native restaking
* Beacon chain proofs
* Stake / proof / withdrawal flows

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`EigenPodManager.sol`](../../src/contracts/pods/EigenPodManager.sol) | Singleton | Transparent proxy |
| [`EigenPod.sol`](../../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy |
| [`DelayedWithdrawalRouter.sol`](../../src/contracts/pods/DelayedWithdrawalRouter.sol) | Singleton | Transparent proxy |
| [TODO - BeaconChainOracle](#TODO) | TODO | TODO |

The `EigenPodManager` and `EigenPod` contracts allow Stakers to restake beacon chain ETH which can then be delegated to Operators via the `DelegationManager`.

These two contracts do stuff!

### Stakers (Before Restaking)

Before a Staker "enters" the EigenLayer system, they need to:
* deploy an `EigenPod`
* start a beacon chain validator and point its withdrawal credentials at the `EigenPod`
* provide a beacon chain state proof that shows their validator has sufficient balance and has withdrawal credentials pointed at their `EigenPod`

The following methods concern these steps:

#### `EigenPodManager.createPod`

```solidity
function createPod() external
```

*Effects*:

*Requirements*:

#### `EigenPodManager.stake`

```solidity
function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable
```

*Effects*:

*Requirements*:

#### `EigenPod.withdrawBeforeRestaking`

```solidity
function withdrawBeforeRestaking() external onlyEigenPodOwner hasNeverRestaked
```

*Effects*:

*Requirements*:

#### `EigenPod.activateRestaking`

```solidity
function activateRestaking() external onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS) 
    onlyEigenPodOwner 
    hasNeverRestaked
```

This final method allows a Pod Owner to designate their pod (and any future ETH sent to it) as being restaked. Calling this method first withdraws any ETH in the `EigenPod` via the `DelayedWithdrawalRouter`, and then prevents further calls to `withdrawBeforeRestaking`.

Withdrawing any future ETH sent via beacon chain withdrawal to the `EigenPod` requires providing beacon chain state proofs. However, sending ETH to

*Effects*:

*Requirements*:
* Caller MUST be the Pod Owner
* Pod MUST NOT have already activated restaking