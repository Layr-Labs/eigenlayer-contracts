#!/bin/bash

# Unapply the patch to the contracts directory
echo "Unapplying the patch for SplitContractMixin"
git apply -R certora/patches/SplitContractMixin.patch

echo "Unapplying the patch for AllocationManager"
git apply -R certora/patches/AllocationManager.patch