// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

struct ConsolidationReq {
    address source;
    bytes sourcePubkey;
    bytes targetPubkey;
}

contract EIP_7251_Mock {

    address constant SYSTEM_ADDRESS = 0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE;

    uint constant EXCESS_INHIBITOR = type(uint).max;
    uint constant MIN_CONSOLIDATION_REQUEST_FEE = 1;
    uint constant CONSOLIDATION_REQUEST_FEE_UPDATE_FRACTION = 17;

    uint constant TARGET_CONSOLIDATION_REQUESTS_PER_BLOCK = 1;

    uint excessRequests;
    uint consolidationRequestCount;

    uint queueHead;
    uint queueTail;
    mapping(uint => ConsolidationReq) requestQueue;

    fallback() external payable {
        if (msg.sender == SYSTEM_ADDRESS) {
            _doReadRequests();
        } else if (msg.data.length == 0) {
            _doGetFee();
        } else if (msg.data.length == 96) {
            _doAddRequest();
        } else {
            revert("EIP_7251_Mock: unknown function");
        }
    }

    /**
     *
     *                            "PUBLIC" METHODS
     *
     */

    function _doAddRequest() internal {
        uint fee = _getFee();
        require(msg.value >= fee, "EIP_7251_Mock: insufficient value for fee");

        consolidationRequestCount++;

        bytes memory sourcePubkey = msg.data[0:48];
        bytes memory targetPubkey = msg.data[48:];

        ConsolidationReq storage request = requestQueue[queueTail];
        request.source = msg.sender;
        request.sourcePubkey = sourcePubkey;
        request.targetPubkey = targetPubkey;

        queueTail++;
    }

    function _doGetFee() internal view {
        uint fee = _getFee();

        assembly {
            mstore(0, fee)
            return(0, 32)
        }
    }

    /// (see https://eips.ethereum.org/EIPS/eip-7251#system-call)
    function _doReadRequests() internal {
        ConsolidationReq[] memory reqs = _dequeueConsolidationRequests();
        _updateExcessConsolidationRequests();
        _resetConsolidationRequestCount();

        bytes memory returnData = abi.encode(reqs);
        assembly { return(add(32, returnData), mload(returnData)) }
    }

    /**
     *
     *                            PRIVATE METHODS
     *
     */

    function _dequeueConsolidationRequests() internal returns (ConsolidationReq[] memory reqs) {
        uint numInQueue = queueTail - queueHead;
        // The actual spec caps this to MAX_CONSOLIDATION_REQUESTS_PER_BLOCK, but that's not super useful for tests
        uint numToDequeue = numInQueue;

        reqs = new ConsolidationReq[](numToDequeue);
        for (uint i = 0; i < numToDequeue; i++) {
            reqs[i] = requestQueue[queueHead + i];
        }

        uint newQueueHead = queueHead + numToDequeue;
        if (newQueueHead == queueTail) {
            queueHead = 0;
            queueTail = 0;
        } else {
            queueHead = newQueueHead;
        }

        return reqs;
    }

    function _updateExcessConsolidationRequests() internal {
        uint prevExcess = excessRequests;
        if (prevExcess == EXCESS_INHIBITOR) {
            prevExcess = 0;
        }

        uint newExcess = 0;
        if (prevExcess + consolidationRequestCount > TARGET_CONSOLIDATION_REQUESTS_PER_BLOCK) {
            newExcess = prevExcess + consolidationRequestCount - TARGET_CONSOLIDATION_REQUESTS_PER_BLOCK;
        }

        excessRequests = newExcess;
    }

    function _resetConsolidationRequestCount() internal {
        consolidationRequestCount = 0;
    }

    function _getFee() private view returns (uint) {
        require(excessRequests != EXCESS_INHIBITOR, "EIP_7251_Mock: excess inhibitor reached");

        return _fakeExponential({
            factor: MIN_CONSOLIDATION_REQUEST_FEE,
            numerator: excessRequests,
            denominator: CONSOLIDATION_REQUEST_FEE_UPDATE_FRACTION
        });
    }

    /// (see https://eips.ethereum.org/EIPS/eip-7251#fee-calculation)
    function _fakeExponential(uint factor, uint numerator, uint denominator) private pure returns (uint) {
        uint i = 1;
        uint output = 0;
        uint numeratorAccum = factor * denominator;

        while (numeratorAccum > 0) {
            output += numeratorAccum;
            numeratorAccum = (numeratorAccum * numerator) / (denominator * i);
            i++;
        }

        return output / denominator;
    }
}
