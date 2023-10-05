## StrategyManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`StrategyManager.sol`](../../src/contracts/core/StrategyManager.sol) | Singleton | Transparent proxy |
| [`StrategyBaseTVLLimits.sol`](../../src/contracts/strategies/StrategyBaseTVLLimits.sol) | 3 instances (for cbETH, rETH, stETH) | Transparent proxy |

The primary function of the `StrategyManager` is to handle accounting for individual Stakers as they deposit and withdraw LSTs from their corresponding strategies. It is responsible for (i) allowing Stakers to deposit/withdraw LSTs into the corresponding strategy, (ii) managing a queue of Staker withdrawals with an associated withdrawal delay, and (iii) keeping the `DelegationManager` updated as Stakers' shares change during these two operations.

As of M2, three LSTs are supported and each has its own instance of `StrategyBaseTVLLimits`: cbETH, rETH, and stETH. Each `StrategyBaseTVLLimits` has two main functions (`deposit` and `withdraw`), both of which can only be called by the `StrategyManager`. These `StrategyBaseTVLLimits` contracts are fairly simple deposit/withdraw contracts that hold tokens deposited by Stakers. Because these strategies are essentially extensions of the `StrategyManager`, their functions are documented in this file (see below).

*Important state variables*:
* `mapping(address => mapping(IStrategy => uint256)) public stakerStrategyShares`: Tracks the current balance a Staker holds in a given strategy. Updated on deposit/withdraw.
* `mapping(address => IStrategy[]) public stakerStrategyList`: Maintains a list of the strategies a Staker holds a nonzero number of shares in.
    * Updated as needed when Stakers deposit and withdraw: if a Staker has a zero balance in a Strategy, it is removed from the list. Likewise, if a Staker deposits into a Strategy and did not previously have a balance, it is added to the list.
* `mapping(bytes32 => bool) public withdrawalRootPending`: `QueuedWithdrawals` are hashed and set to `true` in this mapping when a withdrawal is initiated. The hash is set to false again when the withdrawal is completed. A per-staker nonce provides a way to distinguish multiple otherwise-identical withdrawals.
* `mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit`: The `strategyWhitelister` is (as of M2) a permissioned role that can be changed by the contract owner. The `strategyWhitelister` has currently whitelisted 3 `StrategyBaseTVLLimits` contracts in this mapping, one for each supported LST.

*Helpful definitions*:
* `stakerStrategyListLength(address staker) -> (uint)`:
    * Gives `stakerStrategyList[staker].length`
    * Used (especially by the `DelegationManager`) to determine whether a Staker has shares in any strategy in the `StrategyManager` (will be 0 if not)
* `uint withdrawalDelayBlocks`:
    * As of M2, this is 50400 (roughly 1 week)
    * Stakers must wait this amount of time before a withdrawal can be completed

### Stakers

The following methods are called by Stakers as they (i) deposit LSTs into strategies to receive shares, (ii) initiate withdrawals of shares, and (iii) finalize withdrawals to receive either shares or tokens.

#### `depositIntoStrategy`

```solidity
function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (uint256 shares)
```

Allows a Staker to deposit some `amount` of `token` into the specified `strategy` in exchange for shares of that strategy. The underlying `strategy` must be one of the three whitelisted `StrategyBaseTVLLimits` instances, and the `token` being deposited must correspond to that `strategy's` underlying token (cbETH, rETH, or stETH).

The number of shares received is calculated by the `strategy` using an internal exchange rate that depends on the previous number of tokens deposited.

If the Staker is delegated to an Operator, the Operator's delegated shares are increased in the `DelegationManager`.

