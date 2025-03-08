#!/bin/bash

# Run all security tests
cd /home/ubuntu/repos/eigenlayer-contracts
forge test --match-path src/test/security/*.t.sol -vv
