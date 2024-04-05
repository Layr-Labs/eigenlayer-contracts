// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPodPausingConstants.sol";

import "src/test/events/IEigenPodManagerEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/harnesses/EigenPodManagerWrapper.sol";
import "src/test/mocks/EigenPodMock.sol";
import "src/test/mocks/ETHDepositMock.sol";

contract EigenPodManagerUnitTests is EigenLayerUnitTestSetup {
    // Contracts Under Test: EigenPodManager
    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManager public eigenPodManager;

    using stdStorage for StdStorage;

    // Mocks
    IETHPOSDeposit public ethPOSMock;
    IEigenPod public eigenPodMockImplementation;
    IBeacon public eigenPodBeacon; // Proxy for eigenPodMockImplementation
    
    // Constants
    uint256 public constant GWEI_TO_WEI = 1e9;
    address public defaultStaker = address(this);
    IEigenPod public defaultPod;
    address public initialOwner = address(this);

    function setUp() virtual override public {
        EigenLayerUnitTestSetup.setUp();

        // Deploy Mocks
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
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        EigenPodManager.initialize.selector,
                        IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
                        initialOwner,
                        pauserRegistry,
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );

        // Set defaultPod
        defaultPod = eigenPodManager.getPod(defaultStaker);

        // Exclude the zero address, and the eigenPodManager itself from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManager)] = true;
    }

    /*******************************************************************************
                            Helper Functions/Modifiers
    *******************************************************************************/

    function _initializePodWithShares(address podOwner, int256 shares) internal {
        // Deploy pod
        IEigenPod deployedPod = _deployAndReturnEigenPodForStaker(podOwner);
        
        // Set shares
        cheats.prank(address(deployedPod));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(podOwner, shares);
    }


    modifier deployPodForStaker(address staker) {
        _deployAndReturnEigenPodForStaker(staker);
        _;
    }

    function _deployAndReturnEigenPodForStaker(address staker) internal returns (IEigenPod deployedPod) {
        deployedPod = eigenPodManager.getPod(staker);
        cheats.prank(staker);
        eigenPodManager.createPod();
        return deployedPod;
    }

    function _checkPodDeployed(address staker, address expectedPod, uint256 numPodsBefore) internal {
        assertEq(address(eigenPodManager.ownerToPod(staker)), expectedPod, "Expected pod not deployed");
        assertEq(eigenPodManager.numPods(), numPodsBefore + 1, "Num pods not incremented");
    }
}

contract EigenPodManagerUnitTests_Initialization_Setters is EigenPodManagerUnitTests, IEigenPodManagerEvents {

    /*******************************************************************************
                                Initialization Tests
    *******************************************************************************/

    function test_initialization() public {
        // Check max pods, beacon chain, owner, and pauser
        assertEq(address(eigenPodManager.beaconChainOracle()), address(IBeaconChainOracle(address(0))), "Initialization: beacon chain oracle incorrect");
        assertEq(eigenPodManager.owner(), initialOwner, "Initialization: owner incorrect");
        assertEq(address(eigenPodManager.pauserRegistry()), address(pauserRegistry), "Initialization: pauser registry incorrect");
        assertEq(eigenPodManager.paused(), 0, "Initialization: paused value not 0");

        // Check storage variables
        assertEq(address(eigenPodManager.ethPOS()), address(ethPOSMock), "Initialization: ethPOS incorrect");
        assertEq(address(eigenPodManager.eigenPodBeacon()), address(eigenPodBeacon), "Initialization: eigenPodBeacon incorrect");
        assertEq(address(eigenPodManager.strategyManager()), address(strategyManagerMock), "Initialization: strategyManager incorrect");
        assertEq(address(eigenPodManager.slasher()), address(slasherMock), "Initialization: slasher incorrect");
        assertEq(address(eigenPodManager.delegationManager()), address(delegationManagerMock), "Initialization: delegationManager incorrect");
    }

    function test_initialize_revert_alreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        eigenPodManager.initialize(
            IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
            initialOwner,
            pauserRegistry,
            0 /*initialPausedStatus*/);
    }

    /*******************************************************************************
                                     Setters
    *******************************************************************************/

    function testFuzz_updateBeaconChainOracle_revert_notOwner(address notOwner) public filterFuzzedAddressInputs(notOwner) {
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
        assertEq(address(eigenPodManager.beaconChainOracle()), address(newBeaconChainOracle), "Beacon chain oracle not updated");
    }

    function test_setDenebForkTimestamp(uint64 denebForkTimestamp) public {
        cheats.assume(denebForkTimestamp != 0);
        cheats.assume(denebForkTimestamp != type(uint64).max);
        cheats.prank(initialOwner);

        cheats.expectEmit(true, true, true, true);
        emit DenebForkTimestampUpdated(denebForkTimestamp);
        eigenPodManager.setDenebForkTimestamp(denebForkTimestamp);
        assertEq(eigenPodManager.denebForkTimestamp(), denebForkTimestamp, "fork timestamp not set correctly");
    }

    function test_setDenebForkTimestamp_Twice(uint64 timestamp1, uint64 timestamp2) public {
        cheats.assume(timestamp1 != 0);
        cheats.assume(timestamp2 != 0);
        cheats.assume(timestamp1 != type(uint64).max);
        cheats.assume(timestamp2 != type(uint64).max);
        cheats.prank(initialOwner);

        eigenPodManager.setDenebForkTimestamp(timestamp1);
        cheats.expectRevert(bytes("EigenPodManager.setDenebForkTimestamp: cannot set denebForkTimestamp more than once"));
        eigenPodManager.setDenebForkTimestamp(timestamp2);
    }
}

