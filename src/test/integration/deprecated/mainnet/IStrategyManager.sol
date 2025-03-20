// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";

/**
 * @notice M2 DEPRECATED INTERFACE at commit hash https://github.com/Layr-Labs/eigenlayer-contracts/tree/426f461c59b4f0e16f8becdffd747075edcaded8
 * @title Interface for the primary entrypoint for funds into EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `StrategyManager` contract itself for implementation details.
 */
interface IStrategyManager_DeprecatedM2 {
    /**
     * @notice Deposits `amount` of `token` into the specified `strategy`, with the resultant shares credited to `msg.sender`
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the staker
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
     *
     * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy.
     */
    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint amount) external returns (uint shares);

    /**
     * @notice Used for depositing an asset into the specified strategy with the resultant shares credited to `staker`,
     * who must sign off on the action.
     * Note that the assets are transferred out/from the `msg.sender`, not from the `staker`; this function is explicitly designed
     * purely to help one address deposit 'for' another.
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the staker
     * @param staker the staker that the deposited assets will be credited to
     * @param expiry the timestamp at which the signature expires
     * @param signature is a valid signature from the `staker`. either an ECDSA signature if the `staker` is an EOA, or data to forward
     * following EIP-1271 if the `staker` is a contract
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
     * targeting stakers who may be attempting to undelegate.
     * @dev Cannot be called if thirdPartyTransfersForbidden is set to true for this strategy
     *
     *  WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy
     */
    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint amount,
        address staker,
        uint expiry,
        bytes memory signature
    ) external returns (uint shares);

    /// @notice Used by the DelegationManager to remove a Staker's shares from a particular strategy when entering the withdrawal queue
    function removeShares(address staker, IStrategy strategy, uint shares) external;

    /// @notice Used by the DelegationManager to award a Staker some shares that have passed through the withdrawal queue
    function addShares(address staker, IERC20 token, IStrategy strategy, uint shares) external;

    /// @notice Used by the DelegationManager to convert withdrawn shares to tokens and send them to a recipient
    function withdrawSharesAsTokens(address recipient, IStrategy strategy, uint shares, IERC20 token) external;

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerStrategyShares(address user, IStrategy strategy) external view returns (uint shares);

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @param staker The staker of interest, whose deposits this function will fetch
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(address staker) external view returns (IStrategy[] memory, uint[] memory);

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address staker) external view returns (uint);

    /**
     * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
     * @param thirdPartyTransfersForbiddenValues bool values to set `thirdPartyTransfersForbidden` to for each strategy
     */
    function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist, bool[] calldata thirdPartyTransfersForbiddenValues)
        external;

    /**
     * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
     */
    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external;

    /**
     * If true for a strategy, a user cannot depositIntoStrategyWithSignature into that strategy for another staker
     * and also when performing DelegationManager.queueWithdrawals, a staker can only withdraw to themselves.
     * Defaulted to false for all existing strategies.
     * @param strategy The strategy to set `thirdPartyTransfersForbidden` value to
     * @param value bool value to set `thirdPartyTransfersForbidden` to
     */
    function setThirdPartyTransfersForbidden(IStrategy strategy, bool value) external;

    /// @notice Returns the single, central Delegation contract of EigenLayer
    function delegation() external view returns (IDelegationManager);

    /// @notice Returns the EigenPodManager contract of EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager);

    /// @notice Returns the address of the `strategyWhitelister`
    function strategyWhitelister() external view returns (address);

    /// @notice Returns bool for whether or not `strategy` is whitelisted for deposit
    function strategyIsWhitelistedForDeposit(IStrategy strategy) external view returns (bool);

    /**
     * @notice Owner-only function to change the `strategyWhitelister` address.
     * @param newStrategyWhitelister new address for the `strategyWhitelister`.
     */
    function setStrategyWhitelister(address newStrategyWhitelister) external;

    /**
     * @notice Returns bool for whether or not `strategy` enables credit transfers. i.e enabling
     * depositIntoStrategyWithSignature calls or queueing withdrawals to a different address than the staker.
     */
    function thirdPartyTransfersForbidden(IStrategy strategy) external view returns (bool);

    /**
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     */
    function domainSeparator() external view returns (bytes32);

    /// VIEW FUNCTIONS FOR PUBLIC VARIABLES, NOT IN ORIGINAL INTERFACE
    function nonces(address user) external view returns (uint);

    function DEPOSIT_TYPEHASH() external view returns (bytes32);
}

