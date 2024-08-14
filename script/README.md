## Release Scripting

Desired usecases:
* Creating a new release script from a template (e.g. `zeus new v0.4.2-pepe`)
* Being able to see whether a release script has been run for a given environment (`zeus status v0.4.2-pepe`)
    * Since release scripts are split into 3 parts (`deploy`, `queue`, `execute`), being able to see

`zeus run deploy pepe --sender "0x1234"`

`zeus run deploy pepe --live --ledger`

`zeus run queue pepe --live --ledger`
* https://docs.safe.global/sdk/api-kit
* For proposing txns to the Safe UI

`zeus status pepe`

### Creating a New Release Script

```
zeus new $VERSION $RELEASE_NAME
```

This command will generate a new release script based on `Release_Template.s.sol`. The new script is placed in the `/releases` folder and named accordingly (e.g. `zeus new v0.4.2-pepe` will create `releases/v0.4.2-pepe/script.s.sol`).

Release scripts look like this:

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./utils/Releasoor.s.sol";

contract Release_TEMPLATE is Releasoor {

    using TxBuilder for *;
    using AddressUtils for *;

    function deploy(Addresses memory addrs) internal override {
        // If you're deploying contracts, do that here
    }

    function queueUpgrade(Addresses memory addrs) internal override {
        // If you're queueing an upgrade via the timelock, you can
        // define and encode those transactions here
    }

    function executeUpgrade(Addresses memory addrs) internal override {
        // Whether you are using the timelock or just making transactions
        // from the ops multisig, you can define/encode those transactions here
    }
}
```

### Implementing Your Release Script

Release scripts have three functions which may or may not be used, depending on your needs. Generally, a release consists of three steps:
1. Deploying new implementation contracts
2. Queueing an upgrade to these contracts in the timelock
3. Executing the queued timelock upgrade after the timelock delay has passed

However, your release may not require deploying or queueing an upgrade in the timelock. For example, if all you want to do is have the ops multisig call `StrategyFactory.whitelistStrategies`, you don't need to deploy any contracts or queue a timelock transaction. In this case, you don't need to fill in the `deploy` or `queueUpgrade` methods - all you need to implement is `executeUpgrade`.

#### Deploying a New Contract

Typically, we deploy new contracts via EOA. As such, deploying new contracts can be done directly in the `deploy` method of your release script by invoking `vm.startBroadcast()`, deploying your contract(s), then invoking `vm.stopBroadcast()`.

**Important**: In order to keep our configs up to date, ensure your deploy scripts call `addrs.X.setPending(newImpl)`! As you run this script for various environments, this will automatically update each environment's config to track these new deployments as "pending" upgrades to the contract in question.

```solidity
function deploy(Addresses memory addrs) public override {
    vm.startBroadcast();

    // Deploy new implementation
    EigenPod newEigenPodImpl = new EigenPod(
        IETHPOSDeposit(params.ethPOS),
        IEigenPodManager(addrs.eigenPodManager.proxy),
        params.EIGENPOD_GENESIS_TIME
    );

    vm.stopBroadcast();

    // Update the current environment's config
    // (sets eigenPod.pendingImpl = newEigenPodImpl)
    addrs.eigenPod.setPending(address(newEigenPodImpl));
}
```

You can run your deploy script on a forked version of an existing environment with the following command:
* `make deploy <release name> <environment>`

e.g. `make deploy pepe preprod` will:
* Read `preprod.config` and populate [your script helpers](#script-helpers)
* Spin up an anvil fork of the preprod environment
* Deploy your contract(s) to the anvil fork
* Update `preprod.config` with the `pendingImpl` addresses (via `setPending`)
<!-- 
TODO - should config be updated if we know we're on a local fork? 

Ideally, yes -- because it's helpful to see a git diff that shows the config got updated
However, we don't want people committing pendingImpls for local tests...
 -->

After you've tested the script using a forked environment, you can run it on a live network with the following command:
* `make deploy <release name> <environment> --live`

e.g. `make deploy pepe preprod --live` will:
* Read `preprod.config` and populate [your script helpers](#script-helpers)
* Deploy your contracts to `preprod` using the configured RPC endpoint and a ledger connection
* Verify the deployed contracts on etherscan
* Update `preprod.config` with the `pendingImpl` addresses (via `setPending`)

#### Queueing a Timelock Transaction

The most common thing you'll need to do when deploying a release is using the ops multisig to queue transactions via the timelock. Defining transactions for this stage of the process can involve multiple contracts. Ultimately, the goal is to *define one or more contract calls that the executor multisig will carry out*.

However, the actual output of this part of the process will be the contract call made by the ops multisig to the timelock queueing a transaction that (when the delay has elapsed) can be executed to trigger the executor multisig to make the aforementioned calls. Whew.

Instead of trying to keep track of all of this, `queueUpgrade` can take care of this for you. All you need to define (and return) is:
1. The ETA (TODO different ETA between environments)
2. The contract calls you want the executor multisig to execute

```solidity
function queueUpgrade(Addresses memory addrs) public override returns (Tx[] memory executorTxns, uint eta) {
    Txs storage txs = _newTxs();
    eta = env.isMainnet() ? 12351235 : 0;

    txs.append({
        to: addrs.eigenPod.beacon,
        data: EncBeacon.upgradeTo(addrs.eigenPod.getPending())
    });

    txs.append({
        to: addrs.eigenPodManager.proxy,
        data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.getPending())
    });

    return (txs.toArray(), eta);
}
```

The above function defines two transactions for the executor multisig:
* A call to `eigenPod.beacon.upgradeTo()`, upgrading the `eigenPod` beacon
* A call to `proxyAdmin.upgrade()`, upgrading the `eigenPodManager` proxy

**Important**: Note that the implementation addresses for the `eigenPod` and `eigenPodManager` are fetched using `getPending()`, which gets the `pendingImpl` for both addresses (recall that this `pendingImpl` was set during the `deploy` part of this release script). If `getPending()` finds that no `pendingImpl` address has been set, it will revert. This prevents `queueUpgrade` from running if it has a dependency on `deploy`.

You can run the `queueUpgrade` script on a forked version of an existing environment with the following command:
* `make queue <release name> <environment>`

e.g. `make queue pepe preprod` will:
* Read `preprod.config` and populate [your script helpers](#script-helpers)
* Spin up an anvil fork of the preprod environment
* 
* Update `preprod.config` with the `pendingImpl` addresses (via `setPending`)

### Script Helpers
