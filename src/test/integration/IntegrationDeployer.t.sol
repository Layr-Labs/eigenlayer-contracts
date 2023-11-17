// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

// Imports
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "forge-std/Test.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/core/Slasher.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/DelayedWithdrawalRouter.sol";
import "src/contracts/permissions/PauserRegistry.sol";

import "src/test/mocks/EmptyContract.sol";
import "src/test/mocks/ETHDepositMock.sol";
import "src/test/mocks/BeaconChainOracleMock.sol";

abstract contract IntegrationDeployer is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    // Core contracts to deploy
    DelegationManager delegationManager;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;
    PauserRegistry pauserRegistry;
    Slasher slasher;
    IBeacon eigenPodBeacon;
    EigenPod pod;
    DelayedWithdrawalRouter delayedWithdrawalRouter;

    // Base strategy implementation in case we want to create more strategies later
    StrategyBase baseStrategyImplementation;
    // Lists of strategies and corresponding underlying tokens
    IStrategy[] _strategies;
    IERC20[] _underlyingTokens;

    // Mock Contracts to deploy
    ETHPOSDepositMock ethPOSDeposit;
    BeaconChainOracleMock beaconChainOracle;

    // ProxyAdmin
    ProxyAdmin eigenLayerProxyAdmin;
    // Admin Addresses
    address eigenLayerReputedMultisig = address(this); // admin address
    address constant pauser = address(555);
    address constant unpauser = address(556);

    // Constants
    uint64 constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    uint64 constant GOERLI_GENESIS_TIME = 1616508000;
    IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    function setUp() public virtual {
        // Deploy ProxyAdmin
        eigenLayerProxyAdmin = new ProxyAdmin();

        // Deploy PauserRegistry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        // Deploy mocks
        EmptyContract emptyContract = new EmptyContract();
        ethPOSDeposit = new ETHPOSDepositMock();
        beaconChainOracle = new BeaconChainOracleMock();

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
        slasher = Slasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Deploy EigenPod Contracts
        pod = new EigenPod(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            eigenPodManager,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(pod));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        StrategyManager strategyManagerImplementation = new StrategyManager(delegationManager, eigenPodManager, slasher);
        Slasher slasherImplementation = new Slasher(strategyManager, delegationManager);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegationManager
        );
        DelayedWithdrawalRouter delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(eigenPodManager);

        // Third, upgrade the proxy contracts to point to the implementations
        // DelegationManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                pauserRegistry,
                0 // initialPausedStatus
            )
        );
        // StrategyManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                eigenLayerReputedMultisig, //initialOwner
                eigenLayerReputedMultisig, //initial whitelister
                pauserRegistry,
                0 // initialPausedStatus
            )
        );
        // Slasher
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(
                Slasher.initialize.selector,
                eigenLayerReputedMultisig,
                pauserRegistry,
                0 // initialPausedStatus
            )
        );
        // EigenPodManager
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                type(uint256).max, // maxPods
                address(beaconChainOracle),
                eigenLayerReputedMultisig, // initialOwner
                pauserRegistry,
                0 // initialPausedStatus
            )
        );
        // Delayed Withdrawal Router
        uint256 withdrawalDelayBlocks = 7 days / 12 seconds;
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation),
            abi.encodeWithSelector(
                DelayedWithdrawalRouter.initialize.selector,
                eigenLayerReputedMultisig, // initialOwner
                pauserRegistry,
                0, // initialPausedStatus
                withdrawalDelayBlocks
            )
        );

        // Create base strategy implementation and deploy a few strategies
        baseStrategyImplementation = new StrategyBase(strategyManager);

        _newStrategyAndToken("Strategy1Token", "str1", 10e50, address(this)); // initialSupply, owner
        _newStrategyAndToken("Strategy2Token", "str2", 10e50, address(this)); // initialSupply, owner
        _newStrategyAndToken("Strategy3Token", "str3", 10e50, address(this)); // initialSupply, owner
    }

    /// @dev Deploy a strategy and its underlying token, push to global lists of tokens/strategies, and whitelist
    /// strategy in strategyManager
    function _newStrategyAndToken(string memory tokenName, string memory tokenSymbol, uint initialSupply, address owner) internal {
        IERC20 underlyingToken = new ERC20PresetFixedSupply(tokenName, tokenSymbol, 10e50, address(this)); 
        StrategyBase strategy = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, pauserRegistry)
                )
            )
        );

        // Push to list of tokens/strategies
        _underlyingTokens.push(underlyingToken);
        _strategies.push(strategy);

        // Whitelist strategy
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategy;
        strategyManager.addStrategiesToDepositWhitelist(strategies);
    }
}