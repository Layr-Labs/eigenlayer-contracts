# Solidity API

## IStrategy

Custom `Strategy` implementations may expand extensively on this interface.

### deposit

```solidity
function deposit(contract IERC20 token, uint256 amount) external returns (uint256)
```

Used to deposit tokens into this Strategy

_This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
`depositIntoStrategy` function, and individual share balances are recorded in the strategyManager as well._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| token | contract IERC20 | is the ERC20 token being deposited |
| amount | uint256 | is the amount of token being deposited |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | newShares is the number of new shares issued at the current exchange ratio. |

### withdraw

```solidity
function withdraw(address depositor, contract IERC20 token, uint256 amountShares) external
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

### sharesToUnderlying

```solidity
function sharesToUnderlying(uint256 amountShares) external returns (uint256)
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

### underlyingToShares

```solidity
function underlyingToShares(uint256 amountUnderlying) external returns (uint256)
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
| [0] | uint256 | The amount of underlying tokens corresponding to the input `amountShares` |

### userUnderlying

```solidity
function userUnderlying(address user) external returns (uint256)
```

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlyingView`, this function **may** make state modifications

### sharesToUnderlyingView

```solidity
function sharesToUnderlyingView(uint256 amountShares) external view returns (uint256)
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
| [0] | uint256 | The amount of shares corresponding to the input `amountUnderlying` |

### underlyingToSharesView

```solidity
function underlyingToSharesView(uint256 amountUnderlying) external view returns (uint256)
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

### userUnderlyingView

```solidity
function userUnderlyingView(address user) external view returns (uint256)
```

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlying`, this function guarantees no state modifications

### underlyingToken

```solidity
function underlyingToken() external view returns (contract IERC20)
```

The underlying token for shares in this Strategy

### totalShares

```solidity
function totalShares() external view returns (uint256)
```

The total number of extant shares in this Strategy

### explanation

```solidity
function explanation() external view returns (string)
```

Returns either a brief string explaining the strategy's goal & purpose, or a link to metadata that explains in more detail.

