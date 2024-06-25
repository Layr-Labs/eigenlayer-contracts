// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/EigenPod_Checkpoint_Deploy_Preprod.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/EigenPod_Checkpoint_Deploy_Preprod.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 *
 */
contract EigenPod_Checkpoint_Deploy_Preprod is ExistingDeploymentParser {

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    address initOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/holesky/eigenlayer.config.json"
        );
        _parseDeployedContracts(
            "script/output/holesky/eigenlayer_addresses.config.json"
        );

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _upgradeEigenPodAndEPM();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/EigenPod_Checkpoint_Deploy_Preprod.output.json");
    }

    /**
     * @notice Deploy EigenPod and EPM Implementation for Holesky preprod and upgrade the beacon/proxy
     */
    function _upgradeEigenPodAndEPM() internal {
        // Deploy RewardsCoordinator proxy and implementation
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );

        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
            address(rewardsCoordinatorImplementation)
        );
    }
}
