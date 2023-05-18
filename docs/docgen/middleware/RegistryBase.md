# Solidity API

## RegistryBase

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator

_This contract is missing key functions. See `BLSRegistry` or `ECDSARegistry` for examples that inherit from this contract._

### minimumStakeFirstQuorum

```solidity
uint128 minimumStakeFirstQuorum
```

In order to register, an operator must have at least `minimumStakeFirstQuorum` or `minimumStakeSecondQuorum`, as
evaluated by this contract's 'VoteWeigher' logic.

### minimumStakeSecondQuorum

```solidity
uint128 minimumStakeSecondQuorum
```

### registry

```solidity
mapping(address => struct IQuorumRegistry.Operator) registry
```

used for storing Operator info on each operator while registration

### operatorList

```solidity
address[] operatorList
```

used for storing the list of current and past registered operators

### totalStakeHistory

```solidity
struct IQuorumRegistry.OperatorStake[] totalStakeHistory
```

array of the history of the total stakes -- marked as internal since getTotalStakeFromIndex is a getter for this

### totalOperatorsHistory

```solidity
struct IQuorumRegistry.OperatorIndex[] totalOperatorsHistory
```

array of the history of the number of operators, and the taskNumbers at which the number of operators changed

### pubkeyHashToStakeHistory

```solidity
mapping(bytes32 => struct IQuorumRegistry.OperatorStake[]) pubkeyHashToStakeHistory
```

mapping from operator's pubkeyhash to the history of their stake updates

### pubkeyHashToIndexHistory

```solidity
mapping(bytes32 => struct IQuorumRegistry.OperatorIndex[]) pubkeyHashToIndexHistory
```

mapping from operator's pubkeyhash to the history of their index in the array of all operators

### SocketUpdate

```solidity
event SocketUpdate(address operator, string socket)
```

emitted when `operator` updates their socket address to `socket`

### StakeUpdate

```solidity
event StakeUpdate(address operator, uint96 firstQuorumStake, uint96 secondQuorumStake, uint32 updateBlockNumber, uint32 prevUpdateBlockNumber)
```

emitted whenever the stake of `operator` is updated

### Deregistration

```solidity
event Deregistration(address operator, address swapped)
```

Emitted whenever an operator deregisters.
The `swapped` address is the address returned by an internal call to the `_popRegistrant` function.

### constructor

```solidity
constructor(contract IStrategyManager _strategyManager, contract IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS) internal
```

Irrevocably sets the (immutable) `delegation` & `strategyManager` addresses, and `NUMBER_OF_QUORUMS` variable.

### _initialize

```solidity
function _initialize(uint256[] _quorumBips, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _firstQuorumStrategiesConsideredAndMultipliers, struct VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] _secondQuorumStrategiesConsideredAndMultipliers) internal virtual
```

Adds empty first entries to the dynamic arrays `totalStakeHistory` and `totalOperatorsHistory`,
to record an initial condition of zero operators with zero total stake.
Adds `_firstQuorumStrategiesConsideredAndMultipliers` and `_secondQuorumStrategiesConsideredAndMultipliers` to the dynamic arrays
`strategiesConsideredAndMultipliers[0]` and `strategiesConsideredAndMultipliers[1]` (i.e. to the weighing functions of the quorums)

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

### isActiveOperator

```solidity
function isActiveOperator(address operator) external view virtual returns (bool)
```

Returns whether or not the `operator` is currently an active operator, i.e. is "registered".

### getOperatorPubkeyHash

```solidity
function getOperatorPubkeyHash(address operator) public view returns (bytes32)
```

Returns the stored pubkeyHash for the specified `operator`.

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
| [0] | bool | 'true' if it is successfully proven that  the `operator` was active at the `blockNumber`, and 'false' otherwise |

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
| [0] | bool | 'true' if it is successfully proven that  the `operator` was inactive at the `blockNumber`, and 'false' otherwise |

### getMostRecentStakeByOperator

```solidity
function getMostRecentStakeByOperator(address operator) public view returns (struct IQuorumRegistry.OperatorStake)
```

Returns the most recent stake weight for the `operator`

_Function returns an OperatorStake struct with **every entry equal to 0** in the event that the operator has no stake history_

### getStakeHistoryLength

```solidity
function getStakeHistoryLength(bytes32 pubkeyHash) external view returns (uint256)
```

### firstQuorumStakedByOperator

```solidity
function firstQuorumStakedByOperator(address operator) external view returns (uint96)
```

### secondQuorumStakedByOperator

```solidity
function secondQuorumStakedByOperator(address operator) external view returns (uint96)
```

### operatorStakes

```solidity
function operatorStakes(address operator) public view returns (uint96, uint96)
```

Returns the most recent stake weights for the `operator`

_Function returns weights of **0** in the event that the operator has no stake history_

### totalStake

