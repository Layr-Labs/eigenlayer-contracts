// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";

abstract contract AVSDirectoryStorage is IAVSDirectory {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

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

    /// @notice Delay before deallocations are completable and can be added back into freeMagnitude
    /// In this window, deallocations still remain slashable by the operatorSet they were allocated to.
    uint32 public immutable DEALLOCATION_DELAY;

    // Mutatables

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal __deprecated_DOMAIN_SEPARATOR;

    /// @notice Mapping: avs => operator => OperatorAVSRegistrationStatus struct
    /// @dev This storage will be deprecated once M2 based deregistration is deprecated.
    mapping(address => mapping(address => OperatorAVSRegistrationStatus)) public avsOperatorStatus;

    /// @notice Mapping: admin => salt => Whether the salt has been used or not.
    mapping(address => mapping(bytes32 => bool)) public adminSaltIsSpent;

    /// @notice Mapping: avs => Whether it is a an operator set AVS or not.
    mapping(address => bool) public isOperatorSetAVS;

    /// @notice Mapping: avs => operatorSetId => Whether or not an operator set is valid.
    mapping(address => mapping(uint32 => bool)) public isOperatorSet;

    /// @notice Mapping: operator => List of operator sets that operator is registered to.
    /// @dev Each item is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(address => EnumerableSet.Bytes32Set) internal _operatorSetsMemberOf;

    /// @notice Mapping: operatorSet => List of operators that are registered to the operatorSet
    /// @dev Each key is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(bytes32 => EnumerableSet.AddressSet) internal _operatorSetMembers;

    /// @notice Mapping: operatorSet => List of strategies that the operatorSet contains
    /// @dev Each key is formatted as such: bytes32(abi.encodePacked(avs, uint96(operatorSetId)))
    mapping(bytes32 => EnumerableSet.AddressSet) internal _operatorSetStrategies;

    /// @notice Mapping: operator => avs => operatorSetId => operator registration status
    mapping(address => mapping(address => mapping(uint32 => OperatorSetRegistrationStatus))) public operatorSetStatus;

    // Construction

    constructor(IDelegationManager _delegation, uint32 _DEALLOCATION_DELAY) {
        delegation = _delegation;
        DEALLOCATION_DELAY = _DEALLOCATION_DELAY;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[41] private __gap;
}
