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
        strategyManagerImplementation = new StrategyManager(IDelegationManager(address(delegationManagerMock)), pauserRegistry, "v9.9.9");
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
        StrategyBase newStrategyImplementation = new StrategyBase(_strategyManager, _pauserRegistry, "v9.9.9");
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
                keccak256(bytes(bytes.concat(v[0], v[1]))),
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

    // TODO: test depositing into multiple strategies
    // function testFuzz_depositIntoStrategy_MultipleStrategies()
    // /// @notice deploys 'numStratsToAdd' strategies using '_testAddStrategy' and then deposits '1e18' to each of them from 'getOperatorAddress(0)'
    // /// @param numStratsToAdd is the number of strategies being added and deposited into
    // function testDepositStrategies(uint8 numStratsToAdd) public {
    //     _testDepositStrategies(getOperatorAddress(0), 1e18, numStratsToAdd);
    // }

    // TODO: fix old stETH fork test
    //     /// @notice Shadow-forks mainnet and tests depositing stETH tokens into a "StrategyBase" contract.
    // function testForkMainnetDepositSteth() public {
    //     // hard-coded inputs
    //     // address sender = address(this);
    //     uint64 amountToDeposit = 1e12;

    //     // shadow-fork mainnet
    //     try cheats.createFork("mainnet") returns (uint256 forkId) {
    //         cheats.selectFork(forkId);
    //     // If RPC_MAINNET ENV not set, default to this mainnet RPC endpoint
    //     } catch  {
    //         cheats.createSelectFork("https://eth.llamarpc.com");
    //     }

    //     // cast mainnet stETH address to IERC20 interface
    //     // IERC20 steth = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    //     IERC20 underlyingToken = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);

    //     // deploy necessary contracts on the shadow-forked network
    //     // deploy proxy admin for ability to upgrade proxy contracts
    //     eigenLayerProxyAdmin = new ProxyAdmin();
    //     //deploy pauser registry
    //     address[] memory pausers = new address[](1);
    //     pausers[0] = pauser;
    //     eigenLayerPauserReg = new PauserRegistry(pausers, unpauser);
    //     /**
    //      * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
    //      * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
    //      */
    //     emptyContract = new EmptyContract();
    //     delegation = DelegationManager(
    //         address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
    //     );
    //     strategyManager = StrategyManager(
    //         address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
    //     );
    //     eigenPodManager = EigenPodManager(
    //         address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
    //     );

    //     ethPOSDeposit = new ETHPOSDepositMock();
    //     pod = new EigenPod(ethPOSDeposit, eigenPodManager, GOERLI_GENESIS_TIME);

    //     eigenPodBeacon = new UpgradeableBeacon(address(pod));
    //     // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
    //     DelegationManager delegationImplementation = new DelegationManager(avsDirectory, strategyManager, eigenPodManager, allocationManager, MIN_WITHDRAWAL_DELAY);
    //     StrategyManager strategyManagerImplementation = new StrategyManager(delegation);
    //     EigenPodManager eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, delegation);
    //     // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
    //     eigenLayerProxyAdmin.upgradeAndCall(
    //         ITransparentUpgradeableProxy(payable(address(delegation))),
    //         address(delegationImplementation),
    //         abi.encodeWithSelector(
    //             DelegationManager.initialize.selector,
    //             eigenLayerReputedMultisig,
    //             eigenLayerPauserReg,
    //             0 /*initialPausedStatus*/,
    //             minWithdrawalDelayBlocks,
    //             initializeStrategiesToSetDelayBlocks,
    //             initializeWithdrawalDelayBlocks
    //         )
    //     );
    //     eigenLayerProxyAdmin.upgradeAndCall(
    //         ITransparentUpgradeableProxy(payable(address(strategyManager))),
    //         address(strategyManagerImplementation),
    //         abi.encodeWithSelector(
    //             StrategyManager.initialize.selector,
    //             eigenLayerReputedMultisig,
    //             eigenLayerReputedMultisig,
    //             eigenLayerPauserReg,
    //             0/*initialPausedStatus*/
    //         )
    //     );
    //     eigenLayerProxyAdmin.upgradeAndCall(
    //         ITransparentUpgradeableProxy(payable(address(eigenPodManager))),
    //         address(eigenPodManagerImplementation),
    //         abi.encodeWithSelector(
    //             EigenPodManager.initialize.selector,
    //             eigenLayerReputedMultisig,
    //             eigenLayerPauserReg,
    //             0/*initialPausedStatus*/
    //         )
    //     );

    //     // cheat a bunch of ETH to this address
    //     cheats.deal(address(this), 1e20);
    //     // deposit a huge amount of ETH to get ample stETH
    //     {
    //         (bool success, bytes memory returnData) = address(underlyingToken).call{value: 1e20}("");
    //         require(success, "depositing stETH failed");
    //         returnData;
    //     }

    //     // deploy StrategyBase contract implementation, then create upgradeable proxy that points to implementation and initialize it
    //     baseStrategyImplementation = new StrategyBase(strategyManager);
    //     IStrategy stethStrategy = StrategyBase(
    //             address(
    //                 new TransparentUpgradeableProxy(
    //                     address(baseStrategyImplementation),
    //                     address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, eigenLayerPauserReg)
    //                 )
    //             )
    //         );

    //     // REMAINDER OF CODE ADAPTED FROM `_testDepositToStrategy`
    //     // _testDepositToStrategy(sender, amountToDeposit, underlyingToken, stethStrategy);

    //     // whitelist the strategy for deposit, in case it wasn't before
    //     {
    //         cheats.startPrank(strategyManager.strategyWhitelister());
    //         IStrategy[] memory _strategy = new IStrategy[](1);
    //         _strategy[0] = stethStrategy;
    //         strategyManager.addStrategiesToDepositWhitelist(_strategy);
    //         cheats.stopPrank();
    //     }

    //     uint256 operatorSharesBefore = strategyManager.stakerDepositShares(address(this), stethStrategy);
    //     // check the expected output
    //     uint256 expectedSharesOut = stethStrategy.underlyingToShares(amountToDeposit);

    //     underlyingToken.transfer(address(this), amountToDeposit);
    //     cheats.startPrank(address(this));
    //     underlyingToken.approve(address(strategyManager), type(uint256).max);
    //     strategyManager.depositIntoStrategy(stethStrategy, underlyingToken, amountToDeposit);

    //     //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
    //     if (operatorSharesBefore == 0) {
    //         // check that strategy is appropriately added to dynamic array of all of sender's strategies
    //         assertTrue(
    //             strategyManager.stakerStrategyList(address(this), strategyManager.stakerStrategyListLength(address(this)) - 1)
    //                 == stethStrategy,
    //             "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
    //         );
    //     }
    //     // check that the shares out match the expected amount out
    //     // the actual transfer in will be lower by 1-2 wei than expected due to stETH's internal rounding
    //     // to account for this we check approximate rather than strict equivalence here
    //     {
    //         uint256 actualSharesOut = strategyManager.stakerDepositShares(address(this), stethStrategy) - operatorSharesBefore;
    //         require(actualSharesOut >= expectedSharesOut, "too few shares");
    //         require((actualSharesOut * 1000) / expectedSharesOut < 1003, "too many shares");

    //         // additional sanity check for deposit not increasing in value
    //         require(stethStrategy.sharesToUnderlying(actualSharesOut) <= amountToDeposit, "value cannot have increased");
    //         // slippage check
    //         require((stethStrategy.sharesToUnderlying(actualSharesOut) * 1e6) / amountToDeposit >= (1e6 - 1), "bad slippage on first deposit");
    //     }
    //     cheats.stopPrank();
    // }

    // TODO: fix old frontrun depositor test
    // function testFrontrunFirstDepositor(/*uint256 depositAmount*/) public {

    //     //setup addresses
    //     address attacker = address(100);
    //     address user = address(200);

    //     //give 2 ether to attacker and user
    //     weth.transfer(attacker,2 ether);
    //     weth.transfer(user,2 ether);

    //     //attacker FRONTRUN: deposit 1 wei (receive 1 share)
    //     StrategyManager _strategyManager = _whitelistStrategy(strategyManager, wethStrat);

    //     cheats.startPrank(attacker);
    //     weth.approve(address(strategyManager), type(uint256).max);
    //     _strategyManager.depositIntoStrategy(wethStrat, weth, 1 wei);
    //     cheats.stopPrank();

    //     //attacker FRONTRUN: transfer 1 ether into strategy directly to manipulate the value of shares
    //     cheats.prank(attacker);
    //     weth.transfer(address(wethStrat),1 ether);

    //     //user deposits 2 eth into strategy - only gets 1 share due to rounding
    //     cheats.startPrank(user);
    //     weth.approve(address(_strategyManager), type(uint256).max);
    //     _strategyManager.depositIntoStrategy(wethStrat, weth, 2 ether);
    //     cheats.stopPrank();

    //     //attacker deposited 1 ether and 1 wei - received 1 share
    //     //user deposited 2 ether - received X shares
    //     //user has lost 0.5 ether?
    //     (, uint256[] memory shares) = _strategyManager.getDeposits(attacker);
    //     uint256 attackerValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
    //     require(attackerValueWeth >= (1), "attacker got zero shares");

    //     (, shares) = _strategyManager.getDeposits(user);
    //     uint256 userValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
    //     require(userValueWeth >= (1900000000000000000), "user has lost more than 0.1 eth from frontrunning");

    //     uint256 attackerLossesWeth = (2 ether + 1 wei) - attackerValueWeth;
    //     uint256 userLossesWeth = 2 ether - userValueWeth;
    //     require(attackerLossesWeth > userLossesWeth, "griefing attack deals more damage than cost");
    // }

    // TODO: fix old testFrontrunFirstDepositorFuzzed
    // function testFrontrunFirstDepositorFuzzed(uint96 firstDepositAmount, uint96 donationAmount, uint96 secondDepositAmount) public {
    //     // want to only use nonzero amounts or else we'll get reverts
    //     cheats.assume(firstDepositAmount != 0 && secondDepositAmount != 0);

    //     // setup addresses
    //     address attacker = address(100);
    //     address user = address(200);

    //     // attacker makes first deposit
    //     _testDepositToStrategy(attacker, firstDepositAmount, weth, wethStrat);

    //     // transfer tokens into strategy directly to manipulate the value of shares
    //     weth.transfer(address(wethStrat), donationAmount);

    //     // filter out calls that would revert for minting zero shares
    //     cheats.assume(wethStrat.underlyingToShares(secondDepositAmount) != 0);

    //     // user makes 2nd deposit into strategy - gets diminished shares due to rounding
    //     _testDepositToStrategy(user, secondDepositAmount, weth, wethStrat);

    //     // check for griefing
    //     (, uint256[] memory shares) = strategyManager.getDeposits(attacker);
    //     uint256 attackerValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
    //     (, shares) = strategyManager.getDeposits(user);
    //     uint256 userValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);

    //     uint256 attackerCost = uint256(firstDepositAmount) + uint256(donationAmount);
    //     require(attackerCost >= attackerValueWeth, "attacker gained value?");
    //     // uint256 attackerLossesWeth = attackerValueWeth > attackerCost ? 0 : (attackerCost - attackerValueWeth);
    //     uint256 attackerLossesWeth = attackerCost - attackerValueWeth;
    //     uint256 userLossesWeth = secondDepositAmount - userValueWeth;

    //     emit log_named_uint("attackerLossesWeth", attackerLossesWeth);
    //     emit log_named_uint("userLossesWeth", userLossesWeth);

    //     // use '+1' here to account for rounding. given the attack will cost ETH in the form of gas, this is fine.
    //     require(attackerLossesWeth + 1 >= userLossesWeth, "griefing attack deals more damage than cost");
    // }

    // TODO: testDepositTokenWithOneWeiFeeOnTransfer
    // function testDepositTokenWithOneWeiFeeOnTransfer(address sender, uint64 amountToDeposit) public fuzzedAddress(sender) {
    //     cheats.assume(amountToDeposit != 0);

    //     IERC20 underlyingToken;

    //     {
    //         uint256 initSupply = 1e50;
    //         address initOwner = address(this);
    //         ERC20_OneWeiFeeOnTransfer oneWeiFeeOnTransferToken = new ERC20_OneWeiFeeOnTransfer(initSupply, initOwner);
    //         underlyingToken = IERC20(address(oneWeiFeeOnTransferToken));
    //     }

    //     // need to transfer extra here because otherwise the `sender` won't have enough tokens
    //     underlyingToken.transfer(sender, 1000);

    //     IStrategy oneWeiFeeOnTransferTokenStrategy = StrategyBase(
    //             address(
    //                 new TransparentUpgradeableProxy(
    //                     address(baseStrategyImplementation),
    //                     address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, eigenLayerPauserReg)
    //                 )
    //             )
    //         );

    //     // REMAINDER OF CODE ADAPTED FROM `_testDepositToStrategy`
    //     // _testDepositToStrategy(sender, amountToDeposit, underlyingToken, oneWeiFeeOnTransferTokenStrategy);

    //     // whitelist the strategy for deposit, in case it wasn't before
    //     {
    //         cheats.startPrank(strategyManager.strategyWhitelister());
    //         IStrategy[] memory _strategy = new IStrategy[](1);
    //         _strategy[0] = oneWeiFeeOnTransferTokenStrategy;
    //         strategyManager.addStrategiesToDepositWhitelist(_strategy);
    //         cheats.stopPrank();
    //     }
    //     uint256 operatorSharesBefore = strategyManager.stakerDepositShares(sender, oneWeiFeeOnTransferTokenStrategy);
    //     // check the expected output
    //     uint256 expectedSharesOut = oneWeiFeeOnTransferTokenStrategy.underlyingToShares(amountToDeposit);

    //     underlyingToken.transfer(sender, amountToDeposit);
    //     cheats.startPrank(sender);
    //     underlyingToken.approve(address(strategyManager), type(uint256).max);
    //     strategyManager.depositIntoStrategy(oneWeiFeeOnTransferTokenStrategy, underlyingToken, amountToDeposit);

    //     //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
    //     if (operatorSharesBefore == 0) {
    //         // check that strategy is appropriately added to dynamic array of all of sender's strategies
    //         assertTrue(
    //             strategyManager.stakerStrategyList(sender, strategyManager.stakerStrategyListLength(sender) - 1)
    //                 == oneWeiFeeOnTransferTokenStrategy,
    //             "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
    //         );
    //     }
    //     // check that the shares out match the expected amount out
    //     // the actual transfer in will be lower by 1 wei than expected due to stETH's internal rounding
    //     // to account for this we check approximate rather than strict equivalence here
    //     {
    //         uint256 actualSharesOut = strategyManager.stakerDepositShares(sender, oneWeiFeeOnTransferTokenStrategy) - operatorSharesBefore;
    //         require((actualSharesOut * 1000) / expectedSharesOut > 998, "too few shares");
    //         require((actualSharesOut * 1000) / expectedSharesOut < 1002, "too many shares");

    //         // additional sanity check for deposit not increasing in value
    //         require(oneWeiFeeOnTransferTokenStrategy.sharesToUnderlying(actualSharesOut) <= amountToDeposit, "value cannot have increased");
    //     }
    //     cheats.stopPrank();
    // }

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

