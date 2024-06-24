// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/core/RewardsCoordinator.sol";

contract RewardsCoordinatorHarness is RewardsCoordinator {
    constructor(
        IDelegationManager _delegationManager,
        IStrategyManager _strategyManager,
        uint32 _CALCULATION_INTERVAL_SECONDS,
        uint32 _MAX_PAYMENT_DURATION,
        uint32 _MAX_RETROACTIVE_LENGTH,
        uint32 _MAX_FUTURE_LENGTH,
        uint32 _GENESIS_PAYMENT_TIMESTAMP
    ) RewardsCoordinator(
        _delegationManager,
        _strategyManager,
        _CALCULATION_INTERVAL_SECONDS,
        _MAX_PAYMENT_DURATION,
        _MAX_RETROACTIVE_LENGTH,
        _MAX_FUTURE_LENGTH,
        _GENESIS_PAYMENT_TIMESTAMP
    ) {}
}