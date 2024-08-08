# Timing

The timing and delays imposed in the slashing protocol are intricate and very important to enforcing the correct checks and balances. This document explains them.

## Slashing Requests

* Slashing requests are executed synchronously

## Magnitude Updates

* Updates to magnitudes take effect after 21 days
* There can be at most 3 pending allocations/deallocations for a certain strategy, operator, and operatorSet at a time

## Withdrawals

* Withdrawals are completable 17.5 days after they're queued
* Withdrawals are slashable until they're completable

## Properties

* A staker has 3.5 days to undelegate and queue withdrawals in response to magnitude updates they are not comfortable with
* AVS can always generate a forecast of a useful minimum the slashable stake they will have over the next 17.5 days
