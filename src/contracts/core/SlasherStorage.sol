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
    
    // system contracts
    IStrategyManager public immutable strategyManager;
    IDelegationManager public immutable delegation;
    IOperatorSetManager public immutable operatorSetManager;

    struct SlashingRequest {
        uint32 id; // an incrementing ID for each slashing request
        uint64 slashingRate; // This is parts per (BIPS_FACTOR**2), i.e. parts per 1e8, pphm = parts per hundred million, to slash upon execution
        uint64 scalingFactor; // the scaling factor to apply to the operator's shares. this is only set upon execution
    }

    struct SlashingRequestIds {
        uint32 lastCreatedSlashingRequestId; // the last slashing request ID that was created
        uint32 lastExecutedSlashingRequestId; // the last slashing request ID that was executed
    }

    /// @notice Mapping: Operator => Strategy => SlashingRequestIds
    mapping(address => mapping(IStrategy => SlashingRequestIds)) public slashingRequestIds;

    /// @notice Mapping: Operator => Strategy => epoch => SlashingRequest
    mapping(address => mapping(IStrategy => mapping(uint32 => SlashingRequest))) public slashingRequests;

    /**
     * @notice Mapping: operator => strategy => share scalingFactor,
     * stored in the "SHARE_CONVERSION_SCALE", i.e. scalingFactor = 2 * SHARE_CONVERSION_SCALE indicates a scalingFactor of "2".
     * Note that a value of zero is treated as one, since this means that the operator has never been slashed
     */
    mapping(address => mapping(IStrategy => uint64)) internal _shareScalingFactor;

    /// @notice Mapping: operator => strategy => epochs in which the strategy was slashed for the operator
    // TODO: note that since default will be 0, we should probably make the "first epoch" actually be epoch 1 or something
    mapping(address => mapping(IStrategy => uint32[])) public slashedEpochHistory;

    /**
     * @notice Mapping: Operator => Strategy => epoch => operator set hash => requested slashed bips
     * @dev Note that this is *independent* of the slashable bips that the operator has determined!
     *      The amount that the operator's delegated shares will actually get slashed (which goes into `pendingSlashingRate`) is linearly proportional
     *      to *both* this number *and* the slashable bips.
     */
    mapping(address => mapping(IStrategy => mapping(uint32 => mapping(bytes32 => uint32)))) public requestedSlashedBips;

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager,
        IOperatorSetManager _operatorSetManager
    ) {
        strategyManager = _strategyManager;
        delegation = _delegationManager;
        operatorSetManager = _operatorSetManager;
    }
}
