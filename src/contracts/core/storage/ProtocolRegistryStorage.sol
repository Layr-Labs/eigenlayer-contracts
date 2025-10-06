// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../interfaces/IAllocationManager.sol";
// import "../../interfaces/IAVSDirectory.sol";
// import "../../interfaces/IDelegationManager.sol";
import "../../interfaces/IProtocolRegistry.sol";
// import "../../interfaces/IReleaseManager.sol";
// import "../../interfaces/IRewardsCoordinator.sol";
// import "../../interfaces/IStrategyManager.sol";

abstract contract ProtocolRegistryStorage is IProtocolRegistry, IAllocationManagerView {
    // Immutables

    IAllocationManager public immutable allocationManager;

    // Mutatables

    // Construction

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
