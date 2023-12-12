// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

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
    function array_exhibits_properties(address staker) public view returns (bool) {
        uint256 length = stakerStrategyList[staker].length;
        uint256 res = 0;
        // loop for each strategy in array
        for (uint256 i = 0; i < length; ++i) {
            IStrategy strategy = stakerStrategyList[staker][i];
            // check that staker's shares in strategy are nonzero
            if (stakerStrategyShares[staker][strategy] == 0) {
                return false;
            }
            // check that strategy is not duplicated in array
            if (num_times_strategy_is_in_stakers_array(staker, strategy) != 1) {
                return false;
            }
        }
        return true;
    }

    function totalShares(address strategy) public view returns (uint256) {
        return IStrategy(strategy).totalShares();
    }

    function get_stakerStrategyShares(address staker, IStrategy strategy) public view returns (uint256) {
        return stakerStrategyShares[staker][strategy];
    }
}