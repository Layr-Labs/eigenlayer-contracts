// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IStrategy.sol";

/**
 * @title Interface for an AVS viewer which should implement all of the functions required to display the AVS details on the EigenLayer UI.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IAVSStateViewer {
    /**
     * @notice Returns 
     * 1) an array of the quorums that are being served by an operator
     * 2) an array of strategies for each quorum in sequence in a 2 dimensional array
     * @param operator the address of the operator to get the quorum details for
     */
    function strategiesStakedForEachQuorum(address operator) external view returns (uint256[] memory, IStrategy[][] memory);
}