contract EigenPodManagerUnitTests_CreationTests is EigenPodManagerUnitTests, IEigenPodManagerEvents {

    function test_createPod() public {
        // Get expected pod address and pods before
        IEigenPod expectedPod = eigenPodManager.getPod(defaultStaker);
        uint256 numPodsBefore = eigenPodManager.numPods();

        // Create pod
        cheats.expectEmit(true, true, true, true);
        emit PodDeployed(address(expectedPod), defaultStaker);
        eigenPodManager.createPod();

        // Check pod deployed
        _checkPodDeployed(defaultStaker, address(defaultPod), numPodsBefore);
    }

    function test_createPod_revert_alreadyCreated() public deployPodForStaker(defaultStaker) {
        cheats.expectRevert("EigenPodManager.createPod: Sender already has a pod");
        eigenPodManager.createPod();
    }
}

contract EigenPodManagerUnitTests_StakeTests is EigenPodManagerUnitTests {

    function test_stake_podAlreadyDeployed() deployPodForStaker(defaultStaker) public {
        // Declare dummy variables
        bytes memory pubkey = bytes("pubkey");
        bytes memory sig = bytes("sig");
        bytes32 depositDataRoot = bytes32("depositDataRoot");

        // Stake
        eigenPodManager.stake{value: 32 ether}(pubkey, sig, depositDataRoot);

        // Expect pod has 32 ether
        assertEq(address(defaultPod).balance, 32 ether, "ETH not staked in EigenPod");
    }

    function test_stake_newPodDeployed() public {
        // Declare dummy variables
        bytes memory pubkey = bytes("pubkey");
        bytes memory sig = bytes("sig");
        bytes32 depositDataRoot = bytes32("depositDataRoot");

        // Stake
        eigenPodManager.stake{value: 32 ether}(pubkey, sig, depositDataRoot);

        // Check pod deployed
        _checkPodDeployed(defaultStaker, address(defaultPod), 0); // staker, defaultPod, numPodsBefore
        
        // Expect pod has 32 ether
        assertEq(address(defaultPod).balance, 32 ether, "ETH not staked in EigenPod");
    }
}

