// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/ProtocolRegistry.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract ProtocolRegistryUnitTests is EigenLayerUnitTestSetup, IProtocolRegistryEvents, IProtocolRegistryErrors {
    ProtocolRegistry protocolRegistry;
    ProxyAdminMock proxyAdminMock;

    address defaultOwner;
    address nonOwner;
    address pauserMultisig;

    string internal currentVersion;
    uint internal currentMajorVersion;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        defaultOwner = address(this);
        nonOwner = cheats.randomAddress();
        pauserMultisig = cheats.randomAddress();

        proxyAdminMock = new ProxyAdminMock();
        protocolRegistry = _deployProtocolRegistry(address(proxyAdminMock));
    }

    function _deployProtocolRegistry(address proxyAdmin) internal returns (ProtocolRegistry registry) {
        registry = ProtocolRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(new ProtocolRegistry()),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(ProtocolRegistry.initialize.selector, defaultOwner, pauserMultisig)
                )
            )
        );
        isExcludedFuzzAddress[address(registry)] = true;
    }

    /// -----------------------------------------------------------------------
    /// initialize()
    /// -----------------------------------------------------------------------

    function test_initialize_Correctness() public {
        assertEq(protocolRegistry.hasRole(protocolRegistry.DEFAULT_ADMIN_ROLE(), defaultOwner), true);
        assertEq(protocolRegistry.hasRole(protocolRegistry.PAUSER_ROLE(), pauserMultisig), true);
        assertEq(protocolRegistry.getRoleMemberCount(protocolRegistry.DEFAULT_ADMIN_ROLE()), 1);
        assertEq(protocolRegistry.getRoleMemberCount(protocolRegistry.PAUSER_ROLE()), 1);
        assertEq(protocolRegistry.getRoleMember(protocolRegistry.DEFAULT_ADMIN_ROLE(), 0), defaultOwner);
        assertEq(protocolRegistry.getRoleMember(protocolRegistry.PAUSER_ROLE(), 0), pauserMultisig);
        cheats.expectRevert("Initializable: contract is already initialized");
        protocolRegistry.initialize(defaultOwner, pauserMultisig);
    }

    /// -----------------------------------------------------------------------
    /// ship()
    /// -----------------------------------------------------------------------

    function test_ship_OnlyOwner() public {
        address[] memory addresses = new address[](1);
        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        string[] memory names = new string[](1);

        cheats.prank(nonOwner);
        cheats.expectRevert();
        protocolRegistry.ship(addresses, configs, names, "1.0.0");
    }

    function test_ship_SingleDeployment() public {
        address addr = cheats.randomAddress();
        address[] memory addresses = new address[](1);
        addresses[0] = addr;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        string[] memory names = new string[](1);
        names[0] = "TestContract";

        string memory version = "1.0.0";

        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit SemanticVersionUpdated("", version);
        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentShipped(addr, configs[0]);

        protocolRegistry.ship(addresses, configs, names, version);

        // Check version getter
        assertEq(protocolRegistry.version(), version);

        // Check major version getter
        assertEq(protocolRegistry.majorVersion(), "1");

        currentVersion = version;
        currentMajorVersion = 1;

        assertEq(protocolRegistry.totalDeployments(), 1);
        assertEq(protocolRegistry.getAddress("TestContract"), addr);
    }

    function test_ship_MultipleDeployments() public {
        address addr1 = address(0x1);
        address addr2 = address(0x2);
        address addr3 = address(0x3);

        address[] memory addresses = new address[](3);
        addresses[0] = addr1;
        addresses[1] = addr2;
        addresses[2] = addr3;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](3);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[2] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: true});

        string[] memory names = new string[](3);
        names[0] = "Contract1";
        names[1] = "Contract2";
        names[2] = "Contract3";

        string memory version = "2.0.0";

        protocolRegistry.ship(addresses, configs, names, version);

        // Check version getter
        assertEq(protocolRegistry.version(), version);
        // Check major version getter
        assertEq(protocolRegistry.majorVersion(), "2");

        currentVersion = version;
        currentMajorVersion = 2;

        assertEq(protocolRegistry.totalDeployments(), 3);
        assertEq(protocolRegistry.getAddress("Contract1"), addr1);
        assertEq(protocolRegistry.getAddress("Contract2"), addr2);
        assertEq(protocolRegistry.getAddress("Contract3"), addr3);
    }

    /// -----------------------------------------------------------------------
    /// configure()
    /// -----------------------------------------------------------------------

    function test_configure_OnlyOwner() public {
        address addr = cheats.randomAddress();
        IProtocolRegistryTypes.DeploymentConfig memory config = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        cheats.prank(nonOwner);
        cheats.expectRevert();
        protocolRegistry.configure(addr, config);
    }

    function test_configure_Correctness() public {
        // First ship a deployment
        address addr = address(0x123);
        address[] memory addresses = new address[](1);
        addresses[0] = addr;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](1);
        names[0] = "TestContract";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check version getter
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        // Update config
        IProtocolRegistryTypes.DeploymentConfig memory newConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: true});

        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentConfigured(addr, newConfig);

        protocolRegistry.configure(addr, newConfig);

        (, IProtocolRegistryTypes.DeploymentConfig memory retrievedConfig) = protocolRegistry.getDeployment("TestContract");
        assertEq(retrievedConfig.pausable, true);
        assertEq(retrievedConfig.deprecated, true);
    }

    /// -----------------------------------------------------------------------
    /// pauseAll()
    /// -----------------------------------------------------------------------

    function test_pauseAll_OnlyPauserMultisig() public {
        cheats.prank(nonOwner);
        cheats.expectRevert();
        protocolRegistry.pauseAll();
    }

    function test_pauseAll_PausableContracts() public {
        PausableMock pausable1 = new PausableMock();
        PausableMock pausable2 = new PausableMock();
        address nonPausable = address(emptyContract);

        address[] memory addresses = new address[](3);
        addresses[0] = address(pausable1);
        addresses[1] = nonPausable;
        addresses[2] = address(pausable2);

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](3);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[2] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        string[] memory names = new string[](3);
        names[0] = "Pausable1";
        names[1] = "NonPausable";
        names[2] = "Pausable2";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check hecks for version and major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        cheats.prank(pauserMultisig);
        protocolRegistry.pauseAll();

        assertTrue(pausable1.paused());
        assertTrue(pausable2.paused());
    }

    function test_pauseAll_SkipsDeprecated() public {
        PausableMock pausable = new PausableMock();
        PausableMock deprecated = new PausableMock();

        address[] memory addresses = new address[](2);
        addresses[0] = address(pausable);
        addresses[1] = address(deprecated);

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: true});

        string[] memory names = new string[](2);
        names[0] = "Active";
        names[1] = "Deprecated";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check hecks for version and major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        cheats.prank(pauserMultisig);
        protocolRegistry.pauseAll();

        assertTrue(pausable.paused());
        assertFalse(deprecated.paused());
    }

    /// -----------------------------------------------------------------------
    /// getAddress()
    /// -----------------------------------------------------------------------

    function test_getAddress_ExistingDeployment() public {
        address addr = address(0x456);
        address[] memory addresses = new address[](1);
        addresses[0] = addr;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](1);
        names[0] = "MyContract";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check version and major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        assertEq(protocolRegistry.getAddress("MyContract"), addr);
    }

    /// -----------------------------------------------------------------------
    /// getDeployment()
    /// -----------------------------------------------------------------------

    function test_getDeployment() public {
        address addr = address(0x789);
        address[] memory addresses = new address[](1);
        addresses[0] = addr;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        string[] memory names = new string[](1);
        names[0] = "MyContract";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check version and major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        (address retAddr, IProtocolRegistryTypes.DeploymentConfig memory config) = protocolRegistry.getDeployment("MyContract");

        assertEq(retAddr, addr);
        assertEq(config.pausable, true);
        assertEq(config.deprecated, false);
    }

    /// -----------------------------------------------------------------------
    /// getAllDeployments()
    /// -----------------------------------------------------------------------

    function test_getAllDeployments_Empty() public {
        (string[] memory names, address[] memory addresses, IProtocolRegistryTypes.DeploymentConfig[] memory configs) =
            protocolRegistry.getAllDeployments();

        assertEq(names.length, 0);
        assertEq(addresses.length, 0);
        assertEq(configs.length, 0);
    }

    function test_getAllDeployments_Multiple() public {
        address addr1 = address(0x1111);
        address addr2 = address(0x2222);

        address[] memory addresses = new address[](2);
        addresses[0] = addr1;
        addresses[1] = addr2;

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: true});

        string[] memory names = new string[](2);
        names[0] = "First";
        names[1] = "Second";

        string memory version = "3.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Check version & major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "3");

        (string[] memory retNames, address[] memory retAddresses, IProtocolRegistryTypes.DeploymentConfig[] memory retConfigs) =
            protocolRegistry.getAllDeployments();

        assertEq(retNames.length, 2);
        assertEq(retAddresses.length, 2);
        assertEq(retConfigs.length, 2);

        assertEq(retNames[0], "First");
        assertEq(retNames[1], "Second");
        assertEq(retAddresses[0], addr1);
        assertEq(retAddresses[1], addr2);
        assertEq(retConfigs[0].pausable, true);
    }

    /// -----------------------------------------------------------------------
    /// totalDeployments()
    /// -----------------------------------------------------------------------

    function test_totalDeployments_InitiallyZero() public {
        assertEq(protocolRegistry.totalDeployments(), 0);
    }

    function test_totalDeployments_IncreasesWithShip() public {
        address[] memory addresses = new address[](2);
        addresses[0] = address(0x1);
        addresses[1] = address(0x2);

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](2);
        names[0] = "A";
        names[1] = "B";

        string memory version = "1.0.0";
        protocolRegistry.ship(addresses, configs, names, version);

        // Test version and major version
        assertEq(protocolRegistry.version(), version);
        assertEq(protocolRegistry.majorVersion(), "1");

        assertEq(protocolRegistry.totalDeployments(), 2);
    }

    /// -----------------------------------------------------------------------
    /// Fix 7: ship() validation tests
    /// -----------------------------------------------------------------------

    function test_ship_revert_arrayLengthMismatch_addressesConfigs() public {
        address[] memory addresses = new address[](2);
        addresses[0] = address(0x1);
        addresses[1] = address(0x2);

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](2);
        names[0] = "A";
        names[1] = "B";

        cheats.expectRevert(ArrayLengthMismatch.selector);
        protocolRegistry.ship(addresses, configs, names, "1.0.0");
    }

    function test_ship_revert_arrayLengthMismatch_configsNames() public {
        address[] memory addresses = new address[](2);
        addresses[0] = address(0x1);
        addresses[1] = address(0x2);

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](1);
        names[0] = "A";

        cheats.expectRevert(ArrayLengthMismatch.selector);
        protocolRegistry.ship(addresses, configs, names, "1.0.0");
    }

    function test_ship_revert_zeroAddress() public {
        address[] memory addresses = new address[](2);
        addresses[0] = address(0x1);
        addresses[1] = address(0); // Zero address

        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](2);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        configs[1] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});

        string[] memory names = new string[](2);
        names[0] = "A";
        names[1] = "B";

        cheats.expectRevert(InputAddressZero.selector);
        protocolRegistry.ship(addresses, configs, names, "1.0.0");
    }

    /// -----------------------------------------------------------------------
    /// Fix 8: ship() orphaned configs cleanup test
    /// -----------------------------------------------------------------------

    function test_ship_overwriteName_deletesOldConfig() public {
        address oldAddr = address(0x111);
        address newAddr = address(0x222);
        string memory name = "TestContract";

        // Ship with old address
        address[] memory addresses1 = new address[](1);
        addresses1[0] = oldAddr;
        IProtocolRegistryTypes.DeploymentConfig[] memory configs1 = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs1[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        string[] memory names1 = new string[](1);
        names1[0] = name;
        protocolRegistry.ship(addresses1, configs1, names1, "1.0.0");

        // Verify old address is set
        assertEq(protocolRegistry.getAddress(name), oldAddr);

        // Ship same name with new address - should emit DeploymentConfigDeleted for old address
        address[] memory addresses2 = new address[](1);
        addresses2[0] = newAddr;
        IProtocolRegistryTypes.DeploymentConfig[] memory configs2 = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs2[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: true});
        string[] memory names2 = new string[](1);
        names2[0] = name;

        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentConfigDeleted(oldAddr);
        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentShipped(newAddr, configs2[0]);

        protocolRegistry.ship(addresses2, configs2, names2, "1.1.0");

        // Verify new address is set
        assertEq(protocolRegistry.getAddress(name), newAddr);
        // Verify total deployments is still 1 (not 2)
        assertEq(protocolRegistry.totalDeployments(), 1);
    }

    function test_ship_sameNameSameAddress_noDeleteEvent() public {
        address addr = address(0x111);
        string memory name = "TestContract";

        // Ship with address
        address[] memory addresses1 = new address[](1);
        addresses1[0] = addr;
        IProtocolRegistryTypes.DeploymentConfig[] memory configs1 = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs1[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});
        string[] memory names1 = new string[](1);
        names1[0] = name;
        protocolRegistry.ship(addresses1, configs1, names1, "1.0.0");

        // Ship same name with same address - should NOT emit DeploymentConfigDeleted
        IProtocolRegistryTypes.DeploymentConfig[] memory configs2 = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs2[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: true});

        // We only expect DeploymentShipped, not DeploymentConfigDeleted
        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentShipped(addr, configs2[0]);

        protocolRegistry.ship(addresses1, configs2, names1, "1.1.0");
    }

    /// -----------------------------------------------------------------------
    /// Fix 9: configure() for unshipped addresses test
    /// -----------------------------------------------------------------------

    function test_configure_revert_unshippedAddress() public {
        address unshipped = address(0x999);
        IProtocolRegistryTypes.DeploymentConfig memory config = IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: false});

        cheats.expectRevert(DeploymentNotShipped.selector);
        protocolRegistry.configure(unshipped, config);
    }

    function test_configure_succeeds_shippedAddress() public {
        address shipped = address(0x111);

        // First ship the address
        address[] memory addresses = new address[](1);
        addresses[0] = shipped;
        IProtocolRegistryTypes.DeploymentConfig[] memory configs = new IProtocolRegistryTypes.DeploymentConfig[](1);
        configs[0] = IProtocolRegistryTypes.DeploymentConfig({pausable: false, deprecated: false});
        string[] memory names = new string[](1);
        names[0] = "Test";
        protocolRegistry.ship(addresses, configs, names, "1.0.0");

        // Now configure should succeed
        IProtocolRegistryTypes.DeploymentConfig memory newConfig =
            IProtocolRegistryTypes.DeploymentConfig({pausable: true, deprecated: true});

        cheats.expectEmit(true, true, true, true, address(protocolRegistry));
        emit DeploymentConfigured(shipped, newConfig);

        protocolRegistry.configure(shipped, newConfig);

        (, IProtocolRegistryTypes.DeploymentConfig memory retrievedConfig) = protocolRegistry.getDeployment("Test");
        assertEq(retrievedConfig.pausable, true);
        assertEq(retrievedConfig.deprecated, true);
    }
}

// Mock contracts for testing
contract PausableMock {
    bool private _paused;

    function pauseAll() external {
        _paused = true;
    }

    function paused() external view returns (bool) {
        return _paused;
    }
}

contract ProxyAdminMock {
    mapping(address => address) private _implementations;

    function setImplementation(address proxy, address implementation) external {
        _implementations[proxy] = implementation;
    }

    function getProxyImplementation(address proxy) external view returns (address) {
        return _implementations[proxy];
    }
}
