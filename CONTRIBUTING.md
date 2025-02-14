# Contribution Guide

Thank you for considering contributing to EigenLayer! To ensure a smooth and efficient collaboration, please adhere to the following guidelines when submitting a pull request (PR):

## Requirements

- Foundry  
- Git  
- Node.js  

## Setup Repo

```bash
git clone https://github.com/Layr-Labs/eigenlayer-contracts.git
```

### Install Dependencies

```bash
npm install
npx husky install
forge install
```

### Environment Variables

Some of the tests and features of this repository require environment variables to be set up.

Copy the .env.example file to create a .env and populate it with the appropriate environment values that you have control over.

## Development Workflow

- **Testing**: Ensure that your code is well-tested. Write unit and integration tests as appropriate.
- **Branching**: Create a feature branch from the `dev` branch for your work. The `dev` branch contains the most up-to-date code for upcoming releases.
- **Pull Requests**: Once your feature or fix is complete and tested, submit a PR to merge your feature branch into the `dev` branch.
- **Code Review**: After your PR is approved, squash your commits into a single commit before merging to maintain a clean and concise commit history.

## Commit Message Format

When squashing and merging your commits, please follow the commit message convention below:

```
<type>(<scope>): <subject> (#<PR number>)

<motivation>

<modifications>

<result>
```

- **Type**: Indicates the nature of the change.
- **Scope**: Specifies the section or module of the codebase affected (optional).
- **Subject**: A brief description of the change.
- **PR Number**: Reference to the pull request associated with this change.
- **Motivation**: Describe the context and reason for the change.
- **Modifications**: Detail the specific changes made.
- **Result**: Explain the outcome or effects of the change.

### Example

An example of a commit message:

```
feat(StrategyManager): Implement new withdrawal flow (#123)

**Motivation:**

The current withdrawal process is inefficient and leads to delays for users.

**Modifications:**

- Refactored the `withdraw` function in `StrategyManager.sol`.
- Updated associated unit tests to reflect changes.
- Modified documentation to explain the new withdrawal process.

**Result:**

The withdrawal process is now more efficient, reducing wait times for users.
```

### Commit Types

Use the following types for your commit messages:

- **feat**: A new feature.
- **fix**: A bug fix.
- **docs**: Documentation-only changes.
- **style**: Changes that do not affect the meaning of the code (e.g., formatting).
- **refactor**: A code change that neither fixes a bug nor adds a feature.
- **perf**: A code change that improves performance.
- **test**: Adding missing or correcting existing tests.
- **chore**: Changes to the build process or auxiliary tools and libraries.

By following these guidelines, you help maintain the quality and readability of the EigenLayer codebase. We appreciate your contributions and look forward to collaborating with you!

## Building and Running Tests

This repository uses Foundry. See the [Foundry docs](https://book.getfoundry.sh/) for more info on installation and usage. If you already have foundry, you can build this project and run tests with these commands:

```
foundryup

forge build
forge test
```

### Contributor Setup

To set up this repo for the first time, run:

```bash
make deps
```

This will:
* Install the pre-commit hook
* Install foundry and its tools
* Install abigen

### Running Fork Tests

We have a few fork tests against ETH mainnet. Passing these requires the environment variable `RPC_MAINNET` to be set. See `.env.example` for an example. Once you've set up your environment, `forge test` should show these fork tests passing.

Additionally, to run all tests in a forked environment, [install yq](https://mikefarah.gitbook.io/yq/v/v3.x/). Then, set up your environment by running the following command.

`source bin/source-env.sh [goerli|local]`

Then run the tests:

`forge test --fork-url [RPC_URL]`

### Running Static Analysis

1. Install [solhint](https://github.com/protofire/solhint), then run:

`solhint 'src/contracts/**/*.sol'`

2. Install [slither](https://github.com/crytic/slither), then run:

`slither .`

### Generate Inheritance and Control-Flow Graphs

1. Install [surya](https://github.com/ConsenSys/surya/) and graphviz:

```
npm i -g surya

apt install graphviz
```

2. Then, run:

```
surya inheritance ./src/contracts/**/*.sol | dot -Tpng > InheritanceGraph.png

surya mdreport surya_report.md ./src/contracts/**/*.sol
```

### Generate Go bindings

```bash
make bindings
```