// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";

import "forge-std/Test.sol";

import "../../contracts/core/StrategyManager.sol";
import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/Reenterer.sol";
import "../mocks/Reverter.sol";

import "../mocks/ERC20Mock.sol";

import "./Utils.sol";

contract StrategyManagerUnitTests is Test, Utils {

    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 public REQUIRED_BALANCE_WEI = 31 ether;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    StrategyManager public strategyManagerImplementation;
    StrategyManager public strategyManager;
    DelegationMock public delegationMock;
    SlasherMock public slasherMock;
    EigenPodManagerMock public eigenPodManagerMock;

    StrategyBase public dummyStrat;
    StrategyBase public dummyStrat2;

    IStrategy public beaconChainETHStrategy;

    IERC20 public dummyToken;

    Reenterer public reenterer;

    uint256 GWEI_TO_WEI = 1e9;

    address public pauser = address(555);
    address public unpauser = address(999);

    address initialOwner = address(this);

    uint256[] public emptyUintArray;

    // used as transient storage to fix stack-too-deep errors
    IStrategy public _tempStrategyStorage;
    address public _tempStakerStorage;
    uint256 public privateKey = 111111;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    /**
     * @notice Emitted when a new deposit occurs on behalf of `depositor`.
     * @param depositor Is the staker who is depositing funds into EigenLayer.
     * @param strategy Is the strategy that `depositor` has deposited into.
     * @param token Is the token that `depositor` deposited.
     * @param shares Is the number of new shares `depositor` has been granted in `strategy`.
     */
    event Deposit(
        address depositor, IERC20 token, IStrategy strategy, uint256 shares
    );

    /**
     * @notice Emitted when a new withdrawal occurs on behalf of `depositor`.
     * @param depositor Is the staker who is queuing a withdrawal from EigenLayer.
     * @param nonce Is the withdrawal's unique identifier (to the depositor).
     * @param strategy Is the strategy that `depositor` has queued to withdraw from.
     * @param shares Is the number of shares `depositor` has queued to withdraw.
     */
    event ShareWithdrawalQueued(
        address depositor, uint96 nonce, IStrategy strategy, uint256 shares
    );

    /**
     * @notice Emitted when a new withdrawal is queued by `depositor`.
     * @param depositor Is the staker who is withdrawing funds from EigenLayer.
     * @param nonce Is the withdrawal's unique identifier (to the depositor).
     * @param withdrawer Is the party specified by `staker` who will be able to complete the queued withdrawal and receive the withdrawn funds.
     * @param delegatedAddress Is the party who the `staker` was delegated to at the time of creating the queued withdrawal
     * @param withdrawalRoot Is a hash of the input data for the withdrawal.
     */
    event WithdrawalQueued(
        address depositor, uint96 nonce, address withdrawer, address delegatedAddress, bytes32 withdrawalRoot
    );

    /// @notice Emitted when a queued withdrawal is completed
    event WithdrawalCompleted(address indexed depositor, uint96 nonce, address indexed withdrawer, bytes32 withdrawalRoot);

    /// @notice Emitted when the `strategyWhitelister` is changed
    event StrategyWhitelisterChanged(address previousAddress, address newAddress);

    /// @notice Emitted when a strategy is added to the approved list of strategies for deposit
    event StrategyAddedToDepositWhitelist(IStrategy strategy);

    /// @notice Emitted when a strategy is removed from the approved list of strategies for deposit
    event StrategyRemovedFromDepositWhitelist(IStrategy strategy);

    /// @notice Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.
    event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        slasherMock = new SlasherMock();
        delegationMock = new DelegationMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        strategyManagerImplementation = new StrategyManager(delegationMock, eigenPodManagerMock, slasherMock);
        strategyManager = StrategyManager(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyManagerImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        StrategyManager.initialize.selector,
                        initialOwner,
                        initialOwner,
                        pauserRegistry,
                        0/*initialPausedStatus*/,
                        0/*withdrawalDelayBlocks*/
                    )
                )
            )
        );
        dummyToken = new ERC20Mock();
        dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyStrat2 = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](2);
        _strategy[0] = dummyStrat;
        _strategy[1] = dummyStrat2;
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        beaconChainETHStrategy = strategyManager.beaconChainETHStrategy();

        // excude the zero address, the proxyAdmin and the eigenPodManagerMock from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManagerMock)] = true;
    }

    function testCannotReinitialize() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategyManager.initialize(initialOwner, initialOwner, pauserRegistry, 0, 0);
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

    function testDepositIntoStrategySuccessfully(address staker, uint256 amount) public filterFuzzedAddressInputs(staker) {
        IERC20 token = dummyToken;    
        IStrategy strategy = dummyStrat;

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // filter out zero address because the mock ERC20 we are using will revert on using it
        cheats.assume(staker != address(0));
        // filter out the strategy itself from fuzzed inputs
        cheats.assume(staker != address(strategy));
        // sanity check / filter
        cheats.assume(amount <= token.balanceOf(address(this)));
        cheats.assume(amount >= 1);

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        // needed for expecting an event with the right parameters
        uint256 expectedShares = strategy.underlyingToShares(amount);

        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, token, strategy, expectedShares);
        uint256 shares = strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

        require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
        if (sharesBefore == 0) {
            require(stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1, "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1");
            require(strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == strategy,
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy");
        }
    }

    function testDepositIntoStrategySuccessfullyTwice() public {
        address staker = address(this);
        uint256 amount = 1e18;
        testDepositIntoStrategySuccessfully(staker, amount);
        testDepositIntoStrategySuccessfully(staker, amount);
    }

    function testDepositIntoStrategyFailsWhenDepositsPaused() public {
        uint256 amount = 1e18;

        // pause deposits
        cheats.startPrank(pauser);
        strategyManager.pause(1);
        cheats.stopPrank();

        cheats.expectRevert(bytes("Pausable: index is paused"));
        strategyManager.depositIntoStrategy(dummyStrat, dummyToken, amount);
    }

    function testDepositIntoStrategyFailsWhenStakerFrozen() public {
        uint256 amount = 1e18;
        address staker = address(this);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        strategyManager.depositIntoStrategy(dummyStrat, dummyToken, amount);
    }

    function testDepositIntoStrategyFailsWhenReentering() public {
        uint256 amount = 1e18;

        reenterer = new Reenterer();

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = IStrategy(address(reenterer));
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;
        bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.depositIntoStrategy.selector, address(reenterer), dummyToken, amount);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        strategyManager.depositIntoStrategy(IStrategy(address(reenterer)), dummyToken, amount);
    }

    function testDepositIntoStrategyWithSignatureSuccessfully(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage;
        _depositIntoStrategyWithSignature(staker, amount, expiry, expectedRevertMessage);
    }

    function testDepositIntoStrategyWithSignatureReplay(uint256 amount, uint256 expiry) public {
         // min shares must be minted on strategy
        cheats.assume(amount >= 1);
        cheats.assume(expiry > block.timestamp);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        bytes memory signature = _depositIntoStrategyWithSignature(staker, amount, expiry, "");

        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer"));
        strategyManager.depositIntoStrategyWithSignature(dummyStrat, dummyToken, amount, staker, expiry, signature);

    }

    // tries depositing using a signature and an EIP 1271 compliant wallet
    function testDepositIntoStrategyWithSignature_WithContractWallet_Successfully(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage;
        _depositIntoStrategyWithSignature(staker, amount, expiry, expectedRevertMessage);
    }

    // tries depositing using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature
    function testDepositIntoStrategyWithSignature_WithContractWallet_BadSignature(uint256 amount) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= token.balanceOf(address(this)));

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = type(uint256).max;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed"));
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    // tries depositing using a wallet that does not comply with EIP 1271
    function testDepositIntoStrategyWithSignature_WithContractWallet_NonconformingWallet(uint256 amount, uint8 v, bytes32 r, bytes32 s) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
        cheats.stopPrank();
        staker = address(wallet);

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= token.balanceOf(address(this)));

        uint256 expiry = type(uint256).max;
        bytes memory signature = abi.encodePacked(r, s, v);

        cheats.expectRevert();
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    function testDepositIntoStrategyWithSignatureFailsWhenDepositsPaused() public {
        address staker = cheats.addr(privateKey);

        // pause deposits
        cheats.startPrank(pauser);
        strategyManager.pause(1);
        cheats.stopPrank();

        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage = "Pausable: index is paused";
        _depositIntoStrategyWithSignature(staker, 1e18, type(uint256).max, expectedRevertMessage);
    }

    function testDepositIntoStrategyWithSignatureFailsWhenStakerFrozen() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = type(uint256).max;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
    }

    function testDepositIntoStrategyWithSignatureFailsWhenReentering() public {
        reenterer = new Reenterer();

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = IStrategy(address(reenterer));
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = cheats.addr(privateKey);
        IStrategy strategy = IStrategy(address(reenterer));
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = type(uint256).max;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        uint256 shareAmountToReturn = amount;
        reenterer.prepareReturnData(abi.encode(shareAmountToReturn));

        {
            address targetToUse = address(strategyManager);
            uint256 msgValueToUse = 0;
            bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.depositIntoStrategy.selector, address(reenterer), dummyToken, amount);
            reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));
        }

        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore + shareAmountToReturn, "sharesAfter != sharesBefore + shareAmountToReturn");
        require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
    }

    function testDepositIntoStrategyWithSignatureFailsWhenSignatureExpired() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = 5555;
        // warp to 1 second after expiry
        cheats.warp(expiry + 1);
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        cheats.expectRevert(bytes("StrategyManager.depositIntoStrategyWithSignature: signature expired"));
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
    }

    function testDepositIntoStrategyWithSignatureFailsWhenSignatureInvalid() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = block.timestamp;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer"));
        // call with `notStaker` as input instead of `staker` address
        address notStaker = address(3333);
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, notStaker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
    }

    function testUndelegate() public {
        strategyManager.undelegate();
    }

    function testUndelegateRevertsWithActiveDeposits() public {
        address staker = address(this);
        uint256 amount = 1e18;

        testDepositIntoStrategySuccessfully(staker, amount);
        require(strategyManager.stakerStrategyListLength(staker) != 0, "test broken in some way, length shouldn't be 0");

        cheats.expectRevert(bytes("StrategyManager._undelegate: depositor has active deposits"));
        strategyManager.undelegate();
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
            strategyArray[0] = strategyManager.beaconChainETHStrategy();
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
            strategyArray[0] = strategyManager.beaconChainETHStrategy();
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
            strategyArray[1] = strategyManager.beaconChainETHStrategy();
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
            strategyArray[0] = strategyManager.beaconChainETHStrategy();
            shareAmounts[0] = REQUIRED_BALANCE_WEI - 1243895959494;
            strategyIndexes[0] = 0;
        }

        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot queue a withdrawal of Beacon Chain ETH for an non-whole amount of gwei"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(this), undelegateIfPossible);
    }

    function testQueueWithdrawalMismatchedIndexAndStrategyArrayLength() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](2);
        uint256[] memory strategyIndexes = new uint256[](1);
        bool undelegateIfPossible = false;

        {
            strategyArray[0] = strategyManager.beaconChainETHStrategy();
            shareAmounts[0] = 1;    
            shareAmounts[1] = 1;    
            strategyIndexes[0] = 0;
        }

        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: input length mismatch"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(this), undelegateIfPossible);
    }

    function testQueueWithdrawalWithZeroAddress() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        bool undelegateIfPossible = false;

        cheats.expectRevert(bytes("StrategyManager.queueWithdrawal: cannot withdraw to zero address"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(0), undelegateIfPossible);
    }

    function testQueueWithdrawalWithFrozenAddress(address frozenAddress) external filterFuzzedAddressInputs(frozenAddress) {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        uint256[] memory strategyIndexes = new uint256[](1);
        bool undelegateIfPossible = false;

        slasherMock.freezeOperator(frozenAddress);

        cheats.startPrank(frozenAddress);
        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, address(0), undelegateIfPossible);
        cheats.stopPrank();

    }

    function testQueueWithdrawal_ToSelf_NotBeaconChainETH(uint256 depositAmount, uint256 withdrawalAmount, bool undelegateIfPossible) public
        returns (IStrategyManager.QueuedWithdrawal memory /* queuedWithdrawal */, IERC20[] memory /* tokensArray */, bytes32 /* withdrawalRoot */)
    {
        // filtering of fuzzed inputs
        cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);

        // address staker = address(this);
        _tempStrategyStorage = dummyStrat;
        // IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(/*staker*/ address(this), depositAmount);

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, bytes32 withdrawalRoot) =
            _setUpQueuedWithdrawalStructSingleStrat(/*staker*/ address(this), /*withdrawer*/ address(this), dummyToken, _tempStrategyStorage, withdrawalAmount);

        uint256 sharesBefore = strategyManager.stakerStrategyShares(/*staker*/ address(this), _tempStrategyStorage);
        uint256 nonceBefore = strategyManager.numWithdrawalsQueued(/*staker*/ address(this));

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

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

            uint256[] memory strategyIndexes = new uint256[](1);
            strategyIndexes[0] = 0;
            strategyManager.queueWithdrawal(
                strategyIndexes,
                queuedWithdrawal.strategies,
                queuedWithdrawal.shares,
                /*withdrawer*/ address(this),
                undelegateIfPossible
            );
        }

        uint256 sharesAfter = strategyManager.stakerStrategyShares(/*staker*/ address(this), _tempStrategyStorage);
        uint256 nonceAfter = strategyManager.numWithdrawalsQueued(/*staker*/ address(this));

        require(strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is false!");
        require(sharesAfter == sharesBefore - withdrawalAmount, "sharesAfter != sharesBefore - withdrawalAmount");
        require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");

        return (queuedWithdrawal, tokensArray, withdrawalRoot);
    }

    function testQueueWithdrawal_ToSelf_NotBeaconChainETHTwoStrategies(uint256 depositAmount, uint256 withdrawalAmount, bool undelegateIfPossible) public
        returns (IStrategyManager.QueuedWithdrawal memory /* queuedWithdrawal */, IERC20[] memory /* tokensArray */, bytes32 /* withdrawalRoot */)
    {
        // filtering of fuzzed inputs
        cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);


        IStrategy[] memory strategies = new IStrategy[](2);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;

        IERC20[] memory tokens = new IERC20[](2);
        tokens[0] = dummyToken;
        tokens[1] = dummyToken;

        uint256[] memory amounts = new uint256[](2);
        amounts[0] = withdrawalAmount;
        amounts[1] = withdrawalAmount;

        _depositIntoStrategySuccessfully(dummyStrat, /*staker*/ address(this), depositAmount);
        _depositIntoStrategySuccessfully(dummyStrat2, /*staker*/ address(this), depositAmount);

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) =
            _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(/*staker*/ address(this), /*withdrawer*/ address(this), strategies, amounts);

        // uint256 sharesBefore = strategyManager.stakerStrategyShares(/*staker*/ address(this), strategies[0]) + strategyManager.stakerStrategyShares(/*staker*/ address(this), strategies[1]);
        // uint256 nonceBefore = strategyManager.numWithdrawalsQueued(/*staker*/ address(this));

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

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

            uint256[] memory strategyIndexes = new uint256[](2);
            strategyIndexes[0] = 0;
            strategyIndexes[1] = 0;
            strategyManager.queueWithdrawal(
                strategyIndexes,
                queuedWithdrawal.strategies,
                queuedWithdrawal.shares,
                /*withdrawer*/ address(this),
                undelegateIfPossible
            );
        }

        return (queuedWithdrawal, tokens, withdrawalRoot);
    }

    function testQueueWithdrawal_ToDifferentAddress_NotBeaconChainETH(address withdrawer, uint256 amount)
        external filterFuzzedAddressInputs(withdrawer)
    {
        address staker = address(this);
        _tempStrategyStorage = dummyStrat;

        testDepositIntoStrategySuccessfully(staker, amount);

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
            _setUpQueuedWithdrawalStructSingleStrat(staker, withdrawer, /*token*/ dummyToken, _tempStrategyStorage, amount);

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
        uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

        bool undelegateIfPossible = false;
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

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
        uint256 nonceAfter = strategyManager.numWithdrawalsQueued(staker);

        require(strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is false!");
        require(sharesAfter == sharesBefore - amount, "sharesAfter != sharesBefore - amount");
        require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
    }


    // TODO: set up delegation for the following three tests and check afterwords
    function testQueueWithdrawal_WithdrawEverything_DontUndelegate(uint256 amount) external {
        // delegate to self
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationMock.delegateTo(address(this), signatureWithExpiry, bytes32(0));
        require(delegationMock.isDelegated(address(this)), "delegation mock setup failed");
        bool undelegateIfPossible = false;
        // deposit and withdraw the same amount, don't undelegate
        testQueueWithdrawal_ToSelf_NotBeaconChainETH(amount, amount, undelegateIfPossible);
        require(delegationMock.isDelegated(address(this)) == !undelegateIfPossible, "undelegation mock failed");
    }

    function testQueueWithdrawal_WithdrawEverything_DoUndelegate(uint256 amount) external {
        bool undelegateIfPossible = true;
        // deposit and withdraw the same amount, do undelegate if possible
        testQueueWithdrawal_ToSelf_NotBeaconChainETH(amount, amount, undelegateIfPossible);
        require(delegationMock.isDelegated(address(this)) == !undelegateIfPossible, "undelegation mock failed");
    }

    function testQueueWithdrawal_DontWithdrawEverything_MarkUndelegateIfPossibleAsTrue(uint128 amount) external {
        bool undelegateIfPossible = true;
        // deposit and withdraw only half, do undelegate if possible
        testQueueWithdrawal_ToSelf_NotBeaconChainETH(uint256(amount) * 2, amount, undelegateIfPossible);
        require(!delegationMock.isDelegated(address(this)), "undelegation mock failed");
    }

    function testQueueWithdrawalFailsWhenStakerFrozen() public {
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;

        testDepositIntoStrategySuccessfully(staker, depositAmount);

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
            _setUpQueuedWithdrawalStructSingleStrat(staker, /*withdrawer*/ staker, token, strategy, withdrawalAmount);

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

        // freeze the staker
        slasherMock.freezeOperator(staker);

        // bool undelegateIfPossible = false;
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        strategyManager.queueWithdrawal(strategyIndexes, queuedWithdrawal.strategies, queuedWithdrawal.shares, /*withdrawer*/ staker, /*undelegateIfPossible*/ false);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 nonceAfter = strategyManager.numWithdrawalsQueued(address(this));

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is true!");
        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
    }

    function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedFalse() external {
        address staker = address(this);
        uint256 withdrawalAmount = 1e18;
        IStrategy strategy = dummyStrat;

        {
            uint256 depositAmount = 1e18;
            bool undelegateIfPossible = false;
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);
        }

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        {
            strategyArray[0] = strategy;
            shareAmounts[0] = withdrawalAmount;
            tokensArray[0] = dummyToken;
        }

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

        {
            uint256 nonce = strategyManager.numWithdrawalsQueued(staker);

            IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
                withdrawer: staker,
                nonce: (uint96(nonce) - 1)
            });
            queuedWithdrawal = 
                IStrategyManager.QueuedWithdrawal({
                    strategies: strategyArray,
                    shares: shareAmounts,
                    depositor: staker,
                    withdrawerAndNonce: withdrawerAndNonce,
                    withdrawalStartBlock: uint32(block.number),
                    delegatedAddress: strategyManager.delegation().delegatedTo(staker)
                }
            );
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(staker));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalCompleted(
            queuedWithdrawal.depositor,
            queuedWithdrawal.withdrawerAndNonce.nonce,
            queuedWithdrawal.withdrawerAndNonce.withdrawer,
            strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
        );
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(staker));

        require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue_NotWithdrawingBeaconChainETH() external {
        address staker = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;
        _tempStrategyStorage = dummyStrat;

        testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        {
            strategyArray[0] = _tempStrategyStorage;
            shareAmounts[0] = withdrawalAmount;
            tokensArray[0] = dummyToken;
        }

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

        {
            uint256 nonce = strategyManager.numWithdrawalsQueued(staker);

            IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
                withdrawer: staker,
                nonce: (uint96(nonce) - 1)
            });
            queuedWithdrawal = 
                IStrategyManager.QueuedWithdrawal({
                    strategies: strategyArray,
                    shares: shareAmounts,
                    depositor: staker,
                    withdrawerAndNonce: withdrawerAndNonce,
                    withdrawalStartBlock: uint32(block.number),
                    delegatedAddress: strategyManager.delegation().delegatedTo(staker)
                }
            );
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
        uint256 balanceBefore = dummyToken.balanceOf(address(staker));

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalCompleted(
            queuedWithdrawal.depositor,
            queuedWithdrawal.withdrawerAndNonce.nonce,
            queuedWithdrawal.withdrawerAndNonce.withdrawer,
            strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
        );
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, /*middlewareTimesIndex*/ 0, /*receiveAsTokens*/ true);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
        uint256 balanceAfter = dummyToken.balanceOf(address(staker));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore + withdrawalAmount, "balanceAfter != balanceBefore + withdrawalAmount");
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

    function testCompleteQueuedWithdrawalFailsWhenWithdrawalsPaused() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy strategy = queuedWithdrawal.strategies[0];

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        // pause withdrawals
        cheats.startPrank(pauser);
        strategyManager.pause(2);
        cheats.stopPrank();

        cheats.expectRevert(bytes("Pausable: index is paused"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testCompleteQueuedWithdrawalFailsWhenDelegatedAddressFrozen() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy strategy = queuedWithdrawal.strategies[0];

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        // freeze the delegatedAddress
        slasherMock.freezeOperator(strategyManager.delegation().delegatedTo(_tempStakerStorage));

        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testCompleteQueuedWithdrawalFailsWhenAttemptingReentrancy() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;
        IStrategy strategy = dummyStrat;

        reenterer.prepareReturnData(abi.encode(depositAmount));

        testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        {
            strategyArray[0] = strategy;
            shareAmounts[0] = withdrawalAmount;
            tokensArray[0] = dummyToken;
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

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;
        bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.completeQueuedWithdrawal.selector, queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    }

    function testCompleteQueuedWithdrawalFailsWhenWithdrawalDoesNotExist() external {
        _tempStakerStorage = address(this);
        uint256 withdrawalAmount = 1e18;
        IStrategy strategy = dummyStrat;

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        {
            strategyArray[0] = strategy;
            shareAmounts[0] = withdrawalAmount;
            tokensArray[0] = dummyToken;
        }

        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

        {
            IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
                withdrawer: _tempStakerStorage,
                nonce: 0
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

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawal is not pending"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testCompleteQueuedWithdrawalFailsWhenCanWithdrawReturnsFalse() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy strategy = queuedWithdrawal.strategies[0];

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        // prepare mock
        slasherMock.setCanWithdrawResponse(false);

        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: shares pending withdrawal are still slashable"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testUndelegateWithFrozenStaker() public {
        slasherMock.setOperatorFrozenStatus(address(this), true);
        cheats.expectRevert(bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing"));
        cheats.startPrank(address(this));
        strategyManager.undelegate();
        cheats.stopPrank();
    }

    function testCompleteQueuedWithdrawalFailsWhenNotCallingFromWithdrawerAddress() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy strategy = queuedWithdrawal.strategies[0];

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        cheats.startPrank(address(123456));
        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: only specified withdrawer can complete a queued withdrawal"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    }

    function testCompleteQueuedWithdrawalFailsWhenTryingToCompleteSameWithdrawal2X() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        IStrategy strategy = queuedWithdrawal.strategies[0];

        uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalCompleted(
            queuedWithdrawal.depositor,
            queuedWithdrawal.withdrawerAndNonce.nonce,
            queuedWithdrawal.withdrawerAndNonce.withdrawer,
            strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
        );
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
        uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

        require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");

        // try to complete same withdrawal again
        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawal is not pending"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    }

    function testCompleteQueuedWithdrawalFailsWhenWithdrawalDelayBlocksHasNotPassed() external {
        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        uint256 valueToSet = 1;
        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(strategyManager.owner());
        uint256 previousValue = strategyManager.withdrawalDelayBlocks();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
        strategyManager.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
        require(strategyManager.withdrawalDelayBlocks() == valueToSet, "strategyManager.withdrawalDelayBlocks() != valueToSet");

        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    }

    function testCompleteQueuedWithdrawalWithNonzeroWithdrawalDelayBlocks(uint16 valueToSet) external {
        // filter fuzzed inputs to allowed *and nonzero* amounts
        cheats.assume(valueToSet <= strategyManager.MAX_WITHDRAWAL_DELAY_BLOCKS() && valueToSet != 0);

        _tempStakerStorage = address(this);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = 1e18;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 middlewareTimesIndex = 0;
        bool receiveAsTokens = false;

        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(strategyManager.owner());
        uint256 previousValue = strategyManager.withdrawalDelayBlocks();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
        strategyManager.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
        require(strategyManager.withdrawalDelayBlocks() == valueToSet, "strategyManager.withdrawalDelayBlocks() != valueToSet");

        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);


        // roll block number forward to one block before the withdrawal should be completeable and attempt again
        uint256 originalBlockNumber = block.number;
        cheats.roll(originalBlockNumber + valueToSet - 1);
        cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed"));
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

        // roll block number forward to the block at which the withdrawal should be completeable, and complete it
        cheats.roll(originalBlockNumber + valueToSet);
    }

    function testSlashSharesNotBeaconChainETHFuzzed(uint64 withdrawalAmount) external {
        // cannot cause share value to increase too drastically
        cheats.assume(withdrawalAmount <= 1e9 || withdrawalAmount == 1e18);
        _tempStakerStorage = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        {
            uint256 depositAmount = 1e18;
            // filter fuzzed input
            cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
            testDepositIntoStrategySuccessfully(_tempStakerStorage, depositAmount);
        }

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = uint256(withdrawalAmount);

        // freeze the staker
        slasherMock.freezeOperator(_tempStakerStorage);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        uint256 sharesBefore = strategyManager.stakerStrategyShares(_tempStakerStorage, strategy);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(_tempStakerStorage);
        uint256 balanceBefore = dummyToken.balanceOf(recipient);

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(_tempStakerStorage, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(_tempStakerStorage);
        uint256 balanceAfter = dummyToken.balanceOf(recipient);

        require(sharesAfter == sharesBefore - uint256(withdrawalAmount), "sharesAfter != sharesBefore - uint256(withdrawalAmount)");
        require(balanceAfter == balanceBefore + uint256(withdrawalAmount), "balanceAfter != balanceBefore + uint256(withdrawalAmount)");
        if (sharesAfter == 0) {
            require(stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1, "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1");
        }
    }

    function testSlashSharesNotBeaconChainETH_AllShares() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully( staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        // slash the same amount as deposited
        shareAmounts[0] = amount;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 balanceBefore = dummyToken.balanceOf(recipient);

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 balanceAfter = dummyToken.balanceOf(recipient);

        require(sharesAfter == sharesBefore - amount, "sharesAfter != sharesBefore - amount");
        require(balanceAfter == balanceBefore + amount, "balanceAfter != balanceBefore + amount");
        require(sharesAfter == 0, "sharesAfter != 0");
        require(stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1, "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1");
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

    function testSlashSharesMixIncludingBeaconChainETH() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);
        testDepositBeaconChainETHSuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](2);
        IERC20[] memory tokensArray = new IERC20[](2);
        uint256[] memory shareAmounts = new uint256[](2);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = amount;
        strategyArray[1] = beaconChainETHStrategy;
        tokensArray[1] = token;
        shareAmounts[1] = amount;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](2);
        strategyIndexes[0] = 0;
        // this index is also zero, since the other strategy will be removed!
        strategyIndexes[1] = 0;

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 balanceBefore = dummyToken.balanceOf(recipient);

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 balanceAfter = dummyToken.balanceOf(recipient);

        require(sharesAfter == sharesBefore - amount, "sharesAfter != sharesBefore - amount");
        require(balanceAfter == balanceBefore + amount, "balanceAfter != balanceBefore + amount");
    }
    

    function testSlashSharesRevertsWhenCalledByNotOwner() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);

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

        // recipient is not the owner
        cheats.startPrank(recipient);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }

    function testSlashSharesRevertsWhenStakerNotFrozen() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = amount;

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager.onlyFrozen: staker has not been frozen"));
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }

    function testSlashSharesRevertsWhenAttemptingReentrancy() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        reenterer.prepareReturnData(abi.encode(amount));

        testDepositIntoStrategySuccessfully(staker, amount);
        testDepositBeaconChainETHSuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](2);
        IERC20[] memory tokensArray = new IERC20[](2);
        uint256[] memory shareAmounts = new uint256[](2);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = amount;
        strategyArray[1] = beaconChainETHStrategy;
        tokensArray[1] = token;
        shareAmounts[1] = amount;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](2);
        strategyIndexes[0] = 0;
        // this index is also zero, since the other strategy will be removed!
        strategyIndexes[1] = 0;

        // transfer strategyManager's ownership to the reenterer
        cheats.startPrank(strategyManager.owner());
        strategyManager.transferOwnership(address(reenterer));
        cheats.stopPrank();

        // prepare for reentrant call, expecting revert for reentrancy
        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;
        bytes memory calldataToUse =
            abi.encodeWithSelector(StrategyManager.slashShares.selector, slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }

    function testSlashQueuedWithdrawalNotBeaconChainETH() external {
        address recipient = address(333);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) = testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 balanceBefore = dummyToken.balanceOf(address(recipient));

        // slash the delegatedOperator
        slasherMock.freezeOperator(queuedWithdrawal.delegatedAddress);

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        cheats.stopPrank();

        uint256 balanceAfter = dummyToken.balanceOf(address(recipient));
        require(balanceAfter == balanceBefore + withdrawalAmount, "balanceAfter != balanceBefore + withdrawalAmount");
        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is true!");
    }

    /// @notice this function is to test for a bug identified in the Code4Rena audit (H-205).  This bug essentially
    ///         allowed a strategy that is meant to be skipped, to actually be withdrawn from.  This is a regression test
    ///         to ensure that this bug does not reappear.
    function testSlashQueuedWithdrawalIncrementor() external {
        address recipient = address(333);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
testQueueWithdrawal_ToSelf_NotBeaconChainETHTwoStrategies(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 balanceBefore = dummyToken.balanceOf(address(recipient));

        // slash the delegatedOperator
        slasherMock.freezeOperator(queuedWithdrawal.delegatedAddress);

        uint256[] memory indicesToSkip = new uint256[](2);

        indicesToSkip[0] = 0;
        indicesToSkip[1] = 1;

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustTwoDummyTokens(), indicesToSkip);
        cheats.stopPrank();

        uint256 balanceAfter = dummyToken.balanceOf(address(recipient));

        /**
         * This check ensures that the strategy has not been withdrawn from.  If the incrementor is misplaced inside
         * the else statement (as it was before the fix was made), the withdrawal would have been triggered for the 
         * the strategy that we intended to skip, i.e., the check indicesToSkip[indicesToSkipIndex] == i would have 
         * failed, triggering the else logic to withdraw from the strategy that was at index 0.
         */
        require(balanceAfter == balanceBefore, "withdrawal should not have been processed");

        require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is true!");
    }

    function testSlashQueuedWithdrawalFailsWhenNotCallingFromOwnerAddress() external {
        address recipient = address(333);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 balanceBefore = dummyToken.balanceOf(address(recipient));

        // slash the delegatedOperator
        slasherMock.freezeOperator(queuedWithdrawal.delegatedAddress);

        // recipient is not strategyManager.owner()
        cheats.startPrank(recipient);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        cheats.stopPrank();

        uint256 balanceAfter = dummyToken.balanceOf(address(recipient));

        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
        require(strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is false");
    }

    function testSlashQueuedWithdrawalFailsWhenDelegatedAddressNotFrozen() external {
        address recipient = address(333);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;
        bool undelegateIfPossible = false;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, bytes32 withdrawalRoot) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        uint256 balanceBefore = dummyToken.balanceOf(address(recipient));

        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager.onlyFrozen: staker has not been frozen"));
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        cheats.stopPrank();

        uint256 balanceAfter = dummyToken.balanceOf(address(recipient));

        require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
        require(strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is false");
    }

    function testSlashQueuedWithdrawalFailsWhenAttemptingReentrancy() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        address recipient = address(333);
        uint256 depositAmount = 1e18;
        uint256 withdrawalAmount = depositAmount;
        bool undelegateIfPossible = false;

        reenterer.prepareReturnData(abi.encode(depositAmount));

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*IERC20[] memory tokensArray*/, /*bytes32 withdrawalRoot*/) =
            testQueueWithdrawal_ToSelf_NotBeaconChainETH(depositAmount, withdrawalAmount, undelegateIfPossible);

        // freeze the delegatedAddress
        slasherMock.freezeOperator(strategyManager.delegation().delegatedTo(staker));

        // transfer strategyManager's ownership to the reenterer
        cheats.startPrank(strategyManager.owner());
        strategyManager.transferOwnership(address(reenterer));
        cheats.stopPrank();

        // prepare for reentrant call, expecting revert for reentrancy
        address targetToUse = address(strategyManager);
        uint256 msgValueToUse = 0;
        bytes memory calldataToUse =
            abi.encodeWithSelector(StrategyManager.slashQueuedWithdrawal.selector, recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        cheats.stopPrank();
    }

    function testSlashQueuedWithdrawalFailsWhenWithdrawalDoesNotExist() external {
        address recipient = address(333);
        uint256 amount = 1e18;

        (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, /*bytes32 withdrawalRoot*/) =
            // convert wei to gwei for test input
            testQueueWithdrawalBeaconChainETHToSelf(uint128(amount / 1e9));

        // slash the delegatedOperator
        slasherMock.freezeOperator(queuedWithdrawal.delegatedAddress);

        // modify the queuedWithdrawal data so the root won't exist
        queuedWithdrawal.shares[0] = (amount * 2);

        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager.slashQueuedWithdrawal: withdrawal is not pending"));
        strategyManager.slashQueuedWithdrawal(recipient, queuedWithdrawal, _arrayWithJustDummyToken(), emptyUintArray);
        cheats.stopPrank();
    }

    function test_addSharesRevertsWhenSharesIsZero() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        reenterer.prepareReturnData(abi.encode(uint256(0)));

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("StrategyManager._addShares: shares should not be zero!"));
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_addSharesRevertsWhenDepositWouldExeedMaxArrayLength() external {
        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        // uint256 MAX_STAKER_STRATEGY_LIST_LENGTH = strategyManager.MAX_STAKER_STRATEGY_LIST_LENGTH();
        uint256 MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

        // loop that deploys a new strategy and deposits into it
        for (uint256 i = 0; i < MAX_STAKER_STRATEGY_LIST_LENGTH; ++i) {
            cheats.startPrank(staker);
            strategyManager.depositIntoStrategy(strategy, token, amount);
            cheats.stopPrank();

            dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategy = dummyStrat;

            // whitelist the strategy for deposit
            cheats.startPrank(strategyManager.owner());
            IStrategy[] memory _strategy = new IStrategy[](1);
            _strategy[0] = dummyStrat;
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(dummyStrat);
            strategyManager.addStrategiesToDepositWhitelist(_strategy);
            cheats.stopPrank();
        }

        require(strategyManager.stakerStrategyListLength(staker) == MAX_STAKER_STRATEGY_LIST_LENGTH, 
            "strategyManager.stakerStrategyListLength(staker) != MAX_STAKER_STRATEGY_LIST_LENGTH");

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("StrategyManager._addShares: deposit would exceed MAX_STAKER_STRATEGY_LIST_LENGTH"));
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_depositIntoStrategyRevertsWhenTokenSafeTransferFromReverts() external {
        // replace 'dummyStrat' with one that uses a reverting token
        dummyToken = IERC20(address(new Reverter()));
        dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.startPrank(staker);
        cheats.expectRevert();
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_depositIntoStrategyRevertsWhenTokenDoesNotExist() external {
        // replace 'dummyStrat' with one that uses a non-existent token
        dummyToken = IERC20(address(5678));
        dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.startPrank(staker);
        cheats.expectRevert();
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_depositIntoStrategyRevertsWhenStrategyDepositFunctionReverts() external {
        // replace 'dummyStrat' with one that always reverts
        dummyStrat = StrategyBase(
            address(
                new Reverter()
            )
        );

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.startPrank(staker);
        cheats.expectRevert();
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_depositIntoStrategyRevertsWhenStrategyDoesNotExist() external {
        // replace 'dummyStrat' with one that does not exist
        dummyStrat = StrategyBase(
            address(5678)
        );

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(dummyStrat);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.startPrank(staker);
        cheats.expectRevert();
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_depositIntoStrategyRevertsWhenStrategyNotWhitelisted() external {
        // replace 'dummyStrat' with one that is not whitelisted
        dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.startPrank(staker);
        cheats.expectRevert("StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted");
        strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();
    }

    function test_removeSharesRevertsWhenShareAmountIsZero() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = 0;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount should not be zero!"));
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }

    function test_removeSharesRevertsWhenShareAmountIsTooLarge() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);

        IStrategy[] memory strategyArray = new IStrategy[](1);
        IERC20[] memory tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = amount + 1;

        // freeze the staker
        slasherMock.freezeOperator(staker);

        address slashedAddress = address(this);
        address recipient = address(333);
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;

        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount too high"));
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();
    }

    function test_removeStrategyFromStakerStrategyListWorksWithIncorrectIndexInput() external {
        uint256 amount = 1e18;
        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        testDepositIntoStrategySuccessfully(staker, amount);
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
        strategyIndexes[0] = 1;

        // check that we are actually supplying an incorrect index!
        require(strategyManager.stakerStrategyList(staker, strategyIndexes[0]) != strategyArray[0],
            "we want to supply an incorrect index but have supplied a correct one");

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 balanceBefore = dummyToken.balanceOf(recipient);

        cheats.startPrank(strategyManager.owner());
        strategyManager.slashShares(slashedAddress, recipient, strategyArray, tokensArray, strategyIndexes, shareAmounts);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 balanceAfter = dummyToken.balanceOf(recipient);

        require(sharesAfter == sharesBefore - amount, "sharesAfter != sharesBefore - amount");
        require(balanceAfter == balanceBefore + amount, "balanceAfter != balanceBefore + amount");
    }

    function testSetWithdrawalDelayBlocks(uint16 valueToSet) external {
        // filter fuzzed inputs to allowed amounts
        cheats.assume(valueToSet <= strategyManager.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(strategyManager.owner());
        uint256 previousValue = strategyManager.withdrawalDelayBlocks();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
        strategyManager.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
        require(strategyManager.withdrawalDelayBlocks() == valueToSet, "strategyManager.withdrawalDelayBlocks() != valueToSet");
    }

    function testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner(address notOwner) filterFuzzedAddressInputs(notOwner) external {
        cheats.assume(notOwner != strategyManager.owner());

        uint256 valueToSet = 1;
        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        strategyManager.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
    }

    function testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh(uint256 valueToSet) external {
        // filter fuzzed inputs to disallowed amounts
        cheats.assume(valueToSet > strategyManager.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // attempt to set the `withdrawalDelayBlocks` variable
        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert(bytes("StrategyManager.setWithdrawalDelay: _withdrawalDelayBlocks too high"));
        strategyManager.setWithdrawalDelayBlocks(valueToSet);
    }

    function testSetStrategyWhitelister(address newWhitelister) external {
        address previousStrategyWhitelister = strategyManager.strategyWhitelister();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyWhitelisterChanged(previousStrategyWhitelister, newWhitelister);
        strategyManager.setStrategyWhitelister(newWhitelister);
        require(strategyManager.strategyWhitelister() == newWhitelister, "strategyManager.strategyWhitelister() != newWhitelister");
    }

    function testSetStrategyWhitelisterRevertsWhenCalledByNotOwner(address notOwner)
        external filterFuzzedAddressInputs(notOwner)
    {
        cheats.assume(notOwner != strategyManager.owner());
        address newWhitelister = address(this);
        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        strategyManager.setStrategyWhitelister(newWhitelister);
        cheats.stopPrank();
    }

    function testAddStrategiesToDepositWhitelist(uint8 numberOfStrategiesToAdd) public returns (IStrategy[] memory) {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);

        IStrategy[] memory strategyArray = new IStrategy[](numberOfStrategiesToAdd);
        // loop that deploys a new strategy and adds it to the array
        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            IStrategy _strategy = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategyArray[i] = _strategy;
            require(!strategyManager.strategyIsWhitelistedForDeposit(_strategy), "strategy improperly whitelisted?");
        }

        cheats.startPrank(strategyManager.strategyWhitelister());
        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(strategyArray[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        cheats.stopPrank();

        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            require(strategyManager.strategyIsWhitelistedForDeposit(strategyArray[i]), "strategy not properly whitelisted");
        }

        return strategyArray;
    }

    function testAddStrategiesToDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(address notStrategyWhitelister)
        external filterFuzzedAddressInputs(notStrategyWhitelister)
    {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IStrategy _strategy = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = _strategy;

        cheats.startPrank(notStrategyWhitelister);
        cheats.expectRevert(bytes("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"));
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        cheats.stopPrank();
    }

    function testRemoveStrategiesFromDepositWhitelist(uint8 numberOfStrategiesToAdd, uint8 numberOfStrategiesToRemove) external {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);
        cheats.assume(numberOfStrategiesToRemove <= 16);
        cheats.assume(numberOfStrategiesToRemove <= numberOfStrategiesToAdd);

        IStrategy[] memory strategiesAdded = testAddStrategiesToDepositWhitelist(numberOfStrategiesToAdd);

        IStrategy[] memory strategiesToRemove = new IStrategy[](numberOfStrategiesToRemove);
        // loop that selectively copies from array to other array
        for (uint256 i = 0; i < numberOfStrategiesToRemove; ++i) {
            strategiesToRemove[i] = strategiesAdded[i];
        }

        cheats.startPrank(strategyManager.strategyWhitelister());
        for (uint256 i = 0; i < strategiesToRemove.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyRemovedFromDepositWhitelist(strategiesToRemove[i]);
        }
        strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemove);
        cheats.stopPrank();

        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            if (i < numberOfStrategiesToRemove) {
                require(!strategyManager.strategyIsWhitelistedForDeposit(strategiesToRemove[i]), "strategy not properly removed from whitelist");
            } else {
                require(strategyManager.strategyIsWhitelistedForDeposit(strategiesAdded[i]), "strategy improperly removed from whitelist?");                
            }
        }
    }

    function testRemoveStrategiesFromDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(address notStrategyWhitelister)
        external filterFuzzedAddressInputs(notStrategyWhitelister)
    {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = testAddStrategiesToDepositWhitelist(1);

        cheats.startPrank(notStrategyWhitelister);
        cheats.expectRevert(bytes("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"));
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        cheats.stopPrank();
    }

    // INTERNAL / HELPER FUNCTIONS
    function _beaconChainReentrancyTestsSetup() internal {
        // prepare StrategyManager with EigenPodManager and Delegation replaced with a Reenterer contract
        reenterer = new Reenterer();
        strategyManagerImplementation = new StrategyManager(IDelegationManager(address(reenterer)), IEigenPodManager(address(reenterer)), slasherMock);
        strategyManager = StrategyManager(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyManagerImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyManager.initialize.selector, initialOwner, initialOwner, pauserRegistry, 0, 0)
                )
            )
        );
    }

    function _setUpQueuedWithdrawalStructSingleStrat(address staker, address withdrawer, IERC20 token, IStrategy strategy, uint256 shareAmount)
        internal view returns (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, bytes32 withdrawalRoot)
    {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = shareAmount;
        IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
            withdrawer: withdrawer,
            nonce: uint96(strategyManager.numWithdrawalsQueued(staker))
        });
        queuedWithdrawal = 
            IStrategyManager.QueuedWithdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                depositor: staker,
                withdrawerAndNonce: withdrawerAndNonce,
                withdrawalStartBlock: uint32(block.number),
                delegatedAddress: strategyManager.delegation().delegatedTo(staker)
            }
        );
        // calculate the withdrawal root
        withdrawalRoot = strategyManager.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, tokensArray, withdrawalRoot);
    }

    function _depositIntoStrategySuccessfully(IStrategy strategy, address staker, uint256 amount) internal filterFuzzedAddressInputs(staker) {
        IERC20 token = dummyToken;    

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // filter out zero address because the mock ERC20 we are using will revert on using it
        cheats.assume(staker != address(0));
        // sanity check / filter
        cheats.assume(amount <= token.balanceOf(address(this)));

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        // needed for expecting an event with the right parameters
        uint256 expectedShares = amount;

        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, token, strategy, expectedShares);
        uint256 shares = strategyManager.depositIntoStrategy(strategy, token, amount);
        cheats.stopPrank();

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

        require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
        if (sharesBefore == 0) {
            require(stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1, "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1");
            require(strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == strategy,
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy");
        }
    }

    function _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(
        address staker, 
        address withdrawer, 
        IStrategy[] memory strategyArray, 
        uint256[] memory shareAmounts
    )
        internal view returns (IStrategyManager.QueuedWithdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    {
        IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
            withdrawer: withdrawer,
            nonce: uint96(strategyManager.numWithdrawalsQueued(staker))
        });
        queuedWithdrawal = 
            IStrategyManager.QueuedWithdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                depositor: staker,
                withdrawerAndNonce: withdrawerAndNonce,
                withdrawalStartBlock: uint32(block.number),
                delegatedAddress: strategyManager.delegation().delegatedTo(staker)
            }
        );
        // calculate the withdrawal root
        withdrawalRoot = strategyManager.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, withdrawalRoot);
    }

    function _arrayWithJustDummyToken() internal view returns (IERC20[] memory) {
        IERC20[] memory array = new IERC20[](1);
        array[0] = dummyToken;
        return array;
    }

    function _arrayWithJustTwoDummyTokens() internal view returns (IERC20[] memory) {
        IERC20[] memory array = new IERC20[](2);
        array[0] = dummyToken;
        array[1] = dummyToken;
        return array;
    }

    function _calculateSharesDelta(uint256 newAmountGwei, uint256 currentAmountGwei) internal view returns(uint256, bool) {
        uint256 sharesDelta;
        bool isNegative;
        if (currentAmountGwei > newAmountGwei){
            sharesDelta = currentAmountGwei - newAmountGwei;
            isNegative = true;
        } else {
            sharesDelta = newAmountGwei - currentAmountGwei;
        }
        return (sharesDelta * GWEI_TO_WEI, isNegative);
    }

    // internal function for de-duping code. expects success if `expectedRevertMessage` is empty and expiry is valid.
    function _depositIntoStrategyWithSignature(address staker, uint256 amount, uint256 expiry, string memory expectedRevertMessage) internal returns (bytes memory) {

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= dummyToken.balanceOf(address(this)));

        uint256 nonceBefore = strategyManager.nonces(staker);
        bytes memory signature;

        {
            bytes32 structHash = keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), dummyStrat, dummyToken, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);

        bool expectedRevertMessageIsempty;
        {
            string memory emptyString;
            expectedRevertMessageIsempty = keccak256(abi.encodePacked(expectedRevertMessage)) == keccak256(abi.encodePacked(emptyString));
        }
        if (!expectedRevertMessageIsempty) {
            cheats.expectRevert(bytes(expectedRevertMessage));
        } else if (expiry < block.timestamp) {
            cheats.expectRevert("StrategyManager.depositIntoStrategyWithSignature: signature expired");
        } else {
            // needed for expecting an event with the right parameters
            uint256 expectedShares = amount;
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit Deposit(staker, dummyToken, dummyStrat, expectedShares);
        }
        uint256 shares = strategyManager.depositIntoStrategyWithSignature(dummyStrat, dummyToken, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        uint256 nonceAfter = strategyManager.nonces(staker);

        if (expiry >= block.timestamp && expectedRevertMessageIsempty) {
            require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
            require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
        }
        return signature;
    }
}