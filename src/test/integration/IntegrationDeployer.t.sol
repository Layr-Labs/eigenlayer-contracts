// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {MockERC20} from "forge-std/mocks/MockERC20.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/test/integration/users/AVS.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";
import "src/test/integration/users/User_M2.t.sol";
import "src/test/mocks/BeaconChainMock.t.sol";

import "src/test/mocks/EmptyContract.sol";
import "src/test/utils/Constants.t.sol";

import "script/utils/ExistingDeploymentParser.sol";

abstract contract IntegrationDeployer is ExistingDeploymentParser {
    using StdStyle for *;
    using ArrayLib for *;

    /// -----------------------------------------------------------------------
    /// State
    /// -----------------------------------------------------------------------

    /// @notice Returns the genesis time for the beacon chain. Depends on the fork type.
    uint64 public BEACON_GENESIS_TIME;

    /// @notice Returns whether the contracts have been upgraded or not.
    bool public isUpgraded;
    /// @notice Returns the allowed asset types for the tests.
    uint public assetTypes = HOLDS_LST | HOLDS_ETH | HOLDS_ALL;
    /// @notice Returns types of users to be randomly selected during tests
    uint public userTypes = DEFAULT | ALT_METHODS;
    /// @notice Returns the fork type, set only once in setUp if FORK_MAINNET env is set
    uint public forkType;

    /// @notice Returns an array of deployed LST strategies.
    IStrategy[] public lstStrats;
    /// @notice Returns an array of all deployed strategies.
    IStrategy[] public allStrats;
    /// @notice Returns an array of all underlying tokens corresponding to strategies in `allStrats`.
    IERC20[] public allTokens;

    /// @notice Returns the maximum number of unique assets a user holds.
    uint public maxUniqueAssetsHeld;
    /// @dev Returns true if a token should be excluded from testing
    /// If a token is in this mapping, we will ignore this LST as it causes issues with reading balanceOf
    mapping(address => bool) public tokensNotTested;

    /// -----------------------------------------------------------------------
    /// Setup
    /// -----------------------------------------------------------------------

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

        // Use current contracts by default. Upgrade tests are only run with mainnet fork tests
        // using the `UpgradeTest.t.sol` mixin.
        isUpgraded = true;
    }

    /**
     * @dev Anyone who wants to test using this contract in a separate repo via submodules may have to
     * override this function to set the correct paths for the deployment info files.
     *
     * This setUp function will account for specific --fork-url flags and deploy/upgrade contracts accordingly.
     * Note that forkIds are also created so you can make explicit fork tests using cheats.selectFork(forkId)
     */
    function setUp() public virtual {
        SEMVER = "v9.9.9";

        bool forkMainnet = isForktest();

        if (forkMainnet) {
            forkType = MAINNET;
            _setUpMainnet();
        } else {
            forkType = LOCAL;
            _setUpLocal();
        }
        _init();
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /// @dev Override this method in derived test contracts to implement custom initialization logic.
    /// This method is called at the end of the setUp() function after all contracts are deployed and initialized.
    /// It allows test contracts to perform additional setup steps without having to override the entire setUp() function.
    function _init() internal virtual {
        return;
    }

    /**
     * env FOUNDRY_PROFILE=forktest forge t --mc Integration
     *
     * Running foundry like this will trigger the fork test profile,
     * lowering fuzz runs and using a remote RPC to test against mainnet state
     */
    function isForktest() public view returns (bool) {
        return _hash("forktest") == _hash(cheats.envOr(string("FOUNDRY_PROFILE"), string("default")));
    }

    /// Deploy EigenLayer locally
    function _setUpLocal() public virtual noTracing {
        console.log("Setting up `%s` integration tests:", "LOCAL".yellow().bold());

        // Deploy ProxyAdmin
        eigenLayerProxyAdmin = new ProxyAdmin();
        executorMultisig = address(eigenLayerProxyAdmin.owner());

        // Deploy PauserRegistry
        address[] memory pausers = new address[](1);
        pausers[0] = PAUSER;
        eigenLayerPauserReg = new PauserRegistry(pausers, UNPAUSER);

        // Deploy mocks
        emptyContract = new EmptyContract();

        // Matching parameters to testnet
        DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS = 50;
        DEALLOCATION_DELAY = 50;
        ALLOCATION_CONFIGURATION_DELAY = 75;

        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = 86_400;
        REWARDS_COORDINATOR_MAX_REWARDS_DURATION = 6_048_000;
        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = 7_776_000;
        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = 2_592_000;
        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP = 1_710_979_200;

        _deployProxies();
        _deployImplementations();
        _upgradeProxies();
        _initializeProxies();

        // Place native ETH first in `allStrats`
        // This ensures when we select a nonzero number of strategies from this array, we always
        // have beacon chain ETH
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);

        // Deploy and configure strategies and tokens
        for (uint i = 1; i < NUM_LST_STRATS + 1; ++i) {
            string memory name = string.concat("LST-Strat", cheats.toString(i), " token");
            string memory symbol = string.concat("lstStrat", cheats.toString(i));
            // Deploy half of the strategies using the factory.
            _newStrategyAndToken(name, symbol, 10e50, address(this), i % 2 == 0);
        }

        maxUniqueAssetsHeld = allStrats.length;

        // Create time machine and beacon chain. Set block time to beacon chain genesis time and starting block number
        BEACON_GENESIS_TIME = GENESIS_TIME_LOCAL;
        cheats.warp(BEACON_GENESIS_TIME);
        cheats.roll(10_000);

        cheats.etch(address(timeMachine), type(TimeMachine).runtimeCode);
        cheats.etch(address(beaconChain), type(BeaconChainMock).runtimeCode);
        cheats.allowCheatcodes(address(timeMachine));
        cheats.allowCheatcodes(address(beaconChain));
        beaconChain.initialize(eigenPodManager, BEACON_GENESIS_TIME);
    }

    /// Parse existing contracts from mainnet
    function _setUpMainnet() public virtual noTracing {
        console.log("Setting up `%s` integration tests:", "MAINNET_FORK".green().bold());
        console.log("RPC:", cheats.rpcUrl("mainnet"));
        console.log("Block:", MAINNET_FORK_BLOCK);

        cheats.createSelectFork(cheats.rpcUrl("mainnet"), MAINNET_FORK_BLOCK);

        string memory deploymentInfoPath = "script/configs/mainnet/mainnet-addresses.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        string memory existingDeploymentParams = "script/configs/mainnet.json";
        _parseParamsForIntegrationUpgrade(existingDeploymentParams);

        // Place native ETH first in `allStrats`
        // This ensures when we select a nonzero number of strategies from this array, we always
        // have beacon chain ETH
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);

        // Add deployed strategies to lstStrats and allStrats
        for (uint i; i < deployedStrategyArray.length; i++) {
            IStrategy strategy = IStrategy(deployedStrategyArray[i]);

            if (tokensNotTested[address(strategy.underlyingToken())]) continue;

            // Add to lstStrats and allStrats
            lstStrats.push(strategy);
            allStrats.push(strategy);
            allTokens.push(strategy.underlyingToken());
        }

        maxUniqueAssetsHeld = allStrats.length;

        // Create time machine and mock beacon chain
        BEACON_GENESIS_TIME = GENESIS_TIME_MAINNET;
        cheats.etch(address(timeMachine), type(TimeMachine).runtimeCode);
        cheats.etch(address(beaconChain), type(BeaconChainMock).runtimeCode);
        cheats.allowCheatcodes(address(timeMachine));
        cheats.allowCheatcodes(address(beaconChain));
        beaconChain.initialize(eigenPodManager, BEACON_GENESIS_TIME);

        // Since we haven't done the slashing upgrade on mainnet yet, upgrade mainnet contracts
        // prior to test. `isUpgraded` is true by default, but is set to false in `UpgradeTest.t.sol`
        if (isUpgraded) _upgradeMainnetContracts();
    }

    function _upgradeMainnetContracts() public virtual {
        cheats.startPrank(address(executorMultisig));

        // First, deploy the new contracts as empty contracts
        emptyContract = new EmptyContract();
        allocationManager =
            AllocationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        permissionController =
            PermissionController(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));

        emit log_named_uint("EPM pause status", eigenPodManager.paused());

        // Deploy new implementation contracts and upgrade all proxies to point to them
        _deployImplementations();
        _upgradeProxies();

        emit log_named_uint("EPM pause status", eigenPodManager.paused());

        // Initialize the newly-deployed proxy
        allocationManager.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});

        cheats.stopPrank();
    }

    function _deployProxies() public {
        delegationManager =
            DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        strategyManager =
            StrategyManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        eigenPodManager =
            EigenPodManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        rewardsCoordinator =
            RewardsCoordinator(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        avsDirectory = AVSDirectory(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        strategyFactory =
            StrategyFactory(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        allocationManager =
            AllocationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        permissionController =
            PermissionController(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        eigenPodBeacon = new UpgradeableBeacon(address(emptyContract));
        strategyBeacon = new UpgradeableBeacon(address(emptyContract));
    }

    /// Deploy an implementation contract for each contract in the system
    function _deployImplementations() public {
        allocationManagerImplementation = new AllocationManager(
            delegationManager, eigenLayerPauserReg, permissionController, DEALLOCATION_DELAY, ALLOCATION_CONFIGURATION_DELAY, SEMVER
        );
        permissionControllerImplementation = new PermissionController(SEMVER);
        delegationManagerImplementation = new DelegationManager(
            strategyManager,
            eigenPodManager,
            allocationManager,
            eigenLayerPauserReg,
            permissionController,
            DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
            SEMVER
        );
        strategyManagerImplementation = new StrategyManager(delegationManager, eigenLayerPauserReg, SEMVER);
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                delegationManager: delegationManager,
                strategyManager: strategyManager,
                allocationManager: allocationManager,
                pauserRegistry: eigenLayerPauserReg,
                permissionController: permissionController,
                CALCULATION_INTERVAL_SECONDS: REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
                MAX_REWARDS_DURATION: REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
                MAX_RETROACTIVE_LENGTH: REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
                MAX_FUTURE_LENGTH: REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
                GENESIS_REWARDS_TIMESTAMP: REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP,
                version: SEMVER
            })
        );
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg, SEMVER);
        eigenPodManagerImplementation =
            new EigenPodManager(DEPOSIT_CONTRACT, eigenPodBeacon, delegationManager, eigenLayerPauserReg, "v9.9.9");
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg, "v9.9.9");

        // Beacon implementations
        eigenPodImplementation = new EigenPod(DEPOSIT_CONTRACT, eigenPodManager, BEACON_GENESIS_TIME, "v9.9.9");
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg, "v9.9.9");

        // Pre-longtail StrategyBaseTVLLimits implementation
        // TODO - need to update ExistingDeploymentParser
    }

    function _upgradeProxies() public noTracing {
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

        // RewardsCoordinator
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(rewardsCoordinator))), address(rewardsCoordinatorImplementation)
        );

        // AVSDirectory
        eigenLayerProxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(avsDirectory))), address(avsDirectoryImplementation));

        // AllocationManager
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(allocationManager))), address(allocationManagerImplementation)
        );

        // PermissionController
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(permissionController))), address(permissionControllerImplementation)
        );

        // StrategyFactory
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(strategyFactory))), address(strategyFactoryImplementation)
        );

        // EigenPod beacon
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));

        // StrategyBase Beacon
        strategyBeacon.upgradeTo(address(baseStrategyImplementation));

        // Upgrade All deployed strategy contracts to new base strategy
        for (uint i = 0; i < numStrategiesDeployed; i++) {
            // Upgrade existing strategy
            eigenLayerProxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(deployedStrategyArray[i]))), address(baseStrategyImplementation)
            );
        }
    }

    function _initializeProxies() public noTracing {
        allocationManager.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        avsDirectory.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        delegationManager.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        eigenPodManager.initialize({initialOwner: executorMultisig, _initPausedStatus: 0});
        strategyFactory.initialize({_initialOwner: executorMultisig, _initialPausedStatus: 0, _strategyBeacon: strategyBeacon});
        strategyManager.initialize({
            initialOwner: executorMultisig,
            initialStrategyWhitelister: address(strategyFactory),
            initialPausedStatus: 0
        });
    }

    /// @dev Deploy a strategy and its underlying token, push to global lists of tokens/strategies, and whitelist
    /// strategy in strategyManager
    function _newStrategyAndToken(string memory tokenName, string memory tokenSymbol, uint initialSupply, address owner, bool useFactory)
        internal
        noTracing
    {
        MockERC20 token = new MockERC20();
        token.initialize(tokenName, tokenSymbol, 18);
        IERC20 underlyingToken = IERC20(address(token));
        deal(address(underlyingToken), address(owner), initialSupply);

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

        cheats.prank(strategyManager.strategyWhitelister());
        strategyManager.addStrategiesToDepositWhitelist(strategies);

        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
    }

    function _configRand(uint _assetTypes, uint _userTypes) private noTracing {
        // Set asset types and user types directly
        assetTypes = _assetTypes;
        userTypes = _userTypes;

        assertTrue(assetTypes != 0, "_configRand: no asset types selected");
        assertTrue(userTypes != 0, "_configRand: no user types selected");
    }

    function _configAssetTypes(uint _assetTypes) internal {
        assetTypes = _assetTypes;
        assertTrue(assetTypes != 0, "_configRand: no asset types selected");
    }

    function _configAssetAmounts(uint _maxUniqueAssetsHeld) internal {
        if (_maxUniqueAssetsHeld > allStrats.length) _maxUniqueAssetsHeld = allStrats.length;

        maxUniqueAssetsHeld = _maxUniqueAssetsHeld;
        require(maxUniqueAssetsHeld != 0, "_configAssetAmounts: invalid 0");
    }

    function _configUserTypes(uint _userTypes) internal {
        userTypes = _userTypes;
        assertTrue(userTypes != 0, "_configRand: no user types selected");
    }

    /**
     * @dev Create a new User with a random config using the range defined in `_configRand`
     *
     * Assets are pulled from `strategies` based on a random staker/operator `assetType`
     */
    function _randUser(string memory name) internal noTracing returns (User, IStrategy[] memory, uint[] memory) {
        // Deploy new User contract
        uint userType = _randUserType();
        User user = _genRandUser(name, userType);

        // For the specific asset selection we made, get a random assortment of strategies
        // and deal the user some corresponding underlying token balances
        uint assetType = _randAssetType();
        IStrategy[] memory strategies = _selectRandAssets(assetType);
        uint[] memory tokenBalances = _dealRandAmounts(user, strategies);

        print.user(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
    }

    function _randUser(string memory name, IStrategy[] memory strategies) internal noTracing returns (User, uint[] memory) {
        // Deploy new User contract
        uint userType = _randUserType();
        User user = _genRandUser(name, userType);

        // Deal the user some corresponding underlying token balances
        uint[] memory tokenBalances = _dealRandAmounts(user, strategies);

        print.user(name, HOLDS_ALL, userType, strategies, tokenBalances);
        return (user, tokenBalances);
    }

    /// @dev Creates a new user without any assets
    function _randUser_NoAssets(string memory name) internal noTracing returns (User) {
        // Deploy new User contract
        uint userType = _randUserType();
        User user = _genRandUser(name, userType);

        print.user(name, NO_ASSETS, userType, new IStrategy[](0), new uint[](0));
        return user;
    }

    function _genRandUser(string memory name, uint userType) internal returns (User user) {
        // Create User contract based on userType:
        if (forkType == LOCAL || (forkType == MAINNET && isUpgraded)) {
            user = new User(name);

            if (userType == DEFAULT) {
                user = new User(name);
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_AltMethods(name));
            } else {
                revert("_randUser: unimplemented userType");
            }
        } else if (forkType == MAINNET && !isUpgraded) {
            if (userType == DEFAULT) {
                user = User(new User_M2(name));
            } else if (userType == ALT_METHODS) {
                // User will use nonstandard methods like `depositIntoStrategyWithSignature`
                user = User(new User_M2(name));
            } else {
                revert("_randUser: unimplemented userType");
            }
        } else {
            revert("_randUser: unimplemented forkType");
        }
    }

    function _genRandAVS(string memory name) internal returns (AVS avs) {
        if (forkType == LOCAL) avs = new AVS(name);
        else if (forkType == MAINNET) avs = new AVS(name);
        else revert("_genRandAVS: unimplemented forkType");
    }

    /// Given an assetType, select strategies the user will be dealt assets in
    function _selectRandAssets(uint assetType) internal noTracing returns (IStrategy[] memory strategies) {
        if (assetType == NO_ASSETS) return new IStrategy[](0);
        if (assetType == HOLDS_ETH) return BEACONCHAIN_ETH_STRAT.toArray();

        // Select number of assets:
        // HOLDS_LST can hold at most all LSTs. HOLDS_ALL and HOLDS_MAX also hold ETH.
        // Clamp number of assets to maxUniqueAssetsHeld (guaranteed to be at least 1)
        uint assetPoolSize = assetType == HOLDS_LST ? lstStrats.length : allStrats.length;
        uint maxAssets = assetPoolSize > maxUniqueAssetsHeld ? maxUniqueAssetsHeld : assetPoolSize;
        uint numAssets = assetType == HOLDS_MAX ? maxAssets : _randUint(1, maxAssets);

        strategies = new IStrategy[](numAssets);
        for (uint i = 0; i < strategies.length; i++) {
            strategies[i] = assetType == HOLDS_LST ? lstStrats[i] : allStrats[i];
        }

        return strategies.sort();
    }

    /// Given an input list of strategies, deal random underlying token amounts to a user
    function _dealRandAmounts(User user, IStrategy[] memory strategies) internal noTracing returns (uint[] memory) {
        uint[] memory tokenBalances = new uint[](strategies.length);

        for (uint i = 0; i < tokenBalances.length; i++) {
            IStrategy strategy = strategies[i];
            uint balance;

            if (strategy == BEACONCHAIN_ETH_STRAT) {
                // Award the user with a random amount of ETH
                // This guarantees a multiple of 32 ETH (at least 1, up to/incl 5)
                balance = 32 ether * _randUint({min: 1, max: 5});
                cheats.deal(address(user), balance);
            } else {
                IERC20 underlyingToken = strategy.underlyingToken();
                balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});

                StdCheats.deal(address(underlyingToken), address(user), balance);
            }

            tokenBalances[i] = balance;
        }

        return tokenBalances;
    }

    /// Given an array of strategies and an array of amounts, deal the amounts to the user
    function _dealAmounts(User user, IStrategy[] memory strategies, uint[] memory amounts) internal noTracing {
        for (uint i = 0; i < amounts.length; i++) {
            IStrategy strategy = strategies[i];

            if (strategy == BEACONCHAIN_ETH_STRAT) cheats.deal(address(user), amounts[i]);
            else deal(address(strategy.underlyingToken()), address(user), amounts[i]);
        }
    }

    function _randUint(uint min, uint max) internal returns (uint) {
        return cheats.randomUint(min, max);
    }

    function _randBool() internal returns (bool) {
        return cheats.randomBool();
    }

    function _randAssetType() internal returns (uint) {
        // Overflow is not possible given the constraints of the assetTypes bitmap.
        // Underflow is only possible if the bitmap is 0, which is checked for above.
        unchecked {
            uint[] memory options = new uint[](5); // We have 5 possible asset types
            uint count = 0;
            if (assetTypes & NO_ASSETS != 0) options[count++] = NO_ASSETS;
            if (assetTypes & HOLDS_LST != 0) options[count++] = HOLDS_LST;
            if (assetTypes & HOLDS_ETH != 0) options[count++] = HOLDS_ETH;
            if (assetTypes & HOLDS_ALL != 0) options[count++] = HOLDS_ALL;
            if (assetTypes & HOLDS_MAX != 0) options[count++] = HOLDS_MAX;
            return options[cheats.randomUint(0, count - 1)];
        }
    }

    function _randUserType() internal returns (uint) {
        // Overflow is not possible given the constraints of the userTypes bitmap.
        // Underflow is only possible if the bitmap is 0, which is checked for above.
        unchecked {
            uint[] memory options = new uint[](2); // We have 2 possible user types
            uint count = 0;
            if (userTypes & DEFAULT != 0) options[count++] = DEFAULT;
            if (userTypes & ALT_METHODS != 0) options[count++] = ALT_METHODS;
            return options[cheats.randomUint(0, count - 1)];
        }
    }

    function _hash(string memory x) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }

    function NAME() public view virtual override returns (string memory) {
        return "Integration Deployer";
    }
}