contract StrategyManagerUnitTests_increaseBurnableShares is StrategyManagerUnitTests {
    function test_Revert_DelegationManagerModifier() external {
        DelegationManagerMock invalidDelegationManager = new DelegationManagerMock();
        cheats.prank(address(invalidDelegationManager));
        cheats.expectRevert(IStrategyManagerErrors.OnlyDelegationManager.selector);
        strategyManager.increaseBurnableShares(dummyStrat, 1);
    }

    function testFuzz_increaseBurnableShares(uint addedSharesToBurn) external {
        IStrategy strategy = dummyStrat;

        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesIncreased(strategy, addedSharesToBurn);
        cheats.prank(address(delegationManagerMock));
        strategyManager.increaseBurnableShares(strategy, addedSharesToBurn);
        assertEq(
            strategyManager.getBurnableShares(strategy), addedSharesToBurn, "strategyManager.burnableShares(strategy) != addedSharesToBurn"
        );
    }

    function testFuzz_increaseBurnableShares_existingShares(uint existingBurnableShares, uint addedSharesToBurn) external {
        // preventing fuzz overflow, in practice StrategyBase has a 1e38 - 1 maxShares limit so this won't
        // be an issue on mainnet/testnet environments
        existingBurnableShares = bound(existingBurnableShares, 1, type(uint).max / 2);
        addedSharesToBurn = bound(addedSharesToBurn, 1, type(uint).max / 2);

        IStrategy strategy = dummyStrat;
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesIncreased(strategy, existingBurnableShares);
        strategyManager.increaseBurnableShares(strategy, existingBurnableShares);
        assertEq(
            strategyManager.getBurnableShares(strategy),
            existingBurnableShares,
            "strategyManager.burnableShares(strategy) != existingBurnableShares"
        );

        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesIncreased(strategy, addedSharesToBurn);
        strategyManager.increaseBurnableShares(strategy, addedSharesToBurn);

        assertEq(
            strategyManager.getBurnableShares(strategy),
            existingBurnableShares + addedSharesToBurn,
            "strategyManager.burnableShares(strategy) != existingBurnableShares + addedSharesToBurn"
        );
    }
}

