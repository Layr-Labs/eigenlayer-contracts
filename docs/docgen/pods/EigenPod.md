# Solidity API

## EigenPod

The main functionalities are:
- creating new ETH validators with their withdrawal credentials pointed to this contract
- proving from beacon chain state roots that withdrawal credentials are pointed to this contract
- proving from beacon chain state roots the balances of ETH validators with their withdrawal credentials
  pointed to this contract
- updating aggregate balances in the EigenPodManager
- withdrawing eth when withdrawals are initiated

_Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
  to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts_

### GWEI_TO_WEI

```solidity
uint256 GWEI_TO_WEI
```

### VERIFY_OVERCOMMITTED_WINDOW_BLOCKS

```solidity
uint256 VERIFY_OVERCOMMITTED_WINDOW_BLOCKS
```

Maximum "staleness" of a Beacon Chain state root against which `verifyOvercommittedStake` may be proven. 7 days in blocks.

### ethPOS

```solidity
contract IETHPOSDeposit ethPOS
```

This is the beacon chain deposit contract

### delayedWithdrawalRouter

```solidity
contract IDelayedWithdrawalRouter delayedWithdrawalRouter
```

Contract used for withdrawal routing, to provide an extra "safety net" mechanism

### eigenPodManager

```solidity
contract IEigenPodManager eigenPodManager
```

The single EigenPodManager for EigenLayer

### REQUIRED_BALANCE_GWEI

```solidity
uint64 REQUIRED_BALANCE_GWEI
```

The amount of eth, in gwei, that is restaked per validator

### REQUIRED_BALANCE_WEI

```solidity
uint256 REQUIRED_BALANCE_WEI
```

The amount of eth, in wei, that is restaked per ETH validator into EigenLayer

### podOwner

```solidity
address podOwner
```

The owner of this EigenPod

### mostRecentWithdrawalBlockNumber

```solidity
uint64 mostRecentWithdrawalBlockNumber
```

The latest block number at which the pod owner withdrew the balance of the pod.

_This variable is only updated when the `withdraw` function is called, which can only occur before `hasRestaked` is set to true for this pod.
Proofs for this pod are only valid against Beacon Chain state roots corresponding to blocks after the stored `mostRecentWithdrawalBlockNumber`._

### restakedExecutionLayerGwei

```solidity
uint64 restakedExecutionLayerGwei
```

the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from the Beacon Chain but not from EigenLayer),

### hasRestaked

```solidity
bool hasRestaked
```

an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.

### validatorStatus

```solidity
mapping(uint40 => enum IEigenPod.VALIDATOR_STATUS) validatorStatus
```

this is a mapping of validator indices to a Validator struct containing pertinent info about the validator

### provenPartialWithdrawal

```solidity
mapping(uint40 => mapping(uint64 => bool)) provenPartialWithdrawal
```

This is a mapping of validatorIndex to withdrawalIndex to whether or not they have proven a withdrawal for that index

### EigenPodStaked

```solidity
event EigenPodStaked(bytes pubkey)
```

Emitted when an ETH validator stakes via this eigenPod

### ValidatorRestaked

```solidity
event ValidatorRestaked(uint40 validatorIndex)
```

Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod

### ValidatorOvercommitted

```solidity
event ValidatorOvercommitted(uint40 validatorIndex)
```

Emitted when an ETH validator is proven to have a balance less than `REQUIRED_BALANCE_GWEI` in the beacon chain

### FullWithdrawalRedeemed

```solidity
event FullWithdrawalRedeemed(uint40 validatorIndex, address recipient, uint64 withdrawalAmountGwei)
```

Emitted when an ETH validator is prove to have withdrawn from the beacon chain

### PartialWithdrawalRedeemed

```solidity
event PartialWithdrawalRedeemed(uint40 validatorIndex, address recipient, uint64 partialWithdrawalAmountGwei)
```

Emitted when a partial withdrawal claim is successfully redeemed

### RestakedBeaconChainETHWithdrawn

```solidity
event RestakedBeaconChainETHWithdrawn(address recipient, uint256 amount)
```

Emitted when restaked beacon chain ETH is withdrawn from the eigenPod.

### onlyEigenPodManager

```solidity
modifier onlyEigenPodManager()
```

### onlyEigenPodOwner

```solidity
modifier onlyEigenPodOwner()
```

### onlyNotFrozen

```solidity
modifier onlyNotFrozen()
```

### hasNeverRestaked

```solidity
modifier hasNeverRestaked()
```

### proofIsForValidBlockNumber

```solidity
modifier proofIsForValidBlockNumber(uint64 blockNumber)
```

Checks that `blockNumber` is strictly greater than the value stored in `mostRecentWithdrawalBlockNumber`

### onlyWhenNotPaused

```solidity
modifier onlyWhenNotPaused(uint8 index)
```

Based on 'Pausable' code, but uses the storage of the EigenPodManager instead of this contract. This construction
is necessary for enabling pausing all EigenPods at the same time (due to EigenPods being Beacon Proxies).
Modifier throws if the `indexed`th bit of `_paused` in the EigenPodManager is 1, i.e. if the `index`th pause switch is flipped.

### constructor

