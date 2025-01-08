[middleware-repo]: https://github.com/Layr-Labs/eigenlayer-middleware/
[elip-002]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-002.md

## EigenLayer Docs - v1.0.0 Release

This repo contains the EigenLayer core contracts, which enable restaking of liquid staking tokens (LSTs), beacon chain ETH, and permissionlessly deployed ERC20 Strategies to secure new services called AVSs (actively validated services). For more info on AVSs, check out the EigenLayer middleware contracts [here][middleware-repo].

This document provides an overview of system components, contracts, and user roles and is up-to-date with the latest [ELIP-002][elip-002]. Further documentation on the major system contracts can be found in [/core](./core/).

#### Contents

* [System Components](#system-components)
    * [`EigenPodManager`](#eigenpodmanager)
    * [`StrategyManager`](#strategymanager)
    * [`DelegationManager`](#delegationmanager)
    * [`RewardsCoordinator`](#rewardscoordinator)
    * [`AVSDirectory`](#avsdirectory)
    * [`AllocationManager`](#allocationmanager)
    * [`PermissionController`](#permissioncontroller)
* [Roles and Actors](#roles-and-actors)
* [Common User Flows](#common-user-flows)
    * [Depositing Into EigenLayer](#depositing-into-eigenlayer)
    * [Delegating to an Operator](#delegating-to-an-operator)
    * [Undelegating or Queueing a Withdrawal](#undelegating-or-queueing-a-withdrawal)
    * [Completing a Withdrawal as Shares](#completing-a-withdrawal-as-shares)
    * [Completing a Withdrawal as Tokens](#completing-a-withdrawal-as-tokens)
    * [Withdrawal Processing: Validator Exits](#withdrawal-processing-validator-exits)
    * [Withdrawal Processing: Partial Beacon Chain Withdrawals](#withdrawal-processing-partial-beacon-chain-withdrawals)

### System Components

#### EigenPodManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`EigenPodManager.sol`](../src/contracts/pods/EigenPodManager.sol) | Singleton | Transparent proxy |
| [`EigenPod.sol`](../src/contracts/pods/EigenPod.sol) | Instanced, deployed per-user | Beacon proxy |

These contracts work together to enable native ETH restaking:
* Users deploy `EigenPods` via the `EigenPodManager`, which contain beacon chain state proof logic used to verify a validator's withdrawal credentials and current balances. An `EigenPod's` main role is to serve as the fee recipient and/or withdrawal credentials for one or more of a user's validators.
* The `EigenPodManager` handles `EigenPod` creation and accounting+interactions between users with restaked native ETH and the `DelegationManager`.

See full documentation in:
* [`/core/EigenPodManager.md`](./core/EigenPodManager.md)
* [`/core/EigenPod.md`](./core/EigenPod.md)

#### StrategyManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`StrategyManager.sol`](../src/contracts/core/StrategyManager.sol) | Singleton | Transparent proxy |
| [`StrategyFactory.sol`](../../src/contracts/core/StrategyFactory.sol) | Singleton | Transparent proxy |
| [`StrategyBaseTVLLimits.sol`](../../src/contracts/strategies/StrategyBaseTVLLimits.sol) | Instanced, one per supported token | - Strategies deployed outside the `StrategyFactory` use transparent proxies <br /> - Anything deployed via the `StrategyFactory` uses a Beacon proxy |

These contracts work together to enable restaking for ERC20 tokens supported by EigenLayer:
* The `StrategyManager` acts as the entry and exit point for any supported tokens in EigenLayer. It handles deposits into LST-specific strategies, and manages accounting+interactions between users with restaked LSTs and the `DelegationManager`.
* `StrategyFactory` allows anyone to deploy strategies to support deposits/withdrawals for new ERC20 tokens
* `StrategyBaseTVLLimits` is deployed as multiple separate instances, one for each supported token. When a user deposits into a strategy through the `StrategyManager`, this contract receives the tokens and awards the user with a proportional quantity of deposit shares in the strategy. When a user withdraws, the strategy contract sends the LSTs back to the user.

See full documentation in [`/core/StrategyManager.md`](./core/StrategyManager.md).

#### DelegationManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](../src/contracts/core/DelegationManager.sol) | Singleton | Transparent proxy |

The `DelegationManager` sits between the `EigenPodManager` and `StrategyManager` to manage delegation and undelegation of stakers to operators. Its primary features are to allow users to become operators, to keep track of delegated shares to operators across different strategies, and to manage withdrawals on behalf of stakers via the `EigenPodManager` and `StrategyManager`.

The `DelegationManager` is tightly coupled with the `AllocationManager`. The `DelegationManager` ingests information about slashing as part of managing share accounting for stakers whose operators have been slashed. It also receives directives to slash/burn operator shares when an AVS slashes an operator.

See:
* full documentation in [`/core/DelegationManager.md`](./core/DelegationManager.md)
* share accounting documentation in [`/core/accounting/SharesAccounting.md`](./core/accounting/SharesAccounting.md)

#### RewardsCoordinator

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`RewardsCoordinator.sol`](../src/contracts/core/RewardsCoordinator.sol) | Singleton | Transparent proxy |

The `RewardsCoordinator` is the main entry point of submission and claiming of ERC20 rewards in EigenLayer. It carries out three basic functions:
* AVSs (via the AVS's contracts) submit "rewards submissions" to their registered operators and stakers over a specific time period 
* *Off-chain*, the rewards updater will use each RewardsSubmission time period to apply reward amounts to historical staker/operator stake weights. This is consolidated into a merkle root that is posted *on-chain* to the `RewardsCoordinator`, allowing stakers/operators to claim their allocated rewards.
* Stakers/operators can claim rewards posted by the rewards updater.

See full documentation in [`/core/RewardsCoordinator.md`](./core/RewardsCoordinator.md).

#### AVSDirectory

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`AVSDirectory.sol`](../src/contracts/core/AVSDirectory.sol) | Singleton | Transparent proxy |

##### Note: This contract is left unchanged for backwards compatability. Operator<>AVS Registrations are to be replaced entirely with the `AllocationManager` and this contract will be deprecated in a future release.

Previously, the `AVSDirectory` handled interactions between AVSs and the EigenLayer core contracts. Once registered as an operator in EigenLayer core (via the `DelegationManager`), operators could register with one or more AVSs (via the AVS's contracts) to begin providing services to them offchain. As a part of registering with an AVS, the AVS would record this registration in the core contracts by calling into the `AVSDirectory`. As of the slashing release, this process is now managed by the [`AllocationManager`](#allocationmanager).

See full documentation in [`/core/AVSDirectory.md`](./core/AVSDirectory.md).

For more information on AVS contracts, see the [middleware repo][middleware-repo].

#### AllocationManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`AllocationManager.sol`](../src/contracts/core/AllocationManager.sol) | Singleton | Transparent proxy |

The `AllocationManager` is replaces the AVSDirectory with the introduction of _operator sets_ and slashing. It handles several use cases:
* AVSs can create operator sets and can define the EigenLayer Strategies within them
* Operators can register to or deregister from an AVS's operator sets
* Operators can make slashable security commitments to an operator set by allocating a proportion of their total delegated stake for a Strategy to be slashable. Ex. As an operator, I can allocate 50% of my delegated stETH to be slashable by a specific operator set
* AVSs can slash an operator who has allocated to and is registered for one of the AVS's operator sets

See full documentation in [`/core/AllocationManager.md`](./core/AllocationManager.md).

#### PermissionController

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`PermissionController.sol`](../src/contracts/permissions/PermissionController.sol) | Singleton | Transparent proxy |

The `PermissionController` allows AVSs and operators to delegate the ability to call certain core contract functions to other addresses. This delegation ability is not available to stakers, and is not available in ALL core contract functions.

The following core contracts use the `PermissionController` in certain methods:
* `DelegationManager`
* `AllocationManager`
* `RewardsCoordinator`

See full documentation in [`/permissions/PermissionController.md`](./permissions/PermissionController.md).

---

#### Roles and Actors

To see an example of the user flows described in this section, check out our integration tests: [/src/test/integration](../src/test/integration/).

##### Staker

A staker is any party who has assets deposited (or "restaked") into EigenLayer. Currently, these assets can be:
* Native beacon chain ETH (via the EigenPodManager)
* Arbitrary ERC20s (via the StrategyManager)

Stakers can restake any combination of these: a staker may hold ALL of these assets, or only one of them.

*Flows:*
* Stakers **deposit** assets into EigenLayer via either the StrategyManager (for ERC20s) or the EigenPodManager (for beacon chain ETH)
* Stakers **withdraw** assets via the DelegationManager, *no matter what assets they're withdrawing*
* Stakers **delegate** to an operator via the DelegationManager

##### Operator

An operator is a user who helps run the software built on top of EigenLayer (AVSs). operators register in EigenLayer and allow stakers to delegate to them, then opt in to provide various services built on top of EigenLayer. operators may themselves be stakers; these are not mutually exclusive.

*Flows:*
* Users can **register** as an operator via the DelegationManager
* Operators can **deposit** and **withdraw** assets just like stakers can
* Operators can opt in to providing services for an AVS using that AVS's middleware contracts. See the [EigenLayer middleware][middleware-repo] repo for more details.

---

#### Common User Flows

##### Depositing Into EigenLayer

Depositing into EigenLayer varies depending on whether the staker is depositing Native ETH or LSTs:

![.](./images/Staker%20Flow%20Diagrams/Depositing.png)

##### Delegating to an Operator

![.](./images/Staker%20Flow%20Diagrams/Delegating.png)

##### Undelegating or Queueing a Withdrawal

Undelegating from an operator automatically queues a withdrawal that needs to go through the `DelegationManager's` withdrawal delay. Stakers that want to withdraw can choose to `undelegate`, or can simply call `queueWithdrawals` directly.

![.](./images/Staker%20Flow%20Diagrams/Queue%20Withdrawal.png)

##### Completing a Withdrawal as Shares

This flow is mostly useful if a staker wants to change which operator they are delegated to. The staker first needs to undelegate (see above). At this point, they can delegate to a different operator. However, the new operator will only be awarded shares once the staker completes their queued withdrawal "as shares":

![.](./images/Staker%20Flow%20Diagrams/Complete%20Withdrawal%20as%20Shares.png)

##### Completing a Withdrawal as Tokens

Completing a queued withdrawal as tokens is roughly the same for both native ETH and LSTs. 

However, note that *before* a withdrawal can be completed, native ETH stakers will need to perform additional steps, detailed in the diagrams below. 

![.](./images/Staker%20Flow%20Diagrams/Complete%20Withdrawal%20as%20Tokens.png)

##### `EigenPods`: Processing Validator Exits

If a staker wants to fully withdraw from the beacon chain, they need to perform these additional steps before their withdrawal is completable:

![.](./images/Staker%20Flow%20Diagrams/Validator%20Exits.png)

##### `EigenPods`: Processing Validator Yield

As the staker's `EigenPod` accumulates consensus layer or execution layer yield, the `EigenPod's` balance will increase. The staker can Checkpoint their validator to claim this yield as shares, which can either remain staked in EigenLayer or be withdrawn via the `DelegationManager` withdrawal queue:

![.](./images/Staker%20Flow%20Diagrams/Validator%20Yield.png)