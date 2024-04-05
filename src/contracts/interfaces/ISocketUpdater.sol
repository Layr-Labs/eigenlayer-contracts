// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Interface for an `ISocketUpdater` where operators can update their sockets.
 * @author Layr Labs, Inc.
 */
interface ISocketUpdater {
    // EVENTS

    event OperatorSocketUpdate(bytes32 indexed operatorId, string socket);

    // FUNCTIONS

    /**
     * @notice Updates the socket of the msg.sender given they are a registered operator
     * @param socket is the new socket of the operator
     */
    function updateSocket(string memory socket) external;
}

