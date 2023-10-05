
# EigenPods: Handling Beacon Chain ETH

## Overview

This document explains *EigenPods*, the mechanism by which EigenLayer facilitates the restaking of native beacon chain ether.  The EigenPods subprotocol allows entities that own validators that are a part of Ethereum's consensus to repoint their withdrawal credentials (explained later) to contracts in the EigenPods subprotocol.

It is important to contrast this with the restaking of liquid staking derivatives (LSDs) on EigenLayer. EigenLayer will integrate with liquid staking protocols "above the hood", meaning that withdrawal credentials will be pointed to EigenLayer at the smart contract layer rather than the consensus layer. This is because liquid staking protocols need their contracts to be in possession of the withdrawal credentials in order to not have platform risk on EigenLayer. As always, this means that value of liquid staking derivatives carries a discount due to additional smart contract risk.

The architectural design of the EigenPods system is inspired by various liquid staking protocols, particularly Rocket Pool 🚀.

## The EigenPodManager

The EigenPodManager facilitates the higher level functionality of EigenPods and their interactions with the rest of the EigenLayer smart contracts (the StrategyManager and the StrategyManager's owner). Stakers can call the EigenPodManager to create pods (whose addresses are deterministically calculated via the Create2 OZ library) and stake on the Beacon Chain through them. The EigenPodManager also handles the 'overcommitements' of all EigenPods and coordinates processing of overcommitments with the StrategyManager. 

Any user that wants to participate in native restaking first deploys an EigenPod contract by calling createPod() on the EigenPodManager. This deploys an EigenPod contract which is a BeaconProxy in the Beacon Proxy pattern. The user is called the pod owner of the EigenPod they deploy.

This flow is live on Mainnet.

## The EigenPod

The EigenPod is the contract that a staker must set their Ethereum validators' withdrawal credentials to. EigenPods can be created by stakers through a call to the EigenPodManager. EigenPods are deployed using the beacon proxy pattern to have flexible global upgradability for future changes to the Ethereum specification. Stakers can stake for an Ethereum validator when they create their EigenPod, through further calls to their EigenPod, and through parallel deposits to the Beacon Chain deposit contract.

### Beacon State Root Oracle

EigenPods extensively use a Beacon State Root Oracle that will bring beacon state roots into Ethereum for every [`SLOTS_PER_HISTORICAL_ROOT`](https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#time-parameters) slots (currently 8192 slots or ~27 hours) so that all intermediate state roots can be proven against the ones posted on the execution layer.

The following sections are all related to managing Consensus Layer (CL) and Execution Layer (EL) balances via proofs against the beacon state root brought to the EL by the oracle. The below diagram will be of great help to understanding their functioning.

![EigenPods_Architecture drawio](./images/EL_eigenpods_architecture.png)


## EigenPods before Restaking
When EigenPod contracts are initially deployed, the "restaking" functionality is turned off - the withdrawal credential proof has not been initiated yet.  In this "non-restaking" mode, the contract may be used by its owner freely to withdraw validator balances from the beacon chain via the `withdrawBeforeRestaking` function. This function routes the withdrawn balance directly to the `DelayedWithdrawalRouter` contract.  Once the EigenPod's owner verifies that their withdrawal credentials are pointed to the EigenPod via `verifyWithdrawalCredentialsAndBalance`, the `hasRestaked` flag will be set to true and any withdrawals must now be proven for via the `verifyAndProcessWithdrawal` function.  

### Merkle Proof of Correctly Pointed Withdrawal Credentials
After staking an Ethereum validator with its withdrawal credentials pointed to their EigenPod, a staker must show that the new validator exists and has its withdrawal credentials pointed to the EigenPod, by proving it against a beacon state root with a call to `verifyWithdrawalCredentialsAndBalance`. The EigenPod will verify the proof (along with checking for replays and other conditions) and, if the ETH validator's effective balance is proven to be greater than or equal to `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR`, then the EigenPod will pass the validator's effective balance value through its own hysteresis calculation (see [here](#hysteresis)), which effectively underestimates the effective balance of the validator by 1 ETH.  Then a call is made to the EigenPodManager to forward a call to the StrategyManager, crediting the staker with those shares of the virtual beacon chain ETH strategy.

### Effective Restaked Balance - Hysteresis {#hysteresis}
To convey to EigenLayer that an EigenPod has validator(s) restaked on it, anyone can submit a proof against a beacon chain state root the proves that a validator has their withdrawal credentials pointed to the pod.  The proof is verified and the EigenPod calls the EigenPodMananger that calls the StrategyManager which records the validators proven balance run through the hysteresis function worth of ETH in the "beaconChainETH" strategy.  Each EigenPod keeps track of all of the validators by the hash of their public key. For each validator, their validator index and current balance in EigenLayer is kept track of.

### Proofs of Validator Balance Updates
EigenLayer pessimistically assumes the validator has less ETH that they actually have restaked in order for the protocol to have an accurate view of the validator's restaked assets even in the case of an uncorrelated slashing event, for which the penalty is >=1 ETH.
In the case that a validator's balance drops close to or below what is noted in EigenLayer, AVSs need to be notified of that ASAP, in order to get an accurate view of their security. 
In the case that a validator's balance, when run through the hysteresis function, is lower or higher than what is restaked on EigenLayer, anyone is allowed to permissionlessly prove the new balance of the validator, triggering an update in EigenLayer. If the proof is valid, the StrategyManager decrements the pod owners beacon chain ETH shares by however much is staked on EigenLayer and adds the new proven stake, i.e., the strategyManager's view of the staker's shares is an accurate representation of the consensus layer as long as timely balance update proofs are submitted.  

### Proofs of Full/Partial Withdrawals

Whenever a staker withdraws one of their validators from the beacon chain to provide liquidity, they have a few options. Stakers could keep the ETH in the EigenPod and continue staking on EigenLayer, in which case their ETH, when withdrawn to the EigenPod, will not earn any additional Ethereum staking rewards, but may continue to earn rewards on EigenLayer. Stakers could also queue withdrawals on EigenLayer for their virtual beacon chain ETH strategy shares, which will be fullfilled once their obligations to EigenLayer have ended and their EigenPod has enough balance to complete the withdrawal.

In this second case, in order to withdraw their balance from the EigenPod, stakers must provide a valid proof of their full withdrawal or partial withdrawal (withdrawals of beacon chain rewards).  In the case of the former, withdrawals are processed via the queued withdrawal system while in the latter, the balance is instantly withdrawable (as it is technically not being restaked). We distinguish between partial and full withdrawals by checking the `validator.withdrawableEpoch`.  If the `validator.withdrawableEpoch <= executionPayload.slot/SLOTS_PER_EPOCH` then it is classified as a full withdrawal (here `executionPayload` contains the withdrawal being proven).  This is because the `validator.withdrawableEpoch` is set when a validator submits a signed exit transaction.  It is only after this that their withdrawal can be picked up by a sweep and be processed.  In the case of a partial withdrawal, `validator.withdrawableEpoch` is set to FFE (far future epoch). 

We also must prove the `executionPayload.blockNumber > mostRecentWithdrawalBlockNumber`, which is stored in the contract.  `mostRecentWithdrawalBlockNumber` is set when a validator makes a withdrawal in the pre-restaking phase of the EigenPod deployment.  Without this check, a validator can make a partial withdrawal in the EigenPod's pre-restaking mode, withdraw it and then try to prove the same partial withdrawal once withdrawal credentials have been repointed and proven, thus double withdrawing (assuming that they have restaked balance in the EigenPod during the second withdrawal).

In this second case, in order to withdraw their balance from the EigenPod, stakers must provide a valid proof of their full withdrawal against a beacon chain state root. Full withdrawals are differentiated from partial withdrawals by checking against the validator in question's 'withdrawable epoch'; if the validator's withdrawable epoch is less than or equal to the slot's epoch, then the validator has fully withdrawn because a full withdrawal is only processable at or after the withdrawable epoch has passed. Once the full withdrawal is successfully verified, there are 2 cases, each handled slightly differently:

1. If the withdrawn amount is greater than `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR`, then `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR` is held for processing through EigenLayer's normal withdrawal path, while the excess amount above `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR` is marked as instantly withdrawable.

2. If the withdrawn amount is less than `MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR`, then the amount being withdrawn is held for processing through EigenLayer's normal withdrawal path.

### The EigenPod Invariant
The core complexity of the EigenPods system is to ensure that EigenLayer continuously has an accurate picture of the state of the beacon chain balances repointed to it.  In other words, the invariant that governs this system is:
`sum(shares_in_beaconChainETHStrategy) * WEI_TO_GWEI = sum(validator_restakedBalanceGwei) + withdrawableRestakedExecutionLayerGwei`

Essentially this states that the podOwner's shares in the strategyManager's beaconChainETHStrategy must be equal to sum of all the podOwner's restakedBalanceGwei + any withdrawableRestakedExecutionLayerGwei they may have after proving full withdrawals.  


![Beacon Chain Withdrawal Proofs drawio](./images/Withdrawal_Proof_Diagram.png)


## Merkle Proof Specifics

### verifyValidatorFields
This function is used to verify any of the fields, such as withdrawal credentials or slashed status, in the `Validator` container of the Beacon State.  The user provides the validatorFields and the index of the validator they're proving for, and the function verifies this against the Beacon State Root.  

### verifyWithdrawal
This function verifies several proofs related to a withdrawal:
1. It verifies the slot of the withdrawal
2. It verifies the block number of the withdrawal
3. It verifies that the withdrawal fields provided are correct.

### verifyValidatorBalance
The `Validator` container in the Beacon State only contains the effective balance of the validator. The effective balance is used to determine the size of a reward or penalty a validator receives (refer [here](https://kb.beaconcha.in/glossary#current-balance-and-effective-balance) for more information).  The actual balance of the validator is stored in a separate array of `Balance` containers.  Thus we require a separate proof to verify the validator's actual balance, which is verified in `verifyValidatorBalance`.

