// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationChecks.t.sol";

contract IntegrationMainnetFork_UpgradeSetup is IntegrationCheckUtils {

    // /// @notice Test upgrade setup is correct
    // /// forge-config: default.fuzz.runs = 1
    // function test_mainnet_upgrade_setup(uint24 _random) public {
    //     _configRand({
    //         _randomSeed: _random,
    //         _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
    //         _userTypes: DEFAULT | ALT_METHODS,
    //         _forkTypes: MAINNET
    //     });

    //     // // 1. Check proper state pre-upgrade
    //     // _verifyContractPointers();
    //     // _verifyImplementations();
    //     // _verifyContractsInitialized({isInitialDeployment: true});
    //     // _verifyInitializationParams();

    //     // 2. Upgrade mainnet contracts
    //     _upgradeEigenLayerContracts();
    //     _parseInitialDeploymentParams("script/configs/mainnet/M2_mainnet_upgrade.config.json");

    //     // 2. Verify upgrade setup
    //     _verifyContractPointers();
    //     _verifyImplementations();
    //     _verifyContractsInitialized({isInitialDeployment: true});
    //     _verifyInitializationParams();
    // }

    // /// @notice Test upgrade setup is correct
    // /// forge-config: default.fuzz.runs = 1
    // function test_holesky_upgrade_setup(uint24 _random) public {
    //     _configRand({
    //         _randomSeed: _random,
    //         _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
    //         _userTypes: DEFAULT | ALT_METHODS,
    //         _forkTypes: HOLESKY
    //     });

    //     // // 1. Check proper state pre-upgrade
    //     // _verifyContractPointers();
    //     // _verifyImplementations();
    //     // _verifyContractsInitialized({isInitialDeployment: true});
    //     // _verifyInitializationParams();

    //     // 2. Upgrade holesky contracts
    //     _upgradeEigenLayerContracts();
    //     _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_from_scratch.holesky.config.json");

    //     // 3. Verify upgrade setup
    //     _verifyContractPointers();
    //     _verifyImplementations();
    //     _verifyContractsInitialized({isInitialDeployment: true});
    //     _verifyInitializationParams();
    // }

    /// @notice Ensure contracts point at each other correctly via constructors
    /// override to remove ethPOSDeposit contract check
    function _verifyContractPointers() internal virtual override view {
        // AVSDirectory
        require(
            avsDirectory.delegation() == delegationManager,
            "avsDirectory: delegationManager address not set correctly"
        );
        // DelegationManager
        require(delegationManager.slasher() == slasher, "delegationManager: slasher address not set correctly");
        require(
            delegationManager.strategyManager() == strategyManager,
            "delegationManager: strategyManager address not set correctly"
        );
        require(
            delegationManager.eigenPodManager() == eigenPodManager,
            "delegationManager: eigenPodManager address not set correctly"
        );
        // StrategyManager
        require(strategyManager.slasher() == slasher, "strategyManager: slasher address not set correctly");
        require(
            strategyManager.delegation() == delegationManager,
            "strategyManager: delegationManager address not set correctly"
        );
        require(
            strategyManager.eigenPodManager() == eigenPodManager,
            "strategyManager: eigenPodManager address not set correctly"
        );
        // EPM
        require(
            eigenPodManager.eigenPodBeacon() == eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        require(
            eigenPodManager.strategyManager() == strategyManager,
            "eigenPodManager: strategyManager contract address not set correctly"
        );
        require(eigenPodManager.slasher() == slasher, "eigenPodManager: slasher contract address not set correctly");
        require(
            eigenPodManager.delegationManager() == delegationManager,
            "eigenPodManager: delegationManager contract address not set correctly"
        );
        // DelayedWithdrawalRouter
        require(
            delayedWithdrawalRouter.eigenPodManager() == eigenPodManager,
            "delayedWithdrawalRouterContract: eigenPodManager address not set correctly"
        );
    }
}