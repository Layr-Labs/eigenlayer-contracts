# EigenLayer Release Scripts

This document describes the various release scripts available in the `script/releases/` directory.

## Multichain Deployer Scripts

### 1. multichain-deployer-mainnet
Deploys a multichain deployer multisig to mainnet environments (e.g., Base, Ethereum Mainnet). 
- Initializes the multisig to be 3/7
- 6 of the 7 signers come from the ops multisig

### 2. multichain-deployer-testnet-preprod
Deploys a multichain deployer multisig to testnet and preprod environments.
- Configured as a 1/n multisig
- Has 0xDA as the signer

## Destination Chain Deployment Scripts

### 3. v1.6.0-destination-genesis-mainnet
Deploys foundational contracts to a chain (useful for destination chains):
- Proxy Admin
- Ops Multisig
- Pauser Multisig

### 4. v1.6.0-destination-governance-mainnet
Deploys governance infrastructure to a chain (useful for destination chains):
- Timelock Controller (with 1 day delay)
- Protocol Council Multisig
- Community Multisig
- Executor Multisig
- The protocol council and community multisigs are initialized to be the mainnet ops multisig signers with 3/n quorum

### 5. v1.6.0-destination-governance-testnet
Same as the mainnet governance deployment above, but with testnet-specific configurations:
- Timelock delay set to 1 second
- The owner of the multisig is the 0xDA address
- Threshold is 1

## Protocol Deployment Scripts

### 6. v1.6.0-protocol-from-scratch
Deploys the entire EigenLayer protocol from scratch to v1.6.0, including:
- All core contracts
- Governance infrastructure
- Token contracts
- Strategy contracts

### 7. v1.7.0-v1.8.0-multichain-hourglass-combined
Upgrades the protocol to support:
- Multichain functionality
- Hourglass features
