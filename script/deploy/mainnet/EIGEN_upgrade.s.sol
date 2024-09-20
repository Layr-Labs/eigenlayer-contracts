// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

import "../../../src/contracts/token/BackingEigen.sol";
import "../../../src/contracts/token/Eigen.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/mainnet/EIGEN_upgrade.s.sol:EIGEN_upgrade -vvvv --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast
contract EIGEN_upgrade is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    BackingEigen public bEIGEN_proxy = BackingEigen(0x83E9115d334D248Ce39a6f36144aEaB5b3456e75);
    BackingEigen public bEIGEN_implementation;
    Eigen public EIGEN_proxy = Eigen(0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83);
    Eigen public EIGEN_implementation;
    ProxyAdmin public EIGEN_ProxyAdmin = ProxyAdmin(0xB8915E195121f2B5D989Ec5727fd47a5259F1CEC);
    TimelockController public EIGEN_TimelockController = TimelockController(payable(0x2520C6b2C1FBE1813AB5c7c1018CDa39529e9FF2));
    address public EIGEN_TimelockAdmin = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;

    IERC20 public bEIGEN_addressBefore;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId == 1) {
            // rpcUrl = "RPC_MAINNET";
        } else {
            revert("Chain not supported");
        }

        bEIGEN_addressBefore = EIGEN_proxy.bEIGEN();

        require(bEIGEN_addressBefore == IERC20(0x83E9115d334D248Ce39a6f36144aEaB5b3456e75),
            "something horribly wrong");

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contract
        EIGEN_implementation = new Eigen({
            _bEIGEN: bEIGEN_addressBefore
        });

        vm.stopBroadcast();

        emit log_named_address("EIGEN_implementation", address(EIGEN_implementation));

        // Perform post-upgrade tests
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
        simulateWrapAndUnwrap();
    }

    function simulatePerformingUpgrade() public {
        // Upgrade beacon
        uint256 delay = EIGEN_TimelockController.getMinDelay();
        bytes memory data = abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                TransparentUpgradeableProxy(payable(address(EIGEN_proxy))),
                EIGEN_implementation
        );
        emit log_named_bytes("data", data);

        vm.startPrank(EIGEN_TimelockAdmin);
        EIGEN_TimelockController.schedule({
            target: address(EIGEN_ProxyAdmin),
            value: 0,
            data: data,
            predecessor: bytes32(0),
            salt: bytes32(0),
            delay: delay
        });

        vm.warp(block.timestamp + delay);
        EIGEN_TimelockController.execute({
            target: address(EIGEN_ProxyAdmin),
            value: 0,
            payload: data,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        cheats.stopPrank();
    }

    function checkUpgradeCorrectness() public {
        vm.prank(address(EIGEN_TimelockController));
        require(EIGEN_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(EIGEN_proxy)))) == address(EIGEN_implementation),
            "implementation set incorrectly");
        require(EIGEN_proxy.bEIGEN() == bEIGEN_addressBefore,
            "bEIGEN address changed unexpectedly");
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