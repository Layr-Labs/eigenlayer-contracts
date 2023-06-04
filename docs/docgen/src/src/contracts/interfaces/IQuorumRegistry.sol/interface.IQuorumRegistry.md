# IQuorumRegistry
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IQuorumRegistry.sol)

**Inherits:**
[IRegistry](/docs/docgen/src/src/contracts/interfaces/IRegistry.sol/interface.IRegistry.md)

**Author:**
Layr Labs, Inc.

This contract does not currently support n-quorums where n >= 3.
Note in particular the presence of only `firstQuorumStake` and `secondQuorumStake` in the `OperatorStake` struct.


## Functions
### getLengthOfTotalStakeHistory


```solidity
function getLengthOfTotalStakeHistory() external view returns (uint256);
```

### getTotalStakeFromIndex

Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory`.

*Function will revert in the event that `index` is out-of-bounds.*


```solidity
function getTotalStakeFromIndex(uint256 index) external view returns (OperatorStake memory);
```

### getOperatorPubkeyHash

Returns the stored pubkeyHash for the specified `operator`.


```solidity
function getOperatorPubkeyHash(address operator) external view returns (bytes32);
```

### getFromTaskNumberForOperator

Returns task number from when `operator` has been registered.


```solidity
function getFromTaskNumberForOperator(address operator) external view returns (uint32);
```

### getStakeFromPubkeyHashAndIndex

Returns the stake weight corresponding to `pubkeyHash`, at the
`index`-th entry in the `pubkeyHashToStakeHistory[pubkeyHash]` array.

*Function will revert if `index` is out-of-bounds.*


```solidity
function getStakeFromPubkeyHashAndIndex(bytes32 pubkeyHash, uint256 index)
    external
    view
    returns (OperatorStake memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`pubkeyHash`|`bytes32`|Hash of the public key of the operator of interest.|
|`index`|`uint256`|Array index for lookup, within the dynamic array `pubkeyHashToStakeHistory[pubkeyHash]`.|


### checkOperatorActiveAtBlockNumber

Checks that the `operator` was active at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.

*In order for this function to return 'true', the inputs must satisfy all of the following list:
1) `pubkeyHashToStakeHistory[pubkeyHash][index].updateBlockNumber <= blockNumber`
2) `pubkeyHashToStakeHistory[pubkeyHash][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
is must be strictly greater than `blockNumber`
3) `pubkeyHashToStakeHistory[pubkeyHash][index].firstQuorumStake > 0`
or `pubkeyHashToStakeHistory[pubkeyHash][index].secondQuorumStake > 0`, i.e. the operator had nonzero stake*

*Note that a return value of 'false' does not guarantee that the `operator` was inactive at `blockNumber`, since a
bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.*


```solidity
function checkOperatorActiveAtBlockNumber(address operator, uint256 blockNumber, uint256 stakeHistoryIndex)
    external
    view
    returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator of interest|
|`blockNumber`|`uint256`|is the block number of interest|
|`stakeHistoryIndex`|`uint256`|specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up in `registry[operator].pubkeyHash`|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|'true' if it is successfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise|


### checkOperatorInactiveAtBlockNumber

Checks that the `operator` was inactive at the `blockNumber`, using the specified `stakeHistoryIndex` as proof.

*In order for this function to return 'true', the inputs must satisfy all of the following list:
1) `pubkeyHashToStakeHistory[pubkeyHash][index].updateBlockNumber <= blockNumber`
2) `pubkeyHashToStakeHistory[pubkeyHash][index].nextUpdateBlockNumber` must be either `0` (signifying no next update) or
is must be strictly greater than `blockNumber`
3) `pubkeyHashToStakeHistory[pubkeyHash][index].firstQuorumStake > 0`
or `pubkeyHashToStakeHistory[pubkeyHash][index].secondQuorumStake > 0`, i.e. the operator had nonzero stake*

*Note that a return value of 'false' does not guarantee that the `operator` was active at `blockNumber`, since a
bad `stakeHistoryIndex` can be supplied in order to obtain a response of 'false'.*


```solidity
function checkOperatorInactiveAtBlockNumber(address operator, uint256 blockNumber, uint256 stakeHistoryIndex)
    external
    view
    returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`|is the operator of interest|
|`blockNumber`|`uint256`|is the block number of interest|
|`stakeHistoryIndex`|`uint256`|specifies an index in `pubkeyHashToStakeHistory[pubkeyHash]`, where `pubkeyHash` is looked up in `registry[operator].pubkeyHash`|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|'true' if it is successfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise|


### getOperatorIndex

Looks up the `operator`'s index in the dynamic array `operatorList` at the specified `blockNumber`.

*Function will revert in the event that the specified `index` input does not identify the appropriate entry in the
array `pubkeyHashToIndexHistory[pubkeyHash]` to pull the info from.*


```solidity
function getOperatorIndex(address operator, uint32 blockNumber, uint32 index) external view returns (uint32);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`operator`|`address`||
|`blockNumber`|`uint32`|Is the desired block number at which we wish to query the operator's position in the `operatorList` array|
|`index`|`uint32`|Used to specify the entry within the dynamic array `pubkeyHashToIndexHistory[pubkeyHash]` to read data from, where `pubkeyHash` is looked up from `operator`'s registration info|


### getTotalOperators

Looks up the number of total operators at the specified `blockNumber`.

*This function will revert if the provided `index` is out of bounds.*


```solidity
function getTotalOperators(uint32 blockNumber, uint32 index) external view returns (uint32);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`blockNumber`|`uint32`||
|`index`|`uint32`|Input used to specify the entry within the dynamic array `totalOperatorsHistory` to read data from.|


### numOperators

Returns the current number of operators of this service.


```solidity
function numOperators() external view returns (uint32);
```

### operatorStakes

Returns the most recent stake weights for the `operator`

*Function returns weights of **0** in the event that the operator has no stake history*


```solidity
function operatorStakes(address operator) external view returns (uint96, uint96);
```

### totalStake

Returns the stake amounts from the latest entry in `totalStakeHistory`.


```solidity
function totalStake() external view returns (uint96, uint96);
```

## Structs
### Operator
Data structure for storing info on operators to be used for:
- sending data by the sequencer
- payment and associated challenges


```solidity
struct Operator {
    bytes32 pubkeyHash;
    uint32 fromTaskNumber;
    Status status;
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
struct used to store the stakes of an individual operator or the sum of all operators' stakes, for storage


```solidity
struct OperatorStake {
    uint32 updateBlockNumber;
    uint32 nextUpdateBlockNumber;
    uint96 firstQuorumStake;
    uint96 secondQuorumStake;
}
```

## Enums
### Status

```solidity
enum Status {
    INACTIVE,
    ACTIVE
}
```

