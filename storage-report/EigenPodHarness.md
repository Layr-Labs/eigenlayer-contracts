| Name                                            | Type                                               | Slot | Offset | Bytes | Contract                                               |
|-------------------------------------------------|----------------------------------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized                                    | uint8                                              | 0    | 0      | 1     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| _initializing                                   | bool                                               | 0    | 1      | 1     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| _status                                         | uint256                                            | 1    | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __gap                                           | uint256[49]                                        | 2    | 0      | 1568  | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| podOwner                                        | address                                            | 51   | 0      | 20    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __deprecated_mostRecentWithdrawalTimestamp      | uint64                                             | 51   | 20     | 8     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| withdrawableRestakedExecutionLayerGwei          | uint64                                             | 52   | 0      | 8     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __deprecated_hasRestaked                        | bool                                               | 52   | 8      | 1     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __deprecated_provenWithdrawal                   | mapping(bytes32 => mapping(uint64 => bool))        | 53   | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| _validatorPubkeyHashToInfo                      | mapping(bytes32 => struct IEigenPod.ValidatorInfo) | 54   | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __deprecated_nonBeaconChainETHBalanceWei        | uint256                                            | 55   | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __deprecated_sumOfPartialWithdrawalsClaimedGwei | uint64                                             | 56   | 0      | 8     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| activeValidatorCount                            | uint256                                            | 57   | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| lastCheckpointTimestamp                         | uint64                                             | 58   | 0      | 8     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| currentCheckpointTimestamp                      | uint64                                             | 58   | 8      | 8     | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| checkpointBalanceExitedGwei                     | mapping(uint64 => uint64)                          | 59   | 0      | 32    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| _currentCheckpoint                              | struct IEigenPod.Checkpoint                        | 60   | 0      | 64    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| proofSubmitter                                  | address                                            | 62   | 0      | 20    | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
| __gap                                           | uint256[36]                                        | 63   | 0      | 1152  | src/test/harnesses/EigenPodHarness.sol:EigenPodHarness |
