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
    uint8 constant PAUSED_CHAIN_WHITELIST = 3;

    // Test state variables
    CrossChainRegistry crossChainRegistry;
    CrossChainRegistry crossChainRegistryImplementation;
    address defaultAVS;
    address notPermissioned = address(0xDEAD);
    OperatorSet defaultOperatorSet;
    OperatorTableCalculatorMock defaultCalculator;
    OperatorSetConfig defaultConfig;
    address defaultOperatorTableUpdater = address(0x1);
    uint[] defaultChainIDs;
    address[] defaultOperatorTableUpdaters;

    function setUp() public virtual override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Deploy CrossChainRegistry implementation
        crossChainRegistryImplementation = new CrossChainRegistry(
            IAllocationManager(address(allocationManagerMock)),
            IKeyRegistrar(address(keyRegistrar)),
            IPermissionController(address(permissionController)),
            pauserRegistry,
            "1.0.0"
        );

        // Deploy CrossChainRegistry proxy
        crossChainRegistry = CrossChainRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(crossChainRegistryImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        CrossChainRegistry.initialize.selector,
                        address(this), // initial owner
                        1 days, // initial table update cadence
                        0 // initial paused status
                    )
                )
            )
        );

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

        // Set the key type for the operator set in KeyRegistrar
        keyRegistrar.configureOperatorSet(defaultOperatorSet, CurveType.BN254);

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
        crossChainRegistry.initialize(address(this), 0, 0);
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
        uint32 initialTableUpdateCadence = 1 days;
        uint initialPausedStatus = (1 << PAUSED_GENERATION_RESERVATIONS) | (1 << PAUSED_OPERATOR_TABLE_CALCULATOR);

        CrossChainRegistry freshRegistry = CrossChainRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(freshImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(CrossChainRegistry.initialize.selector, newOwner, initialTableUpdateCadence, initialPausedStatus)
                )
            )
        );

        assertEq(freshRegistry.owner(), newOwner, "Owner not set correctly");
        assertEq(freshRegistry.getTableUpdateCadence(), initialTableUpdateCadence, "Table update cadence not set correctly");
        assertTrue(freshRegistry.paused(PAUSED_GENERATION_RESERVATIONS), "PAUSED_GENERATION_RESERVATIONS not set");
        assertTrue(freshRegistry.paused(PAUSED_OPERATOR_TABLE_CALCULATOR), "PAUSED_OPERATOR_TABLE_CALCULATOR should not be set");
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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
    }

    function test_Revert_NotPermissioned() public {
        cheats.prank(notPermissioned);
        cheats.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
    }

    function test_Revert_InvalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = _createOperatorSet(cheats.randomAddress(), 999);

        // Grant permission for the invalid operator set's AVS
        _grantUAMRole(address(this), invalidOperatorSet.avs);

        cheats.expectRevert(InvalidOperatorSet.selector);
        crossChainRegistry.createGenerationReservation(invalidOperatorSet, defaultCalculator, defaultConfig);
    }

    function test_Revert_GenerationReservationAlreadyExists() public {
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        cheats.expectRevert(GenerationReservationAlreadyExists.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
    }

    function test_Revert_InvalidStalenessPeriod() public {
        // Set a table update cadence
        uint32 tableUpdateCadence = 7 days;
        crossChainRegistry.setTableUpdateCadence(tableUpdateCadence);

        // Try to create with a config that has staleness period less than cadence (but not 0)
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(cheats.randomAddress(), 1 days);

        cheats.expectRevert(InvalidStalenessPeriod.selector);
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, invalidConfig);
    }

    function test_Revert_KeyTypeNotSet() public {
        // Create a new operator set that hasn't been configured in KeyRegistrar
        OperatorSet memory unconfiguredOperatorSet = _createOperatorSet(defaultAVS, 999);

        // Make the operator set valid in AllocationManager but don't configure it in KeyRegistrar
        allocationManagerMock.setIsOperatorSet(unconfiguredOperatorSet, true);

        // Expect the KeyTypeNotSet error when trying to create generation reservation
        cheats.expectRevert(KeyTypeNotSet.selector);
        crossChainRegistry.createGenerationReservation(unconfiguredOperatorSet, defaultCalculator, defaultConfig);
    }

    function test_createGenerationReservation_Success() public {
        // Expect events
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit GenerationReservationCreated(defaultOperatorSet);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorTableCalculatorSet(defaultOperatorSet, defaultCalculator);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorSetConfigSet(defaultOperatorSet, defaultConfig);

        // Make the call
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
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
        emit OperatorTableCalculatorRemoved(defaultOperatorSet);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit OperatorSetConfigRemoved(defaultOperatorSet);

        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit GenerationReservationRemoved(defaultOperatorSet);

        // Remove the reservation
        crossChainRegistry.removeGenerationReservation(defaultOperatorSet);

        // Verify state
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        assertEq(activeReservations.length, 0, "Should have no active reservations");

        assertEq(address(crossChainRegistry.getOperatorTableCalculator(defaultOperatorSet)), address(0), "Calculator should be removed");

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.owner, address(0), "Config owner should be zero");
        assertEq(retrievedConfig.maxStalenessPeriod, 0, "Config staleness period should be zero");
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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);
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
        stalenessPeriod = uint32(bound(stalenessPeriod, 1 days, 365 days));
        OperatorSetConfig memory fuzzConfig = _createOperatorSetConfig(cheats.randomAddress(), stalenessPeriod);

        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, fuzzConfig);

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.maxStalenessPeriod, stalenessPeriod, "Staleness period not set correctly");
    }

    function test_Revert_InvalidStalenessPeriod() public {
        // Set a table update cadence
        uint32 tableUpdateCadence = 7 days;
        crossChainRegistry.setTableUpdateCadence(tableUpdateCadence);

        // Try to set config with staleness period less than minimum (but not 0)
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(cheats.randomAddress(), 1 days);

        cheats.expectRevert(InvalidStalenessPeriod.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, invalidConfig);
    }

    function test_setOperatorSetConfig_ZeroStalenessPeriod() public {
        // Set a table update cadence
        uint32 tableUpdateCadence = 7 days;
        crossChainRegistry.setTableUpdateCadence(tableUpdateCadence);

        // Create config with 0 staleness period (should be allowed as special case)
        OperatorSetConfig memory zeroStalenessConfig = _createOperatorSetConfig(cheats.randomAddress(), 0);

        // Should succeed
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, zeroStalenessConfig);

        // Verify the config was set
        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.maxStalenessPeriod, 0, "Zero staleness period should be allowed");
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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

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
            // Set the key type for the operator set in KeyRegistrar
            keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

            crossChainRegistry.createGenerationReservation(operatorSet, defaultCalculator, defaultConfig);
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
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        // Set up mock data
        defaultCalculator.setOperatorTableBytes(defaultOperatorSet, testOperatorTableBytes);
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

