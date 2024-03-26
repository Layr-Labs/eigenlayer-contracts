// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

//small library for dealing with efficiently-packed signatures, where parameters v,r,s are packed into vs and r (64 bytes instead of 65)
library SignatureCompaction {
    bytes32 internal constant HALF_CURVE_ORDER = 0x7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0;

    function ecrecoverPacked(bytes32 hash, bytes32 r, bytes32 vs) internal pure returns (address) {
        (address recovered, ECDSA.RecoverError err) = ECDSA.tryRecover(hash, r, vs);
        require(err == ECDSA.RecoverError.NoError, "error in ecrecoverPacked");
        return recovered;
    }

    function packSignature(bytes32 r, bytes32 s, uint8 v) internal pure returns (bytes32, bytes32) {
        require(s <= HALF_CURVE_ORDER, "malleable signature, s too high");
        //v parity is a single bit, encoded as either v = 27 or v = 28 -- in order to recover the bit we subtract 27
        bytes32 vs = bytes32(uint256(bytes32(uint256(v) - 27) << 255) | uint256(s));
        return (r, vs);
    }

    //same as above, except doesn't take 'r' as argument since it is unneeded
    function packVS(bytes32 s, uint8 v) internal pure returns (bytes32) {
        require(s <= HALF_CURVE_ORDER, "malleable signature, s too high");
        //v parity is a single bit, encoded as either v = 27 or v = 28 -- in order to recover the bit we subtract 27
        return bytes32(uint256(bytes32(uint256(v) - 27) << 255) | uint256(s));
    }
}
