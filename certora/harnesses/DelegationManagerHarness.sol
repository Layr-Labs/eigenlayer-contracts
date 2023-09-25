// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../munged/core/DelegationManager.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(IStrategyManager _strategyManager, ISlasher _slasher, IEigenPodManager _eigenPodManager)
        DelegationManager(_strategyManager, _slasher, _eigenPodManager) {}


    /// Harnessed functions
    function decreaseDelegatedShares(
        address staker,
        IStrategy strategy1,
        IStrategy strategy2,
        uint256 share1,
        uint256 share2
        ) external {
            IStrategy[] memory strategies = new IStrategy[](2);
            uint256[] memory shares = new uint256[](2);
            strategies[0] = strategy1;
            strategies[1] = strategy2;
            shares[0] = share1;
            shares[1] = share2;
            super.decreaseDelegatedShares(staker, strategies, shares);
    }

    function get_operatorShares(address operator, IStrategy strategy) public view returns(uint256) {
        return operatorShares[operator][strategy];
    }
}