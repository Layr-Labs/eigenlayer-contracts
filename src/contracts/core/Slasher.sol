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
	) external {
        require(bipsToIncrease > 0, "Slasher.increaseRequestedBipsToSlash: bipsToIncrease must be positive");
        _modifyRequestedBipsToSlash(operator, operatorSetID, strategies, EpochUtils.currentEpoch(), bipsToIncrease);
    }

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
    ) external {
        uint32 currentEpoch = EpochUtils.currentEpoch();
        require(
            epoch == currentEpoch || epoch + 1 == currentEpoch,
            "Slasher.reduceRequestedBipsToSlash: can only reduce for current or previous epoch"
        );
        require(bipsToReduce < 0, "Slasher.reduceRequestedBipsToSlash: bipsToReduce must be negative");
        _modifyRequestedBipsToSlash(operator, operatorSetID, strategies, epoch, bipsToReduce);
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
            // make sure the slashing request is in order and overwrite the last executed slashing request ID
            SlashingRequest memory slashingRequest = slashingRequests[operator][strategy][epoch];
            require(
                slashingRequestIds[operator][strategy].lastExecutedSlashingRequestId + 1 == slashingRequest.id,
                "Slasher.executeSlashing: must execute slashings in order"
            );
            slashingRequestIds[operator][strategy].lastExecutedSlashingRequestId = slashingRequest.id;

            // truncate to BIPS_FACTOR_SQUARED if it exceeds it
            if (slashingRequest.slashingRate > SlashingAccountingUtils.BIPS_FACTOR_SQUARED) {
                slashingRequest.slashingRate = SlashingAccountingUtils.BIPS_FACTOR_SQUARED;
            } else if (slashingRequest.slashingRate == 0) {
                // if the slashing rate is 0 due to cancellation, skip the slashing
                continue;
            } else {
                uint256 scalingFactorBefore = shareScalingFactor(operator, strategy);
                uint256 scalingFactorAfter = SlashingAccountingUtils.findNewScalingFactor({
                    scalingFactorBefore: scalingFactorBefore,
                    rateToSlash: slashingRequest.slashingRate
                });
                // update storage to reflect the slashing
                slashedEpochHistory[operator][strategy].push(epoch);
                _shareScalingFactor[operator][strategy] = scalingFactorAfter;
                // TODO: note that this behavior risks off-by-one errors. it needs to be clearly defined precisely how the historical storage is supposed to work
                shareScalingFactorHistory[operator][strategy][epoch] = scalingFactorAfter;
            }

            emit SlashingExecuted(epoch, operator, strategy, slashingRequest.slashingRate);
        }
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
     * modifications to requested slashing rate by operatorSet
     *
     * @param operator the operator to get the pending slashing rate for
     * @param strategy the strategy to get the pending slashing rate for
     * @param operatorSet the operator set to get the pending slashing rate for
     *
     * @return the parts per hundred million that will be slashed for the given
     * operator, strategy, epoch, and operator set assuming no further
     * modifications to requested slashing rate by operatorSet
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
     * no more modifications to requested slashing rate for the operator.
     *
     * @param operator the operator to get the pending slashing rate for
     * @param strategy the strategy to get the pending slashing rate for
     * @param epoch the epoch to get the pending slashing rate for
     *
     * @return the parts per hundred million that will be slashed for the
     * given operator, strategy, and epoch, across all operator set assuming
     * no more modifications to requested slashing rate for the operator.
     */
    function getTotalPendingSlashingRate(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) external view returns (uint32) {
        uint64 totalSlashingRate = slashingRequests[operator][strategy][epoch].slashingRate;
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
            rateToSlash: slashingRequests[operator][strategy][EpochUtils.currentEpoch()].slashingRate
        });
        return pendingScalingFactor;
    }

    // TODO: this is a "naive" search since it brute-force backwards searches; we might technically want a binary search for completeness
    function shareScalingFactorAtEpoch(
        address operator,
        IStrategy strategy,
        uint32 epoch
    ) public view returns (uint256) {
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

    function _modifyRequestedBipsToSlash(
        address operator,
        bytes4 operatorSetID,
        IStrategy[] memory strategies,
        uint32 epoch,
        int32 bipsToModify
    ) internal {
        require(bipsToModify != 0, "Slasher._modifyRequestedBipsToSlash: cannot modify slashing by 0");

        IOperatorSetManager.OperatorSet memory operatorSet =
            IOperatorSetManager.OperatorSet({avs: msg.sender, id: operatorSetID});
        bytes32 operatorSetHash = _hashOperatorSet(operatorSet);
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

            SlashingRequest memory slashingRequest = slashingRequests[operator][strategy][epoch];
            // if this is the first request, increment the lastCreatedSlashingRequestId and set the id
            if (slashingRequest.id == 0) {
                uint32 requestId = slashingRequestIds[operator][strategy].lastCreatedSlashingRequestId + 1;
                slashingRequestIds[operator][strategy].lastCreatedSlashingRequestId = requestId;
                slashingRequest.id = requestId;
            }

            // pending slashing rate is modified by the bips to modify * the slashable bips
            // when divided by BIPS_FACTOR**2, this will give the actual proportion of the
            // stake is being modified. this is done for accuracy when slashable bips are small
            slashingRequest.slashingRate = uint64(
                int64(slashingRequest.slashingRate)
                    + int64(bipsToModify)
                        * int64(uint64(operatorSetManager.getSlashableBips(operator, operatorSet, strategy, epoch)))
            );
            slashingRequests[operator][strategy][epoch] = slashingRequest;
        }

        emit RequestedBipsToSlashModified(epoch, operator, operatorSet, strategies, bipsToModify);
    }

    function _hashOperatorSet(IOperatorSetManager.OperatorSet memory operatorSet) internal pure returns (bytes32) {
        IOperatorSetManager.OperatorSet[] memory operatorSetArray = new IOperatorSetManager.OperatorSet[](1);
        operatorSetArray[0] = operatorSet;
        return keccak256(abi.encode(operatorSetArray));
    }
}
