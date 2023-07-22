# MiddlewareUtils
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/libraries/MiddlewareUtils.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service


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

