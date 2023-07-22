# EigenPodTests
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/EigenPod.t.sol)

**Inherits:**
[ProofParsing](/src/test/utils/ProofParsing.sol/contract.ProofParsing.md), [EigenPodPausingConstants](/src/contracts/pods/EigenPodPausingConstants.sol/abstract.EigenPodPausingConstants.md)


## State Variables
### GWEI_TO_WEI

```solidity
uint256 internal constant GWEI_TO_WEI = 1e9;
```


### pubkey

```solidity
bytes pubkey = hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
```


### validatorIndex0

```solidity
uint40 validatorIndex0 = 0;
```


### validatorIndex1

```solidity
uint40 validatorIndex1 = 1;
```


### validatorTreeRoot

```solidity
bytes32 validatorTreeRoot;
```


### validatorRoot

```solidity
bytes32 validatorRoot;
```


### podOwner

```solidity
address podOwner = address(42000094993494);
```


### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### delegation

```solidity
DelegationManager public delegation;
```


### strategyManager

```solidity
IStrategyManager public strategyManager;
```


### slasher

```solidity
Slasher public slasher;
```


### pauserReg

```solidity
PauserRegistry public pauserReg;
```


### eigenLayerProxyAdmin

```solidity
ProxyAdmin public eigenLayerProxyAdmin;
```


### blsPkCompendium

```solidity
IBLSPublicKeyCompendium public blsPkCompendium;
```


### eigenPodManager

```solidity
IEigenPodManager public eigenPodManager;
```


### podImplementation

```solidity
IEigenPod public podImplementation;
```


### delayedWithdrawalRouter

```solidity
IDelayedWithdrawalRouter public delayedWithdrawalRouter;
```


### ethPOSDeposit

```solidity
IETHPOSDeposit public ethPOSDeposit;
```


### eigenPodBeacon

```solidity
IBeacon public eigenPodBeacon;
```


### beaconChainOracle

```solidity
IBeaconChainOracleMock public beaconChainOracle;
```


### generalReg1

```solidity
MiddlewareRegistryMock public generalReg1;
```


### generalServiceManager1

```solidity
ServiceManagerMock public generalServiceManager1;
```


### slashingContracts

```solidity
address[] public slashingContracts;
```


### pauser

```solidity
address pauser = address(69);
```


### unpauser

```solidity
address unpauser = address(489);
```


### podManagerAddress

```solidity
address podManagerAddress = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;
```


### podAddress

```solidity
address podAddress = address(123);
```


### stakeAmount

```solidity
uint256 stakeAmount = 32e18;
```


### fuzzedAddressMapping

```solidity
mapping(address => bool) fuzzedAddressMapping;
```


### signature

```solidity
bytes signature;
```


### depositDataRoot

```solidity
bytes32 depositDataRoot;
```


### withdrawalFields

```solidity
bytes32[] withdrawalFields;
```


### validatorFields

```solidity
bytes32[] validatorFields;
```


### WITHDRAWAL_DELAY_BLOCKS

```solidity
uint32 WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
```


### REQUIRED_BALANCE_WEI

```solidity
uint256 REQUIRED_BALANCE_WEI = 31 ether;
```


## Functions
### fuzzedAddress


```solidity
modifier fuzzedAddress(address addr) virtual;
```

### setUp


```solidity
function setUp() public;
```

### testStaking

First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.


```solidity
function testStaking() public;
```

### testWithdrawBeforeRestaking


```solidity
function testWithdrawBeforeRestaking() public;
```

### testWithdrawBeforeRestakingAfterRestaking


```solidity
function testWithdrawBeforeRestakingAfterRestaking() public;
```

### testWithdrawFromPod


```solidity
function testWithdrawFromPod() public;
```

### testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials


```solidity
function testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials() public;
```

### testFullWithdrawalProof


```solidity
function testFullWithdrawalProof() public;
```

### testFullWithdrawalFlow

This test is to ensure the full withdrawal flow works


```solidity
function testFullWithdrawalFlow() public returns (IEigenPod);
```

