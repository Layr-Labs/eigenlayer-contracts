// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./BLSSignatureChecker.sol";
import "../permissions/Pausable.sol";

import "../interfaces/IDelegationManager.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IStrategyManager.sol";

/**
 * @title Base implementation of `IServiceManager` interface, designed to be inherited from by more complex ServiceManagers.
 * @author Layr Labs, Inc.
 * @notice This contract is used for:
 * - proxying calls to the Slasher contract
 * - implementing the two most important functionalities of a ServiceManager:
 *  - freezing operators as the result of various "challenges"
 *  - defining the latestServeUntilBlock which is used by the Slasher to determine whether a withdrawal can be completed
 */
abstract contract ServiceManagerBase is
    IServiceManager,
    Initializable,
    OwnableUpgradeable,
    BLSSignatureChecker,
    Pausable
{
    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev This function should contain slashing logic, to make sure operators are not needlessly being slashed
    //       hence it is marked as virtual and must be implemented in each avs' respective service manager contract
    function freezeOperator(address operatorAddr) external virtual;

    /// @notice Returns the block until which operators must serve.
    /// @dev The block until which the stake accounted for in stake updates is slashable by this middleware
    function latestServeUntilBlock() public view virtual returns (uint32);

    ISlasher public immutable slasher;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyRegistryCoordinator() {
        require(
            msg.sender == address(registryCoordinator),
            "onlyRegistryCoordinator: not from registry coordinator"
        );
        _;
    }

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    /// or by StakeRegistry
    modifier onlyRegistryCoordinatorOrStakeRegistry() {
        require(
            (msg.sender == address(registryCoordinator)) ||
                (msg.sender ==
                    address(
                        IBLSRegistryCoordinatorWithIndices(
                            address(registryCoordinator)
                        ).stakeRegistry()
                    )),
            "onlyRegistryCoordinatorOrStakeRegistry: not from registry coordinator or stake registry"
        );
        _;
    }

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
        ISlasher _slasher
    ) BLSSignatureChecker(_registryCoordinator) {
        slasher = _slasher;
        _disableInitializers();
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
    }

    // PROXY CALLS TO EQUIVALENT SLASHER FUNCTIONS

    /**
     * @notice Called by the Registry in the event of a new registration, to forward a call to the Slasher
     * @param operator The operator whose stake is being updated
     */
    function recordFirstStakeUpdate(
        address operator,
        uint32 serveUntilBlock
    ) external virtual onlyRegistryCoordinator {
        slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
    }

    /**
     * @notice Called by the registryCoordinator, in order to forward a call to the Slasher, informing it of a stake update
     * @param operator The operator whose stake is being updated
     * @param updateBlock The block at which the update is being made
     * @param prevElement The value of the previous element in the linked list of stake updates (generated offchain)
     */
    function recordStakeUpdate(
        address operator,
        uint32 updateBlock,
        uint32 serveUntilBlock,
        uint256 prevElement
    ) external virtual onlyRegistryCoordinatorOrStakeRegistry {
        slasher.recordStakeUpdate(
            operator,
            updateBlock,
            serveUntilBlock,
            prevElement
        );
    }

    /**
     * @notice Called by the registryCoordinator in the event of deregistration, to forward a call to the Slasher
     * @param operator The operator being deregistered
     */
    function recordLastStakeUpdateAndRevokeSlashingAbility(
        address operator,
        uint32 serveUntilBlock
    ) external virtual onlyRegistryCoordinator {
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(
            operator,
            serveUntilBlock
        );
    }

    // VIEW FUNCTIONS

    /// @dev need to override function here since its defined in both these contracts
    function owner()
        public
        view
        override(OwnableUpgradeable, IServiceManager)
        returns (address)
    {
        return OwnableUpgradeable.owner();
    }
}
