# Reenterer
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/Reenterer.sol)

**Inherits:**
Test


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### target

```solidity
address public target;
```


### msgValue

```solidity
uint256 public msgValue;
```


### callData

```solidity
bytes public callData;
```


### expectedRevertData

```solidity
bytes public expectedRevertData;
```


### dataToReturn

```solidity
bytes public dataToReturn;
```


## Functions
### prepare


```solidity
function prepare(address targetToUse, uint256 msgValueToUse, bytes memory callDataToUse) external;
```

### prepare


```solidity
function prepare(
    address targetToUse,
    uint256 msgValueToUse,
    bytes memory callDataToUse,
    bytes memory expectedRevertDataToUse
) external;
```

### prepareReturnData


```solidity
function prepareReturnData(bytes memory returnDataToUse) external;
```

### receive


```solidity
receive() external payable;
```

### fallback


```solidity
fallback() external payable;
```

## Events
### Reentered

```solidity
event Reentered(bytes returnData);
```

