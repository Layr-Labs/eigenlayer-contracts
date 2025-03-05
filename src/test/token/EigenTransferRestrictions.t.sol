// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "../harnesses/EigenHarness.sol";

contract EigenTransferRestrictionsTest is Test {
    mapping(address => bool) fuzzedOutAddresses;

    address minter1 = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;
    address minter2 = 0x87787389BB2Eb2EC8Fe4aA6a2e33D671d925A60f;

    ProxyAdmin proxyAdmin;

    EigenHarness eigenImpl;
    Eigen eigen;
    uint totalSupply = 1.67e9 ether;

    // EVENTS FROM EIGEN.sol
    /// @notice event emitted when the allowedFrom status of an address is set
    event SetAllowedFrom(address indexed from, bool isAllowedFrom);
    /// @notice event emitted when the allowedTo status of an address is set
    event SetAllowedTo(address indexed to, bool isAllowedTo);
    /// @notice event emitted when a minter mints
    event Mint(address indexed minter, uint amount);
    /// @notice event emitted when the transfer restrictions are disabled
    event TransferRestrictionsDisabled();

    modifier filterAddress(address fuzzedAddress) {
        vm.assume(!fuzzedOutAddresses[fuzzedAddress]);
        _;
    }

    function setUp() public {
        vm.startPrank(minter1);
        proxyAdmin = new ProxyAdmin();
        // initialize with dummy BackingEigen address

        eigenImpl =
            new EigenHarness(new ERC20PresetFixedSupply({name: "bEIGEN", symbol: "bEIGEN", initialSupply: totalSupply, owner: minter1}));
        eigen = Eigen(address(new TransparentUpgradeableProxy(address(eigenImpl), address(proxyAdmin), "")));
        eigen.bEIGEN().transfer(address(eigen), totalSupply);
        vm.stopPrank();

        fuzzedOutAddresses[minter1] = true;
        fuzzedOutAddresses[minter2] = true;
        fuzzedOutAddresses[address(proxyAdmin)] = true;
        fuzzedOutAddresses[address(eigen)] = true;
        fuzzedOutAddresses[address(0)] = true;
    }

    function test_AllowedFromCanSendAnywhere(address to) public filterAddress(to) {
        _simulateMint();

        // minter1 and eigenminter2 are already allowedFrom
        vm.startPrank(minter1);
        eigen.transfer(to, eigen.balanceOf(minter1) / 2);

        vm.startPrank(minter2);
        eigen.transfer(to, eigen.balanceOf(minter2) / 2);
    }

    function test_CanSetAllowedFrom(address from, address to) public filterAddress(from) filterAddress(to) {
        _simulateMint();

        // set allowedFrom[from] = true
        vm.startPrank(minter1);
        vm.expectEmit(true, true, true, true, address(eigen));
        emit SetAllowedFrom(from, true);
        eigen.setAllowedFrom(from, true);
        assertEq(eigen.allowedFrom(from), true, "EigenTest.test_CanSetAllowedFrom: allowedFrom was not set correctly");

        // new allowedFrom can send
        vm.startPrank(from);
        eigen.transfer(to, eigen.balanceOf(from) / 2);

        // set allowedFrom[from] = false
        vm.startPrank(minter1);
        vm.expectEmit(true, true, true, true, address(eigen));
        emit SetAllowedFrom(from, false);
        eigen.setAllowedFrom(from, false);
        assertEq(eigen.allowedFrom(from), false, "EigenTest.test_CanSetAllowedFrom: allowedFrom was not set correctly");
    }

    function test_OnlyOwnerCanSetAllowedFrom(address notOwner) public filterAddress(notOwner) {
        _simulateMint();

        // set allowedFrom[from] = true
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        eigen.setAllowedFrom(notOwner, true);
    }

    function test_NotAllowedFromCannotSendIfNoAllowedTos(address from, address to) public filterAddress(from) filterAddress(to) {
        _simulateMint();

        // send other tokens from minter1
        vm.startPrank(minter1);
        eigen.transfer(from, eigen.balanceOf(minter1) / 2);

        // sending from other will revert
        vm.startPrank(from);
        uint fromBalance = eigen.balanceOf(from);
        vm.expectRevert("Eigen._beforeTokenTransfer: from or to must be whitelisted");
        eigen.transfer(to, fromBalance / 2);
    }

    function test_CanSetAllowedTo(address from, address to) public filterAddress(from) filterAddress(to) {
        _simulateMint();

        // set allowedFrom[from] = true
        vm.startPrank(minter1);
        vm.expectEmit(true, true, true, true, address(eigen));
        emit SetAllowedTo(to, true);
        eigen.setAllowedTo(to, true);
        assertEq(eigen.allowedTo(to), true, "EigenTest.test_CanSetAllowedTo: allowedTo was not set correctly");

        // new allowedFrom can send
        vm.startPrank(from);
        eigen.transfer(to, eigen.balanceOf(to) / 2);

        // set allowedFrom[from] = false
        vm.startPrank(minter1);
        vm.expectEmit(true, true, true, true, address(eigen));
        emit SetAllowedTo(to, false);
        eigen.setAllowedTo(to, false);
        assertEq(eigen.allowedTo(to), false, "EigenTest.test_CanSetAllowedTo: allowedTo was not set correctly");
    }

    function test_OnlyOwnerCanSetAllowedTo(address notOwner) public filterAddress(notOwner) {
        _simulateMint();

        // set allowedFrom[from] = true
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        eigen.setAllowedTo(notOwner, true);
    }

    function test_disableTransferRestrictions(address from, address to) public filterAddress(from) filterAddress(to) {
        _simulateMint();

        vm.startPrank(minter1);
        eigen.transfer(from, eigen.balanceOf(minter1) / 2);

        // transfer will revert
        vm.startPrank(from);
        uint fromBalance = eigen.balanceOf(from);
        vm.expectRevert("Eigen._beforeTokenTransfer: from or to must be whitelisted");
        eigen.transfer(to, fromBalance / 2);

        assertEq(eigen.transferRestrictionsDisabledAfter(), type(uint).max, "invalid test setup");

        // set transfer restrictions to be disabled after one year in the future
        vm.startPrank(minter1);
        vm.expectEmit(true, true, true, true, address(eigen));
        emit TransferRestrictionsDisabled();
        eigen.disableTransferRestrictions();
        assertEq(
            eigen.transferRestrictionsDisabledAfter(),
            0,
            "EigenTest.test_disableTransferRestrictions: transfer restrictions were not disabled correctly"
        );

        vm.startPrank(from);
        // transfer restrictions are disabled
        assertTrue(eigen.transfer(to, fromBalance / 2), "transfer returned false");
    }

    function _simulateMint() internal {
        // dummy mint
        EigenHarness(address(eigen)).mint(minter1, totalSupply / 2);
        EigenHarness(address(eigen)).mint(minter2, totalSupply / 2);

        // set allowed froms
        EigenHarness(address(eigen)).setAllowedFromPermissionless(minter1, true);
        EigenHarness(address(eigen)).setAllowedFromPermissionless(minter2, true);

        // set transfer restrictions to be disabled after to max
        EigenHarness(address(eigen)).setTransferRestrictionsDisabledAfterToMax();

        // set owner to minter1
        EigenHarness(address(eigen)).transferOwnershipPermissionless(minter1);
    }
}
