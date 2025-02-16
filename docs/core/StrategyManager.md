## StrategyManager

| File | Notes |
| -------- | -------- |
| [`StrategyManager.sol`](../../src/contracts/core/StrategyManager.sol) | singleton share manager hooked into core |
| [`StrategyManagerStorage.sol`](../../src/contracts/core/StrategyManagerStorage.sol) | state variables |
| [`IStrategyManager.sol`](../../src/contracts/interfaces/IStrategyManager.sol) | interface |

StrategyFactory:

| File | Notes |
| -------- | -------- |
| [`StrategyFactory.sol`](../../src/contracts/core/StrategyFactory.sol) | allows deployment of `StrategyBase` for ERC20 tokens |
| [`StrategyBase.sol`](../../src/contracts/strategies/StrategyBase.sol) | deployed as a beacon proxy via `StrategyFactory` |

Individual strategies:

| File | Notes |
| -------- | -------- |
| [`StrategyBaseTVLLimits.sol`](../../src/contracts/strategies/StrategyBaseTVLLimits.sol) | Pre-StrategyFactory, deployed for certain LSTs. Each instances uses a transparent proxy pattern |
| [`EigenStrategy.sol`](../../src/contracts/strategies/EigenStrategy.sol) | One-off strategy deployed to support EIGEN/bEIGEN |

## Overview

