# Solidity API

## StrategyBase

Simple, basic, "do-nothing" Strategy that holds a single underlying token and returns it on withdrawals.
Implements minimal versions of the IStrategy functions, this contract is designed to be inherited by
more complex strategies, which can then override its functions as necessary.
This contract functions similarly to an ERC4626 vault, only without issuing a token.
To mitigate against the common "inflation attack" vector, we have chosen to use the 'virtual shares' mitigation route,
similar to [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/ERC4626.sol).
We acknowledge that this mitigation has the known downside of the virtual shares causing some losses to users, which are pronounced
particularly in the case of the share exchange rate changing signficantly, either positively or negatively.
For a fairly thorough discussion of this issue and our chosen mitigation strategy, we recommend reading through
[this thread](https://github.com/OpenZeppelin/openzeppelin-contracts/issues/3706) on the OpenZeppelin repo.
We specifically use a share offset of `SHARES_OFFSET` and a balance offset of `BALANCE_OFFSET`.

_Note that some functions have their mutability restricted; developers inheriting from this contract cannot broaden
the mutability without modifying this contract itself.
This contract is expressly *not* intended for use with 'fee-on-transfer'-type tokens.
Setting the `underlyingToken` to be a fee-on-transfer token may result in improper accounting._

### PAUSED_DEPOSITS

```solidity
uint8 PAUSED_DEPOSITS
```

### PAUSED_WITHDRAWALS

```solidity
uint8 PAUSED_WITHDRAWALS
```

### SHARES_OFFSET

```solidity
uint256 SHARES_OFFSET
```

virtual shares used as part of the mitigation of the common 'share inflation' attack vector.
Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
incurring reasonably small losses to depositors

### BALANCE_OFFSET

```solidity
uint256 BALANCE_OFFSET
```

virtual balance used as part of the mitigation of the common 'share inflation' attack vector
Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
incurring reasonably small losses to depositors

### strategyManager

```solidity
contract IStrategyManager strategyManager
```

EigenLayer's StrategyManager contract

### underlyingToken

```solidity
contract IERC20 underlyingToken
```

The underlying token for shares in this Strategy

### totalShares

```solidity
uint256 totalShares
```

The total number of extant shares in this Strategy

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
function initialize(contract IERC20 _underlyingToken, contract IPauserRegistry _pauserRegistry) public virtual
```

### _initializeStrategyBase

```solidity
function _initializeStrategyBase(contract IERC20 _underlyingToken, contract IPauserRegistry _pauserRegistry) internal
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
| depositor | address | is the address to receive the withdrawn funds |
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

_Implementation for these functions in particular may vary significantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountShares | uint256 | is the amount of shares to calculate its conversion into the underlying token |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | The amount of underlying tokens corresponding to the input `amountShares` |

### sharesToUnderlying

```solidity
function sharesToUnderlying(uint256 amountShares) public view virtual returns (uint256)
```

Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.
In contrast to `sharesToUnderlyingView`, this function **may** make state modifications

_Implementation for these functions in particular may vary significantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountShares | uint256 | is the amount of shares to calculate its conversion into the underlying token |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | The amount of underlying tokens corresponding to the input `amountShares` |

### underlyingToSharesView

```solidity
function underlyingToSharesView(uint256 amountUnderlying) public view virtual returns (uint256)
```

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
In contrast to `underlyingToShares`, this function guarantees no state modifications

_Implementation for these functions in particular may vary significantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountUnderlying | uint256 | is the amount of `underlyingToken` to calculate its conversion into strategy shares |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | The amount of shares corresponding to the input `amountUnderlying` |

### underlyingToShares

```solidity
function underlyingToShares(uint256 amountUnderlying) external view virtual returns (uint256)
```

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.
In contrast to `underlyingToSharesView`, this function **may** make state modifications

_Implementation for these functions in particular may vary significantly for different strategies_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amountUnderlying | uint256 | is the amount of `underlyingToken` to calculate its conversion into strategy shares |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | The amount of shares corresponding to the input `amountUnderlying` |

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

