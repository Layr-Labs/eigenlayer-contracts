| Name                       | Type                                      | Slot | Offset | Bytes | Contract                               |
|----------------------------|-------------------------------------------|------|--------|-------|----------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/mocks/Reenterer.sol:Reenterer |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/mocks/Reenterer.sol:Reenterer |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/mocks/Reenterer.sol:Reenterer |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/mocks/Reenterer.sol:Reenterer |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/mocks/Reenterer.sol:Reenterer |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/mocks/Reenterer.sol:Reenterer |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/mocks/Reenterer.sol:Reenterer |
| cheats                     | contract Vm                               | 27   | 0      | 20    | src/test/mocks/Reenterer.sol:Reenterer |
| target                     | address                                   | 28   | 0      | 20    | src/test/mocks/Reenterer.sol:Reenterer |
| msgValue                   | uint256                                   | 29   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| callData                   | bytes                                     | 30   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| expectedRevertData         | bytes                                     | 31   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
| dataToReturn               | bytes                                     | 32   | 0      | 32    | src/test/mocks/Reenterer.sol:Reenterer |
