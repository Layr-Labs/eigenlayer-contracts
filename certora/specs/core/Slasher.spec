
methods {
	//// External Calls
	// external calls to DelegationManager 
    undelegate(address) => DISPATCHER(true)
    isDelegated(address) returns (bool) => DISPATCHER(true)
    delegatedTo(address) returns (address) => DISPATCHER(true)
	decreaseDelegatedShares(address,address[],uint256[]) => DISPATCHER(true)
	increaseDelegatedShares(address,address,uint256) => DISPATCHER(true)
	_delegationReceivedHook(address,address,address[],uint256[]) => NONDET
    _delegationWithdrawnHook(address,address,address[],uint256[]) => NONDET

	// external calls to Slasher
    isFrozen(address) returns (bool) envfree 
	canWithdraw(address,uint32,uint256) returns (bool) 

	// external calls to StrategyManager
    getDeposits(address) returns (address[],uint256[]) => DISPATCHER(true)
    slasher() returns (address) => DISPATCHER(true)
	deposit(address,uint256) returns (uint256) => DISPATCHER(true)
	withdraw(address,address,uint256) => DISPATCHER(true)

	// external calls to EigenPodManager
	withdrawBeaconChainETH(address,address,uint256) => DISPATCHER(true)
	
    // external calls to EigenPod
	withdrawBeaconChainETH(address,uint256) => DISPATCHER(true)
    
    // external calls to IDelegationTerms
    onDelegationWithdrawn(address,address[],uint256[]) => CONSTANT
    onDelegationReceived(address,address[],uint256[]) => CONSTANT
    
    // external calls to PauserRegistry
    pauser() returns (address) => DISPATCHER(true)
	unpauser() returns (address) => DISPATCHER(true)
	
    //// Harnessed Functions
    // Harnessed calls
    // Harnessed getters
	get_is_operator(address) returns (bool) envfree
	get_is_delegated(address) returns (bool) envfree
	get_list_exists(address) returns (bool) envfree
	get_next_node_exists(address, uint256) returns (bool) envfree
	get_next_node(address, uint256) returns (uint256) envfree
	get_previous_node_exists(address, uint256) returns (bool) envfree
	get_previous_node(address, uint256) returns (uint256) envfree
	get_node_exists(address, address) returns (bool) envfree
	get_list_head(address) returns (uint256) envfree
	get_lastest_update_block_at_node(address, uint256) returns (uint256) envfree
	get_lastest_update_block_at_head(address) returns (uint256) envfree
	get_linked_list_entry(address operator, uint256 node, bool direction) returns (uint256) envfree

	// nodeDoesExist(address operator, uint256 node) returns (bool) envfree
	nodeIsWellLinked(address operator, uint256 node) returns (bool) envfree
	
	//// Normal Functions
	owner() returns(address) envfree
	contractCanSlashOperatorUntil(address, address) returns (uint32) envfree
	paused(uint8) returns (bool) envfree
}

// uses that _HEAD = 0. Similar to StructuredLinkedList.nodeExists but slightly better defined
definition nodeDoesExist(address operator, uint256 node) returns bool =
	(get_next_node(operator, node) == 0 && get_previous_node(operator, node) == 0) 
		=> (get_next_node(operator, 0) == node && get_previous_node(operator, 0) == node);

definition nodeIsWellLinked(address operator, uint256 node) returns bool =
	// node is not linked to itself
	get_previous_node(operator, node) != node && get_next_node(operator, node) != node
	// node is the previous node's next node and the next node's previous node
	&& get_linked_list_entry(operator, get_previous_node(operator, node), true) == node
	&& get_linked_list_entry(operator, get_next_node(operator, node), false) == node;

