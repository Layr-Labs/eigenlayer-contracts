// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/multichain/CrossChainRegistry.sol";

/**
 * @title CrossChainRegistryWrapper
 * @notice Exposes internal checkCanCall logic as an external function
 */
contract CrossChainRegistryHarness is CrossChainRegistry {


    /**
     * @dev Constructor that passes required parameters to the CrossChainRegistry base constructor
     * @param _allocationManager Address of the allocation manager contract
     * @param _keyRegistrar Address of the key registrar contract
     * @param _permissionController Address of the permission controller contract
     * @param _pauserRegistry Address of the pauser registry contract
     * @param _version Semantic version string
     */
    constructor(
        IAllocationManager _allocationManager,
        IKeyRegistrar _keyRegistrar,
        IPermissionController _permissionController,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        CrossChainRegistry(
            _allocationManager,
            _keyRegistrar,
            _permissionController,
            _pauserRegistry,
            _version
        )
    {}


    /**
     * @notice External function to check whether a given address can call privileged functions
     * @param caller The address to check permissions for
     */
    function canCall(address caller, address sender, uint32 selector) external returns (bool) {

        return permissionController.canCall(caller, sender, address(this), bytes4(selector));

    }
}   