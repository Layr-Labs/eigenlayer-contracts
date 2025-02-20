// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Imports
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/strategies/StrategyFactory.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/permissions/PermissionController.sol";

import "src/test/mocks/EmptyContract.sol";
import "src/test/mocks/ETHDepositMock.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";

import "src/test/integration/users/AVS.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";
import "src/test/integration/users/User_M2.t.sol";

import "script/utils/ExistingDeploymentParser.sol";

// DelegationManager
uint8 constant PAUSED_NEW_DELEGATION = 0;
uint8 constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;
uint8 constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;
// StrategyManager
uint8 constant PAUSED_DEPOSITS = 0;
// EigenpodManager
uint8 constant PAUSED_NEW_EIGENPODS = 0;
uint8 constant PAUSED_WITHDRAW_RESTAKED_ETH = 1;
uint8 constant PAUSED_EIGENPODS_VERIFY_CREDENTIALS = 2;
uint8 constant PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE = 3;
uint8 constant PAUSED_EIGENPODS_VERIFY_WITHDRAWAL = 4;
uint8 constant PAUSED_NON_PROOF_WITHDRAWALS = 5;

IStrategy constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

abstract contract IntegrationDeployer is ExistingDeploymentParser {
    using StdStyle for *;

    // Fork ids for specific fork tests
    bool isUpgraded;
    uint mainnetForkBlock = 20_847_130; // Post PI upgrade
    uint mainnetForkId;
    uint holeskyForkBLock = 1_213_950;
    uint holeskyForkId;
    uint64 constant DENEB_FORK_TIMESTAMP = 1_705_473_120;

    // Beacon chain genesis time when running locally
    // Multiple of 12 for sanity's sake
    uint64 constant GENESIS_TIME_LOCAL = 1 hours * 12;
    uint64 constant GENESIS_TIME_MAINNET = 1_606_824_023;

    uint8 constant NUM_LST_STRATS = 32;

    TimeMachine public timeMachine;

    // Lists of strategies used in the system
    //
    // When we select random user assets, we use the `assetType` to determine
    // which of these lists to select user assets from.
    IStrategy[] lstStrats;
    IStrategy[] ethStrats; // only has one strat tbh
    IStrategy[] allStrats; // just a combination of the above 2 lists
    IERC20[] allTokens; // `allStrats`, but contains all of the underlying tokens instead

    // If a token is in this mapping, then we will ignore this LST as it causes issues with reading balanceOf
    mapping(address => bool) public tokensNotTested;

    // Mock Contracts to deploy
    ETHPOSDepositMock ethPOSDeposit;
    BeaconChainMock public beaconChain;

    // Admin Addresses
    address eigenLayerReputedMultisig = address(this); // admin address
    address constant pauser = address(555);
    address constant unpauser = address(556);

    // Randomness state vars
    bytes32 random;
    // After calling `_configRand`, these are the allowed "variants" on users that will
    // be returned from `_randUser`.
    bytes assetTypes;
    bytes userTypes;
    // Set only once in setUp, if FORK_MAINNET env is set
    uint forkType;

    // Constants
    uint64 constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;

    constructor() {
        address stETH_Holesky = 0x3F1c547b21f65e10480dE3ad8E19fAAC46C95034;
        address stETH_Mainnet = 0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84;
        address OETH_Mainnet = 0x856c4Efb76C1D1AE02e20CEB03A2A6a08b0b8dC3;
        address osETH_Holesky = 0xF603c5A3F774F05d4D848A9bB139809790890864;
        address osETH_Mainnet = 0xf1C9acDc66974dFB6dEcB12aA385b9cD01190E38;
        address cbETH_Holesky = 0x8720095Fa5739Ab051799211B146a2EEE4Dd8B37;
        tokensNotTested[stETH_Holesky] = true;
        tokensNotTested[stETH_Mainnet] = true;
        tokensNotTested[OETH_Mainnet] = true;
        tokensNotTested[osETH_Holesky] = true;
        tokensNotTested[osETH_Mainnet] = true;
        tokensNotTested[cbETH_Holesky] = true;
    }

    function NAME() public view virtual override returns (string memory) {
        return "Integration Deployer";
    }

    /**
     * @dev Anyone who wants to test using this contract in a separate repo via submodules may have to
     * override this function to set the correct paths for the deployment info files.
     *
     * This setUp function will account for specific --fork-url flags and deploy/upgrade contracts accordingly.
     * Note that forkIds are also created so you can make explicit fork tests using cheats.selectFork(forkId)
     */
    function setUp() public virtual {
        isUpgraded = false;

        /**
         * env FOUNDRY_PROFILE=forktest forge t --mc Integration
         *
         * Running foundry like this will trigger the fork test profile,
         * lowering fuzz runs and using a remote RPC to test against mainnet state
         */
        bool forkMainnet = _hash("forktest") == _hash(cheats.envOr(string("FOUNDRY_PROFILE"), string("default")));

        if (forkMainnet) {
            console.log("Setting up `%s` integration tests:", "MAINNET_FORK".green().bold());
            console.log("RPC:", cheats.rpcUrl("mainnet"));
            console.log("Block:", mainnetForkBlock);

            cheats.createSelectFork(cheats.rpcUrl("mainnet"), mainnetForkBlock);
            forkType = MAINNET;
        } else {
            console.log("Setting up `%s` integration tests:", "LOCAL".yellow().bold());

            forkType = LOCAL;
        }

        _deployOrFetchContracts();
    }

    function _setUpLocal() public virtual {
        // Deploy ProxyAdmin
        eigenLayerProxyAdmin = new ProxyAdmin();
        executorMultisig = address(eigenLayerProxyAdmin.owner());

        // Deploy PauserRegistry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        eigenLayerPauserReg = new PauserRegistry(pausers, unpauser);

        // Deploy mocks
        EmptyContract emptyContract = new EmptyContract();
        ethPOSDeposit = new ETHPOSDepositMock();

        // Matching parameters to testnet
        DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS = 50;
        DEALLOCATION_DELAY = 50;
        ALLOCATION_CONFIGURATION_DELAY = 75;

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        delegationManager = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyFactory = StrategyFactory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        allocationManager = AllocationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        permissionController = PermissionController(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GENESIS_TIME_LOCAL);

        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));
        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, permissionController, DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            delegationManager,
            eigenLayerPauserReg
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, delegationManager, eigenLayerPauserReg);
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg);
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg);
        allocationManagerImplementation = new AllocationManager(delegationManager, eigenLayerPauserReg, permissionController, DEALLOCATION_DELAY, ALLOCATION_CONFIGURATION_DELAY);
        permissionControllerImplementation = new PermissionController();

        // Third, upgrade the proxy contracts to point to the implementations
        // DelegationManager
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                0 /* initialPausedStatus */
            )
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                eigenLayerReputedMultisig, //initialOwner
                eigenLayerReputedMultisig, //initial whitelister
                0 // initialPausedStatus
            )
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                0 // initialPausedStatus
            )
        );
        // AVSDirectory
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(
                AVSDirectory.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                0 // initialPausedStatus
            )
        );
        // AllocationManager
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(allocationManager))),
            address(allocationManagerImplementation),
            abi.encodeWithSelector(
                AllocationManager.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                0 // initialPausedStatus
            )
        );
        //PermissionController
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(permissionController))),
            address(permissionControllerImplementation)
        );
        // Create base strategy implementation and deploy a few strategies
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg);

        // Create a proxy beacon for base strategy implementation
        strategyBeacon = new UpgradeableBeacon(address(baseStrategyImplementation));

        // Strategy Factory, upgrade and initalized
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyFactory))),
            address(strategyFactoryImplementation),
            abi.encodeWithSelector(
                StrategyFactory.initialize.selector,
                eigenLayerReputedMultisig,
                0, // initial paused status
                IBeacon(strategyBeacon)
            )
        );

        cheats.prank(eigenLayerReputedMultisig);
        strategyManager.setStrategyWhitelister(address(strategyFactory));

        for (uint i = 1; i < NUM_LST_STRATS + 1; ++i) {
            string memory name = string.concat("LST-Strat", cheats.toString(i), " token");
            string memory symbol = string.concat("lstStrat", cheats.toString(i));
            // Deploy half of the strategies using the factory.
            _newStrategyAndToken(name, symbol, 10e50, address(this), i % 2 == 0);
        }

        ethStrats.push(BEACONCHAIN_ETH_STRAT);
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);

        // Create time machine and beacon chain. Set block time to beacon chain genesis time
        // TODO: update if needed to sane timestamp
        cheats.warp(GENESIS_TIME_LOCAL);
        timeMachine = new TimeMachine();
        beaconChain = new BeaconChainMock(eigenPodManager, GENESIS_TIME_LOCAL);
    }

    /**
     * @notice deploy current implementation contracts and upgrade the existing proxy EigenLayer contracts
     * on Mainnet. Setup for integration tests on mainnet fork.
     *
     * Note that beacon chain oracle and eth deposit contracts are mocked and pointed to different addresses for these tests.
     */
    function _upgradeMainnetContracts() public virtual {
        cheats.startPrank(address(executorMultisig));

        ethPOSDeposit = new ETHPOSDepositMock();
        ETHPOSDepositAddress = address(ethPOSDeposit); // overwrite for upgrade checks later

        // First, deploy the new contracts as empty contracts
        emptyContract = new EmptyContract();
        allocationManager = AllocationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        permissionController = PermissionController(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        allocationManagerImplementation = new AllocationManager(delegationManager, eigenLayerPauserReg, permissionController, DEALLOCATION_DELAY, ALLOCATION_CONFIGURATION_DELAY);
        permissionControllerImplementation = new PermissionController();
        delegationManagerImplementation = new DelegationManager(strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, permissionController, DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            allocationManager,
            eigenLayerPauserReg,
            permissionController,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            delegationManager,
            eigenLayerPauserReg
        );
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GENESIS_TIME_MAINNET);
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg);
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg);

        // Third, upgrade the proxy contracts to point to the implementations
        
        // Initialize the newly deployed contracts
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(allocationManager))),
            address(allocationManagerImplementation),
            abi.encodeWithSelector(
                AllocationManager.initialize.selector,
                executorMultisig,
                0 // initialPausedStatus
            )
        );
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(permissionController))),
            address(permissionControllerImplementation)
        );

        // DelegationManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation)
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(strategyManager))), address(strategyManagerImplementation)
        );
        // RewardsCoordinator
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(rewardsCoordinator))), address(rewardsCoordinatorImplementation)
        );
        // AVSDirectory 
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(avsDirectory))), address(avsDirectoryImplementation)
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(eigenPodManager))), address(eigenPodManagerImplementation)
        );
        // EigenPod
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
        // StrategyFactory
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(strategyFactory))), address(strategyFactoryImplementation)
        );
        // Strategy Beacon
        strategyBeacon.upgradeTo(address(baseStrategyImplementation));

        // Upgrade All deployed strategy contracts to new base strategy
        for (uint i = 0; i < numStrategiesDeployed; i++) {
            // Upgrade existing strategy
            eigenLayerProxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(deployedStrategyArray[i]))),
                address(baseStrategyImplementation)
            );
        }

        // Third, unpause core contracts
        delegationManager.unpause(0);
        eigenPodManager.unpause(0);
        strategyManager.unpause(0);

        cheats.stopPrank();

        ethStrats.push(BEACONCHAIN_ETH_STRAT);
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);
    }

    /**
     * @notice deploy current implementation contracts and upgrade the existing proxy EigenLayer contracts
     * on Holesky. Setup for integration tests on Holesky fork.
     *
     * Note that beacon chain oracle and eth deposit contracts are mocked and pointed to different addresses for these tests.
     */
    function _upgradeHoleskyContracts() public virtual {
        cheats.startPrank(address(executorMultisig));

        ethPOSDeposit = new ETHPOSDepositMock();
        ETHPOSDepositAddress = address(ethPOSDeposit); // overwrite for upgrade checks later

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, 0);
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
        // Deploy AVSDirectory, contract has not been deployed on mainnet yet
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // First, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, permissionController, DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS);
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            delegationManager,
            eigenLayerPauserReg
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, delegationManager, eigenLayerPauserReg);
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg);

        // Second, upgrade the proxy contracts to point to the implementations
        // DelegationManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation)
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(strategyManager))), address(strategyManagerImplementation)
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(eigenPodManager))), address(eigenPodManagerImplementation)
        );
        // AVSDirectory, upgrade and initalized
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(
                AVSDirectory.initialize.selector,
                executorMultisig,
                0 // initialPausedStatus
            )
        );

        // Create base strategy implementation and deploy a few strategies
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg);

        // Upgrade All deployed strategy contracts to new base strategy
        for (uint i = 0; i < numStrategiesDeployed; i++) {
            // Upgrade existing strategy
            eigenLayerProxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(deployedStrategyArray[i]))),
                address(baseStrategyImplementation)
            );
        }

        // Third, unpause core contracts
        delegationManager.unpause(0);
        eigenPodManager.unpause(0);
        strategyManager.unpause(0);

        cheats.stopPrank();

        ethStrats.push(BEACONCHAIN_ETH_STRAT);
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);
    }

    /// @dev Deploy a strategy and its underlying token, push to global lists of tokens/strategies, and whitelist
    /// strategy in strategyManager
    function _newStrategyAndToken(
        string memory tokenName,
        string memory tokenSymbol,
        uint initialSupply,
        address owner,
        bool useFactory
    ) internal {
        IERC20 underlyingToken = new ERC20PresetFixedSupply(tokenName, tokenSymbol, initialSupply, owner);

        StrategyBase strategy;

        if (useFactory) {
            strategy = StrategyBase(address(strategyFactory.deployNewStrategy(underlyingToken)));
        } else {
            strategy = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                    )
                )
            );
        }

        // Whitelist strategy
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;

        if (forkType == MAINNET) {
            cheats.prank(strategyManager.strategyWhitelister());
            IStrategyManager_DeprecatedM1(address(strategyManager)).addStrategiesToDepositWhitelist(strategies);
            cheats.prank(eigenLayerPauserReg.unpauser());
            StrategyBaseTVLLimits(address(strategy)).setTVLLimits(type(uint).max, type(uint).max);
        } else {
            cheats.prank(strategyManager.strategyWhitelister());
            strategyManager.addStrategiesToDepositWhitelist(strategies);
        }

        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
    }

    function _configRand(uint24 _randomSeed, uint _assetTypes, uint _userTypes) internal {
        // Using uint24 for the seed type so that if a test fails, it's easier
        // to manually use the seed to replay the same test.
        random = _hash(_randomSeed);
        
        // Convert flag bitmaps to bytes of set bits for easy use with _randUint
        assetTypes = _bitmapToBytes(_assetTypes);
        userTypes = _bitmapToBytes(_userTypes);

        assertTrue(assetTypes.length != 0, "_configRand: no asset types selected");
        assertTrue(userTypes.length != 0, "_configRand: no user types selected");
    }

    /**
     * Depending on the forkType, either deploy contracts locally or parse existing contracts
     * from network.
     *
     * Note: for non-LOCAL forktypes, upgrade of contracts will be peformed after user initialization.
     */
    function _deployOrFetchContracts() internal {
        if (forkType == LOCAL) {
            _setUpLocal();
            // Set Upgraded as local setup deploys most up to date contracts
            isUpgraded = true;
        } else if (forkType == MAINNET) {
            // cheats.selectFork(mainnetForkId);
            string memory deploymentInfoPath = "script/configs/mainnet/mainnet-addresses.config.json";
            _parseDeployedContracts(deploymentInfoPath);
            string memory existingDeploymentParams = "script/configs/mainnet.json";
            _parseParamsForIntegrationUpgrade(existingDeploymentParams);

            // Unpause to enable deposits and withdrawals for initializing random user state
            cheats.prank(eigenLayerPauserReg.unpauser());
            strategyManager.unpause(0);

            // Add deployed strategies to lstStrats and allStrats
            for (uint i; i < deployedStrategyArray.length; i++) {
                IStrategy strategy = IStrategy(deployedStrategyArray[i]);

                if (tokensNotTested[address(strategy.underlyingToken())]) {
                    continue;
                }

                // Add to lstStrats and allStrats
                lstStrats.push(strategy);
                allStrats.push(strategy);
                allTokens.push(strategy.underlyingToken());
            }

            // Create time machine and mock beacon chain
            timeMachine = new TimeMachine();
            beaconChain = new BeaconChainMock(eigenPodManager, GENESIS_TIME_MAINNET);
        } else if (forkType == HOLESKY) {
            revert("_deployOrFetchContracts - holesky tests currently broken sorry");
            // // cheats.selectFork(holeskyForkId);
            // string memory deploymentInfoPath = "script/configs/holesky/Holesky_current_deployment.config.json";
            // _parseDeployedContracts(deploymentInfoPath);

            // // Add deployed strategies to lstStrats and allStrats
            // for (uint i; i < deployedStrategyArray.length; i++) {
            //     IStrategy strategy = IStrategy(deployedStrategyArray[i]);

            //     if (tokensNotTested[address(strategy.underlyingToken())]) {
            //         continue;
            //     }

            //     // Add to lstStrats and allStrats
            //     lstStrats.push(strategy);
            //     allStrats.push(strategy);
            //     allTokens.push(strategy.underlyingToken());
            // }

            // // Update deposit contract to be a mock
            // ethPOSDeposit = new ETHPOSDepositMock();
            // eigenPodImplementation = new EigenPod(
            //     ethPOSDeposit,
            //     eigenPodImplementation.delayedWithdrawalRouter(),
            //     eigenPodImplementation.eigenPodManager(),
            //     eigenPodImplementation.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR(),
            //     0
            // );
            // // Create time machine and set block timestamp forward so we can create EigenPod proofs in the past
            // timeMachine = new TimeMachine();
            // beaconChainOracle = new BeaconChainOracleMock();
            // // Create mock beacon chain / proof gen interface
            // beaconChain = new BeaconChainMock(timeMachine, beaconChainOracle, eigenPodManager);

            // cheats.startPrank(executorMultisig);
            // eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
            // eigenPodManager.updateBeaconChainOracle(beaconChainOracle);
            // cheats.stopPrank();
        } else {
            revert("_deployOrFetchContracts: unimplemented forkType");
        }
    }

    /**
     * @dev Create a new User with a random config using the range defined in `_configRand`
     *
     * Assets are pulled from `strategies` based on a random staker/operator `assetType`
     */
    function _randUser(
        string memory name
    ) internal returns (User, IStrategy[] memory, uint[] memory) {
        // For the new user, select what type of assets they'll have and whether
        // they'll use `xWithSignature` methods.
        //
        // The values selected here are in the ranges configured via `_configRand`
        uint assetType = _randAssetType();
        uint userType = _randUserType();

        // Deploy new User contract
        User user = _genRandUser(name, userType);

        // For the specific asset selection we made, get a random assortment of
        // strategies and deal the user some corresponding underlying token balances
        (IStrategy[] memory strategies, uint[] memory tokenBalances) = _dealRandAssets(user, assetType);

        print.user(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
    }

    /// @dev Create a new user without native ETH. See _randUser above for standard usage
    function _randUser_NoETH(
        string memory name
    ) internal returns (User, IStrategy[] memory, uint[] memory) {
        // For the new user, select what type of assets they'll have and whether
        // they'll use `xWithSignature` methods.
        //
        // The values selected here are in the ranges configured via `_configRand`
        uint userType = _randUserType();

        // Pick the user's asset distribution, removing "native ETH" as an option
        // I'm sorry if this eventually leads to a bug that's really hard to track down
        uint assetType = _randAssetType();
        if (assetType == HOLDS_ETH) {
            assetType = NO_ASSETS;
        } else if (assetType == HOLDS_ALL) {
            assetType = HOLDS_LST;
        }

        // Deploy new User contract
        User user = _genRandUser(name, userType);

        // For the specific asset selection we made, get a random assortment of
        // strategies and deal the user some corresponding underlying token balances
        (IStrategy[] memory strategies, uint[] memory tokenBalances) = _dealRandAssets(user, assetType);

        print.user(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
    }

    function _genRandUser(string memory name, uint userType) internal returns (User user) {
        // Create User contract based on userType:
        if (forkType == LOCAL) {
            user = new User(name);

            if (userType == DEFAULT) {
                user = new User(name);
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_AltMethods(name));
            } else {
                revert("_randUser: unimplemented userType");
            }
        } else if (forkType == MAINNET) {
            if (userType == DEFAULT) {
                user = User(new User_M2(name));
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_M2(name));
            } else {
                revert("_randUser: unimplemented userType");
            }
        } else if (forkType == HOLESKY) {
            // User deployment for Holesky is exact same as holesky.
            // Current Holesky deployment is up to date and no deprecated interfaces have been added.

            user = new User(name);

            if (userType == DEFAULT) {
                user = new User(name);
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_AltMethods(name));
            } else {
                revert("_randUser: unimplemented userType");
            }
        } else {
            revert("_randUser: unimplemented forkType");
        }
    }

    function _genRandAVS(
        string memory name
    ) internal returns (AVS avs) {
        if (forkType == LOCAL) {
            avs = new AVS(name);
        } else if (forkType == MAINNET) {
            avs = new AVS(name);
        } else if (forkType == HOLESKY) {
            avs = new AVS(name);
        } else {
            revert("_genRandAVS: unimplemented forkType");
        }
    }

    /// @dev For a given `assetType`, select a random assortment of strategies and assets
    /// NO_ASSETS - return will be empty
    /// HOLDS_LST - `strategies` will be a random subset of initialized strategies
    ///             `tokenBalances` will be the user's balances in each token
    /// HOLDS_ETH - `strategies` will only contain BEACONCHAIN_ETH_STRAT, and
    ///             `tokenBalances` will contain the user's eth balance
    /// HOLDS_ALL - `strategies` will contain ALL initialized strategies AND BEACONCHAIN_ETH_STRAT, and
    ///             `tokenBalances` will contain random token/eth balances accordingly
    function _dealRandAssets(User user, uint assetType) internal returns (IStrategy[] memory, uint[] memory) {
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;
        if (assetType == NO_ASSETS) {
            strategies = new IStrategy[](0);
            tokenBalances = new uint[](0);
        } else if (assetType == HOLDS_LST) {
            assetType = HOLDS_LST;
            // Select a random number of assets
            uint numAssets = _randUint({min: 1, max: lstStrats.length});
            strategies = new IStrategy[](numAssets);
            tokenBalances = new uint[](numAssets);

            // For each asset, award the user a random balance of the underlying token
            for (uint i = 0; i < numAssets; i++) {
                IStrategy strat = lstStrats[i];
                IERC20 underlyingToken = strat.underlyingToken();
                uint balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

                StdCheats.deal(address(underlyingToken), address(user), balance);
                tokenBalances[i] = balance;
                strategies[i] = strat;
            }
        } else if (assetType == HOLDS_ETH) {
            strategies = new IStrategy[](1);
            tokenBalances = new uint[](1);

            // Award the user with a random amount of ETH
            // This guarantees a multiple of 32 ETH (at least 1, up to/incl 5)
            uint amount = 32 ether * _randUint({min: 1, max: 5});
            cheats.deal(address(user), amount);

            strategies[0] = BEACONCHAIN_ETH_STRAT;
            tokenBalances[0] = amount;
        } else if (assetType == HOLDS_ALL || assetType == HOLDS_MAX) {
            uint numLSTs = assetType == HOLDS_MAX ? lstStrats.length : 5;
            strategies = new IStrategy[](numLSTs + 1);
            tokenBalances = new uint[](numLSTs + 1);

            // For each LST, award the user a random balance of the underlying token
            for (uint i = 0; i < numLSTs; i++) {
                IStrategy strat = lstStrats[i];
                IERC20 underlyingToken = strat.underlyingToken();
                uint balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

                StdCheats.deal(address(underlyingToken), address(user), balance);
                tokenBalances[i] = balance;
                strategies[i] = strat;
            }

            // Award the user with a random amount of ETH
            // This guarantees a multiple of 32 ETH (at least 1, up to/incl 5)
            uint amount = 32 ether * _randUint({min: 1, max: 5});
            cheats.deal(address(user), amount);

            // Add BEACONCHAIN_ETH_STRAT and eth balance
            strategies[numLSTs] = BEACONCHAIN_ETH_STRAT;
            tokenBalances[numLSTs] = amount;
        } else {
            revert("_dealRandAssets: assetType unimplemented");
        }

        return (strategies, tokenBalances);
    }

    /// @dev By default will have a assetType of HOLDS_LST
    function _dealRandAssets_M1(
        User user
    ) internal returns (IStrategy[] memory, uint[] memory) {
        // Select a random number of assets
        uint numAssets = _randUint({min: 1, max: lstStrats.length});

        IStrategy[] memory strategies = new IStrategy[](numAssets);
        uint[] memory tokenBalances = new uint[](numAssets);

        // For each asset, award the user a random balance of the underlying token
        for (uint i = 0; i < numAssets; i++) {
            IStrategy strat = lstStrats[i];
            IERC20 underlyingToken = strat.underlyingToken();
            uint balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

            StdCheats.deal(address(underlyingToken), address(user), balance);
            tokenBalances[i] = balance;
            strategies[i] = strat;
        }

        return (strategies, tokenBalances);
    }

    /// @dev Uses `random` to return a random uint, with a range given by `min` and `max` (inclusive)
    /// @return `min` <= result <= `max`
    function _randUint(uint min, uint max) internal returns (uint) {
        uint range = max - min + 1;

        // calculate the number of bits needed for the range
        uint bitsNeeded = 0;
        uint tempRange = range;
        while (tempRange > 0) {
            bitsNeeded++;
            tempRange >>= 1;
        }

        // create a mask for the required number of bits
        // and extract the value from the hash
        uint mask = (1 << bitsNeeded) - 1;
        uint value = uint(random) & mask;

        // in case value is out of range, wrap around or retry
        while (value >= range) {
            value = (value - range) & mask;
        }

        // Hash `random` with itself so the next value we generate is different
        random = _hash(uint(random));
        return min + value;
    }

    function _randBool() internal returns (bool) {
        return _randUint({min: 0, max: 1}) == 0;
    }

    function _randAssetType() internal returns (uint) {
        uint idx = _randUint({min: 0, max: assetTypes.length - 1});
        uint assetType = uint(uint8(assetTypes[idx]));

        return assetType;
    }

    function _randUserType() internal returns (uint) {
        uint idx = _randUint({min: 0, max: userTypes.length - 1});
        uint userType = uint(uint8(userTypes[idx]));

        return userType;
    }

    function _shuffle(IStrategy[] memory strats) internal returns (IStrategy[] memory) {        
        // Fisher-Yates shuffle algorithm
        for (uint i = strats.length - 1; i > 0; i--)  {
            uint randomIndex = _randUint({ min: 0, max: i });
            
            // Swap elements
            IStrategy temp = strats[i];
            strats[i] = strats[randomIndex];
            strats[randomIndex] = temp;
        }
        
        return strats;
    }

    function _randomStrategies() internal returns (IStrategy[][] memory strategies) {
        uint numOpSets = _randUint({ min: 1, max: 5 });

        strategies = new IStrategy[][](numOpSets);
        
        for (uint i; i < numOpSets; ++i) {
            IStrategy[] memory randomStrategies = _shuffle(allStrats);

            uint numStrategies = _randUint({ min: 1, max: allStrats.length });

            // Modify the length of the array in memory (thus ignoring remaining elements).
            assembly {
                mstore(randomStrategies, numStrategies)
            }

            strategies[i] = randomStrategies;
        }
    }

    /**
     * @dev Converts a bitmap into an array of bytes
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap
     */
    function _bitmapToBytes(
        uint bitmap
    ) internal pure returns (bytes memory bytesArray) {
        for (uint i = 0; i < 256; ++i) {
            // Mask for i-th bit
            uint mask = uint(1 << i);

            // If the i-th bit is flipped, add a byte to the return array
            if (bitmap & mask != 0) {
                bytesArray = bytes.concat(bytesArray, bytes1(uint8(1 << i)));
            }
        }
        return bytesArray;
    }

    function _hash(
        string memory x
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }

    function _hash(
        uint x
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }
}
