// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../harnesses/EigenHarness.sol";

import "../../contracts/token/BackingEigen.sol";

contract bEIGENTest is Test {
    mapping(address => bool) fuzzedOutAddresses;

    address initialOwner = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;
    address minterToSet = address(500);
    address mintTo = address(12_345);

    ProxyAdmin proxyAdmin;

    EigenHarness eigenImpl;
    Eigen eigen;

    BackingEigen bEIGENImpl;
    BackingEigen bEIGEN;

    function setUp() public {
        vm.startPrank(initialOwner);
        proxyAdmin = new ProxyAdmin();

        // deploy proxies
        eigen = Eigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));
        bEIGEN = BackingEigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));

        // deploy impls
        eigenImpl = new EigenHarness(IERC20(address(bEIGEN)));
        bEIGENImpl = new BackingEigen(IERC20(address(eigen)));

        // upgrade proxies
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(eigen))), address(eigenImpl));
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(bEIGEN))), address(bEIGENImpl));

        vm.stopPrank();
    }

    function test_Initialize() public {
        bEIGEN.initialize(initialOwner);

        // check that the owner is initialOwner
        assertEq(bEIGEN.owner(), initialOwner);
        // check the transfer restrictions are disabled after one year in the future
        assertEq(bEIGEN.transferRestrictionsDisabledAfter(), type(uint).max);
    }

    function testFuzz_CanBackTheEigenToken(uint eigenSupply) public {
        StdCheats.deal(address(eigen), address(this), eigenSupply);

        bEIGEN.initialize(initialOwner);

        // check that the total supply of bEIGEN is equal to the total supply of EIGEN
        assertEq(bEIGEN.totalSupply(), eigen.totalSupply());
        assertEq(bEIGEN.balanceOf(address(eigen)), bEIGEN.totalSupply());
    }

    function test_setIsMinterAndMint() public {
        bEIGEN.initialize(initialOwner);

        vm.prank(initialOwner);
        bEIGEN.setIsMinter(minterToSet, true);
        require(bEIGEN.isMinter(minterToSet), "minter not set correctly");

        uint amountToMint = 5e25;
        uint balanceBefore = bEIGEN.balanceOf(mintTo);
        vm.prank(minterToSet);
        bEIGEN.mint(mintTo, amountToMint);

        uint balanceAfter = bEIGEN.balanceOf(mintTo);
        uint balanceDiff = balanceAfter - balanceBefore;
        assertEq(balanceDiff, amountToMint, "mint not working correctly");
    }

    function test_setIsMinter_revertsWhenNotCalledByOwner() public {
        bEIGEN.initialize(initialOwner);

        vm.prank(mintTo);
        vm.expectRevert("Ownable: caller is not the owner");
        bEIGEN.setIsMinter(minterToSet, true);
    }

    function test_burn() public {
        test_setIsMinterAndMint();
        vm.prank(initialOwner);
        bEIGEN.setAllowedFrom(mintTo, true);

        uint amountToBurn = 1005e18;
        uint balanceBefore = bEIGEN.balanceOf(mintTo);
        vm.prank(mintTo);
        bEIGEN.burn(amountToBurn);

        uint balanceAfter = bEIGEN.balanceOf(mintTo);
        uint balanceDiff = balanceBefore - balanceAfter;
        assertEq(balanceDiff, amountToBurn, "mint not working correctly");
    }

    function test_mint_revertsWhenNotCalledByMinter() public {
        test_setIsMinterAndMint();

        uint amountToMint = 5e25;
        vm.expectRevert("BackingEigen.mint: caller is not a minter");
        bEIGEN.mint(mintTo, amountToMint);
    }
}
