// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/multichain/CrossChainRegistry.sol";
import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/test/mocks/OperatorTableCalculatorMock.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";

/**
 * @title CrossChainRegistryUnitTests
 * @notice Base contract for all CrossChainRegistry unit tests
 */
contract CrossChainRegistryUnitTests is
    EigenLayerMultichainUnitTestSetup,
    ICrossChainRegistryErrors,
    ICrossChainRegistryTypes,
    ICrossChainRegistryEvents,
    IPermissionControllerErrors,
    IKeyRegistrarTypes
{
    // Constants from CrossChainRegistryStorage
    uint8 constant PAUSED_GENERATION_RESERVATIONS = 0;
    uint8 constant PAUSED_OPERATOR_TABLE_CALCULATOR = 1;
    uint8 constant PAUSED_OPERATOR_SET_CONFIG = 2;
    uint8 constant PAUSED_TRANSPORT_DESTINATIONS = 3;
    uint8 constant PAUSED_CHAIN_WHITELIST = 4;

    // Test state variables
    address defaultAVS;
    address notPermissioned = address(0xDEAD);
    OperatorSet defaultOperatorSet;
    OperatorTableCalculatorMock defaultCalculator;
    OperatorSetConfig defaultConfig;
    address defaultOperatorTableUpdater = address(0x1);
    uint[] defaultChainIDs;
    address[] defaultOperatorTableUpdaters;
    uint[] emptyChainIDs;

    function setUp() public virtual override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Set up default test values
        defaultAVS = cheats.randomAddress();
        defaultOperatorSet = OperatorSet({avs: defaultAVS, id: 1});

        defaultCalculator = new OperatorTableCalculatorMock();
        defaultConfig = OperatorSetConfig({owner: cheats.randomAddress(), maxStalenessPeriod: 1 days});

        defaultChainIDs = new uint[](2);
        defaultChainIDs[0] = 1;
        defaultChainIDs[1] = 10;
        defaultOperatorTableUpdaters = new address[](2);
        defaultOperatorTableUpdaters[0] = defaultOperatorTableUpdater;
        defaultOperatorTableUpdaters[1] = defaultOperatorTableUpdater;

        // Setup default permissions
        _grantUAMRole(address(this), defaultAVS);

        // Make the operator set valid in AllocationManager
        allocationManagerMock.setIsOperatorSet(defaultOperatorSet, true);

        // Whitelist chain IDs
        crossChainRegistry.addChainIDsToWhitelist(defaultChainIDs, defaultOperatorTableUpdaters);
    }

    // Helper functions
    function _grantUAMRole(address target, address avs) internal {
        // Grant admin role first
        cheats.prank(avs);
        permissionController.addPendingAdmin(avs, avs);
        cheats.prank(avs);
        permissionController.acceptAdmin(avs);

        // Set appointee for all CrossChainRegistry functions
        cheats.startPrank(avs);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.createGenerationReservation.selector);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.removeGenerationReservation.selector);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.setOperatorTableCalculator.selector);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.setOperatorSetConfig.selector);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.addTransportDestinations.selector);
        permissionController.setAppointee(avs, target, address(crossChainRegistry), crossChainRegistry.removeTransportDestinations.selector);

        // Set appointee for KeyRegistrar functions
        permissionController.setAppointee(avs, target, address(keyRegistrar), keyRegistrar.configureOperatorSet.selector);
        cheats.stopPrank();
    }

    function _createOperatorSet(address avs, uint32 operatorSetId) internal pure returns (OperatorSet memory) {
        return OperatorSet({avs: avs, id: operatorSetId});
    }

    function _createOperatorSetConfig(address owner, uint32 stalenessPeriod) internal pure returns (OperatorSetConfig memory) {
        return OperatorSetConfig({owner: owner, maxStalenessPeriod: stalenessPeriod});
    }

    function _createAndWhitelistChainIDs(uint count) internal returns (uint[] memory) {
        uint[] memory chainIDs = new uint[](count);
        address[] memory operatorTableUpdaters = new address[](count);
        for (uint i = 0; i < count; i++) {
            chainIDs[i] = 100 + i;
            operatorTableUpdaters[i] = address(uint160(uint(100 + i)));
        }
        crossChainRegistry.addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);
        return chainIDs;
    }
}

/**
 * @title CrossChainRegistryUnitTests_initialize
 * @notice Unit tests for CrossChainRegistry.initialize
 */
