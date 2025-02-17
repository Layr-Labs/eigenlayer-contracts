// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../../src/contracts/libraries/StructuredLinkedList.sol";

/**
 * @title StructuredLinkedList
 * @author Vittorio Minacori (https://github.com/vittominacori)
 * @dev An utility library for using sorted linked list data structures in your Solidity project.
 * @notice Adapted from https://github.com/vittominacori/solidity-linked-list/blob/master/contracts/StructuredLinkedList.sol
 */
contract StructuredLinkedListHarness {

    using StructuredLinkedList for StructuredLinkedList.List;

    StructuredLinkedList.List public listStorage;

    /// Getters with single value returns.
    function getAdjacentStrict(uint256 _node, bool _direction) public view returns (uint256 adj) {
        if (!nodeExists(_node)) {
            revert();
        } else {
            adj = listStorage.list[_node][_direction];
        }
    }

    // Generic setters that cover all use cases
    function insert(uint256 _node, uint256 _new, bool _dir) public {
        if (_dir) {
            require(listStorage.insertAfter(_node, _new), "_insert failed");
        } else {
            require(listStorage.insertBefore(_node, _new), "_insert failed");
        }
        
    }

    function remove(uint256 _node) public {
        require(listStorage.remove(_node) > 0, "remove failed");
    }

    /// View functions made public.

    function listExists() public view returns (bool) {
        return listStorage.listExists();
    }

    function nodeExists(uint256 _node) public view returns (bool) {
        return listStorage.nodeExists(_node);
    }

    function sizeOf() public view returns (uint256) {
        return listStorage.sizeOf();
    }

    function getHead() public view returns (uint256) {
        return listStorage.getHead();
    }

    function getNode(uint256 _node) public view returns (bool, uint256, uint256) {
        return listStorage.getNode(_node);
    }

    function getAdjacent(uint256 _node, bool _direction) public view returns (bool, uint256) {
        return listStorage.getAdjacent(_node, _direction);
    }

    function getNextNode(uint256 _node) public view returns (bool, uint256) {
        return listStorage.getNextNode(_node);
    }

    function getPreviousNode(uint256 _node) public view returns (bool, uint256) {
        return listStorage.getPreviousNode(_node);
    }

    /// Setters made public that don't add extra functionality
    /* commented out as some invariants require a specific preserve block for each function and creating one for each of these functions is redundant
    function insertAfter(uint256 _node, uint256 _new) public returns (bool) {
        return listStorage.insertAfter(_node, _new);
    }

    function insertBefore(uint256 _node, uint256 _new) public returns (bool) {
        return listStorage.insertBefore(_node, _new);
    }

    function pushFront(uint256 _node) internal returns (bool) {
        return listStorage.pushFront(_node);
    }

    function pushBack(uint256 _node) public returns (bool) {
        return listStorage.pushBack(_node);
    }

    function popFront() public returns (uint256) {
        return listStorage.popFront();
    }

    function popBack() public returns (uint256) {
        return listStorage.popBack();
    }
    */

}