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

    /// EVENTS

    event OperatorSlashingCapBipsUpdated(address operator, IStrategy[] strategies, uint16[] slashingCapBips);

    event OperatorSetSlashingParametersUpdated(
        address operator, OperatorSet operatorSet, IStrategy[] strategies, uint16[] slashableBips, uint32 effectEpoch
    );

    /// EXTERNAL - STATE MODIFYING

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
        SignatureWithExpiry calldata allocatorSignature
    ) external returns (uint32 effectEpoch);

    /**
     * @notice Updates slashableBips of the provided strategies
     * for an operator set for a given operator
     *
     * @param operator the operator to change slashing parameters for
     * @param operatorSets the operator sets to change slashing parameters for
     * @param strategies the EigenLayer strategies to be update slashable bips for
     * @param slashableBips the maximum basis points of the operator's stake in
     * the given strategy can be slashed per epoch
     * @param allocatorSignature if non-empty is the signature of the allocator on
     * the modification. if empty, the msg.sender must be the allocator for the
     * operator
     *
     * @return effectEpoch the epoch in which the change to parameters will
     * take effect
     */
    function updateOperatorSetSlashingParameters(
        address operator,
        OperatorSet[] calldata operatorSets,
        IStrategy[][] calldata strategies,
        uint16[][] calldata slashableBips,
        SignatureWithExpiry calldata allocatorSignature
    ) external returns (uint32 effectEpoch);

    // /**
    //  * @notice A batch call of updateSlashingCapBips and updateOperatorSetSlashingParameters
    //  */
    // function updateSlashingCapBipsAndOperatorSetSlashingParameters(
    //     address operator,
    //     IStrategy[] calldata strategiesSlashingCaps,
    //     uint16[] calldata slashingCapBips,
    //     OperatorSet[] calldata operatorSets,
    //     IStrategy[][] calldata strategiesSlashableBips,
    //     uint16[][] calldata slashableBips,
    //     SignatureWithExpiry calldata allocatorSignature
    // ) external returns (uint32 effectEpoch);

    /// VIEW

    /**
     * @param operator the operator to get the slashable bips for
     * @param strategy the strategy to get the slashable bips for
     * @param epoch the epoch to get the slashable bips for for
     *
     * @return slashableCapBips the total slashable bips of the
     * given strategy allocated across all operator sets
     */
    function getSlashingCapBips(address operator, IStrategy strategy, uint32 epoch) external view returns (uint16);

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
