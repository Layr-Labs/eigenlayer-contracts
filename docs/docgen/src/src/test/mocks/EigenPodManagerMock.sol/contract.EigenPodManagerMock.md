# EigenPodManagerMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/EigenPodManagerMock.sol)

**Inherits:**
[IEigenPodManager](/docs/docgen/src/src/contracts/interfaces/IEigenPodManager.sol/interface.IEigenPodManager.md), Test


## Functions
### slasher


```solidity
function slasher() external view returns (ISlasher);
```

### createPod


```solidity
function createPod() external pure;
```

### stake


```solidity
function stake(bytes calldata, bytes calldata, bytes32) external payable;
```

### restakeBeaconChainETH


```solidity
function restakeBeaconChainETH(address, uint256) external pure;
```

### recordOvercommittedBeaconChainETH


```solidity
function recordOvercommittedBeaconChainETH(address, uint256, uint256) external pure;
```

### withdrawRestakedBeaconChainETH


```solidity
function withdrawRestakedBeaconChainETH(address, address, uint256) external pure;
```

### updateBeaconChainOracle


```solidity
function updateBeaconChainOracle(IBeaconChainOracle) external pure;
```

### ownerToPod


```solidity
function ownerToPod(address) external pure returns (IEigenPod);
```

### getPod


```solidity
function getPod(address podOwner) external pure returns (IEigenPod);
```

### beaconChainOracle


```solidity
function beaconChainOracle() external pure returns (IBeaconChainOracle);
```

### getBeaconChainStateRoot


```solidity
function getBeaconChainStateRoot(uint64) external pure returns (bytes32);
```

### strategyManager


```solidity
function strategyManager() external pure returns (IStrategyManager);
```

### hasPod


```solidity
function hasPod(address) external pure returns (bool);
```

### pause


```solidity
function pause(uint256) external;
```

### pauseAll


```solidity
function pauseAll() external;
```

### paused


```solidity
function paused() external pure returns (uint256);
```

### paused


```solidity
function paused(uint8) external pure returns (bool);
```

### pauserRegistry


```solidity
function pauserRegistry() external pure returns (IPauserRegistry);
```

### unpause


```solidity
function unpause(uint256) external;
```

