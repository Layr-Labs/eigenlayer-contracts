// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../libraries/SlashingLib.sol";
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
 * - keeping track of the restaked balances of all EigenPod owners
 * - withdrawing eth when withdrawals are completed
 */
contract EigenPodManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    EigenPodPausingConstants,
    EigenPodManagerStorage,
    ReentrancyGuardUpgradeable
{
    using SlashingLib for *;

    modifier onlyEigenPod(
        address podOwner
    ) {
        require(address(ownerToPod[podOwner]) == msg.sender, OnlyEigenPod());
        _;
    }

    modifier onlyDelegationManager() {
        require(msg.sender == address(delegationManager), OnlyDelegationManager());
        _;
    }

    constructor(
        IETHPOSDeposit _ethPOS,
        IBeacon _eigenPodBeacon,
        IStrategyManager _strategyManager,
        IDelegationManager _delegationManager,
        IPauserRegistry _pauserRegistry
    )
        EigenPodManagerStorage(_ethPOS, _eigenPodBeacon, _strategyManager, _delegationManager)
        Pausable(_pauserRegistry)
    {
        _disableInitializers();
    }

    function initialize(address initialOwner, uint256 _initPausedStatus) external initializer {
        _transferOwnership(initialOwner);
        _setPausedStatus(_initPausedStatus);
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     * @dev Returns EigenPod address
     */
    function createPod() external onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) returns (address) {
        require(!hasPod(msg.sender), EigenPodAlreadyExists());
        // deploy a pod if the sender doesn't have one already
        IEigenPod pod = _deployPod();

        return address(pod);
    }

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod.
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) {
        IEigenPod pod = ownerToPod[msg.sender];
        if (address(pod) == address(0)) {
            //deploy a pod if the sender doesn't have one already
            pod = _deployPod();
        }
        pod.stake{value: msg.value}(pubkey, signature, depositDataRoot);
    }

    /**
     * @notice Changes the `podOwner`'s shares by `sharesDelta` and performs a call to the DelegationManager
     * to ensure that delegated shares are also tracked correctly
     * @param podOwner is the pod owner whose balance is being updated.
     * @param sharesDelta is the change in podOwner's beaconChainETHStrategy shares
     * @param proportionOfOldBalance is the proportion (of WAD) of the podOwner's previous balance before the delta
     * @dev Callable only by the podOwner's EigenPod contract.
     * @dev Reverts if `sharesDelta` is not a whole Gwei amount
     */
    function recordBeaconChainETHBalanceUpdate(
        address podOwner,
        int256 sharesDelta,
        uint64 proportionOfOldBalance
    ) external onlyEigenPod(podOwner) nonReentrant {
        require(podOwner != address(0), InputAddressZero());
        require(sharesDelta % int256(GWEI_TO_WEI) == 0, SharesNotMultipleOfGwei());
        // shares can only be negative if they were due to negative shareDeltas after queued withdrawals in before
        // the slashing upgrade. Make people complete queued withdrawals before completing any further checkpoints.
        // the only effects podOwner UX, not AVS UX, since the podOwner already has 0 shares in the DM if they
        // have a negative shares in EPM.
        require(podOwnerDepositShares[podOwner] >= 0, LegacyWithdrawalsNotCompleted());

        if (sharesDelta > 0) {
            (uint256 existingDepositShares, uint256 addedShares) = _addShares(podOwner, uint256(sharesDelta));
            delegationManager.increaseDelegatedShares({
                staker: podOwner,
                strategy: beaconChainETHStrategy,
                existingDepositShares: existingDepositShares,
                addedShares: addedShares
            });
        } else if (sharesDelta < 0 && podOwnerDepositShares[podOwner] > 0) {
            delegationManager.decreaseBeaconChainScalingFactor(
                podOwner, uint256(podOwnerDepositShares[podOwner]), proportionOfOldBalance
            );
        }
    }

    /**
     * @notice Used by the DelegationManager to remove a pod owner's deposit shares when they enter the withdrawal queue.
     * Simply decreases the `podOwner`'s shares by `shares`, down to a minimum of zero.
     * @dev This function reverts if it would result in `podOwnerDepositShares[podOwner]` being less than zero, i.e. it is forbidden for this function to
     * result in the `podOwner` incurring a "share deficit". This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
     * shares from the operator to whom the staker is delegated.
     * @dev The delegation manager validates that the podOwner is not address(0)
     */
    function removeDepositShares(
        address staker,
        IStrategy strategy,
        uint256 depositSharesToRemove
    ) external onlyDelegationManager {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        int256 updatedShares = podOwnerDepositShares[staker] - int256(depositSharesToRemove);
        require(updatedShares >= 0, SharesNegative());
        podOwnerDepositShares[staker] = updatedShares;

        emit NewTotalShares(staker, updatedShares);
    }

    /**
     * @notice Increases the `podOwner`'s shares by `shares`, paying off deficit if possible.
     * Used by the DelegationManager to award a pod owner shares on exiting the withdrawal queue
     * @dev Reverts if `shares` is not a whole Gwei amount
     * @return existingDepositShares the pod owner's shares prior to any additions. Returns 0 if negative
     * @return addedShares the number of shares added to the staker's balance above 0. This means that if,
     * after shares are added, the staker's balance is non-positive, this will return 0.
     */
    function addShares(
        address staker,
        IStrategy strategy,
        IERC20,
        uint256 shares
    ) external onlyDelegationManager returns (uint256, uint256) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        return _addShares(staker, shares);
    }

    /**
     * @notice Used by the DelegationManager to complete a withdrawal, sending tokens to some destination address
     * @dev Prioritizes decreasing the podOwner's share deficit, if they have one
     * @dev Reverts if `shares` is not a whole Gwei amount
     * @dev This function assumes that `removeShares` has already been called by the delegationManager, hence why
     *      we do not need to update the podOwnerDepositShares if `currentpodOwnerDepositShares` is positive
     */
    function withdrawSharesAsTokens(
        address staker,
        IStrategy strategy,
        IERC20,
        uint256 shares
    ) external onlyDelegationManager {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        require(staker != address(0), InputAddressZero());
        require(int256(shares) > 0, SharesNegative());

        int256 currentDepositShares = podOwnerDepositShares[staker];
        uint256 sharesToWithdraw;
        // if there is an existing shares deficit, prioritize decreasing the deficit first
        // this is an M2 legacy codepath. TODO: gross
        if (currentDepositShares < 0) {
            uint256 currentDepositShareDeficit = uint256(-currentDepositShares);
            uint256 depositSharesToAdd;
            if (shares > currentDepositShareDeficit) {
                // get rid of the whole deficit if possible, and pass any remaining shares onto destination
                depositSharesToAdd = currentDepositShareDeficit;
                sharesToWithdraw = shares - currentDepositShareDeficit;
            } else {
                // otherwise get rid of as much deficit as possible, and return early, since there is nothing left over to forward on
                depositSharesToAdd = shares;
                sharesToWithdraw = 0;
            }

            int256 updatedShares = currentDepositShares + int256(depositSharesToAdd);
            podOwnerDepositShares[staker] = updatedShares;
            emit PodSharesUpdated(staker, int256(depositSharesToAdd));
            emit NewTotalShares(staker, updatedShares);
        }
        if (sharesToWithdraw > 0) {
            // Actually withdraw to the destination
            ownerToPod[staker].withdrawRestakedBeaconChainETH(staker, sharesToWithdraw);
        }
    }

    /// @inheritdoc IEigenPodManager
    function setPectraForkTimestamp(uint64 newPectraForkTimestamp) external onlyOwner {
        uint64 currentPectraForkTimestamp = getPectraForkTimestamp();
        require(currentPectraForkTimestamp == type(uint64).max, ForkTimestampAlreadySet());
        require(newPectraForkTimestamp != 0, InvalidForkTimestamp());

        _pectraForkTimestamp = newPectraForkTimestamp;
        emit PectraForkTimestampSet(newPectraForkTimestamp);
    }

    /// @inheritdoc IEigenPodManager
    function getPectraForkTimestamp() public view returns (uint64) {
        /// Initial value is 0, return type(uint64).max if not set
        if (_pectraForkTimestamp == 0) {
            return type(uint64).max;
        } else {
            return _pectraForkTimestamp;
        }
    }

    /// @notice Returns the current shares of `user` in `strategy`
    /// @dev strategy must be beaconChainETH when talking to the EigenPodManager
    /// @dev returns 0 if the user has negative shares
    function stakerDepositShares(address user, IStrategy strategy) public view returns (uint256 depositShares) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        return podOwnerDepositShares[user] < 0 ? 0 : uint256(podOwnerDepositShares[user]);
    }

    // INTERNAL FUNCTIONS

    function _deployPod() internal returns (IEigenPod) {
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

    /// @dev Adds the shares to the staker's balance, returning their current/added shares
    /// NOTE: if the staker ends with a non-positive balance, this returns (0, 0)
    /// @return existingDepositShares the shares the staker had before any were added
    /// @return addedShares the shares added to the staker's balance
    function _addShares(address staker, uint256 shares) internal returns (uint256, uint256) {
        require(staker != address(0), InputAddressZero());
        require(int256(shares) >= 0, SharesNegative());

        int256 sharesToAdd = int256(shares);
        int256 currentDepositShares = podOwnerDepositShares[staker];
        int256 updatedDepositShares = currentDepositShares + sharesToAdd;
        podOwnerDepositShares[staker] = updatedDepositShares;

        emit PodSharesUpdated(staker, sharesToAdd);
        emit NewTotalShares(staker, updatedDepositShares);

        // If we haven't added enough shares to go positive, return (0, 0)
        if (updatedDepositShares <= 0) {
            return (0, 0);
        }

        return (currentDepositShares < 0 ? 0 : uint256(currentDepositShares), shares);
    }

    // VIEW FUNCTIONS
    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(
        address podOwner
    ) public view returns (IEigenPod) {
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
    function hasPod(
        address podOwner
    ) public view returns (bool) {
        return address(ownerToPod[podOwner]) != address(0);
    }
}
