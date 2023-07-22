# ETHPOSDepositMock
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/mocks/ETHDepositMock.sol)

**Inherits:**
[IETHPOSDeposit](/src/contracts/interfaces/IETHPOSDeposit.sol/interface.IETHPOSDeposit.md)


## Functions
### deposit


```solidity
function deposit(
    bytes calldata pubkey,
    bytes calldata withdrawal_credentials,
    bytes calldata signature,
    bytes32 deposit_data_root
) external payable;
```

### get_deposit_root


```solidity
function get_deposit_root() external pure returns (bytes32);
```

### get_deposit_count

Query the current deposit count.


```solidity
function get_deposit_count() external pure returns (bytes memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The deposit count encoded as a little endian 64-bit number.|


