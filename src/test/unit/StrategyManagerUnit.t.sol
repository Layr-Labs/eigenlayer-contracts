// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "src/test/mocks/ERC20Mock.sol";
import "src/test/mocks/ERC20_SetTransferReverting_Mock.sol";
import "src/test/mocks/Reverter.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/MockDecimals.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the StrategyManager contract, entire withdrawal tests related to the
 * DelegationManager are not tested here but callable functions by the DelegationManager are mocked and tested here.
 * Contracts tested: StrategyManager.sol
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract StrategyManagerUnitTests is EigenLayerUnitTestSetup, IStrategyManagerEvents {
    address constant DEFAULT_BURN_ADDRESS = address(0x00000000000000000000000000000000000E16E4);

    OperatorSet defaultOperatorSet = OperatorSet(address(this), 0);
    uint defaultSlashId = 1;

    StrategyManager public strategyManagerImplementation;
    StrategyManager public strategyManager;

    IERC20 public dummyToken;
    ERC20_SetTransferReverting_Mock public revertToken;
    StrategyBase public dummyStrat;
    StrategyBase public dummyStrat2;
    StrategyBase public dummyStrat3;

    Reenterer public reenterer;

    address initialOwner = address(this);
    uint public privateKey = 111_111;
    address constant dummyAdmin = address(uint160(uint(keccak256("DummyAdmin"))));
    uint constant MAX_STRATEGY_TOTAL_SHARES = 1e38 - 1;

    function setUp() public override {
        EigenLayerUnitTestSetup.setUp();
        strategyManagerImplementation = new StrategyManager(
            IDelegationManager(address(delegationManagerMock)),
            ISlashEscrowFactory(address(slashEscrowFactoryMock)),
            pauserRegistry,
            "9.9.9"
        );
        strategyManager = StrategyManager(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyManager.initialize.selector, initialOwner, initialOwner, 0 /*initialPausedStatus*/ )
                )
            )
        );
        dummyToken = new ERC20PresetFixedSupply("mock token", "MOCK", MAX_STRATEGY_TOTAL_SHARES, address(this));
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
        for (uint i = 0; i < _strategies.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategies[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategies);

        isExcludedFuzzAddress[address(reenterer)] = true;
    }

    // INTERNAL / HELPER FUNCTIONS
    function _deployNewStrategy(IERC20 _token, IStrategyManager _strategyManager, IPauserRegistry _pauserRegistry, address admin)
        public
        returns (StrategyBase)
    {
        StrategyBase newStrategyImplementation = new StrategyBase(_strategyManager, _pauserRegistry, "9.9.9");
        StrategyBase newStrategy =
            StrategyBase(address(new TransparentUpgradeableProxy(address(newStrategyImplementation), address(admin), "")));
        newStrategy.initialize(_token);
        return newStrategy;
    }

    function _depositIntoStrategySuccessfully(IStrategy strategy, address staker, uint amount) internal filterFuzzedAddressInputs(staker) {
        IERC20 token = dummyToken;

        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // filter out zero address because the mock ERC20 we are using will revert on using it
        cheats.assume(staker != address(0));
        // sanity check / filter
        cheats.assume(amount <= token.balanceOf(address(this)));

        token.transfer(staker, amount);

        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        // needed for expecting an event with the right parameters
        uint expectedDepositShares = amount;

        cheats.startPrank(staker);
        token.approve(address(strategyManager), amount);

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, strategy, expectedDepositShares);
        uint shares = strategyManager.depositIntoStrategy(strategy, token, amount);

        cheats.stopPrank();

        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

        assertEq(depositSharesAfter, depositSharesBefore + shares, "depositSharesAfter != depositSharesBefore + shares");
        if (depositSharesBefore == 0) {
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
    function _depositIntoStrategyWithSignature(address staker, uint amount, uint expiry, bytes4 expectedRevertMessage)
        internal
        returns (bytes memory)
    {
        // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
        cheats.assume(amount != 0);
        // sanity check / filter
        cheats.assume(amount <= dummyToken.balanceOf(address(this)));

        dummyToken.approve(address(strategyManager), amount);

        uint nonceBefore = strategyManager.nonces(staker);
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, dummyStrat, dummyToken, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, dummyStrat);

        bool expectedRevertMessageIsempty = expectedRevertMessage == bytes4(0x00000000);
        if (!expectedRevertMessageIsempty) {
            cheats.expectRevert(expectedRevertMessage);
        } else if (expiry < block.timestamp) {
            cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        } else {
            // needed for expecting an event with the right parameters
            uint expectedDepositShares = amount;
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit Deposit(staker, dummyStrat, expectedDepositShares);
        }
        uint shares = strategyManager.depositIntoStrategyWithSignature(dummyStrat, dummyToken, amount, staker, expiry, signature);

        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, dummyStrat);
        uint nonceAfter = strategyManager.nonces(staker);

        if (expiry >= block.timestamp && expectedRevertMessageIsempty) {
            assertEq(depositSharesAfter, depositSharesBefore + shares, "depositSharesAfter != depositSharesBefore + shares");
            assertEq(nonceAfter, nonceBefore + 1, "nonceAfter != nonceBefore + 1");
        }
        return signature;
    }

    /**
     * @notice internal function to help check if a strategy is part of list of deposited strategies for a staker
     * Used to check if removed correctly after withdrawing all shares for a given strategy
     */
    function _isDepositedStrategy(address staker, IStrategy strategy) internal view returns (bool) {
        uint stakerStrategyListLength = strategyManager.stakerStrategyListLength(staker);
        for (uint i = 0; i < stakerStrategyListLength; ++i) {
            if (strategyManager.stakerStrategyList(staker, i) == strategy) return true;
        }
        return false;
    }

    /**
     * @notice Deploys numberOfStrategiesToAdd new strategies and adds them to the whitelist
     */
    function _addStrategiesToWhitelist(uint8 numberOfStrategiesToAdd) internal returns (IStrategy[] memory) {
        IStrategy[] memory strategyArray = new IStrategy[](numberOfStrategiesToAdd);
        // loop that deploys a new strategy and adds it to the array
        for (uint i = 0; i < numberOfStrategiesToAdd; ++i) {
            IStrategy _strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategyArray[i] = _strategy;
            assertFalse(strategyManager.strategyIsWhitelistedForDeposit(_strategy), "strategy improperly whitelisted?");
        }

        cheats.prank(strategyManager.strategyWhitelister());
        for (uint i = 0; i < numberOfStrategiesToAdd; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(strategyArray[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);

        for (uint i = 0; i < numberOfStrategiesToAdd; ++i) {
            assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategyArray[i]), "strategy not whitelisted");
        }

        return strategyArray;
    }
}

