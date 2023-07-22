# SlasherUnitTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/unit/SlasherUnit.t.sol)

**Inherits:**
Test, [Utils](/src/test/unit/Utils.sol/contract.Utils.md)


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### HEAD

```solidity
uint256 private constant HEAD = 0;
```


### _NULL

```solidity
uint256 private constant _NULL = 0;
```


### _HEAD

```solidity
uint256 private constant _HEAD = 0;
```


### _PREV

```solidity
bool private constant _PREV = false;
```


### _NEXT

```solidity
bool private constant _NEXT = true;
```


### PAUSED_OPT_INTO_SLASHING

```solidity
uint8 internal constant PAUSED_OPT_INTO_SLASHING = 0;
```


### PAUSED_FIRST_STAKE_UPDATE

```solidity
uint8 internal constant PAUSED_FIRST_STAKE_UPDATE = 1;
```


### PAUSED_NEW_FREEZING

```solidity
uint8 internal constant PAUSED_NEW_FREEZING = 2;
```


### MAX_CAN_SLASH_UNTIL

```solidity
uint32 internal constant MAX_CAN_SLASH_UNTIL = type(uint32).max;
```


### proxyAdmin

```solidity
ProxyAdmin public proxyAdmin;
```


### pauserRegistry

```solidity
PauserRegistry public pauserRegistry;
```


### slasherImplementation

```solidity
Slasher public slasherImplementation;
```


### slasher

```solidity
Slasher public slasher;
```


### strategyManagerMock

```solidity
StrategyManagerMock public strategyManagerMock;
```


### delegationMock

```solidity
DelegationMock public delegationMock;
```


### eigenPodManagerMock

```solidity
EigenPodManagerMock public eigenPodManagerMock;
```


### reenterer

```solidity
Reenterer public reenterer;
```


### pauser

```solidity
address public pauser = address(555);
```


### unpauser

```solidity
address public unpauser = address(999);
```


### initialOwner

```solidity
address initialOwner = address(this);
```


### dummyToken

```solidity
IERC20 public dummyToken;
```


### dummyStrat

```solidity
StrategyBase public dummyStrat;
```


### emptyUintArray

```solidity
uint256[] public emptyUintArray;
```


### contractCanSlashOperatorUntilBefore

```solidity
uint32 contractCanSlashOperatorUntilBefore;
```


### linkedListLengthBefore

```solidity
uint256 linkedListLengthBefore;
```


### middlewareTimesLengthBefore

```solidity
uint256 middlewareTimesLengthBefore;
```


### nodeExists

```solidity
bool nodeExists;
```


### prevNode

```solidity
uint256 prevNode;
```


### nextNode

```solidity
uint256 nextNode;
```


### middlewareDetailsBefore

```solidity
ISlasher.MiddlewareDetails middlewareDetailsBefore;
```


### middlewareDetailsAfter

```solidity
ISlasher.MiddlewareDetails middlewareDetailsAfter;
```


### addressIsExcludedFromFuzzedInputs

```solidity
mapping(address => bool) public addressIsExcludedFromFuzzedInputs;
```


## Functions
### filterFuzzedAddressInputs


```solidity
modifier filterFuzzedAddressInputs(address fuzzedAddress);
```

### setUp


```solidity
function setUp() public virtual;
```

### testCanCompleteNewQueuedWithdrawalAfterTotalDeregistration

Regression test for SigP's EGN2-01 issue, "Middleware can Deny Withdrawals by Revoking Slashing Prior to Queueing Withdrawal".
This test checks that a new queued withdrawal after total deregistration (i.e. queued *after* totally de-registering from all AVSs) can still eventually be completed.