/**
 * @title CrossChainRegistryUnitTests_setTableUpdateCadence
 * @notice Unit tests for CrossChainRegistry.setTableUpdateCadence
 */
contract CrossChainRegistryUnitTests_setTableUpdateCadence is CrossChainRegistryUnitTests {
    function test_Revert_NotOwner() public {
        uint32 newTableUpdateCadence = 14 days;

        cheats.prank(notPermissioned);
        cheats.expectRevert("Ownable: caller is not the owner");
        crossChainRegistry.setTableUpdateCadence(newTableUpdateCadence);
    }

    function test_setTableUpdateCadence_Success() public {
        uint32 newTableUpdateCadence = 14 days;

        // Expect event
        cheats.expectEmit(true, true, true, true, address(crossChainRegistry));
        emit TableUpdateCadenceSet(newTableUpdateCadence);

        // Set table update cadence
        crossChainRegistry.setTableUpdateCadence(newTableUpdateCadence);

        // Verify state
        assertEq(crossChainRegistry.getTableUpdateCadence(), newTableUpdateCadence, "Table update cadence not set correctly");
    }

    function test_setTableUpdateCadence_AffectsConfigValidation() public {
        // Create a reservation with a config
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        // Update table update cadence to be higher than existing config
        uint32 newTableUpdateCadence = 2 days;
        crossChainRegistry.setTableUpdateCadence(newTableUpdateCadence);

        // Try to set a config with staleness period less than new cadence
        OperatorSetConfig memory invalidConfig = _createOperatorSetConfig(cheats.randomAddress(), 1 days);
        cheats.expectRevert(InvalidStalenessPeriod.selector);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, invalidConfig);

        // Verify setting a valid config works
        OperatorSetConfig memory validConfig = _createOperatorSetConfig(cheats.randomAddress(), 3 days);
        crossChainRegistry.setOperatorSetConfig(defaultOperatorSet, validConfig);

        OperatorSetConfig memory retrievedConfig = crossChainRegistry.getOperatorSetConfig(defaultOperatorSet);
        assertEq(retrievedConfig.maxStalenessPeriod, 3 days, "Valid config should be set");
    }

    function testFuzz_setTableUpdateCadence(uint32 tableUpdateCadence) public {
        tableUpdateCadence = uint32(bound(tableUpdateCadence, 1, 365 days));

        crossChainRegistry.setTableUpdateCadence(tableUpdateCadence);
        assertEq(crossChainRegistry.getTableUpdateCadence(), tableUpdateCadence, "Table update cadence not set correctly");
    }

    function test_Revert_TableUpdateCadenceZero() public {
        cheats.expectRevert(InvalidTableUpdateCadence.selector);
        crossChainRegistry.setTableUpdateCadence(0);
    }
}

