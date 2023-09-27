## StrategyManager

| File | Type | Proxy? |
| -------- | -------- | -------- |
| [`StrategyManager.sol`](#TODO) | Singleton | Transparent proxy |
| [`StrategyBaseTVLLimits.sol`](#TODO) | 3 instances (for cbETH, rETH, stETH) | Transparent proxy |

<!-- Technical details on the LST subsystem as it functions during M2. Includes:
* StrategyManager
* Strategies (cbETH, rETH, stETH)
* LST restaking
* Stake / withdrawal flows -->

The primary function of the `StrategyManager` is to handle accounting for individual Stakers as they deposit and withdraw LSTs from their corresponding strategies. It is responsible for (i) allowing Stakers to deposit LST tokens into the corresponding strategy, (ii) managing a queue of Staker withdrawals with an associated withdrawal delay, and (iii) keeping the `DelegationManager` updated as Stakers' shares change during these two operations.

As of M2, three LSTs are supported and each has its own instance of `StrategyBaseTVLLimits`: cbETH, rETH, and stETH. Each `StrategyBaseTVLLimits` has two main functions (`deposit` and `withdraw`), both of which can only be called by the `StrategyManager`.

Note: the `StrategyManager` does NOT hold tokens. It does, however, (i) `transferFrom` tokens from a Staker to a strategy on deposits and (ii) direct strategies to transfer tokens back to Stakers on withdrawal.

### Stakers

#### `depositIntoStrategy`

```solidity
function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (uint256 shares)
```

Allows a Staker to deposit some `amount` of `token` into the specified `strategy` in exchange for shares of that strategy. The underlying `strategy` must be one of the three whitelisted `StrategyBaseTVLLimits` instances, and the `token` being deposited must correspond to that `strategy's` underlying token (cbETH, rETH, or stETH). If the Staker is delegated to an Operator, the Operator's shares are increased in the `DelegationManager`.

**Effects**:
* `token.safeTransferFrom`: Transfers `amount` of `token` to `strategy` on behalf of the caller.
* `strategy.deposit`: `StrategyBaseTVLLimits` calculates an exchange rate based on its current token holdings and total share amount. The total share number is then increased, and the number of shares created is returned to the `StrategyManager`. Individual strategies do not track shares on a per-account basis; this is handled in `StrategyManager`.
* `StrategyManager` awards the Staker with the newly-created shares 
* `DelegationManager.increaseDelegatedShares`: If the Staker is delegated to an Operator (or is themselves an Operator), the Operator's shares are increased for the strategy in question

**Requirements**:
* Pause status MUST NOT be set (`StrategyManager`): `PAUSED_DEPOSITS`
* Caller MUST allow at least `amount` of `token` to be transferred by `StrategyManager`
* `strategy` in question MUST be whitelisted for deposits. Additionally:
    * Pause status MUST NOT be set (`StrategyBaseTVLLimits`): `PAUSED_DEPOSITS`
    * `token` must be the correct `underlyingToken` used by `strategy` (i.e. cannot transfer rETH into the cbETH strategy)
    * `amount` of `token` transferred must not put `strategy` above its configured deposit caps
    * `shares` awarded on deposit must be nonzero

*Unimplemented as of M2*:
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

**Effects**: See `depositIntoStrategy` above. Additionally:
* The Staker's nonce is incremented

**Requirements**: See `depositIntoStrategy` above. Additionally:
* Caller MUST provide a valid, unexpired signature over the correct fields

*Unimplemented as of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

#### `undelegate`

```solidity
function undelegate() external
```

Allows a Staker to undelegate from an Operator in the `DelegationManager`, as long as they have no shares in either the `StrategyManager` or `EigenPodManager`. Note that the `StrategyManager` allows undelegation if a Staker is in the withdrawal queue.

**Effects**:
* `DelegationManager.undelegate`: undelegates the Staker from their delegated Operator

**Requirements**:
* Staker MUST NOT have shares in any strategy within `StrategyManager`
* Staker MUST NOT have shares in the `EigenPodManager`
* Staker MUST NOT be an Operator (in `DelegationManager`)

*Unimplemented as of M2*:
* The `onlyNotFrozen` modifier is currently a no-op

#### `queueWithdrawal`

```solidity
function queueWithdrawal(
    uint256[] calldata strategyIndexes,
    IStrategy[] calldata strategies,
    uint256[] calldata shares,
    address withdrawer,
    bool undelegateIfPossible
)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (bytes32)
```

**Effects**:

**Requirements**:

#### `completeQueuedWithdrawal`

```solidity
function completeQueuedWithdrawal(QueuedWithdrawal calldata queuedWithdrawal, IERC20[] calldata tokens, uint256 middlewareTimesIndex, bool receiveAsTokens)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    nonReentrant
```

**Effects**:

**Requirements**:

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

**Effects**:

**Requirements**:

### DelegationManager

#### `forceTotalWithdrawal`

```solidity
function forceTotalWithdrawal(address staker) external
    onlyDelegationManager
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyNotFrozen(staker)
    nonReentrant
    returns (bytes32)
```

**Effects**:

**Requirements**:

### Operator

#### `setWithdrawalDelayBlocks`

```solidity
function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) external onlyOwner
```

**Effects**:

**Requirements**:

#### `setStrategyWhitelister`

```solidity
function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner
```

**Effects**:

**Requirements**:

#### `slashShares`

```solidity
function slashShares(
    address slashedAddress,
    address recipient,
    IStrategy[] calldata strategies,
    IERC20[] calldata tokens,
    uint256[] calldata strategyIndexes,
    uint256[] calldata shareAmounts
)
    external
    onlyOwner
    onlyFrozen(slashedAddress)
    nonReentrant
```

**Effects**:

**Requirements**:

#### `slashQueuedWithdrawal`

```solidity
function slashQueuedWithdrawal(address recipient, QueuedWithdrawal calldata queuedWithdrawal, IERC20[] calldata tokens, uint256[] calldata indicesToSkip)
    external
    onlyOwner
    onlyFrozen(queuedWithdrawal.delegatedAddress)
    nonReentrant
```

**Effects**:

**Requirements**:

### StrategyWhitelister

#### `addStrategiesToDepositWhitelist`

```solidity
function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external onlyStrategyWhitelister
```

**Effects**:

**Requirements**:

#### `removeStrategiesFromDepositWhitelist`

```solidity
function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external onlyStrategyWhitelister
```

**Effects**:

**Requirements**:

---

## StrategyBaseTVLLimits