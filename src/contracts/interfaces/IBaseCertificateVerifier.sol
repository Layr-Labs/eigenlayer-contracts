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
/// @dev The `CertificateVerifier` verifies certificates for a given operatorSet against an operator table. Operator tables are updated by an offchain service
interface IBaseCertificateVerifier is
    IBaseCertificateVerifierEvents,
    IBaseCertificateVerifierErrors,
    ICrossChainRegistryTypes
{
    /// @notice The following steps describe the certificate verification process, in order:
    /// 1. The AVS configures the following parameters in EigenLayer core:
    ///    a. AllocationManager.createOperatorSet: Creates an operatorSet
    ///    b. KeyRegistrar.configureOperatorSet: Configures the curve type of the operatorSet
    ///    c. CrossChainRegistry.makeGenerationReservation: Registers the operatorSet to be transported by the multichain protocol. This includes
    ///      the `owner`, `maxStalenessPeriod`, and `operatorTableCalculator` for the operatorSet. The output of the `OperatorTableCalculator`
    ///      are the operatorSet's stake weights (i.e. operator table) and is transported by the multichain protocol, along with the `maxStalenessPeriod` and `owner`
    /// 2. The multichain protocol calculates the operatorTable of an operatorSet. The time at which the table is calculated is the reference timesetamp. The protocol
    ///    will then call `updateOperatorTable` to update the operatorSet's operator table for a given referenceTimestamp
    /// 3. A task is created and certificate is generated, off-chain, by the AVS to validate the completion of a task.
    ///    The reference timestamp in the certificate is used to key into the operator table that was updated in step 2.
    /// 4. The certificate is verified, either normally, proportionally, or nominally.
    /// @dev The `referenceTimestamp` is used to key into the operatorSet's stake weights. It is NOT when the certificate was generated off-chain
    /// @dev The `maxStalenessPeriod` configured in step 1c denotes if a certificate is too stale with respect to the `referenceTimestamp`
    /// @dev Operator tables for ALL operatorSets with an active generation reservation are updated at a set cadence. See `crossChainRegistry.tableUpdateCadence` for the frequency of table updates
    /// @dev To ensure that tables do not become stale between table updates (i.e. a large operator has joined or been ejected), the multichain protocol updates tables for operatorSets when the following events are emitted:
    ///      - AllocationManager: `OperatorSlashed`
    ///      - AllocationManager: `OperatorAddedToOperatorSet`
    ///      - AllocationManager: `OperatorRemovedFromOperatorSet`
    /// @dev Certificates can be replayed across all destination chains
    /// @dev Race conditions should be handled by the AVS. The protocol makes no guarantees about how certificates should be verified (eg. preventing certificates against tables that are NOT the latest)
    ///      Some examples of race conditions include:
    ///      a. An in-flight certificate for a past reference timestamp and an operator table update for a newer reference timestamp. The AVS should decide whether it
    ///         wants to only confirm tasks against the *latest* certificate
    ///      b. An in-flight certificate against a stake table with a majority-stake operator that has been slashed or removed from the operatorSet

    /**
     * @notice The address of the owner of the operatorSet
     * @param operatorSet The operatorSet to get the owner of
     * @return The owner
     * @dev The owner of the OperatorSet is not used by this contract, but can be used by periphery contracts
     *      to gate access control for on-chain operations
     * @dev This value is set by the AVS in the `CrossChainRegistry` and transported by the multichain protocol when the operator table is updated
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
     * @dev Staleness periods cannot be greater than 0 and less than the update cadence of the `OperatorTables`, since
     *      certificates would be unable to be validated against. This value is set and bounds enforced in the `CrossChainRegistry`
     * @dev This value is NOT checkpointed. A new staleness period applies to ALL certificates, regardless of a certificate's reference timestamp
     * @dev This value is set by the AVS in the `CrossChainRegistry` and transported by the multichain protocol when the operator table is updated
     */
    function maxOperatorTableStaleness(
        OperatorSet memory operatorSet
    ) external view returns (uint32);

    /**
     * @notice The latest reference timestamp of the operator table for a given operatorSet. This value is
     *     updated each time an operator table is updated
     * @param operatorSet The operatorSet to get the latest reference timestamp of
     * @return The latest reference timestamp, 0 if the operatorSet has never been updated
     * @dev The latest reference timestamp is set when the operator table is updated
     * @dev The reference timestamp denotes the timestamp at which the operator table was calculated by the multichain protocol
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
     * @dev The stake weights are defined in the AVS's `OperatorTableCalculator` and transported by the multichain protocol. An example
     *      of this can be [slashable_stake, delegated_stake, strategy_i_stake, ...], where each stake type is an element in the array.
     *      The stake weights are defined by the operatorSet's `OperatorTableCalculator` and transported by the multichain protocol
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
