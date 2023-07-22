# DelegationUnitTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/unit/DelegationUnit.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### strategyManagerMock

```solidity
StrategyManagerMock strategyManagerMock;
```


### slasherMock

```solidity
SlasherMock slasherMock;
```


### delegationManager

```solidity
DelegationManager delegationManager;
```


### delegationTermsMock

```solidity
DelegationTermsMock delegationTermsMock;
```


### delegationManagerImplementation

```solidity
DelegationManager delegationManagerImplementation;
```


### strategyImplementation

```solidity
StrategyBase strategyImplementation;
```


### strategyMock

```solidity
StrategyBase strategyMock;
```


### GWEI_TO_WEI

```solidity
uint256 GWEI_TO_WEI = 1e9;
```


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testReinitializeDelegation


```solidity
function testReinitializeDelegation() public;
```

### testBadECDSASignatureExpiry


```solidity
function testBadECDSASignatureExpiry(address staker, address operator, uint256 expiry, bytes memory signature) public;
```

### testUndelegateFromNonStrategyManagerAddress


```solidity
function testUndelegateFromNonStrategyManagerAddress(address undelegator) public fuzzedAddress(undelegator);
```

### testUndelegateByOperatorFromThemselves


```solidity
function testUndelegateByOperatorFromThemselves(address operator) public fuzzedAddress(operator);
```

### testIncreaseDelegatedSharesFromNonStrategyManagerAddress


```solidity
function testIncreaseDelegatedSharesFromNonStrategyManagerAddress(address operator, uint256 shares)
    public
    fuzzedAddress(operator);
```

### testDecreaseDelegatedSharesFromNonStrategyManagerAddress


```solidity
function testDecreaseDelegatedSharesFromNonStrategyManagerAddress(
    address operator,
    IStrategy[] memory strategies,
    uint256[] memory shareAmounts
) public fuzzedAddress(operator);
```

### testDelegateWhenOperatorIsFrozen


```solidity
function testDelegateWhenOperatorIsFrozen(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testDelegateWhenStakerHasExistingDelegation


```solidity
function testDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2)
    public
    fuzzedAddress(staker)
    fuzzedAddress(operator)
    fuzzedAddress(operator2);
```

### testDelegationToUnregisteredOperator


```solidity
function testDelegationToUnregisteredOperator(address operator) public;
```

### testDelegationWhenPausedNewDelegationIsSet


```solidity
function testDelegationWhenPausedNewDelegationIsSet(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testRevertingDelegationReceivedHook


```solidity
function testRevertingDelegationReceivedHook(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testRevertingDelegationWithdrawnHook


```solidity
function testRevertingDelegationWithdrawnHook(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testDelegationReceivedHookWithTooMuchReturnData


```solidity
function testDelegationReceivedHookWithTooMuchReturnData(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testDelegationWithdrawnHookWithTooMuchReturnData


```solidity
function testDelegationWithdrawnHookWithTooMuchReturnData(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testDelegationReceivedHookWithNoReturnData


```solidity
function testDelegationReceivedHookWithNoReturnData(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

### testDelegationWithdrawnHookWithNoReturnData


```solidity
function testDelegationWithdrawnHookWithNoReturnData(address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```

## Events
### OnDelegationReceivedCallFailure

```solidity
event OnDelegationReceivedCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
```

### OnDelegationWithdrawnCallFailure

```solidity
event OnDelegationWithdrawnCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);
```

