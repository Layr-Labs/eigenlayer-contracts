// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;
// solhint-disable no-console

import {Script} from "forge-std/Script.sol";
import {VmSafe} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {console2} from "forge-std/console2.sol";

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/Slasher.sol";
import "../../../src/contracts/core/DelegationManager.sol";
import "../../../src/contracts/core/AVSDirectory.sol";
import "../../../src/contracts/core/RewardsCoordinator.sol";

import "../../../src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";

import "../../../src/contracts/permissions/PauserRegistry.sol";

import "../../../src/test/mocks/EmptyContract.sol";
import "../../../src/test/mocks/ETHDepositMock.sol";

import {ContractDeployment, DeploymentDetails} from "./DeploymentDetails.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
contract Deployer_M2 is Script {
    address public constant ETH_DEPOSIT_ADDRESS_MAINNET = 0x00000000219ab540356cBB839Cbe05303d7705Fa;

    // struct used to encode token info in config file
    struct StrategyConfig {
        uint256 maxDeposits;
        uint256 maxPerDeposit;
        address tokenAddress;
        string tokenSymbol;
    }

    // EigenLayer Contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    Slasher public slasher;
    Slasher public slasherImplementation;
    DelegationManager public delegationManager;
    DelegationManager public delegationManagerImplementation;
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

    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;

    // the ETH2 deposit contract -- if not on mainnet, we deploy a mock as stand-in
    IETHPOSDeposit public ethPOSDeposit;

    // strategies deployed
    StrategyBaseTVLLimits[] public deployedStrategyArray;

    // IMMUTABLES TO SET
    uint64 GOERLI_GENESIS_TIME = 1616508000;

    address ETH_DEPOSIT_ADDRESS = ETH_DEPOSIT_ADDRESS_MAINNET;

    // OTHER DEPLOYMENT PARAMETERS
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    uint256 SLASHER_INIT_PAUSED_STATUS;
    uint256 DELEGATION_INIT_PAUSED_STATUS;
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;

    // RewardsCoordinator
    uint32 REWARDS_COORDINATOR_MAX_REWARDS_DURATION;
    uint32 REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH;
    uint32 REWARDS_COORDINATOR_MAX_FUTURE_LENGTH;
    uint32 REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
    address REWARDS_COORDINATOR_UPDATER;
    uint32 REWARDS_COORDINATOR_ACTIVATION_DELAY;
    uint32 REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS;
    uint32 REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS;

    // one week in blocks -- 50400
    uint32 STRATEGY_MANAGER_INIT_WITHDRAWAL_DELAY_BLOCKS;
    uint256 DELEGATION_WITHDRAWAL_DELAY_BLOCKS;

    function run(string memory configFileName) external {
        // read and log the chainID
        console2.log("You are deploying on ChainID", block.chainid);
        console2.log("account is ", msg.sender);

        string memory configData = _readConfig(configFileName);

        // tokens to deploy strategies for
        StrategyConfig[] memory strategyConfigs;
        // load token list
        bytes memory strategyConfigsRaw = stdJson.parseRaw(configData, ".strategies");
        strategyConfigs = abi.decode(strategyConfigsRaw, (StrategyConfig[]));

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

        address nilContract = address(new EmptyContract());

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        delegationManager = DelegationManager(
            address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), ""))
        );
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), ""))
        );
        slasher = Slasher(address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), "")));
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), ""))
        );
        rewardsCoordinator = RewardsCoordinator(
            address(new TransparentUpgradeableProxy(nilContract, address(eigenLayerProxyAdmin), ""))
        );

        // if on mainnet, use the ETH2 deposit contract address
        if (block.chainid == 1) {
            ethPOSDeposit = IETHPOSDeposit(ETH_DEPOSIT_ADDRESS_MAINNET);
            ETH_DEPOSIT_ADDRESS = ETH_DEPOSIT_ADDRESS_MAINNET;
            // if not on mainnet, deploy a mock
        } else {
            ethPOSDeposit = IETHPOSDeposit(ETH_DEPOSIT_ADDRESS);
        }
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GOERLI_GENESIS_TIME);

        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenPodManager, slasher);
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
        slasherImplementation = new Slasher(strategyManager, delegationManager);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        {
            IStrategy[] memory _strategies;
            uint256[] memory _withdrawalDelayBlocks;
            eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(delegationManager))),
                address(delegationManagerImplementation),
                abi.encodeWithSelector(
                    DelegationManager.initialize.selector,
                    executorMultisig,
                    eigenLayerPauserReg,
                    DELEGATION_INIT_PAUSED_STATUS,
                    DELEGATION_WITHDRAWAL_DELAY_BLOCKS,
                    _strategies,
                    _withdrawalDelayBlocks
                )
            );
            delegationManager.transferOwnership(executorMultisig);
        }
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                executorMultisig,
                operationsMultisig,
                eigenLayerPauserReg,
                STRATEGY_MANAGER_INIT_PAUSED_STATUS
            )
        );
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
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(AVSDirectory.initialize.selector, executorMultisig, eigenLayerPauserReg, 0)
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                executorMultisig,
                eigenLayerPauserReg,
                EIGENPOD_MANAGER_INIT_PAUSED_STATUS
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
            address(rewardsCoordinatorImplementation),
            abi.encodeWithSelector(
                RewardsCoordinator.initialize.selector,
                executorMultisig,
                eigenLayerPauserReg,
                REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
                REWARDS_COORDINATOR_UPDATER,
                REWARDS_COORDINATOR_ACTIVATION_DELAY,
                REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS
            )
        );

        // deploy StrategyBaseTVLLimits contract implementation
        baseStrategyImplementation = new StrategyBaseTVLLimits(strategyManager);
        // create upgradeable proxies that each point to the implementation and initialize them
        for (uint256 i = 0; i < strategyConfigs.length; ++i) {
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
                                IERC20(strategyConfigs[i].tokenAddress),
                                eigenLayerPauserReg
                            )
                        )
                    )
                )
            );
        }

        eigenLayerProxyAdmin.transferOwnership(executorMultisig);
        eigenPodBeacon.transferOwnership(executorMultisig);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // CHECK CORRECTNESS OF DEPLOYMENT
        _verifyContractsPointAtOneAnother(
            delegationManagerImplementation,
            strategyManagerImplementation,
            slasherImplementation,
            eigenPodManagerImplementation,
            rewardsCoordinatorImplementation
        );
        _verifyContractsPointAtOneAnother(
            delegationManager,
            strategyManager,
            slasher,
            eigenPodManager,
            rewardsCoordinator
        );
        _verifyImplementationsSetCorrectly();
        _verifyInitialOwners();
        _checkPauserInitializations();
        _verifyInitializationParams(configData);

        _writeJsonDeployment(nilContract);
    }

    function upgrade(string memory configFileName, string memory deploymentDetailsFilename) external {
        // read and log the chainID
        console2.log("You are upgrading ChainID", block.chainid);
        console2.log("account is ", msg.sender);

        string memory configData = _readConfig(configFileName);

        ContractDeployment memory contractDeployment = _readJsonDeployment(deploymentDetailsFilename);
        // tokens to deploy strategies for
        StrategyConfig[] memory strategyConfigs;

        vm.startBroadcast();
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GOERLI_GENESIS_TIME);

        // if on mainnet, use the ETH2 deposit contract address
        if (block.chainid == 1) {
            ethPOSDeposit = IETHPOSDeposit(ETH_DEPOSIT_ADDRESS_MAINNET);
            ETH_DEPOSIT_ADDRESS = ETH_DEPOSIT_ADDRESS_MAINNET;
            // if not on mainnet, deploy a mock
        } else {
            ethPOSDeposit = IETHPOSDeposit(ETH_DEPOSIT_ADDRESS);
        }
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GOERLI_GENESIS_TIME);

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenPodManager, slasher);
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
        slasherImplementation = new Slasher(strategyManager, delegationManager);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );

        // Second, upgrade the proxy contracts to point to the implementations

        console2.log("ProxyAdmin is owned by", eigenLayerProxyAdmin.owner());

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
        // AVSDirectory, upgrade and initalized
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation)
        );

        // Create base strategy implementation and deploy a few strategies
        baseStrategyImplementation = new StrategyBase(strategyManager);

        // Upgrade All deployed strategy contracts to new base strategy
        for (uint i = 0; i < strategyConfigs.length; i++) {
            // Upgrade existing strategy
            eigenLayerProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(deployedStrategyArray[i]))),
                address(baseStrategyImplementation)
            );
        }

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        _writeJsonDeployment(contractDeployment.emptyContract);
    }

    function _verifyContractsPointAtOneAnother(
        DelegationManager delegationContract,
        StrategyManager strategyManagerContract,
        Slasher /*slasherContract*/,
        EigenPodManager eigenPodManagerContract,
        RewardsCoordinator rewardsCoordinatorContract
    ) internal view {
        require(delegationContract.slasher() == slasher, "delegation: slasher address not set correctly");
        require(
            delegationContract.strategyManager() == strategyManager,
            "delegation: strategyManager address not set correctly"
        );

        require(strategyManagerContract.slasher() == slasher, "strategyManager: slasher address not set correctly");
        require(
            strategyManagerContract.delegation() == delegationManager,
            "strategyManager: delegation address not set correctly"
        );
        require(
            strategyManagerContract.eigenPodManager() == eigenPodManager,
            "strategyManager: eigenPodManager address not set correctly"
        );

        // removing slasher requirements because there is no slasher as part of m2-mainnet release
        // require(slasherContract.strategyManager() == strategyManager, "slasher: strategyManager not set correctly");
        // require(slasherContract.delegation() == delegation, "slasher: delegation not set correctly");

        require(
            eigenPodManagerContract.ethPOS() == ethPOSDeposit,
            " eigenPodManager: ethPOSDeposit contract address not set correctly"
        );
        require(
            eigenPodManagerContract.eigenPodBeacon() == eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        require(
            eigenPodManagerContract.strategyManager() == strategyManager,
            "eigenPodManager: strategyManager contract address not set correctly"
        );
        require(
            eigenPodManagerContract.slasher() == slasher,
            "eigenPodManager: slasher contract address not set correctly"
        );

        require(
            rewardsCoordinatorContract.delegationManager() == delegationManager,
            "rewardsCoordinator: delegation address not set correctly"
        );

        require(
            rewardsCoordinatorContract.strategyManager() == strategyManager,
            "rewardsCoordinator: strategyManager address not set correctly"
        );
    }

    function _verifyImplementationsSetCorrectly() internal view {
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(delegationManager)))
            ) == address(delegationManagerImplementation),
            "delegation: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(strategyManager)))
            ) == address(strategyManagerImplementation),
            "strategyManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(slasher)))) ==
                address(slasherImplementation),
            "slasher: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(eigenPodManager)))
            ) == address(eigenPodManagerImplementation),
            "eigenPodManager: implementation set incorrectly"
        );
        require(
            eigenLayerProxyAdmin.getProxyImplementation(
                TransparentUpgradeableProxy(payable(address(rewardsCoordinator)))
            ) == address(rewardsCoordinatorImplementation),
            "rewardsCoordinator: implementation set incorrectly"
        );

        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                eigenLayerProxyAdmin.getProxyImplementation(
                    TransparentUpgradeableProxy(payable(address(deployedStrategyArray[i])))
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
        require(delegationManager.owner() == executorMultisig, "delegation: owner not set correctly");
        // removing slasher requirements because there is no slasher as part of m2-mainnet release
        // require(slasher.owner() == executorMultisig, "slasher: owner not set correctly");
        require(eigenPodManager.owner() == executorMultisig, "eigenPodManager: owner not set correctly");

        require(eigenLayerProxyAdmin.owner() == executorMultisig, "eigenLayerProxyAdmin: owner not set correctly");
        require(eigenPodBeacon.owner() == executorMultisig, "eigenPodBeacon: owner not set correctly");
    }

    function _checkPauserInitializations() internal view {
        require(
            delegationManager.pauserRegistry() == eigenLayerPauserReg,
            "delegation: pauser registry not set correctly"
        );
        require(
            strategyManager.pauserRegistry() == eigenLayerPauserReg,
            "strategyManager: pauser registry not set correctly"
        );
        // removing slasher requirements because there is no slasher as part of m2-mainnet release
        // require(slasher.pauserRegistry() == eigenLayerPauserReg, "slasher: pauser registry not set correctly");
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
            require(
                deployedStrategyArray[i].paused() == 0,
                "StrategyBaseTVLLimits: init paused status set incorrectly"
            );
        }

        // // pause *nothing*
        // uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS = 0;
        // // pause *everything*
        // uint256 SLASHER_INIT_PAUSED_STATUS = type(uint256).max;
        // // pause *everything*
        // uint256 DELEGATION_INIT_PAUSED_STATUS = type(uint256).max;
        // // pause *all of the proof-related functionality* (everything that can be paused other than creation of EigenPods)
        // uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS = (2**1) + (2**2) + (2**3) + (2**4); /* = 30 */
        // // pause *nothing*
        // require(strategyManager.paused() == 0, "strategyManager: init paused status set incorrectly");
        // require(slasher.paused() == type(uint256).max, "slasher: init paused status set incorrectly");
        // require(delegation.paused() == type(uint256).max, "delegation: init paused status set incorrectly");
        // require(eigenPodManager.paused() == 30, "eigenPodManager: init paused status set incorrectly");
    }

    function _verifyInitializationParams(string memory configData) internal {
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

        for (uint i = 0; i < deployedStrategyArray.length; i++) {
            uint256 maxPerDeposit = stdJson.readUint(
                configData,
                string.concat(".strategies[", vm.toString(i), "].max_per_deposit")
            );
            uint256 maxDeposits = stdJson.readUint(
                configData,
                string.concat(".strategies[", vm.toString(i), "].max_deposits")
            );
            (uint256 setMaxPerDeposit, uint256 setMaxDeposits) = deployedStrategyArray[i].getTVLLimits();
            require(setMaxPerDeposit == maxPerDeposit, "setMaxPerDeposit not set correctly");
            require(setMaxDeposits == maxDeposits, "setMaxDeposits not set correctly");
        }
    }

    function _readConfig(string memory configFileName) internal returns (string memory) {
        string memory deployConfigPath = string(bytes(string.concat("script/configs/devnet/", configFileName)));
        string memory configData = vm.readFile(deployConfigPath);

        // READ JSON CONFIG DATA
        // bytes memory parsedData = vm.parseJson(configData);
        STRATEGY_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(configData, ".strategyManager.init_paused_status");
        SLASHER_INIT_PAUSED_STATUS = stdJson.readUint(configData, ".slasher.init_paused_status");
        DELEGATION_INIT_PAUSED_STATUS = stdJson.readUint(configData, ".delegation.init_paused_status");
        DELEGATION_WITHDRAWAL_DELAY_BLOCKS = stdJson.readUint(configData, ".delegation.init_withdrawal_delay_blocks");
        EIGENPOD_MANAGER_INIT_PAUSED_STATUS = stdJson.readUint(configData, ".eigenPodManager.init_paused_status");
        REWARDS_COORDINATOR_INIT_PAUSED_STATUS = stdJson.readUint(configData, ".rewardsCoordinator.init_paused_status");
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.CALCULATION_INTERVAL_SECONDS")
        );
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.MAX_REWARDS_DURATION")
        );
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.MAX_RETROACTIVE_LENGTH")
        );
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.MAX_FUTURE_LENGTH")
        );
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP")
        );
        REWARDS_COORDINATOR_UPDATER = stdJson.readAddress(configData, ".rewardsCoordinator.rewards_updater_address");
        REWARDS_COORDINATOR_ACTIVATION_DELAY = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.activation_delay")
        );
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.calculation_interval_seconds")
        );
        REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS = uint32(
            stdJson.readUint(configData, ".rewardsCoordinator.global_operator_commission_bips")
        );

        STRATEGY_MANAGER_INIT_WITHDRAWAL_DELAY_BLOCKS = uint32(
            stdJson.readUint(configData, ".strategyManager.init_withdrawal_delay_blocks")
        );

        ETH_DEPOSIT_ADDRESS = stdJson.readAddress(configData, ".ethPOSDepositAddress");

        executorMultisig = stdJson.readAddress(configData, ".multisig_addresses.executorMultisig");
        operationsMultisig = stdJson.readAddress(configData, ".multisig_addresses.operationsMultisig");
        pauserMultisig = stdJson.readAddress(configData, ".multisig_addresses.pauserMultisig");

        require(executorMultisig != address(0), "executorMultisig address not configured correctly!");
        require(operationsMultisig != address(0), "operationsMultisig address not configured correctly!");

        return configData;
    }

    function _writeJsonDeployment(address nilContract) internal {
        ContractDeployment memory contractDeployment = ContractDeployment({
            avsDirectory: address(avsDirectory),
            avsDirectoryImplementation: address(avsDirectoryImplementation),
            baseStrategyImplementation: address(baseStrategyImplementation),
            delegationManager: address(delegationManager),
            delegationManagerImplementation: address(delegationManagerImplementation),
            eigenPodBeacon: address(eigenPodBeacon),
            eigenPodImplementation: address(eigenPodImplementation),
            eigenPodManager: address(eigenPodManager),
            eigenPodManagerImplementation: address(eigenPodManagerImplementation),
            emptyContract: nilContract,
            pauserRegistry: address(eigenLayerPauserReg),
            proxyAdmin: address(eigenLayerProxyAdmin),
            rewardsCoordinator: address(rewardsCoordinator),
            rewardsCoordinatorImplementation: address(rewardsCoordinatorImplementation),
            slasher: address(slasher),
            slasherImplementation: address(slasherImplementation),
            strategyManager: address(strategyManager),
            strategyManagerImplementation: address(strategyManagerImplementation),
            ethDepositAddress: ETH_DEPOSIT_ADDRESS
        });
        DeploymentDetails.write(contractDeployment, executorMultisig, operationsMultisig, nilContract);
    }

    function _readJsonDeployment(string memory deploymentDetailsFilename) internal returns (ContractDeployment memory) {
        ContractDeployment memory contractDeployment = DeploymentDetails.read(deploymentDetailsFilename);
        console2.log("proxyAdmin", contractDeployment.proxyAdmin);
        eigenLayerProxyAdmin = ProxyAdmin(contractDeployment.proxyAdmin);
        eigenLayerPauserReg = PauserRegistry(contractDeployment.pauserRegistry);
        slasher = Slasher(contractDeployment.slasher);
        slasherImplementation = Slasher(contractDeployment.slasherImplementation);
        delegationManager = DelegationManager(contractDeployment.delegationManager);
        delegationManagerImplementation = DelegationManager(contractDeployment.delegationManagerImplementation);
        strategyManager = StrategyManager(contractDeployment.strategyManager);
        strategyManagerImplementation = StrategyManager(contractDeployment.strategyManagerImplementation);
        rewardsCoordinator = RewardsCoordinator(contractDeployment.rewardsCoordinator);
        rewardsCoordinatorImplementation = RewardsCoordinator(contractDeployment.rewardsCoordinatorImplementation);
        avsDirectory = AVSDirectory(contractDeployment.avsDirectory);
        avsDirectoryImplementation = AVSDirectory(contractDeployment.avsDirectoryImplementation);
        eigenPodManager = EigenPodManager(contractDeployment.eigenPodManager);
        eigenPodManagerImplementation = EigenPodManager(contractDeployment.eigenPodManagerImplementation);
        eigenPodBeacon = UpgradeableBeacon(contractDeployment.eigenPodBeacon);
        eigenPodImplementation = EigenPod(payable(contractDeployment.eigenPodImplementation));
        baseStrategyImplementation = StrategyBase(contractDeployment.baseStrategyImplementation);
        ETH_DEPOSIT_ADDRESS = contractDeployment.ethDepositAddress;
    }
}
