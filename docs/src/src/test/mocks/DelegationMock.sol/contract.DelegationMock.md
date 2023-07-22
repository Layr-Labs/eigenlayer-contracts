# DelegationMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/DelegationMock.sol)

**Inherits:**
[IDelegationManager](/src/contracts/interfaces/IDelegationManager.sol/interface.IDelegationManager.md), Test


## State Variables
### isOperator

```solidity
mapping(address => bool) public isOperator;
```


### delegatedTo

```solidity
mapping(address => address) public delegatedTo;
```


## Functions
### setIsOperator


```solidity
function setIsOperator(address operator, bool _isOperatorReturnValue) external;
```

### registerAsOperator


```solidity
function registerAsOperator(IDelegationTerms) external;
```

### delegateTo


```solidity
function delegateTo(address operator) external;
```

### delegateToBySignature


```solidity
function delegateToBySignature(address, address, uint256, bytes memory) external;
```

### undelegate


```solidity
function undelegate(address staker) external;
```

### delegationTerms

returns the DelegationTerms of the `operator`, which may mediate their interactions with stakers who delegate to them.


```solidity
function delegationTerms(address) external view returns (IDelegationTerms);
```

### operatorShares

returns the total number of shares in `strategy` that are delegated to `operator`.


```solidity
function operatorShares(address, IStrategy) external view returns (uint256);
```

### increaseDelegatedShares


```solidity
function increaseDelegatedShares(address, IStrategy, uint256) external;
```

### decreaseDelegatedShares


```solidity
function decreaseDelegatedShares(address, IStrategy[] calldata, uint256[] calldata) external;
```

### isDelegated


```solidity
function isDelegated(address staker) external view returns (bool);
```

### isNotDelegated


```solidity
function isNotDelegated(address) external pure returns (bool);
```