contract StrategyManagerUnitTests_initialize is StrategyManagerUnitTests {
    function test_CannotReinitialize() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        strategyManager.initialize(initialOwner, initialOwner, 0);
    }

    function test_InitializedStorageProperly() public view {
        assertTrue(strategyManager.domainSeparator() != bytes32(0), "sanity check");
        assertEq(strategyManager.owner(), initialOwner, "strategyManager.owner() != initialOwner");
        assertEq(strategyManager.strategyWhitelister(), initialOwner, "strategyManager.strategyWhitelister() != initialOwner");
        assertEq(address(strategyManager.pauserRegistry()), address(pauserRegistry), "strategyManager.pauserRegistry() != pauserRegistry");

        bytes memory v = bytes(strategyManager.version());

        bytes32 expectedDomainSeparator = keccak256(
            abi.encode(
                EIP712_DOMAIN_TYPEHASH,
                keccak256(bytes("EigenLayer")),
                keccak256(abi.encodePacked(v[0])),
                block.chainid,
                address(strategyManager)
            )
        );

        assertEq(strategyManager.domainSeparator(), expectedDomainSeparator, "sanity check");
    }
}

contract StrategyManagerUnitTests_depositIntoStrategy is StrategyManagerUnitTests {
    function testFuzz_depositIntoStrategySuccessfully(address staker, uint amount) public filterFuzzedAddressInputs(staker) {
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

        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        // needed for expecting an event with the right parameters
        uint expectedDepositShares = strategy.underlyingToShares(amount);

        uint strategyBalanceBefore = token.balanceOf(address(strategy));

        token.transfer(staker, amount);
        cheats.startPrank(staker);
        token.approve(address(strategyManager), amount);

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit Deposit(staker, strategy, expectedDepositShares);
        uint depositedShares = strategyManager.depositIntoStrategy(strategy, token, amount);

        cheats.stopPrank();
        uint strategyBalanceAfter = token.balanceOf(address(strategy));

        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        assertEq(strategyBalanceBefore + amount, strategyBalanceAfter, "balance of strategy not increased by deposit amount");
        assertEq(depositSharesAfter, depositSharesBefore + depositedShares, "depositSharesAfter != depositSharesBefore + depositedShares");
        if (depositSharesBefore == 0) {
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
        uint amount = 1e18;
        testFuzz_depositIntoStrategySuccessfully(staker, amount);
        testFuzz_depositIntoStrategySuccessfully(staker, amount);
    }

    function test_Revert_WhenDepositsPaused() public {
        uint amount = 1e18;

        // pause deposits
        cheats.prank(pauser);
        strategyManager.pause(1);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        strategyManager.depositIntoStrategy(dummyStrat, dummyToken, amount);
    }

    function test_Revert_WhenReentering() public {
        uint amount = 1e18;

        reenterer = new Reenterer();
        dummyToken.approve(address(strategyManager), MAX_STRATEGY_TOTAL_SHARES);

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = IStrategy(address(reenterer));
        for (uint i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        reenterer.prepareReturnData(abi.encode(amount));

        address targetToUse = address(strategyManager);
        uint msgValueToUse = 0;
        bytes memory calldataToUse =
            abi.encodeWithSelector(StrategyManager.depositIntoStrategy.selector, address(reenterer), dummyToken, amount);
        reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

        strategyManager.depositIntoStrategy(IStrategy(address(reenterer)), dummyToken, amount);
    }

    function test_Revert_WhenTokenSafeTransferFromReverts() external {
        // replace 'dummyStrat' with one that uses a reverting token
        dummyToken = IERC20(address(new ReverterWithDecimals()));
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = address(this);
        IERC20 token = dummyToken;
        uint amount = 1e18;
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
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = address(this);
        IERC20 token = dummyToken;
        uint amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("SafeERC20: low-level call failed");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenStrategyDepositFunctionReverts() external {
        // replace 'dummyStrat' with one that always reverts
        dummyStrat = StrategyBase(address(new Reverter()));

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = address(this);
        dummyToken.approve(address(strategyManager), MAX_STRATEGY_TOTAL_SHARES);
        IERC20 token = dummyToken;
        uint amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert("Reverter: I am a contract that always reverts");
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_Revert_WhenStrategyDoesNotExist() external {
        // replace 'dummyStrat' with one that does not exist
        dummyStrat = StrategyBase(address(5678));

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = address(this);
        IERC20 token = dummyToken;
        uint amount = 1e18;
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
        uint amount = 1e18;
        IStrategy strategy = dummyStrat;

        cheats.prank(staker);
        cheats.expectRevert(IStrategyManagerErrors.StrategyNotWhitelisted.selector);
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }

    function test_addShares_Revert_WhenSharesIsZero() external {
        // replace dummyStrat with Reenterer contract
        reenterer = new Reenterer();
        dummyStrat = StrategyBase(address(reenterer));

        // whitelist the strategy for deposit
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = dummyStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = address(this);
        dummyToken.approve(address(strategyManager), MAX_STRATEGY_TOTAL_SHARES);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint amount = 1e18;

        reenterer.prepareReturnData(abi.encode(uint(0)));

        cheats.prank(staker);
        cheats.expectRevert(IStrategyManagerErrors.SharesAmountZero.selector);
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }
}

contract StrategyManagerUnitTests_depositIntoStrategyWithSignature is StrategyManagerUnitTests {
    function test_Revert_WhenSignatureInvalid() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint amount = 1e18;

        uint nonceBefore = strategyManager.nonces(staker);
        uint expiry = block.timestamp;
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);

        cheats.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        // call with `notStaker` as input instead of `staker` address
        address notStaker = address(3333);
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, notStaker, expiry, signature);

        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        uint nonceAfter = strategyManager.nonces(staker);

        assertEq(depositSharesAfter, depositSharesBefore, "depositSharesAfter != depositSharesBefore");
        assertEq(nonceAfter, nonceBefore, "nonceAfter != nonceBefore");
    }

    function testFuzz_DepositSuccessfully(uint amount, uint expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        _depositIntoStrategyWithSignature(staker, amount, expiry, bytes4(0x00000000));
    }

    function testFuzz_Revert_SignatureReplay(uint amount, uint expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);
        cheats.assume(expiry > block.timestamp);

        address staker = cheats.addr(privateKey);
        // not expecting a revert, so input an empty string
        bytes memory signature = _depositIntoStrategyWithSignature(staker, amount, expiry, "");

        cheats.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        strategyManager.depositIntoStrategyWithSignature(dummyStrat, dummyToken, amount, staker, expiry, signature);
    }

    // tries depositing using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature
    function testFuzz_Revert_WithContractWallet_BadSignature(uint amount) public {
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

        uint nonceBefore = strategyManager.nonces(staker);
        uint expiry = type(uint).max;
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    // tries depositing using a wallet that does not comply with EIP 1271
    function testFuzz_Revert_WithContractWallet_NonconformingWallet(uint amount, uint8 v, bytes32 r, bytes32 s) public {
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

        uint expiry = type(uint).max;
        bytes memory signature = abi.encodePacked(r, s, v);

        cheats.expectRevert();
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    // Tries depositing without token approval and transfer fails. deposit function should also revert
    function test_Revert_WithContractWallet_TokenTransferFails() external {
        address staker = cheats.addr(privateKey);
        uint amount = 1e18;
        uint nonceBefore = strategyManager.nonces(staker);
        uint expiry = block.timestamp + 100;
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, dummyStrat, revertToken, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert("ERC20: insufficient allowance");
        strategyManager.depositIntoStrategyWithSignature(dummyStrat, revertToken, amount, staker, expiry, signature);
    }

    // tries depositing using a signature and an EIP 1271 compliant wallet
    function testFuzz_WithContractWallet_Successfully(uint amount, uint expiry) public {
        // min shares must be minted on strategy
        cheats.assume(amount >= 1);

        address staker = cheats.addr(privateKey);

        // deploy ERC1271WalletMock for staker to use
        cheats.prank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        staker = address(wallet);

        // not expecting a revert, so input an empty string
        _depositIntoStrategyWithSignature(staker, amount, expiry, bytes4(0x00000000));
    }

    function test_Revert_WhenDepositsPaused() public {
        address staker = cheats.addr(privateKey);

        // pause deposits
        cheats.prank(pauser);
        strategyManager.pause(1);

        _depositIntoStrategyWithSignature(staker, 1e18, type(uint).max, IPausable.CurrentlyPaused.selector);
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
        cheats.prank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);

        dummyToken.approve(address(strategyManager), MAX_STRATEGY_TOTAL_SHARES);

        _strategy[0] = IStrategy(address(reenterer));
        for (uint i = 0; i < _strategy.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyAddedToDepositWhitelist(_strategy[i]);
        }
        strategyManager.addStrategiesToDepositWhitelist(_strategy);

        address staker = cheats.addr(privateKey);
        IStrategy strategy = IStrategy(address(reenterer));
        IERC20 token = dummyToken;
        uint amount = 1e18;

        uint nonceBefore = strategyManager.nonces(staker);
        uint expiry = type(uint).max;
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint shareAmountToReturn = amount;
        reenterer.prepareReturnData(abi.encode(shareAmountToReturn));

        {
            address targetToUse = address(strategyManager);
            uint msgValueToUse = 0;
            bytes memory calldataToUse =
                abi.encodeWithSelector(StrategyManager.depositIntoStrategy.selector, address(reenterer), dummyToken, amount);
            reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));
        }
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);
    }

    function test_Revert_WhenSignatureExpired() public {
        address staker = cheats.addr(privateKey);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        uint amount = 1e18;

        uint nonceBefore = strategyManager.nonces(staker);
        uint expiry = 5555;
        // warp to 1 second after expiry
        cheats.warp(expiry + 1);
        bytes memory signature;

        {
            bytes32 structHash =
                keccak256(abi.encode(strategyManager.DEPOSIT_TYPEHASH(), staker, strategy, token, amount, nonceBefore, expiry));
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", strategyManager.domainSeparator(), structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);

        cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        strategyManager.depositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature);

        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        uint nonceAfter = strategyManager.nonces(staker);

        assertEq(depositSharesAfter, depositSharesBefore, "depositSharesAfter != depositSharesBefore");
        assertEq(nonceAfter, nonceBefore, "nonceAfter != nonceBefore");
    }

    function test_Revert_WhenStrategyNotWhitelisted() external {
        // replace 'dummyStrat' with one that is not whitelisted
        dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        dummyToken = dummyStrat.underlyingToken();
        address staker = cheats.addr(privateKey);
        uint amount = 1e18;

        _depositIntoStrategyWithSignature(staker, amount, type(uint).max, IStrategyManagerErrors.StrategyNotWhitelisted.selector);
    }
}

