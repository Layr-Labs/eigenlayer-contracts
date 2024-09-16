# Release Scripting With Zeus

This directory is where you will build [Zeus](https://github.com/Layr-Labs/zeus) scripts to manage core protocol releases. Releases are broken up into multiple steps composed of either a `forge` script or a custom shell script. Zeus's role is:
* Provide an environment to your script via environment variables
* Update environment after a release is run to completion
* Track status of releases to ensure steps are run (in order) only once per environment

**Note about environments:** Zeus scripts are intended to be written once, and run across _any_ environment we use. We currently have 3 live environments (`preprod`, `testnet`, and `mainnet`), and the params/deployment addresses for each live in separate folders in [`layr-labs/eigenlayer-contracts-metadata`](https://github.com/Layr-Labs/eigenlayer-contracts-metadata).

When running or testing a script, _you tell zeus which environment to use,_ and it will fork the corresponding network state and setup environment variables for that environment's params/deployment addresses.

##### Getting Started

* Install [Zeus](https://github.com/Layr-Labs/zeus)
* Run `zeus login`

At this point, you should be able to view an environment's config (try `zeus env show preprod`)

---

### Writing a Script

Scripts are broken up into multiple steps TODO


### EOADeployer

`EOADeployer` is the base contract for performing EOA deploys, providing a `Deployment` struct. The entry function is `deploy()` and requires adding `-s "deploy(string memory)" "<path_to_config>"` as a `forge script` flag. Any inheriting contract is expected to inherit the `_deploy()` internal function and return a `Deployment[]` array, containing all deployed contract details.

The `Deployment` struct contains 3 fields:
* `name`: the name of the contract published
* `address`: the address to which the contract was published
* `envToUpdate`: the environment variable in the config file to update if relevant.
    * Leave this as an empty string if deploying an instance of a contract that doesn't need to be tracked long-term.

If you need to do something with an EOA other than deploy, the base `ConfigParser` provided by `zeus-templates` is the appropriate base contract.

### MultisigBuilder

`MultisigBuilder` is the base class for any action to be taken by a Safe Multisig. The entry function is `execute()` and requires adding `-s "execute(string memory)" "<path_to_config>"` as a `forge script` flag. Any inheriting contract is expected to inherit the `_execute()` internal function and return a `MultisigCall[]` object, containing all intended multisig calls.

The `MultisigCall` struct contains 3 fields:
* `to`: the address to call with the multisig
* `value`: the amount of ETH to send as part of the transaction.
* `data`: the calldata associated with the multisig call.

Once the `_execute()` function is implemented, the base MultisigBuilder contract will combine these calls into one `SafeTx` object, which is intended to pass all calls into the [MultiSendCallOnly](https://github.com/safe-global/safe-smart-account/blob/6fde75d29c8b52d5ac0c93a6fb7631d434b64119/contracts/libraries/MultiSendCallOnly.sol) contract.

## Setting up a Release Directory

### Naming

Every file must be named either `#-eoa.s.sol` or `#-multisig.s.sol`, where `#` represents the script's order in the release process and the choice of `(eoa|multisig)` represents the relevant signing strategy.

## Example: How to write a Deploy-Queue-Upgrade

Before diving into how to do this, a quick context blurb for those unfamiliar. If you are familiar, feel free to jump to the next subsection, [Deploy](#deploy).

### Context

A deploy-queue-execute flow is common given EigenLayer's [multisig governance](https://docs.eigenlayer.xyz/eigenlayer/security/multisig-governance) structure. For context, the Executor Multisig (a 1-of-2 multisig) owns many critical privileges on EigenLayer, such as the ability to upgrade core contracts. An EigenLayer upgrade will often originate from the Operations Multisig, controlled by Eigen Labs, which is 1 of the 2 signers. Well, in spirit at least.

Technically, there is an intermediate contract between the Ops Multisig and Executor known as the Timelock into which actions must be queued for a waiting period (~10 days). After the delay, the Operations Multsig must then poke the Timelock to forward the transaction to the Executor, which will finally take action.

### Deploy

Deploy scripts are expected to inherit the `EOADeployer` script and produce a `Deployment[]` array containing all deployed contract names, addresses, and overridden config values.

An example deploy script looks as follows:

```solidity
pragma solidity ^0.8.12;

import "zeus-templates/templates/EOADeployer.sol";

... // imagine imports here

contract DeployEigenPodManager is EOADeployer {

    function _deploy(Addresses memory, Environment memory, Params memory params) internal override returns (Deployment[] memory) {

        vm.startBroadcast();

        address newEigenPodManager = new EigenPodManager(
            ... // imagine params here
        );

        _deployments.push(Deployment({
            name: type(EigenPodManager).name,
            deployedTo: newEigenPodManager,
            envToUpdate: "eigenPodManager.pendingImpl"
        }));

        vm.stopBroadcast();

        return _deployments;
    }
}
```

Here, the script inherits the `EOADeployer` and implements the `_deploy()` function, creating a new `EigenPodManager` with relevant details and requesting for Zeus to set the `pendingImpl` field of the `eigenPodManager` in the config to the given address. It would be pending because a transaction must be issued to upgrade the EigenPodManager to this new implementation.

### Queue

Once the above `eigenPodManager` is deployed, a `queueTransaction` operation must be taken in the Timelock if upgrading via the Ops Multisig. This could look like something below:

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/templates/OpsTimelockBuilder.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IUpgradeableBeacon} from "script/interfaces/IUpgradeableBeacon.sol";
import "src/contracts/pods/EigenPodManager.sol";

contract QueueEigenPodAndManager is OpsTimelockBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] internal _executorCalls;

    function queue(Addresses memory addrs, Environment memory env, Params memory params) public override returns (MultisigCall[] memory) {

        // construct initialization data for eigenPodManager
        bytes memory eigenPodManagerData = abi.encodeWithSelector(
            EigenPodManager(addrs.eigenPodManager.pendingImpl).initialize.selector,
            ... // imagine initialization data
        );

        // upgrade eigenPodManager
        _executorCalls.append({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgradeAndCall.selector,
                addrs.eigenPodManager.proxy,
                addrs.eigenPodManager.pendingImpl,
                eigenPodManagerData // initialize impl here
            )
        });

        // upgrade eigenPod beacon implementation
        _executorCalls.append({
            to: addrs.eigenPod.beacon,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            )
        });

        return _executorCalls;
    }
}
```

After inheriting the `OpsTimelockBuilder` contract (which is an extension of `MultisigBuilder`), we implement the `queue()` function from the perspective of the Executor Multisig. All calls are eventually encoded as a `SafeTx` from the Ops Multisig to the Timelock.

### Upgrade

Once the delay has passed, `executeTransaction` can be called on the timelock with the appropriate calldata. As an example:

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/templates/OpsTimelockBuilder.sol";

import "src/contracts/pods/EigenPodManager.sol";

import "./2-multisig.s.sol"; // using previous script to avoid rewriting

contract ExecuteEigenPodManager is MultisigBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for *;

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        QueueEigenPodManager queue = new QueueEigenPodManager();

        MultisigCall[] memory _executorCalls = queue.queue(addrs, env, params);

        // steals logic from queue() to perform execute()
        // likely the first step of any _execute() after a _queue()
        bytes memory executorCalldata = queue.makeExecutorCalldata(
            _executorCalls,
            params.multiSendCallOnly,
            addrs.timelock
        );

        // execute queued transaction upgrading eigenPodManager
        _multisigCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });

        return _multisigCalls;
    }
}
```

After reusing the previous script's `queue()` function to recreate the calldata, we wrap it up in an `executeTransaction` encoded data blob and return the `MultisigCall`, which will then be crafted as a `SafeTx` by the parent contract.