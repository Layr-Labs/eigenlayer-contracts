# ERC20_SetTransferReverting_Mock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/ERC20_SetTransferReverting_Mock.sol)

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

