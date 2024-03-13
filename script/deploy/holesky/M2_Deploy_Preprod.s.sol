// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

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
contract M2_Deploy_Holesky_Preprod is ExistingDeploymentParser {
    /// @dev EOAowner is the deployer and owner of the contracts
    address EOAowner;

    function run() external {
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_preprod.holesky.config.json");

        EOAowner = msg.sender;

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

    /**
     * @notice Deploy EigenLayer contracts from scratch for Holesky
     */
    function _deployFromScratch() internal {
        // Deploy ProxyAdmin, later set admins for all proxies to be EOAowner
        eigenLayerProxyAdmin = new ProxyAdmin();

        // Set EOAowners as pausers and unpauser
        address[] memory pausers = new address[](1);
        pausers[0] = EOAowner;
        address unpauser = EOAowner;
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
                EOAowner, // initialOwner
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
                EOAowner, // initialOwner
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
                EOAowner, //initialOwner
                EOAowner, //initial whitelister
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
                EOAowner,
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
                EOAowner,
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
                EOAowner, // initialOwner
                eigenLayerPauserReg,
                DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS,
                DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS
            )
        );

        // Deploy Strategies
        baseStrategyImplementation = new StrategyBaseTVLLimits(strategyManager);
        uint256 numStrategiesToDeploy = strategiesToDeploy.length;
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

            deployedStrategyArray.push(strategy);
        }

        // Fork timestamp config
        eigenPodManager.setDenebForkTimestamp(EIGENPOD_MANAGER_DENEB_FORK_TIMESTAMP);

        // Transfer ownership
        eigenLayerProxyAdmin.transferOwnership(EOAowner);
        eigenPodBeacon.transferOwnership(EOAowner);
    }

    function _verifyInitializationParams() internal view override {
        // AVSDirectory
        require(
            avsDirectory.pauserRegistry() == eigenLayerPauserReg,
            "avsdirectory: pauser registry not set correctly"
        );
        require(avsDirectory.owner() == EOAowner, "avsdirectory: owner not set correctly");
        require(
            avsDirectory.paused() == AVS_DIRECTORY_INIT_PAUSED_STATUS,
            "avsdirectory: init paused status set incorrectly"
        );
        // DelegationManager
        require(
            delegationManager.pauserRegistry() == eigenLayerPauserReg,
            "delegationManager: pauser registry not set correctly"
        );
        require(delegationManager.owner() == EOAowner, "delegationManager: owner not set correctly");
        require(
            delegationManager.paused() == DELEGATION_MANAGER_INIT_PAUSED_STATUS,
            "delegationManager: init paused status set incorrectly"
        );
        require(
            delegationManager.minWithdrawalDelayBlocks() == DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
            "delegationManager: minWithdrawalDelayBlocks not set correctly"
        );
        // StrategyManager
        require(
            strategyManager.pauserRegistry() == eigenLayerPauserReg,
            "strategyManager: pauser registry not set correctly"
        );
        require(strategyManager.owner() == EOAowner, "strategyManager: owner not set correctly");
        require(
            strategyManager.paused() == STRATEGY_MANAGER_INIT_PAUSED_STATUS,
            "strategyManager: init paused status set incorrectly"
        );
        require(
            strategyManager.strategyWhitelister() == EOAowner,
            "strategyManager: strategyWhitelister not set correctly"
        );
        // EigenPodManager
        require(
            eigenPodManager.pauserRegistry() == eigenLayerPauserReg,
            "eigenPodManager: pauser registry not set correctly"
        );
        require(eigenPodManager.owner() == EOAowner, "eigenPodManager: owner not set correctly");
        require(
            eigenPodManager.paused() == EIGENPOD_MANAGER_INIT_PAUSED_STATUS,
            "eigenPodManager: init paused status set incorrectly"
        );
        // EigenPodBeacon
        require(eigenPodBeacon.owner() == EOAowner, "eigenPodBeacon: owner not set correctly");
        // DelayedWithdrawalRouter
        require(
            delayedWithdrawalRouter.pauserRegistry() == eigenLayerPauserReg,
            "delayedWithdrawalRouter: pauser registry not set correctly"
        );
        require(delayedWithdrawalRouter.owner() == EOAowner, "delayedWithdrawalRouter: owner not set correctly");
        require(
            delayedWithdrawalRouter.paused() == DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS,
            "delayedWithdrawalRouter: init paused status set incorrectly"
        );
        require(
            delayedWithdrawalRouter.withdrawalDelayBlocks() == DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS,
            "delayedWithdrawalRouter: withdrawalDelayBlocks not set correctly"
        );
        // Strategies
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                deployedStrategyArray[i].pauserRegistry() == eigenLayerPauserReg,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            require(
                deployedStrategyArray[i].paused() == 0,
                "StrategyBaseTVLLimits: init paused status set incorrectly"
            );
        }

        // Pausing Permissions
        require(eigenLayerPauserReg.isPauser(EOAowner), "pauserRegistry: EOAowner is not pauser");
        require(eigenLayerPauserReg.unpauser() == EOAowner, "pauserRegistry: unpauser not set correctly");
    }
}
