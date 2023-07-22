# StrategyBaseUnitTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/unit/StrategyBaseUnit.t.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### proxyAdmin

```solidity
ProxyAdmin public proxyAdmin;
```


### pauserRegistry

```solidity
PauserRegistry public pauserRegistry;
```


### strategyManager

```solidity
IStrategyManager public strategyManager;
```


### underlyingToken

```solidity
IERC20 public underlyingToken;
```


### strategyImplementation

```solidity
StrategyBase public strategyImplementation;
```


### strategy

```solidity
StrategyBase public strategy;
```


### pauser

```solidity
address public pauser = address(555);
```


### unpauser

```solidity
address public unpauser = address(999);
```


### initialSupply

```solidity
uint256 initialSupply = 1e36;
```


### initialOwner

```solidity
address initialOwner = address(this);
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


## Functions
### setUp


```solidity
function setUp() public virtual;
```

### testCannotReinitialize


```solidity
function testCannotReinitialize() public;
```

### testCannotReceiveZeroShares


```solidity
function testCannotReceiveZeroShares() public;
```

### testDepositWithZeroPriorBalanceAndZeroPriorShares


```solidity
function testDepositWithZeroPriorBalanceAndZeroPriorShares(uint256 amountToDeposit) public virtual;
```

### testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares


```solidity
function testDepositWithNonzeroPriorBalanceAndNonzeroPriorShares(uint256 priorTotalShares, uint256 amountToDeposit)
    public
    virtual;
```

### testDepositFailsWhenDepositsPaused


```solidity
function testDepositFailsWhenDepositsPaused() public;
```

### testDepositFailsWhenCallingFromNotStrategyManager


```solidity
function testDepositFailsWhenCallingFromNotStrategyManager(address caller) public;
```

### testDepositFailsWhenNotUsingUnderlyingToken


```solidity
function testDepositFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public;
```

### testWithdrawWithPriorTotalSharesAndAmountSharesEqual


```solidity
function testWithdrawWithPriorTotalSharesAndAmountSharesEqual(uint256 amountToDeposit) public virtual;
```

### testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual


```solidity
function testWithdrawWithPriorTotalSharesAndAmountSharesNotEqual(uint96 amountToDeposit, uint96 sharesToWithdraw)
    public
    virtual;
```

### testWithdrawZeroAmount


```solidity
function testWithdrawZeroAmount(uint256 amountToDeposit) public;
```

### testWithdrawFailsWhenWithdrawalsPaused


```solidity
function testWithdrawFailsWhenWithdrawalsPaused(uint256 amountToDeposit) public virtual;
```

### testWithdrawalFailsWhenCallingFromNotStrategyManager


```solidity
function testWithdrawalFailsWhenCallingFromNotStrategyManager(address caller) public;
```

### testWithdrawalFailsWhenNotUsingUnderlyingToken


```solidity
function testWithdrawalFailsWhenNotUsingUnderlyingToken(address notUnderlyingToken) public;
```

### testWithdrawFailsWhenSharesGreaterThanTotalShares


```solidity
function testWithdrawFailsWhenSharesGreaterThanTotalShares(uint256 amountToDeposit, uint256 sharesToWithdraw)
    public
    virtual;
```

### testWithdrawalFailsWhenTokenTransferFails


```solidity
function testWithdrawalFailsWhenTokenTransferFails() public;
```

### testIntegrityOfSharesToUnderlyingWithZeroTotalShares


```solidity
function testIntegrityOfSharesToUnderlyingWithZeroTotalShares(uint240 amountSharesToQuery) public view;
```

### testDeposit_ZeroAmount


```solidity
function testDeposit_ZeroAmount() public;
```

### testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares


```solidity
function testIntegrityOfSharesToUnderlyingWithNonzeroTotalShares(
    uint256 amountToDeposit,
    uint256 amountToTransfer,
    uint96 amountSharesToQuery
) public virtual;
```

### testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares


```solidity
function testIntegrityOfUnderlyingToSharesWithNonzeroTotalShares(
    uint256 amountToDeposit,
    uint256 amountToTransfer,
    uint96 amountUnderlyingToQuery
) public virtual;
```

### testCanWithdrawDownToSmallShares


```solidity
function testCanWithdrawDownToSmallShares(uint256 amountToDeposit, uint32 sharesToLeave) public virtual;
```

