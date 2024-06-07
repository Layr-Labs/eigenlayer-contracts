// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./SlasherStorage.sol";

import "../libraries/EpochUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";

/**
 * @notice Slasher
 */ 
contract Slasher is SlasherStorage {

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegationManager, IOperatorSetManager _operatorSetManager) 
        SlasherStorage(_strategyManager, _delegationManager, _operatorSetManager) {}

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
     *     and positive when the AVS wants to increase the amount of slashing
	 */
    function modifyRequestedBipsToSlash(
        address operator,
        bytes4 operatorSetID,
        IStrategy[] memory strategies,
        int32 bipsToModify
    ) external {
        require(bipsToModify != 0, "Slasher._modifyRequestedBipsToSlash: cannot modify slashing by 0");

        IOperatorSetManager.OperatorSet memory operatorSet = IOperatorSetManager.OperatorSet({
            avs: msg.sender,
            id: operatorSetID
        });
        bytes32 operatorSetHash = _hashOperatorSet(operatorSet);
        uint32 epoch = EpochUtils.currentEpoch();
        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy strategy = strategies[i];
            uint32 requestedSlashedBipsBefore = requestedSlashedBips[operator][strategy][epoch][operatorSetHash];
            int32 requestedSlashedBipsAfter = int32(requestedSlashedBipsBefore) + bipsToModify;
            // don't allow underflow
            if (requestedSlashedBipsAfter < 0) {
                bipsToModify = -int32(requestedSlashedBipsBefore);
                requestedSlashedBipsAfter = 0;
            }
            requestedSlashedBips[operator][strategy][epoch][operatorSetHash] = uint32(requestedSlashedBipsAfter);

            // pending slashing rate is modified by the bips to modify * the slashable bips
            // when divided by BIPS_FACTOR**2, this will give the actual proportion of the 
            // stake is being modified. this is done for accuracy when slashable bips are small
            pendingSlashingRate[operator][strategy][epoch] = uint64(
                int64(pendingSlashingRate[operator][strategy][epoch]) 
                    + int64(bipsToModify) * int64(uint64(operatorSetManager.getSlashableBips(operator, operatorSet, strategy, epoch)))
            );
        }
        // TODO: Events
    }

    /**
	 * @notice Permissionlessly called to execute slashing of a given list of 
	 * strategies for a given operator, for the latest unslashed epoch
	 *
	 * @param operator the operator to slash
	 * @param strategies the list of strategies to execute slashing for
	 * @param epoch the epoch in which the slashing requests to execute were made
	 */
    function executeSlashing(address operator, IStrategy[] memory strategies, uint32 epoch) external {
        // TODO: decide if this needs a stonger condition. e.g. the epoch must be further in the past
        require(epoch < EpochUtils.currentEpoch(), "Slasher.executeSlashing: must slash for a previous epoch");
        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy strategy = strategies[i];
            uint64 rateToSlash = pendingSlashingRate[operator][strategy][epoch];
            pendingSlashingRate[operator][strategy][epoch] = 0;
            // truncate to BIPS_FACTOR_SQUARED if it exceeds it
            if(rateToSlash > SlashingAccountingUtils.BIPS_FACTOR_SQUARED) {
                rateToSlash = SlashingAccountingUtils.BIPS_FACTOR_SQUARED;
            }
            // TODO: again note that we need something like the first epoch being epoch 1 here, to allow actually slashing in the first epoch
            require(epoch > lastSlashed(operator, strategy), "Slasher._slashShares: slashing must occur in strictly ascending epoch order");
            uint256 scalingFactorBefore = shareScalingFactor(operator, strategy);
            uint256 scalingFactorAfter = SlashingAccountingUtils.findNewScalingFactor({
                scalingFactorBefore: scalingFactorBefore,
                rateToSlash: rateToSlash
            });
            // update storage to reflect the slashing
            slashedEpochHistory[operator][strategy].push(epoch);
            _shareScalingFactor[operator][strategy] = scalingFactorAfter;
            // TODO: note that this behavior risks off-by-one errors. it needs to be clearly defined precisely how the historical storage is supposed to work
            shareScalingFactorHistory[operator][strategy][epoch] = scalingFactorAfter;
        }
        // TODO: events!
    }

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
	) public view returns (uint32) {
        return requestedSlashedBips[operator][strategy][epoch][_hashOperatorSet(operatorSet)] 
            * operatorSetManager.getSlashableBips(operator, operatorSet, strategy, epoch);
    }
	
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
	) external view returns (uint32) {
        uint32 requestedSlashingRate = getRequestedSlashingRate(operator, strategy, operatorSet, epoch);
        if (requestedSlashingRate >= SlashingAccountingUtils.BIPS_FACTOR_SQUARED) {
            requestedSlashingRate = uint32(SlashingAccountingUtils.BIPS_FACTOR_SQUARED);
        }
        return requestedSlashingRate;
    }
	
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
	) external view returns (uint32) {
        uint64 totalSlashingRate = pendingSlashingRate[operator][strategy][epoch];
        if (totalSlashingRate >= SlashingAccountingUtils.BIPS_FACTOR_SQUARED) {
            totalSlashingRate = SlashingAccountingUtils.BIPS_FACTOR_SQUARED;
        }
        return uint32(totalSlashingRate);
    }

    function shareScalingFactor(address operator, IStrategy strategy) public view returns (uint256) {
        uint256 scalingFactor = _shareScalingFactor[operator][strategy];
        if (scalingFactor == 0) {
            return SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        } else {
            return scalingFactor;
        }
    }

    // includes pending but unexecuted slashings
    function pendingShareScalingFactor(address operator, IStrategy strategy) public view returns (uint256) {
        uint256 pendingScalingFactor = SlashingAccountingUtils.findNewScalingFactor({
            scalingFactorBefore: _shareScalingFactor[operator][strategy],
            // TODO: this particular lookup could potentially get more complex? e.g. also including the trailing epoch's pending amounts
            rateToSlash: pendingSlashingRate[operator][strategy][EpochUtils.currentEpoch()]
        });
        return pendingScalingFactor;
    }

    // TODO: this is a "naive" search since it brute-force backwards searches; we might technically want a binary search for completeness
    function shareScalingFactorAtEpoch(address operator, IStrategy strategy, uint32 epoch) public view returns (uint256) {
        uint256 scalingFactor = SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        // TODO: note the edge case of 0th epoch; need to make sure it's clear how it should be handled
        for (uint256 i = slashedEpochHistory[operator][strategy].length; i > 0; --i) {
            if (slashedEpochHistory[operator][strategy][i] <= epoch) {
                uint32 correctEpochForLookup = slashedEpochHistory[operator][strategy][i];
                scalingFactor = shareScalingFactorHistory[operator][strategy][correctEpochForLookup];
                break;
            }
        }
        return scalingFactor;
    }

    // @notice Returns the epoch in which the operator was last slashed
    function lastSlashed(address operator, IStrategy strategy) public view returns (uint32) {
        uint256 slashedEpochHistoryLength = slashedEpochHistory[operator][strategy].length;
        if (slashedEpochHistoryLength == 0) {
            // TODO: consider if a different return value is more appropriate for this special case
            return 0;
        } else {
            return slashedEpochHistory[operator][strategy][slashedEpochHistoryLength - 1];
        }
    }

    function _hashOperatorSet(
        IOperatorSetManager.OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        IOperatorSetManager.OperatorSet[] memory operatorSetArray = new IOperatorSetManager.OperatorSet[](1);
        operatorSetArray[0] = operatorSet;
        return keccak256(abi.encode(operatorSetArray));
    }
}
