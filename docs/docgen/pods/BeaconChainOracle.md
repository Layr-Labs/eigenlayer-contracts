# Solidity API

## BeaconChainOracle

The owner of this contract can edit a set of 'oracle signers', as well as changing the threshold number of oracle signers that must vote for a
 particular state root at a specified blockNumber before the state root is considered 'confirmed'.

### MINIMUM_THRESHOLD

```solidity
uint256 MINIMUM_THRESHOLD
```

The minimum value which the `threshold` variable is allowed to take.

### totalOracleSigners

```solidity
uint256 totalOracleSigners
```

Total number of members of the set of oracle signers.

### threshold

```solidity
uint256 threshold
```

Number of oracle signers that must vote for a state root in order for the state root to be confirmed.
Adjustable by this contract's owner through use of the `setThreshold` function.

_We note that there is an edge case -- when the threshold is adjusted downward, if a state root already has enough votes to meet the *new* threshold,
the state root must still receive one additional vote from an oracle signer to be confirmed. This behavior is intended, to minimize unexpected root confirmations._

### latestConfirmedOracleBlockNumber

```solidity
uint64 latestConfirmedOracleBlockNumber
```

Largest blockNumber that has been confirmed by the oracle.

### beaconStateRootAtBlockNumber

```solidity
mapping(uint64 => bytes32) beaconStateRootAtBlockNumber
```

Mapping: Beacon Chain blockNumber => the Beacon Chain state root at the specified blockNumber.

_This will return `bytes32(0)` if the state root at the specified blockNumber is not yet confirmed._

### isOracleSigner

```solidity
mapping(address => bool) isOracleSigner
```

Mapping: address => whether or not the address is in the set of oracle signers.

### hasVoted

```solidity
mapping(uint64 => mapping(address => bool)) hasVoted
```

Mapping: Beacon Chain blockNumber => oracle signer address => whether or not the oracle signer has voted on the state root at the blockNumber.

### stateRootVotes

```solidity
mapping(uint64 => mapping(bytes32 => uint256)) stateRootVotes
```

Mapping: Beacon Chain blockNumber => state root => total number of oracle signer votes for the state root at the blockNumber.

### ThresholdModified

```solidity
event ThresholdModified(uint256 previousValue, uint256 newValue)
```

Emitted when the value of the `threshold` variable is changed from `previousValue` to `newValue`.

### StateRootConfirmed

```solidity
event StateRootConfirmed(uint64 blockNumber, bytes32 stateRoot)
```

Emitted when the beacon chain state root at `blockNumber` is confirmed to be `stateRoot`.

### OracleSignerAdded

```solidity
event OracleSignerAdded(address addedOracleSigner)
```

Emitted when `addedOracleSigner` is added to the set of oracle signers.

### OracleSignerRemoved

```solidity
event OracleSignerRemoved(address removedOracleSigner)
```

Emitted when `removedOracleSigner` is removed from the set of oracle signers.

### onlyOracleSigner

```solidity
modifier onlyOracleSigner()
```

Modifier that restricts functions to only be callable by members of the oracle signer set

### constructor

```solidity
constructor(address initialOwner, uint256 initialThreshold, address[] initialOracleSigners) public
```

### setThreshold

```solidity
function setThreshold(uint256 _threshold) external
```

Owner-only function used to modify the value of the `threshold` variable.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _threshold | uint256 | Desired new value for the `threshold` variable. Function will revert if this is set to zero. |

### addOracleSigners

```solidity
function addOracleSigners(address[] _oracleSigners) external
```

Owner-only function used to add a signer to the set of oracle signers.

_Function will have no effect on the i-th input address if `_oracleSigners[i]`is already in the set of oracle signers._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _oracleSigners | address[] | Array of address to be added to the set. |

### removeOracleSigners

```solidity
function removeOracleSigners(address[] _oracleSigners) external
```

Owner-only function used to remove a signer from the set of oracle signers.

_Function will have no effect on the i-th input address if `_oracleSigners[i]`is already not in the set of oracle signers._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| _oracleSigners | address[] | Array of address to be removed from the set. |

### voteForBeaconChainStateRoot

```solidity
function voteForBeaconChainStateRoot(uint64 blockNumber, bytes32 stateRoot) external
```

Called by a member of the set of oracle signers to assert that the Beacon Chain state root is `stateRoot` at `blockNumber`.

_The state root will be confirmed once the total number of votes *for this exact state root at this exact blockNumber* meets the `threshold` value._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| blockNumber | uint64 | The Beacon Chain blockNumber of interest. |
| stateRoot | bytes32 | The Beacon Chain state root that the caller asserts was the correct root, at the specified `blockNumber`. |

### _setThreshold

```solidity
function _setThreshold(uint256 _threshold) internal
```

Internal function used for modifying the value of the `threshold` variable, used in the constructor and the `setThreshold` function

### _addOracleSigners

```solidity
function _addOracleSigners(address[] _oracleSigners) internal
```

Internal counterpart of the `addOracleSigners` function. Also used in the constructor.

