import "../problems.spec";
import "../unresolved.spec";
import "../optimizations.spec";

rule check_builtin_assertions(method f)
    filtered { f -> f.contract == currentContract }
{
    env e;
    calldataarg arg;
    f(e, arg);
    assert true;
}
