# PauserRegistry
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/permissions/PauserRegistry.sol)

**Inherits:**
[IPauserRegistry](/src/contracts/interfaces/IPauserRegistry.sol/interface.IPauserRegistry.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


## State Variables
### isPauser
Mapping of addresses to whether they hold the pauser role.


```solidity
mapping(address => bool) public isPauser;
```


### unpauser
Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.


```solidity
address public unpauser;
```


## Functions
### onlyUnpauser


```solidity
modifier onlyUnpauser();
```

### constructor


```solidity
constructor(address[] memory _pausers, address _unpauser);
```

### setIsPauser

Sets new pauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold


```solidity
function setIsPauser(address newPauser, bool canPause) external onlyUnpauser;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newPauser`|`address`|Address to be added/removed as pauser|
|`canPause`|`bool`|Whether the address should be added or removed as pauser|


### setUnpauser

Sets new unpauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold


```solidity
function setUnpauser(address newUnpauser) external onlyUnpauser;
```

### _setIsPauser


```solidity
function _setIsPauser(address pauser, bool canPause) internal;
```

### _setUnpauser


```solidity
function _setUnpauser(address newUnpauser) internal;
```

## Events
### PauserStatusChanged

```solidity
event PauserStatusChanged(address pauser, bool canPause);
```

### UnpauserChanged

```solidity
event UnpauserChanged(address previousUnpauser, address newUnpauser);
```

