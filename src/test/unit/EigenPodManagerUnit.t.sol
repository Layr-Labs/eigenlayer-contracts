// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPodPausingConstants.sol";

import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/harnesses/EigenPodManagerWrapper.sol";
import "src/test/mocks/EigenPodMock.sol";
import "src/test/mocks/ETHDepositMock.sol";

contract EigenPodManagerUnitTests is EigenLayerUnitTestSetup, IEigenPodManagerEvents {
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

    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

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
            IDelegationManager(address(delegationManagerMock)),
            pauserRegistry
        );
        eigenPodManager = EigenPodManager(
            address(
                new TransparentUpgradeableProxy(
                    address(eigenPodManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        EigenPodManager.initialize.selector,
                        initialOwner,
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );

        // Set defaultPod
        defaultPod = eigenPodManager.getPod(defaultStaker);

        // Exclude the zero address, and the eigenPodManager itself from fuzzed inputs
        isExcludedFuzzAddress[address(0)] = true;
        isExcludedFuzzAddress[address(eigenPodManager)] = true;
    }

    /*******************************************************************************
                            Helper Functions/Modifiers
    *******************************************************************************/

    function _initializePodWithShares(address podOwner, int256 shares) internal {
        // Deploy pod
        _deployAndReturnEigenPodForStaker(podOwner);
        
        if (shares >= 0) {
            cheats.prank(address(delegationManagerMock));
            eigenPodManager.addShares(podOwner, beaconChainETHStrategy, IERC20(address(0)), uint256(shares));
        } else {
            EigenPodManagerWrapper(address(eigenPodManager)).setPodOwnerShares(podOwner, shares);
        }
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

    function _checkPodDeployed(address staker, address expectedPod, uint256 numPodsBefore) internal view {
        assertEq(address(eigenPodManager.ownerToPod(staker)), expectedPod, "Expected pod not deployed");
        assertEq(eigenPodManager.numPods(), numPodsBefore + 1, "Num pods not incremented");
    }
}

contract EigenPodManagerUnitTests_Initialization_Setters is EigenPodManagerUnitTests {

    /*******************************************************************************
                                Initialization Tests
    *******************************************************************************/

    function test_initialization() public view {
        // Check max pods, beacon chain, owner, and pauser
        assertEq(eigenPodManager.owner(), initialOwner, "Initialization: owner incorrect");
        assertEq(address(eigenPodManager.pauserRegistry()), address(pauserRegistry), "Initialization: pauser registry incorrect");
        assertEq(eigenPodManager.paused(), 0, "Initialization: paused value not 0");

        // Check storage variables
        assertEq(address(eigenPodManager.ethPOS()), address(ethPOSMock), "Initialization: ethPOS incorrect");
        assertEq(address(eigenPodManager.eigenPodBeacon()), address(eigenPodBeacon), "Initialization: eigenPodBeacon incorrect");
        assertEq(address(eigenPodManager.delegationManager()), address(delegationManagerMock), "Initialization: delegationManager incorrect");
    }

    function test_initialize_revert_alreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        eigenPodManager.initialize(
            initialOwner,
            0 /*initialPausedStatus*/);
    }
}

contract EigenPodManagerUnitTests_CreationTests is EigenPodManagerUnitTests {

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
        cheats.expectRevert(IEigenPodManagerErrors.EigenPodAlreadyExists.selector);
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
        cheats.expectRevert(IEigenPodManagerErrors.OnlyDelegationManager.selector);
        eigenPodManager.addShares(defaultStaker, IStrategy(address(0)), IERC20(address(0)), 0);
    }
    
    function test_addShares_revert_podOwnerZeroAddress() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodErrors.InputAddressZero.selector);
        eigenPodManager.addShares(address(0), beaconChainETHStrategy, IERC20(address(0)), 0);
    }

    function testFuzz_addShares_revert_sharesNegative(int256 shares) public {
        cheats.assume(shares < 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodManagerErrors.SharesNegative.selector);
        eigenPodManager.addShares(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), uint256(shares));
    }

    function testFuzz_addShares(uint256 shares) public {
        // Fuzz inputs
        cheats.assume(defaultStaker != address(0));
        shares = shares - (shares % GWEI_TO_WEI); // Round down to nearest Gwei
        cheats.assume(int256(shares) >= 0);

        // Add shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.addShares(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), shares);

        // Check storage update
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), int256(shares), "Incorrect number of shares added");
    }

    /*******************************************************************************
                                Remove Shares Tests
    ******************************************************************************/

    function testFuzz_removeDepositShares_revert_notDelegationManager(address notDelegationManager) public filterFuzzedAddressInputs(notDelegationManager) {
        cheats.assume(notDelegationManager != address(delegationManagerMock));
        cheats.prank(notDelegationManager);
        cheats.expectRevert(IEigenPodManagerErrors.OnlyDelegationManager.selector);
        eigenPodManager.removeDepositShares(defaultStaker, beaconChainETHStrategy, 0);
    }

    function testFuzz_removeDepositShares_revert_sharesNegative(uint224 sharesToRemove) public {
        cheats.assume(sharesToRemove > 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodManagerErrors.SharesNegative.selector);
        eigenPodManager.removeDepositShares(defaultStaker, beaconChainETHStrategy, sharesToRemove);
    }

    function testFuzz_removeDepositShares_revert_tooManySharesRemoved(uint224 sharesToAdd, uint224 sharesToRemove) public {
        // Constrain inputs
        cheats.assume(sharesToRemove > sharesToAdd);
        uint256 sharesAdded = sharesToAdd * GWEI_TO_WEI;
        uint256 sharesRemoved = sharesToRemove * GWEI_TO_WEI;

        // Initialize pod with shares
        _initializePodWithShares(defaultStaker,  int256(sharesAdded));

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodManagerErrors.SharesNegative.selector);
        eigenPodManager.removeDepositShares(defaultStaker,  beaconChainETHStrategy, sharesRemoved);
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
        eigenPodManager.removeDepositShares(defaultStaker,  beaconChainETHStrategy, sharesRemoved);

        // Check storage
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), int256(sharesAdded - sharesRemoved), "Incorrect number of shares removed");
    }

    function testFuzz_removeDepositShares_zeroShares(address podOwner, uint256 shares) public filterFuzzedAddressInputs(podOwner) {
        // Constrain inputs
        cheats.assume(podOwner != address(0));
        cheats.assume(shares < type(uint256).max / 2);
        shares = shares - (shares % GWEI_TO_WEI);  // Round down to nearest Gwei
        assertTrue(int256(shares) % int256(GWEI_TO_WEI) == 0, "Shares must be a whole Gwei amount");

        // Initialize pod with shares
        _initializePodWithShares(podOwner, int256(shares));

        // Remove shares
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.removeDepositShares(podOwner,  beaconChainETHStrategy, shares);

        // Check storage update
        assertEq(eigenPodManager.podOwnerDepositShares(podOwner), 0, "Shares not reset to zero");
    }
}

