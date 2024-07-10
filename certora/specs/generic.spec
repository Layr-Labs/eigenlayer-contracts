/*
This rule find which functions are privileged.
A function is privileged if there is only one address that can call it.

The rules finds this by finding which functions can be called by two different users.
*/
rule privilegedOperation(method f, address privileged) {
	env e1;
	calldataarg arg;
	require e1.msg.sender == privileged;

	storage initialStorage = lastStorage;
	f@withrevert(e1, arg); // privileged succeeds executing candidate privileged operation.
	bool firstSucceeded = !lastReverted;

	env e2;
	calldataarg arg2;
	require e2.msg.sender != privileged;
	f@withrevert(e2, arg2) at initialStorage; // unprivileged
	bool secondSucceeded = !lastReverted;

	assert  !(firstSucceeded && secondSucceeded);
}

rule timeoutChecker(method f) {
	storage before = lastStorage;
	env e; calldataarg arg;
	f(e,arg);
	assert before == lastStorage;
}

/*
This rule find which functions that can be called, may fail due to someone else calling a function right before.

This is n expensive rule - might fail on the demo site on big contracts
*/
rule simpleFrontRunning(method f, address privileged) filtered { f-> !f.isView } {
	env e1;
	calldataarg arg;
	require e1.msg.sender == privileged;  
	storage initialStorage = lastStorage;
	f(e1, arg); 
	bool firstSucceeded = !lastReverted;
	env e2;
	calldataarg arg2;
	require e2.msg.sender != e1.msg.sender;
	f(e2, arg2) at initialStorage; 
	f@withrevert(e1, arg);
	bool succeeded = !lastReverted;
	assert succeeded;
}

rule noRevert(method f) {
	env e;
	calldataarg arg;
	require e.msg.value == 0; 
	f@withrevert(e, arg); 
	assert !lastReverted;
}


rule alwaysRevert(method f) {
	env e;
	calldataarg arg;
	f@withrevert(e, arg); 
	assert lastReverted;
}

/* failing CALL should lead to a revert */
ghost bool saw_failing_call;

hook CALL(uint g, address addr, uint value, uint argsOffset, uint argsLength, uint retOffset, uint retLength) uint rc {
	saw_failing_call = saw_failing_call || rc == 0;
}

rule failing_CALL_leads_to_revert(method f) {
	saw_failing_call = false;
	env e;
	calldataarg arg;
	f@withrevert(e, arg);
	bool reverted = lastReverted;
	assert saw_failing_call => reverted;
}

// All usages
use builtin rule sanity;
use builtin rule hasDelegateCalls;
use builtin rule msgValueInLoopRule;
use builtin rule viewReentrancy;

/**

// Integrate rules from generic.spec in importing specs like this:

use builtin rule sanity filtered { f -> f.contract == currentContract }
use builtin rule hasDelegateCalls filtered { f -> f.contract == currentContract }
use builtin rule msgValueInLoopRule;
use builtin rule viewReentrancy;
use rule privilegedOperation filtered { f -> f.contract == currentContract }
use rule timeoutChecker filtered { f -> f.contract == currentContract }
use rule simpleFrontRunning filtered { f -> f.contract == currentContract }
use rule noRevert filtered { f -> f.contract == currentContract }
use rule alwaysRevert filtered { f -> f.contract == currentContract }
use rule failing_CALL_leads_to_revert filtered { f -> f.contract == currentContract }
 **/
