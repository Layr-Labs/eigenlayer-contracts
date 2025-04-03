# Maintenance

This document outlines basic procedures for maintenance of PRs, and releases that maintainers of the repository should follow.


## Enrivonment Network

Before getting into maintenance details, let's define environment network first.

| Environment | Environment Network | Definition and Practices |
|-------------|---------------------|--------------------------|
| mainnet | mainnet-ethereum | - mainnet production env on Ethereum<br>- only one single mainnet in foreseeable future<br>- external facing |
| testnet | testnet-holesky<br>testnet-sepolia<br>testnet-hoodi | - testnet envs on Ethereum testnets like holesky or sepolia<br>- can be 1 or more testnet env networks for EigenLayer<br>- external facing, long lived for both internal and external developers to test on |
| preprod | preprod-holesky<br>preprod-sepolia<br>preprod-hoodi | - preprod envs on Ethereum testnets like holesky or sepolia<br>- can be 1 or more preprod env networks for EigenLayer<br>- long lived for internal development and testing, before moving to testnets |

See blog [The Future of EigenLayer Testing: New & Improved Testnets & Tooling Coming Soon](https://www.blog.eigenlayer.xyz/the-future-of-eigenlayer-testing-new-and-improved-testnets-tooling-coming-soon/) for detailed motivation



##  Development

### Code Review

Code reviewers take responsibility of making sure PRs are rigorously labeled, well formatted, well tested before merging them. 

For first-time external contributors, a maintainer should review their PR carefully and then manually approve CI to run.


### Branching Model

#### Development Branches

- `main (default)`: 
    - The primary development and canonical branch where all new code are merged
    - It should remain stable enough for internal and external testing, ensuring continuous integration (CI) runs smoothly on every commit
    - Security audit will always happen on main by a given commit
- `release-dev/xxx`:
    - For large features with many commits that will be breaking changes and we want to all-or-none in `main`, maintainer should create a `release-dev/xxx` branch to develop on
        - eg. release-dev/slashing-v1, release-dev/rewards-v2
    - These branches branch off from `main` for development, should be constantly rebased to `main` to be compatible with canonical branch, and will be merged back into `main` when the development is done and ready to deploy
    - release-dev branches should be merged into main for auditing, no auditing will be on release-dev branches


### Merge PRs

We classify PRs into two types:

1. **Regular PRs** – Used for bug fixes, small changes, or incremental updates. Each PR is atomic itself.
2. **Release-dev Merges** – Merging a release-dev branch into `main` after completing a major release. The release-dev branch and such a merge contains many regular PRs

We use different merge rules for each PR type.

- For regular PRs:  
    - Always **rebase, squash, and merge**. Since multiple commits in a PR represent a single unit of work, merging them separately can be unnecessary or even break the code. Squashing ensures atomicity and keeps the commit history clean.
    - demonstration: PR -> squash -> `main` or `release-dev`


- For release-dev merges:
    - Always **rebase and create a new commit to merge, never squash**
    - All commit history should be preserved - a release-dev branch contains multiple independent units of work, and squashing would create a massive, monolithic commit that erases history, making it difficult to track and understand individual changes.
    - The additional commit will act as a marker in `main` as end of a release-dev merge. With each release-dev merge having its own end marker, it will significantly simplify the cherry-picking process in release creation
    - demonstration: `release-dev` -> merge -> `main`



## Releases



## Release Manager


