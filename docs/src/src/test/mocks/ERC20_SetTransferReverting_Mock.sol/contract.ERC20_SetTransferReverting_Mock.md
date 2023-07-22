# ERC20_SetTransferReverting_Mock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/ERC20_SetTransferReverting_Mock.sol)

**Inherits:**
ERC20PresetFixedSupply


## State Variables
### transfersRevert

```solidity
bool public transfersRevert;
```


## Functions
### constructor


```solidity
constructor(uint256 initSupply, address initOwner)
    ERC20PresetFixedSupply("ERC20_SetTransferReverting_Mock", "ERC20_SetTransferReverting_Mock", initSupply, initOwner);
```

### setTransfersRevert


```solidity
function setTransfersRevert(bool _transfersRevert) public;
```

### _beforeTokenTransfer


```solidity
function _beforeTokenTransfer(address, address, uint256) internal view override;
```

