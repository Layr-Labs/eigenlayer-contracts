// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/utils/EigenLayerUnitTestSetup.sol";

import "src/test/mocks/BN254CertificateVerifierMock.sol";
import "src/test/mocks/ECDSACertificateVerifierMock.sol";

abstract contract EigenLayerMultichainUnitTestSetup is EigenLayerUnitTestSetup {
    using StdStyle for *;
    using ArrayLib for *;

    BN254CertificateVerifierMock bn254CertificateVerifierMock;
    ECDSACertificateVerifierMock ecdsaCertificateVerifierMock;

    function setUp() public virtual override {
        // Setup mocks
        super.setUp();

        // Deploy mocks
        bn254CertificateVerifierMock = new BN254CertificateVerifierMock();
        ecdsaCertificateVerifierMock = new ECDSACertificateVerifierMock();

        // Filter out mocks
        isExcludedFuzzAddress[address(bn254CertificateVerifierMock)] = true;
        isExcludedFuzzAddress[address(ecdsaCertificateVerifierMock)] = true;
    }
}
