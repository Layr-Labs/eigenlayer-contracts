// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/token/Eigen.sol";
import "../../contracts/token/BackingEigen.sol";

contract EigenWrappingTests is Test {
    mapping(address => bool) fuzzedOutAddresses;

    address minter1 = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;
    address minter2 = 0x87787389BB2Eb2EC8Fe4aA6a2e33D671d925A60f;

    ProxyAdmin proxyAdmin;

    Eigen eigenImpl;
    Eigen eigen;

    BackingEigen bEIGENImpl;
    BackingEigen bEIGEN;

    uint totalSupply = 1.67e9 ether;

    // EVENTS FROM BackingEigen.sol
    /// @notice event emitted when a minter status is modified
    event IsMinterModified(address indexed minterAddress, bool newStatus);

    // EVENTS FROM Eigen.sol
    /// @notice event emitted when bEIGEN tokens are wrapped into EIGEN
    event TokenWrapped(address indexed account, uint amount);
    /// @notice event emitted when EIGEN tokens are unwrapped into bEIGEN
    event TokenUnwrapped(address indexed account, uint amount);

    modifier filterAddress(address fuzzedAddress) {
        vm.assume(!fuzzedOutAddresses[fuzzedAddress]);
        _;
    }

    function setUp() public {
        vm.startPrank(minter1);
        proxyAdmin = new ProxyAdmin();

        // deploy proxies
        eigen = Eigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));
        bEIGEN = BackingEigen(address(new TransparentUpgradeableProxy(address(proxyAdmin), address(proxyAdmin), "")));

        // deploy impls
        eigenImpl = new Eigen(IERC20(address(bEIGEN)));
        bEIGENImpl = new BackingEigen();

        // upgrade proxies
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(eigen))), address(eigenImpl));
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(bEIGEN))), address(bEIGENImpl));

        // initialize bEIGEN - this will mint the entire supply to the Eigen contract
        bEIGEN.initialize(minter1);

        // set minters (for future minting if needed)
        vm.expectEmit(true, false, false, false);
        emit IsMinterModified(minter1, true);
        bEIGEN.setIsMinter(minter1, true);

        vm.expectEmit(true, false, false, false);
        emit IsMinterModified(minter2, true);
        bEIGEN.setIsMinter(minter2, true);

        // initialize eigen with empty arrays since we don't need minting anymore
        eigen.initialize(minter1);

        // Mint and wrap tokens for minter1
        vm.startPrank(minter1);
        bEIGEN.mint(minter1, totalSupply / 2);
        bEIGEN.approve(address(eigen), totalSupply / 2);
        vm.expectEmit(true, false, false, false);
        emit TokenWrapped(minter1, totalSupply / 2);
        eigen.wrap(totalSupply / 2);
        vm.stopPrank();

        // Mint and wrap tokens for minter2
        vm.startPrank(minter2);
        bEIGEN.mint(minter2, totalSupply / 2);
        bEIGEN.approve(address(eigen), totalSupply / 2);
        vm.expectEmit(true, false, false, false);
        emit TokenWrapped(minter2, totalSupply / 2);
        eigen.wrap(totalSupply / 2);
        vm.stopPrank();

        fuzzedOutAddresses[minter1] = true;
        fuzzedOutAddresses[minter2] = true;
        fuzzedOutAddresses[address(proxyAdmin)] = true;
        fuzzedOutAddresses[address(eigen)] = true;
        fuzzedOutAddresses[address(bEIGEN)] = true;
        fuzzedOutAddresses[address(0)] = true;
    }

    function test_AnyoneCanUnwrap(address unwrapper, uint unwrapAmount) public filterAddress(unwrapper) {
        vm.assume(unwrapper != address(0));

        // minter1 balance
        uint minter1Balance = eigen.balanceOf(minter1);

        // send EIGEN to unwrapper
        vm.prank(minter1);
        eigen.transfer(unwrapper, minter1Balance);

        // initial bEIGEN balance
        uint initialBEIGENBalanceOfEigenToken = bEIGEN.balanceOf(address(eigen));
        // initial EIGEN token supply
        assertEq(eigen.totalSupply(), bEIGEN.totalSupply(), "eigen totalSupply changed incorrectly");

        // unwrap
        // unwrap amount should be less than minter1 balance
        unwrapAmount = unwrapAmount % minter1Balance;
        vm.prank(unwrapper);
        vm.expectEmit(true, false, false, false);
        emit TokenUnwrapped(unwrapper, unwrapAmount);
        eigen.unwrap(unwrapAmount);

        // check total supply and balance changes
        assertEq(eigen.totalSupply(), bEIGEN.totalSupply(), "eigen totalSupply changed incorrectly");
        assertEq(
            bEIGEN.balanceOf(address(eigen)),
            initialBEIGENBalanceOfEigenToken - unwrapAmount,
            "beigen balance of EIGEN tokens changed incorrectly"
        );
        assertEq(eigen.balanceOf(address(unwrapper)), minter1Balance - unwrapAmount);
        assertEq(bEIGEN.balanceOf(address(unwrapper)), unwrapAmount);
    }

    function test_AnyoneCanWrap(address wrapper, uint wrapAmount) public filterAddress(wrapper) {
        vm.assume(wrapper != address(0));

        // initial bEIGEN balance
        uint initialBEIGENBalanceOfEigenToken = bEIGEN.balanceOf(address(eigen));
        // minter1 balance
        uint minter1Balance = eigen.balanceOf(minter1);

        // unwrap
        vm.startPrank(minter1);
        eigen.unwrap(minter1Balance);

        // send bEIGEN to wrapper
        bEIGEN.transfer(wrapper, minter1Balance);
        vm.stopPrank();

        // initial EIGEN token supply
        assertEq(eigen.totalSupply(), bEIGEN.totalSupply(), "eigen totalSupply changed incorrectly");

        // wrap
        // wrap amount should be less than minter1 balance
        wrapAmount = wrapAmount % minter1Balance;
        vm.startPrank(wrapper);
        // approve bEIGEN
        bEIGEN.approve(address(eigen), wrapAmount);
        // wrap
        vm.expectEmit(true, false, false, false);
        emit TokenWrapped(wrapper, wrapAmount);
        eigen.wrap(wrapAmount);
        vm.stopPrank();

        // check total supply and balance changes
        assertEq(eigen.totalSupply(), bEIGEN.totalSupply(), "eigen totalSupply changed incorrectly");
        assertEq(bEIGEN.balanceOf(address(eigen)), initialBEIGENBalanceOfEigenToken - minter1Balance + wrapAmount);
        assertEq(eigen.balanceOf(address(wrapper)), wrapAmount);
        assertEq(bEIGEN.balanceOf(address(wrapper)), minter1Balance - wrapAmount);
    }

    function test_CannotUnwrapMoreThanBalance(address unwrapper, uint unwrapAmount) public filterAddress(unwrapper) {
        // unwrap amount should be less than minter1 balance
        unwrapAmount = unwrapAmount % eigen.balanceOf(minter1);

        // send EIGEN to unwrapper
        vm.prank(minter1);
        eigen.transfer(unwrapper, unwrapAmount);

        // unwrap
        vm.prank(unwrapper);
        vm.expectRevert("ERC20: burn amount exceeds balance");
        eigen.unwrap(unwrapAmount + 1);
    }

    function test_CannotWrapMoreThanBalance(address wrapper, uint wrapAmount) public filterAddress(wrapper) {
        // wrap amount should be less than minter1 balance
        wrapAmount = wrapAmount % eigen.balanceOf(minter1);

        // unwrap
        vm.startPrank(minter1);
        eigen.unwrap(wrapAmount);
        bEIGEN.transfer(wrapper, wrapAmount);
        vm.stopPrank();

        // wrap
        vm.startPrank(wrapper);
        // approve bEIGEN
        bEIGEN.approve(address(eigen), type(uint).max);
        // send bEIGEN to wrapper
        vm.expectRevert("ERC20: transfer amount exceeds balance");
        eigen.wrap(wrapAmount + 1);
        vm.stopPrank();
    }
}
