// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./M2_Deploy_From_Scratch.s.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * forge script script/deploy/holesky/M2_Deploy_Preprod.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/M2_Deploy_Preprod.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 * Script for dev environment, exact same as M2_Deploy_From_Scratch.s.sol but with an EOAowner
 * instead of multisig addresses for permissions.
 * Unused config fields:
 * - init_strategy_whitelister
 * - multisig_addresses(operations, pauser, executor, community)
 */
contract M2_Deploy_Holesky_Preprod is M2_Deploy_Holesky_From_Scratch {
    /// @dev EOAowner is the deployer and owner of the contracts
    address EOAowner;

    function run() external virtual override {
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_preprod.holesky.config.json");

        // Overwrite multisig to be EOAowner
        EOAowner = msg.sender;
        executorMultisig = EOAowner;
        operationsMultisig = EOAowner;
        pauserMultisig = EOAowner;
        communityMultisig = EOAowner;
        STRATEGY_MANAGER_WHITELISTER = EOAowner;

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer and EOAowner Address", EOAowner);

        _deployFromScratch();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams(); // override to check contract.owner() is EOAowner instead

        logAndOutputContractAddresses("script/output/holesky/M2_deploy_preprod.holesky.config.json");
    }
}
