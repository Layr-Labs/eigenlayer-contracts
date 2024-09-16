| Name                                        | Type                                                       | Slot | Offset | Bytes | Contract                                               |
|---------------------------------------------|------------------------------------------------------------|------|--------|-------|--------------------------------------------------------|
| _initialized                                | uint8                                                      | 0    | 0      | 1     | src/contracts/core/StrategyManager.sol:StrategyManager |
| _initializing                               | bool                                                       | 0    | 1      | 1     | src/contracts/core/StrategyManager.sol:StrategyManager |
| __gap                                       | uint256[50]                                                | 1    | 0      | 1600  | src/contracts/core/StrategyManager.sol:StrategyManager |
| _owner                                      | address                                                    | 51   | 0      | 20    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __gap                                       | uint256[49]                                                | 52   | 0      | 1568  | src/contracts/core/StrategyManager.sol:StrategyManager |
| _status                                     | uint256                                                    | 101  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __gap                                       | uint256[49]                                                | 102  | 0      | 1568  | src/contracts/core/StrategyManager.sol:StrategyManager |
| pauserRegistry                              | contract IPauserRegistry                                   | 151  | 0      | 20    | src/contracts/core/StrategyManager.sol:StrategyManager |
| _paused                                     | uint256                                                    | 152  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __gap                                       | uint256[48]                                                | 153  | 0      | 1536  | src/contracts/core/StrategyManager.sol:StrategyManager |
| _DOMAIN_SEPARATOR                           | bytes32                                                    | 201  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| nonces                                      | mapping(address => uint256)                                | 202  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| strategyWhitelister                         | address                                                    | 203  | 0      | 20    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __deprecated_withdrawalDelayBlocks          | uint256                                                    | 204  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| stakerStrategyShares                        | mapping(address => mapping(contract IStrategy => uint256)) | 205  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| stakerStrategyList                          | mapping(address => contract IStrategy[])                   | 206  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __deprecated_withdrawalRootPending          | mapping(bytes32 => bool)                                   | 207  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __deprecated_numWithdrawalsQueued           | mapping(address => uint256)                                | 208  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| strategyIsWhitelistedForDeposit             | mapping(contract IStrategy => bool)                        | 209  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| beaconChainETHSharesToDecrementOnWithdrawal | mapping(address => uint256)                                | 210  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| thirdPartyTransfersForbidden                | mapping(contract IStrategy => bool)                        | 211  | 0      | 32    | src/contracts/core/StrategyManager.sol:StrategyManager |
| __gap                                       | uint256[39]                                                | 212  | 0      | 1248  | src/contracts/core/StrategyManager.sol:StrategyManager |
