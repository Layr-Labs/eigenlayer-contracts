# BLSSignatureChecker
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/BLSSignatureChecker.sol)

**Author:**
Layr Labs, Inc.

This is the contract for checking the validity of aggregate operator signatures.


## State Variables
### registry

```solidity
IQuorumRegistry public immutable registry;
```


### BYTE_LENGTH_totalStakeIndex

```solidity
uint256 internal constant BYTE_LENGTH_totalStakeIndex = 6;
```


### BYTE_LENGTH_referenceBlockNumber

```solidity
uint256 internal constant BYTE_LENGTH_referenceBlockNumber = 4;
```


### BYTE_LENGTH_taskNumberToConfirm

```solidity
uint256 internal constant BYTE_LENGTH_taskNumberToConfirm = 4;
```


### BYTE_LENGTH_numberNonSigners

```solidity
uint256 internal constant BYTE_LENGTH_numberNonSigners = 4;
```


### BYTE_LENGTH_G1_POINT

```solidity
uint256 internal constant BYTE_LENGTH_G1_POINT = 64;
```


### BYTE_LENGTH_G2_POINT

```solidity
uint256 internal constant BYTE_LENGTH_G2_POINT = 128;
```


### BYTE_LENGTH_stakeIndex

```solidity
uint256 internal constant BYTE_LENGTH_stakeIndex = 4;
```


### BYTE_LENGTH_NON_SIGNER_INFO

```solidity
uint256 internal constant BYTE_LENGTH_NON_SIGNER_INFO = 68;
```


### BYTE_LENGTH_apkIndex

```solidity
uint256 internal constant BYTE_LENGTH_apkIndex = 4;
```


### BIT_SHIFT_totalStakeIndex

```solidity
uint256 internal constant BIT_SHIFT_totalStakeIndex = 208;
```


### BIT_SHIFT_referenceBlockNumber

```solidity
uint256 internal constant BIT_SHIFT_referenceBlockNumber = 224;
```


### BIT_SHIFT_taskNumberToConfirm

```solidity
uint256 internal constant BIT_SHIFT_taskNumberToConfirm = 224;
```


### BIT_SHIFT_numberNonSigners

```solidity
uint256 internal constant BIT_SHIFT_numberNonSigners = 224;
```


### BIT_SHIFT_stakeIndex

```solidity
uint256 internal constant BIT_SHIFT_stakeIndex = 224;
```


### BIT_SHIFT_apkIndex

```solidity
uint256 internal constant BIT_SHIFT_apkIndex = 224;
```


### CALLDATA_OFFSET_totalStakeIndex

```solidity
uint256 internal constant CALLDATA_OFFSET_totalStakeIndex = 32;
```


### CALLDATA_OFFSET_referenceBlockNumber

```solidity
uint256 internal constant CALLDATA_OFFSET_referenceBlockNumber = 38;
```


### CALLDATA_OFFSET_taskNumberToConfirm

```solidity
uint256 internal constant CALLDATA_OFFSET_taskNumberToConfirm = 42;
```


### CALLDATA_OFFSET_numberNonSigners

```solidity
uint256 internal constant CALLDATA_OFFSET_numberNonSigners = 46;
```


### CALLDATA_OFFSET_NonsignerPubkeys

```solidity
uint256 internal constant CALLDATA_OFFSET_NonsignerPubkeys = 50;
```


## Functions
### constructor


```solidity
constructor(IQuorumRegistry _registry);
```

### checkSignatures

This function is called by disperser when it has aggregated all the signatures of the operators
that are part of the quorum for a particular taskNumber and is asserting them into on-chain. The function
checks that the claim for aggregated signatures are valid.
The thesis of this procedure entails:
- computing the aggregated pubkey of all the operators that are not part of the quorum for
this specific taskNumber (represented by aggNonSignerPubkey)
- getting the aggregated pubkey of all registered nodes at the time of pre-commit by the
disperser (represented by pk),
- do subtraction of aggNonSignerPubkey from pk over Jacobian coordinate system to get aggregated pubkey
of all operators that are part of quorum.
- use this aggregated pubkey to verify the aggregated signature under BLS scheme.