contract StrategyManagerUnitTests_removeDepositShares is StrategyManagerUnitTests {
    /**
     * @notice Should revert if not called by DelegationManager
     */
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(IStrategyManagerErrors.OnlyDelegationManager.selector);
        invalidDelegationManager.removeDepositShares(strategyManager, address(this), dummyStrat, 1);
    }

    /**
     * @notice deposits a single strategy and tests removeDepositShares() function reverts when sharesAmount is 0
     */
    function testFuzz_Revert_ZeroShares(address staker, uint depositAmount) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(IStrategyManagerErrors.SharesAmountZero.selector);
        delegationManagerMock.removeDepositShares(strategyManager, staker, strategy, 0);
    }

    /**
     * @notice deposits a single strategy and tests removeDepositShares() function reverts when sharesAmount is
     * higher than depositAmount
     */
    function testFuzz_Revert_ShareAmountTooHigh(address staker, uint depositAmount, uint removeSharesAmount)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(IStrategyManagerErrors.SharesAmountTooHigh.selector);
        delegationManagerMock.removeDepositShares(strategyManager, staker, strategy, removeSharesAmount);
    }

    /**
     * @notice deposit single strategy and removeDepositShares() for less than the deposited amount
     * Shares should be updated correctly with stakerStrategyListLength unchanged
     */
    function testFuzz_RemoveSharesLessThanDeposit(address staker, uint depositAmount, uint removeSharesAmount)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(0));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply());
        cheats.assume(removeSharesAmount > 0 && removeSharesAmount < depositAmount);
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);
        delegationManagerMock.removeDepositShares(strategyManager, staker, strategy, removeSharesAmount);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        assertEq(depositSharesBefore, depositSharesAfter + removeSharesAmount, "Remove incorrect amount of shares");
        assertEq(stakerStrategyListLengthBefore, stakerStrategyListLengthAfter, "stakerStrategyListLength shouldn't have changed");
    }

    /**
     * @notice testing removeDepositShares()
     * deposits 1 strategy and tests it is removed from staker strategy list after removing all shares
     */
    function testFuzz_RemovesStakerStrategyListSingleStrat(address staker, uint sharesAmount) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply());
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, sharesAmount);

        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, strategy);
        assertEq(depositSharesBefore, sharesAmount, "Staker has not deposited amount into strategy");

        delegationManagerMock.removeDepositShares(strategyManager, staker, strategy, sharesAmount);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, strategy);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        assertEq(depositSharesAfter, 0, "depositSharesAfter != 0");
        assertFalse(_isDepositedStrategy(staker, strategy), "strategy should not be part of staker strategy list");
    }

    /**
     * @notice testing removeDepositShares() function with 3 strategies deposited.
     * Randomly selects one of the 3 strategies to be fully removed from staker strategy list.
     * Only callable by DelegationManager
     */
    function testFuzz_RemovesStakerStrategyListMultipleStrat(address staker, uint[3] memory amounts, uint8 randStrategy)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(0));
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        for (uint i = 0; i < 3; ++i) {
            amounts[i] = bound(amounts[i], 1, dummyToken.totalSupply() - 1);
            _depositIntoStrategySuccessfully(strategies[i], staker, amounts[i]);
        }
        IStrategy removeStrategy = strategies[randStrategy % 3];
        uint removeAmount = amounts[randStrategy % 3];

        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint[] memory depositSharesBefore = new uint[](3);
        for (uint i = 0; i < 3; ++i) {
            depositSharesBefore[i] = strategyManager.stakerDepositShares(staker, strategies[i]);
            assertEq(depositSharesBefore[i], amounts[i], "Staker has not deposited amount into strategy");
            assertTrue(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }

        delegationManagerMock.removeDepositShares(strategyManager, staker, removeStrategy, removeAmount);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, removeStrategy);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore - 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore - 1"
        );
        assertEq(depositSharesAfter, 0, "depositSharesAfter != 0");
        assertFalse(_isDepositedStrategy(staker, removeStrategy), "strategy should not be part of staker strategy list");
    }

    /**
     * @notice testing removeDepositShares() function with 3 strategies deposited.
     * Removing Shares could result in removing from staker strategy list if depositAmounts[i] == sharesAmounts[i].
     * Only callable by DelegationManager
     */
    function testFuzz_RemoveShares(uint[3] memory depositAmounts, uint[3] memory sharesAmounts) external {
        address staker = address(this);
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        uint[] memory depositSharesBefore = new uint[](3);
        for (uint i = 0; i < 3; ++i) {
            depositAmounts[i] = bound(depositAmounts[i], 1, strategies[i].underlyingToken().totalSupply());
            sharesAmounts[i] = bound(sharesAmounts[i], 1, depositAmounts[i]);
            _depositIntoStrategySuccessfully(strategies[i], staker, depositAmounts[i]);
            depositSharesBefore[i] = strategyManager.stakerDepositShares(staker, strategies[i]);
            assertEq(depositSharesBefore[i], depositAmounts[i], "Staker has not deposited amount into strategy");
            assertTrue(_isDepositedStrategy(staker, strategies[i]), "strategy should be deposited");
        }
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

        uint numPoppedStrategies = 0;
        uint[] memory depositSharesAfter = new uint[](3);
        for (uint i = 0; i < 3; ++i) {
            delegationManagerMock.removeDepositShares(strategyManager, staker, strategies[i], sharesAmounts[i]);
        }

        for (uint i = 0; i < 3; ++i) {
            depositSharesAfter[i] = strategyManager.stakerDepositShares(staker, strategies[i]);
            if (sharesAmounts[i] == depositAmounts[i]) {
                ++numPoppedStrategies;
                assertFalse(_isDepositedStrategy(staker, strategies[i]), "strategy should not be part of staker strategy list");
                assertEq(depositSharesAfter[i], 0, "depositSharesAfter != 0");
            } else {
                assertTrue(_isDepositedStrategy(staker, strategies[i]), "strategy should be part of staker strategy list");
                assertEq(
                    depositSharesAfter[i],
                    depositSharesBefore[i] - sharesAmounts[i],
                    "depositSharesAfter != depositSharesBefore - sharesAmounts"
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
        cheats.expectRevert(IStrategyManagerErrors.OnlyDelegationManager.selector);
        invalidDelegationManager.addShares(strategyManager, address(this), dummyStrat, 1);
    }

    function testFuzz_Revert_StakerZeroAddress(uint amount) external {
        cheats.expectRevert(IStrategyManagerErrors.StakerAddressZero.selector);
        delegationManagerMock.addShares(strategyManager, address(0), dummyStrat, amount);
    }

    function testFuzz_Revert_ZeroShares(address staker) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0));
        cheats.expectRevert(IStrategyManagerErrors.SharesAmountZero.selector);
        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, 0);
    }

    function testFuzz_AppendsStakerStrategyList(address staker, uint amount) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0) && amount != 0);
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, dummyStrat);
        assertEq(depositSharesBefore, 0, "Staker has already deposited into this strategy");
        assertFalse(_isDepositedStrategy(staker, dummyStrat), "strategy should not be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, amount);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, dummyStrat);
        assertEq(
            stakerStrategyListLengthAfter,
            stakerStrategyListLengthBefore + 1,
            "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
        );
        assertEq(depositSharesAfter, amount, "depositSharesAfter != amount");
        assertTrue(_isDepositedStrategy(staker, dummyStrat), "strategy should be deposited");
    }

    function testFuzz_AddSharesToExistingShares(address staker, uint sharesAmount) external filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != address(0) && 0 < sharesAmount && sharesAmount <= dummyToken.totalSupply());
        uint initialAmount = 1e18;
        IStrategy strategy = dummyStrat;
        _depositIntoStrategySuccessfully(strategy, staker, initialAmount);
        uint stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesBefore = strategyManager.stakerDepositShares(staker, dummyStrat);
        assertEq(depositSharesBefore, initialAmount, "Staker has not deposited amount into strategy");
        assertTrue(_isDepositedStrategy(staker, strategy), "strategy should be deposited");

        delegationManagerMock.addShares(strategyManager, staker, dummyStrat, sharesAmount);
        uint stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);
        uint depositSharesAfter = strategyManager.stakerDepositShares(staker, dummyStrat);
        assertEq(
            stakerStrategyListLengthAfter, stakerStrategyListLengthBefore, "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore"
        );
        assertEq(depositSharesAfter, depositSharesBefore + sharesAmount, "depositSharesAfter != depositSharesBefore + sharesAmount");
        assertTrue(_isDepositedStrategy(staker, strategy), "strategy should be deposited");
    }

    /**
     * @notice When _addShares() called either by depositIntoStrategy or addShares() results in appending to
     * stakerStrategyListLength when the staker has MAX_STAKER_STRATEGY_LIST_LENGTH strategies, it should revert
     */
    function test_Revert_WhenMaxStrategyListLength() external {
        address staker = address(this);
        IERC20 token = dummyToken;
        uint amount = 1e18;
        IStrategy strategy = dummyStrat;
        uint MAX_STAKER_STRATEGY_LIST_LENGTH = 32;
        cheats.prank(staker);
        token.approve(address(strategyManager), MAX_STRATEGY_TOTAL_SHARES);

        // loop that deploys a new strategy and deposits into it
        for (uint i = 0; i < MAX_STAKER_STRATEGY_LIST_LENGTH; ++i) {
            cheats.prank(staker);
            strategyManager.depositIntoStrategy(strategy, token, amount);

            dummyStrat = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
            strategy = dummyStrat;

            // whitelist the strategy for deposit
            cheats.prank(strategyManager.owner());
            IStrategy[] memory _strategy = new IStrategy[](1);
            _strategy[0] = dummyStrat;
            strategyManager.addStrategiesToDepositWhitelist(_strategy);
        }

        assertEq(
            strategyManager.stakerStrategyListLength(staker),
            MAX_STAKER_STRATEGY_LIST_LENGTH,
            "strategyManager.stakerStrategyListLength(staker) != MAX_STAKER_STRATEGY_LIST_LENGTH"
        );

        cheats.prank(staker);
        cheats.expectRevert(IStrategyManagerErrors.MaxStrategiesExceeded.selector);
        delegationManagerMock.addShares(strategyManager, staker, strategy, amount);

        cheats.expectRevert(IStrategyManagerErrors.MaxStrategiesExceeded.selector);
        strategyManager.depositIntoStrategy(strategy, token, amount);
    }
}

