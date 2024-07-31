import "methodsAndAliases.spec";

definition isPrivileged(method f) returns bool =
	f.selector == sig:DelegationManagerHarnessAlias.transferOwnership(address).selector || 
	f.selector == sig:DelegationManagerHarnessAlias.renounceOwnership().selector ||  
	
	f.selector == sig:DelegationManagerHarness.setMinWithdrawalDelayBlocks(uint256).selector || 
	f.selector == sig:DelegationManagerHarness.setStrategyWithdrawalDelayBlocks(address[],uint256[]).selector ||

	((f.contract == DummyEigenPodAAlias || f.contract == DummyEigenPodBAlias) && //there are methods with the same selector on other contracts so we have to distinguish
	(
		f.selector == sig:DummyEigenPodA.recoverTokens(address[],uint256[],address).selector ||
		f.selector == sig:DummyEigenPodA.withdrawRestakedBeaconChainETH(address,uint256).selector ||
		//f.selector == sig:DummyEigenPodA.withdrawNonBeaconChainETHBalanceWei(address,uint256).selector ||
		//f.selector == sig:DummyEigenPodA.activateRestaking().selector ||
		f.selector == sig:DummyEigenPodA.verifyWithdrawalCredentials(uint64,BeaconChainProofs.StateRootProof,uint40[],bytes[],bytes32[][]).selector ||
		f.selector == sig:DummyEigenPodA.stake(bytes,bytes,bytes32).selector
	)) ||

	(f.contract == EigenPodManagerHarnessAlias && //there are methods with the same selector on other contracts so we have to distinguish
	(
		f.selector == sig:EigenPodManagerHarness.addShares(address,uint256).selector ||
		f.selector == sig:EigenPodManagerHarness.removeShares(address,uint256).selector ||
		//f.selector == sig:EigenPodManagerHarness.setDenebForkTimestamp(uint64).selector ||
		f.selector == sig:EigenPodManagerHarness.initialize(address, address, uint256).selector ||    
		f.selector == sig:EigenPodManagerHarness.withdrawSharesAsTokens(address,address,uint256).selector
	)) ||

	f.selector == sig:EigenStrategy.withdraw(address,address,uint256).selector ||
	f.selector == sig:EigenStrategy.deposit(address,uint256).selector ||

	f.selector == sig:StrategyManagerHarness.removeStrategiesFromDepositWhitelist(address[]).selector ||
	f.selector == sig:StrategyManagerHarness.addShares(address,address,address,uint256).selector ||
	f.selector == sig:StrategyManagerHarness.withdrawSharesAsTokens(address,address,uint256,address).selector ||
	f.selector == sig:StrategyManagerHarness.setThirdPartyTransfersForbidden(address,bool).selector ||
	f.selector == sig:StrategyManagerHarness.setStrategyWhitelister(address).selector ||
	f.selector == sig:StrategyManagerHarness.addStrategiesToDepositWhitelist(address[],bool[]).selector ||
	f.selector == sig:StrategyManagerHarness.removeShares(address,address,uint256).selector ||

	f.selector == sig:PauserRegistry.setIsPauser(address,bool).selector ||
	f.selector == sig:PauserRegistry.setUnpauser(address).selector;

definition isIgnoredMethod(method f) returns bool =
	f.selector == sig:certorafallback_0().selector;

function isPrivilegedSender(env e) returns bool
{
	address  sender = e.msg.sender;
	if (
		sender == StrategyManagerHarnessAlias ||
		sender == StrategyManagerHarnessAlias ||
		sender == DummyEigenPodAAlias ||
		sender == DummyEigenPodBAlias ||
		sender == DummyERC20AAlias ||
		sender == DummyERC20BAlias ||
		sender == DelegationManagerHarnessAlias ||
		sender == EigenStrategyAlias ||
		sender == PausableHarnessAlias ||
		sender == ETHPOSDepositMockAlias ||
		sender == ERC1271WalletMockAlias ||
		sender == PauserRegistryAlias)
		return true;
	return false;
}

function timestampsNotFromFuture(env e) returns bool
{
	return e.block.timestamp > 0 &&
		e.block.timestamp >= require_uint256(get_currentCheckpointTimestamp()) &&
		e.block.timestamp >= require_uint256(get_lastCheckpointTimestamp());
}

function validatorDataNotFromFuture(env e, bytes32 validatorHash) returns bool
{
	return e.block.timestamp >= require_uint256(get_validatorLastCheckpointed(validatorHash));
}
	
definition canIncreasePodOwnerShares(method f) returns bool = 
	f.selector == sig:EigenPodManagerHarness.addShares(address,uint256).selector ||
	f.selector == sig:EigenPodManagerHarness.recordBeaconChainETHBalanceUpdate(address,int256).selector ||
	f.selector == sig:EigenPodManagerHarness.withdrawSharesAsTokens(address,address,uint256).selector ||
	f.selector == sig:DummyEigenPodA.startCheckpoint(bool).selector;

definition canDecreasePodOwnerShares(method f) returns bool = 
	f.selector == sig:EigenPodManagerHarness.recordBeaconChainETHBalanceUpdate(address,int256).selector ||
	f.selector == sig:EigenPodManagerHarness.removeShares(address,uint256).selector;

definition canIncreaseBalanceUpdateTimestamp(method f) returns bool = 
	f.selector == sig:DummyEigenPodA.verifyWithdrawalCredentials(uint64, BeaconChainProofs.StateRootProof, uint40[], bytes[], bytes32[][]).selector;

definition canDecreaseBalanceUpdateTimestamp(method f) returns bool = false;
