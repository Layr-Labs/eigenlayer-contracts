methods {
    function listExists() external returns (bool) envfree;
    function nodeExists(uint256) external returns (bool) envfree;
    function sizeOf() external returns (uint256) envfree;
    function getHead() external returns (uint256) envfree;
    function getNode(uint256) external returns (bool, uint256, uint256) envfree;
    function getAdjacent(uint256,bool) external returns (bool, uint256) envfree;
    function getAdjacentStrict(uint256,bool) external returns (uint256) envfree;
    function getNextNode(uint256) external returns (bool, uint256) envfree;
    function getPreviousNode(uint256) external returns (bool, uint256) envfree;
    function insert(uint256,uint256,bool) external envfree;
    function remove(uint256) external envfree;
}

ghost mapping(uint256 => bool) connectsToHead {
    init_state axiom connectsToHead[0] == true;
}

hook Sstore currentContract.listStorage.list[KEY uint256 node][KEY bool direction] uint256 link (uint256 old_link) {
    connectsToHead[link] = connectsToHead[node];
    connectsToHead[old_link] = old_link == 0;
}

definition isNodeDoubleLinked(uint256 node) returns bool =
    node == getAdjacentStrict(getAdjacentStrict(node, true), false)
    && node == getAdjacentStrict(getAdjacentStrict(node, false), true);

definition doesNodePointToSelf(uint256 node) returns bool =
    getAdjacentStrict(node, true) == node || getAdjacentStrict(node, false) == node;

// if node x points at node y, node y must point back at node x
// ==?
// only two way links exist.
invariant alwaysDoubleLinked(uint256 node)
    nodeExists(node) => isNodeDoubleLinked(node)
    {
        preserved {
            requireInvariant noSelfPoint(node);
            requireInvariant alwaysDoubleLinked(0);
        }
        preserved insert(uint256 _node, uint256 _new, bool _dir) {
            requireInvariant alwaysDoubleLinked(_node);
        }
        preserved remove(uint256 _node) {
            requireInvariant alwaysDoubleLinked(_node);
            requireInvariant alwaysDoubleLinked(0);           
        }
    }

/// Nonhead node can't point to itself.
/// @title noSelfPoint
invariant noSelfPoint(uint256 node)
    node != 0 => !doesNodePointToSelf(node)
    {
        preserved remove(uint256 _node) {
            requireInvariant alwaysDoubleLinked(_node);
            requireInvariant zeroRequiredInCircle(node, _node);
        }
    }

/// A node can not point to 0 if 0 does not point back.
/// @title noDeadEnds
invariant noDeadEnds(uint256 node, uint256 lost, bool dir)
    getAdjacentStrict(node, dir) == 0 
        => getAdjacentStrict(0, !dir) == node 
        || (getAdjacentStrict(lost, dir) != node && getAdjacentStrict(lost, !dir) != node)
    {
        preserved insert(uint256 _node, uint256 _new, bool _dir) {
            requireInvariant alwaysDoubleLinked(_node);
        }
        preserved remove(uint256 _node) {
            requireInvariant alwaysDoubleLinked(_node);
            requireInvariant alwaysDoubleLinked(node);
            requireInvariant alwaysDoubleLinked(0);
        }
    }

/// 0 must point to itself in both directions or not at all.
invariant allOrNothing()
    getAdjacentStrict(0, true) == 0 <=> getAdjacentStrict(0, false) == 0
    {
        preserved insert(uint256 _node, uint256 _new, bool _dir) {
            requireInvariant alwaysDoubleLinked(_node);
        }
        preserved remove(uint256 _node) {
            requireInvariant alwaysDoubleLinked(_node);
        }
    }

