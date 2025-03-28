// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/strategies/StrategyFactory.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "../../contracts/permissions/PauserRegistry.sol";

/**
 * @notice Unit testing of the StrategyFactory contract.
 * Contracts tested: StrategyFactory
 */
contract StrategyFactoryUnitTests is EigenLayerUnitTestSetup {
    // Contract under test
    StrategyFactory public strategyFactory;
    StrategyFactory public strategyFactoryImplementation;

    // Contract dependencies
    StrategyBase public strategyImplementation;
    UpgradeableBeacon public strategyBeacon;
    ERC20PresetFixedSupply public underlyingToken;

    uint initialSupply = 1e36;
    address initialOwner = address(this);
    address beaconProxyOwner = address(this);
    address notOwner = address(7_777_777);

    uint initialPausedStatus = 0;

    /// @notice Emitted when the `strategyBeacon` is changed
    event StrategyBeaconModified(IBeacon previousBeacon, IBeacon newBeacon);

    /// @notice Emitted whenever a slot is set in the `deployedStrategies` mapping
    event StrategySetForToken(IERC20 token, IStrategy strategy);

    event TokenBlacklisted(IERC20 token);

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        underlyingToken = new ERC20PresetFixedSupply("Test Token", "TEST", initialSupply, initialOwner);

        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "v9.9.9");

        strategyBeacon = new UpgradeableBeacon(address(strategyImplementation));
        strategyBeacon.transferOwnership(beaconProxyOwner);

        strategyFactoryImplementation = new StrategyFactory(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "v9.9.9");

        strategyFactory = StrategyFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyFactoryImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyFactory.initialize.selector, initialOwner, initialPausedStatus, strategyBeacon)
                )
            )
        );
    }

    function test_initialization() public view {
        assertEq(
            address(strategyFactory.strategyManager()),
            address(strategyManagerMock),
            "constructor / initializer incorrect, strategyManager set wrong"
        );
        assertEq(
            address(strategyFactory.strategyBeacon().implementation()),
            address(strategyImplementation),
            "constructor / initializer incorrect, strategyImplementation set wrong"
        );
        assertEq(
            address(strategyFactory.strategyBeacon()),
            address(strategyBeacon),
            "constructor / initializer incorrect, strategyBeacon set wrong"
        );
        assertEq(
            address(strategyFactory.pauserRegistry()),
            address(pauserRegistry),
            "constructor / initializer incorrect, pauserRegistry set wrong"
        );
        assertEq(strategyFactory.owner(), initialOwner, "constructor / initializer incorrect, owner set wrong");
        assertEq(strategyFactory.paused(), initialPausedStatus, "constructor / initializer incorrect, paused status set wrong");
        assertEq(strategyBeacon.owner(), beaconProxyOwner, "constructor / initializer incorrect, beaconProxyOwner set wrong");
    }

    function test_initialize_revert_reinitialization() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        strategyFactory.initialize({_initialOwner: initialOwner, _initialPausedStatus: initialPausedStatus, _strategyBeacon: strategyBeacon});
    }

    function test_deployNewStrategy() public {
        // cheats.expectEmit(true, true, true, true, address(strategyFactory));
        // StrategySetForToken(underlyingToken, newStrategy);
        StrategyBase newStrategy = StrategyBase(address(strategyFactory.deployNewStrategy(underlyingToken)));

        require(strategyFactory.deployedStrategies(underlyingToken) == newStrategy, "deployedStrategies mapping not set correctly");
        require(address(newStrategy.strategyManager()) == address(strategyManagerMock), "strategyManager not set correctly");
        require(strategyBeacon.implementation() == address(strategyImplementation), "strategyImplementation not set correctly");
        require(newStrategy.pauserRegistry() == pauserRegistry, "pauserRegistry not set correctly");
        require(newStrategy.underlyingToken() == underlyingToken, "underlyingToken not set correctly");
        require(strategyManagerMock.strategyIsWhitelistedForDeposit(newStrategy), "underlyingToken is not whitelisted");
    }

    function test_deployNewStrategy_revert_StrategyAlreadyExists() public {
        test_deployNewStrategy();
        cheats.expectRevert(IStrategyFactory.StrategyAlreadyExists.selector);
        strategyFactory.deployNewStrategy(underlyingToken);
    }

    function test_blacklistTokens(IERC20 token) public {
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = token;

        vm.prank(strategyFactory.owner());
        cheats.expectEmit(true, false, false, false, address(strategyFactory));
        emit TokenBlacklisted(token);
        strategyFactory.blacklistTokens(tokens);

        cheats.expectRevert(IStrategyFactory.BlacklistedToken.selector);
        strategyFactory.deployNewStrategy(token);
    }

    function test_blacklistTokens_RemovesFromWhitelist() public {
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = underlyingToken;

        IStrategy newStrat = strategyFactory.deployNewStrategy(underlyingToken);
        IStrategy[] memory toRemove = new IStrategy[](1);
        toRemove[0] = newStrat;

        vm.prank(strategyFactory.owner());
        cheats.expectEmit(true, false, false, false, address(strategyFactory));
        emit TokenBlacklisted(underlyingToken);
        cheats.expectCall(
            address(strategyManagerMock),
            abi.encodeWithSelector(strategyManagerMock.removeStrategiesFromDepositWhitelist.selector, toRemove)
        );
        strategyFactory.blacklistTokens(tokens);
    }

    function test_whitelistStrategies() public {
        StrategyBase strategy = _deployStrategy();
        IStrategy[] memory strategiesToWhitelist = new IStrategy[](1);
        strategiesToWhitelist[0] = strategy;
        strategyFactory.whitelistStrategies(strategiesToWhitelist);

        assertTrue(strategyManagerMock.strategyIsWhitelistedForDeposit(strategy), "Strategy not whitelisted");
    }

    function test_whitelistStrategies_revert_notOwner() public {
        IStrategy[] memory strategiesToWhitelist = new IStrategy[](1);

        cheats.expectRevert("Ownable: caller is not the owner");
        cheats.prank(notOwner);
        strategyFactory.whitelistStrategies(strategiesToWhitelist);
    }

    function test_removeStrategiesFromWhitelist_revert_notOwner() public {
        IStrategy[] memory strategiesToRemove = new IStrategy[](1);

        cheats.expectRevert("Ownable: caller is not the owner");
        cheats.prank(notOwner);
        strategyFactory.removeStrategiesFromWhitelist(strategiesToRemove);
    }

    function test_removeStrategiesFromWhitelist() public {
        IStrategy[] memory strategiesToRemove = new IStrategy[](1);
        strategiesToRemove[0] = IStrategy(_deployStrategy());

        strategyFactory.removeStrategiesFromWhitelist(strategiesToRemove);
        assertFalse(strategyManagerMock.strategyIsWhitelistedForDeposit(strategiesToRemove[0]), "Strategy not removed from whitelist");
    }

    function _deployStrategy() internal returns (StrategyBase) {
        return StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken)
                )
            )
        );
    }
}
