// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IIndexRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

/**
 * @title Storage variables for the `IndexRegistry` contract.
 * @author Layr Labs, Inc.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract IndexRegistryStorage is Initializable, IIndexRegistry {

    /// @notice The value that indices of deregistered operators are set to
    uint32 public constant OPERATOR_DEREGISTERED_INDEX = type(uint32).max;

    /// @notice The RegistryCoordinator contract for this middleware
    IRegistryCoordinator public immutable registryCoordinator;

    /// @notice list of all operators ever registered, may include duplicates. used to avoid running an indexer on nodes
    bytes32[] public globalOperatorList;

    /// @notice mapping of operatorId => quorumNumber => index history of that operator
    mapping(bytes32 => mapping(uint8 => OperatorIndexUpdate[])) internal _operatorIdToIndexHistory;
    /// @notice mapping of quorumNumber => history of numbers of unique registered operators
    mapping(uint8 => OperatorIndexUpdate[]) internal _totalOperatorsHistory;

    constructor(
        IRegistryCoordinator _registryCoordinator
    ){
        registryCoordinator = _registryCoordinator;
        // disable initializers so that the implementation contract cannot be initialized
        _disableInitializers();
    }

    // storage gap for upgradeability
    uint256[47] private __GAP;
}
