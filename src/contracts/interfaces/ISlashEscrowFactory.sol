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

    /// @notice Thrown when a redistribution is not mature.
    error RedistributionNotMature();

    /// @notice Thrown when a burn or redistribution delay is less than the minimum burn or redistribution delay.
    error BurnOrRedistributionDelayLessThanMinimum();

    /// @notice Thrown when the escrow delay has not elapsed.
    error EscrowDelayNotElapsed();
}

interface ISlashEscrowFactoryEvents {
    /// @notice Emitted when a redistribution is initiated.
    event StartBurnOrRedistribution(OperatorSet operatorSet, uint256 slashId, IStrategy strategy, uint32 startBlock);

    /// @notice Emitted when a redistribution is released.
    event BurnOrRedistributionComplete(OperatorSet operatorSet, uint256 slashId, IStrategy strategy, address recipient);

    /// @notice Emitted when a redistribution is paused.
    event RedistributionPaused(OperatorSet operatorSet, uint256 slashId);

    /// @notice Emitted when a redistribution is unpaused.
    event RedistributionUnpaused(OperatorSet operatorSet, uint256 slashId);

    /// @notice Emitted when a global burn or redistribution delay is set.
    event GlobalBurnOrRedistributionDelaySet(uint32 delay);

    /// @notice Emitted when a burn or redistribution delay is set.
    event StrategyBurnOrRedistributionDelaySet(IStrategy strategy, uint32 delay);
}

interface ISlashEscrowFactory is ISlashEscrowFactoryErrors, ISlashEscrowFactoryEvents {
    /**
     * @notice Initializes the initial owner and paused status.
     * @param initialOwner The initial owner of the router.
     * @param initialPausedStatus The initial paused status of the router.
     * @param initialGlobalDelayBlocks The initial global burn or redistribution delay.
     */
    function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) external;

    /**
     * @notice Locks up a redistribution.
     * @param operatorSet The operator set whose redistribution is being locked up.
     * @param slashId The slash ID of the redistribution that is being locked up.
     * @param strategy The strategy that whose underlying tokens are being redistributed.
     */
    function initiateSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId, IStrategy strategy) external;

    /**
     * @notice Releases a redistribution.
     * @param operatorSet The operator set whose redistribution is being released.
     * @param slashId The slash ID of the redistribution that is being released.
     * @dev The caller must be the redistribution recipient, unless the redistribution recipient
     * is the default burn address in which case anyone can call.
     * @dev The slash escrow is released once the delay for ALL strategies has elapsed.
     */
    function releaseSlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Deploys a counterfactual `SlashEscrow` if code hasn't already been deployed.
     * @param operatorSet The operator set whose slash escrow is being deployed.
     * @param slashId The slash ID of the slash escrow that is being deployed.
     * @return The deployed `SlashEscrow`.
     */
    function deploySlashEscrow(OperatorSet calldata operatorSet, uint256 slashId) external returns (ISlashEscrow);

    /**
     * @notice Pauses a redistribution.
     * @param operatorSet The operator set whose redistribution is being paused.
     * @param slashId The slash ID of the redistribution that is being paused.
     */
    function pauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Unpauses a redistribution.
     * @param operatorSet The operator set whose redistribution is being unpaused.
     * @param slashId The slash ID of the redistribution that is being unpaused.
     */
    function unpauseRedistribution(OperatorSet calldata operatorSet, uint256 slashId) external;

    /**
     * @notice Sets the delay for the burn or redistribution of a strategies underlying token.
     * @dev If the strategy delay is less than the global delay, the strategy delay will be used.
     * @param strategy The strategy whose burn or redistribution delay is being set.
     * @param delay The delay for the burn or redistribution.
     */
    function setStrategyBurnOrRedistributionDelay(IStrategy strategy, uint32 delay) external;

    /**
     * @notice Sets the delay for the burn or redistribution of all strategies underlying tokens globally.
     * @param delay The delay for the burn or redistribution.
     */
    function setGlobalBurnOrRedistributionDelay(
        uint32 delay
    ) external;

    /**
     * @notice Returns the operator sets that have pending burn or redistributions.
     * @return operatorSets The operator sets that have pending burn or redistributions.
     */
    function getPendingOperatorSets() external view returns (OperatorSet[] memory operatorSets);

    /**
     * @notice Returns the total number of operator sets with pending burn or redistributions.
     * @return The total number of operator sets with pending burn or redistributions.
     */
    function getTotalPendingOperatorSets() external view returns (uint256);

    /**
     * @notice Returns whether an operator set has pending burn or redistributions.
     * @param operatorSet The operator set whose pending burn or redistributions are being queried.
     * @return Whether the operator set has pending burn or redistributions.
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
     * @param slashId The slash ID of the burn or redistribution that is being queried.
     * @param strategy The strategy whose pending underlying amount is being queried.
     * @return The pending underlying amount.
     */
    function getPendingUnderlyingAmountForStrategy(
        OperatorSet calldata operatorSet,
        uint256 slashId,
        IStrategy strategy
    ) external view returns (uint256);

    /**
     * @notice Returns the paused status of a redistribution.
     * @param operatorSet The operator set whose redistribution is being queried.
     * @param slashId The slash ID of the redistribution that is being queried.
     * @return The paused status of the redistribution.
     */
    function isBurnOrRedistributionPaused(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (bool);

    /**
     * @notice Returns the start block for a slash ID.
     * @param operatorSet The operator set whose start block is being queried.
     * @param slashId The slash ID of the start block that is being queried.
     * @return The start block.
     */
    function getBurnOrRedistributionStartBlock(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint256);

    /**
     * @notice Returns the block at which the escrow can be released.
     * @param operatorSet The operator set whose start block is being queried.
     * @param slashId The slash ID of the start block that is being queried.
     * @return The block at which the escrow can be released.
     */
    function getBurnOrRedistributionCompleteBlock(
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (uint32);

    /**
     * @notice Returns the burn or redistribution delay for a strategy.
     * @param strategy The strategy whose burn or redistribution delay is being queried.
     * @return The burn or redistribution delay.
     */
    function getStrategyBurnOrRedistributionDelay(
        IStrategy strategy
    ) external view returns (uint32);

    /**
     * @notice Returns the global burn or redistribution delay.
     * @return The global burn or redistribution delay.
     */
    function getGlobalBurnOrRedistributionDelay() external view returns (uint32);

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
     * @notice Returns whether a slash escrow is deployed or still counterfactual.
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
