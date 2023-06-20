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
    /// @notice the current aggregate pubkey of all operators registered in this contract, regardless of quorum
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
    function registerOperator(address operator, bytes memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        //calculate hash of the operator's pubkey
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        require(pubkeyHash != ZERO_PK_HASH, "BLSPubkeyRegistry.registerOperator: cannot register zero pubkey");
        //ensure that the operator owns their public key by referencing the BLSPubkeyCompendium
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
    function deregisterOperator(address operator, bytes memory quorumNumbers, BN254.G1Point memory pubkey) external onlyRegistryCoordinator returns(bytes32){
        bytes32 pubkeyHash = BN254.hashG1Point(pubkey);

        require(pubkeyCompendium.pubkeyHashToOperator(pubkeyHash) == operator,"BLSPubkeyRegistry.registerOperator: operator does not own pubkey");

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
     * @param blockNumber is the number of the block for which the latest ApkHash will be retrieved
     * @param index is the index of the apkUpdate being retrieved from the list of quorum apkUpdates in storage
     */
    function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) external view returns (bytes32){
        ApkUpdate memory quorumApkUpdate = quorumApkUpdates[quorumNumber][index];
        _validateApkHashForQuorumAtBlockNumber(quorumApkUpdate, blockNumber);
        return quorumApkUpdate.apkHash;
    }

	/**
     * @notice get hash of the global apk among all quorums at `blockNumber` using the provided `index`;
     * called by checkSignatures in BLSSignatureChecker.sol.
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

    function _processGlobalApkUpdate(BN254.G1Point memory point) internal {
        // load and store in memory in case we need to access the length again
        uint256 globalApkUpdatesLength = globalApkUpdates.length;
        // update the nextUpdateBlockNumber of the previous update
        if (globalApkUpdatesLength > 0) {
            globalApkUpdates[globalApkUpdatesLength - 1].nextUpdateBlockNumber = uint32(block.number);
        }

        // accumulate the given point into the globalApk
        globalApk = globalApk.plus(point);
        // add this update to the list of globalApkUpdates
        ApkUpdate memory latestGlobalApkUpdate;
        latestGlobalApkUpdate.apkHash = BN254.hashG1Point(globalApk);
        latestGlobalApkUpdate.updateBlockNumber = uint32(block.number);
        globalApkUpdates.push(latestGlobalApkUpdate);
    }

    function _processQuorumApkUpdate(bytes memory quorumNumbers, BN254.G1Point memory point) internal {
        BN254.G1Point memory apkAfterUpdate;

        for (uint i = 0; i < quorumNumbers.length;) {
            uint8 quorumNumber = uint8(quorumNumbers[i]);

            uint256 quorumApkUpdatesLength = quorumApkUpdates[quorumNumber].length;
            if (quorumApkUpdatesLength > 0) {
                // update nextUpdateBlockNumber of the current latest ApkUpdate
                quorumApkUpdates[quorumNumber][quorumApkUpdatesLength - 1].nextUpdateBlockNumber = uint32(block.number);
            }

            apkAfterUpdate = quorumApk[quorumNumber].plus(point);

            //update aggregate public key for this quorum
            quorumApk[quorumNumber] = apkAfterUpdate;
            //create new ApkUpdate to add to the mapping
            ApkUpdate memory latestApkUpdate;
            latestApkUpdate.apkHash = BN254.hashG1Point(apkAfterUpdate);
            latestApkUpdate.updateBlockNumber = uint32(block.number);
            quorumApkUpdates[quorumNumber].push(latestApkUpdate);

            unchecked{
                ++i;
            }
        }
    }

    function _validateApkHashForQuorumAtBlockNumber(ApkUpdate memory apkUpdate, uint32 blockNumber) internal pure {
        require(
            blockNumber >= apkUpdate.updateBlockNumber, 
            "BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: index too recent"
        );
        /**
        * if there is a next update, check that the blockNumber is before the next update or if 
        * there is no next update, then apkUpdate.nextUpdateBlockNumber is 0.
        */
        require(
            apkUpdate.nextUpdateBlockNumber == 0 || blockNumber < apkUpdate.nextUpdateBlockNumber, 
            "BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: not latest apk update"
        );
    }
}