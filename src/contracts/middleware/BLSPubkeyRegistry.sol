// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";

import "../libraries/BN254.sol";

import "forge-std/Test.sol";


contract BLSPubkeyRegistry is IBLSPubkeyRegistry, Test {
    using BN254 for BN254.G1Point;

    // represents the hash of the zero pubkey aka BN254.G1Point(0,0)
    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    // the current global aggregate pubkey
    BN254.G1Point public globalApk;
    IRegistryCoordinator public registryCoordinator;
    IBLSPublicKeyCompendium public immutable pubkeyCompendium;

    // list of all updates made to the global aggregate pubkey
    ApkUpdate[] public globalApkUpdates;
    // mapping of quorumNumber => ApkUpdate[], tracking the aggregate pubkey updates of every quorum
    mapping(uint8 => ApkUpdate[]) public quorumToApkUpdates;
    // mapping of quorumNumber => current aggregate pubkey of quorum
    mapping(uint8 => BN254.G1Point) public quorumToApk;

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

        _initializeApkUpdates();
    }

    /**
     * @notice registers `operator` with the given `pubkey` for the specified `quorumNumbers`
     * @dev Permissioned by RegistryCoordinator
     */
    function registerOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        //calculate hash of the operator's pubkey
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        require(pubkeyHash != ZERO_PK_HASH, "BLSRegistry._registerOperator: cannot register zero pubkey");
        require(quorumNumbers.length > 0, "BLSRegistry._registerOperator: must register for at least one quorum");
        //ensure that the operator owns their public key by referencing the BLSPubkeyCompendium
        require(pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,"BLSRegistry._registerOperator: operator does not own pubkey");
        // update each quorum's aggregate pubkey
        _processQuorumApkUpdate(quorumNumbers, pubkey);
        // update the global aggregate pubkey
        _processGlobalApkUpdate(pubkey);

        emit PubkeyAdded(operator, pubkey);
        return pubkeyHash;
    }

    /**
     * @notice deregisters `operator` with the given `pubkey` for the quorums specified by `quorumBitmap`
     * @dev Permissioned by RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        require(quorumNumbers.length > 0, "BLSRegistry._deregisterOperator: must register for at least one quorum");
        //ensure that the operator owns their public key by referencing the BLSPubkeyCompendium
        require(pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,"BLSRegistry._deregisterOperator: operator does not own pubkey");
        // update each quorum's aggregate pubkey
        _processQuorumApkUpdate(quorumNumbers, pubkey.negate());
        // update the global aggregate pubkey
        _processGlobalApkUpdate(pubkey.negate());

        emit PubkeyRemoved(operator, pubkey);
        return pubkeyHash;
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory){
        return quorumToApkUpdates[quorumNumber][index];
    }

    function quorumApk(uint8 quorumNumber) external view returns (BN254.G1Point memory){
        return quorumToApk[quorumNumber];
    }

    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory quorumApkUpdate = quorumToApkUpdates[quorumNumber][index];
        _validateApkHashForQuorumAtBlockNumber(quorumApkUpdate, blockNumber);
        return quorumApkUpdate.apkHash;
    }

	/**
     * @notice get hash of the apk among all quourums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory globalApkUpdate = globalApkUpdates[index];
        _validateApkHashForQuorumAtBlockNumber(globalApkUpdate, blockNumber);
        return globalApkUpdate.apkHash;
    }

    function getQuorumApkHistoryLength(uint8 quorumNumber) external view returns(uint32){
        return uint32(quorumToApkUpdates[quorumNumber].length);
    }

    function getGlobalApkHistoryLength() external view returns(uint32){
        return uint32(globalApkUpdates.length);
    }

    function _processGlobalApkUpdate(BN254.G1Point memory point) internal returns(bytes32){
        globalApkUpdates[globalApkUpdates.length - 1].nextUpdateBlockNumber = uint32(block.number);
        globalApk = globalApk.plus(point);

        bytes32 globalApkHash = BN254.hashG1Point(globalApk);

        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = globalApkHash;
        latestGlobalApkUpdate.updateBlockNumber = uint32(block.number);
        globalApkUpdates.push(latestGlobalApkUpdate);

        return globalApkHash;
    }

    function _processQuorumApkUpdate(uint8[] memory quorumNumbers, BN254.G1Point memory point) internal {
        BN254.G1Point memory apkBeforeUpdate;
        BN254.G1Point memory apkAfterUpdate;

        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];
            
            apkBeforeUpdate = quorumToApk[quorumNumber];
            emit log_named_uint("apkBeforeUpdate.X", apkBeforeUpdate.X);
            emit log_named_uint("apkBeforeUpdate.Y", apkBeforeUpdate.Y);
            emit log_named_uint("pubkey.X", point.X);
            emit log_named_uint("pubkey.Y", point.Y);

            apkAfterUpdate = BN254.plus(apkBeforeUpdate, point);
            //update aggregate public key for this quorum
            quorumToApk[quorumNumber] = apkAfterUpdate;
            //update nextUpdateBlockNumber of the current latest ApkUpdate
            quorumToApkUpdates[quorumNumber][quorumToApkUpdates[quorumNumber].length - 1].nextUpdateBlockNumber = uint32(block.number);
            //create new ApkUpdate to add to the mapping
            ApkUpdate memory latestApkUpdate;
            latestApkUpdate.apkHash = BN254.hashG1Point(apkAfterUpdate);
            latestApkUpdate.updateBlockNumber = uint32(block.number);
            quorumToApkUpdates[quorumNumber].push(latestApkUpdate);
        }
    }

    function _initializeApkUpdates() internal {

        BN254.G1Point memory pk = BN254.G1Point(0,0);
        _processGlobalApkUpdate(pk);

        for (uint8 quorumNumber = 0; quorumNumber < 255; quorumNumber++) {

            quorumToApk[quorumNumber] = pk;
            quorumToApkUpdates[quorumNumber].push(ApkUpdate({
                apkHash: BN254.hashG1Point(pk),
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0
            }));
        }
        quorumToApk[255] = pk;
        quorumToApkUpdates[255].push(ApkUpdate({
            apkHash: BN254.hashG1Point(pk),
            updateBlockNumber: uint32(block.number),
            nextUpdateBlockNumber: 0
        }));
    }

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