*Effects*:
* `token.safeTransferFrom`: Transfers `amount` of `token` to `strategy` on behalf of the caller.
* See [`StrategyBaseTVLLimits.deposit`](#strategybasetvllimitsdeposit)
* `StrategyManager` awards the Staker with the newly-created shares 
* See [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_DEPOSITS`
* Caller MUST allow at least `amount` of `token` to be transferred by `StrategyManager` to the strategy
* `strategy` in question MUST be whitelisted for deposits. 
* See [`StrategyBaseTVLLimits.deposit`](#strategybasetvllimitsdeposit)

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

#### `depositIntoStrategyWithSignature`

```solidity
function depositIntoStrategyWithSignature(
    IStrategy strategy,
    IERC20 token,
    uint256 amount,
    address staker,
    uint256 expiry,
    bytes memory signature
)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(staker)
    nonReentrant
    returns (uint256 shares)
```

*Effects*: See `depositIntoStrategy` above. Additionally:
* The Staker's nonce is incremented

*Requirements*: See `depositIntoStrategy` above. Additionally:
* Caller MUST provide a valid, unexpired signature over the correct fields

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

#### `queueWithdrawal`

```solidity
function queueWithdrawal(
    uint256[] calldata strategyIndexes,
    IStrategy[] calldata strategies,
    uint256[] calldata shares,
    address withdrawer
)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (bytes32)
```

Allows a Staker to initiate a withdrawal of their held shares across any strategy. Multiple strategies can be included in this single withdrawal with specific share amounts for each strategy. The Staker must specify a `withdrawer` to receive the funds once the withdrawal is completed (although this can be the Staker itself). Withdrawals are able to be completed by calling `completeQueuedWithdrawal` after sufficient time passes (`withdrawalDelayBlocks`).

Before queueing the withdrawal, this method removes the specified shares from the Staker's `StrategyManager` balances and updates the `DelegationManager` via `decreaseDelegatedShares`. If the Staker is delegated to an Operator, this will remove the shares from the Operator's delegated share balances.

Note that at no point during `queueWithdrawal` are the corresponding `StrategyBaseTVLLimits` contracts called; this only occurs once the withdrawal is completed (see `completeQueuedWithdrawal`).

*Effects*:
* The Staker's balances for each strategy in `strategies` is decreased by the corresponding value in `shares`
    * If any of these decreases results in a balance of 0 for a strategy, the strategy is removed from the Staker's strategy list
* A `QueuedWithdrawal` is created for the Staker, tracking the strategies and shares to be withdrawn. 
    * The Staker's withdrawal nonce is increased.
    * The hash of the `QueuedWithdrawal` is marked as "pending"
* See [`DelegationManager.decreaseDelegatedShares`](./DelegationManager.md#decreasedelegatedshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`
* `strategies.length` MUST equal `shares.length`
* `strategyIndexes.length` MUST be at least equal to `strategies.length`
* The Staker MUST have sufficient share balances in the specified strategies
* The `withdrawer` MUST NOT be 0

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

#### `completeQueuedWithdrawal`

```solidity
function completeQueuedWithdrawal(QueuedWithdrawal calldata queuedWithdrawal, IERC20[] calldata tokens, uint256 middlewareTimesIndex, bool receiveAsTokens)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    nonReentrant
```

After waiting `withdrawalDelayBlocks`, this allows the `withdrawer` of a `QueuedWithdrawal` to finalize a withdrawal and receive either (i) the underlying tokens of the strategies being withdrawn from, or (ii) the shares being withdrawn. This choice is dependent on the passed-in parameter `receiveAsTokens`.

For each strategy/share pair in the `QueuedWithdrawal`:
* If the `withdrawer` chooses to receive tokens from the withdrawal, `StrategyBaseTVLLimits.withdraw` exchanges the shares for tokens and transfers them to the `withdrawer`.
* If the `withdrawer` chooses to receive shares, the `StrategyManager` increases the `withdrawer's` strategy share balance.
    * If the `withdrawer` is delegated to an Operator, `DelegationManager.increaseDelegatedShares` will increase that Operator's delegated share balance for the given strategy.

*Effects*:
* The hash of the `QueuedWithdrawal` is removed from the pending withdrawals
* If `receiveAsTokens`, the tokens are withdrawn and transferred to the `withdrawer`:
    * See [`StrategyBaseTVLLimits.withdraw`](#strategybasetvllimitswithdraw)
* If `!receiveAsTokens`, no tokens are moved. Instead:
    * The shares are added to the `withdrawer's` balance for the corresponding strategy
        * If this balance was zero before, the strategy is added to the `withdrawer's` strategy list
    * See [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`
* The hash of the passed-in `QueuedWithdrawal` MUST correspond to a pending withdrawal
    * At least `withdrawalDelayBlocks` MUST have passed before `completeQueuedWithdrawal` is called
    * Caller MUST be the `withdrawer` specified in the `QueuedWithdrawal`
    * If `receiveAsTokens`, the caller MUST pass in the underlying `IERC20[] tokens` being withdrawn in the order they are listed in the `QueuedWithdrawal`.
* See [`StrategyBaseTVLLimits.withdraw`](#strategybasetvllimitswithdraw)

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op
* The `middlewareTimesIndex` parameter has to do with the Slasher, which currently does nothing. As of M2, this parameter has no bearing on anything and can be ignored. It is passed into a call to the Slasher, but the call is a no-op.

#### `completeQueuedWithdrawals`

```solidity
function completeQueuedWithdrawals(
    QueuedWithdrawal[] calldata queuedWithdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexes,
    bool[] calldata receiveAsTokens
) 
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    nonReentrant
```

This method is a looped version of `completeQueuedWithdrawal` that allows a `withdrawer` to finalize several withdrawals in a single call. See `completeQueuedWithdrawal` for details.

### DelegationManager

These methods are callable ONLY by the `DelegationManager`:

#### `forceTotalWithdrawal`

```solidity
function forceTotalWithdrawal(address staker) external
    onlyDelegationManager
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyNotFrozen(staker)
    nonReentrant
    returns (IStrategy[] memory, uint256[] memory, bytes32)
```

The `DelegationManager` calls this method when a Staker is undelegated from an Operator. If the Staker has shares in any strategies, this method removes the shares from the Staker's balance and creates a `QueuedWithdrawal` for the shares and strategies in question. The Staker must wait for `withdrawalDelayBlocks` before they can complete the withdrawal.

The strategies and shares removed from the Staker are returned to the `DelegationManager`; these values will be removed from the Operator's share count.

*Entry Points*:
* `DelegationManager.undelegate`

*Effects*:
* For each strategy in which the Staker holds a balance, that balance is removed and the corresponding strategy is removed from the Staker's strategy list.
* A `QueuedWithdrawal` is created for the Staker which can be completed after the withdrawal delay

*Requirements*:
* Caller MUST be the `DelegationManager`
* Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`

*As of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

### Operator

#### `setWithdrawalDelayBlocks`

```solidity
function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) external onlyOwner
```

Allows the `owner` to update the number of blocks that must pass before a withdrawal can be completed.

*Effects*:
* Updates `StrategyManager.withdrawalDelayBlocks`

*Requirements*:
* Caller MUST be the `owner`
* `_withdrawalDelayBlocks` MUST NOT be greater than `MAX_WITHDRAWAL_DELAY_BLOCKS` (50400)

#### `setStrategyWhitelister`

```solidity
function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner
```

Allows the `owner` to update the Strategy Whitelister address.

*Effects*:
* Updates `StrategyManager.strategyWhitelister`

*Requirements*:
* Caller MUST be the `owner`

### Strategy Whitelister

#### `addStrategiesToDepositWhitelist`

```solidity
function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external onlyStrategyWhitelister
```

Allows the Strategy Whitelister address to add any number of strategies to the `StrategyManager` whitelist. Strategies on the whitelist are eligible for deposit via `depositIntoStrategy`.

*Effects*:
* Adds entries to `StrategyManager.strategyIsWhitelistedForDeposit`

*Requirements*:
* Caller MUST be the `strategyWhitelister`

#### `removeStrategiesFromDepositWhitelist`

```solidity
function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external onlyStrategyWhitelister
```

Allows the Strategy Whitelister address to remove any number of strategies from the `StrategyManager` whitelist. The removed strategies will no longer be eligible for deposit via `depositIntoStrategy`. However, withdrawals for previously-whitelisted strategies may still be initiated and completed, as long as the Staker has shares to withdraw.

*Effects*:
* Removes entries from `StrategyManager.strategyIsWhitelistedForDeposit`

*Requirements*:
* Caller MUST be the `strategyWhitelister`

---

### StrategyBaseTVLLimits

`StrategyBaseTVLLimits` only has two methods of note, and both can only be called by the `StrategyManager`. Documentation for these methods are included below, rather than in a separate file.

#### `StrategyBaseTVLLimits.deposit`

```solidity
function deposit(IERC20 token, uint256 amount)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyStrategyManager
    returns (uint256 newShares)
```

The `StrategyManager` calls this method when Stakers deposit LSTs into a strategy. At the time this method is called, the tokens have already been transferred to the strategy. The role of this method is to (i) calculate the number of shares the deposited tokens represent according to the exchange rate, and (ii) add the new shares to the strategy's recorded total shares.

The new shares created are returned to the `StrategyManager` to be added to the Staker's strategy share balance.

*Entry Points*:
* `StrategyManager.depositIntoStrategy`
* `StrategyManager.depositIntoStrategyWithSignature`

*Effects*:
* `StrategyBaseTVLLimits.totalShares` is increased to account for the new shares created by the deposit

*Requirements*:
* Caller MUST be the `StrategyManager`
* Pause status MUST NOT be set: `PAUSED_DEPOSITS`
* The passed-in `token` MUST match the strategy's `underlyingToken`
* The token amount being deposited MUST NOT exceed the per-deposit cap
* After deposit, the strategy's current token balance MUST NOT exceed the total-deposit cap
* When converted to shares via the strategy's exchange rate, the `amount` of `token` deposited MUST represent at least 1 new share for the depositor

#### `StrategyBaseTVLLimits.withdraw`

```solidity
function withdraw(address depositor, IERC20 token, uint256 amountShares)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyStrategyManager
```

The `StrategyManager` calls this method when a queued withdrawal from a strategy is completed, assuming the withdrawer specifies they would like to receive the withdrawal as tokens (see `completeQueuedWithdrawal`). 

This method converts the withdrawal shares back into tokens using the strategy's exchange rate. The strategy's total shares are decreased to reflect the withdrawal before transferring the tokens to the withdrawer.

*Entry Points*:
* `StrategyManager.completeQueuedWithdrawal`
* `StrategyManager.completeQueuedWithdrawals`

*Effects*:
* `StrategyBaseTVLLimits.totalShares` is decreased to account for the shares being withdrawn
* `underlyingToken.safeTransfer` is called to transfer the tokens to the withdrawer

*Requirements*:
* Caller MUST be the `StrategyManager`
* Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`
* The passed-in `token` MUST match the strategy's `underlyingToken`
* The `amountShares` being withdrawn MUST NOT exceed the `totalShares` in the strategy
* The tokens represented by `amountShares` MUST NOT exceed the strategy's token balance