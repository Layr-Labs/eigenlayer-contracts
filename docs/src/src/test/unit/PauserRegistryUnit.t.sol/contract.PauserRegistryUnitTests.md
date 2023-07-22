# PauserRegistryUnitTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/unit/PauserRegistryUnit.t.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### pauserRegistry

```solidity
PauserRegistry public pauserRegistry;
```


### pauser

```solidity
address public pauser = address(555);
```


### unpauser

```solidity
address public unpauser = address(999);
```


### addressIsExcludedFromFuzzedInputs

```solidity
mapping(address => bool) public addressIsExcludedFromFuzzedInputs;
```


## Functions
### setUp


```solidity
function setUp() public virtual;
```

### testSetIsPauserTrue


```solidity
function testSetIsPauserTrue(address newPauser) public;
```

### testSetIsPauserFalse


```solidity
function testSetIsPauserFalse() public;
```

### testSetUnpauser


```solidity
function testSetUnpauser(address newUnpauser) public;
```

### testSetPauser_RevertsWhenCallingFromNotUnpauser


```solidity
function testSetPauser_RevertsWhenCallingFromNotUnpauser(address notUnpauser, address newPauser) public;
```

### testSetUnpauser_RevertsWhenCallingFromNotUnpauser


```solidity
function testSetUnpauser_RevertsWhenCallingFromNotUnpauser(address notUnpauser, address newUnpauser) public;
```

### testSetPauser_RevertsWhenSettingToZeroAddress


```solidity
function testSetPauser_RevertsWhenSettingToZeroAddress() public;
```

### testSetUnpauser_RevertsWhenSettingToZeroAddress


```solidity
function testSetUnpauser_RevertsWhenSettingToZeroAddress() public;
```

## Events
### PauserStatusChanged

```solidity
event PauserStatusChanged(address pauser, bool canPause);
```

### UnpauserChanged

```solidity
event UnpauserChanged(address previousUnpauser, address newUnpauser);
```

