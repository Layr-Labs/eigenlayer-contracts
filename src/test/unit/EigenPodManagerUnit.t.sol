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

    function testDepositBeaconChainETHSuccessfully(address staker, uint256 amount) public filterFuzzedAddressInputs(staker) {
        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, beaconChainETHStrategy);

        cheats.startPrank(address(strategyManager.eigenPodManager()));
        strategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, beaconChainETHStrategy);
        require(sharesAfter == sharesBefore + amount, "sharesAfter != sharesBefore + amount");
    }

    function testDepositBeaconChainETHFailsWhenNotCalledByEigenPodManager(address improperCaller) public filterFuzzedAddressInputs(improperCaller) {
        uint256 amount = 1e18;
        address staker = address(this);

        cheats.expectRevert(bytes("StrategyManager.onlyEigenPodManager: not the eigenPodManager"));
        cheats.startPrank(address(improperCaller));
        strategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testDepositBeaconChainETHFailsWhenDepositsPaused() public {
        uint256 amount = 1e18;
        address staker = address(this);

        // pause deposits
        cheats.startPrank(pauser);
        strategyManager.pause(1);
        cheats.stopPrank();

        cheats.expectRevert(bytes("Pausable: index is paused"));
        cheats.startPrank(address(eigenPodManagerMock));
        strategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testDepositBeaconChainETHFailsWhenStakerFrozen() public {
        uint256 amount = 1e18;
        address staker = address(this);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        cheats.startPrank(address(eigenPodManagerMock));
        strategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testDepositBeaconChainETHFailsWhenReentering() public {
        uint256 amount = 1e18;
        address staker = address(this);

        _beaconChainReentrancyTestsSetup();

        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;
        bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.depositBeaconChainETH.selector, staker, amount);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.startPrank(address(reenterer));
        strategyManager.depositBeaconChainETH(staker, amount);
        cheats.stopPrank();
    }

    function testRecordOvercommittedBeaconChainETHFailsWhenNotCalledByEigenPodManager(address improperCaller) public filterFuzzedAddressInputs(improperCaller) {
        uint256 amount = 1e18;
        address staker = address(this);
        uint256 beaconChainETHStrategyIndex = 0;

        testDepositBeaconChainETHSuccessfully(staker, amount);

        cheats.expectRevert(bytes("StrategyManager.onlyEigenPodManager: not the eigenPodManager"));
        cheats.startPrank(address(improperCaller));
        strategyManager.recordBeaconChainETHBalanceUpdate(staker, beaconChainETHStrategyIndex, 0);
        cheats.stopPrank();
    }

    function testRecordBeaconChainETHBalanceUpdateFailsWhenReentering() public {
        uint256 amount = 1e18;
        uint256 amount2 = 2e18;
        address staker = address(this);
        uint256 beaconChainETHStrategyIndex = 0;

        _beaconChainReentrancyTestsSetup();

        testDepositBeaconChainETHSuccessfully(staker, amount);        

        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;

        int256 amountDelta = int256(amount2 - amount);
        // reference: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 sharesDelta, bool isNegative)
        bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.recordBeaconChainETHBalanceUpdate.selector, staker, beaconChainETHStrategyIndex, amountDelta);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.startPrank(address(reenterer));
        strategyManager.recordBeaconChainETHBalanceUpdate(staker, beaconChainETHStrategyIndex, amountDelta);
        cheats.stopPrank();
    }

    // fuzzed input amountGwei is sized-down, since it must be in GWEI and gets sized-up to be WEI
    function testQueueWithdrawalBeaconChainETHToSelf(uint128 amountGwei)
        public returns (IStrategyManager.QueuedWithdrawal memory, bytes32 /*withdrawalRoot*/) 
    {
        // scale fuzzed amount up to be a whole amount of GWEI
        uint256 amount = uint256(amountGwei) * 1e9;
        address staker = address(this);
        address withdrawer = staker;
        IStrategy strategy = beaconChainETHStrategy;
        IERC20 token;
        testDepositBeaconChainETHSuccessfully(staker, amount);
        bool undelegateIfPossible = false;
        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
            _setUpQueuedWithdrawalStructSingleStrat(staker, withdrawer, token, strategy, amount);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);
        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        {
            for (uint256 i = 0; i < queuedWithdrawal.strategies.length; ++i) {
                cheats.expectEmit(true, true, true, true, address(strategyManager));
                emit ShareWithdrawalQueued(
                    /*staker*/ address(this),
                    queuedWithdrawal.withdrawerAndNonce.nonce,
                    queuedWithdrawal.strategies[i],
                    queuedWithdrawal.shares[i]
                );                
            }
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit WithdrawalQueued(
                /*staker*/ address(this),
                queuedWithdrawal.withdrawerAndNonce.nonce,
                queuedWithdrawal.withdrawerAndNonce.withdrawer,
                queuedWithdrawal.delegatedAddress,
                withdrawalRoot
            );
        }
        strategyManager.queueWithdrawal(strategyIndexes, queuedWithdrawal.strategies, queuedWithdrawal.shares, withdrawer, undelegateIfPossible);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.numWithdrawalsQueued(staker);
        require(strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is false!");
        require(sharesAfter == sharesBefore - amount, "sharesAfter != sharesBefore - amount");
        require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
        return (queuedWithdrawal, withdrawalRoot);
    }

    function testQueueWithdrawalBeaconChainETHToDifferentAddress(address withdrawer) external filterFuzzedAddressInputs(withdrawer) {
        // filtering for test flakiness
        cheats.assume(withdrawer != address(this));
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        bool undelegateIfPossible = false;
        {
            strategyArray[0] = eigenPodManagerMock.beaconChainETHStrategy();
            shareAmounts[0] = REQUIRED_BALANCE_WEI;
            strategyIndexes[0] = 0;
        }
        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH to a different address"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, withdrawer, undelegateIfPossible);
    }

    function testQueueWithdrawalMultipleStrategiesWithBeaconChain() external {
        testDepositIntoStrategySuccessfully(address(this), REQUIRED_BALANCE_WEI);
        IStrategy[] memory strategyArray = new IStrategy[](2);
        uint256[] memory shareAmounts = new uint256[](2);
        uint256[] memory strategyIndexes = new uint256[](2);
        bool undelegateIfPossible = false;
        {
            strategyArray[0] = eigenPodManagerMock.beaconChainETHStrategy();
            shareAmounts[0] = REQUIRED_BALANCE_WEI;
            strategyIndexes[0] = 0;
            strategyArray[1] = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            shareAmounts[1] = REQUIRED_BALANCE_WEI;
            strategyIndexes[1] = 1;
        }
        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot queue a withdrawal including Beacon Chain ETH and other tokens"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(this), undelegateIfPossible);
        {
            strategyArray[0] = dummyStrat;
            shareAmounts[0] = 1;
            strategyIndexes[0] = 0;
            strategyArray[1] = eigenPodManagerMock.beaconChainETHStrategy();
            shareAmounts[1] = REQUIRED_BALANCE_WEI;
            strategyIndexes[1] = 1;
        }
        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot queue a withdrawal including Beacon Chain ETH and other tokens"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(this), undelegateIfPossible);
    }

    function testQueueWithdrawalBeaconChainEthNonWholeAmountGwei(uint256 nonWholeAmount) external {
        cheats.assume(nonWholeAmount % GWEI_TO_WEI != 0);
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        bool undelegateIfPossible = false;
        {
            strategyArray[0] = eigenPodManagerMock.beaconChainETHStrategy();
            shareAmounts[0] = REQUIRED_BALANCE_WEI - 1243895959494;
            strategyIndexes[0] = 0;
        }
        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(this), undelegateIfPossible);
    }

    function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_WithdrawingBeaconChainETH() external {
        _tempStakerStorage = address(this);
        uint256 withdrawalAmount = 1e18;
        _tempStrategyStorage = beaconChainETHStrategy;

        // withdrawalAmount is converted to GWEI here
        testQueueWithdrawalBeaconChainETHToSelf(uint128(withdrawalAmount / 1e9));

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        {
            strategyArray[0] = _tempStrategyStorage;
            shareAmounts[0] = withdrawalAmount;
        }

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

        {
            uint256 nonce = strategyManager.numWithdrawalsQueued(_tempStakerStorage);

            IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
                withdrawer: _tempStakerStorage,
                nonce: (uint96(nonce) - 1)
            });
            queuedWithdrawal = 
                IStrategyManager.QueuedWithdrawal({
                    strategies: strategyArray,
                    shares: shareAmounts,
                    depositor: _tempStakerStorage,
                    withdrawerAndNonce: withdrawerAndNonce,
                    withdrawalStartBlock: uint32(block.number),
                    delegatedAddress: strategyManager.delegation().delegatedTo(_tempStakerStorage)
                }
            );
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(_tempStakerStorage, _tempStrategyStorage);
        // uint256 balanceBefore = address(this).balance;

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = true;

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalCompleted(
            queuedWithdrawal.depositor,
            queuedWithdrawal.withdrawerAndNonce.nonce,
            queuedWithdrawal.withdrawerAndNonce.withdrawer,
            strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
        );
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(_tempStakerStorage, _tempStrategyStorage);
        // uint256 balanceAfter = address(this).balance;

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        // require(balanceAfter == balanceBefore + withdrawalAmount, "balanceAfter != balanceBefore + withdrawalAmount");
        // TODO: make EigenPodManagerMock do something so we can verify that it gets called appropriately?
    }

    function testSlashSharesBeaconChainETH() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = beaconChainETHStrategy;
        IERC20 token;

        testDepositBeaconChainETHSuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = amount;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }
}