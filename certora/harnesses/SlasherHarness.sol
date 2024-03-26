// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/core/Slasher.sol";

contract SlasherHarness is Slasher {

    constructor(IStrategyManager _strategyManager, IDelegationManager _delegation) Slasher(_strategyManager, _delegation) {}
    
    /// Harnessed functions
    function get_is_operator(address staker) public returns (bool) {
        return delegation.isOperator(staker);        
    }

    function get_is_delegated(address staker) public returns (bool) {
        return delegation.isDelegated(staker);        
    }


    // Linked List Functions
    function get_list_exists(address operator) public returns (bool) {
        return StructuredLinkedList.listExists(_operatorToWhitelistedContractsByUpdate[operator]);
    }

    function get_next_node_exists(address operator, uint256 node) public returns (bool) {
        (bool res, ) = StructuredLinkedList.getNextNode(_operatorToWhitelistedContractsByUpdate[operator], node);
        return res;
    }

    function get_next_node(address operator, uint256 node) public returns (uint256) {
        (, uint256 res) = StructuredLinkedList.getNextNode(_operatorToWhitelistedContractsByUpdate[operator], node);
        return res;
    }

    function get_previous_node_exists(address operator, uint256 node) public returns (bool) {
        (bool res, ) = StructuredLinkedList.getPreviousNode(_operatorToWhitelistedContractsByUpdate[operator], node);
        return res;
    }

    function get_previous_node(address operator, uint256 node) public returns (uint256) {
        (, uint256 res) = StructuredLinkedList.getPreviousNode(_operatorToWhitelistedContractsByUpdate[operator], node);
        return res;
    }

    function get_list_head(address operator) public returns (uint256) {
        return StructuredLinkedList.getHead(_operatorToWhitelistedContractsByUpdate[operator]);
    }

    function get_lastest_update_block_at_node(address operator, uint256 node) public returns (uint256) {
        return _whitelistedContractDetails[operator][_uintToAddress(node)].latestUpdateBlock;
    }

    function get_lastest_update_block_at_head(address operator) public returns (uint256) {
        return get_lastest_update_block_at_node(operator, get_list_head(operator));
    }

    function get_linked_list_entry(address operator, uint256 node, bool direction) public returns (uint256) {
        return (_operatorToWhitelistedContractsByUpdate[operator].list[node][direction]);
    }

    // // uses that _HEAD = 0. Similar to StructuredLinkedList.nodeExists but slightly better defined
    // function nodeDoesExist(address operator, uint256 node) public returns (bool) {
    //     if (get_next_node(operator, node) == 0 && get_previous_node(operator, node) == 0) {
    //         // slightly stricter check than that defined in StructuredLinkedList.nodeExists
    //         if (get_next_node(operator, 0) == node && get_previous_node(operator, 0) == node) {
    //             return true;
    //         } else {
    //             return false;
    //         }
    //     } else {
    //         return true;
    //     }
    // }

    // // uses that _PREV = false, _NEXT = true
    // function nodeIsWellLinked(address operator, uint256 node) public returns (bool) {
    //     return (
    //         // node is not linked to itself
    //         get_previous_node(operator, node) != node && get_next_node(operator, node) != node
    //         &&
    //         // node is the previous node's next node and the next node's previous node
    //         get_linked_list_entry(operator, get_previous_node(operator, node), true) == node && get_linked_list_entry(operator, get_next_node(operator, node), false) == node
    //     );
    // }
}