contract EigenPodManagerUnitTests_WithdrawSharesAsTokensTests is EigenPodManagerUnitTests {
    // Wrapper contract that exposes the internal `_calculateChangeInDelegatableShares` function
    EigenPodManagerWrapper public eigenPodManagerWrapper;

    function setUp() virtual override public {
        super.setUp();

        // Upgrade eigenPodManager to wrapper
        eigenPodManagerWrapper = new EigenPodManagerWrapper(
            ethPOSMock,
            eigenPodBeacon,
            IDelegationManager(address(delegationManagerMock)),
            pauserRegistry
        );
        eigenLayerProxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(eigenPodManager))), address(eigenPodManagerWrapper));
    }
    /*******************************************************************************
                        WithdrawSharesAsTokens Tests
    ******************************************************************************/

    function test_withdrawSharesAsTokens_revert_invalidStrategy() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodManagerErrors.InvalidStrategy.selector);
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, IStrategy(address(0)), IERC20(address(0)), 0);
    }

    function test_withdrawSharesAsTokens_revert_podOwnerZeroAddress() public {
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodErrors.InputAddressZero.selector);
        eigenPodManager.withdrawSharesAsTokens(address(0), beaconChainETHStrategy, IERC20(address(0)), 0);
    }

    function testFuzz_withdrawSharesAsTokens_revert_sharesNegative(int256 shares) public {
        cheats.assume(shares < 0);
        cheats.prank(address(delegationManagerMock));
        cheats.expectRevert(IEigenPodManagerErrors.SharesNegative.selector);
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), uint256(shares));
    }

    /**
     * @notice The `withdrawSharesAsTokens` is called in the `completeQueuedWithdrawal` function from the 
     *         delegationManager. When a withdrawal is queued in the delegationManager, `removeDepositShares is called`
     */
    function test_withdrawSharesAsTokens_m2NegativeShares_reduceEntireDeficit() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = -100e18;
        uint256 sharesToWithdraw = 101e18;
        
        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), sharesBeginning, "Shares not initialized correctly");

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit PodSharesUpdated(defaultStaker, 100e18);
        cheats.expectEmit(true, true, true, true);
        emit NewTotalShares(defaultStaker, 0);
        // Expect call to EigenPod for the withdrawal
        cheats.expectCall(
            address(defaultPod),
            abi.encodeWithSelector(IEigenPod.withdrawRestakedBeaconChainETH.selector, defaultStaker, 1e18),
            1
        );
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), sharesToWithdraw);

        // Check storage update
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), int256(0), "Shares not reduced to 0");
    }

    function test_withdrawSharesAsTokens_m2NegativeShares_partialDefecitReduction() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = -100e18;
        uint256 sharesToWithdraw = 50e18;

        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), sharesBeginning, "Shares not initialized correctly");

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true);
        emit PodSharesUpdated(defaultStaker, 50e18);
        cheats.expectEmit(true, true, true, true);
        emit NewTotalShares(defaultStaker, -50e18);
        // Assert that no call is made by passing in zero for the count
        bytes memory emptyBytes;
        cheats.expectCall(
            address(defaultPod), 
            emptyBytes, // Cheatcode checks a partial match starting at the first byte of the calldata
            0
        );
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), sharesToWithdraw);

        // Check storage update
        int256 expectedShares = sharesBeginning + int256(sharesToWithdraw);
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), expectedShares, "Shares not reduced to expected amount");
    }

    function test_withdrawSharesAsTokens_withdrawPositive() public {
        // Shares to initialize & withdraw
        int256 sharesBeginning = 100e18;
        uint256 sharesToWithdraw = 50e18;

        // Deploy Pod And initialize with negative shares
        _initializePodWithShares(defaultStaker, sharesBeginning);

        // Withdraw shares
        cheats.prank(address(delegationManagerMock));
        // Expect call to EigenPod for the withdrawal
        cheats.expectCall(
            address(defaultPod),
            abi.encodeWithSelector(IEigenPod.withdrawRestakedBeaconChainETH.selector, defaultStaker, sharesToWithdraw), 
            1
        );
        eigenPodManager.withdrawSharesAsTokens(defaultStaker, beaconChainETHStrategy, IERC20(address(0)), sharesToWithdraw);

        // Check storage remains the same
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), sharesBeginning, "Shares should not be adjusted");
    }
}

