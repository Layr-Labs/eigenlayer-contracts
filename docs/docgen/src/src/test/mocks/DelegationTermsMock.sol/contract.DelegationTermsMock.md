# DelegationTermsMock
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/mocks/DelegationTermsMock.sol)

**Inherits:**
[IDelegationTerms](/docs/docgen/src/src/contracts/interfaces/IDelegationTerms.sol/interface.IDelegationTerms.md), Test


## State Variables
### shouldRevert

```solidity
bool public shouldRevert;
```


### shouldReturnData

```solidity
bool public shouldReturnData;
```


## Functions
### setShouldRevert


```solidity
function setShouldRevert(bool _shouldRevert) external;
```

### setShouldReturnData


```solidity
function setShouldReturnData(bool _shouldReturnData) external;
```

### payForService


```solidity
function payForService(IERC20, uint256) external payable;
```

### onDelegationWithdrawn


```solidity
function onDelegationWithdrawn(address, IStrategy[] memory, uint256[] memory) external view returns (bytes memory);
```

### onDelegationReceived


```solidity
function onDelegationReceived(address, IStrategy[] memory, uint256[] memory) external view returns (bytes memory);
```

