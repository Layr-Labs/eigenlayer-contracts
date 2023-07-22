# BeaconChainProofs
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/libraries/BeaconChainProofs.sol)


## State Variables
### NUM_BEACON_BLOCK_HEADER_FIELDS

```solidity
uint256 internal constant NUM_BEACON_BLOCK_HEADER_FIELDS = 5;
```


### BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT = 3;
```


### NUM_BEACON_BLOCK_BODY_FIELDS

```solidity
uint256 internal constant NUM_BEACON_BLOCK_BODY_FIELDS = 11;
```


### BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT = 4;
```


### NUM_BEACON_STATE_FIELDS

```solidity
uint256 internal constant NUM_BEACON_STATE_FIELDS = 21;
```


### BEACON_STATE_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant BEACON_STATE_FIELD_TREE_HEIGHT = 5;
```


### NUM_ETH1_DATA_FIELDS

```solidity
uint256 internal constant NUM_ETH1_DATA_FIELDS = 3;
```


### ETH1_DATA_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant ETH1_DATA_FIELD_TREE_HEIGHT = 2;
```


### NUM_VALIDATOR_FIELDS

```solidity
uint256 internal constant NUM_VALIDATOR_FIELDS = 8;
```


### VALIDATOR_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant VALIDATOR_FIELD_TREE_HEIGHT = 3;
```


### NUM_EXECUTION_PAYLOAD_HEADER_FIELDS

```solidity
uint256 internal constant NUM_EXECUTION_PAYLOAD_HEADER_FIELDS = 15;
```


### EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT = 4;
```


### NUM_EXECUTION_PAYLOAD_FIELDS

```solidity
uint256 internal constant NUM_EXECUTION_PAYLOAD_FIELDS = 15;
```


### EXECUTION_PAYLOAD_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant EXECUTION_PAYLOAD_FIELD_TREE_HEIGHT = 4;
```


### HISTORICAL_ROOTS_TREE_HEIGHT

```solidity
uint256 internal constant HISTORICAL_ROOTS_TREE_HEIGHT = 24;
```


### HISTORICAL_BATCH_TREE_HEIGHT

```solidity
uint256 internal constant HISTORICAL_BATCH_TREE_HEIGHT = 1;
```


### STATE_ROOTS_TREE_HEIGHT

```solidity
uint256 internal constant STATE_ROOTS_TREE_HEIGHT = 13;
```


### BLOCK_ROOTS_TREE_HEIGHT

```solidity
uint256 internal constant BLOCK_ROOTS_TREE_HEIGHT = 13;
```


### NUM_WITHDRAWAL_FIELDS

```solidity
uint256 internal constant NUM_WITHDRAWAL_FIELDS = 4;
```


### WITHDRAWAL_FIELD_TREE_HEIGHT

```solidity
uint256 internal constant WITHDRAWAL_FIELD_TREE_HEIGHT = 2;
```


### VALIDATOR_TREE_HEIGHT

```solidity
uint256 internal constant VALIDATOR_TREE_HEIGHT = 40;
```


### BALANCE_TREE_HEIGHT

```solidity
uint256 internal constant BALANCE_TREE_HEIGHT = 38;
```


### WITHDRAWALS_TREE_HEIGHT

```solidity
uint256 internal constant WITHDRAWALS_TREE_HEIGHT = 4;
```


### EXECUTION_PAYLOAD_INDEX

```solidity
uint256 internal constant EXECUTION_PAYLOAD_INDEX = 9;
```


### STATE_ROOT_INDEX

```solidity
uint256 internal constant STATE_ROOT_INDEX = 3;
```


### PROPOSER_INDEX_INDEX

```solidity
uint256 internal constant PROPOSER_INDEX_INDEX = 1;
```


### SLOT_INDEX

```solidity
uint256 internal constant SLOT_INDEX = 0;
```


### BODY_ROOT_INDEX

```solidity
uint256 internal constant BODY_ROOT_INDEX = 4;
```


### STATE_ROOTS_INDEX

```solidity
uint256 internal constant STATE_ROOTS_INDEX = 6;
```


### BLOCK_ROOTS_INDEX

```solidity
uint256 internal constant BLOCK_ROOTS_INDEX = 5;
```


### HISTORICAL_ROOTS_INDEX

```solidity
uint256 internal constant HISTORICAL_ROOTS_INDEX = 7;
```


### ETH_1_ROOT_INDEX

```solidity
uint256 internal constant ETH_1_ROOT_INDEX = 8;
```


### VALIDATOR_TREE_ROOT_INDEX

```solidity
uint256 internal constant VALIDATOR_TREE_ROOT_INDEX = 11;
```


### BALANCE_INDEX

```solidity
uint256 internal constant BALANCE_INDEX = 12;
```


### EXECUTION_PAYLOAD_HEADER_INDEX

```solidity
uint256 internal constant EXECUTION_PAYLOAD_HEADER_INDEX = 24;
```


### HISTORICAL_BATCH_STATE_ROOT_INDEX

```solidity
uint256 internal constant HISTORICAL_BATCH_STATE_ROOT_INDEX = 1;
```


### VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX

```solidity
uint256 internal constant VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX = 1;
```


### VALIDATOR_BALANCE_INDEX

```solidity
uint256 internal constant VALIDATOR_BALANCE_INDEX = 2;
```


### VALIDATOR_SLASHED_INDEX

```solidity
uint256 internal constant VALIDATOR_SLASHED_INDEX = 3;
```


### VALIDATOR_WITHDRAWABLE_EPOCH_INDEX

```solidity
uint256 internal constant VALIDATOR_WITHDRAWABLE_EPOCH_INDEX = 7;
```


### BLOCK_NUMBER_INDEX

```solidity
uint256 internal constant BLOCK_NUMBER_INDEX = 6;
```


### WITHDRAWALS_ROOT_INDEX

```solidity
uint256 internal constant WITHDRAWALS_ROOT_INDEX = 14;
```


### WITHDRAWALS_INDEX

```solidity
uint256 internal constant WITHDRAWALS_INDEX = 14;
```


### WITHDRAWAL_VALIDATOR_INDEX_INDEX

```solidity
uint256 internal constant WITHDRAWAL_VALIDATOR_INDEX_INDEX = 1;
```


### WITHDRAWAL_VALIDATOR_AMOUNT_INDEX

```solidity
uint256 internal constant WITHDRAWAL_VALIDATOR_AMOUNT_INDEX = 3;
```


### HISTORICALBATCH_STATEROOTS_INDEX

```solidity
uint256 internal constant HISTORICALBATCH_STATEROOTS_INDEX = 1;
```


### SLOTS_PER_EPOCH

```solidity
uint256 internal constant SLOTS_PER_EPOCH = 32;
```


### UINT64_MASK

```solidity
bytes8 internal constant UINT64_MASK = 0xffffffffffffffff;
```


## Functions
### getBalanceFromBalanceRoot

This function is parses the balanceRoot to get the uint64 balance of a validator.  During merkleization of the
beacon state balance tree, four uint64 values (making 32 bytes) are grouped together and treated as a single leaf in the merkle tree. Thus the
validatorIndex mod 4 is used to determine which of the four uint64 values to extract from the balanceRoot.


```solidity
function getBalanceFromBalanceRoot(uint40 validatorIndex, bytes32 balanceRoot) internal pure returns (uint64);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`validatorIndex`|`uint40`|is the index of the validator being proven for.|
|`balanceRoot`|`bytes32`|is the combination of 4 validator balances being proven for.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint64`|The validator's balance, in Gwei|


### verifyValidatorFields

This function verifies merkle proofs of the fields of a certain validator against a beacon chain state root


```solidity
function verifyValidatorFields(
    uint40 validatorIndex,
    bytes32 beaconStateRoot,
    bytes calldata proof,
    bytes32[] calldata validatorFields
) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`validatorIndex`|`uint40`|the index of the proven validator|
|`beaconStateRoot`|`bytes32`|is the beacon chain state root to be proven against.|
|`proof`|`bytes`|is the data used in proving the validator's fields|
|`validatorFields`|`bytes32[]`|the claimed fields of the validator|


