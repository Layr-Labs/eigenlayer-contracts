# Solidity API

## BLSSignatureChecker

This is the contract for checking the validity of aggregate operator signatures.

### SignatoryTotals

```solidity
struct SignatoryTotals {
  uint256 signedStakeFirstQuorum;
  uint256 signedStakeSecondQuorum;
  uint256 totalStakeFirstQuorum;
  uint256 totalStakeSecondQuorum;
}
```

### SignatoryRecord

```solidity
event SignatoryRecord(bytes32 msgHash, uint32 taskNumber, uint256 signedStakeFirstQuorum, uint256 signedStakeSecondQuorum, bytes32[] pubkeyHashes)
```

used for recording the event that signature has been checked in checkSignatures function.

### registry

```solidity
contract IQuorumRegistry registry
```

### constructor

```solidity
constructor(contract IQuorumRegistry _registry) internal
```

### BYTE_LENGTH_totalStakeIndex

```solidity
uint256 BYTE_LENGTH_totalStakeIndex
```

### BYTE_LENGTH_referenceBlockNumber

```solidity
uint256 BYTE_LENGTH_referenceBlockNumber
```

### BYTE_LENGTH_taskNumberToConfirm

```solidity
uint256 BYTE_LENGTH_taskNumberToConfirm
```

### BYTE_LENGTH_numberNonSigners

```solidity
uint256 BYTE_LENGTH_numberNonSigners
```

### BYTE_LENGTH_G1_POINT

```solidity
uint256 BYTE_LENGTH_G1_POINT
```

### BYTE_LENGTH_G2_POINT

```solidity
uint256 BYTE_LENGTH_G2_POINT
```

### BYTE_LENGTH_stakeIndex

```solidity
uint256 BYTE_LENGTH_stakeIndex
```

### BYTE_LENGTH_NON_SIGNER_INFO

```solidity
uint256 BYTE_LENGTH_NON_SIGNER_INFO
```

### BYTE_LENGTH_apkIndex

```solidity
uint256 BYTE_LENGTH_apkIndex
```

### BIT_SHIFT_totalStakeIndex

```solidity
uint256 BIT_SHIFT_totalStakeIndex
```

### BIT_SHIFT_referenceBlockNumber

```solidity
uint256 BIT_SHIFT_referenceBlockNumber
```

### BIT_SHIFT_taskNumberToConfirm

```solidity
uint256 BIT_SHIFT_taskNumberToConfirm
```

### BIT_SHIFT_numberNonSigners

```solidity
uint256 BIT_SHIFT_numberNonSigners
```

### BIT_SHIFT_stakeIndex

```solidity
uint256 BIT_SHIFT_stakeIndex
```

### BIT_SHIFT_apkIndex

```solidity
uint256 BIT_SHIFT_apkIndex
```

### CALLDATA_OFFSET_totalStakeIndex

```solidity
uint256 CALLDATA_OFFSET_totalStakeIndex
```

### CALLDATA_OFFSET_referenceBlockNumber

```solidity
uint256 CALLDATA_OFFSET_referenceBlockNumber
```

### CALLDATA_OFFSET_taskNumberToConfirm

```solidity
uint256 CALLDATA_OFFSET_taskNumberToConfirm
```

### CALLDATA_OFFSET_numberNonSigners

```solidity
uint256 CALLDATA_OFFSET_numberNonSigners
```

### CALLDATA_OFFSET_NonsignerPubkeys

```solidity
uint256 CALLDATA_OFFSET_NonsignerPubkeys
```

### checkSignatures

```solidity
function checkSignatures(bytes data) public returns (uint32 taskNumberToConfirm, uint32 referenceBlockNumber, bytes32 msgHash, struct BLSSignatureChecker.SignatoryTotals signedTotals, bytes32 compressedSignatoryRecord)
```

_This calldata is of the format:
<
bytes32 msgHash, the taskHash for which disperser is calling checkSignatures
uint48 index of the totalStake corresponding to the dataStoreId in the 'totalStakeHistory' array of the BLSRegistry
uint32 blockNumber, the blockNumber at which the task was initated
uint32 taskNumberToConfirm
uint32 numberOfNonSigners,
{uint256[2] pubkeyG1, uint32 stakeIndex}[numberOfNonSigners] the G1 public key and the index to query of `pubkeyHashToStakeHistory` for each nonsigner,
uint32 apkIndex, the index in the `apkUpdates` array at which we want to load the aggregate public key
uint256[2] apkG1 (G1 aggregate public key, including nonSigners),
uint256[4] apkG2 (G2 aggregate public key, not including nonSigners),
uint256[2] sigma, the aggregate signature itself
>

Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `referenceBlockNumber`
is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update
for the total stake (or the operator) or latest before the referenceBlockNumber.
The next step involves computing the aggregated pub key of all the operators that are not part of the quorum for this specific taskNumber.
We use a loop to iterate through the `nonSignerPK` array, loading each individual public key from calldata. Before the loop, we isolate the first public key
calldataload - this implementation saves us one ecAdd operation, which would be performed in the i=0 iteration otherwise.
Within the loop, each non-signer public key is loaded from the calldata into memory.  The most recent staking-related information is retrieved and is subtracted
from the total stake of validators in the quorum.  Then the aggregate public key and the aggregate non-signer public key is subtracted from it.
Finally  the siganture is verified by computing the elliptic curve pairing._

### _validateOperatorStake

```solidity
function _validateOperatorStake(struct IQuorumRegistry.OperatorStake opStake, uint32 referenceBlockNumber) internal pure
```

