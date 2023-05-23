// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "../interfaces/IBLSPubkeyRegistry.sol";
import "../libraries/BN254.sol";
import "../interfaces/IServiceManager.sol";
import "./RegistryBase.sol";


contract BLSPubkeyRegistry is RegistryBase, IBLSPubkeyRegistry {

    BN254.G1Point globalApk;

    mapping(uint8 => ApkUpdate[]) public quorumToApkUpdates;
    
    mapping(uint8 => BN254.G1Point) public quorumToApk;


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
        BN254.G1Point currentApk;
        uint256 quorumToApkUpdatesLatestIndex = quorumToApkUpdates[quorumNumber].length - 1;
        for (uint8 quorumNumber = 0; quorumNumber < 256; i++) {
            if(quorumBitmap >> quorumNumber & 1 == 1){
                currentApk = quorumToApk[quorumNumber];
                latestApk = BN254.plus(currentApk, pubkey);
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
         globalApk = BN254.plus(globalApk, pubkey);
         bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        _addRegistrant(operator, pubkeyHash, quorumBitmap);
    }

    /**
     * @notice deregisters `operator` with the given `pubkey` for the quorums specified by `quorumBitmap`
     * @dev Permissioned by RegistryCoordinator
     */    
    function deregisterOperator(address operator, uint256 quorumBitmap, BN254.G1Point memory pubkey, uint32 index) external returns(bytes32){
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);
        BN254.G1Point currentApk;
        for (uint8 quorumNumber = 0; i < 256; i++) {
            if(quorumBitmap >> quorumNumber & 1 == 1){
                currentApk = quorumToApk[quorumNumber];
                latestApk = BN254.plus(currentApk, BN254.negate(pubkey));

                quorumToApk = latestApk;



            }
            
        }

        _removeOperator(operator, pubkeyHash, index)
        globalApk = BN254.plus(globalApk, BN254.negate(pubkey));
    }

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates for the `quorumNumber`
    function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) external view returns (ApkUpdate memory);

    /// @notice returns the `ApkUpdate` struct at `index` in the list of APK updates that keep track of the sum of the APK amonng all quorums
    function globalApkUpdates(uint256 index) external view returns (ApkUpdate memory);

    /**
     * @notice get hash of the apk of `quorumNumber` at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32);

	/**
     * @notice get hash of the apk among all quourms at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
     */
    function getGlobalApkHashAtBlockNumberFromIndex(uint32 blockNumber, uint256 index) external view returns (bytes32);
}