### verifyValidatorBalance

Note: the length of the validator merkle proof is BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1.
There is an additional layer added by hashing the root with the length of the validator list

This function verifies merkle proofs of the balance of a certain validator against a beacon chain state root


```solidity
function verifyValidatorBalance(
    uint40 validatorIndex,
    bytes32 beaconStateRoot,
    bytes calldata proof,
    bytes32 balanceRoot
) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`validatorIndex`|`uint40`|the index of the proven validator|
|`beaconStateRoot`|`bytes32`|is the beacon chain state root to be proven against.|
|`proof`|`bytes`|is the proof of the balance against the beacon chain state root|
|`balanceRoot`|`bytes32`|is the serialized balance used to prove the balance of the validator (refer to `getBalanceFromBalanceRoot` above for detailed explanation)|


### verifyWithdrawalProofs

the beacon state's balance list is a list of uint64 values, and these are grouped together in 4s when merkleized.
Therefore, the index of the balance of a validator is validatorIndex/4

This function verifies the slot and the withdrawal fields for a given withdrawal


```solidity
function verifyWithdrawalProofs(
    bytes32 beaconStateRoot,
    WithdrawalProofs calldata proofs,
    bytes32[] calldata withdrawalFields
) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`beaconStateRoot`|`bytes32`|is the beacon chain state root to be proven against.|
|`proofs`|`WithdrawalProofs`|is the provided set of merkle proofs|
|`withdrawalFields`|`bytes32[]`|is the serialized withdrawal container to be proven|


## Structs
### WithdrawalProofs

```solidity
struct WithdrawalProofs {
    bytes blockHeaderProof;
    bytes withdrawalProof;
    bytes slotProof;
    bytes executionPayloadProof;
    bytes blockNumberProof;
    uint64 blockHeaderRootIndex;
    uint64 withdrawalIndex;
    bytes32 blockHeaderRoot;
    bytes32 blockBodyRoot;
    bytes32 slotRoot;
    bytes32 blockNumberRoot;
    bytes32 executionPayloadRoot;
}
```

### ValidatorFieldsAndBalanceProofs

```solidity
struct ValidatorFieldsAndBalanceProofs {
    bytes validatorFieldsProof;
    bytes validatorBalanceProof;
    bytes32 balanceRoot;
}
```

### ValidatorFieldsProof

```solidity
struct ValidatorFieldsProof {
    bytes validatorProof;
    uint40 validatorIndex;
}
```