/**
 * @title CrossChainRegistryUnitTests_getActiveGenerationReservationsByRange
 * @notice Unit tests for CrossChainRegistry.getActiveGenerationReservationsByRange
 */
contract CrossChainRegistryUnitTests_getActiveGenerationReservationsByRange is CrossChainRegistryUnitTests {
    function test_revert_invalidRange_startGreaterThanEnd() public {
        // Create some reservations first
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        // Test startIndex > endIndex - should revert with InvalidRange
        cheats.expectRevert(InvalidRange.selector);
        crossChainRegistry.getActiveGenerationReservationsByRange(1, 0);
    }

    function test_revert_invalidEndIndex_endGreaterThanLength() public {
        // Create some reservations first
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        // Test endIndex > length - should revert with InvalidEndIndex
        cheats.expectRevert(InvalidEndIndex.selector);
        crossChainRegistry.getActiveGenerationReservationsByRange(0, 2); // length is 1, so endIndex 2 is invalid
    }

    function test_getActiveGenerationReservationsByRange_emptyRegistry() public {
        // Test with empty registry
        cheats.expectRevert(InvalidEndIndex.selector);
        crossChainRegistry.getActiveGenerationReservationsByRange(0, 1); // length is 0, so endIndex 1 is invalid
    }

    function test_emptyRange_startEqualEnd() public {
        // Create some reservations first
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        // Test slice(n,n)
        OperatorSet[] memory reservations = crossChainRegistry.getActiveGenerationReservationsByRange(0, 0);
        assertEq(reservations.length, 0, "Should have 0 reservations");
    }

    function test_getActiveGenerationReservationsByRange_fullRange() public {
        // Create multiple reservations
        OperatorSet[] memory operatorSets = new OperatorSet[](5);
        for (uint i = 0; i < 5; i++) {
            operatorSets[i] = _createOperatorSet(cheats.randomAddress(), uint32(i + 1));
            allocationManagerMock.setIsOperatorSet(operatorSets[i], true);
            _grantUAMRole(address(this), operatorSets[i].avs);
            keyRegistrar.configureOperatorSet(operatorSets[i], CurveType.BN254);
            crossChainRegistry.createGenerationReservation(operatorSets[i], defaultCalculator, defaultConfig);
        }

        // Get full range
        uint totalCount = crossChainRegistry.getActiveGenerationReservationCount();
        OperatorSet[] memory rangeResult = crossChainRegistry.getActiveGenerationReservationsByRange(0, totalCount);
        OperatorSet[] memory fullResult = crossChainRegistry.getActiveGenerationReservations();

        // Verify results are the same
        assertEq(rangeResult.length, fullResult.length, "Range result should match full result length");
        assertEq(rangeResult.length, 5, "Should have 5 reservations");

        for (uint i = 0; i < rangeResult.length; i++) {
            assertEq(rangeResult[i].avs, fullResult[i].avs, "AVS should match");
            assertEq(rangeResult[i].id, fullResult[i].id, "ID should match");
        }
    }

    function test_getActiveGenerationReservationsByRange_DivideInFourths() public {
        // Create 12 reservations to divide evenly by 4
        OperatorSet[] memory operatorSets = new OperatorSet[](12);
        for (uint i = 0; i < 12; i++) {
            operatorSets[i] = _createOperatorSet(cheats.randomAddress(), uint32(i + 1));
            allocationManagerMock.setIsOperatorSet(operatorSets[i], true);
            _grantUAMRole(address(this), operatorSets[i].avs);
            keyRegistrar.configureOperatorSet(operatorSets[i], CurveType.BN254);
            crossChainRegistry.createGenerationReservation(operatorSets[i], defaultCalculator, defaultConfig);
        }

        uint totalCount = crossChainRegistry.getActiveGenerationReservationCount();
        uint quarterSize = totalCount / 4;

        // Get all reservations to compare against
        OperatorSet[] memory allReservations = crossChainRegistry.getActiveGenerationReservations();

        // Collect all reservations by quarters
        OperatorSet[] memory collectedReservations = new OperatorSet[](totalCount);
        uint collectedIndex = 0;

        // Get each quarter
        for (uint quarter = 0; quarter < 4; quarter++) {
            uint startIdx = quarter * quarterSize;
            uint endIdx = (quarter == 3) ? totalCount : (quarter + 1) * quarterSize; // Handle last quarter

            OperatorSet[] memory quarterResult = crossChainRegistry.getActiveGenerationReservationsByRange(startIdx, endIdx);

            // Copy to collected array
            for (uint i = 0; i < quarterResult.length; i++) {
                collectedReservations[collectedIndex] = quarterResult[i];
                collectedIndex++;
            }
        }

        // Verify we collected all reservations and they match
        assertEq(collectedIndex, totalCount, "Should have collected all reservations");

        for (uint i = 0; i < totalCount; i++) {
            assertEq(collectedReservations[i].avs, allReservations[i].avs, "Collected AVS should match");
            assertEq(collectedReservations[i].id, allReservations[i].id, "Collected ID should match");
        }
    }

    function test_getActiveGenerationReservationsByRange_SingleItem() public {
        // Create multiple reservations
        OperatorSet[] memory operatorSets = new OperatorSet[](3);
        for (uint i = 0; i < 3; i++) {
            operatorSets[i] = _createOperatorSet(cheats.randomAddress(), uint32(i + 1));
            allocationManagerMock.setIsOperatorSet(operatorSets[i], true);
            _grantUAMRole(address(this), operatorSets[i].avs);
            keyRegistrar.configureOperatorSet(operatorSets[i], CurveType.BN254);
            crossChainRegistry.createGenerationReservation(operatorSets[i], defaultCalculator, defaultConfig);
        }

        // Get each item individually
        for (uint i = 0; i < 3; i++) {
            OperatorSet[] memory singleResult = crossChainRegistry.getActiveGenerationReservationsByRange(i, i + 1);
            assertEq(singleResult.length, 1, "Should return single item");

            // Verify it matches the expected item from full list
            OperatorSet[] memory allReservations = crossChainRegistry.getActiveGenerationReservations();
            assertEq(singleResult[0].avs, allReservations[i].avs, "Single item AVS should match");
            assertEq(singleResult[0].id, allReservations[i].id, "Single item ID should match");
        }
    }
}

