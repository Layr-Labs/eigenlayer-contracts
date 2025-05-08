// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";

struct Validator {
    bool isDummy;
    bool isSlashed;
    bytes32 pubkeyHash;
    bytes withdrawalCreds;
    uint64 effectiveBalanceGwei;
    uint64 activationEpoch;
    uint64 exitEpoch;

    // cumulative unprocessed withdraw requests
    uint64 pendingBalanceToWithdrawGwei;
}

library LibValidator {

    /// @dev Generates a faux-pubkey from a uint40 validator index
    function toPubkey(uint40 index) internal pure returns (bytes memory pubkey) {
        pubkey = new bytes(48);
        assembly { mstore(add(48, pubkey), index) }
    }

    /// @dev Converts a validator pubkey to the index it uses in BeaconChainMock
    /// NOTE this assumes a valid size pubkey
    function toIndex(bytes memory pubkey) internal pure returns (uint40 validatorIndex) {
        assembly { validatorIndex := mload(add(48, pubkey)) }
    }

    /// @dev Computes a pubkey hash from a validator's pubkey
    /// NOTE this assumes a valid size pubkey
    function pubkeyHash(bytes memory pubkey) internal pure returns (bytes32) {
        return sha256(abi.encodePacked(pubkey, bytes16(0)));
    }

    function getValidatorFields(Validator memory self) internal pure returns (bytes32[] memory fields) {
        fields = new bytes32[](8);

        fields[BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = self.pubkeyHash;
        fields[BeaconChainProofs.VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = bytes32(self.withdrawalCreds);
        fields[BeaconChainProofs.VALIDATOR_BALANCE_INDEX] = toLittleEndianUint64(self.effectiveBalanceGwei);
        fields[BeaconChainProofs.VALIDATOR_SLASHED_INDEX] = bytes32(abi.encode(self.isSlashed));
        fields[BeaconChainProofs.VALIDATOR_ACTIVATION_EPOCH_INDEX] = toLittleEndianUint64(self.activationEpoch);
        fields[BeaconChainProofs.VALIDATOR_EXIT_EPOCH_INDEX] = toLittleEndianUint64(self.exitEpoch);
    }

    /// @dev Returns whether the validator is considered "active" at the given epoch
    function isActiveAt(Validator memory self, uint64 epoch) internal pure returns (bool) {
        return self.activationEpoch <= epoch && epoch < self.exitEpoch;
    }

    /// @dev Returns true if the validator has initiated exit
    function isExiting(Validator memory self) internal pure returns (bool) {
        return self.exitEpoch != BeaconChainProofs.FAR_FUTURE_EPOCH;
    }

    /// @dev Returns the withdrawal address of the validator
    /// NOTE this assumes the validator has 0x01 or 0x02 withdrawal credentials
    function withdrawalAddress(Validator memory self) internal pure returns (address addr) {
        bytes32 creds = bytes32(self.withdrawalCreds);
        uint160 mask = type(uint160).max;

        assembly { addr := and(creds, mask) }
    }

    /// @dev Returns true if the validator does not have 0x01/0x02 withdrawal credentials
    function hasBLSWC(Validator memory self) internal pure returns (bool) {
        return self.withdrawalCreds[0] != 0x01 && self.withdrawalCreds[0] != 0x02;
    }

    /// @dev Returns true IFF the validator has 0x02 withdrawal credentials
    function hasCompoundingWC(Validator memory self) internal pure returns (bool) {
        return self.withdrawalCreds[0] == 0x02;
    }

    /// @dev Opposite of Endian.fromLittleEndianUint64
    function toLittleEndianUint64(uint64 num) internal pure returns (bytes32) {
        uint lenum;

        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint((num & 0xFF) << 56);
        lenum |= uint((num & 0xFF00) << 40);
        lenum |= uint((num & 0xFF0000) << 24);
        lenum |= uint((num & 0xFF000000) << 8);
        lenum |= uint((num & 0xFF00000000) >> 8);
        lenum |= uint((num & 0xFF0000000000) >> 24);
        lenum |= uint((num & 0xFF000000000000) >> 40);
        lenum |= uint((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
    }
}