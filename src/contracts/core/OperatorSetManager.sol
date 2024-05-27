// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../interfaces/IOperatorSetManager.sol";
import "../libraries/EpochUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";

abstract contract OperatorSetManager is IOperatorSetManager {
    struct BipUpdate {
        uint32 epoch;
        uint16 bips;
    }

    struct MagnitudeUpdate {
        uint32 epoch;
        uint192 magnitude;
    }

    mapping(address => mapping(IStrategy => BipUpdate[])) private _slashingCapBipsUpdates;
    mapping(address => mapping(IStrategy => mapping(bytes32 => MagnitudeUpdate[]))) private _operatorSetMagnitudeUpdates;
    mapping(address => mapping(IStrategy => MagnitudeUpdate[])) private _totalMagnitudeUpdates;


    /**
	 * @notice Updates the slashingCapBips for a given operator and the provided strategies
	 *
	 * @param operator the operator to change slashingCapBips for
	 * @param strategies the strategies to set slashingCapBips for
	 * @param slashingCapBips the new slashingCapBips to set for each of the strategies
	 * @param allocatorSignature if non-empty is the signature of the allocator on 
	 * the modification. if empty, the msg.sender is must be the operator's allocator
	 * 
	 * @return effectEpoch the epoch in which the change will take effect
	 */
    function updateSlashingCapBips(
        address operator,
        IStrategy[] calldata strategies,
        uint16[] calldata slashingCapBips,
        SignatureWithExpiry calldata allocatorSignature // TODO: Verify sig!
    ) public override returns (uint32 effectEpoch) {
        require(
            _isAllocatorFor(msg.sender, operator),
            "OperatorSetManager.updateSlashingCapBips: Caller is not the allocator for the operator"
        );
        require(
            strategies.length == slashingCapBips.length,
            "OperatorSetManager.updateSlashingCapBips: Strategies and slashingCapBips length mismatch"
        );

        effectEpoch = EpochUtils.currentEpochUint32() + 2; // Change take effect in 2 epochs

        for (uint256 i = 0; i < strategies.length; i++) {
            require(slashingCapBips[i] <= SlashingAccountingUtils.BIPS_FACTOR, "OperatorSetManager.updateSlashingCapBips: bips is invalid");

            BipUpdate[] storage slashingCapBipsUpdates = _slashingCapBipsUpdates[operator][strategies[i]];
            uint256 slashingCapBipsUpdatesLength = slashingCapBipsUpdates.length;

            // Overwrite the last update if it's in the same effect epoch
            if (slashingCapBipsUpdatesLength > 0 && slashingCapBipsUpdates[slashingCapBipsUpdatesLength - 1].epoch == effectEpoch) {
                slashingCapBipsUpdates[slashingCapBipsUpdatesLength - 1].bips = slashingCapBips[i];
            } else {
                slashingCapBipsUpdates.push(BipUpdate({epoch: effectEpoch, bips: slashingCapBips[i]}));
            }
        }

        emit OperatorSlashingCapBipsUpdated(operator, strategies, slashingCapBips);
    }

    function updateOperatorSetSlashingParameters(
        address operator,
        OperatorSet[] calldata operatorSets,
        IStrategy[][] calldata strategies,
        uint16[][] calldata slashableBips,
        SignatureWithExpiry calldata allocatorSignature // TODO: Verify sig!
    ) external override returns (uint32 effectEpoch) {
        require(_isAllocatorFor(msg.sender, operator), "OperatorSetManager.updateOperatorSetSlashingParameters: Caller is not the allocator for the operator");
        require(operatorSets.length == strategies.length, "OperatorSetManager.updateOperatorSetSlashingParameters: OperatorSets and strategies length mismatch");
        require(operatorSets.length == slashableBips.length, "OperatorSetManager.updateOperatorSetSlashingParameters: OperatorSets and slashableBips length mismatch");

        effectEpoch = EpochUtils.currentEpochUint32() + 2; // Change takes effect in 2 epochs

        for (uint256 i = 0; i < operatorSets.length; i++) {
            bytes32 operatorSetHash = _hashOperatorSet(operatorSets[i]);
            require(strategies[i].length == slashableBips[i].length, "OperatorSetManager.updateOperatorSetSlashingParameters: Strategies and slashableBips length mismatch");

            for (uint256 j = 0; j < strategies[i].length; j++) {
                require(j == 0 || uint160(address(strategies[i][j])) > uint160(address(strategies[i][j - 1])), "OperatorSetManager.updateOperatorSetSlashingParameters: not ascending order of strategies");
                // TODO: can only allocate 9999 bips to the operator set
                require(slashableBips[i][j] <= SlashingAccountingUtils.BIPS_FACTOR - 1, "OperatorSetManager.updateOperatorSetSlashingParameters: bips is invalid");

                MagnitudeUpdate memory totalMagnitudeUpdate;
                uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][strategies[i][j]].length;
                if (totalMagnitudeUpdatesLength == 0) {
                    // if this is the first magnitude update,
                    // set the total magnitude to 10000 
                    _totalMagnitudeUpdates[operator][strategies[i][j]].push(MagnitudeUpdate({epoch: effectEpoch, magnitude: 10000}));
                    // and the operator set magnitude to the slashable bips
                    _operatorSetMagnitudeUpdates[operator][strategies[i][j]][operatorSetHash].push(MagnitudeUpdate({epoch: effectEpoch, magnitude: slashableBips[i][j]}));
                    // continue to the next strategy
                    continue;
                } else {
                    totalMagnitudeUpdate = _totalMagnitudeUpdates[operator][strategies[i][j]][totalMagnitudeUpdatesLength - 1];
                }
                
                MagnitudeUpdate memory operatorSetMagnitudeUpdate;
                uint256 operatorSetMagnitudeUpdatesLength = _operatorSetMagnitudeUpdates[operator][strategies[i][j]][operatorSetHash].length;
                if (operatorSetMagnitudeUpdatesLength > 0) {
                    operatorSetMagnitudeUpdate = _operatorSetMagnitudeUpdates[operator][strategies[i][j]][operatorSetHash][operatorSetMagnitudeUpdatesLength - 1];
                }

                // check whether we can just decrease the operator set magnitude
                uint192 newOperatorSetMagnitude = totalMagnitudeUpdate.magnitude * slashableBips[i][j] / uint32(SlashingAccountingUtils.BIPS_FACTOR);
                if (newOperatorSetMagnitude > operatorSetMagnitudeUpdate.magnitude) {
                    // if the new operator set magnitude is greater than the current one,
                    // we need to update the total magnitude

                    // nosm = newOperatorSetMagnitude
                    // tm = totalMagnitudeUpdate.magnitude
                    // osm = operatorSetMagnitudeUpdate.magnitude
                    // sb = slashableBips[i][j]
                    // we are trying to solve:
                    // nosm / (tm - osm + nosm) = sb / BIPS_FACTOR
                    // nosm = (tm - osm + nosm) * sb / BIPS_FACTOR
                    // BIPS_FACTOR * nosm = (tm - osm + nosm) * sb
                    // (BIPS_FACTOR - sb) * nosm = (tm - osm) * sb
                    // nosm = (tm - osm) * sb / (BIPS_FACTOR - sb)
                    // this is why sb cannot be BIPS_FACTOR, lest we divide by zero
                    newOperatorSetMagnitude = (totalMagnitudeUpdate.magnitude - operatorSetMagnitudeUpdate.magnitude) * slashableBips[i][j] / (uint32(SlashingAccountingUtils.BIPS_FACTOR) - slashableBips[i][j]);
                    
                    uint192 newTotalMagnitude = totalMagnitudeUpdate.magnitude - operatorSetMagnitudeUpdate.magnitude + newOperatorSetMagnitude;
                    // update storage in place if last update is for effectEpoch
                    if (totalMagnitudeUpdate.epoch == effectEpoch) {
                        _totalMagnitudeUpdates[operator][strategies[i][j]][totalMagnitudeUpdatesLength - 1].magnitude = newTotalMagnitude;
                    } else {
                        // otherwise, add a new update
                        totalMagnitudeUpdate = MagnitudeUpdate({epoch: effectEpoch, magnitude: newTotalMagnitude});
                        _totalMagnitudeUpdates[operator][strategies[i][j]].push(totalMagnitudeUpdate);
                    }
                }

                // update storage in place if last update is for effectEpoch
                if (operatorSetMagnitudeUpdate.epoch == effectEpoch) {
                    _operatorSetMagnitudeUpdates[operator][strategies[i][j]][operatorSetHash][operatorSetMagnitudeUpdatesLength - 1].magnitude = newOperatorSetMagnitude;
                } else {
                    // otherwise, add a new update
                    operatorSetMagnitudeUpdate = MagnitudeUpdate({epoch: effectEpoch, magnitude: newOperatorSetMagnitude});
                    _operatorSetMagnitudeUpdates[operator][strategies[i][j]][operatorSetHash].push(operatorSetMagnitudeUpdate);
                }
            }

            emit OperatorSetSlashingParametersUpdated(operator, operatorSets[i], strategies[i], slashableBips[i], effectEpoch);
        }
    }

    /// VIEW

    /**
	 * @param operator the operator to get the slashable bips for
	 * @param strategy the strategy to get the slashable bips for 
	 * @param epoch the epoch to get the slashable bips for for
	 * 
	 * @return slashableCapBips the total slashable bips of the 
	 * given strategy allocated across all operator sets  
	 */
    function getSlashingCapBips(address operator, IStrategy strategy, uint32 epoch) external view returns (uint16) {
        require(epoch <= EpochUtils.currentEpochUint32() + 2, "Epoch is more than 2 epochs in the future");

        BipUpdate[] storage slashingCapBipsUpdates = _slashingCapBipsUpdates[operator][strategy];
        uint16 latestBips = 0;

        // iterate from the latest update to the oldest
        for (uint256 i = slashingCapBipsUpdates.length; i > 0; i--) {
            if (slashingCapBipsUpdates[i - 1].epoch <= epoch) {
                latestBips = slashingCapBipsUpdates[i - 1].bips;
                break;
            }
        }

        return latestBips;
    }

    /**
     * @param operator the operator to get the slashable bips for
     * @param operatorSet the operator set to get the slashable bips for
     * @param strategy the strategy to get the slashable bips for
     * @param epoch the epoch to get the slashable bips for for
     *
     * @return slashableBips the slashable bips of the given strategy owned by
     * the given OperatorSet for the given operator and epoch
     */
    function getSlashableBips(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 epoch
    ) external view returns (uint16 slashableBips) {
        require(epoch <= EpochUtils.currentEpochUint32() + 2, "Epoch is more than 2 epochs in the future");

        bytes32 operatorSetHash = _hashOperatorSet(operatorSet);
        MagnitudeUpdate[] storage operatorSetMagnitudeUpdates = _operatorSetMagnitudeUpdates[operator][strategy][operatorSetHash];
        uint192 operatorSetMagnitude = 0;

        // iterate from the latest operator set magnitude update to the oldest
        for (uint256 i = operatorSetMagnitudeUpdates.length; i > 0; i--) {
            if (operatorSetMagnitudeUpdates[i - 1].epoch <= epoch) {
                operatorSetMagnitude = operatorSetMagnitudeUpdates[i - 1].magnitude;
                break;
            }
        }

        uint192 totalMagnitude = 0;
        // iterate through total magnitude updates to get the total magnitude
        for (uint256 i = _totalMagnitudeUpdates[operator][strategy].length; i > 0; i--) {
            if (_totalMagnitudeUpdates[operator][strategy][i - 1].epoch <= epoch) {
                totalMagnitude = _totalMagnitudeUpdates[operator][strategy][i - 1].magnitude;
                break;
            }
        }

        return uint16(operatorSetMagnitude * SlashingAccountingUtils.BIPS_FACTOR / totalMagnitude);
    }

    function _isAllocatorFor(address sender, address operator) internal view returns (bool) {
        // Implement allocator check logic
        return true; // Stubbed for now
    }

    function _hashOperatorSet(
        OperatorSet memory operatorSet
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(operatorSet));
    }
}
