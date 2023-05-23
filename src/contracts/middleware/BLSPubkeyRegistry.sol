// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../libraries/BN254.sol";
import "../interfaces/IServiceManager.sol";


contract BLSPubkeyRegistry is IBLSPubkeyRegistry {

    BN254.G1Point globalApk;

    IRegistryCoordinator public registryCoordinator;

    ApkUpdate[] public globalApkUpdates;

    mapping(uint8 => ApkUpdate[]) public quorumToApkUpdates;
    
    mapping(uint8 => BN254.G1Point) public quorumToApk;

    event RegistrationEvent(
        address indexed operator,
        bytes32 pubkeyHash,
        uint256 quorumBitmap,
        bytes32 globalApkHash
    );

    event DeregistrationEvent(
        address indexed operator,
        bytes32 pubkeyHash,
        uint256 quorumBitmap,
        bytes32 globalApkHash
    );

    modifier onlyRegistryCoordinator() {
        require(msg.sender == address(registryCoordinator), "BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
        _;
    }


    constructor(
        IRegistryCoordinator _registryCoordinator
    ){
        registryCoordinator = _registryCoordinator;
    }


    function registerOperator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        _processQuorumApkUpdate(quorumBitmap, pubkey, true);
        bytes32 globalApkHash = _processGlobalApkUpdate(BN254.plus(globalApk, pubkey));


         bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        emit RegistrationEvent(operator, pubkeyHash, quorumBitmap, globalApkHash);
    }

    /**
     * @notice deregisters `operator` with the given `pubkey` for the quorums specified by `quorumBitmap`
     * @dev Permissioned by RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubkey, uint32 index) external onlyRegistryCoordinator returns(bytes32){
        _processQuorumApkUpdate(quorumBitmap, pubkey, false);

        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        bytes32 globalApkHash = _processGlobalApkUpdate(BN254.plus(globalApk, BN254.negate(pubkey)));

        emit DeregistrationEvent(operator, pubkeyHash, quorumBitmap, globalApkHash);
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory){
        return quorumToApkUpdates[quorumNumber][index];
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates that keep track of the sum of the APK amonng all quorums
    function globalApkUpdates(uint256 index) external view returns (ApkUpdate memory){
        return globalApkUpdates[index];
    }

    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        require(blockNumber >= quorumToApkUpdates[quorumNumber][index].updateBlockNumber, "BLSPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex: index too recent");

        if (index != quorumToApkUpdates[quorumNumber].length - 1){
            require(blockNumber < quorumToApkUpdates[quorumNumber][index].nextUpdateBlockNumber, "BLSPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex: not latest apk update");
        }

        return quorumToApkUpdates[quorumNumber][index].apkHash;
    }

	/**
     * @notice get hash of the apk among all quourums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){
        require(blockNumber >= globalApkUpdates[index].updateBlockNumber, "BLSPubkeyRegistry.getGlobalApkHashAtBlockNumberFromIndex: blockNumber too recent");

        if (index != globalApkUpdates.length - 1){
            require(blockNumber < globalApkUpdates[index].nextUpdateBlockNumber, "getGlobalApkHashAtBlockNumberFromIndex.getGlobalApkHashAtBlockNumberFromIndex: not latest apk update");
        }

        return globalApkUpdates[index].apkHash;
    }


    function _processGlobalApkUpdate(BN254.G1Point memory newGlobalApk) internal returns(bytes32){
        globalApk = newGlobalApk;
        bytes32 globalApkHash = BN254.hashG1Point(globalApk);
        globalApkUpdates[globalApkUpdates.length - 1].nextUpdateBlockNumber = uint32(block.number);

        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = globalApkHash;
        latestGlobalApkUpdate.updateBlockNumber = uint32(block.number);
        globalApkUpdates.push(latestGlobalApkUpdate);

        return globalApkHash;
    }

    function _processQuorumApkUpdate(uint256 quorumBitmap, BN254.G1Point memory pubkey, bool isRegistration) internal {
        BN254.G1Point memory latestApk;
        for (uint8 quorumNumber = 0; quorumNumber < 256; quorumNumber++) {
            if(quorumBitmap >> quorumNumber & 1 == 1){
                BN254.G1Point memory currentApk = quorumToApk[quorumNumber];
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