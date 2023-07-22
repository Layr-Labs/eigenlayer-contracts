# StrategyManager
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/core/StrategyManager.sol)

**Inherits:**
Initializable, OwnableUpgradeable, ReentrancyGuardUpgradeable, [Pausable](/src/contracts/permissions/Pausable.sol/contract.Pausable.md), [StrategyManagerStorage](/src/contracts/core/StrategyManagerStorage.sol/abstract.StrategyManagerStorage.md)

**Author:**
Layr Labs, Inc.

Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service

This contract is for managing deposits in different strategies. The main
functionalities are:
- adding and removing strategies that any delegator can deposit into
- enabling deposit of assets into specified strategy(s)
- enabling withdrawal of assets from specified strategy(s)
- recording deposit of ETH into settlement layer
- slashing of assets for permissioned strategies


## State Variables
### GWEI_TO_WEI

```solidity
uint256 internal constant GWEI_TO_WEI = 1e9;
```


### PAUSED_DEPOSITS

```solidity
uint8 internal constant PAUSED_DEPOSITS = 0;
```


### PAUSED_WITHDRAWALS

```solidity
uint8 internal constant PAUSED_WITHDRAWALS = 1;
```


### ORIGINAL_CHAIN_ID

```solidity
uint256 immutable ORIGINAL_CHAIN_ID;
```


### ERC1271_MAGICVALUE

```solidity
bytes4 internal constant ERC1271_MAGICVALUE = 0x1626ba7e;
```


## Functions
### onlyNotFrozen


```solidity
modifier onlyNotFrozen(address staker);
```

### onlyFrozen


```solidity
modifier onlyFrozen(address staker);
```

### onlyEigenPodManager


```solidity
modifier onlyEigenPodManager();
```

### onlyStrategyWhitelister


```solidity
modifier onlyStrategyWhitelister();
```

### onlyStrategiesWhitelistedForDeposit


```solidity
modifier onlyStrategiesWhitelistedForDeposit(IStrategy strategy);
```

### constructor


```solidity
constructor(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher)
    StrategyManagerStorage(_delegation, _eigenPodManager, _slasher);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_delegation`|`IDelegationManager`|The delegation contract of EigenLayer.|
|`_eigenPodManager`|`IEigenPodManager`|The contract that keeps track of EigenPod stakes for restaking beacon chain ether.|
|`_slasher`|`ISlasher`|The primary slashing contract of EigenLayer.|


### initialize

Initializes the strategy manager contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
and transfers contract ownership to the specified `initialOwner`.


```solidity
function initialize(
    address initialOwner,
    address initialStrategyWhitelister,
    IPauserRegistry _pauserRegistry,
    uint256 initialPausedStatus,
    uint256 _withdrawalDelayBlocks
) external initializer;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`initialOwner`|`address`|Ownership of this contract is transferred to this address.|
|`initialStrategyWhitelister`|`address`|The initial value of `strategyWhitelister` to set.|
|`_pauserRegistry`|`IPauserRegistry`|Used for access control of pausing.|
|`initialPausedStatus`|`uint256`|The initial value of `_paused` to set.|
|`_withdrawalDelayBlocks`|`uint256`|The initial value of `withdrawalDelayBlocks` to set.|


### depositBeaconChainETH

Deposits `amount` of beaconchain ETH into this contract on behalf of `staker`

*Only callable by EigenPodManager.*


```solidity
function depositBeaconChainETH(address staker, uint256 amount)
    external
    onlyEigenPodManager
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(staker)
    nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`staker`|`address`|is the entity that is restaking in eigenlayer,|
|`amount`|`uint256`|is the amount of beaconchain ETH being restaked,|


### recordOvercommittedBeaconChainETH

Records an overcommitment event on behalf of a staker. The staker's beaconChainETH shares are decremented by `amount`.

*Only callable by EigenPodManager.*