/// No loop without o
invariant zeroRequiredInCircle(uint256 node1, uint256 node2)
    node1 != node2 && getAdjacentStrict(node1, true) == node2
        && getAdjacentStrict(node1, false) == node2
        && getAdjacentStrict(node2, true) == node1
        && getAdjacentStrict(node2, false) == node1
        => node1 == 0 || node2 == 0
    {
        preserved remove(uint256 _node) {
            require !((getAdjacentStrict(_node, true) == node1 
                || getAdjacentStrict(_node, false) == node1
                && getAdjacentStrict(_node, true) == node2 ||
                getAdjacentStrict(_node, false) == node2)
                && ((getAdjacentStrict(node1, true) == node2 
                || getAdjacentStrict(node1, false) == node2)
                && (getAdjacentStrict(node2, true) == node1 
                || getAdjacentStrict(node2, false) == node1)));
        }
    }

/* commented out while failing (can reintroduce in a PR)
// in progress
invariant headInList(uint256 node)
    nodeExists(node) => connectsToHead[node]
    {
        preserved insert(uint256 _node, uint256 _new, bool _dir) {
            requireInvariant noSelfPoint(_node);
            requireInvariant alwaysDoubleLinked(_node);
            requireInvariant noSelfPoint(node);
            requireInvariant alwaysDoubleLinked(node);
        }
        preserved remove(uint256 _node) {
            requireInvariant noSelfPoint(_node);
            requireInvariant alwaysDoubleLinked(_node);
            requireInvariant noSelfPoint(node);
            requireInvariant alwaysDoubleLinked(node);
        }
    }



// size == # of nodes with nonzero next == # of nodes with nonzero prev
*/


