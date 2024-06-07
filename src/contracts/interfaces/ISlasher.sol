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
    function strategyManager() external view returns (IStrategyManager);
    function delegation() external view returns (IDelegationManager);
    function operatorSetManager() external view returns (IOperatorSetManager);

    /// EXTERNAL - STATE MODIFYING
	
	// TODO: possibly add a return amount for the bips that this increased the `pendingSlashedBips`?
	/**
	 * @notice Called by an AVS to modify its own slashing request for a given
	 * operator set and operator in the current epoch
	 *
	 * @param operator the operator that the calling AVS is to modify the bips they want to slash
	 * @param operatorSetID the id of the operator set the AVS is modifying their slashing for
	 * @param strategies the list of strategies slashing requested is being modified for
	 * @param bipsToModify the basis points slashing to modify for given strategies
	 *
	 * @dev bipsToModify is negative when the AVS wants to reduce the amount of slashing
     	 *      and positive when the AVS wants to increase the amount of slashing
	 */
    function modifyRequestedBipsToSlash(
        address operator,
        bytes4 operatorSetID,
        IStrategy[] memory strategies,
        int32 bipsToModify
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
	 * modifications to requested slashed bips by operatorSet
	 *
	 * @param operator the operator to get the pending slashing rate for 
	 * @param strategy the strategy to get the pending slashing rate for
	 * @param operatorSet the operator set to get the pending slashing rate for
	 *
	 * @return the parts per hundred million that will be slashed for the given 
	 * operator, strategy, epoch, and operator set assuming no further 
	 * modifications to requested slashed bips by operatorSet
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
	 * no more modifications to requested slashed bips for the operator.
	 *
	 * @param operator the operator to get the pending slashing rate for
	 * @param strategy the strategy to get the pending slashing rate for
	 * @param epoch the epoch to get the pending slashing rate for
	 * 
	 * @return the parts per hundred million that will be slashed for the 
	 * given operator, strategy, and epoch, across all operator set assuming 
	 * no more modifications to requested slashed bips for the operator.
	 */
	function getTotalPendingSlashingRate(
		address operator, 
		IStrategy strategy,
		uint32 epoch
	) external view returns (uint32);

    /// VIEW
    // TODO: documentation
    function shareScalingFactor(address operator, IStrategy strategy) external view returns (uint256);

    // TODO: documentation
    function pendingShareScalingFactor(address operator, IStrategy strategy) external view returns (uint256);

    // TODO: documentation
    function shareScalingFactorAtEpoch(address operator, IStrategy strategy, uint32 epoch) external view returns (uint256);

}