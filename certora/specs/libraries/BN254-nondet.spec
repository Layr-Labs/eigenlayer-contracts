methods {
    // BN254 implements elliptic curve operations, let's NONDET all of them
    function BN254.plus(BN254.G1Point memory a, BN254.G1Point memory b) internal returns (BN254.G1Point memory) => CVL_randomPoint();
    function BN254.pairing(BN254.G1Point memory a1, BN254.G2Point memory a2, BN254.G1Point memory b1, BN254.G2Point memory b2) internal returns (bool) => NONDET;
    function BN254.safePairing(BN254.G1Point memory a, BN254.G2Point memory b, BN254.G1Point memory c, BN254.G2Point memory d, uint256) internal returns (bool,bool) => NONDET;
    function BN254.scalar_mul(BN254.G1Point memory p, uint256 s) internal returns (BN254.G1Point memory) => CVL_randomPoint();

    function BN254.expMod(uint256 b, uint256 e, uint256 m) internal returns (uint256) => CVL_expMod(b,e,m);
    function BN254.hashToG1(bytes32 _x) internal returns (BN254.G1Point memory) => CVL_randomPoint();
}

function CVL_plus(BN254.G1Point a, BN254.G1Point b) returns BN254.G1Point {
    BN254.G1Point res;
    return res;
}

function CVL_randomPoint() returns BN254.G1Point {
    BN254.G1Point res;
    return res;
}

function CVL_expMod(uint256 b, uint256 e, uint256 m) returns uint256 {
    return require_uint256((b^e) % m);
}