```solidity
function recordOvercommittedBeaconChainETH(
    address overcommittedPodOwner,
    uint256 beaconChainETHStrategyIndex,
    uint256 amount
) external onlyEigenPodManager nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`overcommittedPodOwner`|`address`|is the pod owner to be slashed|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy in case it must be removed,|
|`amount`|`uint256`|is the amount to decrement the slashedAddress's beaconChainETHStrategy shares|


### depositIntoStrategy

Deposits `amount` of `token` into the specified `strategy`, with the resultant shares credited to `msg.sender`

*The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.*

*Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
where the token balance and corresponding strategy shares are not in sync upon reentrancy.*


```solidity
function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (uint256 shares);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategy`|`IStrategy`|is the specified strategy where deposit is to be made,|
|`token`|`IERC20`|is the denomination in which the deposit is to be made,|
|`amount`|`uint256`|is the amount of token to be deposited in the strategy by the depositor|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`shares`|`uint256`|The amount of new shares in the `strategy` created as part of the action.|


### depositIntoStrategyWithSignature

Used for depositing an asset into the specified strategy with the resultant shares credited to `staker`,
who must sign off on the action.
Note that the assets are transferred out/from the `msg.sender`, not from the `staker`; this function is explicitly designed
purely to help one address deposit 'for' another.

*The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.*

*A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
targeting stakers who may be attempting to undelegate.*

*Cannot be called on behalf of a staker that is 'frozen' (this function will revert if the `staker` is frozen).
WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
where the token balance and corresponding strategy shares are not in sync upon reentrancy*


```solidity
function depositIntoStrategyWithSignature(
    IStrategy strategy,
    IERC20 token,
    uint256 amount,
    address staker,
    uint256 expiry,
    bytes memory signature
) external onlyWhenNotPaused(PAUSED_DEPOSITS) onlyNotFrozen(staker) nonReentrant returns (uint256 shares);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategy`|`IStrategy`|is the specified strategy where deposit is to be made,|
|`token`|`IERC20`|is the denomination in which the deposit is to be made,|
|`amount`|`uint256`|is the amount of token to be deposited in the strategy by the depositor|
|`staker`|`address`|the staker that the deposited assets will be credited to|
|`expiry`|`uint256`|the timestamp at which the signature expires|
|`signature`|`bytes`|is a valid signature from the `staker`. either an ECDSA signature if the `staker` is an EOA, or data to forward following EIP-1271 if the `staker` is a contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`shares`|`uint256`|The amount of new shares in the `strategy` created as part of the action.|


### undelegate

check validity of signature:
1) if `staker` is an EOA, then `signature` must be a valid ECDSA signature from `staker`,
indicating their intention for this action
2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271

Called by a staker to undelegate entirely from EigenLayer. The staker must first withdraw all of their existing deposits
(through use of the `queueWithdrawal` function), or else otherwise have never deposited in EigenLayer prior to delegating.


```solidity
function undelegate() external;
```

### queueWithdrawal

Called by a staker to queue a withdrawal of the given amount of `shares` from each of the respective given `strategies`.

*Stakers will complete their withdrawal by calling the 'completeQueuedWithdrawal' function.
User shares are decreased in this function, but the total number of shares in each strategy remains the same.
The total number of shares is decremented in the 'completeQueuedWithdrawal' function instead, which is where
the funds are actually sent to the user through use of the strategies' 'withdrawal' function. This ensures
that the value per share reported by each strategy will remain consistent, and that the shares will continue
to accrue gains during the enforced withdrawal waiting period.*

*Strategies are removed from `stakerStrategyList` by swapping the last entry with the entry to be removed, then
popping off the last entry in `stakerStrategyList`. The simplest way to calculate the correct `strategyIndexes` to input
is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
`stakerStrategyList` to lowest index*

*Note that if the withdrawal includes shares in the enshrined 'beaconChainETH' strategy, then it must *only* include shares in this strategy, and
`withdrawer` must match the caller's address. The first condition is because slashing of queued withdrawals cannot be guaranteed
for Beacon Chain ETH (since we cannot trigger a withdrawal from the beacon chain through a smart contract) and the second condition is because shares in
the enshrined 'beaconChainETH' strategy technically represent non-fungible positions (deposits to the Beacon Chain, each pointed at a specific EigenPod).*


```solidity
function queueWithdrawal(
    uint256[] calldata strategyIndexes,
    IStrategy[] calldata strategies,
    uint256[] calldata shares,
    address withdrawer,
    bool undelegateIfPossible
) external onlyWhenNotPaused(PAUSED_WITHDRAWALS) onlyNotFrozen(msg.sender) nonReentrant returns (bytes32);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategyIndexes`|`uint256[]`|is a list of the indices in `stakerStrategyList[msg.sender]` that correspond to the strategies for which `msg.sender` is withdrawing 100% of their shares|
|`strategies`|`IStrategy[]`|The Strategies to withdraw from|
|`shares`|`uint256[]`|The amount of shares to withdraw from each of the respective Strategies in the `strategies` array|
|`withdrawer`|`address`|The address that can complete the withdrawal and will receive any withdrawn funds or shares upon completing the withdrawal|
|`undelegateIfPossible`|`bool`|If this param is marked as 'true' *and the withdrawal will result in `msg.sender` having no shares in any Strategy,* then this function will also make an internal call to `undelegate(msg.sender)` to undelegate the `msg.sender`.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes32`|The 'withdrawalRoot' of the newly created Queued Withdrawal|


