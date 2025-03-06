// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/core/DelegationManager.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _MIN_WITHDRAWAL_DELAY,
        string memory _version
    )
        DelegationManager(
            _strategyManager,
            _eigenPodManager,
            _allocationManager,
            _pauserRegistry,
            _permissionController,
            _MIN_WITHDRAWAL_DELAY,
            _version
        ) {}

    function get_operatorShares(address operator, IStrategy strategy) public view returns (uint256) {
        return operatorShares[operator][strategy];
    }

    function get_stakerDelegateableShares(address staker, IStrategy strategy) public view returns (uint256) {
        // this is the address of the virtual 'beaconChainETH' strategy
        if (address(strategy) == 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0) {
            int256 beaconChainETHShares = eigenPodManager.podOwnerDepositShares(staker);
            if (beaconChainETHShares <= 0) {
                return 0;
            } else {
                return uint256(beaconChainETHShares);
            }
        } else {
            return strategyManager.stakerDepositShares(staker, strategy);
        }
    }

    function get_min_withdrawal_delay_blocks() public view returns (uint32) {
        return MIN_WITHDRAWAL_DELAY_BLOCKS;
    }

    function canCall(address account, address caller, address target, uint32 selector) external returns (bool) {
        return permissionController.canCall(account, caller, target, bytes4(selector));
    }
}
