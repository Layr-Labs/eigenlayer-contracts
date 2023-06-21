# OpenZeppelin_ERC20
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol)

**Inherits:**
[OpenZeppelin_Context](/docs/docgen/src/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/abstract.OpenZeppelin_Context.md), [OpenZeppelin_IERC20](/docs/docgen/src/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/interface.OpenZeppelin_IERC20.md), [OpenZeppelin_IERC20Metadata](/docs/docgen/src/src/test/mocks/ERC20_OneWeiFeeOnTransfer.sol/interface.OpenZeppelin_IERC20Metadata.md)

*Implementation of the {IERC20} interface.
This implementation is agnostic to the way tokens are created. This means
that a supply mechanism has to be added in a derived contract using {_mint}.
For a generic mechanism see {ERC20PresetMinterPauser}.
TIP: For a detailed writeup see our guide
https://forum.openzeppelin.com/t/how-to-implement-erc20-supply-mechanisms/226[How
to implement supply mechanisms].
The default value of {decimals} is 18. To change this, you should override
this function so it returns a different value.
We have followed general OpenZeppelin Contracts guidelines: functions revert
instead returning `false` on failure. This behavior is nonetheless
conventional and does not conflict with the expectations of ERC20
applications.
Additionally, an {Approval} event is emitted on calls to {transferFrom}.
This allows applications to reconstruct the allowance for all accounts just
by listening to said events. Other implementations of the EIP may not emit
these events, as it isn't required by the specification.
Finally, the non-standard {decreaseAllowance} and {increaseAllowance}
functions have been added to mitigate the well-known issues around setting
allowances. See {IERC20-approve}.*


## State Variables
### _balances

```solidity
mapping(address => uint256) internal _balances;
```


### _allowances

```solidity
mapping(address => mapping(address => uint256)) private _allowances;
```


### _totalSupply

```solidity
uint256 private _totalSupply;
```


### _name

```solidity
string private _name;
```


### _symbol

```solidity
string private _symbol;
```


## Functions
### constructor

*Sets the values for {name} and {symbol}.
All two of these values are immutable: they can only be set once during
construction.*


```solidity
constructor(string memory name_, string memory symbol_);
```

### name

*Returns the name of the token.*


```solidity
function name() public view virtual override returns (string memory);
```

### symbol

*Returns the symbol of the token, usually a shorter version of the
name.*


```solidity
function symbol() public view virtual override returns (string memory);
```

### decimals

*Returns the number of decimals used to get its user representation.
For example, if `decimals` equals `2`, a balance of `505` tokens should
be displayed to a user as `5.05` (`505 / 10 ** 2`).
Tokens usually opt for a value of 18, imitating the relationship between
Ether and Wei. This is the default value returned by this function, unless
it's overridden.
NOTE: This information is only used for _display_ purposes: it in
no way affects any of the arithmetic of the contract, including
{IERC20-balanceOf} and {IERC20-transfer}.*


```solidity
function decimals() public view virtual override returns (uint8);
```

### totalSupply

*See {IERC20-totalSupply}.*


```solidity
function totalSupply() public view virtual override returns (uint256);
```

### balanceOf

*See {IERC20-balanceOf}.*


```solidity
function balanceOf(address account) public view virtual override returns (uint256);
```

### transfer

*See {IERC20-transfer}.
Requirements:
- `to` cannot be the zero address.
- the caller must have a balance of at least `amount`.*


```solidity
function transfer(address to, uint256 amount) public virtual override returns (bool);
```

### allowance

*See {IERC20-allowance}.*


```solidity
function allowance(address owner, address spender) public view virtual override returns (uint256);
```

### approve

*See {IERC20-approve}.
NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
`transferFrom`. This is semantically equivalent to an infinite approval.
Requirements:
- `spender` cannot be the zero address.*


```solidity
function approve(address spender, uint256 amount) public virtual override returns (bool);
```

### transferFrom

*See {IERC20-transferFrom}.
Emits an {Approval} event indicating the updated allowance. This is not
required by the EIP. See the note at the beginning of {ERC20}.
NOTE: Does not update the allowance if the current allowance
is the maximum `uint256`.
Requirements:
- `from` and `to` cannot be the zero address.
- `from` must have a balance of at least `amount`.
- the caller must have allowance for ``from``'s tokens of at least
`amount`.*


```solidity
function transferFrom(address from, address to, uint256 amount) public virtual override returns (bool);
```

### increaseAllowance

*Atomically increases the allowance granted to `spender` by the caller.
This is an alternative to {approve} that can be used as a mitigation for
problems described in {IERC20-approve}.
Emits an {Approval} event indicating the updated allowance.
Requirements:
- `spender` cannot be the zero address.*


```solidity
function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool);
```

### decreaseAllowance

*Atomically decreases the allowance granted to `spender` by the caller.
This is an alternative to {approve} that can be used as a mitigation for
problems described in {IERC20-approve}.
Emits an {Approval} event indicating the updated allowance.
Requirements:
- `spender` cannot be the zero address.
- `spender` must have allowance for the caller of at least
`subtractedValue`.*


```solidity
function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool);
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
function _transfer(address from, address to, uint256 amount) internal virtual;
```

### _mint

*Creates `amount` tokens and assigns them to `account`, increasing
the total supply.
Emits a {Transfer} event with `from` set to the zero address.
Requirements:
- `account` cannot be the zero address.*


```solidity
function _mint(address account, uint256 amount) internal virtual;
```

### _burn

*Destroys `amount` tokens from `account`, reducing the
total supply.
Emits a {Transfer} event with `to` set to the zero address.
Requirements:
- `account` cannot be the zero address.
- `account` must have at least `amount` tokens.*


```solidity
function _burn(address account, uint256 amount) internal virtual;
```

### _approve

*Sets `amount` as the allowance of `spender` over the `owner` s tokens.
This internal function is equivalent to `approve`, and can be used to
e.g. set automatic allowances for certain subsystems, etc.
Emits an {Approval} event.
Requirements:
- `owner` cannot be the zero address.
- `spender` cannot be the zero address.*


```solidity
function _approve(address owner, address spender, uint256 amount) internal virtual;
```

### _spendAllowance

*Updates `owner` s allowance for `spender` based on spent `amount`.
Does not update the allowance amount in case of infinite allowance.
Revert if not enough allowance is available.
Might emit an {Approval} event.*


```solidity
function _spendAllowance(address owner, address spender, uint256 amount) internal virtual;
```

### _beforeTokenTransfer

*Hook that is called before any transfer of tokens. This includes
minting and burning.
Calling conditions:
- when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
will be transferred to `to`.
- when `from` is zero, `amount` tokens will be minted for `to`.
- when `to` is zero, `amount` of ``from``'s tokens will be burned.
- `from` and `to` are never both zero.
To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].*


```solidity
function _beforeTokenTransfer(address from, address to, uint256 amount) internal virtual;
```

### _afterTokenTransfer

*Hook that is called after any transfer of tokens. This includes
minting and burning.
Calling conditions:
- when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
has been transferred to `to`.
- when `from` is zero, `amount` tokens have been minted for `to`.
- when `to` is zero, `amount` of ``from``'s tokens have been burned.
- `from` and `to` are never both zero.
To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].*


```solidity
function _afterTokenTransfer(address from, address to, uint256 amount) internal virtual;
```

