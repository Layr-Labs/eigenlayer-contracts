# Solidity API

## IBeaconChainOracle

### latestConfirmedOracleBlockNumber

```solidity
function latestConfirmedOracleBlockNumber() external view returns (uint64)
```

Largest blockNumber that has been confirmed by the oracle.

### beaconStateRootAtBlockNumber

```solidity
function beaconStateRootAtBlockNumber(uint64 blockNumber) external view returns (bytes32)
```

Mapping: Beacon Chain blockNumber => the Beacon Chain state root at the specified blockNumber.

_This will return `bytes32(0)` if the state root is not yet finalized at the blockNumber._

### isOracleSigner

```solidity
function isOracleSigner(address _oracleSigner) external view returns (bool)
```

Mapping: address => whether or not the address is in the set of oracle signers.

### hasVoted

```solidity
function hasVoted(uint64 blockNumber, address oracleSigner) external view returns (bool)
```

Mapping: Beacon Chain blockNumber => oracle signer address => whether or not the oracle signer has voted on the state root at the blockNumber.

### stateRootVotes

```solidity
function stateRootVotes(uint64 blockNumber, bytes32 stateRoot) external view returns (uint256)
```

Mapping: Beacon Chain blockNumber => state root => total number of oracle signer votes for the state root at the blockNumber.

### totalOracleSigners

```solidity
function totalOracleSigners() external view returns (uint256)
```

Total number of members of the set of oracle signers.

### threshold

```solidity
function threshold() external view returns (uint256)
```

Number of oracle signers that must vote for a state root in order for the state root to be confirmed.
Adjustable by this contract's owner through use of the `setThreshold` function.

_We note that there is an edge case -- when the threshold is adjusted downward, if a state root already has enough votes to meet the *new* threshold,
the state root must still receive one additional vote from an oracle signer to be confirmed. This behavior is intended, to minimize unexpected root confirmations._

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

_The state root will be finalized once the total number of votes *for this exact state root at this exact blockNumber* meets the `threshold` value._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| blockNumber | uint64 | The Beacon Chain blockNumber of interest. |
| stateRoot | bytes32 | The Beacon Chain state root that the caller asserts was the correct root, at the specified `blockNumber`. |

