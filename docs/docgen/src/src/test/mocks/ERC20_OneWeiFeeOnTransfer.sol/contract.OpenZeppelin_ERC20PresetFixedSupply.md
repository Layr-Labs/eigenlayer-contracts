# OpenZeppelin_ERC20PresetFixedSupply
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol)

**Inherits:**
[OpenZeppelin_ERC20Burnable](/docs/docgen/src/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/abstract.OpenZeppelin_ERC20Burnable.md)

*{ERC20} token, including:
- Preminted initial supply
- Ability for holders to burn (destroy) their tokens
- No access control mechanism (for minting/pausing) and hence no governance
This contract uses {ERC20Burnable} to include burn capabilities - head to
its documentation for details.
_Available since v3.4._
_Deprecated in favor of https://wizard.openzeppelin.com/[Contracts Wizard]._*


## Functions
### constructor

*Mints `initialSupply` amount of token and transfers them to `owner`.
See {ERC20-constructor}.*


```solidity
constructor(string memory name, string memory symbol, uint256 initialSupply, address owner)
    OpenZeppelin_ERC20(name, symbol);
```

