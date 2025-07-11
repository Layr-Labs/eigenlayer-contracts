// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

// Imports
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/AllocationManager.sol";
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

import "script/utils/ExistingDeploymentParser.sol";

IStrategy constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

abstract contract IntegrationDeployer is ExistingDeploymentParser {
    using StdStyle for *;
    using ArrayLib for *;

    // Fork ids for specific fork tests
    bool isUpgraded;
    uint mainnetForkBlock = 22_896_603; // Post Redistribution upgrade

    string version = "9.9.9";

    // Beacon chain genesis time when running locally
    // Multiple of 12 for sanity's sake
    uint64 constant GENESIS_TIME_LOCAL = 1 hours * 12;
    uint64 constant GENESIS_TIME_MAINNET = 1_606_824_023;
    uint64 BEACON_GENESIS_TIME; // set after forkType is decided

    // Beacon chain deposit contract. The BeaconChainMock contract etchs ETHPOSDepositMock code here.
    IETHPOSDeposit constant DEPOSIT_CONTRACT = IETHPOSDeposit(0x00000000219ab540356cBB839Cbe05303d7705Fa);

    uint8 constant NUM_LST_STRATS = 32;

    uint32 INITIAL_GLOBAL_DELAY_BLOCKS = 4 days / 12 seconds; // 4 days in blocks

    // Lists of strategies used in the system
    //
    // When we select random user assets, we use the `assetType` to determine
    // which of these lists to select user assets from.
    IStrategy[] lstStrats;
    IStrategy[] ethStrats; // only has one strat tbh
    IStrategy[] allStrats; // just a combination of the above 2 lists
    IERC20[] allTokens; // `allStrats`, but contains all of the underlying tokens instead
    uint maxUniqueAssetsHeld;

    // If a token is in this mapping, then we will ignore this LST as it causes issues with reading balanceOf
    mapping(address => bool) public tokensNotTested;

    // Mock Contracts to deploy
    TimeMachine public timeMachine;
    BeaconChainMock public beaconChain;

    // Admin Addresses
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

    /// @dev used to configure randomness and default user/asset types
    ///
    /// Tests that want alternate user/asset types can still use this modifier,
    /// and then configure user/asset types individually using the methods:
    /// _configAssetTypes(...)
    /// _configUserTypes(...)
    ///
    /// (Alternatively, this modifier can be overwritten)
    modifier rand(uint24 r) virtual {
        _configRand({_randomSeed: r, _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL, _userTypes: DEFAULT | ALT_METHODS});

        // Used to create shared setups between tests
        _init();

        _;
    }

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
        bool forkMainnet = isForktest();

        if (forkMainnet) {
            console.log("Forking mainnet");
            forkType = MAINNET;
            _setUpMainnet();
        } else {
            forkType = LOCAL;
            _setUpLocal();
        }
    }

    /// @dev Used to create shared setup between tests. This method is called
    /// when the `rand` modifier is run, before a test starts
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
        pausers[0] = pauser;
        eigenLayerPauserReg = new PauserRegistry(pausers, unpauser);

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
        ethStrats.push(BEACONCHAIN_ETH_STRAT);
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
        timeMachine = new TimeMachine();
        beaconChain = new BeaconChainMock(eigenPodManager, BEACON_GENESIS_TIME);

        // Set the `pectraForkTimestamp` on the EigenPodManager. Use pectra state
        cheats.startPrank(executorMultisig);
        eigenPodManager.setProofTimestampSetter(executorMultisig);
        eigenPodManager.setPectraForkTimestamp(BEACON_GENESIS_TIME);
        cheats.stopPrank();
    }

    /// Parse existing contracts from mainnet
    function _setUpMainnet() public virtual noTracing {
        console.log("Setting up `%s` integration tests:", "MAINNET_FORK".green().bold());
        console.log("RPC:", cheats.rpcUrl("mainnet"));
        console.log("Block:", mainnetForkBlock);

        cheats.createSelectFork(cheats.rpcUrl("mainnet"), mainnetForkBlock);

        string memory deploymentInfoPath = "script/configs/mainnet/mainnet-addresses.config.json";
        _parseDeployedContracts(deploymentInfoPath);
        string memory existingDeploymentParams = "script/configs/mainnet.json";
        _parseParamsForIntegrationUpgrade(existingDeploymentParams);

        // Place native ETH first in `allStrats`
        // This ensures when we select a nonzero number of strategies from this array, we always
        // have beacon chain ETH
        ethStrats.push(BEACONCHAIN_ETH_STRAT);
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
        timeMachine = new TimeMachine();
        beaconChain = new BeaconChainMock(eigenPodManager, BEACON_GENESIS_TIME);

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

    function _upgradeMainnetContracts() public virtual {
        // Warp time to past the slashing queue transaction
        cheats.warp(block.timestamp + 10 days);

        // cast tx 0x8e6f1701abc942d468a5cea427ae16069b5ee6341407e2f3ac64e18e01e06756 --flashbots
        bytes memory payload =
            hex"6a761202000000000000000000000000c06fd4f821eac1ff1ae8067b36342899b57baa2d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f800000000000000000000000000000000000000000000000000000000000000e0401d5062a000000000000000000000000369e6f597e22eab55ffb173c6d9cd234bd699111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d2f000000000000000000000000000000000000000000000000000000000000000d046a76120200000000000000000000000040a2accbd92bca938b02010e17a5b8929b49130d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c800000000000000000000000000000000000000000000000000000000000000b048d80ff0a00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000aa2008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000039053d51b77dc0d36036fc1fcc8cb819df8ef37a0000000000000000000000006eed6c2802df347e05884857cddb2d3e96d12f89008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000948a420b8cc1d6bfd0b6087c2e7c344a2cd0bc39000000000000000000000000c97602648fa52f92b4ee2b0e5a54bd15b6cb0345008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000858646372cc42e1a627fce94aa7a7033e7cf075a00000000000000000000000046aefd30415be99e20169ee7046f65784b46d123008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000091e677b07f7af907ec9a428aafa9fc14a0d3a338000000000000000000000000e48d7caec1790b293667e4bb2de1e00536f2babd008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000acb55c530acdb2849e6d4f36992cd8c9d50ed8f7000000000000000000000000530fdb7adf7d489df49c27e3d3512c0dd64886be000ed6703c298d28ae0878d1b28e88ca87f9662fe9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659cfe6000000000000000000000000d4d1746142642db4c1ab17b03b9c58baac913e5b008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000093c4b944d05dfe6df7645a86cd2206016c51564d00000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000001bee69b7dfffa4e2d53c2a2df135c388ad25dcd200000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000054945180db7943c0ed0fee7edab2bd24620256bc00000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000009d7ed45ee2e8fc5482fa2428f15c971e6369011d00000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000013760f50a9d7377e4f20cb8cf9e4c26586c658ff00000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000a4c637e0f704745d182e4d38cab7e7485321d05900000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec400000000000000000000000057ba429517c3473b6d34ca9acd56c0e735b94c0200000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000000fe4f44bee93503346a3ac9ee5a26b130a5796d600000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000007ca911e83dabf90c90dd3de5411a10f1a611218400000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec40000000000000000000000008ca7a5d6f3acd3a7a8bc468a8cd0fb14b6bd28b600000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000ae60d8180437b5c34bb956822ac271097258447300000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0008b9566ada63b64d1e1dcf1418b43fd1433b724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499a88ec4000000000000000000000000298afb19a105d59e74658c4c334ff360bade6dd200000000000000000000000062f7226fb9d615590eadb539713b250fb2fdf4e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000c06fd4f821eac1ff1ae8067b36342899b57baa2d00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c3bb49554b2a3b939cbafdd44d3c7b917c194e00ad787577d9d7a2200101e85c725db510701d317d1e0c9777e07274cadb70599085e5ac8067f0dfec96b77c6a611b5742f082349d14b95370052049df9eeb66ea4b2800997861a11d6c129d4dbdf971d274f313d02a097d96b65829cae892b514cd49cf8798672c6b049a76716bc61c000000000000000000000000e31ad7cfd94bd74c40b53160aa0e8a0b6d3408300000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000";

        cheats.startPrank(protocolCouncilMultisig);
        timelockController.execute({target: executorMultisig, value: 0, payload: payload, predecessor: 0, salt: 0});
        cheats.stopPrank();
    }

    function _deployProxies() public {
        if (address(delegationManager) == address(0)) {
            delegationManager =
                DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(strategyManager) == address(0)) {
            strategyManager =
                StrategyManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(eigenPodManager) == address(0)) {
            eigenPodManager =
                EigenPodManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(rewardsCoordinator) == address(0)) {
            rewardsCoordinator =
                RewardsCoordinator(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(avsDirectory) == address(0)) {
            avsDirectory = AVSDirectory(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(strategyFactory) == address(0)) {
            strategyFactory =
                StrategyFactory(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(allocationManager) == address(0)) {
            allocationManager =
                AllocationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(permissionController) == address(0)) {
            permissionController =
                PermissionController(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        }
        if (address(eigenPodBeacon) == address(0)) eigenPodBeacon = new UpgradeableBeacon(address(emptyContract));
        if (address(strategyBeacon) == address(0)) strategyBeacon = new UpgradeableBeacon(address(emptyContract));
    }

    /// Deploy an implementation contract for each contract in the system
    function _deployImplementations() public {
        allocationManagerImplementation = new AllocationManager(
            delegationManager,
            eigenStrategy,
            eigenLayerPauserReg,
            permissionController,
            DEALLOCATION_DELAY,
            ALLOCATION_CONFIGURATION_DELAY,
            version
        );
        permissionControllerImplementation = new PermissionController(version);
        delegationManagerImplementation = new DelegationManager(
            strategyManager,
            eigenPodManager,
            allocationManager,
            eigenLayerPauserReg,
            permissionController,
            DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS,
            version
        );
        strategyManagerImplementation = new StrategyManager(allocationManager, delegationManager, eigenLayerPauserReg, version);
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
                version: version
            })
        );
        avsDirectoryImplementation = new AVSDirectory(delegationManager, eigenLayerPauserReg, version);
        eigenPodManagerImplementation =
            new EigenPodManager(DEPOSIT_CONTRACT, eigenPodBeacon, delegationManager, eigenLayerPauserReg, "9.9.9");
        strategyFactoryImplementation = new StrategyFactory(strategyManager, eigenLayerPauserReg, "9.9.9");

        // Beacon implementations
        eigenPodImplementation = new EigenPod(DEPOSIT_CONTRACT, eigenPodManager, BEACON_GENESIS_TIME, "9.9.9");
        baseStrategyImplementation = new StrategyBase(strategyManager, eigenLayerPauserReg, "9.9.9");

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
        delegationManager.initialize({initialPausedStatus: 0});

        strategyManager.initialize({
            initialOwner: executorMultisig,
            initialStrategyWhitelister: address(strategyFactory),
            initialPausedStatus: 0
        });

        eigenPodManager.initialize({initialOwner: executorMultisig, _initPausedStatus: 0});

        avsDirectory.initialize({initialOwner: executorMultisig, initialPausedStatus: 0});

        allocationManager.initialize({initialPausedStatus: 0});

        strategyFactory.initialize({_initialOwner: executorMultisig, _initialPausedStatus: 0, _strategyBeacon: strategyBeacon});
    }

    /// @dev Deploy a strategy and its underlying token, push to global lists of tokens/strategies, and whitelist
    /// strategy in strategyManager
    function _newStrategyAndToken(string memory tokenName, string memory tokenSymbol, uint initialSupply, address owner, bool useFactory)
        internal
        noTracing
    {
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

        cheats.prank(strategyManager.strategyWhitelister());
        strategyManager.addStrategiesToDepositWhitelist(strategies);

        // Add to lstStrats and allStrats
        lstStrats.push(strategy);
        allStrats.push(strategy);
        allTokens.push(underlyingToken);
    }

    function _configRand(uint24 _randomSeed, uint _assetTypes, uint _userTypes) private noTracing {
        // Using uint24 for the seed type so that if a test fails, it's easier
        // to manually use the seed to replay the same test.
        random = _hash(_randomSeed);

        // Convert flag bitmaps to bytes of set bits for easy use with _randUint
        _configAssetTypes(_assetTypes);
        _configUserTypes(_userTypes);
    }

    function _configAssetTypes(uint _assetTypes) internal {
        assetTypes = _bitmapToBytes(_assetTypes);
        assertTrue(assetTypes.length != 0, "_configRand: no asset types selected");
    }

    function _configAssetAmounts(uint _maxUniqueAssetsHeld) internal {
        if (_maxUniqueAssetsHeld > allStrats.length) _maxUniqueAssetsHeld = allStrats.length;

        maxUniqueAssetsHeld = _maxUniqueAssetsHeld;
        require(maxUniqueAssetsHeld != 0, "_configAssetAmounts: invalid 0");
    }

    function _configUserTypes(uint _userTypes) internal {
        userTypes = _bitmapToBytes(_userTypes);
        assertTrue(userTypes.length != 0, "_configRand: no user types selected");
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

    /// @dev Create a new user without native ETH. See _randUser above for standard usage
    function _randUser_NoETH(string memory name) internal noTracing returns (User, IStrategy[] memory, uint[] memory) {
        // Deploy new User contract
        uint userType = _randUserType();
        User user = _genRandUser(name, userType);

        // Pick the user's asset distribution, removing "native ETH" as an option
        // I'm sorry if this eventually leads to a bug that's really hard to track down
        uint assetType = _randAssetType();
        if (assetType == HOLDS_ETH) assetType = NO_ASSETS;
        else if (assetType == HOLDS_ALL || assetType == HOLDS_MAX) assetType = HOLDS_LST;

        // For the specific asset selection we made, get a random assortment of strategies
        // and deal the user some corresponding underlying token balances
        IStrategy[] memory strategies = _selectRandAssets(assetType);
        uint[] memory tokenBalances = _dealRandAmounts(user, strategies);

        print.user(name, assetType, userType, strategies, tokenBalances);
        return (user, strategies, tokenBalances);
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
            // Leaving this if statement for future upgraded users
        } else if (forkType == MAINNET && !isUpgraded) {
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

    function _genRandAVS(string memory name) internal returns (AVS avs) {
        if (forkType == LOCAL) avs = new AVS(name);
        else if (forkType == MAINNET) avs = new AVS(name);
        else revert("_genRandAVS: unimplemented forkType");
    }

    /// Given an assetType, select strategies the user will be dealt assets in
    function _selectRandAssets(uint assetType) internal noTracing returns (IStrategy[] memory) {
        if (assetType == NO_ASSETS) return new IStrategy[](0);

        /// Select only ETH
        if (assetType == HOLDS_ETH) return beaconChainETHStrategy.toArray();

        /// Select multiple LSTs, and maybe add ETH:

        // Select number of assets:
        // HOLDS_LST can hold at most all LSTs. HOLDS_ALL and HOLDS_MAX also hold ETH.
        // Clamp number of assets to maxUniqueAssetsHeld (guaranteed to be at least 1)
        uint assetPoolSize = assetType == HOLDS_LST ? lstStrats.length : allStrats.length;
        uint maxAssets = assetPoolSize > maxUniqueAssetsHeld ? maxUniqueAssetsHeld : assetPoolSize;

        uint numAssets = assetType == HOLDS_MAX ? maxAssets : _randUint(1, maxAssets);

        IStrategy[] memory strategies = new IStrategy[](numAssets);
        for (uint i = 0; i < strategies.length; i++) {
            if (assetType == HOLDS_LST) {
                strategies[i] = lstStrats[i];
            } else {
                // allStrats[0] is the beaconChainETHStrategy
                strategies[i] = allStrats[i];
            }
        }

        return strategies;
    }

    /// Given an input list of strategies, deal random underlying token amounts to a user
    function _dealRandAmounts(User user, IStrategy[] memory strategies) internal noTracing returns (uint[] memory) {
        uint[] memory tokenBalances = new uint[](strategies.length);

        for (uint i = 0; i < tokenBalances.length; i++) {
            IStrategy strategy = strategies[i];
            uint balance;

            if (strategy == BEACONCHAIN_ETH_STRAT) {
                // Award the user with a random amount of ETH
                // This guarantees a multiple of 32 ETH (at least 1, up to/incl 2080)
                uint amount = 32 ether * _randUint({min: 1, max: 65});
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

            if (strategy == BEACONCHAIN_ETH_STRAT) {
                cheats.deal(address(user), amounts[i]);
            } else {
                IERC20 underlyingToken = strategy.underlyingToken();
                StdCheats.deal(address(underlyingToken), address(user), amounts[i]);
            }
        }
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
        while (value >= range) value = (value - range) & mask;

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

    function _randomStrategies() internal returns (IStrategy[][] memory strategies) {
        uint numOpSets = _randUint({min: 1, max: 5});

        strategies = new IStrategy[][](numOpSets);

        for (uint i; i < numOpSets; ++i) {
            IStrategy[] memory randomStrategies = allStrats.shuffle();
            uint numStrategies = _randUint({min: 1, max: maxUniqueAssetsHeld});

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
    function _bitmapToBytes(uint bitmap) internal pure returns (bytes memory bytesArray) {
        for (uint i = 0; i < 256; ++i) {
            // Mask for i-th bit
            uint mask = uint(1 << i);

            // If the i-th bit is flipped, add a byte to the return array
            if (bitmap & mask != 0) bytesArray = bytes.concat(bytesArray, bytes1(uint8(1 << i)));
        }
        return bytesArray;
    }

    function _hash(string memory x) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }

    function _hash(uint x) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(x));
    }
}
