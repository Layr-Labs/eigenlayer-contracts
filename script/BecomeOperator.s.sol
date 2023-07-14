// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./EigenLayerParser.sol";

contract BecomeOperator is Script, DSTest, EigenLayerParser {
    //performs basic deployment before each test
    function run() external {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: msg.sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        parseEigenLayerParams();
        vm.broadcast(msg.sender);
        delegation.registerAsOperator(operatorDetails);
    }
}
