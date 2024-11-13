# Release Template

This template is an example of the deploy-queue-upgrade structure that can be used for executing EigenLayer upgrades via the Ops Timelock Multisig, as described [here](https://docs.eigenlayer.xyz/eigenlayer/security/multisig-governance).

Zeus can take this template and instantiate a new release template, allowing for quick setup of a common upgrade process.

Note that the names follow the syntax `#-<signer>`, where `#` is the sequence number within the release (i.e. the order in which actions are taken), and `<signer>` is the signing strategy to be used for that script.

## 1-eoa.s.sol

This contract allows for deployments to be broadcast. An EOA is expected to take these actions.

## 2-multisig.s.sol

This contract allows for actions to be written from the perspective of the Executor Multisig, which will then be wrapped for the Timelock, then for the Ops Multisig to execute.

## 3-multisig.s.sol

The final execution step can be written in this contract. It will reuse the logic from the previous step by importing the contract, then allow for additional calls to be appended from the perspective of the Ops Multisig.