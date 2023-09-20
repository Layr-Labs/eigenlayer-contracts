// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IEigenPodManager.sol";

contract EigenPodManagerMock is IEigenPodManager, Test {
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    function slasher() external view returns(ISlasher) {}

    function createPod() external pure {}

    function stake(bytes calldata /*pubkey*/, bytes calldata /*signature*/, bytes32 /*depositDataRoot*/) external payable {}

    function restakeBeaconChainETH(address /*podOwner*/, uint256 /*amount*/) external pure {}

    function recordBeaconChainETHBalanceUpdate(address /*podOwner*/, int256 /*sharesDelta*/) external pure {}
    
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

    function getBlockRootAtTimestamp(uint64 /*timestamp*/) external pure returns(bytes32) {
        return bytes32(0);
    }

    function strategyManager() external pure returns(IStrategyManager) {
        return IStrategyManager(address(0));
    }
    
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

    function podOwnerShares(address podOwner) external returns (uint256){}

    function queueWithdrawal(uint256 amountWei, address withdrawer, bool undelegateIfPossible) external returns(bytes32) {}

    function forceIntoUndelegationLimbo(address podOwner) external {}

    function slashQueuedWithdrawal(address slashedFundsRecipient, BeaconChainQueuedWithdrawal memory queuedWithdrawal) external{}

    function slashShares(address slashedPodOwner, address slashedFundsRecipient, uint256 shareAmount) external{}

    function completeQueuedWithdrawal(BeaconChainQueuedWithdrawal memory queuedWithdrawal, uint256 middlewareTimesIndex) external{}

    // @notice Returns 'true' if `staker` can undelegate and false otherwise
    function stakerHasActiveShares(address /*staker*/) external pure returns (bool) {
        return true;
    }

    // @notice Returns 'true' if `staker` has removed all of their shares from delegation, either by queuing a withdrawal for them or by going into "undelegation limbo"
    function podOwnerHasNoDelegatedShares(address /*staker*/) external pure returns (bool) {
        return true;
    }

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.    
    function calculateWithdrawalRoot(BeaconChainQueuedWithdrawal memory queuedWithdrawal) external pure returns (bytes32) {}

    // @notice Getter function for the internal `_podOwnerUndelegationLimboStatus` mapping.
    function podOwnerUndelegationLimboStatus(address podOwner) external view returns (UndelegationLimboStatus memory) {}

    // @notice Getter function for `_podOwnerUndelegationLimboStatus.undelegationLimboActive`.
    function isInUndelegationLimbo(address podOwner) external view returns (bool) {}
}