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
  - All CI workflows must pass before and after every commit
  - Security audits are performed against specific commits on this branch

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


#### 3. Merge multiple release-dev branches

A release may contain multiple features that each is developed on its own `release-dev/*` branch. 

When merging to `main`, atomic commits on each `release-dev/*` branch should be preserved, in order to track commit history of each individual feature scope in the release.

This can be achieved by either merge commit each `release-dev/*` branch into `main` respectively, or merge commit into a single `release-dev/*` branch (consolidation) then commit merge the single one into `main`.


### Best Practices

- **Branch Hygiene**: Delete branches after merging to avoid clutter
- **CI Validation**: Never merge code that fails CI checks
- **Documentation**: Update documentation alongside code changes
- **Breaking Changes**: Clearly document any breaking changes in PR descriptions


### Auto Remove Stale Branches

Many stale branches lives in the repo and never got cleaned up. to balance developer convenience vs being organized/legible, we have implemented a CI workflow to identify and remove stable branches.

The CI workflow will:

- scan stale branches that not hv new commits since 90 days ago, and notify owner if there's one
- if still no changes coming in, delete the stable branch 7 days later

Following branches and regex are excluded

- `main`
- `release-dev/*` (release development branches)
- `v*.*.*` (release branches)

Regex rules can be changed by modifying `exempt-branches-regex` config in the [yml](https://github.com/Layr-Labs/eigenlayer-contracts/blob/main/.github/workflows/remove-stale-branches.yml)


------


## Release Management


### Environment Networks

Before diving into maintenance details, it's important to understand our environment network structure.

| Environment | Environment Network | Definition and Practices |
|-------------|---------------------|--------------------------|
| **mainnet** | mainnet-ethereum | - Production environment on Ethereum mainnet<br>- Single canonical mainnet instance in foreseeable future<br>- External facing with highest security requirements |
| **testnet** | testnet-holesky<br>testnet-sepolia<br>testnet-hoodi | - Test environments on Ethereum testnets (holesky, sepolia, etc.)<br>- Multiple testnet environments may exist simultaneously<br>- External facing, long-lived environments for both internal and external developers<br>- Used for integration testing, AVS onboarding, and community engagement |
| **preprod** | preprod-holesky<br>preprod-sepolia<br>preprod-hoodi | - Pre-production environments on Ethereum testnets<br>- Multiple preprod environments may exist simultaneously<br>- Internal facing for development team use<br>- Used for feature verification before promoting to testnet<br>- Provides safe environment for testing breaking changes |

For more details on our testing infrastructure strategy, see our blog post: [The Future of EigenLayer Testing: New & Improved Testnets & Tooling Coming Soon](https://www.blog.eigenlayer.xyz/the-future-of-eigenlayer-testing-new-and-improved-testnets-tooling-coming-soon/)


### Release Tags and Semver

This section defines release tag semver where `<semver>` = `<major>.<minor>.<patch>`

E.g.
- v1.1.0
- v1.1.1

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


### Release Base and Cherry Pick

All releases will be cut from `main` the base branch, to ensure global consistency, clean legibility, and high code quality with no env-network specific code

Each env_network deployment will pick one release above.


During release, release manager can cherry pick to exclude certain commits, especially in case of a non-sequential feature deployment


#### Non Sequential Feature Release

In smart contracts, there is a common scenario for non sequential feature deployment due to Ethereum hardforks, whose time we cannot control.

E.g. say we have features A, B, and Ethereum hardfork X,  where B is an upgrade to adapt to X. The deployment sequence for testnet is A, B, X. While that of mainnet has to be B, X, A, because 1/ we don’t control the time of X, 2/ B must happen before X, and 3/ A is not ready for mainnet yet which can’t be deployed before B

Concrete example just happened is A = slashing V1, B = prooftra, X = pectra

So our release branch and model must be flexible enough to solve it via cherry picking


### Release and Deployment Scope

- a release would upgrade versions for all code, even though there's no change to a file from last release
- a deployment would adopt a release as a whole, so all contracts would upgrade to that release, instead of "some file on release v1.3 and some on release v1.4"
    - there's no phrasing of, for example, "all non-EigenPod contracts be on `v1.3.0` and  EigenPods contract on `v1.4.0`" - they would all be on v1.4.0, just the non-eigenpod contract are not changed at all from v1.3.0
    - as implementation detail, we can skip upgrade the contracts that have not changed, but the codebase and releases are on new release version
- all changes have to be deployed to testnet first, before mainnet. No skipping from preprod to mainnet directly


### Changelog, Release Note, Release Announcement Blog

Each release in testnet and mainnet must have corresponding release note, changelog, and announcement blog

- Changelog
    - lives in the repo `/CHANGELOG` dir
    - exact commit history diff from last release
    - See examples in Kubernetes https://github.com/kubernetes/kubernetes/tree/master/CHANGELOG
- Release note:
    - lives in the repo as part of release description under “[Releases](https://github.com/Layr-Labs/eigenlayer-contracts/releases)”
    - high level summary of key changes users need to be aware of
    - See template below
    - Owner: release manager and PM
- Release announcement blog
    - published to [blog.eigenlayer.xyz](http://blog.eigenlayer.xyz)
    - an announcement with polished language to introduce the release to users. It will include content in both release notes and changelog
    - See examples in Golang https://tip.golang.org/doc/go1.24
    - Owner: release manager and PM

Release Note must follow template:

```jsx
<release-version>

🚀 New Features – Highlight major new functionality.
	- ...
	- ...

⛔ Breaking Changes – Call out backward-incompatible changes.
	- ...
	- ...

📌 Deprecations – Mention features that are being phased out.
	- ...
	- ...

🛠️ Security Fixes – Specify patched vulnerabilities.
	- ...
	- ...

🔧 Improvements – Enhancements to existing features.
	- ...
	- ...

🐛 Bug Fixes – List resolved issues.
	- ...
	- ...
```

When writing **release notes** and **changelog entries**, follow these best practices to ensure clarity, usefulness, and professionalism.

**Required Best Practices**

1. **Be Clear and Concise** – Use simple, direct language. Avoid jargon unless necessary.
2. **Categorize Changes** – Use standard sectioned template above.
3. **Provide Context** – Briefly explain why the change was made, not just what changed.
4. **Use a Consistent Format** – Follow a structured template for every release.
5. **Include Relevant Links** – Link to issue trackers, PRs, or documentation for more details.
6. **Highlight User Impact** – Mention how the update affects users and any required actions.
7. **Keep It Versioned** – Use semantic versioning (e.g., `v2.3.1`) and list versions in descending order.
8. **Avoid Internal-Only Details** – Don't include internal project discussions that aren’t relevant to users.
9. **Maintain a Single Source of Truth** – Store changelogs in a `CHANGELOG.md` file or an equivalent documentation system.


## Release Manager Responsibilities

We introduce a new critical role called release manager (RM)

Each release should have a primary RM and a secondary RM. RM is on rotating base, not necessarily always be the tech lead of the project, so we can amortize the cost and onboard all developers with same standards

Primary Release Manager is responsible for

- Managing the release-dev branch during development phase
    - create it from `main` and eventually merge back to `main`
    - constantly rebasing to `main` to resolve conflict early, without required efforts from other developers working on the release-dev branch
    - holding high bar for quality:
        - code quality
        - PR templating
        - test quality (e.g. no flaky or failed tests, keep test runtime under control)
- Manage release branches and deployment
    - create and manage corresponding release branches `Vx.y.z`
    - author and publish release notes and changelog
    - author and publish release blog with PM and marketing team
    - lead communication with marketing and BD team to inform stakerholders for preparation
    - lead deployment to env network
- Mentoring and knowledge sharing
    - mentor and train secondary release manager, who will be the primary manager for next release


By following these maintenance procedures, we ensure a stable, secure, and well-documented codebase that meets the high standards required for EigenLayer's infrastructure.
