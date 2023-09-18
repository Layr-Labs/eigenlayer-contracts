## Roles and Actors

This document describes the different roles and actors that exist in EigenLayer M2, and provides some insight into how they interact with M2's core components.

### Stakers

A **Staker** is any party who has assets deposited (or "restaked") into EigenLayer. Currently, these assets can be:
* Native beacon chain ETH (via the EigenPods subsystem)
* Liquid staking tokens: cbETH, rETH, stETH (via the Strategies subsystem)

Stakers can restake any combination of these. That is, a Staker may hold ALL of these assets, or only one of them.

Once they've deposited, Stakers can delegate their stake to an Operator via the `DelegationManager`, or they can become an Operator by delegating to themselves.

*Flows:*
* Depositing into EigenLayer
* Delegating to an Operator
* Withdrawing out of EigenLayer

*Unimplemented:*
* Stakers earn yield by delegating to an Operator as the Operator provides services to an AVS
* Stakers are at risk of being slashed if the Operator misbehaves

### Operators

An **Operator** is a user who helps run the software build on top of EigenLayer. Operators register in EigenLayer and allow Stakers to delegate to them, then opt in to provide various services built on top of EigenLayer.

*Flows:*
* Registering as an operator
* Opting in to a service
* Exiting from a service

*Unimplemented:*
* Operators earn fees as part of the services they provide
* Operators may be slashed by the services they register with (if they misbehave)
