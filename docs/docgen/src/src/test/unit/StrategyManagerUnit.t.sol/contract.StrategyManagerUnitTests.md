# StrategyManagerUnitTests
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/test/unit/StrategyManagerUnit.t.sol)

**Inherits:**
Test, [Utils](/docs/docgen/src/src/test/unit/Utils.sol/contract.Utils.md)


## State Variables
### cheats

```solidity
Vm cheats = Vm(HEVM_ADDRESS);
```


### REQUIRED_BALANCE_WEI

```solidity
uint256 public REQUIRED_BALANCE_WEI = 31 ether;
```


### proxyAdmin

```solidity
ProxyAdmin public proxyAdmin;
```


### pauserRegistry

```solidity
PauserRegistry public pauserRegistry;
```


### strategyManagerImplementation

```solidity
StrategyManager public strategyManagerImplementation;
```


### strategyManager

```solidity
StrategyManager public strategyManager;
```


### delegationMock

```solidity
DelegationMock public delegationMock;
```


### slasherMock

```solidity
SlasherMock public slasherMock;
```


### eigenPodManagerMock

```solidity
EigenPodManagerMock public eigenPodManagerMock;
```


### dummyStrat

```solidity
StrategyBase public dummyStrat;
```


### dummyStrat2

```solidity
StrategyBase public dummyStrat2;
```


### beaconChainETHStrategy

```solidity
IStrategy public beaconChainETHStrategy;
```


### dummyToken

```solidity
IERC20 public dummyToken;
```


### reenterer

```solidity
Reenterer public reenterer;
```


### GWEI_TO_WEI

```solidity
uint256 GWEI_TO_WEI = 1e9;
```


### pauser

```solidity
address public pauser = address(555);
```


### unpauser

```solidity
address public unpauser = address(999);
```


### initialOwner

```solidity
address initialOwner = address(this);
```


### emptyUintArray

```solidity
uint256[] public emptyUintArray;
```


### _tempStrategyStorage

```solidity
IStrategy public _tempStrategyStorage;
```


### _tempStakerStorage

```solidity
address public _tempStakerStorage;
```


### privateKey

```solidity
uint256 public privateKey = 111111;
```


### addressIsExcludedFromFuzzedInputs

```solidity
mapping(address => bool) public addressIsExcludedFromFuzzedInputs;
```


## Functions
### filterFuzzedAddressInputs


```solidity
modifier filterFuzzedAddressInputs(address fuzzedAddress);
```

### setUp


```solidity
function setUp() public virtual;
```

### testCannotReinitialize


```solidity
function testCannotReinitialize() public;
```

### testDepositBeaconChainETHSuccessfully


```solidity
function testDepositBeaconChainETHSuccessfully(address staker, uint256 amount)
    public
    filterFuzzedAddressInputs(staker);
```

### testDepositBeaconChainETHFailsWhenNotCalledByEigenPodManager


```solidity
function testDepositBeaconChainETHFailsWhenNotCalledByEigenPodManager(address improperCaller)
    public
    filterFuzzedAddressInputs(improperCaller);
```

### testDepositBeaconChainETHFailsWhenDepositsPaused


```solidity
function testDepositBeaconChainETHFailsWhenDepositsPaused() public;
```

### testDepositBeaconChainETHFailsWhenStakerFrozen


```solidity
function testDepositBeaconChainETHFailsWhenStakerFrozen() public;
```

### testDepositBeaconChainETHFailsWhenReentering


```solidity
function testDepositBeaconChainETHFailsWhenReentering() public;
```

### testRecordOvercommittedBeaconChainETHSuccessfully


```solidity
function testRecordOvercommittedBeaconChainETHSuccessfully(uint256 amount_1, uint256 amount_2) public;
```

### testRecordOvercommittedBeaconChainETHFailsWhenNotCalledByEigenPodManager


```solidity
function testRecordOvercommittedBeaconChainETHFailsWhenNotCalledByEigenPodManager(address improperCaller)
    public
    filterFuzzedAddressInputs(improperCaller);
```

### testRecordOvercommittedBeaconChainETHFailsWhenReentering


```solidity
function testRecordOvercommittedBeaconChainETHFailsWhenReentering() public;
```

### testDepositIntoStrategySuccessfully


```solidity
function testDepositIntoStrategySuccessfully(address staker, uint256 amount) public filterFuzzedAddressInputs(staker);
```

