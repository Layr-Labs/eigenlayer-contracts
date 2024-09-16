| Name                            | Type                                   | Slot | Offset | Bytes | Contract                                                             |
|---------------------------------|----------------------------------------|------|--------|-------|----------------------------------------------------------------------|
| __deprecated_beaconChainOracle  | address                                | 0    | 0      | 20    | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| ownerToPod                      | mapping(address => contract IEigenPod) | 1    | 0      | 32    | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| numPods                         | uint256                                | 2    | 0      | 32    | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| __deprecated_maxPods            | uint256                                | 3    | 0      | 32    | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| podOwnerShares                  | mapping(address => int256)             | 4    | 0      | 32    | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| __deprecated_denebForkTimestamp | uint64                                 | 5    | 0      | 8     | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
| __gap                           | uint256[44]                            | 6    | 0      | 1408  | src/contracts/pods/EigenPodManagerStorage.sol:EigenPodManagerStorage |
