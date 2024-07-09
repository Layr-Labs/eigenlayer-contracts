// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/ISlasher.sol";


contract SlasherMock is ISlasher, Test {
    function strategyManager() external view override returns (IStrategyManager) {}
    function delegation() external view override returns (IDelegationManager) {}
    function operatorSetManager() external view override returns (IOperatorSetManager) {}
	
	/**
	 * @notice Called by an AVS to slash an operator for given operatorSetId,
	 * list of strategies, and bipsToSlash.
	 * For each given (operator, operatorSetId, Strategy) tuple, bipsToSlash will be used to slash.
	 * @param operator address to slash
	 * @param operatorSetId which operator set operator is being slashed from
	 * @param strategies set of strategies to slash
	 * @param bipsToSlash number of bips to slash, this will be proportional to the
	 * operator's configured magnitude for the operatorSet
	 */
    function slashOperator(
        address operator,
        bytes4 operatorSetId,
        IStrategy[] memory strategies,
        uint32 bipsToSlash
    ) external {}
	
	/// VIEW

    function shareScalingFactor(address operator, IStrategy strategy) public view returns (uint64) {}


	function shareScalingFactorAtEpoch(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) public view returns (uint64) {}

	function getSlashedRate(
		address operator,
		IStrategy strategy,
		IOperatorSetManager.OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint64) {}
	
    function lastSlashed(address operator, IStrategy strategy) public view returns (uint32) {}
}
