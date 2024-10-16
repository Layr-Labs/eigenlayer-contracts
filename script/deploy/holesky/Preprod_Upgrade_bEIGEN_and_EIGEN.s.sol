// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "../../../src/contracts/token/BackingEigen.sol";
import "../../../src/contracts/token/Eigen.sol";

// forge script script/deploy/holesky/Preprod_Upgrade_bEIGEN_and_EIGEN.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --verify --etherscan-api-key $ETHERSCAN_API_KEY
contract Preprod_Upgrade_bEIGEN_and_EIGEN is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    BackingEigen public bEIGEN_proxy = BackingEigen(0xA72942289a043874249E60469F68f08B8c6ECCe8);
    BackingEigen public bEIGEN_implementation;
    Eigen public EIGEN_proxy = Eigen(0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926);
    Eigen public EIGEN_implementation;
    ProxyAdmin public EIGEN_ProxyAdmin = ProxyAdmin(0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B);
    address public proxyAdminOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    IERC20 public bEIGEN_addressBefore;
    IERC20 public EIGEN_addressBefore;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId != 17000) {
            revert("Chain not supported");
        }

        bEIGEN_addressBefore = EIGEN_proxy.bEIGEN();
        require(bEIGEN_addressBefore == IERC20(0xA72942289a043874249E60469F68f08B8c6ECCe8),
            "something horribly wrong");

        EIGEN_addressBefore = bEIGEN_proxy.EIGEN();
        require(EIGEN_addressBefore == IERC20(0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926),
            "something horribly wrong");

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contracts
        EIGEN_implementation = new Eigen({
            _bEIGEN: bEIGEN_addressBefore
        });
        bEIGEN_implementation = new BackingEigen({
            _EIGEN: EIGEN_addressBefore
        });

        vm.stopBroadcast();

        emit log_named_address("EIGEN_implementation", address(EIGEN_implementation));
        emit log_named_address("bEIGEN_implementation", address(bEIGEN_implementation));

        // Perform upgrade
        vm.startBroadcast();
        EIGEN_ProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(bEIGEN_proxy))),
                address(bEIGEN_implementation)
            );
        EIGEN_ProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(EIGEN_proxy))),
                address(EIGEN_implementation)
            );
        vm.stopBroadcast();

        // Perform post-upgrade tests
        checkUpgradeCorrectness();
        simulateWrapAndUnwrap();
    }

    function checkUpgradeCorrectness() public {
        cheats.startPrank(address(proxyAdminOwner));
        require(EIGEN_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(EIGEN_proxy)))) == address(EIGEN_implementation),
            "implementation set incorrectly");
        require(EIGEN_proxy.bEIGEN() == bEIGEN_addressBefore,
            "bEIGEN address changed unexpectedly");
        require(EIGEN_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(bEIGEN_proxy)))) == address(bEIGEN_implementation),
            "implementation set incorrectly");
        require(bEIGEN_proxy.EIGEN() == EIGEN_addressBefore,
            "EIGEN address changed unexpectedly");
        cheats.stopPrank();
    }

    function simulateWrapAndUnwrap() public {
        uint256 amount = 1e18;
        cheats.prank(address(EIGEN_proxy));
        bEIGEN_proxy.transfer(address(this), amount);

        bEIGEN_proxy.approve(address(EIGEN_proxy), amount);
        uint256 bEIGEN_balanceStart = bEIGEN_proxy.balanceOf(address(this));
        uint256 EIGEN_balanceStart = EIGEN_proxy.balanceOf(address(this));
        EIGEN_proxy.wrap(amount);
        uint256 bEIGEN_balanceMiddle = bEIGEN_proxy.balanceOf(address(this));
        uint256 EIGEN_balanceMiddle = EIGEN_proxy.balanceOf(address(this));
        EIGEN_proxy.unwrap(amount);
        uint256 bEIGEN_balanceAfter = bEIGEN_proxy.balanceOf(address(this));
        uint256 EIGEN_balanceAfter = EIGEN_proxy.balanceOf(address(this));

        require(bEIGEN_balanceMiddle + amount == bEIGEN_balanceStart, "wrapping did not transfer out bEIGEN");
        require(EIGEN_balanceMiddle == EIGEN_balanceStart + amount, "wrapping did not transfer in EIGEN");

        require(bEIGEN_balanceAfter == bEIGEN_balanceStart, "unwrapping did not transfer in bEIGEN");
        require(EIGEN_balanceAfter == EIGEN_balanceStart, "unwrapping did not transfer out EIGEN");
    }
}
