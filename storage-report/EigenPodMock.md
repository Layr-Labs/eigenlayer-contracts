| Name                       | Type                                      | Slot | Offset | Bytes | Contract                                     |
|----------------------------|-------------------------------------------|------|--------|-------|----------------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/mocks/EigenPodMock.sol:EigenPodMock |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/mocks/EigenPodMock.sol:EigenPodMock |
