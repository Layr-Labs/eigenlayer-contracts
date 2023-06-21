# SigPDelegationTerms
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/SigP/DelegationTerms.sol)

**Inherits:**
[IDelegationTerms](/docs/docgen/src/src/contracts/interfaces/IDelegationTerms.sol/interface.IDelegationTerms.md)


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

