// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "src/test/mocks/ERC20Mock.sol";
import "src/test/mocks/ERC20_SetTransferReverting_Mock.sol";
import "src/test/mocks/Reverter.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/MockDecimals.sol";
import "src/test/events/IStrategyManagerEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the StrategyManager contract, entire withdrawal tests related to the
 * DelegationManager are not tested here but callable functions by the DelegationManager are mocked and tested here.
 * Contracts tested: StrategyManager.sol
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract StrategyManagerUnitTests is EigenLayerUnitTestSetup, IStrategyManagerEvents {
    StrategyManager public strategyManagerImplementation;
    StrategyManager public strategyManager;

    IERC20 public dummyToken;
    ERC20_SetTransferReverting_Mock public revertToken;
    StrategyBase public dummyStrat;
    StrategyBase public dummyStrat2;
    StrategyBase public dummyStrat3;

    Reenterer public reenterer;

    address initialOwner = address(this);
    uint256 public privateKey = 111111;
    address constant dummyAdmin = address(uint160(uint256(keccak256("DummyAdmin"))));

    function setUp() public override {
        EigenLayerUnitTestSetup.setUp();
        strategyManagerImplementation = new StrategyManager(delegationManagerMock, eigenPodManagerMock, slasherMock);
        strategyManager = StrategyManager(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyManagerImplementation),
                    address(eigenLayerProxyAdmin),
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
        revertToken = new ERC20_SetTransferReverting_Mock(1000e18, address(this));
        revertToken.setTransfersRevert(true);
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyStrat2 = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyStrat3 = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategies = new IStrategy[](3);
        _strategies[0] = dummyStrat;
        _strategies[1] = dummyStrat2;
        _strategies[2] = dummyStrat3;
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](3);
        _thirdPartyTransfersForbiddenValues[0] = false;
        _thirdPartyTransfersForbiddenValues[1] = false;
        _thirdPartyTransfersForbiddenValues[2] = false;
        for (uint256 i = 0; i < _strategies.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategies[i]);
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit UpdatedThirdPartyTransfersForbidden(_strategies[i], _thirdPartyTransfersForbiddenValues[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategies, _thirdPartyTransfersForbiddenValues);

        addressIsExcludedFromFuzzedInputs[address(reenterer)] = true;
    }

    // INTERNAL / HELPER FUNCTIONS
    function _deployNewStrategy(
        IERC20 _token,
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        address admin
    ) public returns (StrategyBase) {
        StrategyBase newStrategy = new StrategyBase(_strategyManager);
        newStrategy = StrategyBase(address(new TransparentUpgradeableProxy(address(newStrategy), address(admin), "")));
        newStrategy.initialize(_token, _pauserRegistry);
        return newStrategy;
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

        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, token, strategy, expectedShares);
        uint256 shares = strategyManager.depositIntoStrategy(strategy, token, amount);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

        assertEq(sharesAfter, sharesBefore + shares, "sharesAfter != sharesBefore + shares");
        if (sharesBefore == 0) {
            assertEq(
                stakerStrategyListLengthAfter,
                stakerStrategyListLengthBefore + 1,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
            );
            assertEq(
                address(strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1)),
                address(strategy),
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
            );
        }
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
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, dummyStrat, dummyToken, amount, nonceBefore, expiry)
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
            assertEq(sharesAfter, sharesBefore + shares, "sharesAfter != sharesBefore + shares");
            assertEq(nonceAfter, nonceBefore + 1, "nonceAfter != nonceBefore + 1");
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

    /**
     * @notice Deploys numberOfStrategiesToAdd new strategies and adds them to the whitelist
     */
    function _addStrategiesToWhitelist(uint8 numberOfStrategiesToAdd) internal returns (IStrategy[] memory) {
        IStrategy[] memory strategyArray = new IStrategy[](numberOfStrategiesToAdd);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](numberOfStrategiesToAdd);
        // loop that deploys a new strategy and adds it to the array
        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            IStrategy _strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategyArray[i] = _strategy;
            assertFalse(strategyManager.strategyIsWhitelistedForDeposit(_strategy), "strategy improperly whitelisted?");
        }

        cheats.prank(strategyManager.strategyWhitelister());
        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(strategyArray[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);

        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategyArray[i]), "strategy not whitelisted");
        }

        return strategyArray;
    }
}

