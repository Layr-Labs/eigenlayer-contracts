import "../erc20cvl.spec";
import "../ptaHelpers.spec";

using StrategyBase as StrategyBase;

methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external => DISPATCHER(true);
    function _.isDelegated(address) external => DISPATCHER(true);
    function _.delegatedTo(address) external => DISPATCHER(true);
	function _.decreaseDelegatedShares(address,uint256,uint64) external => DISPATCHER(true);
	function _.increaseDelegatedShares(address,address,uint256,uint256) external => DISPATCHER(true);

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.addShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, address, uint256) external => DISPATCHER(true);
    function _.underlyingToken() external => DISPATCHER(true);

	// external calls to EigenPodManager
//    function EigenPodManager.addShares(address,address,address,uint256) external => DISPATCHER(true); //summarized under StrategyManager
//    function _.withdrawSharesAsTokens(address, address,address,uint256) external => DISPATCHER(true); //summarized under StrategyManager

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);
    function _.deposit(address,uint256) external => DISPATCHER(true);

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);

    function StrategyBase.sharesToUnderlyingView(uint256) external returns (uint256) envfree;

    function _.sharesToUnderlyingView(uint256 shares) external => cvlSharesToUnderlyingView(shares) expect uint256;
    //// Harnessed Functions
    // Harnessed calls
    function _.totalShares() external => DISPATCHER(true);
    // Harnessed getters
    function strategy_is_in_stakers_array(address, address) external returns (bool) envfree;
    function num_times_strategy_is_in_stakers_array(address, address) external returns (uint256) envfree;
    function totalShares(address) external returns (uint256) envfree;
    function get_stakerDepositShares(address, address) external returns (uint256) envfree;
    function sharesToUnderlyingView(address, uint256) external returns (uint256) envfree;

	//// Normal Functions
	function stakerStrategyListLength(address) external returns (uint256) envfree;
    function stakerStrategyList(address, uint256) external returns (address) envfree;
    function stakerDepositShares(address, address) external returns (uint256) envfree;
    function array_exhibits_properties(address) external returns (bool) envfree;
    function getBurnableShares(address) external returns (uint256) envfree;

    //// Normal getters
    function DEFAULT_BURN_ADDRESS() external returns (address) envfree;
}

function cvlSharesToUnderlyingView(uint256 shares) returns uint256 {
    return StrategyBase.sharesToUnderlyingView(shares);
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
// The proof of this invariant uses only solidity semantics hardcoded to the prover.  If the index is out of bounds,
// solidity will revert so the statement is vacuously true ([Accessing an array past its end causes a failing assertion](https://docs.soliditylang.org/en/latest/types.html#arrays))
//invariant strategiesNotInArrayHaveZeroShares(address staker, uint256 index)
//    (index >= stakerStrategyListLength(staker)) => (stakerDepositShares(staker, stakerStrategyList(staker, index)) == 0);

/**
* a staker's amount of shares in a strategy (i.e. `stakerDepositShares[staker][strategy]`) should only increase when
* `depositIntoStrategy`, `depositIntoStrategyWithSignature`, or `depositBeaconChainETH` has been called
* *OR* when completing a withdrawal
*/
definition methodCanIncreaseShares(method f) returns bool =
    f.selector == sig:depositIntoStrategy(address,address,uint256).selector
    || f.selector == sig:depositIntoStrategyWithSignature(address,address,uint256,address,uint256,bytes).selector
    || f.selector == sig:withdrawSharesAsTokens(address,address,address,uint256).selector
    || f.selector == sig:addShares(address,address,uint256).selector;

/**
* a staker's amount of shares in a strategy (i.e. `stakerDepositShares[staker][strategy]`) should only decrease when
* `queueWithdrawal`, `slashShares`, or `recordBeaconChainETHBalanceUpdate` has been called
*/
definition methodCanDecreaseShares(method f) returns bool =
    f.selector == sig:removeDepositShares(address,address,uint256).selector;

rule sharesAmountsChangeOnlyWhenAppropriateFunctionsCalled(address staker, address strategy) {
    uint256 sharesBefore = stakerDepositShares(staker, strategy);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    uint256 sharesAfter = stakerDepositShares(staker, strategy);
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
    uint256 stakerDepositSharesBefore = get_stakerDepositShares(e.msg.sender, strategy);
    uint256 totalSharesBefore = totalShares(strategy);
    if (
        f.selector == sig:addShares(address, address, uint256).selector
//        || f.selector == sig:removeShares(address, address, uint256).selector
    ) {
        uint256 totalSharesAfter = totalShares(strategy);
        assert(totalSharesAfter == totalSharesBefore, "total shares changed unexpectedly");
    } else {
        uint256 stakerDepositSharesAfter = get_stakerDepositShares(e.msg.sender, strategy);
        uint256 totalSharesAfter = totalShares(strategy);
        assert(stakerDepositSharesAfter - stakerDepositSharesBefore == totalSharesAfter - totalSharesBefore, "diffs don't match");
    }
}

/**
 * Verifies that ERC20 tokens are transferred out of the account only of the msg.sender.
 * Called 'safeApprovalUse' since approval-related vulnerabilities in general allow a caller to transfer tokens out of a different account.
 * This behavior is not always unsafe, but since we don't ever use it (at present) we can do a blanket-check against it.
 */
rule safeApprovalUse(address user) {
    address token;
    uint256 tokenBalanceBefore = balanceOfCVL(token, user);
    method f;
    env e;
    // special case logic, to handle an edge case
    if (f.selector == sig:withdrawSharesAsTokens(address,address,address,uint256).selector) {
        address recipient;
        address strategy;
        uint256 shares;
        // filter out case where the 'user' is the strategy itself
        require(user != strategy);
        uint256 sharesInTokens = sharesToUnderlyingView(strategy, shares);
        require tokenBalanceBefore + sharesInTokens <= max_uint256; // Require no overflows on balances as valid env assumption
        withdrawSharesAsTokens(e, recipient, strategy, token, shares);
    } else if (f.selector == sig:burnShares(address).selector) {
        address strategy;
        require user != strategy;
        // Using simple overflow check does not work since we have token (the underlying) and shares which might be worth more than one token
        uint256 burnableShares = getBurnableShares(strategy);
        uint256 burnableInTokens = sharesToUnderlyingView(strategy, burnableShares);
        require burnableInTokens + tokenBalanceBefore <= max_uint256; // Require no overflows on balances as valid env assumption
        burnShares(e, strategy); // Anybody can burn strategy's slashed shares
    } else if (f.selector == sig:depositIntoStrategy(address,address,uint256).selector) {
        address strategy;
        uint256 amount;
        require balanceOfCVL(token, strategy) + amount <= max_uint256; // Require no overflows on balances as valid env assumption
        depositIntoStrategy(e, strategy, token, amount);
    } else if (f.selector == sig:depositIntoStrategyWithSignature(address,address,uint256,address,uint256,bytes).selector) {
        address strategy;
        uint256 amount;
        address staker;
        uint256 expiry;
        bytes signature;
        require balanceOfCVL(token, strategy) + amount <= max_uint256; // Require no overflows on balances as valid env assumption
        depositIntoStrategyWithSignature(e, strategy, token, amount, staker, expiry, signature);
    }
    // otherwise just perform an arbitrary function call 
    else {
        calldataarg args;
        f(e,args);
    }
    uint256 tokenBalanceAfter = balanceOfCVL(token, user);
    if (tokenBalanceAfter < tokenBalanceBefore) {
        assert(e.msg.sender == user, "unsafeApprovalUse?");
    }
    assert true;
}