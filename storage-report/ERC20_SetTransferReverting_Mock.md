| Name            | Type                                            | Slot | Offset | Bytes | Contract                                                                           |
|-----------------|-------------------------------------------------|------|--------|-------|------------------------------------------------------------------------------------|
| _balances       | mapping(address => uint256)                     | 0    | 0      | 32    | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
| _allowances     | mapping(address => mapping(address => uint256)) | 1    | 0      | 32    | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
| _totalSupply    | uint256                                         | 2    | 0      | 32    | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
| _name           | string                                          | 3    | 0      | 32    | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
| _symbol         | string                                          | 4    | 0      | 32    | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
| transfersRevert | bool                                            | 5    | 0      | 1     | src/test/mocks/ERC20_SetTransferReverting_Mock.sol:ERC20_SetTransferReverting_Mock |
