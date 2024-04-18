[middleware-repo]: https://github.com/Layr-Labs/eigenlayer-middleware/

## EigenLayer M2 Docs

This repo contains the EigenLayer core contracts, which enable restaking of liquid staking tokens (LSTs) and beacon chain ETH to secure new services, called AVSs (actively validated services). For more info on AVSs, check out the EigenLayer middleware contracts [here][middleware-repo].

This document provides an overview of system components, contracts, and user roles. Further documentation on the major system contracts can be found in [/core](./core/).

#### Contents

* [System Components](#system-components)
    * [`EigenPodManager`](#eigenpodmanager)
    * [`StrategyManager`](#strategymanager)
    * [`DelegationManager`](#delegationmanager)
    * [`AVSDirectory`](#avsdirectory)
    * [`Slasher`](#slasher)
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
| [`DelayedWithdrawalRouter.sol`](../src/contracts/pods/DelayedWithdrawalRouter.sol) | Singleton | Transparent proxy |
| [`succinctlabs/EigenLayerBeaconOracle.sol`](https://github.com/succinctlabs/telepathy-contracts/blob/main/external/integrations/eigenlayer/EigenLayerBeaconOracle.sol) | Singleton | UUPS proxy | [`0x40B1...9f2c`](https://goerli.etherscan.io/address/0x40B10ddD29a2cfF33DBC420AE5bbDa0649049f2c) |

These contracts work together to enable native ETH restaking:
* Users deploy `EigenPods` via the `EigenPodManager`, which contain beacon chain state proof logic used to verify a validator's withdrawal credentials, balance, and exit. An `EigenPod's` main role is to serve as the withdrawal address for one or more of a user's validators.
* The `EigenPodManager` handles `EigenPod` creation and accounting+interactions between users with restaked native ETH and the `DelegationManager`.
* The `DelayedWithdrawalRouter` imposes a 7-day delay on completing partial beacon chain withdrawals from an `EigenPod`. This is primarily to add a stopgap against a hack being able to instantly withdraw funds (note that all withdrawals from EigenLayer -- other than partial withdrawals earned by validators -- are initiated via the `DelegationManager`).
* The `EigenLayerBeaconOracle` provides beacon chain block roots for use in various proofs. The oracle is supplied by Succinct's Telepathy protocol ([docs link](https://docs.telepathy.xyz/)).

See full documentation in [`/core/EigenPodManager.md`](./core/EigenPodManager.md).

#### StrategyManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`StrategyManager.sol`](../src/contracts/core/StrategyManager.sol) | Singleton | Transparent proxy |
| [`StrategyBaseTVLLimits.sol`](../src/contracts/strategies/StrategyBaseTVLLimits.sol) | One instance per supported LST | Transparent proxy |

These contracts work together to enable restaking for LSTs:
* The `StrategyManager` acts as the entry and exit point for LSTs in EigenLayer. It handles deposits into LST-specific strategies, and manages accounting+interactions between users with restaked LSTs and the `DelegationManager`.
* `StrategyBaseTVLLimits` is deployed as multiple separate instances, one for each supported LST. When a user deposits into a strategy through the `StrategyManager`, this contract receives the tokens and awards the user with a proportional quantity of shares in the strategy. When a user withdraws, the strategy contract sends the LSTs back to the user.

See full documentation in [`/core/StrategyManager.md`](./core/StrategyManager.md).

#### DelegationManager

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`DelegationManager.sol`](../src/contracts/core/DelegationManager.sol) | Singleton | Transparent proxy |

The `DelegationManager` sits between the `EigenPodManager` and `StrategyManager` to manage delegation and undelegation of Stakers to Operators. Its primary features are to allow Operators to register as Operators (`registerAsOperator`), to keep track of shares being delegated to Operators across different strategies, and to manage withdrawals on behalf of the `EigenPodManager` and `StrategyManager`.

See full documentation in [`/core/DelegationManager.md`](./core/DelegationManager.md).

#### AVSDirectory

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`AVSDirectory.sol`](../src/contracts/core/AVSDirectory.sol) | Singleton | Transparent proxy |

The `AVSDirectory` handles interactions between AVSs and the EigenLayer core contracts. Once registered as an Operator in EigenLayer core (via the `DelegationManager`), Operators can register with one or more AVSs (via the AVS's contracts) to begin providing services to them offchain. As a part of registering with an AVS, the AVS will record this registration in the core contracts by calling into the `AVSDirectory`.

See full documentation in [`/core/AVSDirectory.md`](./core/AVSDirectory.md).

For more information on AVS contracts, see the [middleware repo][middleware-repo].

#### Slasher

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`Slasher.sol`](../src/contracts/core/Slasher.sol) | - | - |

<p align="center"><b>
🚧 The Slasher contract is under active development and its interface expected to change. We recommend writing slashing logic without integrating with the Slasher at this point in time. Although the Slasher is deployed, it will remain completely paused/unusable during M2. No contracts interact with it, and its design is not finalized. 🚧
</b><p>

---

#### Roles and Actors

To see an example of the user flows described in this section, check out our integration tests: [/src/test/integration](../src/test/integration/).

##### Staker

A Staker is any party who has assets deposited (or "restaked") into EigenLayer. Currently, these assets can be:
* Native beacon chain ETH (via the EigenPodManager)
* Liquid staking tokens (via the StrategyManager): cbETH, rETH, stETH, ankrETH, OETH, osETH, swETH, wBETH

Stakers can restake any combination of these: a Staker may hold ALL of these assets, or only one of them.

*Flows:*
* Stakers **deposit** assets into EigenLayer via either the StrategyManager (for LSTs) or EigenPodManager (for beacon chain ETH)
* Stakers **withdraw** assets via the DelegationManager, *no matter what assets they're withdrawing*
* Stakers **delegate** to an Operator via the DelegationManager

Unimplemented as of M2:
* Stakers earn yield by delegating to an Operator as the Operator provides services to an AVS
* Stakers are at risk of being slashed if the Operator misbehaves

##### Operator

An Operator is a user who helps run the software built on top of EigenLayer (AVSs). Operators register in EigenLayer and allow Stakers to delegate to them, then opt in to provide various services built on top of EigenLayer. Operators may themselves be Stakers; these are not mutually exclusive.

*Flows:*
* User can **register** as an Operator via the DelegationManager
* Operators can **deposit** and **withdraw** assets just like Stakers can
* Operators can opt in to providing services for an AVS using that AVS's middleware contracts. See the [EigenLayer middleware][middleware-repo] repo for more details.

*Unimplemented as of M2:*
* Operators earn fees as part of the services they provide
* Operators may be slashed by the services they register with (if they misbehave)

---

#### Common User Flows

##### Depositing Into EigenLayer

Depositing into EigenLayer varies depending on whether the Staker is depositing Native ETH or LSTs:

![.](./images/Staker%20Flow%20Diagrams/Depositing.png)

##### Delegating to an Operator

![.](./images/Staker%20Flow%20Diagrams/Delegating.png)

##### Undelegating or Queueing a Withdrawal

Undelegating from an Operator automatically queues a withdrawal that needs to go through the `DelegationManager's` withdrawal delay. Stakers that want to withdraw can choose to `undelegate`, or can simply call `queueWithdrawals` directly.

![.](./images/Staker%20Flow%20Diagrams/Queue%20Withdrawal.png)

##### Completing a Withdrawal as Shares

This flow is mostly useful if a Staker wants to change which Operator they are delegated to. The Staker first needs to undelegate (see above). At this point, they can delegate to a different Operator. However, the new Operator will only be awarded shares once the Staker completes their queued withdrawal "as shares":

![.](./images/Staker%20Flow%20Diagrams/Complete%20Withdrawal%20as%20Shares.png)

##### Completing a Withdrawal as Tokens

Completing a queued withdrawal as tokens is roughly the same for both native ETH and LSTs. 

However, note that *before* a withdrawal can be completed, native ETH stakers will need to perform additional steps, detailed in the "Withdrawal Processing" diagrams below. 

![.](./images/Staker%20Flow%20Diagrams/Complete%20Withdrawal%20as%20Tokens.png)

##### Withdrawal Processing: Validator Exits

If a Staker wants to fully withdraw from the beacon chain, they need to perform these additional steps before their withdrawal is completable:

![.](./images/Staker%20Flow%20Diagrams/Validator%20Exits.png)

##### Withdrawal Processing: Partial Beacon Chain Withdrawals

If a Staker wants to withdraw consensus rewards from the beacon chain, they do NOT go through the `DelegationManager`. This is the only withdrawal type that is not initiated in the `DelegationManager`:

![.](./images/Staker%20Flow%20Diagrams/Partial%20Withdrawals.png)