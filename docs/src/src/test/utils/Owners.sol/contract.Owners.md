# Owners
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/utils/Owners.sol)

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

