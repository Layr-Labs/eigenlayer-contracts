// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";

import {BN254} from "src/contracts/libraries/BN254.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";
import {IBN254CertificateVerifier} from "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import {IKeyRegistrar} from "src/contracts/interfaces/IKeyRegistrar.sol";

// Test utilities (FFI-based G2 mul and add)
import {BN256G2} from "src/test/utils/BN256G2.sol";

contract GenerateBN254Cert is Script {
    using BN254 for BN254.G1Point;

    function _g2Mul(
        uint256 sk
    ) internal returns (BN254.G2Point memory g2) {
        // Multiply G2 generator by private key using FFI helper
        // BN256G2 has no mul API; reuse OperatorWalletLib.mul logic via BN256G2 + ffi helper
        // We reconstruct the same flow inline here to avoid importing the VM cheat directly.
        // BN256G2.ECTwistMul is not available, so we rely on the go helper invoked by BN256G2 via ffi
        // However BN256G2 library exposes only add; we keep the FFI path present in src/test/utils/BN256G2.go via forge --ffi.

        // We call the same go helper that tests use, via vm.ffi under the hood.
        // Here we replicate it by invoking the cheatcode address directly.
        address HEVM = address(uint160(uint256(keccak256("hevm cheat code"))));
        bytes memory out;

        // x1 (index 1)
        {
            bytes memory input = abi.encodeWithSignature("ffi(string[])", _ffiArgs(sk, "1"));
            (bool ok, bytes memory res) = HEVM.call(input);
            require(ok, "ffi x1 failed");
            out = res;
            g2.X[1] = abi.decode(out, (uint256));
        }
        // x0 (index 2)
        {
            bytes memory input = abi.encodeWithSignature("ffi(string[])", _ffiArgs(sk, "2"));
            (bool ok, bytes memory res) = HEVM.call(input);
            require(ok, "ffi x0 failed");
            out = res;
            g2.X[0] = abi.decode(out, (uint256));
        }
        // y1 (index 3)
        {
            bytes memory input = abi.encodeWithSignature("ffi(string[])", _ffiArgs(sk, "3"));
            (bool ok, bytes memory res) = HEVM.call(input);
            require(ok, "ffi y1 failed");
            out = res;
            g2.Y[1] = abi.decode(out, (uint256));
        }
        // y0 (index 4)
        {
            bytes memory input = abi.encodeWithSignature("ffi(string[])", _ffiArgs(sk, "4"));
            (bool ok, bytes memory res) = HEVM.call(input);
            require(ok, "ffi y0 failed");
            out = res;
            g2.Y[0] = abi.decode(out, (uint256));
        }
    }

    function _ffiArgs(
        uint256 sk,
        string memory which
    ) internal pure returns (string[] memory args) {
        args = new string[](5);
        args[0] = "go";
        args[1] = "run";
        args[2] = "src/test/utils/g2mul.go";
        args[3] = vm.toString(sk);
        args[4] = which;
    }

    // Single-signer helper: gets apk from KeyRegistrar, no FFI, no broadcasts
    function runSingle(
        address bn254Verifier,
        address keyRegistrar,
        address avs,
        uint32 setId,
        address operator,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint256 signerSk
    ) external view returns (bytes32 digest, BN254.G1Point memory sigAgg, BN254.G2Point memory apkAgg) {
        digest = IBN254CertificateVerifier(bn254Verifier).calculateCertificateDigest(referenceTimestamp, messageHash);

        BN254.G1Point memory h = BN254.hashToG1(digest);
        sigAgg = h.scalar_mul(signerSk);

        (, BN254.G2Point memory pkG2) = IKeyRegistrar(keyRegistrar).getBN254Key(OperatorSet(avs, setId), operator);
        apkAgg = pkG2;

        console2.log("digest");
        console2.logBytes32(digest);
        console2.log("sigAgg.X");
        console2.logUint(sigAgg.X);
        console2.log("sigAgg.Y");
        console2.logUint(sigAgg.Y);
        console2.log("apk.X0");
        console2.logUint(apkAgg.X[0]);
        console2.log("apk.X1");
        console2.logUint(apkAgg.X[1]);
        console2.log("apk.Y0");
        console2.logUint(apkAgg.Y[0]);
        console2.log("apk.Y1");
        console2.logUint(apkAgg.Y[1]);
    }

    function run(
        address bn254Verifier,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint256 signerSk1,
        uint256 signerSk2
    ) external returns (bytes32 digest, BN254.G1Point memory sigAgg, BN254.G2Point memory apkAgg) {
        digest = IBN254CertificateVerifier(bn254Verifier).calculateCertificateDigest(referenceTimestamp, messageHash);

        BN254.G1Point memory h = BN254.hashToG1(digest);
        BN254.G1Point memory sig1 = h.scalar_mul(signerSk1);
        BN254.G1Point memory sig2 = h.scalar_mul(signerSk2);
        sigAgg = sig1.plus(sig2);

        BN254.G2Point memory pk1 = _g2Mul(signerSk1);
        BN254.G2Point memory pk2 = _g2Mul(signerSk2);
        // Aggregate in G2: apk = pk1 + pk2
        (apkAgg.X[0], apkAgg.X[1], apkAgg.Y[0], apkAgg.Y[1]) =
            BN256G2.ECTwistAdd(pk1.X[0], pk1.X[1], pk1.Y[0], pk1.Y[1], pk2.X[0], pk2.X[1], pk2.Y[0], pk2.Y[1]);

        console2.log("digest");
        console2.logBytes32(digest);
        console2.log("sigAgg.X");
        console2.logUint(sigAgg.X);
        console2.log("sigAgg.Y");
        console2.logUint(sigAgg.Y);
        console2.log("apk.X0");
        console2.logUint(apkAgg.X[0]);
        console2.log("apk.X1");
        console2.logUint(apkAgg.X[1]);
        console2.log("apk.Y0");
        console2.logUint(apkAgg.Y[0]);
        console2.log("apk.Y1");
        console2.logUint(apkAgg.Y[1]);
    }

    // Fully offline: compute digest locally and accept G2 pubkey directly.
    function runRaw(
        uint32 referenceTimestamp,
        bytes32 messageHash,
        uint256 signerSk,
        uint256 apkX0,
        uint256 apkX1,
        uint256 apkY0,
        uint256 apkY1
    ) external view returns (bytes32 digest, BN254.G1Point memory sig, BN254.G2Point memory apk) {
        bytes32 TYPEHASH = keccak256("BN254Certificate(uint32 referenceTimestamp,bytes32 messageHash)");
        digest = keccak256(abi.encode(TYPEHASH, referenceTimestamp, messageHash));

        BN254.G1Point memory h = BN254.hashToG1(digest);
        sig = h.scalar_mul(signerSk);

        apk.X[0] = apkX0;
        apk.X[1] = apkX1;
        apk.Y[0] = apkY0;
        apk.Y[1] = apkY1;
    }
}
