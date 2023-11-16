// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/Staker.sol";
import "src/test/integration/GlobalRefs.sol";

contract Operator is Staker {
    // Array of addresses who are delegated to staker
    address[] public delegatedStaker;
    mapping(address => bool) public isDelegatedStaker;


    address operator = address(this);

    // Delegation Signer Private Key
    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

    constructor(GlobalRefs _globalRefs) Staker (_globalRefs){}

    // Registration Functions
    function register() public {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operatorDetails);
    }

    function _registerOperator(
        IDelegationManager.OperatorDetails memory operatorDetails
    ) internal {
        string memory emptyStringForMetadataURI;
        globalRefs.delegationManager().registerAsOperator(operatorDetails, emptyStringForMetadataURI); 

        // Checks
        assertEq(
            operatorDetails.earningsReceiver,
            globalRefs.delegationManager().earningsReceiver(operator),
            "earningsReceiver not set correctly"
        );
        assertEq(
            operatorDetails.delegationApprover,
            globalRefs.delegationManager().delegationApprover(operator),
            "delegationApprover not set correctly"
        );
        assertEq(
            operatorDetails.stakerOptOutWindowBlocks,
            globalRefs.delegationManager().stakerOptOutWindowBlocks(operator),
            "stakerOptOutWindowBlocks not set correctly"
        );
        assertEq(globalRefs.delegationManager().delegatedTo(operator), operator, "operator not delegated to self");       
    }

    // Override staker functions that operator shouldn't call
    function delegate(address operator) public override {
        revert("Operator cannot delegate since it's delegated to itself");
    }

    // Helper functions to update stakers that are delegated to this operator
    function addDelegatedStaker(address staker) public {
        if (!isDelegatedStaker[staker]) {
            isDelegatedStaker[staker] = true;
            delegatedStaker.push(staker);
        }
    }

    function removeDelegatedStaker(address staker) public {
        if (isDelegatedStaker[staker]) {
            isDelegatedStaker[staker] = false;
            for (uint256 i = 0; i < delegatedStaker.length; i++) {
                if (delegatedStaker[i] == staker) {
                    delegatedStaker[i] = delegatedStaker[delegatedStaker.length - 1];
                    delegatedStaker.pop();
                    break;
                }
            }
        }
    }
}