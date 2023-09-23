// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPod.sol";

contract EigenPodMock is IEigenPod, Test {
    /// @notice The max amount of eth, in gwei, that can be restaked per validator
    function MAX_VALIDATOR_BALANCE_GWEI() external view returns(uint64) {}

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from beaconchain but not EigenLayer), 
    function withdrawableRestakedExecutionLayerGwei() external view returns(uint64) {}

    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(address owner) external {}

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable {}

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev Called during withdrawal or slashing.
     * @dev Note that this function is marked as non-reentrant to prevent the recipient calling back into it
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) external {}

    /// @notice The single EigenPodManager for EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager) {}

    /// @notice The owner of this EigenPod
    function podOwner() external view returns (address) {}

    /// @notice an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.
    function hasRestaked() external view returns (bool) {}

    /// @notice block timestamp of the most recent withdrawal
    function mostRecentWithdrawalTimestamp() external view returns (uint64) {}

    /// @notice Returns the validatorInfo struct for the provided pubkeyHash
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external view returns (ValidatorInfo memory) {}


    ///@notice mapping that tracks proven withdrawals
    function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) external view returns (bool) {}

    /// @notice This returns the status of a given validator
    function validatorStatus(bytes32 pubkeyHash) external view returns(VALIDATOR_STATUS) {}


    /**
     * @notice This function verifies that the withdrawal credentials of validator(s) owned by the podOwner are pointed to
     * this contract. It also verifies the effective balance  of the validator.  It verifies the provided proof of the ETH validator against the beacon chain state
     * root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.
     * @param oracleBlockNumber is the Beacon Chain blockNumber whose state root the `proof` will be proven against.
     * @param validatorIndices is the list of indices of the validators being proven, refer to consensus specs 
     * @param withdrawalCredentialProofs is an array of proofs, where each proof proves each ETH validator's balance and withdrawal credentials against a beacon chain state root
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs 
     * for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyWithdrawalCredentials(
        uint64 oracleBlockNumber,
        uint40[] calldata validatorIndices,
        bytes[] calldata validatorPubkeys,
        BeaconChainProofs.WithdrawalCredentialProof[] calldata withdrawalCredentialProofs,
        bytes32[][] calldata validatorFields
    ) external {}

    
    /**
     * @notice This function records an overcommitment of stake to EigenLayer on behalf of a certain ETH validator.
     *         If successful, the overcommitted balance is penalized (available for withdrawal whenever the pod's balance allows).
     *         The ETH validator's shares in the enshrined beaconChainETH strategy are also removed from the StrategyManager and undelegated.
     * @param oracleBlockNumber The oracleBlockNumber whose state root the `proof` will be proven against.
     *        Must be within `VERIFY_OVERCOMMITTED_WINDOW_BLOCKS` of the current block.
     * @param validatorIndex is the index of the validator being proven, refer to consensus specs 
     * @param balanceUpdateProofs is the proof of the validator's balance and validatorFields in the balance tree and the balanceRoot to prove for
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs
     * @dev For more details on the Beacon Chain spec, see: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyBalanceUpdate(
        uint64 oracleBlockNumber,
        uint40 validatorIndex,
        BeaconChainProofs.BalanceUpdateProof calldata balanceUpdateProofs,
        bytes32[] calldata validatorFields
    ) external {}

    /**
     * @notice This function records a full withdrawal on behalf of one of the Ethereum validators for this EigenPod
     * @param oracleTimestamp is the timestamp of the oracle slot that the withdrawal is being proven against
     * @param withdrawalProofs is the information needed to check the veracity of the block number and withdrawal being proven
     * @param validatorFieldsProofs is the proof of the validator's fields in the validator tree
     * @param withdrawalFields are the fields of the withdrawal being proven
     * @param validatorFields are the fields of the validator being proven
     */
    function verifyAndProcessWithdrawals(
        uint64 oracleTimestamp,
        BeaconChainProofs.WithdrawalProof[] calldata withdrawalProofs, 
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields,
        bytes32[][] calldata withdrawalFields
    ) external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function activateRestaking() external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function withdrawBeforeRestaking() external {}
    
    /// @notice called by the eigenPodManager to decrement the withdrawableRestakedExecutionLayerGwei 
    /// in the pod, to reflect a queued withdrawal from the beacon chain strategy
    function decrementWithdrawableRestakedExecutionLayerGwei(uint256 amountWei) external {}

    /// @notice called by the eigenPodManager to increment the withdrawableRestakedExecutionLayerGwei 
    /// in the pod, to reflect a completion of a queued withdrawal as shares
    function incrementWithdrawableRestakedExecutionLayerGwei(uint256 amountWei) external {}

    /// @notice Called by the pod owner to withdraw the nonBeaconChainETHBalanceWei
    function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) external {}

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(IERC20[] memory tokenList, uint256[] memory amountsToWithdraw, address recipient) external {}
}