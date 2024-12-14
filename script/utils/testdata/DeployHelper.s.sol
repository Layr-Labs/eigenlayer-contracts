// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./utils/RelayScript.s.sol";
import "./utils/StrategyToken.sol";

import "../../../src/contracts/strategies/StrategyFactory.sol";
import "../../../src/contracts/core/DelegationManager.sol";
import "../../../src/contracts/core/StrategyManager.sol";

abstract contract DeployHelper is RelayScript {

    function _deployStrategy(string memory name, string memory symbol, uint256 initialSupply) internal returns (IERC20, IStrategy) {
        StrategyToken token = new StrategyToken(
            name,
            symbol,
            initialSupply,
            msg.sender
        );
        uint64 nonce = vm.getNonce(msg.sender);
        IStrategy strategy = _strategyFactory().deployNewStrategy(IERC20(token));
        if (vm.getNonce(msg.sender) == nonce) {
            vm.setNonce(msg.sender, nonce + 1);
        }
        return (token, strategy);
    }

    function _strategyFactory() internal pure returns (StrategyFactory) {
        return StrategyFactory(0xad4A89E3cA9b3dc25AABe0aa7d72E61D2Ec66052);
    }

    function _delegationManager() internal pure returns (DelegationManager) {
        return DelegationManager(0x75dfE5B44C2E530568001400D3f704bC8AE350CC);
    }

    function _strategyManager() internal pure returns (StrategyManager) {
        return StrategyManager(0xF9fbF2e35D8803273E214c99BF15174139f4E67a);
    }
} 