contract CrossChainRegistryUnitTests_initialize is CrossChainRegistryUnitTests {
    function test_initialize_AlreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        crossChainRegistry.initialize(address(this), 0);
    }

    function test_initialize_CorrectOwnerAndPausedStatus() public {
        // Deploy new implementation and proxy to test initialization
        CrossChainRegistry freshImplementation = new CrossChainRegistry(
            IAllocationManager(address(allocationManagerMock)),
            IKeyRegistrar(address(keyRegistrar)),
            IPermissionController(address(permissionController)),
            pauserRegistry,
            "1.0.0"
        );

        address newOwner = cheats.randomAddress();
        uint initialPausedStatus = (1 << PAUSED_GENERATION_RESERVATIONS) | (1 << PAUSED_TRANSPORT_DESTINATIONS);

        CrossChainRegistry freshRegistry = CrossChainRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(freshImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(CrossChainRegistry.initialize.selector, newOwner, initialPausedStatus)
                )
            )
        );

        assertEq(freshRegistry.owner(), newOwner, "Owner not set correctly");
        assertTrue(freshRegistry.paused(PAUSED_GENERATION_RESERVATIONS), "PAUSED_GENERATION_RESERVATIONS not set");
        assertTrue(freshRegistry.paused(PAUSED_TRANSPORT_DESTINATIONS), "PAUSED_TRANSPORT_DESTINATIONS not set");
        assertFalse(freshRegistry.paused(PAUSED_OPERATOR_TABLE_CALCULATOR), "PAUSED_OPERATOR_TABLE_CALCULATOR should not be set");
    }
}

/**
 * @title CrossChainRegistryUnitTests_createGenerationReservation
 * @notice Unit tests for CrossChainRegistry.createGenerationReservation
 */
contract CrossChainRegistryUnitTests_createGenerationReservation is CrossChainRegistryUnitTests {
    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_GENERATION_RESERVATIONS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.createGenerationReservation(invalidOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
    }

    function test_Revert_EmptyChainIDs() public {
        cheats.expectRevert(EmptyChainIDsArray.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, emptyChainIDs);
    }

    function test_Revert_GenerationReservationAlreadyExists() public {
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        cheats.expectRevert(GenerationReservationAlreadyExists.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
    }

    function test_Revert_InvalidOperatorTableCalculator() public {
        cheats.expectRevert(InvalidOperatorTableCalculator.selector);
        crossChainRegistry.createGenerationReservation(
            defaultOperatorSet, IOperatorTableCalculator(address(0)), defaultConfig, defaultChainIDs
        );
    }

    function test_Revert_InvalidOperatorSetConfig_ZeroOwner() public {
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(address(0), 1 days);

        cheats.expectRevert(IPausable.InputAddressZero.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, invalidConfig, defaultChainIDs);
    }

    function test_Revert_InvalidOperatorSetConfig_ZeroStalenessPeriod() public {
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(cheats.randomAddress(), 0);

        cheats.expectRevert(StalenessPeriodZero.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, invalidConfig, defaultChainIDs);
    }

    function test_Revert_ChainIDNotWhitelisted() public {
        uint[] memory nonWhitelistedChains = new uint[](1);
        nonWhitelistedChains[0] = 999;

        cheats.expectRevert(ChainIDNotWhitelisted.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, nonWhitelistedChains);
    }

    function test_createGenerationReservation_Success() public {
        // Expect events
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit GenerationReservationCreated(defaultOperatorSet);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorTableCalculatorSet(defaultOperatorSet, defaultCalculator);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorSetConfigSet(defaultOperatorSet, defaultConfig);

        for (uint i = 0; i < defaultChainIDs.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit TransportDestinationAdded(defaultOperatorSet, defaultChainIDs[i]);
        }

        // Make the call
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        // Verify state
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(activeReservations.length, 1, "Should have 1 active reservation");
        assertEq(activeReservations[0].avs, defaultOperatorSet.avs, "AVS mismatch");
        assertEq(activeReservations[0].id, defaultOperatorSet.id, "OperatorSetId mismatch");

        assertEq(
            address(crossChainRegistry.getOperatorTableCalculator(defaultOperatorSet)), address(defaultCalculator), "Calculator not set"
        );

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.owner, defaultConfig.owner, "Config owner mismatch");
        assertEq(retrievedConfig.maxStalenessPeriod, defaultConfig.maxStalenessPeriod, "Config staleness period mismatch");

        uint[] memory destinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(destinations.length, defaultChainIDs.length, "Transport destinations length mismatch");
        for (uint i = 0; i < destinations.length; i++) {
            assertEq(destinations[i], defaultChainIDs[i], "Transport destination mismatch");
        }
    }

    function testFuzz_createGenerationReservation_MultipleChainIDs(uint8 numChainIDs) public {
        numChainIDs = uint8(bound(numChainIDs, 1, 10));
        uint[] memory chainIDs = _createAndWhitelistChainIDs(numChainIDs);

        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, chainIDs);

        uint[] memory retrievedChainIDs = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(retrievedChainIDs.length, chainIDs.length, "Chain IDs length mismatch");
        for (uint i = 0; i < chainIDs.length; i++) {
            assertEq(retrievedChainIDs[i], chainIDs[i], "Chain ID mismatch");
        }
    }
}

/**
 * @title CrossChainRegistryUnitTests_removeGenerationReservation
 * @notice Unit tests for CrossChainRegistry.removeGenerationReservation
 */