contract StrategyManagerUnitTests_withdrawSharesAsTokens is StrategyManagerUnitTests {
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.expectRevert(IStrategyManagerErrors.OnlyDelegationManager.selector);
        invalidDelegationManager.removeDepositShares(strategyManager, address(this), dummyStrat, 1);
    }

    /**
     * @notice deposits a single strategy and withdrawSharesAsTokens() function reverts when sharesAmount is
     * higher than depositAmount
     */
    function testFuzz_Revert_ShareAmountTooHigh(address staker, uint depositAmount, uint sharesAmount)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(this));
        cheats.assume(staker != address(0));
        cheats.assume(staker != address(dummyStrat));
        cheats.assume(depositAmount > 0 && depositAmount < dummyToken.totalSupply() && depositAmount < sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        cheats.expectRevert(IStrategyErrors.WithdrawalAmountExceedsTotalDeposits.selector);
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
    }

    function testFuzz_SingleStrategyDeposited(address staker, uint depositAmount, uint sharesAmount)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(this));
        cheats.assume(staker != address(0));
        cheats.assume(staker != address(dummyStrat));
        cheats.assume(sharesAmount > 0 && sharesAmount < dummyToken.totalSupply() && depositAmount >= sharesAmount);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);
        uint balanceBefore = token.balanceOf(staker);
        delegationManagerMock.withdrawSharesAsTokens(strategyManager, staker, strategy, sharesAmount, token);
        uint balanceAfter = token.balanceOf(staker);
        assertEq(balanceAfter, balanceBefore + sharesAmount, "balanceAfter != balanceBefore + sharesAmount");
    }
}

