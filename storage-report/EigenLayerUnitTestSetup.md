| Name                              | Type                                      | Slot | Offset | Bytes | Contract                                                           |
|-----------------------------------|-------------------------------------------|------|--------|-------|--------------------------------------------------------------------|
| IS_TEST                           | bool                                      | 0    | 0      | 1     | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _failed                           | bool                                      | 0    | 1      | 1     | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| stdChainsInitialized              | bool                                      | 0    | 2      | 1     | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| chains                            | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| defaultRpcUrls                    | mapping(string => string)                 | 2    | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| idToAlias                         | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| fallbackToDefaultRpcUrls          | bool                                      | 4    | 0      | 1     | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| gasMeteringOff                    | bool                                      | 4    | 1      | 1     | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| stdstore                          | struct StdStorage                         | 5    | 0      | 224   | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _excludedContracts                | address[]                                 | 12   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _excludedSenders                  | address[]                                 | 13   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _targetedContracts                | address[]                                 | 14   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _targetedSenders                  | address[]                                 | 15   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _excludedArtifacts                | string[]                                  | 16   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _targetedArtifacts                | string[]                                  | 17   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _targetedArtifactSelectors        | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| _targetedSelectors                | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| stdstore                          | struct StdStorage                         | 20   | 0      | 224   | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| cheats                            | contract Vm                               | 27   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| pauserRegistry                    | contract PauserRegistry                   | 28   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| eigenLayerProxyAdmin              | contract ProxyAdmin                       | 29   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| addressIsExcludedFromFuzzedInputs | mapping(address => bool)                  | 30   | 0      | 32    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| strategyManagerMock               | contract StrategyManagerMock              | 31   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| delegationManagerMock             | contract DelegationManagerMock            | 32   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| slasherMock                       | contract SlasherMock                      | 33   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| eigenPodManagerMock               | contract EigenPodManagerMock              | 34   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
| avsDirectoryMock                  | contract AVSDirectoryMock                 | 35   | 0      | 20    | src/test/utils/EigenLayerUnitTestSetup.sol:EigenLayerUnitTestSetup |