/*
TODO: sort out if `isFrozen` can also be marked as envfree -- currently this is failing with the error
could not type expression "isFrozen(staker)", message: Could not find an overloading of method isFrozen that matches
the given arguments: address. Method is not envfree; did you forget to provide the environment as the first function argument?
rule cantBeUnfrozen(method f) {
	address staker;

	bool _frozen = isFrozen(staker);
	require _frozen;

	env e; calldataarg args;
	require e.msg.sender != owner();
	f(e,args);

	bool frozen_ = isFrozen(staker);
	assert frozen_, "frozen stakers must stay frozen";
}

/*
verifies that `contractCanSlashOperatorUntil[operator][contractAddress]` only changes when either:
the `operator` themselves calls `allowToSlash`
or
the `contractAddress` calls `recordLastStakeUpdateAndRevokeSlashingAbility`
*/
rule canOnlyChangecontractCanSlashOperatorUntilWithSpecificFunctions(address operator, address contractAddress) {
	uint256 valueBefore = contractCanSlashOperatorUntil(operator, contractAddress);
    // perform arbitrary function call
    method f;
    env e;
    if (f.selector == recordLastStakeUpdateAndRevokeSlashingAbility(address, uint32).selector) {
        address operator2;
		uint32 serveUntil;
        recordLastStakeUpdateAndRevokeSlashingAbility(e, operator2, serveUntil);
		uint256 valueAfter = contractCanSlashOperatorUntil(operator, contractAddress);
        if (e.msg.sender == contractAddress && operator2 == operator/* TODO: proper check */) {
			/* TODO: proper check */
            assert (true, "failure in recordLastStakeUpdateAndRevokeSlashingAbility");
        } else {
            assert (valueBefore == valueAfter, "bad permissions on recordLastStakeUpdateAndRevokeSlashingAbility?");
        }
	} else if (f.selector == optIntoSlashing(address).selector) {
		address arbitraryContract;
		optIntoSlashing(e, arbitraryContract);
		uint256 valueAfter = contractCanSlashOperatorUntil(operator, contractAddress);
		// uses that the `PAUSED_OPT_INTO_SLASHING` index is 0, as an input to the `paused` function
		if (e.msg.sender == operator && arbitraryContract == contractAddress && get_is_operator(operator) && !paused(0)) {
			// uses that `MAX_CAN_SLASH_UNTIL` is equal to max_uint32
			assert(valueAfter == max_uint32, "MAX_CAN_SLASH_UNTIL different than max_uint32?");
		} else {
            assert(valueBefore == valueAfter, "bad permissions on optIntoSlashing?");
		}
	} else {
		calldataarg arg;
		f(e, arg);
		uint256 valueAfter = contractCanSlashOperatorUntil(operator, contractAddress);
        assert(valueBefore == valueAfter, "bondedAfter value changed when it shouldn't have!");
	}
}

/*
checks that the entry in the linked list _whitelistedContractDetails[operator] with the **smallest** value of 'latestUpdateBlock'
is always at the 'HEAD' position in the linked list
*/
/* TODO: modify rule so it works! This seems to make too broad assumptions about initial state (i.e. isn't strict enough)
invariant listHeadHasSmallestValueOfLatestUpdateBlock(address operator, uint256 node)
	(
	get_list_exists(operator) && get_next_node_exists(operator, get_list_head(operator)) => 
		get_lastest_update_block_at_head(operator) <= get_lastest_update_block_at_node(operator, get_next_node(operator, get_list_head(operator)))
	)
*/

/*
TODO: rule doesn't pass.
key properties seem to be that
1) `StructuredLinkedList._createLink` creates only two-way links
2) `StructuredLinkedList.remove` removes both links from a node, and stiches together its existing links (which it breaks)
3) `StructuredLinkedList._insert` similarly inserts a new node 'between' nodes, ensuring that the new node is well-linked
*/
invariant consistentListStructure(address operator, uint256 node1)
	(
	// either node1 doesn't exist
	!nodeDoesExist(operator, node1)
	// or node1 is consistently two-way linked
	||
	nodeIsWellLinked(operator, node1)
	)

/* TODO: assess if this rule is salvageable. seems to have poor storage assumptions due to the way 'node existence' is defined
rule cannotAddSameContractTwice(address operator, address contractAddress) {
	bool nodeExistsBefore = get_node_exists(operator, contractAddress);
	env e;
	uint32 serveUntil;
	recordFirstStakeUpdate(e, operator, serveUntil);
	if (nodeExistsBefore) {
		bool callReverted = lastReverted;
		assert (callReverted, "recordFirstStakeUpdate didn't revert!");
	} else {
		bool nodeExistsAfter = get_node_exists(operator, contractAddress);
		if (e.msg.sender == contractAddress) {
			assert(nodeExistsAfter, "node not added correctly");
		} else {
			assert(!nodeExistsAfter, "node added incorrectly");
		}
	}
}
*/
/*
## Slashing

- slashing happens if and only if a provably malicious action by an operator took place
- operator may be slashed only if allowToSlash() for that particular contract was called
- slashing cannot happen after contractCanSlashOperatorUntil[operator][contractAddress] timestamp
- contractCanSlashOperatorUntil[operator][contractAddress] changed  => allowToSlash() or recordLastStakeUpdateAndRevokeSlashingAbility() was called
- recordLastStakeUpdateAndRevokeSlashingAbility() should only be callable when contractCanSlashOperatorUntil[operator][contractAddress] == MAX_CAN_SLASH_UNTIL, and only by the contractAddress
- Any contractAddress for which contractCanSlashOperatorUntil[operator][contractAddress] > current time can call freezeOperator(operator).
- frozen operator cannot make deposits/withdrawals, cannot complete queued withdrawals
- slashing and unfreezing is performed by the StrategyManager contract owner (is it permanent or configurable?)
- frozenStatus[operator] changed => freezeOperator() or resetFrozenStatus() were called
*/