// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IOwnershipManager.sol";

abstract contract OwnershipManagerStorage is IOwnershipManager {
    /// @notice The admin of the identifier
    mapping(bytes20 => address) public identifierToAdmin;

    /// @notice Mapping from (identifier, target, selector) tuple to its pendingAgent
    /// @notice Mapping(bytes20 identifier => address target => bytes4 functionSelector => address pendingAgent)
    mapping(bytes20 => mapping(address => mapping(bytes4 => address))) public identifierToPendingAgent;

    /// @notice Mapping from agent to identifier that it can act on behalf of for a given target and selector
    /// @notice Mapping(address agent => address target => bytes4 functionSelector => bytes20 identifier)
    mapping(address => mapping(address => mapping(bytes4 => bytes20))) public agentToIdentifier;

    /// @notice Mapping from identifier to agent that can act on its behalf for a given target and selector
    /// @notice Mapping(bytes20 identifier => address target => bytes4 functionSelector => address agent)
    mapping(bytes20 => mapping(address => mapping(bytes4 => address))) public identifierToAgent;
}