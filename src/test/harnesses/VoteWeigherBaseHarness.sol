// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/VoteWeigherBase.sol";

// wrapper around the StakeRegistry contract that exposes the internal functions for unit testing.
contract VoteWeigherBaseHarness is VoteWeigherBase {
    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) VoteWeigherBase(_strategyManager, _serviceManager) {}

    function getMaxWeighingFunctionLength() public pure returns (uint8) {
        return MAX_WEIGHING_FUNCTION_LENGTH;
    }
}