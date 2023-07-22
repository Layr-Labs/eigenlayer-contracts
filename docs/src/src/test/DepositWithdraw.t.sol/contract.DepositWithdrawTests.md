# DepositWithdrawTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/DepositWithdraw.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### emptyUintArray

```solidity
uint256[] public emptyUintArray;
```


## Functions
### testWethDeposit

Verifies that it is possible to deposit WETH


```solidity
function testWethDeposit(uint256 amountToDeposit) public returns (uint256 amountDeposited);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amountToDeposit`|`uint256`|Fuzzed input for amount of WETH to deposit|


### testPreventSlashing


```solidity
function testPreventSlashing() public;
```

### testWithdrawalSequences


```solidity
function testWithdrawalSequences() public;
```

### testDepositStrategies

deploys 'numStratsToAdd' strategies using '_testAddStrategy' and then deposits '1e18' to each of them from 'getOperatorAddress(0)'


```solidity
function testDepositStrategies(uint8 numStratsToAdd) public;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`numStratsToAdd`|`uint8`|is the number of strategies being added and deposited into|


### testDepositEigen

Verifies that it is possible to deposit eigen.


```solidity
function testDepositEigen(uint96 eigenToDeposit) public;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`eigenToDeposit`|`uint96`|is amount of eigen to deposit into the eigen strategy|


### testDepositUnsupportedToken

Tries to deposit an unsupported token into an `StrategyBase` contract by calling `strategyManager.depositIntoStrategy`.
Verifies that reversion occurs correctly.


```solidity
function testDepositUnsupportedToken() public;
```

### testDepositNonexistentStrategy

Tries to deposit into an unsupported strategy by calling `strategyManager.depositIntoStrategy`.
Verifies that reversion occurs correctly.


```solidity
function testDepositNonexistentStrategy(address nonexistentStrategy) public fuzzedAddress(nonexistentStrategy);
```

### testRevertOnZeroDeposit

verify that trying to deposit an amount of zero will correctly revert


```solidity
function testRevertOnZeroDeposit() public;
```

### _createOnlyQueuedWithdrawal

Modified from existing _createQueuedWithdrawal, skips delegation and deposit steps so that we can isolate the withdrawal step

Creates a queued withdrawal from `staker`, queues a withdrawal using
`strategyManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawer)`

After initiating a queued withdrawal, this test checks that `strategyManager.canCompleteQueuedWithdrawal` immediately returns the correct
response depending on whether `staker` is delegated or not.


```solidity
function _createOnlyQueuedWithdrawal(
    address staker,
    bool,
    uint256 amountToDeposit,
    IStrategy[] memory strategyArray,
    IERC20[] memory,
    uint256[] memory shareAmounts,
    uint256[] memory strategyIndexes,
    address withdrawer
) internal returns (bytes32 withdrawalRoot, IStrategyManager.QueuedWithdrawal memory queuedWithdrawal);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|The address to initiate the queued withdrawal|
|`<none>`|`bool`||
|`amountToDeposit`|`uint256`|The amount of WETH to deposit|
|`strategyArray`|`IStrategy[]`||
|`<none>`|`IERC20[]`||
|`shareAmounts`|`uint256[]`||
|`strategyIndexes`|`uint256[]`||
|`withdrawer`|`address`||


### testFrontrunFirstDepositor


```solidity
function testFrontrunFirstDepositor() public;
```

### testFrontrunFirstDepositorFuzzed


```solidity
function testFrontrunFirstDepositorFuzzed(uint96 firstDepositAmount, uint96 donationAmount, uint96 secondDepositAmount)
    public;
```

### testDepositTokenWithOneWeiFeeOnTransfer


```solidity
function testDepositTokenWithOneWeiFeeOnTransfer(address sender, uint64 amountToDeposit) public fuzzedAddress(sender);
```

### testForkMainnetDepositSteth

Shadow-forks mainnet and tests depositing stETH tokens into a "StrategyBase" contract.


```solidity
function testForkMainnetDepositSteth() public;
```

### _whitelistStrategy

First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.


```solidity
function _whitelistStrategy(StrategyManager _strategyManager, StrategyBase _strategyBase)
    internal
    returns (StrategyManager);
```

