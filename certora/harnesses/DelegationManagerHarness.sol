// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/DelegationManager.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(IStrategyManager _strategyManager, ISlasher _slasher, IEigenPodManager _eigenPodManager)
        DelegationManager(_strategyManager, _slasher, _eigenPodManager) {}


}
