// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./IStrategyManager.sol";
import "./IEigenPod.sol";
import "./IBeaconChainOracle.sol";
import "src/contracts/interfaces/IPausable.sol";

/**
 * @notice DEPRECATED INTERFACE at commit hash https://github.com/Layr-Labs/eigenlayer-contracts/tree/0139d6213927c0a7812578899ddd3dda58051928
 * @title Interface for factory that creates and manages solo staking pods that have their withdrawal credentials pointed to EigenLayer.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */

interface IEigenPodManager_DeprecatedM1 is IPausable {
    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external;

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod. 
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.
     * @param podOwner The owner of the pod whose balance must be deposited.
     * @param amount The amount of ETH to 'deposit' (i.e. be credited to the podOwner).
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function restakeBeaconChainETH(address podOwner, uint256 amount) external;

    /**
     * @notice Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
     *         balance of a validator is lower than how much stake they have committed to EigenLayer
     * @param podOwner The owner of the pod whose balance must be removed.
     * @param beaconChainETHStrategyIndex is the index of the beaconChainETHStrategy for the pod owner for the callback to 
     *                                    the StrategyManager in case it must be removed from the list of the podOwner's strategies
     * @param amount The amount of ETH to remove.
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function recordOvercommittedBeaconChainETH(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 amount) external;
    
    /**
     * @notice Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.
     * @param podOwner The owner of the pod whose balance must be withdrawn.
     * @param recipient The recipient of the withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the StrategyManager contract.
     */
    function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount) external;

    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of this contract (i.e. governance)
     */
    function updateBeaconChainOracle(IBeaconChainOracle_DeprecatedM1 newBeaconChainOracle) external;

    /// @notice Returns the address of the `podOwner`'s EigenPod if it has been deployed.
    function ownerToPod(address podOwner) external view returns(IEigenPod_DeprecatedM1);

    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) external view returns(IEigenPod_DeprecatedM1);

    /// @notice Oracle contract that provides updates to the beacon chain's state
    function beaconChainOracle() external view returns(IBeaconChainOracle_DeprecatedM1);    

    /// @notice Returns the Beacon Chain state root at `blockNumber`. Reverts if the Beacon Chain state root at `blockNumber` has not yet been finalized.
    function getBeaconChainStateRoot(uint64 blockNumber) external view returns(bytes32);

    /// @notice EigenLayer's StrategyManager contract
    function strategyManager() external view returns(IStrategyManager_DeprecatedM1);

    /// @notice EigenLayer's Slasher contract
    function slasher() external view returns(ISlasher);

    function hasPod(address podOwner) external view returns (bool);
}
