// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract SlashEscrowFactoryMock is Test {
    receive() external payable {}
    fallback() external payable {}

    function getSlashEscrow(OperatorSet calldata operatorSet, uint slashId) public view returns (address) {
        // Hash the operatorSet and slashId to get a random address
        return address(uint160(uint(keccak256(abi.encode(operatorSet, slashId)))));
    }
}