/**
 * @notice M1 DEPRECATED INTERFACE at commit hash https://github.com/Layr-Labs/eigenlayer-contracts/tree/0139d6213927c0a7812578899ddd3dda58051928
 * @title Interface for the primary entrypoint for funds into EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `StrategyManager` contract itself for implementation details.
 */
interface IStrategyManager_DeprecatedM1 {
    // packed struct for queued withdrawals; helps deal with stack-too-deep errors
    struct WithdrawerAndNonce {
        address withdrawer;
        uint96 nonce;
    }

    /**
     * Struct type used to specify an existing queued withdrawal. Rather than storing the entire struct, only a hash is stored.
     * In functions that operate on existing queued withdrawals -- e.g. `startQueuedWithdrawalWaitingPeriod` or `completeQueuedWithdrawal`,
     * the data is resubmitted and the hash of the submitted data is computed by `calculateWithdrawalRoot` and checked against the
     * stored hash in order to confirm the integrity of the submitted data.
     */
    struct QueuedWithdrawal {
        IStrategy[] strategies;
        uint[] shares;
        address depositor;
        WithdrawerAndNonce withdrawerAndNonce;
        uint32 withdrawalStartBlock;
        address delegatedAddress;
    }

    /**
     * @notice Deposits `amount` of `token` into the specified `strategy`, with the resultant shares credited to `msg.sender`
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the depositor
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
     *
     * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy.
     */
    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint amount) external returns (uint shares);

    /**
     * @notice Deposits `amount` of beaconchain ETH into this contract on behalf of `staker`
     * @param staker is the entity that is restaking in eigenlayer,
     * @param amount is the amount of beaconchain ETH being restaked,
     * @dev Only callable by EigenPodManager.
     */
    function depositBeaconChainETH(address staker, uint amount) external;

    /**
     * @notice Records an overcommitment event on behalf of a staker. The staker's beaconChainETH shares are decremented by `amount`.
     * @param overcommittedPodOwner is the pod owner to be slashed
     * @param beaconChainETHStrategyIndex is the index of the beaconChainETHStrategy in case it must be removed,
     * @param amount is the amount to decrement the slashedAddress's beaconChainETHStrategy shares
     * @dev Only callable by EigenPodManager.
     */
    function recordOvercommittedBeaconChainETH(address overcommittedPodOwner, uint beaconChainETHStrategyIndex, uint amount) external;

    /**
     * @notice Used for depositing an asset into the specified strategy with the resultant shares credited to `staker`,
     * who must sign off on the action.
     * Note that the assets are transferred out/from the `msg.sender`, not from the `staker`; this function is explicitly designed
     * purely to help one address deposit 'for' another.
     * @param strategy is the specified strategy where deposit is to be made,
     * @param token is the denomination in which the deposit is to be made,
     * @param amount is the amount of token to be deposited in the strategy by the depositor
     * @param staker the staker that the deposited assets will be credited to
     * @param expiry the timestamp at which the signature expires
     * @param signature is a valid signature from the `staker`. either an ECDSA signature if the `staker` is an EOA, or data to forward
     * following EIP-1271 if the `staker` is a contract
     * @return shares The amount of new shares in the `strategy` created as part of the action.
     * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
     * @dev A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
     * targeting stakers who may be attempting to undelegate.
     * @dev Cannot be called on behalf of a staker that is 'frozen' (this function will revert if the `staker` is frozen).
     *
     *  WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy
     */
    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint amount,
        address staker,
        uint expiry,
        bytes memory signature
    ) external returns (uint shares);

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerDepositShares(address user, IStrategy strategy) external view returns (uint shares);

    /**
     * @notice Get all details on the depositor's deposits and corresponding shares
     * @return (depositor's strategies, shares in these strategies)
     */
    function getDeposits(address depositor) external view returns (IStrategy[] memory, uint[] memory);

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(address staker) external view returns (uint);

    /**
     * @notice Called by a staker to queue a withdrawal of the given amount of `shares` from each of the respective given `strategies`.
     * @dev Stakers will complete their withdrawal by calling the 'completeQueuedWithdrawal' function.
     * User shares are decreased in this function, but the total number of shares in each strategy remains the same.
     * The total number of shares is decremented in the 'completeQueuedWithdrawal' function instead, which is where
     * the funds are actually sent to the user through use of the strategies' 'withdrawal' function. This ensures
     * that the value per share reported by each strategy will remain consistent, and that the shares will continue
     * to accrue gains during the enforced withdrawal waiting period.
     * @param strategyIndexes is a list of the indices in `stakerStrategyList[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @param strategies The Strategies to withdraw from
     * @param shares The amount of shares to withdraw from each of the respective Strategies in the `strategies` array
     * @param withdrawer The address that can complete the withdrawal and will receive any withdrawn funds or shares upon completing the withdrawal
     * @param undelegateIfPossible If this param is marked as 'true' *and the withdrawal will result in `msg.sender` having no shares in any Strategy,*
     * then this function will also make an internal call to `undelegate(msg.sender)` to undelegate the `msg.sender`.
     * @return The 'withdrawalRoot' of the newly created Queued Withdrawal
     * @dev Strategies are removed from `stakerStrategyList` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `stakerStrategyList`. The simplest way to calculate the correct `strategyIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `stakerStrategyList` to lowest index
     * @dev Note that if the withdrawal includes shares in the enshrined 'beaconChainETH' strategy, then it must *only* include shares in this strategy, and
     * `withdrawer` must match the caller's address. The first condition is because slashing of queued withdrawals cannot be guaranteed
     * for Beacon Chain ETH (since we cannot trigger a withdrawal from the beacon chain through a smart contract) and the second condition is because shares in
     * the enshrined 'beaconChainETH' strategy technically represent non-fungible positions (deposits to the Beacon Chain, each pointed at a specific EigenPod).
     */
    function queueWithdrawal(
        uint[] calldata strategyIndexes,
        IStrategy[] calldata strategies,
        uint[] calldata shares,
        address withdrawer,
        bool undelegateIfPossible
    ) external returns (bytes32);

    /**
     * @notice Used to complete the specified `queuedWithdrawal`. The function caller must match `queuedWithdrawal.withdrawer`
     * @param queuedWithdrawal The QueuedWithdrawal to complete.
     * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `strategies` array
     * of the `queuedWithdrawal`. This input can be provided with zero length if `receiveAsTokens` is set to 'false' (since in that case, this input will be unused)
     * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     * @param receiveAsTokens If true, the shares specified in the queued withdrawal will be withdrawn from the specified strategies themselves
     * and sent to the caller, through calls to `queuedWithdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies
     * will simply be transferred to the caller directly.
     * @dev middlewareTimesIndex should be calculated off chain before calling this function by finding the first index that satisfies `slasher.canWithdraw`
     */
    function completeQueuedWithdrawal(
        QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint middlewareTimesIndex,
        bool receiveAsTokens
    ) external;

    /**
     * @notice Used to complete the specified `queuedWithdrawals`. The function caller must match `queuedWithdrawals[...].withdrawer`
     * @param queuedWithdrawals The QueuedWithdrawals to complete.
     * @param tokens Array of tokens for each QueuedWithdrawal. See `completeQueuedWithdrawal` for the usage of a single array.
     * @param middlewareTimesIndexes One index to reference per QueuedWithdrawal. See `completeQueuedWithdrawal` for the usage of a single index.
     * @param receiveAsTokens If true, the shares specified in the queued withdrawal will be withdrawn from the specified strategies themselves
     * and sent to the caller, through calls to `queuedWithdrawal.strategies[i].withdraw`. If false, then the shares in the specified strategies
     * will simply be transferred to the caller directly.
     * @dev Array-ified version of `completeQueuedWithdrawal`
     * @dev middlewareTimesIndex should be calculated off chain before calling this function by finding the first index that satisfies `slasher.canWithdraw`
     */
    function completeQueuedWithdrawals(
        QueuedWithdrawal[] calldata queuedWithdrawals,
        IERC20[][] calldata tokens,
        uint[] calldata middlewareTimesIndexes,
        bool[] calldata receiveAsTokens
    ) external;

    /**
     * @notice Slashes the shares of a 'frozen' operator (or a staker delegated to one)
     * @param slashedAddress is the frozen address that is having its shares slashed
     * @param recipient is the address that will receive the slashed funds, which could e.g. be a harmed party themself,
     * or a MerkleDistributor-type contract that further sub-divides the slashed funds.
     * @param strategies Strategies to slash
     * @param shareAmounts The amount of shares to slash in each of the provided `strategies`
     * @param tokens The tokens to use as input to the `withdraw` function of each of the provided `strategies`
     * @param strategyIndexes is a list of the indices in `stakerStrategyList[msg.sender]` that correspond to the strategies
     * for which `msg.sender` is withdrawing 100% of their shares
     * @param recipient The slashed funds are withdrawn as tokens to this address.
     * @dev strategies are removed from `stakerStrategyList` by swapping the last entry with the entry to be removed, then
     * popping off the last entry in `stakerStrategyList`. The simplest way to calculate the correct `strategyIndexes` to input
     * is to order the strategies *for which `msg.sender` is withdrawing 100% of their shares* from highest index in
     * `stakerStrategyList` to lowest index
     */
    function slashShares(
        address slashedAddress,
        address recipient,
        IStrategy[] calldata strategies,
        IERC20[] calldata tokens,
        uint[] calldata strategyIndexes,
        uint[] calldata shareAmounts
    ) external;

    /**
     * @notice Slashes an existing queued withdrawal that was created by a 'frozen' operator (or a staker delegated to one)
     * @param recipient The funds in the slashed withdrawal are withdrawn as tokens to this address.
     * @param queuedWithdrawal The previously queued withdrawal to be slashed
     * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Strategy in the `strategies`
     * array of the `queuedWithdrawal`.
     * @param indicesToSkip Optional input parameter -- indices in the `strategies` array to skip (i.e. not call the 'withdraw' function on). This input exists
     * so that, e.g., if the slashed QueuedWithdrawal contains a malicious strategy in the `strategies` array which always reverts on calls to its 'withdraw' function,
     * then the malicious strategy can be skipped (with the shares in effect "burned"), while the non-malicious strategies are still called as normal.
     */
    function slashQueuedWithdrawal(
        address recipient,
        QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint[] calldata indicesToSkip
    ) external;

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.
    function calculateWithdrawalRoot(QueuedWithdrawal memory queuedWithdrawal) external pure returns (bytes32);

    /**
     * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
     */
    function addStrategiesToDepositWhitelist(IStrategy[] calldata strategiesToWhitelist) external;

    /**
     * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
     */
    function removeStrategiesFromDepositWhitelist(IStrategy[] calldata strategiesToRemoveFromWhitelist) external;

    /// @notice Returns the single, central Delegation contract of EigenLayer
    function delegation() external view returns (IDelegationManager);

    /// @notice returns the enshrined, virtual 'beaconChainETH' Strategy
    function beaconChainETHStrategy() external view returns (IStrategy);

    /// @notice Returns the number of blocks that must pass between the time a withdrawal is queued and the time it can be completed
    function withdrawalDelayBlocks() external view returns (uint);

    /// VIEW FUNCTIONS FOR PUBLIC VARIABLES, NOT IN ORIGINAL INTERFACE
    function withdrawalRootPending(bytes32 withdrawalRoot) external view returns (bool);

    function numWithdrawalsQueued(address staker) external view returns (uint);
}
