// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import "./ConfigParser.sol";
import "./StringUtils.sol";
import "./AddressUtils.sol";
import "./TxBuilder.sol";

contract OpsMultisigBuilder is Releasoor {

}

contract Releasoor is ConfigParser {

    using StringUtils for *;
    
    Environment env;
    Params params;

    uint txsCount;
    mapping(uint => Txs) transactions;

    string constant DEPLOY = "deploy";
    string constant QUEUE = "queue";
    string constant EXECUTE = "execute";

    string constant ENV_MAINNET = "mainnet";
    string constant ENV_HOLESKY = "testnet-holesky";
    string constant ENV_PREPROD = "preprod-holesky";

    string constant CONFIG_MAINNET = "script/configs/mainnet.json";
    string constant CONFIG_HOLESKY = "script/configs/holesky.json";
    string constant CONFIG_PREPROD = "script/configs/preprod.json";

    string constant CONFIG_ZIPZOOP = "script/configs/zipzoop.json";
    
    function deploy(string memory jsonFile) public returns (Deployment[] memory deployments) {

        (Addresses, Env, Params) = _readConfig();

        return _deploy(addrs, env, params);
    }

    /// $ run execute mainnet
    /// -- requires that you have previously run deploy and queue
    /// ... except that sometimes we don't need to run deploy OR queue. For example:
    /// I want to write/run a script that just has the ops multisig call `addStratsToWhitelist`.
    /// This doesn't involve a timelock, as the ops multisig holds the whitelister role
    /// We may have deployed something in the deploy script, but we probably didn't queue anything
    /// for this.
    ///
    /// So, it's mainly important that execute fails if:
    /// - We reference a queued transaction that doesn't exist in the timelock
    ///   - This would mean we haven't called `queue` yet
    /// - We reference a pendingImpl that doesn't exist in the config
    ///   - This would mean we haven't called `deploy` yet
    function run(string memory action, string memory _env) public returns (bytes memory) {
        _log(action);
        _log(_env);

        _log("Running script: ".concat(action).concat(" for env: ").concat(_env));

        uint currentChainId = block.chainid;
        emit log_named_uint("You are currently on chain id", currentChainId);

        // string memory filePath;
        // if (_env.eq(ENV_MAINNET)) {
        //     filePath = CONFIG_MAINNET;
        // } else if (_env.eq(ENV_HOLESKY)) {
        //     filePath = CONFIG_HOLESKY;
        // } else if (_env.eq(ENV_PREPROD)) {
        //     filePath = CONFIG_PREPROD;
        // } else {
        //     revert("invalid env");
        // }

        (
            Addresses memory addrs, 
            Environment memory _env,
            Params memory _params
        ) = _readConfigFile(CONFIG_ZIPZOOP);
        _printEnv(_env);

        env = _env;
        params = _params;
        require(_env.chainid == currentChainId, "You are on the wrong chain for this config");
        
        bytes32 initialAddrs = keccak256(abi.encode(addrs));
        
        if (action.eq(DEPLOY)) {
            (string[] memory names, ..) = deploy(addrs);
        } else if (action.eq(QUEUE)) {
            (Txs storage executorTxns, uint eta) = queueUpgrade(addrs);

            _printQueueUpgradeSummary(executorTxns, eta);
        } else if (action.eq(EXECUTE)) {
            executeUpgrade(addrs);
        } else {
            revert("invalid action");
        }

        _log("Script complete");

        // Check if we made changes to config
        if (initialAddrs != keccak256(abi.encode(addrs))) {
            _writeConfigFile(addrs, _env, _params);
        }
    }

    function deploy(Addresses memory addrs) internal virtual {
        revert("deploy not implemented");
    }

    function queueUpgrade(Addresses memory addrs) internal virtual returns (Txs storage executorTxns, uint eta) {
        revert("queueUpgrade not implemented");
    }

    function executeUpgrade(Addresses memory addrs) internal virtual {
        revert("executeUpgrade not implemented");
    }

    function _newTxs() internal returns (Txs storage) {
        Txs storage txs = transactions[txsCount];
        txsCount++;

        return txs;
    }

    function _printQueueUpgradeSummary(Txs storage executorTxns, uint eta) internal {
        bytes memory calldata_to_multisend_contract = EncMultiSendCallOnly.multiSend(executorTxns);
        emit log_named_bytes("calldata_to_multisend_contract", calldata_to_multisend_contract);

        bytes memory final_calldata_to_executor_multisig = EncGnosisSafe.execTransaction({
            from: addrs.timelock,
            to: params.multiSendCallOnly,
            data: calldata_to_multisend_contract,
            op: EncGnosisSafe.Operation.DelegateCall
        });
        emit log_named_bytes("final_calldata_to_executor_multisig", final_calldata_to_executor_multisig);

        bytes memory calldata_to_timelock_queueing_action = EncTimelock.queueTransaction({
            target: addrs.executorMultisig,
            data: final_calldata_to_executor_multisig,
            eta: 0
        });
        emit log_named_bytes("calldata_to_timelock_queueing_action", calldata_to_timelock_queueing_action);
    }

    function _log(string memory s) internal {
        emit log(s);
    }

    function _log(string memory s, address a) internal {
        emit log_named_address(s, a);
    }
}