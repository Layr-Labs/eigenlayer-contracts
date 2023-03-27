# Solidity API

## MiddlewareUtils

### computeSignatoryRecordHash

```solidity
function computeSignatoryRecordHash(uint32 globalDataStoreId, bytes32[] nonSignerPubkeyHashes, uint256 signedStakeFirstQuorum, uint256 signedStakeSecondQuorum) internal pure returns (bytes32)
```

Finds the `signatoryRecordHash`, used for fraudproofs.

