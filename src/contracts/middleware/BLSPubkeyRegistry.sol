// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";

import "../libraries/BN254.sol";

import "forge-std/Test.sol";


contract BLSPubkeyRegistry is IBLSPubkeyRegistry, Test {
    using BN254 for BN254.G1Point;

    /// @notice the hash of the zero pubkey aka BN254.G1Point(0,0)
    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    /// @notice the current global aggregate pubkey among all of the quorums
    BN254.G1Point public globalApk;
    /// @notice the registry coordinator contract
    IRegistryCoordinator public registryCoordinator;
    /// @notice the BLSPublicKeyCompendium contract against which pubkey ownership is checked
    IBLSPublicKeyCompendium public immutable pubkeyCompendium;

    // list of all updates made to the global aggregate pubkey
    ApkUpdate[] public globalApkUpdates;
    // mapping of quorumNumber => ApkUpdate[], tracking the aggregate pubkey updates of every quorum
    mapping(uint8 => ApkUpdate[]) public quorumApkUpdates;
    // mapping of quorumNumber => current aggregate pubkey of quorum
    mapping(uint8 => BN254.G1Point) private quorumApk;

    event PubkeyAdded(
        address operator,
        BN254.G1Point pubkey
    );
    event PubkeyRemoved(
        address operator,
        BN254.G1Point pubkey
    );  

    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        IBLSPublicKeyCompendium _pubkeyCompendium
    ){
        registryCoordinator = _registryCoordinator;
        pubkeyCompendium = _pubkeyCompendium;
    }

    /**
     * @notice Registers the `operator`'s pubkey for the specified `quorumNumbers`.
     * @param operator The address of the operator to register.
     * @param quorumNumbers The quorum numbers the operator is registering for, where each byte is an 8 bit integer quorumNumber.
     * @param pubkey The operator's BLS public key.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions:
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already registered
     */
    function registerOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        // calculate hash of the operator's pubkey
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        require(pubkeyHash != ZERO_PK_HASH, "BLSPubkeyRegistry.registerOperator: cannot register zero pubkey");
        require(quorumNumbers.length > 0, "BLSPubkeyRegistry.registerOperator: must register for at least one quorum");
        // ensure that the operator owns their public key by referencing the BLSPubkeyCompendium
        require(pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,"BLSPubkeyRegistry.registerOperator: operator does not own pubkey");
        // update each quorum's aggregate pubkey
        _processQuorumApkUpdate(quorumNumbers, pubkey);
        // update the global aggregate pubkey
        _processGlobalApkUpdate(pubkey);
        // emit event so offchain actors can update their state
        emit PubkeyAdded(operator, pubkey);
        return pubkeyHash;
    }

    /**
     * @notice Deregisters the `operator`'s pubkey for the specified `quorumNumbers`.
     * @param operator The address of the operator to deregister.
     * @param quorumNumbers The quourm numbers the operator is deregistering from, where each byte is an 8 bit integer quorumNumber.
     * @param pubkey The public key of the operator.
     * @dev access restricted to the RegistryCoordinator
     * @dev Preconditions:
     *         1) `quorumNumbers` has no duplicates
     *         2) `quorumNumbers.length` != 0
     *         3) `quorumNumbers` is ordered in ascending order
     *         4) the operator is not already deregistered
     *         5) `quorumNumbers` is the same as the parameter use when registering
     *         6) `pubkey` is the same as the parameter used when registering
     */   
    function deregisterOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        require(quorumNumbers.length > 0, "BLSPubkeyRegistry.deregisterOperator: must register for at least one quorum");
        //ensure that the operator owns their public key by referencing the BLSPubkeyCompendium
        // TODO: Do we need this check given precondition?
        require(pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,"BLSPubkeyRegistry.deregisterOperator: operator does not own pubkey");
        // update each quorum's aggregate pubkey
        _processQuorumApkUpdate(quorumNumbers, pubkey.negate());
        // update the global aggregate pubkey
        _processGlobalApkUpdate(pubkey.negate());
        // emit event so offchain actors can update their state
        emit PubkeyRemoved(operator, pubkey);
        return pubkeyHash;
    }

    /// @notice Returns the current APK for the provided `quorumNumber `
    function getApkForQuorum(uint8 quorumNumber) external view returns(BN254.G1Point memory) {
        return quorumApk[quorumNumber];
    }

    /// @notice Returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory){
        return quorumApkUpdates[quorumNumber][index];
    }

    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     * @param quorumNumber is the quorum whose ApkHash is being retrieved
     * @param blockNumber is the number of the block for which the latest ApkHash muust be retrieved
     * @param index is the provided witness of the onchain index calculated offchain
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory quorumApkUpdate = quorumApkUpdates[quorumNumber][index];
        _validateApkHashForQuorumAtBlockNumber(quorumApkUpdate, blockNumber);
        return quorumApkUpdate.apkHash;
    }

	/**
     * @notice get hash of the apk among all quourums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     * @param blockNumber is the number of the block for which the latest ApkHash muust be retrieved
     * @param index is the provided witness of the onchain index calculated offchain
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory globalApkUpdate = globalApkUpdates[index];
        _validateApkHashForQuorumAtBlockNumber(globalApkUpdate, blockNumber);
        return globalApkUpdate.apkHash;
    }
    
    /// @notice Returns the length of ApkUpdates for the provided `quorumNumber`
    function getQuorumApkHistoryLength(uint8 quorumNumber) external view returns(uint32){
        return uint32(quorumApkUpdates[quorumNumber].length);
    }

    /// @notice Returns the length of ApkUpdates for the global APK
    function getGlobalApkHistoryLength() external view returns(uint32){
        return uint32(globalApkUpdates.length);
    }

    function _processGlobalApkUpdate(BN254.G1Point memory point) internal returns(bytes32){
        if (globalApkUpdates.length > 0) {
            // update the previous global apk update with the current block number
            globalApkUpdates[globalApkUpdates.length - 1].nextUpdateBlockNumber = uint32(block.number);
        }
        globalApk = globalApk.plus(point);
        bytes32 globalApkHash = BN254.hashG1Point(globalApk);

        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = BN254.hashG1Point(globalApk);
        latestGlobalApkUpdate.updateBlockNumber = uint32(block.number);
        globalApkUpdates.push(latestGlobalApkUpdate);
    }

    function _processQuorumApkUpdate(uint8[] memory quorumNumbers, BN254.G1Point memory point) internal {
        BN254.G1Point memory apkBeforeUpdate;
        BN254.G1Point memory apkAfterUpdate;

        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];
            // load and store in memory in common case we need to access the length again
            uint256 quorumApkUpdatesLength = quorumApkUpdates[quorumNumber].length;
            if (quorumApkUpdatesLength > 0) {
                // update nextUpdateBlockNumber of the current latest ApkUpdate
                quorumApkUpdates[quorumNumber][quorumApkUpdates[quorumNumber].length - 1].nextUpdateBlockNumber = uint32(block.number);
            }

            // fetch the most apk before adding point
            apkBeforeUpdate = quorumApk[quorumNumber];
            // accumulate the given point into the quorum apk
            apkAfterUpdate = apkBeforeUpdate.plus(point);
            // update aggregate public key for this quorum
            quorumApk[quorumNumber] = apkAfterUpdate;
            
            // create new ApkUpdate to add to the mapping
            ApkUpdate memory latestApkUpdate;
            latestApkUpdate.apkHash = BN254.hashG1Point(apkAfterUpdate);
            latestApkUpdate.updateBlockNumber = uint32(block.number);
            quorumApkUpdates[quorumNumber].push(latestApkUpdate);
        }
    }

    /// @notice validates the the ApkUpdate was in fact the latest update at the given blockNumber
    function _validateApkHashForQuorumAtBlockNumber(ApkUpdate memory apkUpdate, uint32 blockNumber) internal pure {
        require(
            blockNumber >= apkUpdate.updateBlockNumber, 
            "BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: index too recent"
        );
        require(
            apkUpdate.nextUpdateBlockNumber == 0 || blockNumber < apkUpdate.nextUpdateBlockNumber, 
            "BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: not latest apk update"
        );
    }
}