// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../libraries/BN254.sol";
import "../interfaces/IServiceManager.sol";
import "./RegistryBase.sol";


contract BLSPubkeyRegistry is RegistryBase, IBLSPubkeyRegistry {

    BN254.G1Point globalApk;

    ApkUpdate[] public globalApkUpdates;

    mapping(uint8 => ApkUpdate[]) public quorumToApkUpdates;
    
    mapping(uint8 => BN254.G1Point) public quorumToApk;

    event RegistrationEvent(
        address indexed operator,
        bytes32 pubkeyHash,
        uint256 quorumBitmap,
        bytes32 globalApkHash;
    )

    event Operator


    constructor(
        IStrategyManager strategyManager,
        IServiceManager serviceManager
    )RegistryBase(strategyManager, serviceManager){
    }

    function initialize(
        uint96[] memory _minimumStakeForQuorum,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) public initializer {
        RegistryBase._initialize(_minimumStakeForQuorum, _quorumStrategiesConsideredAndMultipliers);
    }

    function registerOperator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubkey) external returns(bytes32){
        uint256 quorumToApkUpdatesLatestIndex = quorumToApkUpdates[quorumNumber].length - 1;
        for (uint8 quorumNumber = 0; quorumNumber < 256; i++) {
            if(quorumBitmap >> quorumNumber & 1 == 1){
                _processQuorumApkUpdate(quorumNumber, quorumToApkUpdatesLatestIndex, true);
            }
        }
         globalApk = BN254.plus(globalApk, pubkey);
        bytes32 globalApkHash = _processApkUpdate(globalApk);


         bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        _addRegistrant(operator, pubkeyHash, quorumBitmap);

        emit RegistrationEvent(operator, pubkeyHash, quorumBitmap, globalApkHash);
    }

    /**
     * @notice deregisters `operator` with the given `pubkey` for the quorums specified by `quorumBitmap`
     * @dev Permissioned by RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubkey, uint32 index) external returns(bytes32){
        uint256 quorumToApkUpdatesLatestIndex = quorumToApkUpdates[quorumNumber].length - 1;
        for (uint8 quorumNumber = 0; i < 256; i++) {
            if(quorumBitmap >> quorumNumber & 1 == 1){
                _processQuorumApkUpdate(quorumNumber, quorumToApkUpdatesLatestIndex, false);
            }
            
        }
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        _removeOperator(operator, pubkeyHash, index)
        globalApk = BN254.plus(globalApk, BN254.negate(pubkey));
        bytes32 globalApkHash = _processApkUpdate(globalApk);

        emit RegistrationEvent(operator, pubkeyHash, quorumBitmap, globalApkHash);
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
    }

	/**
     * @notice get hash of the apk among all quourms at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32){

    }


    function _processGlobalApkUpdate(BN254.G1Point globalApk) internal returns(bytes32){
        globalApkHash = BN254.hashG1Point(globalApk);
        globalApkUpdates[globalApkUpdates.length - 1].nextUpdateBlock = block.number

        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = globalApkHash;
        latestGlobalApkUpdate.updateBlockNumber = block.number;
        globalApkUpdates.push(latestGlobalApkUpdate);

        return globalApkHash;
    }

    function _processQuorumApkUpdate(uint8 quorumNumber, uint256 quorumToApkUpdatesLatestIndex, bool isRegistration) internal {
        BN254.G1Point currentApk = quorumToApk[quorumNumber];
        if(isRegistration){
            latestApk = BN254.plus(currentApk, pubkey);
        } else {
            latestApk = BN254.plus(currentApk, BN254.negate(pubkey));
        }

        //update aggregate public key for this quorum
        quorumToApk[quorumNumber] = latestApk;
        //update nextUpdateBlockNumber of the current latest ApkUpdate
        quorumToApkUpdates[quorumNumber][quorumToApkUpdatesLatestIndex].nextUpdateBlockNumber = block.number;
        //create new ApkUpdate to add to the mapping
        ApkUpdate memory latestApkUpdate;
        latestApkUpdate.apkHash = BN254.hashG1Point(latestApk);
        latestApkUpdate.updateBlockNumber = block.number;
        quorumToApkUpdates[quorumNumber].push(latestApkUpdate);
    }
}