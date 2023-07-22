# OpenZeppelin_ERC20Burnable
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol)

**Inherits:**
[OpenZeppelin_Context](/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/abstract.OpenZeppelin_Context.md), [OpenZeppelin_ERC20](/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/contract.OpenZeppelin_ERC20.md)

*Extension of {ERC20} that allows token holders to destroy both their own
tokens and those that they have an allowance for, in a way that can be
recognized off-chain (via event analysis).*


## Functions
### burn

*Destroys `amount` tokens from the caller.
See {ERC20-_burn}.*


```solidity
function burn(uint256 amount) public virtual;
```

### burnFrom

*Destroys `amount` tokens from `account`, deducting from the caller's
allowance.
See {ERC20-_burn} and {ERC20-allowance}.
Requirements:
- the caller must have allowance for ``accounts``'s tokens of at least
`amount`.*


```solidity
function burnFrom(address account, uint256 amount) public virtual;
```

