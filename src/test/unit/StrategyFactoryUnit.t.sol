// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/strategies/StrategyFactory.sol";
import "src/contracts/strategies/DurationVaultStrategy.sol";
import "../../contracts/interfaces/IDurationVaultStrategy.sol";
import "../../contracts/interfaces/IDelegationManager.sol";
import "../../contracts/interfaces/IAllocationManager.sol";
import "../../contracts/libraries/OperatorSetLib.sol";
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
    DurationVaultStrategy public durationVaultImplementation;
    UpgradeableBeacon public strategyBeacon;
    UpgradeableBeacon public durationVaultBeacon;
    ERC20PresetFixedSupply public underlyingToken;

    uint initialSupply = 1e36;
    address initialOwner = address(this);
    address beaconProxyOwner = address(this);
    address notOwner = address(7_777_777);

    uint initialPausedStatus = 0;
    address internal constant OPERATOR_SET_AVS = address(0xABCD);
    uint32 internal constant OPERATOR_SET_ID = 9;
    address internal constant DELEGATION_APPROVER = address(0x5151);
    uint32 internal constant OPERATOR_ALLOCATION_DELAY = 4;
    string internal constant OPERATOR_METADATA_URI = "ipfs://factory-duration-vault";
    bytes internal constant REGISTRATION_DATA = hex"F00D";

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

        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "9.9.9");

        strategyBeacon = new UpgradeableBeacon(address(strategyImplementation));
        strategyBeacon.transferOwnership(beaconProxyOwner);

        durationVaultImplementation = new DurationVaultStrategy(
            IStrategyManager(address(strategyManagerMock)),
            pauserRegistry,
            "9.9.9",
            IDelegationManager(address(delegationManagerMock)),
            IAllocationManager(address(allocationManagerMock))
        );
        durationVaultBeacon = new UpgradeableBeacon(address(durationVaultImplementation));
        durationVaultBeacon.transferOwnership(beaconProxyOwner);

        strategyFactoryImplementation = new StrategyFactory(IStrategyManager(address(strategyManagerMock)), pauserRegistry, "9.9.9");

        strategyFactory = StrategyFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyFactoryImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyFactory.initialize.selector, initialOwner, initialPausedStatus, strategyBeacon)
                )
            )
        );

        strategyFactory.setDurationVaultBeacon(durationVaultBeacon);
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

    function test_deployDurationVaultStrategy() public {
        IDurationVaultStrategy.VaultConfig memory config = IDurationVaultStrategy.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            duration: uint32(30 days),
            maxPerDeposit: 10 ether,
            stakeCap: 100 ether,
            metadataURI: "ipfs://duration",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
            operatorAllocationDelay: OPERATOR_ALLOCATION_DELAY,
            operatorMetadataURI: OPERATOR_METADATA_URI
        });

        DurationVaultStrategy vault = DurationVaultStrategy(address(strategyFactory.deployDurationVaultStrategy(config)));

        IDurationVaultStrategy[] memory deployedVaults = strategyFactory.getDurationVaults(underlyingToken);
        require(deployedVaults.length == 1, "vault not tracked");
        require(address(deployedVaults[0]) == address(vault), "vault mismatch");
        assertTrue(strategyManagerMock.strategyIsWhitelistedForDeposit(vault), "duration vault not whitelisted");
    }

    function test_deployDurationVaultStrategy_revertBeaconNotSet() public {
        strategyFactory.setDurationVaultBeacon(IBeacon(address(0)));

        IDurationVaultStrategy.VaultConfig memory config = IDurationVaultStrategy.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            duration: uint32(30 days),
            maxPerDeposit: 10 ether,
            stakeCap: 100 ether,
            metadataURI: "ipfs://duration",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
            operatorAllocationDelay: OPERATOR_ALLOCATION_DELAY,
            operatorMetadataURI: OPERATOR_METADATA_URI
        });

        cheats.expectRevert(IStrategyFactory.DurationVaultBeaconNotSet.selector);
        strategyFactory.deployDurationVaultStrategy(config);
    }

    function test_deployDurationVaultStrategy_withExistingStrategy() public {
        StrategyBase base = StrategyBase(address(strategyFactory.deployNewStrategy(underlyingToken)));
        require(strategyFactory.deployedStrategies(underlyingToken) == base, "base strategy missing");

        IDurationVaultStrategy.VaultConfig memory config = IDurationVaultStrategy.VaultConfig({
            underlyingToken: underlyingToken,
            vaultAdmin: address(this),
            duration: uint32(7 days),
            maxPerDeposit: 5 ether,
            stakeCap: 50 ether,
            metadataURI: "ipfs://duration",
            operatorSet: OperatorSet({avs: OPERATOR_SET_AVS, id: OPERATOR_SET_ID}),
            operatorSetRegistrationData: REGISTRATION_DATA,
            delegationApprover: DELEGATION_APPROVER,
            operatorAllocationDelay: OPERATOR_ALLOCATION_DELAY,
            operatorMetadataURI: OPERATOR_METADATA_URI
        });

        strategyFactory.deployDurationVaultStrategy(config);

        IDurationVaultStrategy[] memory deployedVaults = strategyFactory.getDurationVaults(underlyingToken);
        require(deployedVaults.length == 1, "duration vault missing");
        assertTrue(
            strategyManagerMock.strategyIsWhitelistedForDeposit(IDurationVaultStrategy(address(deployedVaults[0]))), "vault not whitelisted"
        );

        // Base mapping should remain untouched.
        require(strategyFactory.deployedStrategies(underlyingToken) == base, "base strategy overwritten");
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
