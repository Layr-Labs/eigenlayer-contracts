# BeaconChainOracleMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/BeaconChainOracleMock.sol)

**Inherits:**
[IBeaconChainOracleMock](/docs/docgen/src/src/test/mocks/IBeaconChainOracleMock.sol/interface.IBeaconChainOracleMock.md)


## State Variables
### mockBeaconChainStateRoot

```solidity
bytes32 public mockBeaconChainStateRoot;
```


## Functions
### getBeaconChainStateRoot


```solidity
function getBeaconChainStateRoot() external view returns (bytes32);
```

### setBeaconChainStateRoot


```solidity
function setBeaconChainStateRoot(bytes32 beaconChainStateRoot) external;
```

### beaconStateRootAtBlockNumber


```solidity
function beaconStateRootAtBlockNumber(uint64) external view returns (bytes32);
```

### isOracleSigner


```solidity
function isOracleSigner(address) external pure returns (bool);
```

### hasVoted


```solidity
function hasVoted(uint64, address) external pure returns (bool);
```

### stateRootVotes


```solidity
function stateRootVotes(uint64, bytes32) external pure returns (uint256);
```

### totalOracleSigners


```solidity
function totalOracleSigners() external pure returns (uint256);
```

### threshold


```solidity
function threshold() external pure returns (uint256);
```

### setThreshold


```solidity
function setThreshold(uint256) external pure;
```

### addOracleSigners


```solidity
function addOracleSigners(address[] memory) external pure;
```

### removeOracleSigners


```solidity
function removeOracleSigners(address[] memory) external pure;
```

### voteForBeaconChainStateRoot


```solidity
function voteForBeaconChainStateRoot(uint64, bytes32) external pure;
```

### latestConfirmedOracleBlockNumber


```solidity
function latestConfirmedOracleBlockNumber() external pure returns (uint64);
```

