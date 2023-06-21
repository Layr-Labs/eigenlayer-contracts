# ProofParsing
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/utils/ProofParsing.sol)

**Inherits:**
Test


## State Variables
### proofConfigJson

```solidity
string internal proofConfigJson;
```


### prefix

```solidity
string prefix;
```


### blockHeaderProof

```solidity
bytes32[18] blockHeaderProof;
```


### slotProof

```solidity
bytes32[3] slotProof;
```


### withdrawalProof

```solidity
bytes32[9] withdrawalProof;
```


### validatorProof

```solidity
bytes32[46] validatorProof;
```


### executionPayloadProof

```solidity
bytes32[7] executionPayloadProof;
```


### blockNumberProofs

```solidity
bytes32[4] blockNumberProofs;
```


### slotRoot

```solidity
bytes32 slotRoot;
```


### executionPayloadRoot

```solidity
bytes32 executionPayloadRoot;
```


### blockNumberRoots

```solidity
bytes32 blockNumberRoots;
```


## Functions
### setJSON


```solidity
function setJSON(string memory path) public;
```

### getSlot


```solidity
function getSlot() public returns (uint256);
```

### getValidatorIndex


```solidity
function getValidatorIndex() public returns (uint256);
```

### getWithdrawalIndex


```solidity
function getWithdrawalIndex() public returns (uint256);
```

### getBlockHeaderRootIndex


```solidity
function getBlockHeaderRootIndex() public returns (uint256);
```

### getBeaconStateRoot


```solidity
function getBeaconStateRoot() public returns (bytes32);
```

### getBlockHeaderRoot


```solidity
function getBlockHeaderRoot() public returns (bytes32);
```

### getBlockBodyRoot


```solidity
function getBlockBodyRoot() public returns (bytes32);
```

### getSlotRoot


```solidity
function getSlotRoot() public returns (bytes32);
```

### getBalanceRoot


```solidity
function getBalanceRoot() public returns (bytes32);
```

### getBlockNumberRoot


```solidity
function getBlockNumberRoot() public returns (bytes32);
```

### getExecutionPayloadRoot


```solidity
function getExecutionPayloadRoot() public returns (bytes32);
```

### getExecutionPayloadProof


```solidity
function getExecutionPayloadProof() public returns (bytes32[7] memory);
```

### getBlockNumberProof


```solidity
function getBlockNumberProof() public returns (bytes32[4] memory);
```

### getBlockHeaderProof


```solidity
function getBlockHeaderProof() public returns (bytes32[18] memory);
```

### getSlotProof


```solidity
function getSlotProof() public returns (bytes32[3] memory);
```

### getWithdrawalProof


```solidity
function getWithdrawalProof() public returns (bytes32[9] memory);
```

### getValidatorProof


```solidity
function getValidatorProof() public returns (bytes32[46] memory);
```

### getWithdrawalFields


```solidity
function getWithdrawalFields() public returns (bytes32[] memory);
```

### getValidatorFields


```solidity
function getValidatorFields() public returns (bytes32[] memory);
```

### getValidatorBalanceProof


```solidity
function getValidatorBalanceProof() public returns (bytes32[] memory);
```

### getWithdrawalCredentialProof


```solidity
function getWithdrawalCredentialProof() public returns (bytes32[] memory);
```

