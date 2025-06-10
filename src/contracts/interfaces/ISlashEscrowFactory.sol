// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../interfaces/IStrategy.sol";
import "../libraries/OperatorSetLib.sol";
import "../interfaces/ISlashEscrow.sol";

interface ISlashEscrowFactoryErrors {
    /// @notice Thrown when a caller is not the strategy manager.
    error OnlyStrategyManager();

    /// @notice Thrown when a caller is not the redistribution recipient.
    error OnlyRedistributionRecipient();

    /// @notice Thrown when a escrow is not mature.
    error EscrowNotMature();

    /// @notice Thrown when the escrow delay has not elapsed.
    error EscrowDelayNotElapsed();
}

interface ISlashEscrowFactoryEvents {
    /// @notice Emitted when a escrow is initiated.
    event StartEscrow(OperatorSet operatorSet, uint256 slashId, IStrategy strategy, uint32 startBlock);

    /// @notice Emitted when a escrow is released.
    event EscrowComplete(OperatorSet operatorSet, uint256 slashId, IStrategy strategy, address recipient);

    /// @notice Emitted when a escrow is paused.
    event EscrowPaused(OperatorSet operatorSet, uint256 slashId);

    /// @notice Emitted when a escrow is unpaused.
    event EscrowUnpaused(OperatorSet operatorSet, uint256 slashId);

    /// @notice Emitted when a global escrow delay is set.
    event GlobalEscrowDelaySet(uint32 delay);

    /// @notice Emitted when a escrow delay is set.
    event StrategyEscrowDelaySet(IStrategy strategy, uint32 delay);
}

