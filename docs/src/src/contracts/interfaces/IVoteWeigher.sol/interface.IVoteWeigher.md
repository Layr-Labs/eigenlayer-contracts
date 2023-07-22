# IVoteWeigher
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IVoteWeigher.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

Note that `NUMBER_OF_QUORUMS` is expected to remain constant, as suggested by its uppercase formatting.


## Functions
### weightOfOperator

This function computes the total weight of the @param operator in the quorum @param quorumNumber.

*returns zero in the case that `quorumNumber` is greater than or equal to `NUMBER_OF_QUORUMS`*


```solidity
function weightOfOperator(address operator, uint256 quorumNumber) external returns (uint96);
```

### NUMBER_OF_QUORUMS

Number of quorums that are being used by the middleware.


```solidity
function NUMBER_OF_QUORUMS() external view returns (uint256);
```

### quorumBips

This defines the earnings split between different quorums. Mapping is quorumNumber => BIPS which the quorum earns, out of the total earnings.

*The sum of all entries, i.e. sum(quorumBips[0] through quorumBips[NUMBER_OF_QUORUMS - 1]) should *always* be 10,000!*


```solidity
function quorumBips(uint256 quorumNumber) external view returns (uint256);
```

