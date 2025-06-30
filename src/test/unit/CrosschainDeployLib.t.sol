// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "script/utils/CrosschainDeployLib.sol";

contract CrosschainDeployLibUnitTests is Test {
    address deployer;

    function setUp() public {
        if (block.chainid != 1) vm.skip(true);
        vm.label(address(createx), "createx");
        vm.label(address(this), "test");
    }

    function test_deployCrosschain() public {
        address deployer = msg.sender;
        address computedAddress = CrosschainDeployLib.computeCrosschainAddress(deployer, keccak256(type(EmptyContract).creationCode), "EmptyContract");
        address actualAddress = CrosschainDeployLib.deployCrosschain(deployer, type(EmptyContract).creationCode, "EmptyContract");
        assertEq(computedAddress, actualAddress);
    }
}
