# Solidity API

## PauserRegistry

### pauser

```solidity
address pauser
```

Unique address that holds the pauser role.

### unpauser

```solidity
address unpauser
```

Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.

### PauserChanged

```solidity
event PauserChanged(address previousPauser, address newPauser)
```

### UnpauserChanged

```solidity
event UnpauserChanged(address previousUnpauser, address newUnpauser)
```

### onlyUnpauser

```solidity
modifier onlyUnpauser()
```

### constructor

```solidity
constructor(address _pauser, address _unpauser) public
```

### setPauser

```solidity
function setPauser(address newPauser) external
```

Sets new pauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold

### setUnpauser

```solidity
function setUnpauser(address newUnpauser) external
```

Sets new unpauser - only callable by unpauser, as the unpauser is expected to be kept more secure, e.g. being a multisig with a higher threshold

### _setPauser

```solidity
function _setPauser(address newPauser) internal
```

### _setUnpauser

```solidity
function _setUnpauser(address newUnpauser) internal
```

