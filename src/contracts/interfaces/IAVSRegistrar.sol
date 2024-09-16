// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

interface IAVSRegistrar {
    /**
     * @notice Called by the AllocationManager when an operator wants to register
     * for one or more operator sets. This method should revert if registration
     * is unsuccessful.
     * @param operator the registering operator
     * @param operatorSetIds the list of operator set ids being registered for
     * @param data arbitrary data the operator can provide as part of registration
     */
    function registerOperator(address operator, uint32[] calldata operatorSetIds, bytes calldata data) external;

    /**
     * @notice Called by the AllocationManager when an operator is deregistered from
     * one or more operator sets. If this method reverts, it is ignored.
     * @param operator the deregistering operator
     * @param operatorSetIds the list of operator set ids being deregistered from
     */
    function deregisterOperator(address operator, uint32[] calldata operatorSetIds) external;
}