interface ISlashEscrowFactory is ISlashEscrowFactoryErrors, ISlashEscrowFactoryEvents {
    /**
     * @notice Initializes the initial owner and paused status.
     * @param initialOwner The initial owner of the router.
     * @param initialPausedStatus The initial paused status of the router.
     * @param initialGlobalDelayBlocks The initial global escrow delay.
     */
    function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) external;

    /**
     * @notice Initiates a slash escrow.
     * @param operatorSet The operator set whose escrow is being locked up.
     * @param slashId The slash ID of the escrow that is being locked up.
     * @param strategy The strategy that whose underlying tokens are being redistributed.
     * @dev This function can be called multiple times for a given `operatorSet` and `slashId`.
     */
    function initiateSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId, IStrategy strategy) external;

    /**
     * @notice Releases an escrow by transferring tokens from the `SlashEscrow` to the operator set's redistribution recipient.
     * @param operatorSet The operator set whose escrow is being released.
     * @param slashId The slash ID of the escrow that is being released.
     * @dev The caller must be the escrow recipient, unless the escrow recipient
     * is the default burn address in which case anyone can call.
     * @dev The slash escrow is released once the delay for ALL strategies has elapsed.
     */
    function releaseSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Releases an escrow for a single strategy in a slash.
     * @param operatorSet The operator set whose escrow is being released.
     * @param slashId The slash ID of the escrow that is being released.
     * @param strategy The strategy whose escrow is being released.
     * @dev The caller must be the redistribution recipient, unless the redistribution recipient
     * is the default burn address in which case anyone can call.
     * @dev The slash escrow is released once the delay for ALL strategies has elapsed.
     */
    function releaseSlashEscrowByStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external;

    /**
     * @notice Pauses a escrow.
     * @param operatorSet The operator set whose escrow is being paused.
     * @param slashId The slash ID of the escrow that is being paused.
     */
    function pauseEscrow(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Unpauses a escrow.
     * @param operatorSet The operator set whose escrow is being unpaused.
     * @param slashId The slash ID of the escrow that is being unpaused.
     */
    function unpauseEscrow(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Sets the delay for the escrow of a strategies underlying token.
     * @dev The largest of all strategy delays or global delay will be used.
     * @dev This delay setting only applies to new slashes and does not affect existing ones.
     * @param strategy The strategy whose escrow delay is being set.
     * @param delay The delay for the escrow.
     */
    function setStrategyEscrowDelay(IStrategy strategy, uint32 delay) external;

    /**
     * @notice Sets a global delay applicable to all strategies.
     * @dev This delay setting only applies to new slashes and does not affect existing ones.
     * @param delay The delay for the escrow.
     */
    function setGlobalEscrowDelay(
        uint32 delay
    ) external;

    /**
     * @notice Returns the operator sets that have pending escrows.
     * @return operatorSets The operator sets that have pending escrows.
     */
    function getPendingOperatorSets() external view returns (OperatorSet[] memory operatorSets);

    /**
     * @notice Returns the total number of operator sets with pending escrows.
     * @return The total number of operator sets with pending escrows.
     */
    function getTotalPendingOperatorSets() external view returns (uint256);

    /**
     * @notice Returns whether an operator set has pending escrows.
     * @param operatorSet The operator set whose pending escrows are being queried.
     * @return Whether the operator set has pending escrows.
     */
    function isPendingOperatorSet(
        OperatorSet calldata operatorSet
    ) external view returns (bool);

    /**
     * @notice Returns the pending slash IDs for an operator set.
     * @param operatorSet The operator set whose pending slash IDs are being queried.
     */
    function getPendingSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (uint256[] memory);

    /**
     * @notice Returns the pending escrows and their release blocks.
     * @return operatorSets The pending operator sets.
     * @return isRedistributing Whether the operator set is redistributing.
     * @return slashIds The pending slash IDs for each operator set. Indexed by operator set.
     * @return completeBlocks The block at which a slashID can be released. Indexed by [operatorSet][slashId]
     */
    function getPendingEscrows()
        external
        view
        returns (
            OperatorSet[] memory operatorSets,
            bool[] memory isRedistributing,
            uint256[][] memory slashIds,
            uint32[][] memory completeBlocks
        );

    /**
     * @notice Returns the total number of slash IDs for an operator set.
     * @param operatorSet The operator set whose total slash IDs are being queried.
     * @return The total number of slash IDs for the operator set.
     */
    function getTotalPendingSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (uint256);

    /**
     * @notice Returns whether a slash ID is pending for an operator set.
     * @param operatorSet The operator set whose pending slash IDs are being queried.
     * @param slashId The slash ID of the slash that is being queried.
     * @return Whether the slash ID is pending for the operator set.
     */
    function isPendingSlashId(OperatorSet calldata operatorSet, uint256 slashId) external view returns (bool);

    /**
     * @notice Returns the pending strategies for a slash ID for an operator set.
     * @dev This is a variant that returns the pending strategies for a slash ID for an operator set.
     * @param operatorSet The operator set whose pending strategies are being queried.
     * @param slashId The slash ID of the strategies that are being queried.
     * @return strategies The strategies that are pending strategies.
     */
    function getPendingStrategiesForSlashId(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (IStrategy[] memory strategies);

    /**
     * @notice Returns all pending strategies for all slash IDs for an operator set.
     * @dev This is a variant that returns all pending strategies for all slash IDs for an operator set.
     * @param operatorSet The operator set whose pending strategies are being queried.
     * @return strategies The strategies that are pending strategies.
     */
    function getPendingStrategiesForSlashIds(
        OperatorSet calldata operatorSet
    ) external view returns (IStrategy[][] memory strategies);

    /**
     * @notice Returns the number of pending strategies for a slash ID for an operator set.
     * @param operatorSet The operator set whose pending strategies are being queried.
     * @param slashId The slash ID of the strategies that are being queried.
     * @return The number of pending strategies.
     */
    function getTotalPendingStrategiesForSlashId(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint256);

    /**
     * @notice Returns the pending underlying amount for a strategy for an operator set and slash ID.
     * @param operatorSet The operator set whose pending underlying amount is being queried.
     * @param slashId The slash ID of the escrow that is being queried.
     * @param strategy The strategy whose pending underlying amount is being queried.
     * @return The pending underlying amount.
     */
    function getPendingUnderlyingAmountForStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external view returns (uint256);

    /**
     * @notice Returns the paused status of a escrow.
     * @param operatorSet The operator set whose escrow is being queried.
     * @param slashId The slash ID of the escrow that is being queried.
     * @return The paused status of the escrow.
     */
    function isEscrowPaused(OperatorSet calldata operatorSet, uint256 slashId) external view returns (bool);

    /**
     * @notice Returns the block at which the escrow can be released.
     * @param operatorSet The operator set whose start block is being queried.
     * @param slashId The slash ID of the start block that is being queried.
     * @return The block at which the escrow can be released.
     */
    function getEscrowCompleteBlock(OperatorSet calldata operatorSet, uint256 slashId) external view returns (uint32);

    /**
     * @notice Returns the escrow delay for a strategy.
     * @param strategy The strategy whose escrow delay is being queried.
     * @return The escrow delay.
     */
    function getStrategyEscrowDelay(
        IStrategy strategy
    ) external view returns (uint32);

    /**
     * @notice Returns the global escrow delay.
     * @return The global escrow delay.
     */
    function getGlobalEscrowDelay() external view returns (uint32);

    /**
     * @notice Returns the salt for a slash escrow.
     * @param operatorSet The operator set whose slash escrow is being queried.
     * @param slashId The slash ID of the slash escrow that is being queried.
     * @return The salt for the slash escrow.
     */
    function computeSlashEscrowSalt(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external pure returns (bytes32);

    /**
     * @notice Returns whether a slash escrow is deployed or not.
     * @param operatorSet The operator set whose slash escrow is being queried.
     * @param slashId The slash ID of the slash escrow that is being queried.
     * @return Whether the slash escrow is deployed.
     */
    function isDeployedSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) external view returns (bool);

    /**
     * @notice Returns whether a slash escrow is deployed.
     * @param slashEscrow The slash escrow that is being queried.
     * @return Whether the slash escrow is deployed.
     */
    function isDeployedSlashEscrow(
        ISlashEscrow slashEscrow
    ) external view returns (bool);

    /**
     * @notice Returns the slash escrow for an operator set and slash ID.
     * @param operatorSet The operator set whose slash escrow is being queried.
     * @param slashId The slash ID of the slash escrow that is being queried.
     * @return The slash escrow.
     */
    function getSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) external view returns (ISlashEscrow);
}
