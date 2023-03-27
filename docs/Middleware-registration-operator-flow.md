# Middleware Registration Operator Flow

EigenLayer is a three-sided platform, bringing together middlewares/services, operators, and stakers. Middlewares have an attached risk-reward profile, based on the risk of getting slashed while running their service, and the rewards that they offer operators taking on those risks. Operators have full control over the middlewares that they decide to register with and run, and have an aggregated risk-reward profile of their own, based on these choices and their personal reputation. All of a users' assets that are deposited into EigenLayer are delegated to a single operator.

Operators must thus make an informed decision as to which middlewares to run. This document is meant to give a high-level overview of the registration process that they must then follow in order to register with the middlewares that they have decided to operate.

## Registration Process

Any operator must go through the following sequence of steps:
- [Register as an operator](./EigenLayer-delegation-flow.md#operator-registration) in EigenLayer
- Get stakers to [delegate to them](./EigenLayer-delegation-flow.md#staker-delegation)
- For each middleware:
    - Opt into [slashing](./EigenLayer-tech-spec.md#slasher) by the middleware's [ServiceManager](../src/contracts/interfaces/IServiceManager.sol) contract
    - Make sure to respect any other preconditions set by the middleware — for [EigenDA](https://docs.eigenda.xyz/), these are registering a BLS key in the [BLSPublicKeyCompendium](](../src/contracts/interfaces/IBLSPublicKeyCompendium.sol) and having sufficient (delegated) stake in EigenLayer to meet the minimum stake requirements set by EigenDA
    - Call a function on one of the middleware's contracts to complete the registration — for EigenDA, this is registering on its [BLSRegistry](../src/contracts/interfaces/IBLSRegistry.sol)