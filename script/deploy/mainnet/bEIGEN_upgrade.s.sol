// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

import "../../../src/contracts/token/BackingEigen.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/mainnet/bEIGEN_upgrade.s.sol:bEIGEN_upgrade -vvvv --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast
contract bEIGEN_upgrade is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    BackingEigen public bEIGEN_proxy = BackingEigen(0x83E9115d334D248Ce39a6f36144aEaB5b3456e75);
    BackingEigen public bEIGEN_implementation;
    ProxyAdmin public bEIGEN_ProxyAdmin = ProxyAdmin(0x3f5Ab2D4418d38568705bFd6672630fCC3435CC9);
    TimelockController public bEIGEN_TimelockController = TimelockController(payable(0xd6EC41E453C5E7dA5494f4d51A053Ab571712E6f));
    address public bEIGEN_TimelockAdmin = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;

    // // RPC url to fork from for pre-upgrade state change tests
    // string public rpcUrl;

    IERC20 public EIGEN_addressBefore;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId == 1) {
            // rpcUrl = "RPC_MAINNET";
        } else {
            revert("Chain not supported");
        }

        EIGEN_addressBefore = bEIGEN_proxy.EIGEN();

        require(EIGEN_addressBefore == IERC20(0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83),
            "something horribly wrong");

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contract
        bEIGEN_implementation = new BackingEigen({
            _EIGEN: EIGEN_addressBefore
        });

        vm.stopBroadcast();

        emit log_named_address("bEIGEN_implementation", address(bEIGEN_implementation));

        // Perform post-upgrade tests
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
    }

    function simulatePerformingUpgrade() public {
        // Upgrade beacon
        uint256 delay = bEIGEN_TimelockController.getMinDelay();
        bytes memory data = abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                TransparentUpgradeableProxy(payable(address(bEIGEN_proxy))),
                bEIGEN_implementation
        );
        emit log_named_bytes("data", data);

        vm.startPrank(bEIGEN_TimelockAdmin);
        bEIGEN_TimelockController.schedule({
            target: address(bEIGEN_ProxyAdmin),
            value: 0,
            data: data,
            predecessor: bytes32(0),
            salt: bytes32(0),
            delay: delay
        });

        vm.warp(block.timestamp + delay);
        bEIGEN_TimelockController.execute({
            target: address(bEIGEN_ProxyAdmin),
            value: 0,
            payload: data,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        cheats.stopPrank();
    }

    function checkUpgradeCorrectness() public {
        vm.prank(address(bEIGEN_TimelockController));
        require(bEIGEN_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(bEIGEN_proxy)))) == address(bEIGEN_implementation),
            "implementation set incorrectly");
        require(bEIGEN_proxy.EIGEN() == EIGEN_addressBefore,
            "EIGEN address changed unexpectedly");
    }
}