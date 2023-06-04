# ERC20_OneWeiFeeOnTransfer
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol)

**Inherits:**
[OpenZeppelin_ERC20PresetFixedSupply](/docs/docgen/src/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/contract.OpenZeppelin_ERC20PresetFixedSupply.md)


## Functions
### constructor


```solidity
constructor(uint256 initSupply, address initOwner)
    OpenZeppelin_ERC20PresetFixedSupply(
        "ERC20_OneWeiFeeOnTransfer_Mock",
        "ERC20_OneWeiFeeOnTransfer_Mock",
        initSupply,
        initOwner
    );
```

### _transfer

*Moves `amount` of tokens from `from` to `to`.
This internal function is equivalent to {transfer}, and can be used to
e.g. implement automatic token fees, slashing mechanisms, etc.
Emits a {Transfer} event.
Requirements:
- `from` cannot be the zero address.
- `to` cannot be the zero address.
- `from` must have a balance of at least `amount`.*


```solidity
function _transfer(address from, address to, uint256 amount) internal virtual override;
```