contract StrategyManagerUnitTests_initialize is StrategyManagerUnitTests {
    function test_CannotReinitialize() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        strategyManager.initialize(initialOwner, initialOwner, pauserRegistry, 0);
    }

    function test_InitializedStorageProperly() public {
        assertEq(strategyManager.owner(), initialOwner, "strategyManager.owner() != initialOwner");
        assertEq(
            strategyManager.strategyWhitelister(),
            initialOwner,
            "strategyManager.strategyWhitelister() != initialOwner"
        );
        assertEq(
            address(strategyManager.pauserRegistry()),
            address(pauserRegistry),
            "strategyManager.pauserRegistry() != pauserRegistry"
        );
    }
}

contract StrategyManagerUnitTests_depositIntoStrategy is StrategyManagerUnitTests {
    function testFuzz_depositIntoStrategySuccessfully(
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

        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, token, strategy, expectedShares);
        uint256 shares = strategyManager.depositIntoStrategy(strategy, token, amount);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

        assertEq(sharesAfter, sharesBefore + shares, "sharesAfter != sharesBefore + shares");
        if (sharesBefore == 0) {
            assertEq(
                stakerStrategyListLengthAfter,
                stakerStrategyListLengthBefore + 1,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
            );
            assertEq(
                address(strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1)),
                address(strategy),
                "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
            );
        } else {
            assertEq(
                stakerStrategyListLengthAfter,
                stakerStrategyListLengthBefore,
                "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore"
            );
        }
    }

    function test_DepositWhenStrategySharesExist() public {
        address staker = address(this);
        uint256 amount = 1e18;
        testFuzz_depositIntoStrategySuccessfully(staker, amount);
        testFuzz_depositIntoStrategySuccessfully(staker, amount);
    }

    function test_Revert_WhenDepositsPaused() public {
        uint256 amount = 1e18;

        // pause deposits
        cheats.prank(pauser);
        strategyManager.pause(1);

        cheats.expectRevert("Pausable: index is paused");
        strategyManager.depositIntoStrategy(dummyStrat, dummyToken, amount);
    }

    function test_Revert_WhenReentering() public {
        uint256 amount = 1e18;

        reenterer = new Reenterer();

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = IStrategy(address(reenterer));
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy, thirdPartyTransfersForbiddenValues);
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

    function test_Revert_WhenTokenSafeTransferFromReverts() external {
        // replace 'dummyStrat' with one that uses a reverting token
        dummyToken = IERC20(address(new ReverterWithDecimals()));
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);

        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("Reverter: I am a contract that always reverts");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenTokenDoesNotExist() external {
        // replace 'dummyStrat' with one that uses a non-existent token, but will pass the initializer decimals check
        dummyToken = IERC20(address(new MockDecimals()));
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("SafeERC20: low-level call failed");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenStrategyDepositFunctionReverts() external {
        // replace 'dummyStrat' with one that always reverts
        dummyStrat = StrategyBase(address(new Reverter()));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("Reverter: I am a contract that always reverts");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenStrategyDoesNotExist() external {
        // replace 'dummyStrat' with one that does not exist
        dummyStrat = StrategyBase(address(5678));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);

        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert();
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenStrategyNotWhitelisted() external {
        // replace 'dummyStrat' with one that is not whitelisted
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_addShares_Revert_WhenSharesIsZero() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);

        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        address staker = address(this);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        reenterer.prepareReturnData(abi.encode(uint256(0)));

        cheats.prank(staker);
        cheats.expectRevert("StrategyManager._addShares: shares should not be zero!");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }
}