### completeQueuedWithdrawal

Ensure that if the withdrawal includes beacon chain ETH, the specified 'withdrawer' is not different than the caller.
This is because shares in the enshrined `beaconChainETHStrategy` ultimately represent tokens in **non-fungible** EigenPods,
while other share in all other strategies represent purely fungible positions.
Checking that `stakerStrategyList[msg.sender].length == 0` is not strictly necessary here, but prevents reverting very late in logic,
in the case that 'undelegate' is set to true but the `msg.sender` still has active deposits in EigenLayer.

Used to complete the specified `queuedWithdrawal`. The function caller must match `queuedWithdrawal.withdrawer`

*middlewareTimesIndex should be calculated off chain before calling this function by finding the first index that satisfies `slasher.canWithdraw`*


```solidity
function completeQueuedWithdrawal(
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) external onlyWhenNotPaused(PAUSED_WITHDRAWALS) nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`queuedWithdrawal`|`QueuedWithdrawal`|The QueuedWithdrawal to complete.|
|`tokens`|`IERC20[]`|Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `strategies` array of the `queuedWithdrawal`. This input can be provided with zero length if `receiveAsTokens` is set to 'false' (since in that case, this input will be unused)|
|`middlewareTimesIndex`|`uint256`|is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array|
|`receiveAsTokens`|`bool`|If true, the shares specified in the queued withdrawal will be withdrawn from the specified strategies themselves and sent to the caller, through calls to `queuedWithdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies will simply be transferred to the caller directly.|


### completeQueuedWithdrawals

Used to complete the specified `queuedWithdrawals`. The function caller must match `queuedWithdrawals[...].withdrawer`

*Array-ified version of `completeQueuedWithdrawal`*

*middlewareTimesIndex should be calculated off chain before calling this function by finding the first index that satisfies `slasher.canWithdraw`*


```solidity
function completeQueuedWithdrawals(
    QueuedWithdrawal[] calldata queuedWithdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexes,
    bool[] calldata receiveAsTokens
) external onlyWhenNotPaused(PAUSED_WITHDRAWALS) nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`queuedWithdrawals`|`QueuedWithdrawal[]`|The QueuedWithdrawals to complete.|
|`tokens`|`IERC20[][]`|Array of tokens for each QueuedWithdrawal. See `completeQueuedWithdrawal` for the usage of a single array.|
|`middlewareTimesIndexes`|`uint256[]`|One index to reference per QueuedWithdrawal. See `completeQueuedWithdrawal` for the usage of a single index.|
|`receiveAsTokens`|`bool[]`|If true, the shares specified in the queued withdrawal will be withdrawn from the specified strategies themselves and sent to the caller, through calls to `queuedWithdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies will simply be transferred to the caller directly.|


### slashShares

Slashes the shares of a 'frozen' operator (or a staker delegated to one)

*strategies are removed from `stakerStrategyList` by swapping the last entry with the entry to be removed, then
popping off the last entry in `stakerStrategyList`. The simplest way to calculate the correct `strategyIndexes` to input
is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
`stakerStrategyList` to lowest index*


