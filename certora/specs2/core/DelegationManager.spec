
methods {
    //// External Calls
	// external calls to DelegationManager 
    function undelegate(address) external;
    function decreaseDelegatedShares(address,address[],uint256[]) external;
	function increaseDelegatedShares(address,address,uint256) external;

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.slasher() external => DISPATCHER(true);
	function _.deposit(address,uint256) external => DISPATCHER(true);
	function _.withdraw(address,address,uint256) external => DISPATCHER(true);

	// external calls to EigenPodManager
	function _.withdrawBeaconChainETH(address,address,uint256) external => DISPATCHER(true);
	
    // external calls to EigenPod
	function _.withdrawBeaconChainETH(address,uint256) external => DISPATCHER(true);
    
    // external calls to PauserRegistry
    function _.pauser() external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);
	
    //// Harnessed Functions
    // Harnessed calls
    function decreaseDelegatedShares(address,address,address,uint256,uint256) external;
    // Harmessed getters
    function get_operatorShares(address,address) external returns (uint256) envfree;

    //// Summarized Functions
    function _._delegationReceivedHook(address,address,address[] memory, uint256[] memory) internal => NONDET;
    function _._delegationWithdrawnHook(address,address,address[]memory, uint256[] memory) internal => NONDET;

    //envfree functions
    function isDelegated(address staker) external returns (bool) envfree;
    function isNotDelegated(address staker) external returns (bool) envfree;
    function isOperator(address operator) external returns (bool) envfree;
    function delegatedTo(address staker) external returns (address) envfree;
    function delegationTerms(address operator) external returns (address) envfree;
    function operatorShares(address operator, address strategy) external returns (uint256) envfree;
    function owner() external returns (address) envfree;
    function strategyManager() external returns (address) envfree;
}

/*
LEGAL STATE TRANSITIONS:
1)
FROM not delegated -- defined as delegatedTo(staker) == address(0), likewise returned by isNotDelegated(staker)--
AND not registered as an operator -- defined as isOperator(operator) == false, or equivalently, delegationTerms(operator) == 0,
TO delegated but not an operator
in this case, the end state is that:
isOperator(staker) == false,
delegatedTo(staker) != staker && delegatedTo(staker) != 0,
and isDelegated(staker) == true (redundant with above)
-only allowed when calling `delegateTo` or `delegateToBySignature`

2)
FROM not delegated AND not registered as an operator
TO an operator
in this case, the end state is that:
isOperator(staker) == true,
delegatedTo(staker) == staker,
and isDelegated(staker) == true (redundant with above)
-only allowed when calling `registerAsOperator`

3)
FROM not registered as an operator AND delegated
TO not delegated (and still *not* registered as an operator)
in this case, the end state is that:
isOperator(staker) == false,
delegatedTo(staker) == 0,
and isDelegated(staker) == false (redundant with above)

ILLEGAL STATE TRANSITIONS:
A)
FROM registered as an operator
TO not registered as an operator

B) 
FROM registered as an operator (implies they are also delegated to themselves)
TO not delegated to themselves

C)
FROM delegated to an operator
TO delegated to another operator
(without undelegating in-between)

FORBIDDEN STATES:
-an address cannot be simultaneously (classified as an operator) and (not delegated to themselves)
-an address cannot be (delegated to themselves) and (not classified as an operator)
Combining the above, an address can be (classified as an operator) *iff* they are (delegated to themselves).
The exception is the zero address, since by default an address is 'delegated to the zero address' when they are not delegated at all
*/
//definition notDelegated -- defined as delegatedTo(staker) == address(0), likewise returned by isNotDelegated(staker)--

// verify that anyone who is registered as an operator is also always delegated to themselves
invariant operatorsAlwaysDelegatedToSelf(address operator)
    isOperator(operator) <=> delegatedTo(operator) == operator
    { preserved {
        require operator != 0;
    } }

// verify that once registered as an operator, a person cannot 'unregister' from being an operator
// proving this rule in concert with 'operatorsAlwaysDelegatedToSelf' proves that an operator can never change their delegation
rule operatorCannotUnregister(address operator) {
    requireInvariant operatorsAlwaysDelegatedToSelf(operator);
    // assume `operator` starts in a state of being registered as an operator
    require(isOperator(operator));
    // perform arbitrary function call
    method f;
    env e;
    calldataarg arg;
    f(e,arg);
    // verify that `operator` is still registered as an operator
    assert(isOperator(operator), "operator was able to deregister!");
}