contract StrategyManagerUnitTests_depositIntoStrategyWithSignature is StrategyManagerUnitTests {
    function test_Revert_WhenSignatureInvalid() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint256 amount = 1e18;

        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = block.timestamp;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
        // call with `notStaker` as input instead of `staker` address
        address notStaker = address(3333);
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, notStaker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        assertEq(sharesAfter, sharesBefore, "sharesAfter != sharesBefore");
        assertEq(nonceAfter, nonceBefore, "nonceAfter != nonceBefore");
    }

    function testFuzz_DepositSuccessfully(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage;
        _depositIntoStrategyWithSignature(staker, amount, expiry, expectedRevertMessage);
    }

    function testFuzz_Revert_SignatureReplay(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);
        cheats.assume(expiry > block.timestamp);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        bytes memory signature = _depositIntoStrategyWithSignature(staker, amount, expiry, "");

        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
        strategyManager.depositIntoStrategyWithSignature(dummyStrat, dummyToken, amount, staker, expiry, signature);
    }

    // tries depositing using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature
    function testFuzz_Revert_WithContractWallet_BadSignature(uint256 amount) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;

        // deploy ERC1271WalletMock for staker to use
        cheats.prank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
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
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed");
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    // tries depositing using a wallet that does not comply with EIP 1271
    function testFuzz_Revert_WithContractWallet_NonconformingWallet(
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
        cheats.prank(staker);
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
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

    // Tries depositing without token approval and transfer fails. deposit function should also revert
    function test_Revert_WithContractWallet_TokenTransferFails() external {
        address staker = cheats.addr(privateKey);
        uint256 amount = 1e18;
        uint256 nonceBefore = strategyManager.nonces(staker);
        uint256 expiry = block.timestamp + 100;
        bytes memory signature;

        {
            bytes32 structHash = keccak256(
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, dummyStrat, revertToken, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert("ERC20: insufficient allowance");
        strategyManager.depositIntoStrategyWithSignature(dummyStrat, revertToken, amount, staker, expiry, signature);
    }

    // tries depositing using a signature and an EIP 1271 compliant wallet
    function testFuzz_WithContractWallet_Successfully(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);

        // deploy ERC1271WalletMock for staker to use
        cheats.prank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        staker = address(wallet);

        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage;
        _depositIntoStrategyWithSignature(staker, amount, expiry, expectedRevertMessage);
    }

    function test_Revert_WhenDepositsPaused() public {
        address staker = cheats.addr(privateKey);

        // pause deposits
        cheats.prank(pauser);
        strategyManager.pause(1);

        string memory expectedRevertMessage = "Pausable: index is paused";
        _depositIntoStrategyWithSignature(staker, 1e18, type(uint256).max, expectedRevertMessage);
    }

    /**
     * @notice reenterer contract which is configured as the strategy contract
     * is configured to call depositIntoStrategy after reenterer.deposit() is called from the
     * depositIntoStrategyWithSignature() is called from the StrategyManager. Situation is not likely to occur given
     * the strategy has to be whitelisted but it at least protects from reentrant attacks
     */
    function test_Revert_WhenReentering() public {
        reenterer = new Reenterer();

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);

        
        _strategy[0] = IStrategy(address(reenterer));
        for (uint256 i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
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
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

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
    }

    function test_Revert_WhenSignatureExpired() public {
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
                abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);

        cheats.expectRevert("StrategyManager.depositIntoStrategyWithSignature: signature expired");
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        uint256 nonceAfter = strategyManager.nonces(staker);

        assertEq(sharesAfter, sharesBefore, "sharesAfter != sharesBefore");
        assertEq(nonceAfter, nonceBefore, "nonceAfter != nonceBefore");
    }

    function test_Revert_WhenStrategyNotWhitelisted() external {
        // replace 'dummyStrat' with one that is not whitelisted
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyToken = dummyStrat.underlyingToken();
        address staker = cheats.addr(privateKey);
        uint256 amount = 1e18;

        string
            memory expectedRevertMessage = "StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted";
        _depositIntoStrategyWithSignature(staker, amount, type(uint256).max, expectedRevertMessage);
    }
    
    function testFuzz_Revert_WhenThirdPartyTransfersForbidden(uint256 amount, uint256 expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        cheats.prank(strategyManager.strategyWhitelister());
        strategyManager.setThirdPartyTransfersForbidden(dummyStrat, true);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        string memory expectedRevertMessage = "StrategyManager.depositIntoStrategyWithSignature: third transfers disabled";
        _depositIntoStrategyWithSignature(staker, amount, expiry, expectedRevertMessage);
    }
}

