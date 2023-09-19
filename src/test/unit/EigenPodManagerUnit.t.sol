// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "forge-std/Test.sol";

import "../../contracts/pods/EigenPodManager.sol";
import "../../contracts/pods/EigenPodPausingConstants.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../mocks/delegationManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodMock.sol";
import "../mocks/ETHDepositMock.sol";
import "../mocks/Reenterer.sol";
import "../mocks/Reverter.sol";

contract EigenPodManagerUnitTests is Test, EigenPodPausingConstants {

    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 public REQUIRED_BALANCE_WEI = 31 ether;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManager public eigenPodManager;

    StrategyManagerMock public strategyManagerMock;
    DelegationManagerMock public delegationManagerMock;
    SlasherMock public slasherMock;
    IETHPOSDeposit public ethPOSMock;
    IEigenPod public eigenPodImplementation;
    IBeacon public eigenPodBeacon;

    IStrategy public beaconChainETHStrategy;

    Reenterer public reenterer;

    uint256 GWEI_TO_WEI = 1e9;

    address public pauser = address(555);
    address public unpauser = address(999);

    address initialOwner = address(this);

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

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

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        slasherMock = new SlasherMock();
        delegationManagerMock = new DelegationManagerMock();
        strategyManagerMock = new StrategyManagerMock();
        ethPOSMock = new ETHPOSDepositMock();
        eigenPodImplementation = new EigenPodMock();
        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));

        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSMock,
            eigenPodBeacon,
            strategyManagerMock,
            slasherMock,
            delegationManagerMock
        );
        eigenPodManager = EigenPodManager(
            address(
                new TransparentUpgradeableProxy(
                    address(eigenPodManagerImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        EigenPodManager.initialize.selector,
                        type(uint256).max /*maxPods*/,
                        IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
                        initialOwner,
                        pauserRegistry,
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );

        beaconChainETHStrategy = eigenPodManager.beaconChainETHStrategy();

        // excude the zero address, the proxyAdmin and the eigenPodManager itself from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManager)] = true;
    }

    function testRestakeBeaconChainETHSuccessfully(address staker, uint256 amount) public filterFuzzedAddressInputs(staker) {
        // filter out zero case since it will revert with "EigenPodManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);

        IEigenPod eigenPod = _deployEigenPodForStaker(staker);
        uint256 sharesBefore = eigenPodManager.podOwnerShares(staker);

        cheats.startPrank(address(eigenPod));
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BeaconChainETHDeposited(staker, amount);
        eigenPodManager.restakeBeaconChainETH(staker, amount);
        cheats.stopPrank();

        uint256 sharesAfter = eigenPodManager.podOwnerShares(staker);
        require(sharesAfter == sharesBefore + amount, "sharesAfter != sharesBefore + amount");
    }

    function testRestakeBeaconChainETHFailsWhenNotCalledByEigenPod(address improperCaller) public filterFuzzedAddressInputs(improperCaller) {
        uint256 amount = 1e18;
        address staker = address(this);

        IEigenPod eigenPod = _deployEigenPodForStaker(staker);
        cheats.assume(improperCaller != address(eigenPod));

        cheats.expectRevert(bytes("EigenPodManager.onlyEigenPod: not a pod"));
        cheats.startPrank(address(improperCaller));
        eigenPodManager.restakeBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testRestakeBeaconChainETHFailsWhenDepositsPaused() public {
        uint256 amount = 1e18;
        address staker = address(this);
        IEigenPod eigenPod = _deployEigenPodForStaker(staker);

        // pause deposits
        cheats.startPrank(pauser);
        eigenPodManager.pause(2 ** PAUSED_DEPOSITS);
        cheats.stopPrank();

        cheats.startPrank(address(eigenPod));
        cheats.expectRevert(bytes("Pausable: index is paused"));
        eigenPodManager.restakeBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testRestakeBeaconChainETHFailsWhenStakerFrozen() public {
        uint256 amount = 1e18;
        address staker = address(this);
        IEigenPod eigenPod = _deployEigenPodForStaker(staker);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.startPrank(address(eigenPod));
        cheats.expectRevert(bytes("EigenPodManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        eigenPodManager.restakeBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

// TODO: salvage / re-implement a check for reentrancy guard on functions, as possible
    // function testRestakeBeaconChainETHFailsWhenReentering() public {
    //     uint256 amount = 1e18;
    //     address staker = address(this);
    //     IEigenPod eigenPod = _deployEigenPodForStaker(staker);

    //     _beaconChainReentrancyTestsSetup();

    //     address targetToUse = address(eigenPodManager);
    //     uint256 msgValueToUse = 0;
    //     bytes memory calldataToUse = abi.encodeWithSelector(EigenPodManager.restakeBeaconChainETH.selector, staker, amount);
    //     reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

    //     // etch the EigenPod to instead contain Reenterer code
    //     vm.etch(address(eigenPod), address(reenterer).code);

    //     cheats.startPrank(address(eigenPod));
    //     eigenPodManager.restakeBeaconChainETH(staker, amount);
    //     cheats.stopPrank();
    // }

    function testRecordBeaconChainETHBalanceUpdateFailsWhenNotCalledByEigenPod(address improperCaller) public filterFuzzedAddressInputs(improperCaller) {
        address staker = address(this);
        IEigenPod eigenPod = _deployEigenPodForStaker(staker);
        cheats.assume(improperCaller != address(eigenPod));

        cheats.expectRevert(bytes("EigenPodManager.onlyEigenPod: not a pod"));
        cheats.startPrank(address(improperCaller));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(staker, int256(0));
        cheats.stopPrank();
    }

// TODO: salvage / re-implement a check for reentrancy guard on functions, as possible
    // function testRecordBeaconChainETHBalanceUpdateFailsWhenReentering() public {
    //     uint256 amount = 1e18;
    //     uint256 amount2 = 2e18;
    //     address staker = address(this);
    //     uint256 beaconChainETHStrategyIndex = 0;

    //     _beaconChainReentrancyTestsSetup();

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);        

    //     address targetToUse = address(strategyManager);
    //     uint256 msgValueToUse = 0;

    //     int256 amountDelta = int256(amount2 - amount);
    //     // reference: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 sharesDelta, bool isNegative)
    //     bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.recordBeaconChainETHBalanceUpdate.selector, staker, beaconChainETHStrategyIndex, amountDelta);
    //     reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

    //     cheats.startPrank(address(reenterer));
    //     eigenPodManager.recordBeaconChainETHBalanceUpdate(staker, amountDelta);
    //     cheats.stopPrank();
    // }

    // queues a withdrawal of "beacon chain ETH shares" from this address to itself
    // fuzzed input amountGwei is sized-down, since it must be in GWEI and gets sized-up to be WEI
    function testQueueWithdrawalBeaconChainETHToSelf(uint128 amountGwei)
        public returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    {
        // scale fuzzed amount up to be a whole amount of GWEI
        uint256 amount = uint256(amountGwei) * 1e9;
        address staker = address(this);
        address withdrawer = staker;

        testRestakeBeaconChainETHSuccessfully(staker, amount);

        // TODO: fuzz this param and check behavior
        bool undelegateIfPossible = false;
        (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
            _createQueuedWithdrawal(staker, amount, withdrawer, undelegateIfPossible);

        return (queuedWithdrawal, withdrawalRoot);
    }

    function testQueueWithdrawalBeaconChainETHToDifferentAddress(address withdrawer, uint128 amountGwei)
        public
        filterFuzzedAddressInputs(withdrawer)
        returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    {
        // scale fuzzed amount up to be a whole amount of GWEI
        uint256 amount = uint256(amountGwei) * 1e9;
        address staker = address(this);

        testRestakeBeaconChainETHSuccessfully(staker, amount);

        // TODO: fuzz this param and check behavior
        bool undelegateIfPossible = false;
        (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
            _createQueuedWithdrawal(staker, amount, withdrawer, undelegateIfPossible);

        return (queuedWithdrawal, withdrawalRoot);
    }

    function testQueueWithdrawalBeaconChainETHFailsNonWholeAmountGwei(uint256 nonWholeAmount) external {
        // this also filters out the zero case, which will revert separately
        cheats.assume(nonWholeAmount % GWEI_TO_WEI != 0);
        bool undelegateIfPossible = false;
        cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei"));
        eigenPodManager.queueWithdrawal(nonWholeAmount, address(this), undelegateIfPossible);
    }

    function testQueueWithdrawalBeaconChainETHFailsZeroAmount() external {
        bool undelegateIfPossible = false;
        cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: amount must be greater than zero"));
        eigenPodManager.queueWithdrawal(0, address(this), undelegateIfPossible);
    }

    function testCompleteQueuedWithdrawal() external {
        address staker = address(this);
        uint256 withdrawalAmount = 1e18;

        // withdrawalAmount is converted to GWEI here
        (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) = 
            testQueueWithdrawalBeaconChainETHToSelf(uint128(withdrawalAmount / 1e9));

        IEigenPod eigenPod = eigenPodManager.getPod(staker);
        uint256 eigenPodBalanceBefore = address(eigenPod).balance;

        uint256 middlewareTimesIndex = 0;

        // actually complete the withdrawal
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BeaconChainETHWithdrawalCompleted(
            queuedWithdrawal.podOwner,
            queuedWithdrawal.shares,
            queuedWithdrawal.nonce,
            queuedWithdrawal.delegatedAddress,
            queuedWithdrawal.withdrawer,
            withdrawalRoot
        );
        eigenPodManager.completeQueuedWithdrawal(queuedWithdrawal, middlewareTimesIndex);
        cheats.stopPrank();

        // TODO: make EigenPodMock do something so we can verify that it gets called appropriately?
        uint256 eigenPodBalanceAfter = address(eigenPod).balance;

        // verify that the withdrawal root does bit exist after queuing
        require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    }

// TODO: update tests from here
    function testSlashSharesBeaconChainETH() external {
        uint256 amount = 1e18;
        address staker = address(this);

        testRestakeBeaconChainETHSuccessfully(staker, amount);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = staker;
        address recipient = address(333);

        cheats.startPrank(eigenPodManager.owner());
        eigenPodManager.slashShares(slashedAddress, recipient, amount);
        cheats.stopPrank();

        // TODO: add before/after checks!
    }

    // INTERNAL / HELPER FUNCTIONS
    // deploy an EigenPod for the staker and check the emitted event
    function _deployEigenPodForStaker(address staker) internal returns (IEigenPod deployedPod) {
        deployedPod = eigenPodManager.getPod(staker);
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit PodDeployed(address(deployedPod), staker);
        eigenPodManager.createPod();
        cheats.stopPrank();
        return deployedPod;
    }


    // creates a queued withdrawal of "beacon chain ETH shares", from `staker`, of `amountWei`, "to" the `withdrawer`, passing param `undelegateIfPossible`
    function _createQueuedWithdrawal(address staker, uint256 amountWei, address withdrawer, bool undelegateIfPossible)
        internal
        returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    {
        // create the struct, for reference / to return
        queuedWithdrawal = IEigenPodManager.BeaconChainQueuedWithdrawal({
            shares: amountWei,
            podOwner: staker,
            nonce: uint96(eigenPodManager.numWithdrawalsQueued(staker)),
            withdrawalStartBlock: uint32(block.number),
            delegatedAddress: delegationManagerMock.delegatedTo(staker),
            withdrawer: withdrawer
        });

        // verify that the withdrawal root does not exist before queuing
        require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

        // get staker nonce and shares before queuing
        uint256 nonceBefore = eigenPodManager.numWithdrawalsQueued(staker);
        uint256 sharesBefore = eigenPodManager.podOwnerShares(staker);

        // actually create the queued withdrawal, and check for event emission
        cheats.startPrank(staker);
    
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BeaconChainETHWithdrawalQueued(
            queuedWithdrawal.podOwner,
            queuedWithdrawal.shares,
            queuedWithdrawal.nonce,
            queuedWithdrawal.delegatedAddress,
            queuedWithdrawal.withdrawer,
            eigenPodManager.calculateWithdrawalRoot(queuedWithdrawal)
        );
        withdrawalRoot = eigenPodManager.queueWithdrawal(amountWei, withdrawer, undelegateIfPossible);
        cheats.stopPrank();

        // verify that the withdrawal root does exist after queuing
        require(eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is false!");

        // verify that staker nonce incremented correctly and shares decremented correctly
        uint256 nonceAfter = eigenPodManager.numWithdrawalsQueued(staker);
        uint256 sharesAfter = eigenPodManager.podOwnerShares(staker);
        require(nonceAfter == nonceBefore + 1, "nonce did not increment correctly on queuing withdrawal");
        require(sharesAfter + amountWei == sharesBefore, "shares did not decrement correctly on queuing withdrawal");

        return (queuedWithdrawal, withdrawalRoot);
    }

    function _beaconChainReentrancyTestsSetup() internal {
        // prepare EigenPodManager with StrategyManager and Delegation replaced with a Reenterer contract
        reenterer = new Reenterer();
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSMock,
            eigenPodBeacon,
            IStrategyManager(address(reenterer)),
            slasherMock,
            IDelegationManager(address(reenterer))
        );
        eigenPodManager = EigenPodManager(
            address(
                new TransparentUpgradeableProxy(
                    address(eigenPodManagerImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        EigenPodManager.initialize.selector,
                        type(uint256).max /*maxPods*/,
                        IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
                        initialOwner,
                        pauserRegistry,
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );
    }
}