// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Deploy_Test_PaymentCoordinator.s.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/Deploy_Preprod_PaymentCoordinator.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/Deploy_Preprod_PaymentCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract Deploy_Preprod_PaymentCoordinator is Deploy_Test_PaymentCoordinator {
    function run() external virtual override {
        _parseInitialDeploymentParams("script/configs/holesky/Deploy_PaymentCoordinator.holesky.config.json");
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

        _deployPaymentCoordinator();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/Deploy_PaymentCoordinator_Preprod.holesky.config.json");
    }
}
