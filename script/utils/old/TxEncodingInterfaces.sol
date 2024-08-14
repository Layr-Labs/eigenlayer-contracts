// SPDX-License-Identifier: LGPL AND BSD 3-Clause
pragma solidity >=0.5.0;

// based on https://github.com/safe-global/safe-smart-account/blob/v1.3.0/contracts/GnosisSafe.sol
interface ISafe {
    function execTransaction(
        address to,
        uint256 value,
        bytes calldata data,
        uint8 operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address payable refundReceiver,
        bytes memory signatures
    ) external;

    enum Operation {Call, DelegateCall}
}

// based on https://github.com/compound-finance/compound-protocol/blob/master/contracts/Timelock.sol
interface ITimelock {
    function queueTransaction(address target, uint value, string memory signature, bytes memory data, uint eta) external returns (bytes32);
    function executeTransaction(address target, uint value, string memory signature, bytes memory data, uint eta) external payable returns (bytes memory);
    function queuedTransactions(bytes32) external view returns (bool);
}

