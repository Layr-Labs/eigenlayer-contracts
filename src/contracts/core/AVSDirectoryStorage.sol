// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";

abstract contract AVSDirectoryStorage is IAVSDirectory {
    // Constants

    /// @notice The EIP-712 typehash for the `Registration` struct used by the contract
    bytes32 public constant OPERATOR_AVS_REGISTRATION_TYPEHASH =
        keccak256("OperatorAVSRegistration(address operator,address avs,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetRegistration` struct used by the contract
    bytes32 public constant OPERATOR_SET_REGISTRATION_TYPEHASH =
        keccak256("OperatorSetRegistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetMembership` struct used by the contract
    bytes32 public constant OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH =
        keccak256("OperatorSetForceDeregistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SET_REGISTRATION_AND_DEREGISTRATION = 1;

    // Immutables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    // Mutatables

    /// @dev Do not remove, deprecated storage.
    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /// @notice Returns the registration status of each `operator` for a given `avs`.
    /// @dev This storage will be deprecated once M2-based deregistration is removed.
    mapping(address avs => mapping(address operator => OperatorAVSRegistrationStatus)) public avsOperatorStatus;

    /// @notice Returns whether a `salt` has been used by a given `operator`.
    mapping(address operator => mapping(bytes32 salt => bool isSpent)) public operatorSaltIsSpent;

    // Construction

    constructor(
        IDelegationManager _delegation
    ) {
        delegation = _delegation;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[47] private __gap;
}
