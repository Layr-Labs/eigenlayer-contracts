
# Withdrawal Flow

Withdrawals from EigenLayer are a multi-step process. This is necessary in order to ensure that funds can only be withdrawn once they are no longer placed 'at stake' on an active task of a service built on top of EigenLayer. For more details on the design of withdrawals and how they guarantee this, see the [Withdrawals Design Doc](./Guaranteed-stake-updates.md).

The first step of any withdrawal involves "queuing" the withdrawal itself. The staker who is withdrawing their assets can specify the Strategy(s) they would like to withdraw from, as well as the respective amount of shares to withdraw from each of these strategies. Additionally, the staker can specify the address that will ultimately be able to withdraw the funds. Being able to specify an address different from their own allows stakers to "point their withdrawal" to a smart contract, which can potentially facilitate faster/instant withdrawals in the future.

## Queueing a Withdrawal

![Queuing a Withdrawal](images/EL_queuing_a_withdrawal.png?raw=true "Queuing a Withdrawal")

1. The staker starts a queued withdrawal by calling the `StrategyManager.queueWithdrawal` function.  They set the receiver of the withdrawn funds as `withdrawer` address. Calling `queueWithdrawal` also removes the user's shares in staker-specific storage and the corresponding shares delegated to the operator. Shares in the strategies being withdrawn from, however, technically remain (i.e. the total number of shares in each strategy does not change).  This ensures that the value per share reported by each strategy will remain consistent, and that the shares will continue to accrue gains (or losses!) from any strategy management until the withdrawal is completed.
2. Prior to actually performing the above processing, the StrategyManager calls `Slasher.isFrozen` to ensure that the staker is not 'frozen' in EigenLayer (due to them or the operator they delegate to being slashed).
3. The StrategyManager calls `DelegationManager.decreaseDelegatedShares` to account for any necessary decrease in delegated shares (the DelegationManager contract will not modify its storage if the staker is not an operator and not actively delegated to one).
4. The StrategyManager queries `DelegationManager.delegatedTo` to get the account that the caller is *currently delegated to*. A hash of the withdrawal's details – including the account that the caller is currently delegated to – is stored in the StrategyManager, to record that the queued withdrawal has been created and to store details which can be checked against when the withdrawal is completed.

## Completing a Queued Withdrawal

![Completing a Queued Withdrawal](images/EL_completing_queued_withdrawal.png?raw=true "Completing a Queued Withdrawal")

1. The withdrawer completes the queued withdrawal after the stake is inactive, by calling `StrategyManager.completeQueuedWithdrawal`. They specify whether they would like the withdrawal in shares (to be redelegated in the future) or in tokens (to be removed from the EigenLayer platform), through the `withdrawAsTokens` input flag. The withdrawer must also specify an appropriate `middlewareTimesIndex` which proves that the withdrawn funds are no longer at stake on any active task. The appropriate index can be calculated off-chain and checked using the `Slasher.canWithdraw` function. For more details on this design, see the [Withdrawals Design Doc](./Guaranteed-stake-updates.md).
2. The StrategyManager calls `Slasher.isFrozen` to ensure that the staker who initiated the withdrawal is not 'frozen' in EigenLayer (due to them or the operator they delegate to being slashed). In the event that they are frozen, this indicates that the to-be-withdrawn funds are likely subject to slashing.
3. Depending on the value of the supplied `withdrawAsTokens` input flag:
* If `withdrawAsTokens` is set to 'true', then StrategyManager calls `Strategy.withdraw` on each of the strategies being withdrawn from, causing the withdrawn funds to be transferred from each of the strategies to the withdrawer.
OR
* If `withdrawAsTokens` is set to 'false', then StrategyManager increases the stored share amounts that the withdrawer has in the strategies in question (effectively completing the transfer of shares from the initiator of the withdrawal to the withdrawer), and then calls `DelegationManager.increaseDelegatedShares` to trigger any appropriate updates to delegated share amounts.

## Special Case -- Beacon Chain Full Withdrawals

If a withdrawal includes withdrawing 'Beacon Chain ETH' from EigenLayer, then, it must be limited to *only* Beacon Chain ETH. In addition, before *completing* the withdrawal, the staker must trigger a full withdrawal from the Beacon Chain (as of now this must be originated from the validating keys, but details could change with  Beacon Chain withdrawals) on behalf of enough of their validators to provide sufficient liquidity for their withdrawal.
The staker's EigenPod's balance will eventually increase by the amount withdrawn, and the withdrawals will be reflected in a BeaconChainOracle state root update.
At that point, the staker will prove their full withdrawals (differentiated from partial withdrawals by comparing the amount withdrawn against a hardcoded threshold) credited to the EigenPod against the beacon chain state root via the `verifyAndProcessWithdrawal` function. If the withdrawal's amount is greater than or equal to how much the corresponding Ethereum validator had restaked on EigenLayer, then the excess amount gets instantly withdrawn. If the withdrawal amount is less than the amount restaked on behalf of the validator in EigenLayer, the EigenPod will remove virtual 'beaconChainETH' shares accordingly, by calling the `StrategyManager.recordOvercommittedBeaconChainETH` function.

Once the above is done, then when the withdrawal is completed through calling `StrategyManager.completeQueuedWithdrawal` function (as above), the StrategyManager will pass a call to `EigenPodManager.withdrawRestakedBeaconChainETH`, which will in turn pass a call onto the staker's EigenPod itself, invoking the `EigenPod.withdrawRestakedBeaconChainETH` function and triggering the actual transfer of ETH from the EigenPod to the withdrawer. This final call will only fail if the  full withdrawals made and proven to the EigenPod do not provide sufficient liquidity for the EigenLayer withdrawal to occur.

There exists an edge case in which a staker queues a withdrawal for all (or almost all) of their virtual beaconChainETH shares prior to a call to `StrategyManager.recordOvercommittedBeaconChainETH` -- in this case, once the staker's virtual beaconChainETH shares are decreased to zero, a special `beaconChainETHSharesToDecrementOnWithdrawal` variable is incremented, and in turn when the staker completes their queued withdrawal, the amount will be subtracted from their withdrawal amount. In other words, if the staker incurs a nonzero `beaconChainETHSharesToDecrementOnWithdrawal` amount, then withdrawals of the staker's beaconChainETH shares will prioritize decrementing this amount, prior to sending the staker themselves any funds.
