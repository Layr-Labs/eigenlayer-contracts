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

    /// @notice Emitted to notify the update of the beaconChainOracle address
    event BeaconOracleUpdated(address indexed newOracleAddress);

    /// @notice Emitted to notify the deployment of an EigenPod
    event PodDeployed(address indexed eigenPod, address indexed podOwner);

    /// @notice Emitted to notify a deposit of beacon chain ETH recorded in the strategy manager
    event BeaconChainETHDeposited(address indexed podOwner, uint256 amount);

    /// @notice Emitted when `maxPods` value is updated from `previousValue` to `newValue`
    event MaxPodsUpdated(uint256 previousValue, uint256 newValue);

    /// @notice Emitted when a withdrawal of beacon chain ETH is queued
    event BeaconChainETHWithdrawalQueued(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot);
    
    /// @notice Emitted when a withdrawal of beacon chain ETH is completed
    event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot);

    // @notice Emitted when `podOwner` enters the "undelegation limbo" mode
    event UndelegationLimboEntered(address indexed podOwner);

    // @notice Emitted when `podOwner` exits the "undelegation limbo" mode
    event UndelegationLimboExited(address indexed podOwner);

    modifier onlyEigenPod(address podOwner) {
        require(address(ownerToPod[podOwner]) == msg.sender, "EigenPodManager.onlyEigenPod: not a pod");
        _;
    }

    modifier onlyStrategyManager {
        require(msg.sender == address(strategyManager), "EigenPodManager.onlyStrategyManager: not strategyManager");
        _;
    }

    modifier onlyDelegationManager {
        require(msg.sender == address(delegationManager), "EigenPodManager.onlyDelegationManager: not the DelegationManager");
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
    function restakeBeaconChainETH(address podOwner, uint256 amountWei) 
        external 
        onlyEigenPod(podOwner) 
        onlyNotFrozen(podOwner)
        nonReentrant
    {
        _addShares(podOwner, amountWei);
        emit BeaconChainETHDeposited(podOwner, amountWei);
    }

    /**
     * @notice Removes beacon chain ETH from EigenLayer on behalf of the owner of an EigenPod, when the
     *         balance of a validator is lower than how much stake they have committed to EigenLayer
     * @param podOwner is the pod owner whose balance is being updated.
     * @param sharesDelta is the change in podOwner's beaconChainETHStrategy shares
     * @dev Callable only by the podOwner's EigenPod contract.
     */
    function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) external onlyEigenPod(podOwner) nonReentrant {
       _recordBeaconChainETHBalanceUpdate(podOwner, sharesDelta);
    }

    /**
     * @notice Called by a podOwner to queue a withdrawal of some (or all) of their virtual beacon chain ETH shares.
     * @param amountWei The amount of ETH to withdraw.
     * @param withdrawer The address that can complete the withdrawal and receive the withdrawn funds.
     * @param undelegateIfPossible If marked as 'true', the podOwner will be undelegated from their operator in EigenLayer, if possible.
     */
    function queueWithdrawal(
        uint256 amountWei,
        address withdrawer,
        bool undelegateIfPossible
    ) 
        external
        onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
        onlyNotFrozen(msg.sender)
        nonReentrant
        returns(bytes32)
    {
        return _queueWithdrawal(msg.sender, amountWei, withdrawer, undelegateIfPossible);
    }

    /**
     * @notice Completes an existing BeaconChainQueuedWithdrawal by sending the ETH to the 'withdrawer'
     * @param queuedWithdrawal is the queued withdrawal to be completed
     * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
     */
    function completeQueuedWithdrawal(
        BeaconChainQueuedWithdrawal memory queuedWithdrawal,
        uint256 middlewareTimesIndex
    )
        external
        onlyNotFrozen(queuedWithdrawal.delegatedAddress)
        nonReentrant
        onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
    {
        _completeQueuedWithdrawal(queuedWithdrawal, middlewareTimesIndex);
    }

    /**
     * @notice Called by a staker who owns an EigenPod to enter the "undelegation limbo" mode.
     * @dev Undelegation limbo is a mode which a staker can enter into, in which they remove their virtual "beacon chain ETH shares" from EigenLayer's delegation
     * system but do not necessarily withdraw the associated ETH from EigenLayer itself. This mode allows users who have restaked native ETH a route via
     * which they can undelegate from an operator without needing to exit any of their validators from the Consensus Layer.
     */
    function enterUndelegationLimbo()
        external
        onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
        onlyNotFrozen(msg.sender)
        nonReentrant
    {
        _enterUndelegationLimbo(msg.sender);
    }

    /**
     * @notice Called by a staker who owns an EigenPod to exit the "undelegation limbo" mode.
     * @param middlewareTimesIndex Passed on as an input to the `slasher.canWithdraw` function, to ensure that the caller can exit undelegation limbo.
     * This is because undelegation limbo is subject to the same restrictions as completing a withdrawal
     * @param withdrawFundsFromEigenLayer If marked as 'true', then the caller's beacon chain ETH shares will be withdrawn from EigenLayer, as they are when
     * completing a withdrawal. If marked as 'false', then the caller's shares are simply returned to EigenLayer's delegation system.
     */
    function exitUndelegationLimbo(uint256 middlewareTimesIndex, bool withdrawFundsFromEigenLayer)
        external
        onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
        onlyNotFrozen(msg.sender)
        nonReentrant
    {
        require(isInUndelegationLimbo(msg.sender),
                    "EigenPodManager.exitUndelegationLimbo: must be in undelegation limbo");

        uint32 limboStartBlock = _podOwnerUndelegationLimboStatus[msg.sender].startBlock;
        require(
            slasher.canWithdraw(
                _podOwnerUndelegationLimboStatus[msg.sender].delegatedAddress,
                limboStartBlock,
                middlewareTimesIndex
            ),
            "EigenPodManager.exitUndelegationLimbo: shares in limbo are still slashable"
        );

        require(limboStartBlock + strategyManager.withdrawalDelayBlocks() <= block.number,
            "EigenPodManager.exitUndelegationLimbo: withdrawalDelayBlocks period has not yet passed"
        );

        // delete the pod owner's undelegation limbo details
        delete _podOwnerUndelegationLimboStatus[msg.sender];

        // emit event
        emit UndelegationLimboExited(msg.sender);

        // either withdraw the funds entirely from EigenLayer
        if (withdrawFundsFromEigenLayer) {
            // store the pod owner's current share amount in memory, for gas efficiency
            uint256 currentPodOwnerShares = podOwnerShares[msg.sender];
            // remove the shares from the podOwner and reduce the `withdrawableRestakedExecutionLayerGwei` in the pod
            podOwnerShares[msg.sender] = 0;
            _decrementWithdrawableRestakedExecutionLayerGwei(msg.sender, currentPodOwnerShares);
            // withdraw through the ETH from the EigenPod
            _withdrawRestakedBeaconChainETH(msg.sender, msg.sender, currentPodOwnerShares);
        // or else return the "shares" to the delegation system
        } else {
            delegationManager.increaseDelegatedShares(msg.sender, beaconChainETHStrategy, podOwnerShares[msg.sender]);
        }
    }

    /**
     * @notice forces the podOwner into the "undelegation limbo" mode
     * @param podOwner is the staker to be forced into undelegation limbo
     * @dev This function can only be called by the DelegationManager contract
     */
    function forceIntoUndelegationLimbo(address podOwner) external 
        onlyDelegationManager 
        onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
        onlyNotFrozen(podOwner)
        nonReentrant
    {
        _enterUndelegationLimbo(podOwner);
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
    function _queueWithdrawal(
        address podOwner,
        uint256 amountWei,
        address withdrawer,
        bool undelegateIfPossible
    ) 
        internal
        returns (bytes32)
    {
        require(amountWei > 0, "EigenPodManager._queueWithdrawal: amount must be greater than zero");

        require(amountWei % GWEI_TO_WEI == 0,
            "EigenPodManager._queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei");

        require(!isInUndelegationLimbo(podOwner),
            "EigenPodManager._queueWithdrawal: cannot queue a withdrawal when in undelegation limbo");


        _removeShares(podOwner, amountWei);  
        /**
        * This decrements the withdrawableRestakedExecutionLayerGwei which is incremented only when a podOwner proves a full withdrawal.
        * Remember that withdrawableRestakedExecutionLayerGwei tracks the currently withdrawable ETH from the EigenPod.  
        * By doing this, we ensure that the number of shares in EigenLayer matches the amount of withdrawable ETH in 
        * the pod plus any ETH still staked on the beacon chain via other validators pointed to the pod. As a result, a validator 
        * must complete a full withdrawal from the execution layer prior to queuing a withdrawal of 'beacon chain ETH shares' 
        * via EigenLayer, since otherwise withdrawableRestakedExecutionLayerGwei will be 0.
        */      
        _decrementWithdrawableRestakedExecutionLayerGwei(podOwner, amountWei);  

        address delegatedAddress = delegationManager.delegatedTo(podOwner);

        uint96 nonce = uint96(numWithdrawalsQueued[podOwner]);
        unchecked {
            numWithdrawalsQueued[podOwner] = nonce + 1;
        }

        BeaconChainQueuedWithdrawal memory queuedWithdrawal = BeaconChainQueuedWithdrawal({
            shares: amountWei,
            podOwner: podOwner,
            nonce: nonce,
            withdrawalStartBlock: uint32(block.number),
            delegatedAddress: delegatedAddress,
            withdrawer: withdrawer
        });

        // If the `staker` has withdrawn all of their funds from EigenLayer in this transaction, then they can choose to also undelegate
        /**
         * Checking that `stakerCanUndelegate` is not strictly necessary here, but prevents reverting very late in logic,
         * in the case that `undelegateIfPossible` is set to true but the `staker` still has active deposits in EigenLayer.
         */
        if (undelegateIfPossible && stakerCanUndelegate(podOwner)) {
            _undelegate(podOwner);
        }

        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);
        withdrawalRootPending[withdrawalRoot] = true;

        emit BeaconChainETHWithdrawalQueued(
            podOwner,
            amountWei,
            nonce,
            delegatedAddress,
            withdrawer,
            withdrawalRoot
        );

        return withdrawalRoot;
    }

    // TODO: add documentation to this function
    function _completeQueuedWithdrawal(
        BeaconChainQueuedWithdrawal memory queuedWithdrawal,
        uint256 middlewareTimesIndex
    )
        internal
    {
        // find the withdrawalRoot
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);

        // verify that the queued withdrawal is pending
        require(
            withdrawalRootPending[withdrawalRoot],
            "EigenPodManager._completeQueuedWithdrawal: withdrawal is not pending"
        );

        // verify that the withdrawal is completable
        require(
            slasher.canWithdraw(queuedWithdrawal.delegatedAddress, queuedWithdrawal.withdrawalStartBlock, middlewareTimesIndex),
            "EigenPodManager._completeQueuedWithdrawal: shares pending withdrawal are still slashable"
        );

        // enforce minimum delay lag
        require(queuedWithdrawal.withdrawalStartBlock + strategyManager.withdrawalDelayBlocks() <= block.number,
            "EigenPodManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed"
        );

        // verify that the caller is the specified 'withdrawer'
        require(
            msg.sender == queuedWithdrawal.withdrawer,
            "EigenPodManager._completeQueuedWithdrawal: caller must be the withdrawer"
        );

        // reset the storage slot in mapping of queued withdrawals
        withdrawalRootPending[withdrawalRoot] = false;

        // withdraw the ETH from the EigenPod to the caller
        _withdrawRestakedBeaconChainETH(queuedWithdrawal.podOwner, msg.sender, queuedWithdrawal.shares);

        emit BeaconChainETHWithdrawalCompleted(
            queuedWithdrawal.podOwner,
            queuedWithdrawal.shares,
            queuedWithdrawal.nonce,
            queuedWithdrawal.delegatedAddress,
            queuedWithdrawal.withdrawer,
            withdrawalRoot
        );
    }

    function _deployPod() internal onlyWhenNotPaused(PAUSED_NEW_EIGENPODS) returns (IEigenPod) {
        // check that the limit of EigenPods has not been hit, and increment the EigenPod count
        require(numPods + 1 <= maxPods, "EigenPodManager._deployPod: pod limit reached");
        ++numPods;
        // create the pod
        IEigenPod pod = 
            IEigenPod(
                Create2.deploy(
                    0, 
                    bytes32(uint256(uint160(msg.sender))), 
                    // set the beacon address to the eigenPodBeacon and initialize it
                    abi.encodePacked(
                        beaconProxyBytecode, 
                        abi.encode(eigenPodBeacon, "")
                    )
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


    // @notice Increases the `podOwner`'s shares by `shareAmount` and performs a call to the DelegationManager to ensure delegated shares are also tracked correctly
    function _addShares(address podOwner, uint256 shareAmount) internal {
        require(podOwner != address(0), "EigenPodManager._addShares: podOwner cannot be zero address");
        require(shareAmount > 0, "EigenPodManager._addShares: amount must be greater than zero");

        podOwnerShares[podOwner] += shareAmount;

        // if the podOwner has delegated shares, record an increase in their delegated shares
        if (!isInUndelegationLimbo(podOwner)) {
            delegationManager.increaseDelegatedShares(podOwner, beaconChainETHStrategy, shareAmount);
        }
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

        // if the podOwner has delegated shares, record a decrease in their delegated shares
        if (!isInUndelegationLimbo(podOwner)) {
            uint256[] memory shareAmounts = new uint256[](1);
            shareAmounts[0] = shareAmount;
            IStrategy[] memory strategies = new IStrategy[](1);
            strategies[0] = beaconChainETHStrategy;
            delegationManager.decreaseDelegatedShares(podOwner, strategies, shareAmounts);
        }
    }

    // @notice Changes the `podOwner`'s shares by `sharesDelta` and performs a call to the DelegationManager to ensure delegated shares are also tracked correctly
    function _recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) internal {
        if (sharesDelta < 0) {
            // if change in shares is negative, remove the shares
            _removeShares(podOwner, uint256(-sharesDelta));
        } else {
            // if change in shares is positive, add the shares
            _addShares(podOwner, uint256(sharesDelta));
        }      
    }

    /**
     * @notice This function is called to decrement withdrawableRestakedExecutionLayerGwei when a validator queues a withdrawal.
     * @param amountWei is the amount of ETH in wei to decrement withdrawableRestakedExecutionLayerGwei by
     */
    function _decrementWithdrawableRestakedExecutionLayerGwei(address podOwner, uint256 amountWei) 
        internal
    {
        ownerToPod[podOwner].decrementWithdrawableRestakedExecutionLayerGwei(amountWei);
    }

    /**
     * @notice Withdraws ETH from an EigenPod. The ETH must have first been withdrawn from the beacon chain.
     * @param podOwner The owner of the pod whose balance must be withdrawn.
     * @param recipient The recipient of the withdrawn ETH.
     * @param amount The amount of ETH to withdraw.
     * @dev Callable only by the StrategyManager contract.
     */
    function _withdrawRestakedBeaconChainETH(address podOwner, address recipient, uint256 amount)
        internal
    {
        ownerToPod[podOwner].withdrawRestakedBeaconChainETH(recipient, amount);
    }

    /**
     * @notice If the `staker` has no existing shares, then they can `undelegate` themselves.
     * This allows people a "hard reset" in their relationship with EigenLayer after withdrawing all of their stake.
     * @param staker The address to undelegate. Passed on as an input to the `delegationManager.undelegate` function.
     */
    function _undelegate(address staker) internal onlyNotFrozen(staker) {
        require(stakerCanUndelegate(staker), "EigenPodManager._undelegate: staker has active deposits");
        delegationManager.undelegate(staker);
    }

    /**
     * @notice Internal function to enter `podOwner` into undelegation limbo
     * @dev Does nothing if the `podOwner` has no delegated shares (i.e. they are already in undelegation limbo or have no shares)
     * OR if they are not actively delegated to any operator.
     */
    function _enterUndelegationLimbo(address podOwner) internal {
        if (!stakerHasNoDelegatedShares(podOwner) && delegationManager.isDelegated(podOwner)) {
            // look up the address that the pod owner is currrently delegated to in EigenLayer
            address delegatedAddress = delegationManager.delegatedTo(podOwner);

            // store the undelegation limbo details
            _podOwnerUndelegationLimboStatus[podOwner].active = true;
            _podOwnerUndelegationLimboStatus[podOwner].startBlock = uint32(block.number);
            _podOwnerUndelegationLimboStatus[podOwner].delegatedAddress = delegatedAddress;

            // emit event
            emit UndelegationLimboEntered(podOwner);

            // undelegate all shares
            uint256[] memory shareAmounts = new uint256[](1);
            shareAmounts[0] = podOwnerShares[podOwner];
            IStrategy[] memory strategies = new IStrategy[](1);
            strategies[0] = beaconChainETHStrategy;
            delegationManager.decreaseDelegatedShares(podOwner, strategies, shareAmounts);

            // undelegate the pod owner, if possible
            if (stakerCanUndelegate(podOwner)) {
                _undelegate(podOwner);
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
                    keccak256(abi.encodePacked(
                        beaconProxyBytecode, 
                        abi.encode(eigenPodBeacon, "")
                    )) //bytecode
                ));
        }
        return pod;
    }

    /// @notice Returns 'true' if the `podOwner` has created an EigenPod, and 'false' otherwise.
    function hasPod(address podOwner) public view returns (bool) {
        return address(ownerToPod[podOwner]) != address(0);
    }

    /// @notice Returns the Beacon block root at `timestamp`. Reverts if the Beacon block root at `timestamp` has not yet been finalized.
    function getBlockRootAtTimestamp(uint64 timestamp) external view returns(bytes32) {
        bytes32 stateRoot = beaconChainOracle.beaconStateRootAtBlockNumber(timestamp);
        require(stateRoot != bytes32(0), "EigenPodManager.getBlockRootAtTimestamp: state root at timestamp not yet finalized");
        return stateRoot;
    }

    /// @notice Returns the keccak256 hash of `queuedWithdrawal`.    
    function calculateWithdrawalRoot(BeaconChainQueuedWithdrawal memory queuedWithdrawal) public pure returns (bytes32) {
        return (
            keccak256(
                abi.encode(
                    queuedWithdrawal.shares,
                    queuedWithdrawal.podOwner,
                    queuedWithdrawal.nonce,
                    queuedWithdrawal.withdrawalStartBlock,
                    queuedWithdrawal.delegatedAddress,
                    queuedWithdrawal.withdrawer
                )
            )
        );
    }

    // @notice Returns 'true' if `staker` can undelegate and false otherwise
    function stakerCanUndelegate(address staker) public view returns (bool) {
        return (stakerHasNoDelegatedShares(staker) && strategyManager.stakerStrategyListLength(staker) == 0);
    }

    /**
     * @notice Returns 'true' if `staker` has removed all of their beacon chain ETH "shares" from delegation, either by queuing a withdrawal for them
     * OR by going into "undelegation limbo", and 'false' otherwise
     */
    function stakerHasNoDelegatedShares(address staker) public view returns (bool) {
        return (podOwnerShares[staker] == 0 || isInUndelegationLimbo(staker));
    }

    // @notice Getter function for the internal `_podOwnerUndelegationLimboStatus` mapping.
    function podOwnerUndelegationLimboStatus(address podOwner) external view returns (UndelegationLimboStatus memory) {
        return _podOwnerUndelegationLimboStatus[podOwner];
    }

    // @notice Getter function for `_podOwnerUndelegationLimboStatus.active`.
    function isInUndelegationLimbo(address podOwner) public view returns (bool) {
        return _podOwnerUndelegationLimboStatus[podOwner].active;
    }

}