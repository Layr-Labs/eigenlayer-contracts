// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "src/contracts/core/ReleaseManager.sol";
import "src/contracts/permissions/PermissionController.sol";
import "../utils/CrosschainDeployLib.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract EmptyContract {}

contract DeployFromScratch is Script, Test {
    using CrosschainDeployLib for *;

    Vm cheats = Vm(VM_ADDRESS);

    PermissionController constant permissionController =
        PermissionController(0x0000000000000000000000000000000000000000);
    string semver = "0.0.0";

    ReleaseManager releaseManager;
    ReleaseManager releaseManagerImplementation;
    ProxyAdmin eigenLayerProxyAdmin;
    address emptyContract;

    function run() public {
        // Deploy the empty contract, MUST be on all chains at the same address.
        emptyContract = type(EmptyContract).creationCode.deployCrosschain();
        // Deploy the release manager implementation.
        releaseManagerImplementation = new ReleaseManager(permissionController, semver);
        // Deploy the proxy with EOA as admin, and empty contract as implementation.
        releaseManager = ReleaseManager(address(emptyContract.deployCrosschainProxy(msg.sender, "")));

        // Upgrade the proxy to the release manager implementation and set the admin to the proxy admin.
        ITransparentUpgradeableProxy proxy = ITransparentUpgradeableProxy(address(releaseManager));
        proxy.upgradeTo(address(releaseManagerImplementation));
        proxy.changeAdmin(address(eigenLayerProxyAdmin));
    }
}
