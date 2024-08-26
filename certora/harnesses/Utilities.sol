contract Utilities {
    function havocAll() external {
        (bool success, ) = address(0xdeadbeef).call(abi.encodeWithSelector(0x12345678));
        require(success);
    }

    function justRevert() external {
        revert();
    }

    function nop() external {}
}