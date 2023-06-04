# IStrategyManager
[Git Source](https://github.com/Sabnock01/eigenlayer-contracts/blob/fa80db0202cf74fb2bae3ffc6aa6db988074a698/src/contracts/interfaces/IStrategyManager.sol)

**Author:**
Layr Labs, Inc.

See the `StrategyManager` contract itself for implementation details.


## Functions
### depositIntoStrategy

Deposits `amount` of `token` into the specified `strategy`, with the resultant shares credited to `msg.sender`

*The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.*

*Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
where the token balance and corresponding strategy shares are not in sync upon reentrancy.*


```solidity
function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount) external returns (uint256 shares);
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


### depositBeaconChainETH

Deposits `amount` of beaconchain ETH into this contract on behalf of `staker`

*Only callable by EigenPodManager.*


```solidity
function depositBeaconChainETH(address staker, uint256 amount) external;
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
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`overcommittedPodOwner`|`address`|is the pod owner to be slashed|
|`beaconChainETHStrategyIndex`|`uint256`|is the index of the beaconChainETHStrategy in case it must be removed,|
|`amount`|`uint256`|is the amount to decrement the slashedAddress's beaconChainETHStrategy shares|


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
) external returns (uint256 shares);
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


### stakerStrategyShares

Returns the current shares of `user` in `strategy`


```solidity
function stakerStrategyShares(address user, IStrategy strategy) external view returns (uint256 shares);
```

### getDeposits

Get all details on the depositor's deposits and corresponding shares


```solidity
function getDeposits(address depositor) external view returns (IStrategy[] memory, uint256[] memory);
```
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
) external returns (bytes32);
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

Used to complete the specified `queuedWithdrawal`. The function caller must match `queuedWithdrawal.withdrawer`

*middlewareTimesIndex should be calculated off chain before calling this function by finding the first index that satisfies `slasher.canWithdraw`*


```solidity
function completeQueuedWithdrawal(
    QueuedWithdrawal calldata queuedWithdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
) external;
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
) external;
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
) external;
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
) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The funds in the slashed withdrawal are withdrawn as tokens to this address.|
|`queuedWithdrawal`|`QueuedWithdrawal`|The previously queued withdrawal to be slashed|
|`tokens`|`IERC20[]`|Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `strategies` array of the `queuedWithdrawal`.|
|`indicesToSkip`|`uint256[]`|Optional input parameter -- indices in the `strategies` array to skip (i.e. not call the 'withdraw' function on). This input exists so that, e.g., if the slashed QueuedWithdrawal contains a malicious strategy in the `strategies` array which always reverts on calls to its 'withdraw' function, then the malicious strategy can be skipped (with the shares in effect "burned"), while the non-malicious strategies are still called as normal.|


### calculateWithdrawalRoot

Returns the keccak256 hash of `queuedWithdrawal`.


```solidity
function calculateWithdrawalRoot(QueuedWithdrawal memory queuedWithdrawal) external pure returns (bytes32);
```

### addStrategiesToDepositWhitelist

Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into


```solidity
function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategiesToWhitelist`|`IStrategy[]`|Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)|


### removeStrategiesFromDepositWhitelist

Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into


```solidity
function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`strategiesToRemoveFromWhitelist`|`IStrategy[]`|Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)|


### delegation

Returns the single, central Delegation contract of EigenLayer


```solidity
function delegation() external view returns (IDelegationManager);
```

### slasher

Returns the single, central Slasher contract of EigenLayer


```solidity
function slasher() external view returns (ISlasher);
```

### beaconChainETHStrategy

returns the enshrined, virtual 'beaconChainETH' Strategy


```solidity
function beaconChainETHStrategy() external view returns (IStrategy);
```

### withdrawalDelayBlocks

Returns the number of blocks that must pass between the time a withdrawal is queued and the time it can be completed


```solidity
function withdrawalDelayBlocks() external view returns (uint256);
```

## Structs
### WithdrawerAndNonce

```solidity
struct WithdrawerAndNonce {
    address withdrawer;
    uint96 nonce;
}
```

### QueuedWithdrawal
Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
In functions that operate on existing queued withdrawals -- e.g. `startQueuedWithdrawalWaitingPeriod` or `completeQueuedWithdrawal`,
the data is resubmitted and the hash of the submitted data is computed by `calculateWithdrawalRoot` and checked against the
stored hash in order to confirm the integrity of the submitted data.


```solidity
struct QueuedWithdrawal {
    IStrategy[] strategies;
    uint256[] shares;
    address depositor;
    WithdrawerAndNonce withdrawerAndNonce;
    uint32 withdrawalStartBlock;
    address delegatedAddress;
}
```

