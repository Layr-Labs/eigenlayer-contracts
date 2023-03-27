# Solidity API

## DelayedWithdrawalRouter

### WithdrawalDelayBlocksSet

```solidity
event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
```

Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.

### PAUSED_DELAYED_WITHDRAWAL_CLAIMS

```solidity
uint8 PAUSED_DELAYED_WITHDRAWAL_CLAIMS
```

### withdrawalDelayBlocks

```solidity
uint256 withdrawalDelayBlocks
```

Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).

### MAX_WITHDRAWAL_DELAY_BLOCKS

```solidity
uint256 MAX_WITHDRAWAL_DELAY_BLOCKS
```

### eigenPodManager

```solidity
contract IEigenPodManager eigenPodManager
```

The EigenPodManager contract of EigenLayer.

### _userWithdrawals

```solidity
mapping(address => struct IDelayedWithdrawalRouter.UserDelayedWithdrawals) _userWithdrawals
```

Mapping: user => struct storing all delayedWithdrawal info. Marked as internal with an external getter function named `userWithdrawals`

### DelayedWithdrawalCreated

```solidity
event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
```

event for delayedWithdrawal creation

### DelayedWithdrawalsClaimed

```solidity
event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
```

event for the claiming of delayedWithdrawals

### onlyEigenPod

```solidity
modifier onlyEigenPod(address podOwner)
```

Modifier used to permission a function to only be called by the EigenPod of the specified `podOwner`

### constructor

```solidity
constructor(contract IEigenPodManager _eigenPodManager) public
```

### initialize

```solidity
function initialize(address initOwner, contract IPauserRegistry _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) external
```

### createDelayedWithdrawal

```solidity
function createDelayedWithdrawal(address podOwner, address recipient) external payable
```

Creates a delayed withdrawal for `msg.value` to the `recipient`.

_Only callable by the `podOwner`'s EigenPod contract._

### claimDelayedWithdrawals

```solidity
function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) external
```

Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | address | The address to claim delayedWithdrawals for. |
| maxNumberOfDelayedWithdrawalsToClaim | uint256 | Used to limit the maximum number of delayedWithdrawals to loop through claiming. |

### claimDelayedWithdrawals

```solidity
function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) external
```

Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| maxNumberOfDelayedWithdrawalsToClaim | uint256 | Used to limit the maximum number of delayedWithdrawals to loop through claiming. |

### setWithdrawalDelayBlocks

```solidity
function setWithdrawalDelayBlocks(uint256 newValue) external
```

Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.

### userWithdrawals

```solidity
function userWithdrawals(address user) external view returns (struct IDelayedWithdrawalRouter.UserDelayedWithdrawals)
```

Getter function for the mapping `_userWithdrawals`

### claimableUserDelayedWithdrawals

```solidity
function claimableUserDelayedWithdrawals(address user) external view returns (struct IDelayedWithdrawalRouter.DelayedWithdrawal[])
```

Getter function to get all delayedWithdrawals that are currently claimable by the `user`

### userDelayedWithdrawalByIndex

```solidity
function userDelayedWithdrawalByIndex(address user, uint256 index) external view returns (struct IDelayedWithdrawalRouter.DelayedWithdrawal)
```

Getter function for fetching the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array

### userWithdrawalsLength

```solidity
function userWithdrawalsLength(address user) external view returns (uint256)
```

Getter function for fetching the length of the delayedWithdrawals array of a specific user

### canClaimDelayedWithdrawal

```solidity
function canClaimDelayedWithdrawal(address user, uint256 index) external view returns (bool)
```

Convenience function for checking whethere or not the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array is currently claimable

### _claimDelayedWithdrawals

```solidity
function _claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) internal
```

internal function used in both of the overloaded `claimDelayedWithdrawals` functions

### _setWithdrawalDelayBlocks

```solidity
function _setWithdrawalDelayBlocks(uint256 newValue) internal
```

internal function for changing the value of `withdrawalDelayBlocks`. Also performs sanity check and emits an event.

### __gap

```solidity
uint256[48] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