### testDepositIntoStrategySuccessfullyTwice


```solidity
function testDepositIntoStrategySuccessfullyTwice() public;
```

### testDepositIntoStrategyFailsWhenDepositsPaused


```solidity
function testDepositIntoStrategyFailsWhenDepositsPaused() public;
```

### testDepositIntoStrategyFailsWhenStakerFrozen


```solidity
function testDepositIntoStrategyFailsWhenStakerFrozen() public;
```

### testDepositIntoStrategyFailsWhenReentering


```solidity
function testDepositIntoStrategyFailsWhenReentering() public;
```

### testDepositIntoStrategyWithSignatureSuccessfully


```solidity
function testDepositIntoStrategyWithSignatureSuccessfully(uint256 amount, uint256 expiry) public;
```

### testDepositIntoStrategyWithSignatureReplay


```solidity
function testDepositIntoStrategyWithSignatureReplay(uint256 amount, uint256 expiry) public;
```

### testDepositIntoStrategyWithSignature_WithContractWallet_Successfully


```solidity
function testDepositIntoStrategyWithSignature_WithContractWallet_Successfully(uint256 amount, uint256 expiry) public;
```

### testDepositIntoStrategyWithSignature_WithContractWallet_BadSignature


```solidity
function testDepositIntoStrategyWithSignature_WithContractWallet_BadSignature(uint256 amount) public;
```

### testDepositIntoStrategyWithSignature_WithContractWallet_NonconformingWallet


```solidity
function testDepositIntoStrategyWithSignature_WithContractWallet_NonconformingWallet(
    uint256 amount,
    uint8 v,
    bytes32 r,
    bytes32 s
) public;
```

### testDepositIntoStrategyWithSignatureFailsWhenDepositsPaused


```solidity
function testDepositIntoStrategyWithSignatureFailsWhenDepositsPaused() public;
```

### testDepositIntoStrategyWithSignatureFailsWhenStakerFrozen


```solidity
function testDepositIntoStrategyWithSignatureFailsWhenStakerFrozen() public;
```

### testDepositIntoStrategyWithSignatureFailsWhenReentering


```solidity
function testDepositIntoStrategyWithSignatureFailsWhenReentering() public;
```

### testDepositIntoStrategyWithSignatureFailsWhenSignatureExpired


```solidity
function testDepositIntoStrategyWithSignatureFailsWhenSignatureExpired() public;
```

### testDepositIntoStrategyWithSignatureFailsWhenSignatureInvalid


```solidity
function testDepositIntoStrategyWithSignatureFailsWhenSignatureInvalid() public;
```

### testUndelegate


```solidity
function testUndelegate() public;
```

### testUndelegateRevertsWithActiveDeposits


```solidity
function testUndelegateRevertsWithActiveDeposits() public;
```

### testQueueWithdrawalBeaconChainETHToSelf


```solidity
function testQueueWithdrawalBeaconChainETHToSelf(uint128 amountGwei)
    public
    returns (IStrategyManager.QueuedWithdrawal memory, bytes32);
```

### testQueueWithdrawalBeaconChainETHToDifferentAddress


```solidity
function testQueueWithdrawalBeaconChainETHToDifferentAddress(address withdrawer)
    external
    filterFuzzedAddressInputs(withdrawer);
```

### testQueueWithdrawalMultipleStrategiesWithBeaconChain


```solidity
function testQueueWithdrawalMultipleStrategiesWithBeaconChain() external;
```

### testQueueWithdrawalBeaconChainEthNonWholeAmountGwei


```solidity
function testQueueWithdrawalBeaconChainEthNonWholeAmountGwei(uint256 nonWholeAmount) external;
```

### testQueueWithdrawalMismatchedIndexAndStrategyArrayLength


```solidity
function testQueueWithdrawalMismatchedIndexAndStrategyArrayLength() external;
```

### testQueueWithdrawalWithZeroAddress


```solidity
function testQueueWithdrawalWithZeroAddress() external;
```

### testQueueWithdrawalWithFrozenAddress


```solidity
function testQueueWithdrawalWithFrozenAddress(address frozenAddress)
    external
    filterFuzzedAddressInputs(frozenAddress);
```

### testQueueWithdrawal_ToSelf_NotBeaconChainETH


