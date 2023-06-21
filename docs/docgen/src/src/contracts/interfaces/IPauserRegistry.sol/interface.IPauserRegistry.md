# IPauserRegistry
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IPauserRegistry.sol)

**Author:**
Layr Labs, Inc.


## Functions
### pauser

Unique address that holds the pauser role.


```solidity
function pauser() external view returns (address);
```

### unpauser

Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.


```solidity
function unpauser() external view returns (address);
```

