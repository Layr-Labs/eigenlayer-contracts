methods {
    // avoids linking messages upon upgradeToAndCall
    function _._upgradeToAndCall(address,bytes,bool) external => HAVOC_ECF;
    function _._upgradeToAndCallUUPS(address,bytes,bool) external => HAVOC_ECF;
    // view function
    function _.proxiableUUID() external => NONDET; // expect bytes32
}
