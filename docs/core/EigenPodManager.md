# EigenPods

Technical details on the EigenPod subsystem as it functions during M2. Includes:
* EigenPodManager
* EigenPods
* Native restaking
* Beacon chain proofs
* Stake / proof / withdrawal flows

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`EigenPodManager.sol`](#TODO) | Singleton | Transparent proxy |
| [`EigenPod.sol`](#TODO) | Instanced, deployed per-user | Beacon proxy |
| [`DelayedWithdrawalRouter.sol`](#TODO) | Singleton | Transparent proxy |
| [TODO - BeaconChainOracle](#TODO) | TODO | TODO |

## EigenPodManager

### Operators

### Stakers