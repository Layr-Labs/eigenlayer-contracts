# EigenLayer Contracts

EigenLayer is a set of smart contracts deployed on Ethereum that enable restaking of assets to secure new services. This repo contains the EigenLayer core contracts supporting beacon chain ETH and liquid staking tokens (LSTs).

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
  - `core/` - DelegationManager, StrategyManager, AVSDirectory, Slasher, RewardsCoordinator
  - `pods/` - EigenPodManager, EigenPod, DelayedWithdrawalRouter
  - `strategies/` - StrategyBase, StrategyFactory, EigenStrategy
  - `token/` - Eigen, BackingEigen tokens
  - `permissions/` - PauserRegistry
- `src/test/` - Test files
  - `integration/` - Integration tests showing user interactions
- `script/` - Deployment and utility scripts
- `docs/` - Technical documentation
- `pkg/` - Go bindings

## Key Branches

- `main` - Work-in-progress for upcoming releases (default)

## Testing Notes

- Fork tests require `RPC_MAINNET` environment variable
- See `.env.example` for environment setup
- Use `source source-env.sh [goerli|local]` with yq installed for forked environment tests
