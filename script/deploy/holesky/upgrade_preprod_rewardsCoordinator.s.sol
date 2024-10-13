// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Deploy_Test_RewardsCoordinator.s.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 *
 * Holesky testnet: Deploy/Upgrade RewardsCoordinator
 * forge script script/deploy/holesky/upgrade_preprod_rewardsCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv 
 *
 */
contract Upgrade_Preprod_RewardsCoordinator is Deploy_Test_RewardsCoordinator {
    function run() external virtual override {
        _parseInitialDeploymentParams("script/configs/holesky/eigenlayer_preprod.config.json");
        _parseDeployedContracts("script/configs/holesky/eigenlayer_addresses_preprod.config.json");

        // Overwrite testAddress and multisigs to be EOAowner
        testAddress = msg.sender;
        executorMultisig = testAddress;
        operationsMultisig = testAddress;
        pauserMultisig = testAddress;
        communityMultisig = testAddress;
        STRATEGY_MANAGER_WHITELISTER = testAddress;

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _upgradeRewardsCoordinator();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();

    }
}
