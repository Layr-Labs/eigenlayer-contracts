// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";
import "../../utils/TimelockEncoding.sol";

// forge script script/admin/mainnet/Mainnet_Unpause_Deposits.s.sol:Mainnet_Unpause_Deposits --fork-url $RPC_MAINNET -vvvv
contract Mainnet_Unpause_Deposits is ExistingDeploymentParser, TimelockEncoding {
    Vm cheats = Vm(HEVM_ADDRESS);

    // Tues Apr 16 2024 12:00:00 GMT-0700 (Pacific Daylight Time)
    uint256 timelockEta = 1713250800;

    function run() external virtual {
        _parseDeployedContracts("script/output/mainnet/M1_deployment_mainnet_2023_6_9.json");

        bytes memory final_calldata_to_executor_multisig = encodeForExecutor({
            // call to executor will be from the timelock
            from: timelock,
            // performing single pause operation
            to: address(strategyManager),
            // value to send in tx
            value: 0,
            // calldata for the operation
            data: abi.encodeWithSelector(Pausable.unpause.selector, 0),
            // operation type (for performing single operation)
            operation: ISafe.Operation.Call
        });

        (bytes memory calldata_to_timelock_queuing_action, bytes memory calldata_to_timelock_executing_action) = encodeForTimelock({
            // address to be called from the timelock
            to: executorMultisig,
            // value to send in tx
            value: 0,
            // calldata for the operation
            data: final_calldata_to_executor_multisig,
            // time at which the tx will become executable
            timelockEta: timelockEta
        });

        bytes32 expectedTxHash = getTxHash({
            target: executorMultisig,
            _value: 0,
            _data: final_calldata_to_executor_multisig,
            eta: timelockEta
        });
        emit log_named_bytes32("expectedTxHash", expectedTxHash);

        cheats.prank(operationsMultisig);
        (bool success, ) = timelock.call(calldata_to_timelock_queuing_action);
        require(success, "call to timelock queuing action failed");

        require(ITimelock(timelock).queuedTransactions(expectedTxHash), "expectedTxHash not queued");

        // test performing the upgrade
        cheats.warp(timelockEta);
        cheats.prank(operationsMultisig);
        (success, ) = timelock.call(calldata_to_timelock_executing_action);
        require(success, "call to timelock executing action failed");

        // Check correctness after upgrade
        require(strategyManager.paused() == 0, "unpausing was not completed correctly");
    }

    function getTxHash(address target, uint256 _value, bytes memory _data, uint256 eta) public pure returns (bytes32) {
        // empty bytes
        bytes memory signature;
        bytes32 txHash = keccak256(abi.encode(target, _value, signature, _data, eta));
        return txHash;        
    }
}