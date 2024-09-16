| Name                       | Type                                      | Slot | Offset | Bytes | Contract                         |
|----------------------------|-------------------------------------------|------|--------|-------|----------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/utils/Owners.sol:Owners |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/utils/Owners.sol:Owners |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/utils/Owners.sol:Owners |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/utils/Owners.sol:Owners |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/utils/Owners.sol:Owners |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/utils/Owners.sol:Owners |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/utils/Owners.sol:Owners |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/utils/Owners.sol:Owners |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/utils/Owners.sol:Owners |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/utils/Owners.sol:Owners |
| ownersConfigJson           | string                                    | 27   | 0      | 32    | src/test/utils/Owners.sol:Owners |
| addresses                  | address[]                                 | 28   | 0      | 32    | src/test/utils/Owners.sol:Owners |
