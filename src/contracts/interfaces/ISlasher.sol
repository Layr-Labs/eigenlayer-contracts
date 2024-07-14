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

    /*******************************************************************************
                                 STRUCTS
    *******************************************************************************/

	struct OperatorSet {
        address avs;
        uint32 id;
    }

    enum MagnitudeAdjustmentType {
        ALLOCATION,
        DEALLOCATION
    }

    /**
     * @notice this struct is used in MagnitudeAdjustmentsParam in order to specify
     * an operator's slashability for a certain operator set
     *
     * @param operatorSet the operator set to change slashing parameters for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that the operator set is getting. This struct is used either in allocating or deallocating
     * slashable stake to an operator set. Slashable stake for an operator set is
     * (slashableMagnitude / sum of all slashableMagnitudes for the strategy/operator + nonslashableMagnitude) of
     * an operator's delegated stake.
     */
    struct MagnitudeAdjustment {
        OperatorSet operatorSet;
        uint64 magnitudeDiff;
    }

    /**
     * @notice Input param struct used for functions queueAllocation and queueDeallocation.
     * A structure defining a set of operator-based slashing configurations
     * to manage slashable stake.
     * @param strategy the strategy to update magnitudes for
     * @param magnitudeAdjustmentType the type of adjustment to make to the operator's slashable stake
     * @param magnitudeAdjustments the input magnitude parameters defining the adjustments to the proportion
     * the operator set is able to slash an operator. This is passed as an array of
     * operator sets, magnitudes diffs, MagnitudeAdjustmentType enum for a given strategy
     */
    struct MagnitudeAdjustmentsParam {
        IStrategy strategy;
        MagnitudeAdjustmentType magnitudeAdjustmentType;
        MagnitudeAdjustment[] magnitudeAdjustments;
    }

    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => avs => operatorSetId => MagnitudeUpdate[]
     * New updates are pushed whenever queueDeallocations or queueAllocations is called
     * @param timestamp timestamp of MagnitudeUpdate, if timestamp > block.timestamp then it is currently pending
     * @param magnitude the magnitude/proportion value of the (operator, Strategy, operatorSet) at timestamp
     */
    struct MagnitudeUpdate {
        uint32 timestamp;
        uint64 magnitude;
    }
    
    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => TotalMagnitudeUpdate[]
     * New total magnitude updates are pushed whenever magnitude changing functions are called
     * @param timestamp timestamp of TotalAndNonslashableUpdate, if timestamp > block.timestamp then it is currently pending
     * @param totalMagnitude total magnitude amount for a strategy which equals
     * nonslashableMagnitude amount + sum of each operatorSet's allocated magnitudes
     * @param nonslashableMagnitude nonslashable magnitude that CANNOT be slashed if timestamp <= block.timestamp
     * Note if timestamp > block.timestamp, this is an upper bound on the slashable amount
     * as it may still be a pending deallocation (still slashable)
     * @param cumulativeAllocationSum monotonically increasing sum of all magnitudes when allocating
     * required to ensure all allocations are backed by nonslashable magnitude
     */
    struct TotalAndNonslashableUpdate {
        uint32 timestamp;
        uint64 totalMagnitude;
        uint64 nonslashableMagnitude;
        uint64 cumulativeAllocationSum;
    }

    /*******************************************************************************
                                EVENTS
    *******************************************************************************/

    event MagnitudeUpdated(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint32 effectTimestamp,
        uint64 slashableMagnitude
    );

    event TotalAndNonslashableMagnitudeUpdated(
        address operator,
        IStrategy strategy,
        uint32 effectTimestamp,
        uint64 totalSlashableMagnitude,
        uint64 nonslashableMagnitude,
        uint64 cumulativeAllocationSum
    );

    /// @notice emitted when slashOperator is called 
    event OperatorSlashed(
        address operator,
        OperatorSet operatorSet,
        uint32 bipsToSlash,
        IStrategy[] strategies,
        uint64[] slashingRates
    );
    event MagnitudeDecremented(
        address operator,
        OperatorSet operatorSet,
        IStrategy strategy,
        uint64 updatedSlashableMagnitude,
        uint32 effectTimestamp
    );
    event NonslashableMagnitudeDecremented(
        address operator,
        IStrategy strategy,
        uint64 updatedNonslashableMagnitude,
        uint32 effectTimestamp
    );

    /*******************************************************************************
                                EXTERNAL FUNCTIONS
    *******************************************************************************/

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
    ) external returns (uint32 effectTimestamp);

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
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude);

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
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude);

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