contract StrategyManagerUnitTests_increaseBurnOrRedistributableShares is StrategyManagerUnitTests {
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.prank(address(invalidDelegationManager));
        cheats.expectRevert(IStrategyManagerErrors.OnlyDelegationManager.selector);
        strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, dummyStrat, 1);
    }

    // function testFuzz_increaseBurnOrRedistributableShares(uint addedSharesToBurn) external {
    //     IStrategy strategy = dummyStrat;
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit BurnOrRedistributableSharesIncreased(defaultOperatorSet, defaultSlashId, strategy, addedSharesToBurn);
    //     cheats.prank(address(delegationManagerMock));
    //     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategy, addedSharesToBurn);
    //     assertEq(
    //         strategyManager.getBurnableShares(strategy), addedSharesToBurn, "strategyManager.burnableShares(strategy) != addedSharesToBurn"
    //     );
    // }

    // function testFuzz_increaseBurnOrRedistributableShares_existingShares(uint existingBurnableShares, uint addedSharesToBurn) external {
    //     // preventing fuzz overflow, in practice StrategyBase has a 1e38 - 1 maxShares limit so this won't
    //     // be an issue on mainnet/testnet environments
    //     existingBurnableShares = bound(existingBurnableShares, 1, type(uint).max / 2);
    //     addedSharesToBurn = bound(addedSharesToBurn, 1, type(uint).max / 2);
    //     IStrategy strategy = dummyStrat;
    //     cheats.prank(address(delegationManagerMock));
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit BurnOrRedistributableSharesIncreased(defaultOperatorSet, defaultSlashId, strategy, existingBurnableShares);
    //     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategy, existingBurnableShares);
    //     assertEq(
    //         strategyManager.getBurnableShares(strategy),
    //         existingBurnableShares,
    //         "strategyManager.burnableShares(strategy) != existingBurnableShares"
    //     );

    //     cheats.prank(address(delegationManagerMock));
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit BurnOrRedistributableSharesIncreased(defaultOperatorSet, defaultSlashId, strategy, addedSharesToBurn);
    //     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategy, addedSharesToBurn);

    //     assertEq(
    //         strategyManager.getBurnableShares(strategy),
    //         existingBurnableShares + addedSharesToBurn,
    //         "strategyManager.burnableShares(strategy) != existingBurnableShares + addedSharesToBurn"
    //     );
    // }
}

