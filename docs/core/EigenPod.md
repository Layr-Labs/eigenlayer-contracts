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

## Proposed Changes (DONT READ ON IF YOU DONT WANT WEEDS)

This protocol is nice in several ways and leaves much to be desired in others. What's nice about it is that it's (hopefully) safe. This is not a small statement, since we've spent a bunch of time designing it and making sure it is safe.

There is simplicity and functionality that is left to be desired.

### beaconChainETHSharesToDecrementOnWithdrawal Complexity

The beaconChainETHSharesToDecrementOnWithdrawal accounting is a complex feature of the protocol that came up numerous times in our code4rena audit. This is a prime example of [systemic complexity](https://vitalik.ca/general/2022/02/28/complexity.html). We've connected these two parts of the protocol in a way that leads to more problems down the line. Ideally, we'd get rid of this complexity.

At the core of this complexity is the fact that withdrawals from the EigenLayer can happen before withdrawals from the beacon chain. Since we allow this to happen, one can queue a withdrawal for more beacon chain ETH *than they actually have on the beacon chain*. This requires that we have a system that accounts for the difference between actual beacon chain balance and the balance that EigenLayer beleives the pod owner has.

Being able to have ones shares not be staked on EigenLayer but still be staked on the beacon chain is a nice feature, but is likely a bunch of added complexity for not so much added value. Instead, when a withdrawal is queued, the amount of beacon chain shares in the withdrawal should be decremented from the "restaked balance" of proven full withdrawals to the pod. This means the "restaked balance" is equal to the shares that a pod owner can queue withdrawals for. This means that whatever is withdrawn from EigenLayer has to previously have been proven to be withdrawn from the beacon chain.

This eliminates the previous scenario where there are no shares to be deducted on EigenLayer because they are in queued withdrawal. Any funds from an overcommitted validator (not fully withdrawn) must still be accounted for in the StrategyManager because they can't have been proven to be full withdrawn yet.

Another case that needs to be handled is when the shares are withdrawn as shares, instead of tokens, they need to be added back to the "restaked balance" of the EigenPod.

This change gets rid of the ability to queue withdrawal from EigenLayer without having withdrawn from the beacon chain, but significantly simplifies the communication that needs to occur between the StrategyManager and the EigenPods protocol.

### Harsh Treatment of Balance Drops

One functionality issue that is undesirable is the fact that there are only 2 effective balances that a validator can have on EigenLayer: 0 ETH and 31 ETH. This is fine for the most part, but is completely intolerant to any correlated slashing at all (in fact we may need to reduce this to 30.5 ETH or something). In the case of a correlated slashing event, it would not be ideal for all slashed validators to have their balances drop to 0.

Instead, we should consider using a [hysteresis](https://eth2book.info/capella/part2/incentives/balances/#hysteresis) inspired balance tracking approach.
![](https://hackmd.io/_uploads/ByJGFnCP2.png)

$\min\left(32,\operatorname{floor}\left(x-0.25\right)\right)$

Where the x axis is the proven balance and the y axis is the amount we consider restaked on EigenLayer. Note that 0.25 is a parameter that should be thought more about and then set.

This is attractive for a couple reasons:

1. Whenever balance drops occur slowly (due to inactivation), proofs of overcommitment can be relayed while the balance is dropping by 0.25 ETH (say when a validators balance is falling from 31.24 ETH down). This allows EigenLayer to be consistently pessimistic, not catching up in this case. Note that we could acheive the same in the current implementation by dropping the beacon chain shares to 0 whenever a balance was proven < 31.25 ETH, but still only counting 31 ETH.
2. It retains granularity on the amount of stake that operators have after they get slashed. This makes EigenLayer still functional in the case of mass slashing.

Although this method is more complicated then what we currently have implemented, its complexity is entirely encapsulated rather than systemic with downstream effects (when implemented with the solution to beaconChainETHSharesToDecrementOnWithdrawal). It has notable functionality enhancements to its benefit.

The change would be to record the output of this function in the StrategyManager, rather than 0, whenever a validator's overcommitment is proven. It would also allow overcommitment to be proven multiple times for the same validator.

### Not Factoring in Balance Increases

Another place the protocol leaves functionality to be desired is that it doesn't take into account validator balance increases after slashing or proof of overcommitment. 

This one is easy to explain. Instead of only allowing proofs of overcommitment, we should consider allowing proofs of *undercommitment*. If a validator's balance is every greater than 31 ETH and they are considered overcommitted, or their hysteresis function above shows an amount lower than it would with current balances, they can prove that they are undercommitted and have their effective balance increased on EigenLayer.

This would require 2 small changes. It would require the storage of the validators effective balance in the pod (to subtract from the StrategyManager before adding new proven balances) and it would require an extra function for proving undercommitment.

We should make the first of these changes in the EigenPod just to be forwards compatible with this behaviour, but we should discuss the merits and downstream effects of this solution more before going ahead with implementing it.

### Using Validator Index Instead of Validator Pubkey

[EIP-7002](https://github.com/ethereum/EIPs/pull/7002), which allows smart contract triggered validator exits requires that the contract specifies the pubkey of the validator. For this reason, instead of having noting down the status of each validator based on their index, we should consider switching to:

```solidity
struct ValidatorInfo {
    uint64 index;
    uint64 eigenlayerBalance;
    uint32 lastWithdrawalBlockNumber;
}
mapping(bytes32 => ValidatorInfo) public pubkeyHashToValidatorInfo;
```

### Lack of Incentives for Overcommitment Proofs

Ok. This is the controversial one. There is no explicit incentive for overcommitment proofs. Note that the reason for overcommitment proofs. The implicit incentives for overcommitment proofs are

1. Layr Labs, Inc. would submit overcommitment proofs since AVSs having an accurate view of the amount of stake securing them is integral to the EigenLayer platform in which they hold a lot of emotional and financial value
2. AVSs would submit overcommitment proofs because they are the specific party that is harmed by overestimating the stake they have securing them

The first reason is a centralized party entrusted with the safety of the system. This is not ideal right now and untenable once governance is decentralized.

The second reason is subject to the tragedy of the commons, high communication barrier, and general unlikeliness to actually get running in practice. It's bad UX for AVSs as well. In terms for propagating the updates from the StrategyManager to the AVS, that is up to the AVS and we should think about solutions for that.

Drawing an analogy to a similar problem encountered in a different part of blockchain applications:

1. BGD Labs could submit all liquidiations on AAVE because solvency is integral to the AAVE platform in which they hold a lot of emotional and financial value
2. Depositors could submit liquididations on AAVE because they are the specific party that is harmed by not selling the assets dropping in value

Both solutions are not optimal for the same reasons described above. Instead, they went with a solution that incentivised a permissionless set of actors with the funds of those harming the platform.

In the same way, we should consider permissionlessly incentivising overcommitment proofs by slashing a small portion of validator funds whenever an overcommitment fraud proof is submitted. This would make it permissionless and incentivised to run a "watcher" (which would be very easy to run alongside a beaconchain validator), leading to the fastest, most robust system for overcommitment proofs.

The hard part of rewarding overcommitment provers is that the EigenPod may not have enough funds for the reward!

One way to account for this is to keep track of the "total penalties" in the pod, which is increased 0.5 ETH for each overcommitment proof (for simplicity). In addition, the overcommitment prover's "penalties owed" for teh EigenPod is increased by 0.25 ETH in a new contract, the EigenPodPenaltyManager. We essentially burn 0.25 ETH to disincentivise validators getting slashed on the beacon chain and submitting their own overcommitment proof. Finally, we factor the penalties into the the amount which the StrategyManager considered restaked by the validator (it is subtracted).

Whenever withdrawals are proven (remember this is permissionless), partial or full, we decrement from the "total penalties" before incrementing "restaked balance" or sending partial withdrawals to the pod owner through the DelayedWithdrawalRouter. This means that penalties are paid before the pod owner is. When decremented, the penalties are sent the the EigenPodPenaltyManager on behalf of the pod. Entities that are owed penalties then have the ability to withdraw whatever they can. Note that this isn't perfect because it leads to many entities fighting over the same funds in case withdrawals occur rarely on the EigenPod.

This substantially increases the complexity of the system. Separating the EigenPodPenaltyManager from the EigenPod makes them easy to test independently. The new logic that needs to be tested on the EigenPods system is 
- Penalties are correctly factored into the share update on the StrategyManager
- Penalties are always paid before partial withdrawals or "restaked balance" is incremented

This is an important step in the automation and trustlessness of the protocol. We should decide whether this is something we want to implement now.

### Other notes

There were many issues brought up relevant to slashing on the code4rena audit. My assessment is that all of these are moot if the beaconChainETHSharesToDecrementOnWithdrawal change is made.

### Meeting Notes

- Implement beaconChainETHSharesToDecrementOnWithdrawal change as recommended
- Implement storing validators based on pubkeyHash
    - Per validator, store the eigenLayerBalance and the index
    - TODO: Figure out rest of the struct
- Implement hysteresis proposal as recommended both over and undercommitment proofs (validatorBalanceUpdate)
    - Debate constant to subtract before flooring
    - Need compare to stored eigenLayerBalance to determine how much of the full withdrawal is instantly withdrawable
- Abandon penalties and rewards for overcommitment proofs. Too much complexity for now. This can be unincentivised for mainnet for a while.
- Come up with interface for proofs from Succinct
- Propose invariants for testings/formal verif/monitoring

## Invariants
- 