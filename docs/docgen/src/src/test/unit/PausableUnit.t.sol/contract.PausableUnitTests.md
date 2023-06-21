# PausableUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/PausableUnit.t.sol)

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


### pausable

```solidity
PausableHarness public pausable;
```


### pauser

```solidity
address public pauser = address(555);
```


### unpauser

```solidity
address public unpauser = address(999);
```


### initPausedStatus

```solidity
uint256 public initPausedStatus = 0;
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

### testCannotReinitialize


```solidity
function testCannotReinitialize(address _pauserRegistry, uint256 _initPausedStatus) public;
```

### testCannotInitializeWithZeroAddress


```solidity
function testCannotInitializeWithZeroAddress(uint256 _initPausedStatus) public;
```

### testPause


```solidity
function testPause(uint256 previousPausedStatus, uint256 newPausedStatus) public;
```

### testPause_RevertsWhenCalledByNotPauser


```solidity
function testPause_RevertsWhenCalledByNotPauser(address notPauser, uint256 newPausedStatus) public;
```

### testPauseAll


```solidity
function testPauseAll(uint256 previousPausedStatus) public;
```

### testPauseAll_RevertsWhenCalledByNotPauser


```solidity
function testPauseAll_RevertsWhenCalledByNotPauser(address notPauser) public;
```

### testPause_RevertsWhenTryingToUnpause


```solidity
function testPause_RevertsWhenTryingToUnpause(uint256 previousPausedStatus, uint256 newPausedStatus) public;
```

### testPauseSingleIndex


```solidity
function testPauseSingleIndex(uint256 previousPausedStatus, uint8 indexToPause) public;
```

### testUnpause


```solidity
function testUnpause(uint256 previousPausedStatus, uint256 newPausedStatus) public;
```

### testUnpause_RevertsWhenCalledByNotUnpauser


```solidity
function testUnpause_RevertsWhenCalledByNotUnpauser(address notUnpauser, uint256 newPausedStatus) public;
```

### testUnpause_RevertsWhenTryingToPause


```solidity
function testUnpause_RevertsWhenTryingToPause(uint256 previousPausedStatus, uint256 newPausedStatus) public;
```

## Events
### Paused
Emitted when the pause is triggered by `account`, and changed to `newPausedStatus`.


```solidity
event Paused(address indexed account, uint256 newPausedStatus);
```

### Unpaused
Emitted when the pause is lifted by `account`, and changed to `newPausedStatus`.


```solidity
event Unpaused(address indexed account, uint256 newPausedStatus);
```