contract StrategyManagerUnitTests_decreaseBurnOrRedistributableShares is StrategyManagerUnitTests {
// function testFuzz_SingleStrategyDeposited(address staker, uint depositAmount, uint sharesToBurn)
//     external
//     filterFuzzedAddressInputs(staker)
// {
//     cheats.assume(staker != address(0));
//     cheats.assume(staker != address(dummyStrat));
//     cheats.assume(sharesToBurn > 0 && sharesToBurn < dummyToken.totalSupply() && depositAmount >= sharesToBurn);
//     IStrategy strategy = dummyStrat;
//     IERC20 token = dummyToken;
//     _depositIntoStrategySuccessfully(strategy, staker, depositAmount);

//     // slash shares and increase amount to burn from DelegationManager
//     cheats.prank(address(delegationManagerMock));
//     cheats.expectEmit(true, true, true, true, address(strategyManager));
//     emit BurnOrRedistributableSharesIncreased(defaultOperatorSet, defaultSlashId, strategy, sharesToBurn);
//     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategy, sharesToBurn);

//     uint strategyBalanceBefore = token.balanceOf(address(strategy));
//     uint burnAddressBalanceBefore = token.balanceOf(strategyManager.DEFAULT_BURN_ADDRESS());
//     cheats.prank(address(delegationManagerMock));
//     // cheats.expectEmit(true, true, true, true, address(strategyManager));
//     // emit BurnOrRedistributableSharesDecreased(
//     //     OperatorSet(address(strategyManager), type(uint32).max), type(uint).max, strategy, DEFAULT_BURN_ADDRESS, sharesToBurn
//     // );
//     strategyManager.burnShares(strategy);
//     uint strategyBalanceAfter = token.balanceOf(address(strategy));
//     uint burnAddressBalanceAfter = token.balanceOf(strategyManager.DEFAULT_BURN_ADDRESS());

//     assertEq(strategyBalanceBefore - sharesToBurn, strategyBalanceAfter, "strategyBalanceBefore - sharesToBurn != strategyBalanceAfter");
//     assertEq(burnAddressBalanceAfter, burnAddressBalanceBefore + sharesToBurn, "balanceAfter != balanceBefore + sharesAmount");

//     // Verify strategy was removed from burnable sharesc
//     (address[] memory strategiesAfterBurn,) = strategyManager.getStrategiesWithBurnableShares();
//     assertEq(strategiesAfterBurn.length, 0, "Should have no strategies after burning");
//     assertEq(strategyManager.getBurnableShares(strategy), 0, "getBurnableShares should return 0 after burning");
// }

// /// @notice check that balances are unchanged with a reverting token but burnShares doesn't revert
// function testFuzz_BurnableSharesUnchangedWithRevertToken(address staker, uint depositAmount, uint sharesToBurn)
//     external
//     filterFuzzedAddressInputs(staker)
// {
//     cheats.assume(staker != address(0));
//     cheats.assume(sharesToBurn > 0 && sharesToBurn < dummyToken.totalSupply() && depositAmount >= sharesToBurn);
//     IStrategy strategy = dummyStrat;
//     IERC20 token = dummyToken;
//     _depositIntoStrategySuccessfully(strategy, staker, depositAmount);

//     // slash shares and increase amount to burn from DelegationManager
//     cheats.prank(address(delegationManagerMock));
//     cheats.expectEmit(true, true, true, true, address(strategyManager));
//     emit BurnOrRedistributableSharesIncreased(defaultOperatorSet, defaultSlashId, strategy, sharesToBurn);
//     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategy, sharesToBurn);

//     // Now set token to be contract that reverts simulating an upgrade
//     cheats.etch(address(token), address(revertToken).code);
//     ERC20_SetTransferReverting_Mock(address(token)).setTransfersRevert(true);

//     cheats.expectRevert("SafeERC20: low-level call failed");
//     cheats.prank(address(delegationManagerMock));
//     strategyManager.burnShares(strategy);

//     assertEq(strategyManager.getBurnableShares(strategy), sharesToBurn, "burnable shares should be unchanged");
// }
}

