# IDelayedService
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IDelayedService.sol)

**Author:**
Layr Labs, Inc.

Specifically, this interface is designed for services that consult stake amounts up to `BLOCK_STALE_MEASURE`
blocks in the past. This may be necessary due to, e.g., network processing & communication delays, or to avoid race conditions
that could be present with coordinating aggregate operator signatures while service operators are registering & de-registering.

*To clarify edge cases, the middleware can look `BLOCK_STALE_MEASURE` blocks into the past, i.e. it may trust stakes from the interval
[block.number - BLOCK_STALE_MEASURE, block.number] (specifically, *inclusive* of the block that is `BLOCK_STALE_MEASURE` before the current one)*


## Functions
### BLOCK_STALE_MEASURE

The maximum amount of blocks in the past that the service will consider stake amounts to still be 'valid'.


```solidity
function BLOCK_STALE_MEASURE() external view returns (uint32);
```

