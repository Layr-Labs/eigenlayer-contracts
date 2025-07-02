// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/core/AllocationManager.sol";
import "forge-std/Test.sol";
import "../TestConstants.sol";

contract AllocationManagerHarness is AllocationManager {
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;

    constructor(
        IDelegationManager _delegation,
        IStrategy _eigenStrategy,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    )
        AllocationManager(
            _delegation,
            _eigenStrategy,
            _pauserRegistry,
            _permissionController,
            _DEALLOCATION_DELAY,
            _ALLOCATION_CONFIGURATION_DELAY,
            TestConstants.TEST_VERSION
        )
    {}

    function deallocationQueueAtIndex(address operator, IStrategy strategy, uint index) external view returns (bytes32) {
        return deallocationQueue[operator][strategy].at(index);
    }
}
