import "../shared.spec";

using Utilities as utils;

methods {
    // e.g. cbETH (Coinbase Wrapped Staked ETH), WBETH (Wrapped Beacon ETH)
    function _.mint(address _to, uint256 _amount) external => HAVOC_ECF; // expect void;
    function _.exchangeRate() external => NONDET; // expect (uint256 _exchangeRate);
    
    // WBETH
    function _.deposit(address referral) external with (env e) => pay_and_havoc(calledContract, e); // payable, expect void
}

