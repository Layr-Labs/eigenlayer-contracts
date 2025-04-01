// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "src/test/utils/Logger.t.sol";
import "./IDeploymentParser.sol";

contract JsonDeploymentParser is IDeploymentParser, Script, Logger {
    using stdJson for string;

    string private jsonPath;
    string private SEMVER;

    // Core contracts
    AllocationManager private _allocationManager;
    AVSDirectory private _avsDirectory;
    DelegationManager private _delegationManager;
    RewardsCoordinator private _rewardsCoordinator;
    StrategyManager private _strategyManager;
    EigenPodManager private _eigenPodManager;
    PermissionController private _permissionController;
    PauserRegistry private _pauserRegistry;

    // Strategy contracts
    StrategyFactory private _strategyFactory;
    EigenStrategy private _eigenStrategy;
    StrategyBase private _strategyBase;
    StrategyBase[] private _deployedStrategyArray;

    // Token contracts
    IEigen private _EIGEN;
    IBackingEigen private _bEIGEN;

    // Admin contracts
    ProxyAdmin private _eigenLayerProxyAdmin;
    ProxyAdmin private _tokenProxyAdmin;
    UpgradeableBeacon private _eigenPodBeacon;

    // Implementation contracts
    AllocationManager private _allocationManagerImplementation;
    AVSDirectory private _avsDirectoryImplementation;
    DelegationManager private _delegationManagerImplementation;
    RewardsCoordinator private _rewardsCoordinatorImplementation;
    StrategyManager private _strategyManagerImplementation;
    EigenPodManager private _eigenPodManagerImplementation;
    PermissionController private _permissionControllerImplementation;
    StrategyFactory private _strategyFactoryImplementation;
    EigenStrategy private _eigenStrategyImplementation;
    StrategyBase private _baseStrategyImplementation;
    EigenPod private _eigenPodImplementation;
    IEigen private _EIGENImplementation;
    IBackingEigen private _bEIGENImplementation;

    // Multisig addresses
    address private _executorMultisig;
    address private _operationsMultisig;
    address private _communityMultisig;
    address private _pauserMultisig;
    address private _timelock;

    // Parameters
    uint256 private _STRATEGY_MANAGER_INIT_PAUSED_STATUS;
    uint256 private _DELEGATION_MANAGER_INIT_PAUSED_STATUS;
    uint256 private _AVS_DIRECTORY_INIT_PAUSED_STATUS;
    uint256 private _REWARDS_COORDINATOR_INIT_PAUSED_STATUS;
    uint256 private _EIGENPOD_MANAGER_INIT_PAUSED_STATUS;
    uint256 private _ALLOCATION_MANAGER_INIT_PAUSED_STATUS;

    constructor(
        string memory _jsonPath
    ) {
        jsonPath = _jsonPath;
    }

    function NAME() public view virtual override returns (string memory) {
        return "JsonDeploymentParser";
    }

    function initialize() external {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        console.log("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory json = cheats.readFile(jsonPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = json.readUint(".chainInfo.chainId");
        assertEq(configChainId, currentChainId, "You are on the wrong chain for this config");

        console.log("Using addresses file", jsonPath);
        console.log("- Last Updated", stdJson.readString(json, ".lastUpdated"));

        SEMVER = stdJson.readString(json, ".parameters.semver");

        // read all of the deployed addresses
        _executorMultisig = json.readAddress(".parameters.executorMultisig");
        _operationsMultisig = json.readAddress(".parameters.operationsMultisig");
        _communityMultisig = json.readAddress(".parameters.communityMultisig");
        _pauserMultisig = json.readAddress(".parameters.pauserMultisig");
        _timelock = json.readAddress(".parameters.timelock");

        _eigenLayerProxyAdmin = ProxyAdmin(json.readAddress(".addresses.eigenLayerProxyAdmin"));
        _pauserRegistry = PauserRegistry(json.readAddress(".addresses.eigenLayerPauserReg"));

        // Core contracts
        _allocationManager = AllocationManager(json.readAddress(".addresses.allocationManager"));
        _allocationManagerImplementation =
            AllocationManager(json.readAddress(".addresses.allocationManagerImplementation"));
        _avsDirectory = AVSDirectory(json.readAddress(".addresses.avsDirectory"));
        _avsDirectoryImplementation = AVSDirectory(json.readAddress(".addresses.avsDirectoryImplementation"));
        _delegationManager = DelegationManager(json.readAddress(".addresses.delegationManager"));
        _delegationManagerImplementation =
            DelegationManager(json.readAddress(".addresses.delegationManagerImplementation"));
        _rewardsCoordinator = RewardsCoordinator(json.readAddress(".addresses.rewardsCoordinator"));
        _rewardsCoordinatorImplementation =
            RewardsCoordinator(json.readAddress(".addresses.rewardsCoordinatorImplementation"));
        _strategyManager = StrategyManager(json.readAddress(".addresses.strategyManager"));
        _strategyManagerImplementation = StrategyManager(json.readAddress(".addresses.strategyManagerImplementation"));
        _eigenPodManager = EigenPodManager(json.readAddress(".addresses.eigenPodManager"));
        _eigenPodManagerImplementation = EigenPodManager(json.readAddress(".addresses.eigenPodManagerImplementation"));
        _permissionController = PermissionController(json.readAddress(".addresses.permissionController"));
        _permissionControllerImplementation =
            PermissionController(json.readAddress(".addresses.permissionControllerImplementation"));

        // Strategy contracts
        _strategyFactory = StrategyFactory(json.readAddress(".addresses.strategyFactory"));
        _strategyFactoryImplementation = StrategyFactory(json.readAddress(".addresses.strategyFactoryImplementation"));
        _strategyBase = StrategyBase(json.readAddress(".addresses.baseStrategyImplementation"));
        _eigenStrategy = EigenStrategy(json.readAddress(".addresses.token.eigenStrategy"));
        _eigenStrategyImplementation = EigenStrategy(json.readAddress(".addresses.token.eigenStrategyImpl"));

        // Token contracts
        _EIGEN = IEigen(json.readAddress(".addresses.token.EIGEN"));
        _EIGENImplementation = IEigen(json.readAddress(".addresses.token.EIGENImpl"));
        _bEIGEN = IBackingEigen(json.readAddress(".addresses.token.bEIGEN"));
        _bEIGENImplementation = IBackingEigen(json.readAddress(".addresses.token.bEIGENImpl"));

        // Admin contracts
        _tokenProxyAdmin = ProxyAdmin(json.readAddress(".addresses.token.tokenProxyAdmin"));
        _eigenPodBeacon = UpgradeableBeacon(json.readAddress(".addresses.eigenPodBeacon"));
        _eigenPodImplementation = EigenPod(payable(json.readAddress(".addresses.eigenPodImplementation")));

        // Strategies Deployed
        uint256 numStrategiesDeployed = json.readUint(".addresses.numStrategiesDeployed");
        for (uint256 i = 0; i < numStrategiesDeployed; ++i) {
            string memory key = string.concat(".addresses.strategyAddresses[", cheats.toString(i), "]");
            address strategyAddress = abi.decode(json.parseRaw(key), (address));
            _deployedStrategyArray.push(StrategyBase(strategyAddress));
        }

        // Read initialize params
        _STRATEGY_MANAGER_INIT_PAUSED_STATUS = json.readUint(".strategyManager.init_paused_status");
        _DELEGATION_MANAGER_INIT_PAUSED_STATUS = json.readUint(".delegationManager.init_paused_status");
        _AVS_DIRECTORY_INIT_PAUSED_STATUS = json.readUint(".avsDirectory.init_paused_status");
        _REWARDS_COORDINATOR_INIT_PAUSED_STATUS = json.readUint(".rewardsCoordinator.init_paused_status");
        _EIGENPOD_MANAGER_INIT_PAUSED_STATUS = json.readUint(".eigenPodManager.init_paused_status");
        _ALLOCATION_MANAGER_INIT_PAUSED_STATUS = json.readUint(".allocationManager.init_paused_status");
    }

    // Core contracts
    function allocationManager() external view returns (AllocationManager) {
        return _allocationManager;
    }

    function avsDirectory() external view returns (AVSDirectory) {
        return _avsDirectory;
    }

    function delegationManager() external view returns (DelegationManager) {
        return _delegationManager;
    }

    function rewardsCoordinator() external view returns (RewardsCoordinator) {
        return _rewardsCoordinator;
    }

    function strategyManager() external view returns (StrategyManager) {
        return _strategyManager;
    }

    function eigenPodManager() external view returns (EigenPodManager) {
        return _eigenPodManager;
    }

    function permissionController() external view returns (PermissionController) {
        return _permissionController;
    }

    function pauserRegistry() external view returns (PauserRegistry) {
        return _pauserRegistry;
    }

    // Strategy contracts
    function strategyFactory() external view returns (StrategyFactory) {
        return _strategyFactory;
    }

    function eigenStrategy() external view returns (EigenStrategy) {
        return _eigenStrategy;
    }

    function strategyBase() external view returns (StrategyBase) {
        return _strategyBase;
    }

    function deployedStrategyArray(
        uint256 index
    ) external view returns (StrategyBase) {
        return _deployedStrategyArray[index];
    }

    // Token contracts
    function EIGEN() external view returns (IEigen) {
        return _EIGEN;
    }

    function bEIGEN() external view returns (IBackingEigen) {
        return _bEIGEN;
    }

    // Admin contracts
    function eigenLayerProxyAdmin() external view returns (ProxyAdmin) {
        return _eigenLayerProxyAdmin;
    }

    function tokenProxyAdmin() external view returns (ProxyAdmin) {
        return _tokenProxyAdmin;
    }

    function eigenPodBeacon() external view returns (UpgradeableBeacon) {
        return _eigenPodBeacon;
    }

    // Implementation contracts
    function allocationManagerImplementation() external view returns (AllocationManager) {
        return _allocationManagerImplementation;
    }

    function avsDirectoryImplementation() external view returns (AVSDirectory) {
        return _avsDirectoryImplementation;
    }

    function delegationManagerImplementation() external view returns (DelegationManager) {
        return _delegationManagerImplementation;
    }

    function rewardsCoordinatorImplementation() external view returns (RewardsCoordinator) {
        return _rewardsCoordinatorImplementation;
    }

    function strategyManagerImplementation() external view returns (StrategyManager) {
        return _strategyManagerImplementation;
    }

    function eigenPodManagerImplementation() external view returns (EigenPodManager) {
        return _eigenPodManagerImplementation;
    }

    function permissionControllerImplementation() external view returns (PermissionController) {
        return _permissionControllerImplementation;
    }

    function strategyFactoryImplementation() external view returns (StrategyFactory) {
        return _strategyFactoryImplementation;
    }

    function eigenStrategyImplementation() external view returns (EigenStrategy) {
        return _eigenStrategyImplementation;
    }

    function baseStrategyImplementation() external view returns (StrategyBase) {
        return _strategyBase;
    }

    function eigenPodImplementation() external view returns (EigenPod) {
        return _eigenPodImplementation;
    }

    function EIGENImplementation() external view returns (IEigen) {
        return _EIGENImplementation;
    }

    function bEIGENImplementation() external view returns (IBackingEigen) {
        return _bEIGENImplementation;
    }

    // Multisig addresses
    function executorMultisig() external view returns (address) {
        return _executorMultisig;
    }

    function operationsMultisig() external view returns (address) {
        return _operationsMultisig;
    }

    function communityMultisig() external view returns (address) {
        return _communityMultisig;
    }

    function pauserMultisig() external view returns (address) {
        return _pauserMultisig;
    }

    function timelock() external view returns (address) {
        return _timelock;
    }

    // Verification functions
    function verifyContractPointers() external view {
        // AVSDirectory
        assertTrue(
            _avsDirectory.delegation() == _delegationManager,
            "avsDirectory: delegationManager address not set correctly"
        );
        // RewardsCoordinator
        assertTrue(
            _rewardsCoordinator.delegationManager() == _delegationManager,
            "rewardsCoordinator: delegationManager address not set correctly"
        );
        assertTrue(
            _rewardsCoordinator.strategyManager() == _strategyManager,
            "rewardsCoordinator: strategyManager address not set correctly"
        );
        // DelegationManager
        assertTrue(
            _delegationManager.strategyManager() == _strategyManager,
            "delegationManager: strategyManager address not set correctly"
        );
        assertTrue(
            _delegationManager.eigenPodManager() == _eigenPodManager,
            "delegationManager: eigenPodManager address not set correctly"
        );
        // StrategyManager
        assertTrue(
            _strategyManager.delegation() == _delegationManager,
            "strategyManager: delegationManager address not set correctly"
        );
        // EPM
        assertTrue(
            _eigenPodManager.eigenPodBeacon() == _eigenPodBeacon,
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        assertTrue(
            _eigenPodManager.delegationManager() == _delegationManager,
            "eigenPodManager: delegationManager contract address not set correctly"
        );
    }

    function verifyImplementations() external view {
        assertEq(
            _eigenLayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(_avsDirectory)))),
            address(_avsDirectoryImplementation),
            "avsDirectory: implementation set incorrectly"
        );
        assertEq(
            _eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(_rewardsCoordinator)))
            ),
            address(_rewardsCoordinatorImplementation),
            "rewardsCoordinator: implementation set incorrectly"
        );
        assertEq(
            _eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(_delegationManager)))
            ),
            address(_delegationManagerImplementation),
            "delegationManager: implementation set incorrectly"
        );
        assertEq(
            _eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(_strategyManager)))
            ),
            address(_strategyManagerImplementation),
            "strategyManager: implementation set incorrectly"
        );
        assertEq(
            _eigenLayerProxyAdmin.getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(_eigenPodManager)))
            ),
            address(_eigenPodManagerImplementation),
            "eigenPodManager: implementation set incorrectly"
        );

        for (uint256 i = 0; i < _deployedStrategyArray.length; ++i) {
            assertEq(
                _eigenLayerProxyAdmin.getProxyImplementation(
                    ITransparentUpgradeableProxy(payable(address(_deployedStrategyArray[i])))
                ),
                address(_strategyBase),
                "strategy: implementation set incorrectly"
            );
        }

        assertEq(
            _eigenPodBeacon.implementation(),
            address(_eigenPodImplementation),
            "eigenPodBeacon: implementation set incorrectly"
        );
    }

    function verifyContractsInitialized(
        bool isInitialDeployment
    ) external {
        // AVSDirectory
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        _avsDirectory.initialize(address(0), _AVS_DIRECTORY_INIT_PAUSED_STATUS);
        // RewardsCoordinator
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        _rewardsCoordinator.initialize(
            address(0),
            0, // initialPausedStatus
            address(0), // rewardsUpdater
            0, // activationDelay
            0 // defaultSplitBips
        );
        // DelegationManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        _delegationManager.initialize(address(0), 0);
        // StrategyManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        _strategyManager.initialize(address(0), address(0), _STRATEGY_MANAGER_INIT_PAUSED_STATUS);
        // EigenPodManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        _eigenPodManager.initialize(address(0), _EIGENPOD_MANAGER_INIT_PAUSED_STATUS);
        // Strategies
        for (uint256 i = 0; i < _deployedStrategyArray.length; ++i) {
            cheats.expectRevert(bytes("Initializable: contract is already initialized"));
            StrategyBaseTVLLimits(address(_deployedStrategyArray[i])).initialize(0, 0, IERC20(address(0)));
        }
    }

    function verifyInitializationParams() external view {
        // AVSDirectory
        assertTrue(_avsDirectory.pauserRegistry() == _pauserRegistry, "avsdirectory: pauser registry not set correctly");
        assertEq(_avsDirectory.owner(), _executorMultisig, "avsdirectory: owner not set correctly");
        assertEq(
            _avsDirectory.paused(),
            _AVS_DIRECTORY_INIT_PAUSED_STATUS,
            "avsdirectory: init paused status set incorrectly"
        );

        // RewardsCoordinator
        assertTrue(
            _rewardsCoordinator.pauserRegistry() == _pauserRegistry,
            "rewardsCoordinator: pauser registry not set correctly"
        );

        // DelegationManager
        assertTrue(
            _delegationManager.pauserRegistry() == _pauserRegistry,
            "delegationManager: pauser registry not set correctly"
        );
        assertEq(_delegationManager.owner(), _executorMultisig, "delegationManager: owner not set correctly");
        assertEq(
            _delegationManager.paused(),
            _DELEGATION_MANAGER_INIT_PAUSED_STATUS,
            "delegationManager: init paused status set incorrectly"
        );

        // StrategyManager
        assertTrue(
            _strategyManager.pauserRegistry() == _pauserRegistry, "strategyManager: pauser registry not set correctly"
        );
        assertEq(_strategyManager.owner(), _executorMultisig, "strategyManager: owner not set correctly");
        assertEq(
            _strategyManager.paused(),
            _STRATEGY_MANAGER_INIT_PAUSED_STATUS,
            "strategyManager: init paused status set incorrectly"
        );

        // EigenPodManager
        assertTrue(
            _eigenPodManager.pauserRegistry() == _pauserRegistry, "eigenPodManager: pauser registry not set correctly"
        );
        assertEq(_eigenPodManager.owner(), _executorMultisig, "eigenPodManager: owner not set correctly");
        assertEq(
            _eigenPodManager.paused(),
            _EIGENPOD_MANAGER_INIT_PAUSED_STATUS,
            "eigenPodManager: init paused status set incorrectly"
        );

        // EigenPodBeacon
        assertEq(_eigenPodBeacon.owner(), _executorMultisig, "eigenPodBeacon: owner not set correctly");

        // Strategies
        for (uint256 i = 0; i < _deployedStrategyArray.length; ++i) {
            assertTrue(
                _deployedStrategyArray[i].pauserRegistry() == _pauserRegistry,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            assertEq(_deployedStrategyArray[i].paused(), 0, "StrategyBaseTVLLimits: init paused status set incorrectly");
            assertTrue(
                _strategyManager.strategyIsWhitelistedForDeposit(_deployedStrategyArray[i]),
                "StrategyBaseTVLLimits: strategy should be whitelisted"
            );
        }

        // Pausing Permissions
        assertTrue(_pauserRegistry.isPauser(_operationsMultisig), "pauserRegistry: operationsMultisig is not pauser");
        assertTrue(_pauserRegistry.isPauser(_executorMultisig), "pauserRegistry: executorMultisig is not pauser");
        assertTrue(_pauserRegistry.isPauser(_pauserMultisig), "pauserRegistry: pauserMultisig is not pauser");
        assertEq(_pauserRegistry.unpauser(), _executorMultisig, "pauserRegistry: unpauser not set correctly");
    }
}
