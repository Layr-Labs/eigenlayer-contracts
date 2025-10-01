// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

abstract contract SplitContractMixin {
    /// @notice The address of the second half of the contract.
    address internal immutable secondHalf;

    constructor(
        address _secondHalf
    ) {
        secondHalf = _secondHalf;
    }

    /**
     * @dev Delegates the current call to `implementation`.
     *
     * This function does not return to its internal call site, it will return directly to the external caller.
     *
     * Copied from OpenZeppelin Contracts v4.9.0 (proxy/Proxy.sol).
     */
    function _delegate(
        address implementation
    ) internal virtual {
        assembly {
            // Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code. We overwrite the
            // Solidity scratch pad at memory position 0.
            calldatacopy(0, 0, calldatasize())

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(gas(), implementation, 0, calldatasize(), 0, 0)

            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // delegatecall returns 0 on error.
            case 0 { revert(0, returndatasize()) }
            default { return(0, returndatasize()) }
        }
    }

    /**
     * @dev Fallback function that delegates calls to the address returned by `_implementation()`. Will run if no other
     * function in the contract matches the call data.
     * TODO: Explore if we want to static-delegatecall to ensure no state mutations are possible.
     */
    fallback() external virtual {
        _delegate(secondHalf);
    }
}
