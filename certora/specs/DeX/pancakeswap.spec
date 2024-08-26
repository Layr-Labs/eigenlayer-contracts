methods {
    // interface IPancackeV3SwapRouter
    function _.WETH9() external => HAVOC_ECF; // expect (address); // xxx not marked as view but suspect it is...
    function _.unwrapWETH9(uint256 amountMinimum, address recipient) external => HAVOC_ECF; // payable, expect void;
    // xxx to use this, must import IPancackeV3SwapRouter
    // function _.exactInputSingle(IPancackeV3SwapRouter.ExactInputSingleParams /* calldata */ params) external => HAVOC_ECF; // payable, expect (uint256 amountOut);
}