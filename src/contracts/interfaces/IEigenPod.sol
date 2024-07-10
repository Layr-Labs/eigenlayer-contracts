// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../libraries/BeaconChainProofs.sol";
import "./IEigenPodManager.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title The implementation contract used for restaking beacon chain ETH on EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice The main functionalities are:
 * - creating new ETH validators with their withdrawal credentials pointed to this contract
 * - proving from beacon chain state roots that withdrawal credentials are pointed to this contract
 * - proving from beacon chain state roots the balances of ETH validators with their withdrawal credentials
 *   pointed to this contract
 * - updating aggregate balances in the EigenPodManager
 * - withdrawing eth when withdrawals are initiated
 * @dev Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
 *   to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts
 */
interface IEigenPod {
    enum VALIDATOR_STATUS {
        INACTIVE, // doesnt exist
        ACTIVE, // staked on ethpos and withdrawal credentials are pointed to the EigenPod
        WITHDRAWN // withdrawn from the Beacon Chain

    }

    struct ValidatorInfo {
        // index of the validator in the beacon chain
        uint64 validatorIndex;
        // amount of beacon chain ETH restaked on EigenLayer in gwei
        uint64 restakedBalanceGwei;
        //timestamp of the validator's most recent balance update
        uint64 mostRecentBalanceUpdateTimestamp;
        // status of the validator
        VALIDATOR_STATUS status;
    }

    /**
     * @notice struct used to store amounts related to proven withdrawals in memory. Used to help
     * manage stack depth and optimize the number of external calls, when batching withdrawal operations.
     */
    struct VerifiedWithdrawal {
        // amount to send to a podOwner from a proven withdrawal
        uint256 amountToSendGwei;
        // difference in shares to be recorded in the eigenPodManager, as a result of the withdrawal
        int256 sharesDeltaGwei;
    }

    enum PARTIAL_WITHDRAWAL_CLAIM_STATUS {
        REDEEMED,
        PENDING,
        FAILED
    }

    /// @notice Emitted when an ETH validator stakes via this eigenPod
    event EigenPodStaked(bytes pubkey);

    /// @notice Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod
    event ValidatorRestaked(uint40 validatorIndex);

