# ServiceManagerMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/ServiceManagerMock.sol)

**Inherits:**
[IServiceManager](/docs/docgen/src/src/contracts/interfaces/IServiceManager.sol/interface.IServiceManager.md), DSTest


## State Variables
### slasher

```solidity
ISlasher public slasher;
```


## Functions
### constructor


```solidity
constructor(ISlasher _slasher);
```

### taskNumber

Returns the current 'taskNumber' for the middleware


```solidity
function taskNumber() external pure returns (uint32);
```

### freezeOperator

Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract


```solidity
function freezeOperator(address operator) external;
```

### recordFirstStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntil) external pure;
```

### recordStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement)
    external
    pure;
```

### recordLastStakeUpdateAndRevokeSlashingAbility

Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntil) external pure;
```

### paymentChallengeToken

Token used for placing guarantee on challenges & payment commits


```solidity
function paymentChallengeToken() external pure returns (IERC20);
```

### delegationManager

The Delegation contract of EigenLayer.


```solidity
function delegationManager() external pure returns (IDelegationManager);
```

### latestServeUntilBlock

Returns the `latestServeUntilBlock` until which operators must serve.


```solidity
function latestServeUntilBlock() external pure returns (uint32);
```

### owner


```solidity
function owner() external pure returns (address);
```

