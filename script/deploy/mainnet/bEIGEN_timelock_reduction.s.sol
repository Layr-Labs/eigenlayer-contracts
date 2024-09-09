// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/governance/TimelockController.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/mainnet/bEIGEN_timelock_reduction.s.sol:bEIGEN_timelock_reduction -vvvv --rpc-url $RPC_URL
contract bEIGEN_timelock_reduction is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    TimelockController public bEIGEN_TimelockController = TimelockController(payable(0xd6EC41E453C5E7dA5494f4d51A053Ab571712E6f));
    address public bEIGEN_TimelockAdmin = 0xbb00DDa2832850a43840A3A86515E3Fe226865F2;

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

        uint256 minDelayBefore = bEIGEN_TimelockController.getMinDelay();

        require(minDelayBefore == 24 days,
            "something horribly wrong");

        bytes memory proposalData = abi.encodeWithSelector(
                TimelockController.updateDelay.selector,
                newDelay
        );
        emit log_named_bytes("proposalData", proposalData);

        // propose change to zero delay
        vm.startPrank(bEIGEN_TimelockAdmin);
        bEIGEN_TimelockController.schedule({
            target: address(bEIGEN_TimelockController),
            value: 0,
            data: proposalData,
            predecessor: bytes32(0),
            salt: bytes32(0),
            delay: minDelayBefore
        });

        // fast-forward to after current delay and execute
        vm.warp(block.timestamp + minDelayBefore);
        bEIGEN_TimelockController.execute({
            target: address(bEIGEN_TimelockController),
            value: 0,
            payload: proposalData,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        cheats.stopPrank();

        uint256 minDelayAfter = bEIGEN_TimelockController.getMinDelay();

        require(minDelayAfter == 0,
            "min delay not set to zero");
    }
}