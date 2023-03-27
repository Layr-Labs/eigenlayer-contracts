# Solidity API

## IRegistry

Functions related to the registration process itself have been intentionally excluded
because their function signatures may vary significantly.

### isActiveOperator

```solidity
function isActiveOperator(address operator) external view returns (bool)
```

Returns 'true' if `operator` is registered as an active operator, and 'false' otherwise.

