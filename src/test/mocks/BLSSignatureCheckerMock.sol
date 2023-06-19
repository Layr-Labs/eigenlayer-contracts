// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/BLSSignatureChecker.sol";



contract BLSSignatureCheckerWrapper is BLSSignatureChecker {
    constructor(IBLSStakeRegistryCoordinator _registryCoordinator) BLSSignatureChecker(_registryCoordinator) {}
    // function checkSignature(
    //     bytes32 domainSeparator,
    //     bytes32 messageHash,
    //     bytes memory signature
    // ) public view returns (bool) {
    //     return _checkSignature(domainSeparator, messageHash, signature);
    // }
}