contract StrategyManagerUnitTests_setStrategyWhitelister is StrategyManagerUnitTests {
    function testFuzz_SetStrategyWhitelister(address newWhitelister) external filterFuzzedAddressInputs(newWhitelister) {
        address previousStrategyWhitelister = strategyManager.strategyWhitelister();
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyWhitelisterChanged(previousStrategyWhitelister, newWhitelister);
        strategyManager.setStrategyWhitelister(newWhitelister);
        assertEq(strategyManager.strategyWhitelister(), newWhitelister, "strategyManager.strategyWhitelister() != newWhitelister");
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
    function testFuzz_Revert_WhenCalledByNotStrategyWhitelister(address notStrategyWhitelister)
        external
        filterFuzzedAddressInputs(notStrategyWhitelister)
    {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = new IStrategy[](1);

        IStrategy _strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = _strategy;

        cheats.prank(notStrategyWhitelister);
        cheats.expectRevert(IStrategyManagerErrors.OnlyStrategyWhitelister.selector);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
    }

    function test_AddSingleStrategyToWhitelist() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);

        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");
    }

    function test_AddAlreadyWhitelistedStrategy() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);

        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");
        // Make sure event not emitted by checking logs length
        cheats.recordLogs();
        uint numLogsBefore = cheats.getRecordedLogs().length;
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        uint numLogsAfter = cheats.getRecordedLogs().length;
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
    function testFuzz_Revert_WhenCalledByNotStrategyWhitelister(address notStrategyWhitelister)
        external
        filterFuzzedAddressInputs(notStrategyWhitelister)
    {
        cheats.assume(notStrategyWhitelister != strategyManager.strategyWhitelister());
        IStrategy[] memory strategyArray = _addStrategiesToWhitelist(1);

        cheats.prank(notStrategyWhitelister);
        cheats.expectRevert(IStrategyManagerErrors.OnlyStrategyWhitelister.selector);
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
        uint numLogsBefore = cheats.getRecordedLogs().length;
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        uint numLogsAfter = cheats.getRecordedLogs().length;
        assertEq(numLogsBefore, numLogsAfter, "event emitted when strategy already not whitelisted");
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy still should not be whitelisted");
    }

    /**
     * @notice testing that strategy is removed from whitelist and event is emitted
     */
    function test_RemoveWhitelistedStrategy() external {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        IStrategy strategy = _deployNewStrategy(dummyToken, strategyManager, pauserRegistry, dummyAdmin);
        strategyArray[0] = strategy;

        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should not be whitelisted");
        // Add strategy to whitelist first
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyAddedToDepositWhitelist(strategy);
        strategyManager.addStrategiesToDepositWhitelist(strategyArray);
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should be whitelisted");

        // Now remove strategy from whitelist
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit StrategyRemovedFromDepositWhitelist(strategy);
        strategyManager.removeStrategiesFromDepositWhitelist(strategyArray);
        assertFalse(strategyManager.strategyIsWhitelistedForDeposit(strategy), "strategy should no longer be whitelisted");
    }

    function testFuzz_RemoveStrategiesFromDepositWhitelist(uint8 numberOfStrategiesToAdd, uint8 numberOfStrategiesToRemove) external {
        // sanity filtering on fuzzed input
        cheats.assume(numberOfStrategiesToAdd <= 16);
        cheats.assume(numberOfStrategiesToRemove <= 16);
        cheats.assume(numberOfStrategiesToRemove <= numberOfStrategiesToAdd);

        IStrategy[] memory strategiesAdded = _addStrategiesToWhitelist(numberOfStrategiesToAdd);

        IStrategy[] memory strategiesToRemove = new IStrategy[](numberOfStrategiesToRemove);
        // loop that selectively copies from array to other array
        for (uint i = 0; i < numberOfStrategiesToRemove; ++i) {
            strategiesToRemove[i] = strategiesAdded[i];
        }

        cheats.prank(strategyManager.strategyWhitelister());
        for (uint i = 0; i < strategiesToRemove.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(strategyManager));
            emit StrategyRemovedFromDepositWhitelist(strategiesToRemove[i]);
        }
        strategyManager.removeStrategiesFromDepositWhitelist(strategiesToRemove);

        for (uint i = 0; i < numberOfStrategiesToAdd; ++i) {
            if (i < numberOfStrategiesToRemove) {
                assertFalse(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesToRemove[i]), "strategy not properly removed from whitelist"
                );
            } else {
                assertTrue(
                    strategyManager.strategyIsWhitelistedForDeposit(strategiesAdded[i]), "strategy improperly removed from whitelist?"
                );
            }
        }
    }
}

