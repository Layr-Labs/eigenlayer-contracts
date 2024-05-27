// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "../interfaces/IOperatorSetManager.sol";
import "../libraries/EpochUtils.sol";
import "../libraries/SlashingAccountingUtils.sol";

contract OperatorSetManager is IOperatorSetManager {
    struct BipUpdate {
        uint32 epoch;
        uint16 bips;
    }

    struct TotalMagnitudeUpdate {
        uint32 epoch;
        uint64 totalMagnitude;
        uint64 totalAllocatedMagnitude;
    }

    struct MagnitudeUpdate {
        uint32 epoch;
        uint64 magnitude;
    }

    mapping(address => mapping(IStrategy => BipUpdate[])) private _slashingCapBipsUpdates;
    mapping(address => mapping(IStrategy => mapping(bytes32 => MagnitudeUpdate[]))) private _operatorSetMagnitudeUpdates;
    mapping(address => mapping(IStrategy => TotalMagnitudeUpdate[])) private _totalMagnitudeUpdates;

    /**
	 * @notice Updates slashableBips of the provided strategies
	 * for an operator set for a given operator
	 * 
	 * @param operator the operator to change slashing parameters for
	 * @param slashingMagnitudeParameters the new slashing magnitude parameters
	 * @param allocatorSignature if non-empty is the signature of the allocator on 
	 * the modification. if empty, the msg.sender must be the allocator for the 
	 * operator
	 * 
	 * @return effectEpoch the epoch in which the change to parameters will 
	 * take effect 
	 */
	function updateSlashingParameters(
		address operator,
		SlashingMagnitudeParameters[] calldata slashingMagnitudeParameters,
		SignatureWithExpiry calldata allocatorSignature
	) external returns(uint32 effectEpoch) {
        require(_isAllocatorFor(msg.sender, operator), "OperatorSetManager.updateOperatorSetSlashingParameters: Caller is not the allocator for the operator");
        // Change takes effect in 2 epochs
        effectEpoch = EpochUtils.currentEpochUint32() + 2; 

        for (uint256 i = 0; i < slashingMagnitudeParameters.length; i++) {
            require(slashingMagnitudeParameters[i].operatorSets.length == slashingMagnitudeParameters[i].slashableMagnitudes.length, "OperatorSetManager.updateOperatorSetSlashingParameters: operatorSets and slashableMagnitudes length mismatch");

            TotalMagnitudeUpdate memory totalMagnitudeUpdate;
            uint256 totalMagnitudeUpdatesLength = _totalMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy].length;
            if (totalMagnitudeUpdatesLength != 0) {
                totalMagnitudeUpdate = _totalMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][totalMagnitudeUpdatesLength - 1];
            }

            for (uint256 j = 0; j < slashingMagnitudeParameters[i].operatorSets.length; j++) {
                bytes32 operatorSetHash = _hashOperatorSet(slashingMagnitudeParameters[i].operatorSets[j]);
                
                // set the magnitude to the slashable magnitude
                MagnitudeUpdate memory operatorSetMagnitudeUpdate;
                uint256 operatorSetMagnitudeUpdatesLength = _operatorSetMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][operatorSetHash].length;
                if (operatorSetMagnitudeUpdatesLength != 0) {
                    operatorSetMagnitudeUpdate = _operatorSetMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][operatorSetHash][operatorSetMagnitudeUpdatesLength - 1];
                }

                // keep track of the lower bound on total magnitude
                totalMagnitudeUpdate.totalAllocatedMagnitude -= operatorSetMagnitudeUpdate.magnitude;
                totalMagnitudeUpdate.totalAllocatedMagnitude += slashingMagnitudeParameters[i].slashableMagnitudes[j];

                // update storage in place if last update is for effectEpoch
                if (operatorSetMagnitudeUpdate.epoch == effectEpoch) {
                    _operatorSetMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][operatorSetHash][operatorSetMagnitudeUpdatesLength - 1].magnitude = slashingMagnitudeParameters[i].slashableMagnitudes[j];
                } else {
                    // otherwise, add a new update
                    operatorSetMagnitudeUpdate = MagnitudeUpdate({epoch: effectEpoch, magnitude: slashingMagnitudeParameters[i].slashableMagnitudes[j]});
                    _operatorSetMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][operatorSetHash].push(operatorSetMagnitudeUpdate);
                }

                emit SlashableMagnitudeUpdated(operator, slashingMagnitudeParameters[i].strategy, slashingMagnitudeParameters[i].operatorSets[j], slashingMagnitudeParameters[i].slashableMagnitudes[j], effectEpoch);
            }

            require(totalMagnitudeUpdate.totalAllocatedMagnitude <= slashingMagnitudeParameters[i].totalMagnitude, "OperatorSetManager.updateOperatorSetSlashingParameters: totalAllocatedMagnitude exceeds totalMagnitude");
            totalMagnitudeUpdate.totalMagnitude = slashingMagnitudeParameters[i].totalMagnitude;
            totalMagnitudeUpdate.totalAllocatedMagnitude = totalMagnitudeUpdate.totalAllocatedMagnitude;

            // update storage in place if last update is for effectEpoch
            if (totalMagnitudeUpdate.epoch == effectEpoch) {
                _totalMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy][totalMagnitudeUpdatesLength - 1] = totalMagnitudeUpdate;
            } else {
                // otherwise, add a new update
                totalMagnitudeUpdate.epoch = effectEpoch;
                _totalMagnitudeUpdates[operator][slashingMagnitudeParameters[i].strategy].push(totalMagnitudeUpdate);
            }

            emit TotalMagnitudeUpdated(operator, slashingMagnitudeParameters[i].strategy, slashingMagnitudeParameters[i].totalMagnitude, effectEpoch);
        } 

        return effectEpoch;
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
                totalMagnitude = _totalMagnitudeUpdates[operator][strategy][i - 1].totalMagnitude;
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
