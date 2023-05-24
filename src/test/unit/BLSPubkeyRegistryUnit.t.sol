//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";


contract BLSPubkeyRegistryUnitTests is Test {

    BLSPubkeyRegistry public blsPubkeyRegistry;
    IRegistryCoordinator public registryCoordinator;

    address registryCoordinatorAddress = address(555);

    setUp() external {
        registryCoordinator = IRegistryCoordinator(registryCoordinatorAddress);
        blsPubkeyRegistry = new BLSPubkeyRegistry(registryCoordinator);

    }

}