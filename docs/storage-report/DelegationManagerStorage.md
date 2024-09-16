| Name                          | Type                                                          | Slot | Offset | Bytes | Contract                                                                 |
|-------------------------------|---------------------------------------------------------------|------|--------|-------|--------------------------------------------------------------------------|
| _DOMAIN_SEPARATOR             | bytes32                                                       | 0    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| operatorShares                | mapping(address => mapping(contract IStrategy => uint256))    | 1    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| _operatorDetails              | mapping(address => struct IDelegationManager.OperatorDetails) | 2    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| delegatedTo                   | mapping(address => address)                                   | 3    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| stakerNonce                   | mapping(address => uint256)                                   | 4    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| delegationApproverSaltIsSpent | mapping(address => mapping(bytes32 => bool))                  | 5    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| minWithdrawalDelayBlocks      | uint256                                                       | 6    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| pendingWithdrawals            | mapping(bytes32 => bool)                                      | 7    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| cumulativeWithdrawalsQueued   | mapping(address => uint256)                                   | 8    | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| __deprecated_stakeRegistry    | address                                                       | 9    | 0      | 20    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| strategyWithdrawalDelayBlocks | mapping(contract IStrategy => uint256)                        | 10   | 0      | 32    | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
| __gap                         | uint256[39]                                                   | 11   | 0      | 1248  | src/contracts/core/DelegationManagerStorage.sol:DelegationManagerStorage |
