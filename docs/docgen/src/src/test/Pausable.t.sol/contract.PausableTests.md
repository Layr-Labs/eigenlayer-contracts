# PausableTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/Pausable.t.sol)

**Inherits:**
[EigenLayerTestHelper](/docs/docgen/src/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


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