contract StrategyManagerUnitTests_burnShares is StrategyManagerUnitTests {
    function testFuzz_SingleStrategyDeposited(address staker, uint depositAmount, uint sharesToBurn)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(0));
        cheats.assume(staker != address(dummyStrat));
        cheats.assume(sharesToBurn > 0 && sharesToBurn < dummyToken.totalSupply() && depositAmount >= sharesToBurn);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);

        // slash shares and increase amount to burn from DelegationManager
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesIncreased(strategy, sharesToBurn);
        strategyManager.increaseBurnableShares(strategy, sharesToBurn);

        uint strategyBalanceBefore = token.balanceOf(address(strategy));
        uint burnAddressBalanceBefore = token.balanceOf(strategyManager.DEFAULT_BURN_ADDRESS());
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesDecreased(strategy, sharesToBurn);
        strategyManager.burnShares(strategy);
        uint strategyBalanceAfter = token.balanceOf(address(strategy));
        uint burnAddressBalanceAfter = token.balanceOf(strategyManager.DEFAULT_BURN_ADDRESS());

        assertEq(strategyBalanceBefore - sharesToBurn, strategyBalanceAfter, "strategyBalanceBefore - sharesToBurn != strategyBalanceAfter");
        assertEq(burnAddressBalanceAfter, burnAddressBalanceBefore + sharesToBurn, "balanceAfter != balanceBefore + sharesAmount");

        // Verify strategy was removed from burnable shares
        (address[] memory strategiesAfterBurn,) = strategyManager.getStrategiesWithBurnableShares();
        assertEq(strategiesAfterBurn.length, 0, "Should have no strategies after burning");
        assertEq(strategyManager.getBurnableShares(strategy), 0, "getBurnableShares should return 0 after burning");
    }

    /// @notice check that balances are unchanged with a reverting token but burnShares doesn't revert
    function testFuzz_BurnableSharesUnchangedWithRevertToken(address staker, uint depositAmount, uint sharesToBurn)
        external
        filterFuzzedAddressInputs(staker)
    {
        cheats.assume(staker != address(0));
        cheats.assume(sharesToBurn > 0 && sharesToBurn < dummyToken.totalSupply() && depositAmount >= sharesToBurn);
        IStrategy strategy = dummyStrat;
        IERC20 token = dummyToken;
        _depositIntoStrategySuccessfully(strategy, staker, depositAmount);

        // slash shares and increase amount to burn from DelegationManager
        cheats.prank(address(delegationManagerMock));
        cheats.expectEmit(true, true, true, true, address(strategyManager));
        emit BurnableSharesIncreased(strategy, sharesToBurn);
        strategyManager.increaseBurnableShares(strategy, sharesToBurn);

        // Now set token to be contract that reverts simulating an upgrade
        cheats.etch(address(token), address(revertToken).code);
        ERC20_SetTransferReverting_Mock(address(token)).setTransfersRevert(true);

        cheats.expectRevert("SafeERC20: low-level call failed");
        cheats.prank(address(delegationManagerMock));
        strategyManager.burnShares(strategy);

        assertEq(strategyManager.getBurnableShares(strategy), sharesToBurn, "burnable shares should be unchanged");
    }
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

    function testFuzz_getStrategiesWithBurnableShares_Single(uint sharesToAdd) public {
        //ensure non-zero
        cheats.assume(sharesToAdd > 0);

        // Add burnable shares
        cheats.prank(address(delegationManagerMock));
        strategyManager.increaseBurnableShares(dummyStrat, sharesToAdd);

        // Get strategies with burnable shares
        (address[] memory strategies, uint[] memory shares) = strategyManager.getStrategiesWithBurnableShares();

        // Verify results
        assertEq(strategies.length, 1, "Should have one strategy");
        assertEq(shares.length, 1, "Should have one share amount");
        assertEq(strategies[0], address(dummyStrat), "Wrong strategy address");
        assertEq(shares[0], sharesToAdd, "Wrong shares amount");
    }

    function testFuzz_getStrategiesWithBurnableShares_Multiple(uint[3] calldata sharesToAdd) public {
        IStrategy[] memory strategies = new IStrategy[](3);
        strategies[0] = dummyStrat;
        strategies[1] = dummyStrat2;
        strategies[2] = dummyStrat3;
        uint[3] memory expectedShares;
        uint expectedLength = 0;

        // Add non-zero shares to strategies
        for (uint i = 0; i < 3; i++) {
            expectedShares[i] = sharesToAdd[i];
            if (sharesToAdd[i] > 0) {
                expectedLength++;
                cheats.prank(address(delegationManagerMock));
                strategyManager.increaseBurnableShares(strategies[i], sharesToAdd[i]);
            }
        }

        // Get strategies with burnable shares
        (address[] memory returnedStrategies, uint[] memory returnedShares) = strategyManager.getStrategiesWithBurnableShares();

        // Verify lengths match
        assertEq(returnedStrategies.length, expectedLength, "Wrong number of strategies returned");
        assertEq(returnedShares.length, expectedLength, "Wrong number of share amounts returned");

        // For all strategies with non-zero shares, verify they are in the returned arrays
        uint foundCount = 0;
        for (uint i = 0; i < 3; i++) {
            if (expectedShares[i] > 0) {
                bool found = false;
                for (uint j = 0; j < returnedStrategies.length; j++) {
                    if (returnedStrategies[j] == address(strategies[i])) {
                        assertEq(returnedShares[j], expectedShares[i], "Wrong share amount");
                        found = true;
                        foundCount++;
                        break;
                    }
                }
                assertTrue(found, "Strategy with non-zero shares not found in returned array");
            }
        }
        assertEq(foundCount, expectedLength, "Number of found strategies doesn't match expected length");
    }
}
