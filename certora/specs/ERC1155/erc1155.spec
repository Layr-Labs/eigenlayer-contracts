methods {
    function _.onERC1155Received(
        address operator,
        address from,
        uint256 id,
        uint256 value,
        bytes /* calldata */ data
    ) external => HAVOC_ECF; // expect (bytes4);
    function _.onERC1155BatchReceived(
        address operator,
        address from,
        uint256[] /* calldata */ ids,
        uint256[] /* calldata */ values,
        bytes /* calldata */ data
    ) external => HAVOC_ECF; // expect (bytes4);
}