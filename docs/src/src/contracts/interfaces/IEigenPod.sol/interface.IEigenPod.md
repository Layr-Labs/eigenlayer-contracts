# IEigenPod
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/interfaces/IEigenPod.sol)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

The main functionalities are:
- creating new ETH validators with their withdrawal credentials pointed to this contract
- proving from beacon chain state roots that withdrawal credentials are pointed to this contract
- proving from beacon chain state roots the balances of ETH validators with their withdrawal credentials
pointed to this contract
- updating aggregate balances in the EigenPodManager
- withdrawing eth when withdrawals are initiated

*Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts*


## Functions
### REQUIRED_BALANCE_GWEI

The amount of eth, in gwei, that is restaked per validator


```solidity
function REQUIRED_BALANCE_GWEI() external view returns (uint64);
```

### REQUIRED_BALANCE_WEI

The amount of eth, in wei, that is restaked per validator


```solidity
function REQUIRED_BALANCE_WEI() external view returns (uint256);
```

### validatorStatus

this is a mapping of validator indices to a Validator struct containing pertinent info about the validator


```solidity
function validatorStatus(uint40 validatorIndex) external view returns (VALIDATOR_STATUS);
```

### restakedExecutionLayerGwei

the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from beaconchain but not EigenLayer),


```solidity
function restakedExecutionLayerGwei() external view returns (uint64);
```

### initialize

Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager


```solidity
function initialize(address owner) external;
```

### stake

Called by EigenPodManager when the owner wants to create another ETH validator.


```solidity
function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;
```

### withdrawRestakedBeaconChainETH

Transfers `amountWei` in ether from this contract to the specified `recipient` address

Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.

*Called during withdrawal or slashing.*

*Note that this function is marked as non-reentrant to prevent the recipient calling back into it*


```solidity
function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) external;
```

### eigenPodManager

The single EigenPodManager for EigenLayer


```solidity
function eigenPodManager() external view returns (IEigenPodManager);
```

### podOwner

The owner of this EigenPod


```solidity
function podOwner() external view returns (address);
```

### hasRestaked

an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.


```solidity
function hasRestaked() external view returns (bool);
```

### mostRecentWithdrawalBlockNumber

block number of the most recent withdrawal


```solidity
function mostRecentWithdrawalBlockNumber() external view returns (uint64);
```

### provenPartialWithdrawal

mapping that tracks proven partial withdrawals


```solidity
function provenPartialWithdrawal(uint40 validatorIndex, uint64 slot) external view returns (bool);
```

### verifyWithdrawalCredentialsAndBalance

This function verifies that the withdrawal credentials of the podOwner are pointed to
this contract. It also verifies the current (not effective) balance  of the validator.  It verifies the provided proof of the ETH validator against the beacon chain state
root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.


```solidity
function verifyWithdrawalCredentialsAndBalance(
    uint64 oracleBlockNumber,
    uint40 validatorIndex,
    BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory proofs,
    bytes32[] calldata validatorFields
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oracleBlockNumber`|`uint64`|is the Beacon Chain blockNumber whose state root the `proof` will be proven against.|
|`validatorIndex`|`uint40`|is the index of the validator being proven, refer to consensus specs|
|`proofs`|`BeaconChainProofs.ValidatorFieldsAndBalanceProofs`|is the bytes that prove the ETH validator's balance and withdrawal credentials against a beacon chain state root|
|`validatorFields`|`bytes32[]`|are the fields of the "Validator Container", refer to consensus specs for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator|


### verifyOvercommittedStake

This function records an overcommitment of stake to EigenLayer on behalf of a certain ETH validator.
If successful, the overcommitted balance is penalized (available for withdrawal whenever the pod's balance allows).
The ETH validator's shares in the enshrined beaconChainETH strategy are also removed from the StrategyManager and undelegated.

*For more details on the Beacon Chain spec, see: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator*


```solidity
function verifyOvercommittedStake(
    uint40 validatorIndex,
    BeaconChainProofs.ValidatorFieldsAndBalanceProofs calldata proofs,
    bytes32[] calldata validatorFields,
    uint256 beaconChainETHStrategyIndex,
    uint64 oracleBlockNumber
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`validatorIndex`|`uint40`|is the index of the validator being proven, refer to consensus specs|
|`proofs`|`BeaconChainProofs.ValidatorFieldsAndBalanceProofs`|is the proof of the validator's balance and validatorFields in the balance tree and the balanceRoot to prove for|
|`validatorFields`|`bytes32[]`|are the fields of the "Validator Container", refer to consensus specs|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the StrategyManager in case it must be removed from the list of the podOwners strategies|
|`oracleBlockNumber`|`uint64`|The oracleBlockNumber whose state root the `proof` will be proven against. Must be within `VERIFY_OVERCOMMITTED_WINDOW_BLOCKS` of the current block.|


### verifyAndProcessWithdrawal

This function records a full withdrawal on behalf of one of the Ethereum validators for this EigenPod


```solidity
function verifyAndProcessWithdrawal(
    BeaconChainProofs.WithdrawalProofs calldata withdrawalProofs,
    bytes calldata validatorFieldsProof,
    bytes32[] calldata validatorFields,
    bytes32[] calldata withdrawalFields,
    uint256 beaconChainETHStrategyIndex,
    uint64 oracleBlockNumber
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`withdrawalProofs`|`BeaconChainProofs.WithdrawalProofs`|is the information needed to check the veracity of the block number and withdrawal being proven|
|`validatorFieldsProof`|`bytes`|is the proof of the validator's fields in the validator tree|
|`validatorFields`|`bytes32[]`|are the fields of the validator being proven|
|`withdrawalFields`|`bytes32[]`|are the fields of the withdrawal being proven|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the EigenPodManager to the StrategyManager in case it must be removed from the podOwner's list of strategies|
|`oracleBlockNumber`|`uint64`||


### withdrawBeforeRestaking

Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false


```solidity
function withdrawBeforeRestaking() external;
```

## Structs
### PartialWithdrawalClaim

```solidity
struct PartialWithdrawalClaim {
    PARTIAL_WITHDRAWAL_CLAIM_STATUS status;
    uint32 creationBlockNumber;
    uint32 fraudproofPeriodEndBlockNumber;
    uint64 partialWithdrawalAmountGwei;
}
```

## Enums
### VALIDATOR_STATUS

```solidity
enum VALIDATOR_STATUS {
    INACTIVE,
    ACTIVE,
    OVERCOMMITTED,
    WITHDRAWN
}
```

### PARTIAL_WITHDRAWAL_CLAIM_STATUS

```solidity
enum PARTIAL_WITHDRAWAL_CLAIM_STATUS {
    REDEEMED,
    PENDING,
    FAILED
}
```

