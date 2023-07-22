# SlasherTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/Slasher.t.sol)

**Inherits:**
[EigenLayerTestHelper](/src/test/EigenLayerTestHelper.t.sol/contract.EigenLayerTestHelper.md)


## State Variables
### instance

```solidity
ISlasher instance;
```


### HEAD

```solidity
uint256 constant HEAD = 0;
```


### middleware

```solidity
address middleware = address(0xdeadbeef);
```


### middleware_2

```solidity
address middleware_2 = address(0x009849);
```


### middleware_3

```solidity
address middleware_3 = address(0x001000);
```


### middleware_4

```solidity
address middleware_4 = address(0x002000);
```


### delegationTerms

```solidity
MerkleDelegationTerms delegationTerms;
```


## Functions
### setUp


```solidity
function setUp() public override;
```

### testSlashing

this function tests the slashing process by first freezing
the operator and then calling the strategyManager.slashShares()
to actually enforce the slashing conditions.


```solidity
function testSlashing() public;
```

### testOnlyOwnerFunctions

testing ownable permissions for slashing functions
addPermissionedContracts(), removePermissionedContracts()
and resetFrozenStatus().


```solidity
function testOnlyOwnerFunctions(address incorrectCaller, address inputAddr)
    public
    fuzzedAddress(incorrectCaller)
    fuzzedAddress(inputAddr);
```

### testRecursiveCallRevert


```solidity
function testRecursiveCallRevert() public;
```

### testRecordFirstStakeUpdate


```solidity
function testRecordFirstStakeUpdate() public;
```

### testRecordStakeUpdate


```solidity
function testRecordStakeUpdate() public;
```

### testOrderingRecordStakeUpdateVuln

Register and opt into slashing with operator


```solidity
function testOrderingRecordStakeUpdateVuln() public;
```

### testOnlyRegisteredForService

Register and opt into slashing with operator


```solidity
function testOnlyRegisteredForService(address _slasher, uint32 _serveUntilBlock) public fuzzedAddress(_slasher);
```

### testOptIn


```solidity
function testOptIn(address _operator, address _slasher) public fuzzedAddress(_slasher) fuzzedAddress(_operator);
```

### testFreezeOperator


```solidity
function testFreezeOperator() public;
```

### testResetFrozenOperator


```solidity
function testResetFrozenOperator(address _attacker) public fuzzedAddress(_attacker);
```

### testRecordLastStakeUpdateAndRevokeSlashingAbility


```solidity
function testRecordLastStakeUpdateAndRevokeSlashingAbility() public;
```

