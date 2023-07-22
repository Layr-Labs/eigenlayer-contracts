# StrategyTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/Strategy.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## Functions
### testCannotInitMultipleTimesDelegation

This function tests to ensure that a delegation contract
cannot be intitialized multiple times


```solidity
function testCannotInitMultipleTimesDelegation() public cannotReinit;
```

### testInvalidCalltoDeposit

This function tests to ensure that only the strategyManager
can deposit into a strategy


```solidity
function testInvalidCalltoDeposit(address invalidDepositor) public fuzzedAddress(invalidDepositor);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`invalidDepositor`|`address`|is the non-registered depositor|


### testInvalidCalltoWithdraw

This function tests to ensure that only the strategyManager
can deposit into a strategy


```solidity
function testInvalidCalltoWithdraw(address depositor, address invalidWithdrawer)
    public
    fuzzedAddress(invalidWithdrawer);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|is the depositor for which the shares are being withdrawn|
|`invalidWithdrawer`|`address`|is the non-registered withdrawer|


### testWithdrawalExceedsTotalShares

This function tests ensures that withdrawing for a depositor that never
actually deposited fails.


```solidity
function testWithdrawalExceedsTotalShares(address depositor, uint256 shares) public fuzzedAddress(depositor);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|is the depositor for which the shares are being withdrawn|
|`shares`|`uint256`||


