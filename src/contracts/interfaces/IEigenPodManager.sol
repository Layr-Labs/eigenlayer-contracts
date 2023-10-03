// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "./IETHPOSDeposit.sol";
import "./IStrategyManager.sol";
import "./IDelegationManager.sol";
import "./IEigenPod.sol";
import "./IBeaconChainOracle.sol";
import "./IPausable.sol";
import "./ISlasher.sol";
import "./IStrategy.sol";

/**
 * @title Interface for factory that creates and manages solo staking pods that have their withdrawal credentials pointed to EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */

interface IEigenPodManager is IPausable {
    /**
     * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
     * In functions that operate on existing queued withdrawals -- e.g. `startQueuedWithdrawalWaitingPeriod` or `completeQueuedWithdrawal`,
     * the data is resubmitted and the hash of the submitted data is computed by `calculateWithdrawalRoot` and checked against the
     * stored hash in order to confirm the integrity of the submitted data.
     */
    struct BeaconChainQueuedWithdrawal {
        // @notice Number of "beacon chain ETH" virtual shares in the withdrawal.
        uint256 shares;
        // @notice Owner of the EigenPod who initiated the withdrawal.
        address podOwner;
        // @notice Nonce of the podOwner when the withdrawal was queued. Used to help ensure uniqueness of the hash of the withdrawal.
        uint96 nonce;
        // @notice Block number at which the withdrawal was initiated.
        uint32 withdrawalStartBlock;
        // @notice The operator to which the podOwner was delegated in EigenLayer when the withdrawal was created.
        address delegatedAddress;
        // @notice The address that can complete the withdrawal and receive the withdrawn funds.
        address withdrawer;
    }

    /**
     * @notice Struct used to track a pod owner's "undelegation limbo" status and associated variables.
     * @dev Undelegation limbo is a mode which a staker can enter into, in which they remove their virtual "beacon chain ETH shares" from EigenLayer's delegation
     * system but do not necessarily withdraw the associated ETH from EigenLayer itself. This mode allows users who have restaked native ETH a route via
     * which they can undelegate from an operator without needing to exit any of their validators from the Consensus Layer.
     */
    struct UndelegationLimboStatus {
        // @notice Whether or not the pod owner is in the "undelegation limbo" mode.
        bool active;
        // @notice The block at which the pod owner entered "undelegation limbo". Should be zero if `podOwnerIsInUndelegationLimbo` is marked as 'false'
        uint32 startBlock;
        // @notice The address which the pod owner was delegated to at the time that they entered "undelegation limbo".
        address delegatedAddress;
    }

    /// @notice Emitted to notify the update of the beaconChainOracle address
    event BeaconOracleUpdated(address indexed newOracleAddress);

    /// @notice Emitted to notify the deployment of an EigenPod
    event PodDeployed(address indexed eigenPod, address indexed podOwner);

