// SPDX-License-Identifier: MIT
// several functions are taken or adapted from https://github.com/HarryR/solcrypto/blob/master/contracts/altbn128.sol (MIT license):
// Copyright 2017 Christian Reitwiessner
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// The remainder of the code in this library is written by LayrLabs Inc. and is also under an MIT license

pragma solidity ^0.8.27;

/**
 * @title Library for operations on the BN254 elliptic curve.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Contains BN254 parameters, common operations (addition, scalar mul, pairing), and BLS signature functionality.
 */
library BN254 {
    // modulus for the underlying field F_p of the elliptic curve
    uint256 internal constant FP_MODULUS =
        21_888_242_871_839_275_222_246_405_745_257_275_088_696_311_157_297_823_662_689_037_894_645_226_208_583;
    // modulus for the underlying field F_r of the elliptic curve
    uint256 internal constant FR_MODULUS =
        21_888_242_871_839_275_222_246_405_745_257_275_088_548_364_400_416_034_343_698_204_186_575_808_495_617;

    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    // Encoding of field elements is: X[1] * i + X[0]
    struct G2Point {
        uint256[2] X;
        uint256[2] Y;
    }

    /// @dev Thrown when the sum of two points of G1 fails
    error ECAddFailed();
    /// @dev Thrown when the scalar multiplication of a point of G1 fails
    error ECMulFailed();
    /// @dev Thrown when the scalar is too large.
    error ScalarTooLarge();
    /// @dev Thrown when the pairing check fails
    error ECPairingFailed();
    /// @dev Thrown when the exponentiation mod fails
    error ExpModFailed();

    function generatorG1() internal pure returns (G1Point memory) {
        return G1Point(1, 2);
    }

    // generator of group G2
    /// @dev Generator point in F_q2 is of the form: (x0 + ix1, y0 + it1).
    uint256 internal constant G2x1 =
        11_559_732_032_986_387_107_991_004_021_392_285_783_925_812_861_821_192_530_917_403_151_452_391_805_634;
    uint256 internal constant G2x0 =
        10_857_046_999_023_057_135_944_570_762_232_829_481_370_756_359_578_518_086_990_519_993_285_655_852_781;
    uint256 internal constant G2y1 =
        4_082_367_875_863_433_681_332_203_403_145_435_568_316_851_327_593_401_208_105_741_076_214_120_093_531;
    uint256 internal constant G2y0 =
        8_495_653_923_123_431_417_604_973_247_489_272_438_418_190_587_263_600_148_770_280_649_306_958_101_930;

    /// @notice returns the G2 generator
    /// @dev mind the ordering of the 1s and 0s!
    ///      this is because of the (unknown to us) convention used in the bn254 pairing precompile contract
    ///      "Elements a * i + b of F_p^2 are encoded as two elements of F_p, (a, b)."
    ///      https://github.com/ethereum/EIPs/blob/master/EIPS/eip-197.md#encoding
    function generatorG2() internal pure returns (G2Point memory) {
        return G2Point([G2x1, G2x0], [G2y1, G2y0]);
    }

    // negation of the generator of group G2
    /// @dev Generator point in F_q2 is of the form: (x0 + ix1, y0 + it1).
    uint256 internal constant nG2x1 =
        11_559_732_032_986_387_107_991_004_021_392_285_783_925_812_861_821_192_530_917_403_151_452_391_805_634;
    uint256 internal constant nG2x0 =
        10_857_046_999_023_057_135_944_570_762_232_829_481_370_756_359_578_518_086_990_519_993_285_655_852_781;
    uint256 internal constant nG2y1 =
        17_805_874_995_975_841_540_914_202_342_111_839_520_379_459_829_704_422_454_583_296_818_431_106_115_052;
    uint256 internal constant nG2y0 =
        13_392_588_948_715_843_804_641_432_497_768_002_650_278_120_570_034_223_513_918_757_245_338_268_106_653;

    function negGeneratorG2() internal pure returns (G2Point memory) {
        return G2Point([nG2x1, nG2x0], [nG2y1, nG2y0]);
    }

    bytes32 internal constant powersOfTauMerkleRoot = 0x22c998e49752bbb1918ba87d6d59dd0e83620a311ba91dd4b2cc84990b31b56f;

    /**
     * @param p Some point in G1.
     * @return The negation of `p`, i.e. p.plus(p.negate()) should be zero.
     */
    function negate(
        G1Point memory p
    ) internal pure returns (G1Point memory) {
        // The prime q in the base field F_q for G1
        if (p.X == 0 && p.Y == 0) {
            return G1Point(0, 0);
        } else {
            return G1Point(p.X, FP_MODULUS - (p.Y % FP_MODULUS));
        }
    }

    /**
     * @return r the sum of two points of G1
     */
    function plus(G1Point memory p1, G1Point memory p2) internal view returns (G1Point memory r) {
        uint256[4] memory input;
        input[0] = p1.X;
        input[1] = p1.Y;
        input[2] = p2.X;
        input[3] = p2.Y;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 6, input, 0x80, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 { invalid() }
        }

        require(success, ECAddFailed());
    }

    /**
     * @notice an optimized ecMul implementation that takes O(log_2(s)) ecAdds
     * @param p the point to multiply
     * @param s the scalar to multiply by
     * @dev this function is only safe to use if the scalar is 9 bits or less
     */
    function scalar_mul_tiny(BN254.G1Point memory p, uint16 s) internal view returns (BN254.G1Point memory) {
        require(s < 2 ** 9, ScalarTooLarge());

        // if s is 1 return p
        if (s == 1) {
            return p;
        }

        // the accumulated product to return
        BN254.G1Point memory acc = BN254.G1Point(0, 0);
        // the 2^n*p to add to the accumulated product in each iteration
        BN254.G1Point memory p2n = p;
        // value of most significant bit
        uint16 m = 1;
        // index of most significant bit
        uint8 i = 0;

        //loop until we reach the most significant bit
        while (s >= m) {
            unchecked {
                // if the  current bit is 1, add the 2^n*p to the accumulated product
                if ((s >> i) & 1 == 1) {
                    acc = plus(acc, p2n);
                }
                // double the 2^n*p for the next iteration
                p2n = plus(p2n, p2n);

                // increment the index and double the value of the most significant bit
                m <<= 1;
                ++i;
            }
        }

        // return the accumulated product
        return acc;
    }

    /**
     * @return r the product of a point on G1 and a scalar, i.e.
     *         p == p.scalar_mul(1) and p.plus(p) == p.scalar_mul(2) for all
     *         points p.
     */
    function scalar_mul(G1Point memory p, uint256 s) internal view returns (G1Point memory r) {
        uint256[3] memory input;
        input[0] = p.X;
        input[1] = p.Y;
        input[2] = s;
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 7, input, 0x60, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 { invalid() }
        }
        require(success, ECMulFailed());
    }

    /**
     *  @return The result of computing the pairing check
     *         e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
     *         For example,
     *         pairing([P1(), P1().negate()], [P2(), P2()]) should return true.
     */
    function pairing(
        G1Point memory a1,
        G2Point memory a2,
        G1Point memory b1,
        G2Point memory b2
    ) internal view returns (bool) {
        G1Point[2] memory p1 = [a1, b1];
        G2Point[2] memory p2 = [a2, b2];

        uint256[12] memory input;

        for (uint256 i = 0; i < 2; i++) {
            uint256 j = i * 6;
            input[j + 0] = p1[i].X;
            input[j + 1] = p1[i].Y;
            input[j + 2] = p2[i].X[0];
            input[j + 3] = p2[i].X[1];
            input[j + 4] = p2[i].Y[0];
            input[j + 5] = p2[i].Y[1];
        }

        uint256[1] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 8, input, mul(12, 0x20), out, 0x20)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 { invalid() }
        }

        require(success, ECPairingFailed());

        return out[0] != 0;
    }

    /**
     * @notice This function is functionally the same as pairing(), however it specifies a gas limit
     *         the user can set, as a precompile may use the entire gas budget if it reverts.
     */
    function safePairing(
        G1Point memory a1,
        G2Point memory a2,
        G1Point memory b1,
        G2Point memory b2,
        uint256 pairingGas
    ) internal view returns (bool, bool) {
        G1Point[2] memory p1 = [a1, b1];
        G2Point[2] memory p2 = [a2, b2];

        uint256[12] memory input;

        for (uint256 i = 0; i < 2; i++) {
            uint256 j = i * 6;
            input[j + 0] = p1[i].X;
            input[j + 1] = p1[i].Y;
            input[j + 2] = p2[i].X[0];
            input[j + 3] = p2[i].X[1];
            input[j + 4] = p2[i].Y[0];
            input[j + 5] = p2[i].Y[1];
        }

        uint256[1] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(pairingGas, 8, input, mul(12, 0x20), out, 0x20)
        }

        //Out is the output of the pairing precompile, either 0 or 1 based on whether the two pairings are equal.
        //Success is true if the precompile actually goes through (aka all inputs are valid)

        return (success, out[0] != 0);
    }

    /// @return hashedG1 the keccak256 hash of the G1 Point
    /// @dev used for BLS signatures
    function hashG1Point(
        BN254.G1Point memory pk
    ) internal pure returns (bytes32 hashedG1) {
        assembly {
            mstore(0, mload(pk))
            mstore(0x20, mload(add(0x20, pk)))
            hashedG1 := keccak256(0, 0x40)
        }
    }

    /// @return the keccak256 hash of the G2 Point
    /// @dev used for BLS signatures
    function hashG2Point(
        BN254.G2Point memory pk
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(pk.X[0], pk.X[1], pk.Y[0], pk.Y[1]));
    }

    /**
     * @notice adapted from https://github.com/HarryR/solcrypto/blob/master/contracts/altbn128.sol
     */
    function hashToG1(
        bytes32 _x
    ) internal view returns (G1Point memory) {
        uint256 beta = 0;
        uint256 y = 0;

        uint256 x = uint256(_x) % FP_MODULUS;

        while (true) {
            (beta, y) = findYFromX(x);

            // y^2 == beta
            if (beta == mulmod(y, y, FP_MODULUS)) {
                return G1Point(x, y);
            }

            x = addmod(x, 1, FP_MODULUS);
        }
        return G1Point(0, 0);
    }

    /**
     * Given X, find Y
     *
     *   where y = sqrt(x^3 + b)
     *
     * Returns: (x^3 + b), y
     */
    function findYFromX(
        uint256 x
    ) internal view returns (uint256, uint256) {
        // beta = (x^3 + b) % p
        uint256 beta = addmod(mulmod(mulmod(x, x, FP_MODULUS), x, FP_MODULUS), 3, FP_MODULUS);

        // y^2 = x^3 + b
        // this acts like: y = sqrt(beta) = beta^((p+1) / 4)
        uint256 y = expMod(beta, 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52, FP_MODULUS);

        return (beta, y);
    }

    function expMod(uint256 _base, uint256 _exponent, uint256 _modulus) internal view returns (uint256 retval) {
        bool success;
        uint256[1] memory output;
        uint256[6] memory input;
        input[0] = 0x20; // baseLen = new(big.Int).SetBytes(getData(input, 0, 32))
        input[1] = 0x20; // expLen  = new(big.Int).SetBytes(getData(input, 32, 32))
        input[2] = 0x20; // modLen  = new(big.Int).SetBytes(getData(input, 64, 32))
        input[3] = _base;
        input[4] = _exponent;
        input[5] = _modulus;
        assembly {
            success := staticcall(sub(gas(), 2000), 5, input, 0xc0, output, 0x20)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 { invalid() }
        }
        require(success, ExpModFailed());
        return output[0];
    }
}
