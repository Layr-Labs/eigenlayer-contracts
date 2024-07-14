// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/ISlasher.sol";


contract SlasherMock is ISlasher, Test {
    function strategyManager() external view override returns (IStrategyManager) {}
    function delegation() external view override returns (IDelegationManager) {}
	
	/**
     * @notice Queues magnitude adjustment updates of type ALLOCATION or DEALLOCATION
     * The magnitude allocation takes 21 days from time when it is queued to take effect.
     *
     * @param operator the operator whom the magnitude parameters are being adjusted
     * @param adjustmentParams allocation adjustment params with differences to add/subtract to current magnitudes
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @dev pushes a MagnitudeUpdate and TotalAndNonslashableUpdate to take effect in 21 days
     * @dev If ALLOCATION adjustment type:
     * reverts if sum of magnitudeDiffs > nonslashable magnitude for the latest pending update - sum of all pending allocations
     * since one cannot allocate more than is nonslahsable
     * @dev if DEALLOCATION adjustment type:
     * reverts if magnitudeDiff > allocated magnitude for the latest pending update
     * since one cannot deallocate more than is already allocated
     * @dev reverts if there are more than 3 pending allocations/deallocations for the given (operator, strategy, operatorSet) tuple
     * @dev emits events MagnitudeUpdated, TotalAndNonslashableMagnitudeUpdated
     */
    function queueReallocation(
        address operator,
        MagnitudeAdjustmentsParam[] calldata adjustmentParams,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns (uint32 effectTimestamp) {}

    /**
     * @notice Add to nonslashable magnitude to dilute relative magnitude proportions
     * of all operator set magnitude allocations. This is an efficient way of adjusting
     * total magnitude and stake slashability risk.
     *
     * @dev TODO hardcode a limit on configurable total Magnitude
     */
    function queueMagnitudeDilution(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToAdd,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {}

    /**
     * @notice For a strategy, decrement from nonslashable magnitude to concentrate all relative magnitude
     * proportions for already allocated magnitudes. Efficient way to adjust total magnitude
     * and increase risk expsore across all operatorSets. 
     *
     * @dev reverts if nonslashableDecremented > nonslashableMagnitude
     */
    function queueMagnitudeConcentration(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToDecrement,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {}

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
        uint32 operatorSetId,
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
		OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint64) {}
	
    function lastSlashed(address operator, IStrategy strategy) public view returns (uint32) {}
}
