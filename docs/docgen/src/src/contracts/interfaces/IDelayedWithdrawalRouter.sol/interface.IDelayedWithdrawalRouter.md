# IDelayedWithdrawalRouter
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IDelayedWithdrawalRouter.sol)


## Functions
### createDelayedWithdrawal

Creates an delayed withdrawal for `msg.value` to the `recipient`.

*Only callable by the `podOwner`'s EigenPod contract.*


```solidity
function createDelayedWithdrawal(address podOwner, address recipient) external payable;
```

### claimDelayedWithdrawals

Called in order to withdraw delayed withdrawals made to the `recipient` that have passed the `withdrawalDelayBlocks` period.


```solidity
function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfWithdrawalsToClaim) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The address to claim delayedWithdrawals for.|
|`maxNumberOfWithdrawalsToClaim`|`uint256`|Used to limit the maximum number of withdrawals to loop through claiming.|


### claimDelayedWithdrawals

Called in order to withdraw delayed withdrawals made to the caller that have passed the `withdrawalDelayBlocks` period.


```solidity
function claimDelayedWithdrawals(uint256 maxNumberOfWithdrawalsToClaim) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`maxNumberOfWithdrawalsToClaim`|`uint256`|Used to limit the maximum number of withdrawals to loop through claiming.|


### setWithdrawalDelayBlocks

Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.


```solidity
function setWithdrawalDelayBlocks(uint256 newValue) external;
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

### withdrawalDelayBlocks

Delay enforced by this contract for completing any delayedWithdrawal. Measured in blocks, and adjustable by this contract's owner,
up to a maximum of `MAX_WITHDRAWAL_DELAY_BLOCKS`. Minimum value is 0 (i.e. no delay enforced).


```solidity
function withdrawalDelayBlocks() external view returns (uint256);
```

## Structs
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
    DelayedWithdrawal[] delayedWithdrawals;
}
```

