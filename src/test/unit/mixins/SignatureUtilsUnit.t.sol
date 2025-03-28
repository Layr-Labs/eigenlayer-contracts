// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/mixins/SignatureUtilsMixin.sol";

contract MockSigner {
    mapping(bytes32 => mapping(bytes => bool)) public validSignatures;

    function setValidSignature(bytes32 digest, bytes memory signature, bool valid) public {
        validSignatures[digest][signature] = valid;
    }

    function isValidSignatureNow(bytes32 digest, bytes memory signature) public view returns (bool) {
        return validSignatures[digest][signature];
    }
}

contract SignatureUtilsMixinUnit is Test, SignatureUtilsMixin("v0.0.0") {
    uint signerPk;
    address signer;
    bytes32 hash;
    bytes32 digest;
    bytes32 expectedDomainSeparator;

    function setUp() public {
        vm.chainId(1);

        signerPk = 1;
        signer = vm.addr(signerPk);

        hash = keccak256("");
        digest = _calculateSignableDigest(hash);

        expectedDomainSeparator = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256(bytes("EigenLayer")),
                keccak256(bytes(_majorVersion())),
                block.chainid,
                address(this)
            )
        );
    }

    function test_domainSeparator_NonZero() public view {
        assertTrue(domainSeparator() != 0, "The domain separator should be non-zero");
        assertTrue(domainSeparator() == expectedDomainSeparator, "The domain separator should be as expected");
    }

    function test_domainSeparator_NewChainId() public {
        bytes32 initialDomainSeparator = domainSeparator();

        // Change the chain ID
        vm.chainId(9999);

        bytes32 newDomainSeparator = domainSeparator();

        assertTrue(newDomainSeparator != 0, "The new domain separator should be non-zero");
        assertTrue(initialDomainSeparator != newDomainSeparator, "The domain separator should change when the chain ID changes");
    }

    /// forge-config: default.allow_internal_expect_revert = true
    function test_checkIsValidSignatureNow_Expired() public {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPk, digest);

        vm.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        _checkIsValidSignatureNow(signer, digest, abi.encode(r, s, v), block.timestamp - 1);
    }

    // function testFail_checkIsValidSignatureNow_InvalidSignature() public {
    //     _checkIsValidSignatureNow(signer, digest, "", block.timestamp);
    // }
}
