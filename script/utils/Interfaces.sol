// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

interface IProxyAdmin {
    function upgrade(address proxy, address implementation) external;
    function getProxyImplementation(address proxy) external returns (address);
}

interface IUpgradeableBeacon {
    function upgradeTo(address newImplementation) external;
    function implementation() external returns (address);
}

interface ISafe {
    /// @dev Allows to execute a Safe transaction confirmed by required number of owners and then pays the account that submitted the transaction.
    ///      Note: The fees are always transferred, even if the user transaction fails.
    /// @param to Destination address of Safe transaction.
    /// @param value Ether value of Safe transaction.
    /// @param data Data payload of Safe transaction.
    /// @param operation Operation type of Safe transaction.
    /// @param safeTxGas Gas that should be used for the Safe transaction.
    /// @param baseGas Gas costs that are independent of the transaction execution(e.g. base transaction fee, signature check, payment of the refund)
    /// @param gasPrice Gas price that should be used for the payment calculation.
    /// @param gasToken Token address (or 0 if ETH) that is used for the payment.
    /// @param refundReceiver Address of receiver of gas payment (or 0 if tx.origin).
    /// @param signatures Packed signature data ({bytes32 r}{bytes32 s}{uint8 v})
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
}

interface ITimelock {
    function queuedTransactions(bytes32) external view returns (bool);

    function queueTransaction(
        address target, 
        uint value, 
        string memory signature, 
        bytes memory data, 
        uint eta
    ) external returns (bytes32);

    function executeTransaction(
        address target, 
        uint value, 
        string memory signature, 
        bytes memory data, 
        uint eta
    ) external payable returns (bytes memory);

    function cancelTransaction(
        address target, 
        uint value, 
        string memory signature, 
        bytes memory data, 
        uint eta
    ) external;
}

interface IMultiSend {
    function multiSend(bytes memory transactions) external payable;
}