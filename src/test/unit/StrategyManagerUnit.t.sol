// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";

import "forge-std/Test.sol";

import "../../contracts/core/StrategyManager.sol";
import "../../contracts/strategies/StrategyBase.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../mocks/DelegationManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/Reenterer.sol";
import "../mocks/Reverter.sol";

import "../mocks/ERC20Mock.sol";

import "./Utils.sol";

contract StrategyManagerUnitTests is Test, Utils {
    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 public REQUIRED_BALANCE_WEI = 32 ether;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    StrategyManager public strategyManagerImplementation;
    StrategyManager public strategyManager;
    DelegationManagerMock public delegationManagerMock;
    SlasherMock public slasherMock;
    EigenPodManagerMock public eigenPodManagerMock;

    StrategyBase public dummyStrat;
    StrategyBase public dummyStrat2;
    StrategyBase public dummyStrat3;

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
    event Deposit(address depositor, IERC20 token, IStrategy strategy, uint256 shares);

    /**
     * @notice Emitted when a new withdrawal occurs on behalf of `depositor`.
     * @param depositor Is the staker who is queuing a withdrawal from EigenLayer.
     * @param nonce Is the withdrawal's unique identifier (to the depositor).
     * @param strategy Is the strategy that `depositor` has queued to withdraw from.
     * @param shares Is the number of shares `depositor` has queued to withdraw.
     */
    event ShareWithdrawalQueued(address depositor, uint96 nonce, IStrategy strategy, uint256 shares);

    /**
     * @notice Emitted when a new withdrawal is queued by `depositor`.
     * @param depositor Is the staker who is withdrawing funds from EigenLayer.
     * @param nonce Is the withdrawal's unique identifier (to the depositor).
     * @param withdrawer Is the party specified by `staker` who will be able to complete the queued withdrawal and receive the withdrawn funds.
     * @param delegatedAddress Is the party who the `staker` was delegated to at the time of creating the queued withdrawal
     * @param withdrawalRoot Is a hash of the input data for the withdrawal.
     */
    event WithdrawalQueued(
        address depositor,
        uint96 nonce,
        address withdrawer,
        address delegatedAddress,
        bytes32 withdrawalRoot
    );

    /// @notice Emitted when a queued withdrawal is completed
    event WithdrawalCompleted(
        address indexed depositor,
        uint96 nonce,
        address indexed withdrawer,
        bytes32 withdrawalRoot
    );

    /// @notice Emitted when the `strategyWhitelister` is changed
    event StrategyWhitelisterChanged(address previousAddress, address newAddress);

    /// @notice Emitted when a strategy is added to the approved list of strategies for deposit
    event StrategyAddedToDepositWhitelist(IStrategy strategy);

    /// @notice Emitted when a strategy is removed from the approved list of strategies for deposit
    event StrategyRemovedFromDepositWhitelist(IStrategy strategy);

    /// @notice Emitted when the `withdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.
    event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue);

    function setUp() public virtual {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        slasherMock = new SlasherMock();
        delegationManagerMock = new DelegationManagerMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        strategyManagerImplementation = new StrategyManager(delegationManagerMock, eigenPodManagerMock, slasherMock);
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
                        0 /*initialPausedStatus*/
                    )
                )
            )
        );
        dummyToken = new ERC20Mock();
        dummyStrat = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyStrat2 = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyStrat3 = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](3);
        _strategy[0] = dummyStrat;
        _strategy[1] = dummyStrat2;
        _strategy[2] = dummyStrat3;
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        beaconChainETHStrategy = eigenPodManagerMock.beaconChainETHStrategy();

        // excude the zero address, the proxyAdmin and the eigenPodManagerMock from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManagerMock)] = true;
    }

    function testCannotReinitialize() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        strategyManager.initialize(initialOwner, initialOwner, pauserRegistry, 0);
    }

    function testDepositIntoStrategySuccessfully(
        address staker,
        uint256 amount
    ) public filterFuzzedAddressInputs(staker) {
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
            require(
                stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
            );
            require(
                strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == strategy,
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
            );
        }
    }

    function testDepositIntoStrategySuccessfullyTwice() public {
        address staker = address(this);
        uint256 amount = 1e18;
        testDepositIntoStrategySuccessfully(staker, amount);
        testDepositIntoStrategySuccessfully(staker, amount);
    }

    function testDepositIntoStrategyRevertsWhenDepositsPaused() public {
        uint256 amount = 1e18;

        // pause deposits
        cheats.startPrank(pauser);
        strategyManager.pause(1);
        cheats.stopPrank();

        cheats.expectRevert(bytes("Pausable: index is paused"));
        strategyManager.depositIntoStrategy(dummyStrat, dummyToken, amount);
    }

    function testDepositIntoStrategyRevertsWhenReentering() public {
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
        bytes memory calldataToUse = abi.encodeWithSelector(
            StrategyManager.depositIntoStrategy.selector,
            address(reenterer),
            dummyToken,
            amount
        );
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
    function testDepositIntoStrategyWithSignature_WithContractWallet_Successfully(
        uint256 amount,
        uint256 expiry
    ) public {
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
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert(
            bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed")
        );
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    // tries depositing using a wallet that does not comply with EIP 1271
    function testDepositIntoStrategyWithSignature_WithContractWallet_NonconformingWallet(
        uint256 amount,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public {
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

    function testDepositIntoStrategyWithSignatureRevertsWhenDepositsPaused() public {
        address staker = cheats.addr(privateKey);

        // pause deposits
        cheats.startPrank(pauser);
        strategyManager.pause(1);
        cheats.stopPrank();

        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage = "Pausable: index is paused";
        _depositIntoStrategyWithSignature(staker, 1e18, type(uint256).max, expectedRevertMessage);
    }

    function testDepositIntoStrategyWithSignatureRevertsWhenReentering() public {
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
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry)
            );
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
            bytes memory calldataToUse = abi.encodeWithSelector(
                StrategyManager.depositIntoStrategy.selector,
                address(reenterer),
                dummyToken,
                amount
            );
            reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));
        }

        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore + shareAmountToReturn, "sharesAfter != sharesBefore + shareAmountToReturn");
        require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
    }

    function testDepositIntoStrategyWithSignatureRevertsWhenSignatureExpired() public {
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
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry)
            );
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

    function testDepositIntoStrategyWithSignatureRevertsWhenSignatureInvalid() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = block.timestamp;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), strategy, token, amount, nonceBefore, expiry)
            );
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

        require(
            strategyManager.stakerStrategyListLength(staker) == MAX_STAKER_STRATEGY_LIST_LENGTH,
            "strategyManager.stakerStrategyListLength(staker) != MAX_STAKER_STRATEGY_LIST_LENGTH"
        );

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
        dummyStrat = StrategyBase(address(new Reverter()));

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
        dummyStrat = StrategyBase(address(5678));

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

    function testSetStrategyWhitelister(address newWhitelister) external {
        address previousStrategyWhitelister = strategyManager.strategyWhitelister();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyWhitelisterChanged(previousStrategyWhitelister, newWhitelister);
        strategyManager.setStrategyWhitelister(newWhitelister);
        require(
            strategyManager.strategyWhitelister() == newWhitelister,
            "strategyManager.strategyWhitelister() != newWhitelister"
        );
    }

    function testSetStrategyWhitelisterRevertsWhenCalledByNotOwner(
        address notOwner
    ) external filterFuzzedAddressInputs(notOwner) {
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
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategyArray[i]),
                "strategy not properly whitelisted"
            );
        }

        return strategyArray;
    }

    function testAddStrategiesToDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(
        address notStrategyWhitelister
    ) external filterFuzzedAddressInputs(notStrategyWhitelister) {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IStrategy _strategy = deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = _strategy;

        cheats.startPrank(notStrategyWhitelister);
        cheats.expectRevert(bytes("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"));
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        cheats.stopPrank();
    }

    function testRemoveStrategiesFromDepositWhitelist(
        uint8 numberOfStrategiesToAdd,
        uint8 numberOfStrategiesToRemove
    ) external {
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
                require(
                    !strategyManager.strategyIsWhitelistedForDeposit(strategiesToRemove[i]),
                    "strategy not properly removed from whitelist"
                );
            } else {
                require(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesAdded[i]),
                    "strategy improperly removed from whitelist?"
                );
            }
        }
    }

    function testRemoveStrategiesFromDepositWhitelistRevertsWhenCalledByNotStrategyWhitelister(
        address notStrategyWhitelister
    ) external filterFuzzedAddressInputs(notStrategyWhitelister) {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = testAddStrategiesToDepositWhitelist(1);

        cheats.startPrank(notStrategyWhitelister);
        cheats.expectRevert(bytes("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister"));
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        cheats.stopPrank();
    }

    function testAddSharesRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.addShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testAddSharesRevertsStakerZeroAddress(uint256 amount) external {
        cheats.expectRevert(bytes("StrategyManager._addShares: staker cannot be zero address"));
        delegationManagerMock.addShares(strategyManager, address(0), dummyStrat, amount);
    }

    function testAddSharesRevertsZeroShares(address staker) external {
        cheats.assume(staker != address(0));
        cheats.expectRevert(bytes("StrategyManager._addShares: shares should not be zero!"));
        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, 0);
    }

    function testAddSharesAppendsStakerStrategyList(address staker, uint256 amount) external {
        cheats.assume(staker != address(0) && amount != 0);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(sharesBefore == 0, "Staker has already deposited into this strategy");
        require(!_isDepositedStrategy(staker, dummyStrat), "strategy shouldn't be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, amount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
        );
        require(sharesAfter == amount, "sharesAfter != amount");
        require(_isDepositedStrategy(staker, dummyStrat), "strategy should be deposited");
    }

    function testAddSharesExistingShares(address staker, uint256 sharesAmount) external {
        cheats.assume(staker != address(0) && 0 < sharesAmount && sharesAmount <= dummyToken.totalSupply());
        uint256 initialAmount = 1e18;
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, initialAmount);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(sharesBefore == initialAmount, "Staker has not deposited into strategy");
        require(_isDepositedStrategy(staker, strategy), "strategy should be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore"
        );
        require(sharesAfter == sharesBefore + sharesAmount, "sharesAfter != sharesBefore + amount");
        require(_isDepositedStrategy(staker, strategy), "strategy should be deposited");
    }

    function testRemoveSharesRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testRemoveSharesRevertsShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 removeSharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount too high"));
        delegationManagerMock.removeShares(strategyManager, staker, strategy, removeSharesAmount);
    }

    function testRemoveSharesRemovesStakerStrategyListSingleStrat(address staker, uint256 sharesAmount) external {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, sharesAmount);

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        require(sharesBefore == sharesAmount, "Staker has not deposited amount into strategy");

        delegationManagerMock.removeShares(strategyManager, staker, strategy, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        require(sharesAfter == 0, "sharesAfter != 0");
        require(!_isDepositedStrategy(staker, strategy), "strategy should not be part of staker strategy list");
    }

    // Remove Strategy from staker strategy list with multiple strategies deposited
    function testRemoveSharesRemovesStakerStrategyListMultipleStrat(
        address staker,
        uint256[3] memory amounts,
        uint8 randStrategy
    ) external {
        cheats.assume(staker != address(0));
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        for (uint256 i = 0; i < 3; ++i) {
            cheats.assume(amounts[i] > 0 && amounts[i] < dummyToken.totalSupply());
            _depositIntoStrategySuccessfully(strategies[i], staker, amounts[i]);
        }
        IStrategy removeStrategy = strategies[randStrategy % 3];
        uint256 removeAmount = amounts[randStrategy % 3];

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            require(sharesBefore[i] == amounts[i], "Staker has not deposited amount into strategy");
            require(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }

        delegationManagerMock.removeShares(strategyManager, staker, removeStrategy, removeAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, removeStrategy);
        require(
            stakerStrategyListLengthAfter == stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        require(sharesAfter == 0, "sharesAfter != 0");
        require(!_isDepositedStrategy(staker, removeStrategy), "strategy should not be part of staker strategy list");
    }

    // removeShares() from staker strategy list with multiple strategies deposited. Only callable by DelegationManager
    function testRemoveShares(uint256[3] memory depositAmounts, uint256[3] memory sharesAmounts) external {
        address staker = address(this);
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            cheats.assume(sharesAmounts[i] > 0 && sharesAmounts[i] <= depositAmounts[i]);
            _depositIntoStrategySuccessfully(strategies[i], staker, depositAmounts[i]);
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            require(sharesBefore[i] == depositAmounts[i], "Staker has not deposited amount into strategy");
            require(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        uint256 numPoppedStrategies = 0;
        uint256[] memory sharesAfter = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            delegationManagerMock.removeShares(strategyManager, staker, strategies[i], sharesAmounts[i]);
            sharesAfter[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            if (sharesAmounts[i] == depositAmounts[i]) {
                ++numPoppedStrategies;
                require(
                    !_isDepositedStrategy(staker, strategies[i]),
                    "strategy should not be part of staker strategy list"
                );
                require(sharesAfter[i] == 0, "sharesAfter != 0");
            } else {
                require(_isDepositedStrategy(staker, strategies[i]), "strategy should be part of staker strategy list");
                require(
                    sharesAfter[i] == sharesBefore[i] - sharesAmounts[i],
                    "sharesAfter != sharesBefore - sharesAmounts"
                );
            }
        }
        require(
            stakerStrategyListLengthBefore - numPoppedStrategies == strategyManager.stakerStrategyListLength(staker),
            "stakerStrategyListLengthBefore - numPoppedStrategies != strategyManager.stakerStrategyListLength(staker)"
        );
    }

    function testWithdrawSharesAsTokensRevertsDelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(bytes("StrategyManager.onlyDelegationManager: not the DelegationManager"));
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testWithdrawSharesAsTokensRevertsShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply() && depositAmount < sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(bytes("StrategyBase.withdraw: amountShares must be less than or equal to totalShares"));
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
    }

    function testWithdrawSharesAsTokensSingleStrategyDeposited(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply() && depositAmount >= sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint256 balanceBefore = token.balanceOf(staker);
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
        uint256 balanceAfter = token.balanceOf(staker);
        require(balanceAfter == balanceBefore + sharesAmount, "balanceAfter != balanceBefore + sharesAmount");
    }

    // INTERNAL / HELPER FUNCTIONS
    function _setUpQueuedWithdrawalStructSingleStrat(
        address staker,
        address withdrawer,
        IERC20 token,
        IStrategy strategy,
        uint256 shareAmount
    )
        internal
        view
        returns (
            IDelegationManager.Withdrawal memory queuedWithdrawal,
            IERC20[] memory tokensArray,
            bytes32 withdrawalRoot
        )
    {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = shareAmount;
        queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: staker,
            withdrawer: withdrawer,
            nonce: delegationManagerMock.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            delegatedTo: strategyManager.delegation().delegatedTo(staker)
        });
        // calculate the withdrawal root
        withdrawalRoot = delegationManagerMock.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, tokensArray, withdrawalRoot);
    }

    function _depositIntoStrategySuccessfully(
        IStrategy strategy,
        address staker,
        uint256 amount
    ) internal filterFuzzedAddressInputs(staker) {
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
            require(
                stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
            );
            require(
                strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == strategy,
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
            );
        }
    }

    function _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(
        address staker,
        address withdrawer,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts
    ) internal view returns (IDelegationManager.Withdrawal memory queuedWithdrawal, bytes32 withdrawalRoot) {
        queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: staker,
            withdrawer: withdrawer,
            nonce: delegationManagerMock.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            delegatedTo: strategyManager.delegation().delegatedTo(staker)
        });
        // calculate the withdrawal root
        withdrawalRoot = delegationManagerMock.calculateWithdrawalRoot(queuedWithdrawal);
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

    // internal function for de-duping code. expects success if `expectedRevertMessage` is empty and expiry is valid.
    function _depositIntoStrategyWithSignature(
        address staker,
        uint256 amount,
        uint256 expiry,
        string memory expectedRevertMessage
    ) internal returns (bytes memory) {
        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= dummyToken.balanceOf(address(this)));

        uint256 nonceBefore = strategyManager.nonces(staker);
        bytes memory signature;

        {
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), dummyStrat, dummyToken, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);

        bool expectedRevertMessageIsempty;
        {
            string memory emptyString;
            expectedRevertMessageIsempty =
                keccak256(abi.encodePacked(expectedRevertMessage)) == keccak256(abi.encodePacked(emptyString));
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
        uint256 shares = strategyManager.depositIntoStrategyWithSignature(
            dummyStrat,
            dummyToken,
            amount,
            staker,
            expiry,
            signature
        );

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        uint256 nonceAfter = strategyManager.nonces(staker);

        if (expiry >= block.timestamp && expectedRevertMessageIsempty) {
            require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
            require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
        }
        return signature;
    }

    /**
     * @notice internal function to help check if a strategy is part of list of deposited strategies for a staker
     * Used to check if removed correctly after withdrawing all shares for a given strategy
     */
    function _isDepositedStrategy(address staker, IStrategy strategy) internal view returns (bool) {
        uint256 stakerStrategyListLength = strategyManager.stakerStrategyListLength(staker);
        for (uint256 i = 0; i < stakerStrategyListLength; ++i) {
            if (strategyManager.stakerStrategyList(staker, i) == strategy) {
                return true;
            }
        }
        return false;
    }
}
