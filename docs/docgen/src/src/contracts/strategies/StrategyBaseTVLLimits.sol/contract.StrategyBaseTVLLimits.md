# StrategyBaseTVLLimits
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/strategies/StrategyBaseTVLLimits.sol)

**Inherits:**
[StrategyBase](/docs/docgen/src/src/contracts/strategies/StrategyBase.sol/contract.StrategyBase.md)

**Author:**
Layr Labs, Inc.

*Note that this implementation still converts between any amount of shares or underlying tokens in its view functions;
these functions purposefully do not take the TVL limit into account.*


## State Variables
### maxPerDeposit
The maximum deposit (in underlyingToken) that this strategy will accept per deposit


```solidity
uint256 public maxPerDeposit;
```


### maxTotalDeposits
The maximum deposits (in underlyingToken) that this strategy will hold


```solidity
uint256 public maxTotalDeposits;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[48] private __gap;
```


## Functions
### constructor


```solidity
constructor(IStrategyManager _strategyManager) StrategyBase(_strategyManager);
```

### initialize


```solidity
function initialize(
    uint256 _maxPerDeposit,
    uint256 _maxTotalDeposits,
    IERC20 _underlyingToken,
    IPauserRegistry _pauserRegistry
) public virtual initializer;
```

### setTVLLimits

Sets the maximum deposits (in underlyingToken) that this strategy will hold and accept per deposit

*Callable only by the pauser of this contract*

*We note that there is a potential race condition between a call to this function that lowers either or both of these limits and call(s)
to `deposit`, that may result in some calls to `deposit` reverting.*


```solidity
function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) external onlyPauser;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newMaxPerDeposit`|`uint256`||
|`newMaxTotalDeposits`|`uint256`|The new maximum deposits|


### getTVLLimits

Simple getter function that returns the current values of `maxPerDeposit` and `maxTotalDeposits`.


```solidity
function getTVLLimits() external view returns (uint256, uint256);
```

### _setTVLLimits

Internal setter for TVL limits


```solidity
function _setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) internal;
```

### _beforeDeposit

Called in the external `deposit` function, before any logic is executed. Makes sure that deposits don't exceed configured maximum.

*Unused token param is the token being deposited. This is already checked in the `deposit` function.*

*Note that the `maxTotalDeposits` is purely checked against the current `_tokenBalance()`, since by this point in the deposit flow, the
tokens should have already been transferred to this Strategy by the StrategyManager*

*We note as well that this makes it possible for various race conditions to occur:
a) multiple simultaneous calls to `deposit` may result in some of these calls reverting due to `maxTotalDeposits` being reached.
b) transferring funds directly to this Strategy (although not generally in someone's economic self interest) in order to reach `maxTotalDeposits`
is a route by which someone can cause calls to `deposit` to revert.
c) increases in the token balance of this contract through other effects – including token rebasing – may cause similar issues to (a) and (b).*


```solidity
function _beforeDeposit(IERC20, uint256 amount) internal virtual override;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`IERC20`||
|`amount`|`uint256`|The amount of `token` being deposited|


## Events
### MaxPerDepositUpdated
Emitted when `maxPerDeposit` value is updated from `previousValue` to `newValue`


```solidity
event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue);
```

### MaxTotalDepositsUpdated
Emitted when `maxTotalDeposits` value is updated from `previousValue` to `newValue`


```solidity
event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue);
```

