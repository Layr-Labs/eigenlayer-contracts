// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodPausingConstants.sol";

import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/utils/ProofParsing.sol";
import "src/test/harnesses/EigenPodHarness.sol";
import "src/test/mocks/EigenPodMock.sol";
import "src/test/mocks/Dummy.sol";
import "src/test/mocks/ETHDepositMock.sol";
import "src/test/mocks/DelayedWithdrawalRouterMock.sol";
import "src/test/mocks/ERC20Mock.sol";
import "src/test/events/IEigenPodEvents.sol";
import "src/test/events/IEigenPodManagerEvents.sol";

contract EigenPod_PodManager_UnitTests is EigenLayerUnitTestSetup {
    // Contracts Under Test: EigenPodManager & EigenPod
    EigenPod public eigenPod;
    EigenPod public podImplementation;
    IBeacon public eigenPodBeacon;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;

    // Mocks
    IETHPOSDeposit public ethPOSMock;
    IDelayedWithdrawalRouter public delayedWithdrawalRouterMock;
    
    // Constants
    uint64 public constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    uint64 public constant GOERLI_GENESIS_TIME = 1616508000;
    address public initialOwner = address(this);

    // Owner for which proofs are generated; eigenPod above is owned by this address
    address public constant podManagerAddress = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;
    address public constant podOwner = address(42000094993494);
    address public constant expectedPodAddress = address(0x49c486E3f4303bc11C02F952Fe5b08D0AB22D443);

    function setUp() public override virtual {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy Mocks
        ethPOSMock = new ETHPOSDepositMock();
        delayedWithdrawalRouterMock = new DelayedWithdrawalRouterMock();

        // Deploy EigenPod Implementation and beacon
        podImplementation = new EigenPod(
            ethPOSMock,
            delayedWithdrawalRouterMock,
            eigenPodManagerMock,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        // Deploy EPM Implementation and Proxy
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
                        type(uint256).max /*maxPods*/,
                        IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
                        initialOwner,
                        pauserRegistry,
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );

        // Set EPM address
        bytes memory code = address(eigenPodManager).code;
        cheats.etch(podManagerAddress, code);
        eigenPodManager = EigenPodManager(podManagerAddress);

        // Deploy eigenPod
        cheats.prank(podOwner);
        eigenPod = EigenPod(payable(eigenPodManager.createPod()));
        // TODO: make sure eigenPod is deployed to address for which proofs are generated
    }
}

contract EigenPod_PodManager_UnitTests_EigenPodPausing is EigenPod_PodManager_UnitTests {
/**
 * 1. verifyBalanceUpdates revert when PAUSED_EIGENPODS_VERIFY_BALANCE_UPDATE set
 * 2. verifyAndProcessWithdrawals revert when PAUSED_EIGENPODS_VERIFY_WITHDRAWAL set
 * 3. verifyWithdrawalCredentials revert when PAUSED_EIGENPODS_VERIFY_CREDENTIALS set
 * 4. activateRestaking revert when PAUSED_EIGENPODS_VERIFY_CREDENTIALS set
 */
}

contract EigenPod_PodManager_UnitTests_EigenPod is EigenPod_PodManager_UnitTests {
/**
 * @notice Tests function calls from EPM to EigenPod
 * 1. Stake works when pod is deployed
 * 2. Stake when pod is not deployed -> check that ethPOS deposit contract is correct for this and above test
 * 3. Withdraw restakedBeaconChainETH -> ensure we do a full withdrawal before
 */
}

contract EigenPod_PodManager_UnitTests_EigenPodManager is EigenPod_PodManager_UnitTests {
/**
 * @notice Tests function calls from EigenPod to EigenPodManager
 * 1. Do a full withdrawal and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
 * 2. Do a partial withdrawal and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
 * 3. Verify balance updates and call `recordBeaconChainEThBalanceUpdate` -> assert shares are updated
 * 4. Verify withdrawal credentials and call `recordBeaconChainETHBalanceUpdate` -> assert shares are updated
 * 3. Reentrancy check (first function in below commented out tests)
 */
}

///@notice Placeholder for future unit tests that combine interactions between the EigenPod & EigenPodManager

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

    // function _beaconChainReentrancyTestsSetup() internal {
    //     // prepare EigenPodManager with StrategyManager and Delegation replaced with a Reenterer contract
    //     reenterer = new Reenterer();
    //     eigenPodManagerImplementation = new EigenPodManager(
    //         ethPOSMock,
    //         eigenPodBeacon,
    //         IStrategyManager(address(reenterer)),
    //         slasherMock,
    //         IDelegationManager(address(reenterer))
    //     );
    //     eigenPodManager = EigenPodManager(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(eigenPodManagerImplementation),
    //                 address(proxyAdmin),
    //                 abi.encodeWithSelector(
    //                     EigenPodManager.initialize.selector,
    //                     type(uint256).max /*maxPods*/,
    //                     IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
    //                     initialOwner,
    //                     pauserRegistry,
    //                     0 /*initialPausedStatus*/
    //                 )
    //             )
    //         )
    //     );
    // }
