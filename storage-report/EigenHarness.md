| Name                              | Type                                                          | Slot | Offset | Bytes | Contract                                         |
|-----------------------------------|---------------------------------------------------------------|------|--------|-------|--------------------------------------------------|
| _initialized                      | uint8                                                         | 0    | 0      | 1     | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _initializing                     | bool                                                          | 0    | 1      | 1     | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[50]                                                   | 1    | 0      | 1600  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _owner                            | address                                                       | 51   | 0      | 20    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[49]                                                   | 52   | 0      | 1568  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _balances                         | mapping(address => uint256)                                   | 101  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _allowances                       | mapping(address => mapping(address => uint256))               | 102  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _totalSupply                      | uint256                                                       | 103  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _name                             | string                                                        | 104  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _symbol                           | string                                                        | 105  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[45]                                                   | 106  | 0      | 1440  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _hashedName                       | bytes32                                                       | 151  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _hashedVersion                    | bytes32                                                       | 152  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _name                             | string                                                        | 153  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _version                          | string                                                        | 154  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[48]                                                   | 155  | 0      | 1536  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _nonces                           | mapping(address => struct CountersUpgradeable.Counter)        | 203  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _PERMIT_TYPEHASH_DEPRECATED_SLOT  | bytes32                                                       | 204  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[49]                                                   | 205  | 0      | 1568  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _delegates                        | mapping(address => address)                                   | 254  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _checkpoints                      | mapping(address => struct ERC20VotesUpgradeable.Checkpoint[]) | 255  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| _totalSupplyCheckpoints           | struct ERC20VotesUpgradeable.Checkpoint[]                     | 256  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| __gap                             | uint256[47]                                                   | 257  | 0      | 1504  | src/test/harnesses/EigenHarness.sol:EigenHarness |
| mintAllowedAfter                  | mapping(address => uint256)                                   | 304  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| mintingAllowance                  | mapping(address => uint256)                                   | 305  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| transferRestrictionsDisabledAfter | uint256                                                       | 306  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| allowedFrom                       | mapping(address => bool)                                      | 307  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
| allowedTo                         | mapping(address => bool)                                      | 308  | 0      | 32    | src/test/harnesses/EigenHarness.sol:EigenHarness |
