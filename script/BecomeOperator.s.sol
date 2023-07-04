// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./EigenLayerParser.sol";

contract BecomeOperator is Script, DSTest, EigenLayerParser {
    //performs basic deployment before each test
    function run() external {
        parseEigenLayerParams();
        vm.broadcast(msg.sender);
        delegation.registerAsOperator(IDelegationTerms(msg.sender));
    }
}



