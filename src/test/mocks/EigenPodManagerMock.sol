// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IStrategy.sol";
import "../../contracts/permissions/Pausable.sol";

contract EigenPodManagerMock is Test, Pausable {
    receive() external payable {}
    fallback() external payable {}

    mapping(address => int) public podOwnerDepositShares;

    mapping(address => uint) public podOwnerSharesWithdrawn;

    uint64 public pectraForkTimestamp;

    struct BeaconChainSlashingFactor {
        bool isSet;
        uint64 slashingFactor;
    }

    mapping(address => BeaconChainSlashingFactor) _beaconChainSlashingFactor;

    constructor(IPauserRegistry _pauserRegistry) Pausable(_pauserRegistry) {
        _setPausedStatus(0);
        pectraForkTimestamp = 1 hours * 12;
    }

    function podOwnerShares(address podOwner) external view returns (int) {
        return podOwnerDepositShares[podOwner];
    }

    function stakerDepositShares(address user, address) public view returns (uint depositShares) {
        return podOwnerDepositShares[user] < 0 ? 0 : uint(podOwnerDepositShares[user]);
    }

    function setPodOwnerShares(address podOwner, int shares) external {
        podOwnerDepositShares[podOwner] = shares;
    }

    function addShares(address podOwner, IStrategy, uint shares) external returns (uint, uint) {
        uint existingDepositShares = uint(podOwnerDepositShares[podOwner]);
        podOwnerDepositShares[podOwner] += int(shares);
        return (existingDepositShares, shares);
    }

    function removeDepositShares(
        address podOwner,
        IStrategy, // strategy
        uint shares
    ) external returns (uint) {
        int updatedShares = podOwnerDepositShares[podOwner] - int(shares);
        podOwnerDepositShares[podOwner] = updatedShares;
        return uint(updatedShares);
    }

    function denebForkTimestamp() external pure returns (uint64) {
        return type(uint64).max;
    }

    function withdrawSharesAsTokens(
        address podOwner,
        address,
        /**
         * strategy
         */
        address,
        /**
         * token
         */
        uint shares
    ) external {
        podOwnerSharesWithdrawn[podOwner] += shares;
    }

    function setBeaconChainSlashingFactor(address staker, uint64 bcsf) external {
        _beaconChainSlashingFactor[staker] = BeaconChainSlashingFactor({isSet: true, slashingFactor: bcsf});
    }

    function beaconChainSlashingFactor(address staker) external view returns (uint64) {
        BeaconChainSlashingFactor memory bsf = _beaconChainSlashingFactor[staker];
        return bsf.isSet ? bsf.slashingFactor : WAD;
    }

    function setPectraForkTimestamp(uint64 _pectraForkTimestamp) external {
        pectraForkTimestamp = _pectraForkTimestamp;
    }
}
