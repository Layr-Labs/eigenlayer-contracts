# PauserRegistryUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/PauserRegistryUnit.t.sol)

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

### testSetPauser


```solidity
function testSetPauser(address newPauser) public;
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
### PauserChanged

```solidity
event PauserChanged(address previousPauser, address newPauser);
```

### UnpauserChanged

```solidity
event UnpauserChanged(address previousUnpauser, address newUnpauser);
```