/*
1) `StructuredLinkedList._createLink` creates only two-way links
2) `StructuredLinkedList.remove` removes both links from a node, and stitches together its existing links (which it breaks)
3) `StructuredLinkedList._insert` similarly inserts a new node 'between' nodes, ensuring that the new node is well-linked
*/
/*
// DEFINITIONS
definition isInDLL(address _id) returns bool =
    getValueOf(_id) != 0;
definition isLinked(address _id) returns bool =
    getPrev(_id) != 0 || getNext(_id) != 0;
definition isEmpty(address _id) returns bool =
    ! isInDLL(_id) && ! isLinked(_id);
definition isTwoWayLinked(address first, address second) returns bool =
    first != 0 && second != 0 => (getNext(first) == second <=> getPrev(second) == first);
definition isHeadWellFormed() returns bool =
    getPrev(getHead()) == 0 && (getHead() != 0 => isInDLL(getHead()));
definition isTailWellFormed() returns bool =
    getNext(getTail()) == 0 && (getTail() != 0 => isInDLL(getTail()));
definition hasNoPrevIsHead(address addr) returns bool =
    isInDLL(addr) && getPrev(addr) == 0 => addr == getHead();
definition hasNoNextIsTail(address addr) returns bool =
    isInDLL(addr) && getNext(addr) == 0 => addr == getTail();
function safeAssumptions() { 
    requireInvariant zeroEmpty();
    requireInvariant headWellFormed();
    requireInvariant tailWellFormed();
    requireInvariant tipsZero();
}
// INVARIANTS & RULES
// Notice that some invariants have the preservation proof separated for some public functions,
// or even all of the public functions (in that last case they are still relevant for proving 
// the property at initial state).

invariant zeroEmpty()
    isEmpty(0)
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }

rule zeroEmptyPreservedInsertSorted(address _id, uint256 _value) {
    address prev;

    require isEmpty(0);

    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant noNextIsTail(prev);
    requireInvariant tipsZero();

    insertSorted(_id, _value, maxIterations());

    require prev == getInsertedAfter();

    assert isEmpty(0);
}

invariant headWellFormed()
    isHeadWellFormed()
    { preserved remove(address _id) {
        requireInvariant zeroEmpty();
        requireInvariant twoWayLinked(getPrev(_id), _id);
        requireInvariant twoWayLinked(_id, getNext(_id));
        requireInvariant linkedIsInDLL(getNext(_id));
      }
    }

invariant tailWellFormed()
    isTailWellFormed()
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }
    { preserved remove(address _id) {
        requireInvariant zeroEmpty();
        requireInvariant twoWayLinked(getPrev(_id), _id);
        requireInvariant twoWayLinked(_id, getNext(_id));
        requireInvariant linkedIsInDLL(getPrev(_id));
      }
    }

rule tailWellFormedPreservedInsertSorted(address _id, uint256 _value) {
    address next; address prev;

    require isTailWellFormed();

    requireInvariant zeroEmpty();
    requireInvariant twoWayLinked(getPrev(next), next);
    requireInvariant twoWayLinked(prev, getNext(prev));

    insertSorted(_id, _value, maxIterations());
    
    require prev == getInsertedAfter();
    require next == getInsertedBefore();
    
    assert isTailWellFormed();
}

invariant tipsZero()
    getTail() == 0 <=> getHead() == 0
    { preserved remove(address _id) {
        requireInvariant zeroEmpty();
        requireInvariant noNextIsTail(_id);
        requireInvariant noPrevIsHead(_id);
      }
    }

invariant noPrevIsHead(address addr)
    hasNoPrevIsHead(addr)
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }
    { preserved remove(address _id) {
        safeAssumptions();
        requireInvariant twoWayLinked(_id, getNext(_id));
        requireInvariant twoWayLinked(getPrev(_id), _id);
        requireInvariant noPrevIsHead(_id);
      }
    }

rule noPrevIsHeadPreservedInsertSorted(address _id, uint256 _value) {
    address addr; address next; address prev;

    require hasNoPrevIsHead(addr);

    safeAssumptions();
    requireInvariant twoWayLinked(getPrev(next), next);
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant noNextIsTail(prev);

    insertSorted(_id, _value, maxIterations());
    
    require prev == getInsertedAfter();
    require next == getInsertedBefore();
    
    assert hasNoPrevIsHead(addr);
}

invariant noNextIsTail(address addr)
    hasNoNextIsTail(addr)
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }
    { preserved remove(address _id) {
        safeAssumptions();
        requireInvariant twoWayLinked(_id, getNext(_id));
        requireInvariant twoWayLinked(getPrev(_id), _id);
        requireInvariant noNextIsTail(_id);
      }
    }

rule noNextisTailPreservedInsertSorted(address _id, uint256 _value) {
    address addr; address next; address prev;

    require hasNoNextIsTail(addr);

    safeAssumptions();
    requireInvariant twoWayLinked(getPrev(next), next);
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant forwardLinked(getTail());

    insertSorted(_id, _value, maxIterations());
    
    require prev == getInsertedAfter();
    require next == getInsertedBefore();
    
    assert hasNoNextIsTail(addr);
}

invariant linkedIsInDLL(address addr)
    isLinked(addr) => isInDLL(addr)
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }
    { preserved remove(address _id) {
        safeAssumptions();
        requireInvariant twoWayLinked(_id, getNext(_id));
        requireInvariant twoWayLinked(getPrev(_id), _id);
      }
    }

rule linkedIsInDllPreservedInsertSorted(address _id, uint256 _value) {
    address addr; address next; address prev;

    require isLinked(addr) => isInDLL(addr);
    require isLinked(getPrev(next)) => isInDLL(getPrev(next));

    safeAssumptions();
    requireInvariant twoWayLinked(getPrev(next), next);
    requireInvariant noPrevIsHead(next);
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant noNextIsTail(prev);

    insertSorted(_id, _value, maxIterations());

    require prev == getInsertedAfter();
    require next == getInsertedBefore();

    assert isLinked(addr) => isInDLL(addr);
}

invariant twoWayLinked(address first, address second)
    isTwoWayLinked(first, second)
    filtered { f -> f.selector != sig:insertSorted(address, uint256, uint256).selector }
    { preserved remove(address _id) {
        safeAssumptions();
        requireInvariant twoWayLinked(getPrev(_id), _id);
        requireInvariant twoWayLinked(_id, getNext(_id));
      }
    }

rule twoWayLinkedPreservedInsertSorted(address _id, uint256 _value) {
    address first; address second; address next;

    require isTwoWayLinked(first, second);
    require isTwoWayLinked(getPrev(next), next);

    safeAssumptions();
    requireInvariant linkedIsInDLL(_id);

    insertSorted(_id, _value, maxIterations());

    require next == getInsertedBefore();

    assert isTwoWayLinked(first, second);
}

invariant forwardLinked(address addr)
    isInDLL(addr) => isForwardLinkedBetween(getHead(), addr)
    filtered { f -> f.selector != sig:remove(address).selector &&
                    f.selector != sig:insertSorted(address, uint256, uint256).selector }

rule forwardLinkedPreservedInsertSorted(address _id, uint256 _value) {
    address addr; address prev;

    require isInDLL(addr) => isForwardLinkedBetween(getHead(), addr);
    require isInDLL(getTail()) => isForwardLinkedBetween(getHead(), getTail());

    safeAssumptions();
    requireInvariant linkedIsInDLL(_id);
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant noNextIsTail(prev);

    insertSorted(_id, _value, maxIterations());

    require prev == getInsertedAfter();

    assert isInDLL(addr) => isForwardLinkedBetween(getHead(), addr);
}

rule forwardLinkedPreservedRemove(address _id) {
    address addr; address prev;

    require prev == getPreceding(_id);

    require isInDLL(addr) => isForwardLinkedBetween(getHead(), addr);

    safeAssumptions();
    requireInvariant noPrevIsHead(_id);
    requireInvariant twoWayLinked(getPrev(_id), _id);
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant noNextIsTail(_id);

    remove(_id);

    assert isInDLL(addr) => isForwardLinkedBetween(getHead(), addr);
}

rule removeRemoves(address _id) {
    safeAssumptions();

    remove(_id);

    assert !isInDLL(_id);
}

rule insertSortedInserts(address _id, uint256 _value) {
    safeAssumptions();

    insertSorted(_id, _value, maxIterations());

    assert isInDLL(_id);
}

rule insertSortedDecreasingOrder(address _id, uint256 _value) {
    address prev;

    uint256 maxIter = maxIterations();

    safeAssumptions();
    requireInvariant twoWayLinked(prev, getNext(prev));
    requireInvariant linkedIsInDLL(_id);

    insertSorted(_id, _value, maxIter);

    require prev == getInsertedAfter();

    uint256 positionInDLL = lenUpTo(_id);

    assert positionInDLL > maxIter => greaterThanUpTo(_value, 0, maxIter) && _id == getTail();
    assert positionInDLL <= maxIter => greaterThanUpTo(_value, _id, getLength()) && _value > getValueOf(getNext(_id));
}

// DERIVED RESULTS

// result: isForwardLinkedBetween(getHead(), getTail())
// explanation: if getTail() == 0, then from tipsZero() we know that getHead() == 0 so the result follows
// otherwise, from tailWellFormed(), we know that isInDLL(getTail()) so the result follows from forwardLinked(getTail()).

// result: forall addr. isForwardLinkedBetween(addr, getTail())
// explanation: it can be obtained from the previous result and forwardLinked.
// Going from head to tail should lead to addr in between (otherwise addr is never reached because there is nothing after the tail).

// result: "BackwardLinked" dual results
// explanation: it can be obtained from ForwardLinked and twoWayLinked.

// result: there is only one list
// explanation: it comes from the fact that every non zero address that is in the DLL is linked to getHead().

// result: there are no cycles that do not contain the 0 address
// explanation: let N be a node in a cycle. Since there is a link from getHead() to N, it means that getHead()
// is part of the cycle. This is absurd because we know from headWellFormed() that the previous element of
// getHead() is the 0 address. 