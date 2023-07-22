# SigPDelegationTerms
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/SigP/DelegationTerms.sol)

**Inherits:**
[IDelegationTerms](/src/contracts/interfaces/IDelegationTerms.sol/interface.IDelegationTerms.md)


## State Variables
### paid

```solidity
uint256 public paid;
```


### isDelegationWithdrawn

```solidity
bytes public isDelegationWithdrawn;
```


### isDelegationReceived

```solidity
bytes public isDelegationReceived;
```


## Functions
### payForService


```solidity
function payForService(IERC20, uint256) external payable;
```

### onDelegationWithdrawn


```solidity
function onDelegationWithdrawn(address, IStrategy[] memory, uint256[] memory) external returns (bytes memory);
```

### onDelegationReceived


```solidity
function onDelegationReceived(address, IStrategy[] memory, uint256[] memory) external returns (bytes memory);
```

### delegate


```solidity
function delegate() external;
```

