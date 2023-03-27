# Solidity API

## StructuredLinkedList

Adapted from https://github.com/vittominacori/solidity-linked-list/blob/master/contracts/StructuredLinkedList.sol

_An utility library for using sorted linked list data structures in your Solidity project._

### _NULL

```solidity
uint256 _NULL
```

### _HEAD

```solidity
uint256 _HEAD
```

### _PREV

```solidity
bool _PREV
```

### _NEXT

```solidity
bool _NEXT
```

### List

```solidity
struct List {
  uint256 size;
  mapping(uint256 => mapping(bool => uint256)) list;
}
```

### listExists

```solidity
function listExists(struct StructuredLinkedList.List self) internal view returns (bool)
```

_Checks if the list exists_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if list exists, false otherwise |

### nodeExists

```solidity
function nodeExists(struct StructuredLinkedList.List self, uint256 _node) internal view returns (bool)
```

_Checks if the node exists_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | a node to search for |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if node exists, false otherwise |

### sizeOf

```solidity
function sizeOf(struct StructuredLinkedList.List self) internal view returns (uint256)
```

_Returns the number of elements in the list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 |

### getHead

```solidity
function getHead(struct StructuredLinkedList.List self) internal view returns (uint256)
```

_Gets the head of the list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 the head of the list |

### getNode

```solidity
function getNode(struct StructuredLinkedList.List self, uint256 _node) internal view returns (bool, uint256, uint256)
```

_Returns the links of a node as a tuple_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | id of the node to get |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool, uint256, uint256 true if node exists or false otherwise, previous node, next node |
| [1] | uint256 |  |
| [2] | uint256 |  |

### getAdjacent

```solidity
function getAdjacent(struct StructuredLinkedList.List self, uint256 _node, bool _direction) internal view returns (bool, uint256)
```

_Returns the link of a node `_node` in direction `_direction`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | id of the node to step from |
| _direction | bool | direction to step in |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool, uint256 true if node exists or false otherwise, node in _direction |
| [1] | uint256 |  |

### getNextNode

```solidity
function getNextNode(struct StructuredLinkedList.List self, uint256 _node) internal view returns (bool, uint256)
```

_Returns the link of a node `_node` in direction `_NEXT`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | id of the node to step from |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool, uint256 true if node exists or false otherwise, next node |
| [1] | uint256 |  |

### getPreviousNode

```solidity
function getPreviousNode(struct StructuredLinkedList.List self, uint256 _node) internal view returns (bool, uint256)
```

_Returns the link of a node `_node` in direction `_PREV`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | id of the node to step from |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool, uint256 true if node exists or false otherwise, previous node |
| [1] | uint256 |  |

### insertAfter

```solidity
function insertAfter(struct StructuredLinkedList.List self, uint256 _node, uint256 _new) internal returns (bool)
```

_Insert node `_new` beside existing node `_node` in direction `_NEXT`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | existing node |
| _new | uint256 | new node to insert |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### insertBefore

```solidity
function insertBefore(struct StructuredLinkedList.List self, uint256 _node, uint256 _new) internal returns (bool)
```

_Insert node `_new` beside existing node `_node` in direction `_PREV`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | existing node |
| _new | uint256 | new node to insert |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### remove

```solidity
function remove(struct StructuredLinkedList.List self, uint256 _node) internal returns (uint256)
```

_Removes an entry from the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | node to remove from the list |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 the removed node |

### pushFront

```solidity
function pushFront(struct StructuredLinkedList.List self, uint256 _node) internal returns (bool)
```

_Pushes an entry to the head of the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | new entry to push to the head |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### pushBack

```solidity
function pushBack(struct StructuredLinkedList.List self, uint256 _node) internal returns (bool)
```

_Pushes an entry to the tail of the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | new entry to push to the tail |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### popFront

```solidity
function popFront(struct StructuredLinkedList.List self) internal returns (uint256)
```

_Pops the first entry from the head of the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 the removed node |

### popBack

```solidity
function popBack(struct StructuredLinkedList.List self) internal returns (uint256)
```

_Pops the first entry from the tail of the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 the removed node |

### _push

```solidity
function _push(struct StructuredLinkedList.List self, uint256 _node, bool _direction) private returns (bool)
```

_Pushes an entry to the head of the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | new entry to push to the head |
| _direction | bool | push to the head (_NEXT) or tail (_PREV) |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### _pop

```solidity
function _pop(struct StructuredLinkedList.List self, bool _direction) private returns (uint256)
```

_Pops the first entry from the linked list_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _direction | bool | pop from the head (_NEXT) or the tail (_PREV) |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 the removed node |

### _insert

```solidity
function _insert(struct StructuredLinkedList.List self, uint256 _node, uint256 _new, bool _direction) private returns (bool)
```

_Insert node `_new` beside existing node `_node` in direction `_direction`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | existing node |
| _new | uint256 | new node to insert |
| _direction | bool | direction to insert node in |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if success, false otherwise |

### _createLink

```solidity
function _createLink(struct StructuredLinkedList.List self, uint256 _node, uint256 _link, bool _direction) private
```

_Creates a bidirectional link between two nodes on direction `_direction`_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | struct StructuredLinkedList.List | stored linked list from contract |
| _node | uint256 | existing node |
| _link | uint256 | node to link to in the _direction |
| _direction | bool | direction to insert node in |

