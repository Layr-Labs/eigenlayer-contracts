# Solidity API

## ISafe

### Operation

```solidity
enum Operation {
  Call,
  DelegateCall
}
```

### setup

```solidity
function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address payable paymentReceiver) external
```

### execTransaction

```solidity
function execTransaction(address to, uint256 value, bytes data, enum ISafe.Operation operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address payable refundReceiver, bytes signatures) external payable returns (bytes)
```

### checkSignatures

```solidity
function checkSignatures(bytes32 dataHash, bytes signatures) external view
```

### approveHash

```solidity
function approveHash(bytes32 hashToApprove) external
```

