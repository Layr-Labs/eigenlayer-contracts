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
// forge script script/deploy/local/deploy_from_scratch.slashing.s.sol --rpc-url $RPC_URL -vvvv --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/deploy_from_scratch.slashing.anvil.config.json
contract DeployFromScratch is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    // struct used to encode token info in config file
    struct StrategyConfig {
        uint256 maxDeposits;
        uint256 maxPerDeposit;
        address tokenAddress;
        string tokenSymbol;
    }

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
    StrategyBase public baseStrategyImplementation;
    AllocationManager public allocationManagerImplementation;
    AllocationManager public allocationManager;
    PermissionController public permissionControllerImplementation;
    PermissionController public permissionController;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;

    string SEMVER;

    // the ETH2 deposit contract -- if not on mainnet, we deploy a mock as stand-in
    IETHPOSDeposit public ethPOSDeposit;

    // strategies deployed
    StrategyBaseTVLLimits[] public deployedStrategyArray;

    // IMMUTABLES TO SET
    uint64 HOLESKY_GENESIS_TIME = 1_616_508_000;

    // OTHER DEPLOYMENT PARAMETERS
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    uint256 DELEGATION_INIT_PAUSED_STATUS;
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;

    // DelegationManager
    uint32 MIN_WITHDRAWAL_DELAY;

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
    uint32 REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS;
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

        MIN_WITHDRAWAL_DELAY = uint32(stdJson.readUint(config_data, ".delegation.withdrawal_delay_blocks"));
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
        REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS =
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

        // tokens to deploy strategies for
        StrategyConfig[] memory strategyConfigs;

        executorMultisig = stdJson.readAddress(config_data, ".multisig_addresses.executorMultisig");
        operationsMultisig = stdJson.readAddress(config_data, ".multisig_addresses.operationsMultisig");
        pauserMultisig = stdJson.readAddress(config_data, ".multisig_addresses.pauserMultisig");
        // load token list
        bytes memory strategyConfigsRaw = stdJson.parseRaw(config_data, ".strategies");
        strategyConfigs = abi.decode(strategyConfigsRaw, (StrategyConfig[]));

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

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        {
            IStrategy[] memory _strategies;
            uint256[] memory _withdrawalDelayBlocks;
            eigenLayerProxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(payable(address(delegation))),
                address(delegationImplementation),
                abi.encodeWithSelector(
                    DelegationManager.initialize.selector,
                    executorMultisig,
                    DELEGATION_INIT_PAUSED_STATUS,
                    DELEGATION_WITHDRAWAL_DELAY_BLOCKS,
                    _strategies,
                    _withdrawalDelayBlocks
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
            abi.encodeWithSelector(AVSDirectory.initialize.selector, executorMultisig, 0)
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
                REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS
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

        // deploy StrategyBaseTVLLimits contract implementation
        baseStrategyImplementation = new StrategyBaseTVLLimits(strategyManager, eigenLayerPauserReg, SEMVER);
        // create upgradeable proxies that each point to the implementation and initialize them
        for (uint256 i = 0; i < strategyConfigs.length; ++i) {
            if (strategyConfigs[i].tokenAddress == address(0)) {
                strategyConfigs[i].tokenAddress = address(
                    new ERC20PresetFixedSupply("TestToken", "TEST", uint256(type(uint128).max), executorMultisig)
                );
            }
            deployedStrategyArray.push(
                StrategyBaseTVLLimits(
                    address(
                        new TransparentUpgradeableProxy(
                            address(baseStrategyImplementation),
                            address(eigenLayerProxyAdmin),
                            abi.encodeWithSelector(
                                StrategyBaseTVLLimits.initialize.selector,
                                strategyConfigs[i].maxPerDeposit,
                                strategyConfigs[i].maxDeposits,
                                IERC20(strategyConfigs[i].tokenAddress)
                            )
                        )
                    )
                )
            );
        }

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
            rewardsCoordinatorImplementation
        );
        _verifyContractsPointAtOneAnother(delegation, strategyManager, eigenPodManager, rewardsCoordinator);
        _verifyImplementationsSetCorrectly();
        _verifyInitialOwners();
        _checkPauserInitializations();
        _verifyInitializationParams();

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_strategies = "strategies";
        for (uint256 i = 0; i < strategyConfigs.length; ++i) {
            vm.serializeAddress(deployed_strategies, strategyConfigs[i].tokenSymbol, address(deployedStrategyArray[i]));
        }
        string memory deployed_strategies_output = strategyConfigs.length == 0
            ? ""
            : vm.serializeAddress(
                deployed_strategies,
                strategyConfigs[strategyConfigs.length - 1].tokenSymbol,
                address(deployedStrategyArray[strategyConfigs.length - 1])
            );

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
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodManagerImplementation", address(eigenPodManagerImplementation));
        vm.serializeAddress(deployed_addresses, "rewardsCoordinator", address(rewardsCoordinator));
        vm.serializeAddress(
            deployed_addresses, "rewardsCoordinatorImplementation", address(rewardsCoordinatorImplementation)
        );
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "eigenPodImplementation", address(eigenPodImplementation));
        vm.serializeAddress(deployed_addresses, "baseStrategyImplementation", address(baseStrategyImplementation));
        vm.serializeAddress(deployed_addresses, "emptyContract", address(emptyContract));

        vm.serializeAddress(deployed_addresses, "strategy", address(deployedStrategyArray[0]));
        vm.serializeAddress(deployed_addresses, "TestToken", address(strategyConfigs[0].tokenAddress));

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
        vm.writeJson(finalJson, "script/output/devnet/SLASHING_deploy_from_scratch_deployment_data.json");
    }

    function _verifyContractsPointAtOneAnother(
        DelegationManager delegationContract,
        StrategyManager strategyManagerContract,
        EigenPodManager eigenPodManagerContract,
        RewardsCoordinator rewardsCoordinatorContract
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

        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                eigenLayerProxyAdmin.getProxyImplementation(
                    ITransparentUpgradeableProxy(payable(address(deployedStrategyArray[i])))
                ) == address(baseStrategyImplementation),
                "strategy: implementation set incorrectly"
            );
        }

        require(
            eigenPodBeacon.implementation() == address(eigenPodImplementation),
            "eigenPodBeacon: implementation set incorrectly"
        );
    }

    function _verifyInitialOwners() internal view {
        require(strategyManager.owner() == executorMultisig, "strategyManager: owner not set correctly");
        require(delegation.owner() == executorMultisig, "delegation: owner not set correctly");
        require(eigenPodManager.owner() == executorMultisig, "eigenPodManager: owner not set correctly");

        require(eigenLayerProxyAdmin.owner() == executorMultisig, "eigenLayerProxyAdmin: owner not set correctly");
        require(eigenPodBeacon.owner() == executorMultisig, "eigenPodBeacon: owner not set correctly");
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
            strategyManager.strategyWhitelister() == operationsMultisig,
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

        string memory config_data = vm.readFile(deployConfigPath);
        for (uint256 i = 0; i < deployedStrategyArray.length; i++) {
            uint256 maxPerDeposit =
                stdJson.readUint(config_data, string.concat(".strategies[", vm.toString(i), "].max_per_deposit"));
            uint256 maxDeposits =
                stdJson.readUint(config_data, string.concat(".strategies[", vm.toString(i), "].max_deposits"));
            (uint256 setMaxPerDeposit, uint256 setMaxDeposits) = deployedStrategyArray[i].getTVLLimits();
            require(setMaxPerDeposit == maxPerDeposit, "setMaxPerDeposit not set correctly");
            require(setMaxDeposits == maxDeposits, "setMaxDeposits not set correctly");
        }
    }
}
