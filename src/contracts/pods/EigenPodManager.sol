// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";

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
    modifier onlyEigenPod(address podOwner) {
        require(address(ownerToPod[podOwner]) == msg.sender, "EigenPodManager.onlyEigenPod: not a pod");
        _;
    }

    modifier onlyDelegationManager() {
        require(
            msg.sender == address(delegationManager), "EigenPodManager.onlyDelegationManager: not the DelegationManager"
        );
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
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 _initPausedStatus
    ) external initializer {
        _transferOwnership(initialOwner);
        _initializePauser(_pauserRegistry, _initPausedStatus);
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     * @dev Returns EigenPod address
     */
    function createPod() external onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) returns (address) {
        require(!hasPod(msg.sender), "EigenPodManager.createPod: Sender already has a pod");
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
     * @dev Callable only by the podOwner's EigenPod contract.
     * @dev Reverts if `sharesDelta` is not a whole Gwei amount
     */
    function recordBeaconChainETHBalanceUpdate(
        address podOwner,
        int256 sharesDelta
    ) external onlyEigenPod(podOwner) nonReentrant {
        require(
            podOwner != address(0), "EigenPodManager.recordBeaconChainETHBalanceUpdate: podOwner cannot be zero address"
        );
        require(
            sharesDelta % int256(GWEI_TO_WEI) == 0,
            "EigenPodManager.recordBeaconChainETHBalanceUpdate: sharesDelta must be a whole Gwei amount"
        );
        int256 currentPodOwnerShares = podOwnerShares[podOwner];
        int256 updatedPodOwnerShares = currentPodOwnerShares + sharesDelta;
        podOwnerShares[podOwner] = updatedPodOwnerShares;

        // inform the DelegationManager of the change in delegateable shares
        int256 changeInDelegatableShares = _calculateChangeInDelegatableShares({
            sharesBefore: currentPodOwnerShares,
            sharesAfter: updatedPodOwnerShares
        });
        // skip making a call to the DelegationManager if there is no change in delegateable shares
        if (changeInDelegatableShares != 0) {
            if (changeInDelegatableShares < 0) {
                delegationManager.decreaseDelegatedShares({
                    staker: podOwner,
                    strategy: beaconChainETHStrategy,
                    shares: uint256(-changeInDelegatableShares)
                });
            } else {
                delegationManager.increaseDelegatedShares({
                    staker: podOwner,
                    strategy: beaconChainETHStrategy,
                    shares: uint256(changeInDelegatableShares)
                });
            }
        }
        emit PodSharesUpdated(podOwner, sharesDelta);
        emit NewTotalShares(podOwner, updatedPodOwnerShares);
    }

    /**
     * @notice Used by the DelegationManager to remove a pod owner's shares while they're in the withdrawal queue.
     * Simply decreases the `podOwner`'s shares by `shares`, down to a minimum of zero.
     * @dev This function reverts if it would result in `podOwnerShares[podOwner]` being less than zero, i.e. it is forbidden for this function to
     * result in the `podOwner` incurring a "share deficit". This behavior prevents a Staker from queuing a withdrawal which improperly removes excessive
     * shares from the operator to whom the staker is delegated.
     * @dev Reverts if `shares` is not a whole Gwei amount
     * @dev The delegation manager validates that the podOwner is not address(0)
     */
    function removeShares(address podOwner, uint256 shares) external onlyDelegationManager {
        require(int256(shares) >= 0, "EigenPodManager.removeShares: shares cannot be negative");
        require(shares % GWEI_TO_WEI == 0, "EigenPodManager.removeShares: shares must be a whole Gwei amount");
        int256 updatedPodOwnerShares = podOwnerShares[podOwner] - int256(shares);
        require(
            updatedPodOwnerShares >= 0,
            "EigenPodManager.removeShares: cannot result in pod owner having negative shares"
        );
        podOwnerShares[podOwner] = updatedPodOwnerShares;

        emit NewTotalShares(podOwner, updatedPodOwnerShares);
    }

    /**
     * @notice Increases the `podOwner`'s shares by `shares`, paying off deficit if possible.
     * Used by the DelegationManager to award a pod owner shares on exiting the withdrawal queue
     * @dev Returns the number of shares added to `podOwnerShares[podOwner]` above zero, which will be less than the `shares` input
     * in the event that the podOwner has an existing shares deficit (i.e. `podOwnerShares[podOwner]` starts below zero)
     * @dev Reverts if `shares` is not a whole Gwei amount
     */
    function addShares(address podOwner, uint256 shares) external onlyDelegationManager returns (uint256) {
        require(podOwner != address(0), "EigenPodManager.addShares: podOwner cannot be zero address");
        require(int256(shares) >= 0, "EigenPodManager.addShares: shares cannot be negative");
        require(shares % GWEI_TO_WEI == 0, "EigenPodManager.addShares: shares must be a whole Gwei amount");
        int256 currentPodOwnerShares = podOwnerShares[podOwner];
        int256 updatedPodOwnerShares = currentPodOwnerShares + int256(shares);
        podOwnerShares[podOwner] = updatedPodOwnerShares;

        emit PodSharesUpdated(podOwner, int256(shares));
        emit NewTotalShares(podOwner, updatedPodOwnerShares);

        return uint256(
            _calculateChangeInDelegatableShares({
                sharesBefore: currentPodOwnerShares,
                sharesAfter: updatedPodOwnerShares
            })
        );
    }

    /**
     * @notice Used by the DelegationManager to complete a withdrawal, sending tokens to some destination address
     * @dev Prioritizes decreasing the podOwner's share deficit, if they have one
     * @dev Reverts if `shares` is not a whole Gwei amount
     * @dev This function assumes that `removeShares` has already been called by the delegationManager, hence why
     *      we do not need to update the podOwnerShares if `currentPodOwnerShares` is positive
     */
    function withdrawSharesAsTokens(
        address podOwner,
        address destination,
        uint256 shares
    ) external onlyDelegationManager {
        require(podOwner != address(0), "EigenPodManager.withdrawSharesAsTokens: podOwner cannot be zero address");
        require(destination != address(0), "EigenPodManager.withdrawSharesAsTokens: destination cannot be zero address");
        require(int256(shares) >= 0, "EigenPodManager.withdrawSharesAsTokens: shares cannot be negative");
        require(shares % GWEI_TO_WEI == 0, "EigenPodManager.withdrawSharesAsTokens: shares must be a whole Gwei amount");
        int256 currentPodOwnerShares = podOwnerShares[podOwner];

        // if there is an existing shares deficit, prioritize decreasing the deficit first
        if (currentPodOwnerShares < 0) {
            uint256 currentShareDeficit = uint256(-currentPodOwnerShares);

            if (shares > currentShareDeficit) {
                // get rid of the whole deficit if possible, and pass any remaining shares onto destination
                podOwnerShares[podOwner] = 0;
                shares -= currentShareDeficit;
                emit PodSharesUpdated(podOwner, int256(currentShareDeficit));
                emit NewTotalShares(podOwner, 0);
            } else {
                // otherwise get rid of as much deficit as possible, and return early, since there is nothing left over to forward on
                int256 updatedPodOwnerShares = podOwnerShares[podOwner] + int256(shares);
                podOwnerShares[podOwner] = updatedPodOwnerShares;
                emit PodSharesUpdated(podOwner, int256(shares));
                emit NewTotalShares(podOwner, updatedPodOwnerShares);
                return;
            }
        }
        // Actually withdraw to the destination
        ownerToPod[podOwner].withdrawRestakedBeaconChainETH(destination, shares);
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

    /**
     * @notice Calculates the change in a pod owner's delegateable shares as a result of their beacon chain ETH shares changing
     * from `sharesBefore` to `sharesAfter`. The key concept here is that negative/"deficit" shares are not delegateable.
     */
    function _calculateChangeInDelegatableShares(
        int256 sharesBefore,
        int256 sharesAfter
    ) internal pure returns (int256) {
        if (sharesBefore <= 0) {
            if (sharesAfter <= 0) {
                // if the shares started negative and stayed negative, then there cannot have been an increase in delegateable shares
                return 0;
            } else {
                // if the shares started negative and became positive, then the increase in delegateable shares is the ending share amount
                return sharesAfter;
            }
        } else {
            if (sharesAfter <= 0) {
                // if the shares started positive and became negative, then the decrease in delegateable shares is the starting share amount
                return (-sharesBefore);
            } else {
                // if the shares started positive and stayed positive, then the change in delegateable shares
                // is the difference between starting and ending amounts
                return (sharesAfter - sharesBefore);
            }
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
}
