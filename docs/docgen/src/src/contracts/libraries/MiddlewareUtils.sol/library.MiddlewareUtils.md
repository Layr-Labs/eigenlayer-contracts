# MiddlewareUtils
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/libraries/MiddlewareUtils.sol)

**Author:**
Layr Labs, Inc.


## Functions
### computeSignatoryRecordHash

Finds the `signatoryRecordHash`, used for fraudproofs.


```solidity
function computeSignatoryRecordHash(
    uint32 globalDataStoreId,
    bytes32[] memory nonSignerPubkeyHashes,
    uint256 signedStakeFirstQuorum,
    uint256 signedStakeSecondQuorum
) internal pure returns (bytes32);
```

