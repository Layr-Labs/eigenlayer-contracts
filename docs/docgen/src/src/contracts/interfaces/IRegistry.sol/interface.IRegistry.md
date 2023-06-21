# IRegistry
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IRegistry.sol)

**Author:**
Layr Labs, Inc.

Functions related to the registration process itself have been intentionally excluded
because their function signatures may vary significantly.


## Functions
### isActiveOperator

Returns 'true' if `operator` is registered as an active operator, and 'false' otherwise.


```solidity
function isActiveOperator(address operator) external view returns (bool);
```

