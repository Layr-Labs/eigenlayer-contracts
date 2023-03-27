
methods {
    //// External Calls
	// external calls to DelegationManager 
    undelegate(address) 
    decreaseDelegatedShares(address,address[],uint256[]) 
	increaseDelegatedShares(address,address,uint256) 

	// external calls to Slasher
    isFrozen(address) returns (bool) => DISPATCHER(true)
	canWithdraw(address,uint32,uint256) returns (bool) => DISPATCHER(true)

	// external calls to StrategyManager
    getDeposits(address) returns (address[],uint256[]) => DISPATCHER(true)
    slasher() returns (address) => DISPATCHER(true)
	deposit(address,uint256) returns (uint256) => DISPATCHER(true)
	withdraw(address,address,uint256) => DISPATCHER(true)

	// external calls to EigenPodManager
	withdrawBeaconChainETH(address,address,uint256) => DISPATCHER(true)
	
    // external calls to EigenPod
	withdrawBeaconChainETH(address,uint256) => DISPATCHER(true)
    
    // external calls to PauserRegistry
    pauser() returns (address) => DISPATCHER(true)
	unpauser() returns (address) => DISPATCHER(true)

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    isValidSignature(bytes32, bytes) returns (bytes4) => DISPATCHER(true)
	
    //// Harnessed Functions
    // Harnessed calls
    decreaseDelegatedShares(address,address,address,uint256,uint256)
    // Harmessed getters
    get_operatorShares(address,address) returns(uint256) envfree

    //// Summarized Functions
    _delegationReceivedHook(address,address,address[],uint256[]) => NONDET
    _delegationWithdrawnHook(address,address,address[],uint256[]) => NONDET

    //envfree functions
    isDelegated(address staker) returns (bool) envfree
    isNotDelegated(address staker) returns (bool) envfree
    isOperator(address operator) returns (bool) envfree
    delegatedTo(address staker) returns (address) envfree
    delegationTerms(address operator) returns (address) envfree
    operatorShares(address operator, address strategy) returns (uint256) envfree
    owner() returns (address) envfree
    strategyManager() returns (address) envfree
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
    if (f.selector == undelegate(address).selector) {
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
    if (f.selector == delegateTo(address).selector) {
        address operator;
        delegateTo(e, operator);
        // we check against operator being the zero address here, since we view being delegated to the zero address as *not* being delegated
        if (e.msg.sender == staker && isOperator(operator) && operator != 0) {
            assert (isDelegated(staker) && delegatedTo(staker) == operator, "failure in delegateTo");
        } else {
            assert (isNotDelegated(staker), "staker delegated to inappropriate address?");
        }
    } else if (f.selector == delegateToBySignature(address, address, uint256, bytes).selector) {
        address toDelegateFrom;
        address operator;
        uint256 expiry;
        bytes signature;
        delegateToBySignature(e, toDelegateFrom, operator, expiry, signature);
        // TODO: this check could be stricter! need to filter when the block timestamp is appropriate for expiry and signature is valid
        assert (isNotDelegated(staker) || delegatedTo(staker) == operator, "delegateToBySignature bug?");
    } else if (f.selector == registerAsOperator(address).selector) {
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