# Solidity API

## IDelayedWithdrawalRouter

### DelayedWithdrawal

```solidity
struct DelayedWithdrawal {
  uint224 amount;
  uint32 blockCreated;
}
```

### UserDelayedWithdrawals

```solidity
struct UserDelayedWithdrawals {
  uint256 delayedWithdrawalsCompleted;
  struct IDelayedWithdrawalRouter.DelayedWithdrawal[] delayedWithdrawals;
}
```

### createDelayedWithdrawal

```solidity
function createDelayedWithdrawal(address podOwner, address recipient) external payable
```

Creates an delayed withdrawal for `msg.value` to the `recipient`.

_Only callable by the `podOwner`'s EigenPod contract._

### claimDelayedWithdrawals

```solidity
function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) external
```

Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| recipient | address | The address to claim delayedWithdrawals for. |
| maxNumberOfWithdrawalsToClaim | uint256 | Used to limit the maximum number of withdrawals to loop through claiming. |

### claimDelayedWithdrawals

```solidity
function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) external
```

Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| maxNumberOfWithdrawalsToClaim | uint256 | Used to limit the maximum number of withdrawals to loop through claiming. |

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

Convenience function for checking whether or not the delayedWithdrawal at the `index`th entry from the `_userWithdrawals[user].delayedWithdrawals` array is currently claimable

### withdrawalDelayBlocks

```solidity
function withdrawalDelayBlocks() external view returns (uint256)
```

Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).

