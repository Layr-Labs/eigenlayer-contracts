# StrategyBaseTVLLimitsUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/StrategyBaseTVLLimitsUnit.sol)

**Inherits:**
[StrategyBaseUnitTests](/docs/docgen/src/src/test/unit/StrategyBaseUnit.t.sol/contract.StrategyBaseUnitTests.md)


## State Variables
### strategyBaseTVLLimitsImplementation

```solidity
StrategyBaseTVLLimits public strategyBaseTVLLimitsImplementation;
```


### strategyWithTVLLimits

```solidity
StrategyBaseTVLLimits public strategyWithTVLLimits;
```


### maxTotalDeposits

```solidity
uint256 maxTotalDeposits = 3200e18;
```


### maxPerDeposit

```solidity
uint256 maxPerDeposit = 32e18;
```


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testSetTVLLimits


```solidity
function testSetTVLLimits(uint256 maxPerDepositFuzzedInput, uint256 maxTotalDepositsFuzzedInput) public;
```

### testSetTVLLimitsFailsWhenNotCalledByPauser


```solidity
function testSetTVLLimitsFailsWhenNotCalledByPauser(
    uint256 maxPerDepositFuzzedInput,
    uint256 maxTotalDepositsFuzzedInput,
    address notPauser
) public;
```

### testSetInvalidMaxPerDepositAndMaxDeposits


```solidity
function testSetInvalidMaxPerDepositAndMaxDeposits(
    uint256 maxPerDepositFuzzedInput,
    uint256 maxTotalDepositsFuzzedInput
) public;
```

### testDepositMoreThanMaxPerDeposit


```solidity
function testDepositMoreThanMaxPerDeposit(
    uint256 maxPerDepositFuzzedInput,
    uint256 maxTotalDepositsFuzzedInput,
    uint256 amount
) public;
```

### testDepositMorethanMaxDeposits


```solidity
function testDepositMorethanMaxDeposits() public;
```

### testDepositValidAmount


```solidity
function testDepositValidAmount(uint256 depositAmount) public;
```

### testDepositTVLLimit_ThenChangeTVLLimit


```solidity
function testDepositTVLLimit_ThenChangeTVLLimit(
    uint256 maxTotalDepositsFuzzedInput,
    uint256 newMaxTotalDepositsFuzzedInput
) public;
```

### testDeposit_WithTVLLimits

General-purpose test, re-useable, handles whether the deposit should revert or not and returns 'true' if it did revert.


```solidity
function testDeposit_WithTVLLimits(
    uint256 maxPerDepositFuzzedInput,
    uint256 maxTotalDepositsFuzzedInput,
    uint256 depositAmount
) public returns (bool depositReverted);
```

### _setTVLLimits


```solidity
function _setTVLLimits(uint256 _maxPerDeposit, uint256 _maxTotalDeposits) internal;
```

### filterToValidDepositAmounts

OVERRIDING EXISTING TESTS TO FILTER INPUTS THAT WOULD FAIL DUE TO DEPOSIT-LIMITING


```solidity
modifier filterToValidDepositAmounts(uint256 amountToDeposit);
```

### testCanWithdrawDownToSmallShares


```solidity
function testCanWithdrawDownToSmallShares(uint256 amountToDeposit, uint32 sharesToLeave)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

### testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares


```solidity
function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint256 priorTotalShares, uint256 amountToDeposit)
    public
    virtual
    override
    filterToValidDepositAmounts(priorTotalShares)
    filterToValidDepositAmounts(amountToDeposit);
```

### testDepositWithZeroPriorBalanceAndZeroPriorShares


```solidity
function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint256 amountToDeposit)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

### testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares


```solidity
function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(
    uint256 amountToDeposit,
    uint256 amountToTransfer,
    uint96 amountSharesToQuery
) public virtual override filterToValidDepositAmounts(amountToDeposit);
```

### testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares


```solidity
function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(
    uint256 amountToDeposit,
    uint256 amountToTransfer,
    uint96 amountUnderlyingToQuery
) public virtual override filterToValidDepositAmounts(amountToDeposit);
```

### testWithdrawFailsWhenSharesGreaterThanTotalShares


```solidity
function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint256 amountToDeposit, uint256 sharesToWithdraw)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

### testWithdrawFailsWhenWithdrawalsPaused


```solidity
function testWithdrawFailsWhenWithdrawalsPaused(uint256 amountToDeposit)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

### testWithdrawWithPriorTotalSharesAndAmountSharesEqual


```solidity
function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint256 amountToDeposit)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

### testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual


```solidity
function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw)
    public
    virtual
    override
    filterToValidDepositAmounts(amountToDeposit);
```

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

