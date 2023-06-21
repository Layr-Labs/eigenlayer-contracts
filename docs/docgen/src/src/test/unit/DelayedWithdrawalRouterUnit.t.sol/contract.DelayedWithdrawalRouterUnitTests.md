# DelayedWithdrawalRouterUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/DelayedWithdrawalRouterUnit.t.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### proxyAdmin

```solidity
ProxyAdmin public proxyAdmin;
```


### pauserRegistry

```solidity
PauserRegistry public pauserRegistry;
```


### eigenPodManagerMock

```solidity
EigenPodManagerMock public eigenPodManagerMock;
```


### delayedWithdrawalRouterImplementation

```solidity
DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
```


### delayedWithdrawalRouter

```solidity
DelayedWithdrawalRouter public delayedWithdrawalRouter;
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


### delayedWithdrawalAmounts

```solidity
uint224[] public delayedWithdrawalAmounts;
```


### _pseudorandomNumber

```solidity
uint256 internal _pseudorandomNumber;
```


### PAUSED_PAYMENT_CLAIMS

```solidity
uint8 internal constant PAUSED_PAYMENT_CLAIMS = 0;
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
function setUp() external;
```

### testCannotReinitialize


```solidity
function testCannotReinitialize() external;
```

### testCreateDelayedWithdrawalNonzeroAmount


```solidity
function testCreateDelayedWithdrawalNonzeroAmount(uint224 delayedWithdrawalAmount, address podOwner, address recipient)
    public
    filterFuzzedAddressInputs(podOwner);
```

### testCreateDelayedWithdrawalZeroAmount


```solidity
function testCreateDelayedWithdrawalZeroAmount(address podOwner, address recipient)
    public
    filterFuzzedAddressInputs(podOwner);
```

### testCreateDelayedWithdrawalZeroAddress


```solidity
function testCreateDelayedWithdrawalZeroAddress(address podOwner) external filterFuzzedAddressInputs(podOwner);
```

### testCreateDelayedWithdrawalFromNonPodAddress


```solidity
function testCreateDelayedWithdrawalFromNonPodAddress(address podOwner, address nonPodAddress)
    external
    filterFuzzedAddressInputs(podOwner)
    filterFuzzedAddressInputs(nonPodAddress);
```

### testClaimDelayedWithdrawals


```solidity
function testClaimDelayedWithdrawals(
    uint8 delayedWithdrawalsToCreate,
    uint8 maxNumberOfDelayedWithdrawalsToClaim,
    uint256 pseudorandomNumber_,
    address recipient,
    bool useOverloadedFunction
) public filterFuzzedAddressInputs(recipient);
```

### testDelayedWithdrawalsGetterFunctions

This function is used to test the getter function 'getClaimableDelayedWithdrawals'


```solidity
function testDelayedWithdrawalsGetterFunctions(
    uint8 delayedWithdrawalsToCreate,
    uint224 delayedWithdrawalAmount,
    address recipient
) public filterFuzzedAddressInputs(recipient);
```

### testClaimDelayedWithdrawalsSomeUnclaimable

Creates a set of delayedWithdrawals of length (2 * `delayedWithdrawalsToCreate`),
where only the first half is claimable, claims using `maxNumberOfDelayedWithdrawalsToClaim` input,
and checks that only appropriate delayedWithdrawals were claimed.


```solidity
function testClaimDelayedWithdrawalsSomeUnclaimable(
    uint8 delayedWithdrawalsToCreate,
    uint8 maxNumberOfDelayedWithdrawalsToClaim,
    bool useOverloadedFunction
) external;
```

### testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimZero


```solidity
function testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimZero(
    uint256 pseudorandomNumber_,
    bool useOverloadedFunction
) external;
```

### testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimNonzero


```solidity
function testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimNonzero(
    uint256 pseudorandomNumber_,
    bool useOverloadedFunction
) external;
```

### testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimZero


```solidity
function testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimZero(
    uint256 pseudorandomNumber_,
    bool useOverloadedFunction
) external;
```

### testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimNonzero


```solidity
function testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimNonzero(
    uint8 maxNumberOfDelayedWithdrawalsToClaim,
    uint256 pseudorandomNumber_,
    bool useOverloadedFunction
) external;
```

### testClaimDelayedWithdrawals_RevertsOnAttemptingReentrancy


```solidity
function testClaimDelayedWithdrawals_RevertsOnAttemptingReentrancy(bool useOverloadedFunction) external;
```

### testClaimDelayedWithdrawals_RevertsWhenPaused


```solidity
function testClaimDelayedWithdrawals_RevertsWhenPaused(bool useOverloadedFunction) external;
```

### testSetWithdrawalDelayBlocks


```solidity
function testSetWithdrawalDelayBlocks(uint16 valueToSet) external;
```

### testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner


```solidity
function testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner(address notOwner)
    external
    filterFuzzedAddressInputs(notOwner);
```

### testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh


```solidity
function testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh(uint256 valueToSet) external;
```

### _getPseudorandomNumber


```solidity
function _getPseudorandomNumber() internal returns (uint256);
```

## Events
### WithdrawalDelayBlocksSet
Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.


```solidity
event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);
```

### DelayedWithdrawalCreated
event for delayedWithdrawal creation


```solidity
event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index);
```

### DelayedWithdrawalsClaimed
event for the claiming of delayedWithdrawals


```solidity
event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted);
```

