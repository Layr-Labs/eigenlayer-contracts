# PauserRegistry
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/permissions/PauserRegistry.sol)

**Inherits:**
[IPauserRegistry](/docs/docgen/src/src/contracts/interfaces/IPauserRegistry.sol/interface.IPauserRegistry.md)

**Author:**
Layr Labs, Inc.


## State Variables
### pauser
Unique address that holds the pauser role.


```solidity
address public pauser;
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
constructor(address _pauser, address _unpauser);
```

### setPauser

Sets new pauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold


```solidity
function setPauser(address newPauser) external onlyUnpauser;
```

### setUnpauser

Sets new unpauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold


```solidity
function setUnpauser(address newUnpauser) external onlyUnpauser;
```

### _setPauser


```solidity
function _setPauser(address newPauser) internal;
```

### _setUnpauser


```solidity
function _setUnpauser(address newUnpauser) internal;
```

## Events
### PauserChanged

```solidity
event PauserChanged(address previousPauser, address newPauser);
```

### UnpauserChanged

```solidity
event UnpauserChanged(address previousUnpauser, address newUnpauser);
```

