// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";

import "./IECDSATableCalculator.sol";
import "./IBN254TableCalculator.sol";
import "./IECDSACertificateVerifier.sol";
import "./IBN254CertificateVerifier.sol";
import "./ICrossChainRegistry.sol";

interface IOperatorTableUpdaterErrors {
    /// @notice Thrown when the GlobalTableRoot update fails
    error GlobalTableRootUpdateFailed();
}

interface IOperatorTableUpdaterEvents {
    /**
     * @notice Emitted when a new global table root is set
     * @param referenceTimestamp the timestamp of the global table root
     * @param globalTableRoot the root of the global table
     */
    event NewglobalTableRoot(uint32 referenceTimestamp, bytes32 globalTableRoot);
}

interface IOperatorTableUpdater is
    IOperatorTableUpdaterErrors,
    IOperatorTableUpdaterEvents,
    IECDSACertificateVerifierTypes,
    IBN254CertificateVerifierTypes,
    ICrossChainRegistryTypes
{
    /**
     * @notice Confirms Global operator table root
     * @param globalTableRootCert certificate of the root
     * @param globalTableRoot merkle root of all operatorSet tables
     * @param referenceTimestamp timestamp of the root
     * @dev Any entity can submit with a valid certificate signed off by the `globalRootConfirmerSet`
     * @dev The `msgHash` in the `globalOperatorTableRootCert` is the hash of the `globalOperatorTableRoot`
     */
    function confirmGlobalTableRoot(
        BN254Certificate calldata globalTableRootCert,
        bytes32 globalTableRoot,
        uint32 referenceTimestamp
    ) external;

    /**
     * @notice Set the operatorSet which certifies against global roots
     * @param operatorSet the operatorSet which certifies against global roots
     * @dev The `operatorSet` is used to verify the certificate of the global table root
     * @dev Only callable by the owner of the contract
     */
    function setGlobalRootConfirmerSet(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice The threshold, in bps, for a global root to be signed off on and updated
     * @dev Only callable by the owner of the contract
     */
    function setGlobalRootConfirmationThreshold(
        uint16 bps
    ) external;

    /**
     * @notice update a BN254 operator table in the CertificateVerifier
     * @param referenceTimestamp the reference block number of the globalTableRoot
     * @param globalTableRoot the new globalTableRoot
     * @param operatorSetIndex the index of the given operatorSet being updated
     * @param proof the proof of the leaf at index against the globalTableRoot
     * @param operatorSet the operatorSet being proven
     * @param operatorSetInfo the operatorSetInfo of the operator table
     * @param config the configuration of the operatorSet
     * @dev globalTableRoot must be confirmed majority certified by globalTableRootSet
     */
    function updateBN254OperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        OperatorSet calldata operatorSet,
        BN254OperatorSetInfo calldata operatorSetInfo,
        OperatorSetConfig calldata config
    ) external;

    /**
     * @notice updates an ECDSA operator table in the CertificateVerifier
     * @param referenceTimestamp the reference block number of the globalTableRoot
     * @param globalTableRoot the new globalTableRoot
     * @param operatorSetIndex the index of the given operatorSet being updated
     * @param proof the proof of the leaf at index against the globalTableRoot
     * @param operatorSet the operatorSet being proven
     * @param operatorInfos the operatorInfos of the operator table
     * @param config the configuration of the operatorSet
     * @dev globalTableRoot must be confirmed majority certified by globalTableRootSet
     */
    function updateECDSAOperatorTable(
        uint32 referenceTimestamp,
        bytes32 globalTableRoot,
        uint32 operatorSetIndex,
        bytes calldata proof,
        OperatorSet calldata operatorSet,
        ECDSAOperatorInfo[] calldata operatorInfos,
        OperatorSetConfig calldata config
    ) external;

    /**
     * @notice Get the current global table root
     * @return globalTableRoot the current global table root
     */
    function getCurrentGlobalTableRoot() external view returns (bytes32 globalTableRoot);

    /**
     * @notice Get the table root by timestamp
     * @param referenceTimestamp the timestamp of the table root
     * @return tableRoot the table root at the given timestamp
     */
    function getTableRootByTimestamp(
        uint32 referenceTimestamp
    ) external view returns (bytes32 tableRoot);
}
