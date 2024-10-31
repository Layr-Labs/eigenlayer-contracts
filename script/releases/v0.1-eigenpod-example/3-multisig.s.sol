// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

import "src/contracts/pods/EigenPodManager.sol";

import "./2-multisig.s.sol";

contract ExecuteEigenPodAndManager is QueueEigenPodAndManager {
    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for *;

    MultisigCall[] private _multisigCalls;

    function _execute() internal override returns (MultisigCall[] memory) {
        MultisigCall[] memory _executorCalls = _queue();

        // steals logic from queue() to perform execute()
        // likely the first step of any _execute() after a _queue()
        bytes memory executorCalldata = _executorCalls.makeExecutorCalldata(_multiSendCallOnly(), _timelock());

        // execute queued transaction upgrading eigenPodManager and eigenPod
        _multisigCalls.append({
            to: _timelock(),
            value: 0,
            data: abi.encodeWithSelector(ITimelock.executeTransaction.selector, executorCalldata)
        });

        // after queued transaction, renounce ownership from eigenPodManager
        _multisigCalls.append({
            to: _eigenPodManagerProxy(),
            value: 0,
            data: abi.encodeWithSelector(EigenPodManager(_eigenPodManagerProxy()).renounceOwnership.selector)
        });

        return _multisigCalls;
    }

    function zeusTest() public override {
        // Test function implementation
    }
}
