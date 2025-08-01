# Operations Scripting With Zeus

This directory is where you will build [Zeus](https://github.com/Layr-Labs/zeus) scripts to do one-off protocol interactions. It *should not* be used for pauser multisig interactions - see the Pauser guide for more information. 

When running or testing a script, _you tell zeus which environment to use,_ and it will fork the corresponding network state and setup environment variables for that environment's params/deployment addresses.

zeus script --env testnet-hoodi --multisig script/releases/tests/ExecuteTransferOwnership.s.sol

##### Getting Started

* Install [Zeus](https://github.com/Layr-Labs/zeus)
* Run `zeus login`

At this point, you should be able to view an environment's config (try `zeus env show preprod`)

---

### Writing a Script

Operations scripts are generally single script that require a multisig step. The step should include comprehensive test functions to validate the changes at each stage of the upgrade process.

### Upgrade.json
The `upgrade.json` file denotes the steps of the script. See previous upgrade scripts for examples

### MultisigBuilder

`MultisigBuilder` is the base class for any action to be taken by a Safe Multisig. The entry function is `execute()` and requires adding `-s "execute(string memory)" "<path_to_config>"` as a `forge script` flag. Any inheriting contract is expected to inherit the `_execute()` internal function and return a `MultisigCall[]` object, containing all intended multisig calls.

The `MultisigCall` struct contains 3 fields:
* `to`: the address to call with the multisig
* `value`: the amount of ETH to send as part of the transaction.
* `data`: the calldata associated with the multisig call.

Once the `_execute()` function is implemented, the base MultisigBuilder contract will combine these calls into one `SafeTx` object, which is intended to pass all calls into the [MultiSendCallOnly](https://github.com/safe-global/safe-smart-account/blob/6fde75d29c8b52d5ac0c93a6fb7631d434b64119/contracts/libraries/MultiSendCallOnly.sol) contract.

### Running A Script

Example script:

```bash
zeus script --env <env> --<eoa or multisig> <path to script>

# zeus script --env testnet-sepolia --multisig script/operations/update-generator/1-updateGenerator.s.sol
```

To run tests:

```bash
zeus test --env <env> <path to script> --rpcUrl <rpc_url>

# zeus test script/operations/update-generator/1-updateGenerator.s.sol --env testnet-sepolia --rpcUrl $RPC_SEPOLIA
```