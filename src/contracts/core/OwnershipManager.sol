// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./OwnershipManagerStorage.sol";

contract OwnershipManager is OwnershipManagerStorage {
    /// @inheritdoc IOwnershipManager
    function setAdmin(bytes20 identifier, address newAdmin) external {
        // Sender must be the admin of the identifier OR the identifier itself if the admin is not set
        address currentAdmin = identifierToAdmin[identifier];
        if (currentAdmin != address(0)) {
            require(msg.sender == currentAdmin, NotAdmin());
        } else {
            require(msg.sender == address(identifier), SenderMustBeIdentifierIfAdminNotSet());
        }

        identifierToAdmin[identifier] = newAdmin;

        emit AdminSet(identifier, currentAdmin, newAdmin);
    }

    /// @inheritdoc IOwnershipManager
    function setPendingAgent(bytes20 identifier, address target, bytes4 selector, address pendingAgent) external {
        require(msg.sender == identifierToAdmin[identifier], NotAdmin());

        identifierToPendingAgent[identifier][target][selector] = pendingAgent;

        emit PendingAgentSet(identifier, target, selector, pendingAgent);
    }

    /// @inheritdoc IOwnershipManager
    function acceptAgency(bytes20 identifier, address target, bytes4 selector) external {
        require(msg.sender == identifierToPendingAgent[identifier][target][selector], NotPendingAdmin());
        require(agentToIdentifier[msg.sender][target][selector] == bytes20(0), AgentAlreadyOccupied());

        // Set agentToIdentifier and identifierToAgent mappings
        agentToIdentifier[msg.sender][target][selector] = identifier;
        identifierToAgent[identifier][target][selector] = msg.sender;

        // Reset pending agent
        identifierToPendingAgent[identifier][target][selector] = address(0);

        emit AgencyAccepted(identifier, target, selector, msg.sender);
    }

    /// @inheritdoc IOwnershipManager
    function revokeAgent(bytes20 identifier, address target, bytes4 selector, address agent) external {
        // Only the agent of the (identifier, target, selector) tuple or the admin can revoke the agent
        if (msg.sender == agent) {
            require(msg.sender == identifierToAgent[identifier][target][selector], NotAgent());
        } else {
            require(msg.sender == identifierToAdmin[identifier], NotAdmin());
        }

        // Reset agentToIdentifier and identifierToAgent mappings
        agentToIdentifier[agent][target][selector] = bytes20(0);
        identifierToAgent[identifier][target][selector] = address(0);

        emit AgencyRevoked(identifier, target, selector, agent, msg.sender);
    }

    /// @inheritdoc IOwnershipManager
    function getIdentifier(address agent, address target, bytes4 selector) external view returns (bytes20) {
        bytes20 identifier = agentToIdentifier[agent][target][selector];
        return identifier == bytes20(0) ? bytes20(agent) : identifier;
    }
}