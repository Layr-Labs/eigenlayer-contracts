// to allow calling ERC20 token within this spec
using ERC20 as token;

methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external => DISPATCHER(true);
    function _.isDelegated(address) external => DISPATCHER(true);
    function _.delegatedTo(address) external => DISPATCHER(true);
	function _.decreaseDelegatedShares(address,address,uint256) external => DISPATCHER(true);
	function _.increaseDelegatedShares(address,address,uint256) external => DISPATCHER(true);

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.slasher() external => DISPATCHER(true);
    function _.addShares(address,address,address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256, address) external => DISPATCHER(true);

	// external calls to EigenPodManager
    function _.addShares(address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256) external => DISPATCHER(true);

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);
    function _.deposit(address,uint256) external => DISPATCHER(true);

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
    function _.totalShares() external => DISPATCHER(true);
    // Harnessed getters
    function strategy_is_in_stakers_array(address, address) external returns (bool) envfree;
    function num_times_strategy_is_in_stakers_array(address, address) external returns (uint256) envfree;
    function totalShares(address) external returns (uint256) envfree;
    function get_stakerStrategyShares(address, address) external returns (uint256) envfree;

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
    || f.selector == sig:withdrawSharesAsTokens(address,address,uint256,address).selector
    || f.selector == sig:addShares(address,address,address,uint256).selector;

/**
* a staker's amount of shares in a strategy (i.e. `stakerStrategyShares[staker][strategy]`) should only decrease when
* `queueWithdrawal`, `slashShares`, or `recordBeaconChainETHBalanceUpdate` has been called
*/
definition methodCanDecreaseShares(method f) returns bool =
    f.selector == sig:removeShares(address,address,uint256).selector;

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

/**
* Verifies that the `totalShares` returned by an Strategy increases appropriately when new Strategy shares are issued by the StrategyManager
* contract (specifically as a result of a call to `StrategyManager.depositIntoStrategy` or `StrategyManager.depositIntoStrategyWithSignature`).
* This rule excludes the `addShares` and `removeShares` functions, since these are called only by the DelegationManager, and do not
* "create new shares", but rather represent existing shares being "moved into a withdrawal".
*/
rule newSharesIncreaseTotalShares(address strategy) {
    method f;
    env e;
    uint256 stakerStrategySharesBefore = get_stakerStrategyShares(e.msg.sender, strategy);
    uint256 totalSharesBefore = totalShares(strategy);
    if (
        f.selector == sig:addShares(address, address, address, uint256).selector
        || f.selector == sig:removeShares(address, address, uint256).selector
    ) {
        uint256 totalSharesAfter = totalShares(strategy);
        assert(totalSharesAfter == totalSharesBefore, "total shares changed unexpectedly");
    } else {
        uint256 stakerStrategySharesAfter = get_stakerStrategyShares(e.msg.sender, strategy);
        uint256 totalSharesAfter = totalShares(strategy);
        assert(stakerStrategySharesAfter - stakerStrategySharesBefore == totalSharesAfter - totalSharesBefore, "diffs don't match");
    }
}

/**
 * Verifies that ERC20 tokens are transferred out of the account only of the msg.sender.
 * Called 'safeApprovalUse' since approval-related vulnerabilities in general allow a caller to transfer tokens out of a different account.
 * This behavior is not always unsafe, but since we don't ever use it (at present) we can do a blanket-check against it.
 */
rule safeApprovalUse(address user) {
    uint256 tokenBalanceBefore = token.balanceOf(user);
    method f;
    env e;
    // special case logic, to handle an edge case
    if (f.selector == sig:withdrawSharesAsTokens(address,address,uint256,address).selector) {
        address recipient;
        address strategy;
        uint256 shares;
        // filter out case where the 'user' is the strategy itself
        require(user != strategy);
        withdrawSharesAsTokens(e, recipient, strategy, shares, token);
    // otherwise just perform an arbitrary function call
    } else {
        calldataarg args;
        f(e,args);
    }
    uint256 tokenBalanceAfter = token.balanceOf(user);
    if (tokenBalanceAfter < tokenBalanceBefore) {
        assert(e.msg.sender == user, "unsafeApprovalUse?");
    }
    assert true;
}