contract CrossChainRegistryUnitTests_removeGenerationReservation is CrossChainRegistryUnitTests {
    function setUp() public override {
        super.setUp();
        // Create a default reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_GENERATION_RESERVATIONS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.removeGenerationReservation(defaultOperatorSet);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.removeGenerationReservation(defaultOperatorSet);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.removeGenerationReservation(invalidOperatorSet);
    }

    function test_Revert_GenerationReservationDoesNotExist() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(nonExistentOperatorSet, true);

        cheats.expectRevert(GenerationReservationDoesNotExist.selector);
        crossChainRegistry.removeGenerationReservation(nonExistentOperatorSet);
    }

    function test_removeGenerationReservation_Success() public {
        // Expect events
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit GenerationReservationRemoved(defaultOperatorSet);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorTableCalculatorSet(defaultOperatorSet, IOperatorTableCalculator(address(0)));

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorSetConfigSet(defaultOperatorSet, OperatorSetConfig(address(0), 0));

        for (uint i = 0; i < defaultChainIDs.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit TransportDestinationRemoved(defaultOperatorSet, defaultChainIDs[i]);
        }

        // Remove the reservation
        crossChainRegistry.removeGenerationReservation(defaultOperatorSet);

        // Verify state
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(activeReservations.length, 0, "Should have no active reservations");

        assertEq(address(crossChainRegistry.getOperatorTableCalculator(defaultOperatorSet)), address(0), "Calculator should be removed");

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.owner, address(0), "Config owner should be zero");
        assertEq(retrievedConfig.maxStalenessPeriod, 0, "Config staleness period should be zero");

        uint[] memory destinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(destinations.length, 0, "Should have no transport destinations");
    }
}

/**
 * @title CrossChainRegistryUnitTests_setOperatorTableCalculator
 * @notice Unit tests for CrossChainRegistry.setOperatorTableCalculator
 */
contract CrossChainRegistryUnitTests_setOperatorTableCalculator is CrossChainRegistryUnitTests {
    OperatorTableCalculatorMock newCalculator;

    function setUp() public override {
        super.setUp();
        // Create a default reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
        newCalculator = new OperatorTableCalculatorMock();
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_OPERATOR_TABLE_CALCULATOR);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.setOperatorTableCalculator(defaultOperatorSet, newCalculator);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.setOperatorTableCalculator(defaultOperatorSet, newCalculator);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.setOperatorTableCalculator(invalidOperatorSet, newCalculator);
    }

    function test_Revert_GenerationReservationDoesNotExist() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(nonExistentOperatorSet, true);

        cheats.expectRevert(GenerationReservationDoesNotExist.selector);
        crossChainRegistry.setOperatorTableCalculator(nonExistentOperatorSet, newCalculator);
    }

    function test_Revert_InvalidOperatorTableCalculator() public {
        cheats.expectRevert(InvalidOperatorTableCalculator.selector);
        crossChainRegistry.setOperatorTableCalculator(defaultOperatorSet, IOperatorTableCalculator(address(0)));
    }

    function test_setOperatorTableCalculator_Success() public {
        // Expect event
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorTableCalculatorSet(defaultOperatorSet, newCalculator);

        // Set new calculator
        crossChainRegistry.setOperatorTableCalculator(defaultOperatorSet, newCalculator);

        // Verify state
        assertEq(
            address(crossChainRegistry.getOperatorTableCalculator(defaultOperatorSet)), address(newCalculator), "Calculator not updated"
        );
    }

    function testFuzz_setOperatorTableCalculator_MultipleUpdates(uint8 numUpdates) public {
        numUpdates = uint8(bound(numUpdates, 1, 10));

        for (uint i = 0; i < numUpdates; i++) {
            OperatorTableCalculatorMock calc = new OperatorTableCalculatorMock();
            crossChainRegistry.setOperatorTableCalculator(defaultOperatorSet, calc);
            assertEq(address(crossChainRegistry.getOperatorTableCalculator(defaultOperatorSet)), address(calc), "Calculator not updated");
        }
    }
}

/**
 * @title CrossChainRegistryUnitTests_setOperatorSetConfig
 * @notice Unit tests for CrossChainRegistry.setOperatorSetConfig
 */
contract CrossChainRegistryUnitTests_setOperatorSetConfig is CrossChainRegistryUnitTests {
    OperatorSetConfig newConfig;

    function setUp() public override {
        super.setUp();
        // Create a default reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
        newConfig = _createOperatorSetConfig(cheats.randomAddress(), 2 days);
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_OPERATOR_SET_CONFIG);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, newConfig);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, newConfig);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.setOperatorSetConfig(invalidOperatorSet, newConfig);
    }

    function test_Revert_GenerationReservationDoesNotExist() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(nonExistentOperatorSet, true);

        cheats.expectRevert(GenerationReservationDoesNotExist.selector);
        crossChainRegistry.setOperatorSetConfig(nonExistentOperatorSet, newConfig);
    }

    function test_Revert_InvalidOperatorSetConfig_ZeroOwner() public {
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(address(0), 1 days);

        cheats.expectRevert(IPausable.InputAddressZero.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, invalidConfig);
    }

    function test_Revert_InvalidOperatorSetConfig_ZeroStalenessPeriod() public {
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(cheats.randomAddress(), 0);

        cheats.expectRevert(StalenessPeriodZero.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, invalidConfig);
    }

    function test_setOperatorSetConfig_Success() public {
        // Expect event
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorSetConfigSet(defaultOperatorSet, newConfig);

        // Set new config
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, newConfig);

        // Verify state
        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.owner, newConfig.owner, "Config owner not updated");
        assertEq(retrievedConfig.maxStalenessPeriod, newConfig.maxStalenessPeriod, "Config staleness period not updated");
    }

    function testFuzz_setOperatorSetConfig_StalenessPeriod(uint32 stalenessPeriod) public {
        stalenessPeriod = uint32(bound(stalenessPeriod, 1, 365 days));
        OperatorSetConfig memory fuzzConfig = _createOperatorSetConfig(cheats.randomAddress(), stalenessPeriod);

        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, fuzzConfig);

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.maxStalenessPeriod, stalenessPeriod, "Staleness period not set correctly");
    }
}

