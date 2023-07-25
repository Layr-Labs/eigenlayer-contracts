// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/ISlasher.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../interfaces/IServiceManager.sol";

/**
 * @title Abstract contract for implementing some of the Service Manager functionality that will be used by AVSs.
 * @author Layr Labs, Inc.
 * @notice This contract has mostly disabled functionality.
 */
abstract contract ServiceManagerBase is Initializable, OwnableUpgradeable, IServiceManager {

    IStrategyManager public immutable strategyManager;

    IDelegationManager public immutable delegationManager;

    ISlasher public immutable slasher;

    IAVSStateViewer public immutable avsStateViewer;

    IBLSRegistryCoordinatorWithIndices public immutable registryCoordinator;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "onlyRegistryCoordinator: not from registry coordinator");
        _;
    }

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegationMananger,
        ISlasher _slasher,
        IAVSStateViewer _avsStateViewer,
        IBLSRegistryCoordinatorWithIndices _registryCoordinator
    ) {
        strategyManager = _strategyManager;
        delegationManager = _delegationMananger;
        slasher = _slasher;
        avsStateViewer = _avsStateViewer;
        registryCoordinator = _registryCoordinator;
        _disableInitializers();
    }

    function _initialize(
        address initialOwner
    )
        public
        initializer
    {
        _transferOwnership(initialOwner);
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    function freezeOperator(address /*operator*/) external pure {
        revert("EigenDAServiceManager.freezeOperator: not implemented");
        // require(
        //     msg.sender == address(eigenDAChallenge)
        //         || msg.sender == address(eigenDABombVerifier),
        //     "EigenDAServiceManager.freezeOperator: Only challenge resolvers can slash operators"
        // );
        // slasher.freezeOperator(operator);
    }

    /**
     * @notice Called by the Registry in the event of a new registration, to forward a call to the Slasher
     * @param operator The operator whose stake is being updated
     * @param serveUntilBlock The block until which the stake accounted for in the first update is slashable by this middleware
     */ 
    function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) external onlyRegistryCoordinator {
        // slasher.recordFirstStakeUpdate(operator, serveUntilBlock);
    }

    /** 
     * @notice Called by the registryCoordinator, in order to forward a call to the Slasher, informing it of a stake update
     * @param operator The operator whose stake is being updated
     * @param updateBlock The block at which the update is being made
     * @param serveUntilBlock The block until which the stake withdrawn from the operator in this update is slashable by this middleware
     * @param prevElement The value of the previous element in the linked list of stake updates (generated offchain)
     */
    function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 prevElement) external onlyRegistryCoordinator {
        // slasher.recordStakeUpdate(operator, updateBlock, serveUntilBlock, prevElement);
    }

    /**
     * @notice Called by the registryCoordinator in the event of deregistration, to forward a call to the Slasher
     * @param operator The operator being deregistered
     * @param serveUntilBlock The block until which the stake delegated to the operator is slashable by this middleware
     */ 
    function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) external onlyRegistryCoordinator {
        // slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntilBlock);
    }

    // VIEW Functions

    /// @dev need to override function here since its defined in both these contracts
    function owner() public view override(OwnableUpgradeable, IServiceManager) returns (address) {
        return OwnableUpgradeable.owner();
    }
}