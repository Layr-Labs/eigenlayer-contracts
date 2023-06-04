# EigenPod
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/pods/EigenPod.sol)

**Inherits:**
[IEigenPod](/docs/docgen/src/src/contracts/interfaces/IEigenPod.sol/interface.IEigenPod.md), Initializable, ReentrancyGuardUpgradeable, [EigenPodPausingConstants](/docs/docgen/src/src/contracts/pods/EigenPodPausingConstants.sol/abstract.EigenPodPausingConstants.md)

**Author:**
Layr Labs, Inc.

The main functionalities are:
- creating new ETH validators with their withdrawal credentials pointed to this contract
- proving from beacon chain state roots that withdrawal credentials are pointed to this contract
- proving from beacon chain state roots the balances of ETH validators with their withdrawal credentials
pointed to this contract
- updating aggregate balances in the EigenPodManager
- withdrawing eth when withdrawals are initiated

*Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts*


## State Variables
### GWEI_TO_WEI

```solidity
uint256 internal constant GWEI_TO_WEI = 1e9;
```


### VERIFY_OVERCOMMITTED_WINDOW_BLOCKS
Maximum "staleness" of a Beacon Chain state root against which `verifyOvercommittedStake` may be proven. 7 days in blocks.


```solidity
uint256 internal constant VERIFY_OVERCOMMITTED_WINDOW_BLOCKS = 50400;
```


### ethPOS
This is the beacon chain deposit contract


```solidity
IETHPOSDeposit public immutable ethPOS;
```


### delayedWithdrawalRouter
Contract used for withdrawal routing, to provide an extra "safety net" mechanism


```solidity
IDelayedWithdrawalRouter public immutable delayedWithdrawalRouter;
```


### eigenPodManager
The single EigenPodManager for EigenLayer


```solidity
IEigenPodManager public immutable eigenPodManager;
```


### REQUIRED_BALANCE_GWEI
The amount of eth, in gwei, that is restaked per validator


```solidity
uint64 public immutable REQUIRED_BALANCE_GWEI;
```


### REQUIRED_BALANCE_WEI
The amount of eth, in wei, that is restaked per ETH validator into EigenLayer


```solidity
uint256 public immutable REQUIRED_BALANCE_WEI;
```


### podOwner
The owner of this EigenPod


```solidity
address public podOwner;
```


### mostRecentWithdrawalBlockNumber
The latest block number at which the pod owner withdrew the balance of the pod.

*This variable is only updated when the `withdraw` function is called, which can only occur before `hasRestaked` is set to true for this pod.
Proofs for this pod are only valid against Beacon Chain state roots corresponding to blocks after the stored `mostRecentWithdrawalBlockNumber`.*


```solidity
uint64 public mostRecentWithdrawalBlockNumber;
```


### restakedExecutionLayerGwei
the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from the Beacon Chain but not from EigenLayer),


```solidity
uint64 public restakedExecutionLayerGwei;
```


### hasRestaked
an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.


```solidity
bool public hasRestaked;
```


### validatorStatus
this is a mapping of validator indices to a Validator struct containing pertinent info about the validator


```solidity
mapping(uint40 => VALIDATOR_STATUS) public validatorStatus;
```


### provenPartialWithdrawal
This is a mapping of validatorIndex to withdrawalIndex to whether or not they have proven a withdrawal for that index


```solidity
mapping(uint40 => mapping(uint64 => bool)) public provenPartialWithdrawal;
```


### __gap
*This empty reserved space is put in place to allow future versions to add new
variables without shifting down storage in the inheritance chain.
See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps*


```solidity
uint256[46] private __gap;
```


## Functions
### onlyEigenPodManager


```solidity
modifier onlyEigenPodManager();
```

### onlyEigenPodOwner


```solidity
modifier onlyEigenPodOwner();
```

### onlyNotFrozen


```solidity
modifier onlyNotFrozen();
```

### hasNeverRestaked


```solidity
modifier hasNeverRestaked();
```

### proofIsForValidBlockNumber

Checks that `blockNumber` is strictly greater than the value stored in `mostRecentWithdrawalBlockNumber`


```solidity
modifier proofIsForValidBlockNumber(uint64 blockNumber);
```

### onlyWhenNotPaused

Based on 'Pausable' code, but uses the storage of the EigenPodManager instead of this contract. This construction
is necessary for enabling pausing all EigenPods at the same time (due to EigenPods being Beacon Proxies).
Modifier throws if the `indexed`th bit of `_paused` in the EigenPodManager is 1, i.e. if the `index`th pause switch is flipped.


```solidity
modifier onlyWhenNotPaused(uint8 index);
```

### constructor


```solidity
constructor(
    IETHPOSDeposit _ethPOS,
    IDelayedWithdrawalRouter _delayedWithdrawalRouter,
    IEigenPodManager _eigenPodManager,
    uint256 _REQUIRED_BALANCE_WEI
);
```

### initialize

Used to initialize the pointers to addresses crucial to the pod's functionality. Called on construction by the EigenPodManager.


```solidity
function initialize(address _podOwner) external initializer;
```

### stake

Called by EigenPodManager when the owner wants to create another ETH validator.


```solidity
function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot)
    external
    payable
    onlyEigenPodManager;
```

### verifyWithdrawalCredentialsAndBalance

This function verifies that the withdrawal credentials of the podOwner are pointed to
this contract. It also verifies the current (not effective) balance  of the validator.  It verifies the provided proof of the ETH validator against the beacon chain state
root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.


