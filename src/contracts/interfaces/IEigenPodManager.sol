// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategyManager.sol";
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
        uint256 shares;
        address podOwner;
        uint96 nonce;
        uint32 withdrawalStartBlock;
        address delegatedAddress;
        bool alsoWithdraw;
    }

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
     * @notice queues a withdrawal beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod.
     * @param amountWei is the amount of ETH to withdraw
     * @param undelegateIfPossible is whether or not to undelegate the shares if possible
     * @param alsoWithdraw is whether or not to also withdraw the ETH from the beacon chain vs. redelegating to a new operator in EL
     */
    function queueWithdrawal(uint256 amountWei, bool undelegateIfPossible, bool alsoWithdraw) external returns(bytes32);

    /**
     * @notice forces a withdrawal of the podOwner's beaconChainETHStrategy shares
     * @param podOwner is the pod owner whose shares are to be removed
     */
    function forceWithdrawal(address podOwner) external returns (bytes32);


    /** 
     * @notice slashes a pending queued withdrawal of the podOwner's beaconChainETHStrategy shares
     * @param slashedFundsRecipient is the address to receive the slashed funds
     * @param queuedWithdrawal is the queued withdrawal to be slashed
     */
    function slashQueuedWithdrawal(address slashedFundsRecipient, BeaconChainQueuedWithdrawal memory queuedWithdrawal) external;

    /**
     * @notice slashes shares of the podOwner and sends them to the slashedFundsRecipient
     * @param slashedPodOwner is the address of the pod owner whose shares are to be slashed
     * @param slashedFundsRecipient is the address to receive the slashed funds
     * @param shareAmount is the amount of shares to be slashed     
     */
    function slashShares(address slashedPodOwner, address slashedFundsRecipient, uint256 shareAmount) external;


    /**
     * @notice Completes an existing queuedWithdrawal either by sending the ETH to the recipient or allowing the podOwner to re-delegate it
     * @param queuedWithdrawal is the queued withdrawal to be completed
     * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function completeWithdrawal(BeaconChainQueuedWithdrawal memory queuedWithdrawal, uint256 middlewareTimesIndex) external;
    
    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of this contract (i.e. governance)
     */
    function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external;

    /// @notice Returns the address of the `podOwner`'s EigenPod if it has been deployed.
    function ownerToPod(address podOwner) external view returns(IEigenPod);

    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) external view returns(IEigenPod);

    /// @notice Oracle contract that provides updates to the beacon chain's state
    function beaconChainOracle() external view returns(IBeaconChainOracle);    

    /// @notice Returns the Beacon Chain state root at `blockNumber`. Reverts if the Beacon Chain state root at `blockNumber` has not yet been finalized.
    function getBeaconChainStateRoot(uint64 blockNumber) external view returns(bytes32);

    /// @notice EigenLayer's StrategyManager contract
    function strategyManager() external view returns(IStrategyManager);

    /// @notice EigenLayer's Slasher contract
    function slasher() external view returns(ISlasher);

    function hasPod(address podOwner) external view returns (bool);

    /// @notice returns shares of provided podOwner
    function podOwnerShares(address podOwner) external returns (uint256);

    /// @notice returns canonical beaconChainETH strategy
    function beaconChainETHStrategy() external view returns (IStrategy);
}