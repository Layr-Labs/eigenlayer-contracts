# DelegationTermsMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/DelegationTermsMock.sol)

**Inherits:**
[IDelegationTerms](/src/contracts/interfaces/IDelegationTerms.sol/interface.IDelegationTerms.md), Test


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

