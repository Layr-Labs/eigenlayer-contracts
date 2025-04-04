# Maintenance Guide

This document outlines essential procedures for repository maintenance, including PR management, branching strategies, and release processes that maintainers should follow to ensure code quality and stability.

## Code Review Process

Code reviewers are responsible for ensuring that PRs meet the following criteria before merging:

- **Properly labeled** according to our labeling conventions
- **Well-formatted** code adhering to our style guidelines
- **Comprehensive test coverage** including unit and integration tests
- **Documentation updates** where appropriate
- **Security considerations** addressed and potential vulnerabilities mitigated

For first-time external contributors, a maintainer must:
1. Review the PR with additional scrutiny
2. Verify the contributor has signed the CLA (Contributor License Agreement)
3. Manually approve CI workflows to run
4. Provide detailed feedback to help onboard the contributor


## Branching and Merge Strategy for Development

### Core Branches

- **`main` (default)**: 
  - The canonical branch where all approved code changes are merged
  - Must remain stable and suitable for internal/external testing
  - All CI workflows must pass on every commit
  - Security audits are performed against specific commits on this branch
  - All deployments originate from this branch

- **`release-dev/xxx`**:
  - Used for large features requiring multiple commits that introduce breaking changes
  - Examples: `release-dev/slashing-v1`, `release-dev/rewards-v2`
  - Branch from `main` for isolated development
  - Must be regularly rebased against `main` to reduce merge conflicts
  - Merged back into `main` as a single logical unit when complete
  - No security audits are performed directly on these branches


### Merge Strategies

We employ different merge strategies based on PR type:

#### 1. Regular PRs
Regular PRs include bug fixes, small features, or incremental updates where each PR represents an atomic change.

**Merge Strategy**: Always use **Squash and Merge**
- Combines all commits into a single, atomic unit
- Maintains a clean, linear history on the target branch
- Preserves PR number in commit message for traceability
- Example: `PR #123 -> [squash] -> main`

#### 2. Release-Dev Merges
When merging a complete `release-dev` branch into `main` after a major feature is completed.

**Merge Strategy**: Always use **Merge Commit (no squash)**
- Preserves the complete commit history of the feature branch
- Creates a merge commit that serves as a feature completion marker
- Simplifies cherry-picking specific changes for hotfixes
- Example: `release-dev/feature -> [merge commit] -> main`


### Best Practices

- **Branch Hygiene**: Delete branches after merging to avoid clutter
- **CI Validation**: Never merge code that fails CI checks
- **Documentation**: Update documentation alongside code changes
- **Breaking Changes**: Clearly document any breaking changes in PR descriptions



## Release Management


### Environment Networks

Before diving into maintenance details, it's important to understand our environment network structure.

| Environment | Environment Network | Definition and Practices |
|-------------|---------------------|--------------------------|
| **mainnet** | mainnet-ethereum | - Production environment on Ethereum mainnet<br>- Single canonical mainnet instance in foreseeable future<br>- External facing with highest security requirements |
| **testnet** | testnet-holesky<br>testnet-sepolia<br>testnet-hoodi | - Test environments on Ethereum testnets (holesky, sepolia, etc.)<br>- Multiple testnet environments may exist simultaneously<br>- External facing, long-lived environments for both internal and external developers<br>- Used for integration testing, AVS onboarding, and community engagement |
| **preprod** | preprod-holesky<br>preprod-sepolia<br>preprod-hoodi | - Pre-production environments on Ethereum testnets<br>- Multiple preprod environments may exist simultaneously<br>- Internal facing for development team use<br>- Used for feature verification before promoting to testnet<br>- Provides safe environment for testing breaking changes |

For more details on our testing infrastructure strategy, see our blog post: [The Future of EigenLayer Testing: New & Improved Testnets & Tooling Coming Soon](https://www.blog.eigenlayer.xyz/the-future-of-eigenlayer-testing-new-and-improved-testnets-tooling-coming-soon/)


### Base Branch For Environment Networks

Following are base branches to corresponding env networks

- `main` ⇒ canonical branch
- `preprod` ⇒ preprod env network
- `testnet` ⇒ testnet env network
- `mainnet` ⇒ mainnet env network

Specific release on a given env network will be forked from these base branches accordingly.


### Release Tags and Semver

This section defines release tag.

format: `<semver>_<env_network>` where `<semver>` = `<major>.<minor>.<patch>`

E.g.
- v1.1.0_mainnet.ethereum
- v1.1.0_testnet.holesky
- v1.1.0_testnet.sepolia

`<semver>` should follow [best practices of semver definition](https://semver.org/)


In observing semver, we define the following;

1. A **major** upgrade
    1. An ideologically breaking upgrade, which may require re-learning parts of the system. 
    2. “Upgrade Required” — Major upgrades are **assumed to break** all tooling, integrations, and downstream consumers unless otherwise stated.
    3. Major upgrades are announced ahead of time via blog posts, the ELIP process, and onchain via the Eigen timelock.
    4. example: slashing, which changes most of our codebase
2. A **minor** upgrade
    1. A conceptually breaking upgrade, which require adapting break changes and APIs, upgrading tools, CLIs, or integrations with EigenLayer.
    2. “Upgrade carefully” — Minor upgrades are **assumed not to break** any tooling, CLIs, or interactions with EigenLayer unless otherwise stated in release notes.
3. A **patch** upgrade
    1. A “silent” upgrade, which addresses a bug or inconsistency in the system, but does not change any interfaces.
        1. i.e. fixes reveal the need for a rounding fix in a contract. Notably **abi** and **tooling** remains stable across patches.
    2. “Upgrade suggested” — Patch upgrades may come with upgrades to tooling, CLIs, or other Eigen infrastructure. You are encouraged to read the patch notes and upgrade.


### Release and Deployment Scope

- a release would upgrade versions for all code, even though there's no change to a file from last release
- a deployment would adopt a release as a whole, so all contracts would upgrade to that release, instead of "some file on release v1.3 and some on release v1.4"
    - there's no phrasing of, for example, "all non-EigenPod contracts be on `v1.3.0` and  EigenPods contract on `v1.4.0`" - they would all be on v1.4.0, just the non-eigenpod contract are not changed at all from v1.3.0
    - as implementation detail, we can skip upgrade the contracts that have not changed, but the codebase and releases are on new release version
- all changes have to be deployed to testnet first, before mainnet. No skipping from preprod to mainnet directly


### Release by Cherry-Pick
 
We need a branching model that supports promoting features from `main` to `preprod` to`testnet` and then to `mainnet`, while allowing `mainnet` to have features in a different order than `testnet`. A flexible approach is to use **feature branches** and **cherry-picking** to control what gets promoted.

Take testnet release and workflow for example  

- **`testnet`**
    - Rebase from `preprod`
        - note: fast forward merge is also ok in the case, and have same effect as rebase
    - Cut a release branch from `testnet`
- **`release/V<x.y.z>-testnet.<network>`**
    - Created from `testnet` and cherry-pick commits based on features to include/exclude

The same applies to preprod and mainnet release

We admit that there will always be edge cases in above scenario which cannot be resolved easily, especially in a large release like `slashing` which touches all code base. In that case, we have to spend extra efforts by leveraging release manager


## Release Manager Responsibilities


------

By following these maintenance procedures, we ensure a stable, secure, and well-documented codebase that meets the high standards required for EigenLayer's infrastructure.
