// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./SlasherStorage.sol";

import "../libraries/EpochUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";

/**
 * @notice Slasher
 */
contract Slasher is SlasherStorage {
    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager,
        IOperatorSetManager _operatorSetManager
    ) SlasherStorage(_strategyManager, _delegationManager, _operatorSetManager) {}

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
    ) external {
        require(
            0 < bipsToSlash && bipsToSlash < SlashingAccountingUtils.BIPS_FACTOR,
            "Slasher._slashRequestedBips: invalid bipsToSlash"
        );

        IOperatorSetManager.OperatorSet memory operatorSet = IOperatorSetManager.OperatorSet({
            avs: msg.sender,
            id: operatorSetId
        });
        uint32 epoch = EpochUtils.currentEpoch();

        // slash operator for each strategy
        for (uint256 i = 0; i < strategies.length; ++i) {
            IStrategy strategy = strategies[i];

            // calculate slashingRate based off allocated magnitude
            uint64 slashingRate = 
                uint64(bipsToSlash) * uint64(operatorSetManager.getSlashableBips(operator, operatorSet, strategy, epoch));
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

            emit OperatorSlashed(epoch, operator, strategy, operatorSet, bipsToSlash, slashingRate);
        }
    }

    /// VIEW

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
		IOperatorSetManager.OperatorSet calldata operatorSet,
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
