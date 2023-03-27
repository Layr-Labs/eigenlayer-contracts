# Solidity API

## IServiceManager

### taskNumber

```solidity
function taskNumber() external view returns (uint32)
```

Returns the current 'taskNumber' for the middleware

### freezeOperator

```solidity
function freezeOperator(address operator) external
```

Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract

### recordFirstStakeUpdate

```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntil) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)

### recordStakeUpdate

```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntil, uint256 prevElement) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update

### recordLastStakeUpdateAndRevokeSlashingAbility

```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntil) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)

### paymentChallengeToken

```solidity
function paymentChallengeToken() external view returns (contract IERC20)
```

Token used for placing a gurantee on challenges & payment commits

### latestTime

```solidity
function latestTime() external view returns (uint32)
```

Returns the `latestTime` until which operators must serve.

### owner

```solidity
function owner() external view returns (address)
```

