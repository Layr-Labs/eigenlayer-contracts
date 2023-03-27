# Solidity API

## IQuorumRegistry

This contract does not currently support n-quorums where n >= 3.
Note in particular the presence of only `firstQuorumStake` and `secondQuorumStake` in the `OperatorStake` struct.

### Status

```solidity
enum Status {
  INACTIVE,
  ACTIVE
}
```

### Operator

```solidity
struct Operator {
  bytes32 pubkeyHash;
  uint32 fromTaskNumber;
  enum IQuorumRegistry.Status status;
}
```

### OperatorIndex

```solidity
struct OperatorIndex {
  uint32 toBlockNumber;
  uint32 index;
}
```

### OperatorStake

```solidity
struct OperatorStake {
  uint32 updateBlockNumber;
  uint32 nextUpdateBlockNumber;
  uint96 firstQuorumStake;
  uint96 secondQuorumStake;
}
```

### getLengthOfTotalStakeHistory

```solidity
function getLengthOfTotalStakeHistory() external view returns (uint256)
```

### getTotalStakeFromIndex

```solidity
function getTotalStakeFromIndex(uint256 index) external view returns (struct IQuorumRegistry.OperatorStake)
```

Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory`.

_Function will revert in the event that `index` is out-of-bounds._

### getOperatorPubkeyHash

```solidity
function getOperatorPubkeyHash(address operator) external view returns (bytes32)
```

Returns the stored pubkeyHash for the specified `operator`.

### getFromTaskNumberForOperator

```solidity
function getFromTaskNumberForOperator(address operator) external view returns (uint32)
```

Returns task number from when `operator` has been registered.

### getStakeFromPubkeyHashAndIndex

```solidity
function getStakeFromPubkeyHashAndIndex(bytes32 pubkeyHash, uint256 index) external view returns (struct IQuorumRegistry.OperatorStake)
```

Returns the stake weight corresponding to `pubkeyHash`, at the
`index`-th entry in the `pubkeyHashToStakeHistory[pubkeyHash]` array.

_Function will revert if `index` is out-of-bounds._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pubkeyHash | bytes32 | Hash of the public key of the operator of interest. |
| index | uint256 | Array index for lookup, within the dynamic array `pubkeyHashToStakeHistory[pubkeyHash]`. |

### checkOperatorActiveAtBlockNumber

```solidity
function checkOperatorActiveAtBlockNumber(address operator, uint256 blockNumber, uint256 stakeHistoryIndex) external view returns (bool)
```

Checks that the `operator` was active at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.

_In order for this function to return 'true', the inputs must satisfy all of the following list:
1) `pubkeyHashToStakeHistory[pubkeyHash][index].updateBlockNumber <= blockNumber`
2) `pubkeyHashToStakeHistory[pubkeyHash][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
is must be strictly greater than `blockNumber`
3) `pubkeyHashToStakeHistory[pubkeyHash][index].firstQuorumStake > 0`
or `pubkeyHashToStakeHistory[pubkeyHash][index].secondQuorumStake > 0`, i.e. the operator had nonzero stake
Note that a return value of 'false' does not guarantee that the `operator` was inactive at `blockNumber`, since a
bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the operator of interest |
| blockNumber | uint256 | is the block number of interest |
| stakeHistoryIndex | uint256 | specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up in `registry[operator].pubkeyHash` |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | 'true' if it is succesfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise |

### checkOperatorInactiveAtBlockNumber

```solidity
function checkOperatorInactiveAtBlockNumber(address operator, uint256 blockNumber, uint256 stakeHistoryIndex) external view returns (bool)
```

Checks that the `operator` was inactive at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.

_In order for this function to return 'true', the inputs must satisfy all of the following list:
1) `pubkeyHashToStakeHistory[pubkeyHash][index].updateBlockNumber <= blockNumber`
2) `pubkeyHashToStakeHistory[pubkeyHash][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
is must be strictly greater than `blockNumber`
3) `pubkeyHashToStakeHistory[pubkeyHash][index].firstQuorumStake > 0`
or `pubkeyHashToStakeHistory[pubkeyHash][index].secondQuorumStake > 0`, i.e. the operator had nonzero stake
Note that a return value of 'false' does not guarantee that the `operator` was active at `blockNumber`, since a
bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address | is the operator of interest |
| blockNumber | uint256 | is the block number of interest |
| stakeHistoryIndex | uint256 | specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up in `registry[operator].pubkeyHash` |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | 'true' if it is succesfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise |

### getOperatorIndex

```solidity
function getOperatorIndex(address operator, uint32 blockNumber, uint32 index) external view returns (uint32)
```

Looks up the `operator`'s index in the dynamic array `operatorList` at the specified `blockNumber`.

_Function will revert in the event that the specified `index` input does not identify the appropriate entry in the
array `pubkeyHashToIndexHistory[pubkeyHash]` to pull the info from._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| operator | address |  |
| blockNumber | uint32 | Is the desired block number at which we wish to query the operator's position in the `operatorList` array |
| index | uint32 | Used to specify the entry within the dynamic array `pubkeyHashToIndexHistory[pubkeyHash]` to  read data from, where `pubkeyHash` is looked up from `operator`'s registration info |

### getTotalOperators

```solidity
function getTotalOperators(uint32 blockNumber, uint32 index) external view returns (uint32)
```

Looks up the number of total operators at the specified `blockNumber`.

_This function will revert if the provided `index` is out of bounds._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| blockNumber | uint32 |  |
| index | uint32 | Input used to specify the entry within the dynamic array `totalOperatorsHistory` to read data from. |

### numOperators

```solidity
function numOperators() external view returns (uint32)
```

Returns the current number of operators of this service.

### operatorStakes

```solidity
function operatorStakes(address operator) external view returns (uint96, uint96)
```

Returns the most recent stake weights for the `operator`

_Function returns weights of **0** in the event that the operator has no stake history_

### totalStake

```solidity
function totalStake() external view returns (uint96, uint96)
```

Returns the stake amounts from the latest entry in `totalStakeHistory`.

