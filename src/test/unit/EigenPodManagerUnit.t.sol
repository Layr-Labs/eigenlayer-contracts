// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "forge-std/Test.sol";

import "../../contracts/pods/EigenPodManager.sol";
import "../../contracts/pods/EigenPodPausingConstants.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../utils/EigenLayerUnitTestSetup.sol";
import "../mocks/DelegationManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodMock.sol";
import "../mocks/ETHDepositMock.sol";
import "../mocks/Reenterer.sol";

contract EigenPodManagerUnitTests is EigenLayerUnitTestSetup {
    // Contracts Under Test: EigenPodManager
    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManager public eigenPodManager;

    // Proxy Admin & Pauser Registry
    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    // Mocks
    StrategyManagerMock public strategyManagerMock;
    DelegationManagerMock public delegationManagerMock;
    SlasherMock public slasherMock;
    IETHPOSDeposit public ethPOSMock;
    IEigenPod public eigenPodMockImplementation;
    IBeacon public eigenPodBeacon; // Proxy for eigenPodMockImplementation
    IStrategy public beaconChainETHStrategy;

    Reenterer public reenterer;

    // Constants
    uint256 public constant GWEI_TO_WEI = 1e9;
    uint256 public constant REQUIRED_BALANCE_WEI = 31 ether;
    address public defaultStaker = address(this);
    IEigenPod public defaultPod;
    address public initialOwner = address(this);

    function setUp() virtual public {
        // Deploy ProxyAdmin
        proxyAdmin = new ProxyAdmin();

        // Initialize PauserRegistry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        // Deploy Mocks
        slasherMock = new SlasherMock();
        delegationManagerMock = new DelegationManagerMock();
        strategyManagerMock = new StrategyManagerMock();
        ethPOSMock = new ETHPOSDepositMock();
        eigenPodMockImplementation = new EigenPodMock();
        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodMockImplementation));

        // Deploy EPM Implementation & Proxy
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

        // Set beaconChainETHStrategy
        beaconChainETHStrategy = eigenPodManager.beaconChainETHStrategy();

        // Set defaultPod
        defaultPod = eigenPodManager.getPod(defaultStaker);

        // Exclude the zero address, the proxyAdmin and the eigenPodManager itself from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManager)] = true;
    }

    /*******************************************************************************
                                 Initialization Tests
    *******************************************************************************/

    function test_initialization() external {
        // Check max pods, beacon chain, owner, and pauser
        assertEq(eigenPodManager.maxPods(), type(uint256).max);
        assertEq(address(eigenPodManager.beaconChainOracle()), address(IBeaconChainOracle(address(0))));
        assertEq(eigenPodManager.owner(), initialOwner);
        assertEq(address(eigenPodManager.pauserRegistry()), address(pauserRegistry));
        assertEq(eigenPodManager.paused(), 0);

        // Check storage variables
        assertEq(address(eigenPodManager.ethPOS()), address(ethPOSMock));
        assertEq(address(eigenPodManager.eigenPodBeacon()), address(eigenPodBeacon));
        assertEq(address(eigenPodManager.strategyManager()), address(strategyManagerMock));
        assertEq(address(eigenPodManager.slasher()), address(slasherMock));
        assertEq(address(eigenPodManager.delegationManager()), address(delegationManagerMock));
    }

    function test_initialize_revert_alreadyInitialized() external {
        cheats.expectRevert("Initializable: contract is already initialized");
        eigenPodManager.initialize(type(uint256).max /*maxPods*/,
            IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
            initialOwner,
            pauserRegistry,
            0 /*initialPausedStatus*/);
    }

    /*******************************************************************************
                                EigenPod Creation Tests
    *******************************************************************************/

    function test_createPod() external {
        // Get expected pod address and pods before
        IEigenPod expectedPod = eigenPodManager.getPod(defaultStaker);
        uint256 numPodsBefore = eigenPodManager.numPods();

        // Create pod
        cheats.expectEmit(true, true, true, true);
        emit PodDeployed(address(expectedPod), defaultStaker);
        eigenPodManager.createPod();

        // Check pod deployed
        _checkPodDeployed(defaultStaker, numPodsBefore);
    }

    function test_createPod_revert_alreadyCreated() external deployPodForStaker(defaultStaker) {
        cheats.expectRevert("EigenPodManager.createPod: Sender already has a pod");
        eigenPodManager.createPod();
    }

    function test_createPod_revert_maxPodsCreated() external {
        // Write numPods into storage. Num pods is at slot 153
        bytes32 slot = bytes32(uint256(153)); 
        bytes32 value = bytes32(eigenPodManager.maxPods());
        cheats.store(address(eigenPodManager), slot, value);

        //Create pod & expect revert based on maxPods size
        if (eigenPodManager.maxPods() == type(uint256).max) {
            // Arithmetic over/underflow not working with foundry
            cheats.expectRevert();
        } else {
            cheats.expectRevert("EigenPodManager._deployPod: pod limit reached");
        }
        eigenPodManager.createPod();
    }

    /*******************************************************************************
                                    Stake Tests
    *******************************************************************************/

    function test_stake_podAlreadyDeployed() deployPodForStaker(defaultStaker) external {
        // Declare dummy variables
        bytes memory pubkey = bytes("pubkey");
        bytes memory sig = bytes("sig");
        bytes32 depositDataRoot = bytes32("depositDataRoot");

        // Stake
        eigenPodManager.stake{value: 32 ether}(pubkey, sig, depositDataRoot);

        // Expect pod has 32 ether
        assertEq(address(defaultPod).balance, 32 ether);
    }

    function test_stake_newPodDeployed() external {
        // Declare dummy variables
        bytes memory pubkey = bytes("pubkey");
        bytes memory sig = bytes("sig");
        bytes32 depositDataRoot = bytes32("depositDataRoot");

        // Stake
        eigenPodManager.stake{value: 32 ether}(pubkey, sig, depositDataRoot);

        // Check pod deployed
        _checkPodDeployed(defaultStaker, 0); // staker, numPodsBefore
        
        // Expect pod has 32 ether
        assertEq(address(defaultPod).balance, 32 ether);
    }

    /*******************************************************************************
                                   Share Update Tests
    *******************************************************************************/
    
    function testFuzz_addShares_revert_notDelegationManager(address notDelegationManager) public {
        cheats.assume(notDelegationManager != address(delegationManagerMock));
        cheats.prank(notDelegationManager);
        cheats.expectRevert("EigenPodManager.onlyDelegationManager: not the DelegationManager");
        eigenPodManager.addShares(defaultStaker, 0);
    }

    function test_addShares_revert_podOwnerZeroAddress() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.addShares: podOwner cannot be zero address");
        eigenPodManager.addShares(address(0), 0);
    }

    function testFuzz_addShares_revert_sharesNegative(int256 shares) public {
        cheats.assume(shares < 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.addShares: shares cannot be negative");
        eigenPodManager.addShares(defaultStaker, uint256(shares));
    }

    function testFuzz_addShares_revert_sharesNotWholeGwei(uint256 shares) public {
        cheats.assume(int256(shares) >= 0);
        cheats.assume(shares % GWEI_TO_WEI != 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.addShares: shares must be a whole Gwei amount");
        eigenPodManager.addShares(defaultStaker, shares);
    }

    function testFuzz_addShares(uint256 shares) public {
        // Fuzz inputs
        cheats.assume(defaultStaker != address(0));
        cheats.assume(shares % GWEI_TO_WEI == 0);
        cheats.assume(int256(shares) >= 0);

        // Add shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.addShares(defaultStaker, shares);

        // Check storage update
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(shares));
    }

    // function testFuzz_addShares_shares

    function testFuzz_removeShares_revert_notDelegationManager(address notDelegationManager) public {
        cheats.assume(notDelegationManager != address(delegationManagerMock));
        cheats.prank(notDelegationManager);
        cheats.expectRevert("EigenPodManager.onlyDelegationManager: not the DelegationManager");
        eigenPodManager.removeShares(defaultStaker, 0);
    }

    function testFuzz_removeShares_revert_sharesNegative(int256 shares) public {
        cheats.assume(shares < 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.removeShares: shares cannot be negative");
        eigenPodManager.removeShares(defaultStaker, uint256(shares));
    }
    
    function testFuzz_removeShares_revert_sharesNotWholeGwei(uint256 shares) public {
        cheats.assume(int256(shares) >= 0);
        cheats.assume(shares % GWEI_TO_WEI != 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.removeShares: shares must be a whole Gwei amount");
        eigenPodManager.removeShares(defaultStaker, shares);
    }

    // Fuzzer rejects too many inputs, unit testsing
    function test_removeShares_revert_tooManySharesRemoved() public {
        uint256 sharesToRemove = GWEI_TO_WEI;

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.removeShares: cannot result in pod owner having negative shares");
        eigenPodManager.removeShares(defaultStaker, sharesToRemove);
    }

    // Fuzzer rejects too many inputs, unit testing
    function test_removeShares() public {
        uint256 sharesToAdd = GWEI_TO_WEI * 2;
        uint256 sharesToRemove = GWEI_TO_WEI;

        // Add shares
        cheats.startPrank(address(delegationManagerMock));
        eigenPodManager.addShares(defaultStaker, sharesToAdd);

        // Remove shares
        eigenPodManager.removeShares(defaultStaker, sharesToRemove);
        cheats.stopPrank();

        // Check storage
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(sharesToAdd - sharesToRemove));
    }

    // function testFuzz_removeShares(address podOwner, uint256 shares) public {
    //     // Fuzz inputs
    //     cheats.assume(podOwner != address(0));
    //     cheats.assume(shares % GWEI_TO_WEI == 0);
    //     cheats.assume(int256(shares) >= 0);

    //     // Add shares for user
    //     cheats.startPrank(address(delegationManagerMock));
    //     eigenPodManager.addShares(podOwner, shares);

    //     // Remove shares
    //     eigenPodManager.removeShares(podOwner, shares);
    //     cheats.stopPrank();

    //     // Check storage update
    //     assertEq(eigenPodManager.podOwnerShares(podOwner), -int256(shares));
    // }



    /*******************************************************************************
                                    Setter Tests
    *******************************************************************************/

    function testFuzz_setMaxPods_revert_notUnpauser(address notUnpauser) public {
        cheats.assume(notUnpauser != unpauser);
        cheats.prank(notUnpauser);
        cheats.expectRevert("msg.sender is not permissioned as unpauser");
        eigenPodManager.setMaxPods(0);
    }

    function test_setMaxPods() public {
        // Set max pods
        uint256 newMaxPods = 0;
        cheats.expectEmit(true, true, true, true);
        emit MaxPodsUpdated(eigenPodManager.maxPods(), newMaxPods);
        cheats.prank(unpauser);
        eigenPodManager.setMaxPods(newMaxPods);

        // Check storage update
        assertEq(eigenPodManager.maxPods(), newMaxPods);
    }

    function testFuzz_updateBeaconChainOracle_revert_notOwner(address notOwner) public {
        cheats.assume(notOwner != initialOwner);
        cheats.prank(notOwner);
        cheats.expectRevert("Ownable: caller is not the owner");
        eigenPodManager.updateBeaconChainOracle(IBeaconChainOracle(address(1)));
    }

    function test_updateBeaconChainOracle() public {
        // Set new beacon chain oracle
        IBeaconChainOracle newBeaconChainOracle = IBeaconChainOracle(address(1));
        cheats.prank(initialOwner);
        cheats.expectEmit(true, true, true, true);
        emit BeaconOracleUpdated(address(newBeaconChainOracle));
        eigenPodManager.updateBeaconChainOracle(newBeaconChainOracle);

        // Check storage update
        assertEq(address(eigenPodManager.beaconChainOracle()), address(newBeaconChainOracle));
    }

    function testRecordBeaconChainETHBalanceUpdateFailsWhenNotCalledByEigenPod(address improperCaller) public filterFuzzedAddressInputs(improperCaller) {
        IEigenPod eigenPod = _deployAndReturnEigenPodForStaker(defaultStaker);
        cheats.assume(improperCaller != address(eigenPod));

        cheats.expectRevert(bytes("EigenPodManager.onlyEigenPod: not a pod"));
        cheats.prank(address(improperCaller));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, int256(0));
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
// TODO: reimplement similar test
    // function testQueueWithdrawalBeaconChainETHToSelf(uint128 amountGwei)
    //     public returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    // {
    //     // scale fuzzed amount up to be a whole amount of GWEI
    //     uint256 amount = uint256(amountGwei) * 1e9;
    //     address staker = address(this);
    //     address withdrawer = staker;

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);

    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
    //         _createQueuedWithdrawal(staker, amount, withdrawer);

    //     return (queuedWithdrawal, withdrawalRoot);
    // }
// TODO: reimplement similar test
    // function testQueueWithdrawalBeaconChainETHToDifferentAddress(address withdrawer, uint128 amountGwei)
    //     public
    //     filterFuzzedAddressInputs(withdrawer)
    //     returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    // {
    //     // scale fuzzed amount up to be a whole amount of GWEI
    //     uint256 amount = uint256(amountGwei) * 1e9;
    //     address staker = address(this);

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);

    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
    //         _createQueuedWithdrawal(staker, amount, withdrawer);

    //     return (queuedWithdrawal, withdrawalRoot);
    // }
