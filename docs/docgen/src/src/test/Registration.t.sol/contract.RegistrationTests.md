# RegistrationTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/Registration.t.sol)

**Inherits:**
[EigenLayerTestHelper](/docs/docgen/src/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### dlRegImplementation

```solidity
BLSRegistry public dlRegImplementation;
```


### pubkeyCompendium

```solidity
BLSPublicKeyCompendiumMock public pubkeyCompendium;
```


### pubkeyCompendiumImplementation

```solidity
BLSPublicKeyCompendiumMock public pubkeyCompendiumImplementation;
```


### dlReg

```solidity
BLSRegistry public dlReg;
```


### dataLayrProxyAdmin

```solidity
ProxyAdmin public dataLayrProxyAdmin;
```


### dlsm

```solidity
ServiceManagerMock public dlsm;
```


### strategyManagerMock

```solidity
StrategyManagerMock public strategyManagerMock;
```


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### initializeMiddlewares


```solidity
function initializeMiddlewares() public;
```

### testRegisterOperator


```solidity
function testRegisterOperator(address operator, uint32 operatorIndex, string calldata socket)
    public
    fuzzedAddress(operator);
```

### testDeregisterOperator


```solidity
function testDeregisterOperator(address operator, uint32 operatorIndex, string calldata socket)
    public
    fuzzedAddress(operator);
```

