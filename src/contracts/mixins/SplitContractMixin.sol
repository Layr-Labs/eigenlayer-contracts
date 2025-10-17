// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

abstract contract SplitContractMixin {
    /// @notice The address of the second half of the contract.
    address internal immutable viewImplementation;

    constructor(
        address _viewImplementation
    ) {
        viewImplementation = _viewImplementation;
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
     * @dev Performs a delegate call to `implementation` in the context of a view function.
     *
     * This function typecasts the non-view `_delegate` function to a view function in order to
     * allow its invocation from a view context. This is required because the EVM itself does not
     * enforce view/pure mutability, and using inline assembly, it is possible to cast a function
     * pointer to a view (read-only) signature. This pattern is sometimes used for readonly proxies,
     * but it should be used cautiously since any state-modifying logic in the underlying delegate
     * violates the spirit of a view call.
     *
     * @param implementation The address to which the call should be delegated.
     */
    function _delegateView(
        address implementation
    ) internal view virtual {
        function(address) fn = _delegate;
        function(address) view fnView;
        /// @solidity memory-safe-assembly
        assembly {
            fnView := fn
        }
        fnView(implementation);
    }
}