contract EigenPodManagerUnitTests_ShareUpdateTests is EigenPodManagerUnitTests {

    /*******************************************************************************
                                Add Shares Tests
    *******************************************************************************/

    function testFuzz_addShares_revert_notDelegationManager(address notDelegationManager) public filterFuzzedAddressInputs(notDelegationManager){
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
        shares = shares - (shares % GWEI_TO_WEI); // Round down to nearest Gwei
        cheats.assume(int256(shares) >= 0);

        // Add shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.addShares(defaultStaker, shares);

        // Check storage update
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(shares), "Incorrect number of shares added");
    }

    /*******************************************************************************
                                Remove Shares Tests
    ******************************************************************************/

    function testFuzz_removeShares_revert_notDelegationManager(address notDelegationManager) public filterFuzzedAddressInputs(notDelegationManager) {
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

    function testFuzz_removeShares_revert_tooManySharesRemoved(uint224 sharesToAdd, uint224 sharesToRemove) public {
        // Constrain inputs
        cheats.assume(sharesToRemove > sharesToAdd);
        uint256 sharesAdded = sharesToAdd * GWEI_TO_WEI;
        uint256 sharesRemoved = sharesToRemove * GWEI_TO_WEI;

        // Initialize pod with shares
        _initializePodWithShares(defaultStaker, int256(sharesAdded));

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.removeShares: cannot result in pod owner having negative shares");
        eigenPodManager.removeShares(defaultStaker, sharesRemoved);
    }

    function testFuzz_removeShares(uint224 sharesToAdd, uint224 sharesToRemove) public {
        // Constain inputs
        cheats.assume(sharesToRemove <= sharesToAdd);
        uint256 sharesAdded = sharesToAdd * GWEI_TO_WEI;
        uint256 sharesRemoved = sharesToRemove * GWEI_TO_WEI;

        // Initialize pod with shares
        _initializePodWithShares(defaultStaker, int256(sharesAdded));

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.removeShares(defaultStaker, sharesRemoved);

        // Check storage
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(sharesAdded - sharesRemoved), "Incorrect number of shares removed");
    }

    function testFuzz_removeShares_zeroShares(address podOwner, uint256 shares) public filterFuzzedAddressInputs(podOwner) {
        // Constrain inputs
        cheats.assume(podOwner != address(0));
        cheats.assume(shares < type(uint256).max / 2);
        shares = shares - (shares % GWEI_TO_WEI);  // Round down to nearest Gwei
        assertTrue(int256(shares) % int256(GWEI_TO_WEI) == 0, "Shares must be a whole Gwei amount");

        // Initialize pod with shares
        _initializePodWithShares(podOwner, int256(shares));

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.removeShares(podOwner, shares);

        // Check storage update
        assertEq(eigenPodManager.podOwnerShares(podOwner), 0, "Shares not reset to zero");
    }

    /*******************************************************************************
                        WithdrawSharesAsTokens Tests
    ******************************************************************************/

    function test_withdrawSharesAsTokens_revert_podOwnerZeroAddress() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.withdrawSharesAsTokens: podOwner cannot be zero address");
        eigenPodManager.withdrawSharesAsTokens(address(0), address(0), 0);
    }

    function test_withdrawSharesAsTokens_revert_destinationZeroAddress() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.withdrawSharesAsTokens: destination cannot be zero address");
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, address(0), 0);  
    }

    function testFuzz_withdrawSharesAsTokens_revert_sharesNegative(int256 shares) public {
        cheats.assume(shares < 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.withdrawSharesAsTokens: shares cannot be negative");
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, defaultStaker, uint256(shares));
    }

    function testFuzz_withdrawSharesAsTokens_revert_sharesNotWholeGwei(uint256 shares) public {
        cheats.assume(int256(shares) >= 0);
        cheats.assume(shares % GWEI_TO_WEI != 0);

        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert("EigenPodManager.withdrawSharesAsTokens: shares must be a whole Gwei amount");
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, defaultStaker, shares);
    }

    /**
     * @notice The `withdrawSharesAsTokens` is called in the `completeQueuedWithdrawal` function from the 
     *         delegationManager. When a withdrawal is queued in the delegationManager, `removeShares is called`
     */
    function test_withdrawSharesAsTokens_reduceEntireDeficit() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = -100e18;
        uint256 sharesToWithdraw = 101e18;
        
        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, defaultStaker, sharesToWithdraw);

        // Check storage update
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(0), "Shares not reduced to 0");
    }

    function test_withdrawSharesAsTokens_partialDefecitReduction() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = -100e18;
        uint256 sharesToWithdraw = 50e18;

        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, defaultStaker, sharesToWithdraw);

        // Check storage update
        int256 expectedShares = sharesBeginning + int256(sharesToWithdraw);
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), expectedShares, "Shares not reduced to expected amount");
    }

    function test_withdrawSharesAsTokens_withdrawPositive() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = 100e18;
        uint256 sharesToWithdraw = 50e18;

        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, defaultStaker, sharesToWithdraw);

        // Check storage remains the same
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), sharesBeginning, "Shares should not be adjusted");
    }
}

