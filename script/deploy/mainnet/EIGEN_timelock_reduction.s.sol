// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/governance/TimelockController.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/mainnet/EIGEN_timelock_reduction.s.sol:EIGEN_timelock_reduction -vvvv --rpc-url $RPC_URL
contract EIGEN_timelock_reduction is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    TimelockController public EIGEN_TimelockController = TimelockController(payable(0x2520C6b2C1FBE1813AB5c7c1018CDa39529e9FF2));
    address public EIGEN_TimelockAdmin = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;

    uint256 public newDelay = 0;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId == 1) {
            // rpcUrl = "RPC_MAINNET";
        } else {
            revert("Chain not supported");
        }

        uint256 minDelayBefore = EIGEN_TimelockController.getMinDelay();

        require(minDelayBefore == 10 days,
            "something horribly wrong");

        bytes memory proposalData = abi.encodeWithSelector(
                TimelockController.updateDelay.selector,
                newDelay
        );
        emit log_named_bytes("proposalData", proposalData);

        // propose change to zero delay
        vm.startPrank(EIGEN_TimelockAdmin);
        EIGEN_TimelockController.schedule({
            target: address(EIGEN_TimelockController),
            value: 0,
            data: proposalData,
            predecessor: bytes32(0),
            salt: bytes32(0),
            delay: minDelayBefore
        });

        // fast-forward to after current delay and execute
        vm.warp(block.timestamp + minDelayBefore);
        EIGEN_TimelockController.execute({
            target: address(EIGEN_TimelockController),
            value: 0,
            payload: proposalData,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        cheats.stopPrank();

        uint256 minDelayAfter = EIGEN_TimelockController.getMinDelay();

        require(minDelayAfter == 0,
            "min delay not set to zero");
    }
}