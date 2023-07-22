# Operators
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/utils/Operators.sol)

**Inherits:**
Test


## State Variables
### operatorConfigJson

```solidity
string internal operatorConfigJson;
```


## Functions
### constructor


```solidity
constructor();
```

### operatorPrefix


```solidity
function operatorPrefix(uint256 index) public pure returns (string memory);
```

### getNumOperators


```solidity
function getNumOperators() public returns (uint256);
```

### getOperatorAddress


```solidity
function getOperatorAddress(uint256 index) public returns (address);
```

### getOperatorSchnorrSignature


```solidity
function getOperatorSchnorrSignature(uint256 index) public returns (uint256, BN254.G1Point memory);
```

### getOperatorSecretKey


```solidity
function getOperatorSecretKey(uint256 index) public returns (uint256);
```

### getOperatorPubkeyG1


```solidity
function getOperatorPubkeyG1(uint256 index) public returns (BN254.G1Point memory);
```

### getOperatorPubkeyG2


```solidity
function getOperatorPubkeyG2(uint256 index) public returns (BN254.G2Point memory);
```

### readUint


```solidity
function readUint(string memory json, uint256 index, string memory key) public returns (uint256);
```

### stringToUint


```solidity
function stringToUint(string memory s) public pure returns (uint256);
```

### setOperatorJsonFilePath


```solidity
function setOperatorJsonFilePath(string memory filepath) public;
```

