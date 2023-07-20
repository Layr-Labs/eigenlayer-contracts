
# EigenLayer Technical Specification

## Overview
EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services.
**Restaking** is the process of staking an asset that has already been staked in another protocol into EigenLayer. The canonical example of restaking is ETH restaking, in which an existing Ethereum validator restakes the ETH they have staked to secure Ethereum Proof-of-Stake consensus, but restaking also encompasses actions such as depositing Liquid Staked Tokens into EigenLayer.
**Restaked assets** are placed under the control of EigenLayer’s smart contracts, enabling them to act as stake securing additional services, such as rollups, bridges, and data availability networks.
EigenLayer connects stakers who are willing to provide these additional services to consumers – typically protocols or companies – who want secure services with decentralized validator networks. These consumers pay for the services delivered to them, enabling stakers to earn returns on their staked assets, *in addition* to their existing staking rewards. Thus with restaking, stakers can augment their rewards in exchange for the services they opt-in to providing.

These returns provide an economic incentive for stakers to opt-in and act as “operators” for services. In order to disincentivize malicious actions and deliver *cryptoeconomic security* to services, services built on EigenLayer also impose **slashing conditions** in which provably bad behavior is punished, through the 'slashing' of malicious operators' deposited funds.

EigenLayer is built to be permissionless – anyone can join as a staker, consume a service, or even create their own service – with no external approval acquired. We term this **open innovation**. Being an operator for any service on EigenLayer is strictly **opt-in**. Stakers can choose to serve a single service, many (compatible) services, or simply delegate their stake to an operator *whom they trust to not get slashed*, that can earn rewards using the staker's restaked assets (and presumably somehow share the rewards).

New services built on EigenLayer can define their own, arbitrary *slashing conditions*, which allows services to potentially slash their operators for any action that is **on-chain checkable**. In particular, this is compatible with the permissionless ability to launch a new service on EigenLayer only because all services are opt-in; if a staker believes a service has an unsafe slashing mechanism, then they can simply not opt-in to serving that application.

## Actors in the System

### Stakers
A **staker** is any party who has assets deposited into EigenLayer. In general, these could be any mix of ERC20 tokens and/or staked ETH itself (deposited by transferring withdrawal credentials to EigenLayer or depositing to the Beacon Chain through EigenLayer). Stakers can delegate their stake to an operator, or act as an operator themselves.

### Operators
**Operators** in EigenLayer are those users who actually run the software built on top of EigenLayer. Operators register in EigenLayer, allowing stakers to delegate to them, and then opt-in to any mix of services built on top of EigenLayer; each service that an operator chooses to serve may impose its own slashing conditions on the operator.

### Watchers
**NOTE: at present, EigenLayer does not feature any optimistically rolled up claims. This paragraph reflects a potential future state of the system.**

Some operations in EigenLayer are "**optimistically rolled up**". This is a design pattern used where it is either impossible or infeasible to prove that some claim is true, but *easy to check a counterexample that proves the claim is false*. The general pattern is:
1. A "rolled-up" claim is made, asserting that some condition is true.
2. There is a "fraudproof period", during which anyone can *disprove* the claim with a single counterexample. If a claim is disproven, then the original claimant is punished in some way (e.g. by forfeiting some amount or being slashed).
3. If the claim is *not* disproved during the fraudproof period, then it is assumed to be true, and the system proceeds from this assumption.

**Watchers** are parties who passively observe these "rolled up" claims, and step in only in the case of an invalid or false claim. In such a case, an honest watcher will perform the fraudproof, disproving the claim.

### Services / Middleware
We refer to software built on top of EigenLayer as either **services** or **middleware**. Since we anticipate a wide variety of services built on top of EigenLayer, the EigenLayer team has endeavored to make a minimal amount of assumptions about the structure of services.

## Key Assumptions
### Discretization of Services ("Tasks")
We assume that services manage **tasks**. In other words, we assume that services discretize commitments undertaken by operators, with each task defining the time period for which the service's operators' stakes are placed "at stake", i.e. potentially subject to slashing.

### Delegation "Trust Network" Structure
It is assumed that any staker who delegates their stake to an operator is in the same "trust network" as their chosen operator. In other words, the Staker-DelegatedOperator relationship is assumed to have a significant *trust component*. Operators may have the ability to steal the rewards that they earn from the deposited funds of stakers who delegate to them, as well as imposing other negative externalities on those delegated to them.

### Non-Compromise of Trusted Roles
We assume that all trusted roles (multisigs, etc) remain solely in the hands of honest parties.

### Honest Watcher Assumption
**NOTE: at present, EigenLayer does not feature any optimistically rolled up claims. This paragraph reflects a potential future state of the system.**