/**
 * @title CrossChainRegistryUnitTests_addTransportDestinations
 * @notice Unit tests for CrossChainRegistry.addTransportDestinations
 */
contract CrossChainRegistryUnitTests_addTransportDestinations is CrossChainRegistryUnitTests {
    uint[] newChainIDs;
    address[] newOperatorTableUpdaters;

    function setUp() public override {
        super.setUp();
        // Create a default reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        // Setup new chain IDs to add
        newChainIDs = new uint[](2);
        newChainIDs[0] = 20;
        newChainIDs[1] = 30;
        newOperatorTableUpdaters = new address[](2);
        newOperatorTableUpdaters[0] = defaultOperatorTableUpdater;
        newOperatorTableUpdaters[1] = defaultOperatorTableUpdater;
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, newOperatorTableUpdaters);
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_TRANSPORT_DESTINATIONS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, newChainIDs);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, newChainIDs);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.addTransportDestinations(invalidOperatorSet, newChainIDs);
    }

    function test_Revert_EmptyChainIDs() public {
        cheats.expectRevert(EmptyChainIDsArray.selector);
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, emptyChainIDs);
    }

    function test_Revert_GenerationReservationDoesNotExist() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(nonExistentOperatorSet, true);

        cheats.expectRevert(GenerationReservationDoesNotExist.selector);
        crossChainRegistry.addTransportDestinations(nonExistentOperatorSet, newChainIDs);
    }

    function test_Revert_ChainIDNotWhitelisted() public {
        uint[] memory nonWhitelistedChains = new uint[](1);
        nonWhitelistedChains[0] = 999;

        cheats.expectRevert(ChainIDNotWhitelisted.selector);
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, nonWhitelistedChains);
    }

    function test_Revert_TransportDestinationAlreadyAdded() public {
        cheats.expectRevert(TransportDestinationAlreadyAdded.selector);
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, defaultChainIDs);
    }

    function test_addTransportDestinations_Success() public {
        // Expect events
        for (uint i = 0; i < newChainIDs.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit TransportDestinationAdded(defaultOperatorSet, newChainIDs[i]);
        }

        // Add new destinations
        crossChainRegistry.addTransportDestinations(defaultOperatorSet, newChainIDs);

        // Verify state
        uint[] memory destinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(destinations.length, defaultChainIDs.length + newChainIDs.length, "Destinations count mismatch");

        // Check all destinations are present
        bool[] memory found = new bool[](defaultChainIDs.length + newChainIDs.length);
        for (uint i = 0; i < destinations.length; i++) {
            for (uint j = 0; j < defaultChainIDs.length; j++) {
                if (destinations[i] == defaultChainIDs[j]) found[j] = true;
            }
            for (uint j = 0; j < newChainIDs.length; j++) {
                if (destinations[i] == newChainIDs[j]) found[defaultChainIDs.length + j] = true;
            }
        }
        for (uint i = 0; i < found.length; i++) {
            assertTrue(found[i], "Chain ID not found");
        }
    }

    function testFuzz_addTransportDestinations_MultipleChainIDs(uint8 numNewChainIDs) public {
        numNewChainIDs = uint8(bound(numNewChainIDs, 1, 10));
        uint[] memory fuzzChainIDs = _createAndWhitelistChainIDs(numNewChainIDs);

        crossChainRegistry.addTransportDestinations(defaultOperatorSet, fuzzChainIDs);

        uint[] memory destinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(destinations.length, defaultChainIDs.length + fuzzChainIDs.length, "Destinations count mismatch");
    }
}

/**
 * @title CrossChainRegistryUnitTests_removeTransportDestinations
 * @notice Unit tests for CrossChainRegistry.removeTransportDestinations
 */
