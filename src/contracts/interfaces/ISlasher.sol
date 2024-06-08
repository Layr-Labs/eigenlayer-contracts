// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategyManager.sol";
import "./IDelegationManager.sol";
import "./IOperatorSetManager.sol";

/**
 * @title Interface for the primary 'slashing' contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `Slasher` contract itself for implementation details.
 */
interface ISlasher {

    /// EVENTS
	
	event RequestedBipsToSlashModified(
		uint32 epoch,
		address operator,
        IOperatorSetManager.OperatorSet operatorSet,
        IStrategy[] strategies,
        int32 bipsToModify
    );
    
    event SlashingExecuted(
        uint32 epoch, // epoch in which the slashing was requested
        address operator,
        IStrategy strategy,
        uint64 slashingRate
    );

    /// EXTERNAL - STATE MODIFYING
	
	// TODO: possibly add a return amount for the bips that this increased the `pendingSlashedBips`?
	/**
	 * @notice Called by an AVS to increase its own slashing request for a given
	 * operator set and operator in the current epoch
	 *
	 * @param operator the operator that the calling AVS is to increase the bips they want to slash
	 * @param operatorSetID the id of the operator set the AVS is increasing their slashing for
	 * @param strategies the list of strategies slashing requested is being modified for
	 * @param bipsToIncrease the basis points slashing to modify for given strategies
	 *
	 * @dev bipsToModify must be positive
	 */
	function increaseRequestedBipsToSlash(
		address operator,
		bytes4 operatorSetID,
		IStrategy[] memory strategies,
		int32 bipsToIncrease
	) external;

	/**
	 * @notice Called by an AVS to reduce its own slashing request for a given
	 * operator set and operator in the current or previous epoch
	 *
	 * @param operator the operator that the calling AVS is to reduce the bips they want to slash
	 * @param operatorSetID the id of the operator set the AVS is reducing their slashing for
	 * @param strategies the list of strategies slashing requested is being reduced for
	 * @param epoch the epoch in which slashing was requested
	 * @param bipsToReduce the basis points slashing to reduced for given strategies
	 *
	 * @dev bipsToReduce must be negative
	 */
	function reduceRequestedBipsToSlash(
		address operator,
		bytes4 operatorSetID,
		IStrategy[] memory strategies,
		uint32 epoch,
		int32 bipsToReduce
	) external;
	
	/**
	 * @notice Permissionlessly called to execute slashing of a given list of 
	 * strategies for a given operator, for the latest unslashed epoch
	 *
	 * @param operator the operator to slash
	 * @param strategies the list of strategies to execute slashing for
	 * @param epoch the epoch in which the slashing requests to execute were made
	 */
	function executeSlashing(
		address operator, 
		IStrategy[] memory strategies,
		uint32 epoch
	) external;
	
	/// VIEW

	function strategyManager() external view returns (IStrategyManager);
    function delegation() external view returns (IDelegationManager);
    function operatorSetManager() external view returns (IOperatorSetManager);
	
	/**
	 * @notice fetches the requested parts per hundred million to slash for the 
	 * given operator, strategy, epoch, and operator set
	 *
	 * @param operator the operator to get the requested slashing rate for
	 * @param strategy the strategy to get the requested slashing rate for
	 * @param operatorSet the operator set to get the requested requested slashing rate for
	 * @param epoch the epoch to get the requested slashing rate  for
	 * 
	 * @return the requested parts per hundred million to slash for the given 
	 * operator, strategy, epoch, and operator set
	 * 
	 * @dev may exceed the AVS operator set's allowed slashing per epoch; 
	 * the `getPendingSlashedPPHM` will accurately reflect this ceiling though.
	 */
	function getRequestedSlashingRate(
		address operator, 
		IStrategy strategy, 
		IOperatorSetManager.OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint32);
	
	/**
	 * @notice fetches the parts per hundred million that will be slashed for the 
	 * given operator, strategy, epoch, and operator set assuming no further 
	 * modifications to requested slashing rate by operatorSet
	 *
	 * @param operator the operator to get the pending slashing rate for 
	 * @param strategy the strategy to get the pending slashing rate for
	 * @param operatorSet the operator set to get the pending slashing rate for
	 *
	 * @return the parts per hundred million that will be slashed for the given 
	 * operator, strategy, epoch, and operator set assuming no further 
	 * modifications to requested slashing rate by operatorSet
	 */
	function getPendingSlashingRate(
		address operator, 
		IStrategy strategy,
		IOperatorSetManager.OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint32);
	
	/**
	 * @notice fetches the parts per hundred million that will be slashed for 
	 * the given operator, strategy, and epoch, across all operator set assuming 
	 * no more modifications to requested slashing rate for the operator.
	 *
	 * @param operator the operator to get the pending slashing rate for
	 * @param strategy the strategy to get the pending slashing rate for
	 * @param epoch the epoch to get the pending slashing rate for
	 * 
	 * @return the parts per hundred million that will be slashed for the 
	 * given operator, strategy, and epoch, across all operator set assuming 
	 * no more modifications to requested slashing rate for the operator.
	 */
	function getTotalPendingSlashingRate(
		address operator, 
		IStrategy strategy,
		uint32 epoch
	) external view returns (uint32);

	/**
     * @notice gets whether withdrawals of the given strategy delegated to the given operator can be withdrawn and the scaling factor
     * @param operator the operator the withdrawal is delegated to
     * @param strategy the strategy the withdrawal is from
     * @param epoch the last epoch the withdrawal was slashable until
     * @return whether the withdrawal can be executed
     * @return whether there was a slashing request for the given operator and strategy at the given epoch
     */
    function getWithdrawabilityAndScalingFactorAtEpoch(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) external view returns (bool, uint64);

	/**
     * @notice gets whether withdrawals of the given strategy delegated to the given operator can be withdrawn
     * @param operator the operator the withdrawal is delegated to
     * @param strategy the strategy the withdrawal is from
     * @param epoch the last epoch the withdrawal was slashable until
     * @return whether the withdrawal can be executed
     */
    function canWithdraw(address operator, IStrategy strategy, uint32 epoch) external view returns (bool);

    /**
     * @notice gets the scaling factor for the given operator and strategy
     * @param operator the operator to get the scaling factor for
     * @param strategy the strategy to get the scaling factor for
     * @return the scaling factor for the given operator and strategy
     */
    function shareScalingFactor(address operator, IStrategy strategy) external view returns (uint64);

    // TODO: documentation
    function pendingShareScalingFactor(address operator, IStrategy strategy) external view returns (uint64);

    /**
     * @notice gets the scaling factor for the given operator and strategy at the given epoch
     * @param operator the operator to get the scaling factor for
     * @param strategy the strategy to get the scaling factor for
     * @param epoch the epoch to get the scaling factor for
     * @return the scaling factor for the given operator and strategy at the given epoch
     */
    function shareScalingFactorAtEpoch(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) external view returns (uint64);

}