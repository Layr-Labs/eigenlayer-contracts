// to allow calling ERC20 token within this spec
using ERC20 as token;

methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external => DISPATCHER(true);
    function _.isDelegated(address) external => DISPATCHER(true);
    function _.delegatedTo(address) external => DISPATCHER(true);
	function _.decreaseDelegatedShares(address,address[],uint256[],bool) external => DISPATCHER(true);
	function _.increaseDelegatedShares(address,address,uint256) external => DISPATCHER(true);

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

	// external calls to StrategyManager
    function _.getDeposits(address) external;
    function _.slasher() external;
	function _.deposit(address,uint256) external;
	function _.withdraw(address,address,uint256) external;

	// external calls to Strategy
    function _.deposit(address, uint256) external => DISPATCHER(true);
    function _.withdraw(address, address, uint256) external => DISPATCHER(true);
    function _.totalShares() external => DISPATCHER(true);

	// external calls to EigenPodManager
	function _.withdrawRestakedBeaconChainETH(address,address,uint256) external => DISPATCHER(true);
    // call made to EigenPodManager by DelayedWithdrawalRouter
    function _.getPod(address) external => DISPATCHER(true);

    // external calls to EigenPod (from EigenPodManager)
    function _.withdrawRestakedBeaconChainETH(address, uint256) external => DISPATCHER(true);
	    
    // external calls to DelayedWithdrawalRouter (from EigenPod)
    function _.createDelayedWithdrawal(address, address) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

    // calls to ERC20 in this spec
    function token.balanceOf(address) external returns(uint256) envfree;

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);

    //// Harnessed Functions
    // Harnessed calls
    // Harnessed getters
    function strategy_is_in_stakers_array(address, address) external returns (bool) envfree;
    function num_times_strategy_is_in_stakers_array(address, address) external returns (uint256) envfree;
    function totalShares(address) external returns (uint256) envfree;

	//// Normal Functions
	function stakerStrategyListLength(address) external returns (uint256) envfree;
    function stakerStrategyList(address, uint256) external returns (address) envfree;
    function stakerStrategyShares(address, address) external returns (uint256) envfree;
    function array_exhibits_properties(address) external returns (bool) envfree;
}

invariant stakerStrategyListLengthLessThanOrEqualToMax(address staker)
	stakerStrategyListLength(staker) <= 32;

// verifies that strategies in the staker's array of strategies are not duplicated, and that the staker has nonzero shares in each one
invariant arrayExhibitsProperties(address staker)
    array_exhibits_properties(staker) == true
        {
            preserved
            {
                requireInvariant stakerStrategyListLengthLessThanOrEqualToMax(staker);
            }
        }

// if a strategy is *not* in staker's array of strategies, then the staker should have precisely zero shares in that strategy
invariant strategiesNotInArrayHaveZeroShares(address staker, uint256 index)
    (index >= stakerStrategyListLength(staker)) => (stakerStrategyShares(staker, stakerStrategyList(staker, index)) == 0);

/**
* a staker's amount of shares in a strategy (i.e. `stakerStrategyShares[staker][strategy]`) should only increase when
* `depositIntoStrategy`, `depositIntoStrategyWithSignature`, or `depositBeaconChainETH` has been called
* *OR* when completing a withdrawal
*/
definition methodCanIncreaseShares(method f) returns bool =
    f.selector == sig:depositIntoStrategy(address,address,uint256).selector
    || f.selector == sig:depositIntoStrategyWithSignature(address,address,uint256,address,uint256,bytes).selector
    || f.selector == sig:completeQueuedWithdrawal(IStrategyManager.QueuedWithdrawal,address[],uint256,bool).selector;
    // || f.selector == sig:slashQueuedWithdrawal(address,bytes,address[],uint256[]).selector
    // || f.selector == sig:slashShares(address,address,address[],address[],uint256[],uint256[]).selector;

