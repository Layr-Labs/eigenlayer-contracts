# Multichain Deployer Setup

This directory contains scripts for deploying and configuring the multichain deployer multisig on destination chains. The multichain deployer is responsible for deploying EigenLayer contracts across multiple blockchain networks.

## Overview

The deployment process consists of two phases:

1. **Deploy Multichain Deployer** (`1-deployMultichainDeployer.s.sol`): Deploys a Safe multisig contract at a deterministic address
2. **Transfer Ownership** (`2-transferDeployerOwnership.s.sol`): Adds additional owners and updates the threshold for security

*Note: The initialOwner (`0x792e42f05E87Fb9D8b8F9FdFC598B1de20507964`) is not removed from the Multisig. It can be done in a separate tx*

## Deployment Details

### Safe Multisig Configuration
- **Expected Address**: `0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9`
- **Safe Version**: 1.4.1
- **Final Threshold**: 3 (after ownership transfer)
- **Final Owners**: `0x792e42f05E87Fb9D8b8F9FdFC598B1de20507964` + Ops msig signers (see `owners.toml`)