For any "optimistically-rolled-up" process that relies on fraudproofs (i.e. in which someone makes an "optimistic claim" that can then be *disproven* within some window, and is otherwise treated as true), we **assume there is at least one honest watcher** who will step in to fraudproof false claims when they are made.
We assume that such an honest watcher will fraudproof *all false claims*, regardless of the size and independent of any financial incentive that may or may not be present for the watcher.
Efforts have been made to relax this assumption, but work is still ongoing.

## Overview of Contracts
The `StrategyManager` contract is the primary coordinator for inflows and outflows of tokens to/from EigenLayer itself. The StrategyManager hands restaked assets over to `Strategy` contracts, which may perform targeted management of restaked assets in order to earn returns outside of EigenLayer (e.g. by lending the assets out on a lending protocol) -- more details on `Strategies` to follow.

Any staker in EigenLayer can choose *either* to register as an operator *or* to delegate their restaked assets to an existing operator. These actions are performed on the `DelegationManager` contract. 

Withdrawals and undelegation are handled through the `StrategyManager`. Both *necessitate delays*, since it is infeasible to immediately know whether or not specific restaked funds are "at stake" on any existing tasks created by services. Instead, stakers who wish to withdraw and/or undelegate must go through a *queued withdrawal* process, in which they:
1. Begin the withdrawal, signaling that the funds they are withdrawing should no longer be placed "at stake" on new tasks.
2. Push any necessary updates to middlewares (or wait for someone else to do so), recording the decrease in funds to be placed at stake on new tasks.
3. Complete their withdrawal after an appropriate delay, i.e. once all tasks have been completed upon which the to-be-withdrawn funds were placed at stake.

## Contract-Specific Overview

### StrategyManager
The StrategyManager contract keeps track of all stakers’ deposits, in the form of “shares” in the Strategy contracts. Stakers who wish to deposit ERC20 tokens can do so by calling the StrategyManager, which will transfer the depositor’s tokens to a user-specified Strategy contract, which in turn manages the tokens to generate rewards in the deposited token (or just passively holds them, if the depositor is risk-averse or if the token lacks good reward-generating opportunities).

As the arbiter of share amounts, the StrategyManager is also the main interaction point for withdrawals from EigenLayer. In general, withdrawals from EigenLayer must ensure that restaked assets cannot be withdrawn until they are no longer placed at risk of slashing by securing some service on EigenLayer. To accomplish this, EigenLayer enforces "guaranteed stake updates on withdrawals". The full withdrawal process is outlined in [the withdrawal flow doc](./EigenLayer-withdrawal-flow.md).

Lastly, the StrategyManager processes slashing actions, in which some (or all) of a user's shares are transferred to a specified address. Slashing of this kind should only ever occur as the result of an operator taking a provably malicious action.

## Strategy(s)
Each `Strategy` contract is expected to manage a single, underlying ERC20 token, known as the `underlyingToken`. Each user's holdings in the strategy is expected to be reflected in a number of `shares`, and the strategy is expected to define methods for converting between an amount of underlying tokens and an amount of shares (and vice versa), somewhat similar to an [ERC4626 Vault](https://eips.ethereum.org/EIPS/eip-4626) but without most of the tokenizing aspects of EIP-4626 (e.g. no `transfer` or `transferFrom` functions are expected).
Assets *may* be depositable or withdrawable to a single `Strategy` contract in multiple forms, and the strategy *may* either actively or passively manage the funds.
Since individual users' share amounts are stored in the `StrategyManager` itself, it is generally expected that each strategy's `deposit` and `withdraw` functions are restricted to only be callable by the `StrategyManager` itself.

### DelegationManager
The DelegationManager contract handles delegation of stakers’ deposited funds to “operators”, who actually serve the applications built on EigenLayer. While delegation to someone else is entirely optional, any operator on EigenLayer must also "register as an operator" by calling the `registerAsOperator` function of this contract.

Any staker in EigenLayer may choose to become *either*:
1. an **operator**, allowing other stakers to delegate to them, and potentially earning a share of the funds generated from using the restaked assets of stakers who delegate to them

OR

2. a **delegator**, choosing to allow an operator to use their restaked assets in securing applications built on EigenLayer

Stakers can choose which path they’d like to take by interacting with the DelegationManager contract. Stakers who wish to delegate select an operator whom they trust to use their restaked assets to serve applications, while operators register to allow others to delegate to them, specifying their `OperatorDetails` and (optionally) providing a `metadataURI` to help structure and explain their relationship with any stakers who delegate to them.

#### Storage in DelegationManager

The `DelegationManager` contract relies heavily upon the `StrategyManager` contract. It keeps track of all active operators -- specifically by storing the `Delegation Terms` for each operator -- as well as storing what operator each staker is delegated to.
A **staker** becomes an **operator** by calling `registerAsOperator`. By design, registered as an operator, an address can never "deregister" as an operator in EigenLayer.
The mapping `delegatedTo` stores which operator each staker is delegated to. Querying `delegatedTo(staker)` will return the *address* of the operator that `staker` is delegated to. Note that operators are *always considered to be delegated to themselves*.

