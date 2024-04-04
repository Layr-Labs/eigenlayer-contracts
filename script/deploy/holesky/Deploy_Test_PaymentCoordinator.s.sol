// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * forge script script/deploy/holesky/M2_Deploy_From_Scratch.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract Deploy_PaymentCoordinator is ExistingDeploymentParser {
    function run() external virtual {
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_from_scratch.holesky.config.json");
        _parseDeployedContracts("script/output/mainnet/M1_deployment_mainnet_2023_6_9.json");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        // _deployPaymentCoordinator();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/M2_deploy_from_scratch.holesky.config.json");
    }

    // /**
    //  * @notice Deploy PaymentCoordinator for Holesky
    //  */
    // function _deployPaymentCoordinator() internal {
    //     // Deploy PaymentCoordinator proxy and implementation
    //     paymentCoordinatorImplementation = new PaymentCoordinator(
    //         delegationManager,
    //         strategyManager,
    //         MAX_PAYMENT_DURATION,
    //         MAX_RETROACTIVE_LENGTH,
    //         MAX_FUTURE_LENGTH,
    //         GENESIS_PAYMENT_TIMESTAMP
    //     );
    //     paymentCoordinator = PaymentCoordinator(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(paymentCoordinatorImplementation),
    //                 address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(
    //                     PaymentCoordinator.initialize.selector,
    //                     address(this), // initOwner
    //                     pauserRegistry,
    //                     0, // 0 is initialPausedStatus
    //                     paymentUpdater,
    //                     activationDelay,
    //                     calculationIntervalSeconds,
    //                     globalCommissionBips
    //                 )
    //             )
    //         )
    //     );


    // }
}
