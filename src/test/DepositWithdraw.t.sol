// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

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
        cheats.assume(amountToDeposit >= 1e9);
        return _testDepositWeth(getOperatorAddress(0), amountToDeposit);
    }

    function testPreventSlashing() public {
        //use preexisting helper function to set up a withdrawal
        address middleware = address(0xdeadbeef);
        address staker = getOperatorAddress(0);
        uint256 depositAmount = 1 ether;
        IStrategy strategy = wethStrat;
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;

        //invalid token
        IERC20[] memory tokensArray = new IERC20[](1);
        tokensArray[0] = IERC20(address(0));
        
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = depositAmount - 1 gwei; //leave some shares behind so we don't get undelegation issues
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        address withdrawer = staker;

        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;

        ( ,queuedWithdrawal) = _createQueuedWithdrawal(staker, 
                                true,
                                depositAmount,
                                strategyArray,
                                shareAmounts,
                                strategyIndexes,
                                withdrawer
                                );

        cheats.startPrank(staker);
        //opt in staker to restake for the two middlewares we are using
        slasher.optIntoSlashing(middleware);
        cheats.stopPrank();

        //move ahead a block after queuing the withdrawal
        cheats.roll(2);

        cheats.startPrank(middleware);
        // stake update with updateBlock = 2, serveUntilBlock = 5
        uint32 serveUntilBlock = 5;
        slasher.recordFirstStakeUpdate(staker, serveUntilBlock);
        cheats.stopPrank();

        cheats.roll(6);

        // freeze the staker
        cheats.startPrank(middleware);
        slasher.freezeOperator(staker);
        cheats.stopPrank();

        // attempt to slash - reverts
        cheats.startPrank(strategyManager.owner());
        cheats.expectRevert("StrategyBase.withdraw: Can only withdraw the strategy token");
        strategyManager.slashQueuedWithdrawal(address(slasher), queuedWithdrawal, tokensArray, emptyUintArray);
        cheats.stopPrank();

        //staker is unfrozen at a future date
        address[] memory addressArray = new address[](1);
        addressArray[0] = staker;
        cheats.startPrank(slasher.owner());
        slasher.resetFrozenStatus(addressArray);
        cheats.stopPrank();

        // staker can still withdraw shares
        cheats.startPrank(staker);
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, 0, false);
        cheats.stopPrank();
    }

    function testWithdrawalSequences() public {
        //use preexisting helper function to set up a withdrawal
        address middleware = address(0xdeadbeef);
        address middleware_2 = address(0x009849);
        address staker = getOperatorAddress(0);
        IStrategyManager.QueuedWithdrawal memory queuedWithdrawal;
        
        uint256 depositAmount = 1 ether;
        IStrategy strategy = wethStrat;
        IERC20 underlyingToken = weth;
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;
        IERC20[] memory tokensArray = new IERC20[](1);
        tokensArray[0] = underlyingToken;
        {
        uint256[] memory shareAmounts = new uint256[](1);
        shareAmounts[0] = depositAmount - 1 gwei; //leave some shares behind so we don't get undelegation issues
        uint256[] memory strategyIndexes = new uint256[](1);
        strategyIndexes[0] = 0;
        address withdrawer = staker;
        
        {   
            assertTrue(!delegation.isDelegated(staker), "_createQueuedWithdrawal: staker is already delegated");
            _testRegisterAsOperator(staker, IDelegationTerms(staker));
            assertTrue(
                delegation.isDelegated(staker), "_createQueuedWithdrawal: staker isn't delegated when they should be"
            );
        
            //make deposit in WETH strategy
            uint256 amountDeposited = _testDepositWeth(staker, depositAmount);
            // We can't withdraw more than we deposit
            if (shareAmounts[0] > amountDeposited) {
                cheats.expectRevert("StrategyManager._removeShares: shareAmount too high");
            }
        }

        
            cheats.startPrank(staker);
            //opt in staker to restake for the two middlewares we are using
            slasher.optIntoSlashing(middleware);
            slasher.optIntoSlashing(middleware_2);
            cheats.stopPrank();

            cheats.startPrank(middleware);
            // first stake update with updateBlock = 1, serveUntilBlock = 5
      
            uint32 serveUntilBlock = 5;
            slasher.recordFirstStakeUpdate(staker, serveUntilBlock);
            cheats.stopPrank();
            //check middlewareTimes entry is correct
            require(slasher.getMiddlewareTimesIndexBlock(staker, 0) == 1, "middleware updateBlock update incorrect");
            require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 0) == 5, "middleware serveUntil update incorrect");
            

            cheats.startPrank(middleware_2);
            // first stake update with updateBlock = 1, serveUntilBlock = 6
            slasher.recordFirstStakeUpdate(staker, serveUntilBlock+1);
            cheats.stopPrank();
            //check middlewareTimes entry is correct
            require(slasher.getMiddlewareTimesIndexBlock(staker, 1) == 1, "middleware updateBlock update incorrect");
            require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 1) == 6, "middleware serveUntil update incorrect");
            //check old entry has not changed
            require(slasher.getMiddlewareTimesIndexBlock(staker, 0) == 1, "middleware updateBlock update incorrect");
            require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 0) == 5, "middleware serveUntil update incorrect");

            //move ahead a block before queuing the withdrawal
            cheats.roll(2);
            //cheats.startPrank(staker);
            //queue the withdrawal
            ( ,queuedWithdrawal) = _createOnlyQueuedWithdrawal(staker, 
                                true,
                                depositAmount,
                                strategyArray,
                                tokensArray,
                                shareAmounts,
                                strategyIndexes,
                                withdrawer
                                );
        
            }
        //Because the staker has queued a withdrawal both currently staked middlewares must issued an update as required for the completion of the withdrawal
        //to be realistic we move ahead a block before updating middlewares
        cheats.roll(3);

        cheats.startPrank(middleware);
        // stake update with updateBlock = 3, serveUntilBlock = 7
        uint32 serveUntilBlock = 7;
        uint32 updateBlock = 3;
        uint256 insertAfter = 1;
        slasher.recordStakeUpdate(staker, updateBlock, serveUntilBlock, insertAfter);
        cheats.stopPrank();
        //check middlewareTimes entry is correct
        require(slasher.getMiddlewareTimesIndexBlock(staker, 2) == 1, "middleware updateBlock update incorrect");
        require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 2) == 7, "middleware serveUntil update incorrect");

        cheats.startPrank(middleware_2);
        // stake update with updateBlock = 3, serveUntilBlock = 10
        slasher.recordStakeUpdate(staker, updateBlock, serveUntilBlock+3, insertAfter);
        cheats.stopPrank();
        //check middlewareTimes entry is correct
        require(slasher.getMiddlewareTimesIndexBlock(staker, 3) == 3, "middleware updateBlock update incorrect");
        require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 3) == 10, "middleware serveUntil update incorrect");

        cheats.startPrank(middleware);
        // stake update with updateBlock = 3, serveUntilBlock = 7
        serveUntilBlock = 7;
        updateBlock = 3;
        insertAfter = 2;
        slasher.recordStakeUpdate(staker, updateBlock, serveUntilBlock, insertAfter);
        cheats.stopPrank();
        //check middlewareTimes entry is correct
        require(slasher.getMiddlewareTimesIndexBlock(staker, 4) == 3, "middleware updateBlock update incorrect");
        require(slasher.getMiddlewareTimesIndexServeUntilBlock(staker, 4) == 10, "middleware serveUntil update incorrect");

        //move timestamp to 6, one middleware is past serveUntilBlock but the second middleware is still using the restaked funds.
        cheats.warp(8);
        //Also move the current block ahead one
        cheats.roll(4);
        
        cheats.startPrank(staker);
        //when called with the correct middlewareTimesIndex the call reverts

        slasher.getMiddlewareTimesIndexBlock(staker, 3);
        
        
        {
        uint256 correctMiddlewareTimesIndex = 4;
        cheats.expectRevert("StrategyManager.completeQueuedWithdrawal: shares pending withdrawal are still slashable");
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, correctMiddlewareTimesIndex, false);
        }

        //When called with a stale index the call should also revert.
        {
        uint256 staleMiddlewareTimesIndex = 2;
        cheats.expectRevert("StrategyManager.completeQueuedWithdrawal: shares pending withdrawal are still slashable");
        strategyManager.completeQueuedWithdrawal(queuedWithdrawal, tokensArray, staleMiddlewareTimesIndex, false);
        }
        
        
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
        cheats.assume(eigenToDeposit >= 1e9);
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
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = wethStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
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
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = IStrategy(nonexistentStrategy);
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        cheats.expectRevert();
        strategyManager.depositIntoStrategy(IStrategy(nonexistentStrategy), token, testDepositAmount);
    }

    /// @notice verify that trying to deposit an amount of zero will correctly revert
    function testRevertOnZeroDeposit() public {
        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = wethStrat;
        strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        cheats.expectRevert(bytes("StrategyBase.deposit: newShares cannot be zero"));
        strategyManager.depositIntoStrategy(wethStrat, weth, 0);
        cheats.stopPrank();
    }


     /**
     * @notice Modified from existing _createQueuedWithdrawal, skips delegation and deposit steps so that we can isolate the withdrawal step
     * @notice Creates a queued withdrawal from `staker`, queues a withdrawal using
     * `strategyManager.queueWithdrawal(strategyIndexes, strategyArray, tokensArray, shareAmounts, withdrawer)`
     * @notice After initiating a queued withdrawal, this test checks that `strategyManager.canCompleteQueuedWithdrawal` immediately returns the correct
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
        uint256[] memory strategyIndexes,
        address withdrawer
    )
        internal returns(bytes32 withdrawalRoot, IStrategyManager.QueuedWithdrawal memory queuedWithdrawal)
    {
        require(amountToDeposit >= shareAmounts[0], "_createQueuedWithdrawal: sanity check failed");

        IStrategyManager.WithdrawerAndNonce memory withdrawerAndNonce = IStrategyManager.WithdrawerAndNonce({
            withdrawer: withdrawer,
            nonce: uint96(strategyManager.numWithdrawalsQueued(staker))
        });

        queuedWithdrawal = IStrategyManager.QueuedWithdrawal({
            strategies: strategyArray,
            shares: shareAmounts,
            depositor: staker,
            withdrawerAndNonce: withdrawerAndNonce,
            delegatedAddress: delegation.delegatedTo(staker),
            withdrawalStartBlock: uint32(block.number)
        });

        

        //queue the withdrawal
        cheats.startPrank(staker);
        withdrawalRoot = strategyManager.queueWithdrawal(strategyIndexes, strategyArray, shareAmounts, withdrawer, true);
        cheats.stopPrank();
        return (withdrawalRoot, queuedWithdrawal);
     }
        
        
    function testFrontrunFirstDepositor(/*uint256 depositAmount*/) public {

        //setup addresses
        address attacker = address(100);
        address user = address(200);

        //give 2 ether to attacker and user
        weth.transfer(attacker,2 ether);
        weth.transfer(user,2 ether);


        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.

        //attacker FRONTRUN: deposit 1 wei (receive 1 share)
        StrategyManager _strategyManager = _whitelistStrategy(strategyManager, wethStrat);

        cheats.startPrank(attacker);
        weth.approve(address(strategyManager), type(uint256).max);
        cheats.expectRevert("StrategyBase.deposit: updated totalShares amount would be nonzero but below MIN_NONZERO_TOTAL_SHARES");
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
        //user deposited 2 ether - received 1 share
        //user has lost 0.5 ether
        (, uint256[] memory shares) = _strategyManager.getDeposits(attacker);
        require(shares.length == 0, "Attacker deposit should fail due to minimum balances");

        (, shares) = _strategyManager.getDeposits(user);
        require(wethStrat.sharesToUnderlyingView(shares[0]) >= (1900000000000000000), "user has lost more than 0.1 eth from frontrunning");

    }

    function testDepositTokenWithOneWeiFeeOnTransfer(address sender, uint64 amountToDeposit) public fuzzedAddress(sender) {
        // MIN_NONZERO_TOTAL_SHARES = 1e9
        cheats.assume(amountToDeposit >= 1e9);

        uint256 initSupply = 1e50;
        address initOwner = address(this);

        ERC20_OneWeiFeeOnTransfer oneWeiFeeOnTransferToken = new ERC20_OneWeiFeeOnTransfer(initSupply, initOwner);
        IERC20 underlyingToken = IERC20(address(oneWeiFeeOnTransferToken));

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

        _testDepositToStrategy(sender, amountToDeposit, underlyingToken, oneWeiFeeOnTransferTokenStrategy);
    }

    /// @notice Shadow-forks mainnet and tests depositing stETH tokens into a "StrategyBase" contract.
    function testForkMainnetDepositSteth() public {
        // hard-coded inputs
        address sender = address(this);
        uint64 amountToDeposit = 1e12;

        // shadow-fork mainnet
        uint256 forkId = cheats.createFork("mainnet");
        cheats.selectFork(forkId);

        // cast mainnet stETH address to IERC20 interface
        IERC20 steth = IERC20(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
        IERC20 underlyingToken = steth;

        // deploy necessary contracts on the shadow-forked network
        // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();
        //deploy pauser registry
        eigenLayerPauserReg = new PauserRegistry(pauser, unpauser);
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
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        address[] memory initialOracleSignersArray = new address[](0);
        beaconChainOracle = new BeaconChainOracle(eigenLayerReputedMultisig, initialBeaconChainOracleThreshold, initialOracleSignersArray);

        ethPOSDeposit = new ETHPOSDepositMock();
        pod = new EigenPod(ethPOSDeposit, delayedWithdrawalRouter, eigenPodManager, REQUIRED_BALANCE_WEI);

        eigenPodBeacon = new UpgradeableBeacon(address(pod));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher);
        StrategyManager strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        Slasher slasherImplementation = new Slasher(strategyManager, delegation);
        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(ethPOSDeposit, eigenPodBeacon, strategyManager, slasher);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0/*initialPausedStatus*/
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
                0/*initialPausedStatus*/,
                0/*withdrawalDelayBlocks*/
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
                type(uint256).max,
                beaconChainOracle,
                eigenLayerReputedMultisig,
                eigenLayerPauserReg,
                0/*initialPausedStatus*/
            )
        );

        // cheat a bunch of ETH to this address
        cheats.deal(address(this), 1e20);
        // deposit a huge amount of ETH to get ample stETH
        (bool success, bytes memory returnData) = address(steth).call{value: 1e20}("");
        require(success, "depositing stETH failed");
        returnData;

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

        _testDepositToStrategy(sender, amountToDeposit, underlyingToken, stethStrategy);

    }

    function _whitelistStrategy(StrategyManager _strategyManager, StrategyBase _strategyBase) internal returns(StrategyManager) {
        // whitelist the strategy for deposit
        cheats.startPrank(strategyManager.owner());
        IStrategy[] memory _strategy = new IStrategy[](1);
        _strategy[0] = IStrategy(_strategyBase);
        _strategyManager.addStrategiesToDepositWhitelist(_strategy);
        cheats.stopPrank();

        return _strategyManager;
    }
}
