# EigenPodPausingConstants
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/pods/EigenPodPausingConstants.sol)

**Author:**
Layr Labs, Inc.


## State Variables
### PAUSED_NEW_EIGENPODS
Index for flag that pauses creation of new EigenPods when set. See EigenPodManager code for details.


```solidity
uint8 internal constant PAUSED_NEW_EIGENPODS = 0;
```


### PAUSED_WITHDRAW_RESTAKED_ETH
Index for flag that pauses the `withdrawRestakedBeaconChainETH` function *of the EigenPodManager* when set. See EigenPodManager code for details.


```solidity
uint8 internal constant PAUSED_WITHDRAW_RESTAKED_ETH = 1;
```


### PAUSED_EIGENPODS_VERIFY_CREDENTIALS
Index for flag that pauses the `verifyCorrectWithdrawalCredentials` function *of the EigenPods* when set. see EigenPod code for details.


```solidity
uint8 internal constant PAUSED_EIGENPODS_VERIFY_CREDENTIALS = 2;
```


### PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED
Index for flag that pauses the `verifyOvercommittedStake` function *of the EigenPods* when set. see EigenPod code for details.


```solidity
uint8 internal constant PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED = 3;
```


### PAUSED_EIGENPODS_VERIFY_WITHDRAWAL
Index for flag that pauses the `verifyBeaconChainFullWithdrawal` function *of the EigenPods* when set. see EigenPod code for details.


```solidity
uint8 internal constant PAUSED_EIGENPODS_VERIFY_WITHDRAWAL = 4;
```


