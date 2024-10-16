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
// forge script script/deploy/holesky/bEIGEN_and_EIGEN_upgrade.s.sol:bEIGEN_and_EIGEN_upgrade -vvv --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast
contract bEIGEN_and_EIGEN_upgrade is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    BackingEigen public bEIGEN_proxy = BackingEigen(0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27);
    BackingEigen public bEIGEN_implementation;
    Eigen public EIGEN_proxy = Eigen(0x3B78576F7D6837500bA3De27A60c7f594934027E);
    Eigen public EIGEN_implementation;
    ProxyAdmin public token_ProxyAdmin = ProxyAdmin(0x67482666771e82C9a73BB9e9A22B2B597f448BBf);
    address public opsMultisig = 0xfaEF7338b7490b9E272d80A1a39f4657cAf2b97d;

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
        EIGEN_addressBefore = bEIGEN_proxy.EIGEN();

        require(bEIGEN_addressBefore == IERC20(0x275cCf9Be51f4a6C94aBa6114cdf2a4c45B9cb27),
            "something horribly wrong");
        require(EIGEN_addressBefore == IERC20(0x3B78576F7D6837500bA3De27A60c7f594934027E),
            "something horribly wrong");

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implementation contracts
        EIGEN_implementation = new Eigen({
            _bEIGEN: bEIGEN_addressBefore
        });
        bEIGEN_implementation = new BackingEigen({
            _EIGEN: EIGEN_addressBefore
        });
        vm.stopBroadcast();

        emit log_named_address("EIGEN_implementation", address(EIGEN_implementation));
        emit log_named_address("bEIGEN_implementation", address(bEIGEN_implementation));

        // Perform post-upgrade tests
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
        simulateWrapAndUnwrap();
    }

    function simulatePerformingUpgrade() public {
        cheats.startPrank(opsMultisig);
        // Upgrade contracts
        token_ProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(EIGEN_proxy))), address(EIGEN_implementation));
        token_ProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(bEIGEN_proxy))), address(bEIGEN_implementation));
        cheats.stopPrank();
    }

    function checkUpgradeCorrectness() public {
        vm.startPrank(opsMultisig);
        require(token_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(EIGEN_proxy)))) == address(EIGEN_implementation),
            "implementation set incorrectly");
        require(EIGEN_proxy.bEIGEN() == bEIGEN_addressBefore,
            "bEIGEN address changed unexpectedly");
        require(token_ProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(bEIGEN_proxy)))) == address(bEIGEN_implementation),
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