contract CrossChainRegistryUnitTests_removeTransportDestinations is CrossChainRegistryUnitTests {
    uint[] chainIDsToRemove;

    function setUp() public override {
        super.setUp();
        // Create a default reservation with multiple chain IDs
        uint[] memory manyChainIDs = new uint[](4);
        manyChainIDs[0] = 1; // Already whitelisted in base setUp
        manyChainIDs[1] = 10; // Already whitelisted in base setUp
        manyChainIDs[2] = 20;
        manyChainIDs[3] = 30;

        // Only whitelist the new chain IDs
        uint[] memory newChainIDs = new uint[](2);
        newChainIDs[0] = 20;
        newChainIDs[1] = 30;
        address[] memory newOperatorTableUpdaters = new address[](2);
        newOperatorTableUpdaters[0] = defaultOperatorTableUpdater;
        newOperatorTableUpdaters[1] = defaultOperatorTableUpdater;
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, newOperatorTableUpdaters);

        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, manyChainIDs);

        // Setup chain IDs to remove (subset)
        chainIDsToRemove = new uint[](2);
        chainIDsToRemove[0] = 10;
        chainIDsToRemove[1] = 20;
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_TRANSPORT_DESTINATIONS);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, chainIDsToRemove);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, chainIDsToRemove);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.removeTransportDestinations(invalidOperatorSet, chainIDsToRemove);
    }

    function test_Revert_EmptyChainIDs() public {
        cheats.expectRevert(EmptyChainIDsArray.selector);
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, emptyChainIDs);
    }

    function test_Revert_GenerationReservationDoesNotExist() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(nonExistentOperatorSet, true);

        cheats.expectRevert(GenerationReservationDoesNotExist.selector);
        crossChainRegistry.removeTransportDestinations(nonExistentOperatorSet, chainIDsToRemove);
    }

    function test_Revert_TransportDestinationNotFound() public {
        uint[] memory nonExistentChains = new uint[](1);
        nonExistentChains[0] = 999;

        cheats.expectRevert(TransportDestinationNotFound.selector);
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, nonExistentChains);
    }

    function test_Revert_RequireAtLeastOneTransportDestination() public {
        // Get all current destinations
        uint[] memory allDestinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);

        // Try to remove all of them
        cheats.expectRevert(RequireAtLeastOneTransportDestination.selector);
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, allDestinations);
    }

    function test_removeTransportDestinations_Success() public {
        // Expect events
        for (uint i = 0; i < chainIDsToRemove.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit TransportDestinationRemoved(defaultOperatorSet, chainIDsToRemove[i]);
        }

        // Remove destinations
        crossChainRegistry.removeTransportDestinations(defaultOperatorSet, chainIDsToRemove);

        // Verify state
        uint[] memory remainingDestinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(remainingDestinations.length, 2, "Should have 2 remaining destinations");

        // Verify the correct destinations remain (1 and 30)
        assertTrue(
            (remainingDestinations[0] == 1 && remainingDestinations[1] == 30)
                || (remainingDestinations[0] == 30 && remainingDestinations[1] == 1),
            "Incorrect remaining destinations"
        );
    }

    function testFuzz_removeTransportDestinations_PartialRemoval(uint8 numToRemove) public {
        // Setup with many destinations
        uint[] memory manyChainIDs = _createAndWhitelistChainIDs(10);
        OperatorSet memory fuzzOperatorSet = _createOperatorSet(cheats.randomAddress(), 100);
        allocationManagerMock.setIsOperatorSet(fuzzOperatorSet, true);
        _grantUAMRole(address(this), fuzzOperatorSet.avs);

        crossChainRegistry.createGenerationReservation(fuzzOperatorSet, defaultCalculator, defaultConfig, manyChainIDs);

        // Remove some but not all
        numToRemove = uint8(bound(numToRemove, 1, 9)); // Leave at least one
        uint[] memory toRemove = new uint[](numToRemove);
        for (uint i = 0; i < numToRemove; i++) {
            toRemove[i] = manyChainIDs[i];
        }

        crossChainRegistry.removeTransportDestinations(fuzzOperatorSet, toRemove);

        uint[] memory remaining = crossChainRegistry.getTransportDestinations(fuzzOperatorSet);
        assertEq(remaining.length, manyChainIDs.length - numToRemove, "Incorrect remaining count");
    }
}

/**
 * @title CrossChainRegistryUnitTests_addChainIDsToWhitelist
 * @notice Unit tests for CrossChainRegistry.addChainIDsToWhitelist
 */
