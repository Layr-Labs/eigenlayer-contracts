# IRegistry
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IRegistry.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

Functions related to the registration process itself have been intentionally excluded
because their function signatures may vary significantly.


## Functions
### isActiveOperator

Returns 'true' if `operator` is registered as an active operator, and 'false' otherwise.


```solidity
function isActiveOperator(address operator) external view returns (bool);
```

