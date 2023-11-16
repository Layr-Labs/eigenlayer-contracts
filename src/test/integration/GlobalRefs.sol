// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";

contract GlobalRefs {
    EigenPodManager public eigenPodManager;
    DelegationManager public delegationManager;
    StrategyManager public strategyManager;

    constructor(EigenPodManager _eigenPodManager, DelegationManager _delegationManager, StrategyManager _strategyManager) {
        eigenPodManager = _eigenPodManager;
        delegationManager = _delegationManager;
        strategyManager = _strategyManager;
    }
}