contract StrategyManagerUnitTests_removeShares is StrategyManagerUnitTests {
    /**
     * @notice Should revert if not called by DelegationManager
     */
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert("StrategyManager.onlyDelegationManager: not the DelegationManager");
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    /**
     * @notice deposits a single strategy and tests removeShares() function reverts when sharesAmount is 0
     */
    function testFuzz_Revert_ZeroShares(
        address staker,
        uint256 depositAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert("StrategyManager._removeShares: shareAmount should not be zero!");
        delegationManagerMock.removeShares(strategyManager, staker, strategy, 0);
    }

    /**
     * @notice deposits a single strategy and tests removeShares() function reverts when sharesAmount is
     * higher than depositAmount
     */
    function testFuzz_Revert_ShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 removeSharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert("StrategyManager._removeShares: shareAmount too high");
        delegationManagerMock.removeShares(strategyManager, staker, strategy, removeSharesAmount);
    }

    /**
     * @notice deposit single strategy and removeShares() for less than the deposited amount
     * Shares should be updated correctly with stakerStrategyListLength unchanged
     */
    function testFuzz_RemoveSharesLessThanDeposit(
        address staker,
        uint256 depositAmount,
        uint256 removeSharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > 0 && removeSharesAmount < depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        delegationManagerMock.removeShares(strategyManager, staker, strategy, removeSharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        assertEq(sharesBefore, sharesAfter + removeSharesAmount, "Remove incorrect amount of shares");
        assertEq(
            stakerStrategyListLengthBefore,
            stakerStrategyListLengthAfter,
            "stakerStrategyListLength shouldn't have changed"
        );
    }

    /**
     * @notice testing removeShares()
     * deposits 1 strategy and tests it is removed from staker strategy list after removing all shares
     */
    function testFuzz_RemovesStakerStrategyListSingleStrat(
        address staker,
        uint256 sharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, sharesAmount);

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
        assertEq(sharesBefore, sharesAmount, "Staker has not deposited amount into strategy");

        delegationManagerMock.removeShares(strategyManager, staker, strategy, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        assertEq(sharesAfter, 0, "sharesAfter != 0");
        assertFalse(_isDepositedStrategy(staker, strategy), "strategy should not be part of staker strategy list");
    }

    /**
     * @notice testing removeShares() function with 3 strategies deposited.
     * Randomly selects one of the 3 strategies to be fully removed from staker strategy list.
     * Only callable by DelegationManager
     */
    function testFuzz_RemovesStakerStrategyListMultipleStrat(
        address staker,
        uint256[3] memory amounts,
        uint8 randStrategy
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        for (uint256 i = 0; i < 3; ++i) {
            amounts[i] = bound(amounts[i], 1, dummyToken.totalSupply() - 1);
            _depositIntoStrategySuccessfully(strategies[i], staker, amounts[i]);
        }
        IStrategy removeStrategy = strategies[randStrategy % 3];
        uint256 removeAmount = amounts[randStrategy % 3];

        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            assertEq(sharesBefore[i], amounts[i], "Staker has not deposited amount into strategy");
            assertTrue(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }

        delegationManagerMock.removeShares(strategyManager, staker, removeStrategy, removeAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, removeStrategy);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        assertEq(sharesAfter, 0, "sharesAfter != 0");
        assertFalse(
            _isDepositedStrategy(staker, removeStrategy),
            "strategy should not be part of staker strategy list"
        );
    }

    /**
     * @notice testing removeShares() function with 3 strategies deposited.
     * Removing Shares could result in removing from staker strategy list if depositAmounts[i] == sharesAmounts[i].
     * Only callable by DelegationManager
     */
    function testFuzz_RemoveShares(uint256[3] memory depositAmounts, uint256[3] memory sharesAmounts) external {
        address staker = address(this);
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        uint256[] memory sharesBefore = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            depositAmounts[i] = bound(depositAmounts[i], 1, strategies[i].underlyingToken().totalSupply());
            sharesAmounts[i] = bound(sharesAmounts[i], 1, depositAmounts[i]);
            _depositIntoStrategySuccessfully(strategies[i], staker, depositAmounts[i]);
            sharesBefore[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            assertEq(sharesBefore[i], depositAmounts[i], "Staker has not deposited amount into strategy");
            assertTrue(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        uint256 numPoppedStrategies = 0;
        uint256[] memory sharesAfter = new uint256[](3);
        for (uint256 i = 0; i < 3; ++i) {
            delegationManagerMock.removeShares(strategyManager, staker, strategies[i], sharesAmounts[i]);
        }

        for (uint256 i = 0; i < 3; ++i) {
            sharesAfter[i] = strategyManager.stakerStrategyShares(staker, strategies[i]);
            if (sharesAmounts[i] == depositAmounts[i]) {
                ++numPoppedStrategies;
                assertFalse(
                    _isDepositedStrategy(staker, strategies[i]),
                    "strategy should not be part of staker strategy list"
                );
                assertEq(sharesAfter[i], 0, "sharesAfter != 0");
            } else {
                assertTrue(
                    _isDepositedStrategy(staker, strategies[i]),
                    "strategy should be part of staker strategy list"
                );
                assertEq(
                    sharesAfter[i],
                    sharesBefore[i] - sharesAmounts[i],
                    "sharesAfter != sharesBefore - sharesAmounts"
                );
            }
        }
        assertEq(
            stakerStrategyListLengthBefore - numPoppedStrategies,
            strategyManager.stakerStrategyListLength(staker),
            "stakerStrategyListLengthBefore - numPoppedStrategies != strategyManager.stakerStrategyListLength(staker)"
        );
    }
}

contract StrategyManagerUnitTests_addShares is StrategyManagerUnitTests {
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert("StrategyManager.onlyDelegationManager: not the DelegationManager");
        invalidDelegationManager.addShares(strategyManager, address(this), dummyToken, dummyStrat, 1);
    }

    function testFuzz_Revert_StakerZeroAddress(uint256 amount) external {
        cheats.expectRevert("StrategyManager._addShares: staker cannot be zero address");
        delegationManagerMock.addShares(strategyManager, address(0), dummyToken, dummyStrat, amount);
    }

    function testFuzz_Revert_ZeroShares(address staker) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.expectRevert("StrategyManager._addShares: shares should not be zero!");
        delegationManagerMock.addShares(strategyManager, staker, dummyToken, dummyStrat, 0);
    }

    function testFuzz_AppendsStakerStrategyList(
        address staker,
        uint256 amount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0) && amount != 0);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        assertEq(sharesBefore, 0, "Staker has already deposited into this strategy");
        assertFalse(_isDepositedStrategy(staker, dummyStrat), "strategy should not be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyToken, dummyStrat, amount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore + 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
        );
        assertEq(sharesAfter, amount, "sharesAfter != amount");
        assertTrue(_isDepositedStrategy(staker, dummyStrat), "strategy should be deposited");
    }

    function testFuzz_AddSharesToExistingShares(
        address staker,
        uint256 sharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0) && 0 < sharesAmount && sharesAmount <= dummyToken.totalSupply());
        uint256 initialAmount = 1e18;
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, initialAmount);
        uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, dummyStrat);
        assertEq(sharesBefore, initialAmount, "Staker has not deposited amount into strategy");
        assertTrue(_isDepositedStrategy(staker, strategy), "strategy should be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyToken, dummyStrat, sharesAmount);
        uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, dummyStrat);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore"
        );
        assertEq(sharesAfter, sharesBefore + sharesAmount, "sharesAfter != sharesBefore + sharesAmount");
        assertTrue(_isDepositedStrategy(staker, strategy), "strategy should be deposited");
    }

    /**
     * @notice When _addShares() called either by depositIntoStrategy or addShares() results in appending to
     * stakerStrategyListLength when the staker has MAX_STAKER_STRATEGY_LIST_LENGTH strategies, it should revert
     */
    function test_Revert_WhenMaxStrategyListLength() external {
        address staker = address(this);
        IERC20 token = dummyToken;
        uint256 amount = 1e18;
        IStrategy strategy = dummyStrat;
        uint256 MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

        // loop that deploys a new strategy and deposits into it
        for (uint256 i = 0; i < MAX_STAKER_STRATEGY_LIST_LENGTH; ++i) {
            cheats.startPrank(staker);
            strategyManager.depositIntoStrategy(strategy, token, amount);
            cheats.stopPrank();

            dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategy = dummyStrat;

            // whitelist the strategy for deposit
            cheats.startPrank(strategyManager.owner());
            IStrategy[] memory _strategy = new IStrategy[](1);
            bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);

            _strategy[0] = dummyStrat;
            strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
            cheats.stopPrank();
        }

        assertEq(
            strategyManager.stakerStrategyListLength(staker),
            MAX_STAKER_STRATEGY_LIST_LENGTH,
            "strategyManager.stakerStrategyListLength(staker) != MAX_STAKER_STRATEGY_LIST_LENGTH"
        );

        cheats.prank(staker);
        cheats.expectRevert("StrategyManager._addShares: deposit would exceed MAX_STAKER_STRATEGY_LIST_LENGTH");
        delegationManagerMock.addShares(strategyManager, staker, dummyToken, strategy, amount);

        cheats.expectRevert("StrategyManager._addShares: deposit would exceed MAX_STAKER_STRATEGY_LIST_LENGTH");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }
}

