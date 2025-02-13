// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/mixins/SignatureUtils.sol";

contract MockSigner {
    mapping(bytes32 => mapping(bytes => bool)) public validSignatures;

    function setValidSignature(bytes32 digest, bytes memory signature, bool valid) public {
        validSignatures[digest][signature] = valid;
    }

    function isValidSignatureNow(bytes32 digest, bytes memory signature) public view returns (bool) {
        return validSignatures[digest][signature];
    }
}

contract SignatureUtilsHarness is SignatureUtils {
    function calculateSignableDigest(bytes32 hash) public view returns (bytes32) {
        return _calculateSignableDigest(hash);
    }

    function checkIsValidSignatureNow(
        address signer,
        bytes32 digest,
        bytes memory signature,
        uint256 expiry
    ) public view {
        _checkIsValidSignatureNow(signer, digest, signature, expiry);
    }
}

contract SignatureUtilsUnit is Test {
    SignatureUtilsHarness harness;
    MockSigner mockSigner;
    uint256 signerPk;
    address signer;
    bytes32 hash;
    bytes32 digest;
    bytes32 expectedDomainSeparator;

    function setUp() public {
        vm.chainId(1);

        harness = new SignatureUtilsHarness();
        mockSigner = new MockSigner();
        signerPk = 1;
        signer = vm.addr(signerPk);
        
        hash = keccak256("");
        digest = harness.calculateSignableDigest(hash);

        expectedDomainSeparator = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)"),
                keccak256(bytes("EigenLayer")), 
                block.chainid, 
                address(harness)
            )
        );
    }

    function test_domainSeparator_NonZero() public view {
        assertTrue(harness.domainSeparator() != 0, "The domain separator should be non-zero");
        assertTrue(harness.domainSeparator() == expectedDomainSeparator, "The domain separator should be as expected");
    }

    function test_domainSeparator_NewChainId() public {
        bytes32 initialDomainSeparator = harness.domainSeparator();

        // Change the chain ID
        vm.chainId(9999);

        bytes32 newDomainSeparator = harness.domainSeparator();

        assertTrue(newDomainSeparator != 0, "The new domain separator should be non-zero");
        assertTrue(
            initialDomainSeparator != newDomainSeparator,
            "The domain separator should change when the chain ID changes"
        );
    }

    function test_checkIsValidSignatureNow_Expired() public {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPk, digest);

        vm.expectRevert(ISignatureUtils.SignatureExpired.selector);
        harness.checkIsValidSignatureNow(signer, digest, abi.encode(r, s, v), block.timestamp - 1);
    }

    function test_Revert_checkIsValidSignatureNow_InvalidSignature() public {
        vm.expectRevert(ISignatureUtils.InvalidSignature.selector);
        harness.checkIsValidSignatureNow(signer, digest, "", block.timestamp);
    }
}