contract EigenPodManagerUnitTests_BeaconChainETHBalanceUpdateTests is EigenPodManagerUnitTests {
    // Wrapper contract that exposes the internal `_calculateChangeInDelegatableShares` function
    EigenPodManagerWrapper public eigenPodManagerWrapper;

    function setUp() virtual override public {
        super.setUp();

        // Upgrade eigenPodManager to wrapper
        eigenPodManagerWrapper = new EigenPodManagerWrapper(
            ethPOSMock,
            eigenPodBeacon,
            IDelegationManager(address(delegationManagerMock)),
            pauserRegistry
        );
        eigenLayerProxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(eigenPodManager))), address(eigenPodManagerWrapper));
    }

    function testFuzz_revert_notPod(address invalidCaller) public filterFuzzedAddressInputs(invalidCaller) deployPodForStaker(defaultStaker) {
        cheats.assume(invalidCaller != address(defaultPod));
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodManagerErrors.OnlyEigenPod.selector);
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, 0, 0);
    }

    function test_revert_zeroAddress() public {
        IEigenPod zeroAddressPod = _deployAndReturnEigenPodForStaker(address(0));
        cheats.prank(address(zeroAddressPod));
        cheats.expectRevert(IEigenPodErrors.InputAddressZero.selector);
        eigenPodManager.recordBeaconChainETHBalanceUpdate(address(0), 0, 0);
    }

    function testFuzz_revert_nonWholeGweiAmount(int256 sharesDelta) public deployPodForStaker(defaultStaker) {
        cheats.assume(sharesDelta % int256(GWEI_TO_WEI) != 0);
        cheats.prank(address(defaultPod));
        cheats.expectRevert(IEigenPodManagerErrors.SharesNotMultipleOfGwei.selector);
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, 0, sharesDelta);
    }

    function testFuzz_revert_negativeDepositShares(int224 sharesBefore) public {
        cheats.assume(sharesBefore < 0);
        
        // Initialize shares
        _initializePodWithShares(defaultStaker, sharesBefore);

        // Record balance update
        cheats.prank(address(defaultPod));
        cheats.expectRevert(IEigenPodManagerErrors.LegacyWithdrawalsNotCompleted.selector);
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, 0, 0);
    }

    function testFuzz_recordPositiveBalanceUpdate(
        uint256 sharesBefore, 
        uint256 sharesDelta,
        uint256 prevRestakedBalanceWei
    ) public {
        // Constrain inputs
        sharesBefore = bound(sharesBefore, 0, type(uint224).max) * uint(GWEI_TO_WEI);
        sharesDelta = bound(sharesDelta, 0, type(uint224).max) * uint(GWEI_TO_WEI);
        prevRestakedBalanceWei = bound(prevRestakedBalanceWei, 0, type(uint256).max);

        // Initialize shares
        _initializePodWithShares(defaultStaker, int(sharesBefore));

        uint64 prevSlashingFactor = eigenPodManager.beaconChainSlashingFactor(defaultStaker);

        // Add shares
        cheats.expectEmit(true, true, true, true);
        emit PodSharesUpdated(defaultStaker, int(sharesDelta));
        cheats.expectEmit(true, true, true, true);
        emit NewTotalShares(defaultStaker, int(sharesBefore + sharesDelta));

        cheats.prank(address(defaultPod));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, prevRestakedBalanceWei, int(sharesDelta));

        // Check storage
        // Note that this is a unit test, we don't validate that the withdrawable shares are updated correctly
        // See the integration tests for checking scaling factors and withdrawable shares
        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), int(sharesBefore + sharesDelta), "Shares not updated correctly");
        assertEq(eigenPodManager.beaconChainSlashingFactor(defaultStaker), prevSlashingFactor, "bcsf should not change");
    }

    function testFuzz_recordNegativeBalanceUpdate(
        uint256 sharesBefore,
        uint256 sharesDelta,
        uint256 prevRestakedBalanceWei
    ) public {
        // Constrain inputs
        sharesBefore = bound(sharesBefore, 0, type(uint224).max) * uint(GWEI_TO_WEI);
        prevRestakedBalanceWei = bound(prevRestakedBalanceWei, 1, type(uint224).max);
        sharesDelta = bound(sharesDelta, 1, prevRestakedBalanceWei) * uint(GWEI_TO_WEI);
        prevRestakedBalanceWei *= GWEI_TO_WEI;

        // Initialize shares
        _initializePodWithShares(defaultStaker, int(sharesBefore));

        uint64 prevSlashingFactor = eigenPodManager.beaconChainSlashingFactor(defaultStaker);

        // Not checking the new slashing factor - just checking the invariant that new <= prev
        cheats.expectEmit(true, true, true, false);
        emit BeaconChainSlashingFactorDecreased(defaultStaker, 0, 0);

        cheats.prank(address(defaultPod));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, prevRestakedBalanceWei, -int(sharesDelta));

        assertEq(eigenPodManager.podOwnerDepositShares(defaultStaker), int(sharesBefore), "Shares should not be adjusted");
        assertTrue(eigenPodManager.beaconChainSlashingFactor(defaultStaker) <= prevSlashingFactor, "bcsf should always decrease");
    }
}

