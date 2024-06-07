// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/ISlasher.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @notice SlasherStorage
 */ 
abstract contract SlasherStorage is Initializable, OwnableUpgradeable, ISlasher, Pausable {
    IStrategyManager public strategyManager;

    IDelegationManager public delegation;

    IOperatorSetManager public operatorSetManager;

    /**
     * @notice Mapping: operator => strategy => share scalingFactor,
     * stored in the "SHARE_CONVERSION_SCALE", i.e. scalingFactor = 2 * SHARE_CONVERSION_SCALE indicates a scalingFactor of "2".
     * Note that a value of zero is treated as one, since this means that the operator has never been slashed
     */
    mapping(address => mapping(IStrategy => uint256)) internal _shareScalingFactor;

    /// @notice Mapping: operator => strategy => epochs in which the strategy was slashed for the operator
    // TODO: note that since default will be 0, we should probably make the "first epoch" actually be epoch 1 or something
    mapping(address => mapping(IStrategy => uint32[])) public slashedEpochHistory;

    /**
     * @notice Mapping: operator => strategy => epoch => scaling factor as a result of slashing *in that epoch*
     * @dev Note that this will be zero in the event of no slashing for the (operator, strategy) tuple in the given epoch.
     * You should use `shareScalingFactorAtEpoch` if you want the actual historical value of the share scaling factor in a given epoch.
     */
    mapping(address => mapping(IStrategy => mapping(uint32 => uint256))) public shareScalingFactorHistory;

    /**
     * @notice Mapping: Operator => Strategy => epoch => operator set hash => requested slashed bips
     * @dev Note that this is *independent* of the slashable bips that the operator has determined!
     *      The amount that the operator's delegated shares will actually get slashed (which goes into `pendingSlashingRate`) is linearly proportional
     *      to *both* this number *and* the slashable bips.
     */
    mapping(address => mapping(IStrategy => mapping(uint32 => mapping(bytes32 => uint32)))) public requestedSlashedBips;

    /**
     * @notice Mapping: Operator => Strategy => epoch => pending slashed amount, where pending slashed amount is the
     * amount that will be slashed when slashing is executed for the current epoch, assuming no existing requests are cancelled or nullified. summed over all AVSs
     * @dev Note that this is parts per (BIPS_FACTOR**2), i.e. parts per 1e8
     */
    mapping(address => mapping(IStrategy => mapping(uint32 => uint64))) public pendingSlashingRate;

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegationManager, IOperatorSetManager _operatorSetManager) {
        strategyManager = _strategyManager;
        delegation = _delegationManager;
        operatorSetManager = _operatorSetManager;
    }
}
