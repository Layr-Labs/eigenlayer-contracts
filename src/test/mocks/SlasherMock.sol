// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/ISlasher.sol";


contract SlasherMock is ISlasher, Test {
    function strategyManager() external view override returns (IStrategyManager) {}
    function delegation() external view override returns (IDelegationManager) {}
    function operatorSetManager() external view override returns (IOperatorSetManager) {}

    function shareScalingFactor(address operator, IStrategy strategy) external view returns (uint64) {}
    function pendingShareScalingFactor(address operator, IStrategy strategy) public view returns (uint64) {}
    function shareScalingFactorAtEpoch(address operator, IStrategy strategy, uint32 epoch) public view returns (uint64) {}

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
	) external {}

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
	) external {}
	
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
	) external {}
	
	/// VIEW
	function getRequestedSlashingRate(
		address operator, 
		IStrategy strategy, 
		IOperatorSetManager.OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint32) {}

	function getPendingSlashingRate(
		address operator, 
		IStrategy strategy,
		IOperatorSetManager.OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint32) {}

	function getTotalPendingSlashingRate(
		address operator, 
		IStrategy strategy,
		uint32 epoch
	) external view returns (uint32) {}

}
