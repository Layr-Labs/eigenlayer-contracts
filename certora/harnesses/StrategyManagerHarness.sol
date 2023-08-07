// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../munged/core/StrategyManager.sol";

contract StrategyManagerHarness is StrategyManager {
    constructor(IDelegationManager _delegation, IEigenPodManager _eigenPodManager, ISlasher _slasher)
        StrategyManager(_delegation, _eigenPodManager, _slasher)
        {}

    function slashSharesSinglet(
        address slashedAddress,
        address recipient,
        IStrategy strategy,
        IERC20 token,
        uint256 strategyIndex,
        uint256 shareAmount
    )
        external
        onlyOwner
        onlyFrozen(slashedAddress)
        nonReentrant
    {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = token;
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = strategyIndex;
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = shareAmount;
        require(tokens.length == strategies.length, "StrategyManager.slashShares: input length mismatch");
        uint256 strategyIndexIndex;
        uint256 strategiesLength = strategies.length;
        for (uint256 i = 0; i < strategiesLength;) {
            // the internal function will return 'true' in the event the strategy was
            // removed from the slashedAddress's array of strategies -- i.e. stakerStrategyList[slashedAddress]
            if (_removeShares(slashedAddress, strategyIndexes[strategyIndexIndex], strategies[i], shareAmounts[i])) {
                unchecked {
                    ++strategyIndexIndex;
                }
            }

            if (strategies[i] == beaconChainETHStrategy) {
                 //withdraw the beaconChainETH to the recipient
                eigenPodManager.withdrawRestakedBeaconChainETH(slashedAddress, recipient, shareAmounts[i]);
            }
            else {
                // withdraw the shares and send funds to the recipient
                strategies[i].withdraw(recipient, tokens[i], shareAmounts[i]);
            }

            // increment the loop
            unchecked {
                ++i;
            }
        }

        // modify delegated shares accordingly, if applicable
        delegation.decreaseDelegatedShares(slashedAddress, strategies, shareAmounts);
    }

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
        return  IStrategy(strategy).totalShares();
    }
}