contract CrossChainRegistryUnitTests_addChainIDsToWhitelist is CrossChainRegistryUnitTests {
    uint[] newChainIDs;
    address[] newOperatorTableUpdaters;

    function setUp() public override {
        super.setUp();
        newChainIDs = new uint[](3);
        newChainIDs[0] = 100;
        newChainIDs[1] = 200;
        newChainIDs[2] = 300;
        newOperatorTableUpdaters = new address[](3);
        newOperatorTableUpdaters[0] = defaultOperatorTableUpdater;
        newOperatorTableUpdaters[1] = defaultOperatorTableUpdater;
        newOperatorTableUpdaters[2] = defaultOperatorTableUpdater;
    }

    function test_Revert_NotOwner() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert("Ownable: caller is not the owner");
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, newOperatorTableUpdaters);
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_CHAIN_WHITELIST);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, newOperatorTableUpdaters);
    }

    function test_Revert_ArrayLengthMismatch() public {
        address[] memory invalidOperatorTableUpdaters = new address[](1);
        invalidOperatorTableUpdaters[0] = defaultOperatorTableUpdater;
        cheats.expectRevert(ArrayLengthMismatch.selector);
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, invalidOperatorTableUpdaters);
    }

    function test_Revert_InvalidChainId() public {
        uint[] memory invalidChainIDs = new uint[](1);
        invalidChainIDs[0] = 0;
        address[] memory invalidOperatorTableUpdaters = new address[](1);
        invalidOperatorTableUpdaters[0] = defaultOperatorTableUpdater;

        cheats.expectRevert(InvalidChainId.selector);
        crossChainRegistry.addChainIDsToWhitelist(invalidChainIDs, invalidOperatorTableUpdaters);
    }

    function test_Revert_ChainIDAlreadyWhitelisted() public {
        cheats.expectRevert(ChainIDAlreadyWhitelisted.selector);
        crossChainRegistry.addChainIDsToWhitelist(defaultChainIDs, defaultOperatorTableUpdaters);
    }

    function test_addChainIDsToWhitelist_Success() public {
        // Expect events
        for (uint i = 0; i < newChainIDs.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit ChainIDAddedToWhitelist(newChainIDs[i], newOperatorTableUpdaters[i]);
        }

        // Add to whitelist
        crossChainRegistry.addChainIDsToWhitelist(newChainIDs, newOperatorTableUpdaters);

        // Verify state
        (uint[] memory supportedChains, address[] memory operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
        assertEq(supportedChains.length, defaultChainIDs.length + newChainIDs.length, "Supported chains count mismatch");
        assertEq(operatorTableUpdaters.length, defaultChainIDs.length + newChainIDs.length, "Operator table updaters count mismatch");
        for (uint i = 0; i < supportedChains.length; i++) {
            if (i < defaultChainIDs.length) {
                assertEq(supportedChains[i], defaultChainIDs[i], "Supported chain mismatch");
                assertEq(operatorTableUpdaters[i], defaultOperatorTableUpdaters[i], "Operator table updater mismatch");
            } else {
                assertEq(supportedChains[i], newChainIDs[i - defaultChainIDs.length], "Supported chain mismatch");
                assertEq(operatorTableUpdaters[i], newOperatorTableUpdaters[i - defaultChainIDs.length], "Operator table updater mismatch");
            }
        }
    }

    function testFuzz_addChainIDsToWhitelist_MultipleChainIDs(uint8 numChainIDs) public {
        numChainIDs = uint8(bound(numChainIDs, 1, 50));
        uint[] memory fuzzChainIDs = new uint[](numChainIDs);
        address[] memory fuzzOperatorTableUpdaters = new address[](numChainIDs);

        for (uint i = 0; i < numChainIDs; i++) {
            fuzzChainIDs[i] = 1000 + uint(i);
            fuzzOperatorTableUpdaters[i] = address(uint160(uint(1000 + i)));
        }

        crossChainRegistry.addChainIDsToWhitelist(fuzzChainIDs, fuzzOperatorTableUpdaters);

        (uint[] memory supportedChains, address[] memory operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
        assertEq(supportedChains.length, numChainIDs + defaultChainIDs.length, "Supported chains count mismatch");
        assertEq(operatorTableUpdaters.length, numChainIDs + defaultChainIDs.length, "Operator table updaters count mismatch");
        for (uint i = 0; i < supportedChains.length; i++) {
            if (i < defaultChainIDs.length) {
                assertEq(supportedChains[i], defaultChainIDs[i], "Supported chain mismatch");
                assertEq(operatorTableUpdaters[i], defaultOperatorTableUpdaters[i], "Operator table updater mismatch");
            } else {
                assertEq(supportedChains[i], fuzzChainIDs[i - defaultChainIDs.length], "Supported chain mismatch");
                assertEq(operatorTableUpdaters[i], fuzzOperatorTableUpdaters[i - defaultChainIDs.length], "Operator table updater mismatch");
            }
        }
    }
}

/**
 * @title CrossChainRegistryUnitTests_removeChainIDsFromWhitelist
 * @notice Unit tests for CrossChainRegistry.removeChainIDsFromWhitelist
 */
contract CrossChainRegistryUnitTests_removeChainIDsFromWhitelist is CrossChainRegistryUnitTests {
    function setUp() public override {
        super.setUp();
    }

    function test_Revert_NotOwner() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert("Ownable: caller is not the owner");
        crossChainRegistry.removeChainIDsFromWhitelist(defaultChainIDs);
    }

    function test_Revert_Paused() public {
        cheats.prank(pauser);
        crossChainRegistry.pause(1 << PAUSED_CHAIN_WHITELIST);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        crossChainRegistry.removeChainIDsFromWhitelist(defaultChainIDs);
    }

    function test_Revert_ChainIDNotWhitelisted() public {
        uint[] memory nonWhitelistedChains = new uint[](1);
        nonWhitelistedChains[0] = 999;

        cheats.expectRevert(ChainIDNotWhitelisted.selector);
        crossChainRegistry.removeChainIDsFromWhitelist(nonWhitelistedChains);
    }

    function test_removeChainIDsFromWhitelist_Success() public {
        // Expect events
        for (uint i = 0; i < defaultChainIDs.length; i++) {
            cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
            emit ChainIDRemovedFromWhitelist(defaultChainIDs[i]);
        }

        // Remove from whitelist
        crossChainRegistry.removeChainIDsFromWhitelist(defaultChainIDs);

        // Verify state
        (uint[] memory supportedChains, address[] memory operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
        assertEq(supportedChains.length, 0, "Should have no supported chains");
        assertEq(operatorTableUpdaters.length, 0, "Should have no operator table updaters");
    }

    function test_removeChainIDsFromWhitelist_AffectsTransportDestinations() public {
        // Create a reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        // Remove one chain from whitelist
        uint[] memory chainToRemove = new uint[](1);
        chainToRemove[0] = defaultChainIDs[0];
        crossChainRegistry.removeChainIDsFromWhitelist(chainToRemove);

        // Verify transport destinations only returns whitelisted chains
        uint[] memory destinations = crossChainRegistry.getTransportDestinations(defaultOperatorSet);
        assertEq(destinations.length, 1, "Should only return whitelisted destination");
        assertEq(destinations[0], defaultChainIDs[1], "Wrong destination returned");
    }
}

/**
 * @title CrossChainRegistryUnitTests_getActiveGenerationReservations
 * @notice Unit tests for CrossChainRegistry.getActiveGenerationReservations
 */
contract CrossChainRegistryUnitTests_getActiveGenerationReservations is CrossChainRegistryUnitTests {
    function test_getActiveGenerationReservations_Empty() public {
        OperatorSet[] memory reservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(reservations.length, 0, "Should have no reservations");
    }

    function test_getActiveGenerationReservations_Single() public {
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        OperatorSet[] memory reservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(reservations.length, 1, "Should have 1 reservation");
        assertEq(reservations[0].avs, defaultOperatorSet.avs, "AVS mismatch");
        assertEq(reservations[0].id, defaultOperatorSet.id, "OperatorSetId mismatch");
    }

    function testFuzz_getActiveGenerationReservations_Multiple(uint8 numReservations) public {
        numReservations = uint8(bound(numReservations, 1, 10));

        for (uint i = 0; i < numReservations; i++) {
            OperatorSet memory operatorSet = _createOperatorSet(cheats.randomAddress(), uint32(i));
            allocationManagerMock.setIsOperatorSet(operatorSet, true);
            _grantUAMRole(address(this), operatorSet.avs);

            crossChainRegistry.createGenerationReservation(operatorSet, defaultCalculator, defaultConfig, defaultChainIDs);
        }

        OperatorSet[] memory reservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(reservations.length, numReservations, "Reservation count mismatch");
    }
}

/**
 * @title CrossChainRegistryUnitTests_calculateOperatorTableBytes
 * @notice Unit tests for CrossChainRegistry.calculateOperatorTableBytes
 */
contract CrossChainRegistryUnitTests_calculateOperatorTableBytes is CrossChainRegistryUnitTests {
    bytes testOperatorTableBytes = hex"1234567890";

    function setUp() public override {
        super.setUp();
        // Create a default reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        // Set up mock data
        defaultCalculator.setOperatorTableBytes(defaultOperatorSet, testOperatorTableBytes);
        // Configure operator set in KeyRegistrar (permissions already granted in base setUp)
        keyRegistrar.configureOperatorSet(defaultOperatorSet, CurveType.BN254);
    }

    function test_calculateOperatorTableBytes_Success() public {
        bytes memory result = crossChainRegistry.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode the result
        (
            OperatorSet memory decodedOperatorSet,
            CurveType curveType,
            OperatorSetConfig memory decodedConfig,
            bytes memory decodedOperatorTableBytes
        ) = abi.decode(result, (OperatorSet, CurveType, OperatorSetConfig, bytes));

        // Verify the decoded data
        assertEq(decodedOperatorSet.avs, defaultOperatorSet.avs, "AVS mismatch");
        assertEq(decodedOperatorSet.id, defaultOperatorSet.id, "OperatorSetId mismatch");
        assertTrue(curveType == CurveType.BN254, "CurveType mismatch");
        assertEq(decodedConfig.owner, defaultConfig.owner, "Config owner mismatch");
        assertEq(decodedConfig.maxStalenessPeriod, defaultConfig.maxStalenessPeriod, "Config staleness period mismatch");
        assertEq(decodedOperatorTableBytes, testOperatorTableBytes, "OperatorTableBytes mismatch");
    }

    function test_calculateOperatorTableBytes_NonExistentOperatorSet() public {
        OperatorSet memory nonExistentOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Should revert when trying to call calculateOperatorTableBytes on a null calculator
        cheats.expectRevert();
        crossChainRegistry.calculateOperatorTableBytes(nonExistentOperatorSet);
    }
}

/**
 * @title CrossChainRegistryUnitTests_getActiveTransportReservations
 * @notice Unit tests for CrossChainRegistry.getActiveTransportReservations
 */
contract CrossChainRegistryUnitTests_getActiveTransportReservations is CrossChainRegistryUnitTests {
    function test_getActiveTransportReservations_Empty() public {
        (OperatorSet[] memory operatorSets, uint[][] memory chainIDs) = crossChainRegistry.getActiveTransportReservations();
        assertEq(operatorSets.length, 0, "Should have no transport reservations");
        assertEq(chainIDs.length, 0, "Should have no chain IDs");
    }

    function test_getActiveTransportReservations_Single() public {
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig, defaultChainIDs);

        (OperatorSet[] memory operatorSets, uint[][] memory chainIDs) = crossChainRegistry.getActiveTransportReservations();
        assertEq(operatorSets.length, 1, "Should have 1 transport reservation");
        assertEq(operatorSets[0].avs, defaultOperatorSet.avs, "AVS mismatch");
        assertEq(operatorSets[0].id, defaultOperatorSet.id, "OperatorSetId mismatch");
        assertEq(chainIDs[0].length, defaultChainIDs.length, "Chain IDs length mismatch");
        for (uint i = 0; i < chainIDs[0].length; i++) {
            assertEq(chainIDs[0][i], defaultChainIDs[i], "Chain ID mismatch");
        }
    }

    function test_getActiveTransportReservations_Multiple() public {
        uint numReservations = 3;

        for (uint i = 0; i < numReservations; i++) {
            OperatorSet memory operatorSet = _createOperatorSet(cheats.randomAddress(), uint32(i));
            allocationManagerMock.setIsOperatorSet(operatorSet, true);
            _grantUAMRole(address(this), operatorSet.avs);

            // Create unique chain IDs for each iteration
            uint[] memory chainIDs = new uint[](i + 1);
            address[] memory operatorTableUpdaters = new address[](i + 1);
            for (uint j = 0; j <= i; j++) {
                // Use a formula that ensures unique chainIDs across iterations
                chainIDs[j] = 100 + uint(i * 10 + j);
                operatorTableUpdaters[j] = address(uint160(uint(100 + i * 10 + j)));
            }
            crossChainRegistry.addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);

            crossChainRegistry.createGenerationReservation(operatorSet, defaultCalculator, defaultConfig, chainIDs);
        }

        (OperatorSet[] memory operatorSets, uint[][] memory chainIDs) = crossChainRegistry.getActiveTransportReservations();
        assertEq(operatorSets.length, numReservations, "Transport reservation count mismatch");

        for (uint i = 0; i < numReservations; i++) {
            assertEq(chainIDs[i].length, i + 1, "Chain IDs length mismatch for reservation");
        }
    }
}

