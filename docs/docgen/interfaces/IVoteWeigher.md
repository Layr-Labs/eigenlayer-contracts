# Solidity API

## IVoteWeigher

Note that `NUMBER_OF_QUORUMS` is expected to remain constant, as suggested by its uppercase formatting.

### weightOfOperator

```solidity
function weightOfOperator(address operator, uint256 quorumNumber) external returns (uint96)
```

This function computes the total weight of the @param operator in the quorum @param quorumNumber.

_returns zero in the case that `quorumNumber` is greater than or equal to `NUMBER_OF_QUORUMS`_

### NUMBER_OF_QUORUMS

```solidity
function NUMBER_OF_QUORUMS() external view returns (uint256)
```

Number of quorums that are being used by the middleware.

### quorumBips

```solidity
function quorumBips(uint256 quorumNumber) external view returns (uint256)
```

This defines the earnings split between different quorums. Mapping is quorumNumber => BIPS which the quorum earns, out of the total earnings.

_The sum of all entries, i.e. sum(quorumBips[0] through quorumBips[NUMBER_OF_QUORUMS - 1]) should *always* be 10,000!_

