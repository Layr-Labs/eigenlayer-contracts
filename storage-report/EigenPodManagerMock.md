| Name                       | Type                                      | Slot | Offset | Bytes | Contract                                                   |
|----------------------------|-------------------------------------------|------|--------|-------|------------------------------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| pauserRegistry             | contract IPauserRegistry                  | 27   | 0      | 20    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| _paused                    | uint256                                   | 28   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| __gap                      | uint256[48]                               | 29   | 0      | 1536  | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| eigenPodBeacon             | contract IBeacon                          | 77   | 0      | 20    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| ethPOS                     | contract IETHPOSDeposit                   | 78   | 0      | 20    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
| podShares                  | mapping(address => int256)                | 79   | 0      | 32    | src/test/mocks/EigenPodManagerMock.sol:EigenPodManagerMock |
