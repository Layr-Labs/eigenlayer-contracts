import "setup.spec";

rule noOneCanMakeMyCallRevert(method f, method g) filtered 
{ 
	f -> !f.isView && !isIgnoredMethod(f),
	g -> !g.isView && !isPrivileged(g) && !isIgnoredMethod(g)
} 
{
	require f.contract == g.contract;
	env e1;
	calldataarg args; 
	storage initialStorage = lastStorage;

	f(e1, args); // didn't revert before
	
	env e2;
	calldataarg args2;
	require e2.msg.sender != e1.msg.sender;

	g(e2, args2) at initialStorage;
	f@withrevert(e1, args);
	bool succeeded = !lastReverted;
	assert succeeded;	// didn't revert after g()
}

rule privilegedOperation(method f, address privileged) 
	filtered { f -> !isIgnoredMethod(f) }
{
	env e1;
	calldataarg arg;
	require e1.msg.sender == privileged;

	storage initialStorage = lastStorage;
	f@withrevert(e1, arg); // privileged msg.sender succeeds executing candidate privileged operation.
	bool firstSucceeded = !lastReverted;

	env e2;
	calldataarg arg2;
	require e2.msg.sender != privileged;
	f@withrevert(e2, arg2) at initialStorage; // unprivileged msg.sender should not success
	bool secondSucceeded = !lastReverted;
	bool callableByMultipleSenders = (firstSucceeded && secondSucceeded);
		//each method can either be callable by multiple senders or be marked privileged
		//this ensures that (methods not callable by multiple senders) is subset of (methods marked privileged)
	satisfy callableByMultipleSenders || isPrivileged(f);
		
		//if a method is marked privileged, it cannot be callable by multiple senders
		//this ensures that (methods marked privileged) is subset of (methods not callable by multiple senders)
	assert isPrivileged(f) => !callableByMultipleSenders; 
}

