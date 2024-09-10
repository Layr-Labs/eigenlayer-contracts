// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IDelegationManager.sol";
import {Snapshots} from "../libraries/Snapshots.sol";

abstract contract AllocationManagerStorage is IAllocationManager {
    using Snapshots for Snapshots.History;

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `MagnitudeAdjustments` struct used by the contract
    bytes32 public constant MAGNITUDE_ADJUSTMENT_TYPEHASH = keccak256(
        "MagnitudeAdjustments(address operator,MagnitudeAdjustment(address strategy, OperatorSet(address avs, uint32 operatorSetId)[], uint64[] magnitudeDiffs)[],bytes32 salt,uint256 expiry)"
    );

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice Mapping: operator => salt => Whether the salt has been used or not.
    mapping(address => mapping(bytes32 => bool)) public operatorSaltIsSpent;

    /// @notice Mapping: operator => strategy => snapshotted totalMagnitude
    /// Note that totalMagnitude is monotonically decreasing and only gets updated upon slashing
    mapping(address => mapping(IStrategy => Snapshots.History)) internal _totalMagnitudeUpdate;

    /// @notice Mapping: operator => strategy => operatorSet (encoded) => snapshotted magnitude
    mapping(address => mapping(IStrategy => mapping(bytes32 => Snapshots.History))) internal _magnitudeUpdate;

    /// @notice Mapping: operator => strategy => OperatorMagnitudeInfo to keep track of info regarding pending magnitude allocations.
    mapping(address => mapping(IStrategy => OperatorMagnitudeInfo)) public operatorMagnitudeInfo;

    /// @notice Mapping: operator => strategy => PendingFreeMagnitude[] to keep track of pending free magnitude from deallocations
    mapping(address => mapping(IStrategy => PendingFreeMagnitude[])) internal _pendingFreeMagnitude;

    /// @notice Mapping: operator => strategy => operatorSet (encoded) => list of queuedDeallocation indices
    mapping(address => mapping(IStrategy => mapping(bytes32 => uint256[]))) internal _queuedDeallocationIndices;

    constructor(IDelegationManager _delegation, IAVSDirectory _avsDirectory) {
        delegation = _delegation;
        avsDirectory = _avsDirectory;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