```solidity
function testQueueWithdrawal_ToSelf_NotBeaconChainETH(
    uint256 depositAmount,
    uint256 withdrawalAmount,
    bool undelegateIfPossible
) public returns (IStrategyManager.QueuedWithdrawal memory, IERC20[] memory, bytes32);
```

### testQueueWithdrawal_ToSelf_NotBeaconChainETHTwoStrategies


```solidity
function testQueueWithdrawal_ToSelf_NotBeaconChainETHTwoStrategies(
    uint256 depositAmount,
    uint256 withdrawalAmount,
    bool undelegateIfPossible
) public returns (IStrategyManager.QueuedWithdrawal memory, IERC20[] memory, bytes32);
```

### testQueueWithdrawal_ToDifferentAddress_NotBeaconChainETH


```solidity
function testQueueWithdrawal_ToDifferentAddress_NotBeaconChainETH(address withdrawer, uint256 amount)
    external
    filterFuzzedAddressInputs(withdrawer);
```

### testQueueWithdrawal_WithdrawEverything_DontUndelegate


```solidity
function testQueueWithdrawal_WithdrawEverything_DontUndelegate(uint256 amount) external;
```

### testQueueWithdrawal_WithdrawEverything_DoUndelegate


```solidity
function testQueueWithdrawal_WithdrawEverything_DoUndelegate(uint256 amount) external;
```

### testQueueWithdrawal_DontWithdrawEverything_MarkUndelegateIfPossibleAsTrue


```solidity
function testQueueWithdrawal_DontWithdrawEverything_MarkUndelegateIfPossibleAsTrue(uint128 amount) external;
```

### testQueueWithdrawalFailsWhenStakerFrozen


```solidity
function testQueueWithdrawalFailsWhenStakerFrozen() public;
```

### testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedFalse


```solidity
function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedFalse() external;
```

### testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_NotWithdrawingBeaconChainETH


```solidity
function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_NotWithdrawingBeaconChainETH() external;
```

### testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_WithdrawingBeaconChainETH


```solidity
function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_WithdrawingBeaconChainETH() external;
```

### testCompleteQueuedWithdrawalFailsWhenWithdrawalsPaused


```solidity
function testCompleteQueuedWithdrawalFailsWhenWithdrawalsPaused() external;
```

### testCompleteQueuedWithdrawalFailsWhenDelegatedAddressFrozen


```solidity
function testCompleteQueuedWithdrawalFailsWhenDelegatedAddressFrozen() external;
```

### testCompleteQueuedWithdrawalFailsWhenAttemptingReentrancy


```solidity
function testCompleteQueuedWithdrawalFailsWhenAttemptingReentrancy() external;
```

### testCompleteQueuedWithdrawalFailsWhenWithdrawalDoesNotExist


```solidity
function testCompleteQueuedWithdrawalFailsWhenWithdrawalDoesNotExist() external;
```

### testCompleteQueuedWithdrawalFailsWhenCanWithdrawReturnsFalse


```solidity
function testCompleteQueuedWithdrawalFailsWhenCanWithdrawReturnsFalse() external;
```

### testUndelegateWithFrozenStaker


```solidity
function testUndelegateWithFrozenStaker() public;
```

### testCompleteQueuedWithdrawalFailsWhenNotCallingFromWithdrawerAddress


```solidity
function testCompleteQueuedWithdrawalFailsWhenNotCallingFromWithdrawerAddress() external;
```

### testCompleteQueuedWithdrawalFailsWhenTryingToCompleteSameWithdrawal2X


```solidity
function testCompleteQueuedWithdrawalFailsWhenTryingToCompleteSameWithdrawal2X() external;
```

### testCompleteQueuedWithdrawalFailsWhenWithdrawalDelayBlocksHasNotPassed


```solidity
function testCompleteQueuedWithdrawalFailsWhenWithdrawalDelayBlocksHasNotPassed() external;
```

### testCompleteQueuedWithdrawalWithNonzeroWithdrawalDelayBlocks


```solidity
function testCompleteQueuedWithdrawalWithNonzeroWithdrawalDelayBlocks(uint16 valueToSet) external;
```

### testSlashSharesNotBeaconChainETHFuzzed


```solidity
function testSlashSharesNotBeaconChainETHFuzzed(uint64 withdrawalAmount) external;
```

### testSlashSharesNotBeaconChainETH_AllShares


```solidity
function testSlashSharesNotBeaconChainETH_AllShares() external;
```