contract StrategyManagerUnitTests_withdrawSharesAsTokens is StrategyManagerUnitTests {
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert("StrategyManager.onlyDelegationManager: not the DelegationManager");
        invalidDelegationManager.removeShares(strategyManager, address(this), dummyStrat, 1);
    }

    /**
     * @notice deposits a single strategy and withdrawSharesAsTokens() function reverts when sharesAmount is
     * higher than depositAmount
     */
    function testFuzz_Revert_ShareAmountTooHigh(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply() && depositAmount < sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert("StrategyBase.withdraw: amountShares must be less than or equal to totalShares");
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
    }

    function testFuzz_SingleStrategyDeposited(
        address staker,
        uint256 depositAmount,
        uint256 sharesAmount
    ) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply() && depositAmount >= sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint256 balanceBefore = token.balanceOf(staker);
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
        uint256 balanceAfter = token.balanceOf(staker);
        assertEq(balanceAfter, balanceBefore + sharesAmount, "balanceAfter != balanceBefore + sharesAmount");
    }
}

contract StrategyManagerUnitTests_setStrategyWhitelister is StrategyManagerUnitTests {
    function testFuzz_SetStrategyWhitelister(
        address newWhitelister
    ) external filterFuzzedAddressInputs(newWhitelister) {
        address previousStrategyWhitelister = strategyManager.strategyWhitelister();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyWhitelisterChanged(previousStrategyWhitelister, newWhitelister);
        strategyManager.setStrategyWhitelister(newWhitelister);
        assertEq(
            strategyManager.strategyWhitelister(),
            newWhitelister,
            "strategyManager.strategyWhitelister() != newWhitelister"
        );
    }

    function testFuzz_Revert_WhenCalledByNotOwner(address notOwner) external filterFuzzedAddressInputs(notOwner) {
        cheats.assume(notOwner != strategyManager.owner());
        address newWhitelister = address(this);
        cheats.prank(notOwner);
        cheats.expectRevert("Ownable: caller is not the owner");
        strategyManager.setStrategyWhitelister(newWhitelister);
    }
}

