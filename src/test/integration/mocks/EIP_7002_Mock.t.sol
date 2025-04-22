// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

struct WithdrawalRequest {
    address source;
    bytes validatorPubkey;
    uint64 amount;
}

contract EIP_7002_Mock {

    address constant SYSTEM_ADDRESS = 0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE;

    uint constant EXCESS_INHIBITOR = type(uint).max;
    uint constant MIN_WITHDRAWAL_REQUEST_FEE = 1;
    uint constant WITHDRAWAL_REQUEST_FEE_UPDATE_FRACTION = 17;

    uint MAX_WITHDRAWAL_REQUESTS_PER_BLOCK = 16;
    uint TARGET_WITHDRAWAL_REQUESTS_PER_BLOCK = 2;

    uint excessRequests;
    uint withdrawalRequestCount;

    uint queueHead;
    uint queueTail;
    mapping(uint => WithdrawalRequest) requestQueue;

    fallback() external payable {
        if (msg.sender == SYSTEM_ADDRESS) {
            _doReadRequests();
        } else if (msg.data.length == 0) {
            _doGetFee();
        } else if (msg.data.length == 56) {
            _doAddRequest();
        } else {
            revert("EIP_7002_Mock: unknown function");
        }
    }

    /**
     *
     *                            "PUBLIC" METHODS
     *
     */

    function _doAddRequest() internal {
        uint fee = _getFee();
        require(msg.value >= fee, "EIP_7002_Mock: insuffient value for fee");

        withdrawalRequestCount++;

        bytes memory pubkey = new bytes(48);
        uint64 amount;
        assembly {
            calldatacopy(add(32, pubkey), 0, 48) // copy bytes48 pubkey into memory
            calldatacopy(24, 48, 8)              // copy uint64 amount to memory[0]
            amount := mload(0)
        }

        WithdrawalRequest storage request = requestQueue[queueTail];
        request.source = msg.sender;
        request.validatorPubkey = pubkey;
        request.amount = amount;

        queueTail++;
    }

    function _doGetFee() internal view {
        uint fee = _getFee();

        assembly {
            mstore(0, fee)
            return(0, 32)
        }
    }

    /// (see https://eips.ethereum.org/EIPS/eip-7002#system-call)
    function _doReadRequests() internal {
        WithdrawalRequest[] memory reqs = _dequeueWithdrawalRequests();
        _updateExcessWithdrawalRequests();
        _resetWithdrawalRequestCount();

        bytes memory returnData = abi.encode(reqs);
        assembly { return(add(32, returnData), mload(returnData)) }
    }

    /**
     *
     *                            PRIVATE METHODS
     *
     */

    function _dequeueWithdrawalRequests() internal returns (WithdrawalRequest[] memory reqs) {
        uint numInQueue = queueTail - queueHead;
        uint numToDequeue = 
            numInQueue > MAX_WITHDRAWAL_REQUESTS_PER_BLOCK ? MAX_WITHDRAWAL_REQUESTS_PER_BLOCK : numInQueue;

        reqs = new WithdrawalRequest[](numToDequeue);
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

    function _updateExcessWithdrawalRequests() internal {
        uint prevExcess = excessRequests;
        if (prevExcess == EXCESS_INHIBITOR) {
            prevExcess = 0;
        }

        uint newExcess = 0;
        if (prevExcess + withdrawalRequestCount > TARGET_WITHDRAWAL_REQUESTS_PER_BLOCK) {
            newExcess = prevExcess + withdrawalRequestCount - TARGET_WITHDRAWAL_REQUESTS_PER_BLOCK;
        }

        excessRequests = newExcess;
    }

    function _resetWithdrawalRequestCount() internal {
        withdrawalRequestCount = 0;
    }

    function _getFee() private view returns (uint) {
        require(excessRequests != EXCESS_INHIBITOR, "EIP_7002_Mock: excess inhibitor reached");

        return _fakeExponential({
            factor: MIN_WITHDRAWAL_REQUEST_FEE,
            numerator: excessRequests,
            denominator: WITHDRAWAL_REQUEST_FEE_UPDATE_FRACTION
        });
    }

    /// (see https://eips.ethereum.org/EIPS/eip-7002#fee-calculation)
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
