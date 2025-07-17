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
    /// @notice Thrown when the certificate is too stale, per the max staleness period of the operatorSet
    error CertificateStale();
    /// @notice Thrown when the reference timestamp does not exist
    error ReferenceTimestampDoesNotExist();
    /// @notice Thrown when certificate verification fails
    error VerificationFailed();
    /// @notice Thrown when the global table root is disabled
    error RootDisabled();
}

/// @notice A base interface that verifies certificates for a given operatorSet
/// @notice This is a base interface that all curve certificate verifiers (eg. BN254, ECDSA) must implement
/// @dev A single `CertificateVerifier` is used for all operatorSets for a given key type
interface IBaseCertificateVerifier is
    IBaseCertificateVerifierEvents,
    IBaseCertificateVerifierErrors,
    ICrossChainRegistryTypes
{
    /**
     * @notice The address of the owner of the OperatorSet
     * @param operatorSet The operatorSet to get the owner of
     * @return The owner
     * @dev The owner of the OperatorSet is not used by this contract, but can be used by periphery contracts
     *      to gate access control for on-chain operations
     */
    function getOperatorSetOwner(
        OperatorSet memory operatorSet
    ) external view returns (address);

    /**
     * @notice The max staleness period of the operator table for a given operatorSet. This value is AVS-set and
     *         transported by the multichain protocol
     * @param operatorSet The operatorSet to get the max staleness period of
     * @return The max staleness period
     * @dev A staleness period of 0 allows for certificates to be verified against any timestamp in the past
     * @dev Staleness periods should not be greater than 0 and less than the update cadence of the `OperatorTables`, since
     *      certificates would be unable to be validated against
     */
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external view returns (uint32);

    /**
     * @notice The latest reference timestamp of the operator table for a given operatorSet. This value is
     *         updated each time an operator table is updated
     * @param operatorSet The operatorSet to get the latest reference timestamp of
     * @return The latest reference timestamp, 0 if the operatorSet has never been updated
     * @dev The latest reference timestamp is set when the operator table is updated
     */
    function latestReferenceTimestamp(
        OperatorSet memory operatorSet
    ) external view returns (uint32);

    /**
     * @notice Whether the operator table has been updated for a given reference timestamp
     * @param operatorSet The operatorSet to check
     * @param referenceTimestamp The reference timestamp to check
     * @return Whether the reference timestamp has been updated
     * @dev The reference timestamp is set when the operator table is updated
     */
    function isReferenceTimestampSet(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (bool);

    /**
     * @notice Get the total stake weights for all operators at a given reference timestamp
     * @param operatorSet The operator set to calculate stakes for
     * @param referenceTimestamp The reference timestamp
     * @return The sum of stake weights for each stake type, empty if the operatorSet has not been updated for the given reference timestamp
     * @dev For ECDSA, this function *reverts* if the reference timestamp is not set or the number of operators is 0
     * @dev For BN254, this function returns empty array if the reference timestamp is not set or the number of operators is 0
     */
    function getTotalStakeWeights(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint256[] memory);

    /**
     * @notice Get the number of operators at a given reference timestamp
     * @param operatorSet The operator set to get the number of operators for
     * @param referenceTimestamp The reference timestamp
     * @return The number of operators
     * @dev Returns 0 if the reference timestamp is not set or the number of operators is 0
     */
    function getOperatorCount(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp
    ) external view returns (uint256);
}
