# RegistryBase
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/RegistryBase.sol)

**Inherits:**
[VoteWeigherBase](/docs/docgen/src/src/contracts/middleware/VoteWeigherBase.sol/abstract.VoteWeigherBase.md), [IQuorumRegistry](/docs/docgen/src/src/contracts/interfaces/IQuorumRegistry.sol/interface.IQuorumRegistry.md)

**Author:**
Layr Labs, Inc.

This contract is used for
- registering new operators
- committing to and finalizing de-registration as an operator
- updating the stakes of the operator

*This contract is missing key functions. See `BLSRegistry` or `ECDSARegistry` for examples that inherit from this contract.*


## State Variables
### minimumStakeFirstQuorum
In order to register, an operator must have at least `minimumStakeFirstQuorum` or `minimumStakeSecondQuorum`, as
evaluated by this contract's 'VoteWeigher' logic.


```solidity
uint128 public minimumStakeFirstQuorum = 1 wei;
```


### minimumStakeSecondQuorum

```solidity
uint128 public minimumStakeSecondQuorum = 1 wei;
```


### registry
used for storing Operator info on each operator while registration


```solidity
mapping(address => Operator) public registry;
```


### operatorList
used for storing the list of current and past registered operators


```solidity
address[] public operatorList;
```


### totalStakeHistory
array of the history of the total stakes -- marked as internal since getTotalStakeFromIndex is a getter for this


```solidity
OperatorStake[] internal totalStakeHistory;
```


### totalOperatorsHistory
array of the history of the number of operators, and the taskNumbers at which the number of operators changed


```solidity
OperatorIndex[] public totalOperatorsHistory;
```


### pubkeyHashToStakeHistory
mapping from operator's pubkeyhash to the history of their stake updates


```solidity
mapping(bytes32 => OperatorStake[]) public pubkeyHashToStakeHistory;
```


### pubkeyHashToIndexHistory
mapping from operator's pubkeyhash to the history of their index in the array of all operators


```solidity
mapping(bytes32 => OperatorIndex[]) public pubkeyHashToIndexHistory;
```


## Functions
### constructor

Irrevocably sets the (immutable) `delegation` & `strategyManager` addresses, and `NUMBER_OF_QUORUMS` variable.


```solidity
constructor(IStrategyManager _strategyManager, IServiceManager _serviceManager, uint8 _NUMBER_OF_QUORUMS)
    VoteWeigherBase(_strategyManager, _serviceManager, _NUMBER_OF_QUORUMS);
```

### _initialize

Adds empty first entries to the dynamic arrays `totalStakeHistory` and `totalOperatorsHistory`,
to record an initial condition of zero operators with zero total stake.
Adds `_firstQuorumStrategiesConsideredAndMultipliers` and `_secondQuorumStrategiesConsideredAndMultipliers` to the dynamic arrays
`strategiesConsideredAndMultipliers[0]` and `strategiesConsideredAndMultipliers[1]` (i.e. to the weighing functions of the quorums)


```solidity
function _initialize(
    uint256[] memory _quorumBips,
    StrategyAndWeightingMultiplier[] memory _firstQuorumStrategiesConsideredAndMultipliers,
    StrategyAndWeightingMultiplier[] memory _secondQuorumStrategiesConsideredAndMultipliers
) internal virtual onlyInitializing;
```

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

Since the 'to' field represents the blockNumber at which a new index started, it is OK if the
previous array entry has 'to' == blockNumber, so we check not strict inequality here
When deregistering, the operator does *not* serve the current block number -- 'to' gets set (from zero) to the current block number.
Since the 'to' field represents the blocknumber at which a new index started, we want to check strict inequality here.

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


### isActiveOperator

Since the 'to' field represents the blockNumber at which a new index started, it is OK if the
previous array entry has 'to' == blockNumber, so we check not strict inequality here

Returns whether or not the `operator` is currently an active operator, i.e. is "registered".


```solidity
function isActiveOperator(address operator) external view virtual returns (bool);
```

### getOperatorPubkeyHash

Returns the stored pubkeyHash for the specified `operator`.


```solidity
function getOperatorPubkeyHash(address operator) public view returns (bytes32);
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

verify that the stake was non-zero at the time (note: here was use the assumption that the operator was 'inactive'
once their stake fell to zero)

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


### getMostRecentStakeByOperator

verify that the stake was zero at the time (note: here was use the assumption that the operator was 'inactive'
once their stake fell to zero)

Returns the most recent stake weight for the `operator`

*Function returns an OperatorStake struct with **every entry equal to 0** in the event that the operator has no stake history*


```solidity
function getMostRecentStakeByOperator(address operator) public view returns (OperatorStake memory);
```

### getStakeHistoryLength


```solidity
function getStakeHistoryLength(bytes32 pubkeyHash) external view returns (uint256);
```

### firstQuorumStakedByOperator


```solidity
function firstQuorumStakedByOperator(address operator) external view returns (uint96);
```

### secondQuorumStakedByOperator


```solidity
function secondQuorumStakedByOperator(address operator) external view returns (uint96);
```

### operatorStakes

Returns the most recent stake weights for the `operator`

*Function returns weights of **0** in the event that the operator has no stake history*


```solidity
function operatorStakes(address operator) public view returns (uint96, uint96);
```

### totalStake

Returns the stake amounts from the latest entry in `totalStakeHistory`.


```solidity
function totalStake() external view returns (uint96, uint96);
```

### getLengthOfPubkeyHashStakeHistory


```solidity
function getLengthOfPubkeyHashStakeHistory(bytes32 pubkeyHash) external view returns (uint256);
```

### getLengthOfPubkeyHashIndexHistory


```solidity
function getLengthOfPubkeyHashIndexHistory(bytes32 pubkeyHash) external view returns (uint256);
```

### getLengthOfTotalStakeHistory


```solidity
function getLengthOfTotalStakeHistory() external view returns (uint256);
```

### getLengthOfTotalOperatorsHistory


```solidity
function getLengthOfTotalOperatorsHistory() external view returns (uint256);
```

### getTotalStakeFromIndex

Returns the `index`-th entry in the dynamic array of total stake, `totalStakeHistory`.

*Function will revert in the event that `index` is out-of-bounds.*


```solidity
function getTotalStakeFromIndex(uint256 index) external view returns (OperatorStake memory);
```

### getFromTaskNumberForOperator

Returns task number from when `operator` has been registered.


```solidity
function getFromTaskNumberForOperator(address operator) external view returns (uint32);
```

### numOperators

Returns the current number of operators of this service.


```solidity
function numOperators() public view returns (uint32);
```

### setMinimumStakeFirstQuorum

Adjusts the `minimumStakeFirstQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 1st quorum.


