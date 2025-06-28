// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/multichain/CrossChainRegistry.sol";

contract CrossChainRegistryMock is Test, ICrossChainRegistryTypes {
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}
}
