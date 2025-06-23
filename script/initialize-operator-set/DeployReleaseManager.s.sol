// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "src/contracts/core/ReleaseManager.sol";
import "src/contracts/permissions/PermissionController.sol";
import "script/utils/CrosschainDeployLib.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract DeployFromScratch is Script, Test {
    using CrosschainDeployLib for *;

    Vm cheats = Vm(VM_ADDRESS);

    PermissionController constant permissionController =
        PermissionController(0x0000000000000000000000000000000000000000);

    string semver = "0.0.0";

    ProxyAdmin proxyAdmin;

    function run() public {
        address emptyContract = type(EmptyContract).creationCode.deployCrosschain();
        ReleaseManager proxy =
            ReleaseManager(address(emptyContract.deployCrosschainProxy({salt: bytes11(uint88(0x1234))})));
        ReleaseManager implementation = new ReleaseManager(permissionController, semver);
        ITransparentUpgradeableProxy(address(proxy)).upgradeTo(address(implementation));
        ITransparentUpgradeableProxy(address(proxy)).changeAdmin(address(proxyAdmin));
    }
}