contract EigenPodManagerUnitTests_BeaconChainETHBalanceUpdateTests is EigenPodManagerUnitTests, IEigenPodManagerEvents {

    function testFuzz_recordBalanceUpdate_revert_notPod(address invalidCaller) public filterFuzzedAddressInputs(invalidCaller) deployPodForStaker(defaultStaker) {
        cheats.assume(invalidCaller != address(defaultPod));
        cheats.prank(invalidCaller);
        cheats.expectRevert("EigenPodManager.onlyEigenPod: not a pod");
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, 0);
    }

    function test_recordBalanceUpdate_revert_zeroAddress() public {
        IEigenPod zeroAddressPod = _deployAndReturnEigenPodForStaker(address(0));
        cheats.prank(address(zeroAddressPod));
        cheats.expectRevert("EigenPodManager.recordBeaconChainETHBalanceUpdate: podOwner cannot be zero address");
        eigenPodManager.recordBeaconChainETHBalanceUpdate(address(0), 0);
    }

    function testFuzz_recordBalanceUpdate_revert_nonWholeGweiAmount(int256 sharesDelta) public deployPodForStaker(defaultStaker) {
        cheats.assume(sharesDelta % int256(GWEI_TO_WEI) != 0);
        cheats.prank(address(defaultPod));
        cheats.expectRevert("EigenPodManager.recordBeaconChainETHBalanceUpdate: sharesDelta must be a whole Gwei amount");
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, sharesDelta);
    }

    function testFuzz_recordBalanceUpdateX(int224 sharesBefore, int224 sharesDelta) public {
        // Constrain inputs
        int256 scaledSharesBefore = sharesBefore * int256(GWEI_TO_WEI);
        int256 scaledSharesDelta = sharesDelta * int256(GWEI_TO_WEI);

        // Initialize shares
        _initializePodWithShares(defaultStaker, scaledSharesBefore);

        // Update balance
        cheats.expectEmit(true, true, true, true);
        emit PodSharesUpdated(defaultStaker, scaledSharesDelta);
        cheats.prank(address(defaultPod));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, scaledSharesDelta);

        // Check storage
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), scaledSharesBefore + scaledSharesDelta, "Shares not updated correctly");
    }
}

contract EigenPodManagerUnitTests_ShareAdjustmentCalculationTests is EigenPodManagerUnitTests {
    // Wrapper contract that exposes the internal `_calculateChangeInDelegatableShares` function
    EigenPodManagerWrapper public eigenPodManagerWrapper;

    function setUp() virtual override public {
        super.setUp();

        // Upgrade eigenPodManager to wrapper
        eigenPodManagerWrapper = new EigenPodManagerWrapper(
            ethPOSMock,
            eigenPodBeacon,
            strategyManagerMock,
            slasherMock,
            delegationManagerMock
        );
        eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(eigenPodManager))), address(eigenPodManagerWrapper));
    }

    function testFuzz_shareAdjustment_negativeToNegative(int256 sharesBefore, int256 sharesAfter) public {
        cheats.assume(sharesBefore <= 0);
        cheats.assume(sharesAfter <= 0);
        
        int256 sharesDelta = eigenPodManagerWrapper.calculateChangeInDelegatableShares(sharesBefore, sharesAfter);
        assertEq(sharesDelta, 0, "Shares delta must be 0");
    }

    function testFuzz_shareAdjustment_negativeToPositive(int256 sharesBefore, int256 sharesAfter) public {
        cheats.assume(sharesBefore <= 0);
        cheats.assume(sharesAfter > 0);
        
        int256 sharesDelta = eigenPodManagerWrapper.calculateChangeInDelegatableShares(sharesBefore, sharesAfter);
        assertEq(sharesDelta, sharesAfter, "Shares delta must be equal to sharesAfter");
    }

    function testFuzz_shareAdjustment_positiveToNegative(int256 sharesBefore, int256 sharesAfter) public {
        cheats.assume(sharesBefore > 0);
        cheats.assume(sharesAfter <= 0);
        
        int256 sharesDelta = eigenPodManagerWrapper.calculateChangeInDelegatableShares(sharesBefore, sharesAfter);
        assertEq(sharesDelta, -sharesBefore, "Shares delta must be equal to the negative of sharesBefore");
    }

    function testFuzz_shareAdjustment_positiveToPositive(int256 sharesBefore, int256 sharesAfter) public {
        cheats.assume(sharesBefore > 0);
        cheats.assume(sharesAfter > 0);
        
        int256 sharesDelta = eigenPodManagerWrapper.calculateChangeInDelegatableShares(sharesBefore, sharesAfter);
        assertEq(sharesDelta, sharesAfter - sharesBefore, "Shares delta must be equal to the difference between sharesAfter and sharesBefore");
    }
}
