// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/pods/EigenPod.sol";

import "../mocks/ETHPOSDepositMock.sol";
import "../mocks/BeaconChainOracleMock.sol";
import "../utils/EigenLayerUnitTestSetup.sol";

contract EigenPodUnitTests is EigenLayerUnitTestSetup {
    // Contract Under Test: EigenPod
    EigenPod public eigenPod;

    // Mocks
    IETHPOSDeposit public ethPOSDepositMock;
    IBeaconChainOracleMock public beaconChainOracleMock;
    
    // Constants
    uint32 public constant WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    uint64 public constant MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 31e9;
    uint64 public constant RESTAKED_BALANCE_OFFSET_GWEI = 75e7;
    uint64 public constant GOERLI_GENESIS_TIME = 1616508000;
    uint64 public constant SECONDS_PER_SLOT = 12;
    address public podOwner = address(this);

    function setUp() public override {
        // Setup
        EigenLayerUnitTestSetup.setUp();
        beaconChainOracle = new BeaconChainOracleMock();

        // Deploy mocks
        ethPOSDepositMock = new ETHPOSDepositMock();
        beaconChainOracleMock = new BeaconChainOracleMock();

        // Deploy EigenPod
        podImplementation = new EigenPod(
            ethPOSDepositMock,
            // DelayedWithdrawaRouterMock,
            eigenPodManagerMock,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        )
    }

}