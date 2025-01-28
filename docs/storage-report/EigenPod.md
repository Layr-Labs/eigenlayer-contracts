| Name                                            | Type                                                    | Slot | Offset | Bytes | Contract                                 |
|-------------------------------------------------|---------------------------------------------------------|------|--------|-------|------------------------------------------|
| _initialized                                    | uint8                                                   | 0    | 0      | 1     | src/contracts/pods/EigenPod.sol:EigenPod |
| _initializing                                   | bool                                                    | 0    | 1      | 1     | src/contracts/pods/EigenPod.sol:EigenPod |
| _status                                         | uint256                                                 | 1    | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| __gap                                           | uint256[49]                                             | 2    | 0      | 1568  | src/contracts/pods/EigenPod.sol:EigenPod |
| podOwner                                        | address                                                 | 51   | 0      | 20    | src/contracts/pods/EigenPod.sol:EigenPod |
| __deprecated_mostRecentWithdrawalTimestamp      | uint64                                                  | 51   | 20     | 8     | src/contracts/pods/EigenPod.sol:EigenPod |
| restakedExecutionLayerGwei                      | uint64                                                  | 52   | 0      | 8     | src/contracts/pods/EigenPod.sol:EigenPod |
| __deprecated_hasRestaked                        | bool                                                    | 52   | 8      | 1     | src/contracts/pods/EigenPod.sol:EigenPod |
| __deprecated_provenWithdrawal                   | mapping(bytes32 => mapping(uint64 => bool))             | 53   | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| _validatorPubkeyHashToInfo                      | mapping(bytes32 => struct IEigenPodTypes.ValidatorInfo) | 54   | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| __deprecated_nonBeaconChainETHBalanceWei        | uint256                                                 | 55   | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| __deprecated_sumOfPartialWithdrawalsClaimedGwei | uint64                                                  | 56   | 0      | 8     | src/contracts/pods/EigenPod.sol:EigenPod |
| activeValidatorCount                            | uint256                                                 | 57   | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| lastCheckpointTimestamp                         | uint64                                                  | 58   | 0      | 8     | src/contracts/pods/EigenPod.sol:EigenPod |
| currentCheckpointTimestamp                      | uint64                                                  | 58   | 8      | 8     | src/contracts/pods/EigenPod.sol:EigenPod |
| checkpointBalanceExitedGwei                     | mapping(uint64 => uint64)                               | 59   | 0      | 32    | src/contracts/pods/EigenPod.sol:EigenPod |
| _currentCheckpoint                              | struct IEigenPodTypes.Checkpoint                        | 60   | 0      | 64    | src/contracts/pods/EigenPod.sol:EigenPod |
| proofSubmitter                                  | address                                                 | 62   | 0      | 20    | src/contracts/pods/EigenPod.sol:EigenPod |
| __gap                                           | uint256[35]                                             | 63   | 0      | 1120  | src/contracts/pods/EigenPod.sol:EigenPod |
