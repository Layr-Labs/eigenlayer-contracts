// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";


import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPod.sol";

import "../interfaces/IBeaconChainOracle.sol";

import "../permissions/Pausable.sol";
import "./EigenPodPausingConstants.sol";

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
    IEigenPodManager,
    EigenPodPausingConstants,
    ReentrancyGuardUpgradeable
{
    /**
     * @notice Stored code of type(BeaconProxy).creationCode
     * @dev Maintained as a constant to solve an edge case - changes to OpenZeppelin's BeaconProxy code should not cause
     * addresses of EigenPods that are pre-computed with Create2 to change, even upon upgrading this contract, changing compiler version, etc.
    */
    bytes internal constant beaconProxyBytecode = hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";

    // @notice Internal constant used in calculations, since the beacon chain stores balances in Gwei rather than wei
    uint256 internal constant GWEI_TO_WEI = 1e9;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /// @notice The ETH2 Deposit Contract
    IETHPOSDeposit public immutable ethPOS;
    
    /// @notice Beacon proxy to which the EigenPods point
    IBeacon public immutable eigenPodBeacon;

    /// @notice EigenLayer's StrategyManager contract
    IStrategyManager public immutable strategyManager;

    /// @notice EigenLayer's Slasher contract
    ISlasher public immutable slasher;

    /// @notice EigenLayer's DelegationManager contract
    IDelegationManager public immutable delegationManager;

    /// @notice Oracle contract that provides updates to the beacon chain's state
    IBeaconChainOracle public beaconChainOracle;

    /// @notice Pod owner to deployed EigenPod address
    mapping(address => IEigenPod) public ownerToPod;

    // BEGIN STORAGE VARIABLES ADDED AFTER FIRST TESTNET DEPLOYMENT -- DO NOT SUGGEST REORDERING TO CONVENTIONAL ORDER
    /// @notice The number of EigenPods that have been deployed
    uint256 public numPods;

    /// @notice The maximum number of EigenPods that can be deployed
    uint256 public maxPods;

    // BEGIN STORAGE VARIABLES ADDED AFTER MAINNET DEPLOYMENT -- DO NOT SUGGEST REORDERING TO CONVENTIONAL ORDER
    /// @notice Pod owner to the number of shares they have in the beacon chain ETH strategy
    mapping(address => uint256) public podOwnerShares;

    /// @notice Mapping: podOwner => cumulative number of queued withdrawals of beaconchainETH they have ever initiated. only increments (doesn't decrement)
    mapping(address => uint256) public numWithdrawalsQueued;

    /// @notice Mapping: hash of withdrawal inputs, aka 'withdrawalRoot' => whether the withdrawal is pending
    mapping(bytes32 => bool) public withdrawalRootPending;

    // @notice Mapping: pod owner => UndelegationLimboStatus struct. Mapping is internal so we can have a getter that returns a memory struct.
    mapping(address => UndelegationLimboStatus) internal _podOwnerUndelegationLimboStatus;

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
    ) {
        ethPOS = _ethPOS;
        eigenPodBeacon = _eigenPodBeacon;
        strategyManager = _strategyManager;
        slasher = _slasher;
        delegationManager = _delegationManager;
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
        onlyWhenNotPaused(PAUSED_DEPOSITS)
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
     */
    function queueWithdrawal(
        uint256 amountWei,
        address withdrawer
    ) 
        external
        onlyWhenNotPaused(PAUSED_WITHDRAWALS)
        onlyNotFrozen(msg.sender)
        nonReentrant
        returns(bytes32)
    {
        return _queueWithdrawal(msg.sender, amountWei, withdrawer);
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
        onlyWhenNotPaused(PAUSED_WITHDRAWALS)
    {
        _completeQueuedWithdrawal(queuedWithdrawal, middlewareTimesIndex);
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
        onlyWhenNotPaused(PAUSED_WITHDRAWALS)
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

        // enforce minimum delay lag
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
    
    // EXTERNAL FUNCTIONS PERMISSIONED TO SINGLE PARTIES
    // TODO: write documentation for this function
    function slashShares(
        address slashedPodOwner,
        address slashedFundsRecipient,
        uint256 shareAmount
    )
        external
        onlyOwner
        onlyFrozen(slashedPodOwner)
        nonReentrant
    {
        require(
            // either the pod owner themselves is frozen
            slasher.isFrozen(slashedPodOwner) ||
            // or they are in "undelegation limbo" and the operator who they *were* delegated to is frozen
            (
                isInUndelegationLimbo(slashedPodOwner) &&
                slasher.isFrozen(_podOwnerUndelegationLimboStatus[slashedPodOwner].delegatedAddress)
            ),
            "EigenPodManager.slashShares: cannot slash the specified pod owner"
        );
        require(shareAmount > 0, "EigenPodManager.slashShares: shares must be greater than zero");
        require(podOwnerShares[slashedPodOwner] >= shareAmount,
            "EigenPodManager.slashShares: podOwnerShares[podOwner] must be greater than or equal to shares");

        // remove the shares, without trying to undelegate the `slashedPodOwner`
        _removeShares(slashedPodOwner, shareAmount);
        // TODO: @Sidu28 -- confirm that decrementing `withdrawableRestakedExecutionLayerGwei` is correct/intended here
        _decrementWithdrawableRestakedExecutionLayerGwei(slashedPodOwner, shareAmount);

        // send the ETH to the `slashedFundsRecipient`
        _withdrawRestakedBeaconChainETH(slashedPodOwner, slashedFundsRecipient, shareAmount);
    }

    // TODO: write documentation for this function
    function slashQueuedWithdrawal(
        address slashedFundsRecipient,
        BeaconChainQueuedWithdrawal memory queuedWithdrawal
    )
        external 
        onlyOwner 
        onlyFrozen(queuedWithdrawal.delegatedAddress) 
        nonReentrant 
    {
        // find the withdrawalRoot
        bytes32 withdrawalRoot = calculateWithdrawalRoot(queuedWithdrawal);

         // verify that the queued withdrawal is pending
        require(
            withdrawalRootPending[withdrawalRoot],
            "EigenPodManager.slashQueuedWithdrawal: withdrawal is not pending"
        );

        // reset the storage slot in mapping of queued withdrawals
        withdrawalRootPending[withdrawalRoot] = false;

        _withdrawRestakedBeaconChainETH(queuedWithdrawal.podOwner, slashedFundsRecipient, queuedWithdrawal.shares);
    }

    /**
     * @notice forces the podOwner into the "undelegation limbo" mode, and returns the number of virtual 'beacon chain ETH shares'
     * that the podOwner has, which were entered into undelegation limbo.
     * @param podOwner is the staker to be forced into undelegation limbo
     * @param delegatedTo is the operator the staker is currently delegated to
     * @dev This function can only be called by the DelegationManager contract
     */
    function forceIntoUndelegationLimbo(address podOwner, address delegatedTo)
        external
        onlyDelegationManager 
        onlyWhenNotPaused(PAUSED_WITHDRAWALS)
        onlyNotFrozen(podOwner)
        nonReentrant
        returns (uint256 sharesRemovedFromDelegation)
    {
        // put the `podOwner` into undelegation limbo, and return the amount of shares which were entered into undelegation limbo
        return _enterUndelegationLimbo(podOwner, delegatedTo);
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
    /**
     * @notice Queues a withdrawal of `amountWei` of virtual "beacon chain ETH shares" from `podOwner` to `withdrawer`.
     */
    function _queueWithdrawal(
        address podOwner,
        uint256 amountWei,
        address withdrawer
    ) 
        internal
        returns (bytes32)
    {
        require(amountWei > 0, "EigenPodManager._queueWithdrawal: amount must be greater than zero");

        require(amountWei % GWEI_TO_WEI == 0,
            "EigenPodManager._queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei");

        require(!isInUndelegationLimbo(podOwner),
            "EigenPodManager._queueWithdrawal: cannot queue a withdrawal when in undelegation limbo");


        // Decrease podOwner's shares here (and in DelegationManager if podOwner is delegated)
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
            // if change in shares is negative, remove the shares (and don't bother trying to undelegate the `podOwner`)
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
        internal onlyWhenNotPaused(PAUSED_WITHDRAW_RESTAKED_ETH)
    {
        ownerToPod[podOwner].withdrawRestakedBeaconChainETH(recipient, amount);
    }

    /**
     * @notice Internal function to enter `podOwner` into undelegation limbo, and return the number of shares which were moved into undelegation limbo.
     * @dev Does nothing if the `podOwner` has no delegated shares (i.e. they are already in undelegation limbo or have no shares)
     * OR if they are not actively delegated to any operator.
     *
     * This method assumes the podOwner is delegated, as it's being called by the
     * DelegationManager (and supplied with `delegatedTo`)
     */
    function _enterUndelegationLimbo(address podOwner, address delegatedTo) internal returns (uint256) {
        if (podOwnerShares[podOwner] != 0 && !isInUndelegationLimbo(podOwner)) {
            // store the undelegation limbo details
            _podOwnerUndelegationLimboStatus[podOwner].active = true;
            _podOwnerUndelegationLimboStatus[podOwner].startBlock = uint32(block.number);
            _podOwnerUndelegationLimboStatus[podOwner].delegatedAddress = delegatedTo;

            // emit event
            emit UndelegationLimboEntered(podOwner);

            return podOwnerShares[podOwner];
        } else {
            return 0;
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

    /**
     * @notice Returns 'true' if `staker` has removed all of their beacon chain ETH "shares" from delegation, either by queuing a
     * withdrawal for them OR by going into "undelegation limbo", and 'false' otherwise
     */
    function podOwnerHasNoDelegatedShares(address staker) public view returns (bool) {
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

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[42] private __gap;
}