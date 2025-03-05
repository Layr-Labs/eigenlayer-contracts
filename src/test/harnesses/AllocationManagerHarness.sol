// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/core/AllocationManager.sol";

contract AllocationManagerHarness is AllocationManager {
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;

    constructor(
        IDelegationManager _delegation,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    )
        AllocationManager(_delegation, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY, "v9.9.9")
    {}

    function deallocationQueueAtIndex(address operator, IStrategy strategy, uint index) external view returns (bytes32) {
        return deallocationQueue[operator][strategy].at(index);
    }
}
