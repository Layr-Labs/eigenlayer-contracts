// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/core/DelegationManager.sol";
import "forge-std/Test.sol";

contract DelegationManagerHarness is DelegationManager {

    constructor(
        IStrategyManager _strategyManager,
        IEigenPodManager _eigenPodManager,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
        IPermissionController _permissionController,
        uint32 _MIN_WITHDRAWAL_DELAY
    )
        DelegationManager(
            _strategyManager,
            _eigenPodManager,
            _allocationManager,
            _pauserRegistry,
            _permissionController,
            _MIN_WITHDRAWAL_DELAY
        )
    {}

    function getSlashingFactor(
        address staker,
        IStrategy strategy,
        uint64 operatorMaxMagnitude
    ) external view returns (uint256) {
        return _getSlashingFactor(staker, strategy, operatorMaxMagnitude);
    }

    function getSlashingFactors(
        address staker,
        address operator,
        IStrategy[] memory strategies
    ) external view returns (uint256[] memory) {
        return _getSlashingFactors(staker, operator, strategies);
    }

    function getSlashingFactorsAtBlock(
        address staker,
        address operator,
        IStrategy[] memory strategies,
        uint32 blockNumber
    ) external view returns (uint256[] memory) {
        return _getSlashingFactorsAtBlock(staker, operator, strategies, blockNumber);
    }
}
