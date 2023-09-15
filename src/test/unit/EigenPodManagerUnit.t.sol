// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "forge-std/Test.sol";

import "../../contracts/pods/EigenPodManager.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodMock.sol";
import "../mocks/ETHDepositMock.sol";
import "../mocks/Reenterer.sol";
import "../mocks/Reverter.sol";

contract EigenPodManagerUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 public REQUIRED_BALANCE_WEI = 31 ether;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    EigenPodManager public eigenPodManagerImplementation;
    EigenPodManager public eigenPodManager;

    StrategyManagerMock public strategyManagerMock;
    DelegationMock public delegationMock;
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
    event BeaconChainETHWithdrawalQueued(address indexed podOwner, uint256 amount, uint96 nonce);

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
        delegationMock = new DelegationMock();
        strategyManagerMock = new StrategyManagerMock();
        ethPOSMock = new ETHPOSDepositMock();
        eigenPodImplementation = new EigenPodMock();
        eigenPodBeacon = new UpgradeableBeacon(address(eigenPodImplementation));

        eigenPodManagerImplementation = new EigenPodManager(
            ethPOSMock,
            eigenPodBeacon,
            strategyManagerMock,
            slasherMock,
            delegationMock
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
}