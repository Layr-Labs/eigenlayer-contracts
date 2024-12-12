// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

import "zeus-templates/utils/Encode.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "script/releases/Env.sol";

import "src/contracts/strategies/StrategyFactory.sol";
import "forge-std/Script.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade. 
 */

/**
    anvil --fork-url $RPC_HOLESKY
    forge script script/deploy/devnet/preprod_strategy_hotfix/2-strategyFactoryWhitelist.s.sol:EigenStrategyPreprodHotfix \
        --broadcast \
        --rpc-url http://127.0.0.1:8545 \
        --private-key $PRIVATE_KEY -vvvv
 */
contract EigenStrategyPreprodHotfix is Script, MultisigBuilder {
    using Env for *;
    using Encode for *;

    /// These addresses not found in zeus config so hard coding them here
    address bEIGEN = 0xA72942289a043874249E60469F68f08B8c6ECCe8;
    address EIGEN = 0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926;
    // Deployed EigenStrategy separately, at the end of this script, this should get pushed
    // to zeus preprod metadata
    EigenStrategy eigenStrategyToAdd = EigenStrategy(0x4e0125f8a928Eb1b9dB4BeDd3756BA3c200563C2);

    MultisigCall[] internal _opsCalls;

    function run() public {
        execute();
    }

    function _runAsMultisig() internal override {
        bytes memory calldata_to_operations_multisig = _getCalldataToOpsMultisig();

        _startPrank2(Env.opsMultisig());

        (bool success, bytes memory returnData) = Env.multiSendCallOnly().delegatecall(
            abi.encodeCall(IMultiSend.multiSend, (calldata_to_operations_multisig))
        );

        // (bool success, ) = address(Env.opsMultisig()).call(calldata_to_operations_multisig);
        // require(success, "main transaction failed");
        _stopPrank();
    }

    function _getCalldataToOpsMultisig() internal returns (bytes memory) {
        // 1. prepare calldata for multisig 
        IStrategy[] memory strategiesToRemoveFromWhitelist = new IStrategy[](1);
        strategiesToRemoveFromWhitelist[0] = Env.proxy.eigenStrategy();
        IStrategy[] memory strategiesToWhitelist = new IStrategy[](1);
        strategiesToWhitelist[0] = eigenStrategyToAdd;

        // 2. prepare multisig calls for opsMultisig
        // Remove deprecated strategy from whitelist
        _opsCalls.append({
            to: address(Env.proxy.strategyFactory()),
            data: abi.encodeWithSelector(
                StrategyFactory.removeStrategiesFromWhitelist.selector,
                strategiesToRemoveFromWhitelist
            )
        });
        // Add new strategy to whitelist
        _opsCalls.append({
            to: address(Env.proxy.strategyFactory()),
            data: abi.encodeWithSelector(
                StrategyFactory.whitelistStrategies.selector,
                strategiesToWhitelist
            )
        });

        return (Encode.multiSend(_opsCalls));
    }

    function _startPrank2(address caller) internal {
        require(!hasPranked, "MultisigBuilder._startPrank: called twice in txn");
        hasPranked = true;

        emit ZeusRequireMultisig(caller, Encode.Operation.DelegateCall);
        vm.startPrank(caller);
    }

    function _stopPrank2() internal {
        require(hasPranked, "MultisigBuilder._stopPrank: _startPrank not called");
        vm.stopPrank();
    }
}