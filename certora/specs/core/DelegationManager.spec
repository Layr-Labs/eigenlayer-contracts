
methods {
    //// External Calls
	// external calls to DelegationManager 
    function undelegate(address) external;
    function decreaseDelegatedShares(address,address,uint256) external;
	function increaseDelegatedShares(address,address,uint256) external;

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.slasher() external => DISPATCHER(true);
    function _.addShares(address,address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256, address) external => DISPATCHER(true);

	// external calls to EigenPodManager
    function _.addShares(address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256) external => DISPATCHER(true);

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);
    
    // external calls to DelayedWithdrawalRouter (from EigenPod)
    function _.createDelayedWithdrawal(address, address) external => DISPATCHER(true);

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

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);
	
    //// Harnessed Functions
    // Harnessed getters
    function get_operatorShares(address,address) external returns (uint256) envfree;
    function get_stakerDelegateableShares(address,address) external returns (uint256) envfree;

    //envfree functions
    function delegatedTo(address) external returns (address) envfree;
    function operatorDetails(address) external returns (IDelegationManager.OperatorDetails memory) envfree;
    function earningsReceiver(address) external returns (address) envfree;
    function delegationApprover(address operator) external returns (address) envfree;
    function stakerOptOutWindowBlocks(address operator) external returns (uint256) envfree;
    function operatorShares(address operator, address strategy) external returns (uint256) envfree;
    function isDelegated(address staker) external returns (bool) envfree;
    function isOperator(address operator) external returns (bool) envfree;
    function stakerNonce(address staker) external returns (uint256) envfree;
    function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) external returns (bool) envfree;
    function owner() external returns (address) envfree;
    function strategyManager() external returns (address) envfree;
    function eigenPodManager() external returns (address) envfree;
}

/*
LEGAL STATE TRANSITIONS:
1)
FROM not delegated -- defined as delegatedTo(staker) == address(0), likewise returned by !isDelegated(staker)--
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
//definition notDelegated -- defined as delegatedTo(staker) == address(0), likewise returned by !isDelegated(staker)--

// verify that anyone who is registered as an operator is also always delegated to themselves
// the zero address is an exception to this rule, since it is always "delegated to itself" but not an operator
invariant operatorsAlwaysDelegatedToSelf(address operator)
    operator != 0 => (isOperator(operator) <=> delegatedTo(operator) == operator);

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
    requireInvariant operatorsAlwaysDelegatedToSelf(staker);
    // assume the staker is delegated to begin with
    require(isDelegated(staker));
    address delegatedToBefore = delegatedTo(staker);
    // perform arbitrary function call
    method f;
    env e;
    // the only way the staker can become undelegated is an appropriate function is called
    if (f.selector == sig:undelegate(address).selector) {
        address toUndelegate;
        undelegate(e, toUndelegate);
        // either the `staker` address was an input to `undelegate` AND the caller was allowed to call the function
        if (
            (toUndelegate == staker && (delegatedToBefore != staker)) &&
            (e.msg.sender == staker || e.msg.sender == delegatedToBefore || e.msg.sender == delegationApprover(delegatedToBefore))
        ){
        assert (delegatedTo(staker) == 0, "undelegation did not result in delegation to zero address");
        // or the staker's delegation should have remained the same
        } else {
            address delegatedToAfter = delegatedTo(staker);
            assert (delegatedToAfter == delegatedToBefore, "delegation changed without undelegating -- problem in undelegate permissions?");
        }
        assert(true);
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
    require(!isDelegated(staker));
    // perform arbitrary function call
    method f;
    env e;
    if (f.selector == sig:delegateTo(address, ISignatureUtils.SignatureWithExpiry, bytes32).selector) {
        address operator;
        ISignatureUtils.SignatureWithExpiry approverSignatureAndExpiry;
        bytes32 salt;
        delegateTo(e, operator, approverSignatureAndExpiry, salt);
        // we check against operator being the zero address here, since we view being delegated to the zero address as *not* being delegated
        if (e.msg.sender == staker && isOperator(operator) && operator != 0) {
            assert (isDelegated(staker) && delegatedTo(staker) == operator, "failure in delegateTo");
        } else {
            assert (!isDelegated(staker), "staker delegated to inappropriate address?");
        }
    } else if (f.selector == sig:delegateToBySignature(address, address, ISignatureUtils.SignatureWithExpiry, ISignatureUtils.SignatureWithExpiry, bytes32).selector) {
        address toDelegateFrom;
        address operator;
        ISignatureUtils.SignatureWithExpiry stakerSignatureAndExpiry;
        ISignatureUtils.SignatureWithExpiry approverSignatureAndExpiry;
        bytes32 salt;
        delegateToBySignature(e, toDelegateFrom, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt);
        // TODO: this check could be stricter! need to filter when the block timestamp is appropriate for expiry and signature is valid
        assert (!isDelegated(staker) || delegatedTo(staker) == operator, "delegateToBySignature bug?");
    } else if (f.selector == sig:registerAsOperator(IDelegationManager.OperatorDetails, string).selector) {
        IDelegationManager.OperatorDetails operatorDetails;
        string metadataURI;
        registerAsOperator(e, operatorDetails, metadataURI);
        if (e.msg.sender == staker) {
            assert (isOperator(staker));
        } else {
            assert(!isDelegated(staker));
        }
    } else {
        calldataarg arg;
        f(e,arg);
        assert (!isDelegated(staker), "staker became delegated through inappropriate function call");
    }
}

rule sharesBecomeDelegatedWhenStakerDelegates(address operator, address staker, address strategy) {
    requireInvariant operatorsAlwaysDelegatedToSelf(operator);
    // filter out zero address (not a valid operator)
    require(operator != 0);
    // assume the staker begins as undelegated
    require(!isDelegated(staker));
    mathint stakerDelegateableSharesInStrategy = get_stakerDelegateableShares(staker, strategy);
    mathint operatorSharesBefore = get_operatorShares(operator, strategy);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg arg;
    mathint operatorSharesAfter = get_operatorShares(operator, strategy);
    if (delegatedTo(staker) == operator) {
        assert(operatorSharesAfter == operatorSharesBefore + stakerDelegateableSharesInStrategy, "operator shares did not increase appropriately");
    } else {
        assert(operatorSharesAfter == operatorSharesBefore, "operator shares changed inappropriately");
    }
}

rule sharesBecomeUndelegatedWhenStakerUndelegates(address operator, address staker, address strategy) {
    requireInvariant operatorsAlwaysDelegatedToSelf(operator);
    // filter out zero address (not a valid operator)
    require(operator != 0);
    // assume the staker begins as delegated to the operator
    require(delegatedTo(staker) == operator);
    mathint stakerDelegateableSharesInStrategy = get_stakerDelegateableShares(staker, strategy);
    mathint operatorSharesBefore = get_operatorShares(operator, strategy);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg arg;
    mathint operatorSharesAfter = get_operatorShares(operator, strategy);
    if (!isDelegated(staker)) {
        assert(operatorSharesAfter == operatorSharesBefore - stakerDelegateableSharesInStrategy, "operator shares did not decrease appropriately");
    } else {
        assert(operatorSharesAfter == operatorSharesBefore, "operator shares changed inappropriately");
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