// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/StrategyManager.sol";

contract StrategyManagerHarness is StrategyManager {
    constructor(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher)
        StrategyManager(_delegation, _eigenPodManager, _slasher)
        {}

    function strategy_is_in_stakers_array(address staker, IStrategy strategy) public view returns (bool) {
        uint256 length = stakerStrategyList[staker].length;
        for (uint256 i = 0; i < length; ++i) {
            if (stakerStrategyList[staker][i] == strategy) {
                return true;
            }
        }
        return false;
    }

    function num_times_strategy_is_in_stakers_array(address staker, IStrategy strategy) public view returns (uint256) {
        uint256 length = stakerStrategyList[staker].length;
        uint256 res = 0;
        for (uint256 i = 0; i < length; ++i) {
            if (stakerStrategyList[staker][i] == strategy) {
                res += 1;
            }
        }
        return res;
    }

    // checks that stakerStrategyList[staker] contains no duplicates and that all strategies in array have nonzero shares

    function totalShares(address strategy) public view returns (uint256) {
        return IStrategy(strategy).totalShares();
    }

}
