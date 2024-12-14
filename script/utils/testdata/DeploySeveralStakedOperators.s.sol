// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./DeployHelper.s.sol";

contract DeploySeveralStakedOperators is DeployHelper {
    uint256 constant OPERATOR_TRANSFER_BATCH_SIZE = 10;
    uint256 constant OPERATOR_REGISTRATION_BATCH_SIZE = 10;

    function run(uint256 numStrategies, uint256 numOperators) public relayBroadcast {
        _startBroadcastEOA(msg.sender);

        IERC20[] memory tokens = new IERC20[](numStrategies);
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        for (uint256 i = 0; i < numStrategies; i++) {
            (tokens[i], strategies[i]) = _deployStrategy("Test", "TEST", 1e50);
            console.log("deployed strategy %s", address(strategies[i]));
        }
        _flushQueuedBroadcasts();

        address[] memory operators = new address[](numOperators);
        uint256[][] memory operatorBalances = new uint256[][](numOperators);
        for (uint256 i = 0; i < numOperators; i++) {
            // deploy a new relay
            operators[i] = address(new Relay());
            console.log("deployed operator %s", operators[i]);
            uint64 nonce = vm.getNonce(msg.sender);
            // transfer some random amount of tokens to the relay
            for (uint256 j = 0; j < numStrategies; j++) {
                operatorBalances[i] = new uint256[](numStrategies);
                operatorBalances[i][j] = vm.randomUint(0, 100e18);
                tokens[j].transfer(operators[i], operatorBalances[i][j]);
            }
            if (nonce == vm.getNonce(msg.sender)) {
                vm.setNonce(msg.sender, nonce + uint64(numStrategies));
            }

            if (i % OPERATOR_TRANSFER_BATCH_SIZE == 0) {
                _flushQueuedBroadcasts();
            }
        }
        _stopBroadcast();

        for (uint256 i = 0; i < numOperators; i++) {
            _startBroadcastRelay(operators[i]);
            // approve and deposit into each strategy
            for (uint256 j = 0; j < numStrategies; j++) {
                if (operatorBalances[i][j] > 0) {
                    tokens[j].approve(address(_strategyManager()), type(uint256).max);
                    _strategyManager().depositIntoStrategy(strategies[j], tokens[j], operatorBalances[i][j]);
                }
            }
            // register the relay as an operator
            // 0 allocation delay
            _delegationManager().registerAsOperator(address(0), 0, "https://example.com");
            _stopBroadcast();

            if (i % OPERATOR_REGISTRATION_BATCH_SIZE == 0) {
                _flushQueuedBroadcasts();
            }
        }
    }
} 
