// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Create2.sol";
import "./BeaconProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IDelegationManager.sol";
import "../../contracts/interfaces/IEigenPodManager.sol";
import "../../contracts/interfaces/IETHPOSDeposit.sol";
import "../../contracts/interfaces/IEigenPod.sol";
import "../../contracts/interfaces/IBeaconChainOracle.sol";

// import "forge-std/Test.sol";

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
contract EigenPodManagerNEW is Initializable, OwnableUpgradeable, IEigenPodManager {
    function getBlockRootAtTimestamp(uint64 timestamp) external view returns(bytes32) {}

    function pause(uint256 newPausedStatus) external {}    

    function pauseAll() external {}

    function paused() external view returns (uint256) {}

    function paused(uint8 index) external view returns (bool) {}

    function setPauserRegistry(IPauserRegistry newPauserRegistry) external {}

    function pauserRegistry() external view returns (IPauserRegistry) {}

    function unpause(uint256 newPausedStatus) external {}

    function ownerToPod(address podOwner) external view returns(IEigenPod) {}


    IETHPOSDeposit public immutable ethPOS;
    /// @notice Beacon proxy to which the EigenPods point
    IBeacon public immutable eigenPodBeacon;

    /// @notice EigenLayer's StrategyManager contract
    IStrategyManager public immutable strategyManager;

    /// @notice EigenLayer's Slasher contract
    ISlasher public immutable slasher;

    /// @notice Oracle contract that provides updates to the beacon chain's state
    IBeaconChainOracle public beaconChainOracle;
    
    /// @notice Pod owner to the amount of penalties they have paid that are still in this contract
    mapping(address => uint256) public podOwnerToUnwithdrawnPaidPenalties;

    /// @notice Emitted when an EigenPod pays penalties, on behalf of its owner
    event PenaltiesPaid(address indexed podOwner, uint256 amountPaid);

    modifier onlyEigenPod(address podOwner) {
        require(address(getPod(podOwner)) == msg.sender, "EigenPodManager.onlyEigenPod: not a pod");
        _;
    }

    modifier onlyStrategyManager {
        require(msg.sender == address(strategyManager), "EigenPodManager.onlyStrategyManager: not strategyManager");
        _;
    }

    constructor(IETHPOSDeposit _ethPOS, IBeacon _eigenPodBeacon, IStrategyManager _strategyManager, ISlasher _slasher) {
        ethPOS = _ethPOS;
        eigenPodBeacon = _eigenPodBeacon;
        strategyManager = _strategyManager;
        slasher = _slasher;
        _disableInitializers();
        
    }

    function initialize(IBeaconChainOracle _beaconChainOracle, address initialOwner) public initializer {
        _updateBeaconChainOracle(_beaconChainOracle);
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Creates an EigenPod for the sender.
     * @dev Function will revert if the `msg.sender` already has an EigenPod.
     */
    function createPod() external {
        require(!hasPod(msg.sender), "EigenPodManager.createPod: Sender already has a pod");
        //deploy a pod if the sender doesn't have one already
        IEigenPod pod = _deployPod();

        emit PodDeployed(address(pod), msg.sender);
    }

    /**
     * @notice Stakes for a new beacon chain validator on the sender's EigenPod. 
     * Also creates an EigenPod for the sender if they don't have one already.
     * @param pubkey The 48 bytes public key of the beacon chain validator.
     * @param signature The validator's signature of the deposit data.
     * @param depositDataRoot The root/hash of the deposit data for the validator's deposit.
     */
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable {
        IEigenPod pod = getPod(msg.sender);
        if(!hasPod(msg.sender)) {
            //deploy a pod if the sender doesn't have one already
            pod = _deployPod();
        }
        pod.stake{value: msg.value}(pubkey, signature, depositDataRoot);
    }

    /**
     * @notice Deposits/Restakes beacon chain ETH in EigenLayer on behalf of the owner of an EigenPod.
     * @param podOwner The owner of the pod whose balance must be deposited.
     * @param amount The amount of ETH to 'deposit' (i.e. be credited to the podOwner).
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function restakeBeaconChainETH(address podOwner, uint256 amount) external onlyEigenPod(podOwner) {}

    /**
     * @notice Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
     *         balance of a validator is lower than how much stake they have committed to EigenLayer
     * @param podOwner The owner of the pod whose balance must be removed.
     * @param sharesDelta is the change in podOwner's beaconChainETHStrategy shares
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) external onlyEigenPod(podOwner){}

    /**
     * @notice Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.
     * @param podOwner The owner of the pod whose balance must be withdrawn.
     * @param recipient The recipient of the withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the StrategyManager contract.
     */
    function withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount) external onlyStrategyManager {
        getPod(podOwner).withdrawRestakedBeaconChainETH(recipient, amount);
    }

    /**
     * @notice Records receiving ETH from the `PodOwner`'s EigenPod, paid in order to fullfill the EigenPod's penalties to EigenLayer
     * @param podOwner The owner of the pod whose balance is being sent.
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function payPenalties(address podOwner) external payable onlyEigenPod(podOwner) {
        podOwnerToUnwithdrawnPaidPenalties[podOwner] += msg.value;
        emit PenaltiesPaid(podOwner, msg.value);
    }

    /**
     * @notice Withdraws paid penalties of the `podOwner`'s EigenPod, to the `recipient` address
     * @param recipient The recipient of withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the strategyManager.owner().
     */
    function withdrawPenalties(address podOwner, address recipient, uint256 amount) external {
        require(msg.sender == Ownable(address(strategyManager)).owner(), "EigenPods.withdrawPenalties: only strategyManager owner");
        podOwnerToUnwithdrawnPaidPenalties[podOwner] -= amount;
        // transfer penalties from pod to `recipient`
        Address.sendValue(payable(recipient), amount);
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
    function _deployPod() internal returns (IEigenPod) {
        IEigenPod pod = 
            IEigenPod(
                Create2.deploy(
                    0, 
                    bytes32(uint256(uint160(msg.sender))), 
                    // set the beacon address to the eigenPodBeacon and initialize it
                    abi.encodePacked(
                        type(BeaconProxyNEW).creationCode, 
                        abi.encode(eigenPodBeacon, abi.encodeWithSelector(IEigenPod.initialize.selector, IEigenPodManager(address(this)), msg.sender))
                    )
                )
            );
        return pod;
    }

    function _updateBeaconChainOracle(IBeaconChainOracle newBeaconChainOracle) internal {
        beaconChainOracle = newBeaconChainOracle;
        emit BeaconOracleUpdated(address(newBeaconChainOracle));
    }

    // VIEW FUNCTIONS
    /// @notice Returns the address of the `podOwner`'s EigenPod (whether it is deployed yet or not).
    function getPod(address podOwner) public view returns (IEigenPod) {
        return IEigenPod(
                Create2.computeAddress(
                    bytes32(uint256(uint160(podOwner))), //salt
                    keccak256(abi.encodePacked(
                        type(BeaconProxyNEW).creationCode, 
                        abi.encode(eigenPodBeacon, abi.encodeWithSelector(IEigenPod.initialize.selector, IEigenPodManager(address(this)), podOwner))
                    )) //bytecode
                ));
    }

    /// @notice Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.
    function hasPod(address podOwner) public view returns (bool) {
        return address(getPod(podOwner)).code.length > 0;
    }

    function getBlockRootAtTimestamp() external view returns(bytes32) {
        // return beaconChainOracle.getBlockRootAtTimestamp();
    }

    function podOwnerShares(address podOwner) external returns (uint256){
        // return podOwner[podOwner];
    }

    function queueWithdrawal(uint256 amountWei, address withdrawer) external returns(bytes32){}

    function forceIntoUndelegationLimbo(address podOwner, address delegatedTo) external returns (uint256) {}

    function completeQueuedWithdrawal(BeaconChainQueuedWithdrawal memory queuedWithdrawal, uint256 middlewareTimesIndex) external{}

    function beaconChainETHStrategy() external view returns (IStrategy){}

    function podOwnerHasActiveShares(address staker) external view returns (bool) {}

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.    
    function calculateWithdrawalRoot(BeaconChainQueuedWithdrawal memory queuedWithdrawal) external pure returns (bytes32) {}

    // @notice Getter function for the internal `_podOwnerUndelegationLimboStatus` mapping.
    function podOwnerUndelegationLimboStatus(address podOwner) external view returns (UndelegationLimboStatus memory) {}

    // @notice Getter function for `_podOwnerUndelegationLimboStatus.undelegationLimboActive`.
    function isInUndelegationLimbo(address podOwner) external view returns (bool) {}
}