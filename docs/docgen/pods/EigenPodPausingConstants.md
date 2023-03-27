# Solidity API

## EigenPodPausingConstants

### PAUSED_NEW_EIGENPODS

```solidity
uint8 PAUSED_NEW_EIGENPODS
```

Index for flag that pauses creation of new EigenPods when set. See EigenPodManager code for details.

### PAUSED_WITHDRAW_RESTAKED_ETH

```solidity
uint8 PAUSED_WITHDRAW_RESTAKED_ETH
```

Index for flag that pauses the `withdrawRestakedBeaconChainETH` function *of the EigenPodManager* when set. See EigenPodManager code for details.

### PAUSED_EIGENPODS_VERIFY_CREDENTIALS

```solidity
uint8 PAUSED_EIGENPODS_VERIFY_CREDENTIALS
```

Index for flag that pauses the `verifyCorrectWithdrawalCredentials` function *of the EigenPods* when set. see EigenPod code for details.

### PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED

```solidity
uint8 PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED
```

Index for flag that pauses the `verifyOvercommittedStake` function *of the EigenPods* when set. see EigenPod code for details.

### PAUSED_EIGENPODS_VERIFY_WITHDRAWAL

```solidity
uint8 PAUSED_EIGENPODS_VERIFY_WITHDRAWAL
```

Index for flag that pauses the `verifyBeaconChainFullWithdrawal` function *of the EigenPods* when set. see EigenPod code for details.

