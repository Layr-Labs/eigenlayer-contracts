# IETHPOSDeposit
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IETHPOSDeposit.sol)

This is the Ethereum 2.0 deposit contract interface.
For more information see the Phase 0 specification under https://github.com/ethereum/eth2.0-specs


## Functions
### deposit

Submit a Phase 0 DepositData object.


```solidity
function deposit(
    bytes calldata pubkey,
    bytes calldata withdrawal_credentials,
    bytes calldata signature,
    bytes32 deposit_data_root
) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`pubkey`|`bytes`|A BLS12-381 public key.|
|`withdrawal_credentials`|`bytes`|Commitment to a public key for withdrawals.|
|`signature`|`bytes`|A BLS12-381 signature.|
|`deposit_data_root`|`bytes32`|The SHA-256 hash of the SSZ-encoded DepositData object. Used as a protection against malformed input.|


### get_deposit_root

Query the current deposit root hash.


```solidity
function get_deposit_root() external view returns (bytes32);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes32`|The deposit root hash.|


### get_deposit_count

Query the current deposit count.


```solidity
function get_deposit_count() external view returns (bytes memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The deposit count encoded as a little endian 64-bit number.|


## Events
### DepositEvent
A processed deposit event.


```solidity
event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index);
```

