// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./ISignatureUtils.sol";
import "./IStrategy.sol";

interface IOperatorSetManager is ISignatureUtils {
    /// STRUCTS

    struct OperatorSet {
        address avs;
        bytes4 id;
    }

    /**
     * @notice Parameters for updating slashing magnitudes for an operator
     * @param strategy the strategy to update slashable magnitudes for
     * @param totalMagnitude the total magnitude allocated to all operator sets
     * @param operatorSets the operator sets to change slashing parameters for
     * @param slashableMagnitudes the proportional parts of totalMagnitude that the operator set is getting as part of the strategy
     */
    struct SlashingMagnitudeParameters {
        IStrategy strategy;
        uint64 totalMagnitude;
        OperatorSet[] operatorSets;
        uint64[] slashableMagnitudes;
    }

    /// EVENTS

    event SlashableMagnitudeUpdated(
        address operator, IStrategy strategy, OperatorSet operatorSet, uint64 slashableMagnitude, uint32 effectEpoch
    );

    event TotalMagnitudeUpdated(
        address operator, IStrategy strategy, uint64 totalMagnitude, uint32 effectEpoch
    );

    /// EXTERNAL - STATE MODIFYING

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
	) external returns(uint32 effectEpoch);

    /// VIEW

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
    ) external returns (uint16 slashableBips);
}
