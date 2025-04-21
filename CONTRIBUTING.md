# Contribution Guide

Thank you for considering contributing to EigenLayer! This guide will help you set up your development environment and submit high-quality pull requests (PRs).

## Setting Up Your Environment

We use fork base PR mechanism. Contributions that do not follow our fork base PR practices will be automatically immediately closed and deleted, preventing branch pollution, keeping our repository clean, tidy, more readable and searchable.

### Fork and Clone the Repository

1. Fork the repository following [GitHub's instructions](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo)
2. Clone your forked repository:

```bash
git clone https://github.com/<your_github_id>/eigenlayer-contracts.git
cd eigenlayer-contracts
```

### Install Dependencies

To set up this repo for the first time, run:

```bash
make deps
```

This will:
* Install the pre-commit hook
* Install foundry and its tools
* Install abigen


See the [Foundry docs](https://book.getfoundry.sh/) for more info on installation and usage of foundry. If you already have it, you can build this project with these commands:

```
foundryup

forge build
```


### Run Unit and Integration Tests

Verify your setup by running the test suite:

```bash
forge test
```

Congratulations! You've successfully set up your repository and run the tests.


### Run Fork Tests

We have a few fork tests against ETH mainnet. Passing these requires the environment variable `RPC_MAINNET` to be set. See `.env.example` for an example. Once you've set up your environment, `forge test` should show these fork tests passing.

Additionally, to run all tests in a forked environment, [install yq](https://mikefarah.gitbook.io/yq/v/v3.x/). Then, set up your environment by running the following command.

`source bin/source-env.sh [local]`

Then run the tests:

`forge test --fork-url [RPC_URL]`


### Run Static Analysis

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


### Generate updated Storage Report

To update the storage reports in `/docs/storage-report` run:

```bash
make storage-report
```

------



## Submitting a Pull Request


### Configure Remote Repositories

All PRs should be submitted from your forked repository. Add the official EigenLayer repository as your upstream remote:

```bash
git remote add upstream https://github.com/Layr-Labs/eigenlayer-contracts.git
```

Verify your configuration:

```bash
git remote -v
```

You should see:
```
origin    https://github.com/<your_github_id>/eigenlayer-contracts.git (fetch)
origin    https://github.com/<your_github_id>/eigenlayer-contracts.git (push)
upstream  https://github.com/Layr-Labs/eigenlayer-contracts.git (fetch)
upstream  https://github.com/Layr-Labs/eigenlayer-contracts.git (push)
```


### Create Your PR to Upstream Branches

The following applies if you submit a PR to upstream branch of either `main` or `release-dev/*`, which will refer to as `<upstream-target-branch>`

Create a local branch from the upstream branch:

```
git checkout -B <my-local-feature-branch> upstream/<upstream-target-branch>
```

Make your changes, then commit:

```
git add .
git commit -m "your message"
```

Push to your fork:

```
git push -u origin <my-local-feature-branch>
```

Go to GitHub, and open a PR from:

```
your-fork:<my-local-feature-branch>
**to**
Layr-Labs/eigenlayer-contracts:<upstream-target-branch>
```

GitHub should auto-detect this if branches match.

If you already have an open PR for this branch, just push more commits to `<my-local-feature-branch>`, and the PR will update.


### Keeping Your Branch Updated With Upstream

To stay in sync with upstream's `<upstream-target-branch>`, whether it's `main` or `release-dev/*`:

```
git fetch upstream
git checkout <my-local-feature-branch>
git rebase upstream/<upstream-target-branch>

# or use merge if you prefer
# git merge upstream/<upstream-target-branch>

git push -f origin <my-local-feature-branch>
```

Use --force only if you're rebasing (not with merges).


### PR Title Format

Format your PR title as follows:

```
<type>: <subject>
```

Where:
- **Type**: Indicates the nature of the change
- **Subject**: A brief description of the change

Use one of these types:

| Type | Description |
|------|-------------|
| **feat** | A new feature |
| **fix** | A bug fix |
| **docs** | Documentation-only changes |
| **style** | Changes that don't affect code meaning (formatting) |
| **refactor** | Code changes that neither fix bugs nor add features |
| **ci** | Changes to continuous integration configuration |
| **perf** | Performance improvements |
| **test** | Adding or correcting tests |
| **chore** | Changes to build process or auxiliary tools |
| **release** | Merging a release-dev branch to `main` |

Remember to set the appropriate PR label as well.

### PR Description Template

Use this template for your PR description:

```markdown
**Motivation:**
- Describe the context and reason for the change

**Modifications:**
- Detail the specific changes made

**Result:**
- Explain the outcome or effects of the change
```

#### Example PR:

```markdown
feat: Implement new withdrawal flow in StrategyManager

**Motivation:**
The current withdrawal process is inefficient and leads to delays for users.

**Modifications:**
- Refactored the `withdraw` function in `StrategyManager.sol`
- Updated associated unit tests to reflect changes
- Modified documentation to explain the new withdrawal process

**Result:**
The withdrawal process is now more efficient, reducing wait times for users.
```

## Testing and Continuous Integration

- All PRs should include appropriate test coverage
- Write unit and integration tests that verify your changes
- Ensure all existing tests continue to pass

For first-time external contributors: A maintainer will review your PR and manually approve CI to run.

By following these guidelines, you help maintain the quality and readability of the EigenLayer codebase. We appreciate your contributions and look forward to collaborating with you!
