// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";
import {EmptyContract} from "./EmptyContract.sol";

contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        deployImpl({
            name: type(EmptyContract).name,
            deployedTo: address(new EmptyContract())
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
    }
}
