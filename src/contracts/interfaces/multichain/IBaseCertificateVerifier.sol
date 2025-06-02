// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "./ICrossChainRegistry.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";

interface IBaseCertificateVerifierEvents {
    /// @notice Emitted when the owner of an operatorSet is updated
    event OperatorSetOwnerUpdated(OperatorSet operatorSet, address owner);

    /// @notice Emitted when the max staleness period of an operatorSet is updated
    event MaxStalenessPeriodUpdated(OperatorSet operatorSet, uint32 maxStalenessPeriod);
}

interface IBaseCertificateVerifierErrors {
    /// @notice Thrown when the table updater is not caller
    error OnlyTableUpdater();
    /// @notice Thrown when the table update is stale
    error TableUpdateStale();
    /// @notice Thrown when array lengths mismatch
    error ArrayLengthMismatch();
    /// @notice Thrown when the certificate is too stale
    error CertificateStale();
    /// @notice Thrown when the reference timestamp does not exist
    error ReferenceTimestampDoesNotExist();
    /// @notice Thrown when certificate verification fails
    error VerificationFailed();
}

/// @notice A base interface that verifies certificates for a given operatorSet
/// @notice This is a base interface that all curve certificate verifiers (eg. BN254, ECDSA) must implement
/// @dev A single `CertificateVerifier` can be used for multiple operatorSets, but a single key type
interface IBaseCertificateVerifier is
    IBaseCertificateVerifierEvents,
    IBaseCertificateVerifierErrors,
    ICrossChainRegistryTypes
{
    /// @notice the address of the owner of the OperatorSet
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external returns (address);

    /// @return the maximum amount of seconds that a operator table can be in the past for a given operatorSet
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external returns (uint32);

    /// @notice The latest reference timestamp of the operator table for a given operatorSet
    function latestReferenceTimestamp(
        OperatorSet memory operatorSet
    ) external returns (uint32);
}