    /// @notice Emitted to notify a deposit of beacon chain ETH recorded in the strategy manager
    event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);

    /// @notice Emitted when `maxPods` value is updated from `previousValue` to `newValue`
    event MaxPodsUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when a withdrawal of beacon chain ETH is queued
    event BeaconChainETHWithdrawalQueued(
        address indexed podOwner,
        uint256 shares,
        uint96 nonce,
        address delegatedAddress,
        address withdrawer,
        bytes32 withdrawalRoot
    );

    /// @notice Emitted when a withdrawal of beacon chain ETH is completed
    event BeaconChainETHWithdrawalCompleted(
        address indexed podOwner,
        uint256 shares,
        uint96 nonce,
        address delegatedAddress,
        address withdrawer,
        bytes32 withdrawalRoot
    );

    // @notice Emitted when `podOwner` enters the "undelegation limbo" mode
    event UndelegationLimboEntered(address indexed podOwner);

    // @notice Emitted when `podOwner` exits the "undelegation limbo" mode
    event UndelegationLimboExited(address indexed podOwner);

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external;

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod.
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.
     * @param podOwner The owner of the pod whose balance must be deposited.
     * @param amount The amount of ETH to 'deposit' (i.e. be credited to the podOwner).
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function restakeBeaconChainETH(address podOwner, uint256 amount) external;

    /**
     * @notice Records an update in beacon chain strategy shares in the strategy manager
     * @param podOwner is the pod owner whose shares are to be updated,
     * @param sharesDelta is the change in podOwner's beaconChainETHStrategy shares
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) external;

    /**
     * @notice Called by a podOwner to queue a withdrawal of some (or all) of their virtual beacon chain ETH shares.
     * @param amountWei The amount of ETH to withdraw.
     * @param withdrawer The address that can complete the withdrawal and receive the withdrawn funds.
     */
    function queueWithdrawal(uint256 amountWei, address withdrawer) external returns (bytes32);

    /**
     * @notice Completes an existing BeaconChainQueuedWithdrawal by sending the ETH to the 'withdrawer'
     * @param queuedWithdrawal is the queued withdrawal to be completed
     * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     */
    function completeQueuedWithdrawal(
        BeaconChainQueuedWithdrawal memory queuedWithdrawal,
        uint256 middlewareTimesIndex
    ) external;

    /**
     * @notice forces the podOwner into the "undelegation limbo" mode, and returns the number of virtual 'beacon chain ETH shares'
     * that the podOwner has, which were entered into undelegation limbo.
     * @param podOwner is the staker to be forced into undelegation limbo
     * @param delegatedTo is the operator the staker is currently delegated to
     * @dev This function can only be called by the DelegationManager contract
     */
    function forceIntoUndelegationLimbo(
        address podOwner,
        address delegatedTo
    ) external returns (uint256 sharesRemovedFromDelegation);

    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of this contract (i.e. governance)
     */
    function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external;

    /// @notice Returns the address of the `podOwner`'s EigenPod if it has been deployed.
    function ownerToPod(address podOwner) external view returns (IEigenPod);

    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) external view returns (IEigenPod);

    /// @notice The ETH2 Deposit Contract
    function ethPOS() external view returns (IETHPOSDeposit);

    /// @notice Beacon proxy to which the EigenPods point
    function eigenPodBeacon() external view returns (IBeacon);

    /// @notice Oracle contract that provides updates to the beacon chain's state
    function beaconChainOracle() external view returns (IBeaconChainOracle);

    /// @notice Returns the beacon block root at `timestamp`. Reverts if the Beacon block root at `timestamp` has not yet been finalized.
    function getBlockRootAtTimestamp(uint64 timestamp) external view returns (bytes32);

    /// @notice EigenLayer's StrategyManager contract
    function strategyManager() external view returns (IStrategyManager);

    /// @notice EigenLayer's Slasher contract
    function slasher() external view returns (ISlasher);

    /// @notice EigenLayer's DelegationManager contract
    function delegationManager() external view returns (IDelegationManager);

    /// @notice Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.
    function hasPod(address podOwner) external view returns (bool);

    /// @notice returns shares of provided podOwner
    function podOwnerShares(address podOwner) external returns (uint256);

    /// @notice returns canonical, virtual beaconChainETH strategy
    function beaconChainETHStrategy() external view returns (IStrategy);

    /// @notice The number of EigenPods that have been deployed
    function numPods() external view returns (uint256);

    /// @notice The maximum number of EigenPods that can be deployed
    function maxPods() external view returns (uint256);

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.
    function calculateWithdrawalRoot(
        BeaconChainQueuedWithdrawal memory queuedWithdrawal
    ) external pure returns (bytes32);

    /**
     * @notice Returns 'false' if `staker` has removed all of their beacon chain ETH "shares" from delegation, either by queuing a
     * withdrawal for them OR by going into "undelegation limbo", and 'true' otherwise
     */
    function podOwnerHasActiveShares(address staker) external view returns (bool);

    // @notice Getter function for the internal `_podOwnerUndelegationLimboStatus` mapping.
    function podOwnerUndelegationLimboStatus(address podOwner) external view returns (UndelegationLimboStatus memory);

    // @notice Getter function for `_podOwnerUndelegationLimboStatus.undelegationLimboActive`.
    function isInUndelegationLimbo(address podOwner) external view returns (bool);
}
