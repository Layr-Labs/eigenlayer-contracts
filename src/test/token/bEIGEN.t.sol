// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../harnesses/EigenHarness.sol";

import "../../contracts/token/BackingEigen.sol";


contract bEIGENTest is Test {
    mapping(address => bool) fuzzedOutAddresses;

    address minter1 = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;

    ProxyAdmin proxyAdmin;

    EigenHarness eigenImpl;
    Eigen eigen;

    BackingEigen bEIGENImpl;
    BackingEigen bEIGEN;


    function setUp() public {
        vm.startPrank(minter1);
        proxyAdmin = new ProxyAdmin();

        // deploy proxies
        eigen = Eigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));
        bEIGEN = BackingEigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));

        // deploy impls
        eigenImpl = new EigenHarness(IERC20(address(bEIGEN)));
        bEIGENImpl = new BackingEigen(IERC20(address(eigen)));

        // upgrade proxies
        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(eigen))), address(eigenImpl));
        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(bEIGEN))), address(bEIGENImpl));

        vm.stopPrank();
    }

    function test_Initialize() public {
        bEIGEN.initialize(minter1);
        
        // check that the owner is minter1
        assertEq(bEIGEN.owner(), minter1);
        // check the transfer restrictions are disabled after one year in the future
        assertEq(bEIGEN.transferRestrictionsDisabledAfter(), type(uint256).max);
    }

    function testFuzz_CanBackTheEigenToken(uint eigenSupply) public {
        StdCheats.deal(address(eigen), address(this), eigenSupply);

        bEIGEN.initialize(minter1);
        
        // check that the total supply of bEIGEN is equal to the total supply of EIGEN
        assertEq(bEIGEN.totalSupply(), eigen.totalSupply());
        assertEq(bEIGEN.balanceOf(address(eigen)), bEIGEN.totalSupply());
    }
}
