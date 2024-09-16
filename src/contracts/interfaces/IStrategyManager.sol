// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategy.sol";
import "./IShareManager.sol";
import "./IDelegationManager.sol";
import "./IEigenPodManager.sol";

interface IStrategyManagerErrors {
    /// @dev Thrown when total strategies deployed exceeds max.
    error MaxStrategiesExceeded();
    /// @dev Thrown when call attempted from address that's not delegation manager.
    error OnlyDelegationManager();
    /// @dev Thrown when call attempted from address that's not strategy whitelister.
    error OnlyStrategyWhitelister();
    /// @dev Thrown when provided `shares` amount is too high.
    error SharesAmountTooHigh();
    /// @dev Thrown when provided `shares` amount is zero.
    error SharesAmountZero();
    /// @dev Thrown when provided `staker` address is null.
    error StakerAddressZero();
    /// @dev Thrown when provided `strategy` not found.
    error StrategyNotFound();
    /// @dev Thrown when attempting to deposit to a non-whitelisted strategy.
    error StrategyNotWhitelisted();
}

interface IStrategyManagerEvents {
    /**
     * @notice Emitted when a new deposit occurs on behalf of `staker`.
     * @param staker Is the staker who is depositing funds into EigenLayer.
     * @param strategy Is the strategy that `staker` has deposited into.
     * @param token Is the token that `staker` deposited.
     * @param shares Is the number of new shares `staker` has been granted in `strategy`.
     */
    event Deposit(address staker, IERC20 token, IStrategy strategy, uint256 shares);

    /// @notice Emitted when the `strategyWhitelister` is changed
    event StrategyWhitelisterChanged(address previousAddress, address newAddress);

    /// @notice Emitted when a strategy is added to the approved list of strategies for deposit
    event StrategyAddedToDepositWhitelist(IStrategy strategy);

    /// @notice Emitted when a strategy is removed from the approved list of strategies for deposit
    event StrategyRemovedFromDepositWhitelist(IStrategy strategy);
}

/**
 * @title Interface for the primary entrypoint for funds into EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `StrategyManager` contract itself for implementation details.
 */
interface IStrategyManager is IStrategyManagerErrors, IStrategyManagerEvents, IShareManager {
    /**
     * @notice Initializes the strategy manager contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
     * and transfers contract ownership to the specified `initialOwner`.
     * @param initialOwner Ownership of this contract is transferred to this address.
     * @param initialStrategyWhitelister The initial value of `strategyWhitelister` to set.
     * @param initialPausedStatus The initial value of `_paused` to set.
     */
    function initialize(
        address initialOwner,
        address initialStrategyWhitelister,
        uint256 initialPausedStatus
    ) external;

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
    function depositIntoStrategy(IStrategy strategy, IERC20 token, uint256 amount) external returns (uint256 shares);

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
     *
     *  WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a strategy is not recommended.  This can lead to attack vectors
     *          where the token balance and corresponding strategy shares are not in sync upon reentrancy
     */
    function depositIntoStrategyWithSignature(
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        address staker,
        uint256 expiry,
        bytes memory signature
    ) external returns (uint256 shares);

    /**
     * @notice Burns Strategy shares for the given strategy by calling into the strategy to transfer to the default burn address.
     * @param strategy The strategy to burn shares in.
     * @param sharesToBurn The amount of shares to burn.
     * @dev This function is only called by the DelegationManager when an operator is slashed.
     */
    function burnShares(IStrategy strategy, uint256 sharesToBurn) external;

    /**
     * @notice Owner-only function to change the `strategyWhitelister` address.
     * @param newStrategyWhitelister new address for the `strategyWhitelister`.
     */
    function setStrategyWhitelister(
        address newStrategyWhitelister
    ) external;

    /**
     * @notice Owner-only function that adds the provided Strategies to the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToWhitelist Strategies that will be added to the `strategyIsWhitelistedForDeposit` mapping (if they aren't in it already)
     */
    function addStrategiesToDepositWhitelist(
        IStrategy[] calldata strategiesToWhitelist
    ) external;

    /**
     * @notice Owner-only function that removes the provided Strategies from the 'whitelist' of strategies that stakers can deposit into
     * @param strategiesToRemoveFromWhitelist Strategies that will be removed to the `strategyIsWhitelistedForDeposit` mapping (if they are in it)
     */
    function removeStrategiesFromDepositWhitelist(
        IStrategy[] calldata strategiesToRemoveFromWhitelist
    ) external;

    /// @notice Returns bool for whether or not `strategy` is whitelisted for deposit
    function strategyIsWhitelistedForDeposit(
        IStrategy strategy
    ) external view returns (bool);

    /**
     * @notice Get all details on the staker's deposits and corresponding shares
     * @return (staker's strategies, shares in these strategies)
     */
    function getDeposits(
        address staker
    ) external view returns (IStrategy[] memory, uint256[] memory);

    function getStakerStrategyList(
        address staker
    ) external view returns (IStrategy[] memory);

    /// @notice Simple getter function that returns `stakerStrategyList[staker].length`.
    function stakerStrategyListLength(
        address staker
    ) external view returns (uint256);

    /// @notice Returns the current shares of `user` in `strategy`
    function stakerDepositShares(address user, IStrategy strategy) external view returns (uint256 shares);

    /// @notice Returns the single, central Delegation contract of EigenLayer
    function delegation() external view returns (IDelegationManager);

    /// @notice Returns the address of the `strategyWhitelister`
    function strategyWhitelister() external view returns (address);

    /**
     * @param staker The address of the staker.
     * @param strategy The strategy to deposit into.
     * @param token The token to deposit.
     * @param amount The amount of `token` to deposit.
     * @param nonce The nonce of the staker.
     * @param expiry The expiry of the signature.
     * @return The EIP-712 signable digest hash.
     */
    function calculateStrategyDepositDigestHash(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        uint256 nonce,
        uint256 expiry
    ) external view returns (bytes32);
}
