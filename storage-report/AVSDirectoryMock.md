| Name                       | Type                                      | Slot | Offset | Bytes | Contract                                             |
|----------------------------|-------------------------------------------|------|--------|-------|------------------------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/mocks/AVSDirectoryMock.sol:AVSDirectoryMock |
