# Solidity API

## IPauserRegistry

### pauser

```solidity
function pauser() external view returns (address)
```

Unique address that holds the pauser role.

### unpauser

```solidity
function unpauser() external view returns (address)
```

Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.