DelegationManager defines when an operator is delegated or not, as well as defining what makes someone an operator:
* someone who has registered as an operator *once* is *always* considered to be an operator
* an **operator** is considered to be 'delegated' to themself upon registering as an operator

Similar to withdrawals, **undelegation** in EigenLayer necessitates a delay or clawback mechanism. To elaborate: if a staker is delegated to an operator, and that operator places the staker's assets 'at stake' on some task in which the operator *misbehaves* (i.e. acts in a slashable manner), it is critical that the staker's funds can still be slashed
* stakers can only undelegate by queuing withdrawal(s) for *all of their assets currently deposited in EigenLayer*, ensuring that all existing tasks for which the staker's currently deposited assets are actively at stake are resolved prior to allowing a different operator to place those same assets at stake on other tasks

### Slasher
The `Slasher` contract is the central point for slashing in EigenLayer.
Operators can opt-in to slashing by arbitrary contracts by calling the function `allowToSlash`. A contract with slashing permission can itself revoke its slashing ability *after a specified time* -- named `serveUntil` in the function input -- by calling `recordLastStakeUpdateAndRevokeSlashingAbility`. The time until which `contractAddress` can slash `operator` is stored in `contractCanSlashOperatorUntil[operator][contractAddress]` as a uint32-encoded UTC timestamp, and is set to the `MAX_CAN_SLASH_UNTIL` (i.e. max value of a uint32) when `allowToSlash` is initially called.

At present, slashing in EigenLayer is a multi-step process. When a contract wants to slash an operator, it will call the `freezeOperator` function. Any `contractAddress` for which `contractCanSlashOperatorUntil[operator][contractAddress]` is *strictly greater than the current time* can call `freezeOperator(operator)` and trigger **freezing** of the operator. An operator who is frozen -- *and any staker delegated to them* cannot make new deposits or withdrawals, and cannot complete queued withdrawals, as being frozen signals detection of malicious action and they may be subject to slashing. At present, slashing itself is performed by the owner of the `StrategyManager` contract, who can also 'unfreeze' accounts.

### EigenPodManager
The `EigenPodManager` contract is designed to handle Beacon Chain ETH being staked on EigenLayer. Specifically, it is designed around withdrawal credentials pointed directly to the EigenLayer contracts, i.e. primarily those of "solo stakers". The EigenPodManager creates new EigenPod contracts, and coordinates virtual deposits and withdrawals of shares in an enshrined `beaconChainETH` strategy to and from the StrategyManager. More details on the EigenPodManager and EigenPod contracts can be found in the dedicated [EigenPod Doc](./EigenPods.md).

### EigenPods 
Each staker can deploy a single `EigenPod` contract through the EigenPodManager that allows them to stake ETH into the Beacon Chain and restake their deposits on EigenLayer. A watcher can also prove that an Ethereum validator that is restaked on an EigenPod has a lower balance on the Beacon Chain than its stake in EigenLayer. Finally, EigenPods also facilitate the execution of withdrawals of partially withdrawn rewards from the Beacon Chain on behalf of validators (a major upgrade in the upcoming Capella consensus layer hardfork). Calls are -- in general -- passed from the EigenPod to the EigenPodManager to the StrategyManager, to trigger additional accounting logic within EigenLayer.
EigenPods are deployed using a beacon proxy pattern, allowing simultaneous upgrades of all EigenPods. This upgradeability will likely be necessary in order to more fully integrate Beacon Chain withdrawals through the EigenPods, e.g. if Ethereum upgrades to smart contract-triggered withdrawals.

### BeaconChainOracle
This contract will post periodic Beacon Chain state root updates, for consumption by the EigenPod contracts.
Details TBD.

## High-Level Goals (And How They Affect Design Decisions)
1. Anyone can launch a new service on EigenLayer, permissionlessly
    * all services are opt-in by design, so operators can simply choose to not serve a malicious application
    * operators must signal *specific contracts* that can slash them, potentially limiting the damage that can be done, e.g. by a malicious or poorly-written upgrade to a service's smart contracts
2. Stakers should *not* be able to withdraw any stake that is "active" on a service
    * assuming that services use a "task-denominated" model helps to enable this paradigm
    * the queued withdrawal mechanism is designed to first stop the withdrawn funds from being placed at stake on new tasks, and then to verify when the funds are indeed no longer at stake
    * the undelegation process enforces similar delays -- it is only possible for a staker to undelegate by queuing a withdrawal for all of their assets currently deposited in EigenLayer
