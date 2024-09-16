| Name                                | Type                                     | Slot | Offset | Bytes | Contract                                                   |
|-------------------------------------|------------------------------------------|------|--------|-------|------------------------------------------------------------|
| _initialized                        | uint8                                    | 0    | 0      | 1     | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| _initializing                       | bool                                     | 0    | 1      | 1     | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| __gap                               | uint256[50]                              | 1    | 0      | 1600  | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| _owner                              | address                                  | 51   | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| __gap                               | uint256[49]                              | 52   | 0      | 1568  | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| _status                             | uint256                                  | 101  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| __gap                               | uint256[49]                              | 102  | 0      | 1568  | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| pauserRegistry                      | contract IPauserRegistry                 | 151  | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| _paused                             | uint256                                  | 152  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| __gap                               | uint256[48]                              | 153  | 0      | 1536  | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| delegation                          | contract IDelegationManager              | 201  | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| eigenPodManager                     | contract IEigenPodManager                | 202  | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| slasher                             | contract ISlasher                        | 203  | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| strategyWhitelister                 | address                                  | 204  | 0      | 20    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| strategiesToReturn                  | mapping(address => contract IStrategy[]) | 205  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| sharesToReturn                      | mapping(address => uint256[])            | 206  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| strategyIsWhitelistedForDeposit     | mapping(contract IStrategy => bool)      | 207  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| cumulativeWithdrawalsQueued         | mapping(address => uint256)              | 208  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
| stakerStrategyListLengthReturnValue | uint256                                  | 209  | 0      | 32    | src/test/mocks/StrategyManagerMock.sol:StrategyManagerMock |
