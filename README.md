<a name="introduction"/></a>
# EigenLayer
EigenLayer (formerly 'EigenLayr') is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services.
At present, this repository contains *both* the contracts for EigenLayer *and* a set of general "middleware" contracts, designed to be reuseable across different applications built on top of EigenLayer.

Note that the interactions between middleware and EigenLayer are not yet "set in stone", and may change somewhat prior to the platform being fully live on mainnet; in particular, payment architecture is likely to evolve. As such, the "middleware" contracts should not be treated as definitive, but merely as a helpful reference, at least until the architecture is more settled.

Click the links in the Table of Contents below to access more specific documentation. We recommend starting with the [EigenLayer Technical Specification](docs/EigenLayer-tech-spec.md) to get a better overview before diving into any of the other docs.  For contracts addresses deployed on Goerli, click [here](https://github.com/Layr-Labs/eigenlayer-contracts/blob/master/script/output/M1_deployment_goerli_2023_3_23.json).
## Table of Contents  
* [Introduction](#introduction)
* [Installation and Running Tests / Analyzers](#installation)
* [EigenLayer Technical Specification](docs/EigenLayer-tech-spec.md)

Design Docs
* [Withdrawals Design Doc](docs/Guaranteed-stake-updates.md)
* [EigenPods Design Doc](docs/EigenPods.md)

Flow Docs
* [EigenLayer Withdrawal Flow](docs/EigenLayer-withdrawal-flow.md)
* [EigenLayer Deposit Flow](docs/EigenLayer-deposit-flow.md)
* [EigenLayer Delegation Flow](docs/EigenLayer-delegation-flow.md)
* [Middleware Registration Flow for Operators](docs/Middleware-registration-operator-flow.md)

<a name="installation"/></a>
## Installation and Running Tests / Analyzers

### Installation

`foundryup`

This repository uses Foundry as a smart contract development toolchain.

See the [Foundry Docs](https://book.getfoundry.sh/) for more info on installation and usage.

### Natspec Documentation

You will notice that we also have hardhat installed in this repo. This is only used to generate natspec [docgen](https://github.com/OpenZeppelin/solidity-docgen). This is our workaround until foundry [finishes implementing](https://github.com/foundry-rs/foundry/issues/1675) the `forge doc` command.

To generate the docs, run `npx hardhat docgen` (you may need to run `npm install` first). The output is located in `docs/docgen`

### Run Tests

Prior to running tests, you should set up your environment. At present this repository contains fork tests against ETH mainnet; your environment will need an `RPC_MAINNET` key to run these tests. See the `.env.example` file for an example -- two simple options are to  copy the LlamaNodes RPC url to your `env` or use your own infura API key in the provided format.

The main command to run tests is:

`forge test -vv`

### Run Tests on a Fork
Environment config is contained in config.yml.  Before running the following commands, [install yq](https://mikefarah.gitbook.io/yq/v/v3.x/).  Then set up the environment with this script:

`source source-env.sh [CHAIN]`

for example, on goerli: `source source-env.sh goerli`.  Currently options for `[CHAIN]` are `goerli`, `local`.  Then to run the actual tests:

`forge test --fork-url [RPC_URL]`

### Run Static Analysis

`solhint 'src/contracts/**/*.sol'`

`slither .`

### Generate Inheritance and Control-Flow Graphs

first [install surya](https://github.com/ConsenSys/surya/)

then run

`surya inheritance ./src/contracts/**/*.sol | dot -Tpng > InheritanceGraph.png`

and/or

`surya graph ./src/contracts/middleware/*.sol | dot -Tpng > MiddlewareControlFlowGraph.png`

and/or

`surya mdreport surya_report.md ./src/contracts/**/*.sol`
