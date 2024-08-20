// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./EigenLayerTestHelper.t.sol";
import "../contracts/core/StrategyManagerStorage.sol";
import "./mocks/ERC20_OneWeiFeeOnTransfer.sol";

contract DepositWithdrawTests is EigenLayerTestHelper {
    uint256[] public emptyUintArray;

    /**
     * @notice Verifies that it is possible to deposit WETH
     * @param amountToDeposit Fuzzed input for amount of WETH to deposit
     */
    function testWethDeposit(uint256 amountToDeposit) public returns (uint256 amountDeposited) {
        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.
        cheats.assume(amountToDeposit >= 1);
        cheats.assume(amountToDeposit <= 1e38 - 1);
        return _testDepositWeth(getOperatorAddress(0), amountToDeposit);
    }


    /// @notice deploys 'numStratsToAdd' strategies using '_testAddStrategy' and then deposits '1e18' to each of them from 'getOperatorAddress(0)'
    /// @param numStratsToAdd is the number of strategies being added and deposited into
    function testDepositStrategies(uint8 numStratsToAdd) public {
        _testDepositStrategies(getOperatorAddress(0), 1e18, numStratsToAdd);
    }

    /// @notice Verifies that it is possible to deposit eigen.
    /// @param eigenToDeposit is amount of eigen to deposit into the eigen strategy
    function testDepositEigen(uint96 eigenToDeposit) public {
        // sanity check for inputs; keeps fuzzed tests from failing
        cheats.assume(eigenToDeposit < eigenTotalSupply);
        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.
        cheats.assume(eigenToDeposit >= 1);
        _testDepositEigen(getOperatorAddress(0), eigenToDeposit);
    }

    /**
     * @notice Tries to deposit an unsupported token into an `StrategyBase` contract by calling `strategyManager.depositIntoStrategy`.
     * Verifies that reversion occurs correctly.
     */
    function testDepositUnsupportedToken() public {
        IERC20 token = new ERC20PresetFixedSupply(
            "badToken",
            "BADTOKEN",
            100,
            address(this)
        );
        token.approve(address(strategyManager), type(uint256).max);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.strategyWhitelister());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = wethStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        cheats.expectRevert(bytes("StrategyBase.deposit: Can only deposit underlyingToken"));
        strategyManager.depositIntoStrategy(wethStrat, token, 10);
    }

    /**
     * @notice Tries to deposit into an unsupported strategy by calling `strategyManager.depositIntoStrategy`.
     * Verifies that reversion occurs correctly.
     */
    function testDepositNonexistentStrategy(address nonexistentStrategy) public fuzzedAddress(nonexistentStrategy) {
        // assume that the fuzzed address is not already a contract!
        uint256 size;
        assembly {
            size := extcodesize(nonexistentStrategy)
        }
        cheats.assume(size == 0);
        // check against calls from precompile addresses -- was getting fuzzy failures from this
        cheats.assume(uint160(nonexistentStrategy) > 9);

        // harcoded input
        uint256 testDepositAmount = 10;

        IERC20 token = new ERC20PresetFixedSupply(
            "badToken",
            "BADTOKEN",
            100,
            address(this)
        );
        token.approve(address(strategyManager), type(uint256).max);

        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.strategyWhitelister());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = IStrategy(nonexistentStrategy);
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        cheats.expectRevert();
        strategyManager.depositIntoStrategy(IStrategy(nonexistentStrategy), token, testDepositAmount);
    }

    /// @notice verify that trying to deposit an amount of zero will correctly revert
    function testRevertOnZeroDeposit() public {
        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.strategyWhitelister());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = wethStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        cheats.expectRevert(bytes("StrategyBase.deposit: newShares cannot be zero"));
        strategyManager.depositIntoStrategy(wethStrat, weth, 0);
    }


     /**
     * @notice Modified from existing _createQueuedWithdrawal, skips delegation and deposit steps so that we can isolate the withdrawal step
     * @notice Creates a queued withdrawal from `staker`, queues a withdrawal using
     * `delegation.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawer)`
     * @notice After initiating a queued withdrawal, this test checks that `delegation.canCompleteQueuedWithdrawal` immediately returns the correct
     * response depending on whether `staker` is delegated or not.
     * @param staker The address to initiate the queued withdrawal
     * @param amountToDeposit The amount of WETH to deposit
     */
    function _createOnlyQueuedWithdrawal(
        address staker,
        bool /*registerAsOperator*/,
        uint256 amountToDeposit,
        IStrategy[] memory strategyArray,
        IERC20[] memory /*tokensArray*/,
        uint256[] memory shareAmounts,
        uint256[] memory /*strategyIndexes*/,
        address withdrawer
    )
        internal returns(bytes32 withdrawalRoot, IDelegationManager.Withdrawal memory queuedWithdrawal)
    {
        require(amountToDeposit >= shareAmounts[0], "_createQueuedWithdrawal: sanity check failed");

        queuedWithdrawal = IDelegationManager.Withdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            staker: staker,
            withdrawer: withdrawer,
            nonce: delegation.cumulativeWithdrawalsQueued(staker),
            delegatedTo: delegation.delegatedTo(staker),
            startBlock: uint32(block.number)
        });


        IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
        params[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: withdrawer
        });

        bytes32[] memory withdrawalRoots = new bytes32[](1);

        //queue the withdrawal
        cheats.startPrank(staker);
        withdrawalRoots = delegation.queueWithdrawals(params);
        cheats.stopPrank();
        return (withdrawalRoots[0], queuedWithdrawal);
     }
        
        
    function testFrontrunFirstDepositor(/*uint256 depositAmount*/) public {

        //setup addresses
        address attacker = address(100);
        address user = address(200);

        //give 2 ether to attacker and user
        weth.transfer(attacker,2 ether);
        weth.transfer(user,2 ether);

        //attacker FRONTRUN: deposit 1 wei (receive 1 share)
        StrategyManager _strategyManager = _whitelistStrategy(strategyManager, wethStrat);

        cheats.startPrank(attacker);
        weth.approve(address(strategyManager), type(uint256).max);
        _strategyManager.depositIntoStrategy(wethStrat, weth, 1 wei);
        cheats.stopPrank();

        //attacker FRONTRUN: transfer 1 ether into strategy directly to manipulate the value of shares
        cheats.prank(attacker);
        weth.transfer(address(wethStrat),1 ether);

        //user deposits 2 eth into strategy - only gets 1 share due to rounding 
        cheats.startPrank(user);
        weth.approve(address(_strategyManager), type(uint256).max);
        _strategyManager.depositIntoStrategy(wethStrat, weth, 2 ether);
        cheats.stopPrank();

        //attacker deposited 1 ether and 1 wei - received 1 share 
        //user deposited 2 ether - received X shares
        //user has lost 0.5 ether?
        (, uint256[] memory shares) = _strategyManager.getDeposits(attacker);
        uint256 attackerValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
        require(attackerValueWeth >= (1), "attacker got zero shares");

        (, shares) = _strategyManager.getDeposits(user);
        uint256 userValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
        require(userValueWeth >= (1900000000000000000), "user has lost more than 0.1 eth from frontrunning");

        uint256 attackerLossesWeth = (2 ether + 1 wei) - attackerValueWeth;
        uint256 userLossesWeth = 2 ether - userValueWeth;
        require(attackerLossesWeth > userLossesWeth, "griefing attack deals more damage than cost");
    }

    // fuzzed amounts user uint96's to be more realistic with amounts
    function testFrontrunFirstDepositorFuzzed(uint96 firstDepositAmount, uint96 donationAmount, uint96 secondDepositAmount) public {
        // want to only use nonzero amounts or else we'll get reverts
        cheats.assume(firstDepositAmount != 0 && secondDepositAmount != 0);

        // setup addresses
        address attacker = address(100);
        address user = address(200);

        // attacker makes first deposit
        _testDepositToStrategy(attacker, firstDepositAmount, weth, wethStrat);

        // transfer tokens into strategy directly to manipulate the value of shares
        weth.transfer(address(wethStrat), donationAmount);

        // filter out calls that would revert for minting zero shares
        cheats.assume(wethStrat.underlyingToShares(secondDepositAmount) != 0);

        // user makes 2nd deposit into strategy - gets diminished shares due to rounding
        _testDepositToStrategy(user, secondDepositAmount, weth, wethStrat);

        // check for griefing
        (, uint256[] memory shares) = strategyManager.getDeposits(attacker);
        uint256 attackerValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);
        (, shares) = strategyManager.getDeposits(user);
        uint256 userValueWeth = wethStrat.sharesToUnderlyingView(shares[0]);

        uint256 attackerCost = uint256(firstDepositAmount) + uint256(donationAmount);
        require(attackerCost >= attackerValueWeth, "attacker gained value?");
        // uint256 attackerLossesWeth = attackerValueWeth > attackerCost ? 0 : (attackerCost - attackerValueWeth);
        uint256 attackerLossesWeth = attackerCost - attackerValueWeth;
        uint256 userLossesWeth = secondDepositAmount - userValueWeth;

        emit log_named_uint("attackerLossesWeth", attackerLossesWeth);
        emit log_named_uint("userLossesWeth", userLossesWeth);

        // use '+1' here to account for rounding. given the attack will cost ETH in the form of gas, this is fine.
        require(attackerLossesWeth + 1 >= userLossesWeth, "griefing attack deals more damage than cost");
    }


    function testDepositTokenWithOneWeiFeeOnTransfer(address sender, uint64 amountToDeposit) public fuzzedAddress(sender) {
        cheats.assume(amountToDeposit != 0);

        IERC20 underlyingToken;

        {
            uint256 initSupply = 1e50;
            address initOwner = address(this);
            ERC20_OneWeiFeeOnTransfer oneWeiFeeOnTransferToken = new ERC20_OneWeiFeeOnTransfer(initSupply, initOwner);
            underlyingToken = IERC20(address(oneWeiFeeOnTransferToken));
        }

        // need to transfer extra here because otherwise the `sender` won't have enough tokens
        underlyingToken.transfer(sender, 1000);

        IStrategy oneWeiFeeOnTransferTokenStrategy = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, eigenLayerPauserReg)
                    )
                )
            );

        // REMAINDER OF CODE ADAPTED FROM `_testDepositToStrategy`
        // _testDepositToStrategy(sender, amountToDeposit, underlyingToken, oneWeiFeeOnTransferTokenStrategy);

        // whitelist the strategy for deposit, in case it wasn't before
        {
            cheats.startPrank(strategyManager.strategyWhitelister());
            IStrategy[] memory _strategy = new IStrategy[](1);
            bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
            _strategy[0] = oneWeiFeeOnTransferTokenStrategy;
            strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
            cheats.stopPrank();
        }
    
        uint256 operatorSharesBefore = strategyManager.stakerStrategyShares(sender, oneWeiFeeOnTransferTokenStrategy);
        // check the expected output
        uint256 expectedSharesOut = oneWeiFeeOnTransferTokenStrategy.underlyingToShares(amountToDeposit);

        underlyingToken.transfer(sender, amountToDeposit);
        cheats.startPrank(sender);
        underlyingToken.approve(address(strategyManager), type(uint256).max);
        strategyManager.depositIntoStrategy(oneWeiFeeOnTransferTokenStrategy, underlyingToken, amountToDeposit);

        //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
        if (operatorSharesBefore == 0) {
            // check that strategy is appropriately added to dynamic array of all of sender's strategies
            assertTrue(
                strategyManager.stakerStrategyList(sender, strategyManager.stakerStrategyListLength(sender) - 1)
                    == oneWeiFeeOnTransferTokenStrategy,
                "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
            );
        }
            
        // check that the shares out match the expected amount out
        // the actual transfer in will be lower by 1 wei than expected due to stETH's internal rounding
        // to account for this we check approximate rather than strict equivalence here
        {
            uint256 actualSharesOut = strategyManager.stakerStrategyShares(sender, oneWeiFeeOnTransferTokenStrategy) - operatorSharesBefore;
            require((actualSharesOut * 1000) / expectedSharesOut > 998, "too few shares");
            require((actualSharesOut * 1000) / expectedSharesOut < 1002, "too many shares");

            // additional sanity check for deposit not increasing in value
            require(oneWeiFeeOnTransferTokenStrategy.sharesToUnderlying(actualSharesOut) <= amountToDeposit, "value cannot have increased");
        }
        cheats.stopPrank();
    }

    /// @notice Shadow-forks mainnet and tests depositing stETH tokens into a "StrategyBase" contract.
    function testForkMainnetDepositSteth() public {
        // hard-coded inputs
        // address sender = address(this);
        uint64 amountToDeposit = 1e12;

        // shadow-fork mainnet
        try cheats.createFork("mainnet") returns (uint256 forkId) {
            cheats.selectFork(forkId);
        // If RPC_MAINNET ENV not set, default to this mainnet RPC endpoint
        } catch  {
            cheats.createSelectFork("https://eth.llamarpc.com");
        }

        // cast mainnet stETH address to IERC20 interface
        // IERC20 steth = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
        IERC20 underlyingToken = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);

        // deploy necessary contracts on the shadow-forked network
        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();
        //deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        eigenLayerPauserReg = new PauserRegistry(pausers, unpauser);
        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        emptyContract = new EmptyContract();
        delegation = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManager = StrategyManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        slasher = Slasher(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        ethPOSDeposit = new ETHPOSDepositMock();
        pod = new EigenPod(ethPOSDeposit, eigenPodManager, GOERLI_GENESIS_TIME);

        eigenPodBeacon = new UpgradeableBeacon(address(pod));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        StrategyManager strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        Slasher slasherImplementation = new Slasher(strategyManager, delegation);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, slasher, delegation);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0 /*initialPausedStatus*/,
                minWithdrawalDelayBlocks,
                initializeStrategiesToSetDelayBlocks,
                initializeWithdrawalDelayBlocks
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                eigenLayerReputedMultisig,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0/*initialPausedStatus*/
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(
                Slasher.initialize.selector,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0/*initialPausedStatus*/
            )
        );
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0/*initialPausedStatus*/
            )
        );

        // cheat a bunch of ETH to this address
        cheats.deal(address(this), 1e20);
        // deposit a huge amount of ETH to get ample stETH
        {
            (bool success, bytes memory returnData) = address(underlyingToken).call{value: 1e20}("");
            require(success, "depositing stETH failed");
            returnData;
        }

        // deploy StrategyBase contract implementation, then create upgradeable proxy that points to implementation and initialize it
        baseStrategyImplementation = new StrategyBase(strategyManager);
        IStrategy stethStrategy = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, underlyingToken, eigenLayerPauserReg)
                    )
                )
            );

        // REMAINDER OF CODE ADAPTED FROM `_testDepositToStrategy`
        // _testDepositToStrategy(sender, amountToDeposit, underlyingToken, stethStrategy);

        // whitelist the strategy for deposit, in case it wasn't before
        {
            cheats.startPrank(strategyManager.strategyWhitelister());
            IStrategy[] memory _strategy = new IStrategy[](1);
            bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
            _strategy[0] = stethStrategy;
            strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
            cheats.stopPrank();
        }

        uint256 operatorSharesBefore = strategyManager.stakerStrategyShares(address(this), stethStrategy);
        // check the expected output
        uint256 expectedSharesOut = stethStrategy.underlyingToShares(amountToDeposit);

        underlyingToken.transfer(address(this), amountToDeposit);
        cheats.startPrank(address(this));
        underlyingToken.approve(address(strategyManager), type(uint256).max);
        strategyManager.depositIntoStrategy(stethStrategy, underlyingToken, amountToDeposit);

        //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
        if (operatorSharesBefore == 0) {
            // check that strategy is appropriately added to dynamic array of all of sender's strategies
            assertTrue(
                strategyManager.stakerStrategyList(address(this), strategyManager.stakerStrategyListLength(address(this)) - 1)
                    == stethStrategy,
                "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
            );
        }
            
        // check that the shares out match the expected amount out
        // the actual transfer in will be lower by 1-2 wei than expected due to stETH's internal rounding
        // to account for this we check approximate rather than strict equivalence here
        {
            uint256 actualSharesOut = strategyManager.stakerStrategyShares(address(this), stethStrategy) - operatorSharesBefore;
            require(actualSharesOut >= expectedSharesOut, "too few shares");
            require((actualSharesOut * 1000) / expectedSharesOut < 1003, "too many shares");

            // additional sanity check for deposit not increasing in value
            require(stethStrategy.sharesToUnderlying(actualSharesOut) <= amountToDeposit, "value cannot have increased");
            // slippage check
            require((stethStrategy.sharesToUnderlying(actualSharesOut) * 1e6) / amountToDeposit >= (1e6 - 1), "bad slippage on first deposit");
        }
        cheats.stopPrank();
    }

    function _whitelistStrategy(StrategyManager _strategyManager, StrategyBase _strategyBase) internal returns(StrategyManager) {
        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.strategyWhitelister());
        IStrategy[] memory _strategy = new IStrategy[](1);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
        _strategy[0] = IStrategy(_strategyBase);
        _strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
        cheats.stopPrank();

        return _strategyManager;
    }
}
