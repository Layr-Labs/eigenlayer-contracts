// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/interfaces/IETHPOSDeposit.sol";

contract ETHPOSDepositMock is IETHPOSDeposit {
    function deposit(bytes calldata pubkey, bytes calldata withdrawal_credentials, bytes calldata signature, bytes32 deposit_data_root)
        external
        payable
    {}

    function get_deposit_root() external pure returns (bytes32) {
        bytes32 root;
        return root;
    }

    /// @notice Query the current deposit count.
    /// @return The deposit count encoded as a little endian 64-bit number.
    function get_deposit_count() external pure returns (bytes memory) {
        bytes memory root;
        return root;
    }
}
