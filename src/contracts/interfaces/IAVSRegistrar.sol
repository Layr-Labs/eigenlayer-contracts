// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

interface IAVSRegistrar {
    /**
     * @notice Called by the AllocationManager when an operator wants to register
     * for one or more operator sets. This method should revert if registration
     * is unsuccessful.
     * @param operator the registering operator
     * @param avsIdentifier the AVS the operator is registering for. This should be the same as IAVSRegistrar.avs()
     * @param operatorSetIds the list of operator set ids being registered for
     * @param data arbitrary data the operator can provide as part of registration
     */
    function registerOperator(
        address operator,
        address avsIdentifier,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external;

    /**
     * @notice Called by the AllocationManager when an operator is deregistered from
     * one or more operator sets. If this method reverts, it is ignored.
     * @param operator the deregistering operator
     * @param avsIdentifier the AVS the operator is deregistering from. This should be the same as IAVSRegistrar.avs()
     * @param operatorSetIds the list of operator set ids being deregistered from
     */
    function deregisterOperator(address operator, address avsIdentifier, uint32[] calldata operatorSetIds) external;

    /**
     * @notice Returns the account identifier address of the AVS. This should be the
     * the same address as the AVS's UAM account identifier address.
     * @return address of the AVS
     */
    function avs() external view returns (address);
}