```solidity
function slashShares(
    address slashedAddress,
    address recipient,
    IStrategy[] calldata strategies,
    IERC20[] calldata tokens,
    uint256[] calldata strategyIndexes,
    uint256[] calldata shareAmounts
) external onlyOwner onlyFrozen(slashedAddress) nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`slashedAddress`|`address`|is the frozen address that is having its shares slashed|
|`recipient`|`address`|is the address that will receive the slashed funds, which could e.g. be a harmed party themself, or a MerkleDistributor-type contract that further sub-divides the slashed funds.|
|`strategies`|`IStrategy[]`|Strategies to slash|
|`tokens`|`IERC20[]`|The tokens to use as input to the `withdraw` function of each of the provided `strategies`|
|`strategyIndexes`|`uint256[]`|is a list of the indices in `stakerStrategyList[msg.sender]` that correspond to the strategies for which `msg.sender` is withdrawing 100% of their shares|
|`shareAmounts`|`uint256[]`|The amount of shares to slash in each of the provided `strategies`|


### slashQueuedWithdrawal

Slashes an existing queued withdrawal that was created by a 'frozen' operator (or a staker delegated to one)


```solidity
function slashQueuedWithdrawal(
    address recipient,
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256[] calldata indicesToSkip
) external onlyOwner onlyFrozen(queuedWithdrawal.delegatedAddress) nonReentrant;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The funds in the slashed withdrawal are withdrawn as tokens to this address.|
|`queuedWithdrawal`|`QueuedWithdrawal`|The previously queued withdrawal to be slashed|
|`tokens`|`IERC20[]`|Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `strategies` array of the `queuedWithdrawal`.|
|`indicesToSkip`|`uint256[]`|Optional input parameter -- indices in the `strategies` array to skip (i.e. not call the 'withdraw' function on). This input exists so that, e.g., if the slashed QueuedWithdrawal contains a malicious strategy in the `strategies` array which always reverts on calls to its 'withdraw' function, then the malicious strategy can be skipped (with the shares in effect "burned"), while the non-malicious strategies are still called as normal.|


### setWithdrawalDelayBlocks

Owner-only function for modifying the value of the `withdrawalDelayBlocks` variable.


```solidity
function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_withdrawalDelayBlocks`|`uint256`|new value of `withdrawalDelayBlocks`.|


### setStrategyWhitelister

Owner-only function to change the `strategyWhitelister` address.


```solidity
function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newStrategyWhitelister`|`address`|new address for the `strategyWhitelister`.|


### addStrategiesToDepositWhitelist

Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into


```solidity
function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external onlyStrategyWhitelister;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategiesToWhitelist`|`IStrategy[]`|Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)|


### removeStrategiesFromDepositWhitelist

Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into


```solidity
function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist)
    external
    onlyStrategyWhitelister;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategiesToRemoveFromWhitelist`|`IStrategy[]`|Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)|


### _addShares

This function adds `shares` for a given `strategy` to the `depositor` and runs through the necessary update logic.

*In particular, this function calls `delegation.increaseDelegatedShares(depositor, strategy, shares)` to ensure that all
delegated shares are tracked, increases the stored share amount in `stakerStrategyShares[depositor][strategy]`, and adds `strategy`
to the `depositor`'s list of strategies, if it is not in the list already.*


```solidity
function _addShares(address depositor, IStrategy strategy, uint256 shares) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The address to add shares to|
|`strategy`|`IStrategy`|The Strategy in which the `depositor` is receiving shares|
|`shares`|`uint256`|The amount of shares to grant to the `depositor`|


### _depositIntoStrategy

Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the Strategy-type contract
`strategy`, with the resulting shares credited to `depositor`.


```solidity
function _depositIntoStrategy(address depositor, IStrategy strategy, IERC20 token, uint256 amount)
    internal
    onlyStrategiesWhitelistedForDeposit(strategy)
    returns (uint256 shares);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The address that will be credited with the new shares.|
|`strategy`|`IStrategy`|The Strategy contract to deposit into.|
|`token`|`IERC20`|The ERC20 token to deposit.|
|`amount`|`uint256`|The amount of `token` to deposit.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`shares`|`uint256`|The amount of *new* shares in `strategy` that have been credited to the `depositor`.|


### _removeShares

Decreases the shares that `depositor` holds in `strategy` by `shareAmount`.

