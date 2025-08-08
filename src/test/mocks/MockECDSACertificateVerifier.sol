// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IECDSACertificateVerifier} from "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";

contract MockECDSACertificateVerifier is IECDSACertificateVerifier {
    // Mapping to store operator set owners for testing
    mapping(bytes32 => address) public operatorSetOwners;

    function setOperatorSetOwner(OperatorSet memory operatorSet, address owner) external {
        operatorSetOwners[keccak256(abi.encode(operatorSet.avs, operatorSet.id))] = owner;
    }

    function updateOperatorTable(
        OperatorSet calldata, /*operatorSet*/
        uint32, /*referenceTimestamp*/
        ECDSAOperatorInfo[] calldata, /*operatorInfos*/
        OperatorSetConfig calldata /*operatorSetConfig*/
    ) external pure {}

    function verifyCertificate(OperatorSet calldata, /*operatorSet*/ ECDSACertificate memory /*cert*/ )
        external
        pure
        returns (uint[] memory signedStakes, address[] memory signers)
    {
        return (new uint[](0), new address[](0));
    }

    function verifyCertificateProportion(
        OperatorSet calldata, /*operatorSet*/
        ECDSACertificate memory, /*cert*/
        uint16[] memory /*totalStakeProportionThresholds*/
    ) external pure returns (bool, address[] memory signers) {
        return (true, new address[](0));
    }

    function verifyCertificateNominal(
        OperatorSet calldata, /*operatorSet*/
        ECDSACertificate memory, /*cert*/
        uint[] memory /*totalStakeNominalThresholds*/
    ) external pure returns (bool, address[] memory signers) {
        return (true, new address[](0));
    }

    // Implement IBaseCertificateVerifier required functions
    function operatorTableUpdater(OperatorSet memory /*operatorSet*/ ) external pure returns (address) {
        return address(0);
    }

    function getLatestReferenceTimestamp(OperatorSet memory /*operatorSet*/ ) external pure returns (uint32) {
        return 0;
    }

    function getOperatorSetOwner(OperatorSet memory operatorSet) external view returns (address) {
        // Return the configured owner, or the AVS address by default
        address owner = operatorSetOwners[keccak256(abi.encode(operatorSet.avs, operatorSet.id))];
        return owner != address(0) ? owner : operatorSet.avs;
    }

    function latestReferenceTimestamp(OperatorSet memory /*operatorSet*/ ) external pure returns (uint32) {
        return 0;
    }

    function maxOperatorTableStaleness(OperatorSet memory /*operatorSet*/ ) external pure returns (uint32) {
        return 86_400;
    }

    function getCachedSignerList(OperatorSet calldata, /*operatorSet*/ uint32 /*referenceTimestamp*/ )
        external
        pure
        returns (address[] memory)
    {
        return new address[](0);
    }

    function getOperatorInfos(OperatorSet calldata, /*operatorSet*/ uint32 /*referenceTimestamp*/ )
        external
        pure
        returns (ECDSAOperatorInfo[] memory)
    {
        return new ECDSAOperatorInfo[](0);
    }

    function getOperatorInfo(OperatorSet calldata, /*operatorSet*/ uint32, /*referenceTimestamp*/ uint /*operatorIndex*/ )
        external
        pure
        returns (ECDSAOperatorInfo memory)
    {
        uint[] memory weights = new uint[](0);
        return ECDSAOperatorInfo({pubkey: address(0), weights: weights});
    }

    function getOperatorCount(OperatorSet memory, /*operatorSet*/ uint32 /*referenceTimestamp*/ ) external pure returns (uint) {
        return 0;
    }

    function getTotalStakes(OperatorSet calldata, /*operatorSet*/ uint32 /*referenceTimestamp*/ ) external pure returns (uint[] memory) {
        return new uint[](0);
    }

    function domainSeparator() external pure returns (bytes32) {
        return bytes32(0);
    }

    function calculateCertificateDigest(uint32 referenceTimestamp, bytes32 messageHash) external pure returns (bytes32) {
        return keccak256(abi.encode(referenceTimestamp, messageHash));
    }

    function getTotalStakeWeights(OperatorSet calldata, uint32) external pure returns (uint[] memory) {
        return new uint[](0);
    }

    function calculateCertificateDigestBytes(uint32, bytes32) external pure returns (bytes memory) {
        return "";
    }

    function isReferenceTimestampSet(OperatorSet memory, uint32) external pure returns (bool) {
        return true;
    }
}
