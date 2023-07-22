# PausableTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/Pausable.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## Functions
### testPausingWithdrawalsFromStrategyManager

*test that pausing a contract works*


```solidity
function testPausingWithdrawalsFromStrategyManager(uint256 amountToDeposit, uint256 amountToWithdraw) public;
```

### testUnauthorizedPauserStrategyManager


```solidity
function testUnauthorizedPauserStrategyManager(address unauthorizedPauser) public fuzzedAddress(unauthorizedPauser);
```

### testSetPauser


```solidity
function testSetPauser(address newPauser) public fuzzedAddress(newPauser);
```

### testSetUnpauser


```solidity
function testSetUnpauser(address newUnpauser) public fuzzedAddress(newUnpauser);
```

### testSetPauserUnauthorized


```solidity
function testSetPauserUnauthorized(address fakePauser, address newPauser)
    public
    fuzzedAddress(newPauser)
    fuzzedAddress(fakePauser);
```

### testSetPauserRegistryUnpauser


```solidity
function testSetPauserRegistryUnpauser(IPauserRegistry newPauserRegistry) public;
```

### testSetPauserRegistyUnauthorized


```solidity
function testSetPauserRegistyUnauthorized(IPauserRegistry newPauserRegistry, address notUnpauser)
    public
    fuzzedAddress(notUnpauser);
```

## Events
### PauserRegistrySet
Emitted when the `pauserRegistry` is set to `newPauserRegistry`.


```solidity
event PauserRegistrySet(IPauserRegistry pauserRegistry, IPauserRegistry newPauserRegistry);
```

