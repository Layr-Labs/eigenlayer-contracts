# StructuredLinkedList
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/contracts/libraries/StructuredLinkedList.sol)

**Author:**
Vittorio Minacori (https://github.com/vittominacori)

Adapted from https://github.com/vittominacori/solidity-linked-list/blob/master/contracts/StructuredLinkedList.sol

*An utility library for using sorted linked list data structures in your Solidity project.*


## State Variables
### _NULL

```solidity
uint256 private constant _NULL = 0;
```


### _HEAD

```solidity
uint256 private constant _HEAD = 0;
```


### _PREV

```solidity
bool private constant _PREV = false;
```


### _NEXT

```solidity
bool private constant _NEXT = true;
```


## Functions
### listExists

*Checks if the list exists*


```solidity
function listExists(List storage self) internal view returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if list exists, false otherwise|


### nodeExists

*Checks if the node exists*


```solidity
function nodeExists(List storage self, uint256 _node) internal view returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|a node to search for|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if node exists, false otherwise|


### sizeOf

*Returns the number of elements in the list*


```solidity
function sizeOf(List storage self) internal view returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256|


### getHead

*Gets the head of the list*


```solidity
function getHead(List storage self) internal view returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 the head of the list|


### getNode

*Returns the links of a node as a tuple*


```solidity
function getNode(List storage self, uint256 _node) internal view returns (bool, uint256, uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|id of the node to get|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool, uint256, uint256 true if node exists or false otherwise, previous node, next node|
|`<none>`|`uint256`||
|`<none>`|`uint256`||


### getAdjacent

*Returns the link of a node `_node` in direction `_direction`.*


```solidity
function getAdjacent(List storage self, uint256 _node, bool _direction) internal view returns (bool, uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|id of the node to step from|
|`_direction`|`bool`|direction to step in|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool, uint256 true if node exists or false otherwise, node in _direction|
|`<none>`|`uint256`||


### getNextNode

*Returns the link of a node `_node` in direction `_NEXT`.*


```solidity
function getNextNode(List storage self, uint256 _node) internal view returns (bool, uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|id of the node to step from|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool, uint256 true if node exists or false otherwise, next node|
|`<none>`|`uint256`||


### getPreviousNode

*Returns the link of a node `_node` in direction `_PREV`.*


```solidity
function getPreviousNode(List storage self, uint256 _node) internal view returns (bool, uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|id of the node to step from|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool, uint256 true if node exists or false otherwise, previous node|
|`<none>`|`uint256`||


### insertAfter

*Insert node `_new` beside existing node `_node` in direction `_NEXT`.*


```solidity
function insertAfter(List storage self, uint256 _node, uint256 _new) internal returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|existing node|
|`_new`|`uint256`| new node to insert|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### insertBefore

*Insert node `_new` beside existing node `_node` in direction `_PREV`.*


```solidity
function insertBefore(List storage self, uint256 _node, uint256 _new) internal returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|existing node|
|`_new`|`uint256`| new node to insert|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### remove

*Removes an entry from the linked list*


```solidity
function remove(List storage self, uint256 _node) internal returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|node to remove from the list|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 the removed node|


### pushFront

*Pushes an entry to the head of the linked list*


```solidity
function pushFront(List storage self, uint256 _node) internal returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|new entry to push to the head|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### pushBack

*Pushes an entry to the tail of the linked list*


```solidity
function pushBack(List storage self, uint256 _node) internal returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|new entry to push to the tail|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### popFront

*Pops the first entry from the head of the linked list*


```solidity
function popFront(List storage self) internal returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 the removed node|


### popBack

*Pops the first entry from the tail of the linked list*


```solidity
function popBack(List storage self) internal returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 the removed node|


### _push

*Pushes an entry to the head of the linked list*


```solidity
function _push(List storage self, uint256 _node, bool _direction) private returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|new entry to push to the head|
|`_direction`|`bool`|push to the head (_NEXT) or tail (_PREV)|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### _pop

*Pops the first entry from the linked list*


```solidity
function _pop(List storage self, bool _direction) private returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_direction`|`bool`|pop from the head (_NEXT) or the tail (_PREV)|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 the removed node|


### _insert

*Insert node `_new` beside existing node `_node` in direction `_direction`.*


```solidity
function _insert(List storage self, uint256 _node, uint256 _new, bool _direction) private returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|existing node|
|`_new`|`uint256`| new node to insert|
|`_direction`|`bool`|direction to insert node in|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|bool true if success, false otherwise|


### _createLink

*Creates a bidirectional link between two nodes on direction `_direction`*


```solidity
function _createLink(List storage self, uint256 _node, uint256 _link, bool _direction) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`self`|`List`|stored linked list from contract|
|`_node`|`uint256`|existing node|
|`_link`|`uint256`|node to link to in the _direction|
|`_direction`|`bool`|direction to insert node in|


## Structs
### List

```solidity
struct List {
    uint256 size;
    mapping(uint256 => mapping(bool => uint256)) list;
}
```