### testPartialWithdrawalFlow

This test is to ensure that the partial withdrawal flow works correctly


```solidity
function testPartialWithdrawalFlow() public returns (IEigenPod);
```

### testProvingMultipleWithdrawalsForSameSlot

verifies that multiple partial withdrawals can be made before a full withdrawal


```solidity
function testProvingMultipleWithdrawalsForSameSlot() public;
```

### testDoubleFullWithdrawal

verifies that multiple full withdrawals for a single validator fail


```solidity
function testDoubleFullWithdrawal() public;
```

### testDeployAndVerifyNewEigenPod


```solidity
function testDeployAndVerifyNewEigenPod() public returns (IEigenPod);
```

### testUpdateSlashedBeaconBalance


```solidity
function testUpdateSlashedBeaconBalance() public;
```

### testDeployNewEigenPodWithWrongWithdrawalCreds


```solidity
function testDeployNewEigenPodWithWrongWithdrawalCreds(address wrongWithdrawalAddress) public;
```

### testDeployNewEigenPodWithActiveValidator


```solidity
function testDeployNewEigenPodWithActiveValidator() public;
```

### testVerifyWithdrawalCredentialsWithInadequateBalance


```solidity
function testVerifyWithdrawalCredentialsWithInadequateBalance() public;
```

### testProveOverComittedStakeOnWithdrawnValidator


```solidity
function testProveOverComittedStakeOnWithdrawnValidator() public;
```

### getBeaconChainETHShares


```solidity
function getBeaconChainETHShares(address staker) internal view returns (uint256);
```

### testProveSingleWithdrawalCredential


```solidity
function testProveSingleWithdrawalCredential() public;
```

### testProveOverCommittedBalance


```solidity
function testProveOverCommittedBalance() public;
```

### testDeployingEigenPodRevertsWhenPaused


```solidity
function testDeployingEigenPodRevertsWhenPaused() external;
```

### testCreatePodWhenPaused


```solidity
function testCreatePodWhenPaused() external;
```

### testStakeOnEigenPodFromNonPodManagerAddress


```solidity
function testStakeOnEigenPodFromNonPodManagerAddress(address nonPodManager) external fuzzedAddress(nonPodManager);
```

### testCallWithdrawBeforeRestakingFromNonOwner


```solidity
function testCallWithdrawBeforeRestakingFromNonOwner(address nonPodOwner) external fuzzedAddress(nonPodOwner);
```

### testWithdrawRestakedBeaconChainETHRevertsWhenPaused


```solidity
function testWithdrawRestakedBeaconChainETHRevertsWhenPaused() external;
```

### testVerifyCorrectWithdrawalCredentialsRevertsWhenPaused


```solidity
function testVerifyCorrectWithdrawalCredentialsRevertsWhenPaused() external;
```

### testVerifyOvercommittedStakeRevertsWhenPaused


```solidity
function testVerifyOvercommittedStakeRevertsWhenPaused() external;
```

### _proveOverCommittedStake


```solidity
function _proveOverCommittedStake(IEigenPod newPod) internal;
```

### testStake


```solidity
function testStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot) public;
```

### testVerifyInclusionSha256FailsForEmptyProof

Test that the Merkle proof verification fails when the proof length is 0


```solidity
function testVerifyInclusionSha256FailsForEmptyProof(bytes32 root, bytes32 leaf, uint256 index) public;
```

### testVerifyInclusionSha256FailsForNonMultipleOf32ProofLength

Test that the Merkle proof verification fails when the proof length is not a multple of 32


```solidity
function testVerifyInclusionSha256FailsForNonMultipleOf32ProofLength(
    bytes32 root,
    bytes32 leaf,
    uint256 index,
    bytes memory proof
) public;
```

### testVerifyInclusionKeccakFailsForEmptyProof

Test that the Merkle proof verification fails when the proof length is empty


```solidity
function testVerifyInclusionKeccakFailsForEmptyProof(bytes32 root, bytes32 leaf, uint256 index) public;
```

