import "../optimizations.spec";

using AllocationManager as AllocationManager;
methods {
    function AllocationManager.DEALLOCATION_DELAY() external returns(uint32) envfree;

    // external calls to AVSRegistrar.  Note that the source does not have a proper implementation, the one available always reverts
    function _.registerOperator(address,address,uint32[],bytes) external => DISPATCHER(true);
    function _.deregisterOperator(address,address,uint32[]) external => DISPATCHER(true);
    function _.supportsAVS(address) external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.underlyingToken() external => DISPATCHER(true);
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address,uint256) external => DISPATCHER(true);

    function OperatorSetLib.key(OperatorSetLib.OperatorSet memory os) internal returns (bytes32) => returnOperatorSetKey(os); // expect (bytes32); // return unique bytes32 that is not zero.

    // internal math summary to avoid overflows from the tool;
    // function AllocationManager._addInt128(uint64 a, int128 b) internal returns (uint64) => cvlAddInt128(a, b);
}

function cvlAddInt128(uint64 a, int128 b) returns uint64 {
    require(b >= 0 || to_mathint(a) > to_mathint(-b)); // Prevent underflow
    require(b <= 0 || a < to_mathint(max_uint64) -to_mathint(b)); // Prevent overflow
    return require_uint64(to_mathint(a) + to_mathint(b));
}

function returnOperatorSetKey(OperatorSetLib.OperatorSet os) returns bytes32 {
    return idToKey[os.id];
}

ghost mapping(uint32 => bytes32) idToKey {
    axiom forall uint32 id1 . forall uint32 id2 . (idToKey[id1] != to_bytes32(0) && idToKey[id2] != to_bytes32(0)) &&
        (id1 != id2 => idToKey[id1] != idToKey[id2]);
}

use builtin rule sanity;
