# OpenZeppelin_ERC20PresetFixedSupply
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol)

**Inherits:**
[OpenZeppelin_ERC20Burnable](/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/abstract.OpenZeppelin_ERC20Burnable.md)

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

