| Name                       | Type                                      | Slot | Offset | Bytes | Contract                               |
|----------------------------|-------------------------------------------|------|--------|-------|----------------------------------------|
| IS_TEST                    | bool                                      | 0    | 0      | 1     | src/test/utils/Operators.sol:Operators |
| _failed                    | bool                                      | 0    | 1      | 1     | src/test/utils/Operators.sol:Operators |
| stdChainsInitialized       | bool                                      | 0    | 2      | 1     | src/test/utils/Operators.sol:Operators |
| chains                     | mapping(string => struct StdChains.Chain) | 1    | 0      | 32    | src/test/utils/Operators.sol:Operators |
| defaultRpcUrls             | mapping(string => string)                 | 2    | 0      | 32    | src/test/utils/Operators.sol:Operators |
| idToAlias                  | mapping(uint256 => string)                | 3    | 0      | 32    | src/test/utils/Operators.sol:Operators |
| fallbackToDefaultRpcUrls   | bool                                      | 4    | 0      | 1     | src/test/utils/Operators.sol:Operators |
| gasMeteringOff             | bool                                      | 4    | 1      | 1     | src/test/utils/Operators.sol:Operators |
| stdstore                   | struct StdStorage                         | 5    | 0      | 224   | src/test/utils/Operators.sol:Operators |
| _excludedContracts         | address[]                                 | 12   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _excludedSenders           | address[]                                 | 13   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _targetedContracts         | address[]                                 | 14   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _targetedSenders           | address[]                                 | 15   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _excludedArtifacts         | string[]                                  | 16   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _targetedArtifacts         | string[]                                  | 17   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]        | 18   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]        | 19   | 0      | 32    | src/test/utils/Operators.sol:Operators |
| stdstore                   | struct StdStorage                         | 20   | 0      | 224   | src/test/utils/Operators.sol:Operators |
| operatorConfigJson         | string                                    | 27   | 0      | 32    | src/test/utils/Operators.sol:Operators |
