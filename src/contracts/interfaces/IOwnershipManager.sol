// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IOwnershipManager {
    /// Events
    event AdminSet(bytes20 indexed identifier, address oldAdmin, address newAdmin);
    event PendingAgentSet(bytes20 indexed identifier, address target, bytes4 selector, address pendingAgent);
    event AgencyAccepted(bytes20 indexed identifier, address target, bytes4 selector, address agent);
    event AgencyRevoked(bytes20 indexed identifier, address target, bytes4 selector, address agent, address revoker);

    /// Errors
    error NotAdmin();
    error NotPendingAdmin();
    error AgentAlreadyOccupied();
    error SenderMustBeIdentifierIfAdminNotSet();
    error NotAgent();


    /**
     * @notice Sets the identifier for the agent
     * @param identifier The identifier to set the agent for
     * @param newAdmin The new admin to set for the identifier
     * @dev msg.sender must be the admin of the identifier OR the identifier itself if the admin is not set
     */
    function setAdmin(bytes20 identifier, address newAdmin) external;

    /**
     * @notice Sets the agent of the identifier
     * @param identifier The identifier to set the agent for
     * @param target The target to set the agent for
     * @param selector The selector to set the agent for
     * @param pendingAgent The new agent to set for the (identifier, target, selector) tuple
     * @dev Only callable by the admin 
     * @dev An agent can be pending for multiple (identifier, target, selector) tuples, but only accept agency and be the agent for one tuple
     */
	function setPendingAgent(bytes20 identifier, address target, bytes4 selector, address pendingAgent) external;

    /**
     * @notice Accepts the agency status on a (identifier, target, selector) tuple
     * @param identifier The identifier to become an agent for
     * @param target The target to become an agent for
     * @param selector The selector to become an agent for
     * @dev msg.sender is the pending agent to accept the identifier
     * @dev Reverts if the agent can already act on behalf of an (identifier, target, selector) tuple
     * @dev Reverts if the agent is not in the pending state for the (identifier, target, selector) tuple
     */
    function acceptAgency(bytes20 identifier, address target, bytes4 selector) external;

    /**
     * @notice Revokes the agent for a given (identifier, target, selector) tuple
     * @param identifier The identifier to revoke the agent for
     * @param target The target to revoke the agent for
     * @param selector The selector to revoke the agent for
     * @param agent The agent to revoke
     * @dev Only callable by the admin or the current agent of the (identifier, target, selector) tuple
     */
    function revokeAgent(bytes20 identifier, address target, bytes4 selector, address agent) external;

    /**
     * @notice Returns the identifier for the given agent
     * @param agent The agent to get the identifier for
     * @param target The target contract to check agency for
     * @param selector The selector on the garget contract to check agency for
     * @dev If the identifier is bytes20(0), that means the (agent, target, selector) tuple does
     *      not have an identifier and we return the agent address itself
     */
    function getIdentifier(address agent, address target, bytes4 selector) external view returns (bytes20);
}