contract StrategyManagerUnitTests_addStrategiesToDepositWhitelist is StrategyManagerUnitTests {
    function testFuzz_Revert_WhenCalledByNotStrategyWhitelister(
        address notStrategyWhitelister
    ) external filterFuzzedAddressInputs(notStrategyWhitelister) {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = new IStrategy[](1);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        IStrategy _strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = _strategy;

        cheats.prank(notStrategyWhitelister);
        cheats.expectRevert("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister");
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);
    }

    function test_AddSingleStrategyToWhitelist() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");
    }

    function test_AddAlreadyWhitelistedStrategy() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit UpdatedThirdPartyTransfersForbidden(strategy, false);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");
        // Make sure event not emitted by checking logs length
        cheats.recordLogs();
        uint256 numLogsBefore = cheats.getRecordedLogs().length;
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);
        uint256 numLogsAfter = cheats.getRecordedLogs().length;
        assertEq(numLogsBefore, numLogsAfter, "event emitted when strategy already whitelisted");
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should still be whitelisted");
    }

    function testFuzz_AddStrategiesToDepositWhitelist(uint8 numberOfStrategiesToAdd) external {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);
        _addStrategiesToWhitelist(numberOfStrategiesToAdd);
    }
}

contract StrategyManagerUnitTests_removeStrategiesFromDepositWhitelist is StrategyManagerUnitTests {
    function testFuzz_Revert_WhenCalledByNotStrategyWhitelister(
        address notStrategyWhitelister
    ) external filterFuzzedAddressInputs(notStrategyWhitelister) {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = _addStrategiesToWhitelist(1);

        cheats.prank(notStrategyWhitelister);
        cheats.expectRevert("StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister");
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
    }

    /**
     * @notice testing that mapping is still false and no event emitted
     */
    function test_RemoveNonWhitelistedStrategy() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        // Make sure event not emitted by checking logs length
        cheats.recordLogs();
        uint256 numLogsBefore = cheats.getRecordedLogs().length;
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        uint256 numLogsAfter = cheats.getRecordedLogs().length;
        assertEq(numLogsBefore, numLogsAfter, "event emitted when strategy already not whitelisted");
        assertFalse(
            strategyManager.strategyIsWhitelistedForDeposit(strategy),
            "strategy still should not be whitelisted"
        );
    }

    /**
     * @notice testing that strategy is removed from whitelist and event is emitted
     */
    function test_RemoveWhitelistedStrategy() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        // Add strategy to whitelist first
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray, thirdPartyTransfersForbiddenValues);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");

        // Now remove strategy from whitelist
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyRemovedFromDepositWhitelist(strategy);
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        assertFalse(
            strategyManager.strategyIsWhitelistedForDeposit(strategy),
            "strategy should no longer be whitelisted"
        );
    }

    function testFuzz_RemoveStrategiesFromDepositWhitelist(
        uint8 numberOfStrategiesToAdd,
        uint8 numberOfStrategiesToRemove
    ) external {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);
        cheats.assume(numberOfStrategiesToRemove <= 16);
        cheats.assume(numberOfStrategiesToRemove <= numberOfStrategiesToAdd);

        IStrategy[] memory strategiesAdded = _addStrategiesToWhitelist(numberOfStrategiesToAdd);

        IStrategy[] memory strategiesToRemove = new IStrategy[](numberOfStrategiesToRemove);
        // loop that selectively copies from array to other array
        for (uint256 i = 0; i < numberOfStrategiesToRemove; ++i) {
            strategiesToRemove[i] = strategiesAdded[i];
        }

        cheats.prank(strategyManager.strategyWhitelister());
        for (uint256 i = 0; i < strategiesToRemove.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyRemovedFromDepositWhitelist(strategiesToRemove[i]);
        }
        strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemove);

        for (uint256 i = 0; i < numberOfStrategiesToAdd; ++i) {
            if (i < numberOfStrategiesToRemove) {
                assertFalse(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesToRemove[i]),
                    "strategy not properly removed from whitelist"
                );
            } else {
                assertTrue(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesAdded[i]),
                    "strategy improperly removed from whitelist?"
                );
            }
        }
    }
}
