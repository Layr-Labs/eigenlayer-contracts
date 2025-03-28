// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/DelegationManager.sol";
import "../../../src/contracts/core/AVSDirectory.sol";
import "../../../src/contracts/core/RewardsCoordinator.sol";
import "../../../src/contracts/core/AllocationManager.sol";
import "../../../src/contracts/permissions/PermissionController.sol";

import "../../../src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "../../../src/contracts/strategies/StrategyFactory.sol";
import "../../../src/contracts/strategies/StrategyBase.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";

import "../../../src/contracts/permissions/PauserRegistry.sol";

import "../../../src/test/mocks/EmptyContract.sol";
import "../../../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/devnet/deploy_from_scratch.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/deploy_from_scratch.anvil.config.json
contract DeployFromScratch is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    string public deployConfigPath;

    // EigenLayer Contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    DelegationManager public delegation;
    DelegationManager public delegationImplementation;
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    RewardsCoordinator public rewardsCoordinator;
    RewardsCoordinator public rewardsCoordinatorImplementation;
    AVSDirectory public avsDirectory;
    AVSDirectory public avsDirectoryImplementation;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    UpgradeableBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;
    StrategyFactory public strategyFactory;
    StrategyFactory public strategyFactoryImplementation;
    UpgradeableBeacon public strategyBeacon;
    StrategyBase public baseStrategyImplementation;
    AllocationManager public allocationManagerImplementation;
    AllocationManager public allocationManager;
    PermissionController public permissionController;
    PermissionController public permissionControllerImplementation;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;

    // the ETH2 deposit contract -- if not on mainnet, we deploy a mock as stand-in
    IETHPOSDeposit public ethPOSDeposit;

    // strategies deployed
    StrategyBaseTVLLimits[] public deployedStrategyArray;

    string SEMVER;

    // IMMUTABLES TO SET
    uint64 HOLESKY_GENESIS_TIME = 1_616_508_000;

    // OTHER DEPLOYMENT PARAMETERS
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    uint256 DELEGATION_INIT_PAUSED_STATUS;
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;

    // DelegationManager
    uint32 MIN_WITHDRAWAL_DELAY = 86_400;

    // AllocationManager
    uint32 DEALLOCATION_DELAY;
    uint32 ALLOCATION_CONFIGURATION_DELAY;

    // RewardsCoordinator
    uint32 REWARDS_COORDINATOR_MAX_REWARDS_DURATION;
    uint32 REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH;
    uint32 REWARDS_COORDINATOR_MAX_FUTURE_LENGTH;
    uint32 REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
    address REWARDS_COORDINATOR_UPDATER;
    uint32 REWARDS_COORDINATOR_ACTIVATION_DELAY;
    uint32 REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS;
    uint32 REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS;
    uint32 REWARDS_COORDINATOR_OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP;
    uint32 REWARDS_COORDINATOR_OPERATOR_SET_MAX_RETROACTIVE_LENGTH;

    // AllocationManager
    uint256 ALLOCATION_MANAGER_INIT_PAUSED_STATUS;

    // one week in blocks -- 50400
    uint32 STRATEGY_MANAGER_INIT_WITHDRAWAL_DELAY_BLOCKS;
    uint256 DELEGATION_WITHDRAWAL_DELAY_BLOCKS;

    function run(
        string memory configFileName
    ) public {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        // READ JSON CONFIG DATA
        deployConfigPath = string(bytes(string.concat("script/configs/", configFileName)));
        string memory config_data = vm.readFile(deployConfigPath);
        // bytes memory parsedData = vm.parseJson(config_data);

        SEMVER = stdJson.readString(config_data, ".semver");

        STRATEGY_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(config_data, ".strategyManager.init_paused_status");
        DELEGATION_INIT_PAUSED_STATUS = stdJson.readUint(config_data, ".delegation.init_paused_status");
        DELEGATION_WITHDRAWAL_DELAY_BLOCKS = stdJson.readUint(config_data, ".delegation.init_withdrawal_delay_blocks");
        EIGENPOD_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(config_data, ".eigenPodManager.init_paused_status");
        REWARDS_COORDINATOR_INIT_PAUSED_STATUS = stdJson.readUint(config_data, ".rewardsCoordinator.init_paused_status");
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.CALCULATION_INTERVAL_SECONDS"));
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.MAX_REWARDS_DURATION"));
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.MAX_RETROACTIVE_LENGTH"));
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.MAX_FUTURE_LENGTH"));
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP"));
        REWARDS_COORDINATOR_UPDATER = stdJson.readAddress(config_data, ".rewardsCoordinator.rewards_updater_address");
        REWARDS_COORDINATOR_ACTIVATION_DELAY =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.activation_delay"));
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.calculation_interval_seconds"));
        REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.global_operator_commission_bips"));
        REWARDS_COORDINATOR_OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP"));
        REWARDS_COORDINATOR_OPERATOR_SET_MAX_RETROACTIVE_LENGTH =
            uint32(stdJson.readUint(config_data, ".rewardsCoordinator.OPERATOR_SET_MAX_RETROACTIVE_LENGTH"));

        STRATEGY_MANAGER_INIT_WITHDRAWAL_DELAY_BLOCKS =
            uint32(stdJson.readUint(config_data, ".strategyManager.init_withdrawal_delay_blocks"));

        ALLOCATION_MANAGER_INIT_PAUSED_STATUS =
            uint32(stdJson.readUint(config_data, ".allocationManager.init_paused_status"));
        DEALLOCATION_DELAY = uint32(stdJson.readUint(config_data, ".allocationManager.DEALLOCATION_DELAY"));
        ALLOCATION_CONFIGURATION_DELAY =
            uint32(stdJson.readUint(config_data, ".allocationManager.ALLOCATION_CONFIGURATION_DELAY"));

        executorMultisig = stdJson.readAddress(config_data, ".multisig_addresses.executorMultisig");
        operationsMultisig = stdJson.readAddress(config_data, ".multisig_addresses.operationsMultisig");
        pauserMultisig = stdJson.readAddress(config_data, ".multisig_addresses.pauserMultisig");

        require(executorMultisig != address(0), "executorMultisig address not configured correctly!");
        require(operationsMultisig != address(0), "operationsMultisig address not configured correctly!");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();

        //deploy pauser registry
        {
            address[] memory pausers = new address[](3);
            pausers[0] = executorMultisig;
            pausers[1] = operationsMultisig;
            pausers[2] = pauserMultisig;
            eigenLayerPauserReg = new PauserRegistry(pausers, executorMultisig);
        }

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        emptyContract = new EmptyContract();
        delegation = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        rewardsCoordinator = RewardsCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        allocationManager = AllocationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyFactory = StrategyFactory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        permissionController = PermissionController(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // if on mainnet, use the ETH2 deposit contract address
        if (chainId == 1) ethPOSDeposit = IETHPOSDeposit(0x00000000219ab540356cBB839Cbe05303d7705Fa);
        // if not on mainnet, deploy a mock
        else ethPOSDeposit = IETHPOSDeposit(stdJson.readAddress(config_data, ".ethPOSDepositAddress"));
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, HOLESKY_GENESIS_TIME, SEMVER);

        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs

        delegationImplementation = new DelegationManager(
            strategyManager,
            eigenPodManager,
            allocationManager,
            eigenLayerPauserReg,
            permissionController,
            MIN_WITHDRAWAL_DELAY,
            SEMVER
        );

        strategyManagerImplementation = new StrategyManager(delegation, eigenLayerPauserReg, SEMVER);
        avsDirectoryImplementation = new AVSDirectory(delegation, eigenLayerPauserReg, SEMVER);
        eigenPodManagerImplementation =
            new EigenPodManager(ethPOSDeposit, eigenPodBeacon, delegation, eigenLayerPauserReg, SEMVER);
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams(
                delegation,
                strategyManager,
                allocationManager,
                eigenLayerPauserReg,
                permissionController,
                REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
                REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
                REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
                REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
                REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP,
                SEMVER
            )
        );
        allocationManagerImplementation = new AllocationManager(
            delegation,
            eigenLayerPauserReg,
            permissionController,
            DEALLOCATION_DELAY,
            ALLOCATION_CONFIGURATION_DELAY,
            SEMVER
        );
        permissionControllerImplementation = new PermissionController(SEMVER);
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg, SEMVER);

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        {
            eigenLayerProxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(payable(address(delegation))),
                address(delegationImplementation),
                abi.encodeWithSelector(
                    DelegationManager.initialize.selector, executorMultisig, DELEGATION_INIT_PAUSED_STATUS
                )
            );
        }
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                executorMultisig,
                operationsMultisig,
                STRATEGY_MANAGER_INIT_PAUSED_STATUS
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(AVSDirectory.initialize.selector, executorMultisig, eigenLayerPauserReg, 0)
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector, executorMultisig, EIGENPOD_MANAGER_INIT_PAUSED_STATUS
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
            address(rewardsCoordinatorImplementation),
            abi.encodeWithSelector(
                RewardsCoordinator.initialize.selector,
                executorMultisig,
                REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
                REWARDS_COORDINATOR_UPDATER,
                REWARDS_COORDINATOR_ACTIVATION_DELAY,
                REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS
            )
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(allocationManager))),
            address(allocationManagerImplementation),
            abi.encodeWithSelector(
                AllocationManager.initialize.selector, executorMultisig, ALLOCATION_MANAGER_INIT_PAUSED_STATUS
            )
        );

        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(permissionController))),
            address(permissionControllerImplementation)
        );

        // Deploy strategyFactory & base
        // Create base strategy implementation
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg, SEMVER);

        // Create a proxy beacon for base strategy implementation
        strategyBeacon = new UpgradeableBeacon(address(baseStrategyImplementation));

        // Strategy Factory, upgrade and initialized
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyFactory))),
            address(strategyFactoryImplementation),
            abi.encodeWithSelector(
                StrategyFactory.initialize.selector,
                executorMultisig,
                0, // initial paused status
                IBeacon(strategyBeacon)
            )
        );

        // Set the strategyWhitelister to the factory
        strategyManager.setStrategyWhitelister(address(strategyFactory));

        // Deploy a WETH strategy
        strategyFactory.deployNewStrategy(IERC20(address(0x94373a4919B3240D86eA41593D5eBa789FEF3848)));

        // Transfer ownership
        eigenLayerProxyAdmin.transferOwnership(executorMultisig);
        eigenPodBeacon.transferOwnership(executorMultisig);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // CHECK CORRECTNESS OF DEPLOYMENT
        _verifyContractsPointAtOneAnother(
            delegationImplementation,
            strategyManagerImplementation,
            eigenPodManagerImplementation,
            rewardsCoordinatorImplementation,
            allocationManagerImplementation
        );
        _verifyContractsPointAtOneAnother(
            delegation, strategyManager, eigenPodManager, rewardsCoordinator, allocationManager
        );
        _verifyImplementationsSetCorrectly();
        _verifyInitialOwners();
        _checkPauserInitializations();
        _verifyInitializationParams();

        // Check DM and AM have same withdrawa/deallocation delay
        // TODO: Update after AllocationManager is converted to timestamps as well
        // require(
        //     delegation.minWithdrawalDelayBlocks() == allocationManager.DEALLOCATION_DELAY(),
        //     "DelegationManager and AllocationManager have different withdrawal/deallocation delays"
        // );
        require(allocationManager.DEALLOCATION_DELAY() == 1 days);
        require(allocationManager.ALLOCATION_CONFIGURATION_DELAY() == 10 minutes);

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_strategies_output = "";

        string memory deployed_addresses = "addresses";
        vm.serializeUint(deployed_addresses, "numStrategiesDeployed", 0); // for compatibility with other scripts
        vm.serializeAddress(deployed_addresses, "eigenLayerProxyAdmin", address(eigenLayerProxyAdmin));
        vm.serializeAddress(deployed_addresses, "eigenLayerPauserReg", address(eigenLayerPauserReg));
        vm.serializeAddress(deployed_addresses, "delegationManager", address(delegation));
        vm.serializeAddress(deployed_addresses, "delegationManagerImplementation", address(delegationImplementation));
        vm.serializeAddress(deployed_addresses, "avsDirectory", address(avsDirectory));
        vm.serializeAddress(deployed_addresses, "avsDirectoryImplementation", address(avsDirectoryImplementation));
        vm.serializeAddress(deployed_addresses, "allocationManager", address(allocationManager));
        vm.serializeAddress(
            deployed_addresses, "allocationManagerImplementation", address(allocationManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "permissionController", address(permissionController));
        vm.serializeAddress(
            deployed_addresses, "permissionControllerImplementation", address(permissionControllerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(deployed_addresses, "strategyManagerImplementation", address(strategyManagerImplementation));
        vm.serializeAddress(deployed_addresses, "strategyFactory", address(strategyFactory));
        vm.serializeAddress(deployed_addresses, "strategyFactoryImplementation", address(strategyFactoryImplementation));
        vm.serializeAddress(deployed_addresses, "strategyBeacon", address(strategyBeacon));
        vm.serializeAddress(deployed_addresses, "baseStrategyImplementation", address(baseStrategyImplementation));
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodManagerImplementation", address(eigenPodManagerImplementation));
        vm.serializeAddress(deployed_addresses, "rewardsCoordinator", address(rewardsCoordinator));
        vm.serializeAddress(
            deployed_addresses, "rewardsCoordinatorImplementation", address(rewardsCoordinatorImplementation)
        );
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "eigenPodImplementation", address(eigenPodImplementation));
        vm.serializeAddress(deployed_addresses, "emptyContract", address(emptyContract));

        string memory deployed_addresses_output =
            vm.serializeString(deployed_addresses, "strategies", deployed_strategies_output);

        {
            // dummy token data
            string memory token =
                '{"tokenProxyAdmin": "0x0000000000000000000000000000000000000000", "EIGEN": "0x0000000000000000000000000000000000000000","bEIGEN": "0x0000000000000000000000000000000000000000","EIGENImpl": "0x0000000000000000000000000000000000000000","bEIGENImpl": "0x0000000000000000000000000000000000000000","eigenStrategy": "0x0000000000000000000000000000000000000000","eigenStrategyImpl": "0x0000000000000000000000000000000000000000"}';
            deployed_addresses_output = vm.serializeString(deployed_addresses, "token", token);
        }

        string memory parameters = "parameters";
        vm.serializeAddress(parameters, "executorMultisig", executorMultisig);
        vm.serializeAddress(parameters, "communityMultisig", operationsMultisig);
        vm.serializeAddress(parameters, "pauserMultisig", pauserMultisig);
        vm.serializeAddress(parameters, "timelock", address(0));
        string memory parameters_output = vm.serializeAddress(parameters, "operationsMultisig", operationsMultisig);

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        // serialize all the data
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        vm.serializeString(parent_object, chain_info, chain_info_output);
        string memory finalJson = vm.serializeString(parent_object, parameters, parameters_output);
        // TODO: should output to different file depending on configFile passed to run()
        //       so that we don't override mainnet output by deploying to goerli for eg.
        vm.writeJson(finalJson, "script/output/devnet/slashing_output.json");
    }

    function _verifyContractsPointAtOneAnother(
        DelegationManager delegationContract,
        StrategyManager strategyManagerContract,
        EigenPodManager eigenPodManagerContract,
        RewardsCoordinator rewardsCoordinatorContract,
        AllocationManager allocationManagerContract
    ) internal view {
        require(
            delegationContract.strategyManager() == strategyManager,
            "delegation: strategyManager address not set correctly"
        );

        require(
            strategyManagerContract.delegation() == delegation, "strategyManager: delegation address not set correctly"
        );
        require(
            eigenPodManagerContract.ethPOS() == ethPOSDeposit,
            " eigenPodManager: ethPOSDeposit contract address not set correctly"
        );
        require(
            eigenPodManagerContract.eigenPodBeacon() == eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );

        require(
            rewardsCoordinatorContract.delegationManager() == delegation,
            "rewardsCoordinator: delegation address not set correctly"
        );
        require(
            rewardsCoordinatorContract.strategyManager() == strategyManager,
            "rewardsCoordinator: strategyManager address not set correctly"
        );
        require(
            delegationContract.allocationManager() == allocationManager,
            "delegationManager: allocationManager address not set correctly"
        );
        require(
            allocationManagerContract.delegation() == delegation,
            "allocationManager: delegation address not set correctly"
        );
    }

    function _verifyImplementationsSetCorrectly() internal view {
        require(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(delegation))))
                == address(delegationImplementation),
            "delegation: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyManager))))
                == address(strategyManagerImplementation),
            "strategyManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(eigenPodManager))))
                == address(eigenPodManagerImplementation),
            "eigenPodManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(rewardsCoordinator)))
            ) == address(rewardsCoordinatorImplementation),
            "rewardsCoordinator: implementation set incorrectly"
        );

        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(allocationManager)))
            ) == address(allocationManagerImplementation),
            "allocationManager: implementation set incorrectly"
        );

        require(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyFactory))))
                == address(strategyFactoryImplementation),
            "strategyFactory: implementation set incorrectly"
        );

        require(
            eigenPodBeacon.implementation() == address(eigenPodImplementation),
            "eigenPodBeacon: implementation set incorrectly"
        );

        require(
            strategyBeacon.implementation() == address(baseStrategyImplementation),
            "strategyBeacon: implementation set incorrectly"
        );
    }

    function _verifyInitialOwners() internal view {
        require(strategyManager.owner() == executorMultisig, "strategyManager: owner not set correctly");
        require(delegation.owner() == executorMultisig, "delegation: owner not set correctly");
        require(eigenPodManager.owner() == executorMultisig, "eigenPodManager: owner not set correctly");
        require(allocationManager.owner() == executorMultisig, "allocationManager: owner not set correctly");
        require(eigenLayerProxyAdmin.owner() == executorMultisig, "eigenLayerProxyAdmin: owner not set correctly");
        require(eigenPodBeacon.owner() == executorMultisig, "eigenPodBeacon: owner not set correctly");
        require(strategyBeacon.owner() == executorMultisig, "strategyBeacon: owner not set correctly");
    }

    function _checkPauserInitializations() internal view {
        require(delegation.pauserRegistry() == eigenLayerPauserReg, "delegation: pauser registry not set correctly");
        require(
            strategyManager.pauserRegistry() == eigenLayerPauserReg,
            "strategyManager: pauser registry not set correctly"
        );
        require(
            eigenPodManager.pauserRegistry() == eigenLayerPauserReg,
            "eigenPodManager: pauser registry not set correctly"
        );
        require(
            rewardsCoordinator.pauserRegistry() == eigenLayerPauserReg,
            "rewardsCoordinator: pauser registry not set correctly"
        );
        require(
            allocationManager.pauserRegistry() == eigenLayerPauserReg,
            "allocationManager: pauser registry not set correctly"
        );

        require(eigenLayerPauserReg.isPauser(operationsMultisig), "pauserRegistry: operationsMultisig is not pauser");
        require(eigenLayerPauserReg.isPauser(executorMultisig), "pauserRegistry: executorMultisig is not pauser");
        require(eigenLayerPauserReg.isPauser(pauserMultisig), "pauserRegistry: pauserMultisig is not pauser");
        require(eigenLayerPauserReg.unpauser() == executorMultisig, "pauserRegistry: unpauser not set correctly");

        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                deployedStrategyArray[i].pauserRegistry() == eigenLayerPauserReg,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            require(deployedStrategyArray[i].paused() == 0, "StrategyBaseTVLLimits: init paused status set incorrectly");
        }

        // // pause *nothing*
        // uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS = 0;
        // // pause *everything*
        // // pause *everything*
        // uint256 DELEGATION_INIT_PAUSED_STATUS = type(uint256).max;
        // // pause *all of the proof-related functionality* (everything that can be paused other than creation of EigenPods)
        // uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS = (2**1) + (2**2) + (2**3) + (2**4); /* = 30 */
        // // pause *nothing*
        // require(strategyManager.paused() == 0, "strategyManager: init paused status set incorrectly");
        // require(delegation.paused() == type(uint256).max, "delegation: init paused status set incorrectly");
        // require(eigenPodManager.paused() == 30, "eigenPodManager: init paused status set incorrectly");
    }

    function _verifyInitializationParams() internal view {
        // // one week in blocks -- 50400
        // uint32 STRATEGY_MANAGER_INIT_WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
        // require(strategyManager.withdrawalDelayBlocks() == 7 days / 12 seconds,
        //     "strategyManager: withdrawalDelayBlocks initialized incorrectly");
        // uint256 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32 ether;

        require(
            address(strategyManager.strategyWhitelister()) == address(strategyFactory),
            "strategyManager: strategyWhitelister address not set correctly"
        );

        require(
            baseStrategyImplementation.strategyManager() == strategyManager,
            "baseStrategyImplementation: strategyManager set incorrectly"
        );

        require(
            eigenPodImplementation.ethPOS() == ethPOSDeposit,
            "eigenPodImplementation: ethPOSDeposit contract address not set correctly"
        );
        require(
            eigenPodImplementation.eigenPodManager() == eigenPodManager,
            " eigenPodImplementation: eigenPodManager contract address not set correctly"
        );
    }
}
