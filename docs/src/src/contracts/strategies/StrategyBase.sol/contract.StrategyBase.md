# StrategyBase
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/strategies/StrategyBase.sol)

**Inherits:**
Initializable, [Pausable](/src/contracts/permissions/Pausable.sol/contract.Pausable.md), [IStrategy](/src/contracts/interfaces/IStrategy.sol/interface.IStrategy.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

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

*Note that some functions have their mutability restricted; developers inheriting from this contract cannot broaden
the mutability without modifying this contract itself.*

*This contract is expressly *not* intended for use with 'fee-on-transfer'-type tokens.
Setting the `underlyingToken` to be a fee-on-transfer token may result in improper accounting.*


## State Variables
### PAUSED_DEPOSITS

```solidity
uint8 internal constant PAUSED_DEPOSITS = 0;
```


### PAUSED_WITHDRAWALS

```solidity
uint8 internal constant PAUSED_WITHDRAWALS = 1;
```


### SHARES_OFFSET
virtual shares used as part of the mitigation of the common 'share inflation' attack vector.
Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
incurring reasonably small losses to depositors


```solidity
uint256 internal constant SHARES_OFFSET = 1e3;
```


### BALANCE_OFFSET
virtual balance used as part of the mitigation of the common 'share inflation' attack vector
Constant value chosen to reasonably reduce attempted share inflation by the first depositor, while still
incurring reasonably small losses to depositors


```solidity
uint256 internal constant BALANCE_OFFSET = 1e3;
```


### strategyManager
EigenLayer's StrategyManager contract


```solidity
IStrategyManager public immutable strategyManager;
```


### underlyingToken
The underlying token for shares in this Strategy


```solidity
IERC20 public underlyingToken;
```


### totalShares
The total number of extant shares in this Strategy


```solidity
uint256 public totalShares;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[48] private __gap;
```


## Functions
### onlyStrategyManager

Simply checks that the `msg.sender` is the `strategyManager`, which is an address stored immutably at construction.


```solidity
modifier onlyStrategyManager();
```

### constructor

Since this contract is designed to be initializable, the constructor simply sets `strategyManager`, the only immutable variable.


```solidity
constructor(IStrategyManager _strategyManager);
```

### initialize


```solidity
function initialize(IERC20 _underlyingToken, IPauserRegistry _pauserRegistry) public virtual initializer;
```

### _initializeStrategyBase

Sets the `underlyingToken` and `pauserRegistry` for the strategy.


```solidity
function _initializeStrategyBase(IERC20 _underlyingToken, IPauserRegistry _pauserRegistry) internal onlyInitializing;
```

### deposit

Used to deposit tokens into this Strategy

*This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
`depositIntoStrategy` function, and individual share balances are recorded in the strategyManager as well.*

*Note that the assumption is made that `amount` of `token` has already been transferred directly to this contract
(as performed in the StrategyManager's deposit functions). In particular, setting the `underlyingToken` of this contract
to be a fee-on-transfer token will break the assumption that the amount this contract *received* of the token is equal to
the amount that was input when the transfer was performed (i.e. the amount transferred 'out' of the depositor's balance).*


```solidity
function deposit(IERC20 token, uint256 amount)
    external
    virtual
    override
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyStrategyManager
    returns (uint256 newShares);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`IERC20`|is the ERC20 token being deposited|
|`amount`|`uint256`|is the amount of token being deposited|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`newShares`|`uint256`|is the number of new shares issued at the current exchange ratio.|


### withdraw

calculation of newShares *mirrors* `underlyingToShares(amount)`, but is different since the balance of `underlyingToken`
has already been increased due to the `strategyManager` transferring tokens to this strategy prior to calling this function

Used to withdraw tokens from this Strategy, to the `depositor`'s address

*This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
other functions, and individual share balances are recorded in the strategyManager as well.*


```solidity
function withdraw(address depositor, IERC20 token, uint256 amountShares)
    external
    virtual
    override
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyStrategyManager;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|is the address to receive the withdrawn funds|
|`token`|`IERC20`|is the ERC20 token being transferred out|
|`amountShares`|`uint256`|is the amount of shares being withdrawn|


### _beforeDeposit

calculation of amountToSend *mirrors* `sharesToUnderlying(amountShares)`, but is different since the `totalShares` has already
been decremented. Specifically, notice how we use `priorTotalShares` here instead of `totalShares`.

Called in the external `deposit` function, before any logic is executed. Expected to be overridden if strategies want such logic.


```solidity
function _beforeDeposit(IERC20 token, uint256 amount) internal virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`IERC20`|The token being deposited|
|`amount`|`uint256`|The amount of `token` being deposited|


### _beforeWithdrawal

Called in the external `withdraw` function, before any logic is executed.  Expected to be overridden if strategies want such logic.


```solidity
function _beforeWithdrawal(address depositor, IERC20 token, uint256 amountShares) internal virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The address that will receive the withdrawn tokens|
|`token`|`IERC20`|The token being withdrawn|
|`amountShares`|`uint256`|The amount of shares being withdrawn|


### explanation

Currently returns a brief string explaining the strategy's goal & purpose, but for more complex
strategies, may be a link to metadata that explains in more detail.


```solidity
function explanation() external pure virtual override returns (string memory);
```

### sharesToUnderlyingView

Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.

In contrast to `sharesToUnderlying`, this function guarantees no state modifications

*Implementation for these functions in particular may vary significantly for different strategies*


```solidity
function sharesToUnderlyingView(uint256 amountShares) public view virtual override returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amountShares`|`uint256`|is the amount of shares to calculate its conversion into the underlying token|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The amount of underlying tokens corresponding to the input `amountShares`|


### sharesToUnderlying

Used to convert a number of shares to the equivalent amount of underlying tokens for this strategy.

In contrast to `sharesToUnderlyingView`, this function **may** make state modifications

*Implementation for these functions in particular may vary significantly for different strategies*


```solidity
function sharesToUnderlying(uint256 amountShares) public view virtual override returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amountShares`|`uint256`|is the amount of shares to calculate its conversion into the underlying token|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The amount of underlying tokens corresponding to the input `amountShares`|


### underlyingToSharesView

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.

In contrast to `underlyingToShares`, this function guarantees no state modifications

*Implementation for these functions in particular may vary significantly for different strategies*


```solidity
function underlyingToSharesView(uint256 amountUnderlying) public view virtual returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amountUnderlying`|`uint256`|is the amount of `underlyingToken` to calculate its conversion into strategy shares|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The amount of shares corresponding to the input `amountUnderlying`|


### underlyingToShares

Used to convert an amount of underlying tokens to the equivalent amount of shares in this strategy.

In contrast to `underlyingToSharesView`, this function **may** make state modifications

*Implementation for these functions in particular may vary significantly for different strategies*


```solidity
function underlyingToShares(uint256 amountUnderlying) external view virtual returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amountUnderlying`|`uint256`|is the amount of `underlyingToken` to calculate its conversion into strategy shares|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The amount of shares corresponding to the input `amountUnderlying`|


### userUnderlyingView

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlying`, this function guarantees no state modifications


```solidity
function userUnderlyingView(address user) external view virtual returns (uint256);
```

### userUnderlying

convenience function for fetching the current underlying value of all of the `user`'s shares in
this strategy. In contrast to `userUnderlyingView`, this function **may** make state modifications


```solidity
function userUnderlying(address user) external virtual returns (uint256);
```

### shares

convenience function for fetching the current total shares of `user` in this strategy, by
querying the `strategyManager` contract


```solidity
function shares(address user) public view virtual returns (uint256);
```

### _tokenBalance

Internal function used to fetch this contract's current balance of `underlyingToken`.


```solidity
function _tokenBalance() internal view virtual returns (uint256);
```

