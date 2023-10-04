// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../interfaces/IBeaconChainOracle.sol";

import "../permissions/Pausable.sol";
import "./EigenPodPausingConstants.sol";
import "./EigenPodManagerStorage.sol";

/**
 * @title The contract used for creating and managing EigenPods
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice The main functionalities are:
 * - creating EigenPods
 * - staking for new validators on EigenPods
 * - keeping track of the balances of all validators of EigenPods, and their stake in EigenLayer
 * - withdrawing eth when withdrawals are initiated
 */
contract EigenPodManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    EigenPodPausingConstants,
    EigenPodManagerStorage,
    ReentrancyGuardUpgradeable
{
    
    modifier onlyEigenPod(address podOwner) {
        require(address(ownerToPod[podOwner]) == msg.sender, "EigenPodManager.onlyEigenPod: not a pod");
        _;
    }

    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), "EigenPodManager.onlyStrategyManager: not strategyManager");
        _;
    }

    modifier onlyDelegationManager() {
        require(
            msg.sender == address(delegationManager),
            "EigenPodManager.onlyDelegationManager: not the DelegationManager"
        );
        _;
    }

    modifier onlyNotFrozen(address staker) {
        require(
            !slasher.isFrozen(staker),
            "EigenPodManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"
        );
        _;
    }

    modifier onlyFrozen(address staker) {
        require(slasher.isFrozen(staker), "EigenPodManager.onlyFrozen: staker has not been frozen");
        _;
    }

    constructor(
        IETHPOSDeposit _ethPOS,
        IBeacon _eigenPodBeacon,
        IStrategyManager _strategyManager,
        ISlasher _slasher,
        IDelegationManager _delegationManager
    ) EigenPodManagerStorage(_ethPOS, _eigenPodBeacon, _strategyManager, _slasher, _delegationManager) {
        _disableInitializers();
    }

    function initialize(
        uint256 _maxPods,
        IBeaconChainOracle _beaconChainOracle,
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 _initPausedStatus
    ) external initializer {
        _setMaxPods(_maxPods);
        _updateBeaconChainOracle(_beaconChainOracle);
        _transferOwnership(initialOwner);
        _initializePauser(_pauserRegistry, _initPausedStatus);
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external {
        require(!hasPod(msg.sender), "EigenPodManager.createPod: Sender already has a pod");
        // deploy a pod if the sender doesn't have one already
        _deployPod();
    }

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod.
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable {
        IEigenPod pod = ownerToPod[msg.sender];
        if (address(pod) == address(0)) {
            //deploy a pod if the sender doesn't have one already
            pod = _deployPod();
        }
        pod.stake{value: msg.value}(pubkey, signature, depositDataRoot);
    }

    /**
     * @notice Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.
     * @param podOwner The owner of the pod whose balance must be deposited.
     * @param amountWei The amount of ETH to 'deposit' (i.e. be credited to the podOwner).
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function restakeBeaconChainETH(
        address podOwner,
        uint256 amountWei
    ) external onlyEigenPod(podOwner) onlyNotFrozen(podOwner) nonReentrant {
        _addShares(podOwner, amountWei);
        delegationManager.increaseDelegatedShares({
            staker: podOwner,
            strategy: beaconChainETHStrategy,
            shares: amountWei
        });
        emit BeaconChainETHDeposited(podOwner, amountWei);
    }

    /**
     * @notice Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
     *         balance of a validator is lower than how much stake they have committed to EigenLayer
     * @param podOwner is the pod owner whose balance is being updated.
     * @param sharesDelta is the change in podOwner's beaconChainETHStrategy shares
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function recordBeaconChainETHBalanceUpdate(
        address podOwner,
        int256 sharesDelta
    ) external onlyEigenPod(podOwner) nonReentrant {
        _recordBeaconChainETHBalanceUpdate(podOwner, sharesDelta);
    }

    function removeShares(
        address podOwner, 
        uint256 shares
    ) external onlyDelegationManager {
        /**
         * This decrements the withdrawableRestakedExecutionLayerGwei which is incremented only when a podOwner proves a full withdrawal.
         * Remember that withdrawableRestakedExecutionLayerGwei tracks the currently withdrawable ETH from the EigenPod.
         * By doing this, we ensure that the number of shares in EigenLayer matches the amount of withdrawable ETH in
         * the pod plus any ETH still staked on the beacon chain via other validators pointed to the pod. As a result, a validator
         * must complete a full withdrawal from the execution layer prior to queuing a withdrawal of 'beacon chain ETH shares'
         * via EigenLayer, since otherwise withdrawableRestakedExecutionLayerGwei will be 0.
         */
        ownerToPod[podOwner].decrementWithdrawableRestakedExecutionLayerGwei(shares);
        _removeShares(podOwner, shares);
    }

    function awardShares(
        address podOwner,
        uint256 shares
    ) external onlyDelegationManager {
        _addShares(podOwner, shares);
    }

    // TODO the 2 calls here can probably be combined?
    function withdrawSharesAsTokens(
        address podOwner, 
        address destination, 
        uint256 shares
    ) external onlyDelegationManager {
        // Actually withdraw to the destination
        ownerToPod[podOwner].withdrawRestakedBeaconChainETH(destination, shares);
    }

    /**
     * Sets the maximum number of pods that can be deployed
     * @param newMaxPods The new maximum number of pods that can be deployed
     * @dev Callable by the unpauser of this contract
     */
    function setMaxPods(uint256 newMaxPods) external onlyUnpauser {
        _setMaxPods(newMaxPods);
    }

    /**
     * @notice Updates the oracle contract that provides the beacon chain state root
     * @param newBeaconChainOracle is the new oracle contract being pointed to
     * @dev Callable only by the owner of this contract (i.e. governance)
     */
    function updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) external onlyOwner {
        _updateBeaconChainOracle(newBeaconChainOracle);
    }

    // INTERNAL FUNCTIONS

    function _deployPod() internal onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) returns (IEigenPod) {
        // check that the limit of EigenPods has not been hit, and increment the EigenPod count
        require(numPods + 1 <= maxPods, "EigenPodManager._deployPod: pod limit reached");
        ++numPods;
        // create the pod
        IEigenPod pod = IEigenPod(
            Create2.deploy(
                0,
                bytes32(uint256(uint160(msg.sender))),
                // set the beacon address to the eigenPodBeacon and initialize it
                abi.encodePacked(beaconProxyBytecode, abi.encode(eigenPodBeacon, ""))
            )
        );
        pod.initialize(msg.sender);
        // store the pod in the mapping
        ownerToPod[msg.sender] = pod;
        emit PodDeployed(address(pod), msg.sender);
        return pod;
    }

    /// @notice Internal setter for `beaconChainOracle` that also emits an event
    function _updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) internal {
        beaconChainOracle = newBeaconChainOracle;
        emit BeaconOracleUpdated(address(newBeaconChainOracle));
    }

    /// @notice Internal setter for `maxPods` that also emits an event
    function _setMaxPods(uint256 _maxPods) internal {
        emit MaxPodsUpdated(maxPods, _maxPods);
        maxPods = _maxPods;
    }

    // @notice Increases the `podOwner`'s shares by `shareAmount`
    function _addShares(address podOwner, uint256 shareAmount) internal {
        require(podOwner != address(0), "EigenPodManager._addShares: podOwner cannot be zero address");
        require(shareAmount > 0, "EigenPodManager._addShares: amount must be greater than zero");

        podOwnerShares[podOwner] += shareAmount;
    }

    // @notice Reduces the `podOwner`'s shares by `shareAmount` and performs a call to the DelegationManager to ensure delegated shares are also tracked correctly
    function _removeShares(address podOwner, uint256 shareAmount) internal {
        require(podOwner != address(0), "EigenPodManager._removeShares: depositor cannot be zero address");
        require(shareAmount != 0, "EigenPodManager._removeShares: shareAmount should not be zero!");

        uint256 currentPodOwnerShares = podOwnerShares[podOwner];
        require(shareAmount <= currentPodOwnerShares, "EigenPodManager._removeShares: shareAmount too high");

        unchecked {
            currentPodOwnerShares = currentPodOwnerShares - shareAmount;
        }

        podOwnerShares[podOwner] = currentPodOwnerShares;
    }

    // @notice Changes the `podOwner`'s shares by `sharesDelta` and performs a call to the DelegationManager to ensure delegated shares are also tracked correctly
    function _recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) internal {
        if (sharesDelta < 0) {
            // if change in shares is negative, remove the shares (and don't bother trying to undelegate the `podOwner`)
            uint256 toRemove = uint256(-sharesDelta);
            _removeShares(podOwner, toRemove);

            IStrategy[] memory strategies = new IStrategy[](1);
            uint[] memory shares = new uint[](1);
            strategies[0] = beaconChainETHStrategy;
            shares[0] = toRemove;

            delegationManager.decreaseDelegatedShares({
                staker: podOwner,
                strategies: strategies,
                shares: shares
            });
        } else {
            // if change in shares is positive, add the shares
            uint256 toAdd = uint256(sharesDelta);
            _addShares(podOwner, toAdd);
            delegationManager.increaseDelegatedShares({
                staker: podOwner,
                strategy: beaconChainETHStrategy,
                shares: toAdd
            });
        }
    }

    // VIEW FUNCTIONS
    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) public view returns (IEigenPod) {
        IEigenPod pod = ownerToPod[podOwner];
        // if pod does not exist already, calculate what its address *will be* once it is deployed
        if (address(pod) == address(0)) {
            pod = IEigenPod(
                Create2.computeAddress(
                    bytes32(uint256(uint160(podOwner))), //salt
                    keccak256(abi.encodePacked(beaconProxyBytecode, abi.encode(eigenPodBeacon, ""))) //bytecode
                )
            );
        }
        return pod;
    }

    /// @notice Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.
    function hasPod(address podOwner) public view returns (bool) {
        return address(ownerToPod[podOwner]) != address(0);
    }

    /// @notice Returns the Beacon block root at `timestamp`. Reverts if the Beacon block root at `timestamp` has not yet been finalized.
    function getBlockRootAtTimestamp(uint64 timestamp) external view returns (bytes32) {
        bytes32 stateRoot = beaconChainOracle.timestampToBlockRoot(timestamp);
        require(
            stateRoot != bytes32(0),
            "EigenPodManager.getBlockRootAtTimestamp: state root at timestamp not yet finalized"
        );
        return stateRoot;
    }

    /**
     * @notice Returns 'true' if the pod owner has shares
     */
    function podOwnerHasActiveShares(address staker) public view returns (bool) {
        return podOwnerShares[staker] != 0;
    }
}
