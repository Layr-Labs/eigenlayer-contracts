// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Deploy_Test_RewardsCoordinator.s.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 * Local Fork: Deploy/Upgrade RewardsCoordinator
 * forge script script/deploy/holesky/Deploy_Preprod_RewardsCoordinator.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv --sig "run(string memory deployArg)" deploy
 * forge script script/deploy/holesky/Deploy_Preprod_RewardsCoordinator.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv --sig "run(string memory deployArg)" upgrade
 *
 * Holesky testnet: Deploy/Upgrade RewardsCoordinator
 * forge script script/deploy/holesky/Deploy_Preprod_RewardsCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --sig "run(string memory deployArg)" deploy
 * forge script script/deploy/holesky/Deploy_Preprod_RewardsCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --sig "run(string memory deployArg)" upgrade
 *
 */
contract Deploy_Preprod_RewardsCoordinator is Deploy_Test_RewardsCoordinator {
    function run(string memory deployArg) external virtual {
        _parseInitialDeploymentParams("script/configs/holesky/Deploy_RewardsCoordinator.holesky.config.json");
        _parseDeployedContracts("script/output/holesky/M2_deploy_preprod.output.json");

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

        if (keccak256(abi.encode(deployArg)) == keccak256(abi.encode("upgrade"))) {
            _upgradeRewardsCoordinator();
        } else if (keccak256(abi.encode(deployArg)) == keccak256(abi.encode("deploy"))) {
            _deployRewardsCoordinator();
        }

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/Deploy_RewardsCoordinator_Preprod.holesky.config.json");
    }
}
