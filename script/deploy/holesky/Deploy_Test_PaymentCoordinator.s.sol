// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/Deploy_Test_PaymentCoordinator.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/Deploy_Test_PaymentCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract Deploy_Test_PaymentCoordinator is ExistingDeploymentParser {

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run() external virtual {
        _parseInitialDeploymentParams("script/configs/holesky/Deploy_PaymentCoordinator.holesky.config.json");
        _parseDeployedContracts("script/output/holesky/M2_deploy_from_scratch.output.json");

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

        logAndOutputContractAddresses("script/output/holesky/Deploy_PaymentCoordinator.holesky.config.json");
    }

    /**
     * @notice Deploy PaymentCoordinator for Holesky
     */
    function _deployPaymentCoordinator() internal {
        // Deploy PaymentCoordinator proxy and implementation
        paymentCoordinatorImplementation = new PaymentCoordinator(
            delegationManager,
            strategyManager,
            PAYMENT_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            PAYMENT_COORDINATOR_MAX_PAYMENT_DURATION,
            PAYMENT_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            PAYMENT_COORDINATOR_MAX_FUTURE_LENGTH,
            PAYMENT_COORDINATOR_GENESIS_PAYMENT_TIMESTAMP
        );
        paymentCoordinator = PaymentCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(paymentCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        PaymentCoordinator.initialize.selector,
                        testAddress, // initOwner
                        eigenLayerPauserReg,
                        PAYMENT_COORDINATOR_INIT_PAUSED_STATUS,
                        PAYMENT_COORDINATOR_UPDATER,
                        PAYMENT_COORDINATOR_ACTIVATION_DELAY,
                        PAYMENT_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS
                    )
                )
            )
        );


    }

        /**
     * @notice Deploy PaymentCoordinator Implementation for Holesky and upgrade the proxy
     */
    function _upgradePaymentCoordinator() internal {
        // Deploy PaymentCoordinator proxy and implementation
        paymentCoordinatorImplementation = new PaymentCoordinator(
            delegationManager,
            strategyManager,
            PAYMENT_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            PAYMENT_COORDINATOR_MAX_PAYMENT_DURATION,
            PAYMENT_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            PAYMENT_COORDINATOR_MAX_FUTURE_LENGTH,
            PAYMENT_COORDINATOR_GENESIS_PAYMENT_TIMESTAMP
        );

        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(paymentCoordinator))),
            address(paymentCoordinatorImplementation)
        );
    }
}
