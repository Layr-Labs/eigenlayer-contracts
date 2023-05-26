# Solidity API

## BeaconChainProofs

### NUM_BEACON_BLOCK_HEADER_FIELDS

```solidity
uint256 NUM_BEACON_BLOCK_HEADER_FIELDS
```

### BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT

```solidity
uint256 BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT
```

### NUM_BEACON_BLOCK_BODY_FIELDS

```solidity
uint256 NUM_BEACON_BLOCK_BODY_FIELDS
```

### BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT

```solidity
uint256 BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT
```

### NUM_BEACON_STATE_FIELDS

```solidity
uint256 NUM_BEACON_STATE_FIELDS
```

### BEACON_STATE_FIELD_TREE_HEIGHT

```solidity
uint256 BEACON_STATE_FIELD_TREE_HEIGHT
```

### NUM_ETH1_DATA_FIELDS

```solidity
uint256 NUM_ETH1_DATA_FIELDS
```

### ETH1_DATA_FIELD_TREE_HEIGHT

```solidity
uint256 ETH1_DATA_FIELD_TREE_HEIGHT
```

### NUM_VALIDATOR_FIELDS

```solidity
uint256 NUM_VALIDATOR_FIELDS
```

### VALIDATOR_FIELD_TREE_HEIGHT

```solidity
uint256 VALIDATOR_FIELD_TREE_HEIGHT
```

### NUM_EXECUTION_PAYLOAD_HEADER_FIELDS

```solidity
uint256 NUM_EXECUTION_PAYLOAD_HEADER_FIELDS
```

### EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT

```solidity
uint256 EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT
```

### NUM_EXECUTION_PAYLOAD_FIELDS

```solidity
uint256 NUM_EXECUTION_PAYLOAD_FIELDS
```

### EXECUTION_PAYLOAD_FIELD_TREE_HEIGHT

```solidity
uint256 EXECUTION_PAYLOAD_FIELD_TREE_HEIGHT
```

### HISTORICAL_ROOTS_TREE_HEIGHT

```solidity
uint256 HISTORICAL_ROOTS_TREE_HEIGHT
```

### HISTORICAL_BATCH_TREE_HEIGHT

```solidity
uint256 HISTORICAL_BATCH_TREE_HEIGHT
```

### STATE_ROOTS_TREE_HEIGHT

```solidity
uint256 STATE_ROOTS_TREE_HEIGHT
```

### BLOCK_ROOTS_TREE_HEIGHT

```solidity
uint256 BLOCK_ROOTS_TREE_HEIGHT
```

### NUM_WITHDRAWAL_FIELDS

```solidity
uint256 NUM_WITHDRAWAL_FIELDS
```

### WITHDRAWAL_FIELD_TREE_HEIGHT

```solidity
uint256 WITHDRAWAL_FIELD_TREE_HEIGHT
```

### VALIDATOR_TREE_HEIGHT

```solidity
uint256 VALIDATOR_TREE_HEIGHT
```

### BALANCE_TREE_HEIGHT

```solidity
uint256 BALANCE_TREE_HEIGHT
```

### WITHDRAWALS_TREE_HEIGHT

```solidity
uint256 WITHDRAWALS_TREE_HEIGHT
```

### EXECUTION_PAYLOAD_INDEX

```solidity
uint256 EXECUTION_PAYLOAD_INDEX
```

### STATE_ROOT_INDEX

```solidity
uint256 STATE_ROOT_INDEX
```

### PROPOSER_INDEX_INDEX

```solidity
uint256 PROPOSER_INDEX_INDEX
```

### SLOT_INDEX

```solidity
uint256 SLOT_INDEX
```

### BODY_ROOT_INDEX

```solidity
uint256 BODY_ROOT_INDEX
```

### STATE_ROOTS_INDEX

```solidity
uint256 STATE_ROOTS_INDEX
```

### BLOCK_ROOTS_INDEX

```solidity
uint256 BLOCK_ROOTS_INDEX
```

### HISTORICAL_ROOTS_INDEX

```solidity
uint256 HISTORICAL_ROOTS_INDEX
```

### ETH_1_ROOT_INDEX

```solidity
uint256 ETH_1_ROOT_INDEX
```

### VALIDATOR_TREE_ROOT_INDEX

```solidity
uint256 VALIDATOR_TREE_ROOT_INDEX
```

### BALANCE_INDEX

```solidity
uint256 BALANCE_INDEX
```

### EXECUTION_PAYLOAD_HEADER_INDEX

```solidity
uint256 EXECUTION_PAYLOAD_HEADER_INDEX
```

### HISTORICAL_BATCH_STATE_ROOT_INDEX

```solidity
uint256 HISTORICAL_BATCH_STATE_ROOT_INDEX
```

### VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX

```solidity
uint256 VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX
```

### VALIDATOR_BALANCE_INDEX

```solidity
uint256 VALIDATOR_BALANCE_INDEX
```

### VALIDATOR_SLASHED_INDEX

```solidity
uint256 VALIDATOR_SLASHED_INDEX
```

### VALIDATOR_WITHDRAWABLE_EPOCH_INDEX

```solidity
uint256 VALIDATOR_WITHDRAWABLE_EPOCH_INDEX
```

### BLOCK_NUMBER_INDEX

```solidity
uint256 BLOCK_NUMBER_INDEX
```

### WITHDRAWALS_ROOT_INDEX

```solidity
uint256 WITHDRAWALS_ROOT_INDEX
```

### WITHDRAWALS_INDEX

```solidity
uint256 WITHDRAWALS_INDEX
```

### WITHDRAWAL_VALIDATOR_INDEX_INDEX

```solidity
uint256 WITHDRAWAL_VALIDATOR_INDEX_INDEX
```

### WITHDRAWAL_VALIDATOR_AMOUNT_INDEX

```solidity
uint256 WITHDRAWAL_VALIDATOR_AMOUNT_INDEX
```

### HISTORICALBATCH_STATEROOTS_INDEX

```solidity
uint256 HISTORICALBATCH_STATEROOTS_INDEX
```

### SLOTS_PER_EPOCH

```solidity
uint256 SLOTS_PER_EPOCH
```

### UINT64_MASK

```solidity
bytes8 UINT64_MASK
```

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

### computePhase0BeaconBlockHeaderRoot

```solidity
function computePhase0BeaconBlockHeaderRoot(bytes32[5] blockHeaderFields) internal pure returns (bytes32)
```

### computePhase0BeaconStateRoot

```solidity
function computePhase0BeaconStateRoot(bytes32[21] beaconStateFields) internal pure returns (bytes32)
```

### computePhase0ValidatorRoot

```solidity
function computePhase0ValidatorRoot(bytes32[8] validatorFields) internal pure returns (bytes32)
```

### computePhase0Eth1DataRoot

```solidity
function computePhase0Eth1DataRoot(bytes32[3] eth1DataFields) internal pure returns (bytes32)
```

### getBalanceFromBalanceRoot

```solidity
function getBalanceFromBalanceRoot(uint40 validatorIndex, bytes32 balanceRoot) internal pure returns (uint64)
```

This function is parses the balanceRoot to get the uint64 balance of a validator.  During merkleization of the
beacon state balance tree, four uint64 values (making 32 bytes) are grouped together and treated as a single leaf in the merkle tree. Thus the
validatorIndex mod 4 is used to determine which of the four uint64 values to extract from the balanceRoot.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| validatorIndex | uint40 | is the index of the validator being proven for. |
| balanceRoot | bytes32 | is the combination of 4 validator balances being proven for. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint64 | The validator's balance, in Gwei |

### verifyValidatorFields

```solidity
function verifyValidatorFields(uint40 validatorIndex, bytes32 beaconStateRoot, bytes proof, bytes32[] validatorFields) internal view
```

This function verifies merkle proofs of the fields of a certain validator against a beacon chain state root

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| validatorIndex | uint40 | the index of the proven validator |
| beaconStateRoot | bytes32 | is the beacon chain state root to be proven against. |
| proof | bytes | is the data used in proving the validator's fields |
| validatorFields | bytes32[] | the claimed fields of the validator |

### verifyValidatorBalance

```solidity
function verifyValidatorBalance(uint40 validatorIndex, bytes32 beaconStateRoot, bytes proof, bytes32 balanceRoot) internal view
```

This function verifies merkle proofs of the balance of a certain validator against a beacon chain state root

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| validatorIndex | uint40 | the index of the proven validator |
| beaconStateRoot | bytes32 | is the beacon chain state root to be proven against. |
| proof | bytes | is the proof of the balance against the beacon chain state root |
| balanceRoot | bytes32 | is the serialized balance used to prove the balance of the validator (refer to `getBalanceFromBalanceRoot` above for detailed explanation) |

### verifyWithdrawalProofs

```solidity
function verifyWithdrawalProofs(bytes32 beaconStateRoot, struct BeaconChainProofs.WithdrawalProofs proofs, bytes32[] withdrawalFields) internal view
```

This function verifies the slot and the withdrawal fields for a given withdrawal

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| beaconStateRoot | bytes32 | is the beacon chain state root to be proven against. |
| proofs | struct BeaconChainProofs.WithdrawalProofs | is the provided set of merkle proofs |
| withdrawalFields | bytes32[] | is the serialized withdrawal container to be proven |

