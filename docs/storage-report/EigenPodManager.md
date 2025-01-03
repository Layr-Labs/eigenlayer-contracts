| Name                            | Type                                                                       | Slot | Offset | Bytes | Contract                                               |
|---------------------------------|----------------------------------------------------------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized                    | uint8                                                                      | 0    | 0      | 1     | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| _initializing                   | bool                                                                       | 0    | 1      | 1     | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __gap                           | uint256[50]                                                                | 1    | 0      | 1600  | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| _owner                          | address                                                                    | 51   | 0      | 20    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __gap                           | uint256[49]                                                                | 52   | 0      | 1568  | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __deprecated_pauserRegistry     | contract IPauserRegistry                                                   | 101  | 0      | 20    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| _paused                         | uint256                                                                    | 102  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __gap                           | uint256[48]                                                                | 103  | 0      | 1536  | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __deprecated_beaconChainOracle  | address                                                                    | 151  | 0      | 20    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| ownerToPod                      | mapping(address => contract IEigenPod)                                     | 152  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| numPods                         | uint256                                                                    | 153  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __deprecated_maxPods            | uint256                                                                    | 154  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| podOwnerDepositShares           | mapping(address => int256)                                                 | 155  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __deprecated_denebForkTimestamp | uint64                                                                     | 156  | 0      | 8     | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| _beaconChainSlashingFactor      | mapping(address => struct IEigenPodManagerTypes.BeaconChainSlashingFactor) | 157  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __gap                           | uint256[43]                                                                | 158  | 0      | 1376  | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| _status                         | uint256                                                                    | 201  | 0      | 32    | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
| __gap                           | uint256[49]                                                                | 202  | 0      | 1568  | src/contracts/pods/EigenPodManager.sol:EigenPodManager |