*This calldata is of the format:
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
>*

*Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `referenceBlockNumber`
is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update
for the total stake (or the operator) or latest before the referenceBlockNumber.
The next step involves computing the aggregated pub key of all the operators that are not part of the quorum for this specific taskNumber.
We use a loop to iterate through the `nonSignerPK` array, loading each individual public key from calldata. Before the loop, we isolate the first public key
calldataload - this implementation saves us one ecAdd operation, which would be performed in the i=0 iteration otherwise.
Within the loop, each non-signer public key is loaded from the calldata into memory.  The most recent staking-related information is retrieved and is subtracted
from the total stake of validators in the quorum.  Then the aggregate public key and the aggregate non-signer public key is subtracted from it.
Finally  the siganture is verified by computing the elliptic curve pairing.*


```solidity
function checkSignatures(bytes calldata data)
    public
    returns (
        uint32 taskNumberToConfirm,
        uint32 referenceBlockNumber,
        bytes32 msgHash,
        SignatoryTotals memory signedTotals,
        bytes32 compressedSignatoryRecord
    );
```

### _validateOperatorStake

Get the 32 bytes immediately after the function signature and length + offset encoding of 'bytes
calldata' input type, which represents the msgHash for which the disperser is calling `checkSignatures`

retrieving the pubkey of the node in Jacobian coordinates

retrieving the index of the stake of the operator in pubkeyHashToStakeHistory in
Registry.sol that was recorded at the time of pre-commit.

retrieving the pubkey of the operator that is not part of the quorum

retrieving the index of the stake of the operator in pubkeyHashToStakeHistory in
Registry.sol that was recorded at the time of pre-commit.

Get the aggregated publickey at the moment when pre-commit happened

need to subtract aggNonSignerPubkey from the apk to get aggregate signature of all
operators that are part of the quorum

now we verify that e(sigma + gamma * pk, -g2)e(H(m) + gamma * g1, pkG2) == 1

*The next step involves computing the aggregated pub key of all the operators
that are not part of the quorum for this specific taskNumber.*

*loading pubkey for the first operator that is not part of the quorum as listed in the calldata;
Note that this need not be a special case and *could* be subsumed in the for loop below.
However, this implementation saves one ecAdd operation, which would be performed in the i=0 iteration otherwise.*

*Recall that `placeholder` here is the number of operators *not* included in the quorum*

*(input[0], input[1]) is the aggregated non singer public key*

*store each non signer's public key in (input[2], input[3]) and add them to the aggregate non signer public key*

*keep track of the aggreagate non signing stake too*

*this invariant is used in forceOperatorToDisclose in ServiceManager.sol*

*Aggregated pubkey given as part of calldata instead of being retrieved from voteWeigher reduces number of SLOADs*

*(input[2], input[3]) is the apk*


```solidity
function _validateOperatorStake(IQuorumRegistry.OperatorStake memory opStake, uint32 referenceBlockNumber)
    internal
    pure;
```

## Events
### SignatoryRecord
used for recording the event that signature has been checked in checkSignatures function.


```solidity
event SignatoryRecord(
    bytes32 msgHash,
    uint32 taskNumber,
    uint256 signedStakeFirstQuorum,
    uint256 signedStakeSecondQuorum,
    bytes32[] pubkeyHashes
);
```

## Structs
### SignatoryTotals
this data structure is used for recording the details on the total stake of the registered
operators and those operators who are part of the quorum for a particular taskNumber


```solidity
struct SignatoryTotals {
    uint256 signedStakeFirstQuorum;
    uint256 signedStakeSecondQuorum;
    uint256 totalStakeFirstQuorum;
    uint256 totalStakeSecondQuorum;
}
```