### testSlashSharesBeaconChainETH


```solidity
function testSlashSharesBeaconChainETH() external;
```

### testSlashSharesMixIncludingBeaconChainETH


```solidity
function testSlashSharesMixIncludingBeaconChainETH() external;
```

### testSlashSharesRevertsWhenCalledByNotOwner


```solidity
function testSlashSharesRevertsWhenCalledByNotOwner() external;
```

### testSlashSharesRevertsWhenStakerNotFrozen


```solidity
function testSlashSharesRevertsWhenStakerNotFrozen() external;
```

### testSlashSharesRevertsWhenAttemptingReentrancy


```solidity
function testSlashSharesRevertsWhenAttemptingReentrancy() external;
```

### testSlashQueuedWithdrawalNotBeaconChainETH


```solidity
function testSlashQueuedWithdrawalNotBeaconChainETH() external;
```

### testSlashQueuedWithdrawalIncrementor

this function is to test for a bug identified in the Code4Rena audit (H-205).  This bug essentially
allowed a strategy that is meant to be skipped, to actually be withdrawn from.  This is a regression test
to ensure that this bug does not reappear.


```solidity
function testSlashQueuedWithdrawalIncrementor() external;
```

### testSlashQueuedWithdrawalFailsWhenNotCallingFromOwnerAddress

This check ensures that the strategy has not been withdrawn from.  If the incrementor is misplaced inside
the else statement (as it was before the fix was made), the withdrawal would have been triggered for the
the strategy that we intended to skip, i.e., the check indicesToSkip[indicesToSkipIndex] == i would have
failed, triggering the else logic to withdraw from the strategy that was at index 0.


```solidity
function testSlashQueuedWithdrawalFailsWhenNotCallingFromOwnerAddress() external;
```

### testSlashQueuedWithdrawalFailsWhenDelegatedAddressNotFrozen


```solidity
function testSlashQueuedWithdrawalFailsWhenDelegatedAddressNotFrozen() external;
```

### testSlashQueuedWithdrawalFailsWhenAttemptingReentrancy


```solidity
function testSlashQueuedWithdrawalFailsWhenAttemptingReentrancy() external;
```

### testSlashQueuedWithdrawalFailsWhenWithdrawalDoesNotExist


```solidity
function testSlashQueuedWithdrawalFailsWhenWithdrawalDoesNotExist() external;
```

### test_addSharesRevertsWhenSharesIsZero


```solidity
function test_addSharesRevertsWhenSharesIsZero() external;
```

### test_addSharesRevertsWhenDepositWouldExeedMaxArrayLength


```solidity
function test_addSharesRevertsWhenDepositWouldExeedMaxArrayLength() external;
```

### test_depositIntoStrategyRevertsWhenTokenSafeTransferFromReverts


```solidity
function test_depositIntoStrategyRevertsWhenTokenSafeTransferFromReverts() external;
```

### test_depositIntoStrategyRevertsWhenTokenDoesNotExist


```solidity
function test_depositIntoStrategyRevertsWhenTokenDoesNotExist() external;
```

### test_depositIntoStrategyRevertsWhenStrategyDepositFunctionReverts


```solidity
function test_depositIntoStrategyRevertsWhenStrategyDepositFunctionReverts() external;
```

### test_depositIntoStrategyRevertsWhenStrategyDoesNotExist


```solidity
function test_depositIntoStrategyRevertsWhenStrategyDoesNotExist() external;
```

### test_depositIntoStrategyRevertsWhenStrategyNotWhitelisted


```solidity
function test_depositIntoStrategyRevertsWhenStrategyNotWhitelisted() external;
```

### test_removeSharesRevertsWhenShareAmountIsZero


```solidity
function test_removeSharesRevertsWhenShareAmountIsZero() external;
```

### test_removeSharesRevertsWhenShareAmountIsTooLarge


```solidity
function test_removeSharesRevertsWhenShareAmountIsTooLarge() external;
```

### test_removeStrategyFromStakerStrategyListWorksWithIncorrectIndexInput


```solidity
function test_removeStrategyFromStakerStrategyListWorksWithIncorrectIndexInput() external;
```

### testSetWithdrawalDelayBlocks


```solidity
function testSetWithdrawalDelayBlocks(uint16 valueToSet) external;
```

### testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner


