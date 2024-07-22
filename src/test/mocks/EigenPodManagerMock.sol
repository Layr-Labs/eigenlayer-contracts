// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPodManager.sol";
import "../../contracts/permissions/Pausable.sol";


contract EigenPodManagerMock is IEigenPodManager, Test, Pausable {
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
    IBeacon public eigenPodBeacon;
    IETHPOSDeposit public ethPOS;

    mapping(address => int256) public podShares;

    constructor(IPauserRegistry _pauserRegistry) {
        _initializePauser(_pauserRegistry, 0);
    }

    function slasher() external view returns(ISlasher) {}

    function createPod() external returns(address) {}

    function stake(bytes calldata /*pubkey*/, bytes calldata /*signature*/, bytes32 /*depositDataRoot*/) external payable {}

    function recordBeaconChainETHBalanceUpdate(address /*podOwner*/, int256 /*sharesDelta*/) external pure {}

    function ownerToPod(address /*podOwner*/) external pure returns(IEigenPod) {
        return IEigenPod(address(0));
    }

    function getPod(address podOwner) external pure returns(IEigenPod) {
        return IEigenPod(podOwner);
    }

    function strategyManager() external pure returns(IStrategyManager) {
        return IStrategyManager(address(0));
    }
    
    function hasPod(address /*podOwner*/) external pure returns (bool) {
        return false;
    }

    function podOwnerShares(address podOwner) external view returns (int256) {
        return podShares[podOwner];
    }

    function setPodOwnerShares(address podOwner, int256 shares) external {
        podShares[podOwner] = shares;
    }

    function addShares(address /*podOwner*/, uint256 shares) external pure returns (uint256) {
        // this is the "increase in delegateable tokens"
        return (shares);
    }

    function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) external {}

    function removeShares(address podOwner, uint256 shares) external {}

    function numPods() external view returns (uint256) {}

    function updateStaleValidatorCount(address, int256) external {}

    function denebForkTimestamp() external pure returns (uint64) {
        return type(uint64).max;
    }

    function setDenebForkTimestamp(uint64 timestamp) external{}
}