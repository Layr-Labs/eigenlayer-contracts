// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {MockERC20} from "forge-std/mocks/MockERC20.sol";

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ICoreTypes} from "src/contracts/interfaces/ICore.sol";

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

    string constant deploymentInfoPath = "script/configs/mainnet/mainnet-addresses.config.json";
    string constant existingDeploymentParams = "script/configs/mainnet.json";

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
        tokensNotTested[address(0x3F1c547b21f65e10480dE3ad8E19fAAC46C95034)] = true; // stETH holesky
        tokensNotTested[address(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84)] = true; // stETH mainnet
        tokensNotTested[address(0x856c4Efb76C1D1AE02e20CEB03A2A6a08b0b8dC3)] = true; // oETH mainnet
        tokensNotTested[address(0xF603c5A3F774F05d4D848A9bB139809790890864)] = true; // osETH holesky
        tokensNotTested[address(0xf1C9acDc66974dFB6dEcB12aA385b9cD01190E38)] = true; // osETH mainnet
        tokensNotTested[address(0x8720095Fa5739Ab051799211B146a2EEE4Dd8B37)] = true; // cbETH holesky
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
        emptyContract = new EmptyContract();
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

    /// @dev Configures the possible asset and user types for the random users.
    function _configRand(uint _assetTypes, uint _userTypes) private noTracing {
        _configAssetTypes(_assetTypes);
        _configUserTypes(_userTypes);
    }

    /// @dev Configures the possible asset types for the random users.
    function _configAssetTypes(uint _assetTypes) internal {
        assertTrue((assetTypes = _assetTypes) != 0, "_configAssetTypes: no asset types selected");
    }

    /// @dev Configures the possible user types for the random users.
    function _configUserTypes(uint _userTypes) internal {
        assertTrue((userTypes = _userTypes) != 0, "_configUserTypes: no user types selected");
    }

    /// @dev Configures the maximum number of unique assets a user can hold.
    function _configAssetAmounts(uint _maxUniqueAssetsHeld) internal {
        if (_maxUniqueAssetsHeld > allStrats.length) _maxUniqueAssetsHeld = allStrats.length;
        assertTrue((maxUniqueAssetsHeld = _maxUniqueAssetsHeld) != 0, "_configAssetAmounts: invalid 0");
    }

    /// @dev Run `FOUNDRY_PROFILE=forktest forge t --mc Integration` for integration fork testing.
    function isForktest() public view returns (bool) {
        return _hash("forktest") == _hash(cheats.envOr(string("FOUNDRY_PROFILE"), string("default")));
    }

    /// -----------------------------------------------------------------------
    /// Environment Setup
    /// -----------------------------------------------------------------------

    /// @dev Sets up the integration tests for local.
    function _setUpLocal() public virtual noTracing {
        console.log("Setting up `%s` integration tests:", "LOCAL".yellow().bold());
        // Deploy ProxyAdmin, PauserRegistry, and executorMultisig.
        eigenLayerProxyAdmin = new ProxyAdmin();
        eigenLayerPauserReg = new PauserRegistry(PAUSER.toArray(), UNPAUSER);
        executorMultisig = address(eigenLayerProxyAdmin.owner());

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
        cheats.prank(strategyManager.strategyWhitelister());
        strategyManager.addStrategiesToDepositWhitelist(allStrats);

        maxUniqueAssetsHeld = allStrats.length;

        // Create time machine and beacon chain. Set block time to beacon chain genesis time and starting block number
        BEACON_GENESIS_TIME = GENESIS_TIME_LOCAL;
        cheats.warp(BEACON_GENESIS_TIME);
        cheats.roll(10_000);
        _deployTimeMachineAndBeaconChain();
        // Set the `pectraForkTimestamp` on the EigenPodManager. Use pectra state
        cheats.startPrank(executorMultisig);
        eigenPodManager.setProofTimestampSetter(executorMultisig);
        eigenPodManager.setPectraForkTimestamp(BEACON_GENESIS_TIME);
        cheats.stopPrank();
    }

    /// @dev Sets up the integration tests for mainnet.
    function _setUpMainnet() public virtual noTracing {
        console.log("Setting up `%s` integration tests:", "MAINNET_FORK".green().bold());
        console.log("Block:", MAINNET_FORK_BLOCK);

        cheats.createSelectFork(cheats.rpcUrl("mainnet"), MAINNET_FORK_BLOCK);

        _parseDeployedContracts(deploymentInfoPath);
        _parseParamsForIntegrationUpgrade(existingDeploymentParams);

        // Place native ETH first in `allStrats`
        // This ensures when we select a nonzero number of strategies from this array, we always
        // have beacon chain ETH
        allStrats.push(BEACONCHAIN_ETH_STRAT);
        allTokens.push(NATIVE_ETH);

        // Add deployed strategies to lstStrats and allStrats
        for (uint i; i < deployedStrategyArray.length; ++i) {
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
        _deployTimeMachineAndBeaconChain();

        // Since we haven't done the slashing upgrade on mainnet yet, upgrade mainnet contracts
        // prior to test. `isUpgraded` is true by default, but is set to false in `UpgradeTest.t.sol`
        if (isUpgraded) {
            _upgradeMainnetContracts();

            // Set the `pectraForkTimestamp` on the EigenPodManager. Use pectra state
            cheats.startPrank(executorMultisig);
            eigenPodManager.setProofTimestampSetter(executorMultisig);
            eigenPodManager.setPectraForkTimestamp(BEACON_GENESIS_TIME);
            cheats.stopPrank();
        }
    }

    /// @dev Upgrades the mainnet contracts.
    function _upgradeMainnetContracts() public virtual {
        cheats.startPrank(executorMultisig);
        _deployImplementations();
        _upgradeProxies();
        cheats.stopPrank();
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @dev Returns a new transparent proxy without an implementation set.
    function _emptyProxy() internal returns (address) {
        return address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""));
    }

    /// @dev Upgrades a proxy to a new implementation.
    function _upgradeProxy(address proxy, address implementation) public {
        eigenLayerProxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(proxy)), implementation);
    }

    /// @dev Deploys a new transparent proxy without an implementation set for each contract in the system.
    function _deployProxies() public {
        delegationManager = DelegationManager(_emptyProxy());
        strategyManager = StrategyManager(_emptyProxy());
        eigenPodManager = EigenPodManager(_emptyProxy());
        rewardsCoordinator = RewardsCoordinator(_emptyProxy());
        avsDirectory = AVSDirectory(_emptyProxy());
        strategyFactory = StrategyFactory(_emptyProxy());
        allocationManager = AllocationManager(_emptyProxy());
        permissionController = PermissionController(_emptyProxy());
        eigenPodBeacon = new UpgradeableBeacon(address(emptyContract));
        strategyBeacon = new UpgradeableBeacon(address(emptyContract));
    }

    /// @dev Deploys the implementation of each contract in the system.
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

    /// @dev Upgrades all proxies to their implementation contracts.
    function _upgradeProxies() public noTracing {
        _upgradeProxy(address(allocationManager), address(allocationManagerImplementation));
        _upgradeProxy(address(avsDirectory), address(avsDirectoryImplementation));
        _upgradeProxy(address(delegationManager), address(delegationManagerImplementation));
        _upgradeProxy(address(eigenPodManager), address(eigenPodManagerImplementation));
        _upgradeProxy(address(permissionController), address(permissionControllerImplementation));
        _upgradeProxy(address(rewardsCoordinator), address(rewardsCoordinatorImplementation));
        _upgradeProxy(address(strategyFactory), address(strategyFactoryImplementation));
        _upgradeProxy(address(strategyManager), address(strategyManagerImplementation));
        eigenPodBeacon.upgradeTo(address(eigenPodImplementation));
        strategyBeacon.upgradeTo(address(baseStrategyImplementation));
        for (uint i = 0; i < numStrategiesDeployed; ++i) {
            _upgradeProxy(address(deployedStrategyArray[i]), address(baseStrategyImplementation));
        }
    }

    /// @dev Initializes all proxies.
    function _initializeProxies() public noTracing {
        allocationManager.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        avsDirectory.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        delegationManager.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        eigenPodManager.initialize({initialOwner: executorMultisig, _initPausedStatus: 0});
        // TODO Why is this missing?
        // rewardsCoordinator.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});
        strategyFactory.initialize({_initialOwner: executorMultisig, _initialPausedStatus: 0, _strategyBeacon: strategyBeacon});
        strategyManager.initialize({
            initialOwner: executorMultisig,
            initialStrategyWhitelister: address(strategyFactory),
            initialPausedStatus: 0
        });
    }

    /// @dev Deploys a new strategy and token with given parameters.
    function _newStrategyAndToken(string memory tokenName, string memory tokenSymbol, uint initialSupply, address owner, bool useFactory)
        internal
        noTracing
    {
        // Deploy mock token, avoid using OZ for test speed.
        MockERC20 token = new MockERC20();
        token.initialize(tokenName, tokenSymbol, 18);
        IERC20 underlyingToken = IERC20(address(token));
        deal(address(underlyingToken), address(owner), initialSupply);
        // Deploy strategy using factory or directly.
        StrategyBase strategy = useFactory
            ? StrategyBase(address(strategyFactory.deployNewStrategy(underlyingToken)))
            : StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                    )
                )
            );
        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
    }

    /// @dev Deploy the time machine and beacon chain to fixed addresses.
    function _deployTimeMachineAndBeaconChain() internal {
        cheats.pauseTracing();
        cheats.etch(address(timeMachine), type(TimeMachine).runtimeCode);
        cheats.etch(address(beaconChain), type(BeaconChainMock).runtimeCode);
        cheats.label(address(timeMachine), "TimeMachine");
        cheats.label(address(beaconChain), "BeaconChain");
        cheats.allowCheatcodes(address(timeMachine));
        cheats.allowCheatcodes(address(beaconChain));
        beaconChain.initialize(eigenPodManager, BEACON_GENESIS_TIME);
        cheats.resumeTracing();
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

    /// @dev Creates a new user with a random config.
    function _randUser(string memory name)
        internal
        noTracing
        returns (User user, IStrategy[] memory strategies, uint[] memory tokenBalances)
    {
        (uint userType, uint assetType) = (_randUserType(), _randAssetType());
        user = _genRandUser(name, userType);
        strategies = _selectRandAssets(assetType);
        tokenBalances = _dealRandAmounts(user, strategies);
        print.user(name, assetType, userType, strategies, tokenBalances);
    }

    /// @dev Creates a new user with a random config and a given list of strategies.
    function _randUser(string memory name, IStrategy[] memory strategies)
        internal
        noTracing
        returns (User user, uint[] memory tokenBalances)
    {
        uint userType = _randUserType();
        user = _genRandUser(name, userType);
        tokenBalances = _dealRandAmounts(user, strategies);
        print.user(name, HOLDS_ALL, userType, strategies, tokenBalances);
    }

    /// @dev Creates a new user without any assets
    function _randUser_NoAssets(string memory name) internal noTracing returns (User user) {
        uint userType = _randUserType();
        user = _genRandUser(name, userType);
        print.user(name, NO_ASSETS, userType, new IStrategy[](0), new uint[](0));
    }

    /// @dev Generates a new user with a given name and user type.
    function _genRandUser(string memory name, uint userType) internal returns (User user) {
        assertTrue(forkType == LOCAL || forkType == MAINNET, "_randUser: unimplemented forkType");
        assertTrue(userType == DEFAULT || userType == ALT_METHODS, "_randUser: unimplemented userType");
        if (forkType == LOCAL || (forkType == MAINNET && isUpgraded)) {
            user = userType == DEFAULT ? new User(name) : User(new User_AltMethods(name));
        } else if (forkType == MAINNET && !isUpgraded) {
            user = User(new User_M2(name));
        }
    }

    /// @dev Generates a new AVS.
    function _genRandAVS(string memory name) internal returns (AVS) {
        assertTrue(forkType == LOCAL || forkType == MAINNET, "_genRandAVS: unimplemented forkType");
        return new AVS(name);
    }

    /// -----------------------------------------------------------------------
    ///
    /// -----------------------------------------------------------------------

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

        for (uint i = 0; i < numAssets; ++i) {
            strategies[i] = assetType == HOLDS_LST ? lstStrats[i] : allStrats[i];
        }

        return strategies.sort();
    }

    /// Given an input list of strategies, deal random underlying token amounts to a user
    function _dealRandAmounts(User user, IStrategy[] memory strategies) internal noTracing returns (uint[] memory tokenBalances) {
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