    /// @notice Emitted when an ETH validator's  balance is proven to be updated.  Here newValidatorBalanceGwei
    //  is the validator's balance that is credited on EigenLayer.
    event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei);

    /// @notice Emitted when an ETH validator is prove to have withdrawn from the beacon chain
    event FullWithdrawalRedeemed(
        uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei
    );

    /// @notice Emitted when a partial withdrawal claim is successfully redeemed
    event PartialWithdrawalRedeemed(
        uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei
    );

    /// @notice Emitted when restaked beacon chain ETH is withdrawn from the eigenPod.
    event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount);

    /// @notice Emitted when podOwner enables restaking
    event RestakingActivated(address indexed podOwner);

    /// @notice Emitted when ETH is received via the `receive` fallback
    event NonBeaconChainETHReceived(uint256 amountReceived);

    /// @notice Emitted when ETH that was previously received via the `receive` fallback is withdrawn
    event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn);

    /// @notice The max amount of eth, in gwei, that can be restaked per validator
    function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() external view returns (uint64);

    /// @notice the amount of execution layer ETH in this contract that is staked in EigenLayer (i.e. withdrawn from beaconchain but not EigenLayer),
    function withdrawableRestakedExecutionLayerGwei() external view returns (uint64);

    /// @notice any ETH deposited into the EigenPod contract via the `receive` fallback function
    function nonBeaconChainETHBalanceWei() external view returns (uint256);

    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(address owner) external;

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev The podOwner must have already proved sufficient withdrawals, so that this pod's `withdrawableRestakedExecutionLayerGwei` exceeds the
     * `amountWei` input (when converted to GWEI).
     * @dev Reverts if `amountWei` is not a whole Gwei amount
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) external;

    /// @notice The single EigenPodManager for EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager);

    /// @notice The owner of this EigenPod
    function podOwner() external view returns (address);

    /// @notice an indicator of whether or not the podOwner has ever "fully restaked" by successfully calling `verifyCorrectWithdrawalCredentials`.
    function hasRestaked() external view returns (bool);

    /**
     * @notice The latest timestamp at which the pod owner withdrew the balance of the pod, via calling `withdrawBeforeRestaking`.
     * @dev This variable is only updated when the `withdrawBeforeRestaking` function is called, which can only occur before `hasRestaked` is set to true for this pod.
     * Proofs for this pod are only valid against Beacon Chain state roots corresponding to timestamps after the stored `mostRecentWithdrawalTimestamp`.
     */
    function mostRecentWithdrawalTimestamp() external view returns (uint64);

    /// @notice Returns the validatorInfo struct for the provided pubkeyHash
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external view returns (ValidatorInfo memory);

    /// @notice Returns the validatorInfo struct for the provided pubkey
    function validatorPubkeyToInfo(bytes calldata validatorPubkey) external view returns (ValidatorInfo memory);

    ///@notice mapping that tracks proven withdrawals
    function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) external view returns (bool);

    /// @notice This returns the status of a given validator
    function validatorStatus(bytes32 pubkeyHash) external view returns (VALIDATOR_STATUS);

    /// @notice This returns the status of a given validator pubkey
    function validatorStatus(bytes calldata validatorPubkey) external view returns (VALIDATOR_STATUS);

    /**
     * @notice This function verifies that the withdrawal credentials of validator(s) owned by the podOwner are pointed to
     * this contract. It also verifies the effective balance  of the validator.  It verifies the provided proof of the ETH validator against the beacon chain state
     * root, marks the validator as 'active' in EigenLayer, and credits the restaked ETH in Eigenlayer.
     * @param oracleTimestamp is the Beacon Chain timestamp whose state root the `proof` will be proven against.
     * @param validatorIndices is the list of indices of the validators being proven, refer to consensus specs
     * @param withdrawalCredentialProofs is an array of proofs, where each proof proves each ETH validator's balance and withdrawal credentials
     * against a beacon chain state root
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs
     * for details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyWithdrawalCredentials(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata withdrawalCredentialProofs,
        bytes32[][] calldata validatorFields
    ) external;

    /**
     * @notice This function records an update (either increase or decrease) in the pod's balance in the StrategyManager.
     *            It also verifies a merkle proof of the validator's current beacon chain balance.
     * @param oracleTimestamp The oracleTimestamp whose state root the `proof` will be proven against.
     *        Must be within `VERIFY_BALANCE_UPDATE_WINDOW_SECONDS` of the current block.
     * @param validatorIndices is the list of indices of the validators being proven, refer to consensus specs
     * @param validatorFieldsProofs proofs against the `beaconStateRoot` for each validator in `validatorFields`
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs
     * @dev For more details on the Beacon Chain spec, see: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyBalanceUpdates(
        uint64 oracleTimestamp,
        uint40[] calldata validatorIndices,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields
    ) external;

    /**
     * @notice This function records full and partial withdrawals on behalf of one of the Ethereum validators for this EigenPod
     * @param oracleTimestamp is the timestamp of the oracle slot that the withdrawal is being proven against
     * @param withdrawalProofs is the information needed to check the veracity of the block numbers and withdrawals being proven
     * @param validatorFieldsProofs is the proof of the validator's fields' in the validator tree
     * @param withdrawalFields are the fields of the withdrawals being proven
     * @param validatorFields are the fields of the validators being proven
     */
    function verifyAndProcessWithdrawals(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.WithdrawalProof[] calldata withdrawalProofs,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields,
        bytes32[][] calldata withdrawalFields
    ) external;

    /**
     * @notice Called by the pod owner to activate restaking by withdrawing
     * all existing ETH from the pod and preventing further withdrawals via
     * "withdrawBeforeRestaking()"
     */
    function activateRestaking() external;

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function withdrawBeforeRestaking() external;

    /// @notice Called by the pod owner to withdraw the nonBeaconChainETHBalanceWei
    function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) external;

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(IERC20[] memory tokenList, uint256[] memory amountsToWithdraw, address recipient) external;
}
