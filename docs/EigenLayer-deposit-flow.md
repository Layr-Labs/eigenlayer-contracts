
# Deposit Flow

There are 2 main ways in which a staker can deposit new funds into EigenLayer -- depositing into a Strategy through the StrategyManager, and depositing "Beacon Chain ETH" (or proof thereof) through the EigenPodManager.

## Depositing Into a Strategy Through the StrategyManager
The StrategyManager has two functions for depositing funds into Strategy contracts -- `depositIntoStrategy` and `depositIntoStrategyWithSignature`. In both cases, a specified `amount` of an ERC20 `token` is transferred from the caller to a specified Strategy-type contract `strategy`. New shares in the strategy are created according to the return value of `strategy.deposit`; when calling `depositIntoStrategy` these shares are credited to the caller, whereas when calling `depositIntoStrategyWithSignature` the new shares are credited to a specified `staker`, who must have also signed off on the deposit (this enables more complex, contract-mediated deposits, while a signature is required to mitigate the possibility of griefing or dusting-type attacks).
We note as well that deposits cannot be made to a 'frozen' address, i.e. to the address of an operator who has been slashed or to a staker who is actively delegated to a slashed operator.
When performing a deposit through the StrategyManager, the flow of calls between contracts looks like the following:

![Depositing Into EigenLayer Through the StrategyManager -- Contract Flow](images/EL_depositing.png?raw=true "Title")

1. The depositor makes the initial call to either `StrategyManager.depositIntoStrategy` or `StrategyManager.depositIntoStrategyWithSignature`
2. The StrategyManager calls `Slasher.isFrozen` to verify that the recipient (either the caller or the specified `staker` input) is not 'frozen' on EigenLayer
3. The StrategyManager calls the specified `token` contract, transferring specified `amount` of tokens from the caller to the specified `strategy`
4. The StrategyManager calls `strategy.deposit`, and then credits the returned `shares` value to the recipient
5. The StrategyManager calls `DelegationManager.increaseDelegatedShares` to ensure that -- if the recipient has delegated to an operator -- the operator's delegated share amounts are updated appropriately

## Depositing Beacon Chain ETH Through the EigenPodManager
This section covers depositing *new ETH* into the Beacon Chain, with withdrawal credentials pointed to an EigenLayer-controlled contract (an EigenPod) and proving your deposit so it is credited in EigenLayer; this is a multi-step process. For more details on the EigenPods' design in general, see the [EigenPods doc](./EigenPods.md).

The initial deposit of ETH into the Beacon Chain is performed through the EigenPodManager:

![Depositing ETH Into the Beacon Chain Through the EigenPodManager](images/EL_depositing_BeaconChainETH.png?raw=true "Title")

1. The depositor calls `EigenPodManager.stake`
2. The EigenPodManager deploys a new EigenPod for the caller – if they do not already have one – and then calls `EigenPod.stake`
3. The EigenPod deposits ETH into the Beacon Chain through the "ETH2 Deposit Contract". The deposited ETH was supplied as part of the initial call (1), which was passed along to the EigenPod by the EigenPodManager in its own call (2)

After depositing ETH, the depositor waits for the Beacon Chain state root to be updated through EigenLayer's BeaconChainOracle. After an update has been posted that reflects the EigenPod's increased Beacon Chain balance (resulting from the deposit above), then the depositor can call `EigenPod.verifyWithdrawalCredentials` to initiate the following flow:

![Depositing ETH Into the Beacon Chain Through the EigenPodManager Part 2](images/EL_depositing_BeaconChainETH_2.png?raw=true "Title")

1. The depositor calls EigenPod.verifyWithdrawalCredentials on the EigenPod deployed for them above
2. The EigenPod gets the most recent Beacon Chain state root from the EigenPodManager by calling `EigenPodManager.getBeaconChainStateRootAtTimestamp` (the EigenPodManager further passes this query along to the BeaconChainOracle, prior to returning the most recently-posted state root).
3. The EigenPod calls `EigenPodManager.updateBeaconChainBalance` to update the EigenPodManager's accounting of EigenPod balances
4. The EigenPodManager fetches the Slasher's address from the StrategyManager
4. *If the operator has been slashed on the Beacon Chain* (and this is reflected in the latest BeaconChainOracle update), then the EigenPodManager calls `Slasher.freezeOperator` to freeze the staker
5. The EigenPod calls `EigenPodManager.depositBeaconChainETH` to trigger an update in EigenLayer which will reflect the staker's new beacon chain balance
6. The EigenPodManager forwards the information through a call to `StrategyManager.depositBeaconChainETH`, which updates the staker's balance in the enshrined 'beaconChainETHStrategy' after...
7. The StrategyManager makes a call to `Slasher.isFrozen` to verify that the depositor is not 'frozen' in EigenLayer
