methods {
    function _.exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) external => HAVOC_ECF; // expect (uint256);
    function _.exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) external => HAVOC_ECF; // expect (uint256);
    function _.get_virtual_price() external => NONDET; // expect (uint256);
    function _.get_dy(uint256 i, iunt256 j, uint256 dx) external => NONDET; // expect (uint256);
}