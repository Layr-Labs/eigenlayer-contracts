# WhitelisterTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/Whitelister.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### dummyToken

```solidity
ERC20PresetMinterPauser dummyToken;
```


### dummyStrat

```solidity
IStrategy dummyStrat;
```


### dummyStratImplementation

```solidity
IStrategy dummyStratImplementation;
```


### whiteLister

```solidity
Whitelister whiteLister;
```


### blsRegistry

```solidity
BLSRegistry blsRegistry;
```


### blsRegistryImplementation

```solidity
BLSRegistry blsRegistryImplementation;
```


### dummyServiceManager

```solidity
ServiceManagerMock dummyServiceManager;
```


### dummyCompendium

```solidity
BLSPublicKeyCompendiumMock dummyCompendium;
```


### dummyReg

```solidity
MiddlewareRegistryMock dummyReg;
```


### AMOUNT

```solidity
uint256 AMOUNT;
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

### testWhitelistingOperator


```solidity
function testWhitelistingOperator(address operator) public fuzzedAddress(operator);
```

### testWhitelistDepositIntoStrategy


```solidity
function testWhitelistDepositIntoStrategy(address operator, uint256 depositAmount) external fuzzedAddress(operator);
```

### testCallStakerFromNonWhitelisterAddress


```solidity
function testCallStakerFromNonWhitelisterAddress(address nonWhitelister, bytes memory data)
    external
    fuzzedAddress(nonWhitelister);
```

### testNonWhitelistedOperatorRegistration


```solidity
function testNonWhitelistedOperatorRegistration(BN254.G1Point memory pk, string memory socket) external;
```

### testWhitelistQueueWithdrawal


```solidity
function testWhitelistQueueWithdrawal(address operator) public fuzzedAddress(operator);
```

### _testQueueWithdrawal


```solidity
function _testQueueWithdrawal(
    address staker,
    IStrategy[] memory strategyArray,
    uint256[] memory shareAmounts,
    uint256[] memory strategyIndexes
) internal;
```

### _testCompleteQueuedWithdrawal


```solidity
function _testCompleteQueuedWithdrawal(
    address staker,
    IStrategy[] memory strategyArray,
    IERC20[] memory tokensArray,
    uint256[] memory shareAmounts,
    address delegatedTo,
    IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce,
    uint32 withdrawalStartBlock,
    uint256 middlewareTimesIndex
) internal;
```

### testWhitelistTransfer


```solidity
function testWhitelistTransfer(address operator, address receiver) public fuzzedAddress(receiver);
```

## Structs
### DataForTestWithdrawal

```solidity
struct DataForTestWithdrawal {
    IStrategy[] delegatorStrategies;
    uint256[] delegatorShares;
    IStrategyManager.WithdrawerAndNonce withdrawerAndNonce;
}
```

