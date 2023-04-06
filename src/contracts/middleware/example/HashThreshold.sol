// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "../../interfaces/IServiceManager.sol";
import "./ECDSARegistry.sol";

/**
 * @title An EigenLayer middleware example service manager that slashes validators that sign a message that, when hashed 1000 types starts with less than 5 0s.
 * @author Layr Labs, Inc.
 */
contract HashThreshold is Ownable, IServiceManager {
    uint32 constant disputePeriodBlocks = 1 days / 12 seconds;
    uint8 constant numZeroes = 5;
    ISlasher public immutable slasher;
    ECDSARegistry public immutable registry;


    struct CertifiedMessageMetadata {
        bytes32 signaturesHash;
        uint32 validAfterBlock;
    }

    uint32 public taskNumber = 0;
    uint32 public latestTime = 0;
    mapping(bytes32 => CertifiedMessageMetadata) public certifiedMessageMetadatas;

    event MessageCertified(bytes32);

    modifier onlyRegistry {
        require(msg.sender == address(registry), "Only registry can call this function");
    }

    constructor(
        ISlasher _slasher,
        ECDSARegistry _registry
    ) {
        slasher = _slasher;
        registry = _registry;
    }

    function owner() public view override(Ownable, IServiceManager) returns (address) {
        return Ownable.owner();
    }

    function kiloHash(bytes32 message) public pure returns (bytes32) {
        bytes32 hash = message;
        for (uint256 i = 0; i < 1000; i++) {
            hash = keccak256(abi.encodePacked(hash));
        }
        return hash;
    }

    function submitSignatures(bytes32 message, bytes calldata signatures) external {
        // this makes it so that the signatures are viewable in calldata
        require(msg.sender == tx.origin, "EOA must call this function");
        uint128 stakeSigned = 0;
        for(uint256 i = 0; i < signatures.length; i += 65) {
            // we fetch all the signers and check their signatures and their stake
            address signer = ECDSA.recover(message, signatures[i:i+65]);
            require(registry.isActiveOperator(signer), "Signer is not an active operator");
            stakeSigned += registry.firstQuorumStakedByOperator(signer);
        }
        (uint96 totalStake,) = registry.totalStake();
        require(stakeSigned >= 666667 * uint256(totalStake) / 1000000, "Need more than 2/3 of stake to sign");

        uint32 newLatestTime = uint32(block.number + disputePeriodBlocks);

        certifiedMessageMetadatas[message] = CertifiedMessageMetadata({
            validAfterBlock: newLatestTime,
            signaturesHash: keccak256(signatures)
        });

        // increment global service manager values
        taskNumber++;
        latestTime = newLatestTime;
        
        emit MessageCertified(message);
    }

    function slashSigners(bytes32 message, bytes calldata signatures) external {
        CertifiedMessageMetadata memory certifiedMessageMetadata = certifiedMessageMetadatas[message];
        require(certifiedMessageMetadata.validAfterBlock > block.number, "Dispute period has passed");
        require(certifiedMessageMetadata.signaturesHash == keccak256(signatures), "Signatures do not match");
        require(kiloHash(message) >> (256 - numZeroes) == 0, "Message does not hash to enough zeroes");
        for (uint i = 0; i < signatures.length; i += 65) {
            slasher.freezeOperator(ECDSA.recover(message, signatures[i:i+65]));
        }
        certifiedMessageMetadatas[message].validAfterBlock = type(uint32).max;
    }

    function freezeOperator(address operator) external onlyRegistry {
        slasher.freezeOperator(operator);
    }

    function recordFirstStakeUpdate(address operator, uint32 serveUntil) external onlyRegistry {
        slasher.recordFirstStakeUpdate(operator, serveUntil);
    }

    function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntil) external onlyRegistry {
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntil);
    }

    function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntil, uint256 prevElement) external onlyRegistry {
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, prevElement);
    }
}
