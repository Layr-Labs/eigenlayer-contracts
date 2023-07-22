# IPauserRegistry
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IPauserRegistry.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


## Functions
### isPauser

Mapping of addresses to whether they hold the pauser role.


```solidity
function isPauser(address pauser) external view returns (bool);
```

### unpauser

Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.


```solidity
function unpauser() external view returns (address);
```

