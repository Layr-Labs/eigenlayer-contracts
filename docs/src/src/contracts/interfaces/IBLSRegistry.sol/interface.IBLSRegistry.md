# IBLSRegistry
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IBLSRegistry.sol)

**Inherits:**
[IQuorumRegistry](/src/contracts/interfaces/IQuorumRegistry.sol/interface.IQuorumRegistry.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

Adds BLS-specific functions to the base interface.


## Functions
### getCorrectApkHash

get hash of a historical aggregated public key corresponding to a given index;
called by checkSignatures in BLSSignatureChecker.sol.


```solidity
function getCorrectApkHash(uint256 index, uint32 blockNumber) external returns (bytes32);
```

### apkUpdates

returns the `ApkUpdate` struct at `index` in the list of APK updates


```solidity
function apkUpdates(uint256 index) external view returns (ApkUpdate memory);
```

### apkHashes

returns the APK hash that resulted from the `index`th APK update


```solidity
function apkHashes(uint256 index) external view returns (bytes32);
```

### apkUpdateBlockNumbers

returns the block number at which the `index`th APK update occurred


```solidity
function apkUpdateBlockNumbers(uint256 index) external view returns (uint32);
```

### operatorWhitelister


```solidity
function operatorWhitelister() external view returns (address);
```

### operatorWhitelistEnabled


```solidity
function operatorWhitelistEnabled() external view returns (bool);
```

### whitelisted


```solidity
function whitelisted(address) external view returns (bool);
```

### setOperatorWhitelistStatus


```solidity
function setOperatorWhitelistStatus(bool _operatorWhitelistEnabled) external;
```

### addToOperatorWhitelist


```solidity
function addToOperatorWhitelist(address[] calldata) external;
```

### removeFromWhitelist


```solidity
function removeFromWhitelist(address[] calldata operators) external;
```

## Structs
### ApkUpdate
Data structure used to track the history of the Aggregate Public Key of all operators


```solidity
struct ApkUpdate {
    bytes32 apkHash;
    uint32 blockNumber;
}
```