/**
 * @title CrossChainRegistryUnitTests_getActiveGenerationReservationCount
 * @notice Unit tests for CrossChainRegistry.getActiveGenerationReservationCount
 */
contract CrossChainRegistryUnitTests_getActiveGenerationReservationCount is CrossChainRegistryUnitTests {
    function test_getActiveGenerationReservationCount_Empty() public {
        uint count = crossChainRegistry.getActiveGenerationReservationCount();
        assertEq(count, 0, "Should have 0 reservations initially");
    }

    function test_getActiveGenerationReservationCount_Single() public {
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        uint count = crossChainRegistry.getActiveGenerationReservationCount();
        assertEq(count, 1, "Should have 1 reservation");
    }

    function test_getActiveGenerationReservationCount_Multiple() public {
        // Create multiple reservations
        for (uint i = 0; i < 5; i++) {
            OperatorSet memory operatorSet = _createOperatorSet(cheats.randomAddress(), uint32(i + 1));
            allocationManagerMock.setIsOperatorSet(operatorSet, true);
            _grantUAMRole(address(this), operatorSet.avs);
            keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);
            crossChainRegistry.createGenerationReservation(operatorSet, defaultCalculator, defaultConfig);
        }

        uint count = crossChainRegistry.getActiveGenerationReservationCount();
        assertEq(count, 5, "Should have 5 reservations");
    }

    function test_getActiveGenerationReservationCount_AddAndRemove() public {
        // Create reservation
        crossChainRegistry.createGenerationReservation(defaultOperatorSet, defaultCalculator, defaultConfig);

        uint countAfterAdd = crossChainRegistry.getActiveGenerationReservationCount();
        assertEq(countAfterAdd, 1, "Should have 1 reservation after add");

        // Remove reservation
        crossChainRegistry.removeGenerationReservation(defaultOperatorSet);

        uint countAfterRemove = crossChainRegistry.getActiveGenerationReservationCount();
        assertEq(countAfterRemove, 0, "Should have 0 reservations after remove");
    }

    function testFuzz_getActiveGenerationReservationCount_consistency(uint8 numReservations) public {
        numReservations = uint8(bound(numReservations, 0, 20));

        // Create reservations
        for (uint i = 0; i < numReservations; i++) {
            OperatorSet memory operatorSet = _createOperatorSet(cheats.randomAddress(), uint32(i + 1));
            allocationManagerMock.setIsOperatorSet(operatorSet, true);
            _grantUAMRole(address(this), operatorSet.avs);
            keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);
            crossChainRegistry.createGenerationReservation(operatorSet, defaultCalculator, defaultConfig);
        }

        uint count = crossChainRegistry.getActiveGenerationReservationCount();
        OperatorSet[] memory reservations = crossChainRegistry.getActiveGenerationReservations();

        assertEq(count, reservations.length, "Count should always match array length");
        assertEq(count, numReservations, "Count should match expected number");
    }
}