The primary function of the `StrategyManager` is to handle _deposit share_ accounting for individual stakers as they deposit and withdraw supported tokens from their corresponding strategies. Note that the `StrategyManager` only handles _deposit shares_. When the word _shares_ is used in this document, it refers to _deposit shares,_ specifically. For an explanation of other share types, see [Shares Accounting - Terminology](./accounting/SharesAccounting.md#terminology).

The `StrategyManager` is responsible for (i) allowing stakers to deposit tokens into the corresponding strategy, (ii) allowing the `DelegationManager` to remove deposit shares when a staker queues a withdrawal, and (iii) allowing the `DelegationManager` to complete a withdrawal by either adding deposit shares back to the staker or withdrawing the deposit shares as tokens via the corresponding strategy.

Any ERC20-compatible token can be supported by deploying a `StrategyBase` instance from the `StrategyFactory`. Under the hood, the `StrategyFactory` uses the beacon proxy pattern and only allows a strategy to be deployed once per token. Deployed strategies are automatically whitelists for deposit in the `StrategyManager`. For details, see [Strategies](#strategies) below.

**Note**: for the EIGEN/bEIGEN token specifically, the `EigenStrategy` contract is used instead of `StrategyBase`. Additionally, the EIGEN/bEIGEN token are blacklisted within the `StrategyFactory` to prevent duplicate strategies from being deployed for these tokens.

**Note**: for certain LST tokens, the `StrategyBaseTVLLimits` contract is used instead of `StrategyBase`. These strategies were deployed before the `StrategyFactory` allowed arbitrary ERC20 strategies. Unlike strategies deployed through the `StrategyFactory`, these `StrategyBaseTVLLimits` contracts use the transparent proxy pattern. For all intents and purposes, these instances behave the same as `StrategyBase` instances deployed from the `StrategyFactory`. The "TVL Limits" capability of these instances has never been used. Any tokens using one of these instances are blacklisted in the `StrategyFactory` to prevent duplicate strategies from being deployed for these tokens.

The `StrategyManager's` responsibilities can be broken down into the following concepts:
* [Depositing Into Strategies](#depositing-into-strategies)
* [Withdrawal Processing](#withdrawal-processing)
* [Burning Slashed Shares](#burning-slashed-shares)
* [Strategies](#strategies)
* [System Configuration](#system-configuration)

## Parameterization

* `MAX_TOTAL_SHARES = 1e38 - 1`
    * The maximum total shares a single strategy can handle. This maximum prevents overflow in offchain services. Deposits that would increase a strategy's total shares beyond this value will revert.
* `MAX_STAKER_STRATEGY_LIST_LENGTH = 32`
    * The maximum number of unique `StrategyManager` strategies a staker can have deposits in. Any deposits that cause this number to be exceeded will revert.
* `DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4`
    * When slashed shares are burned, they are converted to tokens and transferred to this address, where they are unrecoverable.

---

## Depositing Into Strategies

The following methods are called by stakers as they (i) deposit ERC20 tokens into strategies to receive deposit shares:

* [`StrategyManager.depositIntoStrategy`](#depositintostrategy)
* [`StrategyManager.depositIntoStrategyWithSignature`](#depositintostrategywithsignature)

Withdrawals are performed through the `DelegationManager` (see [`DelegationManager.md`](./DelegationManager.md)).

#### `depositIntoStrategy`

```solidity
/**
 * @notice Deposits `amount` of `token` into the specified `strategy` and credits shares to the caller
 * @param strategy the strategy that handles `token`
 * @param token the token from which the `amount` will be transferred
 * @param amount the number of tokens to deposit
 * @return depositShares the number of deposit shares credited to the caller
 * @dev The caller must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
 * 
 * WARNING: Be extremely cautious when depositing tokens that do not strictly adhere to ERC20 standards.
 * Tokens that diverge significantly from ERC20 norms can cause unexpected behavior in token balances for
 * that strategy, e.g. ERC-777 tokens allowing cross-contract reentrancy.
 */
function depositIntoStrategy(
    IStrategy strategy,
    IERC20 token,
    uint256 amount
)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    nonReentrant
    returns (uint256 depositShares)
```

Allows a staker to deposit some `amount` of `token` into the specified `strategy` in exchange for deposit shares in that strategy. The underlying `strategy` must be whitelisted for deposits, meaning it has either been deployed via the `StrategyFactory`, or is an existing `StrategyBaseTVLLimits/EigenStrategy` instance. The `token` parameter should correspond to the strategy's supported token.

The number of shares received is calculated by the `strategy` using an internal exchange rate that depends on the previous number of tokens deposited.

After processing a deposit, the `StrategyManager` forwards the deposit information to the `DelegationManager`, which updates the staker's deposit scaling factor and delegates shares to the staker's operator (if applicable). See [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares) for details.

*Effects*:
* `token.safeTransferFrom`: Transfers `amount` of `token` to `strategy` on behalf of the caller.
* `StrategyManager` awards the staker with the newly-created deposit shares
* See [`StrategyBase.deposit`](#strategybasedeposit)
* See [`DelegationManager.increaseDelegatedShares`](./DelegationManager.md#increasedelegatedshares)

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_DEPOSITS`
* Caller MUST allow at least `amount` of `token` to be transferred by `StrategyManager` to the strategy
* `strategy` in question MUST be whitelisted for deposits.
* See [`StrategyBaseTVLLimits.deposit`](#strategybasedeposit)

#### `depositIntoStrategyWithSignature`

```solidity
/**
 * @notice Deposits `amount` of `token` into the specified `strategy` and credits shares to the `staker`
 * Note tokens are transferred from `msg.sender`, NOT from `staker`. This method allows the caller, using a
 * signature, to deposit their tokens to another staker's balance.
 * @param strategy the strategy that handles `token`
 * @param token the token from which the `amount` will be transferred
 * @param amount the number of tokens to transfer from the caller to the strategy
 * @param staker the staker that the deposited assets will be credited to
 * @param expiry the timestamp at which the signature expires
 * @param signature a valid ECDSA or EIP-1271 signature from `staker`
 * @return depositShares the number of deposit shares credited to `staker`
 * @dev The caller must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
 *
 * WARNING: Be extremely cautious when depositing tokens that do not strictly adhere to ERC20 standards.
 * Tokens that diverge significantly from ERC20 norms can cause unexpected behavior in token balances for
 * that strategy, e.g. ERC-777 tokens allowing cross-contract reentrancy.
 */
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
    nonReentrant
    returns (uint256 depositShares)
```

This method works like `depositIntoStrategy()`, transferring tokens _from the caller_ to the `strategy` contract. Unlike `depositIntoStrategy`, the resulting deposit shares are credited to the passed-in `staker` address, which must sign off on this intent.

*Effects*: See `depositIntoStrategy` above. Additionally:
* The staker's nonce is incremented

*Requirements*: See `depositIntoStrategy` above. Additionally:
* Caller MUST provide a valid, unexpired signature over the correct fields

---

## Withdrawal Processing

These methods are callable ONLY by the `DelegationManager`, and are used when processing undelegations and withdrawals:
* [`StrategyManager.removeDepositShares`](#removedepositshares)
* [`StrategyManager.addShares`](#addshares)
* [`StrategyManager.withdrawSharesAsTokens`](#withdrawsharesastokens)

See [`DelegationManager.md`](./DelegationManager.md) for more context on how these methods are used.

#### `removeDepositShares`

```solidity
/// @notice Used by the DelegationManager to remove a Staker's shares from a particular strategy when entering the withdrawal queue
/// @dev strategy must be beaconChainETH when talking to the EigenPodManager
function removeDepositShares(
    address staker,
    IStrategy strategy,
    uint256 depositSharesToRemove
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a staker queues a withdrawal (or undelegates, which also queues a withdrawal). The staker's deposit shares are removed while the withdrawal is in the queue, and when the withdrawal is completed, the staker can choose whether to be re-awarded the shares, or to convert and receive them as tokens (`addShares` and `withdrawSharesAsTokens`, respectively).

The staker's deposit share balance for the `strategy` is decreased by the removed `depositSharesToRemove`. If this causes the staker's share balance to hit zero, the `strategy` is removed from the staker's strategy list.

Note that the amount of deposit shares removed while in the withdrawal queue may not equal the amount credited when the withdrawal is completed. The staker may receive fewer if slashing occurred; see [`DelegationManager.md`](./DelegationManager.md) for details.

*Effects*:
* Decrease the staker's deposit share balance for the given `strategy` by the given `depositSharesToRemove`
    * If this causes the balance to hit zero, the `strategy` is removed from the staker's strategy list

*Requirements*:
* Caller MUST be the `DelegationManager`
* `depositSharesToRemove` parameter MUST NOT be zero
* `staker` MUST have at least `depositSharesToRemove` balance for the given `strategy`

#### `addShares`

```solidity
/// @notice Used by the DelegationManager to award a Staker some shares that have passed through the withdrawal queue
/// @dev strategy must be beaconChainETH when talking to the EigenPodManager
/// @return existingDepositShares the shares the staker had before any were added
/// @return addedShares the new shares added to the staker's balance
function addShares(
    address staker,
    IStrategy strategy,
    uint256 shares
)
    external
    onlyDelegationManager
    returns (uint256, uint256)
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal "as shares" (rather than as the underlying tokens). 

This method credits the input deposit shares to the staker. In most cases, the input `shares` equal the same shares originally removed when the withdrawal was queued. However, if the staker's operator was slashed, they may receive less. See [`DelegationManager.md`](./DelegationManager.md) for details.

**Note** that if the staker has deposits in `MAX_STAKER_STRATEGY_LIST_LENGTH` unique strategies (and the input `strategy` is not among them), this method will revert. The staker can still choose to complete the withdrawal "as tokens" (See [`DelegationManager.completeQueuedWithdrawal`](./DelegationManager.md#completequeuedwithdrawal)).

*Effects*:
* Increase the `staker's` deposit share balance for the given `strategy` by `shares`
    * If the prior balance was zero, the `strategy` is added to the `staker's` strategy list
* Emit a `Deposit` event

*Requirements*:
* Caller MUST be the `DelegationManager`
* `staker` parameter MUST NOT be zero
* `shares` parameter MUST NOT be zero
* Length of `stakerStrategyList` for the `staker` MUST NOT exceed `MAX_STAKER_STRATEGY_LIST_LENGTH`

#### `withdrawSharesAsTokens`

```solidity
/// @notice Used by the DelegationManager to convert deposit shares to tokens and send them to a staker
/// @dev strategy must be beaconChainETH when talking to the EigenPodManager
/// @dev token is not validated when talking to the EigenPodManager
function withdrawSharesAsTokens(
    address staker,
    IStrategy strategy,
    IERC20 token,
    uint256 shares
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when a queued withdrawal is completed and the withdrawer specifies that they want to receive the withdrawal as the tokens underlying the shares. 

This method directs the `strategy` to convert the input deposit shares to tokens and send them to the `staker`. In most cases, the input `shares` equal the same shares originally removed when the withdrawal was queued. However, if the staker's operator was slashed, they may receive less. See [`DelegationManager.md`](./DelegationManager.md) for details.

*Effects*:
* Calls [`StrategyBase.withdraw`](#strategybasewithdraw)

*Requirements*:
* Caller MUST be the `DelegationManager`
* See [`StrategyBase.withdraw`](#strategybasewithdraw)

---

## Burning Slashed Shares

Slashed shares are marked as burnable, and anyone can call `burnShares` to transfer them to the default burn address. Burnable shares are stored in `burnableShares`, an [EnumerableMap](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/utils/structs/EnumerableMap.sol) with strategy contract addresses as keys and associated view functions. The following methods handle burning of slashed shares:
* [`StrategyManager.increaseBurnableShares`](#increaseburnableshares)
* [`StrategyManager.burnShares`](#burnshares)

#### `increaseBurnableShares`

```solidity
/**
 * @notice Increase the amount of burnable shares for a given Strategy. This is called by the DelegationManager
 * when an operator is slashed in EigenLayer.
 * @param strategy The strategy to burn shares in.
 * @param addedSharesToBurn The amount of added shares to burn.
 * @dev This function is only called by the DelegationManager when an operator is slashed.
 */
function increaseBurnableShares(
    IStrategy strategy, 
    uint256 addedSharesToBurn
)
    external
    onlyDelegationManager
```

The `DelegationManager` calls this method when an operator is slashed, calculating the number of slashable shares and marking them for burning here.

Anyone can then convert the shares to tokens and trigger a burn via `burnShares`. This asynchronous burning method was added to mitigate potential DoS vectors when slashing.

*Effects*:
* Increases `burnableShares` for the given `strategy` by `addedSharesToBurn`

*Requirements*:
* Can only be called by the `DelegationManager`

#### `burnShares`

```solidity
/**
 * @notice Burns Strategy shares for the given strategy by calling into the strategy to transfer
 * to the default burn address.
 * @param strategy The strategy to burn shares in.
 */
function burnShares(
    IStrategy strategy
)
    external
```

Anyone can call this method to burn slashed shares previously added by the `DelegationManager` via `increaseBurnableShares`. This method resets the strategy's burnable shares to 0, and directs the corresponding `strategy` to convert the shares to tokens and transfer them to `DEFAULT_BURN_ADDRESS`, rendering them unrecoverable.

The `strategy` is not called if the strategy had no burnable shares.

*Effects*:
* Resets the strategy's burnable shares to 0
* Calls `withdraw` on the `strategy`, withdrawing shares and sending a corresponding amount of tokens to the `DEFAULT_BURN_ADDRESS`

---

## Strategies

**Concepts**:
* [StrategyBase vs StrategyBaseTVLLimits](#strategybase-vs-strategybasetvllimits)

**Methods**:
* [`StrategyBase.deposit`](#strategybasedeposit)
* [`StrategyBase.withdraw`](#strategybasewithdraw)
* [`StrategyFactory.deployNewStrategy`](#strategyfactorydeploynewstrategy)
* [`StrategyFactory.blacklistTokens`](#strategyfactoryblacklisttokens)
* [`StrategyFactory.whitelistStrategies`](#strategyfactorywhiteliststrategies)
* [`StrategyFactory.removeStrategiesFromWhitelist`](#strategyfactoryremovestrategiesfromwhitelist)

#### `StrategyBase` vs `StrategyBaseTVLLimits`

Before the introduction of the `StrategyFactory`, strategies were manually deployed and whitelisted in the `StrategyManager`. These strategies used `StrategyBaseTVLLimits.sol`, and were deployed using the transparent proxy pattern. With the introduction of the `StrategyFactory`, anyone can create a depositable strategy for any ERC20 (provided it does not have a deployed strategy yet). The `StrategyFactory` deploys beacon proxies, each of which points at a single implementation of `StrategyBase.sol`.

Though these are two different contracts, `StrategyBaseTVLLimits` inherits all its basic functionality from `StrategyBase`, and only implements a "TVL limits" capability on top of them. In short, this additional functionality checks, before each deposit, whether:
1. the deposit amount exceeds a configured `maxPerDeposit`
2. the total token balance after the deposit exceeds a configured `maxTotalDeposits`

To this date, however, these "TVL limits" capabilities have _never_ been used. The values for both of the variables mentioned above have been set to `type(uint).max` since deployment, and there is no plan to change these. Effectively, all instances of `StrategyBaseTVLLimits` behave identically to instances of `StrategyBase` - with the exception being that the former uses a transparent proxy, and the latter a beacon proxy.

#### `StrategyBase.deposit`

```solidity
/**
 * @notice Used to deposit tokens into this Strategy
 * @param token is the ERC20 token being deposited
 * @param amount is the amount of token being deposited
 * @dev This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
 * `depositIntoStrategy` function, and individual share balances are recorded in the strategyManager as well.
 * @dev Note that the assumption is made that `amount` of `token` has already been transferred directly to this contract
 * (as performed in the StrategyManager's deposit functions). In particular, setting the `underlyingToken` of this contract
 * to be a fee-on-transfer token will break the assumption that the amount this contract *received* of the token is equal to
 * the amount that was input when the transfer was performed (i.e. the amount transferred 'out' of the depositor's balance).
 * @dev Note that any validation of `token` is done inside `_beforeDeposit`. This can be overridden if needed.
 * @return newShares is the number of new shares issued at the current exchange ratio.
 */
function deposit(
    IERC20 token,
    uint256 amount
)
    external
    onlyWhenNotPaused(PAUSED_DEPOSITS)
    onlyStrategyManager
    returns (uint256 newShares)
```

The `StrategyManager` calls this method when stakers deposit ERC20 tokens into a strategy. At the time this method is called, the tokens have already been transferred to the strategy. The role of this method is to (i) calculate the number of deposit shares the tokens represent according to the exchange rate, and (ii) add the new deposit shares to the strategy's recorded total shares.

The number of new shares created are returned to the `StrategyManager` to be added to the staker's strategy share balance.

*Effects*:
* `StrategyBaseTVLLimits.totalShares` is increased to account for the new shares created by the deposit

*Requirements*:
* Caller MUST be the `StrategyManager`
* Pause status MUST NOT be set: `PAUSED_DEPOSITS`
* The passed-in `token` MUST match the strategy's `underlyingToken`
* The token amount being deposited MUST NOT exceed the per-deposit cap
* When converted to shares via the strategy's exchange rate:
    * The `amount` of `token` deposited MUST represent at least 1 new share for the depositor
    * The new total shares awarded by the strategy MUST NOT exceed `MAX_TOTAL_SHARES`


#### `StrategyBase.withdraw`

```solidity
/**
 * @notice Used to withdraw tokens from this Strategy, to the `recipient`'s address
 * @param recipient is the address to receive the withdrawn funds
 * @param token is the ERC20 token being transferred out
 * @param amountShares is the amount of shares being withdrawn
 * @dev This function is only callable by the strategyManager contract. It is invoked inside of the strategyManager's
 * other functions, and individual share balances are recorded in the strategyManager as well.
 * @dev Note that any validation of `token` is done inside `_beforeWithdrawal`. This can be overridden if needed.
 */
function withdraw(
    address recipient,
    IERC20 token,
    uint256 amountShares
)
    external
    onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    onlyStrategyManager
```

The `StrategyManager` calls this method to convert a number of deposit shares to tokens, and transfer them to a `recipient`. Typically, this method is invoked as part of the withdrawal completion flow (see [`DelegationManager.completeQueuedWithdrawal`](./DelegationManager.md#completequeuedwithdrawal)). However, this method may also be invoked during the share burning flow (see [`StrategyManager.burnShares`](#burnshares)).

This method converts the deposit shares back into tokens using the strategy's exchange rate. The strategy's total shares are decreased to reflect the withdrawal before transferring the tokens to the `recipient`.

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
/**
 * @notice Deploy a new StrategyBase contract for the ERC20 token, using a beacon proxy
 * @dev A strategy contract must not yet exist for the token.
 * @dev Immense caution is warranted for non-standard ERC20 tokens, particularly "reentrant" tokens
 * like those that conform to ERC777.
 */
function deployNewStrategy(IERC20 token)
    external
    onlyWhenNotPaused(PAUSED_NEW_STRATEGIES)
    returns (IStrategy newStrategy)
```

Allows anyone to deploy a new `StrategyBase` instance that supports deposits/withdrawals using the provided `token`. As part of calling this method, the `StrategyFactory` automatically whitelists the new strategy for deposits via the `StrategyManager`.

Note that the `StrategyFactory` only permits ONE strategy deployment per `token`. Once a `token` has an associated strategy deployed via this method, `deployNewStrategy` cannot be used to deploy a strategy for `token` again. Additionally, `deployNewStrategy` will reject any `token` placed onto the `StrategyFactory` blacklist. This feature was added to prevent the deployment of strategies that existed _before_ the `StrategyFactory` was created. For details, see [`StrategyFactory.blacklistTokens`](#strategyfactoryblacklisttokens).

**NOTE: Use caution when deploying strategies for tokens that do not strictly conform to ERC20 standards. Rebasing tokens similar to already-whitelisted LSTs should be supported, but please DYOR if your token falls outside of ERC20 norms.** Specific things to look out for include (but are not limited to): exotic rebasing tokens, tokens that support reentrant behavior (like ERC-777), and other nonstandard ERC20 derivatives.

*Effects*:
* Deploys a new `BeaconProxy` for the `token`, which references the current `StrategyBase` implementation
* Updates the `tokenStrategy` mapping for the `token`, preventing a second strategy deployment for the same token
* See `StrategyManager.addStrategiesToDepositWhitelist`

*Requirements*:
* Pause status MUST NOT be set: `PAUSED_NEW_STRATEGIES`
* `token` MUST NOT be blacklisted within `StrategyFactory`
* `StrategyFactory` MUST NOT have been used to deploy a strategy for `token` already
* See `StrategyManager.addStrategiesToDepositWhitelist`

#### `StrategyFactory.blacklistTokens`

```solidity
/**
 * @notice Owner-only function to prevent strategies from being created for given tokens.
 * @param tokens An array of token addresses to blacklist.
 */
function blacklistTokens(IERC20[] calldata tokens) 
    external 
    onlyOwner
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
/**
 * @notice Owner-only function to pass through a call to `StrategyManager.addStrategiesToDepositWhitelist`
 */
function whitelistStrategies(
    IStrategy[] calldata strategiesToWhitelist
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

#### `StrategyFactory.removeStrategiesFromWhitelist`

```solidity
/**
 * @notice Owner-only function to pass through a call to `StrategyManager.removeStrategiesFromDepositWhitelist`
 */
function removeStrategiesFromWhitelist(
    IStrategy[] calldata strategiesToRemoveFromWhitelist
) 
    external
    onlyOwner
```

Allows the owner to remove strategies from the `StrategyManager` strategy whitelist. This method is used as a passthrough for the `StrategyManager.removeStrategiesFromDepositWhitelist`, in case the owner needs to access this method.

*Effects*:
* See `StrategyManager.removeStrategiesFromDepositWhitelist`

*Requirements*:
* Caller MUST be the owner
* See `StrategyManager.removeStrategiesFromDepositWhitelist`

---

## System Configuration

The Strategy Whitelister role has the ability to permit/remove strategies from being depositable via the `StrategyManager`. This role is held by the `StrategyFactory` (which is fully documented in [Strategies](#strategies)). The following methods concern the Strategy Whitelister role and its abilities within the `StrategyManager`:
* [`StrategyManager.setStrategyWhitelister`](#setstrategywhitelister)
* [`StrategyManager.addStrategiesToDepositWhitelist`](#addstrategiestodepositwhitelist)
* [`StrategyManager.removeStrategiesFromDepositWhitelist`](#removestrategiesfromdepositwhitelist)

#### `setStrategyWhitelister`

```solidity
/**
 * @notice Owner-only function to change the `strategyWhitelister` address.
 * @param newStrategyWhitelister new address for the `strategyWhitelister`.
 */
function setStrategyWhitelister(address newStrategyWhitelister) external onlyOwner
```

Allows the `owner` to update the Strategy Whitelister address. Currently, the Strategy Whitelister role is held by the `StrategyFactory`. See [Strategies](#strategies) for more details.

*Effects*:
* Updates `StrategyManager.strategyWhitelister`

*Requirements*:
* Caller MUST be the `owner`

#### `addStrategiesToDepositWhitelist`

```solidity
/**
 * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
 * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
 */
function addStrategiesToDepositWhitelist(
    IStrategy[] calldata strategiesToWhitelist
)
    external
    onlyStrategyWhitelister
```

Allows the Strategy Whitelister to add any number of strategies to the `StrategyManager` whitelist, and configure whether third party transfers are enabled or disabled for each. Strategies on the whitelist are eligible for deposit via `depositIntoStrategy`.

*Effects*:
* Adds entries to `StrategyManager.strategyIsWhitelistedForDeposit`

*Requirements*:
* Caller MUST be the `strategyWhitelister`

#### `removeStrategiesFromDepositWhitelist`

```solidity
/**
 * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
 * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
 */
function removeStrategiesFromDepositWhitelist(
    IStrategy[] calldata strategiesToRemoveFromWhitelist
)
    external
    onlyStrategyWhitelister
```

Allows the Strategy Whitelister to remove any number of strategies from the `StrategyManager` whitelist. The removed strategies will no longer be eligible for deposit via `depositIntoStrategy`. However, withdrawals for previously-whitelisted strategies may still be initiated and completed, as long as the staker has shares to withdraw.

*Effects*:
* Removes entries from `StrategyManager.strategyIsWhitelistedForDeposit`

*Requirements*:
* Caller MUST be the `strategyWhitelister`
