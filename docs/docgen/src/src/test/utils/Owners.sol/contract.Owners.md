# Owners
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/utils/Owners.sol)

**Inherits:**
Test


## State Variables
### ownersConfigJson

```solidity
string internal ownersConfigJson;
```


### addresses

```solidity
address[] addresses;
```


## Functions
### constructor


```solidity
constructor();
```

### ownerPrefix


```solidity
function ownerPrefix(uint256 index) public pure returns (string memory);
```

### getNumOperators


```solidity
function getNumOperators() public returns (uint256);
```

### getOwnerAddress


```solidity
function getOwnerAddress(uint256 index) public returns (address);
```

### getOwnerAddresses


```solidity
function getOwnerAddresses() public returns (address[] memory);
```

### getReputedOwnerAddresses


```solidity
function getReputedOwnerAddresses() public returns (address[] memory);
```

### resetOwnersConfigJson


```solidity
function resetOwnersConfigJson(string memory newConfig) public;
```

