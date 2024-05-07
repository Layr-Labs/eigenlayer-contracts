// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * forge script script/deploy/holesky/M2_Deploy_From_Scratch.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 * 
 */
contract M2_Deploy_Holesky_From_Scratch is ExistingDeploymentParser {
    function run() external virtual {
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_from_scratch.holesky.config.json");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployFromScratch();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized({isInitialDeployment: true});
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/holesky/M2_deploy_from_scratch.holesky.config.json");
    }

    /**
     * @notice Deploy EigenLayer contracts from scratch for Holesky
     */
    function _deployFromScratch() internal {
        // Deploy ProxyAdmin, later set admins for all proxies to be executorMultisig
        eigenLayerProxyAdmin = new ProxyAdmin();

        // Set multisigs as pausers, executorMultisig as unpauser
        address[] memory pausers = new address[](3);
        pausers[0] = executorMultisig;
        pausers[1] = operationsMultisig;
        pausers[2] = pauserMultisig;
        address unpauser = executorMultisig;
        eigenLayerPauserReg = new PauserRegistry(pausers, unpauser);

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        emptyContract = new EmptyContract();
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        delegationManager = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        slasher = Slasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            delayedWithdrawalRouter,
            eigenPodManager,
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            EIGENPOD_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
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

        // Third, upgrade the proxy contracts to point to the implementations
        IStrategy[] memory initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        uint256[] memory initializeWithdrawalDelayBlocks = new uint256[](0);
        // AVSDirectory
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(
                AVSDirectory.initialize.selector,
                executorMultisig, // initialOwner
                eigenLayerPauserReg,
                AVS_DIRECTORY_INIT_PAUSED_STATUS
            )
        );
        // DelegationManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                executorMultisig, // initialOwner
                eigenLayerPauserReg,
                DELEGATION_MANAGER_INIT_PAUSED_STATUS,
                DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
                initializeStrategiesToSetDelayBlocks,
                initializeWithdrawalDelayBlocks
            )
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                msg.sender, //initialOwner, set to executorMultisig later after whitelisting strategies
                msg.sender, //initial whitelister, set to STRATEGY_MANAGER_WHITELISTER later
                eigenLayerPauserReg,
                STRATEGY_MANAGER_INIT_PAUSED_STATUS
            )
        );
        // Slasher
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(
                Slasher.initialize.selector,
                executorMultisig,
                eigenLayerPauserReg,
                SLASHER_INIT_PAUSED_STATUS
            )
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                beaconOracle,
                msg.sender, // initialOwner is msg.sender for now to set forktimestamp later
                eigenLayerPauserReg,
                EIGENPOD_MANAGER_INIT_PAUSED_STATUS
            )
        );
        // Delayed Withdrawal Router
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation),
            abi.encodeWithSelector(
                DelayedWithdrawalRouter.initialize.selector,
                executorMultisig, // initialOwner
                eigenLayerPauserReg,
                DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS,
                DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS
            )
        );

        // Deploy Strategies
        baseStrategyImplementation = new StrategyBaseTVLLimits(strategyManager);
        uint256 numStrategiesToDeploy = strategiesToDeploy.length;
        // whitelist params
        IStrategy[] memory strategiesToWhitelist = new IStrategy[](numStrategiesToDeploy);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](numStrategiesToDeploy);

        for (uint256 i = 0; i < numStrategiesToDeploy; i++) {
            StrategyUnderlyingTokenConfig memory strategyConfig = strategiesToDeploy[i];

            // Deploy and upgrade strategy
            StrategyBaseTVLLimits strategy = StrategyBaseTVLLimits(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
            );
            eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(strategy))),
                address(baseStrategyImplementation),
                abi.encodeWithSelector(
                    StrategyBaseTVLLimits.initialize.selector,
                    STRATEGY_MAX_PER_DEPOSIT,
                    STRATEGY_MAX_TOTAL_DEPOSITS,
                    IERC20(strategyConfig.tokenAddress),
                    eigenLayerPauserReg
                )
            );

            strategiesToWhitelist[i] = strategy;
            thirdPartyTransfersForbiddenValues[i] = false;

            deployedStrategyArray.push(strategy);
        }

        // Add strategies to whitelist and set whitelister to STRATEGY_MANAGER_WHITELISTER
        strategyManager.addStrategiesToDepositWhitelist(strategiesToWhitelist, thirdPartyTransfersForbiddenValues);
        strategyManager.setStrategyWhitelister(STRATEGY_MANAGER_WHITELISTER);

        // Fork timestamp config
        eigenPodManager.setDenebForkTimestamp(EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP);

        // Transfer ownership
        strategyManager.transferOwnership(executorMultisig);
        eigenLayerProxyAdmin.transferOwnership(executorMultisig);
        eigenPodManager.transferOwnership(executorMultisig);
        eigenPodBeacon.transferOwnership(executorMultisig);
    }
}
