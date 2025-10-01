// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/IBaseCertificateVerifier.sol";

contract OperatorTableUpdaterHarness is OperatorTableUpdater {
    constructor(
        IBN254CertificateVerifier _bn254CertificateVerifier,
        IECDSACertificateVerifier _ecdsaCertificateVerifier,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) OperatorTableUpdater(
        _bn254CertificateVerifier,
        _ecdsaCertificateVerifier,
        _pauserRegistry,
        _version
    ) {}

    /// @notice Decodes operator table bytes and returns the latest reference timestamp
    ///         for the operator set in the corresponding certificate verifier.
    function latestTsForOperatorTableBytes(bytes calldata operatorTableBytes)
        external
        view
        returns (uint32)
    {
        (OperatorSet memory set, IKeyRegistrar.CurveType curve,,) = _decodeOperatorTableBytes(operatorTableBytes);

        return IBaseCertificateVerifier(getCertificateVerifier(curve))
            .latestReferenceTimestamp(set);
    }


    function isReferenceTimestampSet(bytes calldata operatorTableBytes, uint32 referenceTimestamp) external view returns (bool){

        (OperatorSet memory set, IKeyRegistrar.CurveType curve,,) = _decodeOperatorTableBytes(operatorTableBytes);

        return IBaseCertificateVerifier(getCertificateVerifier(curve))
            .isReferenceTimestampSet(set, referenceTimestamp);
    }

    function validCurvetype (bytes calldata operatorTableBytes) external view returns (bool){

        (,IKeyRegistrar.CurveType curvetype,,) = _decodeOperatorTableBytes(operatorTableBytes);
        
        return curvetype == CurveType.BN254 || curvetype == CurveType.ECDSA;
    }

}