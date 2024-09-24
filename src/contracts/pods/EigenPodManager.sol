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
     * @dev Callable only by the podOwner's EigenPod contract.
     * @dev Reverts if `sharesDelta` is not a whole Gwei amount
     */
    function recordBeaconChainETHBalanceUpdate(
        address podOwner,
        int256 sharesDelta
    ) external onlyEigenPod(podOwner) nonReentrant {
        require(podOwner != address(0), InputAddressZero());
        require(sharesDelta % int256(GWEI_TO_WEI) == 0, SharesNotMultipleOfGwei());
        int256 currentPodOwnerShares = podOwnerShares[podOwner];
        int256 updatedPodOwnerShares = currentPodOwnerShares + sharesDelta;
        podOwnerShares[podOwner] = updatedPodOwnerShares;

        // inform the DelegationManager of the change in delegateable shares
        int256 changeInDelegatableShares = _calculateChangeInDelegatableShares({
            sharesBefore: currentPodOwnerShares,
            sharesAfter: updatedPodOwnerShares
        });
        // skip making a call to the DelegationManager if there is no change in delegateable shares
        // or if the currentPodShares < 0 and updatedPodShares is still < 0. Means no update required
        // in delegated shares
        if (changeInDelegatableShares != 0) {
            if (changeInDelegatableShares < 0) {
                delegationManager.decreaseDelegatedShares({
                    staker: podOwner,
                    strategy: beaconChainETHStrategy,
                    removedShares: uint256(-changeInDelegatableShares).wrapWithdrawable()
                });
            } else {
                delegationManager.increaseDelegatedShares({
                    staker: podOwner,
                    strategy: beaconChainETHStrategy,
                    // existing shares from standpoint of the DelegationManager
                    existingPrincipalShares: (currentPodOwnerShares < 0 ? 0 : uint256(currentPodOwnerShares)).wrapShares(),
                    addedShares: uint256(changeInDelegatableShares).wrapWithdrawable()
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
    function removeShares(address staker, IStrategy strategy, Shares shares) external onlyDelegationManager {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        require(int256(shares.unwrap()) >= 0, SharesNegative());
        require(shares.unwrap() % GWEI_TO_WEI == 0, SharesNotMultipleOfGwei());
        int256 updatedShares = podOwnerShares[staker] - int256(shares.unwrap());
        require(updatedShares >= 0, SharesNegative());
        podOwnerShares[staker] = updatedShares;

        emit NewTotalShares(staker, updatedShares);
    }

    /**
     * @notice Increases the `podOwner`'s shares by `shares`, paying off deficit if possible.
     * Used by the DelegationManager to award a pod owner shares on exiting the withdrawal queue
     * @dev Returns the number of shares added to `podOwnerShares[podOwner]` above zero, which will be less than the `shares` input
     * in the event that the podOwner has an existing shares deficit (i.e. `podOwnerShares[podOwner]` starts below zero).
     * Also returns existingPodShares prior to adding shares, this is returned as 0 if the existing podOwnerShares is negative
     * @dev Reverts if `shares` is not a whole Gwei amount
     */
    function addShares(
        address staker,
        IERC20 token,
        IStrategy strategy,
        Shares shares
    ) external onlyDelegationManager {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        require(staker != address(0), InputAddressZero());
        require(int256(shares.unwrap()) >= 0, SharesNegative());
        require(shares.unwrap() % GWEI_TO_WEI == 0, SharesNotMultipleOfGwei());
        int256 currentShares = podOwnerShares[staker];
        int256 updatedShares = currentShares + int256(shares.unwrap());
        podOwnerShares[staker] = updatedShares;

        emit PodSharesUpdated(staker, int256(shares.unwrap()));
        emit NewTotalShares(staker, updatedShares);

        uint256 increaseInDelegateableShares =
            uint256(_calculateChangeInDelegatableShares({sharesBefore: currentShares, sharesAfter: updatedShares}));

        // TODO: ADD SHARES BACK TO DM
    }

    /**
     * @notice Used by the DelegationManager to complete a withdrawal, sending tokens to some destination address
     * @dev Prioritizes decreasing the podOwner's share deficit, if they have one
     * @dev Reverts if `shares` is not a whole Gwei amount
     * @dev This function assumes that `removeShares` has already been called by the delegationManager, hence why
     *      we do not need to update the podOwnerShares if `currentPodOwnerShares` is positive
     */
    function withdrawSharesAsTokens(
        address recipient,
        IStrategy strategy,
        Shares shares,
        IERC20
    ) external onlyDelegationManager {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        require(recipient != address(0), InputAddressZero());
        require(int256(shares.unwrap()) >= 0, SharesNegative());
        require(shares.unwrap() % GWEI_TO_WEI == 0, SharesNotMultipleOfGwei());
        int256 currentShares = podOwnerShares[recipient];

        // if there is an existing shares deficit, prioritize decreasing the deficit first
        if (currentShares < 0) {
            uint256 currentShareDeficit = uint256(-currentShares);

            if (shares.unwrap() > currentShareDeficit) {
                // get rid of the whole deficit if possible, and pass any remaining shares onto destination
                podOwnerShares[recipient] = 0;
                shares = shares.sub(currentShareDeficit).wrapShares();
                emit PodSharesUpdated(recipient, int256(currentShareDeficit));
                emit NewTotalShares(recipient, 0);
            } else {
                // otherwise get rid of as much deficit as possible, and return early, since there is nothing left over to forward on
                int256 updatedShares = podOwnerShares[recipient] + int256(shares.unwrap());
                podOwnerShares[recipient] = updatedShares;
                emit PodSharesUpdated(recipient, int256(shares.unwrap()));
                emit NewTotalShares(recipient, updatedShares);
                return;
            }
        }
        // Actually withdraw to the destination
        ownerToPod[recipient].withdrawRestakedBeaconChainETH(recipient, shares.unwrap());
    }

    /// @notice Returns the current shares of `user` in `strategy`
    /// @dev strategy must be beaconChainETH when talking to the EigenPodManager
    /// @dev returns 0 if the user has negative shares
    function stakerStrategyShares(address user, IStrategy strategy) public view returns (Shares shares) {
        require(strategy == beaconChainETHStrategy, InvalidStrategy());
        return (podOwnerShares[user] < 0 ? 0 : uint256(podOwnerShares[user])).wrapShares();
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
