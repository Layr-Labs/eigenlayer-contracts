| Name                       | Type                                                       | Slot | Offset | Bytes | Contract                                                       |
|----------------------------|------------------------------------------------------------|------|--------|-------|----------------------------------------------------------------|
| IS_TEST                    | bool                                                       | 0    | 0      | 1     | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _failed                    | bool                                                       | 0    | 1      | 1     | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| stdChainsInitialized       | bool                                                       | 0    | 2      | 1     | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| chains                     | mapping(string => struct StdChains.Chain)                  | 1    | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| defaultRpcUrls             | mapping(string => string)                                  | 2    | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| idToAlias                  | mapping(uint256 => string)                                 | 3    | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| fallbackToDefaultRpcUrls   | bool                                                       | 4    | 0      | 1     | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| gasMeteringOff             | bool                                                       | 4    | 1      | 1     | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| stdstore                   | struct StdStorage                                          | 5    | 0      | 224   | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _excludedContracts         | address[]                                                  | 12   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _excludedSenders           | address[]                                                  | 13   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _targetedContracts         | address[]                                                  | 14   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _targetedSenders           | address[]                                                  | 15   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _excludedArtifacts         | string[]                                                   | 16   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _targetedArtifacts         | string[]                                                   | 17   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _targetedArtifactSelectors | struct StdInvariant.FuzzSelector[]                         | 18   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| _targetedSelectors         | struct StdInvariant.FuzzSelector[]                         | 19   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| stdstore                   | struct StdStorage                                          | 20   | 0      | 224   | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| isOperator                 | mapping(address => bool)                                   | 27   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| operatorShares             | mapping(address => mapping(contract IStrategy => uint256)) | 28   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
| delegatedTo                | mapping(address => address)                                | 29   | 0      | 32    | src/test/mocks/DelegationManagerMock.sol:DelegationManagerMock |
