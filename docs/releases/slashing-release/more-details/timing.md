# Timing

The timing and delays imposed in the slashing protocol are intricate and very important to enforcing the correct checks and balances. This document explains them.

## Slashing Requests

* Slashing requests are executed synchronously
* Other operatorSets(AVSs) the operator is registered for do not have their slashable stake affected when their operator is slashed. Each operatorSet has their own designated proportion of stake allocated to them by the discretion of the operator.

## Magnitude Updates

* Magnitude allocations have a default allocation delay of 21 days. An operator can perform a one-time initialization of this allocation delay to their desired value.
* Magnitude deallocations have a fixed delay of 17.5 days. This is to ensure a minimum slashability window for the AVSs and operatorSets the slashable stake is securing.
* There can be at most 1 pending allocation OR deallocation for a certain strategy, operator, and operatorSet at a time

## Withdrawals

* Withdrawals are completable 17.5 days after they're queued
* Withdrawals are slashable until they're completable

## Properties

* A staker has 3.5 days to undelegate and queue withdrawals in response to magnitude updates they are not comfortable with
* AVS can always generate a forecast of a useful minimum the slashable stake they will have over the next 17.5 days
