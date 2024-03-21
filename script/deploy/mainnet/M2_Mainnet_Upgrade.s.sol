// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_MAINNET
 * forge script script/deploy/mainnet/M2_Mainnet_Upgrade.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * 
 * forge script script/deploy/mainnet/M2_Mainnet_Upgrade.s.sol --rpc-url $RPC_MAINNET --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract M2_Mainnet_Upgrade is ExistingDeploymentParser {
    function run() external virtual {
        _parseDeployedContracts("script/output/mainnet/M1_deployment_mainnet_2023_6_9.json");
        _parseInitialDeploymentParams("script/configs/mainnet/M2_mainnet_upgrade.config.json");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployImplementationContracts();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Simulate upgrade of contracts to new implementations
        _simulateUpgrade();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/mainnet/M2_mainnet_upgrade.mainnet.config.json");
    }

    /**
     * @notice Deploy EigenLayer contracts from scratch for Holesky
     */
    function _deployImplementationContracts() internal {
        // 1. Deploy New TUPS
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
        avsDirectory = AVSDirectory(
            address(
                new TransparentUpgradeableProxy(
                    address(avsDirectoryImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        AVSDirectory.initialize.selector,
                        executorMultisig, // initialOwner
                        eigenLayerPauserReg,
                        AVS_DIRECTORY_INIT_PAUSED_STATUS
                    )
                )
            )
        );

        // 2. Deploy Implementations
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            delayedWithdrawalRouter,
            eigenPodManager,
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            EIGENPOD_GENESIS_TIME
        );
        delegationManagerImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenPodManager, slasher);
        slasherImplementation = new Slasher(strategyManager, delegationManager);
        eigenPodManagerImplementation = new EigenPodManager(
            IETHPOSDeposit(ETHPOSDepositAddress),
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
        delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(eigenPodManager);
    }

    function _simulateUpgrade() internal {

        vm.startPrank(executorMultisig);

        // First, upgrade the proxy contracts to point to the implementations
        // AVSDirectory
        // eigenLayerProxyAdmin.upgrade(
        //     TransparentUpgradeableProxy(payable(address(avsDirectory))),
        //     address(avsDirectoryImplementation)
        // );
        // DelegationManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation)
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation)
        );
        // Slasher
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation)
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );
        // Delayed Withdrawal Router
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation)
        );

        // Second, configure additional settings and paused statuses
        delegationManager.setMinWithdrawalDelayBlocks(DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        delegationManager.unpause(0);
        eigenPodManager.unpause(0);

        eigenPodManager.setDenebForkTimestamp(EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP);
        eigenPodManager.updateBeaconChainOracle(beaconOracle);
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        vm.stopPrank();
    }
}
