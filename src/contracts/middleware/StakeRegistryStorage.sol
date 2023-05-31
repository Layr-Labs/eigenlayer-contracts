// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IStakeRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "./VoteWeigherBase.sol";

/**
 * @title Storage variables for the `StakeRegistry` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract StakeRegistryStorage is VoteWeigherBase, IStakeRegistry {
    /// @notice the coordinator contract that this registry is associated with
    IRegistryCoordinator public registryCoordinator;

    // TODO: set these on initialization
    /// @notice In order to register, an operator must have at least `minimumStakeFirstQuorum` or `minimumStakeSecondQuorum`, as
    /// evaluated by this contract's 'VoteWeigher' logic.
    uint96[256] public minimumStakeForQuorum;

    /// @notice array of the history of the total stakes for each quorum -- marked as internal since getTotalStakeFromIndex is a getter for this
    OperatorStakeUpdate[][256] internal totalStakeHistory;

    /// @notice mapping from operator's operatorId to the history of their stake updates
    mapping(bytes32 => mapping(uint8 => OperatorStakeUpdate[])) public operatorIdToStakeHistory;

    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) VoteWeigherBase(_strategyManager, _serviceManager)
    // solhint-disable-next-line no-empty-blocks
    {
    }

    // storage gap
    uint256[50] private __GAP;
}