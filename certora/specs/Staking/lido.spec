methods {
    // Lido
    function _.getTotalPooledEther() external => NONDET; // expect (uint256);
    function _.getTotalShares() external => NONDET; // expect (uint256);
    function _.submit(address _referral) external => HAVOC_ECF; // payable, expect (uint256);

    // may be shared with other contracts XXX
    function _.nonces(address _user) external => NONDET; // expect (uint256);
    function _.DOMAIN_SEPARATOR() external => NONDET; // expect (bytes32);

    // Lido Withdrawal Queue
    function _.MAX_STETH_WITHDRAWAL_AMOUNT() external => NONDET; // expect (uint256);
    function _.MIN_STETH_WITHDRAWAL_AMOUNT() external => NONDET; // expect (uint256);
    function _.requestWithdrawals(uint256[] /* calldata */ _amount, address _depositor) external => HAVOC_ECF; // expect (uint256[] memory);
    function _.claimWithdrawals(uint256[] /* calldata */ _requestIds, uint256[] /* calldata */ _hints) external => HAVOC_ECF; // expect void;
}