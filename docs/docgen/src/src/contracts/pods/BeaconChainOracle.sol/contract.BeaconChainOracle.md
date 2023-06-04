# BeaconChainOracle
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/pods/BeaconChainOracle.sol)

**Inherits:**
[IBeaconChainOracle](/docs/docgen/src/src/contracts/interfaces/IBeaconChainOracle.sol/interface.IBeaconChainOracle.md), Ownable

**Author:**
Layr Labs, Inc.

The owner of this contract can edit a set of 'oracle signers', as well as changing the threshold number of oracle signers that must vote for a
particular state root at a specified blockNumber before the state root is considered 'confirmed'.


## State Variables
### MINIMUM_THRESHOLD
The minimum value which the `threshold` variable is allowed to take.


```solidity
uint256 public constant MINIMUM_THRESHOLD = 1;
```


### totalOracleSigners
Total number of members of the set of oracle signers.


```solidity
uint256 public totalOracleSigners;
```


### threshold
Number of oracle signers that must vote for a state root in order for the state root to be confirmed.
Adjustable by this contract's owner through use of the `setThreshold` function.

*We note that there is an edge case -- when the threshold is adjusted downward, if a state root already has enough votes to meet the *new* threshold,
the state root must still receive one additional vote from an oracle signer to be confirmed. This behavior is intended, to minimize unexpected root confirmations.*


```solidity
uint256 public threshold;
```


### latestConfirmedOracleBlockNumber
Largest blockNumber that has been confirmed by the oracle.


```solidity
uint64 public latestConfirmedOracleBlockNumber;
```


### beaconStateRootAtBlockNumber
Mapping: Beacon Chain blockNumber => the Beacon Chain state root at the specified blockNumber.

*This will return `bytes32(0)` if the state root at the specified blockNumber is not yet confirmed.*


```solidity
mapping(uint64 => bytes32) public beaconStateRootAtBlockNumber;
```


### isOracleSigner
Mapping: address => whether or not the address is in the set of oracle signers.


```solidity
mapping(address => bool) public isOracleSigner;
```


### hasVoted
Mapping: Beacon Chain blockNumber => oracle signer address => whether or not the oracle signer has voted on the state root at the blockNumber.


```solidity
mapping(uint64 => mapping(address => bool)) public hasVoted;
```


### stateRootVotes
Mapping: Beacon Chain blockNumber => state root => total number of oracle signer votes for the state root at the blockNumber.


```solidity
mapping(uint64 => mapping(bytes32 => uint256)) public stateRootVotes;
```


## Functions
### onlyOracleSigner

Modifier that restricts functions to only be callable by members of the oracle signer set


```solidity
modifier onlyOracleSigner();
```

### constructor


```solidity
constructor(address initialOwner, uint256 initialThreshold, address[] memory initialOracleSigners);
```

### setThreshold

Owner-only function used to modify the value of the `threshold` variable.


```solidity
function setThreshold(uint256 _threshold) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_threshold`|`uint256`|Desired new value for the `threshold` variable. Function will revert if this is set to zero.|


### addOracleSigners

Owner-only function used to add a signer to the set of oracle signers.

*Function will have no effect on the i-th input address if `_oracleSigners[i]`is already in the set of oracle signers.*


```solidity
function addOracleSigners(address[] memory _oracleSigners) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_oracleSigners`|`address[]`|Array of address to be added to the set.|


### removeOracleSigners

Owner-only function used to remove a signer from the set of oracle signers.

*Function will have no effect on the i-th input address if `_oracleSigners[i]`is already not in the set of oracle signers.*


```solidity
function removeOracleSigners(address[] memory _oracleSigners) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_oracleSigners`|`address[]`|Array of address to be removed from the set.|


### voteForBeaconChainStateRoot

Called by a member of the set of oracle signers to assert that the Beacon Chain state root is `stateRoot` at `blockNumber`.

*The state root will be confirmed once the total number of votes *for this exact state root at this exact blockNumber* meets the `threshold` value.*


```solidity
function voteForBeaconChainStateRoot(uint64 blockNumber, bytes32 stateRoot) external onlyOracleSigner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`blockNumber`|`uint64`|The Beacon Chain blockNumber of interest.|
|`stateRoot`|`bytes32`|The Beacon Chain state root that the caller asserts was the correct root, at the specified `blockNumber`.|


### _setThreshold

Internal function used for modifying the value of the `threshold` variable, used in the constructor and the `setThreshold` function


```solidity
function _setThreshold(uint256 _threshold) internal;
```

### _addOracleSigners

Internal counterpart of the `addOracleSigners` function. Also used in the constructor.


```solidity
function _addOracleSigners(address[] memory _oracleSigners) internal;
```

## Events
### ThresholdModified
Emitted when the value of the `threshold` variable is changed from `previousValue` to `newValue`.


```solidity
event ThresholdModified(uint256 previousValue, uint256 newValue);
```

### StateRootConfirmed
Emitted when the beacon chain state root at `blockNumber` is confirmed to be `stateRoot`.


```solidity
event StateRootConfirmed(uint64 blockNumber, bytes32 stateRoot);
```

### OracleSignerAdded
Emitted when `addedOracleSigner` is added to the set of oracle signers.


```solidity
event OracleSignerAdded(address addedOracleSigner);
```

### OracleSignerRemoved
Emitted when `removedOracleSigner` is removed from the set of oracle signers.


```solidity
event OracleSignerRemoved(address removedOracleSigner);
```

