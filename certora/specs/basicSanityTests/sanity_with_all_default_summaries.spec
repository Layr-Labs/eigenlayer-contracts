import "../ERC20/erc20cvl.spec";

import "../unresolved.spec";

import "../generic.spec"; // pick additional rules from here

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

rule initChecker() {
    env e;
    calldataarg args;
    initialize(e,args);
    satisfy true;
}