/**
 * @title CrossChainRegistryUnitTests_getSupportedChains
 * @notice Unit tests for CrossChainRegistry.getSupportedChains
 */
contract CrossChainRegistryUnitTests_getSupportedChains is CrossChainRegistryUnitTests {
    function test_getSupportedChains_Initial() public {
        (uint[] memory supportedChains, address[] memory operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
        assertEq(supportedChains.length, defaultChainIDs.length, "Should have default chains");
        assertEq(operatorTableUpdaters.length, defaultChainIDs.length, "Should have default operator table updaters");
        for (uint i = 0; i < supportedChains.length; i++) {
            bool found = false;
            for (uint j = 0; j < defaultChainIDs.length; j++) {
                if (supportedChains[i] == defaultChainIDs[j]) {
                    found = true;
                    break;
                }
            }
            assertTrue(found, "Chain ID not found in supported chains");
        }
    }

    function testFuzz_getSupportedChains_AddAndRemove(uint8 numToAdd, uint8 numToRemove) public {
        numToAdd = uint8(bound(numToAdd, 1, 20));
        numToRemove = uint8(bound(numToRemove, 0, defaultChainIDs.length));

        // Add chains
        uint[] memory newChains = _createAndWhitelistChainIDs(numToAdd);

        (uint[] memory supportedChains, address[] memory operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
        assertEq(supportedChains.length, defaultChainIDs.length + numToAdd, "Chain count after add mismatch");

        // Remove some default chains
        if (numToRemove > 0) {
            uint[] memory chainsToRemove = new uint[](numToRemove);
            for (uint i = 0; i < numToRemove; i++) {
                chainsToRemove[i] = defaultChainIDs[i];
            }
            crossChainRegistry.removeChainIDsFromWhitelist(chainsToRemove);

            (supportedChains, operatorTableUpdaters) = crossChainRegistry.getSupportedChains();
            assertEq(supportedChains.length, defaultChainIDs.length + numToAdd - numToRemove, "Chain count after remove mismatch");
            assertEq(
                operatorTableUpdaters.length,
                defaultChainIDs.length + numToAdd - numToRemove,
                "Operator table updater count after remove mismatch"
            );
        }
    }
}