```solidity
function testCanCompleteNewQueuedWithdrawalAfterTotalDeregistration(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 serveUntilBlock,
    uint32 withdrawalStartBlock
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testCanCompleteExistingQueuedWithdrawalAfterTotalDeregistration

Test related to SigP's EGN2-01 issue, "Middleware can Deny Withdrawals by Revoking Slashing Prior to Queueing Withdrawal", to ensure that the fix does not degrade performance.
This test checks that a *previous* queued withdrawal prior to total deregistration (i.e. queued *before* totally de-registering from all AVSs)
can still be withdrawn at the appropriate time, i.e. that a fix to EGN2-01 does not add any delay to existing withdrawals.


```solidity
function testCanCompleteExistingQueuedWithdrawalAfterTotalDeregistration(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 serveUntilBlock
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testCannotReinitialize


```solidity
function testCannotReinitialize() public;
```

### testOptIntoSlashing


```solidity
function testOptIntoSlashing(address operator, address contractAddress)
    public
    filterFuzzedAddressInputs(operator)
    filterFuzzedAddressInputs(contractAddress);
```

### testOptIntoSlashing_RevertsWhenPaused


```solidity
function testOptIntoSlashing_RevertsWhenPaused() public;
```

### testOptIntoSlashing_RevertsWhenCallerNotOperator


```solidity
function testOptIntoSlashing_RevertsWhenCallerNotOperator(address notOperator)
    public
    filterFuzzedAddressInputs(notOperator);
```

### testFreezeOperator


```solidity
function testFreezeOperator(address toBeFrozen, address freezingContract)
    public
    filterFuzzedAddressInputs(toBeFrozen)
    filterFuzzedAddressInputs(freezingContract);
```

### testFreezeOperatorTwice


```solidity
function testFreezeOperatorTwice(address toBeFrozen, address freezingContract) public;
```

### testFreezeOperator_RevertsWhenPaused


```solidity
function testFreezeOperator_RevertsWhenPaused(address toBeFrozen, address freezingContract)
    external
    filterFuzzedAddressInputs(toBeFrozen)
    filterFuzzedAddressInputs(freezingContract);
```

### testFreezeOperator_WhenCallerDoesntHaveSlashingPermission


```solidity
function testFreezeOperator_WhenCallerDoesntHaveSlashingPermission(address toBeFrozen, address freezingContract)
    external
    filterFuzzedAddressInputs(toBeFrozen)
    filterFuzzedAddressInputs(freezingContract);
```

### testResetFrozenStatus


```solidity
function testResetFrozenStatus(uint8 numberOfOperators, uint256 pseudorandomInput) external;
```

### testResetFrozenStatus_RevertsWhenCalledByNotOwner


```solidity
function testResetFrozenStatus_RevertsWhenCalledByNotOwner(address notOwner)
    external
    filterFuzzedAddressInputs(notOwner);
```

### testRecordFirstStakeUpdate


```solidity
function testRecordFirstStakeUpdate(address operator, address contractAddress, uint32 serveUntilBlock)
    public
    filterFuzzedAddressInputs(operator)
    filterFuzzedAddressInputs(contractAddress);
```

### _testRecordFirstStakeUpdate


```solidity
function _testRecordFirstStakeUpdate(address operator, address contractAddress, uint32 serveUntilBlock)
    internal
    filterFuzzedAddressInputs(operator)
    filterFuzzedAddressInputs(contractAddress);
```

### testRecordFirstStakeUpdate_RevertsWhenPaused


```solidity
function testRecordFirstStakeUpdate_RevertsWhenPaused() external;
```

### testRecordFirstStakeUpdate_RevertsWhenCallerDoesntHaveSlashingPermission


```solidity
function testRecordFirstStakeUpdate_RevertsWhenCallerDoesntHaveSlashingPermission() external;
```

### testRecordFirstStakeUpdate_RevertsWhenCallerAlreadyInList


```solidity
function testRecordFirstStakeUpdate_RevertsWhenCallerAlreadyInList() external;
```

### testRecordStakeUpdate


```solidity
function testRecordStakeUpdate(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 updateBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### _testRecordStakeUpdate


```solidity
function _testRecordStakeUpdate(
    address operator,
    address contractAddress,
    uint32 updateBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) internal filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordStakeUpdate_MultipleLinkedListEntries


```solidity
function testRecordStakeUpdate_MultipleLinkedListEntries(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 updateBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList


```solidity
function testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList(
    address operator,
    address contractAddress,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList_MultipleLinkedListEntries


```solidity
function testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList_MultipleLinkedListEntries(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordStakeUpdate_RevertsWhenCallerCannotSlash


```solidity
function testRecordStakeUpdate_RevertsWhenCallerCannotSlash(
    address operator,
    address contractAddress,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordStakeUpdate_RevertsWhenUpdateBlockInFuture


```solidity
function testRecordStakeUpdate_RevertsWhenUpdateBlockInFuture(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 updateBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordLastStakeUpdateAndRevokeSlashingAbility


```solidity
function testRecordLastStakeUpdateAndRevokeSlashingAbility(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 updateBlock,
    uint32 serveUntilBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerCannotSlash


```solidity
function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerCannotSlash(
    address operator,
    address contractAddress,
    uint32 serveUntilBlock
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerNotAlreadyInList


```solidity
function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerNotAlreadyInList(
    address operator,
    address contractAddress,
    uint32 serveUntilBlock
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenServeUntilBlockInputIsMax


```solidity
function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenServeUntilBlockInputIsMax(
    address operator,
    address contractAddress,
    uint32 prevServeUntilBlock,
    uint32 updateBlock,
    uint256 insertAfter
) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(contractAddress);
```

### _addressToUint


```solidity
function _addressToUint(address addr) internal pure returns (uint256);
```

### _uintToAddress


```solidity
function _uintToAddress(uint256 x) internal pure returns (address);
```

