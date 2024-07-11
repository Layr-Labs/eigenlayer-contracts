// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStrategyManager.sol";
import "./IDelegationManager.sol";
import "./ISignatureUtils.sol";
import "./IStrategy.sol";

/**
 * @title Interface for the primary 'slashing' contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice See the `Slasher` contract itself for implementation details.
 */
interface ISlasher is ISignatureUtils {

	struct OperatorSet {
        address avs;
        uint32 id;
    }

    /// EVENTS

    /// EXTERNAL - STATE MODIFYING

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
    ) external;
	
	/// VIEW

	function strategyManager() external view returns (IStrategyManager);
    function delegation() external view returns (IDelegationManager);

    /**
     * @notice gets the scaling factor for the given operator and strategy
     * @param operator the operator to get the scaling factor for
     * @param strategy the strategy to get the scaling factor for
     * @return the scaling factor for the given operator and strategy
     */
    function shareScalingFactor(address operator, IStrategy strategy) external view returns (uint64);

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

	/**
	 * @notice fetches the requested parts per hundred million to slash for the
	 * given operator, strategy, epoch, and operator set
	 *
	 * @param operator the operator to get the requested slashing rate for
	 * @param strategy the strategy to get the requested slashing rate for
	 * @param operatorSet the operator set to get the requested requested slashing rate for
	 * @param epoch the epoch to get the slashing bips for
	 *
	 * @return slashingRate parts per hundred million to slash for the given
	 * operator, strategy, epoch, operator set, and magnitude.
	 */
	function getSlashedRate(
		address operator,
		IStrategy strategy,
		OperatorSet calldata operatorSet,
		uint32 epoch
	) external view returns (uint64);
}