contract EigenPodManagerUnitTests_increaseBurnableShares is EigenPodManagerUnitTests {
    function testFuzz_onlyDelegationManager(address invalidCaller) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(delegationManagerMock));
        cheats.prank(invalidCaller);
        cheats.expectRevert(IEigenPodManagerErrors.OnlyDelegationManager.selector);
        eigenPodManager.increaseBurnableShares(beaconChainETHStrategy, 1 ether);
    }

    function testFuzz_singleDeposit(uint256 increasedBurnableShares) public {
        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BurnableETHSharesIncreased(increasedBurnableShares);

        cheats.prank(address(delegationManagerMock));
        eigenPodManager.increaseBurnableShares(beaconChainETHStrategy, increasedBurnableShares);

        assertEq(eigenPodManager.burnableETHShares(), increasedBurnableShares, "Burnable shares not updated correctly");
    }

    function testFuzz_existingDeposit(
        uint256 existingBurnableShares,
        uint256 increasedBurnableShares
    ) public {
        // prevent overflow
        cheats.assume(existingBurnableShares < type(uint256).max - increasedBurnableShares);

        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BurnableETHSharesIncreased(existingBurnableShares);
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.increaseBurnableShares(beaconChainETHStrategy, existingBurnableShares);

        assertEq(eigenPodManager.burnableETHShares(), existingBurnableShares, "Burnable shares not setup correctly");

        cheats.expectEmit(true, true, true, true, address(eigenPodManager));
        emit BurnableETHSharesIncreased(increasedBurnableShares);
        cheats.prank(address(delegationManagerMock));
        eigenPodManager.increaseBurnableShares(beaconChainETHStrategy, increasedBurnableShares);

        assertEq(
            eigenPodManager.burnableETHShares(),
            existingBurnableShares + increasedBurnableShares, "Burnable shares not updated correctly"
        );
    }
}
