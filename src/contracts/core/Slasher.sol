// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./SlasherStorage.sol";
import "../libraries/EpochUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";
import "../libraries/EIP1271SignatureUtils.sol";

/**
 * @notice Slasher
 */
contract Slasher is SlasherStorage {

    uint256 constant MAGNITUDE_ALLOCATION_DELAY = 21 days;
    uint256 constant MAGNITUDE_DEALLOCATION_DELAY = 21 days;
    uint256 constant MAX_PENDING_MAGNITUDE_UPDATES = 3;

    /**
     * @notice this struct is used in MagnitudeAdjustmentParam in order to specify
     * an operator's slashability for a certain operator set
     * @param operatorSet the operator set to change slashing parameters for
     * @param magnitudeDiff magnitude difference; the difference in proportional parts of the operator's slashable stake
     * that the operator set is getting. This struct is used either in allocating or deallocating
     * slashable stake to an operator set. Slashable stake for an operator set is
     * (slashableMagnitude / sum of all slashableMagnitudes for the strategy/operator + nonSlashableMagnitude) of
     * an operator's delegated stake.
     */
    struct OperatorSetMagnitudeParam {
        OperatorSet operatorSet;
        uint64 magnitudeDiff; 
    }

    /**
     * @notice Input param struct used for functions queueAllocation and queueDeallocation.
     * A structure defining a set of operator-based slashing configurations
     * to manage slashable stake.
     * @param strategy each magnitude param is defined within a single strategy
     * @param operatorSetMagnitudeParams the input magnitude parameters defining the amount
     * the operator set is able to slash an operator. This is passed as an array of
     * operator sets and magnitudes for a given
     */
    struct MagnitudeAdjustmentParam {
        IStrategy strategy;
        OperatorSetMagnitudeParam[] operatorSetMagnitudeParams;
    }

    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => avs => operatorSetId => MagnitudeUpdate[]
     * New updates are pushed whenever queueDeallocations
     * or completeAllocations is called.
     */
    struct MagnitudeUpdate {
        uint32 timestamp;
        uint64 magnitude;
    }
    
    /**
     * @notice Used for historical magnitude updates in mapping
     * operator => IStrategy => TotalAndNonslashableUpdate[]
     * New total magnitude updates are pushed whenever queueDeallocations
     * or completeAllocations is called.
     */
    struct TotalAndNonslashableUpdate {
        uint32 timestamp;
        uint64 totalMagnitude;
        uint64 nonSlashableMagnitude;
    }

    /// EVENTS
    
    event QueuedAllocation(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint32 effectTimestamp,
        uint64 slashableMagnitude
    );
    
    event QueuedDeallocation(
        address operator,
        IStrategy strategy,
        OperatorSet operatorSet,
        uint32 effectTimestamp,
        uint64 slashableMagnitude
    );

    event TotalAndNonSlashableMagnitudeUpdated(
        address operator,
        IStrategy strategy,
        uint32 effectTimestamp,
        uint64 nonSlashableMagnitude,
        uint64 totalSlashableMagnitude
    );

    /// operator => strategy => avs => operatorSetId => MagnitudeUpdate[]
    mapping(address => mapping(IStrategy => mapping(address => mapping(uint32 => MagnitudeUpdate[])))) private _magnitudeUpdates;

    /// operator => strategy => TotalAndNonslashableUpdate[]
    mapping(address => mapping(IStrategy => TotalAndNonslashableUpdate[])) private _totalMagnitudeUpdates;

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager
    ) SlasherStorage(_strategyManager, _delegationManager) {}

    /**
     * @notice Queues slashing magnitude deallocation for an 
     * (operator, strategy, operatorSet) tuple.
     * Queues magnitude deallocation 21 days from now which will be automatically applied
     * at that respective time.
     * 
     * @param operator the operator whom the slashing parameters are being changed
     * @param deallocationParams deallocation params with differences to subtract from current magnitudes
     * @param allocatorSignature if non-empty is the signature of the allocator on 
     * the modification. if empty, the msg.sender must be the allocator for the 
     * operator
     *
     * @dev pushes a MagnitudeUpdate and TotalAndNonslashableUpdate to take effect in 21 days
     * @dev reverts if magnitudeDiff > allocated magnitude
     * since you cannot deallocate more than is already allocated
     * @dev emits event queueDeallocation
     */
    function queueDeallocation(
        address operator,
        MagnitudeAdjustmentParam[] calldata deallocationParams,
        SignatureWithExpiry calldata allocatorSignature // TODO: implement signature verification
    ) external returns(uint32 effectTimestamp) {
        // perform allocator signature verification if not allocator
        if (msg.sender != getAllocatorFor(operator)) {
            // TODO
        }
        effectTimestamp = uint32(block.timestamp + MAGNITUDE_DEALLOCATION_DELAY);

        for (uint256 i = 0; i < deallocationParams.length; ++i) {
            IStrategy strategy = deallocationParams[i].strategy;
            TotalAndNonslashableUpdate storage totalMagnitudeUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategy);
            uint64 nonslashableMagnitude = totalMagnitudeUpdate.nonSlashableMagnitude;
            
            for (uint256 j = 0; j < deallocationParams[i].operatorSetMagnitudeParams.length; ++j) {
                OperatorSet calldata operatorSet = deallocationParams[i].operatorSetMagnitudeParams[j].operatorSet;
                uint64 decrementedMagnitude = deallocationParams[i].operatorSetMagnitudeParams[j].magnitudeDiff;
                // Get either current or queued allocated magnitude for operatorSet
                MagnitudeUpdate storage prevMagnitudeUpdate = _getLatestMagnitudeUpdate(operator, strategy, operatorSet);
                uint64 allocatedMagnitude = prevMagnitudeUpdate.magnitude;
                require(
                    decrementedMagnitude <= allocatedMagnitude,
                    "Slasher.queueDeallocation: magnitudeDiff exceeds allocated magnitude"
                );
                _updateMagnitudeHistory({
                    operator: operator,
                    strategy: strategy,
                    operatorSet: operatorSet,
                    magnitudeUpdate: prevMagnitudeUpdate,
                    newMagnitude: allocatedMagnitude - decrementedMagnitude,
                    effectTimestamp: effectTimestamp
                });
                emit QueuedDeallocation(operator, strategy, operatorSet, effectTimestamp, allocatedMagnitude - decrementedMagnitude);
                // add decremented magnitude back to nonslashableMagnitude
                nonslashableMagnitude += decrementedMagnitude;
            }

            // After updating magnitudes for each operatorSet, push/update a TotalAndNonslashableUpdate with updated nonslashableMagnitude
            _updateTotalAndNonslashableHistory({
                operator: operator,
                strategy: strategy,
                totalAndNonslashableUpdate: totalMagnitudeUpdate,
                newTotalMagnitude: totalMagnitudeUpdate.totalMagnitude,
                newNonslashableMagnitude: nonslashableMagnitude,
                effectTimestamp: effectTimestamp
            });
            emit TotalAndNonSlashableMagnitudeUpdated(
                operator,
                strategy,
                effectTimestamp,
                nonslashableMagnitude,
                totalMagnitudeUpdate.totalMagnitude
            );
        }

    }

    /**
     * @notice Queues slashing magnitude allocation for an
     * (operator, strategy, operatorSet) tuple.
     * Queues magnitude allocation 21 days from now which will be automatically applied
     * at that respective time.
     * 
     * @param operator the operator whom the slashing parameters are being changed
     * @param allocationParams allocation params with differences to add to current magnitudes
     * @param allocatorSignature if non-empty is the signature of the allocator on 
     * the modification. if empty, the msg.sender must be the allocator for the 
     * operator
     *
     * @dev pushes a MagnitudeUpdate and TotalAndNonslashableUpdate to take effect in 21 days
     * overrwrites update if already exists for that effective timestamp
     * @dev reverts if nonslashableMagnitude < total magnitude increase
     * since queued allocation needs to be backed by nonslashableMagnitude portion
     * @dev emits event queueAllocation
     */
    function queueAllocation(
        address operator,
        MagnitudeAdjustmentParam[] calldata allocationParams,
        SignatureWithExpiry calldata allocatorSignature // TODO: implement signature verification
    ) external returns(uint32 effectTimestamp) {
        // perform allocator signature verification if not allocator
        if (msg.sender != getAllocatorFor(operator)) {
            // TODO
        }
        effectTimestamp = uint32(block.timestamp + MAGNITUDE_ALLOCATION_DELAY);

        for (uint256 i = 0; i < allocationParams.length; ++i) {
            // iterate for each strategy
            IStrategy strategy = allocationParams[i].strategy;
            TotalAndNonslashableUpdate storage totalMagnitudeUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategy);
            uint64 nonslashableMagnitude = totalMagnitudeUpdate.nonSlashableMagnitude;

            for (uint256 j = 0; j < allocationParams[i].operatorSetMagnitudeParams.length; ++j) {
                // iterate for each operatorSet for a given strategy
                OperatorSet calldata operatorSet = allocationParams[i].operatorSetMagnitudeParams[j].operatorSet;
                uint64 addedMagnitude = allocationParams[i].operatorSetMagnitudeParams[j].magnitudeDiff;
                
                // check pending updates, if more than 3 pushed updates than revert
                require(checkPendingUpdates(operator, strategy, operatorSet), "Slasher.queueAllocation: too many pending updates");
                // check magnitude increase is backed by nonslashableMagnitude
                require(
                    addedMagnitude <= nonslashableMagnitude,
                    "Slasher.queueAllocation: magnitudeDiff exceeds nonslashableMagnitude"
                );

                // Get either current or queued allocated magnitude for operatorSet to add to
                MagnitudeUpdate storage prevMagnitudeUpdate = _getLatestMagnitudeUpdate(operator, strategy, operatorSet);
                uint64 prevMagnitude = prevMagnitudeUpdate.magnitude;
                _updateMagnitudeHistory({
                    operator: operator,
                    strategy: strategy,
                    operatorSet: operatorSet,
                    magnitudeUpdate: prevMagnitudeUpdate,
                    newMagnitude: prevMagnitude + addedMagnitude,
                    effectTimestamp: effectTimestamp
                });
                emit QueuedAllocation(operator, strategy, operatorSet, effectTimestamp, prevMagnitude + addedMagnitude);
                // subtract from nonslashable which was allocated to this operatorSet
                nonslashableMagnitude -= addedMagnitude;
            }

            // After updating magnitudes for each operatorSet, push/update a TotalAndNonslashableUpdate with updated nonslashableMagnitude
            _updateTotalAndNonslashableHistory({
                operator: operator,
                strategy: strategy,
                totalAndNonslashableUpdate: totalMagnitudeUpdate,
                newTotalMagnitude: totalMagnitudeUpdate.totalMagnitude,
                newNonslashableMagnitude: nonslashableMagnitude,
                effectTimestamp: effectTimestamp
            });
            emit TotalAndNonSlashableMagnitudeUpdated(
                operator,
                strategy,
                effectTimestamp,
                nonslashableMagnitude,
                totalMagnitudeUpdate.totalMagnitude
            );
        }
    }

    /**
     * @notice Add to nonslashable magnitude to dilute relative magnitude proportions
     * of all operator set magnitude allocations. This is an efficient way of adjusting
     * total magnitude and stake slashability risk.
     *
     * @dev TODO hardcode a limit on configurable total Magnitude
     */
    function queueMagnitudeDilution(
        address operator,
        uint64 nonslashableAdded,
        SignatureWithExpiry calldata allocatorSignature
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {

    }

    /**
     * @notice Decrement from nonslashable magnitude to concentrate all relative magnitude
     * proportions for already allocated magnitudes. Efficient way to adjust total magnitude
     * and increase risk expsore across all operatorSets. 
     *
     * @dev reverts if nonslashableDecremented > nonslashableMagnitude
     */
    function queueMagnitudeConcentration(
        address operator,
        uint64 nonslashableDecremented,
        SignatureWithExpiry calldata allocatorSignature
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {

    }

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
    ) external {
        require(
            0 < bipsToSlash && bipsToSlash < SlashingAccountingUtils.BIPS_FACTOR,
            "Slasher._slashRequestedBips: invalid bipsToSlash"
        );

        OperatorSet memory operatorSet = OperatorSet({
            avs: msg.sender,
            id: operatorSetId
        });
        uint32 epoch = EpochUtils.currentEpoch();

        // slash operator for each strategy
        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy strategy = strategies[i];

            // calculate slashingRate based off allocated magnitude
            uint64 slashingRate = uint64(bipsToSlash) * uint64(getSlashableBips(operator, operatorSet, strategy));
            // calculate and apply new scaling factor
            uint64 scalingFactorBefore = shareScalingFactor(operator, strategy);
            uint64 scalingFactorAfter = SlashingAccountingUtils.findNewScalingFactor({
                scalingFactorBefore: scalingFactorBefore,
                rateToSlash: slashingRate
            });
            // update storage to reflect the slashing
            slashingEpochHistory[operator][strategy].push(epoch);
            slashingUpdates[operator][strategy][epoch] = SlashingUpdate({
                slashingRate: slashingRate,
                scalingFactor: scalingFactorAfter
            });
            _shareScalingFactor[operator][strategy] = scalingFactorAfter;
            // TODO DONE?: note that this behavior risks off-by-one errors. it needs to be clearly defined precisely how the historical storage is supposed to work

        }
    }

    /*******************************************************************************
                                 INTERNAL FUNCTIONS
    *******************************************************************************/

    /// @notice push or update the latest magnitude update for the given operator, strategy, and operatorSet
    /// depending on the current effectTimestamp
    function _updateMagnitudeHistory(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet,
        MagnitudeUpdate storage magnitudeUpdate,
        uint64 newMagnitude,
        uint32 effectTimestamp
    ) internal {
        if (magnitudeUpdate.timestamp == effectTimestamp) {
            magnitudeUpdate.magnitude = newMagnitude;
        } else {
            _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id].push(MagnitudeUpdate({
                timestamp: effectTimestamp,
                magnitude: newMagnitude
            }));
        }
    }

    /// @notice push or update the latest TotalAndNonslashable update for the given operator, strategy
    /// depending on the current effectTimestamp
    function _updateTotalAndNonslashableHistory(
        address operator,
        IStrategy strategy,
        TotalAndNonslashableUpdate storage totalAndNonslashableUpdate,
        uint64 newTotalMagnitude,
        uint64 newNonslashableMagnitude,
        uint32 effectTimestamp
    ) internal {
        if (totalAndNonslashableUpdate.timestamp == effectTimestamp) {
            totalAndNonslashableUpdate.totalMagnitude = newTotalMagnitude;
            totalAndNonslashableUpdate.nonSlashableMagnitude = newNonslashableMagnitude;
        } else {
            _totalMagnitudeUpdates[operator][strategy].push(TotalAndNonslashableUpdate({
                timestamp: effectTimestamp,
                totalMagnitude: newTotalMagnitude,
                nonSlashableMagnitude: newNonslashableMagnitude
            }));
        }
    }

    /// @notice Rate limit number of allocations/deallocation updates and return whether a new update could be pushed
    /// for a given (operator, strategy, and operatorSet)
    function checkPendingUpdates(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) public view returns (bool) {
        uint256 magnitudeUpdatesLength = _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id].length;
        if (magnitudeUpdatesLength < MAX_PENDING_MAGNITUDE_UPDATES) {
            return true;
        }

        return _magnitudeUpdates
            [operator]
            [strategy]
            [operatorSet.avs]
            [operatorSet.id]
            [magnitudeUpdatesLength - MAX_PENDING_MAGNITUDE_UPDATES].timestamp < block.timestamp;
    }

    /*******************************************************************************
                                 VIEW FUNCTIONS
    *******************************************************************************/

    /**
     * @notice Get the current slashable bips an operatorSet has for a given operator and strategy
     * @param operator the operator to get the slashable bips for
     * @param operatorSet the operator set to get the slashable bips for
     * @param strategy the strategy to get the slashable bips for
     *
     * @return slashableBips the slashable bips of the given strategy owned by
     * the given OperatorSet for the given operator and epoch
     */
    function getSlashableBips(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy strategy
    ) public view returns (uint16 slashableBips) {
        MagnitudeUpdate memory magnitudeUpdate = getMagnitudeUpdate(
            operator,
            strategy,
            operatorSet,
            uint32(block.timestamp)
        );
        TotalAndNonslashableUpdate memory totalAndNonslashableUpdate = getTotalAndNonslashableUpdate(
            operator,
            strategy,
            uint32(block.timestamp)
        );

        return uint16(magnitudeUpdate.magnitude * SlashingAccountingUtils.BIPS_FACTOR / totalAndNonslashableUpdate.totalMagnitude);
    }

    /// @notice Get the magnitude for a given operator, strategy, operatorSet, timestamp
    function getMagnitudeUpdate(
        address operator,
        IStrategy strategy,
        OperatorSet memory operatorSet,
        uint32 timestamp
    ) public view returns (MagnitudeUpdate memory magnitudeUpdate) {
        uint256 magnitudeUpdatesLength = _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id].length;
        for (uint256 i = magnitudeUpdatesLength; i > 0; --i) {
            if (_magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id][i - 1].timestamp <= timestamp) {
                magnitudeUpdate = _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id][i - 1];
                break;
            }
        }
        revert("Slasher.getMagnitudeUpdate: no magnitude update found at timestamp");
    }

    /// @notice Get the total and nonslashable magnitudes for a given operator, strategy, timestamp
    function getTotalAndNonslashableUpdate(
        address operator,
        IStrategy strategy,
        uint32 timestamp
    ) public view returns (TotalAndNonslashableUpdate memory totalAndNonslashableUpdate) {
        uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][strategy].length;
        for (uint256 i = totalMagnitudeUpdatesLength; i > 0; --i) {
            if (_totalMagnitudeUpdates[operator][strategy][i - 1].timestamp <= timestamp) {
                totalAndNonslashableUpdate = _totalMagnitudeUpdates[operator][strategy][i - 1];
                break;
            }
        }
        revert("Slasher.getTotalAndNonslashableUpdate: no totalandnonslashable magnitude update found at timestamp");
    }

    /// @notice Return latest magnitude update for the given operator, strategy, and operatorSet
    /// if empty array, push an empty update
    function _getLatestMagnitudeUpdate(
        address operator,
        IStrategy strategy,
        OperatorSet calldata operatorSet
    ) internal returns (MagnitudeUpdate storage) {
        uint256 magnitudeUpdatesLength = _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id].length;
        if (magnitudeUpdatesLength == 0) {
            _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id].push(MagnitudeUpdate({
                timestamp: 0,
                magnitude: 0
            }));
        }
        return _magnitudeUpdates[operator][strategy][operatorSet.avs][operatorSet.id][magnitudeUpdatesLength - 1];
    }

    /// @notice Return latest total and nonslashable update for the given operator and strategy
    /// if empty array, push an empty update
    function _getLatestTotalAndNonslashableUpdate(
        address operator,
        IStrategy strategy
    ) internal returns (TotalAndNonslashableUpdate storage) {
        uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][strategy].length;
        if (totalMagnitudeUpdatesLength == 0) {
            _totalMagnitudeUpdates[operator][strategy].push(TotalAndNonslashableUpdate({
                timestamp: 0,
                totalMagnitude: 0,
                nonSlashableMagnitude: 0
            }));
        }
        return _totalMagnitudeUpdates[operator][strategy][totalMagnitudeUpdatesLength - 1];
    }

    /**
     * @notice gets the scaling factor for the given operator and strategy
     * @param operator the operator to get the scaling factor for
     * @param strategy the strategy to get the scaling factor for
     * @return the scaling factor for the given operator and strategy
     */
    function shareScalingFactor(address operator, IStrategy strategy) public view returns (uint64) {
        uint64 scalingFactor = _shareScalingFactor[operator][strategy];
        if (scalingFactor == 0) {
            return SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        } else {
            return scalingFactor;
        }
    }

    // TODO: this is a "naive" search since it brute-force backwards searches; we might technically want a binary search for completeness
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
    ) public view returns (uint64) {
        uint64 scalingFactor = SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        (uint32 lookupEpoch, bool found) = _getLookupEpoch(operator, strategy, epoch);
        if (found) {
            scalingFactor = slashingUpdates[operator][strategy][lookupEpoch].scalingFactor;
        }
        return scalingFactor;
    }

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
	) external view returns (uint64) {
        uint64 slashingRate;
        (uint32 lookupEpoch, bool found) = _getLookupEpoch(operator, strategy, epoch);
        if (found) {
            slashingRate = slashingUpdates[operator][strategy][lookupEpoch].slashingRate;
        }
        return slashingRate;
    }

    // @notice Returns the epoch in which the operator was last slashed
    function lastSlashed(address operator, IStrategy strategy) public view returns (uint32) {
        uint256 slashingEpochHistoryLength = slashingEpochHistory[operator][strategy].length;
        if (slashingEpochHistoryLength == 0) {
            // TODO: consider if a different return value is more appropriate for this special case
            return 0;
        } else {
            return slashingEpochHistory[operator][strategy][slashingEpochHistoryLength - 1];
        }
    }

    /// @notice Read from operator details in DelegationManager to get operator's allocator
    function getAllocatorFor(address operator) public view returns (address) {
        return delegation.operatorDetails(operator).allocator;
    }

    /// @notice Returns the epoch for lookup in the slashing history and whether
    /// there were any slashing requests that happened before or at the given epoch
    function _getLookupEpoch(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) internal view returns (uint32, bool) {
        // TODO DONE?: note the edge case of 0th epoch; need to make sure it's clear how it should be handled
        uint32 epochForLookup;
        bool found;
        for (uint256 i = slashingEpochHistory[operator][strategy].length; i > 0; --i) {
            epochForLookup = slashingEpochHistory[operator][strategy][i - 1];
            if (epochForLookup <= epoch) {
                found = true;
                break;
            }
        }
        return (epochForLookup, found);
    }
}
