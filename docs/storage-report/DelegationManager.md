| Name                          | Type                                                          | Slot | Offset | Bytes | Contract                                                   |
|-------------------------------|---------------------------------------------------------------|------|--------|-------|------------------------------------------------------------|
| _initialized                  | uint8                                                         | 0    | 0      | 1     | src/contracts/core/DelegationManager.sol:DelegationManager |
| _initializing                 | bool                                                          | 0    | 1      | 1     | src/contracts/core/DelegationManager.sol:DelegationManager |
| __gap                         | uint256[50]                                                   | 1    | 0      | 1600  | src/contracts/core/DelegationManager.sol:DelegationManager |
| _owner                        | address                                                       | 51   | 0      | 20    | src/contracts/core/DelegationManager.sol:DelegationManager |
| __gap                         | uint256[49]                                                   | 52   | 0      | 1568  | src/contracts/core/DelegationManager.sol:DelegationManager |
| pauserRegistry                | contract IPauserRegistry                                      | 101  | 0      | 20    | src/contracts/core/DelegationManager.sol:DelegationManager |
| _paused                       | uint256                                                       | 102  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| __gap                         | uint256[48]                                                   | 103  | 0      | 1536  | src/contracts/core/DelegationManager.sol:DelegationManager |
| _DOMAIN_SEPARATOR             | bytes32                                                       | 151  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| operatorShares                | mapping(address => mapping(contract IStrategy => uint256))    | 152  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| _operatorDetails              | mapping(address => struct IDelegationManager.OperatorDetails) | 153  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| delegatedTo                   | mapping(address => address)                                   | 154  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| stakerNonce                   | mapping(address => uint256)                                   | 155  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| delegationApproverSaltIsSpent | mapping(address => mapping(bytes32 => bool))                  | 156  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| minWithdrawalDelayBlocks      | uint256                                                       | 157  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| pendingWithdrawals            | mapping(bytes32 => bool)                                      | 158  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| cumulativeWithdrawalsQueued   | mapping(address => uint256)                                   | 159  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| __deprecated_stakeRegistry    | address                                                       | 160  | 0      | 20    | src/contracts/core/DelegationManager.sol:DelegationManager |
| strategyWithdrawalDelayBlocks | mapping(contract IStrategy => uint256)                        | 161  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| __gap                         | uint256[39]                                                   | 162  | 0      | 1248  | src/contracts/core/DelegationManager.sol:DelegationManager |
| _status                       | uint256                                                       | 201  | 0      | 32    | src/contracts/core/DelegationManager.sol:DelegationManager |
| __gap                         | uint256[49]                                                   | 202  | 0      | 1568  | src/contracts/core/DelegationManager.sol:DelegationManager |
