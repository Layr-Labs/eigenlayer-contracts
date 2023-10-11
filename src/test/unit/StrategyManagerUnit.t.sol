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

    uint256 public REQUIRED_BALANCE_WEI = 31 ether;

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
                        0/*initialPausedStatus*/
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

    function testDepositIntoStrategyRevertsWhenStakerFrozen() public {
        uint256 amount = 1e18;
        address staker = address(this);

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.expectRevert(
            bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
        );
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

    function testDepositIntoStrategyWithSignatureRevertsWhenStakerFrozen() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
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

        // freeze the staker
        slasherMock.freezeOperator(staker);

        cheats.expectRevert(
            bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
        );
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
        require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
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

    // Comment out withdraw tests to be moved

    // // queue and complete withdrawal. Ensure that strategy is no longer part
    // function testQueueWithdrawalFullyWithdraw(uint256 amount) external {
    //     address staker = address(this);
    //     IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
    //     // deposit and withdraw the same amount
    //     testQueueWithdrawal_ToSelf(amount, amount);
    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = _tempStrategyStorage;
    //         shareAmounts[0] = amount;
    //         tokensArray[0] = dummyToken;
    //     }

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

    //     {
    //         uint256 nonce = strategyManager.numWithdrawalsQueued(staker);

    //         IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
    //             withdrawer: staker,
    //             nonce: (uint96(nonce) - 1)
    //         });
    //         queuedWithdrawal = IStrategyManager.QueuedWithdrawal({
    //             strategies: strategyArray,
    //             shares: shareAmounts,
    //             depositor: staker,
    //             withdrawerAndNonce: withdrawerAndNonce,
    //             withdrawalStartBlock: uint32(block.number),
    //             delegatedAddress: strategyManager.delegation().delegatedTo(staker)
    //         });
    //     }

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(staker));

    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalCompleted(
    //         queuedWithdrawal.depositor,
    //         queuedWithdrawal.withdrawerAndNonce.nonce,
    //         queuedWithdrawal.withdrawerAndNonce.withdrawer,
    //         strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     strategyManager.completeQueuedWithdrawal(
    //         queuedWithdrawal,
    //         tokensArray,
    //         /*middlewareTimesIndex*/ 0,
    //         /*receiveAsTokens*/ true
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore + amount, "balanceAfter != balanceBefore + withdrawalAmount");
    //     require(
    //         !_isDepositedStrategy(staker, strategyArray[0]),
    //         "Strategy still part of staker's deposited strategies"
    //     );
    //     require(sharesAfter == 0, "staker shares is not 0");
    // }

    // function testQueueWithdrawalRevertsWhenStakerFrozen(uint256 depositAmount, uint256 withdrawalAmount) public {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);
    //     IStrategy strategy = dummyStrat;
    //     IERC20 token = dummyToken;

    //     testDepositIntoStrategySuccessfully(staker, depositAmount);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal /*IERC20[] memory tokensArray*/,
    //         ,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat(staker, /*withdrawer*/ staker, token, strategy, withdrawalAmount);

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 nonceBefore = delegationManagerMock.cumulativeWithdrawalsQueued(staker);

    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");

    //     // freeze the staker
    //     slasherMock.freezeOperator(staker);

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;
    //     cheats.expectRevert(
    //         bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
    //     );
    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ staker
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 nonceAfter = delegationManagerMock.cumulativeWithdrawalsQueued(address(this));

    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingAfter is true!");
    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(nonceAfter == nonceBefore, "nonceAfter != nonceBefore");
    // }

    // function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedFalse(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);
    //     IStrategy strategy = dummyStrat;

    //     testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = strategy;
    //         shareAmounts[0] = withdrawalAmount;
    //         tokensArray[0] = dummyToken;
    //     }

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

    //     {
    //         uint256 nonce = delegationManagerMock.cumulativeWithdrawalsQueued(staker);

    //         queuedWithdrawal = 
    //             IDelegationManager.Withdrawal({
    //                 strategies: strategyArray,
    //                 shares: shareAmounts,
    //                 staker: staker,
    //                 withdrawer: staker,
    //                 nonce: (nonce - 1),
    //                 startBlock: uint32(block.number),
    //                 delegatedTo: strategyManager.delegation().delegatedTo(staker)
    //             }
    //         );
    //     }

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(staker));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalCompleted(
    //         queuedWithdrawal.depositor,
    //         queuedWithdrawal.withdrawerAndNonce.nonce,
    //         queuedWithdrawal.withdrawerAndNonce.withdrawer,
    //         strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);
    //     _tempStrategyStorage = dummyStrat;

    //     testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = _tempStrategyStorage;
    //         shareAmounts[0] = withdrawalAmount;
    //         tokensArray[0] = dummyToken;
    //     }

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

    //     {
    //         uint256 nonce = delegationManagerMock.cumulativeWithdrawalsQueued(staker);

    //         queuedWithdrawal = 
    //             IDelegationManager.Withdrawal({
    //                 strategies: strategyArray,
    //                 shares: shareAmounts,
    //                 staker: staker,
    //                 withdrawer: staker,
    //                 nonce: (nonce - 1),
    //                 startBlock: uint32(block.number),
    //                 delegatedTo: strategyManager.delegation().delegatedTo(staker)
    //             }
    //         );
    //     }

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(staker));

    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalCompleted(
    //         queuedWithdrawal.depositor,
    //         queuedWithdrawal.withdrawerAndNonce.nonce,
    //         queuedWithdrawal.withdrawerAndNonce.withdrawer,
    //         strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     strategyManager.completeQueuedWithdrawal(
    //         queuedWithdrawal,
    //         tokensArray,
    //         /*middlewareTimesIndex*/ 0,
    //         /*receiveAsTokens*/ true
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, _tempStrategyStorage);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore + withdrawalAmount, "balanceAfter != balanceBefore + withdrawalAmount");
    //     if (depositAmount == withdrawalAmount) {
    //         // Since receiving tokens instead of shares, if withdrawal amount is entire deposit, then strategy will be removed
    //         // with sharesAfter being 0
    //         require(
    //             !_isDepositedStrategy(staker, _tempStrategyStorage),
    //             "Strategy still part of staker's deposited strategies"
    //         );
    //         require(sharesAfter == 0, "staker shares is not 0");
    //     }
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalsPaused(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     // pause withdrawals
    //     cheats.startPrank(pauser);
    //     strategyManager.pause(2);
    //     cheats.stopPrank();

    //     cheats.expectRevert(bytes("Pausable: index is paused"));
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalFailsWhenTokensInputLengthMismatch(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = true;
    //     // mismatch tokens array by setting tokens array to empty array
    //     tokensArray = new IERC20[](0);

    //     cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: input length mismatch"));
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenDelegatedAddressFrozen(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     // freeze the delegatedAddress
    //     slasherMock.freezeOperator(strategyManager.delegation().delegatedTo(_tempStakerStorage));

    //     cheats.expectRevert(
    //         bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenAttemptingReentrancy(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     // replace dummyStrat with Reenterer contract
    //     reenterer = new Reenterer();
    //     dummyStrat = StrategyBase(address(reenterer));

    //     // whitelist the strategy for deposit
    //     cheats.startPrank(strategyManager.owner());
    //     IStrategy[] memory _strategy = new IStrategy[](1);
    //     _strategy[0] = dummyStrat;
    //     for (uint256 i = 0; i < _strategy.length; ++i) {
    //         cheats.expectEmit(true, true, true, true, address(strategyManager));
    //         emit StrategyAddedToDepositWhitelist(_strategy[i]);
    //     }
    //     strategyManager.addStrategiesToDepositWhitelist(_strategy);
    //     cheats.stopPrank();

    //     _tempStakerStorage = address(this);
    //     IStrategy strategy = dummyStrat;

    //     reenterer.prepareReturnData(abi.encode(depositAmount));

    //     testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = strategy;
    //         shareAmounts[0] = withdrawalAmount;
    //         tokensArray[0] = dummyToken;
    //     }

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

    //     {
    //         uint256 nonce = delegationManagerMock.cumulativeWithdrawalsQueued(_tempStakerStorage);

    //         queuedWithdrawal = 
    //             IDelegationManager.Withdrawal({
    //                 strategies: strategyArray,
    //                 shares: shareAmounts,
    //                 staker: _tempStakerStorage,
    //                 withdrawer: _tempStakerStorage,
    //                 nonce: (nonce - 1),
    //                 startBlock: uint32(block.number),
    //                 delegatedTo: strategyManager.delegation().delegatedTo(_tempStakerStorage)
    //             }
    //         );
    //     }

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     address targetToUse = address(strategyManager);
    //     uint256 msgValueToUse = 0;
    //     bytes memory calldataToUse = abi.encodeWithSelector(
    //         StrategyManager.completeQueuedWithdrawal.selector,
    //         queuedWithdrawal,
    //         tokensArray,
    //         middlewareTimesIndex,
    //         receiveAsTokens
    //     );
    //     reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalDoesNotExist() external {
    //     _tempStakerStorage = address(this);
    //     uint256 withdrawalAmount = 1e18;
    //     IStrategy strategy = dummyStrat;

    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = strategy;
    //         shareAmounts[0] = withdrawalAmount;
    //         tokensArray[0] = dummyToken;
    //     }

    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

    //     {
    //         IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
    //             withdrawer: _tempStakerStorage,
    //             nonce: 0
    //         });
    //         queuedWithdrawal = IStrategyManager.QueuedWithdrawal({
    //             strategies: strategyArray,
    //             shares: shareAmounts,
    //             depositor: _tempStakerStorage,
    //             withdrawerAndNonce: withdrawerAndNonce,
    //             withdrawalStartBlock: uint32(block.number),
    //             delegatedAddress: strategyManager.delegation().delegatedTo(_tempStakerStorage)
    //         });
    //     }

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawal is not pending"));
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenCanWithdrawReturnsFalse(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     // prepare mock
    //     slasherMock.setCanWithdrawResponse(false);

    //     cheats.expectRevert(
    //         bytes("StrategyManager.completeQueuedWithdrawal: shares pending withdrawal are still slashable")
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenNotCallingFromWithdrawerAddress(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.startPrank(address(123456));
    //     cheats.expectRevert(
    //         bytes(
    //             "StrategyManager.completeQueuedWithdrawal: only specified withdrawer can complete a queued withdrawal"
    //         )
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    //     cheats.stopPrank();

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenTryingToCompleteSameWithdrawal2X(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     (IDelegationManager.Withdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, /*bytes32 withdrawalRoot*/) =
    //         testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = queuedWithdrawal.strategies[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalCompleted(
    //         queuedWithdrawal.depositor,
    //         queuedWithdrawal.nonce,
    //         queuedWithdrawal.withdrawer,
    //         strategyManager.calculateWithdrawalRoot(queuedWithdrawal)
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalDelayBlocksHasNotPassed(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    // //     // try to complete same withdrawal again
    // //     cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawal is not pending"));
    // //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // // }

    //     uint256 valueToSet = 1;
    //     // set the `withdrawalDelayBlocks` variable
    //     cheats.startPrank(strategyManager.owner());
    //     uint256 previousValue = strategyManager.withdrawalDelayBlocks();
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
    //     strategyManager.setWithdrawalDelayBlocks(valueToSet);
    //     cheats.stopPrank();
    //     require(
    //         strategyManager.withdrawalDelayBlocks() == valueToSet,
    //         "strategyManager.withdrawalDelayBlocks() != valueToSet"
    //     );

    //     cheats.expectRevert(
    //         bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // }

    // function testCompleteQueuedWithdrawalWithNonzeroWithdrawalDelayBlocks(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount,
    //     uint16 valueToSet
    // ) external {
    //     // filter fuzzed inputs to allowed *and nonzero* amounts
    //     cheats.assume(valueToSet <= strategyManager.MAX_WITHDRAWAL_DELAY_BLOCKS() && valueToSet != 0);
    //     cheats.assume(depositAmount != 0 && withdrawalAmount != 0);
    //     cheats.assume(depositAmount >= withdrawalAmount);
    //     address staker = address(this);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray /*bytes32 withdrawalRoot*/,

    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    // //     cheats.expectRevert(bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed"));
    // //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // // }

    //     // set the `withdrawalDelayBlocks` variable
    //     cheats.startPrank(strategyManager.owner());
    //     uint256 previousValue = strategyManager.withdrawalDelayBlocks();
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
    //     strategyManager.setWithdrawalDelayBlocks(valueToSet);
    //     cheats.stopPrank();
    //     require(
    //         strategyManager.withdrawalDelayBlocks() == valueToSet,
    //         "strategyManager.withdrawalDelayBlocks() != valueToSet"
    //     );

    //     cheats.expectRevert(
    //         bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), dummyStrat);
    //     uint256 balanceBefore = dummyToken.balanceOf(address(staker));

    //     // roll block number forward to one block before the withdrawal should be completeable and attempt again
    //     uint256 originalBlockNumber = block.number;
    //     cheats.roll(originalBlockNumber + valueToSet - 1);
    //     cheats.expectRevert(
    //         bytes("StrategyManager.completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     // roll block number forward to the block at which the withdrawal should be completeable, and complete it
    //     cheats.roll(originalBlockNumber + valueToSet);
    //     strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), dummyStrat);
    //     uint256 balanceAfter = dummyToken.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    

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

    // _removeShares() needs to be tested using a StrategyManager Harness contrat to access internal _removeShares
    // no longer has queueWithdrawal call. Same applies to _removeStrategyFromStakerStrategyList()
    // function test_removeSharesRevertsWhenShareAmountIsZero(uint256 depositAmount) external {
    //     address staker = address(this);
    //     uint256 withdrawalAmount = 0;

    //     testDepositIntoStrategySuccessfully(staker, depositAmount);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             dummyToken,
    //             _tempStrategyStorage,
    //             withdrawalAmount
    //         );
    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount should not be zero!"));
    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );
    // }

    // function test_removeSharesRevertsWhenShareAmountIsTooLarge(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(depositAmount > 0 && withdrawalAmount > depositAmount);
    //     address staker = address(this);

    //     testDepositIntoStrategySuccessfully(staker, depositAmount);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             dummyToken,
    //             _tempStrategyStorage,
    //             withdrawalAmount
    //         );
    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount too high"));
    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );
    // }

    // /**
    //  * Testing that removal of all 3 strategies from a staker's strategy list works even if the strategyIndexes are not sorted
    //  * in descending order, in this test case they are in ascending order [0,1,2].
    //  */
    // function test_removeStrategyFromStakerStrategyListWithAscendingIndexInput(uint256[3] memory amounts) external {
    //     // filtering of fuzzed inputs
    //     cheats.assume(amounts[0] != 0 && amounts[1] != 0 && amounts[2] != 0);
    //     address staker = address(this);

    //     // Setup input params
    //     IStrategy[] memory strategies = new IStrategy[](3);
    //     strategies[0] = dummyStrat;
    //     strategies[1] = dummyStrat2;
    //     strategies[2] = dummyStrat3;
    //     uint256[] memory depositAmounts = new uint256[](3);
    //     depositAmounts[0] = amounts[0];
    //     depositAmounts[1] = amounts[1];
    //     depositAmounts[2] = amounts[2];

    //     _depositIntoStrategySuccessfully(dummyStrat, staker, depositAmounts[0]);
    //     _depositIntoStrategySuccessfully(dummyStrat2, staker, depositAmounts[1]);
    //     _depositIntoStrategySuccessfully(dummyStrat3, staker, depositAmounts[2]);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             strategies,
    //             depositAmounts
    //         );
    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    //     uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);
    //     uint256[] memory sharesBefore = new uint256[](3);
    //     sharesBefore[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesBefore[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesBefore[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);
    //     uint256[] memory strategyIndexes = new uint256[](3);
    //     // Correct index for first but incorrect for second and third after first strategy is removed from the list
    //     strategyIndexes[0] = 0;
    //     strategyIndexes[1] = 1;
    //     strategyIndexes[2] = 2;

    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );

    //     uint256[] memory sharesAfter = new uint256[](3);
    //     sharesAfter[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesAfter[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesAfter[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);

    //     require(!_isDepositedStrategy(staker, strategies[0]), "Strategy still part of staker's deposited strategies");
    //     require(!_isDepositedStrategy(staker, strategies[1]), "Strategy still part of staker's deposited strategies");
    //     require(!_isDepositedStrategy(staker, strategies[2]), "Strategy still part of staker's deposited strategies");
    //     for (uint256 i = 0; i < sharesAfter.length; ++i) {
    //         require(sharesAfter[i] == 0, "Strategy still has shares for staker");
    //     }
    // }

    // /**
    //  * Testing that removal of all 3 strategies from a staker's strategy list works even if the strategyIndexes are not sorted
    //  * in descending order, in this test case they are in ascending order [0,1,2].
    //  */
    // function test_removeStrategyFromStakerStrategyListWithMultipleStrategyIndexes(uint256[3] memory amounts) external {
    //     // filtering of fuzzed inputs
    //     cheats.assume(amounts[0] != 0 && amounts[1] != 0 && amounts[2] != 0);
    //     address staker = address(this);

    //     // Setup input params
    //     IStrategy[] memory strategies = new IStrategy[](3);
    //     strategies[0] = dummyStrat;
    //     strategies[1] = dummyStrat2;
    //     strategies[2] = dummyStrat3;
    //     uint256[] memory depositAmounts = new uint256[](3);
    //     depositAmounts[0] = amounts[0];
    //     depositAmounts[1] = amounts[1];
    //     depositAmounts[2] = amounts[2];

    //     _depositIntoStrategySuccessfully(dummyStrat, staker, depositAmounts[0]);
    //     _depositIntoStrategySuccessfully(dummyStrat2, staker, depositAmounts[1]);
    //     _depositIntoStrategySuccessfully(dummyStrat3, staker, depositAmounts[2]);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat_MultipleStrategies(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             strategies,
    //             depositAmounts
    //         );
    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    //     uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);
    //     uint256[] memory sharesBefore = new uint256[](3);
    //     sharesBefore[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesBefore[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesBefore[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);
    //     uint256[] memory strategyIndexes = new uint256[](3);
    //     // Correct index for first but incorrect for second and third after first strategy is removed from the list
    //     strategyIndexes[0] = 2;
    //     strategyIndexes[1] = 1;
    //     strategyIndexes[2] = 0;

    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );

    //     uint256[] memory sharesAfter = new uint256[](3);
    //     sharesAfter[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesAfter[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesAfter[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);

    //     require(!_isDepositedStrategy(staker, strategies[0]), "Strategy still part of staker's deposited strategies");
    //     require(!_isDepositedStrategy(staker, strategies[1]), "Strategy still part of staker's deposited strategies");
    //     require(!_isDepositedStrategy(staker, strategies[2]), "Strategy still part of staker's deposited strategies");
    //     for (uint256 i = 0; i < sharesAfter.length; ++i) {
    //         require(sharesAfter[i] == 0, "Strategy still has shares for staker");
    //     }
    // }

    // function test_removeStrategyFromStakerStrategyListWithIncorrectIndexInput(
    //     uint256 incorrectIndex,
    //     uint256 amount
    // ) external {
    //     // filtering of fuzzed inputs
    //     cheats.assume(amount != 0 && incorrectIndex != 0);
    //     address staker = address(this);
    //     IERC20 token = dummyToken;
    //     IStrategy strategy = dummyStrat;

    //     _depositIntoStrategySuccessfully(dummyStrat, staker, amount);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         ,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             token,
    //             strategy,
    //             amount
    //         );
    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    //     uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);
    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = incorrectIndex;

    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
    //     require(!_isDepositedStrategy(staker, strategy), "Strategy still part of staker's deposited strategies");
    //     require(sharesAfter == 0, "Strategy still has shares for staker");
    // }

    // function test_removeStrategyFromStakerStrategyListWithCorrectIndexInput(uint256 amount) external {
    //     // filtering of fuzzed inputs
    //     cheats.assume(amount != 0);
    //     address staker = address(this);
    //     IERC20 token = dummyToken;
    //     IStrategy strategy = dummyStrat;

    //     _depositIntoStrategySuccessfully(dummyStrat, staker, amount);

    //     (
    //         IStrategyManager.QueuedWithdrawal memory queuedWithdrawal,
    //         ,
    //         bytes32 withdrawalRoot
    //     ) = _setUpQueuedWithdrawalStructSingleStrat(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             token,
    //             strategy,
    //             amount
    //         );
    //     require(!strategyManager.withdrawalRootPending(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    //     uint256 nonceBefore = strategyManager.numWithdrawalsQueued(staker);
    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256[] memory strategyIndexes = new uint256[](1);
    //     strategyIndexes[0] = 0;

    //     strategyManager.queueWithdrawal(
    //         strategyIndexes,
    //         queuedWithdrawal.strategies,
    //         queuedWithdrawal.shares,
    //         /*withdrawer*/ address(this)
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
    //     require(!_isDepositedStrategy(staker, strategy), "Strategy still part of staker's deposited strategies");
    //     require(sharesAfter == 0, "Strategy still has shares for staker");
    // }

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

    // INTERNAL / HELPER FUNCTIONS
    function _setUpQueuedWithdrawalStructSingleStrat(address staker, address withdrawer, IERC20 token, IStrategy strategy, uint256 shareAmount)
        internal view returns (IDelegationManager.Withdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, bytes32 withdrawalRoot)
    {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = shareAmount;
        queuedWithdrawal = 
            IDelegationManager.Withdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                staker: staker,
                withdrawer: withdrawer,
                nonce: delegationManagerMock.cumulativeWithdrawalsQueued(staker),
                startBlock: uint32(block.number),
                delegatedTo: strategyManager.delegation().delegatedTo(staker)
            }
        );
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
    )
        internal view returns (IDelegationManager.Withdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    {
        queuedWithdrawal = 
            IDelegationManager.Withdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                staker: staker,
                withdrawer: withdrawer,
                nonce: delegationManagerMock.cumulativeWithdrawalsQueued(staker),
                startBlock: uint32(block.number),
                delegatedTo: strategyManager.delegation().delegatedTo(staker)
            }
        );
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
