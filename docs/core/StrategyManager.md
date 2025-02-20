## StrategyManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`StrategyManager.sol`](../../src/contracts/core/StrategyManager.sol) | Singleton | Transparent proxy |
| [`StrategyFactory.sol`](../../src/contracts/core/StrategyFactory.sol) | Singleton | Transparent proxy |
| [`StrategyBaseTVLLimits.sol`](../../src/contracts/strategies/StrategyBaseTVLLimits.sol) | Instanced, one per supported token | - Strategies deployed outside the `StrategyFactory` use transparent proxies <br /> - Anything deployed via the `StrategyFactory` uses a Beacon proxy |

The primary function of the `StrategyManager` is to handle accounting for individual Stakers as they deposit and withdraw supported tokens from their corresponding strategies. It is responsible for (i) allowing Stakers to deposit tokens into the corresponding strategy, (ii) allowing the `DelegationManager` to remove shares when a Staker queues a withdrawal, and (iii) allowing the `DelegationManager` to complete a withdrawal by either adding shares back to the Staker or withdrawing the shares as tokens via the corresponding strategy.

Any ERC20-compatible token can be supported by deploying a `StrategyBaseTVLLimits` instance from the `StrategyFactory`. The `StrategyFactory` only allows a strategy to be deployed once per token, and automatically whitelists newly-deployed strategies. This is further documented in [Strategies](#strategies) below.

Each supported token has its own instance of `StrategyBaseTVLLimits`, has two main functions (`deposit` and `withdraw`), both of which can only be called by the `StrategyManager`. These `StrategyBaseTVLLimits` contracts are fairly simple deposit/withdraw contracts that hold tokens deposited by Stakers. Because these strategies are essentially extensions of the `StrategyManager`, their functions are documented in this file (see [Strategies](#strategies) below).

Note that for the EIGEN/bEIGEN token specifically, the `EigenStrategy` contract is used instead of `StrategyBaseTVLLimits`. Additionally, the EIGEN/bEIGEN token and several LSTs whitelisted prior to the existence of the `StrategyFactory` are blacklisted within the `StrategyFactory` to prevent duplicate strategies from being deployed for these tokens.

#### High-level Concepts

This document organizes methods according to the following themes (click each to be taken to the relevant section):
* [Depositing Into Strategies](#depositing-into-strategies)
* [Withdrawal Processing](#withdrawal-processing)
* [Strategies](#strategies)
* [System Configuration](#system-configuration)

#### Important state variables

* `mapping(address => mapping(IStrategy => uint256)) public stakerStrategyShares`: Tracks the current balance a Staker holds in a given strategy. Updated on deposit/withdraw.
* `mapping(address => IStrategy[]) public stakerStrategyList`: Maintains a list of the strategies a Staker holds a nonzero number of shares in.
    * Updated as needed when Stakers deposit and withdraw: if a Staker has a zero balance in a Strategy, it is removed from the list. Likewise, if a Staker deposits into a Strategy and did not previously have a balance, it is added to the list.
* `mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit`: The `strategyWhitelister` is (as of M2) a permissioned role that can be changed by the contract owner. The `strategyWhitelister` has currently whitelisted 3 `StrategyBaseTVLLimits` contracts in this mapping, one for each supported LST.
* `mapping(IStrategy => bool) public thirdPartyTransfersForbidden`: The `strategyWhitelister` can disable third party transfers for a given strategy. If `thirdPartyTransfersForbidden[strategy] == true`:
    * Users cannot deposit on behalf of someone else (see [`depositIntoStrategyWithSignature`](#depositintostrategywithsignature)). 
    * Users cannot withdraw on behalf of someone else. (see [`DelegationManager.queueWithdrawals`](./DelegationManager.md#queuewithdrawals)) 

#### Helpful definitions

* `stakerStrategyListLength(address staker) -> (uint)`:
    * Gives `stakerStrategyList[staker].length`
    * Used (especially by the `DelegationManager`) to determine whether a Staker has shares in any strategy in the `StrategyManager` (will be 0 if not)
* `uint256 constant MAX_TOTAL_SHARES = 1e38 - 1`
    * The maximum total shares a single strategy can handle. This maximum prevents overflow in offchain services.

---

### Depositing Into Strategies

The following methods are called by Stakers as they (i) deposit LSTs into strategies to receive shares:

* [`StrategyManager.depositIntoStrategy`](#depositintostrategy)
* [`StrategyManager.depositIntoStrategyWithSignature`](#depositintostrategywithsignature)

Withdrawals are performed through the `DelegationManager` (see [`DelegationManager.md`](./DelegationManager.md)).

#### `depositIntoStrategy`

```solidity
function depositIntoStrategy(
    IStrategy strategy, 
    IERC20 token, 
    uint256 amount
)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyNotFrozen(msg.sender)
    nonReentrant
    returns (uint256 shares)
```

Allows a Staker to deposit some `amount` of `token` into the specified `strategy` in exchange for shares of that strategy. The underlying `strategy` must be one of the whitelisted `StrategyBaseTVLLimits` instances, and the `token` parameter corresponds to the actual token being transferred as part of the deposit.

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

This method has a similar purpose as `depositIntoStrategy()`, except it is intended to be used when submitting a deposit on behalf of `staker` which will be credited with the new shares.

*Effects*: See `depositIntoStrategy` above. Additionally:
* The Staker's nonce is incremented

*Requirements*: See `depositIntoStrategy` above. Additionally:
* Caller MUST provide a valid, unexpired signature over the correct fields
* `thirdPartyTransfersForbidden[strategy]` MUST be false

---

### Withdrawal Processing

These methods are callable ONLY by the `DelegationManager`, and are used when processing undelegations and withdrawals:
* [`StrategyManager.removeShares`](#removeshares)
* [`StrategyManager.addShares`](#addshares)
* [`StrategyManager.withdrawSharesAsTokens`](#withdrawsharesastokens)

See [`DelegationManager.md`](./DelegationManager.md) for more context on how these methods are used.

#### `removeShares`

```solidity
function removeShares(
    address staker,
    IStrategy strategy,
    uint256 shares
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a Staker queues a withdrawal (or undelegates, which also queues a withdrawal). The shares are removed while the withdrawal is in the queue, and when the queue completes, the shares will either be re-awarded or withdrawn as tokens (`addShares` and `withdrawSharesAsTokens`, respectively).

The Staker's share balance for the `strategy` is decreased by the removed `shares`. If this causes the Staker's share balance to hit zero, the `strategy` is removed from the Staker's strategy list.

*Entry Points*:
* `DelegationManager.undelegate`
* `DelegationManager.queueWithdrawals`

*Effects*:
* The Staker's share balance for the given `strategy` is decreased by the given `shares`
    * If this causes the balance to hit zero, the `strategy` is removed from the Staker's strategy list

*Requirements*:
* Caller MUST be the `DelegationManager`
* `staker` parameter MUST NOT be zero
* `shares` parameter MUST NOT be zero
* `staker` MUST have at least `shares` balance for the given `strategy`

#### `addShares`

```solidity
function addShares(
    address staker,
    IStrategy strategy,
    uint256 shares
) 
    external 
    onlyDelegationManager
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as "shares" (rather than as the underlying tokens). In this case, the `shares` originally removed (via `removeShares`) are awarded to the `staker` passed in by the `DelegationManager`.

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* The `staker's` share balance for the given `strategy` is increased by `shares`
    * If the prior balance was zero, the `strategy` is added to the `staker's` strategy list

*Requirements*:
* Caller MUST be the `DelegationManager`
* `staker` parameter MUST NOT be zero
* `shares` parameter MUST NOT be zero

#### `withdrawSharesAsTokens`

```solidity
function withdrawSharesAsTokens(
    address recipient,
    IStrategy strategy,
    uint shares,
    IERC20 token
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as the tokens underlying the shares. In this case, the `shares` originally removed (via `removeShares`) are converted to tokens within the `strategy` and sent to the `recipient`.

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* Calls [`StrategyBaseTVLLimits.withdraw`](#strategybasetvllimitswithdraw)

*Requirements*:
* Caller MUST be the `DelegationManager`
* See [`StrategyBaseTVLLimits.withdraw`](#strategybasetvllimitswithdraw)

---

### Strategies

`StrategyBaseTVLLimits` only has two methods of note, and both can only be called by the `StrategyManager`. Documentation for these methods are included below, rather than in a separate file:
* [`StrategyBaseTVLLimits.deposit`](#strategybasetvllimitsdeposit)
* [`StrategyBaseTVLLimits.withdraw`](#strategybasetvllimitswithdraw)

Additionally, using the `StrategyFactory`, anyone can deploy a new `StrategyBaseTVLLimits` instance for a particular token. The `StrategyFactory` manages these deployments and other strategy whitelisting features in the following methods:
* [`StrategyFactory.deployNewStrategy`](#strategyfactorydeploynewstrategy)
* [`StrategyFactory.blacklistTokens`](#strategyfactoryblacklisttokens)
* [`StrategyFactory.whitelistStrategies`](#strategyfactorywhiteliststrategies)
* [`StrategyFactory.setThirdPartyTransfersForbidden`](#strategyfactorysetthirdpartytransfersforbidden)
* [`StrategyFactory.removeStrategiesFromWhitelist`](#strategyfactoryremovestrategiesfromwhitelist)

#### `StrategyBaseTVLLimits.deposit`

```solidity
function deposit(
    IERC20 token, 
    uint256 amount
)
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
* When converted to shares via the strategy's exchange rate:
    * The `amount` of `token` deposited MUST represent at least 1 new share for the depositor
    * The new total shares awarded by the strategy MUST NOT exceed `MAX_TOTAL_SHARES`


#### `StrategyBaseTVLLimits.withdraw`

```solidity
function withdraw(
    address recipient, 
    IERC20 token, 
    uint256 amountShares
)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyStrategyManager
```

The `StrategyManager` calls this method when a queued withdrawal is completed and the withdrawer has specified they would like to convert their withdrawn shares to tokens. 

This method converts the withdrawal shares back into tokens using the strategy's exchange rate. The strategy's total shares are decreased to reflect the withdrawal before transferring the tokens to the `recipient`.

*Entry Points*:
* `DelegationManager.completeQueuedWithdrawal`
* `DelegationManager.completeQueuedWithdrawals`

*Effects*:
* `StrategyBaseTVLLimits.totalShares` is decreased to account for the shares being withdrawn
* `underlyingToken.safeTransfer` is called to transfer the tokens to the `recipient`

*Requirements*:
* Caller MUST be the `StrategyManager`
* Pause status MUST NOT be set: `PAUSED_WITHDRAWALS`
* The passed-in `token` MUST match the strategy's `underlyingToken`
* The `amountShares` being withdrawn MUST NOT exceed the `totalShares` in the strategy
* The tokens represented by `amountShares` MUST NOT exceed the strategy's token balance

#### `StrategyFactory.deployNewStrategy`

```solidity
function deployNewStrategy(IERC20 token)
    external
    onlyWhenNotPaused(PAUSED_NEW_STRATEGIES)
    returns (IStrategy newStrategy)
```

Allows anyone to deploy a new `StrategyBaseTVLLimits` instance that supports deposits/withdrawals using the provided `token`. As part of calling this method, the `StrategyFactory` automatically whitelists the new strategy in the `StrategyManager`.

Note that the `StrategyFactory` only permits ONE strategy deployment per `token`. Once a `token` has an associated strategy deployed via this method, `deployNewStrategy` cannot be used to deploy a strategy for `token` again. Additionally, `deployNewStrategy` will reject any `token` placed onto the `StrategyFactory` blacklist. This feature was added to prevent the deployment of strategies that existed _before_ the `StrategyFactory` was created. For details, see [`StrategyFactory.blacklistTokens`](#strategyfactoryblacklisttokens).

**NOTE: Use caution when deploying strategies for tokens that do not strictly conform to ERC20 standards. Rebasing tokens similar to already-whitelisted LSTs should be supported, but please DYOR if your token falls outside of ERC20 norms.** Specific things to look out for include (but are not limited to): exotic rebasing tokens, tokens that support reentrant behavior (like ERC-777), and other nonstandard ERC20 derivatives.

*Effects*:
* Deploys a new `BeaconProxy` for the `token`, which references the current `StrategyBaseTVLLimits` implementation
* Updates the `tokenStrategy` mapping for the `token`, preventing a second strategy deployment for the same token
* See `StrategyManager.addStrategiesToDepositWhitelist`

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_NEW_STRATEGIES`
* `token` MUST NOT be blacklisted within `StrategyFactory`
* `StrategyFactory` MUST NOT have been used to deploy a strategy for `token` already
* See `StrategyManager.addStrategiesToDepositWhitelist`

#### `StrategyFactory.blacklistTokens`

```solidity
function blacklistTokens(IERC20[] calldata tokens) external onlyOwner
```

Allows the owner to prevent certain `tokens` from having strategies deployed via `StrategyFactory.deployNewStrategy`. This method was added to prevent the deployment of strategies for tokens that already have strategies deployed/whitelisted through other means.

Note that once the owner adds tokens to the blacklist, they cannot be removed. This is a known limitation of the `StrategyFactory`, and can be addressed by upgrading the factory if needed.

*Effects*:
* Adds each token in `tokens` to the `isBlacklisted` mapping

*Requirements*:
* Caller MUST be the owner
* Each passed in `token` MUST NOT already be blacklisted

#### `StrategyFactory.whitelistStrategies`

```solidity
function whitelistStrategies(
    IStrategy[] calldata strategiesToWhitelist,
    bool[] calldata thirdPartyTransfersForbiddenValues
) 
    external 
    onlyOwner
```

Allows the owner to explicitly whitelist strategies in the `StrategyManager`. This method is used as a passthrough for the `StrategyManager.addStrategiesToDepositWhitelist`, in case the owner needs to whitelist strategies not deployed via the `StrategyFactory`.

*Effects*:
* See `StrategyManager.addStrategiesToDepositWhitelist`

*Requirements*:
* Caller MUST be the owner
* See `StrategyManager.addStrategiesToDepositWhitelist`

#### `StrategyFactory.setThirdPartyTransfersForbidden`

```solidity
function setThirdPartyTransfersForbidden(IStrategy strategy, bool value) external onlyOwner
```

Allows the owner to explicitly enable or disable third party transfers in the `StrategyManager`. This method is used as a passthrough for the `StrategyManager.setThirdPartyTransfersForbidden`, in case the owner needs to modify these values.

*Effects*:
* See `StrategyManager.setThirdPartyTransfersForbidden`

*Requirements*:
* Caller MUST be the owner
* See `StrategyManager.setThirdPartyTransfersForbidden`

#### `StrategyFactory.removeStrategiesFromWhitelist`

```solidity
function removeStrategiesFromWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external
```

Allows the owner to remove strategies from the `StrategyManager` strategy whitelist. This method is used as a passthrough for the `StrategyManager.removeStrategiesFromDepositWhitelist`, in case the owner needs to access this method.

*Effects*:
* See `StrategyManager.removeStrategiesFromDepositWhitelist`

*Requirements*:
* Caller MUST be the owner
* See `StrategyManager.removeStrategiesFromDepositWhitelist`

---

### System Configuration

The Strategy Whitelister role has the ability to permit/remove strategies from being depositable via the `StrategyManager`. This role is held by the `StrategyFactory` (which is fully documented in [Strategies](#strategies)). The following methods concern the Strategy Whitelister role and its abilities within the `StrategyManager`:
* [`StrategyManager.setStrategyWhitelister`](#setstrategywhitelister)
* [`StrategyManager.addStrategiesToDepositWhitelist`](#addstrategiestodepositwhitelist)
* [`StrategyManager.removeStrategiesFromDepositWhitelist`](#removestrategiesfromdepositwhitelist)
* [`StrategyManager.setThirdPartyTransfersForbidden`](#setthirdpartytransfersforbidden)

#### `setStrategyWhitelister`

```solidity
function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner
```

Allows the `owner` to update the Strategy Whitelister address. Currently, the Strategy Whitelister role is held by the `StrategyFactory`. See [Strategies](#strategies) for more details.

*Effects*:
* Updates `StrategyManager.strategyWhitelister`

*Requirements*:
* Caller MUST be the `owner`

#### `addStrategiesToDepositWhitelist`

```solidity
function addStrategiesToDepositWhitelist(
    IStrategy[] calldata strategiesToWhitelist,
    bool[] calldata thirdPartyTransfersForbiddenValues
) 
    external 
    onlyStrategyWhitelister
```

Allows the Strategy Whitelister to add any number of strategies to the `StrategyManager` whitelist, and configure whether third party transfers are enabled or disabled for each. Strategies on the whitelist are eligible for deposit via `depositIntoStrategy`.

*Effects*:
* Adds entries to `StrategyManager.strategyIsWhitelistedForDeposit`
* Sets `thirdPartyTransfersForbidden` for each added strategy

*Requirements*:
* Caller MUST be the `strategyWhitelister`

#### `removeStrategiesFromDepositWhitelist`

```solidity
function removeStrategiesFromDepositWhitelist(
    IStrategy[] calldata strategiesToRemoveFromWhitelist
) 
    external 
    onlyStrategyWhitelister
```

Allows the Strategy Whitelister to remove any number of strategies from the `StrategyManager` whitelist. The removed strategies will no longer be eligible for deposit via `depositIntoStrategy`. However, withdrawals for previously-whitelisted strategies may still be initiated and completed, as long as the Staker has shares to withdraw.

*Effects*:
* Removes entries from `StrategyManager.strategyIsWhitelistedForDeposit`

*Requirements*:
* Caller MUST be the `strategyWhitelister`

#### `setThirdPartyTransfersForbidden`

```solidity
function setThirdPartyTransfersForbidden(
    IStrategy strategy,
    bool value
) 
    external 
    onlyStrategyWhitelister
```

Allows the Strategy Whitelister to enable or disable third-party transfers for any `strategy`. If third-party transfers are disabled:
* Deposits via [`depositIntoStrategyWithSiganture`](#depositintostrategywithsignature) are disabled.
* Withdrawals to a different address via [`DelegationManager.queueWithdrawals`](./DelegationManager.md#queuewithdrawals) are disabled.

*Effects*:
* Sets `thirdPartyTransfersForbidden[strategy]`, even if that strategy is not currently whitelisted

*Requirements*:
* Caller MUST be the `strategyWhitelister`
