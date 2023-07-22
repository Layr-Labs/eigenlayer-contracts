# IServiceManager
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IServiceManager.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


## Functions
### taskNumber

Returns the current 'taskNumber' for the middleware


```solidity
function taskNumber() external view returns (uint32);
```

### freezeOperator

Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract


```solidity
function freezeOperator(address operator) external;
```

### recordFirstStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external;
```

### recordStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement)
    external;
```

### recordLastStakeUpdateAndRevokeSlashingAbility

Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external;
```

### latestServeUntilBlock

Returns the latest block until which operators must serve.


```solidity
function latestServeUntilBlock() external view returns (uint32);
```

### owner


```solidity
function owner() external view returns (address);
```