```solidity
function verifyWithdrawalCredentialsAndBalance(
    uint64 oracleBlockNumber,
    uint40 validatorIndex,
    BeaconChainProofs.ValidatorFieldsAndBalanceProofs calldata proofs,
    bytes32[] calldata validatorFields
) external onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS) proofIsForValidBlockNumber(oracleBlockNumber);
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
) external onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_OVERCOMMITTED);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`validatorIndex`|`uint40`|is the index of the validator being proven, refer to consensus specs|
|`proofs`|`BeaconChainProofs.ValidatorFieldsAndBalanceProofs`|is the proof of the validator's balance and validatorFields in the balance tree and the balanceRoot to prove for|
|`validatorFields`|`bytes32[]`|are the fields of the "Validator Container", refer to consensus specs|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the StrategyManager in case it must be removed from the list of the podOwner's strategies|
|`oracleBlockNumber`|`uint64`|The oracleBlockNumber whose state root the `proof` will be proven against. Must be within `VERIFY_OVERCOMMITTED_WINDOW_BLOCKS` of the current block.|


### verifyAndProcessWithdrawal

If validator's balance is zero, then either they have fully withdrawn or they have been slashed down zero.
If the validator *has* been slashed, then this function can proceed. If they have *not* been slashed, then
the `verifyAndProcessWithdrawal` function should be called instead.

This function records a full withdrawal on behalf of one of the Ethereum validators for this EigenPod


```solidity
function verifyAndProcessWithdrawal(
    BeaconChainProofs.WithdrawalProofs calldata withdrawalProofs,
    bytes calldata validatorFieldsProof,
    bytes32[] calldata validatorFields,
    bytes32[] calldata withdrawalFields,
    uint256 beaconChainETHStrategyIndex,
    uint64 oracleBlockNumber
)
    external
    onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_WITHDRAWAL)
    onlyNotFrozen
    proofIsForValidBlockNumber(Endian.fromLittleEndianUint64(withdrawalProofs.blockNumberRoot));
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`withdrawalProofs`|`BeaconChainProofs.WithdrawalProofs`|is the information needed to check the veracity of the block number and withdrawal being proven|
|`validatorFieldsProof`|`bytes`|is the information needed to check the veracity of the validator fields being proven|
|`validatorFields`|`bytes32[]`|are the fields of the validator being proven|
|`withdrawalFields`|`bytes32[]`|are the fields of the withdrawal being proven|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy for the pod owner for the callback to the EigenPodManager to the StrategyManager in case it must be removed from the podOwner's list of strategies|
|`oracleBlockNumber`|`uint64`|is the Beacon Chain blockNumber whose state root the `proof` will be proven against.|


### _processFullWithdrawal

If the validator status is inactive, then withdrawal credentials were never verified for the validator,
and thus we cannot know that the validator is related to this EigenPod at all!
if the validator's withdrawable epoch is less than or equal to the slot's epoch, then the validator has fully withdrawn because
a full withdrawal is only processable after the withdrawable epoch has passed.


```solidity
function _processFullWithdrawal(
    uint64 withdrawalAmountGwei,
    uint40 validatorIndex,
    uint256 beaconChainETHStrategyIndex,
    address recipient,
    VALIDATOR_STATUS status
) internal;
```

### _processPartialWithdrawal

since in `verifyOvercommittedStake` the podOwner's beaconChainETH shares are decremented by `REQUIRED_BALANCE_WEI`, we must reverse the process here,
in order to allow the podOwner to complete their withdrawal through EigenLayer's normal withdrawal process
since in `verifyOvercommittedStake` the podOwner's beaconChainETH shares are decremented by `REQUIRED_BALANCE_WEI`, we must reverse the process here,
in order to allow the podOwner to complete their withdrawal through EigenLayer's normal withdrawal process


```solidity
function _processPartialWithdrawal(
    uint64 withdrawalHappenedSlot,
    uint64 partialWithdrawalAmountGwei,
    uint40 validatorIndex,
    address recipient
) internal;
```

### withdrawRestakedBeaconChainETH

Transfers `amountWei` in ether from this contract to the specified `recipient` address

Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.

*Called during withdrawal or slashing.*


```solidity
function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) external onlyEigenPodManager;
```

### withdrawBeforeRestaking

Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false


```solidity
function withdrawBeforeRestaking() external onlyEigenPodOwner hasNeverRestaked;
```

### _podWithdrawalCredentials


```solidity
function _podWithdrawalCredentials() internal view returns (bytes memory);
```

### _sendETH


```solidity
function _sendETH(address recipient, uint256 amountWei) internal;
```

## Events
### EigenPodStaked
Emitted when an ETH validator stakes via this eigenPod


```solidity
event EigenPodStaked(bytes pubkey);
```

### ValidatorRestaked
Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod


```solidity
event ValidatorRestaked(uint40 validatorIndex);
```

### ValidatorOvercommitted
Emitted when an ETH validator is proven to have a balance less than `REQUIRED_BALANCE_GWEI` in the beacon chain


```solidity
event ValidatorOvercommitted(uint40 validatorIndex);
```

### FullWithdrawalRedeemed
Emitted when an ETH validator is prove to have withdrawn from the beacon chain


```solidity
event FullWithdrawalRedeemed(uint40 validatorIndex, address indexed recipient, uint64 withdrawalAmountGwei);
```

### PartialWithdrawalRedeemed
Emitted when a partial withdrawal claim is successfully redeemed


```solidity
event PartialWithdrawalRedeemed(uint40 validatorIndex, address indexed recipient, uint64 partialWithdrawalAmountGwei);
```

### RestakedBeaconChainETHWithdrawn
Emitted when restaked beacon chain ETH is withdrawn from the eigenPod.


```solidity
event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount);
```

