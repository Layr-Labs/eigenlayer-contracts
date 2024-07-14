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
    uint256 constant MAX_PENDING_TOTAL_MAGNITUDE_UPDATES = 6;
    /// TODO Figure out reasonable limit for this
    uint256 constant FIXED_TOTAL_MAGNITUDE_LIMIT = 1e28;

    /// @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /// operator => strategy => avs => operatorSetId => MagnitudeUpdate[]
    mapping(address => mapping(IStrategy => mapping(address => mapping(uint32 => MagnitudeUpdate[])))) private _magnitudeUpdates;

    /// operator => strategy => TotalAndNonslashableUpdate[]
    mapping(address => mapping(IStrategy => TotalAndNonslashableUpdate[])) private _totalMagnitudeUpdates;

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager
    ) SlasherStorage(_strategyManager, _delegationManager) {
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _transferOwnership(initialOwner);
    }

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
    ) external returns (uint32 effectTimestamp) {
        // perform allocator signature verification if not allocator
        address allocator = getAllocatorFor(operator);
        if (msg.sender != allocator) {
            bytes32 digestHash = calculateReallocationDigestHash(
                operator,
                adjustmentParams,
                allocatorSignature.salt,
                allocatorSignature.expiry
            );
            _verifyAllocatorSignature(allocator, digestHash, allocatorSignature);
        }

        effectTimestamp = uint32(block.timestamp + MAGNITUDE_ALLOCATION_DELAY);
        for (uint256 i = 0; i < adjustmentParams.length; ++i) {
            if (adjustmentParams[i].magnitudeAdjustmentType == MagnitudeAdjustmentType.ALLOCATION) {
                _queueAllocation(operator, adjustmentParams[i], effectTimestamp);
            } else {
                _queueDeallocation(operator, adjustmentParams[i], effectTimestamp);
            }
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
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToAdd,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {
        // perform allocator signature verification if not allocator
        address allocator = getAllocatorFor(operator);
        if (msg.sender != allocator) {
            bytes32 digestHash = calculateMagnitudeDilutionDigestHash(
                operator,
                strategies,
                nonslashableToAdd,
                allocatorSignature.salt,
                allocatorSignature.expiry
            );
            _verifyAllocatorSignature(allocator, digestHash, allocatorSignature);
        }

        uint32 effectTimestamp = uint32(block.timestamp + MAGNITUDE_ALLOCATION_DELAY);
        for (uint256 i = 0; i < strategies.length; ++i) {
            require(
                checkPendingTotalMagnitudeUpdates(operator, strategies[i]),
                "Slasher.queueMagnitudeConcentration: too many pending total magnitude updates"
            );
            TotalAndNonslashableUpdate storage totalMagnitudeUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategies[i]);
            uint64 nonslashableMagnitude = totalMagnitudeUpdate.nonslashableMagnitude;
            uint64 totalMagnitude = totalMagnitudeUpdate.totalMagnitude;

            require(
                nonslashableToAdd[i] <= FIXED_TOTAL_MAGNITUDE_LIMIT,
                "Slasher.queueMagnitudeConcentration: nonslashableAdded exceeds FIXED_TOTAL_MAGNITUDE_LIMIT"
            );
            newTotalMagnitude = totalMagnitude + nonslashableToAdd[i];
            newNonslashableMagnitude = nonslashableMagnitude + nonslashableToAdd[i];
            // update total magnitude with decremented nonslashableMagnitude
            _updateTotalAndNonslashableHistory({
                operator: operator,
                strategy: strategies[i],
                totalAndNonslashableUpdate: totalMagnitudeUpdate,
                newTotalMagnitude: newTotalMagnitude,
                newNonslashableMagnitude: newNonslashableMagnitude,
                newCumulativeAllocationSum: totalMagnitudeUpdate.cumulativeAllocationSum,
                effectTimestamp: effectTimestamp
            });
        }
    }

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
    ) external returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude) {
        // perform allocator signature verification if not allocator
        address allocator = getAllocatorFor(operator);
        if (msg.sender != allocator) {
            bytes32 digestHash = calculateMagnitudeConcentrationDigestHash(
                operator,
                strategies,
                nonslashableToDecrement,
                allocatorSignature.salt,
                allocatorSignature.expiry
            );
            _verifyAllocatorSignature(allocator, digestHash, allocatorSignature);
        }

        uint32 effectTimestamp = uint32(block.timestamp + MAGNITUDE_ALLOCATION_DELAY);
        for (uint256 i = 0; i < strategies.length; ++i) {
            require(
                checkPendingTotalMagnitudeUpdates(operator, strategies[i]),
                "Slasher.queueMagnitudeConcentration: too many pending total magnitude updates"
            );
            TotalAndNonslashableUpdate storage totalMagnitudeUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategies[i]);
            uint64 nonslashableMagnitude = totalMagnitudeUpdate.nonslashableMagnitude;
            uint64 totalMagnitude = totalMagnitudeUpdate.totalMagnitude;

            require(
                nonslashableToDecrement[i] <= nonslashableMagnitude,
                "Slasher.queueMagnitudeConcentration: nonslashableDecremented exceeds nonslashableMagnitude"
            );
            newTotalMagnitude = totalMagnitude - nonslashableToDecrement[i];
            newNonslashableMagnitude = nonslashableMagnitude - nonslashableToDecrement[i];
            // update total magnitude with decremented nonslashableMagnitude
            _updateTotalAndNonslashableHistory({
                operator: operator,
                strategy: strategies[i],
                totalAndNonslashableUpdate: totalMagnitudeUpdate,
                newTotalMagnitude: newTotalMagnitude,
                newNonslashableMagnitude: newNonslashableMagnitude,
                newCumulativeAllocationSum: totalMagnitudeUpdate.cumulativeAllocationSum,
                effectTimestamp: effectTimestamp
            });
        }
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

    function _verifyAllocatorSignature(
        address allocator,
        bytes32 allocatorDigestHash,
        SignatureWithSaltAndExpiry calldata allocatorSignature
    ) internal {
        // check the signature expiry
        require(
            allocatorSignature.expiry >= block.timestamp,
            "Slasher._verifyAllocatorSignature: allocator signature expired"
        );
        // Assert allocator's signature cannot be replayed.
        require(
            !allocatorSaltIsSpent[allocator][allocatorSignature.salt],
            "AVSDirectory.updateStandbyParams: salt spent"
        );

        // Assert allocator's signature is valid.
        EIP1271SignatureUtils.checkSignature_EIP1271(
            allocator,
            allocatorDigestHash,
            allocatorSignature.signature
        );
        // Spend salt.
        allocatorSaltIsSpent[allocator][allocatorSignature.salt] = true;
    }

    /// TODO natspec
    function _queueAllocation(
        address operator,
        MagnitudeAdjustmentsParam calldata allocationParam,
        uint32 effectTimestamp
    ) internal {
        IStrategy strategy = allocationParam.strategy;
        TotalAndNonslashableUpdate storage latestTotalUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategy);
        TotalAndNonslashableUpdate memory currentTotalUpdate = getTotalAndNonslashableUpdate(operator, strategy, uint32(block.timestamp));
        // notice that we need to take the current available nonslashableMagnitude subtracted the pending queued allocations
        // since the queued allocations will also be backed by the current non slashable magnitude and we can't double count the current
        // nonslashable magnitude
        // shouldn't revert since amount is cumulative and monotonically increasing
        // and pending allocations previously queued could only be backed from current nonslashableMagnitude
        uint64 nonslashableMagnitude = currentTotalUpdate.nonslashableMagnitude - (latestTotalUpdate.cumulativeAllocationSum - currentTotalUpdate.cumulativeAllocationSum);
        uint64 allocationSum = 0;

        for (uint256 i = 0; i < allocationParam.magnitudeAdjustments.length; ++i) {
            // iterate for each operatorSet for a given strategy
            OperatorSet calldata operatorSet = allocationParam.magnitudeAdjustments[i].operatorSet;
            uint64 addedMagnitude = allocationParam.magnitudeAdjustments[i].magnitudeDiff;
            
            // check pending updates, if more than 3 pushed updates than revert
            require(checkPendingMagnitudeUpdates(operator, strategy, operatorSet), "Slasher._queueAllocation: too many pending updates");
            // check magnitude increase is backed by nonslashableMagnitude
            require(
                addedMagnitude <= nonslashableMagnitude,
                "Slasher._queueAllocation: magnitudeDiff exceeds nonslashableMagnitude"
            );

            // Get either current or queued allocated magnitude for operatorSet to add to
            MagnitudeUpdate storage prevMagnitudeUpdate = _getLatestMagnitudeUpdate(operator, strategy, operatorSet);
            _updateMagnitudeHistory({
                operator: operator,
                strategy: strategy,
                operatorSet: operatorSet,
                magnitudeUpdate: prevMagnitudeUpdate,
                newMagnitude: prevMagnitudeUpdate.magnitude + addedMagnitude,
                effectTimestamp: effectTimestamp
            });
            // subtract from nonslashable which was allocated to this operatorSet
            nonslashableMagnitude -= addedMagnitude;
            allocationSum += addedMagnitude;
        }

        // After updating magnitudes for each operatorSet, push/update a TotalAndNonslashableUpdate with updated queued nonslashableMagnitude
        _updateTotalAndNonslashableHistory({
            operator: operator,
            strategy: strategy,
            totalAndNonslashableUpdate: latestTotalUpdate,
            newTotalMagnitude: latestTotalUpdate.totalMagnitude,
            newNonslashableMagnitude: nonslashableMagnitude,
            newCumulativeAllocationSum: latestTotalUpdate.cumulativeAllocationSum + allocationSum,
            effectTimestamp: effectTimestamp
        });
    }

    /// TODO natspec
    function _queueDeallocation(
        address operator,
        MagnitudeAdjustmentsParam calldata deallocationParam,
        uint32 effectTimestamp
    ) internal {
        IStrategy strategy = deallocationParam.strategy;
        TotalAndNonslashableUpdate storage latestTotalUpdate = _getLatestTotalAndNonslashableUpdate(operator, strategy);
        uint64 nonslashableMagnitude = latestTotalUpdate.nonslashableMagnitude;
        
        for (uint256 i = 0; i < deallocationParam.magnitudeAdjustments.length; ++i) {
            OperatorSet calldata operatorSet = deallocationParam.magnitudeAdjustments[i].operatorSet;
            uint64 decrementedMagnitude = deallocationParam.magnitudeAdjustments[i].magnitudeDiff;
            // check pending updates, if more than 3 pushed updates than revert
            require(checkPendingMagnitudeUpdates(operator, strategy, operatorSet), "Slasher._queueAllocation: too many pending updates");
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
            emit MagnitudeUpdated(operator, strategy, operatorSet, effectTimestamp, allocatedMagnitude - decrementedMagnitude);
            // add decremented magnitude back to nonslashableMagnitude
            nonslashableMagnitude += decrementedMagnitude;
        }

        // After updating magnitudes for each operatorSet, push/update a TotalAndNonslashableUpdate with updated nonslashableMagnitude
        _updateTotalAndNonslashableHistory({
            operator: operator,
            strategy: strategy,
            totalAndNonslashableUpdate: latestTotalUpdate,
            newTotalMagnitude: latestTotalUpdate.totalMagnitude,
            newNonslashableMagnitude: nonslashableMagnitude,
            newCumulativeAllocationSum: latestTotalUpdate.cumulativeAllocationSum,
            effectTimestamp: effectTimestamp
        });
        emit TotalAndNonslashableMagnitudeUpdated(
            operator,
            strategy,
            effectTimestamp,
            latestTotalUpdate.totalMagnitude,
            nonslashableMagnitude,
            latestTotalUpdate.cumulativeAllocationSum
        );
    }

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
        emit MagnitudeUpdated(operator, strategy, operatorSet, effectTimestamp, newMagnitude);
    }

    /// @notice push or update the latest TotalAndNonslashable update for the given operator, strategy
    /// depending on the current effectTimestamp
    function _updateTotalAndNonslashableHistory(
        address operator,
        IStrategy strategy,
        TotalAndNonslashableUpdate storage totalAndNonslashableUpdate,
        uint64 newTotalMagnitude,
        uint64 newNonslashableMagnitude,
        uint64 newCumulativeAllocationSum,
        uint32 effectTimestamp
    ) internal {
        if (totalAndNonslashableUpdate.timestamp == effectTimestamp) {
            totalAndNonslashableUpdate.totalMagnitude = newTotalMagnitude;
            totalAndNonslashableUpdate.nonslashableMagnitude = newNonslashableMagnitude;
        } else {
            _totalMagnitudeUpdates[operator][strategy].push(TotalAndNonslashableUpdate({
                timestamp: effectTimestamp,
                totalMagnitude: newTotalMagnitude,
                nonslashableMagnitude: newNonslashableMagnitude,
                cumulativeAllocationSum: newCumulativeAllocationSum
            }));
        }
        emit TotalAndNonslashableMagnitudeUpdated(
            operator,
            strategy,
            effectTimestamp,
            newTotalMagnitude,
            newNonslashableMagnitude,
            newCumulativeAllocationSum
        );
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
                nonslashableMagnitude: 0,
                cumulativeAllocationSum: 0
            }));
        }
        return _totalMagnitudeUpdates[operator][strategy][totalMagnitudeUpdatesLength - 1];
    }

    /*******************************************************************************
                                 VIEW FUNCTIONS
    *******************************************************************************/

    /// @notice Rate limit number of allocations/deallocation updates and return whether a new update could be pushed
    /// for a given (operator, strategy, and operatorSet)
    function checkPendingMagnitudeUpdates(
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

    /// @notice Rate limit number of allocations/deallocation updates and return whether a new update could be pushed
    /// for a given (operator, strategy, and operatorSet)
    function checkPendingTotalMagnitudeUpdates(
        address operator,
        IStrategy strategy
    ) public view returns (bool) {
        uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][strategy].length;
        if (totalMagnitudeUpdatesLength < MAX_PENDING_TOTAL_MAGNITUDE_UPDATES) {
            return true;
        }

        return _totalMagnitudeUpdates
            [operator]
            [strategy]
            [totalMagnitudeUpdatesLength - MAX_PENDING_TOTAL_MAGNITUDE_UPDATES].timestamp < block.timestamp;
    }

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

    function calculateReallocationDigestHash(
        address operator,
        MagnitudeAdjustmentsParam[] calldata adjustmentParams,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(REALLOCATION_TYPEHASH, operator, adjustmentParams, salt, expiry))
        );
    }

    function calculateMagnitudeDilutionDigestHash(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToAdd,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(MAGNITUDE_DILUTION_TYPEHASH, operator, strategies, nonslashableToAdd, salt, expiry))
        );
    }

    function calculateMagnitudeConcentrationDigestHash(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata nonslashableToDecrement,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        return _calculateDigestHash(
            keccak256(abi.encode(MAGNITUDE_CONCENTRATION_TYPEHASH, operator, strategies, nonslashableToDecrement, salt, expiry))
        );
    }

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() public view returns (bytes32) {
        return _calculateDomainSeparator();
    }

    /// @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
        }
    }

    function _calculateDigestHash(bytes32 structHash) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", _calculateDomainSeparator(), structHash));
    }
}
