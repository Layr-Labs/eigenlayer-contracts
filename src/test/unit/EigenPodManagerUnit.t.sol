// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "forge-std/Test.sol";

import "../../contracts/pods/EigenPodManager.sol";
import "../../contracts/pods/EigenPodPausingConstants.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../events/IEigenPodManagerEvents.sol";
import "../utils/EigenLayerUnitTestSetup.sol";
import "../mocks/DelegationManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodMock.sol";
import "../mocks/ETHDepositMock.sol";

contract EigenPodManagerUnitTests is EigenLayerUnitTestSetup {
    // Contracts Under Test: EigenPodManager
    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManager public eigenPodManager;

    using stdStorage for StdStorage;

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
                            Helper Functions/Modifiers
    *******************************************************************************/

    function _initializePodWithShares(address podOwner, int256 shares) internal {
        _deployAndReturnEigenPodForStaker(podOwner);
        // Signature of `podOwnerShares(address)
        bytes4 signature = 0x60f4062b;
        bytes32 sharesToSet = bytes32(uint256(shares));
        stdstore.target(address(eigenPodManager)).sig(signature).with_key(podOwner).checked_write(sharesToSet);
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

    function _checkPodDeployed(address staker, uint256 numPodsBefore) internal {
        // Get expected pod address
        IEigenPod expectedPod = eigenPodManager.getPod(staker);

        assertEq(address(eigenPodManager.ownerToPod(staker)), address(expectedPod));
        assertEq(eigenPodManager.numPods(), numPodsBefore + 1);
    }
}

contract EigenPodManagerUnitTests_Initialization_Setters is EigenPodManagerUnitTests, IEigenPodManagerEvents {

    /*******************************************************************************
                                Initialization Tests
    *******************************************************************************/

    function test_initialization() public {
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

    function test_initialize_revert_alreadyInitialized() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        eigenPodManager.initialize(type(uint256).max /*maxPods*/,
            IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
            initialOwner,
            pauserRegistry,
            0 /*initialPausedStatus*/);
    }

    function testFuzz_setMaxPods_revert_notUnpauser(address notUnpauser) public {
        cheats.assume(notUnpauser != unpauser);
        cheats.prank(notUnpauser);
        cheats.expectRevert("msg.sender is not permissioned as unpauser");
        eigenPodManager.setMaxPods(0);
    }

    /*******************************************************************************
                                     Setters
    *******************************************************************************/

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
        _checkPodDeployed(defaultStaker, numPodsBefore);
    }

    function test_createPod_revert_alreadyCreated() public deployPodForStaker(defaultStaker) {
        cheats.expectRevert("EigenPodManager.createPod: Sender already has a pod");
        eigenPodManager.createPod();
    }

    function test_createPod_revert_maxPodsCreated() public {
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
        assertEq(address(defaultPod).balance, 32 ether);
    }

    function test_stake_newPodDeployed() public {
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
}

contract EigenPodManagerUnitTests_ShareUpdateTests is EigenPodManagerUnitTests {

    /*******************************************************************************
                                Add Shares Tests
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

    /*******************************************************************************
                                Remove Shares Tests
    ******************************************************************************/

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

    function testFuzz_removeShares_zeroShares(address podOwner, uint256 shares) public {
        // Constrain inputs
        cheats.assume(podOwner != address(0));
        cheats.assume(shares % GWEI_TO_WEI == 0);
        cheats.assume(int256(shares) >= 0);

        // Add shares for user
        cheats.startPrank(address(delegationManagerMock));
        eigenPodManager.addShares(podOwner, shares);

        // Remove shares
        eigenPodManager.removeShares(podOwner, shares);
        cheats.stopPrank();

        // Check storage update
        assertEq(eigenPodManager.podOwnerShares(podOwner), 0);
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
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), int256(0));
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
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), expectedShares);
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
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), sharesBeginning);
    }
}

contract EigenPodManagerUnitTests_BeaconChainETHBalanceUpdateTests is EigenPodManagerUnitTests {

    function testFuzz_recordBalanceUpdate_revert_notPod(address invalidCaller) public deployPodForStaker(defaultStaker) {
        cheats.assume(invalidCaller != defaultStaker);
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

    function testFuzz_recordBalanceUpdate(int256 sharesBefore, int256 sharesDelta) public {
        // Constrain inputs
        cheats.assume(sharesDelta % int256(GWEI_TO_WEI) == 0);
        cheats.assume(sharesBefore + sharesDelta <= type(int256).max);

        // Initialize shares
        _initializePodWithShares(defaultStaker, sharesBefore);

        // Update balance
        cheats.prank(address(defaultPod));
        eigenPodManager.recordBeaconChainETHBalanceUpdate(defaultStaker, sharesDelta);

        // Check storage
        assertEq(eigenPodManager.podOwnerShares(defaultStaker), sharesBefore + sharesDelta);
    }
}