```solidity
constructor(contract IETHPOSDeposit _ethPOS, contract IDelayedWithdrawalRouter _delayedWithdrawalRouter, contract IEigenPodManager _eigenPodManager, uint256 _REQUIRED_BALANCE_WEI) public
```

### initialize

```solidity
function initialize(address _podOwner) external
```

Used to initialize the pointers to addresses crucial to the pod's functionality. Called on construction by the EigenPodManager.

### stake

```solidity
function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) external payable
```

Called by EigenPodManager when the owner wants to create another ETH validator.

### verifyWithdrawalCredentialsAndBalance

```solidity
function verifyWithdrawalCredentialsAndBalance(uint64 oracleBlockNumber, uint40 validatorIndex, struct BeaconChainProofs.ValidatorFieldsAndBalanceProofs proofs, bytes32[] validatorFields) external
```

This function verifies that the withdrawal credentials of the podOwner are pointed to
this contract. It also verifies the current (not effective) balance  of the validator.  It verifies the provided proof of the ETH validator against the beacon chain state
root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| oracleBlockNumber | uint64 | is the Beacon Chain blockNumber whose state root the `proof` will be proven against. |
| validatorIndex | uint40 | is the index of the validator being proven, refer to consensus specs |
| proofs | struct BeaconChainProofs.ValidatorFieldsAndBalanceProofs | is the bytes that prove the ETH validator's balance and withdrawal credentials against a beacon chain state root |
| validatorFields | bytes32[] | are the fields of the "Validator Container", refer to consensus specs  for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator |

### verifyOvercommittedStake

```solidity
function verifyOvercommittedStake(uint40 validatorIndex, struct BeaconChainProofs.ValidatorFieldsAndBalanceProofs proofs, bytes32[] validatorFields, uint256 beaconChainETHStrategyIndex, uint64 oracleBlockNumber) external
```

This function records an overcommitment of stake to EigenLayer on behalf of a certain ETH validator.
        If successful, the overcommitted balance is penalized (available for withdrawal whenever the pod's balance allows).
        The ETH validator's shares in the enshrined beaconChainETH strategy are also removed from the StrategyManager and undelegated.

_For more details on the Beacon Chain spec, see: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| validatorIndex | uint40 | is the index of the validator being proven, refer to consensus specs |
| proofs | struct BeaconChainProofs.ValidatorFieldsAndBalanceProofs | is the proof of the validator's balance and validatorFields in the balance tree and the balanceRoot to prove for |
| validatorFields | bytes32[] | are the fields of the "Validator Container", refer to consensus specs |
| beaconChainETHStrategyIndex | uint256 | is the index of the beaconChainETHStrategy for the pod owner for the callback to                                     the StrategyManager in case it must be removed from the list of the podOwners strategies |
| oracleBlockNumber | uint64 | The oracleBlockNumber whose state root the `proof` will be proven against.        Must be within `VERIFY_OVERCOMMITTED_WINDOW_BLOCKS` of the current block. |

### verifyAndProcessWithdrawal

```solidity
function verifyAndProcessWithdrawal(struct BeaconChainProofs.WithdrawalProofs withdrawalProofs, bytes validatorFieldsProof, bytes32[] validatorFields, bytes32[] withdrawalFields, uint256 beaconChainETHStrategyIndex, uint64 oracleBlockNumber) external
```

This function records a full withdrawal on behalf of one of the Ethereum validators for this EigenPod

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| withdrawalProofs | struct BeaconChainProofs.WithdrawalProofs | is the information needed to check the veracity of the block number and withdrawal being proven |
| validatorFieldsProof | bytes | is the information needed to check the veracity of the validator fields being proven |
| validatorFields | bytes32[] | are the fields of the validator being proven |
| withdrawalFields | bytes32[] | are the fields of the withdrawal being proven |
| beaconChainETHStrategyIndex | uint256 | is the index of the beaconChainETHStrategy for the pod owner for the callback to         the EigenPodManager to the StrategyManager in case it must be removed from the podOwner's list of strategies |
| oracleBlockNumber | uint64 |  |

### _processFullWithdrawal

```solidity
function _processFullWithdrawal(uint64 withdrawalAmountGwei, uint40 validatorIndex, uint256 beaconChainETHStrategyIndex, address recipient, enum IEigenPod.VALIDATOR_STATUS status) internal
```

### _processPartialWithdrawal

```solidity
function _processPartialWithdrawal(uint64 withdrawalHappenedSlot, uint64 partialWithdrawalAmountGwei, uint40 validatorIndex, address recipient) internal
```

### withdrawRestakedBeaconChainETH

```solidity
function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) external
```

Transfers `amountWei` in ether from this contract to the specified `recipient` address
Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.

_Called during withdrawal or slashing.
Note that this function is marked as non-reentrant to prevent the recipient calling back into it_

### withdrawBeforeRestaking

```solidity
function withdrawBeforeRestaking() external
```

Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false

### _podWithdrawalCredentials

```solidity
function _podWithdrawalCredentials() internal view returns (bytes)
```

### _sendETH

```solidity
function _sendETH(address recipient, uint256 amountWei) internal
```

### __gap

```solidity
uint256[46] __gap
```

_This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps_

