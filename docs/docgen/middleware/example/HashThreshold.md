# Solidity API

## HashThreshold

### disputePeriodBlocks

```solidity
uint32 disputePeriodBlocks
```

### numZeroes

```solidity
uint8 numZeroes
```

### slasher

```solidity
contract ISlasher slasher
```

### registry

```solidity
contract ECDSARegistry registry
```

### CertifiedMessageMetadata

```solidity
struct CertifiedMessageMetadata {
  bytes32 signaturesHash;
  uint32 validAfterBlock;
}
```

### taskNumber

```solidity
uint32 taskNumber
```

Returns the current 'taskNumber' for the middleware

### latestServeUntilBlock

```solidity
uint32 latestServeUntilBlock
```

Returns the latest block until which operators must serve.

### certifiedMessageMetadatas

```solidity
mapping(bytes32 => struct HashThreshold.CertifiedMessageMetadata) certifiedMessageMetadatas
```

### MessageCertified

```solidity
event MessageCertified(bytes32)
```

### onlyRegistry

```solidity
modifier onlyRegistry()
```

### constructor

```solidity
constructor(contract ISlasher _slasher, contract ECDSARegistry _registry) public
```

### owner

```solidity
function owner() public view returns (address)
```

### decaHash

```solidity
function decaHash(bytes32 message) public pure returns (bytes32)
```

### submitSignatures

```solidity
function submitSignatures(bytes32 message, bytes signatures) external
```

This function is called by anyone to certify a message. Signers are certifying that the decahashed message starts with at least `numZeros` 0s.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| message | bytes32 | The message to certify |
| signatures | bytes | The signatures of the message, certifying it |

### slashSigners

```solidity
function slashSigners(bytes32 message, bytes signatures) external
```

This function is called by anyone to slash the signers of an invalid message that has been certified.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| message | bytes32 | The message to slash the signers of |
| signatures | bytes | The signatures that certified the message |

### freezeOperator

```solidity
function freezeOperator(address operator) external
```

Permissioned function that causes the ServiceManager to freeze the operator on EigenLayer, through a call to the Slasher contract

### recordFirstStakeUpdate

```solidity
function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording an initial stake update (on operator registration)

### recordLastStakeUpdateAndRevokeSlashingAbility

```solidity
function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording a final stake update (on operator deregistration)

### recordStakeUpdate

```solidity
function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) external
```

Permissioned function to have the ServiceManager forward a call to the slasher, recording a stake update

