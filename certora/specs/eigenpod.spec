import "unresolved.spec";

methods {
    // Because of create2
    function Create2.deploy(uint256, bytes32, bytes memory) internal returns (address)  => NONDET;
    
    // IEigenPod
    function _._ external => DISPATCH [
      EigenPod.initialize(address),
      EigenPod.stake(bytes, bytes, bytes32),
   ] default NONDET;

}