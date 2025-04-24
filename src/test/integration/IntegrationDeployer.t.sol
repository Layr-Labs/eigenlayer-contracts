// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {MockERC20} from "forge-std/mocks/MockERC20.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

import "src/test/integration/users/AVS.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";
import "src/test/integration/users/User_M2.t.sol";
import "src/test/mocks/BeaconChainMock.t.sol";
import "src/test/mocks/EmptyContract.sol";
import "src/test/utils/Constants.t.sol";
import "src/test/Config.t.sol";

abstract contract IntegrationDeployer is ConfigGetters, Logger {
    using ConfigParser for *;
    using StdStyle for *;
    using ArrayLib for *;

    // TODO get rid of this storage
    /// @dev AllocationManager
    uint32 constant DEALLOCATION_DELAY = 50;
    uint32 constant ALLOCATION_CONFIGURATION_DELAY = 75;

    /// @dev DelegationManager
    uint32 constant DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS = 50;

    /// @dev RewardsCoordinator
    uint32 constant REWARDS_COORDINATOR_MAX_REWARDS_DURATION = 6_048_000;
    uint32 constant REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = 7_776_000;
    uint32 constant REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = 2_592_000;
    uint32 constant REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP = 1_710_979_200;
    uint32 constant REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = 86_400;

    /// -----------------------------------------------------------------------
    /// State
    /// -----------------------------------------------------------------------

    /// @dev Returns the semver for test environments.
    string constant SEMVER = "v9.9.9-test";

    /// @notice Returns the genesis time for the beacon chain. Depends on the fork type.
    uint64 public BEACON_GENESIS_TIME;

    /// @notice Returns whether the contracts have been upgraded or not.
    bool public isUpgraded;
    /// @notice Returns the allowed asset types for the tests.
    uint public assetTypes = HOLDS_LST | HOLDS_ETH | HOLDS_ALL;
    /// @notice Returns types of users to be randomly selected during tests
    uint public userTypes = DEFAULT | ALT_METHODS;

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

    EmptyContract public emptyContract;

    string public profile;

    ForkConfig public forkConfig;

    /// -----------------------------------------------------------------------
    /// Setup
    /// -----------------------------------------------------------------------

    constructor() {
        // QUESTION: Why is this needed? Shouldn't we have coverage for weird ERC20s?
        tokensNotTested[address(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84)] = true; // stETH mainnet
        tokensNotTested[address(0x856c4Efb76C1D1AE02e20CEB03A2A6a08b0b8dC3)] = true; // oETH mainnet
        tokensNotTested[address(0xf1C9acDc66974dFB6dEcB12aA385b9cD01190E38)] = true; // osETH mainnet
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
        _setUp(true);
    }

    function _setUp(bool upgrade) internal virtual {
        profile = FOUNDRY_PROFILE();
        forkConfig = ConfigParser.parseForkConfig(profile);
        emptyContract = new EmptyContract();

        if (eq(profile, "default") || eq(profile, "coverage")) {
            // Assumes nothing has been deployed yet.
            _setUpLocal();
        } else if (eq(profile, "forktest-zeus")) {
            // Assumes the proxy contracts have already been deployed.
            config = ConfigParser.parseZeus();
            config.label();
            _setUpFork(upgrade);
        } else {
            // Assumes the proxy contracts have already been deployed.
            config = ConfigParser.parse(string.concat("script/configs/", profile, ".toml"));
            config.label();
            _setUpFork(upgrade);
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

    /// @dev Configures the possible asset and user types for the random users.
    function _configRand(uint _assetTypes, uint _userTypes) internal virtual {
        _configAssetTypes(_assetTypes);
        _configUserTypes(_userTypes);
    }

    /// @dev Configures the possible asset types for the random users.
    function _configAssetTypes(uint _assetTypes) internal virtual {
        assertTrue((assetTypes = _assetTypes) != 0, "_configAssetTypes: no asset types selected");
    }

    /// @dev Configures the possible user types for the random users.
    function _configUserTypes(uint _userTypes) internal virtual {
        assertTrue((userTypes = _userTypes) != 0, "_configUserTypes: no user types selected");
    }

    /// @dev Configures the maximum number of unique assets a user can hold.
    function _configAssetAmounts(uint _maxUniqueAssetsHeld) internal virtual {
        if (_maxUniqueAssetsHeld > allStrats.length) _maxUniqueAssetsHeld = allStrats.length;
        assertTrue((maxUniqueAssetsHeld = _maxUniqueAssetsHeld) != 0, "_configAssetAmounts: invalid 0");
    }

    function FOUNDRY_PROFILE() internal view returns (string memory) {
        return cheats.envOr(string("FOUNDRY_PROFILE"), string("default"));
    }

    function eq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }

    /// -----------------------------------------------------------------------
    /// Environment Setup
    /// -----------------------------------------------------------------------

    /// @dev Sets up the integration tests for local.
    function _setUpLocal() public virtual {
        console.log("Setting up `%s` integration tests:", "LOCAL".yellow().bold());

        // Deploy ProxyAdmin, PauserRegistry, and executorMultisig.
        config.governance.proxyAdmin = new ProxyAdmin();
        config.governance.pauserRegistry = new PauserRegistry(PAUSER.toArray(), UNPAUSER);
        config.governance.executorMultisig = address(proxyAdmin().owner());

        _deployProxies();
        _upgradeProxies();
        _initializeProxies();

        // Place native ETH first in `allStrats`
        // This ensures when we select a nonzero number of strategies from this array, we always have beacon chain ETH.
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);

        // Deploy and configure strategies and tokens, deploy half of them using the strategy factory.
        for (uint i = 1; i < NUM_LST_STRATS + 1; ++i) {
            string memory name = string.concat("LST-Strat", cheats.toString(i), " token");
            string memory symbol = string.concat("lstStrat", cheats.toString(i));
            _newStrategyAndToken(name, symbol, 10e50, address(this), i % 2 == 0);
        }

        // Whitelist the strategies
        cheats.prank(strategyManager().strategyWhitelister());
        strategyManager().addStrategiesToDepositWhitelist(allStrats);

        maxUniqueAssetsHeld = allStrats.length;

        // Create time machine and beacon chain. Set block time to beacon chain genesis time and starting block number
        BEACON_GENESIS_TIME = GENESIS_TIME_LOCAL;
        cheats.warp(BEACON_GENESIS_TIME);
        cheats.roll(10_000);
        _deployTimeMachineAndBeaconChain();
        // Set the `pectraForkTimestamp` on the EigenPodManager. Use pectra state
        cheats.startPrank(executorMultisig());
        eigenPodManager().setProofTimestampSetter(executorMultisig());
        eigenPodManager().setPectraForkTimestamp(BEACON_GENESIS_TIME);
        cheats.stopPrank();
    }

    function _setUpFork() public virtual {
        _setUpFork(true);
    }

    /// @dev Sets up the integration tests for mainnet.
    function _setUpFork(bool upgrade) public virtual {
        if (forkConfig.forkBlock != 0) cheats.createSelectFork(cheats.rpcUrl(profile), forkConfig.forkBlock);
        else cheats.createSelectFork(cheats.rpcUrl(profile));

        if (forkConfig.upgradeBeforeTesting && upgrade) {
            if (forkConfig.timelockPayload.length > 0 && eq(profile, "mainnet")) {
                _executeTimelockUpgrade();
            } else {
                _deployProxies(); // Only deploys what doesn't exist.
                _upgradeProxies();
            }
        }

        if (forkConfig.supportEigenPodTests) {
            allStrats.push(BEACONCHAIN_ETH_STRAT);
            allTokens.push(NATIVE_ETH);
            config.strategies.strategyAddresses.push(BEACONCHAIN_ETH_STRAT);
        }

        uint n = totalStrategies();
        for (uint i; i < n; ++i) {
            IStrategy strategy = strategyAddresses(i);

            if (strategy == BEACONCHAIN_ETH_STRAT) continue;
            IERC20 token = strategy.underlyingToken();
            if (tokensNotTested[address(token)]) continue;

            // Add to lstStrats and allStrats
            lstStrats.push(strategy);
            allStrats.push(strategy);
            allTokens.push(token);
        }

        maxUniqueAssetsHeld = allStrats.length;

        // Create time machine and mock beacon chain
        BEACON_GENESIS_TIME = GENESIS_TIME_MAINNET;
        _deployTimeMachineAndBeaconChain();

        if (eq(profile, "mainnet") && isUpgraded) _upgradeMainnetContracts();

        // Mainnet doesn't yet support this.
        if (isUpgraded && forkConfig.supportEigenPodTests) {
            // Set the `pectraForkTimestamp` on the EigenPodManager. Use pectra state
            cheats.startPrank(executorMultisig());
            eigenPodManager().setProofTimestampSetter(executorMultisig());
            eigenPodManager().setPectraForkTimestamp(BEACON_GENESIS_TIME);
            cheats.stopPrank();
        }
    }

    /// @dev Upgrades the mainnet contracts.
    function _upgradeMainnetContracts() public virtual {
        cheats.startPrank(executorMultisig());
        _upgradeProxies();
        cheats.stopPrank();
    }

    function _executeTimelockUpgrade() internal {
        // Warp forward to elapse the timelock queue.
        cheats.warp(block.timestamp + 21 days);

        // Execute the timelock upgrade.
        cheats.startPrank(protocolCouncil());
        timelock().execute({target: executorMultisig(), value: 0, payload: forkConfig.timelockPayload, predecessor: 0, salt: 0});
        cheats.stopPrank();

        console.log("SIMULATED TIMELOCK UPGRADE".yellow().bold());

        isUpgraded = true;
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @dev Returns a new transparent proxy without an implementation set.
    function _emptyProxy() internal returns (address) {
        return address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin()), ""));
    }

    /// @dev Upgrades a proxy to a new implementation.
    function _upgradeProxy(address proxy, address implementation) internal {
        proxyAdmin().upgrade(ITransparentUpgradeableProxy(payable(proxy)), implementation);
    }

    /// @dev Deploys a new transparent proxy without an implementation set for each contract in the system.
    function _deployProxies() internal {
        // Core contracts
        if (address(config.core.allocationManager) == address(0)) config.core.allocationManager = AllocationManager(_emptyProxy());
        if (address(config.core.avsDirectory) == address(0)) config.core.avsDirectory = AVSDirectory(_emptyProxy());
        if (address(config.core.delegationManager) == address(0)) config.core.delegationManager = DelegationManager(_emptyProxy());
        if (address(config.core.permissionController) == address(0)) config.core.permissionController = PermissionController(_emptyProxy());
        if (address(config.core.rewardsCoordinator) == address(0)) config.core.rewardsCoordinator = RewardsCoordinator(_emptyProxy());
        if (address(config.core.strategyManager) == address(0)) config.core.strategyManager = StrategyManager(_emptyProxy());
        // Pod contracts
        if (address(config.pods.eigenPodBeacon) == address(0)) config.pods.eigenPodBeacon = new UpgradeableBeacon(address(emptyContract));
        if (address(config.pods.eigenPodManager) == address(0)) config.pods.eigenPodManager = EigenPodManager(_emptyProxy());
        // Strategy contracts
        if (address(config.strategies.strategyFactory) == address(0)) config.strategies.strategyFactory = StrategyFactory(_emptyProxy());
        if (address(config.strategies.strategyFactoryBeacon) == address(0)) {
            config.strategies.strategyFactoryBeacon = new UpgradeableBeacon(address(emptyContract));
        }
    }

    /// @dev Upgrades all proxies to their implementation contracts.
    function _upgradeProxies() public {
        // Core contracts
        _upgradeProxy(
            address(allocationManager()),
            address(
                new AllocationManager(
                    delegationManager(),
                    pauserRegistry(),
                    permissionController(),
                    DEALLOCATION_DELAY,
                    ALLOCATION_CONFIGURATION_DELAY,
                    SEMVER
                )
            )
        );
        _upgradeProxy(address(avsDirectory()), address(new AVSDirectory(delegationManager(), pauserRegistry(), SEMVER)));
        _upgradeProxy(
            address(delegationManager()),
            address(
                new DelegationManager(
                    strategyManager(),
                    eigenPodManager(),
                    allocationManager(),
                    pauserRegistry(),
                    permissionController(),
                    DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
                    SEMVER
                )
            )
        );
        _upgradeProxy(address(permissionController()), address(new PermissionController(SEMVER)));
        _upgradeProxy(
            address(rewardsCoordinator()),
            address(
                new RewardsCoordinator(
                    IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams(
                        delegationManager(),
                        strategyManager(),
                        allocationManager(),
                        pauserRegistry(),
                        permissionController(),
                        REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
                        REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
                        REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
                        REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
                        REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP,
                        SEMVER
                    )
                )
            )
        );
        _upgradeProxy(address(strategyManager()), address(new StrategyManager(delegationManager(), pauserRegistry(), SEMVER)));

        // Pod contracts
        eigenPodBeacon().upgradeTo(address(new EigenPod(DEPOSIT_CONTRACT, eigenPodManager(), BEACON_GENESIS_TIME, "v9.9.9")));
        _upgradeProxy(
            address(eigenPodManager()),
            address(new EigenPodManager(DEPOSIT_CONTRACT, eigenPodBeacon(), delegationManager(), pauserRegistry(), "v9.9.9"))
        );

        // Strategy contracts
        _upgradeProxy(address(strategyFactory()), address(new StrategyFactory(strategyManager(), pauserRegistry(), "v9.9.9")));
        address baseStrategyImpl = address(new StrategyBase(strategyManager(), pauserRegistry(), "v9.9.9"));
        strategyFactoryBeacon().upgradeTo(baseStrategyImpl);
        for (uint i = 0; i < totalStrategies(); ++i) {
            IStrategy strategy = strategyAddresses(i);
            if (strategyAddresses(i) == BEACONCHAIN_ETH_STRAT) continue;
            _upgradeProxy(address(strategyAddresses(i)), address(baseStrategyImpl));
        }

        isUpgraded = true;
    }

    /// @dev Initializes all proxies.
    function _initializeProxies() public {
        address executorMultisig = executorMultisig();
        allocationManager().initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        avsDirectory().initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        delegationManager().initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        eigenPodManager().initialize({initialOwner: executorMultisig, _initPausedStatus: 0});
        rewardsCoordinator().initialize({
            initialOwner: executorMultisig,
            initialPausedStatus: 0,
            _rewardsUpdater: executorMultisig,
            _activationDelay: 0,
            _defaultSplitBips: 0
        });
        strategyFactory().initialize({_initialOwner: executorMultisig, _initialPausedStatus: 0, _strategyBeacon: strategyFactoryBeacon()});
        strategyManager().initialize({
            initialOwner: executorMultisig,
            initialStrategyWhitelister: address(strategyFactory()),
            initialPausedStatus: 0
        });
    }

    /// @dev Deploys a new strategy and token with given parameters.
    function _newStrategyAndToken(string memory tokenName, string memory tokenSymbol, uint initialSupply, address owner, bool useFactory)
        internal
    {
        // Deploy mock token, avoid using OZ for test speed.
        MockERC20 token = new MockERC20();
        token.initialize(tokenName, tokenSymbol, 18);
        IERC20 underlyingToken = IERC20(address(token));
        deal(address(underlyingToken), address(owner), initialSupply);
        // Deploy strategy using factory or directly.
        StrategyBase strategy = useFactory
            ? StrategyBase(address(strategyFactory().deployNewStrategy(underlyingToken)))
            : StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(strategyFactoryBeacon().implementation()),
                        address(proxyAdmin()),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                    )
                )
            );
        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
        config.strategies.strategyAddresses.push(strategy);
    }

    /// @dev Deploy the time machine and beacon chain to fixed addresses.
    function _deployTimeMachineAndBeaconChain() internal virtual {
        cheats.pauseTracing();
        cheats.etch(address(timeMachine), type(TimeMachine).runtimeCode);
        cheats.etch(address(beaconChain), type(BeaconChainMock).runtimeCode);
        cheats.label(address(timeMachine), "TimeMachine");
        cheats.label(address(beaconChain), "BeaconChain");
        cheats.allowCheatcodes(address(timeMachine));
        cheats.allowCheatcodes(address(beaconChain));
        beaconChain.initialize(eigenPodManager(), BEACON_GENESIS_TIME);
        cheats.resumeTracing();
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @dev Creates a new user with a random config.
    function _randUser(string memory name) internal returns (User user, IStrategy[] memory strategies, uint[] memory tokenBalances) {
        (uint userType, uint assetType) = (_randUserType(), _randAssetType());
        user = _genRandUser(name, userType);
        strategies = _selectRandAssets(assetType);
        tokenBalances = _dealRandAmounts(user, strategies);
        print.user(name, assetType, userType, strategies, tokenBalances);
    }

    /// @dev Creates a new user with a random config and a given list of strategies.
    function _randUser(string memory name, IStrategy[] memory strategies) internal returns (User user, uint[] memory tokenBalances) {
        uint userType = _randUserType();
        user = _genRandUser(name, userType);
        tokenBalances = _dealRandAmounts(user, strategies);
        print.user(name, HOLDS_ALL, userType, strategies, tokenBalances);
    }

    /// @dev Creates a new user without any assets
    function _randUser_NoAssets(string memory name) internal returns (User user) {
        uint userType = _randUserType();
        user = _genRandUser(name, userType);
        print.user(name, NO_ASSETS, userType, new IStrategy[](0), new uint[](0));
    }

    /// @dev Generates a new user with a given name and user type.
    function _genRandUser(string memory name, uint userType) internal returns (User user) {
        assertTrue(userType == DEFAULT || userType == ALT_METHODS, "_randUser: unimplemented userType");

        // string memory profile = FOUNDRY_PROFILE();

        user = userType == DEFAULT ? new User(name) : User(new User_AltMethods(name));

        assertTrue(address(user) != address(0), "User is not initialized");
        // if (eq(profile, "default") || eq(profile, "mainnet") || (eq(profile, "forktest") && isUpgraded)) {
        //     user = userType == DEFAULT ? new User(name) : User(new User_AltMethods(name));
        // } else if (eq(profile, "forktest") && !isUpgraded) {
        //     user = User(new User_M2(name));
        // }
    }

    /// @dev Generates a new AVS.
    function _genRandAVS(string memory name) internal returns (AVS) {
        return new AVS(name);
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// Given an assetType, select strategies the user will be dealt assets in
    function _selectRandAssets(uint assetType) internal returns (IStrategy[] memory strategies) {
        if (assetType == NO_ASSETS) return new IStrategy[](0);
        if (assetType == HOLDS_ETH) return BEACONCHAIN_ETH_STRAT.toArray();
        // Select number of assets:
        // HOLDS_LST can hold at most all LSTs. HOLDS_ALL and HOLDS_MAX also hold ETH.
        // Clamp number of assets to maxUniqueAssetsHeld (guaranteed to be at least 1)
        uint assetPoolSize = assetType == HOLDS_LST ? lstStrats.length : allStrats.length;
        uint maxAssets = assetPoolSize > maxUniqueAssetsHeld ? maxUniqueAssetsHeld : assetPoolSize;
        uint numAssets = assetType == HOLDS_MAX ? maxAssets : _randUint(1, maxAssets);

        strategies = new IStrategy[](numAssets);

        for (uint i = 0; i < numAssets; ++i) {
            strategies[i] = assetType == HOLDS_LST ? lstStrats[i] : allStrats[i];
        }

        return strategies.sort();
    }

    /// Given an input list of strategies, deal random underlying token amounts to a user
    function _dealRandAmounts(User user, IStrategy[] memory strategies) internal returns (uint[] memory tokenBalances) {
        assertTrue(address(user) != address(0), "User is not initialized");
        tokenBalances = new uint[](strategies.length);
        for (uint i = 0; i < tokenBalances.length; ++i) {
            IStrategy strategy = strategies[i];
            uint balance;

            if (strategy == BEACONCHAIN_ETH_STRAT) {
                // Award the user with a random amount of ETH
                // This guarantees a multiple of 32 ETH (at least 1, up to/incl 2080)
                uint amount = 32 ether * _randUint({min: 1, max: 65});
                balance = 32 ether * _randUint({min: 1, max: 5});
                cheats.deal(address(user), balance);
            } else {
                balance = _randUint({min: MIN_BALANCE, max: MAX_BALANCE});
                StdCheats.deal(address(strategy.underlyingToken()), address(user), balance);
            }

            tokenBalances[i] = balance;
        }
    }

    function _randUint(uint min, uint max) internal returns (uint) {
        return cheats.randomUint(min, max);
    }

    function _randBool() internal returns (bool) {
        return cheats.randomBool();
    }

    function _randAssetType() internal returns (uint) {
        bool supportsEth = !eq(profile, "sepolia") && !eq(profile, "hoodi");

        uint[] memory options = new uint[](supportsEth ? 5 : 2); // We have 2-5 possible asset types
        uint count = 0;

        if (assetTypes & NO_ASSETS != 0) options[count++] = NO_ASSETS;
        if (assetTypes & HOLDS_LST != 0) options[count++] = HOLDS_LST;

        if (supportsEth) {
            if (assetTypes & HOLDS_ETH != 0) options[count++] = HOLDS_ETH;
            if (assetTypes & HOLDS_ALL != 0) options[count++] = HOLDS_ALL;
            if (assetTypes & HOLDS_MAX != 0) options[count++] = HOLDS_MAX;
        }

        // If there are no options, skip the test.
        if (count == 0) cheats.skip(true);

        return options[cheats.randomUint(0, count - 1)];
    }

    function _randUserType() internal returns (uint) {
        uint[] memory options = new uint[](2); // We have 2 possible user types
        uint count = 0;
        if (userTypes & DEFAULT != 0) options[count++] = DEFAULT;
        if (userTypes & ALT_METHODS != 0) options[count++] = ALT_METHODS;
        return options[cheats.randomUint(0, count - 1)];
    }

    function NAME() public view virtual override returns (string memory) {
        return "Integration Deployer";
    }
}
