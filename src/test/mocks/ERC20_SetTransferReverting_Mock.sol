// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

contract ERC20_SetTransferReverting_Mock is ERC20PresetFixedSupply {
    bool public transfersRevert;

    constructor(uint initSupply, address initOwner)
        ERC20PresetFixedSupply("ERC20_SetTransferReverting_Mock", "ERC20_SetTransferReverting_Mock", initSupply, initOwner)
    {}

    function setTransfersRevert(bool _transfersRevert) public {
        transfersRevert = _transfersRevert;
    }

    function _beforeTokenTransfer(address, address, uint) internal view override {
        if (transfersRevert) {
            // revert without message
            revert();
            // revert("ERC20_SetTransferReverting_Mock._beforeTokenTransfer: transfersRevert set");
        }
    }
}
