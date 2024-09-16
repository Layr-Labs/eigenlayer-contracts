| Name                              | Type                                      | Slot | Offset | Bytes | Contract                                                         |
|-----------------------------------|-------------------------------------------|------|--------|-------|------------------------------------------------------------------|
| IS_TEST                           | bool                                      | 0    | 0      | 1     | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _failed                           | bool                                      | 0    | 1      | 1     | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| stdChainsInitialized              | bool                                      | 0    | 2      | 1     | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| chains                            | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| defaultRpcUrls                    | mapping(string => string)                 | 2    | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| idToAlias                         | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| fallbackToDefaultRpcUrls          | bool                                      | 4    | 0      | 1     | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| gasMeteringOff                    | bool                                      | 4    | 1      | 1     | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| stdstore                          | struct StdStorage                         | 5    | 0      | 224   | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _excludedContracts                | address[]                                 | 12   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _excludedSenders                  | address[]                                 | 13   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _targetedContracts                | address[]                                 | 14   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _targetedSenders                  | address[]                                 | 15   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _excludedArtifacts                | string[]                                  | 16   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _targetedArtifacts                | string[]                                  | 17   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _targetedArtifactSelectors        | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| _targetedSelectors                | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| stdstore                          | struct StdStorage                         | 20   | 0      | 224   | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| cheats                            | contract Vm                               | 27   | 0      | 20    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| pauserRegistry                    | contract PauserRegistry                   | 28   | 0      | 20    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| eigenLayerProxyAdmin              | contract ProxyAdmin                       | 29   | 0      | 20    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
| addressIsExcludedFromFuzzedInputs | mapping(address => bool)                  | 30   | 0      | 32    | src/test/utils/EigenLayerUnitTestBase.sol:EigenLayerUnitTestBase |
