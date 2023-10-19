# EigenPods Subprotocol Design and Proposed Revisions

## Purpose

The EigenPods subprotocol is a protocol within the broader EigenLayer protocol that handles the restaking of Native Beacon Chain ETH. It allows entities that own validators that are apart of Ethereum's consensus to repoint their withdrawal credentials (explained later) to contracts in the EigenPods subprotocol that have certain mechanisms to ensure safe restaking of native ETH.

The purpose of this document is to detail the EigenPods subprotocol's functionality since it is complex and not well documented besides the (not so) "self documenting code". In addition, there are some proposed revisions that are to be considered.

Ultimately, the protocol's values are security, functionality, and simplicity. Security means no loss of user funds, no unforseen scenarios for AVSs, and more. Functionality means that the protocol is acheiving its goals with respected to its expected function to the best of its ability (see [functionality goals](#Functionality-Goals) for more). Simplicity means ease of understanding for stakers and AVSs, ease of integration for AVSs, and more confidence in the protocol's security.

## Challenges
The reasons native restaking is more complex than token restaking are
- Balances aren't available natively in the EVM. They need to be brought over from the beacon chain, i.e., Ethereum's consensus layer.
- Balances can decrease arbitrarily through beacon chain slashing.
- Funds withdrawn from the beacon chain are directly added to pod balances, no execution layer logic is run on withdrawal.

## Functionality Goals

Aside from the broad protocol goals of security and simplicity, we also want the following functionality from the EigenPods subprotocol:
- Native restakers should be able to restake on EigenLayer
- Native restakers should be able to withdraw the partial withdrawals from their validators
- EigenLayer should be able to slash natively restaked ETH
- AVSs should have an accurate view of the amount of natively restaked ETH a certain operator has delegated to them

## Current Protocol Overview

### Creating EigenPods

Any user that wants to participate in native restaking first deploys an EigenPod contract by calling `createPod()` on the EigenPodManager. This deploys an EigenPod contract which is a BeaconProxy in the Beacon Proxy pattern. The user is called the *pod owner* of the EigenPod they deploy.

### Repointing Withdrawan Credentials: BLS to Execution Changes and Deposits

The precise method by which native restaking occurs. Once an EigenPod is created, a user can make deposits for new validators or switch the withdrawal credentials of their existing validators to their EigenPod address. This makes it so that all funds that are ever withdrawn from the beacon chain on behalf of their validators will be sent to their EigenPod. Since the EigenPod is a contract that runs code as part of the EigenLayer protocol, we engineer it in such a way that the funds that can/will flow through it are restaked on EigenLayer.

The ability to create new validators and repoint existing validators withdrawal credentials to EigenPods is live on Mainnet.

### Beacon Chain Oracle

Since the execution layer cannot synchronously read arbitrary state from the consensus layer, we require a Beacon Chain Oracle to bring in block roots from the consensus layer and put them in consensus layer storage. 

This will either be implemented by Succinct Labs or eventually by Ethereum natively in EIP-4788. The interface will have requests be for a block number and responses be block roots.

### Hysteresis

We want to understimate validator balances on EigenLayer to be tolerant to small slashing events that occur on the beacon chain.

We underestimate validator stake on EigenLayer through the following equation: $\text{eigenlayerBalance} = \text{min}(32, \text{floor}(x-0.75))$. Since a validator's effective balance on the beacon chain can at most 0.25 ETH more than its actual balance on the beacon chain, we subtract 0.75 and floor it for simplicity.

### Proofs of Staked Validators

To convey to EigenLayer that an EigenPod has validator's restaked on it, anyone can submit a proof against a beacon chain state root the proves that a validator has their withdrawal credentials pointed to the pod.

The proof is verified and the EigenPod calls the EigenPodMananger that calls the StrategyManager which records the validators proven balance run through the hysteresis function worth of ETH in the "beaconChainETH" strategy.

Each EigenPod keeps track of all of the validators by the hash of their public key. For each validator, their validator index and current balance in EigenLayer is kept track of.

### Proofs of Partial Withdrawals

We will take a brief aside to explain a simpler part of the protocol, partial withdrawals. Partial withdrawals are small withdrawals of validator yield that occur approximately once every 5 days. Anyone can submit a proof to an EigenPod that one of its validators had a partial withdrawal to the pod (note that this proof completely guarantees that the withdrawal was a partial withdrawal since it is more than the simple balance check done in Rocket Pool, for example). Once the partial withdrawal is proven, it is sent directly to the DelayedWithdrawalRouter to be sent to the EigenPod owner.

Note that partial withdrawals can be withdrawn immediately from the system because they are not staked on EigenLayer, unlike the base validator stake that is restaked on EigenLayer

These proofs are fairly expensive to execute, so we would ideally like to lower the gas cost. We have an optimistic fraudproof-based design that could lower gas costs, but have chosen to avoid the added complexity of implementing this design, because it is merely desirable but not necessary.

### Proofs of Validator Balance Updates

EigenLayer pessimistically assumes the validator has less ETH that they actually have restaked in order for the protocol to have an accurate view of the validator's restaked assets even in the case of an uncorrelated slashing event, for which the penalty is >=1 ETH.

In the case that a validator's balance drops close to or below what is noted in EigenLayer, AVSs need to be notified of that ASAP, in order to get an accurate view of their security.

In the case that a validator's balance, when run through the hysteresis function, is lower or higher than what is restaked on EigenLayer, anyone is allowed to permissionlessly prove that the balance of a certain validator. If the proof is valid, the StrategyManager decrements the pod owners beacon chain ETH shares by however much is staked on EigenLayer and adds the new proven stake.

### Proofs of Full Withdrawals

Full withdrawals occur when a validator completely exits and withdraws from the beacon chain. The EigenPod is then credited with the validators balance at exit. A proof of a full withdrawal can be submitted by anyone to an EigenPod. The EigenPod then notes the new withdrawal and increases the "restaked balance" variable of the pod by the amount of the withdrawal.

If the amount of the withdrawal was greater than what the validator has reestaked on EigenLayer (which most often will be), then the excess is immediately sent to the DelayedWithdrawalRouter to be sent to the pod owner.

Once the "restaked balance" of the pod is incremented, the pod owner is able to queue withdrawals for up to the "restaked balance", decrementing the "restaked balance" by the withdrawal amount. When the withdrawal is completed the pod simply sends a payment to the pod owner for the queued funds. Note that if the withdrawer chooses to recieve the withdrawal as shares, the StrategyManager will increase the "restaked balance" by the withdrawal amount.