```solidity
function testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner(address notOwner)
    external
    filterFuzzedAddressInputs(notOwner);
```

### testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh


```solidity
function testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh(uint256 valueToSet) external;
```

### testSetStrategyWhitelister


```solidity
function testSetStrategyWhitelister(address newWhitelister) external;
```

### testSetStrategyWhitelisterRevertsWhenCalledByNotOwner


```solidity
function testSetStrategyWhitelisterRevertsWhenCalledByNotOwner(address notOwner)
    external
    filterFuzzedAddressInputs(notOwner);
```

### testAddStrategiesToDepositWhitelist


```solidity
function testAddStrategiesToDepositWhitelist(uint8 numberOfStrategiesToAdd) public returns (IStrategy[] memory);
```

### testAddStrategiesToDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister


```solidity
function testAddStrategiesToDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(address notStrategyWhitelister)
    external
    filterFuzzedAddressInputs(notStrategyWhitelister);
```

### testRemoveStrategiesFromDepositWhitelist


```solidity
function testRemoveStrategiesFromDepositWhitelist(uint8 numberOfStrategiesToAdd, uint8 numberOfStrategiesToRemove)
    external;
```

### testRemoveStrategiesFromDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister


```solidity
function testRemoveStrategiesFromDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(
    address notStrategyWhitelister
) external filterFuzzedAddressInputs(notStrategyWhitelister);
```

### _beaconChainReentrancyTestsSetup


```solidity
function _beaconChainReentrancyTestsSetup() internal;
```

### _setUpQueuedWithdrawalStructSingleStrat


```solidity
function _setUpQueuedWithdrawalStructSingleStrat(
    address staker,
    address withdrawer,
    IERC20 token,
    IStrategy strategy,
    uint256 shareAmount
)
    internal
    view
    returns (
        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
        IERC20[] memory tokensArray,
        bytes32 withdrawalRoot
    );
```

### _depositIntoStrategySuccessfully


```solidity
function _depositIntoStrategySuccessfully(IStrategy strategy, address staker, uint256 amount)
    internal
    filterFuzzedAddressInputs(staker);
```

### _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies


```solidity
function _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(
    address staker,
    address withdrawer,
    IStrategy[] memory strategyArray,
    uint256[] memory shareAmounts
) internal view returns (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot);
```

### _arrayWithJustDummyToken


```solidity
function _arrayWithJustDummyToken() internal view returns (IERC20[] memory);
```

### _arrayWithJustTwoDummyTokens


```solidity
function _arrayWithJustTwoDummyTokens() internal view returns (IERC20[] memory);
```

### _depositIntoStrategyWithSignature


```solidity
function _depositIntoStrategyWithSignature(
    address staker,
    uint256 amount,
    uint256 expiry,
    string memory expectedRevertMessage
) internal returns (bytes memory);
```

## Events
### Deposit
Emitted when a new deposit occurs on behalf of `depositor`.


```solidity
event Deposit(address depositor, IERC20 token, IStrategy strategy, uint256 shares);
```

### ShareWithdrawalQueued
Emitted when a new withdrawal occurs on behalf of `depositor`.


```solidity
event ShareWithdrawalQueued(address depositor, uint96 nonce, IStrategy strategy, uint256 shares);
```

### WithdrawalQueued
Emitted when a new withdrawal is queued by `depositor`.


```solidity
event WithdrawalQueued(
    address depositor, uint96 nonce, address withdrawer, address delegatedAddress, bytes32 withdrawalRoot
);
```

### WithdrawalCompleted
Emitted when a queued withdrawal is completed


```solidity
event WithdrawalCompleted(address indexed depositor, uint96 nonce, address indexed withdrawer, bytes32 withdrawalRoot);
```

### StrategyWhitelisterChanged
Emitted when the `strategyWhitelister` is changed


```solidity
event StrategyWhitelisterChanged(address previousAddress, address newAddress);
```

### StrategyAddedToDepositWhitelist
Emitted when a strategy is added to the approved list of strategies for deposit


```solidity
event StrategyAddedToDepositWhitelist(IStrategy strategy);
```

### StrategyRemovedFromDepositWhitelist
Emitted when a strategy is removed from the approved list of strategies for deposit


```solidity
event StrategyRemovedFromDepositWhitelist(IStrategy strategy);
```

### WithdrawalDelayBlocksSet
Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.


```solidity
event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);
```

