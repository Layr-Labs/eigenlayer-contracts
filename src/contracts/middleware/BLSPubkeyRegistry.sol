// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../interfaces/IRegistryCoordinator.sol";
import "../interfaces/IBLSPublicKeyCompendium.sol";

import "../libraries/BN254.sol";


contract BLSPubkeyRegistry is IBLSPubkeyRegistry {

    BN254.G1Point internal _globalApk;

    IRegistryCoordinator public registryCoordinator;
    IBLSPublicKeyCompendium public immutable pubkeyCompendium;


    ApkUpdate[] public globalApkUpdateList;

    mapping(uint8 => ApkUpdate[]) public quorumToApkUpdates;
    
    mapping(uint8 => BN254.G1Point) public quorumToApk;

    event RegistrationEvent(
        address indexed operator,
        bytes32 pubkeyHash,
        bytes32 globalApkHash
    );

    event DeregistrationEvent(
        address indexed operator,
        bytes32 pubkeyHash,
        bytes32 globalApkHash
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


    function registerOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        _processQuorumApkUpdate(quorumNumbers, pubkey, true);
        bytes32 globalApkHash = _processGlobalApkUpdate(BN254.plus(_globalApk, pubkey));


         bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        emit RegistrationEvent(operator, pubkeyHash, globalApkHash);
    }

    /**
     * @notice deregisters `operator` with the given `pubkey` for the quorums specified by `quorumBitmap`
     * @dev Permissioned by RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint8[] memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        _processQuorumApkUpdate(quorumNumbers, pubkey, false);

        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        bytes32 globalApkHash = _processGlobalApkUpdate(BN254.plus(_globalApk, BN254.negate(pubkey)));

        emit DeregistrationEvent(operator, pubkeyHash, globalApkHash);
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory){
        return quorumToApkUpdates[quorumNumber][index];
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates that keep track of the sum of the APK amonng all quorums
    function globalApkUpdates(uint256 index) external view returns (ApkUpdate memory){
        return globalApkUpdateList[index];
    }

    function quorumApk(uint8 quorumNumber) external view returns (BN254.G1Point memory){
        return quorumToApk[quorumNumber];
    }

    function globalApk() external view returns (BN254.G1Point memory){
        return _globalApk;
    }


    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory quorumApkUpdate = quorumToApkUpdates[quorumNumber][index];
        require(blockNumber >= quorumApkUpdate.updateBlockNumber, "BLSPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex: index too recent");

        if (index != quorumToApkUpdates[quorumNumber].length - 1){
            require(blockNumber < quorumApkUpdate.nextUpdateBlockNumber, "BLSPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex: not latest apk update");
        }

        return quorumApkUpdate.apkHash;
    }

	/**
     * @notice get hash of the apk among all quourums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory globalApkUpdate = globalApkUpdateList[index];
        require(blockNumber >= globalApkUpdate.updateBlockNumber, "BLSPubkeyRegistry.getGlobalApkHashAtBlockNumberFromIndex: blockNumber too recent");

        if (index != globalApkUpdateList.length - 1){
            require(blockNumber < globalApkUpdate.nextUpdateBlockNumber, "getGlobalApkHashAtBlockNumberFromIndex.getGlobalApkHashAtBlockNumberFromIndex: not latest apk update");
        }

        return globalApkUpdate.apkHash;
    }

    function getQuorumApkHistoryLength(uint8 quorumNumber) external view returns(uint32){
        return uint32(quorumToApkUpdates[quorumNumber].length);
    }

    function getGlobalApkHistoryLength() external view returns(uint32){
        return uint32(globalApkUpdateList.length);
    }




    function _processGlobalApkUpdate(BN254.G1Point memory newGlobalApk) internal returns(bytes32){
        bytes32 globalApkHash = BN254.hashG1Point(newGlobalApk);
        globalApkUpdateList[globalApkUpdateList.length - 1].nextUpdateBlockNumber = uint32(block.number);
        _globalApk = newGlobalApk;

        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = globalApkHash;
        latestGlobalApkUpdate.updateBlockNumber = uint32(block.number);
        globalApkUpdateList.push(latestGlobalApkUpdate);

        return globalApkHash;
    }

    function _processQuorumApkUpdate(uint8[] memory quorumNumbers, BN254.G1Point memory pubkey, bool isRegistration) internal {
        BN254.G1Point memory latestApk;
        BN254.G1Point memory currentApk;
        for (uint i = 0; i < quorumNumbers.length; i++) {
            uint8 quorumNumber = quorumNumbers[i];
            
            currentApk = quorumToApk[quorumNumber];
            if(isRegistration){
                latestApk = BN254.plus(currentApk, pubkey);
            } else {
                latestApk = BN254.plus(currentApk, BN254.negate(pubkey));
            }

            //update aggregate public key for this quorum
            quorumToApk[quorumNumber] = latestApk;
            //update nextUpdateBlockNumber of the current latest ApkUpdate
            quorumToApkUpdates[quorumNumber][quorumToApkUpdates[quorumNumber].length - 1].nextUpdateBlockNumber = uint32(block.number);
            //create new ApkUpdate to add to the mapping
            ApkUpdate memory latestApkUpdate;
            latestApkUpdate.apkHash = BN254.hashG1Point(latestApk);
            latestApkUpdate.updateBlockNumber = uint32(block.number);
            quorumToApkUpdates[quorumNumber].push(latestApkUpdate);
        }
    }

    function _initializeApkUpdates() internal {
        BN254.G1Point memory pk = BN254.G1Point(0,0);
        _processGlobalApkUpdate(pk);

        for (uint8 quorumNumber = 0; quorumNumber < 256; quorumNumber++) {
            quorumToApk[quorumNumber] = pk;
            quorumToApkUpdates[quorumNumber].push(ApkUpdate({
                apkHash: BN254.hashG1Point(pk),
                updateBlockNumber: uint32(block.number),
                nextUpdateBlockNumber: 0
            }));
        }
    }
}