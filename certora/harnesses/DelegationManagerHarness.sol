// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/core/DelegationManager.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(IStrategyManager _strategyManager, ISlasher _slasher, IEigenPodManager _eigenPodManager)
        DelegationManager(_strategyManager, _slasher, _eigenPodManager) {}

    function get_operatorShares(address operator, IStrategy strategy) public view returns (uint256) {
        return operatorShares[operator][strategy];
    }

    function get_stakerDelegateableShares(address staker, IStrategy strategy) public view returns (uint256) {
        // this is the address of the virtual 'beaconChainETH' strategy
        if (address(strategy) == 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0) {
            int256 beaconChainETHShares = eigenPodManager.podOwnerShares(staker);
            if (beaconChainETHShares <= 0) {
                return 0;
            } else {
                return uint256(beaconChainETHShares);
            }
        } else {
            return strategyManager.stakerStrategyShares(staker, strategy);
        }
    }
}
