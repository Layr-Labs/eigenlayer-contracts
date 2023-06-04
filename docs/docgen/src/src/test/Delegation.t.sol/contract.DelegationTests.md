# DelegationTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/Delegation.t.sol)

**Inherits:**
[EigenLayerTestHelper](/docs/docgen/src/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### PRIVATE_KEY

```solidity
uint256 public PRIVATE_KEY = 420;
```


### serveUntil

```solidity
uint32 serveUntil = 100;
```


### serviceManager

```solidity
ServiceManagerMock public serviceManager;
```


### voteWeigher

```solidity
MiddlewareVoteWeigherMock public voteWeigher;
```


### voteWeigherImplementation

```solidity
MiddlewareVoteWeigherMock public voteWeigherImplementation;
```


## Functions
### fuzzedAmounts


```solidity
modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount);
```

### setUp


```solidity
function setUp() public virtual override;
```

### initializeMiddlewares


```solidity
function initializeMiddlewares() public;
```

### testSelfOperatorRegister

testing if an operator can register to themselves.


```solidity
function testSelfOperatorRegister() public;
```

### testSelfOperatorDelegate

testing if an operator can delegate to themselves.


```solidity
function testSelfOperatorDelegate(address sender) public;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|is the address of the operator.|


### testTwoSelfOperatorsRegister


```solidity
function testTwoSelfOperatorsRegister() public;
```

### testDelegation

registers a fixed address as a delegate, delegates to it from a second address,
and checks that the delegate's voteWeights increase properly


```solidity
function testDelegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker)
    fuzzedAmounts(ethAmount, eigenAmount);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`staker`|`address`|is the staker delegating stake to the operator.|
|`ethAmount`|`uint96`||
|`eigenAmount`|`uint96`||


### testDelegationReceived

tests that a when an operator is delegated to, that delegation is properly accounted for.


```solidity
function testDelegationReceived(address _operator, address staker, uint64 ethAmount, uint64 eigenAmount)
    public
    fuzzedAddress(_operator)
    fuzzedAddress(staker)
    fuzzedAmounts(ethAmount, eigenAmount);
```

### testUndelegation

tests that a when an operator is undelegated from, that the staker is properly classified as undelegated.


```solidity
function testUndelegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker)
    fuzzedAmounts(ethAmount, eigenAmount);
```

### testDelegateToBySignature

tests delegation from a staker to operator via ECDSA signature.


```solidity
function testDelegateToBySignature(address operator, uint96 ethAmount, uint96 eigenAmount, uint256 expiry)
    public
    fuzzedAddress(operator);
```

### testDelegateToBySignature_WithContractWallet_Successfully

tries delegating using a signature and an EIP 1271 compliant wallet


```solidity
function testDelegateToBySignature_WithContractWallet_Successfully(
    address operator,
    uint96 ethAmount,
    uint96 eigenAmount
) public fuzzedAddress(operator);
```

### testDelegateToBySignature_WithContractWallet_BadSignature

tries delegating using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature


```solidity
function testDelegateToBySignature_WithContractWallet_BadSignature(
    address operator,
    uint96 ethAmount,
    uint96 eigenAmount
) public fuzzedAddress(operator);
```

### testDelegateToBySignature_WithContractWallet_NonconformingWallet

tries delegating using a wallet that does not comply with EIP 1271


```solidity
function testDelegateToBySignature_WithContractWallet_NonconformingWallet(
    address operator,
    uint96 ethAmount,
    uint96 eigenAmount,
    uint8 v,
    bytes32 r,
    bytes32 s
) public fuzzedAddress(operator);
```

### testDelegateToByInvalidSignature

tests delegation to EigenLayer via an ECDSA signatures with invalid signature


```solidity
function testDelegateToByInvalidSignature(
    address operator,
    uint96 ethAmount,
    uint96 eigenAmount,
    uint8 v,
    bytes32 r,
    bytes32 s
) public fuzzedAddress(operator) fuzzedAmounts(ethAmount, eigenAmount);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|
|`ethAmount`|`uint96`||
|`eigenAmount`|`uint96`||
|`v`|`uint8`||
|`r`|`bytes32`||
|`s`|`bytes32`||


### testDelegationMultipleStrategies

registers a fixed address as a delegate, delegates to it from a second address,
and checks that the delegate's voteWeights increase properly


```solidity
function testDelegationMultipleStrategies(uint8 numStratsToAdd, address operator, address staker)
    public
    fuzzedAddress(operator)
    fuzzedAddress(staker);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`numStratsToAdd`|`uint8`||
|`operator`|`address`|is the operator being delegated to.|
|`staker`|`address`|is the staker delegating stake to the operator.|


### testCannotInitMultipleTimesDelegation

This function tests to ensure that a delegation contract
cannot be intitialized multiple times


```solidity
function testCannotInitMultipleTimesDelegation() public cannotReinit;
```

### testRegisterAsOperatorMultipleTimes

This function tests to ensure that a you can't register as a delegate multiple times


```solidity
function testRegisterAsOperatorMultipleTimes(address operator) public fuzzedAddress(operator);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator being delegated to.|


### testDelegationToUnregisteredDelegate

This function tests to ensure that a staker cannot delegate to an unregistered operator


```solidity
function testDelegationToUnregisteredDelegate(address delegate) public fuzzedAddress(delegate);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`delegate`|`address`|is the unregistered operator|


### testCannotInitMultipleTimesDelegation

This function tests to ensure that a delegation contract
cannot be intitialized multiple times, test with different caller addresses


```solidity
function testCannotInitMultipleTimesDelegation(address _attacker) public;
```

### testCannotSetDelegationTermsZeroAddress

This function tests that the delegationTerms cannot be set to address(0)


```solidity
function testCannotSetDelegationTermsZeroAddress() public;
```

### testCannotRegisterAsOperatorTwice

This function tests to ensure that an address can only call registerAsOperator() once


```solidity
function testCannotRegisterAsOperatorTwice(address _operator, address _dt)
    public
    fuzzedAddress(_operator)
    fuzzedAddress(_dt);
```

### testDelegateToInvalidOperator

This function checks that you can only delegate to an address that is already registered.


```solidity
function testDelegateToInvalidOperator(address _staker, address _unregisteredOperator) public fuzzedAddress(_staker);
```

### testUndelegate_SigP_Version


```solidity
function testUndelegate_SigP_Version(address _operator, address _staker, address _dt) public;
```

### _testRegisterAdditionalOperator


```solidity
function _testRegisterAdditionalOperator(uint256 index, uint32 _serveUntil) internal;
```

### _registerOperatorAndDepositFromStaker


```solidity
function _registerOperatorAndDepositFromStaker(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
    internal;
```

