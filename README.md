<a name="introduction"/></a>
# EigenLayer
EigenLayer (formerly 'EigenLayr') is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services.
At present, this repository contains *both* the contracts for EigenLayer *and* a set of general "middleware" contracts, designed to be reuseable across different applications built on top of EigenLayer.

Click the links in the Table of Contents below to access more specific documentation. We recommend starting with the [EigenLayer Technical Specification](docs/EigenLayer-tech-spec.md) to get a better overview before diving into any of the other docs.

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

`forge test -vv`

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
