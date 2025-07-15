// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ITaskMailbox} from "../../interfaces/ITaskMailbox.sol";
import {IKeyRegistrarTypes} from "../../interfaces/IKeyRegistrar.sol";

/**
 * @title TaskMailboxStorage
 * @author Layr Labs, Inc.
 * @notice Storage contract for the TaskMailbox contract.
 */
abstract contract TaskMailboxStorage is ITaskMailbox {
    /// @notice Equivalent to 100%, but in basis points.
    uint16 internal constant ONE_HUNDRED_IN_BIPS = 10_000;

    /// @notice Immutable BN254 certificate verifier
    address public immutable BN254_CERTIFICATE_VERIFIER;

    /// @notice Immutable ECDSA certificate verifier
    address public immutable ECDSA_CERTIFICATE_VERIFIER;

    /// @notice Global counter for tasks created across the TaskMailbox
    uint256 internal _globalTaskCount;

    /// @notice Mapping from task hash to task details
    mapping(bytes32 taskHash => Task task) internal _tasks;

    /// @notice Mapping to track registered executor operator sets by their keys
    mapping(bytes32 operatorSetKey => bool isRegistered) public isExecutorOperatorSetRegistered;

    /// @notice Mapping from executor operator set key to its task configuration
    mapping(bytes32 operatorSetKey => ExecutorOperatorSetTaskConfig config) internal _executorOperatorSetTaskConfigs;

    /// @notice The fee split percentage in basis points (0-10000)
    uint16 public feeSplit;

    /// @notice The address that receives the fee split
    address public feeSplitCollector;

    /**
     * @notice Constructor for TaskMailboxStorage
     * @param _bn254CertificateVerifier Address of the BN254 certificate verifier
     * @param _ecdsaCertificateVerifier Address of the ECDSA certificate verifier
     */
    constructor(address _bn254CertificateVerifier, address _ecdsaCertificateVerifier) {
        BN254_CERTIFICATE_VERIFIER = _bn254CertificateVerifier;
        ECDSA_CERTIFICATE_VERIFIER = _ecdsaCertificateVerifier;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[45] private __gap;
}