/**
* a staker's amount of shares in a strategy (i.e. `stakerStrategyShares[staker][strategy]`) should only decrease when
* `queueWithdrawal`, `slashShares`, or `recordBeaconChainETHBalanceUpdate` has been called
*/
definition methodCanDecreaseShares(method f) returns bool =
    f.selector == sig:queueWithdrawal(uint256[],address[],uint256[],address,bool).selector
    || f.selector == sig:slashShares(address,address,address[],address[],uint256[],uint256[]).selector
    || f.selector == sig:slashSharesSinglet(address,address,address,address,uint256,uint256).selector;

rule sharesAmountsChangeOnlyWhenAppropriateFunctionsCalled(address staker, address strategy) {
    uint256 sharesBefore = stakerStrategyShares(staker, strategy);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    uint256 sharesAfter = stakerStrategyShares(staker, strategy);
    assert(sharesAfter > sharesBefore => methodCanIncreaseShares(f));
    assert(sharesAfter < sharesBefore => methodCanDecreaseShares(f));
}

// based on Certora's example here https://github.com/Certora/Tutorials/blob/michael/ethcc/EthCC/Ghosts/ghostTest.spec
ghost mapping(address => mathint) sumOfSharesInStrategy {
  init_state axiom forall address strategy. sumOfSharesInStrategy[strategy] == 0;
}

hook Sstore stakerStrategyShares[KEY address staker][KEY address strategy] uint256 newValue (uint256 oldValue) STORAGE {
    sumOfSharesInStrategy[strategy] = sumOfSharesInStrategy[strategy] + newValue - oldValue;
}

/**
* Verifies that the `totalShares` returned by an Strategy is always greater than or equal to the sum of shares in the `stakerStrategyShares`
* mapping -- specifically, that `strategy.totalShares() >= sum_over_all_stakers(stakerStrategyShares[staker][strategy])`
* We cannot show strict equality here, since the withdrawal process first decreases a staker's shares (when `queueWithdrawal` is called) and
* only later is `totalShares` decremented (when `completeQueuedWithdrawal` is called).
*/
invariant totalSharesGeqSumOfShares(address strategy)
    to_mathint(totalShares(strategy)) >= sumOfSharesInStrategy[strategy]
    // preserved block since does not apply for 'beaconChainETH'
    { preserved {
        // 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0 converted to decimal (this is the address of the virtual 'beaconChainETH' strategy)
        // require strategy != beaconChainETHStrategy();
        require strategy != 1088545275507480024404324736574744392984337050304;
    } }

/**
 * Verifies that ERC20 tokens are transferred out of the account only of the msg.sender.
 * Called 'safeApprovalUse' since approval-related vulnerabilites in general allow a caller to transfer tokens out of a different account.
 * This behavior is not always unsafe, but since we don't ever use it (at present) we can do a blanket-check against it.
 */
rule safeApprovalUse(address user) {
    uint256 tokenBalanceBefore = token.balanceOf(user);
    method f;
    env e;
    calldataarg args;
    // need special case for `slashShares` function since otherwise this rule fails by making the user address one of the slashed strategy(s)
    if (
        f.selector == sig:slashShares(address,address,address[],address[],uint256[],uint256[]).selector
        || f.selector == sig:slashSharesSinglet(address,address,address,address,uint256,uint256).selector
    ) {
        address slashedAddress;
        address recipient;
        address strategy;
        address desiredToken;
        uint256 strategyIndex;
        uint256 shareAmount;
        // need this filtering here
        require(strategy != user);
        slashSharesSinglet(e, slashedAddress, recipient, strategy, desiredToken, strategyIndex, shareAmount);
    } else {
        f(e,args);
    }
    uint256 tokenBalanceAfter = token.balanceOf(user);
    if (tokenBalanceAfter < tokenBalanceBefore) {
        assert(e.msg.sender == user, "unsafeApprovalUse?");
    }
    assert true;
}