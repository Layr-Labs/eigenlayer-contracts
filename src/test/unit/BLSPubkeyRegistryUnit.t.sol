//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";
import "../../contracts/middleware/BLSPublicKeyCompendium.sol";
import "../mocks/RegistryCoordinatorMock.sol";


contract BLSPubkeyRegistryUnitTests is Test {

    BLSPubkeyRegistry public blsPubkeyRegistry;
    BLSPublicKeyCompendium public pkCompendium;
    RegistryCoordinatorMock public registryCoordinator;

    function setUp() external {
        registryCoordinator = new RegistryCoordinatorMock();
        pkCompendium = new BLSPublicKeyCompendium();
        blsPubkeyRegistry = new BLSPubkeyRegistry(registryCoordinator, pkCompendium);
    }

    function testConstructorArgs() public {
        require(blsPubkeyRegistry.registryCoordinator() == registryCoordinator, "registryCoordinator not set correctly");
        require(blsPubkeyRegistry.pubkeyCompendium() == pkCompendium, "pubkeyCompendium not set correctly");
    }

}