contract StrategyManagerUnitTests_getStrategiesWithBurnableShares is StrategyManagerUnitTests {
    function test_getStrategiesWithBurnableShares_Empty() public view {
        (address[] memory strategies, uint[] memory shares) = strategyManager.getStrategiesWithBurnableShares();
        assertEq(strategies.length, 0, "Should have no strategies when empty");
        assertEq(shares.length, 0, "Should have no shares when empty");
    }

    // function testFuzz_getStrategiesWithBurnableShares_Single(uint sharesToAdd) public {
    //     //ensure non-zero
    //     cheats.assume(sharesToAdd > 0);

    //     // Add burnable shares
    //     cheats.prank(address(delegationManagerMock));
    //     strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, dummyStrat, sharesToAdd);

    //     // Get strategies with burnable shares
    //     (address[] memory strategies, uint[] memory shares) = strategyManager.getStrategiesWithBurnableShares();

    //     // Verify results
    //     assertEq(strategies.length, 1, "Should have one strategy");
    //     assertEq(shares.length, 1, "Should have one share amount");
    //     assertEq(strategies[0], address(dummyStrat), "Wrong strategy address");
    //     assertEq(shares[0], sharesToAdd, "Wrong shares amount");
    // }

    // function testFuzz_getStrategiesWithBurnableShares_Multiple(uint[3] calldata sharesToAdd) public {
    //     IStrategy[] memory strategies = new IStrategy[](3);
    //     strategies[0] = dummyStrat;
    //     strategies[1] = dummyStrat2;
    //     strategies[2] = dummyStrat3;
    //     uint[3] memory expectedShares;
    //     uint expectedLength = 0;

    //     // Add non-zero shares to strategies
    //     for (uint i = 0; i < 3; i++) {
    //         expectedShares[i] = sharesToAdd[i];
    //         if (sharesToAdd[i] > 0) {
    //             expectedLength++;
    //             cheats.prank(address(delegationManagerMock));
    //             strategyManager.increaseBurnOrRedistributableShares(defaultOperatorSet, defaultSlashId, strategies[i], sharesToAdd[i]);
    //         }
    //     }

    //     // Get strategies with burnable shares
    //     (address[] memory returnedStrategies, uint[] memory returnedShares) = strategyManager.getStrategiesWithBurnableShares();

    //     // Verify lengths match
    //     assertEq(returnedStrategies.length, expectedLength, "Wrong number of strategies returned");
    //     assertEq(returnedShares.length, expectedLength, "Wrong number of share amounts returned");

    //     // For all strategies with non-zero shares, verify they are in the returned arrays
    //     uint foundCount = 0;
    //     for (uint i = 0; i < 3; i++) {
    //         if (expectedShares[i] > 0) {
    //             bool found = false;
    //             for (uint j = 0; j < returnedStrategies.length; j++) {
    //                 if (returnedStrategies[j] == address(strategies[i])) {
    //                     assertEq(returnedShares[j], expectedShares[i], "Wrong share amount");
    //                     found = true;
    //                     foundCount++;
    //                     break;
    //                 }
    //             }
    //             assertTrue(found, "Strategy with non-zero shares not found in returned array");
    //         }
    //     }
    //     assertEq(foundCount, expectedLength, "Number of found strategies doesn't match expected length");
    // }
}