### testVerifyInclusionKeccakFailsForNonMultipleOf32ProofLength

Test that the Merkle proof verification fails when the proof length is not a multiple of 32


```solidity
function testVerifyInclusionKeccakFailsForNonMultipleOf32ProofLength(
    bytes32 root,
    bytes32 leaf,
    uint256 index,
    bytes memory proof
) public;
```

### test_incrementNumPodsOnStake


```solidity
function test_incrementNumPodsOnStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot)
    public;
```

### test_maxPodsEnforcementOnStake


```solidity
function test_maxPodsEnforcementOnStake(bytes calldata _pubkey, bytes calldata _signature, bytes32 _depositDataRoot)
    public;
```

### test_incrementNumPodsOnCreatePod


```solidity
function test_incrementNumPodsOnCreatePod() public;
```

### test_createPodTwiceFails


```solidity
function test_createPodTwiceFails() public;
```

### test_maxPodsEnforcementOnCreatePod


```solidity
function test_maxPodsEnforcementOnCreatePod() public;
```

### test_setMaxPods


```solidity
function test_setMaxPods(uint256 newValue) public;
```

### test_setMaxPods_RevertsWhenNotCalledByUnpauser


```solidity
function test_setMaxPods_RevertsWhenNotCalledByUnpauser(address notUnpauser) public fuzzedAddress(notUnpauser);
```

### _testRegisterAsOperator


```solidity
function _testRegisterAsOperator(address sender, IDelegationTerms dt) internal;
```

### _testDelegateToOperator


```solidity
function _testDelegateToOperator(address sender, address operator) internal;
```

### _testDelegation


```solidity
function _testDelegation(address operator, address staker) internal;
```

### _testDeployAndVerifyNewEigenPod


```solidity
function _testDeployAndVerifyNewEigenPod(address _podOwner, bytes memory _signature, bytes32 _depositDataRoot)
    internal
    returns (IEigenPod);
```

### _testQueueWithdrawal


```solidity
function _testQueueWithdrawal(
    address depositor,
    uint256[] memory strategyIndexes,
    IStrategy[] memory strategyArray,
    uint256[] memory shareAmounts,
    bool undelegateIfPossible
) internal returns (bytes32);
```

### _getLatestDelayedWithdrawalAmount


```solidity
function _getLatestDelayedWithdrawalAmount(address recipient) internal view returns (uint256);
```

### _getValidatorFieldsAndBalanceProof


```solidity
function _getValidatorFieldsAndBalanceProof()
    internal
    returns (BeaconChainProofs.ValidatorFieldsAndBalanceProofs memory);
```

### _getWithdrawalProof

this function just generates a valid proof so that we can test other functionalities of the withdrawal flow


```solidity
function _getWithdrawalProof() internal returns (BeaconChainProofs.WithdrawalProofs memory);
```

### _getValidatorFieldsProof


```solidity
function _getValidatorFieldsProof() internal returns (BeaconChainProofs.ValidatorFieldsProof memory);
```

## Events
### BeaconOracleUpdated
Emitted to notify the update of the beaconChainOracle address


```solidity
event BeaconOracleUpdated(address indexed newOracleAddress);
```

### PodDeployed
Emitted to notify the deployment of an EigenPod


```solidity
event PodDeployed(address indexed eigenPod, address indexed podOwner);
```

### BeaconChainETHDeposited
Emitted to notify a deposit of beacon chain ETH recorded in the strategy manager


```solidity
event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);
```

### MaxPodsUpdated
Emitted when `maxPods` value is updated from `previousValue` to `newValue`


```solidity
event MaxPodsUpdated(uint256 previousValue, uint256 newValue);
```

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

### WithdrawalDelayBlocksSet
Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.


```solidity
event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);
```

### DelayedWithdrawalCreated
event for delayedWithdrawal creation


```solidity
event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index);
```

### DelayedWithdrawalsClaimed
event for the claiming of delayedWithdrawals


```solidity
event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted);
```