// TODO: reimplement similar test

    // function testQueueWithdrawalBeaconChainETHFailsNonWholeAmountGwei(uint256 nonWholeAmount) external {
    //     // this also filters out the zero case, which will revert separately
    //     cheats.assume(nonWholeAmount % GWEI_TO_WEI != 0);
    //     cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei"));
    //     eigenPodManager.queueWithdrawal(nonWholeAmount, address(this));
    // }

    // function testQueueWithdrawalBeaconChainETHFailsZeroAmount() external {
    //     cheats.expectRevert(bytes("EigenPodManager._queueWithdrawal: amount must be greater than zero"));
    //     eigenPodManager.queueWithdrawal(0, address(this));
    // }

// TODO: reimplement similar test
    // function testCompleteQueuedWithdrawal() external {
    //     address staker = address(this);
    //     uint256 withdrawalAmount = 1e18;

    //     // withdrawalAmount is converted to GWEI here
    //     (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) = 
    //         testQueueWithdrawalBeaconChainETHToSelf(uint128(withdrawalAmount / 1e9));

    //     IEigenPod eigenPod = eigenPodManager.getPod(staker);
    //     uint256 eigenPodBalanceBefore = address(eigenPod).balance;

    //     uint256 middlewareTimesIndex = 0;

    //     // actually complete the withdrawal
    //     cheats.startPrank(staker);
    //     cheats.expectEmit(true, true, true, true, address(eigenPodManager));
    //     emit BeaconChainETHWithdrawalCompleted(
    //         queuedWithdrawal.podOwner,
    //         queuedWithdrawal.shares,
    //         queuedWithdrawal.nonce,
    //         queuedWithdrawal.delegatedAddress,
    //         queuedWithdrawal.withdrawer,
    //         withdrawalRoot
    //     );
    //     eigenPodManager.completeQueuedWithdrawal(queuedWithdrawal, middlewareTimesIndex);
    //     cheats.stopPrank();

    //     // TODO: make EigenPodMock do something so we can verify that it gets called appropriately?
    //     uint256 eigenPodBalanceAfter = address(eigenPod).balance;

    //     // verify that the withdrawal root does bit exist after queuing
    //     require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    // }

    // INTERNAL / HELPER FUNCTIONS
    // deploy an EigenPod for the staker and check the emitted event

    // TODO: reimplement similar test
    // // creates a queued withdrawal of "beacon chain ETH shares", from `staker`, of `amountWei`, "to" the `withdrawer`
    // function _createQueuedWithdrawal(address staker, uint256 amountWei, address withdrawer)
    //     internal
    //     returns (IEigenPodManager.BeaconChainQueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    // {
    //     // create the struct, for reference / to return
    //     queuedWithdrawal = IEigenPodManager.BeaconChainQueuedWithdrawal({
    //         shares: amountWei,
    //         podOwner: staker,
    //         nonce: eigenPodManager.cumulativeWithdrawalsQueued(staker),
    //         startBlock: uint32(block.number),
    //         delegatedTo: delegationManagerMock.delegatedTo(staker),
    //         withdrawer: withdrawer
    //     });

    //     // verify that the withdrawal root does not exist before queuing
    //     require(!eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

    //     // get staker nonce and shares before queuing
    //     uint256 nonceBefore = eigenPodManager.cumulativeWithdrawalsQueued(staker);
    //     int256 sharesBefore = eigenPodManager.podOwnerShares(staker);

    //     // actually create the queued withdrawal, and check for event emission
    //     cheats.startPrank(staker);
    
    //     cheats.expectEmit(true, true, true, true, address(eigenPodManager));
    //     emit BeaconChainETHWithdrawalQueued(
    //         queuedWithdrawal.podOwner,
    //         queuedWithdrawal.shares,
    //         queuedWithdrawal.nonce,
    //         queuedWithdrawal.delegatedAddress,
    //         queuedWithdrawal.withdrawer,
    //         eigenPodManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     withdrawalRoot = eigenPodManager.queueWithdrawal(amountWei, withdrawer);
    //     cheats.stopPrank();

    //     // verify that the withdrawal root does exist after queuing
    //     require(eigenPodManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is false!");

    //     // verify that staker nonce incremented correctly and shares decremented correctly
    //     uint256 nonceAfter = eigenPodManager.cumulativeWithdrawalsQueued(staker);
    //     int256 sharesAfter = eigenPodManager.podOwnerShares(staker);
    //     require(nonceAfter == nonceBefore + 1, "nonce did not increment correctly on queuing withdrawal");
    //     require(sharesAfter + amountWei == sharesBefore, "shares did not decrement correctly on queuing withdrawal");

    //     return (queuedWithdrawal, withdrawalRoot);
    // }

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

    function _checkPodDeployed(address staker, uint256 numPodsBefore) internal {
        // Get expected pod address
        IEigenPod expectedPod = eigenPodManager.getPod(staker);

        assertEq(address(eigenPodManager.ownerToPod(staker)), address(expectedPod));
        assertEq(eigenPodManager.numPods(), numPodsBefore + 1);
    }

    function _deployAndReturnEigenPodForStaker(address staker) internal returns (IEigenPod deployedPod) {
        deployedPod = eigenPodManager.getPod(staker);
        cheats.prank(staker);
        eigenPodManager.createPod();
        return deployedPod;
    }

    modifier deployPodForStaker(address staker) {
        cheats.prank(staker);
        eigenPodManager.createPod();
        _;
    }

    /*******************************************************************************
                                       EVENTS
    *******************************************************************************/
    
    /// @notice Emitted to notify the update of the beaconChainOracle address
    event BeaconOracleUpdated(address indexed newOracleAddress);

    /// @notice Emitted to notify the deployment of an EigenPod
    event PodDeployed(address indexed eigenPod, address indexed podOwner);

    /// @notice Emitted when `maxPods` value is updated from `previousValue` to `newValue`
    event MaxPodsUpdated(uint256 previousValue, uint256 newValue);
}