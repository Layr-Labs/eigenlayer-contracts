# EigenLayer Contracts

**EigenLayer** is a protocol built on Ethereum that introduces Restaking, a primitive for app and service builders to make verifiable commitments to their users. EigenLayer brings together Restakers, Operators, and Autonomous Verifiable Services (AVSs) to extend Ethereum's cryptoeconomic security. The protocol supports permissionless security: EIGEN, Native ETH, LSTs, and ERC-20s.

## Build & Test Commands

```bash
# Initial setup (installs pre-commit hook, foundry, abigen)
make deps

# Build contracts
forge b

# Run tests
forge t

# Run unit tests
forge test --no-match-contract Integration

# Run integration tests
forge test --match-contract Integration

# Run fork tests
env FOUNDRY_PROFILE=forktest forge test --match-contract Integration

# Format 
make fmt

# Generate Go bindings
make bindings
```

## Project Structure

- `src/contracts/` - Core Solidity contracts
  - `core/` - DelegationManager, StrategyManager, AVSDirectory, AllocationManager, RewardsCoordinator
  - `pods/` - EigenPodManager, EigenPod
  - `strategies/` - StrategyBase, StrategyFactory, EigenStrategy
  - `token/` - Eigen, BackingEigen tokens
  - `permissions/` - PauserRegistry, PermissionController, KeyRegistrar
  - `multichain/` - CrossChainRegistry, OperatorTableUpdater, CertificateVerifiers
  - `avs/` - TaskMailbox and AVS-related contracts
- `src/test/` - Test files
  - `integration/` - Integration tests for cross-contract interactions
- `script/` - Deployment and utility scripts
- `docs/` - Technical documentation (detailed contract docs in `/docs`)
- `pkg/` - Go bindings

## Key Branches

- `main` - The canonical, most up-to-date branch with work-in-progress for upcoming releases
- `Vx.y.z` - Release branches matching EigenLayer releases, cut from `main` via cherry-picking
- `release-dev/xxx` - Development branches for large features (merged to `main` then deleted)

## Testing Notes

- Fork tests require `RPC_MAINNET` environment variable
- See `.env.example` for environment setup
- Use `source source-env.sh [goerli|local]` with yq installed for forked environment tests
- Check integration tests in `src/test/integration/` for examples of user interactions

## Development Workflow

1. Make Changes

2. Format 
make fmt

3. Update docs

4. Ensure Tests Pass
forge t

5. Commit Changes & Open PR
- All commits MUST follow conventional commit standard
- PRs must be named with (release|feat|fix|chore|docs|refactor|test|style|ci|perf): 