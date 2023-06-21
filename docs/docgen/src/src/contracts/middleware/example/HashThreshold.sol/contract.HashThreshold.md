# HashThreshold
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/middleware/example/HashThreshold.sol)

**Inherits:**
Ownable, [IServiceManager](/docs/docgen/src/src/contracts/interfaces/IServiceManager.sol/interface.IServiceManager.md)

**Author:**
Layr Labs, Inc.


## State Variables
### disputePeriodBlocks

```solidity
uint32 constant disputePeriodBlocks = 1 days / 12 seconds;
```


### numZeroes

```solidity
uint8 constant numZeroes = 5;
```


### slasher

```solidity
ISlasher public immutable slasher;
```


### registry

```solidity
ECDSARegistry public immutable registry;
```


### taskNumber

```solidity
uint32 public taskNumber = 0;
```


### latestServeUntilBlock

```solidity
uint32 public latestServeUntilBlock = 0;
```


### certifiedMessageMetadatas

```solidity
mapping(bytes32 => CertifiedMessageMetadata) public certifiedMessageMetadatas;
```


## Functions
### onlyRegistry


```solidity
modifier onlyRegistry();
```

### constructor


```solidity
constructor(ISlasher _slasher, ECDSARegistry _registry);
```

### owner


```solidity
function owner() public view override(Ownable, IServiceManager) returns (address);
```

### decaHash


```solidity
function decaHash(bytes32 message) public pure returns (bytes32);
```

### submitSignatures

This function is called by anyone to certify a message. Signers are certifying that the decahashed message starts with at least `numZeros` 0s.


```solidity
function submitSignatures(bytes32 message, bytes calldata signatures) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`message`|`bytes32`|The message to certify|
|`signatures`|`bytes`|The signatures of the message, certifying it|


### slashSigners

This function is called by anyone to slash the signers of an invalid message that has been certified.


```solidity
function slashSigners(bytes32 message, bytes calldata signatures) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`message`|`bytes32`|The message to slash the signers of|
|`signatures`|`bytes`|The signatures that certified the message|


### freezeOperator

Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract


```solidity
function freezeOperator(address operator) external onlyRegistry;
```

### recordFirstStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)


```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external onlyRegistry;
```

### recordLastStakeUpdateAndRevokeSlashingAbility

Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)


```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock)
    external
    onlyRegistry;
```

### recordStakeUpdate

Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update


```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement)
    external
    onlyRegistry;
```

## Events
### MessageCertified

```solidity
event MessageCertified(bytes32);
```

## Structs
### CertifiedMessageMetadata

```solidity
struct CertifiedMessageMetadata {
    bytes32 signaturesHash;
    uint32 validAfterBlock;
}
```

