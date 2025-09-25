// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/libraries/OperatorSetLib.sol";

contract OperatorSetHelper {
    using OperatorSetLib for OperatorSet;
    constructor(){}

    /// @notice Returns the operatorSet key derived from avs and id
    function getOperatorSetKey(OperatorSet calldata os) external pure returns (bytes32) {
        return os.key(); // calls OperatorSetLib.key()
    }
}