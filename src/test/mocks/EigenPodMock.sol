// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPod.sol";

contract EigenPodMock is IEigenPod, Test {
    /// @notice The max amount of eth, in gwei, that can be restaked per validator
    function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() external view returns(uint64) {}

    function nonBeaconChainETHBalanceWei() external view returns(uint256) {}

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


    function verifyWithdrawalCredentials(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata withdrawalCredentialProofs,
        bytes32[][] calldata validatorFields
    ) external {}

    
    function verifyBalanceUpdates(
        uint64 oracleTimestamp,
        uint40[] calldata validatorIndices,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields
    ) external {}

    function verifyAndProcessWithdrawals(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.WithdrawalProof[] calldata withdrawalProofs,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields,
        bytes32[][] calldata withdrawalFields
    ) external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function activateRestaking() external {}

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function withdrawBeforeRestaking() external {}

    /// @notice Called by the pod owner to withdraw the nonBeaconChainETHBalanceWei
    function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) external {}

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(IERC20[] memory tokenList, uint256[] memory amountsToWithdraw, address recipient) external {}

    function validatorStatus(bytes calldata pubkey) external view returns (VALIDATOR_STATUS){}
    function validatorPubkeyToInfo(bytes calldata validatorPubkey) external view returns (ValidatorInfo memory){}
}