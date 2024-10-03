import "./../setup.spec";

methods {
    // Internal, NONDET-summarized EigenPod library functions
    function _.verifyValidatorFields(bytes32, bytes32[] calldata, bytes calldata, uint40) internal => NONDET;
    function _.verifyValidatorBalance(bytes32, uint40, BeaconChainProofs.BalanceProof calldata) internal => NONDET;
    function _.verifyStateRoot(bytes32, BeaconChainProofs.StateRootProof calldata) internal => NONDET;

    // Internal, NONDET-summarized "send ETH" function -- unsound summary used to avoid HAVOC behavior
    // when sending ETH using `Address.sendValue()`
    function _._sendETH(address recipient, uint256 amountWei) internal => NONDET;

    //// External Calls

	// external calls to Slasher
    function _.recordStakeUpdate(address,uint32,uint32,uint256) external => NONDET;

    // external calls to Strategy contracts
    //function _.deposit(address, uint256) external => NONDET;
    //function _.withdraw(address, address, uint256) external => NONDET;


    // envfree functions
    //function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() external returns (uint64) envfree;
    function withdrawableRestakedExecutionLayerGwei() external returns (uint64) envfree;
    //function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;
    function eigenPodManager() external returns (address) envfree;
    function podOwner() external returns (address) envfree;
    //function hasRestaked() external returns (bool) envfree;
    //function mostRecentWithdrawalTimestamp() external returns (uint64) envfree;
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external returns (IEigenPod.ValidatorInfo) envfree;
    //function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) external returns (bool) envfree;
    function validatorStatus(bytes32 pubkeyHash) external returns (IEigenPod.VALIDATOR_STATUS) envfree;
    //function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;

    // harnessed functions
    function get_validatorIndex(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_restakedBalanceGwei(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_podOwnerShares() external returns (int256) envfree;
    function get_withdrawableRestakedExecutionLayerGwei() external returns (uint256) envfree;
    function get_ETH_Balance() external returns (uint256) envfree;
    function get_currentCheckpointTimestamp() external returns (uint64) envfree;
    function get_lastCheckpointTimestamp() external returns (uint64) envfree;
    function validatorIsActive(bytes32) external returns (bool) envfree;   
    function get_validatorLastCheckpointed(bytes32) external returns (uint64) envfree;
    function activeValidatorCount() external returns (uint256) envfree;
    function currentCheckpoint() external returns (IEigenPod.Checkpoint) envfree;
}

function isDuringCheckpoint() returns bool
{
    return get_currentCheckpointTimestamp() > 0;
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