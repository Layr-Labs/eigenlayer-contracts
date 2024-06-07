// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/ISlasher.sol";


contract SlasherMock is ISlasher, Test {
    function strategyManager() external view override returns (IStrategyManager) {}
    function delegation() external view override returns (IDelegationManager) {}
    function operatorSetManager() external view override returns (IOperatorSetManager) {}

    function shareScalingFactor(address operator, IStrategy strategy) external view returns (uint256) {}
    function pendingShareScalingFactor(address operator, IStrategy strategy) public view returns (uint256) {}
    function shareScalingFactorAtEpoch(address operator, IStrategy strategy, uint32 epoch) public view returns (uint256) {}

    function modifyRequestedBipsToSlash(
        address operator,
        bytes4 operatorSetID,
        IStrategy[] memory strategies,
        int32 bipsToModify
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
