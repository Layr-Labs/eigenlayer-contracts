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

import "src/test/utils/Logger.t.sol";
import "src/test/mocks/EmptyContract.sol";
import "src/test/mocks/ETHDepositMock.sol";
import "src/test/integration/mocks/BeaconChainMock.t.sol";

import "src/test/integration/users/AVS.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

import "script/utils/ExistingDeploymentParser.sol";

abstract contract IntegrationDeployer is ExistingDeploymentParser, Logger {
    using StdStyle for *;

    Vm cheats = Vm(VM_ADDRESS);

    // Fork ids for specific fork tests
    bool isUpgraded;
    uint256 mainnetForkBlock = 19_280_000;
    uint256 mainnetForkId;
    uint256 holeskyForkBLock = 1_213_950;
    uint256 holeskyForkId;
    uint64 constant DENEB_FORK_TIMESTAMP = 1_705_473_120;

    // Beacon chain genesis time when running locally
    // Multiple of 12 for sanity's sake
    uint64 constant GENESIS_TIME_LOCAL = 1 hours * 12;
    uint64 constant GENESIS_TIME_MAINNET = 1_606_824_023;

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
    uint256 forkType;

    // Constants
    uint64 constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;

    IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    IERC20 constant NATIVE_ETH = IERC20(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    uint256 constant MIN_BALANCE = 1e6;
    uint256 constant MAX_BALANCE = 5e6;
    uint256 constant GWEI_TO_WEI = 1e9;

    // Paused Constants
    // DelegationManager
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;
    // StrategyManager
    uint8 internal constant PAUSED_DEPOSITS = 0;
    // EigenpodManager
    uint8 internal constant PAUSED_NEW_EIGENPODS = 0;
    uint8 internal constant PAUSED_WITHDRAW_RESTAKED_ETH = 1;
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_CREDENTIALS = 2;
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE = 3;
    uint8 internal constant PAUSED_EIGENPODS_VERIFY_WITHDRAWAL = 4;
    uint8 internal constant PAUSED_NON_PROOF_WITHDRAWALS = 5;

    // Flags
    uint256 constant FLAG = 1;

    /// @dev Asset flags
    /// These are used with _configRand to determine what assets are given
    /// to a user when they are created.
    uint256 constant NO_ASSETS = (FLAG << 0); // will have no assets
    uint256 constant HOLDS_LST = (FLAG << 1); // will hold some random amount of LSTs
    uint256 constant HOLDS_ETH = (FLAG << 2); // will hold some random amount of ETH
    uint256 constant HOLDS_ALL = (FLAG << 3); // will hold every LST and ETH

    /// @dev User contract flags
    /// These are used with _configRand to determine what User contracts can be deployed
    uint256 constant DEFAULT = (FLAG << 0);
    uint256 constant ALT_METHODS = (FLAG << 1);

    /// @dev Shadow Fork flags
    /// These are used for upgrade integration testing.
    uint256 constant LOCAL = (FLAG << 0);
    uint256 constant MAINNET = (FLAG << 1);
    uint256 constant HOLESKY = (FLAG << 2);

    // /// @dev Withdrawal flags
    // /// These are used with _configRand to determine how a user conducts a withdrawal
    // uint constant FULL_WITHDRAW_SINGLE = (FLAG << 0); // stakers will withdraw all assets using a single queued withdrawal
    // uint constant FULL_WITHDRAW_MULTI  = (FLAG << 1); // stakers will withdraw all assets using multiple queued withdrawals
    // uint constant PART_WITHDRAW_SINGLE = (FLAG << 2); // stakers will withdraw some, but not all assets

    /// Note: Thought about the following flags (but did not implement) -
    ///
    /// WithdrawerType (SELF_WITHDRAWER, OTHER_WITHDRAWER)
    ///     - especially with EPM share handling, this felt like it deserved its own test rather than a fuzzy state
    /// CompletionType (AS_TOKENS, AS_SHARES)
    ///     - same reason as above
    ///
    /// WithdrawalMethod (QUEUE_WITHDRAWAL, UNDELEGATE, REDELEGATE)
    ///     - could still do this!
    ///     - This would trigger staker.queueWithdrawals to use either `queueWithdrawals` or `undelegate` under the hood
    ///     - "redelegate" would be like the above, but adding a new `delegateTo` step after undelegating

    mapping(uint256 => string) assetTypeToStr;
    mapping(uint256 => string) userTypeToStr;
    mapping(uint256 => string) forkTypeToStr;

    constructor() {
        assetTypeToStr[NO_ASSETS] = "NO_ASSETS";
        assetTypeToStr[HOLDS_LST] = "HOLDS_LST";
        assetTypeToStr[HOLDS_ETH] = "HOLDS_ETH";
        assetTypeToStr[HOLDS_ALL] = "HOLDS_ALL";

        userTypeToStr[DEFAULT] = "DEFAULT";
        userTypeToStr[ALT_METHODS] = "ALT_METHODS";

        forkTypeToStr[LOCAL] = "LOCAL";
        forkTypeToStr[MAINNET] = "MAINNET";
        forkTypeToStr[HOLESKY] = "HOLESKY";

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
            console.log("Setting up integration tests against mainnet fork:");
            console.log("Rpc Url:", cheats.rpcUrl("mainnet"));
            console.log("Fork Block:", mainnetForkBlock);

            cheats.createSelectFork(cheats.rpcUrl("mainnet"), mainnetForkBlock);
            forkType = MAINNET;
        } else {
            console.log("Setting up integration tests locally");

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

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GENESIS_TIME_LOCAL);

        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));
        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(
            avsDirectory, strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, MIN_WITHDRAWAL_DELAY
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation =
            new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, delegationManager, eigenLayerPauserReg);
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg);
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg);
        allocationManagerImplementation = new AllocationManager(
            delegationManager, eigenLayerPauserReg, DEALLOCATION_DELAY, ALLOCATION_CONFIGURATION_DELAY
        );

        // Third, upgrade the proxy contracts to point to the implementations
        uint256 withdrawalDelayBlocks = 7 days / 12 seconds;
        IStrategy[] memory initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        uint256[] memory initializeWithdrawalDelayBlocks = new uint256[](0);
        // DelegationManager
        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                0, /* initialPausedStatus */
                withdrawalDelayBlocks,
                initializeStrategiesToSetDelayBlocks,
                initializeWithdrawalDelayBlocks
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

        // Normal deployments
        _newStrategyAndToken("strategy1.underlyingToken()", "str1", 10e50, address(this), false); // initialSupply, owner
        _newStrategyAndToken("strategy2.underlyingToken()", "str2", 10e50, address(this), false); // initialSupply, owner
        _newStrategyAndToken("strategy3.underlyingToken()", "str3", 10e50, address(this), false); // initialSupply, owner

        // Factory deployments
        _newStrategyAndToken("strategy4.underlyingToken()", "str4", 10e50, address(this), true); // initialSupply, owner
        _newStrategyAndToken("strategy5.underlyingToken()", "str5", 10e50, address(this), true); // initialSupply, owner
        _newStrategyAndToken("strategy6.underlyingToken()", "str6", 10e50, address(this), true); // initialSupply, owner

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

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(ethPOSDeposit, eigenPodManager, GENESIS_TIME_MAINNET);
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
        // Deploy AVSDirectory, contract has not been deployed on mainnet yet
        avsDirectory = AVSDirectory(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // First, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        delegationManagerImplementation = new DelegationManager(
            avsDirectory, strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, MIN_WITHDRAWAL_DELAY
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation =
            new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, delegationManager, eigenLayerPauserReg);
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
        for (uint256 i = 0; i < numStrategiesDeployed; i++) {
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
        delegationManagerImplementation = new DelegationManager(
            avsDirectory, strategyManager, eigenPodManager, allocationManager, eigenLayerPauserReg, MIN_WITHDRAWAL_DELAY
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg);
        eigenPodManagerImplementation =
            new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, delegationManager, eigenLayerPauserReg);
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
        for (uint256 i = 0; i < numStrategiesDeployed; i++) {
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
        uint256 initialSupply,
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
            StrategyBaseTVLLimits(address(strategy)).setTVLLimits(type(uint256).max, type(uint256).max);
        } else {
            cheats.prank(strategyManager.strategyWhitelister());
            strategyManager.addStrategiesToDepositWhitelist(strategies);
        }

        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
    }

    function _configRand(uint24 _randomSeed, uint256 _assetTypes, uint256 _userTypes) internal {
        // Using uint24 for the seed type so that if a test fails, it's easier
        // to manually use the seed to replay the same test.

        random = _hash(_randomSeed);

        // Convert flag bitmaps to bytes of set bits for easy use with _randUint
        assetTypes = _bitmapToBytes(_assetTypes);
        userTypes = _bitmapToBytes(_userTypes);

        console.log("_configRand: Users will be initialized with these asset types:");
        for (uint256 i = 0; i < assetTypes.length; i++) {
            console.log(assetTypeToStr[uint256(uint8(assetTypes[i]))]);
        }

        console.log("_configRand: these User contracts will be initialized:");
        for (uint256 i = 0; i < userTypes.length; i++) {
            console.log(userTypeToStr[uint256(uint8(userTypes[i]))]);
        }

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
        console.log("_deployOrFetchContracts using fork for test", forkTypeToStr[forkType]);

        if (forkType == LOCAL) {
            _setUpLocal();
            // Set Upgraded as local setup deploys most up to date contracts
            isUpgraded = true;
        } else if (forkType == MAINNET) {
            // cheats.selectFork(mainnetForkId);
            string memory deploymentInfoPath = "script/configs/mainnet/mainnet-addresses.config.json";
            _parseDeployedContracts(deploymentInfoPath);

            // Unpause to enable deposits and withdrawals for initializing random user state
            cheats.prank(eigenLayerPauserReg.unpauser());
            strategyManager.unpause(0);

            // Add deployed strategies to lstStrats and allStrats
            for (uint256 i; i < deployedStrategyArray.length; i++) {
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
    ) internal returns (User, IStrategy[] memory, uint256[] memory) {
        // For the new user, select what type of assets they'll have and whether
        // they'll use `xWithSignature` methods.
        //
        // The values selected here are in the ranges configured via `_configRand`
        uint256 assetType = _randAssetType();
        uint256 userType = _randUserType();

        // Deploy new User contract
        User user = _genRandUser(name, userType);

        // For the specific asset selection we made, get a random assortment of
        // strategies and deal the user some corresponding underlying token balances
        (IStrategy[] memory strategies, uint256[] memory tokenBalances) = _dealRandAssets(user, assetType);

        _printUserInfo(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
    }

    /// @dev Create a new user without native ETH. See _randUser above for standard usage
    function _randUser_NoETH(
        string memory name
    ) internal returns (User, IStrategy[] memory, uint256[] memory) {
        // For the new user, select what type of assets they'll have and whether
        // they'll use `xWithSignature` methods.
        //
        // The values selected here are in the ranges configured via `_configRand`
        uint256 userType = _randUserType();

        // Pick the user's asset distribution, removing "native ETH" as an option
        // I'm sorry if this eventually leads to a bug that's really hard to track down
        uint256 assetType = _randAssetType();
        if (assetType == HOLDS_ETH) {
            assetType = NO_ASSETS;
        } else if (assetType == HOLDS_ALL) {
            assetType = HOLDS_LST;
        }

        // Deploy new User contract
        User user = _genRandUser(name, userType);

        // For the specific asset selection we made, get a random assortment of
        // strategies and deal the user some corresponding underlying token balances
        (IStrategy[] memory strategies, uint256[] memory tokenBalances) = _dealRandAssets(user, assetType);

        _printUserInfo(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
    }

    function _genRandUser(string memory name, uint256 userType) internal returns (User user) {
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
                user = User(new User_M1(name));
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_M1_AltMethods(name));
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
    function _dealRandAssets(User user, uint256 assetType) internal returns (IStrategy[] memory, uint256[] memory) {
        IStrategy[] memory strategies;
        uint256[] memory tokenBalances;
        if (assetType == NO_ASSETS) {
            strategies = new IStrategy[](0);
            tokenBalances = new uint256[](0);
        } else if (assetType == HOLDS_LST) {
            assetType = HOLDS_LST;
            // Select a random number of assets
            uint256 numAssets = _randUint({min: 1, max: lstStrats.length});
            strategies = new IStrategy[](numAssets);
            tokenBalances = new uint256[](numAssets);

            // For each asset, award the user a random balance of the underlying token
            for (uint256 i = 0; i < numAssets; i++) {
                IStrategy strat = lstStrats[i];
                IERC20 underlyingToken = strat.underlyingToken();
                uint256 balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

                StdCheats.deal(address(underlyingToken), address(user), balance);
                tokenBalances[i] = balance;
                strategies[i] = strat;
            }
        } else if (assetType == HOLDS_ETH) {
            strategies = new IStrategy[](1);
            tokenBalances = new uint256[](1);

            // Award the user with a random amount of ETH
            // This guarantees a multiple of 32 ETH (at least 1, up to/incl 5)
            uint256 amount = 32 ether * _randUint({min: 1, max: 5});
            cheats.deal(address(user), amount);

            strategies[0] = BEACONCHAIN_ETH_STRAT;
            tokenBalances[0] = amount;
        } else if (assetType == HOLDS_ALL) {
            uint256 numLSTs = lstStrats.length;
            strategies = new IStrategy[](numLSTs + 1);
            tokenBalances = new uint256[](numLSTs + 1);

            // For each LST, award the user a random balance of the underlying token
            for (uint256 i = 0; i < numLSTs; i++) {
                IStrategy strat = lstStrats[i];
                IERC20 underlyingToken = strat.underlyingToken();
                uint256 balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

                StdCheats.deal(address(underlyingToken), address(user), balance);
                tokenBalances[i] = balance;
                strategies[i] = strat;
            }

            // Award the user with a random amount of ETH
            // This guarantees a multiple of 32 ETH (at least 1, up to/incl 5)
            uint256 amount = 32 ether * _randUint({min: 1, max: 5});
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
    ) internal returns (IStrategy[] memory, uint256[] memory) {
        // Select a random number of assets
        uint256 numAssets = _randUint({min: 1, max: lstStrats.length});

        IStrategy[] memory strategies = new IStrategy[](numAssets);
        uint256[] memory tokenBalances = new uint256[](numAssets);

        // For each asset, award the user a random balance of the underlying token
        for (uint256 i = 0; i < numAssets; i++) {
            IStrategy strat = lstStrats[i];
            IERC20 underlyingToken = strat.underlyingToken();
            uint256 balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

            StdCheats.deal(address(underlyingToken), address(user), balance);
            tokenBalances[i] = balance;
            strategies[i] = strat;
        }

        return (strategies, tokenBalances);
    }

    /// @dev Uses `random` to return a random uint, with a range given by `min` and `max` (inclusive)
    /// @return `min` <= result <= `max`
    function _randUint(uint256 min, uint256 max) internal returns (uint256) {
        uint256 range = max - min + 1;

        // calculate the number of bits needed for the range
        uint256 bitsNeeded = 0;
        uint256 tempRange = range;
        while (tempRange > 0) {
            bitsNeeded++;
            tempRange >>= 1;
        }

        // create a mask for the required number of bits
        // and extract the value from the hash
        uint256 mask = (1 << bitsNeeded) - 1;
        uint256 value = uint256(random) & mask;

        // in case value is out of range, wrap around or retry
        while (value >= range) {
            value = (value - range) & mask;
        }

        // Hash `random` with itself so the next value we generate is different
        random = _hash(uint256(random));
        return min + value;
    }

    function _randBool() internal returns (bool) {
        return _randUint({min: 0, max: 1}) == 0;
    }

    function _randAssetType() internal returns (uint256) {
        uint256 idx = _randUint({min: 0, max: assetTypes.length - 1});
        uint256 assetType = uint256(uint8(assetTypes[idx]));

        return assetType;
    }

    function _randUserType() internal returns (uint256) {
        uint256 idx = _randUint({min: 0, max: userTypes.length - 1});
        uint256 userType = uint256(uint8(userTypes[idx]));

        return userType;
    }

    /**
     * @dev Converts a bitmap into an array of bytes
     * @dev Each byte in the input is processed as indicating a single bit to flip in the bitmap
     */
    function _bitmapToBytes(
        uint256 bitmap
    ) internal pure returns (bytes memory bytesArray) {
        for (uint256 i = 0; i < 256; ++i) {
            // Mask for i-th bit
            uint256 mask = uint256(1 << i);

            // console.log("mask: ", mask);

            // If the i-th bit is flipped, add a byte to the return array
            if (bitmap & mask != 0) {
                bytesArray = bytes.concat(bytesArray, bytes1(uint8(1 << i)));
            }
        }
        return bytesArray;
    }

    function _printUserInfo(
        string memory name,
        uint256 assetType,
        uint256 userType,
        IStrategy[] memory strategies,
        uint256[] memory tokenBalances
    ) internal view {
        console.log("\n====== Created User %s ======", name.bold());
        console.log("   Asset Type:", assetTypeToStr[assetType]);
        console.log("   User Type:", userTypeToStr[userType]);
        console.log("   Fork Type:", forkTypeToStr[forkType]);
        console.log("   Total Assets:", strategies.length);

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                console.log("       Native ETH balance: %s", _toStringWad(tokenBalances[i]));
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                console.log(
                    "       %s balance: %s",
                    IERC20Metadata(address(underlyingToken)).name(),
                    cheats.toString(tokenBalances[i])
                );
            }
        }
    }

    function _hash(
        string memory x
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }

    function _hash(
        uint256 x
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }
}