// verifies that in order for an address to change who they are delegated to, `undelegate` must be called
rule cannotChangeDelegationWithoutUndelegating(address staker) {
    // assume the staker is delegated to begin with
    require(isDelegated(staker));
    address delegatedToBefore = delegatedTo(staker);
    // perform arbitrary function call
    method f;
    env e;
    // the only way the staker can become undelegated is if `undelegate` is called
    if (f.selector == sig:undelegate(address).selector) {
        address toUndelegate;
        undelegate(e, toUndelegate);
        // either the `strategyManager` called `undelegate` with the argument `staker` (in which can the staker is now undelegated)
        if (e.msg.sender == strategyManager() && toUndelegate == staker) {
            assert (delegatedTo(staker) == 0, "undelegation did not result in delegation to zero address");
        // or the staker's delegation should have remained the same
        } else {
            address delegatedToAfter = delegatedTo(staker);
            assert (delegatedToAfter == delegatedToBefore, "delegation changed without undelegating -- problem in undelegate permissions?");
        }
    } else {
        calldataarg arg;
        f(e,arg);
        address delegatedToAfter = delegatedTo(staker);
        assert (delegatedToAfter == delegatedToBefore, "delegation changed without undelegating");
    }
}

// verifies that an undelegated address can only delegate when calling `delegateTo`, `delegateToBySignature` or `registerAsOperator`
rule canOnlyDelegateWithSpecificFunctions(address staker) {
    requireInvariant operatorsAlwaysDelegatedToSelf(staker);
    // assume the staker begins as undelegated
    require(isNotDelegated(staker));
    // perform arbitrary function call
    method f;
    env e;
    if (f.selector == sig:delegateTo(address).selector) {
        address operator;
        delegateTo(e, operator);
        // we check against operator being the zero address here, since we view being delegated to the zero address as *not* being delegated
        if (e.msg.sender == staker && isOperator(operator) && operator != 0) {
            assert (isDelegated(staker) && delegatedTo(staker) == operator, "failure in delegateTo");
        } else {
            assert (isNotDelegated(staker), "staker delegated to inappropriate address?");
        }
    } else if (f.selector == sig:delegateToBySignature(address, address, uint256, bytes).selector) {
        address toDelegateFrom;
        address operator;
        uint256 expiry;
        bytes signature;
        delegateToBySignature(e, toDelegateFrom, operator, expiry, signature);
        // TODO: this check could be stricter! need to filter when the block timestamp is appropriate for expiry and signature is valid
        assert (isNotDelegated(staker) || delegatedTo(staker) == operator, "delegateToBySignature bug?");
    } else if (f.selector == sig:registerAsOperator(address).selector) {
        address delegationTerms;
        registerAsOperator(e, delegationTerms);
        if (e.msg.sender == staker && delegationTerms != 0) {
            assert (isOperator(staker));
        } else {
            assert(isNotDelegated(staker));
        }
    } else {
        calldataarg arg;
        f(e,arg);
        assert (isNotDelegated(staker), "staker became delegated through inappropriate function call");
    }
}

/*
rule batchEquivalence {
    env e;
    storage initial = lastStorage;
    address staker;
    address strategy1;
    address strategy2;
    uint256 share1;
    uint256 share2;

    mathint _operatorSharesStrategy1 = get_operatorShares(staker, strategy1);
    mathint _operatorSharesStrategy2 = get_operatorShares(staker, strategy2);

    decreaseDelegatedShares(e,staker,strategy1,strategy2,share1,share2);

    mathint operatorSharesStrategy1_batch = get_operatorShares(staker, strategy1);
    mathint operatorSharesStrategy2_batch = get_operatorShares(staker, strategy2);

    decreaseDelegatedShares(e,staker,strategy1,share1) at initial;
    decreaseDelegatedShares(e,staker,strategy2,share2);

    mathint operatorSharesStrategy1_single = get_operatorShares(staker, strategy1);
    mathint operatorSharesStrategy2_single = get_operatorShares(staker, strategy2);

    assert operatorSharesStrategy1_single == operatorSharesStrategy1_batch 
        && operatorSharesStrategy2_single == operatorSharesStrategy2_batch, 
        "operatorShares must be affected in the same way";
}
*/
/*
invariant zeroAddrHasNoShares(address strategy)
    get_operatorShares(0,strategy) == 0
*/