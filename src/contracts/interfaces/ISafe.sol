// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


interface ISafe {
    enum  Operation {
        Call,
        DelegateCall
    }

    function setup(
        address[] calldata _owners,
        uint256 _threshold,
        address to,
        bytes calldata data,
        address fallbackHandler,
        address paymentToken,
        uint256 payment,
        address payable paymentReceiver
    ) external;

    function execTransaction(
        address to,
        uint256 value,
        bytes calldata data,
        Operation operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address payable refundReceiver,
        bytes calldata signatures
    ) external payable returns (bytes memory);

    function checkSignatures(bytes32 dataHash, bytes calldata signatures) external view;
    
    function approveHash(bytes32 hashToApprove) external;
    
    



}