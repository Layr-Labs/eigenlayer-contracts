// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../munged/core/DelegationManager.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(IStrategyManager _strategyManager, ISlasher _slasher, IEigenPodManager _eigenPodManager)
        DelegationManager(_strategyManager, _slasher, _eigenPodManager) {}

    function get_operatorShares(address operator, IStrategy strategy) public view returns(uint256) {
        return operatorShares[operator][strategy];
    }
}