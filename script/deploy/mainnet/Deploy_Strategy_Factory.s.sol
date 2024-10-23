// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Create2.sol";
import "../../utils/ExistingDeploymentParser.sol";

import "../../../src/contracts/strategies/StrategyFactory.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * FORK LOCAL
 * anvil --fork-url $RPC_MAINNET
 * forge script script/deploy/mainnet/Deploy_Strategy_Factory.s.sol:MainnetStrategyFactoryDeploy --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 * MAINNET
 * forge script script/deploy/mainnet/Deploy_Strategy_Factory.s.sol:MainnetStrategyFactoryDeploy --rpc-url $RPC_MAINNET --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 *
 */

contract MainnetStrategyFactoryDeploy is ExistingDeploymentParser {
    function run() external virtual {
        // Use rewards config
        _parseInitialDeploymentParams(
            "script/configs/mainnet/v0.3.0-mainnet-rewards.config.json"
        );
        _parseDeployedContracts(
            "script/configs/mainnet/Mainnet_curent_deployment.config.json"
        );

                // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployStrategyFactory();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/mainnet/v0.3.2-mainnet-strategy-factory.output.json");
    }

    /**
     * @notice Deploy StrategyFactory for Mainnet
     */

    function _deployStrategyFactory() internal {
        strategyFactoryImplementation = new StrategyFactory(
            strategyManager, eigenLayerPauserReg
        );

        

    }
}