// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/core/AVSDirectory.sol";

contract AVSDirectoryHarness is AVSDirectory {
    constructor(IDelegationManager _delegation) AVSDirectory(_delegation) {}

    function magnitudeUpdate(
        address operator,
        IStrategy strategy,
        OperatorSet memory opSet
    ) public view returns (Checkpoints.History memory) {
        bytes32 opSetKey = _encodeOperatorSet(opSet);
        return _magnitudeUpdate[operator][strategy][opSetKey];
    }

    function totalMagnitude(address operator, IStrategy strategy) public view returns (Checkpoints.History memory) {
        return _totalMagnitudeUpdate[operator][strategy];
    }
}
