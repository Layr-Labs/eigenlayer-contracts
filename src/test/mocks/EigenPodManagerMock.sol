// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPodManager.sol";

contract EigenPodManagerMock is IEigenPodManager, Test {
    function slasher() external view returns(ISlasher) {}

    function createPod() external pure {}

    function stake(bytes calldata /*pubkey*/, bytes calldata /*signature*/, bytes32 /*depositDataRoot*/) external payable {}

    function restakeBeaconChainETH(address /*podOwner*/, uint256 /*amount*/) external pure {}

    function recordBeaconChainETHBalanceUpdate(address /*podOwner*/, uint256 /*beaconChainETHStrategyIndex*/, int256 /*sharesDelta*/) external pure {}
    
    function withdrawRestakedBeaconChainETH(address /*podOwner*/, address /*recipient*/, uint256 /*amount*/) external pure {}

    function updateBeaconChainOracle(IBeaconChainOracle /*newBeaconChainOracle*/) external pure {}

    function ownerToPod(address /*podOwner*/) external pure returns(IEigenPod) {
        return IEigenPod(address(0));
    }

    function getPod(address podOwner) external pure returns(IEigenPod) {
        return IEigenPod(podOwner);
    }

    function beaconChainOracle() external pure returns(IBeaconChainOracle) {
        return IBeaconChainOracle(address(0));
    }   

    function getBeaconChainStateRoot(uint64 /*blockNumber*/) external pure returns(bytes32) {
        return bytes32(0);
    }

    function strategyManager() external pure returns(IStrategyManager) {
        return IStrategyManager(address(0));
    }

    function decrementWithdrawableRestakedExecutionLayerGwei(address podOwner, uint256 amountWei) external {}

    function incrementWithdrawableRestakedExecutionLayerGwei(address podOwner, uint256 amountWei) external {}


    function hasPod(address /*podOwner*/) external pure returns (bool) {
        return false;
    }

    function pause(uint256 /*newPausedStatus*/) external{}

    function pauseAll() external{}

    function paused() external pure returns (uint256) {
        return 0;
    }
    
    function paused(uint8 /*index*/) external pure returns (bool) {
        return false;
    }

    function setPauserRegistry(IPauserRegistry /*newPauserRegistry*/) external {}

    function pauserRegistry() external pure returns (IPauserRegistry) {
        return IPauserRegistry(address(0));
    }

    function unpause(uint256 /*newPausedStatus*/) external{}
}