*If the amount of shares represents all of the depositor`s shares in said strategy,
then the strategy is removed from stakerStrategyList[depositor] and 'true' is returned. Otherwise 'false' is returned.*


```solidity
function _removeShares(address depositor, uint256 strategyIndex, IStrategy strategy, uint256 shareAmount)
    internal
    returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The address to decrement shares from|
|`strategyIndex`|`uint256`|The `strategyIndex` input for the internal `_removeStrategyFromStakerStrategyList`. Used only in the case that the removal of the depositor's shares results in them having zero remaining shares in the `strategy`|
|`strategy`|`IStrategy`|The strategy for which the `depositor`'s shares are being decremented|
|`shareAmount`|`uint256`|The amount of shares to decrement|


### _removeStrategyFromStakerStrategyList

Removes `strategy` from `depositor`'s dynamic array of strategies, i.e. from `stakerStrategyList[depositor]`

*the provided `strategyIndex` input is optimistically used to find the strategy quickly in the list. If the specified
index is incorrect, then we revert to a brute-force search.*


```solidity
function _removeStrategyFromStakerStrategyList(address depositor, uint256 strategyIndex, IStrategy strategy) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The user whose array will have an entry removed|
|`strategyIndex`|`uint256`|Preferably the index of `strategy` in `stakerStrategyList[depositor]`. If the input is incorrect, then a brute-force fallback routine will be used to find the correct input|
|`strategy`|`IStrategy`|The Strategy to remove from `stakerStrategyList[depositor]`|


### _completeQueuedWithdrawal

Internal function for completing the given `queuedWithdrawal`.


```solidity
function _completeQueuedWithdrawal(
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) internal onlyNotFrozen(queuedWithdrawal.delegatedAddress);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`queuedWithdrawal`|`QueuedWithdrawal`|The QueuedWithdrawal to complete|
|`tokens`|`IERC20[]`|The ERC20 tokens to provide as inputs to `Strategy.withdraw`. Only relevant if `receiveAsTokens = true`|
|`middlewareTimesIndex`|`uint256`|Passed on as an input to the `slasher.canWithdraw` function, to ensure the withdrawal is completable.|
|`receiveAsTokens`|`bool`|If marked 'true', then calls will be passed on to the `Strategy.withdraw` function for each strategy. If marked 'false', then the shares will simply be internally transferred to the `msg.sender`.|


### _undelegate

If the `depositor` has no existing shares, then they can `undelegate` themselves.
This allows people a "hard reset" in their relationship with EigenLayer after withdrawing all of their stake.


```solidity
function _undelegate(address depositor) internal onlyNotFrozen(depositor);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The address to undelegate. Passed on as an input to the `delegation.undelegate` function.|


### _withdrawBeaconChainETH


```solidity
function _withdrawBeaconChainETH(address staker, address recipient, uint256 amount) internal;
```

### _setWithdrawalDelayBlocks

internal function for changing the value of `withdrawalDelayBlocks`. Also performs sanity check and emits an event.


```solidity
function _setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_withdrawalDelayBlocks`|`uint256`|The new value for `withdrawalDelayBlocks` to take.|


### _setStrategyWhitelister

Internal function for modifying the `strategyWhitelister`. Used inside of the `setStrategyWhitelister` and `initialize` functions.


```solidity
function _setStrategyWhitelister(address newStrategyWhitelister) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newStrategyWhitelister`|`address`|The new address for the `strategyWhitelister` to take.|


### getDeposits

Get all details on the depositor's deposits and corresponding shares


```solidity
function getDeposits(address depositor) external view returns (IStrategy[] memory, uint256[] memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`depositor`|`address`|The staker of interest, whose deposits this function will fetch|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`IStrategy[]`|(depositor's strategies, shares in these strategies)|
|`<none>`|`uint256[]`||


### stakerStrategyListLength

Simple getter function that returns `stakerStrategyList[staker].length`.


```solidity
function stakerStrategyListLength(address staker) external view returns (uint256);
```

### calculateWithdrawalRoot

Returns the keccak256 hash of `queuedWithdrawal`.


```solidity
function calculateWithdrawalRoot(QueuedWithdrawal memory queuedWithdrawal) public pure returns (bytes32);
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

