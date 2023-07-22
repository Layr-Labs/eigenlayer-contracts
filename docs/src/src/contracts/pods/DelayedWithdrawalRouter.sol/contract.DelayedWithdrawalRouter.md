# DelayedWithdrawalRouter
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/pods/DelayedWithdrawalRouter.sol)

**Inherits:**
Initializable, OwnableUpgradeable, ReentrancyGuardUpgradeable, [Pausable](/src/contracts/permissions/Pausable.sol/contract.Pausable.md), [IDelayedWithdrawalRouter](/src/contracts/interfaces/IDelayedWithdrawalRouter.sol/interface.IDelayedWithdrawalRouter.md)


## State Variables
### PAUSED_DELAYED_WITHDRAWAL_CLAIMS

```solidity
uint8 internal constant PAUSED_DELAYED_WITHDRAWAL_CLAIMS = 0;
```


### withdrawalDelayBlocks
Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).


```solidity
uint256 public withdrawalDelayBlocks;
```


### MAX_WITHDRAWAL_DELAY_BLOCKS

```solidity
uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = 50400;
```


### eigenPodManager
The EigenPodManager contract of EigenLayer.


```solidity
IEigenPodManager public immutable eigenPodManager;
```


### _userWithdrawals
Mapping: user => struct storing all delayedWithdrawal info. Marked as internal with an external getter function named `userWithdrawals`


```solidity
mapping(address => UserDelayedWithdrawals) internal _userWithdrawals;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[48] private __gap;
```


## Functions
### onlyEigenPod

Modifier used to permission a function to only be called by the EigenPod of the specified `podOwner`


```solidity
modifier onlyEigenPod(address podOwner);
```

### constructor


```solidity
constructor(IEigenPodManager _eigenPodManager);
```

### initialize


```solidity
function initialize(
    address initOwner,
    IPauserRegistry _pauserRegistry,
    uint256 initPausedStatus,
    uint256 _withdrawalDelayBlocks
) external initializer;
```

### createDelayedWithdrawal

Creates a delayed withdrawal for `msg.value` to the `recipient`.

*Only callable by the `podOwner`'s EigenPod contract.*


```solidity
function createDelayedWithdrawal(address podOwner, address recipient)
    external
    payable
    onlyEigenPod(podOwner)
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS);
```

### claimDelayedWithdrawals

Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.

*
WARNING: Note that the caller of this function cannot control where the funds are sent, but they can control when the
funds are sent once the withdrawal becomes claimable.*


```solidity
function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim)
    external
    nonReentrant
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The address to claim delayedWithdrawals for.|
|`maxNumberOfDelayedWithdrawalsToClaim`|`uint256`|Used to limit the maximum number of delayedWithdrawals to loop through claiming.|


### claimDelayedWithdrawals

Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.


```solidity
function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim)
    external
    nonReentrant
    onlyWhenNotPaused(PAUSED_DELAYED_WITHDRAWAL_CLAIMS);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`maxNumberOfDelayedWithdrawalsToClaim`|`uint256`|Used to limit the maximum number of delayedWithdrawals to loop through claiming.|


### setWithdrawalDelayBlocks

Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.


```solidity
function setWithdrawalDelayBlocks(uint256 newValue) external onlyOwner;
```

### userWithdrawals

Getter function for the mapping `_userWithdrawals`


```solidity
function userWithdrawals(address user) external view returns (UserDelayedWithdrawals memory);
```

### getUserDelayedWithdrawals

Getter function to get all delayedWithdrawals of the `user`


```solidity
function getUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory);
```

### getClaimableUserDelayedWithdrawals

Getter function to get all delayedWithdrawals that are currently claimable by the `user`


```solidity
function getClaimableUserDelayedWithdrawals(address user) external view returns (DelayedWithdrawal[] memory);
```

### userDelayedWithdrawalByIndex

Getter function for fetching the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array


```solidity
function userDelayedWithdrawalByIndex(address user, uint256 index) external view returns (DelayedWithdrawal memory);
```

### userWithdrawalsLength

Getter function for fetching the length of the delayedWithdrawals array of a specific user


```solidity
function userWithdrawalsLength(address user) external view returns (uint256);
```

### canClaimDelayedWithdrawal

Convenience function for checking whether or not the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array is currently claimable


```solidity
function canClaimDelayedWithdrawal(address user, uint256 index) external view returns (bool);
```

### _claimDelayedWithdrawals

internal function used in both of the overloaded `claimDelayedWithdrawals` functions


```solidity
function _claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) internal;
```

### _setWithdrawalDelayBlocks

internal function for changing the value of `withdrawalDelayBlocks`. Also performs sanity check and emits an event.


```solidity
function _setWithdrawalDelayBlocks(uint256 newValue) internal;
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

