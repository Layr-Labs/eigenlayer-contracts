# WithdrawalTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/Withdrawals.t.sol)

**Inherits:**
[DelegationTests](/src/test/Delegation.t.sol/contract.DelegationTests.md)


## State Variables
### generalReg1

```solidity
MiddlewareRegistryMock public generalReg1;
```


### generalServiceManager1

```solidity
ServiceManagerMock public generalServiceManager1;
```


### generalReg2

```solidity
MiddlewareRegistryMock public generalReg2;
```


### generalServiceManager2

```solidity
ServiceManagerMock public generalServiceManager2;
```


## Functions
### initializeGeneralMiddlewares


```solidity
function initializeGeneralMiddlewares() public;
```

### testWithdrawalWrapper


```solidity
function testWithdrawalWrapper(
    address operator,
    address depositor,
    address withdrawer,
    uint96 ethAmount,
    uint96 eigenAmount,
    bool withdrawAsTokens,
    bool RANDAO
) public fuzzedAddress(operator) fuzzedAddress(depositor) fuzzedAddress(withdrawer);
```

### _testWithdrawalAndDeregistration

test staker's ability to undelegate/withdraw from an operator.


```solidity
function _testWithdrawalAndDeregistration(
    address operator,
    address depositor,
    address withdrawer,
    uint96 ethAmount,
    uint96 eigenAmount,
    bool withdrawAsTokens
) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`depositor`|`address`|is the staker delegating stake to the operator.|
|`withdrawer`|`address`||
|`ethAmount`|`uint96`||
|`eigenAmount`|`uint96`||
|`withdrawAsTokens`|`bool`||


### _testWithdrawalWithStakeUpdate

test staker's ability to undelegate/withdraw from an operator.


```solidity
function _testWithdrawalWithStakeUpdate(
    address operator,
    address depositor,
    address withdrawer,
    uint96 ethAmount,
    uint96 eigenAmount,
    bool withdrawAsTokens
) public;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`depositor`|`address`|is the staker delegating stake to the operator.|
|`withdrawer`|`address`||
|`ethAmount`|`uint96`||
|`eigenAmount`|`uint96`||
|`withdrawAsTokens`|`bool`||


### testRedelegateAfterWithdrawal


```solidity
function testRedelegateAfterWithdrawal(
    address operator,
    address depositor,
    address withdrawer,
    uint96 ethAmount,
    uint96 eigenAmount,
    bool withdrawAsShares
) public fuzzedAddress(operator) fuzzedAddress(depositor) fuzzedAddress(withdrawer);
```

### testSlashedOperatorWithdrawal

test to see if an operator who is slashed/frozen
cannot be undelegated from by their stakers.


```solidity
function testSlashedOperatorWithdrawal(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`staker`|`address`|is the staker delegating stake to the operator.|
|`ethAmount`|`uint96`||
|`eigenAmount`|`uint96`||


## Structs
### DataForTestWithdrawal

```solidity
struct DataForTestWithdrawal {
    IStrategy[] delegatorStrategies;
    uint256[] delegatorShares;
    IStrategyManager.WithdrawerAndNonce withdrawerAndNonce;
}
```