```solidity
function totalStake() external view returns (uint96, uint96)
```

Returns the stake amounts from the latest entry in `totalStakeHistory`.

### getLengthOfPubkeyHashStakeHistory

```solidity
function getLengthOfPubkeyHashStakeHistory(bytes32 pubkeyHash) external view returns (uint256)
```

### getLengthOfPubkeyHashIndexHistory

```solidity
function getLengthOfPubkeyHashIndexHistory(bytes32 pubkeyHash) external view returns (uint256)
```

### getLengthOfTotalStakeHistory

```solidity
function getLengthOfTotalStakeHistory() external view returns (uint256)
```

### getLengthOfTotalOperatorsHistory

```solidity
function getLengthOfTotalOperatorsHistory() external view returns (uint256)
```

### getTotalStakeFromIndex

```solidity
function getTotalStakeFromIndex(uint256 index) external view returns (struct IQuorumRegistry.OperatorStake)
```

Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory`.

_Function will revert in the event that `index` is out-of-bounds._

### getFromTaskNumberForOperator

```solidity
function getFromTaskNumberForOperator(address operator) external view returns (uint32)
```

Returns task number from when `operator` has been registered.

### numOperators

```solidity
function numOperators() public view returns (uint32)
```

Returns the current number of operators of this service.

### setMinimumStakeFirstQuorum

```solidity
function setMinimumStakeFirstQuorum(uint128 _minimumStakeFirstQuorum) external
```

Adjusts the `minimumStakeFirstQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 1st quorum.

### setMinimumStakeSecondQuorum

```solidity
function setMinimumStakeSecondQuorum(uint128 _minimumStakeSecondQuorum) external
```

Adjusts the `minimumStakeSecondQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 2nd quorum.

### updateSocket

```solidity
function updateSocket(string newSocket) external
```

### _updateTotalOperatorsHistory

```solidity
function _updateTotalOperatorsHistory() internal
```

Called when the total number of operators has changed.
Sets the `toBlockNumber` field on the last entry *so far* in thedynamic array `totalOperatorsHistory` to the current `block.number`,
recording that the previous entry is *no longer the latest* and the block number at which the next was added.
Pushes a new entry to `totalOperatorsHistory`, with `index` field set equal to the new amount of operators, recording the new number
of total operators (and leaving the `toBlockNumber` field at zero, signaling that this is the latest entry in the array)

### _removeOperator

```solidity
function _removeOperator(address operator, bytes32 pubkeyHash, uint32 index) internal virtual
```

Remove the operator from active status. Removes the operator with the given `pubkeyHash` from the `index` in `operatorList`,
updates operatorList and index histories, and performs other necessary updates for removing operator

### _removeOperatorStake

```solidity
function _removeOperatorStake(bytes32 pubkeyHash) internal returns (uint32)
```

Removes the stakes of the operator with pubkeyHash `pubkeyHash`

### _popRegistrant

```solidity
function _popRegistrant(uint32 index) internal returns (address swappedOperator)
```

Removes the registrant at the given `index` from the `operatorList`

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| swappedOperator | address | is the operator who was swapped with the removed operator in the operatorList, or the *zero address* in the case that the removed operator was already the list operator in the operatorList. |

### _addRegistrant

```solidity
function _addRegistrant(address operator, bytes32 pubkeyHash, struct IQuorumRegistry.OperatorStake _operatorStake) internal virtual
```

Adds the Operator `operator` with the given `pubkeyHash` to the `operatorList` and performs necessary related updates.

### _registrationStakeEvaluation

```solidity
function _registrationStakeEvaluation(address operator, uint8 operatorType) internal returns (struct IQuorumRegistry.OperatorStake)
```

Used inside of inheriting contracts to validate the registration of `operator` and find their `OperatorStake`.

_This function does **not** update the stored state of the operator's stakes -- storage updates are performed elsewhere._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | struct IQuorumRegistry.OperatorStake | The newly calculated `OperatorStake` for `operator`, stored in memory but not yet committed to storage. |

### _updateOperatorStake

```solidity
function _updateOperatorStake(address operator, bytes32 pubkeyHash, struct IQuorumRegistry.OperatorStake currentOperatorStake, uint256 insertAfter) internal returns (struct IQuorumRegistry.OperatorStake updatedOperatorStake)
```

Finds the updated stake for `operator`, stores it and records the update.

_**DOES NOT UPDATE `totalStake` IN ANY WAY** -- `totalStake` updates must be done elsewhere._

### _recordTotalStakeUpdate

```solidity
function _recordTotalStakeUpdate(struct IQuorumRegistry.OperatorStake _totalStake) internal
```

Records that the `totalStake` is now equal to the input param @_totalStake

### _deregistrationCheck

```solidity
function _deregistrationCheck(address operator, uint32 index) internal view
```

Verify that the `operator` is an active operator and that they've provided the correct `index`

