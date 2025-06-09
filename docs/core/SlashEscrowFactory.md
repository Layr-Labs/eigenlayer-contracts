## SlashEscrowFactory

| File | Notes |
| -------- | -------- |
| [`SlashEscrowFactory.sol`](../../src/contracts/core/SlashEscrowFactory.sol) | single slash escrow factory |
| [`SlashEscrowFactoryStorage.sol`](../../src/contracts/core/SlashEscrowFactory.sol) | state variables |
| [`ISlashEscrowFactory.sol`](../../src/contracts/interfaces/ISlashEscrowFactory.sol) | interface |

SlashEscrow:

| File | Notes |
| -------- | -------- |
| [`SlashEscrow.sol`](../../src/contracts/core/SlashEscrow.sol) | Instance, deployed per `OperatorSet`, `slashId` that stores slashed funds |
| [`ISlashEscrow.sol`](../../src/contracts/interface/ISlashEscrow.sol) | interface |

## Overview
The `SlashEscrowFactory` handles the burning or redistribution of slashed funds out of the EigenLayer protocol. The `SlashEscrowFactory` is responsible for (i) enforcing an escrow delay upon an AVS calling [`slashOperator`](./AllocationManager.md#slashoperator), (ii) deploying the `SlashEscrow` for each slash, and (iii) releasing funds from the escrow contract upon completion of a the escrow delay. 

## Parameterization
* `DEFAULT_BURN_ADDRESS = 0x00000000000000000000000000000000000E16E4`
    * The address to which burnt funds are sent
* `_globalEscrowDelayBlocks = 28800` (4 days in blocks)
    * On testnet this value is `5` (1 minute in blocks)

## Initiating and Releasing Escrow
An escrow is initiated for each slash triggered by an AVS. A slash can contain multiple strategies and is unique per `operatorSet` and `slashId`. 

*[`SlashEscrowFactory.initiateSlashEscrow`](#initiateslashescrow)
*[`SlashEscrowFactory.releaseSlashEscrow`](#releaseslashescrow)
*[`SlashEscrowFactory.releaseSlashEscrowByStrategy`](#releaseslashescrow)

#### `initiateSlashEscrow`
```solidity
/**
 * @notice Locks up a escrow.
 * @param operatorSet The operator set whose escrow is being locked up.
 * @param slashId The slash ID of the escrow that is being locked up.
 * @param strategy The strategy that whose underlying tokens are being redistributed.
 * @dev This function can be called multiple times for a given `operatorSet` and `slashId`.
 */
function initiateSlashEscrow(
    OperatorSet calldata operatorSet, 
    uint256 slashId, 
    IStrategy strategy) 
external onlyStrategyManager;
```
Initiates an escrow for a given `operatorSet`, `slashId`, and `strategy`. This function can be called multiple times in the same transaction by the `StrategyManager`, as a single slash can contain multiple strategies. The `operatorSet`, `slashId`, and `strategy` are each stored in an [EnumerableSet](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/utils/structs/EnumerableSet.sol).

*Effects*:
* If the operatorSet and slashID have not already been set to pending:
    * The `SlashEscrow` contract is [deployed](#deployslashescrow)
    * Adds the `operatorSet `to `pendingOperatorSets`
    * Adds the `slashId` to the operatorSet's `pendingSlashIds`
* Adds the strategy to `pendingStrategiesForSlashId` 
* Emits a `StartEscrow` event

*Requirements*:
* Can only be called by the `StrategyManager`

#### `deploySlashEscrow`

```solidity
/**
 * @notice Deploys a counterfactual `SlashEscrow` if code hasn't already been deployed.
 * @param operatorSet The operator set whose slash escrow is being deployed.
 * @param slashId The slash ID of the slash escrow that is being deployed.
 */
function _deploySlashEscrow(
    OperatorSet calldata operatorSet, 
    uint256 slashId) 
internal;
```

The internal function is called on `initiateEscrow`. A unique slash escrow contract is deployed per `operatorSet` and `slashId`

SlashEscrows are deployed deterministically using [Open Zeppelin's Clones Upgradeable Library](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/f6febd79e2a3a17e26969dd0d450c6ebd64bf459/contracts/proxy/ClonesUpgradeable.sol), which is a minimal, non-upgradeable proxy. The deployment salt is a concatenation of the `operatorSet` and `slashId`, which together are guaranteed to be unique. 

#### `releaseSlashEscrow`

```solidity
/**
 * @notice Releases an escrow by transferring tokens from the `SlashEscrow` to the operator set's redistribution recipient.
 * @param operatorSet The operator set whose escrow is being released.
 * @param slashId The slash ID of the escrow that is being released.
 * @dev The caller must be the escrow recipient, unless the escrow recipient
 * is the default burn address in which case anyone can call.
 * @dev The slash escrow is released once the delay for ALL strategies has elapsed.
 */
function releaseSlashEscrow(
    OperatorSet calldata operatorSet, 
    uint256 slashId)
external onlyWhenNotPaused(PAUSED_RELEASE_ESCROW);

/**
 * @notice Releases an escrow for a single strategy in a slash.
 * @param operatorSet The operator set whose escrow is being released.
 * @param slashId The slash ID of the escrow that is being released.
 * @param strategy The strategy whose escrow is being released.
 * @dev The caller must be the redistribution recipient, unless the redistribution recipient
 * is the default burn address in which case anyone can call.
 * @dev The slash escrow is released once the delay for ALL strategies has elapsed.
 */
function releaseSlashEscrowByStrategy(
    OperatorSet calldata operatorSet,
    uint256 slashId,
    IStrategy strategy
) external;
```
At or after the `getEscrowCompleteBlock`, tokens are transferred from the `SlashEscrow` contract to the [`redistributionRecipient`](./AllocationManager.md#createredistributingoperatorsets) of an `operatorSet`.

For each `strategy` in the `slash`, tokens are transferred from the slash's unique `SlashEscrow` contract to the `redistributionRecipient`.

To accommodate the unlimited number of strategies that can be added to an operatorSet, and a token transfer revert blocking the release of other tokens, a user can release a single strategy from escrow via `releaseSlashEscrowByStrategy`.

*Effects*:
* Call [`StrategyManager.clearBurnOrRedistributableShares`](./StrategyManager.md#clearburnorredistributableshares). This function may have already been called prior and will no-op if so. We call it again for sanity to ensure that all tokens are transferred to the `SlashEscrow` contract. For `releaseEscrowByStrategy` we call the by strategy variant: `StrategyManager.clearBurnOrRedistributableSharesByStrategy`
* For each strategy:
    * Call [`SlashEscrow.releaseTokens`](#releasetokens)
    * Emits an `EscrowComplete`
    * Remove the `strategy` from the `_pendingStrategiesForSlashId`
* If all strategies from an operatorSet/slashId have been released:
    * Remove the `slashId` from `_pendingSlashIds`
    * Delete the complete block for the `slashId`
* If the `operatorSet` has no more pending slashes, remove it from `pendingOperatorSets` 

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_RELEASE_ESCROW`
* The escrow paused status MUST NOT be set: `pauseEscrow`
* If the operatorSet is redistributable, caller MUST be the `redistributionRecipient`
* The escrow delay MUST have elapsed

## SlashEscrow

A [minimal proxy](https://eips.ethereum.org/EIPS/eip-1167) contract deployed from the `SlashEscrowFactory`. This contract releases funds once an escrow has elapsed.

### `releaseTokens`

```solidity
/**
 * @notice Burns or redistributes the underlying tokens of the strategies.
 * @param slashEscrowFactory The factory contract that created the slash escrow.
 * @param slashEscrowImplementation The implementation contract that was used to create the slash escrow.
 * @param operatorSet The operator set that was used to create the slash escrow.
 * @param slashId The slash ID that was used to create the slash escrow.
 * @param recipient The recipient of the underlying tokens.
 * @param strategy The strategy that was used to create the slash escrow.
 */
function releaseTokens(
    ISlashEscrowFactory slashEscrowFactory,
    ISlashEscrow slashEscrowImplementation,
    OperatorSet calldata operatorSet,
    uint256 slashId,
    address recipient,
    IStrategy strategy
) external;
```
Sends this contract's balance for a given `strategy`'s underlyingToken to the `recipient`. The `slashEscrowFactory`, `slashEscrowImplementation`, `operatorSet`, and `slashId` are passed in to validate deployment parameters of the contract, enabling this contract to have no storage. 

*Effects*:
* Transfers the balance of the strategy's `underlyingToken` to `redistributionRecipient` 

*Requirements*:
* Calldata passed in MUST match output of `ClonesUpgradeable.predictDeterministicAddress`
* Caller MUST be the `SlashEscrowFactory`

## System Configuration

The `owner` of the `SlashEscrowFactory` can update a `_globalEscrowDelayBlocks` and a per-strategy `_strategyEscrowDelayBlocks`. For a given strategy, its delay is given by `max(_globalEscrowDelayBlocks, _strategyEscrowDelayBlocks)`. For an escrow with multiple strategies, the total time for it to be released is the maximum of each of its strategy's delays.

The following methods concern the `owner` and its abilities in the `SlashEscrowFactory` 
* [`SlashEscrowFactory.setGlobalEscrowDelay`](#setglobalescrowdelay)
* [`SlashEscrowFactory.setStrategyEscrowDelay`](#setstrategyescrowdelay)

#### `setGlobalEscrowDelay` 

```solidity
/**
 * @notice Sets the delay for the escrow of all strategies underlying tokens globally.
 * @dev This delay setting only applies to new slashes and does not affect existing ones.
 * @param delay The delay for the escrow.
 */
function setGlobalEscrowDelay(
    uint32 delay
) external onlyOwner;
```
Sets the global escrow delay observed by all strategies. *Note: If this value is updated, previously existing escrows will not be affected.*

*Effects*:
* Sets the `_globalEscrowDelayBlocks`
* Emits a `GlobalEscrowDelaySet` event

*Requirements*:
* Caller MUST be the `owner`

#### `setStrategyEscrowDelay`

```solidity
/**
 * @notice Sets the delay for the escrow of a strategies underlying token.
 * @dev The maximum of the strategy and global delay is used
 * @dev This delay setting only applies to new slashes and does not affect existing ones.
 * @param strategy The strategy whose escrow delay is being set.
 * @param delay The delay for the escrow.
 */
function setStrategyEscrowDelay(
    IStrategy strategy, 
    uint32 delay
) external onlyOwner;
```

Sets the delay for a given strategy. *Note: If this value is updated, previously existing escrows will not be affected.*

*Effects*:
* Updates the `_strategyEscrowDelayBlocks` for the given `strategy`
* Emits a `StrategyEscrowDelaySet` event

*Requirements*:
* Caller MUST be the `owner`
