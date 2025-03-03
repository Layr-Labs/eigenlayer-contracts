// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

import "../libraries/SlashingLib.sol";
import "../mixins/SemVerMixin.sol";
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
    ReentrancyGuardUpgradeable,
    SemVerMixin
{
    using SlashingLib for *;
    using Math for *;

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
        IDelegationManager _delegationManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        EigenPodManagerStorage(_ethPOS, _eigenPodBeacon, _delegationManager)
        Pausable(_pauserRegistry)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    function initialize(address initialOwner, uint256 _initPausedStatus) external initializer {
        _transferOwnership(initialOwner);
        _setPausedStatus(_initPausedStatus);
    }

    /// @inheritdoc IEigenPodManager
    function createPod() external onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) nonReentrant returns (address) {
        require(!hasPod(msg.sender), EigenPodAlreadyExists());
        // deploy a pod if the sender doesn't have one already
        IEigenPod pod = _deployPod();

        return address(pod);
    }

    /// @inheritdoc IEigenPodManager
    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) nonReentrant {
        IEigenPod pod = ownerToPod[msg.sender];
        if (address(pod) == address(0)) {
            //deploy a pod if the sender doesn't have one already
            pod = _deployPod();
        }
        pod.stake{value: msg.value}(pubkey, signature, depositDataRoot);
    }

    /// @inheritdoc IEigenPodManager
    function recordBeaconChainETHBalanceUpdate(
        address podOwner,
        uint256 prevRestakedBalanceWei,
        int256 balanceDeltaWei
    ) external onlyEigenPod(podOwner) nonReentrant {
        require(podOwner != address(0), InputAddressZero());
        require(balanceDeltaWei % int256(GWEI_TO_WEI) == 0, SharesNotMultipleOfGwei());
        // Negative shares only exist in certain cases where, prior to the slashing release, negative balance
        // deltas were reported after a pod owner queued a withdrawal for all their shares.
        //
        // The new system treats negative balance deltas differently, decreasing the pod owner's slashing factor
        // proportional to the decrease. This check was added to ensure the new system does not need to handle
        // negative shares - instead, stakers will need to go complete any existing withdrawals before their pod
        // can process a balance update.
        int256 currentDepositShares = podOwnerDepositShares[podOwner];
        require(currentDepositShares >= 0, LegacyWithdrawalsNotCompleted());

        // Shares are only added to the pod owner's balance when `balanceDeltaWei` > 0. When a pod reports
        // a negative balance delta, the pod owner's beacon chain slashing factor is decreased, devaluing
        // their shares. If the delta is zero, then no action needs to be taken.
        if (balanceDeltaWei > 0) {
            (uint256 prevDepositShares, uint256 addedShares) = _addShares(podOwner, uint256(balanceDeltaWei));

            // Update operator shares
            delegationManager.increaseDelegatedShares({
                staker: podOwner,
                strategy: beaconChainETHStrategy,
                prevDepositShares: prevDepositShares,
                addedShares: addedShares
            });
        } else if (balanceDeltaWei < 0) {
            uint64 beaconChainSlashingFactorDecrease = _reduceSlashingFactor({
                podOwner: podOwner,
                prevRestakedBalanceWei: prevRestakedBalanceWei,
                balanceDecreasedWei: uint256(-balanceDeltaWei)
            });

            // Update operator shares
            delegationManager.decreaseDelegatedShares({
                staker: podOwner,
                curDepositShares: uint256(currentDepositShares),
                beaconChainSlashingFactorDecrease: beaconChainSlashingFactorDecrease
            });
        }
    }

    /**
     * @notice Used by the DelegationManager to remove a pod owner's deposit shares when they enter the withdrawal queue.
     * Simply decreases the `podOwner`'s shares by `shares`, down to a minimum of zero.
     * @dev This function reverts if it would result in `podOwnerDepositShares[podOwner]` being less than zero, i.e. it is forbidden for this function to
     * result in the `podOwner` incurring a "share deficit". This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
     * shares from the operator to whom the staker is delegated.
     * @dev The delegation manager validates that the podOwner is not address(0)
     * @return updatedShares the staker's deposit shares after decrement
     */
    function removeDepositShares(
        address staker,
        IStrategy strategy,
        uint256 depositSharesToRemove
    ) external onlyDelegationManager nonReentrant returns (uint256) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        int256 updatedShares = podOwnerDepositShares[staker] - int256(depositSharesToRemove);
        require(updatedShares >= 0, SharesNegative());
        podOwnerDepositShares[staker] = updatedShares;

        emit NewTotalShares(staker, updatedShares);
        return uint256(updatedShares);
    }

    /**
     * @notice Increases the `podOwner`'s shares by `shares`, paying off negative shares if needed.
     * Used by the DelegationManager to award a pod owner shares on exiting the withdrawal queue
     * @return existingDepositShares the pod owner's shares prior to any additions. Returns 0 if negative
     * @return addedShares the number of shares added to the staker's balance above 0. This means that if,
     * after shares are added, the staker's balance is non-positive, this will return 0.
     */
    function addShares(
        address staker,
        IStrategy strategy,
        uint256 shares
    ) external onlyDelegationManager nonReentrant returns (uint256, uint256) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        return _addShares(staker, shares);
    }

    /**
     * @notice Used by the DelegationManager to complete a withdrawal, sending tokens to the pod owner
     * @dev Prioritizes decreasing the podOwner's share deficit, if they have one
     * @dev This function assumes that `removeShares` has already been called by the delegationManager, hence why
     *      we do not need to update the podOwnerDepositShares if `currentpodOwnerDepositShares` is positive
     */
    function withdrawSharesAsTokens(
        address staker,
        IStrategy strategy,
        IERC20,
        uint256 shares
    ) external onlyDelegationManager nonReentrant {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        require(staker != address(0), InputAddressZero());
        require(int256(shares) > 0, SharesNegative());

        int256 currentDepositShares = podOwnerDepositShares[staker];
        uint256 sharesToWithdraw = shares;

        // Negative shares only exist in certain cases where, prior to the slashing release, negative balance
        // deltas were reported after a pod owner queued a withdrawal for all their shares.
        //
        // The new system treats negative balance deltas differently, decreasing the pod owner's slashing factor
        // proportional to the decrease. This legacy codepath handles completion of withdrawals queued before
        // the slashing release.
        if (currentDepositShares < 0) {
            uint256 currentDepositShareDeficit = uint256(-currentDepositShares);
            uint256 depositSharesToAdd;

            if (shares > currentDepositShareDeficit) {
                // Get rid of the whole deficit and withdraw any remaining shares
                depositSharesToAdd = currentDepositShareDeficit;
                sharesToWithdraw = shares - currentDepositShareDeficit;
            } else {
                // Get rid of as much deficit as possible and don't withdraw any shares
                depositSharesToAdd = shares;
                sharesToWithdraw = 0;
            }

            int256 updatedShares = currentDepositShares + int256(depositSharesToAdd);
            podOwnerDepositShares[staker] = updatedShares;
            emit PodSharesUpdated(staker, int256(depositSharesToAdd));
            emit NewTotalShares(staker, updatedShares);
        }

        // Withdraw ETH from EigenPod
        if (sharesToWithdraw > 0) {
            ownerToPod[staker].withdrawRestakedBeaconChainETH(staker, sharesToWithdraw);
        }
    }

    /// @inheritdoc IShareManager
    function increaseBurnableShares(IStrategy, uint256 addedSharesToBurn) external onlyDelegationManager nonReentrant {
        burnableETHShares += addedSharesToBurn;
        emit BurnableETHSharesIncreased(addedSharesToBurn);
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
    /// @return prevDepositShares the shares the staker had before any were added
    /// @return addedShares the shares added to the staker's balance
    function _addShares(address staker, uint256 shares) internal returns (uint256, uint256) {
        require(staker != address(0), InputAddressZero());
        require(int256(shares) >= 0, SharesNegative());

        int256 sharesToAdd = int256(shares);
        int256 prevDepositShares = podOwnerDepositShares[staker];
        int256 updatedDepositShares = prevDepositShares + sharesToAdd;
        podOwnerDepositShares[staker] = updatedDepositShares;

        emit PodSharesUpdated(staker, sharesToAdd);
        emit NewTotalShares(staker, updatedDepositShares);

        // If we haven't added enough shares to go positive, return (0, 0)
        if (updatedDepositShares <= 0) {
            return (0, 0);
        }
        // If we have gone from negative to positive shares, return (0, positive delta)
        else if (prevDepositShares < 0) {
            return (0, uint256(updatedDepositShares));
        }
        // Else, return true previous shares and added shares
        else {
            return (uint256(prevDepositShares), shares);
        }
    }

    /// @dev Calculates the proportion a pod owner's restaked balance has decreased, and
    /// reduces their beacon slashing factor accordingly.
    /// Note: `balanceDecreasedWei` is assumed to be less than `prevRestakedBalanceWei`
    function _reduceSlashingFactor(
        address podOwner,
        uint256 prevRestakedBalanceWei,
        uint256 balanceDecreasedWei
    ) internal returns (uint64) {
        uint256 newRestakedBalanceWei = prevRestakedBalanceWei - balanceDecreasedWei;
        uint64 prevBeaconSlashingFactor = beaconChainSlashingFactor(podOwner);
        // newBeaconSlashingFactor is less than prevBeaconSlashingFactor because
        // newRestakedBalanceWei < prevRestakedBalanceWei
        uint64 newBeaconSlashingFactor =
            uint64(prevBeaconSlashingFactor.mulDiv(newRestakedBalanceWei, prevRestakedBalanceWei));
        uint64 beaconChainSlashingFactorDecrease = prevBeaconSlashingFactor - newBeaconSlashingFactor;
        _beaconChainSlashingFactor[podOwner] =
            BeaconChainSlashingFactor({slashingFactor: newBeaconSlashingFactor, isSet: true});
        emit BeaconChainSlashingFactorDecreased(podOwner, prevBeaconSlashingFactor, newBeaconSlashingFactor);
        return beaconChainSlashingFactorDecrease;
    }

    // VIEW FUNCTIONS

    /// @inheritdoc IEigenPodManager
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

    /// @inheritdoc IEigenPodManager
    function hasPod(
        address podOwner
    ) public view returns (bool) {
        return address(ownerToPod[podOwner]) != address(0);
    }

    /// @notice Returns the current shares of `user` in `strategy`
    /// @dev strategy must be beaconChainETHStrategy
    /// @dev returns 0 if the user has negative shares
    function stakerDepositShares(address user, IStrategy strategy) public view returns (uint256 depositShares) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        return podOwnerDepositShares[user] < 0 ? 0 : uint256(podOwnerDepositShares[user]);
    }

    /// @inheritdoc IEigenPodManager
    function beaconChainSlashingFactor(
        address podOwner
    ) public view returns (uint64) {
        BeaconChainSlashingFactor memory bsf = _beaconChainSlashingFactor[podOwner];
        return bsf.isSet ? bsf.slashingFactor : WAD;
    }
}
