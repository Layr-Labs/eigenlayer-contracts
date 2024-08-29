// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";

abstract contract AVSDirectoryStorage is IAVSDirectory {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `Registration` struct used by the contract
    bytes32 public constant OPERATOR_AVS_REGISTRATION_TYPEHASH =
        keccak256("OperatorAVSRegistration(address operator,address avs,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetRegistration` struct used by the contract
    bytes32 public constant OPERATOR_SET_REGISTRATION_TYPEHASH =
        keccak256("OperatorSetRegistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @notice The EIP-712 typehash for the `OperatorSetMembership` struct used by the contract
    bytes32 public constant OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH =
        keccak256("OperatorSetForceDeregistration(address avs,uint32[] operatorSetIds,bytes32 salt,uint256 expiry)");

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice Mapping: avs => operator => OperatorAVSRegistrationStatus struct
    /// @dev This storage will be deprecated once M2 based deregistration is deprecated.
    mapping(address => mapping(address => OperatorAVSRegistrationStatus)) public avsOperatorStatus;

    /// @notice Mapping: operator => salt => Whether the salt has been used or not.
    mapping(address => mapping(bytes32 => bool)) public operatorSaltIsSpent;

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

    constructor(IDelegationManager _delegation) {
        delegation = _delegation;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
