// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/test/mocks/BN254CertificateVerifierMock.sol";
import "src/test/mocks/ECDSACertificateVerifierMock.sol";
import "src/contracts/multichain/CrossChainRegistry.sol";

abstract contract EigenLayerMultichainUnitTestSetup is EigenLayerUnitTestSetup {
    using StdStyle for *;
    using ArrayLib for *;

    /// @dev In order to test key functionality, we use the actual KeyRegistrar implementation and not a mock
    KeyRegistrar keyRegistrarImplementation;
    KeyRegistrar keyRegistrar;

    /// @dev Mocks
    BN254CertificateVerifierMock bn254CertificateVerifierMock;
    ECDSACertificateVerifierMock ecdsaCertificateVerifierMock;
    CrossChainRegistry crossChainRegistry;
    CrossChainRegistry crossChainRegistryImplementation;

    function setUp() public virtual override {
        // Setup Core Mocks
        super.setUp();

        // Deploy Key Registrar
        keyRegistrarImplementation = new KeyRegistrar(permissionController, IAllocationManager(address(allocationManagerMock)), "9.9.9");
        keyRegistrar =
            KeyRegistrar(address(new TransparentUpgradeableProxy(address(keyRegistrarImplementation), address(eigenLayerProxyAdmin), "")));

        // Deploy mocks
        bn254CertificateVerifierMock = new BN254CertificateVerifierMock();
        ecdsaCertificateVerifierMock = new ECDSACertificateVerifierMock();

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
                        0 // initial paused status
                    )
                )
            )
        );

        // Filter out mocks
        isExcludedFuzzAddress[address(bn254CertificateVerifierMock)] = true;
        isExcludedFuzzAddress[address(ecdsaCertificateVerifierMock)] = true;
        isExcludedFuzzAddress[address(crossChainRegistry)] = true;
    }
}
