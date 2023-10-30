// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./IStakeRegistryStub.sol";

// @notice Stub interface to avoid circular-ish inheritance, where core contracts rely on middleware interfaces
interface IStakeRegistryStub {
    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the addresses of the operators whose stake information is getting updated
     */
    function updateStakes(address[] memory operators) external;
}