```solidity
function setMinimumStakeFirstQuorum(uint128 _minimumStakeFirstQuorum) external onlyServiceManagerOwner;
```

### setMinimumStakeSecondQuorum

Adjusts the `minimumStakeSecondQuorum` -- i.e. the node stake (weight) requirement for inclusion in the 2nd quorum.


```solidity
function setMinimumStakeSecondQuorum(uint128 _minimumStakeSecondQuorum) external onlyServiceManagerOwner;
```

### updateSocket


```solidity
function updateSocket(string calldata newSocket) external;
```

### _updateTotalOperatorsHistory

Called when the total number of operators has changed.
Sets the `toBlockNumber` field on the last entry *so far* in thedynamic array `totalOperatorsHistory` to the current `block.number`,
recording that the previous entry is *no longer the latest* and the block number at which the next was added.
Pushes a new entry to `totalOperatorsHistory`, with `index` field set equal to the new amount of operators, recording the new number
of total operators (and leaving the `toBlockNumber` field at zero, signaling that this is the latest entry in the array)


```solidity
function _updateTotalOperatorsHistory() internal;
```

### _removeOperator

Remove the operator from active status. Removes the operator with the given `pubkeyHash` from the `index` in `operatorList`,
updates operatorList and index histories, and performs other necessary updates for removing operator


```solidity
function _removeOperator(address operator, bytes32 pubkeyHash, uint32 index) internal virtual;
```

### _removeOperatorStake

Removes the stakes of the operator with pubkeyHash `pubkeyHash`


```solidity
function _removeOperatorStake(bytes32 pubkeyHash) internal returns (uint32);
```

### _popRegistrant

recording the information pertaining to change in stake for this operator in the history. operator stakes are set to 0 here.

Removes the registrant at the given `index` from the `operatorList`


```solidity
function _popRegistrant(uint32 index) internal returns (address swappedOperator);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`swappedOperator`|`address`|is the operator who was swapped with the removed operator in the operatorList, or the *zero address* in the case that the removed operator was already the list operator in the operatorList.|


### _addRegistrant

Adds the Operator `operator` with the given `pubkeyHash` to the `operatorList` and performs necessary related updates.


```solidity
function _addRegistrant(address operator, bytes32 pubkeyHash, OperatorStake memory _operatorStake) internal virtual;
```

### _registrationStakeEvaluation

TODO: critique: "Currently only `_registrationStakeEvaluation` uses the `uint8 registrantType` input -- we should **EITHER** store this
and keep using it in other places as well, **OR** stop using it altogether"

Used inside of inheriting contracts to validate the registration of `operator` and find their `OperatorStake`.

*This function does **not** update the stored state of the operator's stakes -- storage updates are performed elsewhere.*


```solidity
function _registrationStakeEvaluation(address operator, uint8 operatorType) internal returns (OperatorStake memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`OperatorStake`|The newly calculated `OperatorStake` for `operator`, stored in memory but not yet committed to storage.|


### _updateOperatorStake

Finds the updated stake for `operator`, stores it and records the update.

***DOES NOT UPDATE `totalStake` IN ANY WAY** -- `totalStake` updates must be done elsewhere.*


```solidity
function _updateOperatorStake(
    address operator,
    bytes32 pubkeyHash,
    OperatorStake memory currentOperatorStake,
    uint256 insertAfter
) internal returns (OperatorStake memory updatedOperatorStake);
```

### _recordTotalStakeUpdate

Records that the `totalStake` is now equal to the input param @_totalStake


```solidity
function _recordTotalStakeUpdate(OperatorStake memory _totalStake) internal;
```

### _deregistrationCheck

Verify that the `operator` is an active operator and that they've provided the correct `index`


```solidity
function _deregistrationCheck(address operator, uint32 index) internal view;
```

## Events
### SocketUpdate
emitted when `operator` updates their socket address to `socket`


```solidity
event SocketUpdate(address operator, string socket);
```

### StakeUpdate
emitted whenever the stake of `operator` is updated


```solidity
event StakeUpdate(
    address operator,
    uint96 firstQuorumStake,
    uint96 secondQuorumStake,
    uint32 updateBlockNumber,
    uint32 prevUpdateBlockNumber
);
```

### Deregistration
Emitted whenever an operator deregisters.
The `swapped` address is the address returned by an internal call to the `_popRegistrant` function.


```solidity
event Deregistration(address operator, address swapped);
```

