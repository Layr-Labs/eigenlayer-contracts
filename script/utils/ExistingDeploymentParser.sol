// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/RewardsCoordinator.sol";
import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/permissions/PermissionController.sol";

import "../../src/contracts/strategies/StrategyFactory.sol";
import "../../src/contracts/strategies/StrategyBase.sol";
import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "../../src/contracts/strategies/EigenStrategy.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/test/mocks/EmptyContract.sol";

import "../../src/contracts/interfaces/IBackingEigen.sol";
import "../../src/contracts/interfaces/IEigen.sol";

import "forge-std/Script.sol";

import "src/test/utils/Logger.t.sol";

struct StrategyUnderlyingTokenConfig {
    address tokenAddress;
    string tokenName;
    string tokenSymbol;
}

struct DeployedEigenPods {
    address[] multiValidatorPods;
    address[] singleValidatorPods;
    address[] inActivePods;
}

contract ExistingDeploymentParser is Script, Logger {
    using stdJson for string;

    /// -----------------------------------------------------------------------
    /// EigenLayer Contract Parameters
    /// -----------------------------------------------------------------------

    string public SEMVER;
    /// @dev AllocationManager
    uint256 ALLOCATION_MANAGER_INIT_PAUSED_STATUS;
    uint32 DEALLOCATION_DELAY;
    uint32 ALLOCATION_CONFIGURATION_DELAY;

    /// @dev AVSDirectory
    uint256 AVS_DIRECTORY_INIT_PAUSED_STATUS;

    /// @dev DelegationManager
    uint256 DELEGATION_MANAGER_INIT_PAUSED_STATUS;
    uint32 DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS;

    /// @dev EigenPodManager
    uint256 EIGENPOD_MANAGER_INIT_PAUSED_STATUS;

    /// @dev EigenPod
    uint64 EIGENPOD_GENESIS_TIME;
    uint64 EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR;
    address ETHPOSDepositAddress;

    /// @dev RewardsCoordinator
    uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;
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

    /// @dev StrategyManager
    uint256 STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    address STRATEGY_MANAGER_WHITELISTER;

    /// @dev Strategy Deployment
    uint256 STRATEGY_MAX_PER_DEPOSIT;
    uint256 STRATEGY_MAX_TOTAL_DEPOSITS;

    /// -----------------------------------------------------------------------
    /// EigenLayer Contracts
    /// -----------------------------------------------------------------------

    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    UpgradeableBeacon public eigenPodBeacon;
    UpgradeableBeacon public strategyBeacon;

    /// @dev AllocationManager
    AllocationManager public allocationManager;
    AllocationManager public allocationManagerImplementation;

    /// @dev AVSDirectory
    AVSDirectory public avsDirectory;
    AVSDirectory public avsDirectoryImplementation;

    /// @dev DelegationManager
    DelegationManager public delegationManager;
    DelegationManager public delegationManagerImplementation;

    /// @dev EigenPods
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    EigenPod public eigenPodImplementation;

    /// @dev PermissionController
    PermissionController public permissionController;
    PermissionController public permissionControllerImplementation;

    /// @dev RewardsCoordinator
    RewardsCoordinator public rewardsCoordinator;
    RewardsCoordinator public rewardsCoordinatorImplementation;

    /// @dev StrategyManager
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;

    /// @dev StrategyFactory
    StrategyFactory public strategyFactory;
    StrategyFactory public strategyFactoryImplementation;
    StrategyBase public baseStrategyImplementation;
    StrategyBase public strategyFactoryBeaconImplementation;

    // Token
    ProxyAdmin public tokenProxyAdmin;
    IEigen public EIGEN;
    IEigen public EIGENImpl;
    IBackingEigen public bEIGEN;
    IBackingEigen public bEIGENImpl;
    EigenStrategy public eigenStrategy;
    EigenStrategy public eigenStrategyImpl;

    /// -----------------------------------------------------------------------
    /// Storage
    /// -----------------------------------------------------------------------

    // EigenPods deployed
    address[] public multiValidatorPods;
    address[] public singleValidatorPods;
    address[] public inActivePods;
    // All eigenpods is just single array list of above eigenPods
    address[] public allEigenPods;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;
    address communityMultisig;
    address pauserMultisig;
    address timelock;

    // strategies deployed
    uint256 numStrategiesDeployed;
    StrategyBase[] public deployedStrategyArray;
    // Strategies to Deploy
    uint256 numStrategiesToDeploy;
    StrategyUnderlyingTokenConfig[] public strategiesToDeploy;

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    function NAME() public view virtual override returns (string memory) {
        return "ExistingDeploymentParser";
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @notice use for parsing already deployed EigenLayer contracts
    function _parseDeployedContracts(
        string memory existingDeploymentInfoPath
    ) internal virtual noTracing {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        console.log("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory json = cheats.readFile(existingDeploymentInfoPath);

        SEMVER = stdJson.readString(json, ".parameters.semver");

        // check that the chainID matches the one in the config
        uint256 configChainId = json.readUint(".chainInfo.chainId");
        assertEq(configChainId, currentChainId, "You are on the wrong chain for this config");

        console.log("Using addresses file", existingDeploymentInfoPath);
        console.log("- Last Updated", stdJson.readString(json, ".lastUpdated"));

        // read all of the deployed addresses
        executorMultisig = json.readAddress(".parameters.executorMultisig");
        operationsMultisig = json.readAddress(".parameters.operationsMultisig");
        communityMultisig = json.readAddress(".parameters.communityMultisig");
        pauserMultisig = json.readAddress(".parameters.pauserMultisig");
        timelock = json.readAddress(".parameters.timelock");

        eigenLayerProxyAdmin = ProxyAdmin(json.readAddress(".addresses.eigenLayerProxyAdmin"));
        eigenLayerPauserReg = PauserRegistry(json.readAddress(".addresses.eigenLayerPauserReg"));

        // FIXME: hotfix - remove later...
        permissionControllerImplementation = new PermissionController(SEMVER);
        permissionController = PermissionController(
            address(
                new TransparentUpgradeableProxy(
                    address(permissionControllerImplementation), address(eigenLayerProxyAdmin), ""
                )
            )
        );

        allocationManagerImplementation = new AllocationManager(
            delegationManager,
            eigenLayerPauserReg,
            permissionController,
            DEALLOCATION_DELAY,
            ALLOCATION_CONFIGURATION_DELAY,
            SEMVER
        );
        allocationManager = AllocationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(allocationManagerImplementation), address(eigenLayerProxyAdmin), ""
                )
            )
        );

        // // AllocationManager
        // allocationManager = AllocationManager(json.readAddress(".addresses.allocationManager"));
        // allocationManagerImplementation = json.readAddress(".addresses.allocationManagerImplementation");

        // AVSDirectory
        avsDirectory = AVSDirectory(json.readAddress(".addresses.avsDirectory"));
        avsDirectoryImplementation = AVSDirectory(json.readAddress(".addresses.avsDirectoryImplementation"));

        // DelegationManager
        delegationManager = DelegationManager(json.readAddress(".addresses.delegationManager"));
        delegationManagerImplementation =
            DelegationManager(json.readAddress(".addresses.delegationManagerImplementation"));

        // // PermissionController
        // permissionController = PermissionController(json.readAddress(".addresses.permissionController"));
        // permissionControllerImplementation = json.readAddress(".addresses.permissionControllerImplementation");

        // RewardsCoordinator
        rewardsCoordinator = RewardsCoordinator(json.readAddress(".addresses.rewardsCoordinator"));
        rewardsCoordinatorImplementation =
            RewardsCoordinator(json.readAddress(".addresses.rewardsCoordinatorImplementation"));

        // StrategyManager
        strategyManager = StrategyManager(json.readAddress(".addresses.strategyManager"));
        strategyManagerImplementation = StrategyManager(json.readAddress(".addresses.strategyManagerImplementation"));

        // StrategyFactory
        strategyFactory = StrategyFactory(json.readAddress(".addresses.strategyFactory"));
        strategyFactoryImplementation = StrategyFactory(json.readAddress(".addresses.strategyFactoryImplementation"));

        // StrategyBeacon
        strategyBeacon = UpgradeableBeacon(json.readAddress(".addresses.strategyFactoryBeacon"));
        strategyFactoryBeaconImplementation =
            StrategyBase(json.readAddress(".addresses.strategyFactoryBeaconImplementation"));
        baseStrategyImplementation = StrategyBase(json.readAddress(".addresses.baseStrategyImplementation"));

        // EigenPodManager
        eigenPodManager = EigenPodManager(json.readAddress(".addresses.eigenPodManager"));
        eigenPodManagerImplementation = EigenPodManager(json.readAddress(".addresses.eigenPodManagerImplementation"));

        // EigenPod
        eigenPodBeacon = UpgradeableBeacon(json.readAddress(".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(payable(json.readAddress(".addresses.eigenPodImplementation")));

        emptyContract = EmptyContract(json.readAddress(".addresses.emptyContract"));

        // Strategies Deployed, load strategy list
        numStrategiesDeployed = json.readUint(".addresses.numStrategiesDeployed");
        for (uint256 i = 0; i < numStrategiesDeployed; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".addresses.strategyAddresses[", cheats.toString(i), "]");
            // Use the key and parse the strategy address
            address strategyAddress = abi.decode(json.parseRaw(key), (address));
            deployedStrategyArray.push(StrategyBase(strategyAddress));
        }

        // token
        tokenProxyAdmin = ProxyAdmin(json.readAddress(".addresses.token.tokenProxyAdmin"));
        EIGEN = IEigen(json.readAddress(".addresses.token.EIGEN"));
        EIGENImpl = IEigen(json.readAddress(".addresses.token.EIGENImpl"));
        bEIGEN = IBackingEigen(json.readAddress(".addresses.token.bEIGEN"));
        bEIGENImpl = IBackingEigen(json.readAddress(".addresses.token.bEIGENImpl"));
        eigenStrategy = EigenStrategy(json.readAddress(".addresses.token.eigenStrategy"));
        eigenStrategyImpl = EigenStrategy(json.readAddress(".addresses.token.eigenStrategyImpl"));
    }

    function _parseDeployedEigenPods(
        string memory existingDeploymentInfoPath
    ) internal returns (DeployedEigenPods memory) {
        uint256 currentChainId = block.chainid;

        // READ JSON CONFIG DATA
        string memory json = cheats.readFile(existingDeploymentInfoPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = json.readUint(".chainInfo.chainId");
        assertEq(configChainId, currentChainId, "You are on the wrong chain for this config");

        multiValidatorPods = json.readAddressArray(".eigenPods.multiValidatorPods");
        singleValidatorPods = json.readAddressArray(".eigenPods.singleValidatorPods");
        inActivePods = json.readAddressArray(".eigenPods.inActivePods");
        allEigenPods = json.readAddressArray(".eigenPods.allEigenPods");
        return DeployedEigenPods({
            multiValidatorPods: multiValidatorPods,
            singleValidatorPods: singleValidatorPods,
            inActivePods: inActivePods
        });
    }

    /// @notice use for deploying a new set of EigenLayer contracts
    /// Note that this does assertEq multisigs to already be deployed
    function _parseInitialDeploymentParams(
        string memory initialDeploymentParamsPath
    ) internal virtual {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        console.log("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory json = cheats.readFile(initialDeploymentParamsPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = json.readUint(".chainInfo.chainId");
        assertEq(configChainId, currentChainId, "You are on the wrong chain for this config");

        console.log("Using config file", initialDeploymentParamsPath);
        console.log("- Last Updated", stdJson.readString(json, ".lastUpdated"));

        // read all of the deployed addresses
        executorMultisig = json.readAddress(".multisig_addresses.executorMultisig");
        operationsMultisig = json.readAddress(".multisig_addresses.operationsMultisig");
        communityMultisig = json.readAddress(".multisig_addresses.communityMultisig");
        pauserMultisig = json.readAddress(".multisig_addresses.pauserMultisig");

        // Strategies to Deploy, load strategy list
        numStrategiesToDeploy = json.readUint(".strategies.numStrategies");
        STRATEGY_MAX_PER_DEPOSIT = json.readUint(".strategies.MAX_PER_DEPOSIT");
        STRATEGY_MAX_TOTAL_DEPOSITS = json.readUint(".strategies.MAX_TOTAL_DEPOSITS");
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            // Form the key for the current element
            string memory key = string.concat(".strategies.strategiesToDeploy[", cheats.toString(i), "]");

            // Use parseJson with the key to get the value for the current element
            bytes memory tokenInfoBytes = stdJson.parseRaw(json, key);

            // Decode the token information into the Token struct
            StrategyUnderlyingTokenConfig memory tokenInfo = abi.decode(tokenInfoBytes, (StrategyUnderlyingTokenConfig));

            strategiesToDeploy.push(tokenInfo);
        }

        // Read initialize params for upgradeable contracts
        STRATEGY_MANAGER_INIT_PAUSED_STATUS = json.readUint(".strategyManager.init_paused_status");
        STRATEGY_MANAGER_WHITELISTER = json.readAddress(".strategyManager.init_strategy_whitelister");
        // DelegationManager
        DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS =
            uint32(json.readUint(".delegationManager.init_minWithdrawalDelayBlocks"));
        DELEGATION_MANAGER_INIT_PAUSED_STATUS = json.readUint(".delegationManager.init_paused_status");
        // RewardsCoordinator

        REWARDS_COORDINATOR_INIT_PAUSED_STATUS = json.readUint(".rewardsCoordinator.init_paused_status");
        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS =
            uint32(json.readUint(".rewardsCoordinator.CALCULATION_INTERVAL_SECONDS"));
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION = uint32(json.readUint(".rewardsCoordinator.MAX_REWARDS_DURATION"));
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = uint32(json.readUint(".rewardsCoordinator.MAX_RETROACTIVE_LENGTH"));
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = uint32(json.readUint(".rewardsCoordinator.MAX_FUTURE_LENGTH"));
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP =
            uint32(json.readUint(".rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP"));
        REWARDS_COORDINATOR_UPDATER = json.readAddress(".rewardsCoordinator.rewards_updater_address");
        REWARDS_COORDINATOR_ACTIVATION_DELAY = uint32(json.readUint(".rewardsCoordinator.activation_delay"));
        REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS =
            uint32(json.readUint(".rewardsCoordinator.default_operator_split_bips"));
        REWARDS_COORDINATOR_OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP =
            uint32(json.readUint(".rewardsCoordinator.OPERATOR_SET_GENESIS_REWARDS_TIMESTAMP"));
        REWARDS_COORDINATOR_OPERATOR_SET_MAX_RETROACTIVE_LENGTH =
            uint32(json.readUint(".rewardsCoordinator.OPERATOR_SET_MAX_RETROACTIVE_LENGTH"));
        // AVSDirectory
        AVS_DIRECTORY_INIT_PAUSED_STATUS = json.readUint(".avsDirectory.init_paused_status");
        // EigenPodManager
        EIGENPOD_MANAGER_INIT_PAUSED_STATUS = json.readUint(".eigenPodManager.init_paused_status");
        // AllocationManager
        ALLOCATION_MANAGER_INIT_PAUSED_STATUS = json.readUint(".allocationManager.init_paused_status");
        // EigenPod
        EIGENPOD_GENESIS_TIME = uint64(json.readUint(".eigenPod.GENESIS_TIME"));
        ETHPOSDepositAddress = json.readAddress(".ethPOSDepositAddress");

        // check that all values are non-zero
        logInitialDeploymentParams();
    }

    /// @notice Ensure contracts point at each other correctly via constructors
    function _verifyContractPointers() internal view virtual {
        // AVSDirectory
        assertTrue(
            avsDirectory.delegation() == delegationManager, "avsDirectory: delegationManager address not set correctly"
        );
        // RewardsCoordinator
        assertTrue(
            rewardsCoordinator.delegationManager() == delegationManager,
            "rewardsCoordinator: delegationManager address not set correctly"
        );
        assertTrue(
            rewardsCoordinator.strategyManager() == strategyManager,
            "rewardsCoordinator: strategyManager address not set correctly"
        );
        // DelegationManager
        assertTrue(
            delegationManager.strategyManager() == strategyManager,
            "delegationManager: strategyManager address not set correctly"
        );
        assertTrue(
            delegationManager.eigenPodManager() == eigenPodManager,
            "delegationManager: eigenPodManager address not set correctly"
        );
        // StrategyManager
        assertTrue(
            strategyManager.delegation() == delegationManager,
            "strategyManager: delegationManager address not set correctly"
        );
        // EPM
        assertTrue(
            address(eigenPodManager.ethPOS()) == ETHPOSDepositAddress,
            "eigenPodManager: ethPOSDeposit contract address not set correctly"
        );
        assertTrue(
            eigenPodManager.eigenPodBeacon() == eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        assertTrue(
            eigenPodManager.delegationManager() == delegationManager,
            "eigenPodManager: delegationManager contract address not set correctly"
        );
    }

    /// @notice verify implementations for Transparent Upgradeable Proxies
    function _verifyImplementations() internal view virtual {
        assertEq(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(avsDirectory)))),
            address(avsDirectoryImplementation),
            "avsDirectory: implementation set incorrectly"
        );
        assertEq(
            eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(rewardsCoordinator)))
            ),
            address(rewardsCoordinatorImplementation),
            "rewardsCoordinator: implementation set incorrectly"
        );
        assertEq(
            eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(delegationManager)))
            ),
            address(delegationManagerImplementation),
            "delegationManager: implementation set incorrectly"
        );
        assertEq(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyManager)))),
            address(strategyManagerImplementation),
            "strategyManager: implementation set incorrectly"
        );
        assertEq(
            eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(eigenPodManager)))),
            address(eigenPodManagerImplementation),
            "eigenPodManager: implementation set incorrectly"
        );

        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            assertEq(
                eigenLayerProxyAdmin.getProxyImplementation(
                    ITransparentUpgradeableProxy(payable(address(deployedStrategyArray[i])))
                ),
                address(baseStrategyImplementation),
                "strategy: implementation set incorrectly"
            );
        }

        assertEq(
            eigenPodBeacon.implementation(),
            address(eigenPodImplementation),
            "eigenPodBeacon: implementation set incorrectly"
        );
    }

    /**
     * @notice Verify initialization of Transparent Upgradeable Proxies. Also check
     * initialization params if this is the first deployment.
     * @dev isInitialDeployment True if this is the first deployment of contracts from scratch
     */
    function _verifyContractsInitialized(
        bool /* isInitialDeployment */
    ) internal virtual {
        // AVSDirectory
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        avsDirectory.initialize(address(0), AVS_DIRECTORY_INIT_PAUSED_STATUS);
        // RewardsCoordinator
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        rewardsCoordinator.initialize(
            address(0),
            0, // initialPausedStatus
            address(0), // rewardsUpdater
            0, // activationDelay
            0 // defaultSplitBips
        );
        // DelegationManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delegationManager.initialize(address(0), 0);
        // StrategyManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategyManager.initialize(address(0), address(0), STRATEGY_MANAGER_INIT_PAUSED_STATUS);
        // EigenPodManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        eigenPodManager.initialize(address(0), EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
        // Strategies
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            cheats.expectRevert(bytes("Initializable: contract is already initialized"));
            StrategyBaseTVLLimits(address(deployedStrategyArray[i])).initialize(0, 0, IERC20(address(0)));
        }
    }

    /// @notice Verify params based on config constants that are updated from calling `_parseInitialDeploymentParams`
    function _verifyInitializationParams() internal view virtual {
        // AVSDirectory
        assertTrue(
            avsDirectory.pauserRegistry() == eigenLayerPauserReg, "avsdirectory: pauser registry not set correctly"
        );
        assertEq(avsDirectory.owner(), executorMultisig, "avsdirectory: owner not set correctly");
        assertEq(
            avsDirectory.paused(), AVS_DIRECTORY_INIT_PAUSED_STATUS, "avsdirectory: init paused status set incorrectly"
        );
        // RewardsCoordinator
        assertTrue(
            rewardsCoordinator.pauserRegistry() == eigenLayerPauserReg,
            "rewardsCoordinator: pauser registry not set correctly"
        );
        // assertEq(
        //     rewardsCoordinator.owner(), executorMultisig,
        //     "rewardsCoordinator: owner not set correctly"
        // );
        // assertEq(
        //     rewardsCoordinator.paused(), REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
        //     "rewardsCoordinator: init paused status set incorrectly"
        // );
        assertEq(
            rewardsCoordinator.MAX_REWARDS_DURATION(),
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            "rewardsCoordinator: maxRewardsDuration not set correctly"
        );
        assertEq(
            rewardsCoordinator.MAX_RETROACTIVE_LENGTH(),
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            "rewardsCoordinator: maxRetroactiveLength not set correctly"
        );
        assertEq(
            rewardsCoordinator.MAX_FUTURE_LENGTH(),
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            "rewardsCoordinator: maxFutureLength not set correctly"
        );
        assertEq(
            rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP(),
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP,
            "rewardsCoordinator: genesisRewardsTimestamp not set correctly"
        );
        // assertEq(
        //     rewardsCoordinator.rewardsUpdater(), REWARDS_COORDINATOR_UPDATER,
        //     "rewardsCoordinator: rewardsUpdater not set correctly"
        // );
        assertEq(
            rewardsCoordinator.activationDelay(),
            REWARDS_COORDINATOR_ACTIVATION_DELAY,
            "rewardsCoordinator: activationDelay not set correctly"
        );
        assertEq(
            rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(),
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            "rewardsCoordinator: CALCULATION_INTERVAL_SECONDS not set correctly"
        );
        assertEq(
            rewardsCoordinator.defaultOperatorSplitBips(),
            REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS,
            "rewardsCoordinator: defaultSplitBips not set correctly"
        );
        // DelegationManager
        assertTrue(
            delegationManager.pauserRegistry() == eigenLayerPauserReg,
            "delegationManager: pauser registry not set correctly"
        );
        assertEq(delegationManager.owner(), executorMultisig, "delegationManager: owner not set correctly");
        assertEq(
            delegationManager.paused(),
            DELEGATION_MANAGER_INIT_PAUSED_STATUS,
            "delegationManager: init paused status set incorrectly"
        );
        // StrategyManager
        assertTrue(
            strategyManager.pauserRegistry() == eigenLayerPauserReg,
            "strategyManager: pauser registry not set correctly"
        );
        assertEq(strategyManager.owner(), executorMultisig, "strategyManager: owner not set correctly");
        assertEq(
            strategyManager.paused(),
            STRATEGY_MANAGER_INIT_PAUSED_STATUS,
            "strategyManager: init paused status set incorrectly"
        );
        if (block.chainid == 1) {
            assertEq(
                strategyManager.strategyWhitelister(),
                address(strategyFactory),
                "strategyManager: strategyWhitelister not set correctly"
            );
        } else if (block.chainid == 17_000) {
            // On holesky, for ease of whitelisting we set to executorMultisig
            // assertEq(
            //     strategyManager.strategyWhitelister(), executorMultisig,
            //     "strategyManager: strategyWhitelister not set correctly"
            // );
        }
        // EigenPodManager
        assertTrue(
            eigenPodManager.pauserRegistry() == eigenLayerPauserReg,
            "eigenPodManager: pauser registry not set correctly"
        );
        assertEq(eigenPodManager.owner(), executorMultisig, "eigenPodManager: owner not set correctly");
        assertEq(
            eigenPodManager.paused(),
            EIGENPOD_MANAGER_INIT_PAUSED_STATUS,
            "eigenPodManager: init paused status set incorrectly"
        );
        assertEq(
            address(eigenPodManager.ethPOS()),
            address(ETHPOSDepositAddress),
            "eigenPodManager: ethPOS not set correctly"
        );
        // EigenPodBeacon
        assertEq(eigenPodBeacon.owner(), executorMultisig, "eigenPodBeacon: owner not set correctly");
        // EigenPodImplementation
        assertEq(
            eigenPodImplementation.GENESIS_TIME(),
            EIGENPOD_GENESIS_TIME,
            "eigenPodImplementation: GENESIS TIME not set correctly"
        );
        assertEq(
            address(eigenPodImplementation.ethPOS()),
            ETHPOSDepositAddress,
            "eigenPodImplementation: ethPOS not set correctly"
        );
        // Strategies
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            assertTrue(
                deployedStrategyArray[i].pauserRegistry() == eigenLayerPauserReg,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            assertEq(deployedStrategyArray[i].paused(), 0, "StrategyBaseTVLLimits: init paused status set incorrectly");
            assertTrue(
                strategyManager.strategyIsWhitelistedForDeposit(deployedStrategyArray[i]),
                "StrategyBaseTVLLimits: strategy should be whitelisted"
            );
        }

        // Pausing Permissions
        assertTrue(eigenLayerPauserReg.isPauser(operationsMultisig), "pauserRegistry: operationsMultisig is not pauser");
        assertTrue(eigenLayerPauserReg.isPauser(executorMultisig), "pauserRegistry: executorMultisig is not pauser");
        assertTrue(eigenLayerPauserReg.isPauser(pauserMultisig), "pauserRegistry: pauserMultisig is not pauser");
        assertEq(eigenLayerPauserReg.unpauser(), executorMultisig, "pauserRegistry: unpauser not set correctly");
    }

    function logInitialDeploymentParams() public {
        console.log("==== Parsed Initialize Params for Initial Deployment,==");

        console.log("executorMultisig", executorMultisig);
        console.log("operationsMultisig", operationsMultisig);
        console.log("communityMultisig", communityMultisig);
        console.log("pauserMultisig", pauserMultisig);

        console.log("STRATEGY_MANAGER_INIT_PAUSED_STATUS", STRATEGY_MANAGER_INIT_PAUSED_STATUS);
        console.log("STRATEGY_MANAGER_WHITELISTER", STRATEGY_MANAGER_WHITELISTER);
        console.log("DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS", DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        console.log("DELEGATION_MANAGER_INIT_PAUSED_STATUS", DELEGATION_MANAGER_INIT_PAUSED_STATUS);
        console.log("AVS_DIRECTORY_INIT_PAUSED_STATUS", AVS_DIRECTORY_INIT_PAUSED_STATUS);
        console.log("REWARDS_COORDINATOR_INIT_PAUSED_STATUS", REWARDS_COORDINATOR_INIT_PAUSED_STATUS);
        // todo log all rewards coordinator params
        console.log("EIGENPOD_MANAGER_INIT_PAUSED_STATUS", EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
        console.log("EIGENPOD_GENESIS_TIME", EIGENPOD_GENESIS_TIME);
        console.log("ETHPOSDepositAddress", ETHPOSDepositAddress);

        console.log("==== Strategies to Deploy,==");
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            // Decode the token information into the Token struct
            StrategyUnderlyingTokenConfig memory tokenInfo = strategiesToDeploy[i];

            strategiesToDeploy.push(tokenInfo);
            console.log("TOKEN ADDRESS", tokenInfo.tokenAddress);
            console.log("TOKEN NAME", tokenInfo.tokenName);
            console.log("TOKEN SYMBOL", tokenInfo.tokenSymbol);
        }
    }

    /**
     * @notice Log contract addresses and write to output json file
     */
    function logAndOutputContractAddresses(
        string memory outputPath
    ) public {
        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_strategies = "strategies";
        for (uint256 i = 0; i < numStrategiesToDeploy; ++i) {
            deployed_strategies.serialize(strategiesToDeploy[i].tokenSymbol, address(deployedStrategyArray[i]));
        }
        string memory deployed_strategies_output = numStrategiesToDeploy == 0
            ? ""
            : deployed_strategies.serialize(
                strategiesToDeploy[numStrategiesToDeploy - 1].tokenSymbol,
                address(deployedStrategyArray[numStrategiesToDeploy - 1])
            );

        string memory deployed_addresses = "addresses";
        deployed_addresses.serialize("eigenLayerProxyAdmin", address(eigenLayerProxyAdmin));
        deployed_addresses.serialize("eigenLayerPauserReg", address(eigenLayerPauserReg));
        deployed_addresses.serialize("avsDirectory", address(avsDirectory));
        deployed_addresses.serialize("avsDirectoryImplementation", address(avsDirectoryImplementation));
        deployed_addresses.serialize("delegationManager", address(delegationManager));
        deployed_addresses.serialize("delegationManagerImplementation", address(delegationManagerImplementation));
        deployed_addresses.serialize("strategyManager", address(strategyManager));
        deployed_addresses.serialize("strategyManagerImplementation", address(strategyManagerImplementation));
        deployed_addresses.serialize("rewardsCoordinator", address(rewardsCoordinator));
        deployed_addresses.serialize("rewardsCoordinatorImplementation", address(rewardsCoordinatorImplementation));
        deployed_addresses.serialize("eigenPodManager", address(eigenPodManager));
        deployed_addresses.serialize("eigenPodManagerImplementation", address(eigenPodManagerImplementation));
        deployed_addresses.serialize("eigenPodBeacon", address(eigenPodBeacon));
        deployed_addresses.serialize("eigenPodImplementation", address(eigenPodImplementation));
        deployed_addresses.serialize("baseStrategyImplementation", address(baseStrategyImplementation));
        deployed_addresses.serialize("emptyContract", address(emptyContract));
        string memory deployed_addresses_output = deployed_addresses.serialize("strategies", deployed_strategies_output);

        string memory parameters = "parameters";
        parameters.serialize("executorMultisig", executorMultisig);
        parameters.serialize("operationsMultisig", operationsMultisig);
        parameters.serialize("communityMultisig", communityMultisig);
        parameters.serialize("pauserMultisig", pauserMultisig);
        parameters.serialize("timelock", timelock);
        string memory parameters_output = parameters.serialize("operationsMultisig", operationsMultisig);

        string memory chain_info = "chainInfo";
        chain_info.serialize("deploymentBlock", block.number);
        string memory chain_info_output = chain_info.serialize("chainId", block.chainid);

        // serialize all the data
        parent_object.serialize(deployed_addresses, deployed_addresses_output);
        parent_object.serialize(chain_info, chain_info_output);
        string memory finalJson = parent_object.serialize(parameters, parameters_output);

        cheats.writeJson(finalJson, outputPath);
    }

    /// @notice used for parsing parameters used in the integration test upgrade
    function _parseParamsForIntegrationUpgrade(
        string memory initialDeploymentParamsPath
    ) internal virtual noTracing {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        console.log("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory json = cheats.readFile(initialDeploymentParamsPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = json.readUint(".config.environment.chainid");
        assertEq(configChainId, currentChainId, "You are on the wrong chain for this config");

        console.log("Using config file", initialDeploymentParamsPath);
        console.log("- Last Updated", stdJson.readString(json, ".config.environment.lastUpdated"));

        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS =
            uint32(json.readUint(".config.params.CALCULATION_INTERVAL_SECONDS"));
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION = uint32(json.readUint(".config.params.MAX_REWARDS_DURATION"));
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = uint32(json.readUint(".config.params.MAX_RETROACTIVE_LENGTH"));
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = uint32(json.readUint(".config.params.MAX_FUTURE_LENGTH"));
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP =
            uint32(json.readUint(".config.params.GENESIS_REWARDS_TIMESTAMP"));

        DEALLOCATION_DELAY = uint32(json.readUint(".config.params.MIN_WITHDRAWAL_DELAY_BLOCKS"));
        ALLOCATION_CONFIGURATION_DELAY = uint32(json.readUint(".config.params.ALLOCATION_CONFIGURATION_DELAY"));
        DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS =
            uint32(json.readUint(".config.params.MIN_WITHDRAWAL_DELAY_BLOCKS"));
    }
}
