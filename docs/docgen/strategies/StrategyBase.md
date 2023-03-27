# Solidity API

## StrategyBase

Simple, basic, "do-nothing" Strategy that holds a single underlying token and returns it on withdrawals.
Implements minimal versions of the IStrategy functions, this contract is designed to be inherited by
more complex strategies, which can then override its functions as necessary.

_This contract is expressly *not* intended for use with 'fee-on-transfer'-type tokens.
Setting the `underlyingToken` to be a fee-on-transfer token may result in improper accounting._

### PAUSED_DEPOSITS

```solidity
uint8 PAUSED_DEPOSITS
```

### PAUSED_WITHDRAWALS

```solidity
uint8 PAUSED_WITHDRAWALS
```

### MIN_NONZERO_TOTAL_SHARES

```solidity
uint96 MIN_NONZERO_TOTAL_SHARES
```

### strategyManager

```solidity
contract IStrategyManager strategyManager
```

EigenLayer's StrategyManager contract

### underlyingToken

```solidity
contract IERC20 underlyingToken
```

The underyling token for shares in this Strategy

### totalShares

```solidity
uint256 totalShares
```

The total number of extant shares in thie Strategy

### onlyStrategyManager

```solidity
modifier onlyStrategyManager()
```

Simply checks that the `msg.sender` is the `strategyManager`, which is an address stored immutably at construction.

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager) public
```

Since this contract is designed to be initializable, the constructor simply sets `strategyManager`, the only immutable variable.

### initialize

```solidity
function initialize(contract IERC20 _underlyingToken, contract IPauserRegistry _pauserRegistry) public
```

Sets the `underlyingToken` and `pauserRegistry` for the strategy.

### deposit

```solidity
function deposit(contract IERC20 token, uint256 amount) external virtual returns (uint256 newShares)
```

Used to deposit tokens into this Strategy

_This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
`depositIntoStrategy` function, and individual share balances are recorded in the strategyManager as well.
Note that the assumption is made that `amount` of `token` has already been transferred directly to this contract
(as performed in the StrategyManager's deposit functions). In particular, setting the `underlyingToken` of this contract
to be a fee-on-transfer token will break the assumption that the amount this contract *received* of the token is equal to
the amount that was input when the transfer was performed (i.e. the amount transferred 'out' of the depositor's balance)._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| token | contract IERC20 | is the ERC20 token being deposited |
| amount | uint256 | is the amount of token being deposited |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| newShares | uint256 | is the number of new shares issued at the current exchange ratio. |

### withdraw

```solidity
function withdraw(address depositor, contract IERC20 token, uint256 amountShares) external virtual
```

Used to withdraw tokens from this Strategy, to the `depositor`'s address

_This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
other functions, and individual share balances are recorded in the strategyManager as well._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| depositor | address |  |
| token | contract IERC20 | is the ERC20 token being transferred out |
| amountShares | uint256 | is the amount of shares being withdrawn |

### explanation

```solidity
function explanation() external pure virtual returns (string)
```

Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
strategies, may be a link to metadata that explains in more detail.

### sharesToUnderlyingView

```solidity
function sharesToUnderlyingView(uint256 amountShares) public view virtual returns (uint256)
```

Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
In contrast to `sharesToUnderlying`, this function guarantees no state modifications

_Implementation for these functions in particular may vary signifcantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountShares | uint256 | is the amount of shares to calculate its conversion into the underlying token |

### sharesToUnderlying

```solidity
function sharesToUnderlying(uint256 amountShares) public view virtual returns (uint256)
```

Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
In contrast to `sharesToUnderlyingView`, this function **may** make state modifications

_Implementation for these functions in particular may vary signifcantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountShares | uint256 | is the amount of shares to calculate its conversion into the underlying token |

### underlyingToSharesView

```solidity
function underlyingToSharesView(uint256 amountUnderlying) public view virtual returns (uint256)
```

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
In contrast to `underlyingToShares`, this function guarantees no state modifications

_Implementation for these functions in particular may vary signifcantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountUnderlying | uint256 | is the amount of `underlyingToken` to calculate its conversion into strategy shares |

### underlyingToShares

```solidity
function underlyingToShares(uint256 amountUnderlying) external view virtual returns (uint256)
```

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
In contrast to `underlyingToSharesView`, this function **may** make state modifications

_Implementation for these functions in particular may vary signifcantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountUnderlying | uint256 | is the amount of `underlyingToken` to calculate its conversion into strategy shares |

### userUnderlyingView

```solidity
function userUnderlyingView(address user) external view virtual returns (uint256)
```

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlying`, this function guarantees no state modifications

### userUnderlying

```solidity
function userUnderlying(address user) external virtual returns (uint256)
```

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlyingView`, this function **may** make state modifications

### shares

```solidity
function shares(address user) public view virtual returns (uint256)
```

convenience function for fetching the current total shares of `user` in this strategy, by
querying the `strategyManager` contract

### _tokenBalance

```solidity
function _tokenBalance() internal view virtual returns (uint256)
```

Internal function used to fetch this contract's current balance of `underlyingToken`.

### __gap

```solidity
uint256[48] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

