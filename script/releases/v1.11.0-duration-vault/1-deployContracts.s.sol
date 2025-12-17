// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: use an EOA to deploy all of the new/updated contracts for the Duration Vault feature.
/// Contracts deployed:
/// /// Core
/// - StrategyManager (updated with beforeAddShares/beforeRemoveShares hooks)
/// /// Strategies
/// - EigenStrategy (inherits from updated StrategyBase)
/// - StrategyBase (updated with beforeAddShares/beforeRemoveShares hooks)
/// - StrategyBaseTVLLimits (inherits from updated StrategyBase)
/// - StrategyFactory (updated with duration vault beacon support)
/// - DurationVaultStrategy (new beacon implementation)
/// - DurationVaultStrategy beacon (new UpgradeableBeacon)
contract DeployContracts is CoreContractsDeployer {
    using Env for *;

    function _runAsEOA() internal override {
        // Only execute on version 1.10.0
        if (!Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

        vm.startBroadcast();

        /// core/
        deployStrategyManager();

        /// strategies/
        deployEigenStrategy();
        deployStrategyBase();
        deployStrategyBaseTVLLimits();
        deployStrategyFactory();

        // Deploy DurationVaultStrategy implementation
        DurationVaultStrategy durationVaultImpl = deployDurationVaultStrategy();

        // Deploy the beacon for DurationVaultStrategy if it doesn't exist
        _deployDurationVaultBeacon(address(durationVaultImpl));

        vm.stopBroadcast();
    }

    /// @notice Deploys the DurationVaultStrategy beacon pointing to the implementation
    function _deployDurationVaultBeacon(
        address implementation
    ) internal onlyEOA {
        // Deploy new beacon
        UpgradeableBeacon beacon = new UpgradeableBeacon(implementation);

        // Transfer ownership to the executor multisig
        beacon.transferOwnership(Env.executorMultisig());

        // Register the beacon in the environment
        deployBeacon({name: type(DurationVaultStrategy).name, deployedTo: address(beacon)});
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed() || !Env.isSource() || !Env._strEq(Env.envVersion(), "1.10.0")) {
            return;
        }

        // Deploy the contracts
        runAsEOA();

        // Run validation tests
        _validateDeployedContracts();
    }

    function _validateDeployedContracts() internal view {
        assertTrue(deploys().length == 7, "Deployed contracts should be 7");
        // Validate StrategyManager
        assertTrue(
            address(Env.impl.strategyManager()) != address(0), "StrategyManager implementation should be deployed"
        );
        assertTrue(
            address(Env.impl.strategyManager().delegation()) == address(Env.proxy.delegationManager()),
            "StrategyManager: delegationManager mismatch"
        );

        // Validate EigenStrategy
        assertTrue(address(Env.impl.eigenStrategy()) != address(0), "EigenStrategy implementation should be deployed");
        assertTrue(
            address(Env.impl.eigenStrategy().strategyManager()) == address(Env.proxy.strategyManager()),
            "EigenStrategy: strategyManager mismatch"
        );

        // Validate StrategyBase
        assertTrue(address(Env.impl.strategyBase()) != address(0), "StrategyBase implementation should be deployed");
        assertTrue(
            address(Env.impl.strategyBase().strategyManager()) == address(Env.proxy.strategyManager()),
            "StrategyBase: strategyManager mismatch"
        );

        // Validate StrategyBaseTVLLimits
        assertTrue(
            address(Env.impl.strategyBaseTVLLimits()) != address(0),
            "StrategyBaseTVLLimits implementation should be deployed"
        );
        assertTrue(
            address(Env.impl.strategyBaseTVLLimits().strategyManager()) == address(Env.proxy.strategyManager()),
            "StrategyBaseTVLLimits: strategyManager mismatch"
        );

        // Validate StrategyFactory
        assertTrue(
            address(Env.impl.strategyFactory()) != address(0), "StrategyFactory implementation should be deployed"
        );
        assertTrue(
            address(Env.impl.strategyFactory().strategyManager()) == address(Env.proxy.strategyManager()),
            "StrategyFactory: strategyManager mismatch"
        );

        // Validate DurationVaultStrategy implementation
        assertTrue(
            address(Env.impl.durationVaultStrategy()) != address(0),
            "DurationVaultStrategy implementation should be deployed"
        );
        assertTrue(
            address(Env.impl.durationVaultStrategy().strategyManager()) == address(Env.proxy.strategyManager()),
            "DurationVaultStrategy: strategyManager mismatch"
        );
        assertTrue(
            address(Env.impl.durationVaultStrategy().delegationManager()) == address(Env.proxy.delegationManager()),
            "DurationVaultStrategy: delegationManager mismatch"
        );
        assertTrue(
            address(Env.impl.durationVaultStrategy().allocationManager()) == address(Env.proxy.allocationManager()),
            "DurationVaultStrategy: allocationManager mismatch"
        );
        assertTrue(
            address(Env.impl.durationVaultStrategy().rewardsCoordinator()) == address(Env.proxy.rewardsCoordinator()),
            "DurationVaultStrategy: rewardsCoordinator mismatch"
        );

        // Validate DurationVaultStrategy beacon
        assertTrue(
            address(Env.beacon.durationVaultStrategy()) != address(0), "DurationVaultStrategy beacon should be deployed"
        );
        assertTrue(
            Env.beacon.durationVaultStrategy().implementation() == address(Env.impl.durationVaultStrategy()),
            "DurationVaultStrategy beacon should point to implementation"
        );
        assertTrue(
            Env.beacon.durationVaultStrategy().owner() == Env.executorMultisig(),
            "DurationVaultStrategy beacon should be owned by executor multisig"
        );
    }
}
