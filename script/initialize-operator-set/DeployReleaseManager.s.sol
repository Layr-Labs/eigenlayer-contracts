// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
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

    EmptyContract emptyContract;

    function run() public {
        emptyContract = new EmptyContract(); // NOTE: need some initial contract that's on all chains...

        // Compute the init code for the release manager implementation given constructor arguments.
        bytes memory implementationInitCode =
            abi.encodePacked(type(ReleaseManager).creationCode, abi.encode(permissionController, semver));

        // Deploy the release manager implementation.
        releaseManagerImplementation = ReleaseManager(implementationInitCode.deployCrosschain());

        // Compute the init code for the proxy given constructor arguments.
        bytes memory proxyInitCode =
            abi.encodePacked(type(TransparentUpgradeableProxy).creationCode, abi.encode(emptyContract, msg.sender, ""));

        // Deploy the proxy with EOA as admin, and empty contract as implementation.
        releaseManager = ReleaseManager(proxyInitCode.deployCrosschain());
    }
}
